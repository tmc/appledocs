// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mIncomingDataListenerRefCon] class.
var (
	MIncomingDataListenerRefConClass     _mIncomingDataListenerRefConClass
	MIncomingDataListenerRefConClassOnce sync.Once
)

func getmIncomingDataListenerRefConClass() _mIncomingDataListenerRefConClass {
	MIncomingDataListenerRefConClassOnce.Do(func() {
		MIncomingDataListenerRefConClass = _mIncomingDataListenerRefConClass{objc.GetClass("mIncomingDataListenerRefCon")}
	})
	return MIncomingDataListenerRefConClass
}

type _mIncomingDataListenerRefConClass struct {
	class objc.Class
}

// An interface definition for the [mIncomingDataListenerRefCon] class.
type ImIncomingDataListenerRefCon interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel/mIncomingDataListenerRefCon
type mIncomingDataListenerRefCon struct {
	objectivec.Object
}

// mIncomingDataListenerRefConFrom constructs a [mIncomingDataListenerRefCon] from an unsafe.Pointer.
func mIncomingDataListenerRefConFrom(ptr unsafe.Pointer) mIncomingDataListenerRefCon {
	return mIncomingDataListenerRefCon{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mIncomingDataListenerRefConClass) Alloc() mIncomingDataListenerRefCon {
	rv := objc.Send[mIncomingDataListenerRefCon](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mIncomingDataListenerRefConClass) New() mIncomingDataListenerRefCon {
	rv := objc.Send[mIncomingDataListenerRefCon](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mIncomingDataListenerRefCon) Init() mIncomingDataListenerRefCon {
	rv := objc.Send[mIncomingDataListenerRefCon](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mIncomingDataListenerRefCon) Autorelease() mIncomingDataListenerRefCon {
	rv := objc.Send[mIncomingDataListenerRefCon](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmIncomingDataListenerRefCon creates a new mIncomingDataListenerRefCon instance.
func NewmIncomingDataListenerRefCon() mIncomingDataListenerRefCon {
	return getmIncomingDataListenerRefConClass().New()
}




