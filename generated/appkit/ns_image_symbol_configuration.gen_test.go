// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewImageSymbolConfiguration

// ExampleNewImageSymbolConfigurationWithPointSizeWeight demonstrates how to create a ImageSymbolConfiguration instance using NewImageSymbolConfigurationWithPointSizeWeight.
// Creates a symbol configuration with the specified point size and font weight.
func ExampleNewImageSymbolConfigurationWithPointSizeWeight() {
	_ = appkit.NewImageSymbolConfigurationWithPointSizeWeight(
		0.0, // pointSize float64
		appkit.FontWeight /* not a class type */{}, // weight FontWeight /* not a class type */
	)
	// Output:
}
