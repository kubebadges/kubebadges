package server

import (
	"embed"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kubebadges/kubebadges"
	"github.com/kubebadges/kubebadges/internal/server/controller"
	"github.com/kubebadges/kubebadges/internal/server/middleware"
)

var contentTypeMap = map[string]string{
	".js":   "application/javascript",
	".html": "text/html",
	".css":  "text/css",
	".json": "application/json",
	".png":  "image/png",
	".otf":  "font/otf",
	".ttf":  "font/ttf",
}

func getContentType(path string) (string, bool) {
	ext := filepath.Ext(path)
	contentType, ok := contentTypeMap[ext]
	return contentType, ok
}

func registerStaticFiles(router *gin.Engine, fs embed.FS, root string) {
	entries, _ := fs.ReadDir(root)
	for _, entry := range entries {
		if entry.IsDir() {
			registerStaticFiles(router, fs, root+"/"+entry.Name())
		} else {
			filePath := root + "/" + entry.Name()
			mineType, mineTypeOK := getContentType(filePath)
			if filePath == "web/index.html" {
				router.GET("/", func(c *gin.Context) {
					data, _ := kubebadges.WebFiles.ReadFile(filePath)
					if mineTypeOK {
						c.Data(http.StatusOK, mineType, data)
					} else {
						c.Data(http.StatusOK, http.DetectContentType(data), data)
					}
				})
			}
			router.GET("/"+strings.TrimPrefix(filePath, "web/"), func(c *gin.Context) {
				c.Header("Cache-Control", "public, max-age=300")
				data, _ := fs.ReadFile(filePath)
				if mineTypeOK {
					c.Data(http.StatusOK, mineType, data)
				} else {
					c.Data(http.StatusOK, http.DetectContentType(data), data)
				}
			})
		}
	}
}

func (s *Server) initRouter() {
	baseController := controller.BaseController{
		ServerContext: s.svcCtx,
	}
	kubeController := controller.NewKubeController(s.svcCtx)
	badgesController := controller.NewBadgesController(s.svcCtx)

	registerStaticFiles(s.internalEngine, kubebadges.WebFiles, "web")

	s.internalEngine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
		AllowHeaders:     []string{"Origin", "Push-Id", "App", "App-Version", "X-Device-Id", "Content-Type", "Content-Length", "Authorization", "X-App-Name"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return strings.Contains(origin, "localhost") || strings.Contains(origin, "127.0.0.1")
		},
	}))

	// admin routes
	api := s.internalEngine.Group("/api")
	{
		api.GET("/nodes", kubeController.ListNodes)
		api.GET("/namespaces", kubeController.ListNamespaces)
		api.GET("/deployments/:namespace", kubeController.ListDeployments)
		api.POST("/badge", kubeController.UpdateBadge)
	}

	badges := s.internalEngine.Group("/badges")
	{
		// badges routes
		badges.GET("/kube/node/:node", badgesController.Node)
		badges.GET("/kube/namespace/:namespace", badgesController.Namespace)
		badges.GET("/kube/deployment/:namespace/:deployment", badgesController.Deployment)
		badges.GET("/kube/pod/:namespace/:pod", badgesController.Pod)
	}

	// for external api
	s.externalEngine.NoRoute(func(ctx *gin.Context) {
		baseController.NotFound(ctx)
	})
	s.externalEngine.Use(middleware.BadgeApiAccessMiddleware(s.svcCtx.KubeBadgesService))
	exBadges := s.externalEngine.Group("/badges")
	{
		// badges routes
		exBadges.GET("/kube/node/:node", badgesController.Node)
		exBadges.GET("/kube/namespace/:namespace", badgesController.Namespace)
		exBadges.GET("/kube/deployment/:namespace/:deployment", badgesController.Deployment)
		exBadges.GET("/kube/pod/:namespace/:pod", badgesController.Pod)
	}

}
