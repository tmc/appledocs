// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CIImageRep] class.
var (
	cIImageRepClass     _CIImageRepClass
	cIImageRepClassOnce sync.Once
)

func getCIImageRepClass() _CIImageRepClass {
	cIImageRepClassOnce.Do(func() {
		cIImageRepClass = _CIImageRepClass{objc.GetClass("NSCIImageRep")}
	})
	return cIImageRepClass
}

type _CIImageRepClass struct {
	class objc.Class
}

// An interface definition for the [CIImageRep] class.
type ICIImageRep interface {
	IImageRep
}

// An object that can render an image from a Core Image object.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCIImageRep
type CIImageRep struct {
	ImageRep
}

// CIImageRepFrom constructs a [CIImageRep] from an unsafe.Pointer.
//
// An object that can render an image from a Core Image object.
func CIImageRepFrom(ptr unsafe.Pointer) CIImageRep {
	return CIImageRep{
		ImageRep: ImageRepFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _CIImageRepClass) Alloc() CIImageRep {
	rv := objc.Send[CIImageRep](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _CIImageRepClass) New() CIImageRep {
	rv := objc.Send[CIImageRep](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ CIImageRep) Init() CIImageRep {
	rv := objc.Send[CIImageRep](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ CIImageRep) Autorelease() CIImageRep {
	rv := objc.Send[CIImageRep](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCIImageRep creates a new CIImageRep instance.
func NewCIImageRep() CIImageRep {
	return getCIImageRepClass().New()
}




