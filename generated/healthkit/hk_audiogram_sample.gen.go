// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [HKAudiogramSample] class.
var (
	HKAudiogramSampleClass     _HKAudiogramSampleClass
	HKAudiogramSampleClassOnce sync.Once
)

func getHKAudiogramSampleClass() _HKAudiogramSampleClass {
	HKAudiogramSampleClassOnce.Do(func() {
		HKAudiogramSampleClass = _HKAudiogramSampleClass{objc.GetClass("HKAudiogramSample")}
	})
	return HKAudiogramSampleClass
}

type _HKAudiogramSampleClass struct {
	class objc.Class
}

// An interface definition for the [HKAudiogramSample] class.
type IHKAudiogramSample interface {
	IHKSample
	// properties:
	SensitivityPoints() IHKAudiogramSensitivityPoint
	SetSensitivityPoints(value IHKAudiogramSensitivityPoint)
	// methods:
}

// A sample that stores an audiogram.
//
// This sample stores the results from a hearing test. The sample stores the audiogram data as an array of sensitivity points.


// A sample that stores an audiogram.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAudiogramSample
type HKAudiogramSample struct {
	HKSample
}

// HKAudiogramSampleFrom constructs a [HKAudiogramSample] from an unsafe.Pointer.
//
// A sample that stores an audiogram.
func HKAudiogramSampleFrom(ptr unsafe.Pointer) HKAudiogramSample {
	return HKAudiogramSample{
		HKSample: HKSampleFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (hc _HKAudiogramSampleClass) Alloc() HKAudiogramSample {
	rv := objc.Send[HKAudiogramSample](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKAudiogramSampleClass) New() HKAudiogramSample {
	rv := objc.Send[HKAudiogramSample](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKAudiogramSample) Init() HKAudiogramSample {
	rv := objc.Send[HKAudiogramSample](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKAudiogramSample) Autorelease() HKAudiogramSample {
	rv := objc.Send[HKAudiogramSample](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKAudiogramSample creates a new HKAudiogramSample instance.
func NewHKAudiogramSample() HKAudiogramSample {
	return getHKAudiogramSampleClass().New()
}



// An array of sensitivity point objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkaudiogramsample/sensitivitypoints
func (h_ HKAudiogramSample) SensitivityPoints() IHKAudiogramSensitivityPoint {
	rv := objc.Send[HKAudiogramSensitivityPoint](h_.ID, objc.Sel("sensitivityPoints"))
	return rv
}


// An array of sensitivity point objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkaudiogramsample/sensitivitypoints
func (h_ HKAudiogramSample) SetSensitivityPoints(value IHKAudiogramSensitivityPoint) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setSensitivityPoints:"), value)
}



