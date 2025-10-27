//go:build darwin && ios

// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for NEHotspotEAPSettings


// Sets the client identity for EAP authentication.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotEAPSettings/setIdentity(_:)
func (n_ NEHotspotEAPSettings) SetIdentity(identity objectivec.IObject) bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("setIdentity:"), identity)
	return rv
}

// Sets trusted EAP server certificates for an enterprise Wi-Fi or Hotspot 2.0 network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotEAPSettings/setTrustedServerCertificates(_:)
func (n_ NEHotspotEAPSettings) SetTrustedServerCertificates(certificates foundation.foundation.INSArray) bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("setTrustedServerCertificates:"), certificates)
	return rv
}

// iOS-only properties

// A Boolean value indicating whether a network requires two-factor authentication or allows zero-factor authentication.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotEAPSettings/isTLSClientCertificateRequired
func (n_ NEHotspotEAPSettings) TlsClientCertificateRequired() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("tlsClientCertificateRequired"))
	return rv
}
func (n_ NEHotspotEAPSettings) SetTlsClientCertificateRequired(value bool) {
	n_.ID.Send(objc.RegisterName("setTlsClientCertificateRequired:"), value)
}

// The identity string to be used in the EAP-Identity/Response packet during outer EAP authentication.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotEAPSettings/outerIdentity
func (n_ NEHotspotEAPSettings) OuterIdentity() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("outerIdentity"))
	return rv
}
func (n_ NEHotspotEAPSettings) SetOuterIdentity(value foundation.foundation.INSString) {
	n_.ID.Send(objc.RegisterName("setOuterIdentity:"), value)
}

// The password component of the IEEE 802.1X authentication credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotEAPSettings/password
func (n_ NEHotspotEAPSettings) Password() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("password"))
	return rv
}
func (n_ NEHotspotEAPSettings) SetPassword(value foundation.foundation.INSString) {
	n_.ID.Send(objc.RegisterName("setPassword:"), value)
}

// The Transport Layer Security (TLS) version to use during a TLS authentication handshake.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotEAPSettings/preferredTLSVersion
func (n_ NEHotspotEAPSettings) PreferredTLSVersion() NEHotspotConfigurationEAPTLSVersion {
	rv := objc.Send[NEHotspotConfigurationEAPTLSVersion](n_.ID, objc.Sel("preferredTLSVersion"))
	return rv
}
func (n_ NEHotspotEAPSettings) SetPreferredTLSVersion(value NEHotspotConfigurationEAPTLSVersion) {
	n_.ID.Send(objc.RegisterName("setPreferredTLSVersion:"), value)
}

// An array of supported EAP types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotEAPSettings/supportedEAPTypes
func (n_ NEHotspotEAPSettings) SupportedEAPTypes() []foundation.Number {
	rv := objc.Send[[]foundation.Number](n_.ID, objc.Sel("supportedEAPTypes"))
	return rv
}
func (n_ NEHotspotEAPSettings) SetSupportedEAPTypes(value []foundation.Number) {
	n_.ID.Send(objc.RegisterName("setSupportedEAPTypes:"), value)
}

// An array of server certificate common name strings used to verify a server’s certificate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotEAPSettings/trustedServerNames
func (n_ NEHotspotEAPSettings) TrustedServerNames() []string {
	rv := objc.Send[[]string](n_.ID, objc.Sel("trustedServerNames"))
	return rv
}
func (n_ NEHotspotEAPSettings) SetTrustedServerNames(value []string) {
	n_.ID.Send(objc.RegisterName("setTrustedServerNames:"), value)
}

// The inner-layer authentication protocol used by a TTLS module.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotEAPSettings/ttlsInnerAuthenticationType-swift.property
func (n_ NEHotspotEAPSettings) TtlsInnerAuthenticationType() NEHotspotConfigurationTTLSInnerAuthenticationType {
	rv := objc.Send[NEHotspotConfigurationTTLSInnerAuthenticationType](n_.ID, objc.Sel("ttlsInnerAuthenticationType"))
	return rv
}
func (n_ NEHotspotEAPSettings) SetTtlsInnerAuthenticationType(value NEHotspotConfigurationTTLSInnerAuthenticationType) {
	n_.ID.Send(objc.RegisterName("setTtlsInnerAuthenticationType:"), value)
}

// The user name string for EAP authentication, encoded as UTF-8.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotEAPSettings/username
func (n_ NEHotspotEAPSettings) Username() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("username"))
	return rv
}
func (n_ NEHotspotEAPSettings) SetUsername(value foundation.foundation.INSString) {
	n_.ID.Send(objc.RegisterName("setUsername:"), value)
}





