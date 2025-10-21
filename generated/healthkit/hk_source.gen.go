// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [HKSource] class.
var (
	HKSourceClass     _HKSourceClass
	HKSourceClassOnce sync.Once
)

func getHKSourceClass() _HKSourceClass {
	HKSourceClassOnce.Do(func() {
		HKSourceClass = _HKSourceClass{objc.GetClass("HKSource")}
	})
	return HKSourceClass
}

type _HKSourceClass struct {
	class objc.Class
}

// An interface definition for the [HKSource] class.
type IHKSource interface {
	objectivec.IObject
}

// An object indicating the app or device that created a HealthKit sample
//
// Sources include apps and devices that save data to the HealthKit store. Currently, HealthKit supports only the direct import of data from Bluetooth LE heart rate monitors. All other devices need a companion app to collect and save the data to HealthKit.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKSource
type HKSource struct {
	objectivec.Object
}

// HKSourceFrom constructs a [HKSource] from an unsafe.Pointer.
//
// An object indicating the app or device that created a HealthKit sample
func HKSourceFrom(ptr unsafe.Pointer) HKSource {
	return HKSource{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (hc _HKSourceClass) Alloc() HKSource {
	rv := objc.Send[HKSource](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKSourceClass) New() HKSource {
	rv := objc.Send[HKSource](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKSource) Init() HKSource {
	rv := objc.Send[HKSource](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKSource) Autorelease() HKSource {
	rv := objc.Send[HKSource](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKSource creates a new HKSource instance.
func NewHKSource() HKSource {
	return getHKSourceClass().New()
}




