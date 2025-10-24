// Code generated from Apple documentation for MultipeerConnectivity. DO NOT EDIT.

package multipeerconnectivity

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MCPeerID */


/* debug [class_header]: Header for MCPeerID */
// The class instance for the [MCPeerID] class.
var (
	MCPeerIDClass     _MCPeerIDClass
	MCPeerIDClassOnce sync.Once
)

func getMCPeerIDClass() _MCPeerIDClass {
	MCPeerIDClassOnce.Do(func() {
		MCPeerIDClass = _MCPeerIDClass{objc.GetClass("MCPeerID")}
	})
	return MCPeerIDClass
}

type _MCPeerIDClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MCPeerID */
// An interface definition for the [MCPeerID] class.
type IMCPeerID interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MCPeerID */
	// properties:
	DisplayName() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MCPeerID */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MCPeerID */
// Alloc allocates a new instance without initialization.
func (mc _MCPeerIDClass) Alloc() MCPeerID {
	rv := objc.Send[MCPeerID](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MCPeerIDClass) New() MCPeerID {
	rv := objc.Send[MCPeerID](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MCPeerID) Init() MCPeerID {
	rv := objc.Send[MCPeerID](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MCPeerID) Autorelease() MCPeerID {
	rv := objc.Send[MCPeerID](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMCPeerID creates a new MCPeerID instance.
func NewMCPeerID() MCPeerID {
	return getMCPeerIDClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MCPeerID */
// An object represents a peer in a multipeer session.
//
// You create a single peer ID object that represents the instance of your app running on the local device. The Multipeer Connectivity framework is responsible for creating peer ID objects that represent other devices. To create a new peer ID for the local app and associate a display name with that ID, call . The peer’s name must be no longer than 63 bytes in UTF-8 encoding. Each peer ID your app creates with is unique, even when supplying the same display name. If you want a device’s peer ID to be stable over time, don’t create a new peer ID every time your app begins advertising or browsing. Instead, archive the ID when you create it, and then unarchive it the next time you need it. If you need the peer ID to be tied to the display name, you can archive the name as well, and only create a new peer ID when the name changes, as illustrated in the following code fragment:


// An object represents a peer in a multipeer session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCPeerID
type MCPeerID struct {
	objectivec.Object
}

// MCPeerIDFrom constructs a [MCPeerID] from an unsafe.Pointer.
//
// An object represents a peer in a multipeer session.
func MCPeerIDFrom(ptr unsafe.Pointer) MCPeerID {
	return MCPeerID{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MCPeerID */

// Initializes a peer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCPeerID/init(displayName:)
func NewMCPeerIDWithDisplayName(myDisplayName objc.IObject /* cross-framework: NSString */) MCPeerID {
	instance := getMCPeerIDClass().Alloc()
	rv := objc.Send[MCPeerID](instance.ID, objc.Sel("initWithDisplayName:"), myDisplayName)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMCPeerIDWithDisplayName */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MCPeerID */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MCPeerID */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MCPeerID */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MCPeerID */

// The display name for this peer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCPeerID/displayName
func (m_ MCPeerID) DisplayName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("displayName"))
	return rv
}/* debug [instance_properties/getter]: displayName */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MCPeerID */


