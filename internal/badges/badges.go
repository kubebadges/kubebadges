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

func (b *BadgesHelper) CreateBadgeProxy(badge *BadgeBuilder, c *gin.Context) {
	label := strings.ReplaceAll(badge.Label, "-", "--")
	label = strings.ReplaceAll(label, "_", "__")
	label = strings.ReplaceAll(label, " ", "_")

	message := strings.ReplaceAll(badge.Message, "-", "--")
	message = strings.ReplaceAll(message, "_", "__")
	message = strings.ReplaceAll(message, " ", "_")

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

	c.Writer.Header().Set("X-App-Name", "KubeBadge")
	c.Writer.Header().Del("Access-Control-Allow-Origin")
	proxy.ServeHTTP(c.Writer, c.Request)
}
