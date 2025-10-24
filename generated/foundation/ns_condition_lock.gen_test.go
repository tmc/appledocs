// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewConditionLock

// ExampleNewConditionLockWithCondition demonstrates how to create a ConditionLock instance using NewConditionLockWithCondition.
// Initializes a newly allocated   object and sets its condition.
func ExampleNewConditionLockWithCondition() {
	_ = foundation.NewConditionLockWithCondition(
		0, // condition int
	)
	// Output:
}
// ExampleConditionLock_TryLock demonstrates using TryLock on a ConditionLock instance.
// Attempts to acquire a lock without regard to the receiver’s condition.
func ExampleConditionLock_TryLock() {
	obj := foundation.NewConditionLock()
	_ = obj.TryLock()
	// Output:
	}

