// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [URLProtectionSpace] class.
var (
	URLProtectionSpaceClass     _URLProtectionSpaceClass
	URLProtectionSpaceClassOnce sync.Once
)

func getURLProtectionSpaceClass() _URLProtectionSpaceClass {
	URLProtectionSpaceClassOnce.Do(func() {
		URLProtectionSpaceClass = _URLProtectionSpaceClass{objc.GetClass("NSURLProtectionSpace")}
	})
	return URLProtectionSpaceClass
}

type _URLProtectionSpaceClass struct {
	class objc.Class
}

// An interface definition for the [URLProtectionSpace] class.
type IURLProtectionSpace interface {
	objectivec.IObject
	// properties:
	IsProxy() bool /* primitive/slice/pointer. */
	AuthenticationMethod() IString
	DistinguishedNames() []Data /* primitive/slice/pointer. */
	Host() IString
	Port() int /* primitive/slice/pointer. */
	Protocol() IString
	ProxyType() IString
	Realm() IString
	ReceivesCredentialSecurely() bool /* primitive/slice/pointer. */
	ServerTrust() unsafe.Pointer
	// methods:
}

// A server or an area on a server, commonly referred to as a realm, that requires authentication.
//
// A protection space defines a series of matching constraints that determine which credential should be provided. For example, if a request provides your delegate with a object that requests a client username and password, your app should provide the correct username and password for the particular host, port, protocol, and realm, as specified in the challenge’s protection space.


// A server or an area on a server, commonly referred to as a realm, that requires authentication.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLProtectionSpace
type URLProtectionSpace struct {
	objectivec.Object
}

// URLProtectionSpaceFrom constructs a [URLProtectionSpace] from an unsafe.Pointer.
//
// A server or an area on a server, commonly referred to as a realm, that requires authentication.
func URLProtectionSpaceFrom(ptr unsafe.Pointer) URLProtectionSpace {
	return URLProtectionSpace{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (uc _URLProtectionSpaceClass) Alloc() URLProtectionSpace {
	rv := objc.Send[URLProtectionSpace](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _URLProtectionSpaceClass) New() URLProtectionSpace {
	rv := objc.Send[URLProtectionSpace](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ URLProtectionSpace) Init() URLProtectionSpace {
	rv := objc.Send[URLProtectionSpace](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ URLProtectionSpace) Autorelease() URLProtectionSpace {
	rv := objc.Send[URLProtectionSpace](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewURLProtectionSpace creates a new URLProtectionSpace instance.
func NewURLProtectionSpace() URLProtectionSpace {
	return getURLProtectionSpaceClass().New()
}



// Creates a protection space object from the given host, port, protocol, realm, and authentication method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLProtectionSpace/init(host:port:protocol:realm:authenticationMethod:)
func NewURLProtectionSpaceWithHostPortProtocolRealmAuthenticationMethod(host IString, port int /* primitive/slice/pointer. */, protocol_ IString, realm IString, authenticationMethod IString) URLProtectionSpace {
	instance := getURLProtectionSpaceClass().Alloc()
	rv := objc.Send[URLProtectionSpace](instance.ID, objc.Sel("initWithHost:port:protocol:realm:authenticationMethod:"), host, port, protocol_, realm, authenticationMethod)
	rv.Autorelease()
	return rv
}


// Creates a protection space object representing a proxy server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLProtectionSpace/init(proxyHost:port:type:realm:authenticationMethod:)
func NewURLProtectionSpaceWithProxyHostPortTypeRealmAuthenticationMethod(host IString, port int /* primitive/slice/pointer. */, type_ IString, realm IString, authenticationMethod IString) URLProtectionSpace {
	instance := getURLProtectionSpaceClass().Alloc()
	rv := objc.Send[URLProtectionSpace](instance.ID, objc.Sel("initWithProxyHost:port:type:realm:authenticationMethod:"), host, port, type_, realm, authenticationMethod)
	rv.Autorelease()
	return rv
}



// A Boolean value that indicates whether the receiver represents a proxy server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLProtectionSpace/isProxy
func (u_ URLProtectionSpace) IsProxy() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](u_.ID, objc.Sel("isProxy"))
	return rv
}


// The authentication method used by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLProtectionSpace/authenticationMethod
func (u_ URLProtectionSpace) AuthenticationMethod() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("authenticationMethod"))
	return rv
}


// The acceptable certificate-issuing authorities for client certificate authentication.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLProtectionSpace/distinguishedNames
func (u_ URLProtectionSpace) DistinguishedNames() []Data /* primitive/slice/pointer. */ {
	rv := objc.Send[[]Data](u_.ID, objc.Sel("distinguishedNames"))
	return rv
}


// The receiver’s host.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLProtectionSpace/host
func (u_ URLProtectionSpace) Host() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("host"))
	return rv
}


// The receiver’s port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLProtectionSpace/port
func (u_ URLProtectionSpace) Port() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](u_.ID, objc.Sel("port"))
	return rv
}


// The receiver’s protocol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLProtectionSpace/protocol
func (u_ URLProtectionSpace) Protocol() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("protocol"))
	return rv
}


// The receiver’s proxy type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLProtectionSpace/proxyType
func (u_ URLProtectionSpace) ProxyType() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("proxyType"))
	return rv
}


// The receiver’s authentication realm
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLProtectionSpace/realm
func (u_ URLProtectionSpace) Realm() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("realm"))
	return rv
}


// A Boolean value that indicates whether the credentials for the protection space can be sent securely.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLProtectionSpace/receivesCredentialSecurely
func (u_ URLProtectionSpace) ReceivesCredentialSecurely() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](u_.ID, objc.Sel("receivesCredentialSecurely"))
	return rv
}


// A representation of the server’s SSL transaction state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLProtectionSpace/serverTrust
func (u_ URLProtectionSpace) ServerTrust() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("serverTrust"))
	return rv
}


