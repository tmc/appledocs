// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mChannelIsOpen] class.
var (
	MChannelIsOpenClass     _mChannelIsOpenClass
	MChannelIsOpenClassOnce sync.Once
)

func getmChannelIsOpenClass() _mChannelIsOpenClass {
	MChannelIsOpenClassOnce.Do(func() {
		MChannelIsOpenClass = _mChannelIsOpenClass{objc.GetClass("mChannelIsOpen")}
	})
	return MChannelIsOpenClass
}

type _mChannelIsOpenClass struct {
	class objc.Class
}

// An interface definition for the [mChannelIsOpen] class.
type ImChannelIsOpen interface {
	objectivec.IObject
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel/mChannelIsOpen
type mChannelIsOpen struct {
	objectivec.Object
}

// mChannelIsOpenFrom constructs a [mChannelIsOpen] from an unsafe.Pointer.
func mChannelIsOpenFrom(ptr unsafe.Pointer) mChannelIsOpen {
	return mChannelIsOpen{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mChannelIsOpenClass) Alloc() mChannelIsOpen {
	rv := objc.Send[mChannelIsOpen](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mChannelIsOpenClass) New() mChannelIsOpen {
	rv := objc.Send[mChannelIsOpen](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mChannelIsOpen) Init() mChannelIsOpen {
	rv := objc.Send[mChannelIsOpen](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mChannelIsOpen) Autorelease() mChannelIsOpen {
	rv := objc.Send[mChannelIsOpen](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmChannelIsOpen creates a new mChannelIsOpen instance.
func NewmChannelIsOpen() mChannelIsOpen {
	return getmChannelIsOpenClass().New()
}




