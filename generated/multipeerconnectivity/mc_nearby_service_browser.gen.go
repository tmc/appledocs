// Code generated from Apple documentation for MultipeerConnectivity. DO NOT EDIT.

package multipeerconnectivity

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MCNearbyServiceBrowser */


/* debug [class_header]: Header for MCNearbyServiceBrowser */
// The class instance for the [MCNearbyServiceBrowser] class.
var (
	MCNearbyServiceBrowserClass     _MCNearbyServiceBrowserClass
	MCNearbyServiceBrowserClassOnce sync.Once
)

func getMCNearbyServiceBrowserClass() _MCNearbyServiceBrowserClass {
	MCNearbyServiceBrowserClassOnce.Do(func() {
		MCNearbyServiceBrowserClass = _MCNearbyServiceBrowserClass{objc.GetClass("MCNearbyServiceBrowser")}
	})
	return MCNearbyServiceBrowserClass
}

type _MCNearbyServiceBrowserClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MCNearbyServiceBrowser */
// An interface definition for the [MCNearbyServiceBrowser] class.
type IMCNearbyServiceBrowser interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MCNearbyServiceBrowser */
	// properties:
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	MyPeerID() IMCPeerID
	ServiceType() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MCNearbyServiceBrowser */
	// methods:
	InvitePeerToSessionWithContextTimeout(peerID IMCPeerID, session IMCSession, context objc.IObject /* cross-framework: NSData */, timeout float64)
	StartBrowsingForPeers()
	StopBrowsingForPeers()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MCNearbyServiceBrowser */
// Alloc allocates a new instance without initialization.
func (mc _MCNearbyServiceBrowserClass) Alloc() MCNearbyServiceBrowser {
	rv := objc.Send[MCNearbyServiceBrowser](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MCNearbyServiceBrowserClass) New() MCNearbyServiceBrowser {
	rv := objc.Send[MCNearbyServiceBrowser](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MCNearbyServiceBrowser) Init() MCNearbyServiceBrowser {
	rv := objc.Send[MCNearbyServiceBrowser](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MCNearbyServiceBrowser) Autorelease() MCNearbyServiceBrowser {
	rv := objc.Send[MCNearbyServiceBrowser](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMCNearbyServiceBrowser creates a new MCNearbyServiceBrowser instance.
func NewMCNearbyServiceBrowser() MCNearbyServiceBrowser {
	return getMCNearbyServiceBrowserClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MCNearbyServiceBrowser */
// Searches (by service type) for services offered by nearby devices using infrastructure Wi-Fi, peer-to-peer Wi-Fi, and Bluetooth (in iOS) or Ethernet (in macOS and tvOS), and provides the ability to easily invite those devices to a Multipeer Connectivity session ( ).


// Searches (by service type) for services offered by nearby devices using infrastructure Wi-Fi, peer-to-peer Wi-Fi, and Bluetooth (in iOS) or Ethernet (in macOS and tvOS), and provides the ability to easily invite those devices to a Multipeer Connectivity session ( ).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCNearbyServiceBrowser
type MCNearbyServiceBrowser struct {
	objectivec.Object
}

// MCNearbyServiceBrowserFrom constructs a [MCNearbyServiceBrowser] from an unsafe.Pointer.
//
// Searches (by service type) for services offered by nearby devices using infrastructure Wi-Fi, peer-to-peer Wi-Fi, and Bluetooth (in iOS) or Ethernet (in macOS and tvOS), and provides the ability to easily invite those devices to a Multipeer Connectivity session ( ).
func MCNearbyServiceBrowserFrom(ptr unsafe.Pointer) MCNearbyServiceBrowser {
	return MCNearbyServiceBrowser{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MCNearbyServiceBrowser */

// Initializes the nearby service browser object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCNearbyServiceBrowser/init(peer:serviceType:)
func NewMCNearbyServiceBrowserWithPeerServiceType(myPeerID IMCPeerID, serviceType objc.IObject /* cross-framework: NSString */) MCNearbyServiceBrowser {
	instance := getMCNearbyServiceBrowserClass().Alloc()
	rv := objc.Send[MCNearbyServiceBrowser](instance.ID, objc.Sel("initWithPeer:serviceType:"), myPeerID, serviceType)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMCNearbyServiceBrowserWithPeerServiceType */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MCNearbyServiceBrowser */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MCNearbyServiceBrowser */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MCNearbyServiceBrowser */

// Invites a discovered peer to join a Multipeer Connectivity session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCNearbyServiceBrowser/invitePeer(_:to:withContext:timeout:)
func (m_ MCNearbyServiceBrowser) InvitePeerToSessionWithContextTimeout(peerID IMCPeerID, session IMCSession, context objc.IObject /* cross-framework: NSData */, timeout float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("invitePeer:toSession:withContext:timeout:"), peerID, session, context, timeout)
}/* debug [instance_methods/method]: InvitePeerToSessionWithContextTimeout */


// Starts browsing for peers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCNearbyServiceBrowser/startBrowsingForPeers()
func (m_ MCNearbyServiceBrowser) StartBrowsingForPeers() {
	objc.Send[objc.ID](m_.ID, objc.Sel("startBrowsingForPeers"))
}/* debug [instance_methods/method]: StartBrowsingForPeers */


// Stops browsing for peers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCNearbyServiceBrowser/stopBrowsingForPeers()
func (m_ MCNearbyServiceBrowser) StopBrowsingForPeers() {
	objc.Send[objc.ID](m_.ID, objc.Sel("stopBrowsingForPeers"))
}/* debug [instance_methods/method]: StopBrowsingForPeers */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MCNearbyServiceBrowser */

// The delegate object that handles browser-related events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCNearbyServiceBrowser/delegate
func (m_ MCNearbyServiceBrowser) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The delegate object that handles browser-related events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCNearbyServiceBrowser/delegate
func (m_ MCNearbyServiceBrowser) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// The local peer ID for this instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCNearbyServiceBrowser/myPeerID
func (m_ MCNearbyServiceBrowser) MyPeerID() IMCPeerID {
	rv := objc.Send[MCPeerID](m_.ID, objc.Sel("myPeerID"))
	return rv
}/* debug [instance_properties/getter]: myPeerID */


// The service type to browse for.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCNearbyServiceBrowser/serviceType
func (m_ MCNearbyServiceBrowser) ServiceType() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("serviceType"))
	return rv
}/* debug [instance_properties/getter]: serviceType */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MCNearbyServiceBrowser */


