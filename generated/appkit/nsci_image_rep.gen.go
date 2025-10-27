// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CIImageRep] class.
var (
	CIImageRepClass     _CIImageRepClass
	CIImageRepClassOnce sync.Once
)

func getCIImageRepClass() _CIImageRepClass {
	CIImageRepClassOnce.Do(func() {
		CIImageRepClass = _CIImageRepClass{objc.GetClass("NSCIImageRep")}
	})
	return CIImageRepClass
}

type _CIImageRepClass struct {
	class objc.Class
}





// An interface definition for the [CIImageRep] class.
type ICIImageRep interface {
	IImageRep
	

	// properties:
	CIImage() IImage


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (ic _CIImageRepClass) Alloc() CIImageRep {
	rv := objc.Send[CIImageRep](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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





// An object that can render an image from a Core Image object.


// An object that can render an image from a Core Image object.
//
// [Full Topic]
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






// Returns a representation of an image initialized to the specified Core Image instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCIImageRep/init(ciImage:)
func NewCIImageRepWithCIImage(image IImage) CIImageRep {
	instance := getCIImageRepClass().Alloc()
	rv := objc.Send[CIImageRep](instance.ID, objc.Sel("initWithCIImage:"), image)
	rv.Autorelease()
	return rv
}







// Creates and returns a representation of an image initialized to the specified Core Image instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCIImageRep/imageRepWithCIImage:
func (ic _CIImageRepClass) ImageRepWithCIImage(image IImage) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ic.class), objc.Sel("imageRepWithCIImage:"), image)
	return rv
}

















// The Core Image instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCIImageRep/ciImage
func (i_ CIImageRep) CIImage() IImage {
	rv := objc.Send[Image](i_.ID, objc.Sel("CIImage"))
	return rv
}







