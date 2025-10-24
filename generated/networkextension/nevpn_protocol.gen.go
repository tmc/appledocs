// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NEVPNProtocol */


/* debug [class_header]: Header for NEVPNProtocol */
// The class instance for the [NEVPNProtocol] class.
var (
	NEVPNProtocolClass     _NEVPNProtocolClass
	NEVPNProtocolClassOnce sync.Once
)

func getNEVPNProtocolClass() _NEVPNProtocolClass {
	NEVPNProtocolClassOnce.Do(func() {
		NEVPNProtocolClass = _NEVPNProtocolClass{objc.GetClass("NEVPNProtocol")}
	})
	return NEVPNProtocolClass
}

type _NEVPNProtocolClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NEVPNProtocol */
// An interface definition for the [NEVPNProtocol] class.
type INEVPNProtocol interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for NEVPNProtocol */
	// properties:
	DisconnectOnSleep() bool
	SetDisconnectOnSleep(value bool)
	EnforceRoutes() bool
	SetEnforceRoutes(value bool)
	ExcludeAPNs() bool
	SetExcludeAPNs(value bool)
	ExcludeCellularServices() bool
	SetExcludeCellularServices(value bool)
	ExcludeDeviceCommunication() bool
	SetExcludeDeviceCommunication(value bool)
	ExcludeLocalNetworks() bool
	SetExcludeLocalNetworks(value bool)
	IdentityData() objc.IObject /* cross-framework: NSData */
	SetIdentityData(value objc.IObject /* cross-framework: NSData */)
	IdentityDataPassword() objc.IObject /* cross-framework: NSString */
	SetIdentityDataPassword(value objc.IObject /* cross-framework: NSString */)
	IdentityReference() objc.IObject /* cross-framework: NSData */
	SetIdentityReference(value objc.IObject /* cross-framework: NSData */)
	IncludeAllNetworks() bool
	SetIncludeAllNetworks(value bool)
	PasswordReference() objc.IObject /* cross-framework: NSData */
	SetPasswordReference(value objc.IObject /* cross-framework: NSData */)
	ProxySettings() INEProxySettings
	SetProxySettings(value INEProxySettings)
	ServerAddress() objc.IObject /* cross-framework: NSString */
	SetServerAddress(value objc.IObject /* cross-framework: NSString */)
	Username() objc.IObject /* cross-framework: NSString */
	SetUsername(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NEVPNProtocol */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NEVPNProtocol */
// Alloc allocates a new instance without initialization.
func (nc _NEVPNProtocolClass) Alloc() NEVPNProtocol {
	rv := objc.Send[NEVPNProtocol](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NEVPNProtocolClass) New() NEVPNProtocol {
	rv := objc.Send[NEVPNProtocol](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEVPNProtocol) Init() NEVPNProtocol {
	rv := objc.Send[NEVPNProtocol](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEVPNProtocol) Autorelease() NEVPNProtocol {
	rv := objc.Send[NEVPNProtocol](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEVPNProtocol creates a new NEVPNProtocol instance.
func NewNEVPNProtocol() NEVPNProtocol {
	return getNEVPNProtocolClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NEVPNProtocol */
// Settings common to both IKEv2 and IPsec VPN configurations.
//
// The class is an abstract base class with one subclass for each type of supported VPN configuration. This class provides properties for configuring the VPN, authenticating network connections, and routing network traffic. You can include all network traffic, with some exceptions, and selectively exclude types of network traffic. Instances of this class are thread-safe.


// Settings common to both IKEv2 and IPsec VPN configurations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocol
type NEVPNProtocol struct {
	objectivec.Object
}

// NEVPNProtocolFrom constructs a [NEVPNProtocol] from an unsafe.Pointer.
//
// Settings common to both IKEv2 and IPsec VPN configurations.
func NEVPNProtocolFrom(ptr unsafe.Pointer) NEVPNProtocol {
	return NEVPNProtocol{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NEVPNProtocol *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NEVPNProtocol */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NEVPNProtocol */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NEVPNProtocol */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NEVPNProtocol */

// A Boolean value that indicates whether the VPN disconnects when the device sleeps.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocol/disconnectOnSleep
func (n_ NEVPNProtocol) DisconnectOnSleep() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("disconnectOnSleep"))
	return rv
}/* debug [instance_properties/getter]: disconnectOnSleep */


// A Boolean value that indicates whether the VPN disconnects when the device sleeps.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocol/disconnectOnSleep
func (n_ NEVPNProtocol) SetDisconnectOnSleep(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDisconnectOnSleep:"), value)
}/* debug [instance_properties/setter]: disconnectOnSleep */


// A Boolean value that indicates whether route rules for the tunnel take precedence over any locally defined routes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocol/enforceRoutes
func (n_ NEVPNProtocol) EnforceRoutes() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("enforceRoutes"))
	return rv
}/* debug [instance_properties/getter]: enforceRoutes */


// A Boolean value that indicates whether route rules for the tunnel take precedence over any locally defined routes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocol/enforceRoutes
func (n_ NEVPNProtocol) SetEnforceRoutes(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setEnforceRoutes:"), value)
}/* debug [instance_properties/setter]: enforceRoutes */


// A Boolean value that indicates whether the system excludes all APNs network traffic from the tunnel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocol/excludeAPNs
func (n_ NEVPNProtocol) ExcludeAPNs() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("excludeAPNs"))
	return rv
}/* debug [instance_properties/getter]: excludeAPNs */


// A Boolean value that indicates whether the system excludes all APNs network traffic from the tunnel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocol/excludeAPNs
func (n_ NEVPNProtocol) SetExcludeAPNs(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setExcludeAPNs:"), value)
}/* debug [instance_properties/setter]: excludeAPNs */


// A Boolean value that indicates whether the system excludes all cellular services network traffic from the tunnel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocol/excludeCellularServices
func (n_ NEVPNProtocol) ExcludeCellularServices() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("excludeCellularServices"))
	return rv
}/* debug [instance_properties/getter]: excludeCellularServices */


// A Boolean value that indicates whether the system excludes all cellular services network traffic from the tunnel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocol/excludeCellularServices
func (n_ NEVPNProtocol) SetExcludeCellularServices(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setExcludeCellularServices:"), value)
}/* debug [instance_properties/setter]: excludeCellularServices */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocol/excludeDeviceCommunication
func (n_ NEVPNProtocol) ExcludeDeviceCommunication() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("excludeDeviceCommunication"))
	return rv
}/* debug [instance_properties/getter]: excludeDeviceCommunication */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocol/excludeDeviceCommunication
func (n_ NEVPNProtocol) SetExcludeDeviceCommunication(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setExcludeDeviceCommunication:"), value)
}/* debug [instance_properties/setter]: excludeDeviceCommunication */


