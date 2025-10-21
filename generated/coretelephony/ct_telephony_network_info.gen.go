// Code generated from Apple documentation for CoreTelephony. DO NOT EDIT.

package coretelephony

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
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




