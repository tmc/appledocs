// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MIDINetworkSession */


/* debug [class_header]: Header for MIDINetworkSession */
// The class instance for the [MIDINetworkSession] class.
var (
	MIDINetworkSessionClass     _MIDINetworkSessionClass
	MIDINetworkSessionClassOnce sync.Once
)

func getMIDINetworkSessionClass() _MIDINetworkSessionClass {
	MIDINetworkSessionClassOnce.Do(func() {
		MIDINetworkSessionClass = _MIDINetworkSessionClass{objc.GetClass("MIDINetworkSession")}
	})
	return MIDINetworkSessionClass
}

type _MIDINetworkSessionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MIDINetworkSession */
// An interface definition for the [MIDINetworkSession] class.
type IMIDINetworkSession interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MIDINetworkSession */
	// properties:
	ConnectionPolicy() MIDINetworkConnectionPolicy
	SetConnectionPolicy(value MIDINetworkConnectionPolicy)
	Enabled() bool
	SetEnabled(value bool)
	LocalName() objc.IObject /* cross-framework: NSString */
	NetworkName() objc.IObject /* cross-framework: NSString */
	NetworkPort() uint
	MIDINetworkNotificationContactsDidChange() objc.IObject /* cross-framework: NSString */
	MIDINetworkNotificationSessionDidChange() objc.IObject /* cross-framework: NSString */
	IsEnabled() bool
	SetIsEnabled(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MIDINetworkSession */
	// methods:
	AddConnection(connection IMIDINetworkConnection) bool
	AddContact(contact IMIDINetworkHost) bool
	Connections() unsafe.Pointer
	Contacts() unsafe.Pointer
	DestinationEndpoint() MIDIEndpointRef /* typedef */
	RemoveConnection(connection IMIDINetworkConnection) bool
	RemoveContact(contact IMIDINetworkHost) bool
	SourceEndpoint() MIDIEndpointRef /* typedef */
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MIDINetworkSession */
// Alloc allocates a new instance without initialization.
func (mc _MIDINetworkSessionClass) Alloc() MIDINetworkSession {
	rv := objc.Send[MIDINetworkSession](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MIDINetworkSessionClass) New() MIDINetworkSession {
	rv := objc.Send[MIDINetworkSession](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MIDINetworkSession) Init() MIDINetworkSession {
	rv := objc.Send[MIDINetworkSession](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MIDINetworkSession) Autorelease() MIDINetworkSession {
	rv := objc.Send[MIDINetworkSession](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMIDINetworkSession creates a new MIDINetworkSession instance.
func NewMIDINetworkSession() MIDINetworkSession {
	return getMIDINetworkSessionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MIDINetworkSession */
// An object that represents a pairing of a source and destination.
//
// A session can have any number of connections. The system broadcasts output to all connections, and merges input from multiple connections.


// An object that represents a pairing of a source and destination.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkSession
type MIDINetworkSession struct {
	objectivec.Object
}

// MIDINetworkSessionFrom constructs a [MIDINetworkSession] from an unsafe.Pointer.
//
// An object that represents a pairing of a source and destination.
func MIDINetworkSessionFrom(ptr unsafe.Pointer) MIDINetworkSession {
	return MIDINetworkSession{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MIDINetworkSession *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MIDINetworkSession */

// Returns the default singleton session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkSession/default()
func (mc _MIDINetworkSessionClass) DefaultSession() MIDINetworkSession {
	rv := objc.Send[MIDINetworkSession](objc.ID(mc.class), objc.Sel("defaultSession"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DefaultSession) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MIDINetworkSession */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MIDINetworkSession */

// Adds a new connection to this session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkSession/addConnection(_:)
func (m_ MIDINetworkSession) AddConnection(connection IMIDINetworkConnection) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("addConnection:"), connection)
	return rv
}/* debug [instance_methods/method]: AddConnection */


// Adds a host as a contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkSession/addContact(_:)
func (m_ MIDINetworkSession) AddContact(contact IMIDINetworkHost) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("addContact:"), contact)
	return rv
}/* debug [instance_methods/method]: AddContact */


// Returns the session’s set of MIDI network connections.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkSession/connections()
func (m_ MIDINetworkSession) Connections() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("connections"))
	return rv
}/* debug [instance_methods/method]: Connections */


// Returns the array of network hosts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkSession/contacts()
func (m_ MIDINetworkSession) Contacts() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("contacts"))
	return rv
}/* debug [instance_methods/method]: Contacts */


// Returns the session’s destination endpoint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkSession/destinationEndpoint()
func (m_ MIDINetworkSession) DestinationEndpoint() MIDIEndpointRef /* typedef */ {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("destinationEndpoint"))
	return rv
}/* debug [instance_methods/method]: DestinationEndpoint */


// Removes a connection from this session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkSession/removeConnection(_:)
func (m_ MIDINetworkSession) RemoveConnection(connection IMIDINetworkConnection) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("removeConnection:"), connection)
	return rv
}/* debug [instance_methods/method]: RemoveConnection */


