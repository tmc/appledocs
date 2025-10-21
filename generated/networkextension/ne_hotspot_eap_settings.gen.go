// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NEHotspotEAPSettings] class.
var (
	NEHotspotEAPSettingsClass     _NEHotspotEAPSettingsClass
	NEHotspotEAPSettingsClassOnce sync.Once
)

func getNEHotspotEAPSettingsClass() _NEHotspotEAPSettingsClass {
	NEHotspotEAPSettingsClassOnce.Do(func() {
		NEHotspotEAPSettingsClass = _NEHotspotEAPSettingsClass{objc.GetClass("NEHotspotEAPSettings")}
	})
	return NEHotspotEAPSettingsClass
}

type _NEHotspotEAPSettingsClass struct {
	class objc.Class
}

// An interface definition for the [NEHotspotEAPSettings] class.
type INEHotspotEAPSettings interface {
	objectivec.IObject
}

// Extensible Authentication Protocol settings for configuring WPA and WPA2 enterprise Wi-Fi networks.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotEAPSettings
type NEHotspotEAPSettings struct {
	objectivec.Object
}

// NEHotspotEAPSettingsFrom constructs a [NEHotspotEAPSettings] from an unsafe.Pointer.
//
// Extensible Authentication Protocol settings for configuring WPA and WPA2 enterprise Wi-Fi networks.
func NEHotspotEAPSettingsFrom(ptr unsafe.Pointer) NEHotspotEAPSettings {
	return NEHotspotEAPSettings{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (nc _NEHotspotEAPSettingsClass) Alloc() NEHotspotEAPSettings {
	rv := objc.Send[NEHotspotEAPSettings](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NEHotspotEAPSettingsClass) New() NEHotspotEAPSettings {
	rv := objc.Send[NEHotspotEAPSettings](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEHotspotEAPSettings) Init() NEHotspotEAPSettings {
	rv := objc.Send[NEHotspotEAPSettings](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEHotspotEAPSettings) Autorelease() NEHotspotEAPSettings {
	rv := objc.Send[NEHotspotEAPSettings](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEHotspotEAPSettings creates a new NEHotspotEAPSettings instance.
func NewNEHotspotEAPSettings() NEHotspotEAPSettings {
	return getNEHotspotEAPSettingsClass().New()
}


// The identity string to be used in the EAP-Identity/Response packet during outer EAP authentication.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspoteapsettings/outeridentity
func (n_ NEHotspotEAPSettings) OuterIdentity() string {
	rv := objc.Send[string](n_.ID, objc.Sel("outerIdentity"))
	return rv
}


// SetOuterIdentity sets the value of the outerIdentity property.
// The identity string to be used in the EAP-Identity/Response packet during outer EAP authentication.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspoteapsettings/outeridentity
func (n_ NEHotspotEAPSettings) SetOuterIdentity(value string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setOuterIdentity:"), objc.String(value))
}

// The password component of the IEEE 802.1X authentication credential.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspoteapsettings/password
func (n_ NEHotspotEAPSettings) Password() string {
	rv := objc.Send[string](n_.ID, objc.Sel("password"))
	return rv
}


// SetPassword sets the value of the password property.
// The password component of the IEEE 802.1X authentication credential.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspoteapsettings/password
func (n_ NEHotspotEAPSettings) SetPassword(value string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setPassword:"), objc.String(value))
}

// The user name string for EAP authentication, encoded as UTF-8.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspoteapsettings/username
func (n_ NEHotspotEAPSettings) Username() string {
	rv := objc.Send[string](n_.ID, objc.Sel("username"))
	return rv
}


// SetUsername sets the value of the username property.
// The user name string for EAP authentication, encoded as UTF-8.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspoteapsettings/username
func (n_ NEHotspotEAPSettings) SetUsername(value string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setUsername:"), objc.String(value))
}

// The Transport Layer Security (TLS) version to use during a TLS authentication handshake.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspoteapsettings/preferredtlsversion
func (n_ NEHotspotEAPSettings) PreferredTLSVersion() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("preferredTLSVersion"))
	return rv
}


// SetPreferredTLSVersion sets the value of the preferredTLSVersion property.
// The Transport Layer Security (TLS) version to use during a TLS authentication handshake.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspoteapsettings/preferredtlsversion
func (n_ NEHotspotEAPSettings) SetPreferredTLSVersion(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setPreferredTLSVersion:"), value)
}

// A Boolean value indicating whether a network requires two-factor authentication or allows zero-factor authentication.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspoteapsettings/istlsclientcertificaterequired
func (n_ NEHotspotEAPSettings) IsTLSClientCertificateRequired() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("isTLSClientCertificateRequired"))
	return rv
}


// SetIsTLSClientCertificateRequired sets the value of the isTLSClientCertificateRequired property.
// A Boolean value indicating whether a network requires two-factor authentication or allows zero-factor authentication.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspoteapsettings/istlsclientcertificaterequired
func (n_ NEHotspotEAPSettings) SetIsTLSClientCertificateRequired(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIsTLSClientCertificateRequired:"), value)
}

// An array of supported EAP types.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspoteapsettings/supportedeaptypes
func (n_ NEHotspotEAPSettings) SupportedEAPTypes() foundation.Number {
	rv := objc.Send[foundation.Number](n_.ID, objc.Sel("supportedEAPTypes"))
	return rv
}


// SetSupportedEAPTypes sets the value of the supportedEAPTypes property.
// An array of supported EAP types.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspoteapsettings/supportedeaptypes
func (n_ NEHotspotEAPSettings) SetSupportedEAPTypes(value foundation.Number) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setSupportedEAPTypes:"), value)
}

// The inner-layer authentication protocol used by a TTLS module.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspoteapsettings/ttlsinnerauthenticationtype-swift.property
func (n_ NEHotspotEAPSettings) TtlsInnerAuthenticationType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("ttlsInnerAuthenticationType"))
	return rv
}


// SetTtlsInnerAuthenticationType sets the value of the ttlsInnerAuthenticationType property.
// The inner-layer authentication protocol used by a TTLS module.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspoteapsettings/ttlsinnerauthenticationtype-swift.property
func (n_ NEHotspotEAPSettings) SetTtlsInnerAuthenticationType(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setTtlsInnerAuthenticationType:"), value)
}

// An array of server certificate common name strings used to verify a server’s certificate.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspoteapsettings/trustedservernames
func (n_ NEHotspotEAPSettings) TrustedServerNames() string {
	rv := objc.Send[string](n_.ID, objc.Sel("trustedServerNames"))
	return rv
}


// SetTrustedServerNames sets the value of the trustedServerNames property.
// An array of server certificate common name strings used to verify a server’s certificate.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspoteapsettings/trustedservernames
func (n_ NEHotspotEAPSettings) SetTrustedServerNames(value string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setTrustedServerNames:"), objc.String(value))
}



