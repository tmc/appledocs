// Code generated from Apple documentation for MultipeerConnectivity. DO NOT EDIT.

package multipeerconnectivity

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [MCNearbyServiceAdvertiser] class.
type IMCNearbyServiceAdvertiser interface {
	objectivec.IObject
	StartAdvertisingPeer()
	StopAdvertisingPeer()
}

// The class publishes an advertisement for a specific service that your app provides through the Multipeer Connectivity framework and notifies its delegate about invitations from nearby peers.
//
// Before you can advertise a service, you must create an object that identifies your app and the user to nearby devices. The parameter is a short text string used to describe the app’s networking protocol. It should be in the same format as a Bonjour service type: 1–15 characters long and valid characters include ASCII lowercase letters, numbers, and the hyphen, containing at least one letter and no adjacent hyphens. A short name that distinguishes itself from unrelated services is recommended; for example, a text chat app made by ABC company could use the service type . For more information about service types, read . The parameter is a dictionary of string key/value pairs that will be advertised for browsers to see. The content of will be advertised within Bonjour TXT records, so you should keep the dictionary small for better discovery performance. For more information about TXT records, read .
//
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

// Alloc allocates a new instance without initialization.
func (mc _MCNearbyServiceAdvertiserClass) Alloc() MCNearbyServiceAdvertiser {
	rv := objc.Send[MCNearbyServiceAdvertiser](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Initializes an advertiser object.
//
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCNearbyServiceAdvertiser/init(peer:discoveryInfo:serviceType:)
func NewMCNearbyServiceAdvertiserWithPeerDiscoveryInfoServiceType(myPeerID IMCPeerID, info unsafe.Pointer, serviceType appkit.string) MCNearbyServiceAdvertiser {
	instance := getMCNearbyServiceAdvertiserClass().Alloc()
	rv := objc.Send[MCNearbyServiceAdvertiser](instance.ID, objc.Sel("initWithPeer:discoveryInfo:serviceType:"), myPeerID, info, serviceType)
	rv.Autorelease()
	return rv
}


// Begins advertising the service provided by a local peer.
//
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCNearbyServiceAdvertiser/startAdvertisingPeer()
func (m_ MCNearbyServiceAdvertiser) StartAdvertisingPeer() {
	objc.Send[objc.ID](m_.ID, objc.Sel("startAdvertisingPeer"))
}

// Stops advertising the service provided by a local peer.
//
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCNearbyServiceAdvertiser/stopAdvertisingPeer()
func (m_ MCNearbyServiceAdvertiser) StopAdvertisingPeer() {
	objc.Send[objc.ID](m_.ID, objc.Sel("stopAdvertisingPeer"))
}

// The delegate object that handles advertising-related events.
//
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCNearbyServiceAdvertiser/delegate
func (m_ MCNearbyServiceAdvertiser) Delegate() objc.ID {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The delegate object that handles advertising-related events.

//
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCNearbyServiceAdvertiser/delegate
func (m_ MCNearbyServiceAdvertiser) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDelegate:"), value)
}

// The dictionary passed when this object was initialized.
//
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCNearbyServiceAdvertiser/discoveryInfo
func (m_ MCNearbyServiceAdvertiser) DiscoveryInfo() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("discoveryInfo"))
	return rv
}

// The local peer ID for this instance.
//
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCNearbyServiceAdvertiser/myPeerID
func (m_ MCNearbyServiceAdvertiser) MyPeerID() MCPeerID {
	rv := objc.Send[MCPeerID](m_.ID, objc.Sel("myPeerID"))
	return rv
}

// The service type that your app is advertising
//
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCNearbyServiceAdvertiser/serviceType
func (m_ MCNearbyServiceAdvertiser) ServiceType() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("serviceType"))
	return rv
}


