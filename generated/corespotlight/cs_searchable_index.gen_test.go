// Code generated from Apple documentation for CoreSpotlight. DO NOT EDIT.

package corespotlight_test

import (
	"github.com/tmc/appledocs/generated/corespotlight"
)

// Suppress unused import errors
var _ = corespotlight.NewCSSearchableIndex

// ExampleNewCSSearchableIndexWithName demonstrates how to create a CSSearchableIndex instance using NewCSSearchableIndexWithName.
// Returns an on-device index with the specified name.
func ExampleNewCSSearchableIndexWithName() {
	_ = corespotlight.NewCSSearchableIndexWithName(
		"name", // name string
	)
	// Output:
}
// ExampleNewCSSearchableIndexWithNameProtectionClass demonstrates how to create a CSSearchableIndex instance using NewCSSearchableIndexWithNameProtectionClass.
// Returns an on-device index with the specified name and data protection class.
func ExampleNewCSSearchableIndexWithNameProtectionClass() {
	_ = corespotlight.NewCSSearchableIndexWithNameProtectionClass(
		"name", // name string
		corespotlight.FileProtectionType{}, // protectionClass FileProtectionType
	)
	// Output:
}
