// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [NEProxyServer] class.
type INEProxyServer interface {
	objectivec.IObject
	

	// properties:
	Address() foundation.foundation.INSString
	AuthenticationRequired() bool
	SetAuthenticationRequired(value bool)
	Password() foundation.foundation.INSString
	SetPassword(value foundation.foundation.INSString)
	Port() int
	Username() foundation.foundation.INSString
	SetUsername(value foundation.foundation.INSString)
	HttpEnabled() bool
	SetHttpEnabled(value bool)
	HttpServer() INEProxyServer
	SetHttpServer(value INEProxyServer)
	HttpsEnabled() bool
	SetHttpsEnabled(value bool)
	HttpsServer() INEProxyServer
	SetHttpsServer(value INEProxyServer)


	

	// methods:


}





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






// Initialize a newly-allocated object
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProxyServer/init(address:port:)
func NewNEProxyServerWithAddressPort(address foundation.foundation.INSString, port int) NEProxyServer {
	instance := getNEProxyServerClass().Alloc()
	rv := objc.Send[NEProxyServer](instance.ID, objc.Sel("initWithAddress:port:"), address, port)
	rv.Autorelease()
	return rv
}






















// The address of the proxy server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProxyServer/address
func (n_ NEProxyServer) Address() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("address"))
	return rv
}


// A Boolean indicating if the server requires authentication credentials.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProxyServer/authenticationRequired
func (n_ NEProxyServer) AuthenticationRequired() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("authenticationRequired"))
	return rv
}


// A Boolean indicating if the server requires authentication credentials.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProxyServer/authenticationRequired
func (n_ NEProxyServer) SetAuthenticationRequired(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setAuthenticationRequired:"), value)
}


// The password portion of the authentication credential to be used to authenticate with the proxy server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProxyServer/password
func (n_ NEProxyServer) Password() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("password"))
	return rv
}


// The password portion of the authentication credential to be used to authenticate with the proxy server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProxyServer/password
func (n_ NEProxyServer) SetPassword(value foundation.foundation.INSString) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setPassword:"), value)
}


// The TCP port on which the proxy server is listening for connections.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProxyServer/port
func (n_ NEProxyServer) Port() int {
	rv := objc.Send[int](n_.ID, objc.Sel("port"))
	return rv
}


// The username portion of the authentication credential to be used to authenticate with the proxy server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProxyServer/username
func (n_ NEProxyServer) Username() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("username"))
	return rv
}


// The username portion of the authentication credential to be used to authenticate with the proxy server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProxyServer/username
func (n_ NEProxyServer) SetUsername(value foundation.foundation.INSString) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setUsername:"), value)
}


// A Boolean indicating if a static HTTP proxy will be used.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neproxysettings/httpenabled
func (n_ NEProxyServer) HttpEnabled() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("httpEnabled"))
	return rv
}


// A Boolean indicating if a static HTTP proxy will be used.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neproxysettings/httpenabled
func (n_ NEProxyServer) SetHttpEnabled(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setHttpEnabled:"), value)
}


// An
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neproxysettings/httpserver
func (n_ NEProxyServer) HttpServer() INEProxyServer {
	rv := objc.Send[NEProxyServer](n_.ID, objc.Sel("httpServer"))
	return rv
}


// An
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neproxysettings/httpserver
func (n_ NEProxyServer) SetHttpServer(value INEProxyServer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setHttpServer:"), value)
}


// A Boolean indicating if a static HTTPS proxy will be used.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neproxysettings/httpsenabled
func (n_ NEProxyServer) HttpsEnabled() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("httpsEnabled"))
	return rv
}


// A Boolean indicating if a static HTTPS proxy will be used.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neproxysettings/httpsenabled
func (n_ NEProxyServer) SetHttpsEnabled(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setHttpsEnabled:"), value)
}


// An
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neproxysettings/httpsserver
func (n_ NEProxyServer) HttpsServer() INEProxyServer {
	rv := objc.Send[NEProxyServer](n_.ID, objc.Sel("httpsServer"))
	return rv
}


// An
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neproxysettings/httpsserver
func (n_ NEProxyServer) SetHttpsServer(value INEProxyServer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setHttpsServer:"), value)
}







