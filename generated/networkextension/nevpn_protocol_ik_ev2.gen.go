// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class NEVPNProtocolIKEv2 */


/* debug [class_header]: Header for NEVPNProtocolIKEv2 */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NEVPNProtocolIKEv2 */
// An interface definition for the [NEVPNProtocolIKEv2] class.
type INEVPNProtocolIKEv2 interface {
	INEVPNProtocolIPSec
	
/* debug [class_interface_properties]: Properties for NEVPNProtocolIKEv2 */
	// properties:
	AllowPostQuantumKeyExchangeFallback() bool
	SetAllowPostQuantumKeyExchangeFallback(value bool)
	CertificateType() NEVPNIKEv2CertificateType
	SetCertificateType(value NEVPNIKEv2CertificateType)
	ChildSecurityAssociationParameters() INEVPNIKEv2SecurityAssociationParameters
	DeadPeerDetectionRate() NEVPNIKEv2DeadPeerDetectionRate
	SetDeadPeerDetectionRate(value NEVPNIKEv2DeadPeerDetectionRate)
	DisableMOBIKE() bool
	SetDisableMOBIKE(value bool)
	DisableRedirect() bool
	SetDisableRedirect(value bool)
	EnablePFS() bool
	SetEnablePFS(value bool)
	EnableRevocationCheck() bool
	SetEnableRevocationCheck(value bool)
	IKESecurityAssociationParameters() INEVPNIKEv2SecurityAssociationParameters
	MaximumTLSVersion() NEVPNIKEv2TLSVersion
	SetMaximumTLSVersion(value NEVPNIKEv2TLSVersion)
	MinimumTLSVersion() NEVPNIKEv2TLSVersion
	SetMinimumTLSVersion(value NEVPNIKEv2TLSVersion)
	Mtu() uint
	SetMtu(value uint)
	PpkConfiguration() INEVPNIKEv2PPKConfiguration
	SetPpkConfiguration(value INEVPNIKEv2PPKConfiguration)
	ServerCertificateCommonName() objc.IObject /* cross-framework: NSString */
	SetServerCertificateCommonName(value objc.IObject /* cross-framework: NSString */)
	ServerCertificateIssuerCommonName() objc.IObject /* cross-framework: NSString */
	SetServerCertificateIssuerCommonName(value objc.IObject /* cross-framework: NSString */)
	StrictRevocationCheck() bool
	SetStrictRevocationCheck(value bool)
	UseConfigurationAttributeInternalIPSubnet() bool
	SetUseConfigurationAttributeInternalIPSubnet(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NEVPNProtocolIKEv2 */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NEVPNProtocolIKEv2 */
// Alloc allocates a new instance without initialization.
func (nc _NEVPNProtocolIKEv2Class) Alloc() NEVPNProtocolIKEv2 {
	rv := objc.Send[NEVPNProtocolIKEv2](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NEVPNProtocolIKEv2 */
// Settings for an IKEv2 VPN configuration.
//
// Instances of this class are thread safe.


// Settings for an IKEv2 VPN configuration.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NEVPNProtocolIKEv2 *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NEVPNProtocolIKEv2 */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NEVPNProtocolIKEv2 */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NEVPNProtocolIKEv2 */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NEVPNProtocolIKEv2 */

// A Boolean value that indicates whether servers that don’t support post-quantum key exchanges can skip them.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/allowPostQuantumKeyExchangeFallback
func (n_ NEVPNProtocolIKEv2) AllowPostQuantumKeyExchangeFallback() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("allowPostQuantumKeyExchangeFallback"))
	return rv
}/* debug [instance_properties/getter]: allowPostQuantumKeyExchangeFallback */


// A Boolean value that indicates whether servers that don’t support post-quantum key exchanges can skip them.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/allowPostQuantumKeyExchangeFallback
func (n_ NEVPNProtocolIKEv2) SetAllowPostQuantumKeyExchangeFallback(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setAllowPostQuantumKeyExchangeFallback:"), value)
}/* debug [instance_properties/setter]: allowPostQuantumKeyExchangeFallback */


// The type of the certificate in the identity configured in or .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/certificateType
func (n_ NEVPNProtocolIKEv2) CertificateType() NEVPNIKEv2CertificateType {
	rv := objc.Send[NEVPNIKEv2CertificateType](n_.ID, objc.Sel("certificateType"))
	return rv
}/* debug [instance_properties/getter]: certificateType */


// The type of the certificate in the identity configured in or .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/certificateType
func (n_ NEVPNProtocolIKEv2) SetCertificateType(value NEVPNIKEv2CertificateType) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setCertificateType:"), value)
}/* debug [instance_properties/setter]: certificateType */


// An object containing the parameters for the child IPSec security associations to be negotiated for each IKEv2 policy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/childSecurityAssociationParameters
func (n_ NEVPNProtocolIKEv2) ChildSecurityAssociationParameters() INEVPNIKEv2SecurityAssociationParameters {
	rv := objc.Send[NEVPNIKEv2SecurityAssociationParameters](n_.ID, objc.Sel("childSecurityAssociationParameters"))
	return rv
}/* debug [instance_properties/getter]: childSecurityAssociationParameters */


// The frequency at which the IKEv2 client will run the dead peer detection algorithm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/deadPeerDetectionRate
func (n_ NEVPNProtocolIKEv2) DeadPeerDetectionRate() NEVPNIKEv2DeadPeerDetectionRate {
	rv := objc.Send[NEVPNIKEv2DeadPeerDetectionRate](n_.ID, objc.Sel("deadPeerDetectionRate"))
	return rv
}/* debug [instance_properties/getter]: deadPeerDetectionRate */


// The frequency at which the IKEv2 client will run the dead peer detection algorithm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/deadPeerDetectionRate
func (n_ NEVPNProtocolIKEv2) SetDeadPeerDetectionRate(value NEVPNIKEv2DeadPeerDetectionRate) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDeadPeerDetectionRate:"), value)
}/* debug [instance_properties/setter]: deadPeerDetectionRate */


