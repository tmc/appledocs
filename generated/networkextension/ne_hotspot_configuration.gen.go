// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NEHotspotConfiguration */


/* debug [class_header]: Header for NEHotspotConfiguration */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NEHotspotConfiguration */
// An interface definition for the [NEHotspotConfiguration] class.
type INEHotspotConfiguration interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for NEHotspotConfiguration */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NEHotspotConfiguration */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NEHotspotConfiguration */
// Alloc allocates a new instance without initialization.
func (nc _NEHotspotConfigurationClass) Alloc() NEHotspotConfiguration {
	rv := objc.Send[NEHotspotConfiguration](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NEHotspotConfiguration */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NEHotspotConfiguration */

// Creates a new hotspot configuration, identified by a domain name, for a Hotspot 2.0 Wi-Fi network with HS 2.0 and EAP settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotConfiguration/init(hs20Settings:eapSettings:)
func NewNEHotspotConfigurationWithHS20SettingsEapSettings(hs20Settings INEHotspotHS20Settings, eapSettings INEHotspotEAPSettings) NEHotspotConfiguration {
	instance := getNEHotspotConfigurationClass().Alloc()
	rv := objc.Send[NEHotspotConfiguration](instance.ID, objc.Sel("initWithHS20Settings:eapSettings:"), hs20Settings, eapSettings)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewNEHotspotConfigurationWithHS20SettingsEapSettings */


// Creates a new hotspot configuration, identified by an SSID, for an open Wi-Fi network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotConfiguration/init(ssid:)
func NewNEHotspotConfigurationWithSSID(SSID objc.IObject /* cross-framework: NSString */) NEHotspotConfiguration {
	instance := getNEHotspotConfigurationClass().Alloc()
	rv := objc.Send[NEHotspotConfiguration](instance.ID, objc.Sel("initWithSSID:"), SSID)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewNEHotspotConfigurationWithSSID */


// Creates a new hotspot configuration, identified by an SSID, for a WPA/WPA2 enterprise Wi-Fi network with EAP settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotConfiguration/init(ssid:eapSettings:)
func NewNEHotspotConfigurationWithSSIDEapSettings(SSID objc.IObject /* cross-framework: NSString */, eapSettings INEHotspotEAPSettings) NEHotspotConfiguration {
	instance := getNEHotspotConfigurationClass().Alloc()
	rv := objc.Send[NEHotspotConfiguration](instance.ID, objc.Sel("initWithSSID:eapSettings:"), SSID, eapSettings)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewNEHotspotConfigurationWithSSIDEapSettings */


// Creates a new hotspot configuration, identified by an SSID, for a protected WEP or WPA/WPA2 personal Wi-Fi network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotConfiguration/init(ssid:passphrase:isWEP:)
func NewNEHotspotConfigurationWithSSIDPassphraseIsWEP(SSID objc.IObject /* cross-framework: NSString */, passphrase objc.IObject /* cross-framework: NSString */, isWEP bool) NEHotspotConfiguration {
	instance := getNEHotspotConfigurationClass().Alloc()
	rv := objc.Send[NEHotspotConfiguration](instance.ID, objc.Sel("initWithSSID:passphrase:isWEP:"), SSID, passphrase, isWEP)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewNEHotspotConfigurationWithSSIDPassphraseIsWEP */


// Creates a new hotspot configuration, identified by an SSID prefix string, for an open Wi-Fi network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotConfiguration/init(ssidPrefix:)
func NewNEHotspotConfigurationWithSSIDPrefix(SSIDPrefix objc.IObject /* cross-framework: NSString */) NEHotspotConfiguration {
	instance := getNEHotspotConfigurationClass().Alloc()
	rv := objc.Send[NEHotspotConfiguration](instance.ID, objc.Sel("initWithSSIDPrefix:"), SSIDPrefix)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewNEHotspotConfigurationWithSSIDPrefix */


// Creates a new hotspot configuration, identified by an SSID prefix string, for a protected WEP or WPA/WPA2 personal Wi-Fi network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotConfiguration/init(ssidPrefix:passphrase:isWEP:)
func NewNEHotspotConfigurationWithSSIDPrefixPassphraseIsWEP(SSIDPrefix objc.IObject /* cross-framework: NSString */, passphrase objc.IObject /* cross-framework: NSString */, isWEP bool) NEHotspotConfiguration {
	instance := getNEHotspotConfigurationClass().Alloc()
	rv := objc.Send[NEHotspotConfiguration](instance.ID, objc.Sel("initWithSSIDPrefix:passphrase:isWEP:"), SSIDPrefix, passphrase, isWEP)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewNEHotspotConfigurationWithSSIDPrefixPassphraseIsWEP */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NEHotspotConfiguration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NEHotspotConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NEHotspotConfiguration */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NEHotspotConfiguration */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NEHotspotConfiguration */


