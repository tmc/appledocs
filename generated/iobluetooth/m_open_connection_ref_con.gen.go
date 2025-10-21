// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mOpenConnectionRefCon] class.
var (
	MOpenConnectionRefConClass     _mOpenConnectionRefConClass
	MOpenConnectionRefConClassOnce sync.Once
)

func getmOpenConnectionRefConClass() _mOpenConnectionRefConClass {
	MOpenConnectionRefConClassOnce.Do(func() {
		MOpenConnectionRefConClass = _mOpenConnectionRefConClass{objc.GetClass("mOpenConnectionRefCon")}
	})
	return MOpenConnectionRefConClass
}

type _mOpenConnectionRefConClass struct {
	class objc.Class
}

// An interface definition for the [mOpenConnectionRefCon] class.
type ImOpenConnectionRefCon interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSession/mOpenConnectionRefCon
type mOpenConnectionRefCon struct {
	objectivec.Object
}

// mOpenConnectionRefConFrom constructs a [mOpenConnectionRefCon] from an unsafe.Pointer.
func mOpenConnectionRefConFrom(ptr unsafe.Pointer) mOpenConnectionRefCon {
	return mOpenConnectionRefCon{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mOpenConnectionRefConClass) Alloc() mOpenConnectionRefCon {
	rv := objc.Send[mOpenConnectionRefCon](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mOpenConnectionRefConClass) New() mOpenConnectionRefCon {
	rv := objc.Send[mOpenConnectionRefCon](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mOpenConnectionRefCon) Init() mOpenConnectionRefCon {
	rv := objc.Send[mOpenConnectionRefCon](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mOpenConnectionRefCon) Autorelease() mOpenConnectionRefCon {
	rv := objc.Send[mOpenConnectionRefCon](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmOpenConnectionRefCon creates a new mOpenConnectionRefCon instance.
func NewmOpenConnectionRefCon() mOpenConnectionRefCon {
	return getmOpenConnectionRefConClass().New()
}




