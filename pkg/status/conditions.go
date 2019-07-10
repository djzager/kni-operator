package status

import (
	"time"

	kniv1alpha1 "github.com/mhrivnak/kni-operator/pkg/apis/kni/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// SetStatusCondition sets the corresponding condition in conditions to newCondition.
func SetStatusCondition(conditions *[]kniv1alpha1.Condition, newCondition kniv1alpha1.Condition) {
	if conditions == nil {
		conditions = &[]kniv1alpha1.Condition{}
	}
	existingCondition := FindStatusCondition(*conditions, newCondition.Type)
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

// RemoveStatusCondition removes the corresponding conditionType from conditions.
func RemoveStatusCondition(conditions *[]kniv1alpha1.Condition, conditionType kniv1alpha1.ConditionType) {
	if conditions == nil {
		return
	}
	newConditions := []kniv1alpha1.Condition{}
	for _, condition := range *conditions {
		if condition.Type != conditionType {
			newConditions = append(newConditions, condition)
		}
	}

	*conditions = newConditions
}

// FindStatusCondition finds the conditionType in conditions.
func FindStatusCondition(conditions []kniv1alpha1.Condition, conditionType kniv1alpha1.ConditionType) *kniv1alpha1.Condition {
	for i := range conditions {
		if conditions[i].Type == conditionType {
			return &conditions[i]
		}
	}

	return nil
}

// IsStatusConditionTrue returns true when the conditionType is present and set to `corev1.ConditionTrue`
func IsStatusConditionTrue(conditions []kniv1alpha1.Condition, conditionType kniv1alpha1.ConditionType) bool {
	return IsStatusConditionPresentAndEqual(conditions, conditionType, corev1.ConditionTrue)
}

// IsStatusConditionFalse returns true when the conditionType is present and set to `corev1.ConditionFalse`
func IsStatusConditionFalse(conditions []kniv1alpha1.Condition, conditionType kniv1alpha1.ConditionType) bool {
	return IsStatusConditionPresentAndEqual(conditions, conditionType, corev1.ConditionFalse)
}

// IsStatusConditionUnknown returns true when the conditionType is present and set to `corev1.ConditionUnknown`
func IsStatusConditionUnknown(conditions []kniv1alpha1.Condition, conditionType kniv1alpha1.ConditionType) bool {
	return IsStatusConditionPresentAndEqual(conditions, conditionType, corev1.ConditionUnknown)
}

// IsStatusConditionPresentAndEqual returns true when conditionType is present and equal to status.
func IsStatusConditionPresentAndEqual(conditions []kniv1alpha1.Condition, conditionType kniv1alpha1.ConditionType, status corev1.ConditionStatus) bool {
	for _, condition := range conditions {
		if condition.Type == conditionType {
			return condition.Status == status
		}
	}
	return false
}
