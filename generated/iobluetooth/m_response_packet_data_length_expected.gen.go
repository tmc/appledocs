// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mResponsePacketDataLengthExpected] class.
var (
	MResponsePacketDataLengthExpectedClass     _mResponsePacketDataLengthExpectedClass
	MResponsePacketDataLengthExpectedClassOnce sync.Once
)

func getmResponsePacketDataLengthExpectedClass() _mResponsePacketDataLengthExpectedClass {
	MResponsePacketDataLengthExpectedClassOnce.Do(func() {
		MResponsePacketDataLengthExpectedClass = _mResponsePacketDataLengthExpectedClass{objc.GetClass("mResponsePacketDataLengthExpected")}
	})
	return MResponsePacketDataLengthExpectedClass
}

type _mResponsePacketDataLengthExpectedClass struct {
	class objc.Class
}

// An interface definition for the [mResponsePacketDataLengthExpected] class.
type ImResponsePacketDataLengthExpected interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSession/mResponsePacketDataLengthExpected
type mResponsePacketDataLengthExpected struct {
	objectivec.Object
}

// mResponsePacketDataLengthExpectedFrom constructs a [mResponsePacketDataLengthExpected] from an unsafe.Pointer.
func mResponsePacketDataLengthExpectedFrom(ptr unsafe.Pointer) mResponsePacketDataLengthExpected {
	return mResponsePacketDataLengthExpected{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mResponsePacketDataLengthExpectedClass) Alloc() mResponsePacketDataLengthExpected {
	rv := objc.Send[mResponsePacketDataLengthExpected](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mResponsePacketDataLengthExpectedClass) New() mResponsePacketDataLengthExpected {
	rv := objc.Send[mResponsePacketDataLengthExpected](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mResponsePacketDataLengthExpected) Init() mResponsePacketDataLengthExpected {
	rv := objc.Send[mResponsePacketDataLengthExpected](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mResponsePacketDataLengthExpected) Autorelease() mResponsePacketDataLengthExpected {
	rv := objc.Send[mResponsePacketDataLengthExpected](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmResponsePacketDataLengthExpected creates a new mResponsePacketDataLengthExpected instance.
func NewmResponsePacketDataLengthExpected() mResponsePacketDataLengthExpected {
	return getmResponsePacketDataLengthExpectedClass().New()
}




