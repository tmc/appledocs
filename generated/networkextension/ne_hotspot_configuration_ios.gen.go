//go:build darwin && ios

// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for NEHotspotConfiguration


// iOS-only properties

// A Boolean value that indicates the visibility of the SSID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotConfiguration/hidden
func (n_ NEHotspotConfiguration) Hidden() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("hidden"))
	return rv
}
func (n_ NEHotspotConfiguration) SetHidden(value bool) {
	n_.ID.Send(objc.RegisterName("setHidden:"), value)
}

// Restricts the lifetime of a configuration to the operating status of the app that created it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotConfiguration/joinOnce
func (n_ NEHotspotConfiguration) JoinOnce() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("joinOnce"))
	return rv
}
func (n_ NEHotspotConfiguration) SetJoinOnce(value bool) {
	n_.ID.Send(objc.RegisterName("setJoinOnce:"), value)
}

// The number of days the network retains the associated configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotConfiguration/lifeTimeInDays
func (n_ NEHotspotConfiguration) LifeTimeInDays() foundation.foundation.INSNumber {
	rv := objc.Send[foundation.NSNumber](n_.ID, objc.Sel("lifeTimeInDays"))
	return rv
}
func (n_ NEHotspotConfiguration) SetLifeTimeInDays(value foundation.foundation.INSNumber) {
	n_.ID.Send(objc.RegisterName("setLifeTimeInDays:"), value)
}

// The SSID of an open, WEP, WPA/WPA2 personal, or WPA/WPA2 enterprise Wi-Fi network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotConfiguration/ssid
func (n_ NEHotspotConfiguration) SSID() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("SSID"))
	return rv
}

// The string used to match networks against a known SSID prefix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotConfiguration/ssidPrefix
func (n_ NEHotspotConfiguration) SSIDPrefix() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("SSIDPrefix"))
	return rv
}




