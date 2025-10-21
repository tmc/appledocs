// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NEHotspotConfiguration] class.
var (
	NEHotspotConfigurationClass     _NEHotspotConfigurationClass
	NEHotspotConfigurationClassOnce sync.Once
)

func getNEHotspotConfigurationClass() _NEHotspotConfigurationClass {
	NEHotspotConfigurationClassOnce.Do(func() {
		NEHotspotConfigurationClass = _NEHotspotConfigurationClass{objc.GetClass("NEHotspotConfiguration")}
	})
	return NEHotspotConfigurationClass
}

type _NEHotspotConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [NEHotspotConfiguration] class.
type INEHotspotConfiguration interface {
	objectivec.IObject
}

// Configuration settings for a Wi-Fi network.
//
// The class contains configuration properties and credentials required to connect to Wi-Fi networks.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotConfiguration
type NEHotspotConfiguration struct {
	objectivec.Object
}

// NEHotspotConfigurationFrom constructs a [NEHotspotConfiguration] from an unsafe.Pointer.
//
// Configuration settings for a Wi-Fi network.
func NEHotspotConfigurationFrom(ptr unsafe.Pointer) NEHotspotConfiguration {
	return NEHotspotConfiguration{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (nc _NEHotspotConfigurationClass) Alloc() NEHotspotConfiguration {
	rv := objc.Send[NEHotspotConfiguration](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NEHotspotConfigurationClass) New() NEHotspotConfiguration {
	rv := objc.Send[NEHotspotConfiguration](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEHotspotConfiguration) Init() NEHotspotConfiguration {
	rv := objc.Send[NEHotspotConfiguration](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEHotspotConfiguration) Autorelease() NEHotspotConfiguration {
	rv := objc.Send[NEHotspotConfiguration](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEHotspotConfiguration creates a new NEHotspotConfiguration instance.
func NewNEHotspotConfiguration() NEHotspotConfiguration {
	return getNEHotspotConfigurationClass().New()
}




// Creates a new hotspot configuration, identified by a domain name, for a Hotspot 2.0 Wi-Fi network with HS 2.0 and EAP settings.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotConfiguration/init(hs20Settings:eapSettings:)
func NewNEHotspotConfigurationWithHS20SettingsEapSettings(hs20Settings unsafe.Pointer, eapSettings unsafe.Pointer) NEHotspotConfiguration {
	instance := getNEHotspotConfigurationClass().Alloc()
	rv := objc.Send[NEHotspotConfiguration](instance.ID, objc.Sel("initWithHS20Settings:eapSettings:"), hs20Settings, eapSettings)
	rv.Autorelease()
	return rv
}



// Creates a new hotspot configuration, identified by an SSID, for an open Wi-Fi network.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotConfiguration/init(ssid:)
func NewNEHotspotConfigurationWithSSID(SSID string) NEHotspotConfiguration {
	instance := getNEHotspotConfigurationClass().Alloc()
	rv := objc.Send[NEHotspotConfiguration](instance.ID, objc.Sel("initWithSSID:"), objc.String(SSID))
	rv.Autorelease()
	return rv
}



// Creates a new hotspot configuration, identified by an SSID, for a WPA/WPA2 enterprise Wi-Fi network with EAP settings.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotConfiguration/init(ssid:eapSettings:)
func NewNEHotspotConfigurationWithSSIDEapSettings(SSID string, eapSettings unsafe.Pointer) NEHotspotConfiguration {
	instance := getNEHotspotConfigurationClass().Alloc()
	rv := objc.Send[NEHotspotConfiguration](instance.ID, objc.Sel("initWithSSID:eapSettings:"), objc.String(SSID), eapSettings)
	rv.Autorelease()
	return rv
}



// Creates a new hotspot configuration, identified by an SSID, for a protected WEP or WPA/WPA2 personal Wi-Fi network.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotConfiguration/init(ssid:passphrase:isWEP:)
func NewNEHotspotConfigurationWithSSIDPassphraseIsWEP(SSID string, passphrase string, isWEP bool) NEHotspotConfiguration {
	instance := getNEHotspotConfigurationClass().Alloc()
	rv := objc.Send[NEHotspotConfiguration](instance.ID, objc.Sel("initWithSSID:passphrase:isWEP:"), objc.String(SSID), objc.String(passphrase), isWEP)
	rv.Autorelease()
	return rv
}



// Creates a new hotspot configuration, identified by an SSID prefix string, for an open Wi-Fi network.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotConfiguration/init(ssidPrefix:)
func NewNEHotspotConfigurationWithSSIDPrefix(SSIDPrefix string) NEHotspotConfiguration {
	instance := getNEHotspotConfigurationClass().Alloc()
	rv := objc.Send[NEHotspotConfiguration](instance.ID, objc.Sel("initWithSSIDPrefix:"), objc.String(SSIDPrefix))
	rv.Autorelease()
	return rv
}



// Creates a new hotspot configuration, identified by an SSID prefix string, for a protected WEP or WPA/WPA2 personal Wi-Fi network.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotConfiguration/init(ssidPrefix:passphrase:isWEP:)
func NewNEHotspotConfigurationWithSSIDPrefixPassphraseIsWEP(SSIDPrefix string, passphrase string, isWEP bool) NEHotspotConfiguration {
	instance := getNEHotspotConfigurationClass().Alloc()
	rv := objc.Send[NEHotspotConfiguration](instance.ID, objc.Sel("initWithSSIDPrefix:passphrase:isWEP:"), objc.String(SSIDPrefix), objc.String(passphrase), isWEP)
	rv.Autorelease()
	return rv
}


// A Boolean value that indicates the visibility of the SSID.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotConfiguration/hidden
func (n_ NEHotspotConfiguration) Hidden() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("hidden"))
	return rv
}


// SetHidden sets the value of the hidden property.
// A Boolean value that indicates the visibility of the SSID.

//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotConfiguration/hidden
func (n_ NEHotspotConfiguration) SetHidden(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setHidden:"), value)
}

// Restricts the lifetime of a configuration to the operating status of the app that created it.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotConfiguration/joinOnce
func (n_ NEHotspotConfiguration) JoinOnce() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("joinOnce"))
	return rv
}


// SetJoinOnce sets the value of the joinOnce property.
// Restricts the lifetime of a configuration to the operating status of the app that created it.

//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotConfiguration/joinOnce
func (n_ NEHotspotConfiguration) SetJoinOnce(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setJoinOnce:"), value)
}

// The number of days the network retains the associated configuration.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotConfiguration/lifeTimeInDays
func (n_ NEHotspotConfiguration) LifeTimeInDays() foundation.Number {
	rv := objc.Send[foundation.Number](n_.ID, objc.Sel("lifeTimeInDays"))
	return rv
}


// SetLifeTimeInDays sets the value of the lifeTimeInDays property.
// The number of days the network retains the associated configuration.

//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotConfiguration/lifeTimeInDays
func (n_ NEHotspotConfiguration) SetLifeTimeInDays(value foundation.Number) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setLifeTimeInDays:"), value)
}

// The SSID of an open, WEP, WPA/WPA2 personal, or WPA/WPA2 enterprise Wi-Fi network.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotConfiguration/ssid
func (n_ NEHotspotConfiguration) SSID() string {
	rv := objc.Send[string](n_.ID, objc.Sel("SSID"))
	return rv
}

// The string used to match networks against a known SSID prefix.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotConfiguration/ssidPrefix
func (n_ NEHotspotConfiguration) SSIDPrefix() string {
	rv := objc.Send[string](n_.ID, objc.Sel("SSIDPrefix"))
	return rv
}


