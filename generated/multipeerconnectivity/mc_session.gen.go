// Code generated from Apple documentation for MultipeerConnectivity. DO NOT EDIT.

package multipeerconnectivity

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MCSession */


/* debug [class_header]: Header for MCSession */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MCSession */
// An interface definition for the [MCSession] class.
type IMCSession interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MCSession */
	// properties:
	ConnectedPeers() []MCPeerID
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	EncryptionPreference() MCEncryptionPreference
	MyPeerID() IMCPeerID
	SecurityIdentity() objc.IObject /* cross-framework: NSArray */
	MCErrorDomain() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MCSession */
	// methods:
	CancelConnectPeer(peerID IMCPeerID)
	ConnectPeerWithNearbyConnectionData(peerID IMCPeerID, data objc.IObject /* cross-framework: NSData */)
	Disconnect()
	NearbyConnectionDataForPeerWithCompletionHandler(peerID IMCPeerID, completionHandler unsafe.Pointer)
	SendDataToPeersWithModeError(data objc.IObject /* cross-framework: NSData */, peerIDs []MCPeerID, mode MCSessionSendDataMode, error_ unsafe.Pointer) bool
	SendResourceAtURLWithNameToPeerWithCompletionHandler(resourceURL objc.IObject /* cross-framework: NSURL */, resourceName objc.IObject /* cross-framework: NSString */, peerID IMCPeerID, completionHandler unsafe.Pointer) foundation.Progress
	StartStreamWithNameToPeerError(streamName objc.IObject /* cross-framework: NSString */, peerID IMCPeerID, error_ unsafe.Pointer) foundation.OutputStream
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MCSession */
// Alloc allocates a new instance without initialization.
func (mc _MCSessionClass) Alloc() MCSession {
	rv := objc.Send[MCSession](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MCSession */
// An object enables and manages communication among all peers in a Multipeer Connectivity session.


// An object enables and manages communication among all peers in a Multipeer Connectivity session.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MCSession */

// Creates a Multipeer Connectivity session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCSession/init(peer:)
func NewMCSessionWithPeer(myPeerID IMCPeerID) MCSession {
	instance := getMCSessionClass().Alloc()
	rv := objc.Send[MCSession](instance.ID, objc.Sel("initWithPeer:"), myPeerID)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMCSessionWithPeer */


// Creates a Multipeer Connectivity session, providing security information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCSession/init(peer:securityIdentity:encryptionPreference:)
func NewMCSessionWithPeerSecurityIdentityEncryptionPreference(myPeerID IMCPeerID, identity objc.IObject /* cross-framework: NSArray */, encryptionPreference MCEncryptionPreference) MCSession {
	instance := getMCSessionClass().Alloc()
	rv := objc.Send[MCSession](instance.ID, objc.Sel("initWithPeer:securityIdentity:encryptionPreference:"), myPeerID, identity, encryptionPreference)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMCSessionWithPeerSecurityIdentityEncryptionPreference */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MCSession */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MCSession */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MCSession */

// Cancels an attempt to connect to a peer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCSession/cancelConnectPeer(_:)
func (m_ MCSession) CancelConnectPeer(peerID IMCPeerID) {
	objc.Send[objc.ID](m_.ID, objc.Sel("cancelConnectPeer:"), peerID)
}/* debug [instance_methods/method]: CancelConnectPeer */


// Call this method to connect a peer to the session when using your own service discovery code instead of an or object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCSession/connectPeer(_:withNearbyConnectionData:)
func (m_ MCSession) ConnectPeerWithNearbyConnectionData(peerID IMCPeerID, data objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("connectPeer:withNearbyConnectionData:"), peerID, data)
}/* debug [instance_methods/method]: ConnectPeerWithNearbyConnectionData */


// Disconnects the local peer from the session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCSession/disconnect()
func (m_ MCSession) Disconnect() {
	objc.Send[objc.ID](m_.ID, objc.Sel("disconnect"))
}/* debug [instance_methods/method]: Disconnect */


// Obtains connection data for the specified peer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCSession/nearbyConnectionData(forPeer:withCompletionHandler:)
func (m_ MCSession) NearbyConnectionDataForPeerWithCompletionHandler(peerID IMCPeerID, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("nearbyConnectionDataForPeer:withCompletionHandler:"), peerID, completionHandler)
}/* debug [instance_methods/method]: NearbyConnectionDataForPeerWithCompletionHandler */


// Sends a message to nearby peers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCSession/send(_:toPeers:with:)
func (m_ MCSession) SendDataToPeersWithModeError(data objc.IObject /* cross-framework: NSData */, peerIDs []MCPeerID, mode MCSessionSendDataMode, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("sendData:toPeers:withMode:error:"), data, peerIDs, mode, error_)
	return rv
}/* debug [instance_methods/method]: SendDataToPeersWithModeError */


// Sends the contents of a URL to a peer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCSession/sendResource(at:withName:toPeer:withCompletionHandler:)
func (m_ MCSession) SendResourceAtURLWithNameToPeerWithCompletionHandler(resourceURL objc.IObject /* cross-framework: NSURL */, resourceName objc.IObject /* cross-framework: NSString */, peerID IMCPeerID, completionHandler unsafe.Pointer) foundation.Progress {
	rv := objc.Send[foundation.Progress](m_.ID, objc.Sel("sendResourceAtURL:withName:toPeer:withCompletionHandler:"), resourceURL, resourceName, peerID, completionHandler)
	return rv
}/* debug [instance_methods/method]: SendResourceAtURLWithNameToPeerWithCompletionHandler */


// Opens a byte stream to a nearby peer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCSession/startStream(withName:toPeer:)
func (m_ MCSession) StartStreamWithNameToPeerError(streamName objc.IObject /* cross-framework: NSString */, peerID IMCPeerID, error_ unsafe.Pointer) foundation.OutputStream {
	rv := objc.Send[foundation.OutputStream](m_.ID, objc.Sel("startStreamWithName:toPeer:error:"), streamName, peerID, error_)
	return rv
}/* debug [instance_methods/method]: StartStreamWithNameToPeerError */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MCSession */

// An array of all peers that are currently connected to this session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCSession/connectedPeers
func (m_ MCSession) ConnectedPeers() []MCPeerID {
	rv := objc.Send[[]MCPeerID](m_.ID, objc.Sel("connectedPeers"))
	return rv
}/* debug [instance_properties/getter]: connectedPeers */


// The delegate object that handles session-related events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCSession/delegate
func (m_ MCSession) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The delegate object that handles session-related events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCSession/delegate
func (m_ MCSession) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// A value indicating whether the connection prefers encrypted connections, unencrypted connections, or has no preference.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCSession/encryptionPreference
func (m_ MCSession) EncryptionPreference() MCEncryptionPreference {
	rv := objc.Send[MCEncryptionPreference](m_.ID, objc.Sel("encryptionPreference"))
	return rv
}/* debug [instance_properties/getter]: encryptionPreference */


// A local identifier that represents the device on which your app is currently running.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCSession/myPeerID
func (m_ MCSession) MyPeerID() IMCPeerID {
	rv := objc.Send[MCPeerID](m_.ID, objc.Sel("myPeerID"))
	return rv
}/* debug [instance_properties/getter]: myPeerID */


// The security identity of the local peer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCSession/securityIdentity
func (m_ MCSession) SecurityIdentity() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("securityIdentity"))
	return rv
}/* debug [instance_properties/getter]: securityIdentity */


// The
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/multipeerconnectivity/mcerrordomain
func (m_ MCSession) MCErrorDomain() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("MCErrorDomain"))
	return rv
}/* debug [instance_properties/getter]: MCErrorDomain */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MCSession */


