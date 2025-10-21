// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [HKAudiogramSensitivityPointClampingRange] class.
type IHKAudiogramSensitivityPointClampingRange interface {
	objectivec.IObject
}

// Defines the range within which an ear’s sensitivity point may have been clamped, if any.
//
// At times, it may be required to indicate that a sensitivity point has been clamped to a range. These reasons include but are not limited to user safety, hardware limitations, or algorithm features.
//
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

// Alloc allocates a new instance without initialization.
func (hc _HKAudiogramSensitivityPointClampingRangeClass) Alloc() HKAudiogramSensitivityPointClampingRange {
	rv := objc.Send[HKAudiogramSensitivityPointClampingRange](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAudiogramSensitivityPointClampingRange/init(lowerBound:upperBound:)
func NewHKAudiogramSensitivityPointClampingRangeWithLowerBoundUpperBoundError(lowerBound unsafe.Pointer, upperBound unsafe.Pointer, errorOut unsafe.Pointer) HKAudiogramSensitivityPointClampingRange {
	rv := objc.Send[HKAudiogramSensitivityPointClampingRange](objc.ID(getHKAudiogramSensitivityPointClampingRangeClass().class), objc.Sel("clampingRangeWithLowerBound:upperBound:error:"), lowerBound, upperBound, errorOut)
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAudiogramSensitivityPointClampingRange/init(lowerBound:upperBound:)
func (hc _HKAudiogramSensitivityPointClampingRangeClass) ClampingRangeWithLowerBoundUpperBoundError(lowerBound unsafe.Pointer, upperBound unsafe.Pointer, errorOut unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("clampingRangeWithLowerBound:upperBound:error:"), lowerBound, upperBound, errorOut)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAudiogramSensitivityPointClampingRange/lowerBound
func (h_ HKAudiogramSensitivityPointClampingRange) LowerBound() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("lowerBound"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAudiogramSensitivityPointClampingRange/upperBound
func (h_ HKAudiogramSensitivityPointClampingRange) UpperBound() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("upperBound"))
	return rv
}


