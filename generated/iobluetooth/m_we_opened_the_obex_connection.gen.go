// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mWeOpenedTheOBEXConnection] class.
var (
	MWeOpenedTheOBEXConnectionClass     _mWeOpenedTheOBEXConnectionClass
	MWeOpenedTheOBEXConnectionClassOnce sync.Once
)

func getmWeOpenedTheOBEXConnectionClass() _mWeOpenedTheOBEXConnectionClass {
	MWeOpenedTheOBEXConnectionClassOnce.Do(func() {
		MWeOpenedTheOBEXConnectionClass = _mWeOpenedTheOBEXConnectionClass{objc.GetClass("mWeOpenedTheOBEXConnection")}
	})
	return MWeOpenedTheOBEXConnectionClass
}

type _mWeOpenedTheOBEXConnectionClass struct {
	class objc.Class
}

// An interface definition for the [mWeOpenedTheOBEXConnection] class.
type ImWeOpenedTheOBEXConnection interface {
	objectivec.IObject
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/mWeOpenedTheOBEXConnection
type mWeOpenedTheOBEXConnection struct {
	objectivec.Object
}

// mWeOpenedTheOBEXConnectionFrom constructs a [mWeOpenedTheOBEXConnection] from an unsafe.Pointer.
func mWeOpenedTheOBEXConnectionFrom(ptr unsafe.Pointer) mWeOpenedTheOBEXConnection {
	return mWeOpenedTheOBEXConnection{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mWeOpenedTheOBEXConnectionClass) Alloc() mWeOpenedTheOBEXConnection {
	rv := objc.Send[mWeOpenedTheOBEXConnection](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mWeOpenedTheOBEXConnectionClass) New() mWeOpenedTheOBEXConnection {
	rv := objc.Send[mWeOpenedTheOBEXConnection](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mWeOpenedTheOBEXConnection) Init() mWeOpenedTheOBEXConnection {
	rv := objc.Send[mWeOpenedTheOBEXConnection](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mWeOpenedTheOBEXConnection) Autorelease() mWeOpenedTheOBEXConnection {
	rv := objc.Send[mWeOpenedTheOBEXConnection](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmWeOpenedTheOBEXConnection creates a new mWeOpenedTheOBEXConnection instance.
func NewmWeOpenedTheOBEXConnection() mWeOpenedTheOBEXConnection {
	return getmWeOpenedTheOBEXConnectionClass().New()
}




