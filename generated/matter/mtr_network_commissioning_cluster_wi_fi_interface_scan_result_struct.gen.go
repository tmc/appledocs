// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRNetworkCommissioningClusterWiFiInterfaceScanResultStruct] class.
var (
	MTRNetworkCommissioningClusterWiFiInterfaceScanResultStructClass     _MTRNetworkCommissioningClusterWiFiInterfaceScanResultStructClass
	MTRNetworkCommissioningClusterWiFiInterfaceScanResultStructClassOnce sync.Once
)

func getMTRNetworkCommissioningClusterWiFiInterfaceScanResultStructClass() _MTRNetworkCommissioningClusterWiFiInterfaceScanResultStructClass {
	MTRNetworkCommissioningClusterWiFiInterfaceScanResultStructClassOnce.Do(func() {
		MTRNetworkCommissioningClusterWiFiInterfaceScanResultStructClass = _MTRNetworkCommissioningClusterWiFiInterfaceScanResultStructClass{objc.GetClass("MTRNetworkCommissioningClusterWiFiInterfaceScanResultStruct")}
	})
	return MTRNetworkCommissioningClusterWiFiInterfaceScanResultStructClass
}

type _MTRNetworkCommissioningClusterWiFiInterfaceScanResultStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRNetworkCommissioningClusterWiFiInterfaceScanResultStruct] class.
type IMTRNetworkCommissioningClusterWiFiInterfaceScanResultStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRNetworkCommissioningClusterWiFiInterfaceScanResultStruct
type MTRNetworkCommissioningClusterWiFiInterfaceScanResultStruct struct {
	objectivec.Object
}

// MTRNetworkCommissioningClusterWiFiInterfaceScanResultStructFrom constructs a [MTRNetworkCommissioningClusterWiFiInterfaceScanResultStruct] from an unsafe.Pointer.
func MTRNetworkCommissioningClusterWiFiInterfaceScanResultStructFrom(ptr unsafe.Pointer) MTRNetworkCommissioningClusterWiFiInterfaceScanResultStruct {
	return MTRNetworkCommissioningClusterWiFiInterfaceScanResultStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRNetworkCommissioningClusterWiFiInterfaceScanResultStructClass) Alloc() MTRNetworkCommissioningClusterWiFiInterfaceScanResultStruct {
	rv := objc.Send[MTRNetworkCommissioningClusterWiFiInterfaceScanResultStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRNetworkCommissioningClusterWiFiInterfaceScanResultStructClass) New() MTRNetworkCommissioningClusterWiFiInterfaceScanResultStruct {
	rv := objc.Send[MTRNetworkCommissioningClusterWiFiInterfaceScanResultStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRNetworkCommissioningClusterWiFiInterfaceScanResultStruct) Init() MTRNetworkCommissioningClusterWiFiInterfaceScanResultStruct {
	rv := objc.Send[MTRNetworkCommissioningClusterWiFiInterfaceScanResultStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRNetworkCommissioningClusterWiFiInterfaceScanResultStruct) Autorelease() MTRNetworkCommissioningClusterWiFiInterfaceScanResultStruct {
	rv := objc.Send[MTRNetworkCommissioningClusterWiFiInterfaceScanResultStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRNetworkCommissioningClusterWiFiInterfaceScanResultStruct creates a new MTRNetworkCommissioningClusterWiFiInterfaceScanResultStruct instance.
func NewMTRNetworkCommissioningClusterWiFiInterfaceScanResultStruct() MTRNetworkCommissioningClusterWiFiInterfaceScanResultStruct {
	return getMTRNetworkCommissioningClusterWiFiInterfaceScanResultStructClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterwifiinterfacescanresultstruct/bssid
func (m_ MTRNetworkCommissioningClusterWiFiInterfaceScanResultStruct) Bssid() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("bssid"))
	return rv
}


// SetBssid sets the value of the bssid property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterwifiinterfacescanresultstruct/bssid
func (m_ MTRNetworkCommissioningClusterWiFiInterfaceScanResultStruct) SetBssid(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBssid:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterwifiinterfacescanresultstruct/channel
func (m_ MTRNetworkCommissioningClusterWiFiInterfaceScanResultStruct) Channel() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("channel"))
	return rv
}


// SetChannel sets the value of the channel property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterwifiinterfacescanresultstruct/channel
func (m_ MTRNetworkCommissioningClusterWiFiInterfaceScanResultStruct) SetChannel(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setChannel:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterwifiinterfacescanresultstruct/rssi
func (m_ MTRNetworkCommissioningClusterWiFiInterfaceScanResultStruct) Rssi() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("rssi"))
	return rv
}


// SetRssi sets the value of the rssi property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterwifiinterfacescanresultstruct/rssi
func (m_ MTRNetworkCommissioningClusterWiFiInterfaceScanResultStruct) SetRssi(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRssi:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterwifiinterfacescanresultstruct/security
func (m_ MTRNetworkCommissioningClusterWiFiInterfaceScanResultStruct) Security() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("security"))
	return rv
}


// SetSecurity sets the value of the security property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterwifiinterfacescanresultstruct/security
func (m_ MTRNetworkCommissioningClusterWiFiInterfaceScanResultStruct) SetSecurity(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSecurity:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterwifiinterfacescanresultstruct/ssid
func (m_ MTRNetworkCommissioningClusterWiFiInterfaceScanResultStruct) Ssid() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("ssid"))
	return rv
}


// SetSsid sets the value of the ssid property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterwifiinterfacescanresultstruct/ssid
func (m_ MTRNetworkCommissioningClusterWiFiInterfaceScanResultStruct) SetSsid(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSsid:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterwifiinterfacescanresultstruct/wifiband
func (m_ MTRNetworkCommissioningClusterWiFiInterfaceScanResultStruct) WiFiBand() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("wiFiBand"))
	return rv
}


// SetWiFiBand sets the value of the wiFiBand property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterwifiinterfacescanresultstruct/wifiband
func (m_ MTRNetworkCommissioningClusterWiFiInterfaceScanResultStruct) SetWiFiBand(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setWiFiBand:"), value)
}



