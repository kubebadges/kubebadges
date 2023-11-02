package middleware

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	v1 "github.com/kubebadges/kubebadges/pkg/apis/kubebadges/v1"
)

type MockKubeBadgesService struct{}

func (s *MockKubeBadgesService) GetKubeBadge(key string, b bool) (*v1.KubeBadge, error) {
	if key == "/authorized" {
		return &v1.KubeBadge{Spec: v1.KubeBadgeSpec{Allowed: true}}, nil
	}
	// 在这里返回你期望的值
	return &v1.KubeBadge{Spec: v1.KubeBadgeSpec{Allowed: false}}, nil
}

func TestBadgeApiAccessMiddleware(t *testing.T) {
	kubeBadgeService := &MockKubeBadgesService{}
	gin.SetMode(gin.TestMode)

	t.Run("unauthorized path", func(t *testing.T) {
		router := gin.New()
		router.Use(BadgeApiAccessMiddleware(kubeBadgeService))
		router.GET("/badges", func(c *gin.Context) {
			c.String(200, "ok")
		})

		w := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/some", nil)
		router.ServeHTTP(w, req)

		if w.Code != http.StatusUnauthorized {
			t.Fatalf("expected status code %d, but got %d", http.StatusUnauthorized, w.Code)
		}

		if !strings.Contains(w.Body.String(), unauthorizedSvg) {
			t.Fatalf("expected response body to contain %s", unauthorizedSvg)
		}
	})

	t.Run("unauthorized badge", func(t *testing.T) {
		router := gin.New()
		router.Use(BadgeApiAccessMiddleware(kubeBadgeService))
		router.GET("/badges/some", func(c *gin.Context) {})

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/badges/unauthorized", nil)
		router.ServeHTTP(w, req)

		if w.Code != http.StatusUnauthorized {
			t.Fatalf("expected status code %d, but got %d", http.StatusUnauthorized, w.Code)
		}

		if !strings.Contains(w.Body.String(), unauthorizedSvg) {
			t.Fatalf("expected response body to contain %s", unauthorizedSvg)
		}
	})

	t.Run("authorized badge", func(t *testing.T) {
		router := gin.New()
		router.Use(BadgeApiAccessMiddleware(kubeBadgeService))
		router.GET("/badges/authorized", func(c *gin.Context) {})

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/badges/authorized", nil)
		router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Fatalf("expected status code %d, but got %d", http.StatusOK, w.Code)
		}
	})
}
