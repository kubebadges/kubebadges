package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/kubebadges/kubebadges"
)

func TestRegisterStaticFiles(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	registerStaticFiles(router, kubebadges.WebFiles, "web")

	tests := []struct {
		name         string
		path         string
		wantStatus   int
		wantResponse string
		wantHeader   string
	}{
		{
			name:         "index.html",
			path:         "/",
			wantStatus:   http.StatusOK,
			wantResponse: "index file content",
			wantHeader:   "Content-Type: text/html",
		},
		{
			name:         "canvaskit.wasm",
			path:         "/canvaskit/canvaskit.wasm",
			wantStatus:   http.StatusOK,
			wantResponse: "style file content",
			wantHeader:   "Content-Type: application/wasm",
		},
		{
			name:         "script.js",
			path:         "/canvaskit/skwasm.js",
			wantStatus:   http.StatusOK,
			wantResponse: "script file content",
			wantHeader:   "Content-Type: application/javascript",
		},
		{
			name:       "not found",
			path:       "/non-existent",
			wantStatus: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", tt.path, nil)
			router.ServeHTTP(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("expected status %d, got %d", tt.wantStatus, w.Code)
			}
			if tt.wantHeader != "" {
				if got := w.Header().Get("Content-Type"); got != tt.wantHeader[14:] {
					t.Errorf("expected header %s, got %s", tt.wantHeader, got)
				}
			}
		})
	}
}
