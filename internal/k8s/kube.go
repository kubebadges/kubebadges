package k8s

import (
	"context"
	"log/slog"
	"os"

	"github.com/kubebadges/kubebadges/pkg/generated/clientset/versioned"
	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

type KubeHelper struct {
	client          *kubernetes.Clientset
	kubeBadgeClient *versioned.Clientset
}

func NewKubeHelper() *KubeHelper {
	return &KubeHelper{}
}

func (k *KubeHelper) Init() {
	config, err := rest.InClusterConfig()
	if err != nil {
		// fallback to kubeconfig
		if kubeconfig := homedir.HomeDir() + "/.kube/config"; os.Getenv("KUBECONFIG") != "" || kubeconfig != "" {
			config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
			if err != nil {
				panic(err.Error())
			}
		} else {
			panic(err.Error())
		}
	}

	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	k.client = clientset

	kubeBadgeClient, err := versioned.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	k.kubeBadgeClient = kubeBadgeClient

	version, err := clientset.Discovery().ServerVersion()
	if err != nil {
		panic(err.Error())
	}

	slog.Info("Connected to kubernetes", "version", version.String())
}

func (k *KubeHelper) GetClient() *kubernetes.Clientset {
	return k.client
}

func (k *KubeHelper) GetNodes() ([]corev1.Node, error) {
	nodes, err := k.client.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return nodes.Items, nil
}

func (k *KubeHelper) GetNode(name string) (*corev1.Node, error) {
	node, err := k.client.CoreV1().Nodes().Get(context.Background(), name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (k *KubeHelper) GetNamespaces() ([]corev1.Namespace, error) {
	namespaces, err := k.client.CoreV1().Namespaces().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return namespaces.Items, nil
}

func (k *KubeHelper) GetNamespace(name string) (*corev1.Namespace, error) {
	namespace, err := k.client.CoreV1().Namespaces().Get(context.Background(), name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return namespace, nil
}

func (k *KubeHelper) GetDeployments(namespace string) ([]v1.Deployment, error) {
	deployments, err := k.client.AppsV1().Deployments(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return deployments.Items, nil
}

func (k *KubeHelper) GetDeployment(namespace string, name string) (*v1.Deployment, error) {
	deployment, err := k.client.AppsV1().Deployments(namespace).Get(context.Background(), name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return deployment, nil
}

func (k *KubeHelper) GetPods(namespace string) ([]corev1.Pod, error) {
	pods, err := k.client.CoreV1().Pods("").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return pods.Items, nil
}

func (k *KubeHelper) GetPod(namespace string, name string) (*corev1.Pod, error) {
	pod, err := k.client.CoreV1().Pods(namespace).Get(context.Background(), name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return pod, nil
}
