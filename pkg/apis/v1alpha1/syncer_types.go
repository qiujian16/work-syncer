package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	workapiv1 "open-cluster-management.io/api/work/v1"
)

// StatusSyncerSpec defines the spec of status syncer
type StatusSyncSpec struct {
	// ObjectReference represents the reference to the object to apply the resources.
	ObjectReference ObjectRef `json:"objectRef"`

	InterestedResources []InterestedResource `json:"interesedResource,omitempty"`
}

type ObjectRef struct {
	Group   string `json:"group"`
	Version string `json:"version"`
	Name    string `json:"name"`
}

type InterestedResource struct {
	// ResourceMeta represents the group, version, kind, name and namespace of a resoure.
	// +required
	ResourceMeta workapiv1.AppliedManifestResourceMeta `json:"resourceMeta"`

	InterestedFieldPaths []string `json:"interestedFieldPaths,omitempty"`
}

type StatusSyncStatus struct {
	SyncedStatuses []SyncStatus       `json:"synStatuses,omitempty"`
	Conditions     []metav1.Condition `json:"conditions"`
}

type SyncStatus struct {
	// ResourceMeta represents the group, version, kind, name and namespace of a resoure.
	// +required
	ResourceMeta workapiv1.AppliedManifestResourceMeta `json:"resourceMeta"`

	InterestedValues []InterestedValue `json:"interestedValues"`

	Conditions []metav1.Condition `json:"conditions"`
}

type InterestedValue struct {
	FieldPath string `json:"fieldPath"`

	Value FieldValue `json:"fieldValue"`
}

type FieldValue struct {
	Type ValueType `json:"type"`

	// +optional
	Integer int32 `json:"integer,omitempty"`

	// +optional
	String string `json:"string,omitempty"`

	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

type ValueType string

const (
	Integer    ValueType = "Integer"
	String     ValueType = "String"
	Conditions ValueType = "Conditions"
)

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
	// +optional
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
