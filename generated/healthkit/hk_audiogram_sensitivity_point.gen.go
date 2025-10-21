// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
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




