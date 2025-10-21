// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mServerDevice] class.
var (
	MServerDeviceClass     _mServerDeviceClass
	MServerDeviceClassOnce sync.Once
)

func getmServerDeviceClass() _mServerDeviceClass {
	MServerDeviceClassOnce.Do(func() {
		MServerDeviceClass = _mServerDeviceClass{objc.GetClass("mServerDevice")}
	})
	return MServerDeviceClass
}

type _mServerDeviceClass struct {
	class objc.Class
}

// An interface definition for the [mServerDevice] class.
type ImServerDevice interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/mServerDevice
type mServerDevice struct {
	objectivec.Object
}

// mServerDeviceFrom constructs a [mServerDevice] from an unsafe.Pointer.
func mServerDeviceFrom(ptr unsafe.Pointer) mServerDevice {
	return mServerDevice{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mServerDeviceClass) Alloc() mServerDevice {
	rv := objc.Send[mServerDevice](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mServerDeviceClass) New() mServerDevice {
	rv := objc.Send[mServerDevice](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mServerDevice) Init() mServerDevice {
	rv := objc.Send[mServerDevice](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mServerDevice) Autorelease() mServerDevice {
	rv := objc.Send[mServerDevice](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmServerDevice creates a new mServerDevice instance.
func NewmServerDevice() mServerDevice {
	return getmServerDeviceClass().New()
}




