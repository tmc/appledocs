// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mRFCOMMConnection] class.
var (
	MRFCOMMConnectionClass     _mRFCOMMConnectionClass
	MRFCOMMConnectionClassOnce sync.Once
)

func getmRFCOMMConnectionClass() _mRFCOMMConnectionClass {
	MRFCOMMConnectionClassOnce.Do(func() {
		MRFCOMMConnectionClass = _mRFCOMMConnectionClass{objc.GetClass("mRFCOMMConnection")}
	})
	return MRFCOMMConnectionClass
}

type _mRFCOMMConnectionClass struct {
	class objc.Class
}

// An interface definition for the [mRFCOMMConnection] class.
type ImRFCOMMConnection interface {
	objectivec.IObject
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/mRFCOMMConnection
type mRFCOMMConnection struct {
	objectivec.Object
}

// mRFCOMMConnectionFrom constructs a [mRFCOMMConnection] from an unsafe.Pointer.
func mRFCOMMConnectionFrom(ptr unsafe.Pointer) mRFCOMMConnection {
	return mRFCOMMConnection{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mRFCOMMConnectionClass) Alloc() mRFCOMMConnection {
	rv := objc.Send[mRFCOMMConnection](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mRFCOMMConnectionClass) New() mRFCOMMConnection {
	rv := objc.Send[mRFCOMMConnection](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mRFCOMMConnection) Init() mRFCOMMConnection {
	rv := objc.Send[mRFCOMMConnection](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mRFCOMMConnection) Autorelease() mRFCOMMConnection {
	rv := objc.Send[mRFCOMMConnection](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmRFCOMMConnection creates a new mRFCOMMConnection instance.
func NewmRFCOMMConnection() mRFCOMMConnection {
	return getmRFCOMMConnectionClass().New()
}




