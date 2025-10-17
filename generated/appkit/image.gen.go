// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Image] class.
var ImageClass objc.Class

func init() {
	ImageClass = objc.GetClass("NSImage")
}

type Image struct {
	objc.ID
}

func ImageFrom(ptr unsafe.Pointer) Image {
	return Image{
		ID: objc.ID(ptr),
	}
}


// Returns the best representation for the device with the specified characteristics. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSImage/bestRepresentationForDevice:
func (i_ Image) BestRepresentationForDevice(deviceDescription unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("bestRepresentationForDevice:")
	ret := i_.ID.Send(sel, deviceDescription)
	return unsafe.Pointer(ret)
}

