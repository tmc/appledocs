// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mIOConnection] class.
var (
	MIOConnectionClass     _mIOConnectionClass
	MIOConnectionClassOnce sync.Once
)

func getmIOConnectionClass() _mIOConnectionClass {
	MIOConnectionClassOnce.Do(func() {
		MIOConnectionClass = _mIOConnectionClass{objc.GetClass("mIOConnection")}
	})
	return MIOConnectionClass
}

type _mIOConnectionClass struct {
	class objc.Class
}

// An interface definition for the [mIOConnection] class.
type ImIOConnection interface {
	objectivec.IObject
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothObject/mIOConnection
type mIOConnection struct {
	objectivec.Object
}

// mIOConnectionFrom constructs a [mIOConnection] from an unsafe.Pointer.
func mIOConnectionFrom(ptr unsafe.Pointer) mIOConnection {
	return mIOConnection{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mIOConnectionClass) Alloc() mIOConnection {
	rv := objc.Send[mIOConnection](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mIOConnectionClass) New() mIOConnection {
	rv := objc.Send[mIOConnection](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mIOConnection) Init() mIOConnection {
	rv := objc.Send[mIOConnection](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mIOConnection) Autorelease() mIOConnection {
	rv := objc.Send[mIOConnection](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmIOConnection creates a new mIOConnection instance.
func NewmIOConnection() mIOConnection {
	return getmIOConnectionClass().New()
}




