// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSCIImageRep */


/* debug [class_header]: Header for NSCIImageRep */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CIImageRep */
// An interface definition for the [CIImageRep] class.
type ICIImageRep interface {
	IImageRep
	
/* debug [class_interface_properties]: Properties for CIImageRep */
	// properties:
	CIImage() IImage
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CIImageRep */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CIImageRep */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CIImageRep */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CIImageRep */

// Returns a representation of an image initialized to the specified Core Image instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCIImageRep/init(ciImage:)
func NewCIImageRepWithCIImage(image IImage) CIImageRep {
	instance := getCIImageRepClass().Alloc()
	rv := objc.Send[CIImageRep](instance.ID, objc.Sel("initWithCIImage:"), image)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCIImageRepWithCIImage */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CIImageRep */

// Creates and returns a representation of an image initialized to the specified Core Image instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCIImageRep/imageRepWithCIImage:
func (ic _CIImageRepClass) ImageRepWithCIImage(image IImage) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ic.class), objc.Sel("imageRepWithCIImage:"), image)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ImageRepWithCIImage) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CIImageRep */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CIImageRep */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CIImageRep */

// The Core Image instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCIImageRep/ciImage
func (i_ CIImageRep) CIImage() IImage {
	rv := objc.Send[Image](i_.ID, objc.Sel("CIImage"))
	return rv
}/* debug [instance_properties/getter]: CIImage */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSCIImageRep */


