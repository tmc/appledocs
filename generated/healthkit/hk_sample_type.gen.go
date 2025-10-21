// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [HKSampleType] class.
type IHKSampleType interface {
	IHKObjectType
}

// An abstract superclass for all classes that identify a specific type of sample when working with the HealthKit store.
//
// The class is an abstract subclass of the class, used to represent data samples. Never instantiate an object directly. Instead, work with one of its concrete subclasses: , , , or classes.
//
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

// Alloc allocates a new instance without initialization.
func (hc _HKSampleTypeClass) Alloc() HKSampleType {
	rv := objc.Send[HKSampleType](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// A Boolean value that indicates whether HealthKit supports recalibrating the prediction algorithm used to produce estimates for this sample type.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKSampleType/allowsRecalibrationForEstimates
func (h_ HKSampleType) AllowsRecalibrationForEstimates() bool {
	rv := objc.Send[bool](h_.ID, objc.Sel("allowsRecalibrationForEstimates"))
	return rv
}



