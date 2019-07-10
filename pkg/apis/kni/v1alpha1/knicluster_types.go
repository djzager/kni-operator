package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// KNIClusterSpec defines the desired state of KNICluster
// +k8s:openapi-gen=true
type KNIClusterSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
}

// KNIClusterStatus defines the observed state of KNICluster
// +k8s:openapi-gen=true
type KNIClusterStatus struct {
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
	// conditions describes the state of the operator's reconciliation functionality.
	// +patchMergeKey=type
	// +patchStrategy=merge
	// +optional

	// Conditions is a list of conditions related to operator reconciliation
	Conditions []Condition `json:"conditions,omitempty"  patchStrategy:"merge" patchMergeKey:"type"`
	// RelatedObjects is a list of objects that are "interesting" or related to this operator.
	RelatedObjects []corev1.ObjectReference `json:"relatedObjects,omitempty"`
}

// Condition represents the state of the operator's
// reconciliation functionality.
// +k8s:deepcopy-gen=true
type Condition struct {
	// type specifies the state of the operator's reconciliation functionality.
	Type ConditionType `json:"type"`

	// status of the condition, one of True, False, Unknown.
	Status corev1.ConditionStatus `json:"status"`

	// lastTransitionTime is the time of the last update to the current status object.
	LastTransitionTime metav1.Time `json:"lastTransitionTime"`

	// reason is the reason for the condition's last transition.  Reasons are CamelCase
	Reason string `json:"reason,omitempty"`

	// message provides additional information about the current condition.
	// This is only to be consumed by humans.
	Message string `json:"message,omitempty"`
}

// ConditionType is the state of the operator's reconciliation functionality.
type ConditionType string

const (
	// ConditionAvailable indicates that the resources maintained by the operator,
	// is functional and available in the cluster.
	ConditionAvailable ConditionType = "Available"

	// ConditionProgressing indicates that the operator is actively making changes to the resources maintained by the
	// operator
	ConditionProgressing ConditionType = "Progressing"

	// ConditionDegraded indicates that the resources maintained by the operator are not functioning completely.
	// An example of a degraded state would be if not all pods in a deployment were running.
	// It may still be available, but it is degraded
	ConditionDegraded ConditionType = "Degraded"

	// ConditionUpgradeable indicates whether the resources maintained by the operator are in a state that is safe to upgrade.
	// When `False`, the resources maintained by the operator should not be upgraded and the
	// message field should contain a human readable description of what the administrator should do to
	// allow the operator to successfully update the resources maintained by the operator.
	ConditionUpgradeable ConditionType = "Upgradeable"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KNICluster is the Schema for the kniclusters API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
type KNICluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KNIClusterSpec   `json:"spec,omitempty"`
	Status KNIClusterStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KNIClusterList contains a list of KNICluster
type KNIClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KNICluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KNICluster{}, &KNIClusterList{})
}
