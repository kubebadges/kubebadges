package service

import (
	"errors"
	"log/slog"
	"time"

	"github.com/kubebadges/kubebadges/internal/k8s"
	v1 "github.com/kubebadges/kubebadges/pkg/apis/kubebadges/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"

	mcache "github.com/kubebadges/kubebadges/internal/cache"
)

const (
	eventActionAdd    = "add"
	eventActionUpdate = "update"
	eventActionDelete = "delete"
)

type queueEvent struct {
	action    string
	kubebadge *v1.KubeBadge
}

type KubeBadgesService struct {
	kubeHelper        *k8s.KubeHelper
	informer          cache.SharedIndexInformer
	queue             workqueue.RateLimitingInterface
	cacheWithKey      *mcache.Cache[string, *v1.KubeBadge] // key is the kubebadge name
	cacheWithAliasURL *mcache.Cache[string, *v1.KubeBadge] // key is the kubebadge's alias url
}

func NewKubeBadgesService(kubeHelper *k8s.KubeHelper) *KubeBadgesService {
	service := &KubeBadgesService{
		kubeHelper:        kubeHelper,
		informer:          kubeHelper.NewKubeBadgeInformer(),
		queue:             workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter()),
		cacheWithKey:      mcache.NewCache[string, *v1.KubeBadge](),
		cacheWithAliasURL: mcache.NewCache[string, *v1.KubeBadge](),
	}
	service.init()

	return service
}

func (k *KubeBadgesService) init() {
	k.informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			if value, ok := obj.(*v1.KubeBadge); ok {

				k.queue.Add(queueEvent{
					action:    eventActionAdd,
					kubebadge: value,
				})
			}
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			if value, ok := newObj.(*v1.KubeBadge); ok {
				k.queue.Add(queueEvent{
					action:    eventActionUpdate,
					kubebadge: value,
				})
			}
		},
		DeleteFunc: func(obj interface{}) {
			if value, ok := obj.(*v1.KubeBadge); ok {
				k.queue.Add(queueEvent{
					action:    eventActionDelete,
					kubebadge: value,
				})
			}
		},
	})
}

func (k *KubeBadgesService) Run() {
	stopCh := make(chan struct{})
	defer close(stopCh)
	defer k.queue.ShutDown()

	go k.informer.Run(stopCh)

	if !cache.WaitForCacheSync(stopCh, k.informer.HasSynced) {
		return
	}

	wait.Until(k.runWorker, time.Second, stopCh)
}

func (k *KubeBadgesService) runWorker() {
	for k.processNextItem() {
	}
}

func (k *KubeBadgesService) processNextItem() bool {
	event, quit := k.queue.Get()
	if quit {
		return false
	}

	defer k.queue.Done(event)
	if qe, ok := event.(queueEvent); ok {

		slog.Info("kubebadges service event handler", slog.String("action", qe.action), slog.String("name", qe.kubebadge.Name))

		if qe.action == eventActionAdd || qe.action == eventActionUpdate {
			k.addOrUpdateKubeBadge(qe.kubebadge)
		} else if qe.action == eventActionDelete {
			k.deleteKubeBadge(qe.kubebadge)
		}
	}

	return true
}

func (k *KubeBadgesService) addOrUpdateKubeBadge(kubebadge *v1.KubeBadge) {
	k.cacheWithKey.Set(kubebadge.ObjectMeta.Name, kubebadge, 48*time.Hour)
	if len(kubebadge.Spec.AliasURL) > 0 {
		k.cacheWithAliasURL.Set(kubebadge.Spec.AliasURL, kubebadge, 48*time.Hour)
	}
}

func (k *KubeBadgesService) deleteKubeBadge(kubebadge *v1.KubeBadge) {
	k.cacheWithKey.Delete(kubebadge.ObjectMeta.Name)
	if len(kubebadge.Spec.AliasURL) > 0 {
		k.cacheWithAliasURL.Delete(kubebadge.Spec.AliasURL)
	}
}

func (k *KubeBadgesService) GenerateKubeBadgeName(name string) string {
	return k.kubeHelper.GenerateKubeName(name)
}

func (k *KubeBadgesService) CreateKubeBadgesSpec() v1.KubeBadgeSpec {
	return v1.KubeBadgeSpec{}
}

func (k *KubeBadgesService) CreateKubeBadge(spec v1.KubeBadgeSpec) (*v1.KubeBadge, error) {
	return k.kubeHelper.CreateKubeBadge(spec)
}

func (k *KubeBadgesService) UpdateKubeBadge(kubeBadge *v1.KubeBadge) (*v1.KubeBadge, error) {
	return k.kubeHelper.UpdateKubeBadge(kubeBadge)
}

func (k *KubeBadgesService) GetKubeBadge(name string, force bool) (*v1.KubeBadge, error) {
	if result, ok := k.cacheWithKey.Get(k.GenerateKubeBadgeName(name)); ok {
		return result, nil
	}
	if !force {
		return nil, errors.New("not found")
	}

	result, err := k.kubeHelper.GetBadge(k.GenerateKubeBadgeName(name))
	if err != nil {
		return nil, err
	}

	k.addOrUpdateKubeBadge(result)

	return result, nil

}

func (k *KubeBadgesService) GetKubeBadgeByAlias(aliasURL string) (*v1.KubeBadge, error) {
	if result, ok := k.cacheWithAliasURL.Get(aliasURL); ok {
		return result, nil
	}

	result, err := k.kubeHelper.GetBadgeByLabel(v1.KubeBadgeLabelAliasURL, k.GenerateKubeBadgeName(aliasURL))
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, errors.New("not found")
	}

	k.addOrUpdateKubeBadge(&result[0])

	return &result[0], nil

}
