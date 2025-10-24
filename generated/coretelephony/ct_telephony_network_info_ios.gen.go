//go:build darwin && ios

// Code generated from Apple documentation for CoreTelephony. DO NOT EDIT.

package coretelephony

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for TelephonyNetworkInfo


// iOS-only properties

// The current radio access technology registered with the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTTelephonyNetworkInfo/currentRadioAccessTechnology
func (t_ TelephonyNetworkInfo) CurrentRadioAccessTechnology() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("currentRadioAccessTechnology"))
	return rv
}

// The identifier of the service that’s currently providing data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTTelephonyNetworkInfo/dataServiceIdentifier
func (t_ TelephonyNetworkInfo) DataServiceIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("dataServiceIdentifier"))
	return rv
}

// A dictionary containing the current radio access technology registered to each service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTTelephonyNetworkInfo/serviceCurrentRadioAccessTechnology
func (t_ TelephonyNetworkInfo) ServiceCurrentRadioAccessTechnology() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](t_.ID, objc.Sel("serviceCurrentRadioAccessTechnology"))
	return rv
}

// A dictionary that contains carrier information about each service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTTelephonyNetworkInfo/serviceSubscriberCellularProviders
func (t_ TelephonyNetworkInfo) ServiceSubscriberCellularProviders() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](t_.ID, objc.Sel("serviceSubscriberCellularProviders"))
	return rv
}

// A block dispatched when there are updates to the user’s cellular provider information for any service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTTelephonyNetworkInfo/serviceSubscriberCellularProvidersDidUpdateNotifier
func (t_ TelephonyNetworkInfo) ServiceSubscriberCellularProvidersDidUpdateNotifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("serviceSubscriberCellularProvidersDidUpdateNotifier"))
	return rv
}
func (t_ TelephonyNetworkInfo) SetServiceSubscriberCellularProvidersDidUpdateNotifier(value unsafe.Pointer) {
	t_.ID.Send(objc.RegisterName("setServiceSubscriberCellularProvidersDidUpdateNotifier:"), value)
}

// Information about the user’s cellular service provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTTelephonyNetworkInfo/subscriberCellularProvider
func (t_ TelephonyNetworkInfo) SubscriberCellularProvider() ICTCarrier {
	rv := objc.Send[Carrier](t_.ID, objc.Sel("subscriberCellularProvider"))
	return rv
}

// A block dispatched when the user’s cellular service provider information changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTTelephonyNetworkInfo/subscriberCellularProviderDidUpdateNotifier
func (t_ TelephonyNetworkInfo) SubscriberCellularProviderDidUpdateNotifier() func(unsafe.Pointer) {
	rv := objc.Send[func(unsafe.Pointer)](t_.ID, objc.Sel("subscriberCellularProviderDidUpdateNotifier"))
	return rv
}
func (t_ TelephonyNetworkInfo) SetSubscriberCellularProviderDidUpdateNotifier(value func(unsafe.Pointer)) {
	t_.ID.Send(objc.RegisterName("setSubscriberCellularProviderDidUpdateNotifier:"), value)
}






