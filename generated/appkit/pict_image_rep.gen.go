// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PICTImageRep] class.
var pICTImageRepClass = _PICTImageRepClass{objc.GetClass("NSPICTImageRep")}

type _PICTImageRepClass struct {
	class objc.Class
}

// An interface definition for the [PICTImageRep] class.
type IPICTImageRep interface {
	IImageRep
}

// An object that renders an image from a PICT format data stream of version 1, version 2, and extended version 2. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPICTImageRep

type PICTImageRep struct {
	ImageRep
}

// PICTImageRepFrom constructs a [PICTImageRep] from an unsafe.Pointer.
//
// An object that renders an image from a PICT format data stream of version 1, version 2, and extended version 2.
func PICTImageRepFrom(ptr unsafe.Pointer) PICTImageRep {
	return PICTImageRep{
		ImageRep: ImageRepFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (pc _PICTImageRepClass) Alloc() PICTImageRep {
	rv := objc.Send[PICTImageRep](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (pc _PICTImageRepClass) New() PICTImageRep {
	rv := objc.Send[PICTImageRep](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PICTImageRep) Init() PICTImageRep {
	rv := objc.Send[PICTImageRep](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PICTImageRep) Autorelease() PICTImageRep {
	rv := objc.Send[PICTImageRep](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPICTImageRep creates a new PICTImageRep instance.
func NewPICTImageRep() PICTImageRep {
	return pICTImageRepClass.New()
}




