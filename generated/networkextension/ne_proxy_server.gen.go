// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NEProxyServer */


/* debug [class_header]: Header for NEProxyServer */
// The class instance for the [NEProxyServer] class.
var (
	NEProxyServerClass     _NEProxyServerClass
	NEProxyServerClassOnce sync.Once
)

func getNEProxyServerClass() _NEProxyServerClass {
	NEProxyServerClassOnce.Do(func() {
		NEProxyServerClass = _NEProxyServerClass{objc.GetClass("NEProxyServer")}
	})
	return NEProxyServerClass
}

type _NEProxyServerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NEProxyServer */
// An interface definition for the [NEProxyServer] class.
type INEProxyServer interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for NEProxyServer */
	// properties:
	Address() objc.IObject /* cross-framework: NSString */
	AuthenticationRequired() bool
	SetAuthenticationRequired(value bool)
	Password() objc.IObject /* cross-framework: NSString */
	SetPassword(value objc.IObject /* cross-framework: NSString */)
	Port() int
	Username() objc.IObject /* cross-framework: NSString */
	SetUsername(value objc.IObject /* cross-framework: NSString */)
	HttpEnabled() bool
	SetHttpEnabled(value bool)
	HttpServer() INEProxyServer
	SetHttpServer(value INEProxyServer)
	HttpsEnabled() bool
	SetHttpsEnabled(value bool)
	HttpsServer() INEProxyServer
	SetHttpsServer(value INEProxyServer)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NEProxyServer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NEProxyServer */
// Alloc allocates a new instance without initialization.
func (nc _NEProxyServerClass) Alloc() NEProxyServer {
	rv := objc.Send[NEProxyServer](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NEProxyServerClass) New() NEProxyServer {
	rv := objc.Send[NEProxyServer](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEProxyServer) Init() NEProxyServer {
	rv := objc.Send[NEProxyServer](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEProxyServer) Autorelease() NEProxyServer {
	rv := objc.Send[NEProxyServer](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEProxyServer creates a new NEProxyServer instance.
func NewNEProxyServer() NEProxyServer {
	return getNEProxyServerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NEProxyServer */
// contains settings for a proxy server.
//
// instances are used inside of instances to configure proxy settings for VPN connections.


// contains settings for a proxy server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProxyServer
type NEProxyServer struct {
	objectivec.Object
}

// NEProxyServerFrom constructs a [NEProxyServer] from an unsafe.Pointer.
//
// contains settings for a proxy server.
func NEProxyServerFrom(ptr unsafe.Pointer) NEProxyServer {
	return NEProxyServer{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NEProxyServer */

// Initialize a newly-allocated object
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProxyServer/init(address:port:)
func NewNEProxyServerWithAddressPort(address objc.IObject /* cross-framework: NSString */, port int) NEProxyServer {
	instance := getNEProxyServerClass().Alloc()
	rv := objc.Send[NEProxyServer](instance.ID, objc.Sel("initWithAddress:port:"), address, port)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewNEProxyServerWithAddressPort */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NEProxyServer */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NEProxyServer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NEProxyServer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NEProxyServer */

// The address of the proxy server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProxyServer/address
func (n_ NEProxyServer) Address() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("address"))
	return rv
}/* debug [instance_properties/getter]: address */


// A Boolean indicating if the server requires authentication credentials.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProxyServer/authenticationRequired
func (n_ NEProxyServer) AuthenticationRequired() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("authenticationRequired"))
	return rv
}/* debug [instance_properties/getter]: authenticationRequired */


// A Boolean indicating if the server requires authentication credentials.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProxyServer/authenticationRequired
func (n_ NEProxyServer) SetAuthenticationRequired(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setAuthenticationRequired:"), value)
}/* debug [instance_properties/setter]: authenticationRequired */


// The password portion of the authentication credential to be used to authenticate with the proxy server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProxyServer/password
func (n_ NEProxyServer) Password() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("password"))
	return rv
}/* debug [instance_properties/getter]: password */


// The password portion of the authentication credential to be used to authenticate with the proxy server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProxyServer/password
func (n_ NEProxyServer) SetPassword(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setPassword:"), value)
}/* debug [instance_properties/setter]: password */


// The TCP port on which the proxy server is listening for connections.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProxyServer/port
func (n_ NEProxyServer) Port() int {
	rv := objc.Send[int](n_.ID, objc.Sel("port"))
	return rv
}/* debug [instance_properties/getter]: port */


// The username portion of the authentication credential to be used to authenticate with the proxy server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProxyServer/username
func (n_ NEProxyServer) Username() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("username"))
	return rv
}/* debug [instance_properties/getter]: username */


// The username portion of the authentication credential to be used to authenticate with the proxy server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProxyServer/username
func (n_ NEProxyServer) SetUsername(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setUsername:"), value)
}/* debug [instance_properties/setter]: username */


// A Boolean indicating if a static HTTP proxy will be used.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neproxysettings/httpenabled
func (n_ NEProxyServer) HttpEnabled() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("httpEnabled"))
	return rv
}/* debug [instance_properties/getter]: httpEnabled */


// A Boolean indicating if a static HTTP proxy will be used.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neproxysettings/httpenabled
func (n_ NEProxyServer) SetHttpEnabled(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setHttpEnabled:"), value)
}/* debug [instance_properties/setter]: httpEnabled */


// An
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neproxysettings/httpserver
func (n_ NEProxyServer) HttpServer() INEProxyServer {
	rv := objc.Send[NEProxyServer](n_.ID, objc.Sel("httpServer"))
	return rv
}/* debug [instance_properties/getter]: httpServer */


// An
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neproxysettings/httpserver
func (n_ NEProxyServer) SetHttpServer(value INEProxyServer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setHttpServer:"), value)
}/* debug [instance_properties/setter]: httpServer */


// A Boolean indicating if a static HTTPS proxy will be used.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neproxysettings/httpsenabled
func (n_ NEProxyServer) HttpsEnabled() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("httpsEnabled"))
	return rv
}/* debug [instance_properties/getter]: httpsEnabled */


// A Boolean indicating if a static HTTPS proxy will be used.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neproxysettings/httpsenabled
func (n_ NEProxyServer) SetHttpsEnabled(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setHttpsEnabled:"), value)
}/* debug [instance_properties/setter]: httpsEnabled */


// An
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neproxysettings/httpsserver
func (n_ NEProxyServer) HttpsServer() INEProxyServer {
	rv := objc.Send[NEProxyServer](n_.ID, objc.Sel("httpsServer"))
	return rv
}/* debug [instance_properties/getter]: httpsServer */


// An
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neproxysettings/httpsserver
func (n_ NEProxyServer) SetHttpsServer(value INEProxyServer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setHttpsServer:"), value)
}/* debug [instance_properties/setter]: httpsServer */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NEProxyServer */


