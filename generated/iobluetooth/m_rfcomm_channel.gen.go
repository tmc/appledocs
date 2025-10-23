// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mRFCOMMChannel] class.
var (
	MRFCOMMChannelClass     _mRFCOMMChannelClass
	MRFCOMMChannelClassOnce sync.Once
)

func getmRFCOMMChannelClass() _mRFCOMMChannelClass {
	MRFCOMMChannelClassOnce.Do(func() {
		MRFCOMMChannelClass = _mRFCOMMChannelClass{objc.GetClass("mRFCOMMChannel")}
	})
	return MRFCOMMChannelClass
}

type _mRFCOMMChannelClass struct {
	class objc.Class
}

// An interface definition for the [mRFCOMMChannel] class.
type ImRFCOMMChannel interface {
	objectivec.IObject
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSession/mRFCOMMChannel
type mRFCOMMChannel struct {
	objectivec.Object
}

// mRFCOMMChannelFrom constructs a [mRFCOMMChannel] from an unsafe.Pointer.
func mRFCOMMChannelFrom(ptr unsafe.Pointer) mRFCOMMChannel {
	return mRFCOMMChannel{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mRFCOMMChannelClass) Alloc() mRFCOMMChannel {
	rv := objc.Send[mRFCOMMChannel](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mRFCOMMChannelClass) New() mRFCOMMChannel {
	rv := objc.Send[mRFCOMMChannel](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mRFCOMMChannel) Init() mRFCOMMChannel {
	rv := objc.Send[mRFCOMMChannel](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mRFCOMMChannel) Autorelease() mRFCOMMChannel {
	rv := objc.Send[mRFCOMMChannel](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmRFCOMMChannel creates a new mRFCOMMChannel instance.
func NewmRFCOMMChannel() mRFCOMMChannel {
	return getmRFCOMMChannelClass().New()
}




