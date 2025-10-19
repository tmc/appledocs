// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CustomImageRep] class.
var customImageRepClass = _CustomImageRepClass{objc.GetClass("NSCustomImageRep")}

type _CustomImageRepClass struct {
	class objc.Class
}

// An interface definition for the [CustomImageRep] class.
type ICustomImageRep interface {
	IImageRep
}

// An object that uses a delegate object to render an image from a custom format. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCustomImageRep

type CustomImageRep struct {
	ImageRep
}

// CustomImageRepFrom constructs a [CustomImageRep] from an unsafe.Pointer.
//
// An object that uses a delegate object to render an image from a custom format.
func CustomImageRepFrom(ptr unsafe.Pointer) CustomImageRep {
	return CustomImageRep{
		ImageRep: ImageRepFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (cc _CustomImageRepClass) Alloc() CustomImageRep {
	rv := objc.Send[CustomImageRep](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (cc _CustomImageRepClass) New() CustomImageRep {
	rv := objc.Send[CustomImageRep](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CustomImageRep) Init() CustomImageRep {
	rv := objc.Send[CustomImageRep](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CustomImageRep) Autorelease() CustomImageRep {
	rv := objc.Send[CustomImageRep](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCustomImageRep creates a new CustomImageRep instance.
func NewCustomImageRep() CustomImageRep {
	return customImageRepClass.New()
}




