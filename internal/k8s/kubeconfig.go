package k8s

import (
	"context"

	"github.com/kubebadges/kubebadges/internal/config"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (k *KubeHelper) GetOrCreateConfig() (*v1.ConfigMap, error) {
	configMap, err := k.client.CoreV1().ConfigMaps(config.KubeBadgeNamespace).Get(context.Background(), config.KubeBadgeConfigName, metav1.GetOptions{})
	if err == nil {
		return k.initConfigMapData(configMap), nil
	}

	configMap = k.createConfigMap()

	configMap, err = k.client.CoreV1().ConfigMaps(config.KubeBadgeNamespace).Create(context.Background(), configMap, metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}

	return k.initConfigMapData(configMap), nil
}

func (k *KubeHelper) GetConfig() (*v1.ConfigMap, error) {
	return k.client.CoreV1().ConfigMaps(config.KubeBadgeNamespace).Get(context.Background(), config.KubeBadgeConfigName, metav1.GetOptions{})
}

func (k *KubeHelper) UpdateConfig(configMap *v1.ConfigMap) (*v1.ConfigMap, error) {
	return k.client.CoreV1().ConfigMaps(config.KubeBadgeNamespace).Update(context.Background(), configMap, metav1.UpdateOptions{})
}

func (k *KubeHelper) DeleteConfig(configMap *v1.ConfigMap) error {
	return k.client.CoreV1().ConfigMaps(config.KubeBadgeNamespace).Delete(context.Background(), configMap.Name, metav1.DeleteOptions{})
}

func (k *KubeHelper) createConfigMap() *v1.ConfigMap {
	return &v1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name: config.KubeBadgeConfigName,
		},
		Data: map[string]string{},
	}
}

func (k *KubeHelper) initConfigMapData(configMap *v1.ConfigMap) *v1.ConfigMap {
	if configMap.Data == nil {
		configMap.Data = map[string]string{}
	}

	return configMap
}
