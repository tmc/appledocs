// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mChannelPSM] class.
var (
	MChannelPSMClass     _mChannelPSMClass
	MChannelPSMClassOnce sync.Once
)

func getmChannelPSMClass() _mChannelPSMClass {
	MChannelPSMClassOnce.Do(func() {
		MChannelPSMClass = _mChannelPSMClass{objc.GetClass("mChannelPSM")}
	})
	return MChannelPSMClass
}

type _mChannelPSMClass struct {
	class objc.Class
}

// An interface definition for the [mChannelPSM] class.
type ImChannelPSM interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/mChannelPSM
type mChannelPSM struct {
	objectivec.Object
}

// mChannelPSMFrom constructs a [mChannelPSM] from an unsafe.Pointer.
func mChannelPSMFrom(ptr unsafe.Pointer) mChannelPSM {
	return mChannelPSM{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mChannelPSMClass) Alloc() mChannelPSM {
	rv := objc.Send[mChannelPSM](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mChannelPSMClass) New() mChannelPSM {
	rv := objc.Send[mChannelPSM](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mChannelPSM) Init() mChannelPSM {
	rv := objc.Send[mChannelPSM](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mChannelPSM) Autorelease() mChannelPSM {
	rv := objc.Send[mChannelPSM](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmChannelPSM creates a new mChannelPSM instance.
func NewmChannelPSM() mChannelPSM {
	return getmChannelPSMClass().New()
}




