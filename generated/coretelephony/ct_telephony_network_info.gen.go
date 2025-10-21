// Code generated from Apple documentation for CoreTelephony. DO NOT EDIT.

package coretelephony

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TelephonyNetworkInfo] class.
var (
	TelephonyNetworkInfoClass     _TelephonyNetworkInfoClass
	TelephonyNetworkInfoClassOnce sync.Once
)

func getTelephonyNetworkInfoClass() _TelephonyNetworkInfoClass {
	TelephonyNetworkInfoClassOnce.Do(func() {
		TelephonyNetworkInfoClass = _TelephonyNetworkInfoClass{objc.GetClass("CTTelephonyNetworkInfo")}
	})
	return TelephonyNetworkInfoClass
}

type _TelephonyNetworkInfoClass struct {
	class objc.Class
}

// An interface definition for the [TelephonyNetworkInfo] class.
type ITelephonyNetworkInfo interface {
	objectivec.IObject
}

// An object that provides notifications of changes to the user’s cellular service provider.
//
// Your app should be able to handle changes to the user’s cellular service provider. For example, the user could swap the device’s SIM card with one from another provider while your app is running. This class also gives you access to the object, which contains information about the user’s home cellular service provider.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTTelephonyNetworkInfo
type TelephonyNetworkInfo struct {
	objectivec.Object
}

// TelephonyNetworkInfoFrom constructs a [TelephonyNetworkInfo] from an unsafe.Pointer.
//
// An object that provides notifications of changes to the user’s cellular service provider.
func TelephonyNetworkInfoFrom(ptr unsafe.Pointer) TelephonyNetworkInfo {
	return TelephonyNetworkInfo{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TelephonyNetworkInfoClass) Alloc() TelephonyNetworkInfo {
	rv := objc.Send[TelephonyNetworkInfo](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TelephonyNetworkInfoClass) New() TelephonyNetworkInfo {
	rv := objc.Send[TelephonyNetworkInfo](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TelephonyNetworkInfo) Init() TelephonyNetworkInfo {
	rv := objc.Send[TelephonyNetworkInfo](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TelephonyNetworkInfo) Autorelease() TelephonyNetworkInfo {
	rv := objc.Send[TelephonyNetworkInfo](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTelephonyNetworkInfo creates a new TelephonyNetworkInfo instance.
func NewTelephonyNetworkInfo() TelephonyNetworkInfo {
	return getTelephonyNetworkInfoClass().New()
}


// The identifier of the service that’s currently providing data.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTTelephonyNetworkInfo/dataServiceIdentifier
func (t_ TelephonyNetworkInfo) DataServiceIdentifier() string {
	rv := objc.Send[string](t_.ID, objc.Sel("dataServiceIdentifier"))
	return rv
}

// A dictionary containing the current radio access technology registered to each service.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTTelephonyNetworkInfo/serviceCurrentRadioAccessTechnology
func (t_ TelephonyNetworkInfo) ServiceCurrentRadioAccessTechnology() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("serviceCurrentRadioAccessTechnology"))
	return rv
}

// The current radio access technology registered with the device.
//
// [Full Topic]: https://developer.apple.com/documentation/coretelephony/cttelephonynetworkinfo/currentradioaccesstechnology
func (t_ TelephonyNetworkInfo) CurrentRadioAccessTechnology() string {
	rv := objc.Send[string](t_.ID, objc.Sel("currentRadioAccessTechnology"))
	return rv
}


// SetCurrentRadioAccessTechnology sets the value of the currentRadioAccessTechnology property.
// The current radio access technology registered with the device.

//
// [Full Topic]: https://developer.apple.com/documentation/coretelephony/cttelephonynetworkinfo/currentradioaccesstechnology
func (t_ TelephonyNetworkInfo) SetCurrentRadioAccessTechnology(value string) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setCurrentRadioAccessTechnology:"), objc.String(value))
}

// An object that the system notifies when the data service identifier changes.
//
// [Full Topic]: https://developer.apple.com/documentation/coretelephony/cttelephonynetworkinfo/delegate
func (t_ TelephonyNetworkInfo) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// An object that the system notifies when the data service identifier changes.

//
// [Full Topic]: https://developer.apple.com/documentation/coretelephony/cttelephonynetworkinfo/delegate
func (t_ TelephonyNetworkInfo) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDelegate:"), value)
}

// A dictionary that contains carrier information about each service.
//
// [Full Topic]: https://developer.apple.com/documentation/coretelephony/cttelephonynetworkinfo/servicesubscribercellularproviders
func (t_ TelephonyNetworkInfo) ServiceSubscriberCellularProviders() string {
	rv := objc.Send[string](t_.ID, objc.Sel("serviceSubscriberCellularProviders"))
	return rv
}


// SetServiceSubscriberCellularProviders sets the value of the serviceSubscriberCellularProviders property.
// A dictionary that contains carrier information about each service.

//
// [Full Topic]: https://developer.apple.com/documentation/coretelephony/cttelephonynetworkinfo/servicesubscribercellularproviders
func (t_ TelephonyNetworkInfo) SetServiceSubscriberCellularProviders(value string) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setServiceSubscriberCellularProviders:"), objc.String(value))
}

// A block dispatched when there are updates to the user’s cellular provider information for any service.
//
// [Full Topic]: https://developer.apple.com/documentation/coretelephony/cttelephonynetworkinfo/servicesubscribercellularprovidersdidupdatenotifier
func (t_ TelephonyNetworkInfo) ServiceSubscriberCellularProvidersDidUpdateNotifier() string {
	rv := objc.Send[string](t_.ID, objc.Sel("serviceSubscriberCellularProvidersDidUpdateNotifier"))
	return rv
}


// SetServiceSubscriberCellularProvidersDidUpdateNotifier sets the value of the serviceSubscriberCellularProvidersDidUpdateNotifier property.
// A block dispatched when there are updates to the user’s cellular provider information for any service.

//
// [Full Topic]: https://developer.apple.com/documentation/coretelephony/cttelephonynetworkinfo/servicesubscribercellularprovidersdidupdatenotifier
func (t_ TelephonyNetworkInfo) SetServiceSubscriberCellularProvidersDidUpdateNotifier(value string) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setServiceSubscriberCellularProvidersDidUpdateNotifier:"), objc.String(value))
}

