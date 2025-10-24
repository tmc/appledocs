// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MKPointAnnotation */


/* debug [class_header]: Header for MKPointAnnotation */
// The class instance for the [MKPointAnnotation] class.
var (
	MKPointAnnotationClass     _MKPointAnnotationClass
	MKPointAnnotationClassOnce sync.Once
)

func getMKPointAnnotationClass() _MKPointAnnotationClass {
	MKPointAnnotationClassOnce.Do(func() {
		MKPointAnnotationClass = _MKPointAnnotationClass{objc.GetClass("MKPointAnnotation")}
	})
	return MKPointAnnotationClass
}

type _MKPointAnnotationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKPointAnnotation */
// An interface definition for the [MKPointAnnotation] class.
type IMKPointAnnotation interface {
	IMKShape
	
/* debug [class_interface_properties]: Properties for MKPointAnnotation */
	// properties:
	Coordinate() LocationCoordinate2D /* not a class type */
	SetCoordinate(value LocationCoordinate2D /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKPointAnnotation */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKPointAnnotation */
// Alloc allocates a new instance without initialization.
func (mc _MKPointAnnotationClass) Alloc() MKPointAnnotation {
	rv := objc.Send[MKPointAnnotation](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MKPointAnnotationClass) New() MKPointAnnotation {
	rv := objc.Send[MKPointAnnotation](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKPointAnnotation) Init() MKPointAnnotation {
	rv := objc.Send[MKPointAnnotation](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKPointAnnotation) Autorelease() MKPointAnnotation {
	rv := objc.Send[MKPointAnnotation](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKPointAnnotation creates a new MKPointAnnotation instance.
func NewMKPointAnnotation() MKPointAnnotation {
	return getMKPointAnnotationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKPointAnnotation */
// A string-based piece of location-specific data that you apply to a specific point on a map.
//
// You use this class, rather than define a custom annotation object, in situations where all you want to do is display a title string at the specified point on the map.


// A string-based piece of location-specific data that you apply to a specific point on a map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPointAnnotation
type MKPointAnnotation struct {
	MKShape
}

// MKPointAnnotationFrom constructs a [MKPointAnnotation] from an unsafe.Pointer.
//
// A string-based piece of location-specific data that you apply to a specific point on a map.
func MKPointAnnotationFrom(ptr unsafe.Pointer) MKPointAnnotation {
	return MKPointAnnotation{
		MKShape: MKShapeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKPointAnnotation */

// Creates a point annotation at the specified coordinate on the map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPointAnnotation/init(coordinate:)
func NewMKPointAnnotationWithCoordinate(coordinate LocationCoordinate2D /* not a class type */) MKPointAnnotation {
	instance := getMKPointAnnotationClass().Alloc()
	rv := objc.Send[MKPointAnnotation](instance.ID, objc.Sel("initWithCoordinate:"), coordinate)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMKPointAnnotationWithCoordinate */


// Creates a point annotation displaying a title and subtitle string at the specified coordinate on the map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPointAnnotation/init(coordinate:title:subtitle:)
func NewMKPointAnnotationWithCoordinateTitleSubtitle(coordinate LocationCoordinate2D /* not a class type */, title objc.IObject /* cross-framework: NSString */, subtitle objc.IObject /* cross-framework: NSString */) MKPointAnnotation {
	instance := getMKPointAnnotationClass().Alloc()
	rv := objc.Send[MKPointAnnotation](instance.ID, objc.Sel("initWithCoordinate:title:subtitle:"), coordinate, title, subtitle)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMKPointAnnotationWithCoordinateTitleSubtitle */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKPointAnnotation */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKPointAnnotation */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKPointAnnotation */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKPointAnnotation */

// The coordinate point of the annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPointAnnotation/coordinate
func (m_ MKPointAnnotation) Coordinate() LocationCoordinate2D /* not a class type */ {
	rv := objc.Send[LocationCoordinate2D](m_.ID, objc.Sel("coordinate"))
	return rv
}/* debug [instance_properties/getter]: coordinate */


// The coordinate point of the annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPointAnnotation/coordinate
func (m_ MKPointAnnotation) SetCoordinate(value LocationCoordinate2D /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCoordinate:"), value)
}/* debug [instance_properties/setter]: coordinate */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKPointAnnotation */


