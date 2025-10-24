// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSEPSImageRep */


/* debug [class_header]: Header for NSEPSImageRep */
// The class instance for the [EPSImageRep] class.
var (
	EPSImageRepClass     _EPSImageRepClass
	EPSImageRepClassOnce sync.Once
)

func getEPSImageRepClass() _EPSImageRepClass {
	EPSImageRepClassOnce.Do(func() {
		EPSImageRepClass = _EPSImageRepClass{objc.GetClass("NSEPSImageRep")}
	})
	return EPSImageRepClass
}

type _EPSImageRepClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for EPSImageRep */
// An interface definition for the [EPSImageRep] class.
type IEPSImageRep interface {
	IImageRep
	
/* debug [class_interface_properties]: Properties for EPSImageRep */
	// properties:
	BoundingBox() Rect /* not a class type */
	EPSRepresentation() objc.IObject /* cross-framework: NSData */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for EPSImageRep */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for EPSImageRep */
// Alloc allocates a new instance without initialization.
func (ec _EPSImageRepClass) Alloc() EPSImageRep {
	rv := objc.Send[EPSImageRep](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ec _EPSImageRepClass) New() EPSImageRep {
	rv := objc.Send[EPSImageRep](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ EPSImageRep) Init() EPSImageRep {
	rv := objc.Send[EPSImageRep](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ EPSImageRep) Autorelease() EPSImageRep {
	rv := objc.Send[EPSImageRep](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewEPSImageRep creates a new EPSImageRep instance.
func NewEPSImageRep() EPSImageRep {
	return getEPSImageRepClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for EPSImageRep */
// An object that can render an image from encapsulated PostScript (EPS) code.


// An object that can render an image from encapsulated PostScript (EPS) code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEPSImageRep
type EPSImageRep struct {
	ImageRep
}

// EPSImageRepFrom constructs a [EPSImageRep] from an unsafe.Pointer.
//
// An object that can render an image from encapsulated PostScript (EPS) code.
func EPSImageRepFrom(ptr unsafe.Pointer) EPSImageRep {
	return EPSImageRep{
		ImageRep: ImageRepFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for EPSImageRep */

// Returns a representation of an image initialized with the specified EPS data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEPSImageRep/init(data:)
func NewEPSImageRepWithData(epsData objc.IObject /* cross-framework: NSData */) EPSImageRep {
	instance := getEPSImageRepClass().Alloc()
	rv := objc.Send[EPSImageRep](instance.ID, objc.Sel("initWithData:"), epsData)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewEPSImageRepWithData */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for EPSImageRep */

// Creates and returns a representation of an image initialized with the specified EPS data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEPSImageRep/imageRepWithData:
func (ec _EPSImageRepClass) ImageRepWithData(epsData objc.IObject /* cross-framework: NSData */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ec.class), objc.Sel("imageRepWithData:"), epsData)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ImageRepWithData) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for EPSImageRep */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for EPSImageRep */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for EPSImageRep */

// The rectangle that bounds the image representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEPSImageRep/boundingBox
func (e_ EPSImageRep) BoundingBox() Rect /* not a class type */ {
	rv := objc.Send[Rect](e_.ID, objc.Sel("boundingBox"))
	return rv
}/* debug [instance_properties/getter]: boundingBox */


// The EPS representation of the image representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEPSImageRep/epsRepresentation
func (e_ EPSImageRep) EPSRepresentation() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](e_.ID, objc.Sel("EPSRepresentation"))
	return rv
}/* debug [instance_properties/getter]: EPSRepresentation */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSEPSImageRep */


