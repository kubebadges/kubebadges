package svc

import (
	"github.com/kubebadges/kubebadges/internal/badges"
	"github.com/kubebadges/kubebadges/internal/config"
	"github.com/kubebadges/kubebadges/internal/k8s"
	"github.com/kubebadges/kubebadges/internal/service"
	"github.com/kubebadges/kubebadges/internal/utils"
)

type ServerContext struct {
	KubeHelper        *k8s.KubeHelper
	BadgesHelper      *badges.BadgesHelper
	Config            *config.Config
	KubeBadgesService *service.KubeBadgesService
}

func NewServerContext() *ServerContext {
	// load data
	config := &config.Config{}
	config.ShieldsScheme = utils.GetEnv("SHIELDS_SCHEME", "http")
	config.ShieldsHost = utils.GetEnv("SHIELDS_HOST", "127.0.0.1:8081")
	config.CacheTime = utils.GetEnvAsInt("CACHE_TIME", 300)
	config.BadgeCacheTime = utils.GetEnvAsInt("BADGE_CACHE_TIME", 300)

	kubeHelper := k8s.NewKubeHelper()
	kubeHelper.Init()

	kubeBadgeService := service.NewKubeBadgesService(kubeHelper)
	go kubeBadgeService.Run()

	return &ServerContext{
		Config:            config,
		KubeHelper:        kubeHelper,
		BadgesHelper:      badges.NewBadgesHelper(config),
		KubeBadgesService: kubeBadgeService,
	}
}
