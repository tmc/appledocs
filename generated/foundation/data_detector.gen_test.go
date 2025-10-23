// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewDataDetector

// ExampleNewDataDetectorWithTypesError demonstrates how to create a DataDetector instance using NewDataDetectorWithTypesError.
// Initializes and returns a data detector instance.
func ExampleNewDataDetectorWithTypesError() {
	_ = foundation.NewDataDetectorWithTypesError(
		foundation.TextCheckingTypes{}, // checkingTypes TextCheckingTypes
		foundation.NSError{}, // error NSError
	)
	// Output:
}
