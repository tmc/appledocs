// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewURLConnection

// ExampleURLConnection_Cancel demonstrates using Cancel on a URLConnection instance.
// Cancels an asynchronous load of a request.
func ExampleURLConnection_Cancel() {
	obj := foundation.NewURLConnection()
	obj.Cancel()
	// Output:
	}

// ExampleURLConnection_Start demonstrates using Start on a URLConnection instance.
// Causes the connection to begin loading data, if it has not already.
func ExampleURLConnection_Start() {
	obj := foundation.NewURLConnection()
	obj.Start()
	// Output:
	}

