// Code generated from Apple documentation for MultipeerConnectivity. DO NOT EDIT.

package multipeerconnectivity

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MCSession] class.
var (
	MCSessionClass     _MCSessionClass
	MCSessionClassOnce sync.Once
)

func getMCSessionClass() _MCSessionClass {
	MCSessionClassOnce.Do(func() {
		MCSessionClass = _MCSessionClass{objc.GetClass("MCSession")}
	})
	return MCSessionClass
}

type _MCSessionClass struct {
	class objc.Class
}

// An interface definition for the [MCSession] class.
type IMCSession interface {
	objectivec.IObject
	CancelConnectPeer(peerID unsafe.Pointer)
	ConnectPeerWithNearbyConnectionData(peerID unsafe.Pointer, data unsafe.Pointer)
	Disconnect()
	NearbyConnectionDataForPeerWithCompletionHandler(peerID unsafe.Pointer, completionHandler unsafe.Pointer)
	SendDataToPeersWithModeError(data unsafe.Pointer, peerIDs unsafe.Pointer, mode unsafe.Pointer, error_ unsafe.Pointer) bool
	SendResourceAtURLWithNameToPeerWithCompletionHandler(resourceURL unsafe.Pointer, resourceName string, peerID unsafe.Pointer, completionHandler unsafe.Pointer) unsafe.Pointer
	StartStreamWithNameToPeerError(streamName string, peerID unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer
}

// An object enables and manages communication among all peers in a Multipeer Connectivity session.
//
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCSession
type MCSession struct {
	objectivec.Object
}

// MCSessionFrom constructs a [MCSession] from an unsafe.Pointer.
//
// An object enables and manages communication among all peers in a Multipeer Connectivity session.
func MCSessionFrom(ptr unsafe.Pointer) MCSession {
	return MCSession{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MCSessionClass) Alloc() MCSession {
	rv := objc.Send[MCSession](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MCSessionClass) New() MCSession {
	rv := objc.Send[MCSession](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MCSession) Init() MCSession {
	rv := objc.Send[MCSession](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MCSession) Autorelease() MCSession {
	rv := objc.Send[MCSession](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMCSession creates a new MCSession instance.
func NewMCSession() MCSession {
	return getMCSessionClass().New()
}




// Creates a Multipeer Connectivity session.
//
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCSession/init(peer:)
func NewMCSessionWithPeer(myPeerID unsafe.Pointer) MCSession {
	instance := getMCSessionClass().Alloc()
	rv := objc.Send[MCSession](instance.ID, objc.Sel("initWithPeer:"), myPeerID)
	rv.Autorelease()
	return rv
}



// Creates a Multipeer Connectivity session, providing security information.
//
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCSession/init(peer:securityIdentity:encryptionPreference:)
func NewMCSessionWithPeerSecurityIdentityEncryptionPreference(myPeerID unsafe.Pointer, identity objc.ID, encryptionPreference unsafe.Pointer) MCSession {
	instance := getMCSessionClass().Alloc()
	rv := objc.Send[MCSession](instance.ID, objc.Sel("initWithPeer:securityIdentity:encryptionPreference:"), myPeerID, identity, encryptionPreference)
	rv.Autorelease()
	return rv
}


// Cancels an attempt to connect to a peer.
//
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCSession/cancelConnectPeer(_:)
func (m_ MCSession) CancelConnectPeer(peerID unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("cancelConnectPeer:"), peerID)
}

// Call this method to connect a peer to the session when using your own service discovery code instead of an or object.
//
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCSession/connectPeer(_:withNearbyConnectionData:)
func (m_ MCSession) ConnectPeerWithNearbyConnectionData(peerID unsafe.Pointer, data unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("connectPeer:withNearbyConnectionData:"), peerID, data)
}

// Disconnects the local peer from the session.
//
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCSession/disconnect()
func (m_ MCSession) Disconnect() {
	objc.Send[objc.ID](m_.ID, objc.Sel("disconnect"))
}

// Obtains connection data for the specified peer.
//
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCSession/nearbyConnectionData(forPeer:withCompletionHandler:)
func (m_ MCSession) NearbyConnectionDataForPeerWithCompletionHandler(peerID unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("nearbyConnectionDataForPeer:withCompletionHandler:"), peerID, completionHandler)
}

// Sends a message to nearby peers.
//
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCSession/send(_:toPeers:with:)
func (m_ MCSession) SendDataToPeersWithModeError(data unsafe.Pointer, peerIDs unsafe.Pointer, mode unsafe.Pointer, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("sendData:toPeers:withMode:error:"), data, peerIDs, mode, error_)
	return rv
}

// Sends the contents of a URL to a peer.
//
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCSession/sendResource(at:withName:toPeer:withCompletionHandler:)
func (m_ MCSession) SendResourceAtURLWithNameToPeerWithCompletionHandler(resourceURL unsafe.Pointer, resourceName string, peerID unsafe.Pointer, completionHandler unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("sendResourceAtURL:withName:toPeer:withCompletionHandler:"), resourceURL, objc.String(resourceName), peerID, completionHandler)
	return rv
}

// Opens a byte stream to a nearby peer.
//
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCSession/startStream(withName:toPeer:)
func (m_ MCSession) StartStreamWithNameToPeerError(streamName string, peerID unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("startStreamWithName:toPeer:error:"), objc.String(streamName), peerID, error_)
	return rv
}

// An array of all peers that are currently connected to this session.
//
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCSession/connectedPeers
func (m_ MCSession) ConnectedPeers() []MCPeerID {
	rv := objc.Send[[]MCPeerID](m_.ID, objc.Sel("connectedPeers"))
	return rv
}

// The delegate object that handles session-related events.
//
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCSession/delegate
func (m_ MCSession) Delegate() objc.ID {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The delegate object that handles session-related events.

//
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCSession/delegate
func (m_ MCSession) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDelegate:"), value)
}

// A value indicating whether the connection prefers encrypted connections, unencrypted connections, or has no preference.
//
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCSession/encryptionPreference
func (m_ MCSession) EncryptionPreference() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("encryptionPreference"))
	return rv
}

// A local identifier that represents the device on which your app is currently running.
//
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCSession/myPeerID
func (m_ MCSession) MyPeerID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("myPeerID"))
	return rv
}

// The security identity of the local peer.
//
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCSession/securityIdentity
func (m_ MCSession) SecurityIdentity() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("securityIdentity"))
	return rv
}


