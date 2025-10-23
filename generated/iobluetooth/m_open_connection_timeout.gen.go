// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mOpenConnectionTimeout] class.
var (
	MOpenConnectionTimeoutClass     _mOpenConnectionTimeoutClass
	MOpenConnectionTimeoutClassOnce sync.Once
)

func getmOpenConnectionTimeoutClass() _mOpenConnectionTimeoutClass {
	MOpenConnectionTimeoutClassOnce.Do(func() {
		MOpenConnectionTimeoutClass = _mOpenConnectionTimeoutClass{objc.GetClass("mOpenConnectionTimeout")}
	})
	return MOpenConnectionTimeoutClass
}

type _mOpenConnectionTimeoutClass struct {
	class objc.Class
}

// An interface definition for the [mOpenConnectionTimeout] class.
type ImOpenConnectionTimeout interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSession/mOpenConnectionTimeout
type mOpenConnectionTimeout struct {
	objectivec.Object
}

// mOpenConnectionTimeoutFrom constructs a [mOpenConnectionTimeout] from an unsafe.Pointer.
func mOpenConnectionTimeoutFrom(ptr unsafe.Pointer) mOpenConnectionTimeout {
	return mOpenConnectionTimeout{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mOpenConnectionTimeoutClass) Alloc() mOpenConnectionTimeout {
	rv := objc.Send[mOpenConnectionTimeout](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mOpenConnectionTimeoutClass) New() mOpenConnectionTimeout {
	rv := objc.Send[mOpenConnectionTimeout](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mOpenConnectionTimeout) Init() mOpenConnectionTimeout {
	rv := objc.Send[mOpenConnectionTimeout](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mOpenConnectionTimeout) Autorelease() mOpenConnectionTimeout {
	rv := objc.Send[mOpenConnectionTimeout](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmOpenConnectionTimeout creates a new mOpenConnectionTimeout instance.
func NewmOpenConnectionTimeout() mOpenConnectionTimeout {
	return getmOpenConnectionTimeoutClass().New()
}




