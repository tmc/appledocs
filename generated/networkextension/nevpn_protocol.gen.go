// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [NEVPNProtocol] class.
type INEVPNProtocol interface {
	objectivec.IObject
	

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
	IdentityData() foundation.foundation.INSData
	SetIdentityData(value foundation.foundation.INSData)
	IdentityDataPassword() foundation.foundation.INSString
	SetIdentityDataPassword(value foundation.foundation.INSString)
	IdentityReference() foundation.foundation.INSData
	SetIdentityReference(value foundation.foundation.INSData)
	IncludeAllNetworks() bool
	SetIncludeAllNetworks(value bool)
	PasswordReference() foundation.foundation.INSData
	SetPasswordReference(value foundation.foundation.INSData)
	ProxySettings() INEProxySettings
	SetProxySettings(value INEProxySettings)
	ServerAddress() foundation.foundation.INSString
	SetServerAddress(value foundation.foundation.INSString)
	Username() foundation.foundation.INSString
	SetUsername(value foundation.foundation.INSString)


	

	// methods:


}





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

























// A Boolean value that indicates whether the VPN disconnects when the device sleeps.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocol/disconnectOnSleep
func (n_ NEVPNProtocol) DisconnectOnSleep() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("disconnectOnSleep"))
	return rv
}


// A Boolean value that indicates whether the VPN disconnects when the device sleeps.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocol/disconnectOnSleep
func (n_ NEVPNProtocol) SetDisconnectOnSleep(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDisconnectOnSleep:"), value)
}


// A Boolean value that indicates whether route rules for the tunnel take precedence over any locally defined routes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocol/enforceRoutes
func (n_ NEVPNProtocol) EnforceRoutes() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("enforceRoutes"))
	return rv
}


// A Boolean value that indicates whether route rules for the tunnel take precedence over any locally defined routes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocol/enforceRoutes
func (n_ NEVPNProtocol) SetEnforceRoutes(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setEnforceRoutes:"), value)
}


// A Boolean value that indicates whether the system excludes all APNs network traffic from the tunnel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocol/excludeAPNs
func (n_ NEVPNProtocol) ExcludeAPNs() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("excludeAPNs"))
	return rv
}


// A Boolean value that indicates whether the system excludes all APNs network traffic from the tunnel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocol/excludeAPNs
func (n_ NEVPNProtocol) SetExcludeAPNs(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setExcludeAPNs:"), value)
}


// A Boolean value that indicates whether the system excludes all cellular services network traffic from the tunnel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocol/excludeCellularServices
func (n_ NEVPNProtocol) ExcludeCellularServices() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("excludeCellularServices"))
	return rv
}


// A Boolean value that indicates whether the system excludes all cellular services network traffic from the tunnel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocol/excludeCellularServices
func (n_ NEVPNProtocol) SetExcludeCellularServices(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setExcludeCellularServices:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocol/excludeDeviceCommunication
func (n_ NEVPNProtocol) ExcludeDeviceCommunication() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("excludeDeviceCommunication"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocol/excludeDeviceCommunication
func (n_ NEVPNProtocol) SetExcludeDeviceCommunication(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setExcludeDeviceCommunication:"), value)
}


// A Boolean value that indicates whether the system excludes all traffic destined for local networks from the tunnel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocol/excludeLocalNetworks
func (n_ NEVPNProtocol) ExcludeLocalNetworks() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("excludeLocalNetworks"))
	return rv
}


// A Boolean value that indicates whether the system excludes all traffic destined for local networks from the tunnel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocol/excludeLocalNetworks
func (n_ NEVPNProtocol) SetExcludeLocalNetworks(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setExcludeLocalNetworks:"), value)
}


// The certificate and private key components of the tunneling protocol authentication credential, in PKCS12 format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocol/identityData
func (n_ NEVPNProtocol) IdentityData() foundation.foundation.INSData {
	rv := objc.Send[foundation.NSData](n_.ID, objc.Sel("identityData"))
	return rv
}


// The certificate and private key components of the tunneling protocol authentication credential, in PKCS12 format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocol/identityData
func (n_ NEVPNProtocol) SetIdentityData(value foundation.foundation.INSData) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIdentityData:"), value)
}


// The password for the PKCS12 tunneling protocol authentication credentials.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocol/identityDataPassword
func (n_ NEVPNProtocol) IdentityDataPassword() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("identityDataPassword"))
	return rv
}


// The password for the PKCS12 tunneling protocol authentication credentials.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocol/identityDataPassword
func (n_ NEVPNProtocol) SetIdentityDataPassword(value foundation.foundation.INSString) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIdentityDataPassword:"), value)
}


// A persistent keychain reference to a keychain item containing the certificate and private key components of the tunneling protocol authentication credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocol/identityReference
func (n_ NEVPNProtocol) IdentityReference() foundation.foundation.INSData {
	rv := objc.Send[foundation.NSData](n_.ID, objc.Sel("identityReference"))
	return rv
}


// A persistent keychain reference to a keychain item containing the certificate and private key components of the tunneling protocol authentication credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocol/identityReference
func (n_ NEVPNProtocol) SetIdentityReference(value foundation.foundation.INSData) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIdentityReference:"), value)
}


// A Boolean value that indicates whether the system sends most network traffic over the tunnel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocol/includeAllNetworks
func (n_ NEVPNProtocol) IncludeAllNetworks() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("includeAllNetworks"))
	return rv
}


// A Boolean value that indicates whether the system sends most network traffic over the tunnel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocol/includeAllNetworks
func (n_ NEVPNProtocol) SetIncludeAllNetworks(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIncludeAllNetworks:"), value)
}


// A persistent keychain reference to a keychain item containing the password component of the tunneling protocol authentication credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocol/passwordReference
func (n_ NEVPNProtocol) PasswordReference() foundation.foundation.INSData {
	rv := objc.Send[foundation.NSData](n_.ID, objc.Sel("passwordReference"))
	return rv
}


// A persistent keychain reference to a keychain item containing the password component of the tunneling protocol authentication credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocol/passwordReference
func (n_ NEVPNProtocol) SetPasswordReference(value foundation.foundation.INSData) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setPasswordReference:"), value)
}


// The proxy settings to use for HTTP and HTTPS connections that route through the VPN.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocol/proxySettings
func (n_ NEVPNProtocol) ProxySettings() INEProxySettings {
	rv := objc.Send[NEProxySettings](n_.ID, objc.Sel("proxySettings"))
	return rv
}


// The proxy settings to use for HTTP and HTTPS connections that route through the VPN.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocol/proxySettings
func (n_ NEVPNProtocol) SetProxySettings(value INEProxySettings) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setProxySettings:"), value)
}


// The address of the VPN server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocol/serverAddress
func (n_ NEVPNProtocol) ServerAddress() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("serverAddress"))
	return rv
}


// The address of the VPN server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocol/serverAddress
func (n_ NEVPNProtocol) SetServerAddress(value foundation.foundation.INSString) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setServerAddress:"), value)
}


// The user name component of the tunneling protocol authentication credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocol/username
func (n_ NEVPNProtocol) Username() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("username"))
	return rv
}


// The user name component of the tunneling protocol authentication credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocol/username
func (n_ NEVPNProtocol) SetUsername(value foundation.foundation.INSString) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setUsername:"), value)
}







