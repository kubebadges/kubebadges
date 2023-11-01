package server

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/kubebadges/kubebadges/internal/server/svc"
)

type Server struct {
	internalEngine *gin.Engine
	externalEngine *gin.Engine
	svcCtx         *svc.ServerContext
}

func NewServer() *Server {
	return &Server{
		internalEngine: gin.Default(),
		externalEngine: gin.Default(),
	}
}

func (s *Server) init() {
	gin.SetMode(gin.ReleaseMode)
	s.svcCtx = svc.NewServerContext()
	s.initRouter()
}

func (s *Server) Start() error {
	s.init()
	go func() {
		slog.Info("run external api", "port", 8080)
		if err := s.externalEngine.Run(":8080"); err != nil {
			panic(err)
		}
	}()
	slog.Info("run internal api", "port", 8090)
	return s.internalEngine.Run(":8090")
}
