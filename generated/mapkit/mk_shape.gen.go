// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MKShape */


/* debug [class_header]: Header for MKShape */
// The class instance for the [MKShape] class.
var (
	MKShapeClass     _MKShapeClass
	MKShapeClassOnce sync.Once
)

func getMKShapeClass() _MKShapeClass {
	MKShapeClassOnce.Do(func() {
		MKShapeClass = _MKShapeClass{objc.GetClass("MKShape")}
	})
	return MKShapeClass
}

type _MKShapeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKShape */
// An interface definition for the [MKShape] class.
type IMKShape interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MKShape */
	// properties:
	Subtitle() objc.IObject /* cross-framework: NSString */
	SetSubtitle(value objc.IObject /* cross-framework: NSString */)
	Title() objc.IObject /* cross-framework: NSString */
	SetTitle(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKShape */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKShape */
// Alloc allocates a new instance without initialization.
func (mc _MKShapeClass) Alloc() MKShape {
	rv := objc.Send[MKShape](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MKShapeClass) New() MKShape {
	rv := objc.Send[MKShape](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKShape) Init() MKShape {
	rv := objc.Send[MKShape](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKShape) Autorelease() MKShape {
	rv := objc.Send[MKShape](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKShape creates a new MKShape instance.
func NewMKShape() MKShape {
	return getMKShapeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKShape */
// An abstract class that defines the basic properties for all shape-based overlay objects.
//
// You can’t instantiate this class directly; use a subclass instead. Subclasses are responsible for defining the geometry of the shape and providing an appropriate value for the coordinate property they inherit from the protocol.


// An abstract class that defines the basic properties for all shape-based overlay objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKShape
type MKShape struct {
	objectivec.Object
}

// MKShapeFrom constructs a [MKShape] from an unsafe.Pointer.
//
// An abstract class that defines the basic properties for all shape-based overlay objects.
func MKShapeFrom(ptr unsafe.Pointer) MKShape {
	return MKShape{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKShape *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKShape */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKShape */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKShape */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKShape */

// The subtitle of the shape annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKShape/subtitle
func (m_ MKShape) Subtitle() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("subtitle"))
	return rv
}/* debug [instance_properties/getter]: subtitle */


// The subtitle of the shape annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKShape/subtitle
func (m_ MKShape) SetSubtitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSubtitle:"), value)
}/* debug [instance_properties/setter]: subtitle */


// The title of the shape annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKShape/title
func (m_ MKShape) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("title"))
	return rv
}/* debug [instance_properties/getter]: title */


// The title of the shape annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKShape/title
func (m_ MKShape) SetTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTitle:"), value)
}/* debug [instance_properties/setter]: title */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKShape */



