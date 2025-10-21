// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mRFCOMMChannelID] class.
var (
	MRFCOMMChannelIDClass     _mRFCOMMChannelIDClass
	MRFCOMMChannelIDClassOnce sync.Once
)

func getmRFCOMMChannelIDClass() _mRFCOMMChannelIDClass {
	MRFCOMMChannelIDClassOnce.Do(func() {
		MRFCOMMChannelIDClass = _mRFCOMMChannelIDClass{objc.GetClass("mRFCOMMChannelID")}
	})
	return MRFCOMMChannelIDClass
}

type _mRFCOMMChannelIDClass struct {
	class objc.Class
}

// An interface definition for the [mRFCOMMChannelID] class.
type ImRFCOMMChannelID interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSession/mRFCOMMChannelID
type mRFCOMMChannelID struct {
	objectivec.Object
}

// mRFCOMMChannelIDFrom constructs a [mRFCOMMChannelID] from an unsafe.Pointer.
func mRFCOMMChannelIDFrom(ptr unsafe.Pointer) mRFCOMMChannelID {
	return mRFCOMMChannelID{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mRFCOMMChannelIDClass) Alloc() mRFCOMMChannelID {
	rv := objc.Send[mRFCOMMChannelID](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mRFCOMMChannelIDClass) New() mRFCOMMChannelID {
	rv := objc.Send[mRFCOMMChannelID](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mRFCOMMChannelID) Init() mRFCOMMChannelID {
	rv := objc.Send[mRFCOMMChannelID](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mRFCOMMChannelID) Autorelease() mRFCOMMChannelID {
	rv := objc.Send[mRFCOMMChannelID](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmRFCOMMChannelID creates a new mRFCOMMChannelID instance.
func NewmRFCOMMChannelID() mRFCOMMChannelID {
	return getmRFCOMMChannelIDClass().New()
}




