// Code generated from Apple documentation for CoreWLAN. DO NOT EDIT.

package corewlan_test

import (
	"github.com/tmc/appledocs/generated/corewlan"
)

// Suppress unused import errors
var _ = corewlan.NewCWConfiguration

// ExampleNewCWConfiguration demonstrates how to create a CWConfiguration instance.
// Creates an empty CWConfiguration object.
func ExampleNewCWConfiguration() {
	_ = corewlan.NewCWConfiguration()
	// Output:
}
// ExampleNewCWConfigurationWithConfiguration demonstrates how to create a CWConfiguration instance using NewCWConfigurationWithConfiguration.
// Creates and returns a CWConfiguration object initialized with the given CWConfiguration object.
func ExampleNewCWConfigurationWithConfiguration() {
	_ = corewlan.NewCWConfigurationWithConfiguration(
		corewlan.CWConfiguration{}, // configuration CWConfiguration
	)
	// Output:
}
