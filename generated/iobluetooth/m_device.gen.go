// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mDevice] class.
var (
	MDeviceClass     _mDeviceClass
	MDeviceClassOnce sync.Once
)

func getmDeviceClass() _mDeviceClass {
	MDeviceClassOnce.Do(func() {
		MDeviceClass = _mDeviceClass{objc.GetClass("mDevice")}
	})
	return MDeviceClass
}

type _mDeviceClass struct {
	class objc.Class
}

// An interface definition for the [mDevice] class.
type ImDevice interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/mDevice
type mDevice struct {
	objectivec.Object
}

// mDeviceFrom constructs a [mDevice] from an unsafe.Pointer.
func mDeviceFrom(ptr unsafe.Pointer) mDevice {
	return mDevice{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mDeviceClass) Alloc() mDevice {
	rv := objc.Send[mDevice](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mDeviceClass) New() mDevice {
	rv := objc.Send[mDevice](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mDevice) Init() mDevice {
	rv := objc.Send[mDevice](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mDevice) Autorelease() mDevice {
	rv := objc.Send[mDevice](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmDevice creates a new mDevice instance.
func NewmDevice() mDevice {
	return getmDeviceClass().New()
}




