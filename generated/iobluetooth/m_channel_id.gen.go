// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mChannelID] class.
var (
	MChannelIDClass     _mChannelIDClass
	MChannelIDClassOnce sync.Once
)

func getmChannelIDClass() _mChannelIDClass {
	MChannelIDClassOnce.Do(func() {
		MChannelIDClass = _mChannelIDClass{objc.GetClass("mChannelID")}
	})
	return MChannelIDClass
}

type _mChannelIDClass struct {
	class objc.Class
}

// An interface definition for the [mChannelID] class.
type ImChannelID interface {
	objectivec.IObject
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel/mChannelID
type mChannelID struct {
	objectivec.Object
}

// mChannelIDFrom constructs a [mChannelID] from an unsafe.Pointer.
func mChannelIDFrom(ptr unsafe.Pointer) mChannelID {
	return mChannelID{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mChannelIDClass) Alloc() mChannelID {
	rv := objc.Send[mChannelID](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mChannelIDClass) New() mChannelID {
	rv := objc.Send[mChannelID](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mChannelID) Init() mChannelID {
	rv := objc.Send[mChannelID](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mChannelID) Autorelease() mChannelID {
	rv := objc.Send[mChannelID](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmChannelID creates a new mChannelID instance.
func NewmChannelID() mChannelID {
	return getmChannelIDClass().New()
}




