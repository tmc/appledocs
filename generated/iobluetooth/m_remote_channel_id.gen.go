// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mRemoteChannelID] class.
var (
	MRemoteChannelIDClass     _mRemoteChannelIDClass
	MRemoteChannelIDClassOnce sync.Once
)

func getmRemoteChannelIDClass() _mRemoteChannelIDClass {
	MRemoteChannelIDClassOnce.Do(func() {
		MRemoteChannelIDClass = _mRemoteChannelIDClass{objc.GetClass("mRemoteChannelID")}
	})
	return MRemoteChannelIDClass
}

type _mRemoteChannelIDClass struct {
	class objc.Class
}

// An interface definition for the [mRemoteChannelID] class.
type ImRemoteChannelID interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/mRemoteChannelID
type mRemoteChannelID struct {
	objectivec.Object
}

// mRemoteChannelIDFrom constructs a [mRemoteChannelID] from an unsafe.Pointer.
func mRemoteChannelIDFrom(ptr unsafe.Pointer) mRemoteChannelID {
	return mRemoteChannelID{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mRemoteChannelIDClass) Alloc() mRemoteChannelID {
	rv := objc.Send[mRemoteChannelID](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mRemoteChannelIDClass) New() mRemoteChannelID {
	rv := objc.Send[mRemoteChannelID](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mRemoteChannelID) Init() mRemoteChannelID {
	rv := objc.Send[mRemoteChannelID](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mRemoteChannelID) Autorelease() mRemoteChannelID {
	rv := objc.Send[mRemoteChannelID](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmRemoteChannelID creates a new mRemoteChannelID instance.
func NewmRemoteChannelID() mRemoteChannelID {
	return getmRemoteChannelIDClass().New()
}




