//go:build darwin && ios

// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for NEHotspotNetwork


// Indicate the level of confidence in being able to handle the network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotNetwork/setConfidence(_:)
func (n_ NEHotspotNetwork) SetConfidence(confidence NEHotspotHelperConfidence) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setConfidence:"), confidence)
}

// Provide the password for a protected network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotNetwork/setPassword(_:)
func (n_ NEHotspotNetwork) SetPassword(password foundation.foundation.INSString) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setPassword:"), password)
}

// iOS-only properties

// The BSSID for the Wi-Fi network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotNetwork/bssid
func (n_ NEHotspotNetwork) BSSID() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("BSSID"))
	return rv
}

// Indicates whether the network was joined automatically or was joined explicitly by the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotNetwork/didAutoJoin
func (n_ NEHotspotNetwork) AutoJoined() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("autoJoined"))
	return rv
}

// Indicates whether the network was just joined.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotNetwork/didJustJoin
func (n_ NEHotspotNetwork) JustJoined() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("justJoined"))
	return rv
}

// Indicates whether the calling Hotspot Helper is the chosen helper for this network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotNetwork/isChosenHelper
func (n_ NEHotspotNetwork) ChosenHelper() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("chosenHelper"))
	return rv
}

// Indicates whether the network is secure
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotNetwork/isSecure
func (n_ NEHotspotNetwork) Secure() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("secure"))
	return rv
}

// The type of security used by the Wi-Fi network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotNetwork/securityType
func (n_ NEHotspotNetwork) SecurityType() NEHotspotNetworkSecurityType {
	rv := objc.Send[NEHotspotNetworkSecurityType](n_.ID, objc.Sel("securityType"))
	return rv
}

// The recent signal strength for the Wi-Fi network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotNetwork/signalStrength
func (n_ NEHotspotNetwork) SignalStrength() float64 {
	rv := objc.Send[float64](n_.ID, objc.Sel("signalStrength"))
	return rv
}

// The SSID for the Wi-Fi network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotNetwork/ssid
func (n_ NEHotspotNetwork) SSID() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("SSID"))
	return rv
}





