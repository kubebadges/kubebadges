package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kubebadges/kubebadges/internal/cache"
	"github.com/kubebadges/kubebadges/internal/model"
	"github.com/kubebadges/kubebadges/internal/server/svc"
)

type KubeController struct {
	BaseController
	*svc.ServerContext
	cache *cache.Cache[string, []model.KubeBadges]
}

func NewKubeController(svc *svc.ServerContext) *KubeController {
	cache := cache.NewCache[string, []model.KubeBadges]()
	return &KubeController{
		ServerContext: svc,
		cache:         cache,
	}
}

func (s *KubeController) ListNodes(c *gin.Context) {
	result, ok := s.cache.Get("nodes")

	if !ok || c.Query("force") == "true" {
		nodes, err := s.KubeHelper.GetNodes()
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
		result = make([]model.KubeBadges, len(nodes))

		for i, node := range nodes {
			result[i] = model.KubeBadges{
				Kind:  "node",
				Name:  node.Name,
				Key:   fmt.Sprintf("/kube/node/%s", node.Name),
				Badge: fmt.Sprintf("/badges/kube/node/%s", node.Name),
			}
		}
		s.cache.Set("nodes", result, time.Minute*60)
	}

	s.populateKubeBadges(c, result)
}

func (s *KubeController) ListNamespaces(c *gin.Context) {
	result, ok := s.cache.Get("namespaces")

	if !ok || c.Query("force") == "true" {
		namespaces, err := s.KubeHelper.GetNamespaces()
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
		result = make([]model.KubeBadges, len(namespaces))

		for i, namespace := range namespaces {
			result[i] = model.KubeBadges{
				Kind:  "namespace",
				Name:  namespace.Name,
				Key:   fmt.Sprintf("/kube/namespace/%s", namespace.Name),
				Badge: fmt.Sprintf("/badges/kube/namespace/%s", namespace.Name),
			}
		}
		s.cache.Set("namespace", result, time.Minute*5)
	}

	s.populateKubeBadges(c, result)
}

func (s *KubeController) ListDeployments(c *gin.Context) {
	namespace := c.Param("namespace")

	result, ok := s.cache.Get(fmt.Sprintf("deployment_%s", namespace))
	if !ok || c.Query("force") == "true" {
		deployments, err := s.KubeHelper.GetDeployments(namespace)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		result = make([]model.KubeBadges, len(deployments))

		for i, deployment := range deployments {
			result[i] = model.KubeBadges{
				Kind:  "deployment",
				Name:  deployment.Name,
				Key:   fmt.Sprintf("/kube/deployment/%s/%s", namespace, deployment.Name),
				Badge: fmt.Sprintf("/badges/kube/deployment/%s/%s", namespace, deployment.Name),
			}
		}
		s.cache.Set(fmt.Sprintf("deployment_%s", namespace), result, time.Minute*2)
	}

	s.populateKubeBadges(c, result)
}

type UpdateBadgeRequest struct {
	DisplayName *string `json:"display_name"`
	Alias       *string `json:"alias"`
	Allowed     *bool   `json:"allowed"`
	Key         string  `json:"key"`
}

func (s *KubeController) UpdateBadge(c *gin.Context) {
	var req UpdateBadgeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.Allowed == nil && req.DisplayName == nil && req.Alias == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "at least one of allowed, display_name or alias should be provided"})
		return
	}

	// get kubebadge crd
	kubeBadge, err := s.KubeBadgesService.GetKubeBadge(req.Key, true)
	if err != nil {
		// create kubebadge crd
		spec := s.KubeBadgesService.CreateKubeBadgesSpec()
		spec.DisplayName = ""
		spec.AliasURL = ""
		spec.Allowed = false
		spec.OriginalURL = req.Key
		spec.Type, _, _ = s.parseKey(req.Key)
		_, spec.OwnerNamespace, _ = s.parseKey(req.Key)

		kubeBadge, err = s.KubeBadgesService.CreateKubeBadge(spec)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	// update kubebadge crd
	if req.Allowed != nil {
		kubeBadge.Spec.Allowed = *req.Allowed
	}
	if req.DisplayName != nil {
		kubeBadge.Spec.DisplayName = *req.DisplayName
	}
	if req.Alias != nil {
		kubeBadge.Spec.AliasURL = *req.Alias
	}

	_, err = s.KubeBadgesService.UpdateKubeBadge(kubeBadge)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (s *KubeController) parseKey(key string) (resourceType string, namespace string, name string) {
	segments := strings.Split(key, "/")
	switch segments[2] {
	case "node":
		resourceType = "node"
		name = segments[3]
	case "namespace":
		resourceType = "namespace"
		name = segments[3]
	case "deployment":
		resourceType = "deployment"
		namespace = segments[3]
		name = segments[4]
	case "pod":
		resourceType = "pod"
		namespace = segments[3]
		name = segments[4]
	}

	return
}

func (s *KubeController) populateKubeBadges(c *gin.Context, result []model.KubeBadges) {
	var wg sync.WaitGroup
	newResult := make([]model.KubeBadges, len(result))

	for i := range result {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			newBadge := result[index]
			kubeBadge, err := s.KubeBadgesService.GetKubeBadge(result[index].Key, false)
			if err == nil {
				newBadge.Allowed = kubeBadge.Spec.Allowed
				newBadge.DisplayName = kubeBadge.Spec.DisplayName
				newBadge.AliasURL = kubeBadge.Spec.AliasURL
			}
			newResult[index] = newBadge
		}(i)
	}
	wg.Wait()

	c.JSON(http.StatusOK, newResult)
}

func (s *KubeController) GetConfig(c *gin.Context) {
	configMap, err := s.KubeHelper.GetOrCreateConfig()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, s.mapToConfig(configMap.Data))
}

func (s *KubeController) UpdateConfig(c *gin.Context) {
	var kubeBadgeConfig model.KubeBadgesConfig
	if err := c.ShouldBindJSON(&kubeBadgeConfig); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	configMap, err := s.KubeHelper.GetOrCreateConfig()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	configMap.Data = s.configToMap(&kubeBadgeConfig)

	configMap, err = s.KubeHelper.UpdateConfig(configMap)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, s.mapToConfig(configMap.Data))
}

func (s *KubeController) configToMap(config *model.KubeBadgesConfig) map[string]string {
	jsonData, _ := json.Marshal(config)
	var configMap map[string]string
	_ = json.Unmarshal(jsonData, &configMap)
	return configMap
}

func (s *KubeController) mapToConfig(configMap map[string]string) *model.KubeBadgesConfig {
	jsonData, _ := json.Marshal(configMap)
	var config model.KubeBadgesConfig
	_ = json.Unmarshal(jsonData, &config)
	return &config
}
