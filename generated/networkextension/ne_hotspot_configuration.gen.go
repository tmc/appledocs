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
	// properties:
	Hidden() bool
	SetHidden(value bool)
	LifeTimeInDays() objc.IObject /* cross-framework: NSNumber */
	SetLifeTimeInDays(value objc.IObject /* cross-framework: NSNumber */)
	Ssid() objc.IObject /* cross-framework: NSString */
	SetSsid(value objc.IObject /* cross-framework: NSString */)
	SsidPrefix() objc.IObject /* cross-framework: NSString */
	SetSsidPrefix(value objc.IObject /* cross-framework: NSString */)
	// methods:
}

// Configuration settings for a Wi-Fi network.
//
// The class contains configuration properties and credentials required to connect to Wi-Fi networks.


// Configuration settings for a Wi-Fi network.
//
// [Full Topic]
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



// Creates a new hotspot configuration, identified by an SSID, for an open Wi-Fi network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotConfiguration/init(ssid:)
func NewNEHotspotConfigurationWithSSID(SSID objc.IObject /* cross-framework: NSString */) NEHotspotConfiguration {
	instance := getNEHotspotConfigurationClass().Alloc()
	rv := objc.Send[NEHotspotConfiguration](instance.ID, objc.Sel("initWithSSID:"), SSID)
	rv.Autorelease()
	return rv
}



// A Boolean value that indicates the visibility of the SSID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspotconfiguration/hidden
func (n_ NEHotspotConfiguration) Hidden() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("hidden"))
	return rv
}


// A Boolean value that indicates the visibility of the SSID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspotconfiguration/hidden
func (n_ NEHotspotConfiguration) SetHidden(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setHidden:"), value)
}


// The number of days the network retains the associated configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspotconfiguration/lifetimeindays
func (n_ NEHotspotConfiguration) LifeTimeInDays() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](n_.ID, objc.Sel("lifeTimeInDays"))
	return rv
}


// The number of days the network retains the associated configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspotconfiguration/lifetimeindays
func (n_ NEHotspotConfiguration) SetLifeTimeInDays(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setLifeTimeInDays:"), value)
}


// The SSID of an open, WEP, WPA/WPA2 personal, or WPA/WPA2 enterprise Wi-Fi network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspotconfiguration/ssid
func (n_ NEHotspotConfiguration) Ssid() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("ssid"))
	return rv
}


// The SSID of an open, WEP, WPA/WPA2 personal, or WPA/WPA2 enterprise Wi-Fi network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspotconfiguration/ssid
func (n_ NEHotspotConfiguration) SetSsid(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setSsid:"), value)
}


// The string used to match networks against a known SSID prefix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspotconfiguration/ssidprefix
func (n_ NEHotspotConfiguration) SsidPrefix() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("ssidPrefix"))
	return rv
}


// The string used to match networks against a known SSID prefix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspotconfiguration/ssidprefix
func (n_ NEHotspotConfiguration) SetSsidPrefix(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setSsidPrefix:"), value)
}


