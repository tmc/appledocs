// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mLastServicesUpdate] class.
var (
	MLastServicesUpdateClass     _mLastServicesUpdateClass
	MLastServicesUpdateClassOnce sync.Once
)

func getmLastServicesUpdateClass() _mLastServicesUpdateClass {
	MLastServicesUpdateClassOnce.Do(func() {
		MLastServicesUpdateClass = _mLastServicesUpdateClass{objc.GetClass("mLastServicesUpdate")}
	})
	return MLastServicesUpdateClass
}

type _mLastServicesUpdateClass struct {
	class objc.Class
}

// An interface definition for the [mLastServicesUpdate] class.
type ImLastServicesUpdate interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/mLastServicesUpdate
type mLastServicesUpdate struct {
	objectivec.Object
}

// mLastServicesUpdateFrom constructs a [mLastServicesUpdate] from an unsafe.Pointer.
func mLastServicesUpdateFrom(ptr unsafe.Pointer) mLastServicesUpdate {
	return mLastServicesUpdate{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mLastServicesUpdateClass) Alloc() mLastServicesUpdate {
	rv := objc.Send[mLastServicesUpdate](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mLastServicesUpdateClass) New() mLastServicesUpdate {
	rv := objc.Send[mLastServicesUpdate](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mLastServicesUpdate) Init() mLastServicesUpdate {
	rv := objc.Send[mLastServicesUpdate](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mLastServicesUpdate) Autorelease() mLastServicesUpdate {
	rv := objc.Send[mLastServicesUpdate](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmLastServicesUpdate creates a new mLastServicesUpdate instance.
func NewmLastServicesUpdate() mLastServicesUpdate {
	return getmLastServicesUpdateClass().New()
}




