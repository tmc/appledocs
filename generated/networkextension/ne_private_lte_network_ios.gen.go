//go:build darwin && ios

// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for NEPrivateLTENetwork


// iOS-only properties

// The Mobile Country Code (MCC) of the private LTE network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEPrivateLTENetwork/mobileCountryCode
func (n_ NEPrivateLTENetwork) MobileCountryCode() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("mobileCountryCode"))
	return rv
}
func (n_ NEPrivateLTENetwork) SetMobileCountryCode(value foundation.foundation.INSString) {
	n_.ID.Send(objc.RegisterName("setMobileCountryCode:"), value)
}

// The Mobile Network Code (MNC) of the private LTE network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEPrivateLTENetwork/mobileNetworkCode
func (n_ NEPrivateLTENetwork) MobileNetworkCode() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("mobileNetworkCode"))
	return rv
}
func (n_ NEPrivateLTENetwork) SetMobileNetworkCode(value foundation.foundation.INSString) {
	n_.ID.Send(objc.RegisterName("setMobileNetworkCode:"), value)
}

// The Tracking Area Code of the private LTE network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEPrivateLTENetwork/trackingAreaCode
func (n_ NEPrivateLTENetwork) TrackingAreaCode() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("trackingAreaCode"))
	return rv
}
func (n_ NEPrivateLTENetwork) SetTrackingAreaCode(value foundation.foundation.INSString) {
	n_.ID.Send(objc.RegisterName("setTrackingAreaCode:"), value)
}





