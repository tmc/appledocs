// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [NEFilterPacketProvider] class.
var (
	NEFilterPacketProviderClass     _NEFilterPacketProviderClass
	NEFilterPacketProviderClassOnce sync.Once
)

func getNEFilterPacketProviderClass() _NEFilterPacketProviderClass {
	NEFilterPacketProviderClassOnce.Do(func() {
		NEFilterPacketProviderClass = _NEFilterPacketProviderClass{objc.GetClass("NEFilterPacketProvider")}
	})
	return NEFilterPacketProviderClass
}

type _NEFilterPacketProviderClass struct {
	class objc.Class
}





// An interface definition for the [NEFilterPacketProvider] class.
type INEFilterPacketProvider interface {
	INEFilterProvider
	

	// properties:
	PacketHandler() objectivec.IObject
	SetPacketHandler(value objectivec.IObject)
	Handler() objectivec.IObject
	SetHandler(value objectivec.IObject)


	

	// methods:
	AllowPacket(packet INEPacket)
	DelayCurrentPacket(context INEFilterPacketContext) INEPacket


}





// Alloc allocates a new instance without initialization.
func (nc _NEFilterPacketProviderClass) Alloc() NEFilterPacketProvider {
	rv := objc.Send[NEFilterPacketProvider](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NEFilterPacketProviderClass) New() NEFilterPacketProvider {
	rv := objc.Send[NEFilterPacketProvider](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEFilterPacketProvider) Init() NEFilterPacketProvider {
	rv := objc.Send[NEFilterPacketProvider](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEFilterPacketProvider) Autorelease() NEFilterPacketProvider {
	rv := objc.Send[NEFilterPacketProvider](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEFilterPacketProvider creates a new NEFilterPacketProvider instance.
func NewNEFilterPacketProvider() NEFilterPacketProvider {
	return getNEFilterPacketProviderClass().New()
}





// A filter provider that evaluates network packets and decides whether to block, allow, or delay the packets.


// A filter provider that evaluates network packets and decides whether to block, allow, or delay the packets.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterPacketProvider
type NEFilterPacketProvider struct {
	NEFilterProvider
}

// NEFilterPacketProviderFrom constructs a [NEFilterPacketProvider] from an unsafe.Pointer.
//
// A filter provider that evaluates network packets and decides whether to block, allow, or delay the packets.
func NEFilterPacketProviderFrom(ptr unsafe.Pointer) NEFilterPacketProvider {
	return NEFilterPacketProvider{
		NEFilterProvider: NEFilterProviderFrom(ptr),
	}
}




















// Allow delivery of a previously-delayed packet.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterPacketProvider/allow(_:)
func (n_ NEFilterPacketProvider) AllowPacket(packet INEPacket) {
	objc.Send[objc.ID](n_.ID, objc.Sel("allowPacket:"), packet)
}


// Delay a packet currently processed by a packet handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterPacketProvider/delayCurrentPacket(_:)
func (n_ NEFilterPacketProvider) DelayCurrentPacket(context INEFilterPacketContext) INEPacket {
	rv := objc.Send[NEPacket](n_.ID, objc.Sel("delayCurrentPacket:"), context)
	return rv
}







// A Swift closure or an ObjectiveC block that handles each packet received by the filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterPacketProvider/packetHandler
func (n_ NEFilterPacketProvider) PacketHandler() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](n_.ID, objc.Sel("packetHandler"))
	return rv
}


// A Swift closure or an ObjectiveC block that handles each packet received by the filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterPacketProvider/packetHandler
func (n_ NEFilterPacketProvider) SetPacketHandler(value objectivec.IObject) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setPacketHandler:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterpacketprovider/handler
func (n_ NEFilterPacketProvider) Handler() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](n_.ID, objc.Sel("handler"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterpacketprovider/handler
func (n_ NEFilterPacketProvider) SetHandler(value objectivec.IObject) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setHandler:"), value)
}








