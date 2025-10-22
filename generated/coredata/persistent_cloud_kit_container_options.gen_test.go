// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata_test

import (
	"github.com/tmc/appledocs/generated/coredata"
)

// Suppress unused import errors
var _ = coredata.NewPersistentCloudKitContainerOptions

// ExampleNewPersistentCloudKitContainerOptionsWithContainerIdentifier demonstrates how to create a PersistentCloudKitContainerOptions instance using NewPersistentCloudKitContainerOptionsWithContainerIdentifier.
// Initializes container options using the given CloudKit container identifier.
func ExampleNewPersistentCloudKitContainerOptionsWithContainerIdentifier() {
	_ = coredata.NewPersistentCloudKitContainerOptionsWithContainerIdentifier(
		"containerIdentifier", // containerIdentifier string
	)
	// Output:
}