// A Boolean value that indicates whether the system excludes all traffic destined for local networks from the tunnel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocol/excludeLocalNetworks
func (n_ NEVPNProtocol) ExcludeLocalNetworks() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("excludeLocalNetworks"))
	return rv
}/* debug [instance_properties/getter]: excludeLocalNetworks */


// A Boolean value that indicates whether the system excludes all traffic destined for local networks from the tunnel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocol/excludeLocalNetworks
func (n_ NEVPNProtocol) SetExcludeLocalNetworks(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setExcludeLocalNetworks:"), value)
}/* debug [instance_properties/setter]: excludeLocalNetworks */


// The certificate and private key components of the tunneling protocol authentication credential, in PKCS12 format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocol/identityData
func (n_ NEVPNProtocol) IdentityData() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](n_.ID, objc.Sel("identityData"))
	return rv
}/* debug [instance_properties/getter]: identityData */


// The certificate and private key components of the tunneling protocol authentication credential, in PKCS12 format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocol/identityData
func (n_ NEVPNProtocol) SetIdentityData(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIdentityData:"), value)
}/* debug [instance_properties/setter]: identityData */


// The password for the PKCS12 tunneling protocol authentication credentials.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocol/identityDataPassword
func (n_ NEVPNProtocol) IdentityDataPassword() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("identityDataPassword"))
	return rv
}/* debug [instance_properties/getter]: identityDataPassword */


// The password for the PKCS12 tunneling protocol authentication credentials.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocol/identityDataPassword
func (n_ NEVPNProtocol) SetIdentityDataPassword(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIdentityDataPassword:"), value)
}/* debug [instance_properties/setter]: identityDataPassword */


// A persistent keychain reference to a keychain item containing the certificate and private key components of the tunneling protocol authentication credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocol/identityReference
func (n_ NEVPNProtocol) IdentityReference() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](n_.ID, objc.Sel("identityReference"))
	return rv
}/* debug [instance_properties/getter]: identityReference */


// A persistent keychain reference to a keychain item containing the certificate and private key components of the tunneling protocol authentication credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocol/identityReference
func (n_ NEVPNProtocol) SetIdentityReference(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIdentityReference:"), value)
}/* debug [instance_properties/setter]: identityReference */


// A Boolean value that indicates whether the system sends most network traffic over the tunnel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocol/includeAllNetworks
func (n_ NEVPNProtocol) IncludeAllNetworks() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("includeAllNetworks"))
	return rv
}/* debug [instance_properties/getter]: includeAllNetworks */


// A Boolean value that indicates whether the system sends most network traffic over the tunnel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocol/includeAllNetworks
func (n_ NEVPNProtocol) SetIncludeAllNetworks(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIncludeAllNetworks:"), value)
}/* debug [instance_properties/setter]: includeAllNetworks */


// A persistent keychain reference to a keychain item containing the password component of the tunneling protocol authentication credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocol/passwordReference
func (n_ NEVPNProtocol) PasswordReference() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](n_.ID, objc.Sel("passwordReference"))
	return rv
}/* debug [instance_properties/getter]: passwordReference */


// A persistent keychain reference to a keychain item containing the password component of the tunneling protocol authentication credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocol/passwordReference
func (n_ NEVPNProtocol) SetPasswordReference(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setPasswordReference:"), value)
}/* debug [instance_properties/setter]: passwordReference */


// The proxy settings to use for HTTP and HTTPS connections that route through the VPN.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocol/proxySettings
func (n_ NEVPNProtocol) ProxySettings() INEProxySettings {
	rv := objc.Send[NEProxySettings](n_.ID, objc.Sel("proxySettings"))
	return rv
}/* debug [instance_properties/getter]: proxySettings */


// The proxy settings to use for HTTP and HTTPS connections that route through the VPN.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocol/proxySettings
func (n_ NEVPNProtocol) SetProxySettings(value INEProxySettings) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setProxySettings:"), value)
}/* debug [instance_properties/setter]: proxySettings */


// The address of the VPN server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocol/serverAddress
func (n_ NEVPNProtocol) ServerAddress() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("serverAddress"))
	return rv
}/* debug [instance_properties/getter]: serverAddress */


// The address of the VPN server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocol/serverAddress
func (n_ NEVPNProtocol) SetServerAddress(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setServerAddress:"), value)
}/* debug [instance_properties/setter]: serverAddress */


// The user name component of the tunneling protocol authentication credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocol/username
func (n_ NEVPNProtocol) Username() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("username"))
	return rv
}/* debug [instance_properties/getter]: username */


// The user name component of the tunneling protocol authentication credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocol/username
func (n_ NEVPNProtocol) SetUsername(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setUsername:"), value)
}/* debug [instance_properties/setter]: username */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NEVPNProtocol */


