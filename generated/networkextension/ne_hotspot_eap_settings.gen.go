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
	// properties:
	IsTLSClientCertificateRequired() bool
	SetIsTLSClientCertificateRequired(value bool)
	OuterIdentity() objc.IObject /* cross-framework: NSString */
	SetOuterIdentity(value objc.IObject /* cross-framework: NSString */)
	Password() objc.IObject /* cross-framework: NSString */
	SetPassword(value objc.IObject /* cross-framework: NSString */)
	PreferredTLSVersion() unsafe.Pointer
	SetPreferredTLSVersion(value unsafe.Pointer)
	SupportedEAPTypes() objc.IObject /* cross-framework: NSNumber */
	SetSupportedEAPTypes(value objc.IObject /* cross-framework: NSNumber */)
	TrustedServerNames() objc.IObject /* cross-framework: NSString */
	SetTrustedServerNames(value objc.IObject /* cross-framework: NSString */)
	TtlsInnerAuthenticationType() unsafe.Pointer
	SetTtlsInnerAuthenticationType(value unsafe.Pointer)
	Username() objc.IObject /* cross-framework: NSString */
	SetUsername(value objc.IObject /* cross-framework: NSString */)
	// methods:
}

// Extensible Authentication Protocol settings for configuring WPA and WPA2 enterprise Wi-Fi networks.


// Extensible Authentication Protocol settings for configuring WPA and WPA2 enterprise Wi-Fi networks.
//
// [Full Topic]
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



// A Boolean value indicating whether a network requires two-factor authentication or allows zero-factor authentication.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspoteapsettings/istlsclientcertificaterequired
func (n_ NEHotspotEAPSettings) IsTLSClientCertificateRequired() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("isTLSClientCertificateRequired"))
	return rv
}


// A Boolean value indicating whether a network requires two-factor authentication or allows zero-factor authentication.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspoteapsettings/istlsclientcertificaterequired
func (n_ NEHotspotEAPSettings) SetIsTLSClientCertificateRequired(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIsTLSClientCertificateRequired:"), value)
}


// The identity string to be used in the EAP-Identity/Response packet during outer EAP authentication.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspoteapsettings/outeridentity
func (n_ NEHotspotEAPSettings) OuterIdentity() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("outerIdentity"))
	return rv
}


// The identity string to be used in the EAP-Identity/Response packet during outer EAP authentication.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspoteapsettings/outeridentity
func (n_ NEHotspotEAPSettings) SetOuterIdentity(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setOuterIdentity:"), value)
}


// The password component of the IEEE 802.1X authentication credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspoteapsettings/password
func (n_ NEHotspotEAPSettings) Password() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("password"))
	return rv
}


// The password component of the IEEE 802.1X authentication credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspoteapsettings/password
func (n_ NEHotspotEAPSettings) SetPassword(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setPassword:"), value)
}


// The Transport Layer Security (TLS) version to use during a TLS authentication handshake.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspoteapsettings/preferredtlsversion
func (n_ NEHotspotEAPSettings) PreferredTLSVersion() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("preferredTLSVersion"))
	return rv
}


// The Transport Layer Security (TLS) version to use during a TLS authentication handshake.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspoteapsettings/preferredtlsversion
func (n_ NEHotspotEAPSettings) SetPreferredTLSVersion(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setPreferredTLSVersion:"), value)
}


// An array of supported EAP types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspoteapsettings/supportedeaptypes
func (n_ NEHotspotEAPSettings) SupportedEAPTypes() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](n_.ID, objc.Sel("supportedEAPTypes"))
	return rv
}


// An array of supported EAP types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspoteapsettings/supportedeaptypes
func (n_ NEHotspotEAPSettings) SetSupportedEAPTypes(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setSupportedEAPTypes:"), value)
}


// An array of server certificate common name strings used to verify a server’s certificate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspoteapsettings/trustedservernames
func (n_ NEHotspotEAPSettings) TrustedServerNames() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("trustedServerNames"))
	return rv
}


// An array of server certificate common name strings used to verify a server’s certificate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspoteapsettings/trustedservernames
func (n_ NEHotspotEAPSettings) SetTrustedServerNames(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setTrustedServerNames:"), value)
}


// The inner-layer authentication protocol used by a TTLS module.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspoteapsettings/ttlsinnerauthenticationtype-swift.property
func (n_ NEHotspotEAPSettings) TtlsInnerAuthenticationType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("ttlsInnerAuthenticationType"))
	return rv
}


// The inner-layer authentication protocol used by a TTLS module.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspoteapsettings/ttlsinnerauthenticationtype-swift.property
func (n_ NEHotspotEAPSettings) SetTtlsInnerAuthenticationType(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setTtlsInnerAuthenticationType:"), value)
}


// The user name string for EAP authentication, encoded as UTF-8.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspoteapsettings/username
func (n_ NEHotspotEAPSettings) Username() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("username"))
	return rv
}


// The user name string for EAP authentication, encoded as UTF-8.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspoteapsettings/username
func (n_ NEHotspotEAPSettings) SetUsername(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setUsername:"), value)
}



