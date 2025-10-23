// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mResponsePacketDataLengthSoFar] class.
var (
	MResponsePacketDataLengthSoFarClass     _mResponsePacketDataLengthSoFarClass
	MResponsePacketDataLengthSoFarClassOnce sync.Once
)

func getmResponsePacketDataLengthSoFarClass() _mResponsePacketDataLengthSoFarClass {
	MResponsePacketDataLengthSoFarClassOnce.Do(func() {
		MResponsePacketDataLengthSoFarClass = _mResponsePacketDataLengthSoFarClass{objc.GetClass("mResponsePacketDataLengthSoFar")}
	})
	return MResponsePacketDataLengthSoFarClass
}

type _mResponsePacketDataLengthSoFarClass struct {
	class objc.Class
}

// An interface definition for the [mResponsePacketDataLengthSoFar] class.
type ImResponsePacketDataLengthSoFar interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSession/mResponsePacketDataLengthSoFar
type mResponsePacketDataLengthSoFar struct {
	objectivec.Object
}

// mResponsePacketDataLengthSoFarFrom constructs a [mResponsePacketDataLengthSoFar] from an unsafe.Pointer.
func mResponsePacketDataLengthSoFarFrom(ptr unsafe.Pointer) mResponsePacketDataLengthSoFar {
	return mResponsePacketDataLengthSoFar{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mResponsePacketDataLengthSoFarClass) Alloc() mResponsePacketDataLengthSoFar {
	rv := objc.Send[mResponsePacketDataLengthSoFar](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mResponsePacketDataLengthSoFarClass) New() mResponsePacketDataLengthSoFar {
	rv := objc.Send[mResponsePacketDataLengthSoFar](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mResponsePacketDataLengthSoFar) Init() mResponsePacketDataLengthSoFar {
	rv := objc.Send[mResponsePacketDataLengthSoFar](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mResponsePacketDataLengthSoFar) Autorelease() mResponsePacketDataLengthSoFar {
	rv := objc.Send[mResponsePacketDataLengthSoFar](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmResponsePacketDataLengthSoFar creates a new mResponsePacketDataLengthSoFar instance.
func NewmResponsePacketDataLengthSoFar() mResponsePacketDataLengthSoFar {
	return getmResponsePacketDataLengthSoFarClass().New()
}




