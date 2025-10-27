// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
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
	

	// properties:
	BoundingBox() corefoundation.CGRect
	PICTRepresentation() foundation.foundation.INSData


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (pc _PICTImageRepClass) Alloc() PICTImageRep {
	rv := objc.Send[PICTImageRep](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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






// Returns a representation of an image from the specified data in the PICT file format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPICTImageRep/init(data:)
func NewPICTImageRepWithData(pictData foundation.foundation.INSData) PICTImageRep {
	instance := getPICTImageRepClass().Alloc()
	rv := objc.Send[PICTImageRep](instance.ID, objc.Sel("initWithData:"), pictData)
	rv.Autorelease()
	return rv
}







// Creates and returns a representation of an image from the specified data in the PICT file format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPICTImageRep/imageRepWithData:
func (pc _PICTImageRepClass) ImageRepWithData(pictData foundation.foundation.INSData) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(pc.class), objc.Sel("imageRepWithData:"), pictData)
	return rv
}

















// The rectangle that bounds the image representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPICTImageRep/boundingBox
func (p_ PICTImageRep) BoundingBox() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](p_.ID, objc.Sel("boundingBox"))
	return rv
}


// The image representation’s PICT data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPICTImageRep/pictRepresentation
func (p_ PICTImageRep) PICTRepresentation() foundation.foundation.INSData {
	rv := objc.Send[foundation.NSData](p_.ID, objc.Sel("PICTRepresentation"))
	return rv
}







