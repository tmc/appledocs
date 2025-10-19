// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [BitmapImageRep] class.
var bitmapImageRepClass = _BitmapImageRepClass{objc.GetClass("NSBitmapImageRep")}

type _BitmapImageRepClass struct {
	class objc.Class
}

// An interface definition for the [BitmapImageRep] class.
type IBitmapImageRep interface {
	IImageRep
}

// An object that renders an image from bitmap data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep

type BitmapImageRep struct {
	ImageRep
}

// BitmapImageRepFrom constructs a [BitmapImageRep] from an unsafe.Pointer.
//
// An object that renders an image from bitmap data.
func BitmapImageRepFrom(ptr unsafe.Pointer) BitmapImageRep {
	return BitmapImageRep{
		ImageRep: ImageRepFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (bc _BitmapImageRepClass) Alloc() BitmapImageRep {
	rv := objc.Send[BitmapImageRep](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (bc _BitmapImageRepClass) New() BitmapImageRep {
	rv := objc.Send[BitmapImageRep](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BitmapImageRep) Init() BitmapImageRep {
	rv := objc.Send[BitmapImageRep](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BitmapImageRep) Autorelease() BitmapImageRep {
	rv := objc.Send[BitmapImageRep](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBitmapImageRep creates a new BitmapImageRep instance.
func NewBitmapImageRep() BitmapImageRep {
	return bitmapImageRepClass.New()
}




