// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [NEVPNProtocolIKEv2] class.
var (
	NEVPNProtocolIKEv2Class     _NEVPNProtocolIKEv2Class
	NEVPNProtocolIKEv2ClassOnce sync.Once
)

func getNEVPNProtocolIKEv2Class() _NEVPNProtocolIKEv2Class {
	NEVPNProtocolIKEv2ClassOnce.Do(func() {
		NEVPNProtocolIKEv2Class = _NEVPNProtocolIKEv2Class{objc.GetClass("NEVPNProtocolIKEv2")}
	})
	return NEVPNProtocolIKEv2Class
}

type _NEVPNProtocolIKEv2Class struct {
	class objc.Class
}

// An interface definition for the [NEVPNProtocolIKEv2] class.
type INEVPNProtocolIKEv2 interface {
	INEVPNProtocolIPSec
}

// Settings for an IKEv2 VPN configuration.
//
// Instances of this class are thread safe.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2
type NEVPNProtocolIKEv2 struct {
	NEVPNProtocolIPSec
}

// NEVPNProtocolIKEv2From constructs a [NEVPNProtocolIKEv2] from an unsafe.Pointer.
//
// Settings for an IKEv2 VPN configuration.
func NEVPNProtocolIKEv2From(ptr unsafe.Pointer) NEVPNProtocolIKEv2 {
	return NEVPNProtocolIKEv2{
		NEVPNProtocolIPSec: NEVPNProtocolIPSecFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (nc _NEVPNProtocolIKEv2Class) Alloc() NEVPNProtocolIKEv2 {
	rv := objc.Send[NEVPNProtocolIKEv2](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NEVPNProtocolIKEv2Class) New() NEVPNProtocolIKEv2 {
	rv := objc.Send[NEVPNProtocolIKEv2](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEVPNProtocolIKEv2) Init() NEVPNProtocolIKEv2 {
	rv := objc.Send[NEVPNProtocolIKEv2](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEVPNProtocolIKEv2) Autorelease() NEVPNProtocolIKEv2 {
	rv := objc.Send[NEVPNProtocolIKEv2](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEVPNProtocolIKEv2 creates a new NEVPNProtocolIKEv2 instance.
func NewNEVPNProtocolIKEv2() NEVPNProtocolIKEv2 {
	return getNEVPNProtocolIKEv2Class().New()
}


// A Boolean value that indicates whether servers that don’t support post-quantum key exchanges can skip them.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/allowPostQuantumKeyExchangeFallback
func (n_ NEVPNProtocolIKEv2) AllowPostQuantumKeyExchangeFallback() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("allowPostQuantumKeyExchangeFallback"))
	return rv
}


// SetAllowPostQuantumKeyExchangeFallback sets the value of the allowPostQuantumKeyExchangeFallback property.
// A Boolean value that indicates whether servers that don’t support post-quantum key exchanges can skip them.

//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/allowPostQuantumKeyExchangeFallback
func (n_ NEVPNProtocolIKEv2) SetAllowPostQuantumKeyExchangeFallback(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setAllowPostQuantumKeyExchangeFallback:"), value)
}

// The type of the certificate in the identity configured in or .
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/certificateType
func (n_ NEVPNProtocolIKEv2) CertificateType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("certificateType"))
	return rv
}


// SetCertificateType sets the value of the certificateType property.
// The type of the certificate in the identity configured in or .

//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/certificateType
func (n_ NEVPNProtocolIKEv2) SetCertificateType(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setCertificateType:"), value)
}

// An object containing the parameters for the child IPSec security associations to be negotiated for each IKEv2 policy.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/childSecurityAssociationParameters
func (n_ NEVPNProtocolIKEv2) ChildSecurityAssociationParameters() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("childSecurityAssociationParameters"))
	return rv
}

// The frequency at which the IKEv2 client will run the dead peer detection algorithm.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/deadPeerDetectionRate
func (n_ NEVPNProtocolIKEv2) DeadPeerDetectionRate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("deadPeerDetectionRate"))
	return rv
}


