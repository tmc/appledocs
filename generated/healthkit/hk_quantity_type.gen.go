// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [HKQuantityType] class.
var (
	HKQuantityTypeClass     _HKQuantityTypeClass
	HKQuantityTypeClassOnce sync.Once
)

func getHKQuantityTypeClass() _HKQuantityTypeClass {
	HKQuantityTypeClassOnce.Do(func() {
		HKQuantityTypeClass = _HKQuantityTypeClass{objc.GetClass("HKQuantityType")}
	})
	return HKQuantityTypeClass
}

type _HKQuantityTypeClass struct {
	class objc.Class
}

// An interface definition for the [HKQuantityType] class.
type IHKQuantityType interface {
	IHKSampleType
	IsCompatibleWithUnit(unit unsafe.Pointer) bool
}

// A type that identifies samples that store numerical values.
//
// The class is a concrete subclass of the class. To create a quantity type instance, use the object type’s convenience method. Use quantity types to: Request permission to read or write matching quantity samples. Create and share matching quantity samples. Query for matching quantity samples.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuantityType
type HKQuantityType struct {
	HKSampleType
}

// HKQuantityTypeFrom constructs a [HKQuantityType] from an unsafe.Pointer.
//
// A type that identifies samples that store numerical values.
func HKQuantityTypeFrom(ptr unsafe.Pointer) HKQuantityType {
	return HKQuantityType{
		HKSampleType: HKSampleTypeFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (hc _HKQuantityTypeClass) Alloc() HKQuantityType {
	rv := objc.Send[HKQuantityType](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKQuantityTypeClass) New() HKQuantityType {
	rv := objc.Send[HKQuantityType](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKQuantityType) Init() HKQuantityType {
	rv := objc.Send[HKQuantityType](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKQuantityType) Autorelease() HKQuantityType {
	rv := objc.Send[HKQuantityType](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKQuantityType creates a new HKQuantityType instance.
func NewHKQuantityType() HKQuantityType {
	return getHKQuantityTypeClass().New()
}


// Returns a Boolean value that indicates whether the quantity type is compatible with the given unit.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuantityType/is(compatibleWith:)
func (h_ HKQuantityType) IsCompatibleWithUnit(unit unsafe.Pointer) bool {
	rv := objc.Send[bool](h_.ID, objc.Sel("isCompatibleWithUnit:"), unit)
	return rv
}

// The aggregation style for the given quantity type.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuantityType/aggregationStyle
func (h_ HKQuantityType) AggregationStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("aggregationStyle"))
	return rv
}



