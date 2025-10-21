// Code generated from Apple documentation for InputMethodKit. DO NOT EDIT.

package inputmethodkit_test

import (
	"github.com/tmc/appledocs/generated/inputmethodkit"
)

// Suppress unused import errors
var _ = inputmethodkit.NewIMKServer

// ExampleNewIMKServerWithNameBundleIdentifier demonstrates how to create a IMKServer instance using NewIMKServerWithNameBundleIdentifier.
// Creates and returns a server object from property list information contained in the provided bundle.
func ExampleNewIMKServerWithNameBundleIdentifier() {
	_ = inputmethodkit.NewIMKServerWithNameBundleIdentifier(
		"name", // name string
		"bundleIdentifier", // bundleIdentifier string
	)
	// Output:
}