// SetDeadPeerDetectionRate sets the value of the deadPeerDetectionRate property.
// The frequency at which the IKEv2 client will run the dead peer detection algorithm.

//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/deadPeerDetectionRate
func (n_ NEVPNProtocolIKEv2) SetDeadPeerDetectionRate(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDeadPeerDetectionRate:"), value)
}

// A Boolean indicating whether or not MOBIKE should be disabled for the IKEv2 sessions.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/disableMOBIKE
func (n_ NEVPNProtocolIKEv2) DisableMOBIKE() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("disableMOBIKE"))
	return rv
}


// SetDisableMOBIKE sets the value of the disableMOBIKE property.
// A Boolean indicating whether or not MOBIKE should be disabled for the IKEv2 sessions.

//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/disableMOBIKE
func (n_ NEVPNProtocolIKEv2) SetDisableMOBIKE(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDisableMOBIKE:"), value)
}

// A Boolean indicating whether or not IKEv2 server redirects are disabled.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/disableRedirect
func (n_ NEVPNProtocolIKEv2) DisableRedirect() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("disableRedirect"))
	return rv
}


// SetDisableRedirect sets the value of the disableRedirect property.
// A Boolean indicating whether or not IKEv2 server redirects are disabled.

//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/disableRedirect
func (n_ NEVPNProtocolIKEv2) SetDisableRedirect(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDisableRedirect:"), value)
}

// A property to enable the use of cellular data when Wi-Fi connectivity is poor.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/enableFallback
func (n_ NEVPNProtocolIKEv2) EnableFallback() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("enableFallback"))
	return rv
}


// SetEnableFallback sets the value of the enableFallback property.
// A property to enable the use of cellular data when Wi-Fi connectivity is poor.

//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/enableFallback
func (n_ NEVPNProtocolIKEv2) SetEnableFallback(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setEnableFallback:"), value)
}

// A Boolean indicating whether or not Perfect Forward Secrecy is enabled.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/enablePFS
func (n_ NEVPNProtocolIKEv2) EnablePFS() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("enablePFS"))
	return rv
}


// SetEnablePFS sets the value of the enablePFS property.
// A Boolean indicating whether or not Perfect Forward Secrecy is enabled.

//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/enablePFS
func (n_ NEVPNProtocolIKEv2) SetEnablePFS(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setEnablePFS:"), value)
}

// Enable revocation checking of the IKEv2 server certificate.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/enableRevocationCheck
func (n_ NEVPNProtocolIKEv2) EnableRevocationCheck() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("enableRevocationCheck"))
	return rv
}


// SetEnableRevocationCheck sets the value of the enableRevocationCheck property.
// Enable revocation checking of the IKEv2 server certificate.

//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/enableRevocationCheck
func (n_ NEVPNProtocolIKEv2) SetEnableRevocationCheck(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setEnableRevocationCheck:"), value)
}

// An object containing the parameters for the initial IKE security association to be negotiated with the IKEv2 server.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/ikeSecurityAssociationParameters
func (n_ NEVPNProtocolIKEv2) IKESecurityAssociationParameters() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("IKESecurityAssociationParameters"))
	return rv
}

// The minimum TLS version to allow for EAP-TLS authentication.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/maximumTLSVersion
func (n_ NEVPNProtocolIKEv2) MaximumTLSVersion() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("maximumTLSVersion"))
	return rv
}


// SetMaximumTLSVersion sets the value of the maximumTLSVersion property.
// The minimum TLS version to allow for EAP-TLS authentication.

//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/maximumTLSVersion
func (n_ NEVPNProtocolIKEv2) SetMaximumTLSVersion(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMaximumTLSVersion:"), value)
}

// The minimum TLS version to allow for EAP-TLS authentication.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/minimumTLSVersion
func (n_ NEVPNProtocolIKEv2) MinimumTLSVersion() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("minimumTLSVersion"))
	return rv
}


// SetMinimumTLSVersion sets the value of the minimumTLSVersion property.
// The minimum TLS version to allow for EAP-TLS authentication.

//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/minimumTLSVersion
func (n_ NEVPNProtocolIKEv2) SetMinimumTLSVersion(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMinimumTLSVersion:"), value)
}

