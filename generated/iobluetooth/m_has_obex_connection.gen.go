// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mHasOBEXConnection] class.
var (
	MHasOBEXConnectionClass     _mHasOBEXConnectionClass
	MHasOBEXConnectionClassOnce sync.Once
)

func getmHasOBEXConnectionClass() _mHasOBEXConnectionClass {
	MHasOBEXConnectionClassOnce.Do(func() {
		MHasOBEXConnectionClass = _mHasOBEXConnectionClass{objc.GetClass("mHasOBEXConnection")}
	})
	return MHasOBEXConnectionClass
}

type _mHasOBEXConnectionClass struct {
	class objc.Class
}

// An interface definition for the [mHasOBEXConnection] class.
type ImHasOBEXConnection interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSession/mHasOBEXConnection
type mHasOBEXConnection struct {
	objectivec.Object
}

// mHasOBEXConnectionFrom constructs a [mHasOBEXConnection] from an unsafe.Pointer.
func mHasOBEXConnectionFrom(ptr unsafe.Pointer) mHasOBEXConnection {
	return mHasOBEXConnection{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mHasOBEXConnectionClass) Alloc() mHasOBEXConnection {
	rv := objc.Send[mHasOBEXConnection](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mHasOBEXConnectionClass) New() mHasOBEXConnection {
	rv := objc.Send[mHasOBEXConnection](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mHasOBEXConnection) Init() mHasOBEXConnection {
	rv := objc.Send[mHasOBEXConnection](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mHasOBEXConnection) Autorelease() mHasOBEXConnection {
	rv := objc.Send[mHasOBEXConnection](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmHasOBEXConnection creates a new mHasOBEXConnection instance.
func NewmHasOBEXConnection() mHasOBEXConnection {
	return getmHasOBEXConnectionClass().New()
}




