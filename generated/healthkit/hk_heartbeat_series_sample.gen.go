// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [HKHeartbeatSeriesSample] class.
var (
	HKHeartbeatSeriesSampleClass     _HKHeartbeatSeriesSampleClass
	HKHeartbeatSeriesSampleClassOnce sync.Once
)

func getHKHeartbeatSeriesSampleClass() _HKHeartbeatSeriesSampleClass {
	HKHeartbeatSeriesSampleClassOnce.Do(func() {
		HKHeartbeatSeriesSampleClass = _HKHeartbeatSeriesSampleClass{objc.GetClass("HKHeartbeatSeriesSample")}
	})
	return HKHeartbeatSeriesSampleClass
}

type _HKHeartbeatSeriesSampleClass struct {
	class objc.Class
}

// An interface definition for the [HKHeartbeatSeriesSample] class.
type IHKHeartbeatSeriesSample interface {
	IHKSeriesSample
	// properties:
	HKMetadataKeyAlgorithmVersion() objc.IObject /* cross-framework: NSString */
	// methods:
}

// A sample that represents a series of heartbeats.
//
// Use a to access the underlying heartbeat data. The class is a subclass of the class. These samples are immutable; you set the sample’s properties when you build them, and they can’t change.


// A sample that represents a series of heartbeats.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHeartbeatSeriesSample
type HKHeartbeatSeriesSample struct {
	HKSeriesSample
}

// HKHeartbeatSeriesSampleFrom constructs a [HKHeartbeatSeriesSample] from an unsafe.Pointer.
//
// A sample that represents a series of heartbeats.
func HKHeartbeatSeriesSampleFrom(ptr unsafe.Pointer) HKHeartbeatSeriesSample {
	return HKHeartbeatSeriesSample{
		HKSeriesSample: HKSeriesSampleFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (hc _HKHeartbeatSeriesSampleClass) Alloc() HKHeartbeatSeriesSample {
	rv := objc.Send[HKHeartbeatSeriesSample](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKHeartbeatSeriesSampleClass) New() HKHeartbeatSeriesSample {
	rv := objc.Send[HKHeartbeatSeriesSample](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKHeartbeatSeriesSample) Init() HKHeartbeatSeriesSample {
	rv := objc.Send[HKHeartbeatSeriesSample](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKHeartbeatSeriesSample) Autorelease() HKHeartbeatSeriesSample {
	rv := objc.Send[HKHeartbeatSeriesSample](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKHeartbeatSeriesSample creates a new HKHeartbeatSeriesSample instance.
func NewHKHeartbeatSeriesSample() HKHeartbeatSeriesSample {
	return getHKHeartbeatSeriesSampleClass().New()
}



// A key that indicates the version number of the algorithm used to calculate the sample’s value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkmetadatakeyalgorithmversion
func (h_ HKHeartbeatSeriesSample) HKMetadataKeyAlgorithmVersion() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKMetadataKeyAlgorithmVersion"))
	return rv
}



