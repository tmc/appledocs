// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [HKDevice] class.
var (
	HKDeviceClass     _HKDeviceClass
	HKDeviceClassOnce sync.Once
)

func getHKDeviceClass() _HKDeviceClass {
	HKDeviceClassOnce.Do(func() {
		HKDeviceClass = _HKDeviceClass{objc.GetClass("HKDevice")}
	})
	return HKDeviceClass
}

type _HKDeviceClass struct {
	class objc.Class
}

// An interface definition for the [HKDevice] class.
type IHKDevice interface {
	objectivec.IObject
}

// A device that generates data for HealthKit.
//
// Devices include Apple Watch, iPhone, and any other health or fitness peripherals that produce the sample data stored in HealthKit. Device objects are immutable: You set the device’s properties when you create the object, and they cannot change.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKDevice
type HKDevice struct {
	objectivec.Object
}

// HKDeviceFrom constructs a [HKDevice] from an unsafe.Pointer.
//
// A device that generates data for HealthKit.
func HKDeviceFrom(ptr unsafe.Pointer) HKDevice {
	return HKDevice{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (hc _HKDeviceClass) Alloc() HKDevice {
	rv := objc.Send[HKDevice](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKDeviceClass) New() HKDevice {
	rv := objc.Send[HKDevice](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKDevice) Init() HKDevice {
	rv := objc.Send[HKDevice](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKDevice) Autorelease() HKDevice {
	rv := objc.Send[HKDevice](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKDevice creates a new HKDevice instance.
func NewHKDevice() HKDevice {
	return getHKDeviceClass().New()
}


// An identifier that uniquely identifies the device object on the hardware running this code.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKDevice/localIdentifier
func (h_ HKDevice) LocalIdentifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("localIdentifier"))
	return rv
}



