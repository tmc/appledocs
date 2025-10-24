// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSPICTImageRep */


/* debug [class_header]: Header for NSPICTImageRep */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PICTImageRep */
// An interface definition for the [PICTImageRep] class.
type IPICTImageRep interface {
	IImageRep
	
/* debug [class_interface_properties]: Properties for PICTImageRep */
	// properties:
	BoundingBox() Rect /* not a class type */
	PICTRepresentation() objc.IObject /* cross-framework: NSData */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PICTImageRep */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PICTImageRep */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PICTImageRep */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PICTImageRep */

// Returns a representation of an image from the specified data in the PICT file format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPICTImageRep/init(data:)
func NewPICTImageRepWithData(pictData objc.IObject /* cross-framework: NSData */) PICTImageRep {
	instance := getPICTImageRepClass().Alloc()
	rv := objc.Send[PICTImageRep](instance.ID, objc.Sel("initWithData:"), pictData)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPICTImageRepWithData */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PICTImageRep */

// Creates and returns a representation of an image from the specified data in the PICT file format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPICTImageRep/imageRepWithData:
func (pc _PICTImageRepClass) ImageRepWithData(pictData objc.IObject /* cross-framework: NSData */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(pc.class), objc.Sel("imageRepWithData:"), pictData)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ImageRepWithData) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PICTImageRep */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PICTImageRep */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PICTImageRep */

// The rectangle that bounds the image representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPICTImageRep/boundingBox
func (p_ PICTImageRep) BoundingBox() Rect /* not a class type */ {
	rv := objc.Send[Rect](p_.ID, objc.Sel("boundingBox"))
	return rv
}/* debug [instance_properties/getter]: boundingBox */


// The image representation’s PICT data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPICTImageRep/pictRepresentation
func (p_ PICTImageRep) PICTRepresentation() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](p_.ID, objc.Sel("PICTRepresentation"))
	return rv
}/* debug [instance_properties/getter]: PICTRepresentation */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSPICTImageRep */


