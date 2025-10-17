// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Color] class.
var colorClass = _ColorClass{objc.GetClass("NSColor")}

type _ColorClass struct {
	class objc.Class
}

// An object that stores color data and sometimes opacity (alpha value). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor

type Color struct {
	objectivec.Object
}

// ColorFrom constructs a [Color] from an unsafe.Pointer.
//
// An object that stores color data and sometimes opacity (alpha value).
func ColorFrom(ptr unsafe.Pointer) Color {
	return Color{objectivec.Object{objc.ID(ptr)}}
}

// Creates a new color object whose color is the same as the receiver’s, except that the new color object is in the specified color space. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/usingColorSpaceName(_:)
func (c_ Color) ColorUsingColorSpaceName(name unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("colorUsingColorSpaceName:"), name)
	return rv
}


