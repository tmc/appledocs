// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NETunnelProviderSession */


/* debug [class_header]: Header for NETunnelProviderSession */
// The class instance for the [NETunnelProviderSession] class.
var (
	NETunnelProviderSessionClass     _NETunnelProviderSessionClass
	NETunnelProviderSessionClassOnce sync.Once
)

func getNETunnelProviderSessionClass() _NETunnelProviderSessionClass {
	NETunnelProviderSessionClassOnce.Do(func() {
		NETunnelProviderSessionClass = _NETunnelProviderSessionClass{objc.GetClass("NETunnelProviderSession")}
	})
	return NETunnelProviderSessionClass
}

type _NETunnelProviderSessionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NETunnelProviderSession */
// An interface definition for the [NETunnelProviderSession] class.
type INETunnelProviderSession interface {
	INEVPNConnection
	
/* debug [class_interface_properties]: Properties for NETunnelProviderSession */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NETunnelProviderSession */
	// methods:
	SendProviderMessageReturnErrorResponseHandler(messageData objc.IObject /* cross-framework: NSData */, error_ objectivec.IObject, responseHandler unsafe.Pointer) bool
	StartTunnelWithOptionsAndReturnError(options foundation.IDictionary, error_ objectivec.IObject) bool
	StopTunnel()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NETunnelProviderSession */
// Alloc allocates a new instance without initialization.
func (nc _NETunnelProviderSessionClass) Alloc() NETunnelProviderSession {
	rv := objc.Send[NETunnelProviderSession](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NETunnelProviderSessionClass) New() NETunnelProviderSession {
	rv := objc.Send[NETunnelProviderSession](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NETunnelProviderSession) Init() NETunnelProviderSession {
	rv := objc.Send[NETunnelProviderSession](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NETunnelProviderSession) Autorelease() NETunnelProviderSession {
	rv := objc.Send[NETunnelProviderSession](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNETunnelProviderSession creates a new NETunnelProviderSession instance.
func NewNETunnelProviderSession() NETunnelProviderSession {
	return getNETunnelProviderSessionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NETunnelProviderSession */
// An object to start and stop a tunnel connection and get its status.
//
// objects control network tunnel connections provided by Tunnel Provider extensions. objects are not instantiated directly. Instead, each object has an associated as a read-only property.


// An object to start and stop a tunnel connection and get its status.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderSession
type NETunnelProviderSession struct {
	NEVPNConnection
}

// NETunnelProviderSessionFrom constructs a [NETunnelProviderSession] from an unsafe.Pointer.
//
// An object to start and stop a tunnel connection and get its status.
func NETunnelProviderSessionFrom(ptr unsafe.Pointer) NETunnelProviderSession {
	return NETunnelProviderSession{
		NEVPNConnection: NEVPNConnectionFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NETunnelProviderSession *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NETunnelProviderSession */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NETunnelProviderSession */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NETunnelProviderSession */

// Send a message to the Tunnel Provider extension. If the extension is not running, it should be launched to handle the message. If this method can’t start sending the message it reports an error in the parameter. If an error occurs while sending the message or returning the result, should be sent to the response handler as notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderSession/sendProviderMessage(_:responseHandler:)
func (n_ NETunnelProviderSession) SendProviderMessageReturnErrorResponseHandler(messageData objc.IObject /* cross-framework: NSData */, error_ objectivec.IObject, responseHandler unsafe.Pointer) bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("sendProviderMessage:returnError:responseHandler:"), messageData, error_, responseHandler)
	return rv
}/* debug [instance_methods/method]: SendProviderMessageReturnErrorResponseHandler */


// Start the process of connecting the tunnel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderSession/startTunnel(options:)
func (n_ NETunnelProviderSession) StartTunnelWithOptionsAndReturnError(options foundation.IDictionary, error_ objectivec.IObject) bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("startTunnelWithOptions:andReturnError:"), options, error_)
	return rv
}/* debug [instance_methods/method]: StartTunnelWithOptionsAndReturnError */


// Start the process of disconnecting the tunnel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderSession/stopTunnel()
func (n_ NETunnelProviderSession) StopTunnel() {
	objc.Send[objc.ID](n_.ID, objc.Sel("stopTunnel"))
}/* debug [instance_methods/method]: StopTunnel */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NETunnelProviderSession */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NETunnelProviderSession */



