// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mIsIncoming] class.
var (
	MIsIncomingClass     _mIsIncomingClass
	MIsIncomingClassOnce sync.Once
)

func getmIsIncomingClass() _mIsIncomingClass {
	MIsIncomingClassOnce.Do(func() {
		MIsIncomingClass = _mIsIncomingClass{objc.GetClass("mIsIncoming")}
	})
	return MIsIncomingClass
}

type _mIsIncomingClass struct {
	class objc.Class
}

// An interface definition for the [mIsIncoming] class.
type ImIsIncoming interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel/mIsIncoming
type mIsIncoming struct {
	objectivec.Object
}

// mIsIncomingFrom constructs a [mIsIncoming] from an unsafe.Pointer.
func mIsIncomingFrom(ptr unsafe.Pointer) mIsIncoming {
	return mIsIncoming{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mIsIncomingClass) Alloc() mIsIncoming {
	rv := objc.Send[mIsIncoming](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mIsIncomingClass) New() mIsIncoming {
	rv := objc.Send[mIsIncoming](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mIsIncoming) Init() mIsIncoming {
	rv := objc.Send[mIsIncoming](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mIsIncoming) Autorelease() mIsIncoming {
	rv := objc.Send[mIsIncoming](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmIsIncoming creates a new mIsIncoming instance.
func NewmIsIncoming() mIsIncoming {
	return getmIsIncomingClass().New()
}




