// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class NEVPNProtocolIPSec */


/* debug [class_header]: Header for NEVPNProtocolIPSec */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NEVPNProtocolIPSec */
// An interface definition for the [NEVPNProtocolIPSec] class.
type INEVPNProtocolIPSec interface {
	INEVPNProtocol
	
/* debug [class_interface_properties]: Properties for NEVPNProtocolIPSec */
	// properties:
	AuthenticationMethod() NEVPNIKEAuthenticationMethod
	SetAuthenticationMethod(value NEVPNIKEAuthenticationMethod)
	LocalIdentifier() objc.IObject /* cross-framework: NSString */
	SetLocalIdentifier(value objc.IObject /* cross-framework: NSString */)
	RemoteIdentifier() objc.IObject /* cross-framework: NSString */
	SetRemoteIdentifier(value objc.IObject /* cross-framework: NSString */)
	SharedSecretReference() objc.IObject /* cross-framework: NSData */
	SetSharedSecretReference(value objc.IObject /* cross-framework: NSData */)
	UseExtendedAuthentication() bool
	SetUseExtendedAuthentication(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NEVPNProtocolIPSec */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NEVPNProtocolIPSec */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NEVPNProtocolIPSec */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NEVPNProtocolIPSec *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NEVPNProtocolIPSec */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NEVPNProtocolIPSec */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NEVPNProtocolIPSec */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NEVPNProtocolIPSec */

// The method used to authenticate the device with the IPSec server. For IKE version 2, when using extended authentication, this authentication method only affects how the client validates the authentication payload presented by the server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIPSec/authenticationMethod
func (n_ NEVPNProtocolIPSec) AuthenticationMethod() NEVPNIKEAuthenticationMethod {
	rv := objc.Send[NEVPNIKEAuthenticationMethod](n_.ID, objc.Sel("authenticationMethod"))
	return rv
}/* debug [instance_properties/getter]: authenticationMethod */


// The method used to authenticate the device with the IPSec server. For IKE version 2, when using extended authentication, this authentication method only affects how the client validates the authentication payload presented by the server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIPSec/authenticationMethod
func (n_ NEVPNProtocolIPSec) SetAuthenticationMethod(value NEVPNIKEAuthenticationMethod) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setAuthenticationMethod:"), value)
}/* debug [instance_properties/setter]: authenticationMethod */


// A string identifying the iOS or macOS device for authentication purposes
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIPSec/localIdentifier
func (n_ NEVPNProtocolIPSec) LocalIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("localIdentifier"))
	return rv
}/* debug [instance_properties/getter]: localIdentifier */


// A string identifying the iOS or macOS device for authentication purposes
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIPSec/localIdentifier
func (n_ NEVPNProtocolIPSec) SetLocalIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setLocalIdentifier:"), value)
}/* debug [instance_properties/setter]: localIdentifier */


// A string identifying the IPSec server for authentication purposes
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIPSec/remoteIdentifier
func (n_ NEVPNProtocolIPSec) RemoteIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("remoteIdentifier"))
	return rv
}/* debug [instance_properties/getter]: remoteIdentifier */


// A string identifying the IPSec server for authentication purposes
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIPSec/remoteIdentifier
func (n_ NEVPNProtocolIPSec) SetRemoteIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setRemoteIdentifier:"), value)
}/* debug [instance_properties/setter]: remoteIdentifier */


// A persistent keychain reference to a keychain item containing the IKE shared secret.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIPSec/sharedSecretReference
func (n_ NEVPNProtocolIPSec) SharedSecretReference() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](n_.ID, objc.Sel("sharedSecretReference"))
	return rv
}/* debug [instance_properties/getter]: sharedSecretReference */


// A persistent keychain reference to a keychain item containing the IKE shared secret.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIPSec/sharedSecretReference
func (n_ NEVPNProtocolIPSec) SetSharedSecretReference(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setSharedSecretReference:"), value)
}/* debug [instance_properties/setter]: sharedSecretReference */


// A flag indicating if extended authentication will be negotiated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIPSec/useExtendedAuthentication
func (n_ NEVPNProtocolIPSec) UseExtendedAuthentication() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("useExtendedAuthentication"))
	return rv
}/* debug [instance_properties/getter]: useExtendedAuthentication */


// A flag indicating if extended authentication will be negotiated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIPSec/useExtendedAuthentication
func (n_ NEVPNProtocolIPSec) SetUseExtendedAuthentication(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setUseExtendedAuthentication:"), value)
}/* debug [instance_properties/setter]: useExtendedAuthentication */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NEVPNProtocolIPSec */



