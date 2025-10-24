// Code generated from Apple documentation for MultipeerConnectivity. DO NOT EDIT.

package multipeerconnectivity

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MCNearbyServiceAdvertiser */


/* debug [class_header]: Header for MCNearbyServiceAdvertiser */
// The class instance for the [MCNearbyServiceAdvertiser] class.
var (
	MCNearbyServiceAdvertiserClass     _MCNearbyServiceAdvertiserClass
	MCNearbyServiceAdvertiserClassOnce sync.Once
)

func getMCNearbyServiceAdvertiserClass() _MCNearbyServiceAdvertiserClass {
	MCNearbyServiceAdvertiserClassOnce.Do(func() {
		MCNearbyServiceAdvertiserClass = _MCNearbyServiceAdvertiserClass{objc.GetClass("MCNearbyServiceAdvertiser")}
	})
	return MCNearbyServiceAdvertiserClass
}

type _MCNearbyServiceAdvertiserClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MCNearbyServiceAdvertiser */
// An interface definition for the [MCNearbyServiceAdvertiser] class.
type IMCNearbyServiceAdvertiser interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MCNearbyServiceAdvertiser */
	// properties:
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	DiscoveryInfo() foundation.IDictionary
	MyPeerID() IMCPeerID
	ServiceType() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MCNearbyServiceAdvertiser */
	// methods:
	StartAdvertisingPeer()
	StopAdvertisingPeer()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MCNearbyServiceAdvertiser */
// Alloc allocates a new instance without initialization.
func (mc _MCNearbyServiceAdvertiserClass) Alloc() MCNearbyServiceAdvertiser {
	rv := objc.Send[MCNearbyServiceAdvertiser](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MCNearbyServiceAdvertiserClass) New() MCNearbyServiceAdvertiser {
	rv := objc.Send[MCNearbyServiceAdvertiser](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MCNearbyServiceAdvertiser) Init() MCNearbyServiceAdvertiser {
	rv := objc.Send[MCNearbyServiceAdvertiser](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MCNearbyServiceAdvertiser) Autorelease() MCNearbyServiceAdvertiser {
	rv := objc.Send[MCNearbyServiceAdvertiser](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMCNearbyServiceAdvertiser creates a new MCNearbyServiceAdvertiser instance.
func NewMCNearbyServiceAdvertiser() MCNearbyServiceAdvertiser {
	return getMCNearbyServiceAdvertiserClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MCNearbyServiceAdvertiser */
// The class publishes an advertisement for a specific service that your app provides through the Multipeer Connectivity framework and notifies its delegate about invitations from nearby peers.
//
// Before you can advertise a service, you must create an object that identifies your app and the user to nearby devices. The parameter is a short text string used to describe the app’s networking protocol. It should be in the same format as a Bonjour service type: 1–15 characters long and valid characters include ASCII lowercase letters, numbers, and the hyphen, containing at least one letter and no adjacent hyphens. A short name that distinguishes itself from unrelated services is recommended; for example, a text chat app made by ABC company could use the service type . For more information about service types, read . The parameter is a dictionary of string key/value pairs that will be advertised for browsers to see. The content of will be advertised within Bonjour TXT records, so you should keep the dictionary small for better discovery performance. For more information about TXT records, read .


// The class publishes an advertisement for a specific service that your app provides through the Multipeer Connectivity framework and notifies its delegate about invitations from nearby peers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCNearbyServiceAdvertiser
type MCNearbyServiceAdvertiser struct {
	objectivec.Object
}

// MCNearbyServiceAdvertiserFrom constructs a [MCNearbyServiceAdvertiser] from an unsafe.Pointer.
//
// The class publishes an advertisement for a specific service that your app provides through the Multipeer Connectivity framework and notifies its delegate about invitations from nearby peers.
func MCNearbyServiceAdvertiserFrom(ptr unsafe.Pointer) MCNearbyServiceAdvertiser {
	return MCNearbyServiceAdvertiser{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MCNearbyServiceAdvertiser */

// Initializes an advertiser object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCNearbyServiceAdvertiser/init(peer:discoveryInfo:serviceType:)
func NewMCNearbyServiceAdvertiserWithPeerDiscoveryInfoServiceType(myPeerID IMCPeerID, info foundation.IDictionary, serviceType objc.IObject /* cross-framework: NSString */) MCNearbyServiceAdvertiser {
	instance := getMCNearbyServiceAdvertiserClass().Alloc()
	rv := objc.Send[MCNearbyServiceAdvertiser](instance.ID, objc.Sel("initWithPeer:discoveryInfo:serviceType:"), myPeerID, info, serviceType)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMCNearbyServiceAdvertiserWithPeerDiscoveryInfoServiceType */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MCNearbyServiceAdvertiser */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MCNearbyServiceAdvertiser */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MCNearbyServiceAdvertiser */

// Begins advertising the service provided by a local peer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCNearbyServiceAdvertiser/startAdvertisingPeer()
func (m_ MCNearbyServiceAdvertiser) StartAdvertisingPeer() {
	objc.Send[objc.ID](m_.ID, objc.Sel("startAdvertisingPeer"))
}/* debug [instance_methods/method]: StartAdvertisingPeer */


// Stops advertising the service provided by a local peer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCNearbyServiceAdvertiser/stopAdvertisingPeer()
func (m_ MCNearbyServiceAdvertiser) StopAdvertisingPeer() {
	objc.Send[objc.ID](m_.ID, objc.Sel("stopAdvertisingPeer"))
}/* debug [instance_methods/method]: StopAdvertisingPeer */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MCNearbyServiceAdvertiser */

// The delegate object that handles advertising-related events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCNearbyServiceAdvertiser/delegate
func (m_ MCNearbyServiceAdvertiser) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The delegate object that handles advertising-related events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCNearbyServiceAdvertiser/delegate
func (m_ MCNearbyServiceAdvertiser) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// The dictionary passed when this object was initialized.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCNearbyServiceAdvertiser/discoveryInfo
func (m_ MCNearbyServiceAdvertiser) DiscoveryInfo() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("discoveryInfo"))
	return rv
}/* debug [instance_properties/getter]: discoveryInfo */


// The local peer ID for this instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCNearbyServiceAdvertiser/myPeerID
func (m_ MCNearbyServiceAdvertiser) MyPeerID() IMCPeerID {
	rv := objc.Send[MCPeerID](m_.ID, objc.Sel("myPeerID"))
	return rv
}/* debug [instance_properties/getter]: myPeerID */


// The service type that your app is advertising
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCNearbyServiceAdvertiser/serviceType
func (m_ MCNearbyServiceAdvertiser) ServiceType() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("serviceType"))
	return rv
}/* debug [instance_properties/getter]: serviceType */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MCNearbyServiceAdvertiser */


