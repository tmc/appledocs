// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mReceivePacketBuffer] class.
var (
	MReceivePacketBufferClass     _mReceivePacketBufferClass
	MReceivePacketBufferClassOnce sync.Once
)

func getmReceivePacketBufferClass() _mReceivePacketBufferClass {
	MReceivePacketBufferClassOnce.Do(func() {
		MReceivePacketBufferClass = _mReceivePacketBufferClass{objc.GetClass("mReceivePacketBuffer")}
	})
	return MReceivePacketBufferClass
}

type _mReceivePacketBufferClass struct {
	class objc.Class
}

// An interface definition for the [mReceivePacketBuffer] class.
type ImReceivePacketBuffer interface {
	objectivec.IObject
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSession/mReceivePacketBuffer
type mReceivePacketBuffer struct {
	objectivec.Object
}

// mReceivePacketBufferFrom constructs a [mReceivePacketBuffer] from an unsafe.Pointer.
func mReceivePacketBufferFrom(ptr unsafe.Pointer) mReceivePacketBuffer {
	return mReceivePacketBuffer{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mReceivePacketBufferClass) Alloc() mReceivePacketBuffer {
	rv := objc.Send[mReceivePacketBuffer](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mReceivePacketBufferClass) New() mReceivePacketBuffer {
	rv := objc.Send[mReceivePacketBuffer](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mReceivePacketBuffer) Init() mReceivePacketBuffer {
	rv := objc.Send[mReceivePacketBuffer](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mReceivePacketBuffer) Autorelease() mReceivePacketBuffer {
	rv := objc.Send[mReceivePacketBuffer](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmReceivePacketBuffer creates a new mReceivePacketBuffer instance.
func NewmReceivePacketBuffer() mReceivePacketBuffer {
	return getmReceivePacketBufferClass().New()
}




