// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NEVPNConnection */


/* debug [class_header]: Header for NEVPNConnection */
// The class instance for the [NEVPNConnection] class.
var (
	NEVPNConnectionClass     _NEVPNConnectionClass
	NEVPNConnectionClassOnce sync.Once
)

func getNEVPNConnectionClass() _NEVPNConnectionClass {
	NEVPNConnectionClassOnce.Do(func() {
		NEVPNConnectionClass = _NEVPNConnectionClass{objc.GetClass("NEVPNConnection")}
	})
	return NEVPNConnectionClass
}

type _NEVPNConnectionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NEVPNConnection */
// An interface definition for the [NEVPNConnection] class.
type INEVPNConnection interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for NEVPNConnection */
	// properties:
	ConnectedDate() objc.IObject /* cross-framework: NSDate */
	Manager() INEVPNManager
	Status() NEVPNStatus
	NEVPNConnectionErrorDomain() objc.IObject /* cross-framework: NSString */
	NEVPNConnectionStartOptionPassword() objc.IObject /* cross-framework: NSString */
	NEVPNConnectionStartOptionUsername() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NEVPNConnection */
	// methods:
	FetchLastDisconnectErrorWithCompletionHandler(handler unsafe.Pointer)
	StartVPNTunnelAndReturnError(error_ objectivec.IObject) bool
	StartVPNTunnelWithOptionsAndReturnError(options foundation.IDictionary, error_ objectivec.IObject) bool
	StopVPNTunnel()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NEVPNConnection */
// Alloc allocates a new instance without initialization.
func (nc _NEVPNConnectionClass) Alloc() NEVPNConnection {
	rv := objc.Send[NEVPNConnection](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NEVPNConnectionClass) New() NEVPNConnection {
	rv := objc.Send[NEVPNConnection](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEVPNConnection) Init() NEVPNConnection {
	rv := objc.Send[NEVPNConnection](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEVPNConnection) Autorelease() NEVPNConnection {
	rv := objc.Send[NEVPNConnection](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEVPNConnection creates a new NEVPNConnection instance.
func NewNEVPNConnection() NEVPNConnection {
	return getNEVPNConnectionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NEVPNConnection */
// An object to start and stop a Personal VPN connection and get its status.
//
// objects are not instantiated directly. Instead, each object has an associated object as a read-only property. The class provides methods for starting and stopping the VPN programmatically. The other way that the VPN can be started and stopped is through VPN On Demand. See the property in and . Instances of this class are thread safe.


// An object to start and stop a Personal VPN connection and get its status.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNConnection
type NEVPNConnection struct {
	objectivec.Object
}

// NEVPNConnectionFrom constructs a [NEVPNConnection] from an unsafe.Pointer.
//
// An object to start and stop a Personal VPN connection and get its status.
func NEVPNConnectionFrom(ptr unsafe.Pointer) NEVPNConnection {
	return NEVPNConnection{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NEVPNConnection *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NEVPNConnection */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NEVPNConnection */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NEVPNConnection */

// Retrives the most recent error that caused the VPN to disconnect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNConnection/fetchLastDisconnectError(completionHandler:)
func (n_ NEVPNConnection) FetchLastDisconnectErrorWithCompletionHandler(handler unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("fetchLastDisconnectErrorWithCompletionHandler:"), handler)
}/* debug [instance_methods/method]: FetchLastDisconnectErrorWithCompletionHandler */


// Start the process of connecting the VPN.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNConnection/startVPNTunnel()
func (n_ NEVPNConnection) StartVPNTunnelAndReturnError(error_ objectivec.IObject) bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("startVPNTunnelAndReturnError:"), error_)
	return rv
}/* debug [instance_methods/method]: StartVPNTunnelAndReturnError */


// Start the process of connecting the VPN.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNConnection/startVPNTunnel(options:)
func (n_ NEVPNConnection) StartVPNTunnelWithOptionsAndReturnError(options foundation.IDictionary, error_ objectivec.IObject) bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("startVPNTunnelWithOptions:andReturnError:"), options, error_)
	return rv
}/* debug [instance_methods/method]: StartVPNTunnelWithOptionsAndReturnError */


// Start the process of disconnecting the VPN.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNConnection/stopVPNTunnel()
func (n_ NEVPNConnection) StopVPNTunnel() {
	objc.Send[objc.ID](n_.ID, objc.Sel("stopVPNTunnel"))
}/* debug [instance_methods/method]: StopVPNTunnel */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NEVPNConnection */

// The date and time when the connection status changed to .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNConnection/connectedDate
func (n_ NEVPNConnection) ConnectedDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](n_.ID, objc.Sel("connectedDate"))
	return rv
}/* debug [instance_properties/getter]: connectedDate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNConnection/manager
func (n_ NEVPNConnection) Manager() INEVPNManager {
	rv := objc.Send[NEVPNManager](n_.ID, objc.Sel("manager"))
	return rv
}/* debug [instance_properties/getter]: manager */


// The current status of the VPN connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNConnection/status
func (n_ NEVPNConnection) Status() NEVPNStatus {
	rv := objc.Send[NEVPNStatus](n_.ID, objc.Sel("status"))
	return rv
}/* debug [instance_properties/getter]: status */


// The domain for errors resulting from VPN connection calls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nevpnconnectionerrordomain
func (n_ NEVPNConnection) NEVPNConnectionErrorDomain() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("NEVPNConnectionErrorDomain"))
	return rv
}/* debug [instance_properties/getter]: NEVPNConnectionErrorDomain */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nevpnconnectionstartoptionpassword
func (n_ NEVPNConnection) NEVPNConnectionStartOptionPassword() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("NEVPNConnectionStartOptionPassword"))
	return rv
}/* debug [instance_properties/getter]: NEVPNConnectionStartOptionPassword */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nevpnconnectionstartoptionusername
func (n_ NEVPNConnection) NEVPNConnectionStartOptionUsername() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("NEVPNConnectionStartOptionUsername"))
	return rv
}/* debug [instance_properties/getter]: NEVPNConnectionStartOptionUsername */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NEVPNConnection */



