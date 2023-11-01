package k8s

import (
	"testing"

	v1 "github.com/kubebadges/kubebadges/pkg/apis/kubebadges/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
)

func TestKubeHelper_Init(t *testing.T) {
	kubeHelper := NewKubeHelper()
	kubeHelper.Init()
}
func TestKubeHelper_GetNamespaces(t *testing.T) {
	kubeHelper := NewKubeHelper()
	kubeHelper.Init()

	namespaces, err := kubeHelper.GetNamespaces()
	if err != nil {
		t.Errorf("Error getting namespaces: %v", err)
	}

	for _, namespace := range namespaces {
		t.Logf("%s-%s\n", namespace.Name, namespace.Status.Phase)
	}

	if len(namespaces) == 0 {
		t.Errorf("Expected at least one namespace, but got none")
	}
}

func TestKubeHelper_GetKubeBadge(t *testing.T) {
	kubeHelper := NewKubeHelper()
	kubeHelper.Init()

	_, err := kubeHelper.GetBadge("test_kube")
	if err != nil {
		t.Errorf("Error getting kubebadge: %v", err)
	}
}

func TestKubeHelper_KubeBadge(t *testing.T) {
	kubeHelper := NewKubeHelper()
	kubeHelper.Init()

	spec := v1.KubeBadgeSpec{
		Type:        "node",
		OriginalURL: "/kube/node/node1",
		AliasURL:    "/alias/testalias",
		Allowed:     true,
	}

	created, err := kubeHelper.CreateKubeBadge(spec)
	if err != nil {
		t.Fatalf("Error create kubebadge: %v", err)
	}

	got, err := kubeHelper.GetBadge(kubeHelper.GenerateKubeName(spec.OriginalURL))
	if err != nil {
		t.Fatalf("Error get kubebadge: %v", err)
	}

	if got.Spec.OriginalURL != created.Spec.OriginalURL {
		t.Fatalf("Error get kubebadge err, get %s but want %s", got.Spec.OriginalURL, created.Spec.OriginalURL)
	}

	got.Spec.Allowed = false
	newGot, err := kubeHelper.UpdateKubeBadge(got)
	if err != nil {
		t.Fatalf("Error get kubebadge: %v", err)
	}

	if newGot.Spec.Allowed {
		t.Fatalf("Error update kubebadge")
	}

	newGot2, err := kubeHelper.GetBadgeByLabel(v1.KubeBadgeLabelAliasURL, kubeHelper.GenerateKubeName(spec.AliasURL))
	if err != nil {
		t.Fatalf("Error get kubebadge: %v", err)
	}
	if len(newGot2) == 0 {
		t.Fatalf("Error get kubebadge by label")
	} else {
		if newGot2[0].Spec.AliasURL != spec.AliasURL {
			t.Fatalf("Error get kubebadge")
		}
	}

	if err := kubeHelper.DeleteKubeBadge(kubeHelper.GenerateKubeName(spec.OriginalURL)); err != nil {
		t.Fatalf("Error delete kubebadge: %v", err)
	}

	_, err = kubeHelper.GetBadge(kubeHelper.GenerateKubeName(spec.OriginalURL))
	if err == nil {
		t.Fatalf("Error delete kubebadge, get again")
	}
	if err != nil && !apierrors.IsNotFound(err) {
		t.Fatalf("Error delete kubebadge: %v", err)
	}
}