// Information about the user’s cellular service provider.
//
// [Full Topic]: https://developer.apple.com/documentation/coretelephony/cttelephonynetworkinfo/subscribercellularprovider
func (t_ TelephonyNetworkInfo) SubscriberCellularProvider() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("subscriberCellularProvider"))
	return rv
}


// SetSubscriberCellularProvider sets the value of the subscriberCellularProvider property.
// Information about the user’s cellular service provider.

//
// [Full Topic]: https://developer.apple.com/documentation/coretelephony/cttelephonynetworkinfo/subscribercellularprovider
func (t_ TelephonyNetworkInfo) SetSubscriberCellularProvider(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSubscriberCellularProvider:"), value)
}

// A block dispatched when the user’s cellular service provider information changes.
//
// [Full Topic]: https://developer.apple.com/documentation/coretelephony/cttelephonynetworkinfo/subscribercellularproviderdidupdatenotifier
func (t_ TelephonyNetworkInfo) SubscriberCellularProviderDidUpdateNotifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("subscriberCellularProviderDidUpdateNotifier"))
	return rv
}


// SetSubscriberCellularProviderDidUpdateNotifier sets the value of the subscriberCellularProviderDidUpdateNotifier property.
// A block dispatched when the user’s cellular service provider information changes.

//
// [Full Topic]: https://developer.apple.com/documentation/coretelephony/cttelephonynetworkinfo/subscribercellularproviderdidupdatenotifier
func (t_ TelephonyNetworkInfo) SetSubscriberCellularProviderDidUpdateNotifier(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSubscriberCellularProviderDidUpdateNotifier:"), value)
}




