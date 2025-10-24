// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewColor

// ExampleNewColorWithCGColor demonstrates how to create a Color instance using NewColorWithCGColor.
// Creates a color object using the specified Core Graphics color.
func ExampleNewColorWithCGColor() {
	_ = appkit.NewColorWithCGColor(
		appkit.ColorRef /* not a class type */{}, // cgColor ColorRef /* not a class type */
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
// ExampleColor_Set demonstrates using Set on a Color instance.
// Sets the color of subsequent drawing to the color that the color object represents.
func ExampleColor_Set() {
	obj := appkit.NewColor()
	obj.Set()
	// Output:
	}

// ExampleColor_SetFill demonstrates using SetFill on a Color instance.
// Sets the fill color of subsequent drawing to the color object’s color.
func ExampleColor_SetFill() {
	obj := appkit.NewColor()
	obj.SetFill()
	// Output:
	}

// ExampleColor_SetStroke demonstrates using SetStroke on a Color instance.
// Sets the stroke color of subsequent drawing to the color object’s color.
func ExampleColor_SetStroke() {
	obj := appkit.NewColor()
	obj.SetStroke()
	// Output:
	}

