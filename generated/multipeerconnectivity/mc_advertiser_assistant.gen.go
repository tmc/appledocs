// Code generated from Apple documentation for MultipeerConnectivity. DO NOT EDIT.

package multipeerconnectivity

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MCAdvertiserAssistant */


/* debug [class_header]: Header for MCAdvertiserAssistant */
// The class instance for the [MCAdvertiserAssistant] class.
var (
	MCAdvertiserAssistantClass     _MCAdvertiserAssistantClass
	MCAdvertiserAssistantClassOnce sync.Once
)

func getMCAdvertiserAssistantClass() _MCAdvertiserAssistantClass {
	MCAdvertiserAssistantClassOnce.Do(func() {
		MCAdvertiserAssistantClass = _MCAdvertiserAssistantClass{objc.GetClass("MCAdvertiserAssistant")}
	})
	return MCAdvertiserAssistantClass
}

type _MCAdvertiserAssistantClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MCAdvertiserAssistant */
// An interface definition for the [MCAdvertiserAssistant] class.
type IMCAdvertiserAssistant interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MCAdvertiserAssistant */
	// properties:
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	DiscoveryInfo() foundation.IDictionary
	ServiceType() objc.IObject /* cross-framework: NSString */
	Session() IMCSession
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MCAdvertiserAssistant */
	// methods:
	Start()
	Stop()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MCAdvertiserAssistant */
// Alloc allocates a new instance without initialization.
func (mc _MCAdvertiserAssistantClass) Alloc() MCAdvertiserAssistant {
	rv := objc.Send[MCAdvertiserAssistant](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MCAdvertiserAssistantClass) New() MCAdvertiserAssistant {
	rv := objc.Send[MCAdvertiserAssistant](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MCAdvertiserAssistant) Init() MCAdvertiserAssistant {
	rv := objc.Send[MCAdvertiserAssistant](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MCAdvertiserAssistant) Autorelease() MCAdvertiserAssistant {
	rv := objc.Send[MCAdvertiserAssistant](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMCAdvertiserAssistant creates a new MCAdvertiserAssistant instance.
func NewMCAdvertiserAssistant() MCAdvertiserAssistant {
	return getMCAdvertiserAssistantClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MCAdvertiserAssistant */
// The is a convenience class that handles advertising, presents incoming invitations to the user, and handles users’ responses. Use this class to provide a user interface for handling invitations when your app does not require programmatic control over the invitation process.
//
// Before you can advertise a service, you must create an object that identifies your app and the user to nearby devices.


// The is a convenience class that handles advertising, presents incoming invitations to the user, and handles users’ responses. Use this class to provide a user interface for handling invitations when your app does not require programmatic control over the invitation process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCAdvertiserAssistant
type MCAdvertiserAssistant struct {
	objectivec.Object
}

// MCAdvertiserAssistantFrom constructs a [MCAdvertiserAssistant] from an unsafe.Pointer.
//
// The is a convenience class that handles advertising, presents incoming invitations to the user, and handles users’ responses. Use this class to provide a user interface for handling invitations when your app does not require programmatic control over the invitation process.
func MCAdvertiserAssistantFrom(ptr unsafe.Pointer) MCAdvertiserAssistant {
	return MCAdvertiserAssistant{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MCAdvertiserAssistant */

// Initializes an advertiser assistant object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCAdvertiserAssistant/init(serviceType:discoveryInfo:session:)
func NewMCAdvertiserAssistantWithServiceTypeDiscoveryInfoSession(serviceType objc.IObject /* cross-framework: NSString */, info foundation.IDictionary, session IMCSession) MCAdvertiserAssistant {
	instance := getMCAdvertiserAssistantClass().Alloc()
	rv := objc.Send[MCAdvertiserAssistant](instance.ID, objc.Sel("initWithServiceType:discoveryInfo:session:"), serviceType, info, session)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMCAdvertiserAssistantWithServiceTypeDiscoveryInfoSession */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MCAdvertiserAssistant */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MCAdvertiserAssistant */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MCAdvertiserAssistant */

// Begins advertising the service provided by a local peer and starts the assistant.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCAdvertiserAssistant/start()
func (m_ MCAdvertiserAssistant) Start() {
	objc.Send[objc.ID](m_.ID, objc.Sel("start"))
}/* debug [instance_methods/method]: Start */


// Stops advertising the service provided by a local peer and stops the assistant.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCAdvertiserAssistant/stop()
func (m_ MCAdvertiserAssistant) Stop() {
	objc.Send[objc.ID](m_.ID, objc.Sel("stop"))
}/* debug [instance_methods/method]: Stop */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MCAdvertiserAssistant */

// The delegate object that handles advertising-assistant-related events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCAdvertiserAssistant/delegate
func (m_ MCAdvertiserAssistant) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The delegate object that handles advertising-assistant-related events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCAdvertiserAssistant/delegate
func (m_ MCAdvertiserAssistant) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// The dictionary that was passed when this object was initialized.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCAdvertiserAssistant/discoveryInfo
func (m_ MCAdvertiserAssistant) DiscoveryInfo() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("discoveryInfo"))
	return rv
}/* debug [instance_properties/getter]: discoveryInfo */


// The service type that your app is advertising.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCAdvertiserAssistant/serviceType
func (m_ MCAdvertiserAssistant) ServiceType() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("serviceType"))
	return rv
}/* debug [instance_properties/getter]: serviceType */


// The session into which new peers are added after accepting an invitation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCAdvertiserAssistant/session
func (m_ MCAdvertiserAssistant) Session() IMCSession {
	rv := objc.Send[MCSession](m_.ID, objc.Sel("session"))
	return rv
}/* debug [instance_properties/getter]: session */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MCAdvertiserAssistant */


