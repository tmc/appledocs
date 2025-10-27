// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [NEDNSOverTLSSettings] class.
var (
	NEDNSOverTLSSettingsClass     _NEDNSOverTLSSettingsClass
	NEDNSOverTLSSettingsClassOnce sync.Once
)

func getNEDNSOverTLSSettingsClass() _NEDNSOverTLSSettingsClass {
	NEDNSOverTLSSettingsClassOnce.Do(func() {
		NEDNSOverTLSSettingsClass = _NEDNSOverTLSSettingsClass{objc.GetClass("NEDNSOverTLSSettings")}
	})
	return NEDNSOverTLSSettingsClass
}

type _NEDNSOverTLSSettingsClass struct {
	class objc.Class
}





// An interface definition for the [NEDNSOverTLSSettings] class.
type INEDNSOverTLSSettings interface {
	INEDNSSettings
	

	// properties:
	IdentityReference() foundation.foundation.INSData
	SetIdentityReference(value foundation.foundation.INSData)
	ServerName() foundation.foundation.INSString
	SetServerName(value foundation.foundation.INSString)
	MatchDomains() foundation.foundation.INSString
	SetMatchDomains(value foundation.foundation.INSString)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (nc _NEDNSOverTLSSettingsClass) Alloc() NEDNSOverTLSSettings {
	rv := objc.Send[NEDNSOverTLSSettings](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NEDNSOverTLSSettingsClass) New() NEDNSOverTLSSettings {
	rv := objc.Send[NEDNSOverTLSSettings](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEDNSOverTLSSettings) Init() NEDNSOverTLSSettings {
	rv := objc.Send[NEDNSOverTLSSettings](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEDNSOverTLSSettings) Autorelease() NEDNSOverTLSSettings {
	rv := objc.Send[NEDNSOverTLSSettings](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEDNSOverTLSSettings creates a new NEDNSOverTLSSettings instance.
func NewNEDNSOverTLSSettings() NEDNSOverTLSSettings {
	return getNEDNSOverTLSSettingsClass().New()
}





// The DNS resolver settings for a DNS-over-TLS server.


// The DNS resolver settings for a DNS-over-TLS server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSOverTLSSettings
type NEDNSOverTLSSettings struct {
	NEDNSSettings
}

// NEDNSOverTLSSettingsFrom constructs a [NEDNSOverTLSSettings] from an unsafe.Pointer.
//
// The DNS resolver settings for a DNS-over-TLS server.
func NEDNSOverTLSSettingsFrom(ptr unsafe.Pointer) NEDNSOverTLSSettings {
	return NEDNSOverTLSSettings{
		NEDNSSettings: NEDNSSettingsFrom(ptr),
	}
}

























// A persistent keychain reference to a keychain item containing the certificate and private key components of the DNS client credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSOverTLSSettings/identityReference
func (n_ NEDNSOverTLSSettings) IdentityReference() foundation.foundation.INSData {
	rv := objc.Send[foundation.NSData](n_.ID, objc.Sel("identityReference"))
	return rv
}


// A persistent keychain reference to a keychain item containing the certificate and private key components of the DNS client credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSOverTLSSettings/identityReference
func (n_ NEDNSOverTLSSettings) SetIdentityReference(value foundation.foundation.INSData) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIdentityReference:"), value)
}


// The TLS name of a DNS-over-TLS server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSOverTLSSettings/serverName
func (n_ NEDNSOverTLSSettings) ServerName() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("serverName"))
	return rv
}


// The TLS name of a DNS-over-TLS server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSOverTLSSettings/serverName
func (n_ NEDNSOverTLSSettings) SetServerName(value foundation.foundation.INSString) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setServerName:"), value)
}


// A list of domain strings used to determine which DNS queries will use the DNS resolver settings contained in this object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nednssettings/matchdomains
func (n_ NEDNSOverTLSSettings) MatchDomains() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("matchDomains"))
	return rv
}


// A list of domain strings used to determine which DNS queries will use the DNS resolver settings contained in this object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nednssettings/matchdomains
func (n_ NEDNSOverTLSSettings) SetMatchDomains(value foundation.foundation.INSString) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMatchDomains:"), value)
}








