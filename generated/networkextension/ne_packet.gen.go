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
	// properties:
	Data() objc.IObject /* cross-framework: Data */
	SetData(value objc.IObject /* cross-framework: Data */)
	Direction() unsafe.Pointer
	SetDirection(value unsafe.Pointer)
	Metadata() objc.IObject /* cross-framework: NEFlowMetaData */
	SetMetadata(value objc.IObject /* cross-framework: NEFlowMetaData */)
	ProtocolFamily() unsafe.Pointer
	SetProtocolFamily(value unsafe.Pointer)
	// methods:
}

// A network packet and its associated properties.


// A network packet and its associated properties.
//
// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nepacket/data
func (n_ NEPacket) Data() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](n_.ID, objc.Sel("data"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nepacket/data
func (n_ NEPacket) SetData(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setData:"), value)
}


// The direction of the packet.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nepacket/direction
func (n_ NEPacket) Direction() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("direction"))
	return rv
}


// The direction of the packet.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nepacket/direction
func (n_ NEPacket) SetDirection(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDirection:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nepacket/metadata
func (n_ NEPacket) Metadata() objc.IObject /* cross-framework: NEFlowMetaData */ {
	rv := objc.Send[NEFlowMetaData](n_.ID, objc.Sel("metadata"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nepacket/metadata
func (n_ NEPacket) SetMetadata(value objc.IObject /* cross-framework: NEFlowMetaData */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMetadata:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nepacket/protocolfamily
func (n_ NEPacket) ProtocolFamily() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("protocolFamily"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nepacket/protocolfamily
func (n_ NEPacket) SetProtocolFamily(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setProtocolFamily:"), value)
}



