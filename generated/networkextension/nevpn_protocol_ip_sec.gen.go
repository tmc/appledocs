// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [NEVPNProtocolIPSec] class.
var (
	NEVPNProtocolIPSecClass     _NEVPNProtocolIPSecClass
	NEVPNProtocolIPSecClassOnce sync.Once
)

func getNEVPNProtocolIPSecClass() _NEVPNProtocolIPSecClass {
	NEVPNProtocolIPSecClassOnce.Do(func() {
		NEVPNProtocolIPSecClass = _NEVPNProtocolIPSecClass{objc.GetClass("NEVPNProtocolIPSec")}
	})
	return NEVPNProtocolIPSecClass
}

type _NEVPNProtocolIPSecClass struct {
	class objc.Class
}





// An interface definition for the [NEVPNProtocolIPSec] class.
type INEVPNProtocolIPSec interface {
	INEVPNProtocol
	

	// properties:
	AuthenticationMethod() NEVPNIKEAuthenticationMethod
	SetAuthenticationMethod(value NEVPNIKEAuthenticationMethod)
	LocalIdentifier() foundation.foundation.INSString
	SetLocalIdentifier(value foundation.foundation.INSString)
	RemoteIdentifier() foundation.foundation.INSString
	SetRemoteIdentifier(value foundation.foundation.INSString)
	SharedSecretReference() foundation.foundation.INSData
	SetSharedSecretReference(value foundation.foundation.INSData)
	UseExtendedAuthentication() bool
	SetUseExtendedAuthentication(value bool)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (nc _NEVPNProtocolIPSecClass) Alloc() NEVPNProtocolIPSec {
	rv := objc.Send[NEVPNProtocolIPSec](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NEVPNProtocolIPSecClass) New() NEVPNProtocolIPSec {
	rv := objc.Send[NEVPNProtocolIPSec](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEVPNProtocolIPSec) Init() NEVPNProtocolIPSec {
	rv := objc.Send[NEVPNProtocolIPSec](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEVPNProtocolIPSec) Autorelease() NEVPNProtocolIPSec {
	rv := objc.Send[NEVPNProtocolIPSec](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEVPNProtocolIPSec creates a new NEVPNProtocolIPSec instance.
func NewNEVPNProtocolIPSec() NEVPNProtocolIPSec {
	return getNEVPNProtocolIPSecClass().New()
}





// Settings for an IPsec VPN configuration.
//
// To configure IKE version 2 (IKEv2), use the subclass. Instantiating directly implies IKE version 1.


// Settings for an IPsec VPN configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIPSec
type NEVPNProtocolIPSec struct {
	NEVPNProtocol
}

// NEVPNProtocolIPSecFrom constructs a [NEVPNProtocolIPSec] from an unsafe.Pointer.
//
// Settings for an IPsec VPN configuration.
func NEVPNProtocolIPSecFrom(ptr unsafe.Pointer) NEVPNProtocolIPSec {
	return NEVPNProtocolIPSec{
		NEVPNProtocol: NEVPNProtocolFrom(ptr),
	}
}

























// The method used to authenticate the device with the IPSec server. For IKE version 2, when using extended authentication, this authentication method only affects how the client validates the authentication payload presented by the server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIPSec/authenticationMethod
func (n_ NEVPNProtocolIPSec) AuthenticationMethod() NEVPNIKEAuthenticationMethod {
	rv := objc.Send[NEVPNIKEAuthenticationMethod](n_.ID, objc.Sel("authenticationMethod"))
	return rv
}


// The method used to authenticate the device with the IPSec server. For IKE version 2, when using extended authentication, this authentication method only affects how the client validates the authentication payload presented by the server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIPSec/authenticationMethod
func (n_ NEVPNProtocolIPSec) SetAuthenticationMethod(value NEVPNIKEAuthenticationMethod) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setAuthenticationMethod:"), value)
}


// A string identifying the iOS or macOS device for authentication purposes
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIPSec/localIdentifier
func (n_ NEVPNProtocolIPSec) LocalIdentifier() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("localIdentifier"))
	return rv
}


// A string identifying the iOS or macOS device for authentication purposes
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIPSec/localIdentifier
func (n_ NEVPNProtocolIPSec) SetLocalIdentifier(value foundation.foundation.INSString) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setLocalIdentifier:"), value)
}


// A string identifying the IPSec server for authentication purposes
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIPSec/remoteIdentifier
func (n_ NEVPNProtocolIPSec) RemoteIdentifier() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("remoteIdentifier"))
	return rv
}


// A string identifying the IPSec server for authentication purposes
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIPSec/remoteIdentifier
func (n_ NEVPNProtocolIPSec) SetRemoteIdentifier(value foundation.foundation.INSString) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setRemoteIdentifier:"), value)
}


// A persistent keychain reference to a keychain item containing the IKE shared secret.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIPSec/sharedSecretReference
func (n_ NEVPNProtocolIPSec) SharedSecretReference() foundation.foundation.INSData {
	rv := objc.Send[foundation.NSData](n_.ID, objc.Sel("sharedSecretReference"))
	return rv
}


// A persistent keychain reference to a keychain item containing the IKE shared secret.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIPSec/sharedSecretReference
func (n_ NEVPNProtocolIPSec) SetSharedSecretReference(value foundation.foundation.INSData) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setSharedSecretReference:"), value)
}


// A flag indicating if extended authentication will be negotiated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIPSec/useExtendedAuthentication
func (n_ NEVPNProtocolIPSec) UseExtendedAuthentication() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("useExtendedAuthentication"))
	return rv
}


// A flag indicating if extended authentication will be negotiated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIPSec/useExtendedAuthentication
func (n_ NEVPNProtocolIPSec) SetUseExtendedAuthentication(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setUseExtendedAuthentication:"), value)
}








