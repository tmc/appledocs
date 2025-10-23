// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewColor

// ExampleNewColorFromPasteboard demonstrates how to create a Color instance using NewColorFromPasteboard.
// Creates a color object from color data currently on the pasteboard.
func ExampleNewColorFromPasteboard() {
	_ = appkit.NewColorFromPasteboard(
		appkit.NSPasteboard{}, // pasteBoard NSPasteboard
	)
	// Output:
}
// ExampleNewColorNamedBundle demonstrates how to create a Color instance using NewColorNamedBundle.
// Creates a color object from the provided name, which corresponds to a color in the default asset catalog of the specified bundle.
func ExampleNewColorNamedBundle() {
	_ = appkit.NewColorNamedBundle(
		appkit.ColorName{}, // name ColorName
		appkit.Bundle{}, // bundle Bundle
	)
	// Output:
}
// ExampleNewColorWithCatalogNameColorName demonstrates how to create a Color instance using NewColorWithCatalogNameColorName.
// Creates a color object using the specified asset catalog and color names.
func ExampleNewColorWithCatalogNameColorName() {
	_ = appkit.NewColorWithCatalogNameColorName(
		appkit.ColorListName{}, // listName ColorListName
		appkit.ColorName{}, // colorName ColorName
	)
	// Output:
}
// ExampleNewColorWithCoder demonstrates how to create a Color instance using NewColorWithCoder.
// Creates a color object from data in an unarchiver.
func ExampleNewColorWithCoder() {
	_ = appkit.NewColorWithCoder(
		appkit.Coder{}, // coder Coder
	)
	// Output:
}
// ExampleNewColorWithColorSpaceHueSaturationBrightnessAlpha demonstrates how to create a Color instance using NewColorWithColorSpaceHueSaturationBrightnessAlpha.
// Creates a color object with the specified color space, hue, saturation, brightness, and alpha channel values.
func ExampleNewColorWithColorSpaceHueSaturationBrightnessAlpha() {
	_ = appkit.NewColorWithColorSpaceHueSaturationBrightnessAlpha(
		appkit.NSColorSpace{}, // space NSColorSpace
		0.0, // hue float64
		0.0, // saturation float64
		0.0, // brightness float64
		0.0, // alpha float64
	)
	// Output:
}
// ExampleNewColorWithDeviceCyanMagentaYellowBlackAlpha demonstrates how to create a Color instance using NewColorWithDeviceCyanMagentaYellowBlackAlpha.
// Creates a color object using the given opacity value and CMYK components.
func ExampleNewColorWithDeviceCyanMagentaYellowBlackAlpha() {
	_ = appkit.NewColorWithDeviceCyanMagentaYellowBlackAlpha(
		0.0, // cyan float64
		0.0, // magenta float64
		0.0, // yellow float64
		0.0, // black float64
		0.0, // alpha float64
	)
	// Output:
}
// ExampleNewColorWithDeviceRedGreenBlueAlpha demonstrates how to create a Color instance using NewColorWithDeviceRedGreenBlueAlpha.
// Creates a color object using the given opacity value and RGB components.
func ExampleNewColorWithDeviceRedGreenBlueAlpha() {
	_ = appkit.NewColorWithDeviceRedGreenBlueAlpha(
		0.0, // red float64
		0.0, // green float64
		0.0, // blue float64
		0.0, // alpha float64
	)
	// Output:
}
// ExampleNewColorWithDeviceWhiteAlpha demonstrates how to create a Color instance using NewColorWithDeviceWhiteAlpha.
// Creates a color object using the given opacity and grayscale values.
func ExampleNewColorWithDeviceWhiteAlpha() {
	_ = appkit.NewColorWithDeviceWhiteAlpha(
		0.0, // white float64
		0.0, // alpha float64
	)
	// Output:
}
// ExampleNewColorWithDisplayP3RedGreenBlueAlpha demonstrates how to create a Color instance using NewColorWithDisplayP3RedGreenBlueAlpha.
// Creates a color object from the specified components in the Display P3 color space.
func ExampleNewColorWithDisplayP3RedGreenBlueAlpha() {
	_ = appkit.NewColorWithDisplayP3RedGreenBlueAlpha(
		0.0, // red float64
		0.0, // green float64
		0.0, // blue float64
		0.0, // alpha float64
	)
	// Output:
}
// ExampleNewColorWithPatternImage demonstrates how to create a Color instance using NewColorWithPatternImage.
// Creates a color object that uses the specified image pattern to paint the target area.
func ExampleNewColorWithPatternImage() {
	_ = appkit.NewColorWithPatternImage(
		appkit.Image{}, // image Image
	)
	// Output:
}
// ExampleNewColorWithRedGreenBlueAlphaExposure demonstrates how to create a Color instance using NewColorWithRedGreenBlueAlphaExposure.
// Generates an HDR color in the extended sRGB colorspace by applying an exposure to the SDR color defined by the red, green, and blue components. The  ,  , and   components have a nominal range of [0..1],   is a value >= 0. To produce an HDR color, we process the given color in a linear color space, multiplying component values by  . The produced color will have a   equal to the linearized exposure value. Each whole value of exposure produces a color that is twice as bright.
func ExampleNewColorWithRedGreenBlueAlphaExposure() {
	_ = appkit.NewColorWithRedGreenBlueAlphaExposure(
		0.0, // red float64
		0.0, // green float64
		0.0, // blue float64
		0.0, // alpha float64
		0.0, // exposure float64
	)
	// Output:
}
// ExampleNewColorWithRedGreenBlueAlphaLinearExposure demonstrates how to create a Color instance using NewColorWithRedGreenBlueAlphaLinearExposure.
// Generates an HDR color in the extended sRGB colorspace by applying an exposure to the SDR color defined by the red, green, and blue components. The  ,  , and   components have a nominal range of [0..1],   is a value >= 1. To produce an HDR color, we process the given color in a linear color space, multiplying component values by  . The produced color will have a   equal to  . Each doubling of   produces a color that is twice as bright.
func ExampleNewColorWithRedGreenBlueAlphaLinearExposure() {
	_ = appkit.NewColorWithRedGreenBlueAlphaLinearExposure(
		0.0, // red float64
		0.0, // green float64
		0.0, // blue float64
		0.0, // alpha float64
		0.0, // linearExposure float64
	)
	// Output:
}
// ExampleNewColorWithWhiteAlpha demonstrates how to create a Color instance using NewColorWithWhiteAlpha.
// Creates a color object with the specified brightness and alpha channel values.
func ExampleNewColorWithWhiteAlpha() {
	_ = appkit.NewColorWithWhiteAlpha(
		0.0, // white float64
		0.0, // alpha float64
	)
	// Output:
}
