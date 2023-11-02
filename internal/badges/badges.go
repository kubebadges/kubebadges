package badges

import (
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kubebadges/kubebadges/internal/config"
)

const (
	AppNameHeader = "X-App-Name"
	AppName       = "KubeBadge"
	AccessControl = "Access-Control-Allow-Origin"
)

type BadgesHelper struct {
	targetHOST   string
	targetScheme string
	cacheTime    int
}

func NewBadgesHelper(config *config.Config) *BadgesHelper {
	return &BadgesHelper{
		targetHOST:   config.ShieldsHost,
		targetScheme: config.ShieldsScheme,
		cacheTime:    config.BadgeCacheTime,
	}
}

func formatString(s string) string {
	s = strings.ReplaceAll(s, "-", "--")
	s = strings.ReplaceAll(s, "_", "__")
	return strings.ReplaceAll(s, " ", "_")
}

func (b *BadgesHelper) CreateBadgeProxy(badge *BadgeBuilder, c *gin.Context) {
	label := formatString(badge.Label)
	message := formatString(badge.Message)

	badgeURL := &url.URL{
		Scheme: b.targetScheme,
		Host:   b.targetHOST,
		Path:   fmt.Sprintf("/badge/%s-%s-%s", label, message, badge.MessageColor),
	}

	q := badgeURL.Query()
	if len(badge.Style) > 0 {
		q.Set("style", badge.Style)
	}

	badgeURL.RawQuery = q.Encode()

	slog.Info("proxy to shields", "url", badgeURL.String())

	proxy := httputil.ReverseProxy{
		Director: func(req *http.Request) {
			req.URL = badgeURL
		},
		ModifyResponse: func(resp *http.Response) error {
			resp.Header.Set("Cache-Control", fmt.Sprintf("max-age=%d, s-maxage=%d", b.cacheTime, b.cacheTime))
			return nil
		},
	}

	c.Writer.Header().Set(AppNameHeader, AppName)
	c.Writer.Header().Del(AccessControl)
	proxy.ServeHTTP(c.Writer, c.Request)
}