// A Boolean indicating whether or not MOBIKE should be disabled for the IKEv2 sessions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/disableMOBIKE
func (n_ NEVPNProtocolIKEv2) DisableMOBIKE() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("disableMOBIKE"))
	return rv
}/* debug [instance_properties/getter]: disableMOBIKE */


// A Boolean indicating whether or not MOBIKE should be disabled for the IKEv2 sessions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/disableMOBIKE
func (n_ NEVPNProtocolIKEv2) SetDisableMOBIKE(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDisableMOBIKE:"), value)
}/* debug [instance_properties/setter]: disableMOBIKE */


// A Boolean indicating whether or not IKEv2 server redirects are disabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/disableRedirect
func (n_ NEVPNProtocolIKEv2) DisableRedirect() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("disableRedirect"))
	return rv
}/* debug [instance_properties/getter]: disableRedirect */


// A Boolean indicating whether or not IKEv2 server redirects are disabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/disableRedirect
func (n_ NEVPNProtocolIKEv2) SetDisableRedirect(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDisableRedirect:"), value)
}/* debug [instance_properties/setter]: disableRedirect */


// A Boolean indicating whether or not Perfect Forward Secrecy is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/enablePFS
func (n_ NEVPNProtocolIKEv2) EnablePFS() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("enablePFS"))
	return rv
}/* debug [instance_properties/getter]: enablePFS */


// A Boolean indicating whether or not Perfect Forward Secrecy is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/enablePFS
func (n_ NEVPNProtocolIKEv2) SetEnablePFS(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setEnablePFS:"), value)
}/* debug [instance_properties/setter]: enablePFS */


// Enable revocation checking of the IKEv2 server certificate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/enableRevocationCheck
func (n_ NEVPNProtocolIKEv2) EnableRevocationCheck() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("enableRevocationCheck"))
	return rv
}/* debug [instance_properties/getter]: enableRevocationCheck */


// Enable revocation checking of the IKEv2 server certificate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/enableRevocationCheck
func (n_ NEVPNProtocolIKEv2) SetEnableRevocationCheck(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setEnableRevocationCheck:"), value)
}/* debug [instance_properties/setter]: enableRevocationCheck */


// An object containing the parameters for the initial IKE security association to be negotiated with the IKEv2 server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/ikeSecurityAssociationParameters
func (n_ NEVPNProtocolIKEv2) IKESecurityAssociationParameters() INEVPNIKEv2SecurityAssociationParameters {
	rv := objc.Send[NEVPNIKEv2SecurityAssociationParameters](n_.ID, objc.Sel("IKESecurityAssociationParameters"))
	return rv
}/* debug [instance_properties/getter]: IKESecurityAssociationParameters */


// The minimum TLS version to allow for EAP-TLS authentication.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/maximumTLSVersion
func (n_ NEVPNProtocolIKEv2) MaximumTLSVersion() NEVPNIKEv2TLSVersion {
	rv := objc.Send[NEVPNIKEv2TLSVersion](n_.ID, objc.Sel("maximumTLSVersion"))
	return rv
}/* debug [instance_properties/getter]: maximumTLSVersion */


// The minimum TLS version to allow for EAP-TLS authentication.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/maximumTLSVersion
func (n_ NEVPNProtocolIKEv2) SetMaximumTLSVersion(value NEVPNIKEv2TLSVersion) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMaximumTLSVersion:"), value)
}/* debug [instance_properties/setter]: maximumTLSVersion */


// The minimum TLS version to allow for EAP-TLS authentication.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/minimumTLSVersion
func (n_ NEVPNProtocolIKEv2) MinimumTLSVersion() NEVPNIKEv2TLSVersion {
	rv := objc.Send[NEVPNIKEv2TLSVersion](n_.ID, objc.Sel("minimumTLSVersion"))
	return rv
}/* debug [instance_properties/getter]: minimumTLSVersion */


