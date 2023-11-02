package controller

import (
	"reflect"
	"testing"

	"github.com/kubebadges/kubebadges/internal/model"
)

func TestKubeController_parseKey(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name             string
		kubeController   *KubeController
		args             args
		wantResourceType string
		wantNamespace    string
		wantName         string
	}{
		{
			name:           "node",
			kubeController: &KubeController{},
			args: args{
				key: "/kube/node/minikube",
			},
			wantResourceType: "node",
			wantNamespace:    "",
			wantName:         "minikube",
		},
		{
			name:           "namespace",
			kubeController: &KubeController{},
			args: args{
				key: "/kube/namespace/default",
			},
			wantResourceType: "namespace",
			wantNamespace:    "",
			wantName:         "default",
		},
		{
			name:           "deployment",
			kubeController: &KubeController{},
			args: args{
				key: "/kube/deployment/default/nginx",
			},
			wantResourceType: "deployment",
			wantNamespace:    "default",
			wantName:         "nginx",
		},
		{
			name:           "pod",
			kubeController: &KubeController{},
			args: args{
				key: "/kube/pod/default/nginx-123",
			},
			wantResourceType: "pod",
			wantNamespace:    "default",
			wantName:         "nginx-123",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResourceType, gotNamespace, gotName := tt.kubeController.parseKey(tt.args.key)
			if gotResourceType != tt.wantResourceType {
				t.Errorf("KubeController.parseKey() gotResourceType = %v, want %v", gotResourceType, tt.wantResourceType)
			}
			if gotNamespace != tt.wantNamespace {
				t.Errorf("KubeController.parseKey() gotNamespace = %v, want %v", gotNamespace, tt.wantNamespace)
			}
			if gotName != tt.wantName {
				t.Errorf("KubeController.parseKey() gotName = %v, want %v", gotName, tt.wantName)
			}
		})
	}
}

func TestKubeController_configToMap(t *testing.T) {
	type args struct {
		config *model.KubeBadgesConfig
	}
	tests := []struct {
		name           string
		kubeController *KubeController
		args           args
		want           map[string]string
	}{
		{
			name:           "empty config",
			kubeController: &KubeController{},
			args: args{
				config: &model.KubeBadgesConfig{},
			},
			want: map[string]string{
				"badge_base_url": "",
			},
		},
		{
			name:           "non-empty config",
			kubeController: &KubeController{},
			args: args{
				config: &model.KubeBadgesConfig{
					BadgeBaseURL: "https://example.com",
				},
			},
			want: map[string]string{
				"badge_base_url": "https://example.com",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.kubeController.configToMap(tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KubeController.configToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKubeController_mapToConfig(t *testing.T) {
	tests := []struct {
		name           string
		kubeController *KubeController
		configMap      map[string]string
		want           *model.KubeBadgesConfig
	}{
		{
			name:           "empty map",
			kubeController: &KubeController{},
			configMap:      map[string]string{},
			want:           &model.KubeBadgesConfig{},
		},
		{
			name:           "non-empty map",
			kubeController: &KubeController{},
			configMap: map[string]string{
				"badge_base_url": "https://example.com",
			},
			want: &model.KubeBadgesConfig{
				BadgeBaseURL: "https://example.com",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.kubeController.mapToConfig(tt.configMap); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KubeController.mapToConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
