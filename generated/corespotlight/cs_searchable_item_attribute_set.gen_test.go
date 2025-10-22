// Code generated from Apple documentation for CoreSpotlight. DO NOT EDIT.

package corespotlight_test

import (
	"github.com/tmc/appledocs/generated/corespotlight"
)

// Suppress unused import errors
var _ = corespotlight.NewCSSearchableItemAttributeSet

// ExampleNewCSSearchableItemAttributeSetWithItemContentType demonstrates how to create a CSSearchableItemAttributeSet instance using NewCSSearchableItemAttributeSetWithItemContentType.
// Creates an attribute set for the specified content type.
func ExampleNewCSSearchableItemAttributeSetWithItemContentType() {
	_ = corespotlight.NewCSSearchableItemAttributeSetWithItemContentType(
		"itemContentType", // itemContentType string
	)
	// Output:
}
