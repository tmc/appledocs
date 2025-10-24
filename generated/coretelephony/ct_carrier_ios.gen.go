//go:build darwin && ios

// Code generated from Apple documentation for CoreTelephony. DO NOT EDIT.

package coretelephony

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for Carrier


// iOS-only properties

// The name of the user’s home cellular service provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTCarrier/carrierName
func (c_ Carrier) CarrierName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("carrierName"))
	return rv
}

// The mobile country code (MCC) for the user’s cellular service provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTCarrier/mobileCountryCode
func (c_ Carrier) MobileCountryCode() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("mobileCountryCode"))
	return rv
}

// The mobile network code for the user’s cellular service provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTCarrier/mobileNetworkCode
func (c_ Carrier) MobileNetworkCode() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("mobileNetworkCode"))
	return rv
}





