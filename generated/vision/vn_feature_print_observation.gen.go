// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VNFeaturePrintObservation */


/* debug [class_header]: Header for VNFeaturePrintObservation */
// The class instance for the [FeaturePrintObservation] class.
var (
	FeaturePrintObservationClass     _FeaturePrintObservationClass
	FeaturePrintObservationClassOnce sync.Once
)

func getFeaturePrintObservationClass() _FeaturePrintObservationClass {
	FeaturePrintObservationClassOnce.Do(func() {
		FeaturePrintObservationClass = _FeaturePrintObservationClass{objc.GetClass("VNFeaturePrintObservation")}
	})
	return FeaturePrintObservationClass
}

type _FeaturePrintObservationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for FeaturePrintObservation */
// An interface definition for the [FeaturePrintObservation] class.
type IFeaturePrintObservation interface {
	IObservation
	
/* debug [class_interface_properties]: Properties for FeaturePrintObservation */
	// properties:
	Data() objc.IObject /* cross-framework: NSData */
	ElementCount() uint
	ElementType() ElementType
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for FeaturePrintObservation */
	// methods:
	ComputeDistanceToFeaturePrintObservationError(outDistance objectivec.IObject, featurePrint IVNFeaturePrintObservation, error_ objectivec.IObject) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for FeaturePrintObservation */
// Alloc allocates a new instance without initialization.
func (fc _FeaturePrintObservationClass) Alloc() FeaturePrintObservation {
	rv := objc.Send[FeaturePrintObservation](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FeaturePrintObservationClass) New() FeaturePrintObservation {
	rv := objc.Send[FeaturePrintObservation](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FeaturePrintObservation) Init() FeaturePrintObservation {
	rv := objc.Send[FeaturePrintObservation](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FeaturePrintObservation) Autorelease() FeaturePrintObservation {
	rv := objc.Send[FeaturePrintObservation](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFeaturePrintObservation creates a new FeaturePrintObservation instance.
func NewFeaturePrintObservation() FeaturePrintObservation {
	return getFeaturePrintObservationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for FeaturePrintObservation */
// An observation that provides the recognized feature print.


// An observation that provides the recognized feature print.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNFeaturePrintObservation
type FeaturePrintObservation struct {
	Observation
}

// FeaturePrintObservationFrom constructs a [FeaturePrintObservation] from an unsafe.Pointer.
//
// An observation that provides the recognized feature print.
func FeaturePrintObservationFrom(ptr unsafe.Pointer) FeaturePrintObservation {
	return FeaturePrintObservation{
		Observation: ObservationFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for FeaturePrintObservation *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for FeaturePrintObservation */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for FeaturePrintObservation */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for FeaturePrintObservation */

// Computes the distance between two feature print observations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNFeaturePrintObservation/computeDistance(_:to:)
func (f_ FeaturePrintObservation) ComputeDistanceToFeaturePrintObservationError(outDistance objectivec.IObject, featurePrint IVNFeaturePrintObservation, error_ objectivec.IObject) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("computeDistance:toFeaturePrintObservation:error:"), outDistance, featurePrint, error_)
	return rv
}/* debug [instance_methods/method]: ComputeDistanceToFeaturePrintObservationError */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for FeaturePrintObservation */

// The feature print data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNFeaturePrintObservation/data
func (f_ FeaturePrintObservation) Data() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](f_.ID, objc.Sel("data"))
	return rv
}/* debug [instance_properties/getter]: data */


// The total number of elements in the data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNFeaturePrintObservation/elementCount
func (f_ FeaturePrintObservation) ElementCount() uint {
	rv := objc.Send[uint](f_.ID, objc.Sel("elementCount"))
	return rv
}/* debug [instance_properties/getter]: elementCount */


// The type of each element in the data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNFeaturePrintObservation/elementType
func (f_ FeaturePrintObservation) ElementType() ElementType {
	rv := objc.Send[ElementType](f_.ID, objc.Sel("elementType"))
	return rv
}/* debug [instance_properties/getter]: elementType */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VNFeaturePrintObservation */



