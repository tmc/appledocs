// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit_test

import (
	"github.com/tmc/appledocs/generated/cloudkit"
)

// Suppress unused import errors
var _ = cloudkit.NewCKSystemSharingUIObserver

// ExampleNewCKSystemSharingUIObserverWithContainer demonstrates how to create a CKSystemSharingUIObserver instance using NewCKSystemSharingUIObserverWithContainer.
// Creates and initializes an observer using the provided container.
func ExampleNewCKSystemSharingUIObserverWithContainer() {
	_ = cloudkit.NewCKSystemSharingUIObserverWithContainer(
		cloudkit.CKContainer{}, // container CKContainer
	)
	// Output:
}
