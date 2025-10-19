// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ImageRep] class.
var (
	imageRepClass     _ImageRepClass
	imageRepClassOnce sync.Once
)

func getImageRepClass() _ImageRepClass {
	imageRepClassOnce.Do(func() {
		imageRepClass = _ImageRepClass{objc.GetClass("NSImageRep")}
	})
	return imageRepClass
}

type _ImageRepClass struct {
	class objc.Class
}

// An interface definition for the [ImageRep] class.
type IImageRep interface {
	objectivec.IObject
}

// A semiabstract superclass that provides subclasses that you use to draw an image from a particular type of source data.
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

// Alloc allocates a new instance without initialization.
func (ic _ImageRepClass) Alloc() ImageRep {
	rv := objc.Send[ImageRep](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _ImageRepClass) New() ImageRep {
	rv := objc.Send[ImageRep](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageRep) Init() ImageRep {
	rv := objc.Send[ImageRep](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageRep) Autorelease() ImageRep {
	rv := objc.Send[ImageRep](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageRep creates a new ImageRep instance.
func NewImageRep() ImageRep {
	return getImageRepClass().New()
}




