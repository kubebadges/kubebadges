package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kubebadges/kubebadges/internal/service"
)

var unauthorizedSvg = `<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="114" height="20" role="img" aria-label="404: Unauthorized">
    <title>404: Unauthorized</title>
    <linearGradient id="s" x2="0" y2="100%">
        <stop offset="0" stop-color="#bbb" stop-opacity=".1"/>
        <stop offset="1" stop-opacity=".1"/>
    </linearGradient>
    <clipPath id="r">
        <rect width="114" height="20" rx="3" fill="#fff"/>
    </clipPath>
    <g clip-path="url(#r)">
        <rect width="31" height="20" fill="#555"/>
        <rect x="31" width="83" height="20" fill="#e05d44"/>
        <rect width="114" height="20" fill="url(#s)"/>
    </g>
    <g fill="#fff" text-anchor="middle" font-family="Verdana,Geneva,DejaVu Sans,sans-serif" text-rendering="geometricPrecision" font-size="110">
        <text aria-hidden="true" x="165" y="150" fill="#010101" fill-opacity=".3" transform="scale(.1)" textLength="210">404</text>
        <text x="165" y="140" transform="scale(.1)" fill="#fff" textLength="210">404</text>
        <text aria-hidden="true" x="715" y="150" fill="#010101" fill-opacity=".3" transform="scale(.1)" textLength="730">Unauthorized</text>
        <text x="715" y="140" transform="scale(.1)" fill="#fff" textLength="730">Unauthorized</text>
    </g>
</svg>
`

func BadgeApiAccessMiddleware(kubeService *service.KubeBadgesService) gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("X-App-Name", "KubeBadge")

		if !strings.HasPrefix(c.Request.URL.Path, "/badges") {
			c.Header("Content-Type", "image/svg+xml")
			c.String(http.StatusUnauthorized, unauthorizedSvg)
			c.Abort()
			return
		}

		key := strings.TrimPrefix(c.Request.URL.Path, "/badges")
		kubeBadge, err := kubeService.GetKubeBadge(key, false)
		if err != nil || !kubeBadge.Spec.Allowed {
			c.Header("Content-Type", "image/svg+xml")
			c.String(http.StatusUnauthorized, unauthorizedSvg)
			c.Abort()
			return
		}

		c.Next()
	}
}
