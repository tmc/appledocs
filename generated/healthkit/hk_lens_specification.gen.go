// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class HKLensSpecification */


/* debug [class_header]: Header for HKLensSpecification */
// The class instance for the [HKLensSpecification] class.
var (
	HKLensSpecificationClass     _HKLensSpecificationClass
	HKLensSpecificationClassOnce sync.Once
)

func getHKLensSpecificationClass() _HKLensSpecificationClass {
	HKLensSpecificationClassOnce.Do(func() {
		HKLensSpecificationClass = _HKLensSpecificationClass{objc.GetClass("HKLensSpecification")}
	})
	return HKLensSpecificationClass
}

type _HKLensSpecificationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKLensSpecification */
// An interface definition for the [HKLensSpecification] class.
type IHKLensSpecification interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for HKLensSpecification */
	// properties:
	AddPower() IHKQuantity
	Axis() IHKQuantity
	Cylinder() IHKQuantity
	Sphere() IHKQuantity
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKLensSpecification */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKLensSpecification */
// Alloc allocates a new instance without initialization.
func (hc _HKLensSpecificationClass) Alloc() HKLensSpecification {
	rv := objc.Send[HKLensSpecification](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKLensSpecificationClass) New() HKLensSpecification {
	rv := objc.Send[HKLensSpecification](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKLensSpecification) Init() HKLensSpecification {
	rv := objc.Send[HKLensSpecification](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKLensSpecification) Autorelease() HKLensSpecification {
	rv := objc.Send[HKLensSpecification](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKLensSpecification creates a new HKLensSpecification instance.
func NewHKLensSpecification() HKLensSpecification {
	return getHKLensSpecificationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKLensSpecification */
// An abstract superclass for lens specifications.
//
// Don’t instantiate this class directly. Instead, use one of its concrete subclasses: or .


// An abstract superclass for lens specifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKLensSpecification
type HKLensSpecification struct {
	objectivec.Object
}

// HKLensSpecificationFrom constructs a [HKLensSpecification] from an unsafe.Pointer.
//
// An abstract superclass for lens specifications.
func HKLensSpecificationFrom(ptr unsafe.Pointer) HKLensSpecification {
	return HKLensSpecification{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKLensSpecification *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKLensSpecification */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKLensSpecification */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKLensSpecification */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKLensSpecification */

// The correction for nearsightedness.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKLensSpecification/addPower
func (h_ HKLensSpecification) AddPower() IHKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("addPower"))
	return rv
}/* debug [instance_properties/getter]: addPower */


// Part of the correction for astigmatism that measures the orientation fo the correction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKLensSpecification/axis
func (h_ HKLensSpecification) Axis() IHKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("axis"))
	return rv
}/* debug [instance_properties/getter]: axis */


// Part of the correction for astigmatism that measures the strength of the correction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKLensSpecification/cylinder
func (h_ HKLensSpecification) Cylinder() IHKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("cylinder"))
	return rv
}/* debug [instance_properties/getter]: cylinder */


// The correction for farsightedness.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKLensSpecification/sphere
func (h_ HKLensSpecification) Sphere() IHKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("sphere"))
	return rv
}/* debug [instance_properties/getter]: sphere */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKLensSpecification */



