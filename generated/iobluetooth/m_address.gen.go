// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mAddress] class.
var (
	MAddressClass     _mAddressClass
	MAddressClassOnce sync.Once
)

func getmAddressClass() _mAddressClass {
	MAddressClassOnce.Do(func() {
		MAddressClass = _mAddressClass{objc.GetClass("mAddress")}
	})
	return MAddressClass
}

type _mAddressClass struct {
	class objc.Class
}

// An interface definition for the [mAddress] class.
type ImAddress interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/mAddress
type mAddress struct {
	objectivec.Object
}

// mAddressFrom constructs a [mAddress] from an unsafe.Pointer.
func mAddressFrom(ptr unsafe.Pointer) mAddress {
	return mAddress{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mAddressClass) Alloc() mAddress {
	rv := objc.Send[mAddress](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mAddressClass) New() mAddress {
	rv := objc.Send[mAddress](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mAddress) Init() mAddress {
	rv := objc.Send[mAddress](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mAddress) Autorelease() mAddress {
	rv := objc.Send[mAddress](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmAddress creates a new mAddress instance.
func NewmAddress() mAddress {
	return getmAddressClass().New()
}




