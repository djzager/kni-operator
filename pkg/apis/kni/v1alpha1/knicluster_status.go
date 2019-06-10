package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"time"
)

// ObjectReference contains enough information to let you inspect or modify the referred object.
type ObjectReference struct {
	// group of the referent.
	APIVersion string `json:"apiVersion"`
	// kind of the referent.
	Kind string `json:"kind"`
	// namespace of the referent.
	// +optional
	Namespace string `json:"namespace,omitempty"`
	// name of the referent.
	Name string `json:"name"`
}

// Condition represents the state of the operator's
// reconciliation functionality.
// +k8s:deepcopy-gen=true
type Condition struct {
	// type specifies the state of the operator's reconciliation functionality.
	Type ConditionType `json:"type"`

	// status of the condition, one of True, False, Unknown.
	Status ConditionStatus `json:"status"`

	// lastTransitionTime is the time of the last update to the current status object.
	LastTransitionTime metav1.Time `json:"lastTransitionTime"`

	// reason is the reason for the condition's last transition.  Reasons are CamelCase
	Reason ConditionReason `json:"reason,omitempty"`

	// message provides additional information about the current condition.
	// This is only to be consumed by humans.
	Message string `json:"message,omitempty"`
}

// ConditionType is the state of the operator's reconciliation functionality.
type ConditionType string

const (
	// ConditionTypeHealthy indicates the operator is healthy
	ConditionTypeHealthy ConditionType = "Healthy"
)

// ConditionStatus is whether, for the given Condition type, the status is True, False, Unknown.
type ConditionStatus string

const (
	// ConditionStatusTrue is it in the condition
	ConditionStatusTrue ConditionStatus = "True"
	// ConditionStatusFalse is it not in the condition
	ConditionStatusFalse ConditionStatus = "False"
	// ConditionStatusUnknown is we have no idea
	ConditionStatusUnknown ConditionStatus = "Unkown"
)

// ConditionReason is the reason a condition type is changing
type ConditionReason string

const (
	// ReconcileFailed means the operator failed to complete reconciliation
	ReconcileFailed ConditionReason = "ReconcileFailed"

	// ReconcileSucceeded means the operator completed reconciliation
	ReconcileSucceeded ConditionReason = "ReconcileSucceeded"
)

// SetObjectReference - updates list of object references based on newObject
func SetObjectReference(objects *[]ObjectReference, newObject ObjectReference) {
	if objects == nil {
		objects = &[]ObjectReference{}
	}
	existingObject := FindObjectReference(*objects, newObject)
	if existingObject == nil {
		*objects = append(*objects, newObject)
		return
	}
}

// FindObjectReference - finds an ObjectReference in a slice of objects
func FindObjectReference(objects []ObjectReference, object ObjectReference) *ObjectReference {
	for i := range objects {
		if objects[i].APIVersion == object.APIVersion && objects[i].Kind == object.Kind {
			return &objects[i]
		}
	}

	return nil
}

// SetCondition - updates conditions based on newCondition
func SetCondition(conditions *[]Condition, newCondition Condition) {
	if conditions == nil {
		conditions = &[]Condition{}
	}
	existingCondition := FindCondition(*conditions, newCondition.Type)
	if existingCondition == nil {
		newCondition.LastTransitionTime = metav1.NewTime(time.Now())
		*conditions = append(*conditions, newCondition)
		return
	}

	if existingCondition.Status != newCondition.Status {
		existingCondition.Status = newCondition.Status
		existingCondition.LastTransitionTime = metav1.NewTime(time.Now())
	}

	existingCondition.Reason = newCondition.Reason
	existingCondition.Message = newCondition.Message
}

// RemoveCondition - removes the condition in conditions matching conditionType
func RemoveCondition(conditions *[]Condition, conditionType ConditionType) {
	if conditions == nil {
		conditions = &[]Condition{}
	}
	newConditions := []Condition{}
	for _, condition := range *conditions {
		if condition.Type != conditionType {
			newConditions = append(newConditions, condition)
		}
	}

	*conditions = newConditions
}

// FindCondition - finds a Condition in a slice of conditions matching conditionType
func FindCondition(conditions []Condition, conditionType ConditionType) *Condition {
	for i := range conditions {
		if conditions[i].Type == conditionType {
			return &conditions[i]
		}
	}

	return nil
}
