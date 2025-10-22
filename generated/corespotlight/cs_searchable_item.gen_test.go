// Code generated from Apple documentation for CoreSpotlight. DO NOT EDIT.

package corespotlight_test

import (
	"github.com/tmc/appledocs/generated/corespotlight"
)

// Suppress unused import errors
var _ = corespotlight.NewCSSearchableItem

// ExampleNewCSSearchableItemWithUniqueIdentifierDomainIdentifierAttributeSet demonstrates how to create a CSSearchableItem instance using NewCSSearchableItemWithUniqueIdentifierDomainIdentifierAttributeSet.
// Returns a searchable item associated with the specified identifier, domain identifier, and attribute set.
func ExampleNewCSSearchableItemWithUniqueIdentifierDomainIdentifierAttributeSet() {
	_ = corespotlight.NewCSSearchableItemWithUniqueIdentifierDomainIdentifierAttributeSet(
		"uniqueIdentifier", // uniqueIdentifier string
		"domainIdentifier", // domainIdentifier string
		corespotlight.CSSearchableItemAttributeSet{}, // attributeSet CSSearchableItemAttributeSet
	)
	// Output:
}
