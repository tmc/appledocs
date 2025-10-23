// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mDataAvailablePort] class.
var (
	MDataAvailablePortClass     _mDataAvailablePortClass
	MDataAvailablePortClassOnce sync.Once
)

func getmDataAvailablePortClass() _mDataAvailablePortClass {
	MDataAvailablePortClassOnce.Do(func() {
		MDataAvailablePortClass = _mDataAvailablePortClass{objc.GetClass("mDataAvailablePort")}
	})
	return MDataAvailablePortClass
}

type _mDataAvailablePortClass struct {
	class objc.Class
}

// An interface definition for the [mDataAvailablePort] class.
type ImDataAvailablePort interface {
	objectivec.IObject
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/mDataAvailablePort
type mDataAvailablePort struct {
	objectivec.Object
}

// mDataAvailablePortFrom constructs a [mDataAvailablePort] from an unsafe.Pointer.
func mDataAvailablePortFrom(ptr unsafe.Pointer) mDataAvailablePort {
	return mDataAvailablePort{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mDataAvailablePortClass) Alloc() mDataAvailablePort {
	rv := objc.Send[mDataAvailablePort](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mDataAvailablePortClass) New() mDataAvailablePort {
	rv := objc.Send[mDataAvailablePort](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mDataAvailablePort) Init() mDataAvailablePort {
	rv := objc.Send[mDataAvailablePort](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mDataAvailablePort) Autorelease() mDataAvailablePort {
	rv := objc.Send[mDataAvailablePort](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmDataAvailablePort creates a new mDataAvailablePort instance.
func NewmDataAvailablePort() mDataAvailablePort {
	return getmDataAvailablePortClass().New()
}




