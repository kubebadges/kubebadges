package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kubebadges/kubebadges/internal/badges"
	"github.com/kubebadges/kubebadges/internal/server/svc"
	"github.com/kubebadges/kubebadges/internal/service"
)

var notFoundSvg = `
<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="132" height="20" role="img" aria-label="404: badge not found">
    <title>404: badge not found</title>
    <linearGradient id="s" x2="0" y2="100%">
        <stop offset="0" stop-color="#bbb" stop-opacity=".1"/>
        <stop offset="1" stop-opacity=".1"/>
    </linearGradient>
    <clipPath id="r">
        <rect width="132" height="20" rx="3" fill="#fff"/>
    </clipPath>
    <g clip-path="url(#r)">
        <rect width="31" height="20" fill="#555"/>
        <rect x="31" width="101" height="20" fill="#e05d44"/>
        <rect width="132" height="20" fill="url(#s)"/>
    </g>
    <g fill="#fff" text-anchor="middle" font-family="Verdana,Geneva,DejaVu Sans,sans-serif" text-rendering="geometricPrecision" font-size="110">
        <text aria-hidden="true" x="165" y="150" fill="#010101" fill-opacity=".3" transform="scale(.1)" textLength="210">404</text>
        <text x="165" y="140" transform="scale(.1)" fill="#fff" textLength="210">404</text>
        <text aria-hidden="true" x="805" y="150" fill="#010101" fill-opacity=".3" transform="scale(.1)" textLength="910">badge not found</text>
        <text x="805" y="140" transform="scale(.1)" fill="#fff" textLength="910">badge not found</text>
    </g>
</svg>
`

type BadgeMessage struct {
	Key          string
	Label        string
	Message      string
	MessageColor string
}

type BaseController struct {
	*svc.ServerContext
}

func (b *BaseController) NotFound(c *gin.Context) {
	resultType := c.DefaultQuery("type", "svg")
	if resultType == "json" {
		c.Header("Content-Type", "application/json")
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.Header("Content-Type", "image/svg+xml")
	c.String(http.StatusNotFound, notFoundSvg)
}

func (b *BaseController) Success(c *gin.Context, kubeBadgesService *service.KubeBadgesService, badgeMessage BadgeMessage) {

	if kubeBadge, err := kubeBadgesService.GetKubeBadge(badgeMessage.Key, false); err == nil {
		if len(kubeBadge.Spec.DisplayName) > 0 {
			badgeMessage.Label = kubeBadge.Spec.DisplayName
		}
	}

	badge := badges.NewBadgeBuilder().
		SetLabel(badgeMessage.Label).
		SetMessage(badgeMessage.Message).
		SetMessageColor(badgeMessage.MessageColor).
		SetStyle(c.Query("style")).
		Build()

	resultType := c.DefaultQuery("type", "svg")
	if resultType == "json" {
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, gin.H{
			"label":   badge.Label,
			"message": badge.Message,
			"color":   badge.MessageColor,
		})
		return
	}

	b.BadgesHelper.CreateBadgeProxy(badge, c)
}
