// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mClassOfDevice] class.
var (
	MClassOfDeviceClass     _mClassOfDeviceClass
	MClassOfDeviceClassOnce sync.Once
)

func getmClassOfDeviceClass() _mClassOfDeviceClass {
	MClassOfDeviceClassOnce.Do(func() {
		MClassOfDeviceClass = _mClassOfDeviceClass{objc.GetClass("mClassOfDevice")}
	})
	return MClassOfDeviceClass
}

type _mClassOfDeviceClass struct {
	class objc.Class
}

// An interface definition for the [mClassOfDevice] class.
type ImClassOfDevice interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/mClassOfDevice
type mClassOfDevice struct {
	objectivec.Object
}

// mClassOfDeviceFrom constructs a [mClassOfDevice] from an unsafe.Pointer.
func mClassOfDeviceFrom(ptr unsafe.Pointer) mClassOfDevice {
	return mClassOfDevice{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mClassOfDeviceClass) Alloc() mClassOfDevice {
	rv := objc.Send[mClassOfDevice](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mClassOfDeviceClass) New() mClassOfDevice {
	rv := objc.Send[mClassOfDevice](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mClassOfDevice) Init() mClassOfDevice {
	rv := objc.Send[mClassOfDevice](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mClassOfDevice) Autorelease() mClassOfDevice {
	rv := objc.Send[mClassOfDevice](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmClassOfDevice creates a new mClassOfDevice instance.
func NewmClassOfDevice() mClassOfDevice {
	return getmClassOfDeviceClass().New()
}




