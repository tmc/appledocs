// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewUserScriptTask

// ExampleNewUserScriptTaskWithURLError demonstrates how to create a UserScriptTask instance using NewUserScriptTaskWithURLError.
// Return a user script task instance given a URL for a script file.
func ExampleNewUserScriptTaskWithURLError() {
	_ = foundation.NewUserScriptTaskWithURLError(
		foundation.URL{}, // url URL
		foundation.NSError{}, // error NSError
	)
	// Output:
}
