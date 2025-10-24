// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MKMapFeatureAnnotation */


/* debug [class_header]: Header for MKMapFeatureAnnotation */
// The class instance for the [MKMapFeatureAnnotation] class.
var (
	MKMapFeatureAnnotationClass     _MKMapFeatureAnnotationClass
	MKMapFeatureAnnotationClassOnce sync.Once
)

func getMKMapFeatureAnnotationClass() _MKMapFeatureAnnotationClass {
	MKMapFeatureAnnotationClassOnce.Do(func() {
		MKMapFeatureAnnotationClass = _MKMapFeatureAnnotationClass{objc.GetClass("MKMapFeatureAnnotation")}
	})
	return MKMapFeatureAnnotationClass
}

type _MKMapFeatureAnnotationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKMapFeatureAnnotation */
// An interface definition for the [MKMapFeatureAnnotation] class.
type IMKMapFeatureAnnotation interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MKMapFeatureAnnotation */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKMapFeatureAnnotation */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKMapFeatureAnnotation */
// Alloc allocates a new instance without initialization.
func (mc _MKMapFeatureAnnotationClass) Alloc() MKMapFeatureAnnotation {
	rv := objc.Send[MKMapFeatureAnnotation](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MKMapFeatureAnnotationClass) New() MKMapFeatureAnnotation {
	rv := objc.Send[MKMapFeatureAnnotation](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKMapFeatureAnnotation) Init() MKMapFeatureAnnotation {
	rv := objc.Send[MKMapFeatureAnnotation](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKMapFeatureAnnotation) Autorelease() MKMapFeatureAnnotation {
	rv := objc.Send[MKMapFeatureAnnotation](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKMapFeatureAnnotation creates a new MKMapFeatureAnnotation instance.
func NewMKMapFeatureAnnotation() MKMapFeatureAnnotation {
	return getMKMapFeatureAnnotationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKMapFeatureAnnotation */
// A class that describes an annotation element on the map’s display such as a point of interest, territorial boundary, or physical feature.


// A class that describes an annotation element on the map’s display such as a point of interest, territorial boundary, or physical feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapFeatureAnnotation
type MKMapFeatureAnnotation struct {
	objectivec.Object
}

// MKMapFeatureAnnotationFrom constructs a [MKMapFeatureAnnotation] from an unsafe.Pointer.
//
// A class that describes an annotation element on the map’s display such as a point of interest, territorial boundary, or physical feature.
func MKMapFeatureAnnotationFrom(ptr unsafe.Pointer) MKMapFeatureAnnotation {
	return MKMapFeatureAnnotation{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKMapFeatureAnnotation *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKMapFeatureAnnotation */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKMapFeatureAnnotation */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKMapFeatureAnnotation */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKMapFeatureAnnotation */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKMapFeatureAnnotation */


