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
// ExampleNewImageSymbolConfigurationWithPaletteColors demonstrates how to create a ImageSymbolConfiguration instance using NewImageSymbolConfigurationWithPaletteColors.
// Creates a color configuration by specifying a palette of colors.
func ExampleNewImageSymbolConfigurationWithPaletteColors() {
	_ = appkit.NewImageSymbolConfigurationWithPaletteColors(
		[]appkit.Color{}, // paletteColors []Color
	)
	// Output:
}
// ExampleNewImageSymbolConfigurationWithPointSizeWeight demonstrates how to create a ImageSymbolConfiguration instance using NewImageSymbolConfigurationWithPointSizeWeight.
// Creates a symbol configuration with the specified point size and font weight.
func ExampleNewImageSymbolConfigurationWithPointSizeWeight() {
	_ = appkit.NewImageSymbolConfigurationWithPointSizeWeight(
		0.0, // pointSize float64
		appkit.FontWeight /* typedef */{}, // weight FontWeight /* typedef */
	)
	// Output:
}
// ExampleNewImageSymbolConfigurationWithPointSizeWeightScale demonstrates how to create a ImageSymbolConfiguration instance using NewImageSymbolConfigurationWithPointSizeWeightScale.
// Creates a symbol configuration with the specified point size, font weight, and symbol scale.
func ExampleNewImageSymbolConfigurationWithPointSizeWeightScale() {
	_ = appkit.NewImageSymbolConfigurationWithPointSizeWeightScale(
		0.0, // pointSize float64
		appkit.FontWeight /* typedef */{}, // weight FontWeight /* typedef */
		appkit.ImageSymbolScale{}, // scale ImageSymbolScale
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
// ExampleNewImageSymbolConfigurationWithTextStyle demonstrates how to create a ImageSymbolConfiguration instance using NewImageSymbolConfigurationWithTextStyle.
// Creates a symbol configuration with the specified text style.
func ExampleNewImageSymbolConfigurationWithTextStyle() {
	_ = appkit.NewImageSymbolConfigurationWithTextStyle(
		appkit.FontTextStyle /* typedef */{}, // style FontTextStyle /* typedef */
	)
	// Output:
}
// ExampleNewImageSymbolConfigurationWithTextStyleScale demonstrates how to create a ImageSymbolConfiguration instance using NewImageSymbolConfigurationWithTextStyleScale.
// Creates a symbol configuration with the specified text style and symbol scale.
func ExampleNewImageSymbolConfigurationWithTextStyleScale() {
	_ = appkit.NewImageSymbolConfigurationWithTextStyleScale(
		appkit.FontTextStyle /* typedef */{}, // style FontTextStyle /* typedef */
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
