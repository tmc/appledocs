// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
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




