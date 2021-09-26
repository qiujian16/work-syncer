package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	workapiv1 "open-cluster-management.io/api/work/v1"
)

// StatusSyncerSpec defines the spec of status syncer
type StatusSyncSpec struct {
	// ObjectReference represents the reference to the object to apply the resources.
	ObjectReference ObjectRef `json:"objectRef"`
}

type ObjectRef struct {
	Group   string `json:"group"`
	Version string `json:"version"`
	Name    string `json:"name"`
}

type StatusSyncStatus struct {
	SyncedStatuses []SyncStatus       `json:"synStatuses,omitempty"`
	Conditions     []metav1.Condition `json:"conditions"`
}

type SyncStatus struct {
	// ResourceMeta represents the group, version, kind, name and namespace of a resoure.
	// +required
	ResourceMeta workapiv1.ManifestResourceMeta `json:"resourceMeta"`

	AvailableReplica *int32 `json:"availableReplica,omitempty"`

	// Conditions represents the conditions of this resource on a managed cluster.
	// +required
	Conditions []metav1.Condition `json:"conditions"`
}

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// StatusSync is the Schema for the StatusSync API
type StatusSync struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// spec defines the what objects to be synced.
	// +optional
	Spec StatusSyncSpec `json:"spec,omitempty"`
	// status defines the status of each applied manifest on the spoke cluster.
	Status StatusSyncStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WorkList contains a list of Work
type StatusSyncList struct {
	metav1.TypeMeta `json:",inline"`
	// Standard list metadata.
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`
	// List of works.
	// +listType=set
	Items []StatusSync `json:"items"`
}
