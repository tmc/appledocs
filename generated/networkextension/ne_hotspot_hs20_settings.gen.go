// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NEHotspotHS20Settings */


/* debug [class_header]: Header for NEHotspotHS20Settings */
// The class instance for the [NEHotspotHS20Settings] class.
var (
	NEHotspotHS20SettingsClass     _NEHotspotHS20SettingsClass
	NEHotspotHS20SettingsClassOnce sync.Once
)

func getNEHotspotHS20SettingsClass() _NEHotspotHS20SettingsClass {
	NEHotspotHS20SettingsClassOnce.Do(func() {
		NEHotspotHS20SettingsClass = _NEHotspotHS20SettingsClass{objc.GetClass("NEHotspotHS20Settings")}
	})
	return NEHotspotHS20SettingsClass
}

type _NEHotspotHS20SettingsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NEHotspotHS20Settings */
// An interface definition for the [NEHotspotHS20Settings] class.
type INEHotspotHS20Settings interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for NEHotspotHS20Settings */
	// properties:
	IsRoamingEnabled() bool
	SetIsRoamingEnabled(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NEHotspotHS20Settings */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NEHotspotHS20Settings */
// Alloc allocates a new instance without initialization.
func (nc _NEHotspotHS20SettingsClass) Alloc() NEHotspotHS20Settings {
	rv := objc.Send[NEHotspotHS20Settings](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NEHotspotHS20SettingsClass) New() NEHotspotHS20Settings {
	rv := objc.Send[NEHotspotHS20Settings](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEHotspotHS20Settings) Init() NEHotspotHS20Settings {
	rv := objc.Send[NEHotspotHS20Settings](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEHotspotHS20Settings) Autorelease() NEHotspotHS20Settings {
	rv := objc.Send[NEHotspotHS20Settings](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEHotspotHS20Settings creates a new NEHotspotHS20Settings instance.
func NewNEHotspotHS20Settings() NEHotspotHS20Settings {
	return getNEHotspotHS20SettingsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NEHotspotHS20Settings */
// Settings for configuring Hotspot 2.0 Wi-Fi networks.


// Settings for configuring Hotspot 2.0 Wi-Fi networks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotHS20Settings
type NEHotspotHS20Settings struct {
	objectivec.Object
}

// NEHotspotHS20SettingsFrom constructs a [NEHotspotHS20Settings] from an unsafe.Pointer.
//
// Settings for configuring Hotspot 2.0 Wi-Fi networks.
func NEHotspotHS20SettingsFrom(ptr unsafe.Pointer) NEHotspotHS20Settings {
	return NEHotspotHS20Settings{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NEHotspotHS20Settings */

// Creates a new hotspot configuration of a legacy Hotspot or HS 2.0 Wi-Fi network. with optional roaming enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotHS20Settings/init(domainName:roamingEnabled:)
func NewNEHotspotHS20SettingsWithDomainNameRoamingEnabled(domainName objc.IObject /* cross-framework: NSString */, roamingEnabled bool) NEHotspotHS20Settings {
	instance := getNEHotspotHS20SettingsClass().Alloc()
	rv := objc.Send[NEHotspotHS20Settings](instance.ID, objc.Sel("initWithDomainName:roamingEnabled:"), domainName, roamingEnabled)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewNEHotspotHS20SettingsWithDomainNameRoamingEnabled */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NEHotspotHS20Settings */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NEHotspotHS20Settings */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NEHotspotHS20Settings */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NEHotspotHS20Settings */

// A Boolean value indicating whether or not roaming is enabled on a Hotspot 2.0 Wi-Fi network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspoths20settings/isroamingenabled
func (n_ NEHotspotHS20Settings) IsRoamingEnabled() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("isRoamingEnabled"))
	return rv
}/* debug [instance_properties/getter]: isRoamingEnabled */


// A Boolean value indicating whether or not roaming is enabled on a Hotspot 2.0 Wi-Fi network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspoths20settings/isroamingenabled
func (n_ NEHotspotHS20Settings) SetIsRoamingEnabled(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIsRoamingEnabled:"), value)
}/* debug [instance_properties/setter]: isRoamingEnabled */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NEHotspotHS20Settings */


