// Code generated from Apple documentation for CoreSpotlight. DO NOT EDIT.

package corespotlight_test

import (
	"github.com/tmc/appledocs/generated/corespotlight"
)

// Suppress unused import errors
var _ = corespotlight.NewCSPerson

// ExampleNewCSPersonWithDisplayNameHandlesHandleIdentifier demonstrates how to create a CSPerson instance using NewCSPersonWithDisplayNameHandlesHandleIdentifier.
// Returns a new   object initialized with the specified display name and contact attributes.
func ExampleNewCSPersonWithDisplayNameHandlesHandleIdentifier() {
	_ = corespotlight.NewCSPersonWithDisplayNameHandlesHandleIdentifier(
		"displayName", // displayName string
		[]corespotlight.string{}, // handles []string
		"handleIdentifier", // handleIdentifier string
	)
	// Output:
}
