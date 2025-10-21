// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mDeviceForService] class.
var (
	MDeviceForServiceClass     _mDeviceForServiceClass
	MDeviceForServiceClassOnce sync.Once
)

func getmDeviceForServiceClass() _mDeviceForServiceClass {
	MDeviceForServiceClassOnce.Do(func() {
		MDeviceForServiceClass = _mDeviceForServiceClass{objc.GetClass("mDeviceForService")}
	})
	return MDeviceForServiceClass
}

type _mDeviceForServiceClass struct {
	class objc.Class
}

// An interface definition for the [mDeviceForService] class.
type ImDeviceForService interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceRecord/mDeviceForService
type mDeviceForService struct {
	objectivec.Object
}

// mDeviceForServiceFrom constructs a [mDeviceForService] from an unsafe.Pointer.
func mDeviceForServiceFrom(ptr unsafe.Pointer) mDeviceForService {
	return mDeviceForService{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mDeviceForServiceClass) Alloc() mDeviceForService {
	rv := objc.Send[mDeviceForService](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mDeviceForServiceClass) New() mDeviceForService {
	rv := objc.Send[mDeviceForService](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mDeviceForService) Init() mDeviceForService {
	rv := objc.Send[mDeviceForService](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mDeviceForService) Autorelease() mDeviceForService {
	rv := objc.Send[mDeviceForService](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmDeviceForService creates a new mDeviceForService instance.
func NewmDeviceForService() mDeviceForService {
	return getmDeviceForServiceClass().New()
}




