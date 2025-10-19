// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewPredicate





// ExampleNewPredicateFromMetadataQueryString demonstrates how to create a Predicate instance using NewPredicateFromMetadataQueryString.
// Creates a predicate with a metadata query string.
func ExampleNewPredicateFromMetadataQueryString() {
	_ = foundation.NewPredicateFromMetadataQueryString(
		"kMDItemFSName == '*.txt'", // queryString string
	)
	// Output:
}

// ExampleNewPredicateWithValue demonstrates how to create a Predicate instance using NewPredicateWithValue.
// Creates and returns a predicate that always evaluates to a specified Boolean value.
func ExampleNewPredicateWithValue() {
	_ = foundation.NewPredicateWithValue(
		false, // value bool
	)
	// Output:
}


