// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewKeyedArchiver

// ExampleNewKeyedArchiver demonstrates how to create a KeyedArchiver instance.
// Initializes an archiver to encode data.
func ExampleNewKeyedArchiver() {
	_ = foundation.NewKeyedArchiver()
	// Output:
}
// ExampleNewKeyedArchiverRequiringSecureCoding demonstrates how to create a KeyedArchiver instance using NewKeyedArchiverRequiringSecureCoding.
// Creates an archiver to encode data, and optionally disables secure coding.
func ExampleNewKeyedArchiverRequiringSecureCoding() {
	_ = foundation.NewKeyedArchiverRequiringSecureCoding(
		false, // requiresSecureCoding bool
	)
	// Output:
}
// ExampleKeyedArchiver_FinishEncoding demonstrates using FinishEncoding on a KeyedArchiver instance.
// Instructs the receiver to construct the final data stream.
func ExampleKeyedArchiver_FinishEncoding() {
	obj := foundation.NewKeyedArchiver()
	obj.FinishEncoding()
	// Output:
	}

