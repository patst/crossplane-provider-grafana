/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type GrafanaDashboardObservation struct {

	// The numeric ID of the dashboard computed by Grafana.
	DashboardID *float64 `json:"dashboardId,omitempty" tf:"dashboard_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// URL friendly version of the dashboard title. This field is deprecated, please use `uid` instead.
	Slug *string `json:"slug,omitempty" tf:"slug,omitempty"`

	// The unique identifier of a dashboard. This is used to construct its URL. It's automatically generated if not provided when creating a dashboard. The uid allows having consistent URLs for accessing dashboards and when syncing dashboards between multiple Grafana installs.
	UID *string `json:"uid,omitempty" tf:"uid,omitempty"`

	// The full URL of the dashboard.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`

	// Whenever you save a version of your dashboard, a copy of that version is saved so that previous versions of your dashboard are not lost.
	Version *float64 `json:"version,omitempty" tf:"version,omitempty"`
}

type GrafanaDashboardParameters struct {

	// The complete dashboard model JSON.
	// +kubebuilder:validation:Required
	ConfigJSON *string `json:"configJson" tf:"config_json,omitempty"`

	// The id of the folder to save the dashboard in. This attribute is a string to reflect the type of the folder's id.
	// +kubebuilder:validation:Optional
	Folder *string `json:"folder,omitempty" tf:"folder,omitempty"`

	// Set a commit message for the version history.
	// +kubebuilder:validation:Optional
	Message *string `json:"message,omitempty" tf:"message,omitempty"`

	// Set to true if you want to overwrite existing dashboard with newer version, same dashboard title in folder or same dashboard uid.
	// +kubebuilder:validation:Optional
	Overwrite *bool `json:"overwrite,omitempty" tf:"overwrite,omitempty"`
}

// GrafanaDashboardSpec defines the desired state of GrafanaDashboard
type GrafanaDashboardSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GrafanaDashboardParameters `json:"forProvider"`
}

// GrafanaDashboardStatus defines the observed state of GrafanaDashboard.
type GrafanaDashboardStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GrafanaDashboardObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GrafanaDashboard is the Schema for the GrafanaDashboards API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,grafana}
type GrafanaDashboard struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GrafanaDashboardSpec   `json:"spec"`
	Status            GrafanaDashboardStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GrafanaDashboardList contains a list of GrafanaDashboards
type GrafanaDashboardList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GrafanaDashboard `json:"items"`
}

// Repository type metadata.
var (
	GrafanaDashboard_Kind             = "GrafanaDashboard"
	GrafanaDashboard_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GrafanaDashboard_Kind}.String()
	GrafanaDashboard_KindAPIVersion   = GrafanaDashboard_Kind + "." + CRDGroupVersion.String()
	GrafanaDashboard_GroupVersionKind = CRDGroupVersion.WithKind(GrafanaDashboard_Kind)
)

func init() {
	SchemeBuilder.Register(&GrafanaDashboard{}, &GrafanaDashboardList{})
}