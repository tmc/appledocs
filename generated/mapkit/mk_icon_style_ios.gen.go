//go:build darwin && ios

// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for MKIconStyle


// iOS-only properties

// The background color of the icon.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKIconStyle/backgroundColor
func (m_ MKIconStyle) BackgroundColor() appkit.Color {
	rv := objc.Send[appkit.Color](m_.ID, objc.Sel("backgroundColor"))
	return rv
}

// The icon image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKIconStyle/image
func (m_ MKIconStyle) Image() appkit.Image {
	rv := objc.Send[appkit.Image](m_.ID, objc.Sel("image"))
	return rv
}





