// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mDeviceConnectNotification] class.
var (
	MDeviceConnectNotificationClass     _mDeviceConnectNotificationClass
	MDeviceConnectNotificationClassOnce sync.Once
)

func getmDeviceConnectNotificationClass() _mDeviceConnectNotificationClass {
	MDeviceConnectNotificationClassOnce.Do(func() {
		MDeviceConnectNotificationClass = _mDeviceConnectNotificationClass{objc.GetClass("mDeviceConnectNotification")}
	})
	return MDeviceConnectNotificationClass
}

type _mDeviceConnectNotificationClass struct {
	class objc.Class
}

// An interface definition for the [mDeviceConnectNotification] class.
type ImDeviceConnectNotification interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/mDeviceConnectNotification
type mDeviceConnectNotification struct {
	objectivec.Object
}

// mDeviceConnectNotificationFrom constructs a [mDeviceConnectNotification] from an unsafe.Pointer.
func mDeviceConnectNotificationFrom(ptr unsafe.Pointer) mDeviceConnectNotification {
	return mDeviceConnectNotification{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mDeviceConnectNotificationClass) Alloc() mDeviceConnectNotification {
	rv := objc.Send[mDeviceConnectNotification](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mDeviceConnectNotificationClass) New() mDeviceConnectNotification {
	rv := objc.Send[mDeviceConnectNotification](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mDeviceConnectNotification) Init() mDeviceConnectNotification {
	rv := objc.Send[mDeviceConnectNotification](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mDeviceConnectNotification) Autorelease() mDeviceConnectNotification {
	rv := objc.Send[mDeviceConnectNotification](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmDeviceConnectNotification creates a new mDeviceConnectNotification instance.
func NewmDeviceConnectNotification() mDeviceConnectNotification {
	return getmDeviceConnectNotificationClass().New()
}




