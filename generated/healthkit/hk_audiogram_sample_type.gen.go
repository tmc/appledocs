// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [HKAudiogramSampleType] class.
var (
	HKAudiogramSampleTypeClass     _HKAudiogramSampleTypeClass
	HKAudiogramSampleTypeClassOnce sync.Once
)

func getHKAudiogramSampleTypeClass() _HKAudiogramSampleTypeClass {
	HKAudiogramSampleTypeClassOnce.Do(func() {
		HKAudiogramSampleTypeClass = _HKAudiogramSampleTypeClass{objc.GetClass("HKAudiogramSampleType")}
	})
	return HKAudiogramSampleTypeClass
}

type _HKAudiogramSampleTypeClass struct {
	class objc.Class
}

// An interface definition for the [HKAudiogramSampleType] class.
type IHKAudiogramSampleType interface {
	IHKSampleType
}

// A type that identifies samples that contain audiogram data.
//
// The class is a concrete subclass of the class. To create an audiogram sample type instance, use the object type’s convenience method. Use audiogram sample types to: Request permission to read or write audiogram samples. Create and share audiogram samples. Query for audiogram samples.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAudiogramSampleType
type HKAudiogramSampleType struct {
	HKSampleType
}

// HKAudiogramSampleTypeFrom constructs a [HKAudiogramSampleType] from an unsafe.Pointer.
//
// A type that identifies samples that contain audiogram data.
func HKAudiogramSampleTypeFrom(ptr unsafe.Pointer) HKAudiogramSampleType {
	return HKAudiogramSampleType{
		HKSampleType: HKSampleTypeFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (hc _HKAudiogramSampleTypeClass) Alloc() HKAudiogramSampleType {
	rv := objc.Send[HKAudiogramSampleType](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKAudiogramSampleTypeClass) New() HKAudiogramSampleType {
	rv := objc.Send[HKAudiogramSampleType](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKAudiogramSampleType) Init() HKAudiogramSampleType {
	rv := objc.Send[HKAudiogramSampleType](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKAudiogramSampleType) Autorelease() HKAudiogramSampleType {
	rv := objc.Send[HKAudiogramSampleType](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKAudiogramSampleType creates a new HKAudiogramSampleType instance.
func NewHKAudiogramSampleType() HKAudiogramSampleType {
	return getHKAudiogramSampleTypeClass().New()
}