// The minimum TLS version to allow for EAP-TLS authentication.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/minimumTLSVersion
func (n_ NEVPNProtocolIKEv2) SetMinimumTLSVersion(value NEVPNIKEv2TLSVersion) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMinimumTLSVersion:"), value)
}/* debug [instance_properties/setter]: minimumTLSVersion */


// The Maximum Transmission Unit (MTU) size in bytes to assign to the tunnel interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/mtu
func (n_ NEVPNProtocolIKEv2) Mtu() uint {
	rv := objc.Send[uint](n_.ID, objc.Sel("mtu"))
	return rv
}/* debug [instance_properties/getter]: mtu */


// The Maximum Transmission Unit (MTU) size in bytes to assign to the tunnel interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/mtu
func (n_ NEVPNProtocolIKEv2) SetMtu(value uint) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMtu:"), value)
}/* debug [instance_properties/setter]: mtu */


// The configuration for a post-quantum pre-shared key (PPK).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/ppkConfiguration
func (n_ NEVPNProtocolIKEv2) PpkConfiguration() INEVPNIKEv2PPKConfiguration {
	rv := objc.Send[NEVPNIKEv2PPKConfiguration](n_.ID, objc.Sel("ppkConfiguration"))
	return rv
}/* debug [instance_properties/getter]: ppkConfiguration */


// The configuration for a post-quantum pre-shared key (PPK).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/ppkConfiguration
func (n_ NEVPNProtocolIKEv2) SetPpkConfiguration(value INEVPNIKEv2PPKConfiguration) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setPpkConfiguration:"), value)
}/* debug [instance_properties/setter]: ppkConfiguration */


// A string containing the value of the Subject Common Name field of the IKEv2 server’s certificate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/serverCertificateCommonName
func (n_ NEVPNProtocolIKEv2) ServerCertificateCommonName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("serverCertificateCommonName"))
	return rv
}/* debug [instance_properties/getter]: serverCertificateCommonName */


// A string containing the value of the Subject Common Name field of the IKEv2 server’s certificate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/serverCertificateCommonName
func (n_ NEVPNProtocolIKEv2) SetServerCertificateCommonName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setServerCertificateCommonName:"), value)
}/* debug [instance_properties/setter]: serverCertificateCommonName */


// A string containing the value of the Subject Common Name field of the Certificate Authority certificate that issued the IKEv2 server’s certificate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/serverCertificateIssuerCommonName
func (n_ NEVPNProtocolIKEv2) ServerCertificateIssuerCommonName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("serverCertificateIssuerCommonName"))
	return rv
}/* debug [instance_properties/getter]: serverCertificateIssuerCommonName */


// A string containing the value of the Subject Common Name field of the Certificate Authority certificate that issued the IKEv2 server’s certificate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/serverCertificateIssuerCommonName
func (n_ NEVPNProtocolIKEv2) SetServerCertificateIssuerCommonName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setServerCertificateIssuerCommonName:"), value)
}/* debug [instance_properties/setter]: serverCertificateIssuerCommonName */


// Require a “not revoked” result when checking if the certificate identifying the server is revoked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/strictRevocationCheck
func (n_ NEVPNProtocolIKEv2) StrictRevocationCheck() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("strictRevocationCheck"))
	return rv
}/* debug [instance_properties/getter]: strictRevocationCheck */


// Require a “not revoked” result when checking if the certificate identifying the server is revoked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/strictRevocationCheck
func (n_ NEVPNProtocolIKEv2) SetStrictRevocationCheck(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setStrictRevocationCheck:"), value)
}/* debug [instance_properties/setter]: strictRevocationCheck */


// A Boolean indicating whether or not the IKEv2 client should use the INTERNAL_IP4_SUBNET and/or INTERNAL_IP6_SUBNET attributes sent by the IKEv2 server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/useConfigurationAttributeInternalIPSubnet
func (n_ NEVPNProtocolIKEv2) UseConfigurationAttributeInternalIPSubnet() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("useConfigurationAttributeInternalIPSubnet"))
	return rv
}/* debug [instance_properties/getter]: useConfigurationAttributeInternalIPSubnet */


// A Boolean indicating whether or not the IKEv2 client should use the INTERNAL_IP4_SUBNET and/or INTERNAL_IP6_SUBNET attributes sent by the IKEv2 server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/useConfigurationAttributeInternalIPSubnet
func (n_ NEVPNProtocolIKEv2) SetUseConfigurationAttributeInternalIPSubnet(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setUseConfigurationAttributeInternalIPSubnet:"), value)
}/* debug [instance_properties/setter]: useConfigurationAttributeInternalIPSubnet */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NEVPNProtocolIKEv2 */


