// Code generated from Apple documentation for CoreSpotlight. DO NOT EDIT.

package corespotlight_test

import (
	"github.com/tmc/appledocs/generated/corespotlight"
)

// Suppress unused import errors
var _ = corespotlight.NewCSCustomAttributeKey

// ExampleNewCSCustomAttributeKeyWithKeyName demonstrates how to create a CSCustomAttributeKey instance using NewCSCustomAttributeKeyWithKeyName.
// Returns a new custom attribute key with the specified name.
func ExampleNewCSCustomAttributeKeyWithKeyName() {
	_ = corespotlight.NewCSCustomAttributeKeyWithKeyName(
		"keyName", // keyName string
	)
	// Output:
}
// ExampleNewCSCustomAttributeKeyWithKeyNameSearchableSearchableByDefaultUniqueMultiValued demonstrates how to create a CSCustomAttributeKey instance using NewCSCustomAttributeKeyWithKeyNameSearchableSearchableByDefaultUniqueMultiValued.
// Returns a new custom attribute key with the specified name and properties.
func ExampleNewCSCustomAttributeKeyWithKeyNameSearchableSearchableByDefaultUniqueMultiValued() {
	_ = corespotlight.NewCSCustomAttributeKeyWithKeyNameSearchableSearchableByDefaultUniqueMultiValued(
		"keyName", // keyName string
		false, // searchable bool
		false, // searchableByDefault bool
		false, // unique bool
		false, // multiValued bool
	)
	// Output:
}
