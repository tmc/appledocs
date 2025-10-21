// Code generated from Apple documentation for CoreBluetooth. DO NOT EDIT.

package corebluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CBPeer] class.
var (
	CBPeerClass     _CBPeerClass
	CBPeerClassOnce sync.Once
)

func getCBPeerClass() _CBPeerClass {
	CBPeerClassOnce.Do(func() {
		CBPeerClass = _CBPeerClass{objc.GetClass("CBPeer")}
	})
	return CBPeerClass
}

type _CBPeerClass struct {
	class objc.Class
}

// An interface definition for the [CBPeer] class.
type ICBPeer interface {
	objectivec.IObject
}

// An object that represents a remote device.
//
// The class is an abstract base class that defines common behavior for objects representing remote devices. You typically don’t create instances of either or its concrete subclasses. Instead, the system creates them for you during the process of peer discovery. Your app takes the role of either a central (by creating an instance of ) or a peripheral (by creating an instance of ), and interacts through the manager with remote devices in the opposite role. During the process of peer discovery, where a central device scans for peripherals advertising services, the system creates objects from the concrete subclasses of to represent discovered remote devices. The concrete subclasses of are and .
//
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBPeer
type CBPeer struct {
	objectivec.Object
}

// CBPeerFrom constructs a [CBPeer] from an unsafe.Pointer.
//
// An object that represents a remote device.
func CBPeerFrom(ptr unsafe.Pointer) CBPeer {
	return CBPeer{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CBPeerClass) Alloc() CBPeer {
	rv := objc.Send[CBPeer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CBPeerClass) New() CBPeer {
	rv := objc.Send[CBPeer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CBPeer) Init() CBPeer {
	rv := objc.Send[CBPeer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CBPeer) Autorelease() CBPeer {
	rv := objc.Send[CBPeer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCBPeer creates a new CBPeer instance.
func NewCBPeer() CBPeer {
	return getCBPeerClass().New()
}


// The UUID associated with the peer.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBPeer/identifier
func (c_ CBPeer) Identifier() foundation.UUID {
	rv := objc.Send[foundation.UUID](c_.ID, objc.Sel("identifier"))
	return rv
}



