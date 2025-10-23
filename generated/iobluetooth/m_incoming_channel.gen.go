// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mIncomingChannel] class.
var (
	MIncomingChannelClass     _mIncomingChannelClass
	MIncomingChannelClassOnce sync.Once
)

func getmIncomingChannelClass() _mIncomingChannelClass {
	MIncomingChannelClassOnce.Do(func() {
		MIncomingChannelClass = _mIncomingChannelClass{objc.GetClass("mIncomingChannel")}
	})
	return MIncomingChannelClass
}

type _mIncomingChannelClass struct {
	class objc.Class
}

// An interface definition for the [mIncomingChannel] class.
type ImIncomingChannel interface {
	objectivec.IObject
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/mIncomingChannel
type mIncomingChannel struct {
	objectivec.Object
}

// mIncomingChannelFrom constructs a [mIncomingChannel] from an unsafe.Pointer.
func mIncomingChannelFrom(ptr unsafe.Pointer) mIncomingChannel {
	return mIncomingChannel{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mIncomingChannelClass) Alloc() mIncomingChannel {
	rv := objc.Send[mIncomingChannel](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mIncomingChannelClass) New() mIncomingChannel {
	rv := objc.Send[mIncomingChannel](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mIncomingChannel) Init() mIncomingChannel {
	rv := objc.Send[mIncomingChannel](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mIncomingChannel) Autorelease() mIncomingChannel {
	rv := objc.Send[mIncomingChannel](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmIncomingChannel creates a new mIncomingChannel instance.
func NewmIncomingChannel() mIncomingChannel {
	return getmIncomingChannelClass().New()
}




