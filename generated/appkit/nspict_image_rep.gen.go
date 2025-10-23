// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [PICTImageRep] class.
var (
	PICTImageRepClass     _PICTImageRepClass
	PICTImageRepClassOnce sync.Once
)

func getPICTImageRepClass() _PICTImageRepClass {
	PICTImageRepClassOnce.Do(func() {
		PICTImageRepClass = _PICTImageRepClass{objc.GetClass("NSPICTImageRep")}
	})
	return PICTImageRepClass
}

type _PICTImageRepClass struct {
	class objc.Class
}

// An interface definition for the [PICTImageRep] class.
type IPICTImageRep interface {
	IImageRep
	PICTRepresentation() foundation.NSData
	BoundingBox() coregraphics.CGRect
	SetBoundingBox(value coregraphics.CGRect)
}

// An object that renders an image from a PICT format data stream of version 1, version 2, and extended version 2.


// An object that renders an image from a PICT format data stream of version 1, version 2, and extended version 2.
//
// [Full Topic]
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

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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
	return getPICTImageRepClass().New()
}



// The image representation’s PICT data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPICTImageRep/pictRepresentation
func (p_ PICTImageRep) PICTRepresentation() foundation.NSData {
	rv := objc.Send[foundation.NSData](p_.ID, objc.Sel("PICTRepresentation"))
	return rv
}


// The rectangle that bounds the image representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspictimagerep/boundingbox
func (p_ PICTImageRep) BoundingBox() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](p_.ID, objc.Sel("boundingBox"))
	return rv
}


// The rectangle that bounds the image representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspictimagerep/boundingbox
func (p_ PICTImageRep) SetBoundingBox(value coregraphics.CGRect) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setBoundingBox:"), value)
}



