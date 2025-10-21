// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit_test

import (
	"github.com/tmc/appledocs/generated/callkit"
)

// Suppress unused import errors
var _ = callkit.NewCXProviderConfiguration

// ExampleNewCXProviderConfiguration demonstrates how to create a CXProviderConfiguration instance.
// Creates the configuration of a provider object.
func ExampleNewCXProviderConfiguration() {
	_ = callkit.NewCXProviderConfiguration()
	// Output:
}
// ExampleNewCXProviderConfigurationWithLocalizedName demonstrates how to create a CXProviderConfiguration instance using NewCXProviderConfigurationWithLocalizedName.
// Initializes a configuration with the specified localized name.
func ExampleNewCXProviderConfigurationWithLocalizedName() {
	_ = callkit.NewCXProviderConfigurationWithLocalizedName(
		"localizedName", // localizedName string
	)
	// Output:
}
