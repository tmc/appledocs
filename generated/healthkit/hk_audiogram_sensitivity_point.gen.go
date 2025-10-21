// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [HKAudiogramSensitivityPoint] class.
var (
	HKAudiogramSensitivityPointClass     _HKAudiogramSensitivityPointClass
	HKAudiogramSensitivityPointClassOnce sync.Once
)

func getHKAudiogramSensitivityPointClass() _HKAudiogramSensitivityPointClass {
	HKAudiogramSensitivityPointClassOnce.Do(func() {
		HKAudiogramSensitivityPointClass = _HKAudiogramSensitivityPointClass{objc.GetClass("HKAudiogramSensitivityPoint")}
	})
	return HKAudiogramSensitivityPointClass
}

type _HKAudiogramSensitivityPointClass struct {
	class objc.Class
}

// An interface definition for the [HKAudiogramSensitivityPoint] class.
type IHKAudiogramSensitivityPoint interface {
	objectivec.IObject
}

// A hearing sensitivity reading associated with a hearing test.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAudiogramSensitivityPoint
type HKAudiogramSensitivityPoint struct {
	objectivec.Object
}

// HKAudiogramSensitivityPointFrom constructs a [HKAudiogramSensitivityPoint] from an unsafe.Pointer.
//
// A hearing sensitivity reading associated with a hearing test.
func HKAudiogramSensitivityPointFrom(ptr unsafe.Pointer) HKAudiogramSensitivityPoint {
	return HKAudiogramSensitivityPoint{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (hc _HKAudiogramSensitivityPointClass) Alloc() HKAudiogramSensitivityPoint {
	rv := objc.Send[HKAudiogramSensitivityPoint](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKAudiogramSensitivityPointClass) New() HKAudiogramSensitivityPoint {
	rv := objc.Send[HKAudiogramSensitivityPoint](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKAudiogramSensitivityPoint) Init() HKAudiogramSensitivityPoint {
	rv := objc.Send[HKAudiogramSensitivityPoint](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKAudiogramSensitivityPoint) Autorelease() HKAudiogramSensitivityPoint {
	rv := objc.Send[HKAudiogramSensitivityPoint](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKAudiogramSensitivityPoint creates a new HKAudiogramSensitivityPoint instance.
func NewHKAudiogramSensitivityPoint() HKAudiogramSensitivityPoint {
	return getHKAudiogramSensitivityPointClass().New()
}


// The sensitivity of the left ear.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkaudiogramsensitivitypoint/leftearsensitivity
func (h_ HKAudiogramSensitivityPoint) LeftEarSensitivity() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("leftEarSensitivity"))
	return rv
}


// SetLeftEarSensitivity sets the value of the leftEarSensitivity property.
// The sensitivity of the left ear.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkaudiogramsensitivitypoint/leftearsensitivity
func (h_ HKAudiogramSensitivityPoint) SetLeftEarSensitivity(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setLeftEarSensitivity:"), value)
}

// The sensitivity of the right ear.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkaudiogramsensitivitypoint/rightearsensitivity
func (h_ HKAudiogramSensitivityPoint) RightEarSensitivity() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("rightEarSensitivity"))
	return rv
}


// SetRightEarSensitivity sets the value of the rightEarSensitivity property.
// The sensitivity of the right ear.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkaudiogramsensitivitypoint/rightearsensitivity
func (h_ HKAudiogramSensitivityPoint) SetRightEarSensitivity(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setRightEarSensitivity:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkaudiogramsensitivitypoint/tests
func (h_ HKAudiogramSensitivityPoint) Tests() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("tests"))
	return rv
}


// SetTests sets the value of the tests property.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkaudiogramsensitivitypoint/tests
func (h_ HKAudiogramSensitivityPoint) SetTests(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setTests:"), value)
}

// The frequency tested in the hearing test.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkaudiogramsensitivitypoint/frequency
func (h_ HKAudiogramSensitivityPoint) Frequency() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("frequency"))
	return rv
}


// SetFrequency sets the value of the frequency property.
// The frequency tested in the hearing test.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkaudiogramsensitivitypoint/frequency
func (h_ HKAudiogramSensitivityPoint) SetFrequency(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setFrequency:"), value)
}



