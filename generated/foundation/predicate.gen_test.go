// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)


// ExampleNewPredicateWithFormatArgumentArray demonstrates how to create a Predicate instance using NewPredicateWithFormatArgumentArray.
// Creates a predicate by substituting the values in a specified array into a format string and parsing the result.
func ExampleNewPredicateWithFormatArgumentArray() {
	_ = foundation.NewPredicateWithFormatArgumentArray(
		"predicateFormat", // predicateFormat string
		nil, // arguments unsafe.Pointer
	)
	// Output:
}

// ExampleNewPredicateWithFormatArguments demonstrates how to create a Predicate instance using NewPredicateWithFormatArguments.
// Creates a predicate by substituting the values in an argument list into a format string and parsing the result.
func ExampleNewPredicateWithFormatArguments() {
	_ = foundation.NewPredicateWithFormatArguments(
		"predicateFormat", // predicateFormat string
		nil, // argList unsafe.Pointer
	)
	// Output:
}

// ExampleNewPredicateFromMetadataQueryString demonstrates how to create a Predicate instance using NewPredicateFromMetadataQueryString.
// Creates a predicate with a metadata query string.
func ExampleNewPredicateFromMetadataQueryString() {
	_ = foundation.NewPredicateFromMetadataQueryString(
		"queryString", // queryString string
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

// ExampleNewPredicateWithBlock demonstrates how to create a Predicate instance using NewPredicateWithBlock.
// Creates a predicate that evaluates using a specified block object and bindings dictionary.
func ExampleNewPredicateWithBlock() {
	_ = foundation.NewPredicateWithBlock(
		nil, // block unsafe.Pointer
	)
	// Output:
}


