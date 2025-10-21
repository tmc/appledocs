// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit_test

import (
	"github.com/tmc/appledocs/generated/callkit"
)

// Suppress unused import errors
var _ = callkit.NewCXProvider

// ExampleNewCXProviderWithConfiguration demonstrates how to create a CXProvider instance using NewCXProviderWithConfiguration.
// Initializes a new provider with the specified configuration.
func ExampleNewCXProviderWithConfiguration() {
	_ = callkit.NewCXProviderWithConfiguration(
		callkit.CXProviderConfiguration{}, // configuration CXProviderConfiguration
	)
	// Output:
}
