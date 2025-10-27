// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [NEFilterPacketContext] class.
var (
	NEFilterPacketContextClass     _NEFilterPacketContextClass
	NEFilterPacketContextClassOnce sync.Once
)

func getNEFilterPacketContextClass() _NEFilterPacketContextClass {
	NEFilterPacketContextClassOnce.Do(func() {
		NEFilterPacketContextClass = _NEFilterPacketContextClass{objc.GetClass("NEFilterPacketContext")}
	})
	return NEFilterPacketContextClass
}

type _NEFilterPacketContextClass struct {
	class objc.Class
}





// An interface definition for the [NEFilterPacketContext] class.
type INEFilterPacketContext interface {
	objectivec.IObject
	

	// properties:
	PacketHandler() objectivec.IObject
	SetPacketHandler(value objectivec.IObject)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (nc _NEFilterPacketContextClass) Alloc() NEFilterPacketContext {
	rv := objc.Send[NEFilterPacketContext](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NEFilterPacketContextClass) New() NEFilterPacketContext {
	rv := objc.Send[NEFilterPacketContext](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEFilterPacketContext) Init() NEFilterPacketContext {
	rv := objc.Send[NEFilterPacketContext](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEFilterPacketContext) Autorelease() NEFilterPacketContext {
	rv := objc.Send[NEFilterPacketContext](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEFilterPacketContext creates a new NEFilterPacketContext instance.
func NewNEFilterPacketContext() NEFilterPacketContext {
	return getNEFilterPacketContextClass().New()
}





// The context object provided to the filter packet handler.


// The context object provided to the filter packet handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterPacketContext
type NEFilterPacketContext struct {
	objectivec.Object
}

// NEFilterPacketContextFrom constructs a [NEFilterPacketContext] from an unsafe.Pointer.
//
// The context object provided to the filter packet handler.
func NEFilterPacketContextFrom(ptr unsafe.Pointer) NEFilterPacketContext {
	return NEFilterPacketContext{objectivec.Object{objc.ID(ptr)}}
}

























// A Swift closure or an ObjectiveC block that handles each packet received by the filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterpacketprovider/packethandler
func (n_ NEFilterPacketContext) PacketHandler() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](n_.ID, objc.Sel("packetHandler"))
	return rv
}


// A Swift closure or an ObjectiveC block that handles each packet received by the filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterpacketprovider/packethandler
func (n_ NEFilterPacketContext) SetPacketHandler(value objectivec.IObject) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setPacketHandler:"), value)
}








