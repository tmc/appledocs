// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ImageRep] class.
var imageRepClass = _ImageRepClass{objc.GetClass("NSImageRep")}

type _ImageRepClass struct {
	class objc.Class
}

// An interface definition for the [ImageRep] class.
type IImageRep interface {
	objectivec.IObject
}

// A semiabstract superclass that provides subclasses that you use to draw an image from a particular type of source data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageRep

type ImageRep struct {
	objectivec.Object
}

// ImageRepFrom constructs a [ImageRep] from an unsafe.Pointer.
//
// A semiabstract superclass that provides subclasses that you use to draw an image from a particular type of source data.
func ImageRepFrom(ptr unsafe.Pointer) ImageRep {
	return ImageRep{objectivec.Object{objc.ID(ptr)}}
}



