// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NEHotspotHelperCommand] class.
var (
	NEHotspotHelperCommandClass     _NEHotspotHelperCommandClass
	NEHotspotHelperCommandClassOnce sync.Once
)

func getNEHotspotHelperCommandClass() _NEHotspotHelperCommandClass {
	NEHotspotHelperCommandClassOnce.Do(func() {
		NEHotspotHelperCommandClass = _NEHotspotHelperCommandClass{objc.GetClass("NEHotspotHelperCommand")}
	})
	return NEHotspotHelperCommandClass
}

type _NEHotspotHelperCommandClass struct {
	class objc.Class
}

// An interface definition for the [NEHotspotHelperCommand] class.
type INEHotspotHelperCommand interface {
	objectivec.IObject
}

// A command for the hotspot helper to handle.
//
// NEHotspostHelperCommand objects are passed to the the Hotspot Helper app’s command handler block. The Hotspot Helper app processes the command, instantiates an object, sets the annotated or ( or commands only), and then delivers the response to the system.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotHelperCommand
type NEHotspotHelperCommand struct {
	objectivec.Object
}

// NEHotspotHelperCommandFrom constructs a [NEHotspotHelperCommand] from an unsafe.Pointer.
//
// A command for the hotspot helper to handle.
func NEHotspotHelperCommandFrom(ptr unsafe.Pointer) NEHotspotHelperCommand {
	return NEHotspotHelperCommand{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (nc _NEHotspotHelperCommandClass) Alloc() NEHotspotHelperCommand {
	rv := objc.Send[NEHotspotHelperCommand](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NEHotspotHelperCommandClass) New() NEHotspotHelperCommand {
	rv := objc.Send[NEHotspotHelperCommand](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEHotspotHelperCommand) Init() NEHotspotHelperCommand {
	rv := objc.Send[NEHotspotHelperCommand](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEHotspotHelperCommand) Autorelease() NEHotspotHelperCommand {
	rv := objc.Send[NEHotspotHelperCommand](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEHotspotHelperCommand creates a new NEHotspotHelperCommand instance.
func NewNEHotspotHelperCommand() NEHotspotHelperCommand {
	return getNEHotspotHelperCommandClass().New()
}


// The type of the command
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspothelpercommand/commandtype
func (n_ NEHotspotHelperCommand) CommandType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("commandType"))
	return rv
}


// SetCommandType sets the value of the commandType property.
// The type of the command

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspothelpercommand/commandtype
func (n_ NEHotspotHelperCommand) SetCommandType(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setCommandType:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspothelpercommand/interface-46dq
func (n_ NEHotspotHelperCommand) Interface() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("interface"))
	return rv
}


// SetInterface sets the value of the interface property.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspothelpercommand/interface-46dq
func (n_ NEHotspotHelperCommand) SetInterface(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setInterface:"), value)
}

// The network associated with the command.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspothelpercommand/network
func (n_ NEHotspotHelperCommand) Network() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("network"))
	return rv
}


// SetNetwork sets the value of the network property.
// The network associated with the command.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspothelpercommand/network
func (n_ NEHotspotHelperCommand) SetNetwork(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setNetwork:"), value)
}

// The list of networks associated with the command.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspothelpercommand/networklist
func (n_ NEHotspotHelperCommand) NetworkList() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("networkList"))
	return rv
}


// SetNetworkList sets the value of the networkList property.
// The list of networks associated with the command.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspothelpercommand/networklist
func (n_ NEHotspotHelperCommand) SetNetworkList(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setNetworkList:"), value)
}



