// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NEPacket] class.
var nEPacketClass = _NEPacketClass{objc.GetClass("NEPacket")}

type _NEPacketClass struct {
	class objc.Class
}

// An interface definition for the [NEPacket] class.
type INEPacket interface {
	objectivec.IObject
}

// A network packet and its associated properties. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEPacket

type NEPacket struct {
	objectivec.Object
}

// NEPacketFrom constructs a [NEPacket] from an unsafe.Pointer.
//
// A network packet and its associated properties.
func NEPacketFrom(ptr unsafe.Pointer) NEPacket {
	return NEPacket{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (nc _NEPacketClass) Alloc() NEPacket {
	rv := objc.Send[NEPacket](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (nc _NEPacketClass) New() NEPacket {
	rv := objc.Send[NEPacket](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEPacket) Init() NEPacket {
	rv := objc.Send[NEPacket](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEPacket) Autorelease() NEPacket {
	rv := objc.Send[NEPacket](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEPacket creates a new NEPacket instance.
func NewNEPacket() NEPacket {
	return nEPacketClass.New()
}




