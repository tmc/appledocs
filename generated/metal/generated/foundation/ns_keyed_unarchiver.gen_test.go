// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewKeyedUnarchiver

// ExampleNewKeyedUnarchiver demonstrates how to create a KeyedUnarchiver instance.
// Initializes an archiver to decode data.
func ExampleNewKeyedUnarchiver() {
	_ = foundation.NewKeyedUnarchiver()
	// Output:
}
// ExampleKeyedUnarchiver_FinishDecoding demonstrates using FinishDecoding on a KeyedUnarchiver instance.
// Tells the receiver that you are finished decoding objects.
func ExampleKeyedUnarchiver_FinishDecoding() {
	obj := foundation.NewKeyedUnarchiver()
	obj.FinishDecoding()
	// Output:
	}

