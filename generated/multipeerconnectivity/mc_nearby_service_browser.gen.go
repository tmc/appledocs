// Code generated from Apple documentation for MultipeerConnectivity. DO NOT EDIT.

package multipeerconnectivity

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [MCNearbyServiceBrowser] class.
type IMCNearbyServiceBrowser interface {
	objectivec.IObject
	InvitePeerToSessionWithContextTimeout(peerID IMCPeerID, session IMCSession, context foundation.IData, timeout foundation.ITimeInterval)
	StartBrowsingForPeers()
	StopBrowsingForPeers()
	Delegate() objc.ID
	SetDelegate(value objc.ID)
	MyPeerID() MCPeerID
	ServiceType() string
}

// Searches (by service type) for services offered by nearby devices using infrastructure Wi-Fi, peer-to-peer Wi-Fi, and Bluetooth (in iOS) or Ethernet (in macOS and tvOS), and provides the ability to easily invite those devices to a Multipeer Connectivity session ( ).
//
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

// Alloc allocates a new instance without initialization.
func (mc _MCNearbyServiceBrowserClass) Alloc() MCNearbyServiceBrowser {
	rv := objc.Send[MCNearbyServiceBrowser](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Initializes the nearby service browser object.
//
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCNearbyServiceBrowser/init(peer:serviceType:)
func NewMCNearbyServiceBrowserWithPeerServiceType(myPeerID IMCPeerID, serviceType string) MCNearbyServiceBrowser {
	instance := getMCNearbyServiceBrowserClass().Alloc()
	rv := objc.Send[MCNearbyServiceBrowser](instance.ID, objc.Sel("initWithPeer:serviceType:"), myPeerID, objc.String(serviceType))
	rv.Autorelease()
	return rv
}


// Invites a discovered peer to join a Multipeer Connectivity session.
//
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCNearbyServiceBrowser/invitePeer(_:to:withContext:timeout:)
func (m_ MCNearbyServiceBrowser) InvitePeerToSessionWithContextTimeout(peerID IMCPeerID, session IMCSession, context foundation.IData, timeout foundation.ITimeInterval) {
	objc.Send[objc.ID](m_.ID, objc.Sel("invitePeer:toSession:withContext:timeout:"), peerID, session, context, timeout)
}

// Starts browsing for peers.
//
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCNearbyServiceBrowser/startBrowsingForPeers()
func (m_ MCNearbyServiceBrowser) StartBrowsingForPeers() {
	objc.Send[objc.ID](m_.ID, objc.Sel("startBrowsingForPeers"))
}

// Stops browsing for peers.
//
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCNearbyServiceBrowser/stopBrowsingForPeers()
func (m_ MCNearbyServiceBrowser) StopBrowsingForPeers() {
	objc.Send[objc.ID](m_.ID, objc.Sel("stopBrowsingForPeers"))
}

// The delegate object that handles browser-related events.
//
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCNearbyServiceBrowser/delegate
func (m_ MCNearbyServiceBrowser) Delegate() objc.ID {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The delegate object that handles browser-related events.

//
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCNearbyServiceBrowser/delegate
func (m_ MCNearbyServiceBrowser) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDelegate:"), value)
}

// The local peer ID for this instance.
//
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCNearbyServiceBrowser/myPeerID
func (m_ MCNearbyServiceBrowser) MyPeerID() MCPeerID {
	rv := objc.Send[MCPeerID](m_.ID, objc.Sel("myPeerID"))
	return rv
}

// The service type to browse for.
//
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCNearbyServiceBrowser/serviceType
func (m_ MCNearbyServiceBrowser) ServiceType() string {
	rv := objc.Send[string](m_.ID, objc.Sel("serviceType"))
	return rv
}


