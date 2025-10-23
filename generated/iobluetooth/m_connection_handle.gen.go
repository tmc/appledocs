// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mConnectionHandle] class.
var (
	MConnectionHandleClass     _mConnectionHandleClass
	MConnectionHandleClassOnce sync.Once
)

func getmConnectionHandleClass() _mConnectionHandleClass {
	MConnectionHandleClassOnce.Do(func() {
		MConnectionHandleClass = _mConnectionHandleClass{objc.GetClass("mConnectionHandle")}
	})
	return MConnectionHandleClass
}

type _mConnectionHandleClass struct {
	class objc.Class
}

// An interface definition for the [mConnectionHandle] class.
type ImConnectionHandle interface {
	objectivec.IObject
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/mConnectionHandle
type mConnectionHandle struct {
	objectivec.Object
}

// mConnectionHandleFrom constructs a [mConnectionHandle] from an unsafe.Pointer.
func mConnectionHandleFrom(ptr unsafe.Pointer) mConnectionHandle {
	return mConnectionHandle{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mConnectionHandleClass) Alloc() mConnectionHandle {
	rv := objc.Send[mConnectionHandle](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mConnectionHandleClass) New() mConnectionHandle {
	rv := objc.Send[mConnectionHandle](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mConnectionHandle) Init() mConnectionHandle {
	rv := objc.Send[mConnectionHandle](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mConnectionHandle) Autorelease() mConnectionHandle {
	rv := objc.Send[mConnectionHandle](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmConnectionHandle creates a new mConnectionHandle instance.
func NewmConnectionHandle() mConnectionHandle {
	return getmConnectionHandleClass().New()
}




