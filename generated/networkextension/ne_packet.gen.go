// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NEPacket] class.
var (
	NEPacketClass     _NEPacketClass
	NEPacketClassOnce sync.Once
)

func getNEPacketClass() _NEPacketClass {
	NEPacketClassOnce.Do(func() {
		NEPacketClass = _NEPacketClass{objc.GetClass("NEPacket")}
	})
	return NEPacketClass
}

type _NEPacketClass struct {
	class objc.Class
}

// An interface definition for the [NEPacket] class.
type INEPacket interface {
	objectivec.IObject
	Data() foundation.NSData
	Direction() NETrafficDirection
	Metadata() NEFlowMetaData
	ProtocolFamily() unsafe.Pointer
}

// A network packet and its associated properties.
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

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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
	return getNEPacketClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEPacket/init(data:protocolFamily:)
func NewNEPacketWithDataProtocolFamily(data foundation.IData, protocolFamily unsafe.Pointer) NEPacket {
	instance := getNEPacketClass().Alloc()
	rv := objc.Send[NEPacket](instance.ID, objc.Sel("initWithData:protocolFamily:"), data, protocolFamily)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEPacket/data
func (n_ NEPacket) Data() foundation.NSData {
	rv := objc.Send[foundation.NSData](n_.ID, objc.Sel("data"))
	return rv
}

// The direction of the packet.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEPacket/direction
func (n_ NEPacket) Direction() NETrafficDirection {
	rv := objc.Send[NETrafficDirection](n_.ID, objc.Sel("direction"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEPacket/metadata
func (n_ NEPacket) Metadata() NEFlowMetaData {
	rv := objc.Send[NEFlowMetaData](n_.ID, objc.Sel("metadata"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEPacket/protocolFamily
func (n_ NEPacket) ProtocolFamily() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("protocolFamily"))
	return rv
}


