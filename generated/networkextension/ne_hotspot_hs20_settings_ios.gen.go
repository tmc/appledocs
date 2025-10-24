//go:build darwin && ios

// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for NEHotspotHS20Settings


// iOS-only properties

// The domain name of a Hotspot 2.0 Wi-Fi Network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotHS20Settings/domainName
func (n_ NEHotspotHS20Settings) DomainName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("domainName"))
	return rv
}

// A Boolean value indicating whether or not roaming is enabled on a Hotspot 2.0 Wi-Fi network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotHS20Settings/isRoamingEnabled
func (n_ NEHotspotHS20Settings) RoamingEnabled() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("roamingEnabled"))
	return rv
}
func (n_ NEHotspotHS20Settings) SetRoamingEnabled(value bool) {
	n_.ID.Send(objc.RegisterName("setRoamingEnabled:"), value)
}

// An array of Mobile Country Code (MCC) and Mobile Network Code (MNC) pairs used for Wi-Fi Hotspot 2.0 negotiation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotHS20Settings/mccAndMNCs
func (n_ NEHotspotHS20Settings) MCCAndMNCs() []string {
	rv := objc.Send[[]string](n_.ID, objc.Sel("MCCAndMNCs"))
	return rv
}
func (n_ NEHotspotHS20Settings) SetMCCAndMNCs(value []string) {
	n_.ID.Send(objc.RegisterName("setMCCAndMNCs:"), value)
}

// An array of Network Access Identifier (NAI) realm name strings used for Wi-Fi Hotspot 2.0 negotiation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotHS20Settings/naiRealmNames
func (n_ NEHotspotHS20Settings) NaiRealmNames() []string {
	rv := objc.Send[[]string](n_.ID, objc.Sel("naiRealmNames"))
	return rv
}
func (n_ NEHotspotHS20Settings) SetNaiRealmNames(value []string) {
	n_.ID.Send(objc.RegisterName("setNaiRealmNames:"), value)
}

// An array of Roaming Consortium Organization (RCO) identifiers used for Wi-Fi Hotspot 2.0 negotiation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotHS20Settings/roamingConsortiumOIs
func (n_ NEHotspotHS20Settings) RoamingConsortiumOIs() []string {
	rv := objc.Send[[]string](n_.ID, objc.Sel("roamingConsortiumOIs"))
	return rv
}
func (n_ NEHotspotHS20Settings) SetRoamingConsortiumOIs(value []string) {
	n_.ID.Send(objc.RegisterName("setRoamingConsortiumOIs:"), value)
}




