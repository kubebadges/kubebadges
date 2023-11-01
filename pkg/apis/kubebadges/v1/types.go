package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	KubeBadgeLabelType           string = "type"
	KubeBadgeLabelAllowed        string = "allowed"
	KubeBadgeLabelAliasURL       string = "aliasURL"
	KubeBadgeLabelOriginalURL    string = "originalURL"
	KubeBadgeLabelDisplayName    string = "displayName"
	KubeBadgeLabelOwnerNamespace string = "ownerNamespace"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:resource:shortName=kubebadge
// +kubebuilder:printcolumn:name="Type",type=string,JSONPath=`.spec.type`
// +kubebuilder:printcolumn:name="OriginalURL",type=string,JSONPath=`.spec.originalURL`
// +kubebuilder:printcolumn:name="DisplayName",type=string,JSONPath=`.spec.displayName`
// +kubebuilder:printcolumn:name="OwnerNamespace",type=string,JSONPath=`.spec.ownerNamespace`
// +kubebuilder:printcolumn:name="Allowed",type=boolean,JSONPath=`.spec.allowed`

// KubeBadge is the Schema for the kubebadges API.
type KubeBadge struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec KubeBadgeSpec `json:"spec"`
}

// KubeBadgeSpec defines the desired state of KubeBadge.
type KubeBadgeSpec struct {
	// +required
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Description="Type represents the type of the badge."
	Type string `json:"type"`

	// +required
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Description="OriginalURL is the original URL of the badge."
	OriginalURL string `json:"originalURL"`

	// +optional
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Description="AliasURL is an optional alias URL of the badge."
	AliasURL string `json:"aliasURL"`

	// +optional
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Description="DisplayName is an optional display name of the badge."
	DisplayName string `json:"displayName"`

	// +optional
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Description="OwnerNamespace is an optional namespace of the badge owner."
	OwnerNamespace string `json:"ownerNamespace"`

	// +required
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Type=boolean
	// +kubebuilder:validation:Description="Allowed specifies if the badge is allowed to public access."
	Allowed bool `json:"allowed"`

	// +optional
	Custom Custom `json:"custom,omitempty"`
}

type Custom struct {
	// +optional
	Type string `json:"type"`

	// +optional
	Address string `json:"address"`

	// +optional
	Port int `json:"port"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// KubeBadgeList contains a list of KubeBadge.
type KubeBadgeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []KubeBadge `json:"items"`
}
