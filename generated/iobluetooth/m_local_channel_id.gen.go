// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mLocalChannelID] class.
var (
	MLocalChannelIDClass     _mLocalChannelIDClass
	MLocalChannelIDClassOnce sync.Once
)

func getmLocalChannelIDClass() _mLocalChannelIDClass {
	MLocalChannelIDClassOnce.Do(func() {
		MLocalChannelIDClass = _mLocalChannelIDClass{objc.GetClass("mLocalChannelID")}
	})
	return MLocalChannelIDClass
}

type _mLocalChannelIDClass struct {
	class objc.Class
}

// An interface definition for the [mLocalChannelID] class.
type ImLocalChannelID interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/mLocalChannelID
type mLocalChannelID struct {
	objectivec.Object
}

// mLocalChannelIDFrom constructs a [mLocalChannelID] from an unsafe.Pointer.
func mLocalChannelIDFrom(ptr unsafe.Pointer) mLocalChannelID {
	return mLocalChannelID{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mLocalChannelIDClass) Alloc() mLocalChannelID {
	rv := objc.Send[mLocalChannelID](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mLocalChannelIDClass) New() mLocalChannelID {
	rv := objc.Send[mLocalChannelID](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mLocalChannelID) Init() mLocalChannelID {
	rv := objc.Send[mLocalChannelID](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mLocalChannelID) Autorelease() mLocalChannelID {
	rv := objc.Send[mLocalChannelID](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmLocalChannelID creates a new mLocalChannelID instance.
func NewmLocalChannelID() mLocalChannelID {
	return getmLocalChannelIDClass().New()
}




