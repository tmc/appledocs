// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class HKSampleType */


/* debug [class_header]: Header for HKSampleType */
// The class instance for the [HKSampleType] class.
var (
	HKSampleTypeClass     _HKSampleTypeClass
	HKSampleTypeClassOnce sync.Once
)

func getHKSampleTypeClass() _HKSampleTypeClass {
	HKSampleTypeClassOnce.Do(func() {
		HKSampleTypeClass = _HKSampleTypeClass{objc.GetClass("HKSampleType")}
	})
	return HKSampleTypeClass
}

type _HKSampleTypeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKSampleType */
// An interface definition for the [HKSampleType] class.
type IHKSampleType interface {
	IHKObjectType
	
/* debug [class_interface_properties]: Properties for HKSampleType */
	// properties:
	AllowsRecalibrationForEstimates() bool
	IsMaximumDurationRestricted() bool
	IsMinimumDurationRestricted() bool
	MaximumAllowedDuration() float64
	MinimumAllowedDuration() float64
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKSampleType */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKSampleType */
// Alloc allocates a new instance without initialization.
func (hc _HKSampleTypeClass) Alloc() HKSampleType {
	rv := objc.Send[HKSampleType](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKSampleTypeClass) New() HKSampleType {
	rv := objc.Send[HKSampleType](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKSampleType) Init() HKSampleType {
	rv := objc.Send[HKSampleType](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKSampleType) Autorelease() HKSampleType {
	rv := objc.Send[HKSampleType](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKSampleType creates a new HKSampleType instance.
func NewHKSampleType() HKSampleType {
	return getHKSampleTypeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKSampleType */
// An abstract superclass for all classes that identify a specific type of sample when working with the HealthKit store.
//
// The class is an abstract subclass of the class, used to represent data samples. Never instantiate an object directly. Instead, work with one of its concrete subclasses: , , , or classes.


// An abstract superclass for all classes that identify a specific type of sample when working with the HealthKit store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKSampleType
type HKSampleType struct {
	HKObjectType
}

// HKSampleTypeFrom constructs a [HKSampleType] from an unsafe.Pointer.
//
// An abstract superclass for all classes that identify a specific type of sample when working with the HealthKit store.
func HKSampleTypeFrom(ptr unsafe.Pointer) HKSampleType {
	return HKSampleType{
		HKObjectType: HKObjectTypeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKSampleType *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKSampleType */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKSampleType */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKSampleType */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKSampleType */

// A Boolean value that indicates whether HealthKit supports recalibrating the prediction algorithm used to produce estimates for this sample type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKSampleType/allowsRecalibrationForEstimates
func (h_ HKSampleType) AllowsRecalibrationForEstimates() bool {
	rv := objc.Send[bool](h_.ID, objc.Sel("allowsRecalibrationForEstimates"))
	return rv
}/* debug [instance_properties/getter]: allowsRecalibrationForEstimates */


// A Boolean value that indicates whether samples of this type have a maximum time interval between the start and end dates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKSampleType/isMaximumDurationRestricted
func (h_ HKSampleType) IsMaximumDurationRestricted() bool {
	rv := objc.Send[bool](h_.ID, objc.Sel("isMaximumDurationRestricted"))
	return rv
}/* debug [instance_properties/getter]: isMaximumDurationRestricted */


// A Boolean value that indicates whether samples of this type have a minimum time interval between the start and end dates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKSampleType/isMinimumDurationRestricted
func (h_ HKSampleType) IsMinimumDurationRestricted() bool {
	rv := objc.Send[bool](h_.ID, objc.Sel("isMinimumDurationRestricted"))
	return rv
}/* debug [instance_properties/getter]: isMinimumDurationRestricted */


// The maximum duration if the sample type has a restricted duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKSampleType/maximumAllowedDuration
func (h_ HKSampleType) MaximumAllowedDuration() float64 {
	rv := objc.Send[float64](h_.ID, objc.Sel("maximumAllowedDuration"))
	return rv
}/* debug [instance_properties/getter]: maximumAllowedDuration */


// The minimum duration if the sample type has a restricted duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKSampleType/minimumAllowedDuration
func (h_ HKSampleType) MinimumAllowedDuration() float64 {
	rv := objc.Send[float64](h_.ID, objc.Sel("minimumAllowedDuration"))
	return rv
}/* debug [instance_properties/getter]: minimumAllowedDuration */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKSampleType */



