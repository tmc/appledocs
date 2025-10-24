// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class HKAudiogramSensitivityPointClampingRange */


/* debug [class_header]: Header for HKAudiogramSensitivityPointClampingRange */
// The class instance for the [HKAudiogramSensitivityPointClampingRange] class.
var (
	HKAudiogramSensitivityPointClampingRangeClass     _HKAudiogramSensitivityPointClampingRangeClass
	HKAudiogramSensitivityPointClampingRangeClassOnce sync.Once
)

func getHKAudiogramSensitivityPointClampingRangeClass() _HKAudiogramSensitivityPointClampingRangeClass {
	HKAudiogramSensitivityPointClampingRangeClassOnce.Do(func() {
		HKAudiogramSensitivityPointClampingRangeClass = _HKAudiogramSensitivityPointClampingRangeClass{objc.GetClass("HKAudiogramSensitivityPointClampingRange")}
	})
	return HKAudiogramSensitivityPointClampingRangeClass
}

type _HKAudiogramSensitivityPointClampingRangeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKAudiogramSensitivityPointClampingRange */
// An interface definition for the [HKAudiogramSensitivityPointClampingRange] class.
type IHKAudiogramSensitivityPointClampingRange interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for HKAudiogramSensitivityPointClampingRange */
	// properties:
	LowerBound() IHKQuantity
	UpperBound() IHKQuantity
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKAudiogramSensitivityPointClampingRange */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKAudiogramSensitivityPointClampingRange */
// Alloc allocates a new instance without initialization.
func (hc _HKAudiogramSensitivityPointClampingRangeClass) Alloc() HKAudiogramSensitivityPointClampingRange {
	rv := objc.Send[HKAudiogramSensitivityPointClampingRange](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKAudiogramSensitivityPointClampingRangeClass) New() HKAudiogramSensitivityPointClampingRange {
	rv := objc.Send[HKAudiogramSensitivityPointClampingRange](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKAudiogramSensitivityPointClampingRange) Init() HKAudiogramSensitivityPointClampingRange {
	rv := objc.Send[HKAudiogramSensitivityPointClampingRange](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKAudiogramSensitivityPointClampingRange) Autorelease() HKAudiogramSensitivityPointClampingRange {
	rv := objc.Send[HKAudiogramSensitivityPointClampingRange](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKAudiogramSensitivityPointClampingRange creates a new HKAudiogramSensitivityPointClampingRange instance.
func NewHKAudiogramSensitivityPointClampingRange() HKAudiogramSensitivityPointClampingRange {
	return getHKAudiogramSensitivityPointClampingRangeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKAudiogramSensitivityPointClampingRange */
// Defines the range within which an ear’s sensitivity point may have been clamped, if any.
//
// At times, it may be required to indicate that a sensitivity point has been clamped to a range. These reasons include but are not limited to user safety, hardware limitations, or algorithm features.


// Defines the range within which an ear’s sensitivity point may have been clamped, if any.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAudiogramSensitivityPointClampingRange
type HKAudiogramSensitivityPointClampingRange struct {
	objectivec.Object
}

// HKAudiogramSensitivityPointClampingRangeFrom constructs a [HKAudiogramSensitivityPointClampingRange] from an unsafe.Pointer.
//
// Defines the range within which an ear’s sensitivity point may have been clamped, if any.
func HKAudiogramSensitivityPointClampingRangeFrom(ptr unsafe.Pointer) HKAudiogramSensitivityPointClampingRange {
	return HKAudiogramSensitivityPointClampingRange{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKAudiogramSensitivityPointClampingRange */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAudiogramSensitivityPointClampingRange/init(lowerBound:upperBound:)
func NewHKAudiogramSensitivityPointClampingRangeWithLowerBoundUpperBoundError(lowerBound objc.IObject /* cross-framework: NSNumber */, upperBound objc.IObject /* cross-framework: NSNumber */, errorOut objectivec.IObject) HKAudiogramSensitivityPointClampingRange {
	rv := objc.Send[HKAudiogramSensitivityPointClampingRange](objc.ID(getHKAudiogramSensitivityPointClampingRangeClass().class), objc.Sel("clampingRangeWithLowerBound:upperBound:error:"), lowerBound, upperBound, errorOut)
	return rv
}/* debug [class_init_methods/constructor]: NewHKAudiogramSensitivityPointClampingRangeWithLowerBoundUpperBoundError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKAudiogramSensitivityPointClampingRange */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAudiogramSensitivityPointClampingRange/init(lowerBound:upperBound:)
func (hc _HKAudiogramSensitivityPointClampingRangeClass) ClampingRangeWithLowerBoundUpperBoundError(lowerBound objc.IObject /* cross-framework: NSNumber */, upperBound objc.IObject /* cross-framework: NSNumber */, errorOut objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("clampingRangeWithLowerBound:upperBound:error:"), lowerBound, upperBound, errorOut)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ClampingRangeWithLowerBoundUpperBoundError) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKAudiogramSensitivityPointClampingRange */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKAudiogramSensitivityPointClampingRange */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKAudiogramSensitivityPointClampingRange */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAudiogramSensitivityPointClampingRange/lowerBound
func (h_ HKAudiogramSensitivityPointClampingRange) LowerBound() IHKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("lowerBound"))
	return rv
}/* debug [instance_properties/getter]: lowerBound */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAudiogramSensitivityPointClampingRange/upperBound
func (h_ HKAudiogramSensitivityPointClampingRange) UpperBound() IHKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("upperBound"))
	return rv
}/* debug [instance_properties/getter]: upperBound */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKAudiogramSensitivityPointClampingRange */


