// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit_test

import (
	"github.com/tmc/appledocs/generated/callkit"
)

// Suppress unused import errors
var _ = callkit.NewCXCallController


// ExampleNewCXCallController demonstrates how to create a CXCallController instance.
// Initializes a new call controller with a private, serial queue, which is used for calling completion blocks.
func ExampleNewCXCallController() {
	_ = callkit.NewCXCallController()
	// Output:
}



