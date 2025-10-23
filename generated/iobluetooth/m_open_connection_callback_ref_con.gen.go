// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mOpenConnectionCallbackRefCon] class.
var (
	MOpenConnectionCallbackRefConClass     _mOpenConnectionCallbackRefConClass
	MOpenConnectionCallbackRefConClassOnce sync.Once
)

func getmOpenConnectionCallbackRefConClass() _mOpenConnectionCallbackRefConClass {
	MOpenConnectionCallbackRefConClassOnce.Do(func() {
		MOpenConnectionCallbackRefConClass = _mOpenConnectionCallbackRefConClass{objc.GetClass("mOpenConnectionCallbackRefCon")}
	})
	return MOpenConnectionCallbackRefConClass
}

type _mOpenConnectionCallbackRefConClass struct {
	class objc.Class
}

// An interface definition for the [mOpenConnectionCallbackRefCon] class.
type ImOpenConnectionCallbackRefCon interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSession/mOpenConnectionCallbackRefCon
type mOpenConnectionCallbackRefCon struct {
	objectivec.Object
}

// mOpenConnectionCallbackRefConFrom constructs a [mOpenConnectionCallbackRefCon] from an unsafe.Pointer.
func mOpenConnectionCallbackRefConFrom(ptr unsafe.Pointer) mOpenConnectionCallbackRefCon {
	return mOpenConnectionCallbackRefCon{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mOpenConnectionCallbackRefConClass) Alloc() mOpenConnectionCallbackRefCon {
	rv := objc.Send[mOpenConnectionCallbackRefCon](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mOpenConnectionCallbackRefConClass) New() mOpenConnectionCallbackRefCon {
	rv := objc.Send[mOpenConnectionCallbackRefCon](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mOpenConnectionCallbackRefCon) Init() mOpenConnectionCallbackRefCon {
	rv := objc.Send[mOpenConnectionCallbackRefCon](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mOpenConnectionCallbackRefCon) Autorelease() mOpenConnectionCallbackRefCon {
	rv := objc.Send[mOpenConnectionCallbackRefCon](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmOpenConnectionCallbackRefCon creates a new mOpenConnectionCallbackRefCon instance.
func NewmOpenConnectionCallbackRefCon() mOpenConnectionCallbackRefCon {
	return getmOpenConnectionCallbackRefConClass().New()
}




