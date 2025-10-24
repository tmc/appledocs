//go:build darwin && ios

// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for NEPrivateLTENetwork


// iOS-only properties

// The Mobile Country Code (MCC) of the private LTE network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEPrivateLTENetwork/mobileCountryCode
func (n_ NEPrivateLTENetwork) MobileCountryCode() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("mobileCountryCode"))
	return rv
}
func (n_ NEPrivateLTENetwork) SetMobileCountryCode(value objc.IObject /* cross-framework: NSString */) {
	n_.ID.Send(objc.RegisterName("setMobileCountryCode:"), value)
}

// The Mobile Network Code (MNC) of the private LTE network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEPrivateLTENetwork/mobileNetworkCode
func (n_ NEPrivateLTENetwork) MobileNetworkCode() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("mobileNetworkCode"))
	return rv
}
func (n_ NEPrivateLTENetwork) SetMobileNetworkCode(value objc.IObject /* cross-framework: NSString */) {
	n_.ID.Send(objc.RegisterName("setMobileNetworkCode:"), value)
}

// The Tracking Area Code of the private LTE network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEPrivateLTENetwork/trackingAreaCode
func (n_ NEPrivateLTENetwork) TrackingAreaCode() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("trackingAreaCode"))
	return rv
}
func (n_ NEPrivateLTENetwork) SetTrackingAreaCode(value objc.IObject /* cross-framework: NSString */) {
	n_.ID.Send(objc.RegisterName("setTrackingAreaCode:"), value)
}





