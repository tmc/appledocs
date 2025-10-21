// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [NEHotspotHS20Settings] class.
type INEHotspotHS20Settings interface {
	objectivec.IObject
}

// Settings for configuring Hotspot 2.0 Wi-Fi networks.
//
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

// Alloc allocates a new instance without initialization.
func (nc _NEHotspotHS20SettingsClass) Alloc() NEHotspotHS20Settings {
	rv := objc.Send[NEHotspotHS20Settings](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The domain name of a Hotspot 2.0 Wi-Fi Network.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspoths20settings/domainname
func (n_ NEHotspotHS20Settings) DomainName() appkit.string {
	rv := objc.Send[appkit.string](n_.ID, objc.Sel("domainName"))
	return rv
}


// SetDomainName sets the value of the domainName property.
// The domain name of a Hotspot 2.0 Wi-Fi Network.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspoths20settings/domainname
func (n_ NEHotspotHS20Settings) SetDomainName(value appkit.string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDomainName:"), value)
}

// A Boolean value indicating whether or not roaming is enabled on a Hotspot 2.0 Wi-Fi network.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspoths20settings/isroamingenabled
func (n_ NEHotspotHS20Settings) IsRoamingEnabled() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("isRoamingEnabled"))
	return rv
}


// SetIsRoamingEnabled sets the value of the isRoamingEnabled property.
// A Boolean value indicating whether or not roaming is enabled on a Hotspot 2.0 Wi-Fi network.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspoths20settings/isroamingenabled
func (n_ NEHotspotHS20Settings) SetIsRoamingEnabled(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIsRoamingEnabled:"), value)
}

// An array of Mobile Country Code (MCC) and Mobile Network Code (MNC) pairs used for Wi-Fi Hotspot 2.0 negotiation.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspoths20settings/mccandmncs
func (n_ NEHotspotHS20Settings) MccAndMNCs() appkit.string {
	rv := objc.Send[appkit.string](n_.ID, objc.Sel("mccAndMNCs"))
	return rv
}


// SetMccAndMNCs sets the value of the mccAndMNCs property.
// An array of Mobile Country Code (MCC) and Mobile Network Code (MNC) pairs used for Wi-Fi Hotspot 2.0 negotiation.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspoths20settings/mccandmncs
func (n_ NEHotspotHS20Settings) SetMccAndMNCs(value appkit.string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMccAndMNCs:"), value)
}

// An array of Network Access Identifier (NAI) realm name strings used for Wi-Fi Hotspot 2.0 negotiation.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspoths20settings/nairealmnames
func (n_ NEHotspotHS20Settings) NaiRealmNames() appkit.string {
	rv := objc.Send[appkit.string](n_.ID, objc.Sel("naiRealmNames"))
	return rv
}


// SetNaiRealmNames sets the value of the naiRealmNames property.
// An array of Network Access Identifier (NAI) realm name strings used for Wi-Fi Hotspot 2.0 negotiation.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspoths20settings/nairealmnames
func (n_ NEHotspotHS20Settings) SetNaiRealmNames(value appkit.string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setNaiRealmNames:"), value)
}

// An array of Roaming Consortium Organization (RCO) identifiers used for Wi-Fi Hotspot 2.0 negotiation.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspoths20settings/roamingconsortiumois
func (n_ NEHotspotHS20Settings) RoamingConsortiumOIs() appkit.string {
	rv := objc.Send[appkit.string](n_.ID, objc.Sel("roamingConsortiumOIs"))
	return rv
}


// SetRoamingConsortiumOIs sets the value of the roamingConsortiumOIs property.
// An array of Roaming Consortium Organization (RCO) identifiers used for Wi-Fi Hotspot 2.0 negotiation.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspoths20settings/roamingconsortiumois
func (n_ NEHotspotHS20Settings) SetRoamingConsortiumOIs(value appkit.string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setRoamingConsortiumOIs:"), value)
}



