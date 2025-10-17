// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Color] class.
var ColorClass objc.Class

func init() {
	ColorClass = objc.GetClass("NSColor")
}

type Color struct {
	objc.ID
}

func ColorFrom(ptr unsafe.Pointer) Color {
	return Color{
		ID: objc.ID(ptr),
	}
}


// Creates a new color object whose color is the same as the receiver’s, except that the new color object is in the specified color space. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSColor/usingColorSpaceName(_:)
func (c_ Color) ColorUsingColorSpaceName(name unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("colorUsingColorSpaceName:")
	ret := c_.ID.Send(sel, name)
	return unsafe.Pointer(ret)
}

