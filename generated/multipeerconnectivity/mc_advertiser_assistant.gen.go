// Code generated from Apple documentation for MultipeerConnectivity. DO NOT EDIT.

package multipeerconnectivity

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [MCAdvertiserAssistant] class.
type IMCAdvertiserAssistant interface {
	objectivec.IObject
	Start()
	Stop()
}

// The is a convenience class that handles advertising, presents incoming invitations to the user, and handles users’ responses. Use this class to provide a user interface for handling invitations when your app does not require programmatic control over the invitation process.
//
// Before you can advertise a service, you must create an object that identifies your app and the user to nearby devices.
//
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

// Alloc allocates a new instance without initialization.
func (mc _MCAdvertiserAssistantClass) Alloc() MCAdvertiserAssistant {
	rv := objc.Send[MCAdvertiserAssistant](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Initializes an advertiser assistant object.
//
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCAdvertiserAssistant/init(serviceType:discoveryInfo:session:)
func NewMCAdvertiserAssistantWithServiceTypeDiscoveryInfoSession(serviceType string, info unsafe.Pointer, session unsafe.Pointer) MCAdvertiserAssistant {
	instance := getMCAdvertiserAssistantClass().Alloc()
	rv := objc.Send[MCAdvertiserAssistant](instance.ID, objc.Sel("initWithServiceType:discoveryInfo:session:"), objc.String(serviceType), info, session)
	rv.Autorelease()
	return rv
}


// Begins advertising the service provided by a local peer and starts the assistant.
//
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCAdvertiserAssistant/start()
func (m_ MCAdvertiserAssistant) Start() {
	objc.Send[objc.ID](m_.ID, objc.Sel("start"))
}

// Stops advertising the service provided by a local peer and stops the assistant.
//
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCAdvertiserAssistant/stop()
func (m_ MCAdvertiserAssistant) Stop() {
	objc.Send[objc.ID](m_.ID, objc.Sel("stop"))
}

// The delegate object that handles advertising-assistant-related events.
//
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCAdvertiserAssistant/delegate
func (m_ MCAdvertiserAssistant) Delegate() objc.ID {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The delegate object that handles advertising-assistant-related events.

//
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCAdvertiserAssistant/delegate
func (m_ MCAdvertiserAssistant) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDelegate:"), value)
}
// The dictionary that was passed when this object was initialized.
//
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCAdvertiserAssistant/discoveryInfo
func (m_ MCAdvertiserAssistant) DiscoveryInfo() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("discoveryInfo"))
	return rv
}

// The service type that your app is advertising.
//
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCAdvertiserAssistant/serviceType
func (m_ MCAdvertiserAssistant) ServiceType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("serviceType"))
	return rv
}

// The session into which new peers are added after accepting an invitation.
//
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCAdvertiserAssistant/session
func (m_ MCAdvertiserAssistant) Session() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("session"))
	return rv
}


