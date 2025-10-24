// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [NEDNSOverHTTPSSettings] class.
var (
	NEDNSOverHTTPSSettingsClass     _NEDNSOverHTTPSSettingsClass
	NEDNSOverHTTPSSettingsClassOnce sync.Once
)

func getNEDNSOverHTTPSSettingsClass() _NEDNSOverHTTPSSettingsClass {
	NEDNSOverHTTPSSettingsClassOnce.Do(func() {
		NEDNSOverHTTPSSettingsClass = _NEDNSOverHTTPSSettingsClass{objc.GetClass("NEDNSOverHTTPSSettings")}
	})
	return NEDNSOverHTTPSSettingsClass
}

type _NEDNSOverHTTPSSettingsClass struct {
	class objc.Class
}

// An interface definition for the [NEDNSOverHTTPSSettings] class.
type INEDNSOverHTTPSSettings interface {
	INEDNSSettings
	// properties:
	IdentityReference() objc.IObject /* cross-framework: NSData */
	SetIdentityReference(value objc.IObject /* cross-framework: NSData */)
	ServerURL() objc.IObject /* cross-framework: URL */
	SetServerURL(value objc.IObject /* cross-framework: URL */)
	MatchDomains() objc.IObject /* cross-framework: NSString */
	SetMatchDomains(value objc.IObject /* cross-framework: NSString */)
	// methods:
}

// The DNS resolver settings for a DNS-over-HTTPS server.


// The DNS resolver settings for a DNS-over-HTTPS server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSOverHTTPSSettings
type NEDNSOverHTTPSSettings struct {
	NEDNSSettings
}

// NEDNSOverHTTPSSettingsFrom constructs a [NEDNSOverHTTPSSettings] from an unsafe.Pointer.
//
// The DNS resolver settings for a DNS-over-HTTPS server.
func NEDNSOverHTTPSSettingsFrom(ptr unsafe.Pointer) NEDNSOverHTTPSSettings {
	return NEDNSOverHTTPSSettings{
		NEDNSSettings: NEDNSSettingsFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (nc _NEDNSOverHTTPSSettingsClass) Alloc() NEDNSOverHTTPSSettings {
	rv := objc.Send[NEDNSOverHTTPSSettings](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NEDNSOverHTTPSSettingsClass) New() NEDNSOverHTTPSSettings {
	rv := objc.Send[NEDNSOverHTTPSSettings](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEDNSOverHTTPSSettings) Init() NEDNSOverHTTPSSettings {
	rv := objc.Send[NEDNSOverHTTPSSettings](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEDNSOverHTTPSSettings) Autorelease() NEDNSOverHTTPSSettings {
	rv := objc.Send[NEDNSOverHTTPSSettings](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEDNSOverHTTPSSettings creates a new NEDNSOverHTTPSSettings instance.
func NewNEDNSOverHTTPSSettings() NEDNSOverHTTPSSettings {
	return getNEDNSOverHTTPSSettingsClass().New()
}



// A persistent keychain reference to a keychain item containing the certificate and private key components of the DNS client credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSOverHTTPSSettings/identityReference
func (n_ NEDNSOverHTTPSSettings) IdentityReference() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](n_.ID, objc.Sel("identityReference"))
	return rv
}


// A persistent keychain reference to a keychain item containing the certificate and private key components of the DNS client credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSOverHTTPSSettings/identityReference
func (n_ NEDNSOverHTTPSSettings) SetIdentityReference(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIdentityReference:"), value)
}


// The URL of a DNS-over-HTTPS server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nednsoverhttpssettings/serverurl
func (n_ NEDNSOverHTTPSSettings) ServerURL() objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[foundation.URL](n_.ID, objc.Sel("serverURL"))
	return rv
}


// The URL of a DNS-over-HTTPS server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nednsoverhttpssettings/serverurl
func (n_ NEDNSOverHTTPSSettings) SetServerURL(value objc.IObject /* cross-framework: URL */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setServerURL:"), value)
}


// A list of domain strings used to determine which DNS queries will use the DNS resolver settings contained in this object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nednssettings/matchdomains
func (n_ NEDNSOverHTTPSSettings) MatchDomains() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("matchDomains"))
	return rv
}


// A list of domain strings used to determine which DNS queries will use the DNS resolver settings contained in this object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nednssettings/matchdomains
func (n_ NEDNSOverHTTPSSettings) SetMatchDomains(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMatchDomains:"), value)
}