// The Maximum Transmission Unit (MTU) size in bytes to assign to the tunnel interface.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/mtu
func (n_ NEVPNProtocolIKEv2) Mtu() uint {
	rv := objc.Send[uint](n_.ID, objc.Sel("mtu"))
	return rv
}


// SetMtu sets the value of the mtu property.
// The Maximum Transmission Unit (MTU) size in bytes to assign to the tunnel interface.

//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/mtu
func (n_ NEVPNProtocolIKEv2) SetMtu(value uint) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMtu:"), value)
}

// The configuration for a post-quantum pre-shared key (PPK).
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/ppkConfiguration
func (n_ NEVPNProtocolIKEv2) PpkConfiguration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("ppkConfiguration"))
	return rv
}


// SetPpkConfiguration sets the value of the ppkConfiguration property.
// The configuration for a post-quantum pre-shared key (PPK).

//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/ppkConfiguration
func (n_ NEVPNProtocolIKEv2) SetPpkConfiguration(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setPpkConfiguration:"), value)
}

// A string containing the value of the Subject Common Name field of the IKEv2 server’s certificate.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/serverCertificateCommonName
func (n_ NEVPNProtocolIKEv2) ServerCertificateCommonName() string {
	rv := objc.Send[string](n_.ID, objc.Sel("serverCertificateCommonName"))
	return rv
}


// SetServerCertificateCommonName sets the value of the serverCertificateCommonName property.
// A string containing the value of the Subject Common Name field of the IKEv2 server’s certificate.

//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/serverCertificateCommonName
func (n_ NEVPNProtocolIKEv2) SetServerCertificateCommonName(value string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setServerCertificateCommonName:"), objc.String(value))
}

// A string containing the value of the Subject Common Name field of the Certificate Authority certificate that issued the IKEv2 server’s certificate.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/serverCertificateIssuerCommonName
func (n_ NEVPNProtocolIKEv2) ServerCertificateIssuerCommonName() string {
	rv := objc.Send[string](n_.ID, objc.Sel("serverCertificateIssuerCommonName"))
	return rv
}


// SetServerCertificateIssuerCommonName sets the value of the serverCertificateIssuerCommonName property.
// A string containing the value of the Subject Common Name field of the Certificate Authority certificate that issued the IKEv2 server’s certificate.

//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/serverCertificateIssuerCommonName
func (n_ NEVPNProtocolIKEv2) SetServerCertificateIssuerCommonName(value string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setServerCertificateIssuerCommonName:"), objc.String(value))
}

// Require a “not revoked” result when checking if the certificate identifying the server is revoked.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/strictRevocationCheck
func (n_ NEVPNProtocolIKEv2) StrictRevocationCheck() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("strictRevocationCheck"))
	return rv
}


// SetStrictRevocationCheck sets the value of the strictRevocationCheck property.
// Require a “not revoked” result when checking if the certificate identifying the server is revoked.

//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/strictRevocationCheck
func (n_ NEVPNProtocolIKEv2) SetStrictRevocationCheck(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setStrictRevocationCheck:"), value)
}

// A Boolean indicating whether or not the IKEv2 client should use the INTERNAL_IP4_SUBNET and/or INTERNAL_IP6_SUBNET attributes sent by the IKEv2 server.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/useConfigurationAttributeInternalIPSubnet
func (n_ NEVPNProtocolIKEv2) UseConfigurationAttributeInternalIPSubnet() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("useConfigurationAttributeInternalIPSubnet"))
	return rv
}


// SetUseConfigurationAttributeInternalIPSubnet sets the value of the useConfigurationAttributeInternalIPSubnet property.
// A Boolean indicating whether or not the IKEv2 client should use the INTERNAL_IP4_SUBNET and/or INTERNAL_IP6_SUBNET attributes sent by the IKEv2 server.

//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/useConfigurationAttributeInternalIPSubnet
func (n_ NEVPNProtocolIKEv2) SetUseConfigurationAttributeInternalIPSubnet(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setUseConfigurationAttributeInternalIPSubnet:"), value)
}



