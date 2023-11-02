package k8s

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/kubebadges/kubebadges/internal/config"
	v1 "github.com/kubebadges/kubebadges/pkg/apis/kubebadges/v1"
	typev1 "github.com/kubebadges/kubebadges/pkg/generated/clientset/versioned/typed/kubebadges/v1"
	informers "github.com/kubebadges/kubebadges/pkg/generated/informers/externalversions/kubebadges/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/cache"
)

func (k *KubeHelper) kubebadge() typev1.KubeBadgeInterface {
	return k.kubeBadgeClient.KubebadgesV1().KubeBadges(config.KubeBadgeNamespace)
}

func (k *KubeHelper) GenerateKubeName(name string) string {
	name = strings.TrimPrefix(name, "/")
	return strings.ReplaceAll(name, "/", "-")
}

func (k *KubeHelper) NewKubeBadgeSpec() v1.KubeBadgeSpec {
	return v1.KubeBadgeSpec{}
}

func (k *KubeHelper) GetBadge(name string) (*v1.KubeBadge, error) {
	return k.kubebadge().Get(context.Background(), name, metav1.GetOptions{})
}

func (k *KubeHelper) GetBadgeByLabel(key string, value string) ([]v1.KubeBadge, error) {
	filter := metav1.ListOptions{
		LabelSelector: key + "=" + value,
	}
	list, err := k.kubebadge().List(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	return list.Items, nil
}

func (k *KubeHelper) CreateKubeBadge(spec v1.KubeBadgeSpec) (*v1.KubeBadge, error) {
	kubeBadgeCR := v1.KubeBadge{
		TypeMeta: metav1.TypeMeta{
			Kind:       config.KubeBadgeCRDKind,
			APIVersion: config.KubeBadgeCRDAPIVersion,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      k.GenerateKubeName(spec.OriginalURL),
			Namespace: config.KubeBadgeNamespace,
			Labels: map[string]string{
				v1.KubeBadgeLabelType:           spec.Type,
				v1.KubeBadgeLabelAllowed:        strconv.FormatBool(spec.Allowed),
				v1.KubeBadgeLabelAliasURL:       k.GenerateKubeName(spec.AliasURL),
				v1.KubeBadgeLabelOriginalURL:    k.GenerateKubeName(spec.OriginalURL),
				v1.KubeBadgeLabelOwnerNamespace: spec.OwnerNamespace,
			},
		},
		Spec: spec,
	}

	return k.kubebadge().Create(context.Background(), &kubeBadgeCR, metav1.CreateOptions{})
}

func (k *KubeHelper) UpdateKubeBadge(kubeBadge *v1.KubeBadge) (*v1.KubeBadge, error) {
	kubeBadge.ObjectMeta.Labels = map[string]string{
		v1.KubeBadgeLabelType:           kubeBadge.Spec.Type,
		v1.KubeBadgeLabelAllowed:        strconv.FormatBool(kubeBadge.Spec.Allowed),
		v1.KubeBadgeLabelAliasURL:       k.GenerateKubeName(kubeBadge.Spec.AliasURL),
		v1.KubeBadgeLabelOriginalURL:    k.GenerateKubeName(kubeBadge.Spec.OriginalURL),
		v1.KubeBadgeLabelOwnerNamespace: kubeBadge.Spec.OwnerNamespace,
	}

	return k.kubebadge().Update(context.Background(), kubeBadge, metav1.UpdateOptions{})
}

func (k *KubeHelper) DeleteKubeBadge(name string) error {
	return k.kubebadge().Delete(context.Background(), name, metav1.DeleteOptions{})
}

func (k *KubeHelper) NewKubeBadgeInformer() cache.SharedIndexInformer {
	return informers.NewKubeBadgeInformer(
		k.kubeBadgeClient,
		config.KubeBadgeNamespace,
		24*time.Hour,
		cache.Indexers{},
	)
}
