// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class VNHumanObservation */


/* debug [class_header]: Header for VNHumanObservation */
// The class instance for the [HumanObservation] class.
var (
	HumanObservationClass     _HumanObservationClass
	HumanObservationClassOnce sync.Once
)

func getHumanObservationClass() _HumanObservationClass {
	HumanObservationClassOnce.Do(func() {
		HumanObservationClass = _HumanObservationClass{objc.GetClass("VNHumanObservation")}
	})
	return HumanObservationClass
}

type _HumanObservationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HumanObservation */
// An interface definition for the [HumanObservation] class.
type IHumanObservation interface {
	IDetectedObjectObservation
	
/* debug [class_interface_properties]: Properties for HumanObservation */
	// properties:
	UpperBodyOnly() bool
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HumanObservation */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HumanObservation */
// Alloc allocates a new instance without initialization.
func (hc _HumanObservationClass) Alloc() HumanObservation {
	rv := objc.Send[HumanObservation](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HumanObservationClass) New() HumanObservation {
	rv := objc.Send[HumanObservation](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HumanObservation) Init() HumanObservation {
	rv := objc.Send[HumanObservation](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HumanObservation) Autorelease() HumanObservation {
	rv := objc.Send[HumanObservation](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHumanObservation creates a new HumanObservation instance.
func NewHumanObservation() HumanObservation {
	return getHumanObservationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HumanObservation */
// An object that represents a person that the request detects.


// An object that represents a person that the request detects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNHumanObservation
type HumanObservation struct {
	DetectedObjectObservation
}

// HumanObservationFrom constructs a [HumanObservation] from an unsafe.Pointer.
//
// An object that represents a person that the request detects.
func HumanObservationFrom(ptr unsafe.Pointer) HumanObservation {
	return HumanObservation{
		DetectedObjectObservation: DetectedObjectObservationFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HumanObservation *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HumanObservation */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HumanObservation */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HumanObservation */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HumanObservation */

// A Boolean value that indicates whether the observation represents an upper-body or full-body rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNHumanObservation/upperBodyOnly
func (h_ HumanObservation) UpperBodyOnly() bool {
	rv := objc.Send[bool](h_.ID, objc.Sel("upperBodyOnly"))
	return rv
}/* debug [instance_properties/getter]: upperBodyOnly */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VNHumanObservation */



