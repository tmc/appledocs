// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MKMapItemAnnotation */


/* debug [class_header]: Header for MKMapItemAnnotation */
// The class instance for the [MKMapItemAnnotation] class.
var (
	MKMapItemAnnotationClass     _MKMapItemAnnotationClass
	MKMapItemAnnotationClassOnce sync.Once
)

func getMKMapItemAnnotationClass() _MKMapItemAnnotationClass {
	MKMapItemAnnotationClassOnce.Do(func() {
		MKMapItemAnnotationClass = _MKMapItemAnnotationClass{objc.GetClass("MKMapItemAnnotation")}
	})
	return MKMapItemAnnotationClass
}

type _MKMapItemAnnotationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKMapItemAnnotation */
// An interface definition for the [MKMapItemAnnotation] class.
type IMKMapItemAnnotation interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MKMapItemAnnotation */
	// properties:
	MapItem() IMKMapItem
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKMapItemAnnotation */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKMapItemAnnotation */
// Alloc allocates a new instance without initialization.
func (mc _MKMapItemAnnotationClass) Alloc() MKMapItemAnnotation {
	rv := objc.Send[MKMapItemAnnotation](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MKMapItemAnnotationClass) New() MKMapItemAnnotation {
	rv := objc.Send[MKMapItemAnnotation](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKMapItemAnnotation) Init() MKMapItemAnnotation {
	rv := objc.Send[MKMapItemAnnotation](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKMapItemAnnotation) Autorelease() MKMapItemAnnotation {
	rv := objc.Send[MKMapItemAnnotation](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKMapItemAnnotation creates a new MKMapItemAnnotation instance.
func NewMKMapItemAnnotation() MKMapItemAnnotation {
	return getMKMapItemAnnotationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKMapItemAnnotation */
// An annotation that represents a map item


// An annotation that represents a map item
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItemAnnotation
type MKMapItemAnnotation struct {
	objectivec.Object
}

// MKMapItemAnnotationFrom constructs a [MKMapItemAnnotation] from an unsafe.Pointer.
//
// An annotation that represents a map item
func MKMapItemAnnotationFrom(ptr unsafe.Pointer) MKMapItemAnnotation {
	return MKMapItemAnnotation{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKMapItemAnnotation */

// Creates a map item annotation
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItemAnnotation/init(mapItem:)
func NewMKMapItemAnnotationWithMapItem(mapItem IMKMapItem) MKMapItemAnnotation {
	instance := getMKMapItemAnnotationClass().Alloc()
	rv := objc.Send[MKMapItemAnnotation](instance.ID, objc.Sel("initWithMapItem:"), mapItem)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMKMapItemAnnotationWithMapItem */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKMapItemAnnotation */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKMapItemAnnotation */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKMapItemAnnotation */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKMapItemAnnotation */

// The map item represented by this annotation
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItemAnnotation/mapItem
func (m_ MKMapItemAnnotation) MapItem() IMKMapItem {
	rv := objc.Send[MKMapItem](m_.ID, objc.Sel("mapItem"))
	return rv
}/* debug [instance_properties/getter]: mapItem */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKMapItemAnnotation */


