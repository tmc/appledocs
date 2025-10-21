// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewSharingService

// ExampleNewSharingServiceNamed demonstrates how to create a SharingService instance using NewSharingServiceNamed.
// Returns a sharing service instance representing the specified service name.
func ExampleNewSharingServiceNamed() {
	_ = appkit.NewSharingServiceNamed(
		appkit.SharingServiceName{}, // serviceName SharingServiceName
	)
	// Output:
}
