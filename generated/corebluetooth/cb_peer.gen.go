// Code generated from Apple documentation for CoreBluetooth. DO NOT EDIT.

package corebluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CBPeer */


/* debug [class_header]: Header for CBPeer */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CBPeer */
// An interface definition for the [CBPeer] class.
type ICBPeer interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CBPeer */
	// properties:
	Identifier() foundation.UUID
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CBPeer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CBPeer */
// Alloc allocates a new instance without initialization.
func (cc _CBPeerClass) Alloc() CBPeer {
	rv := objc.Send[CBPeer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CBPeer */
// An object that represents a remote device.
//
// The class is an abstract base class that defines common behavior for objects representing remote devices. You typically don’t create instances of either or its concrete subclasses. Instead, the system creates them for you during the process of peer discovery. Your app takes the role of either a central (by creating an instance of ) or a peripheral (by creating an instance of ), and interacts through the manager with remote devices in the opposite role. During the process of peer discovery, where a central device scans for peripherals advertising services, the system creates objects from the concrete subclasses of to represent discovered remote devices. The concrete subclasses of are and .


// An object that represents a remote device.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CBPeer *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CBPeer */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CBPeer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CBPeer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CBPeer */

// The UUID associated with the peer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBPeer/identifier
func (c_ CBPeer) Identifier() foundation.UUID {
	rv := objc.Send[foundation.UUID](c_.ID, objc.Sel("identifier"))
	return rv
}/* debug [instance_properties/getter]: identifier */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CBPeer */



