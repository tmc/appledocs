// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewAlert

// ExampleNewAlertWithError demonstrates how to create a Alert instance using NewAlertWithError.
// Returns an alert initialized from information in an error object.
func ExampleNewAlertWithError() {
	_ = appkit.NewAlertWithError(
		appkit.Error{}, // error Error
	)
	// Output:
}
