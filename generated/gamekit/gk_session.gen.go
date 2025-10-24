// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GKSession */


/* debug [class_header]: Header for GKSession */
// The class instance for the [Session] class.
var (
	SessionClass     _SessionClass
	SessionClassOnce sync.Once
)

func getSessionClass() _SessionClass {
	SessionClassOnce.Do(func() {
		SessionClass = _SessionClass{objc.GetClass("GKSession")}
	})
	return SessionClass
}

type _SessionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Session */
// An interface definition for the [Session] class.
type ISession interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Session */
	// properties:
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	DisconnectTimeout() float64
	SetDisconnectTimeout(value float64)
	DisplayName() objc.IObject /* cross-framework: NSString */
	Available() bool
	SetAvailable(value bool)
	PeerID() objc.IObject /* cross-framework: NSString */
	SessionID() objc.IObject /* cross-framework: NSString */
	SessionMode() SessionMode
	IsAvailable() bool
	SetIsAvailable(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Session */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Session */
// Alloc allocates a new instance without initialization.
func (sc _SessionClass) Alloc() Session {
	rv := objc.Send[Session](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SessionClass) New() Session {
	rv := objc.Send[Session](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ Session) Init() Session {
	rv := objc.Send[Session](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ Session) Autorelease() Session {
	rv := objc.Send[Session](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSession creates a new Session instance.
func NewSession() Session {
	return getSessionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Session */
// A object provides the ability to discover and connect to nearby iOS devices using Bluetooth or Wi-fi.
//
// Sessions primarily work with . A peer is any iOS device made visible by creating and configuring a object. Each peer is identified by a unique identifier, called a peer id ( ) string. Your application can use a string to obtain a user-readable name for a remote peer and to attempt to connect to that peer. Similarly, your session’s peer ID is visible to other nearby peers. After a connection is established, your application uses the remote peer’s ID to address data packets that it wants to send. Peers discover other peers by using a unique string to identify the service they implement, called a session ID ( ). Sessions can be configured to broadcast a session ID (as a ), to search for other peers advertising with that session ID (as a ), or to act as both a server and a client simultaneously (as a . Your application controls the behavior of a session through a delegate that implements the protocol. The delegate is called when remote peers are discovered, when those peers attempt to connect to the session, and when the state of a remote peer changes. Your application also provides a data handler to the session so that the session can forward data it receives from remote peers. The data handler can be a separate object or the same object as the delegate. When Bluetooth is turned on, Wi-Fi download speeds drastically decrease while the device is searching for other Bluetooth enabled devices. After the Bluetooth discovery time has completed, Wi-Fi speeds return to normal. methods are thread-safe and may be called from any thread. However, the session always calls its delegate on the main thread.


// A object provides the ability to discover and connect to nearby iOS devices using Bluetooth or Wi-fi.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKSession
type Session struct {
	objectivec.Object
}

// SessionFrom constructs a [Session] from an unsafe.Pointer.
//
// A object provides the ability to discover and connect to nearby iOS devices using Bluetooth or Wi-fi.
func SessionFrom(ptr unsafe.Pointer) Session {
	return Session{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Session */

// Initializes and returns a newly allocated session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKSession/init(sessionID:displayName:sessionMode:)
func NewSessionWithSessionIDDisplayNameSessionMode(sessionID objc.IObject /* cross-framework: NSString */, name objc.IObject /* cross-framework: NSString */, mode SessionMode) Session {
	instance := getSessionClass().Alloc()
	rv := objc.Send[Session](instance.ID, objc.Sel("initWithSessionID:displayName:sessionMode:"), sessionID, name, mode)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewSessionWithSessionIDDisplayNameSessionMode */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Session */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Session */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Session */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Session */

// The delegate of the session object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKSession/delegate
func (s_ Session) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The delegate of the session object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKSession/delegate
func (s_ Session) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// A time interval that expresses how long the session waits before it disconnects a nonresponsive peer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKSession/disconnectTimeout
func (s_ Session) DisconnectTimeout() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("disconnectTimeout"))
	return rv
}/* debug [instance_properties/getter]: disconnectTimeout */


// A time interval that expresses how long the session waits before it disconnects a nonresponsive peer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKSession/disconnectTimeout
func (s_ Session) SetDisconnectTimeout(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDisconnectTimeout:"), value)
}/* debug [instance_properties/setter]: disconnectTimeout */


// The name of the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKSession/displayName
func (s_ Session) DisplayName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("displayName"))
	return rv
}/* debug [instance_properties/getter]: displayName */


// A Boolean value that determines whether or not the session wants to connect to new peers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKSession/isAvailable
func (s_ Session) Available() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("available"))
	return rv
}/* debug [instance_properties/getter]: available */


// A Boolean value that determines whether or not the session wants to connect to new peers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKSession/isAvailable
func (s_ Session) SetAvailable(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAvailable:"), value)
}/* debug [instance_properties/setter]: available */


// A string that identifies your session to other peers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKSession/peerID
func (s_ Session) PeerID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("peerID"))
	return rv
}/* debug [instance_properties/getter]: peerID */


// A string used to filter the list of peers who are allowed to see your session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKSession/sessionID
func (s_ Session) SessionID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("sessionID"))
	return rv
}/* debug [instance_properties/getter]: sessionID */


// The mode the session uses to find other peers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKSession/sessionMode
func (s_ Session) SessionMode() SessionMode {
	rv := objc.Send[SessionMode](s_.ID, objc.Sel("sessionMode"))
	return rv
}/* debug [instance_properties/getter]: sessionMode */


// A Boolean value that determines whether or not the session wants to connect to new peers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gksession/isavailable
func (s_ Session) IsAvailable() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isAvailable"))
	return rv
}/* debug [instance_properties/getter]: isAvailable */


// A Boolean value that determines whether or not the session wants to connect to new peers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gksession/isavailable
func (s_ Session) SetIsAvailable(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsAvailable:"), value)
}/* debug [instance_properties/setter]: isAvailable */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKSession */


