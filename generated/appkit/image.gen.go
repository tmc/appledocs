// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Image] class.
var imageClass = _ImageClass{objc.GetClass("NSImage")}

type _ImageClass struct {
	class objc.Class
}

// An interface definition for the [Image] class.
type IImage interface {
	objectivec.IObject
	BestRepresentationForDevice(deviceDescription unsafe.Pointer) unsafe.Pointer
}

// A high-level interface for manipulating image data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage

type Image struct {
	objectivec.Object
}

// ImageFrom constructs a [Image] from an unsafe.Pointer.
//
// A high-level interface for manipulating image data.
func ImageFrom(ptr unsafe.Pointer) Image {
	return Image{objectivec.Object{objc.ID(ptr)}}
}

// Returns the best representation for the device with the specified characteristics. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/bestRepresentationForDevice:
func (i_ Image) BestRepresentationForDevice(deviceDescription unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("bestRepresentationForDevice:"), deviceDescription)
	return rv
}


