// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewImageSymbolConfiguration

// ExampleNewImageSymbolConfigurationWithColorRenderingMode demonstrates how to create a ImageSymbolConfiguration instance using NewImageSymbolConfigurationWithColorRenderingMode.
// Create a configuration with a specific color rendering mode.
func ExampleNewImageSymbolConfigurationWithColorRenderingMode() {
	_ = appkit.NewImageSymbolConfigurationWithColorRenderingMode(
		appkit.ImageSymbolColorRenderingMode{}, // mode ImageSymbolColorRenderingMode
	)
	// Output:
}
// ExampleNewImageSymbolConfigurationWithHierarchicalColor demonstrates how to create a ImageSymbolConfiguration instance using NewImageSymbolConfigurationWithHierarchicalColor.
// Creates a hierarchical color configuration using the color you specify.
func ExampleNewImageSymbolConfigurationWithHierarchicalColor() {
	_ = appkit.NewImageSymbolConfigurationWithHierarchicalColor(
		appkit.NSColor{}, // hierarchicalColor NSColor
	)
	// Output:
}
// ExampleNewImageSymbolConfigurationWithPaletteColors demonstrates how to create a ImageSymbolConfiguration instance using NewImageSymbolConfigurationWithPaletteColors.
// Creates a color configuration by specifying a palette of colors.
func ExampleNewImageSymbolConfigurationWithPaletteColors() {
	_ = appkit.NewImageSymbolConfigurationWithPaletteColors(
		[]appkit.Color{}, // paletteColors []Color
	)
	// Output:
}
// ExampleNewImageSymbolConfigurationWithScale demonstrates how to create a ImageSymbolConfiguration instance using NewImageSymbolConfigurationWithScale.
// Creates a symbol configuration using the scale you specify.
func ExampleNewImageSymbolConfigurationWithScale() {
	_ = appkit.NewImageSymbolConfigurationWithScale(
		appkit.ImageSymbolScale{}, // scale ImageSymbolScale
	)
	// Output:
}
// ExampleNewImageSymbolConfigurationWithVariableValueMode demonstrates how to create a ImageSymbolConfiguration instance using NewImageSymbolConfigurationWithVariableValueMode.
// Create a configuration with a specified variable value mode.
func ExampleNewImageSymbolConfigurationWithVariableValueMode() {
	_ = appkit.NewImageSymbolConfigurationWithVariableValueMode(
		appkit.ImageSymbolVariableValueMode{}, // variableValueMode ImageSymbolVariableValueMode
	)
	// Output:
}