// Removes a host as a contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkSession/removeContact(_:)
func (m_ MIDINetworkSession) RemoveContact(contact IMIDINetworkHost) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("removeContact:"), contact)
	return rv
}/* debug [instance_methods/method]: RemoveContact */


// Returns the session’s source endpoint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkSession/sourceEndpoint()
func (m_ MIDINetworkSession) SourceEndpoint() MIDIEndpointRef /* typedef */ {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("sourceEndpoint"))
	return rv
}/* debug [instance_methods/method]: SourceEndpoint */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MIDINetworkSession */

// The policy that determines who can connect to this session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkSession/connectionPolicy
func (m_ MIDINetworkSession) ConnectionPolicy() MIDINetworkConnectionPolicy {
	rv := objc.Send[MIDINetworkConnectionPolicy](m_.ID, objc.Sel("connectionPolicy"))
	return rv
}/* debug [instance_properties/getter]: connectionPolicy */


// The policy that determines who can connect to this session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkSession/connectionPolicy
func (m_ MIDINetworkSession) SetConnectionPolicy(value MIDINetworkConnectionPolicy) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setConnectionPolicy:"), value)
}/* debug [instance_properties/setter]: connectionPolicy */


// A Boolean value that determines whether the session is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkSession/isEnabled
func (m_ MIDINetworkSession) Enabled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("enabled"))
	return rv
}/* debug [instance_properties/getter]: enabled */


// A Boolean value that determines whether the session is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkSession/isEnabled
func (m_ MIDINetworkSession) SetEnabled(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEnabled:"), value)
}/* debug [instance_properties/setter]: enabled */


// The name of this session’s entity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkSession/localName
func (m_ MIDINetworkSession) LocalName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("localName"))
	return rv
}/* debug [instance_properties/getter]: localName */


// The name with which this session advertises itself over Bonjour.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkSession/networkName
func (m_ MIDINetworkSession) NetworkName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("networkName"))
	return rv
}/* debug [instance_properties/getter]: networkName */


// The session’s UDP port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkSession/networkPort
func (m_ MIDINetworkSession) NetworkPort() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("networkPort"))
	return rv
}/* debug [instance_properties/getter]: networkPort */


// Indicates that the list of contacts changed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midinetworknotificationcontactsdidchange
func (m_ MIDINetworkSession) MIDINetworkNotificationContactsDidChange() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("MIDINetworkNotificationContactsDidChange"))
	return rv
}/* debug [instance_properties/getter]: MIDINetworkNotificationContactsDidChange */


// Indicates that other aspects of the session changed, such as the connection list, connection policy, and so on.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midinetworknotificationsessiondidchange
func (m_ MIDINetworkSession) MIDINetworkNotificationSessionDidChange() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("MIDINetworkNotificationSessionDidChange"))
	return rv
}/* debug [instance_properties/getter]: MIDINetworkNotificationSessionDidChange */


// A Boolean value that determines whether the session is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midinetworksession/isenabled
func (m_ MIDINetworkSession) IsEnabled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isEnabled"))
	return rv
}/* debug [instance_properties/getter]: isEnabled */


// A Boolean value that determines whether the session is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midinetworksession/isenabled
func (m_ MIDINetworkSession) SetIsEnabled(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsEnabled:"), value)
}/* debug [instance_properties/setter]: isEnabled */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MIDINetworkSession */





