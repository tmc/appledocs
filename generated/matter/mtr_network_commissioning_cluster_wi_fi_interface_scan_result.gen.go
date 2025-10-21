// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRNetworkCommissioningClusterWiFiInterfaceScanResult] class.
var (
	MTRNetworkCommissioningClusterWiFiInterfaceScanResultClass     _MTRNetworkCommissioningClusterWiFiInterfaceScanResultClass
	MTRNetworkCommissioningClusterWiFiInterfaceScanResultClassOnce sync.Once
)

func getMTRNetworkCommissioningClusterWiFiInterfaceScanResultClass() _MTRNetworkCommissioningClusterWiFiInterfaceScanResultClass {
	MTRNetworkCommissioningClusterWiFiInterfaceScanResultClassOnce.Do(func() {
		MTRNetworkCommissioningClusterWiFiInterfaceScanResultClass = _MTRNetworkCommissioningClusterWiFiInterfaceScanResultClass{objc.GetClass("MTRNetworkCommissioningClusterWiFiInterfaceScanResult")}
	})
	return MTRNetworkCommissioningClusterWiFiInterfaceScanResultClass
}

type _MTRNetworkCommissioningClusterWiFiInterfaceScanResultClass struct {
	class objc.Class
}

// An interface definition for the [MTRNetworkCommissioningClusterWiFiInterfaceScanResult] class.
type IMTRNetworkCommissioningClusterWiFiInterfaceScanResult interface {
	IMTRNetworkCommissioningClusterWiFiInterfaceScanResultStruct
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRNetworkCommissioningClusterWiFiInterfaceScanResult
type MTRNetworkCommissioningClusterWiFiInterfaceScanResult struct {
	MTRNetworkCommissioningClusterWiFiInterfaceScanResultStruct
}

// MTRNetworkCommissioningClusterWiFiInterfaceScanResultFrom constructs a [MTRNetworkCommissioningClusterWiFiInterfaceScanResult] from an unsafe.Pointer.
func MTRNetworkCommissioningClusterWiFiInterfaceScanResultFrom(ptr unsafe.Pointer) MTRNetworkCommissioningClusterWiFiInterfaceScanResult {
	return MTRNetworkCommissioningClusterWiFiInterfaceScanResult{
		MTRNetworkCommissioningClusterWiFiInterfaceScanResultStruct: MTRNetworkCommissioningClusterWiFiInterfaceScanResultStructFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRNetworkCommissioningClusterWiFiInterfaceScanResultClass) Alloc() MTRNetworkCommissioningClusterWiFiInterfaceScanResult {
	rv := objc.Send[MTRNetworkCommissioningClusterWiFiInterfaceScanResult](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRNetworkCommissioningClusterWiFiInterfaceScanResultClass) New() MTRNetworkCommissioningClusterWiFiInterfaceScanResult {
	rv := objc.Send[MTRNetworkCommissioningClusterWiFiInterfaceScanResult](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRNetworkCommissioningClusterWiFiInterfaceScanResult) Init() MTRNetworkCommissioningClusterWiFiInterfaceScanResult {
	rv := objc.Send[MTRNetworkCommissioningClusterWiFiInterfaceScanResult](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRNetworkCommissioningClusterWiFiInterfaceScanResult) Autorelease() MTRNetworkCommissioningClusterWiFiInterfaceScanResult {
	rv := objc.Send[MTRNetworkCommissioningClusterWiFiInterfaceScanResult](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRNetworkCommissioningClusterWiFiInterfaceScanResult creates a new MTRNetworkCommissioningClusterWiFiInterfaceScanResult instance.
func NewMTRNetworkCommissioningClusterWiFiInterfaceScanResult() MTRNetworkCommissioningClusterWiFiInterfaceScanResult {
	return getMTRNetworkCommissioningClusterWiFiInterfaceScanResultClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterwifiinterfacescanresult/security
func (m_ MTRNetworkCommissioningClusterWiFiInterfaceScanResult) Security() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("security"))
	return rv
}


// SetSecurity sets the value of the security property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterwifiinterfacescanresult/security
func (m_ MTRNetworkCommissioningClusterWiFiInterfaceScanResult) SetSecurity(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSecurity:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterwifiinterfacescanresult/rssi
func (m_ MTRNetworkCommissioningClusterWiFiInterfaceScanResult) Rssi() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("rssi"))
	return rv
}


// SetRssi sets the value of the rssi property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterwifiinterfacescanresult/rssi
func (m_ MTRNetworkCommissioningClusterWiFiInterfaceScanResult) SetRssi(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRssi:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterwifiinterfacescanresult/wifiband
func (m_ MTRNetworkCommissioningClusterWiFiInterfaceScanResult) WiFiBand() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("wiFiBand"))
	return rv
}


// SetWiFiBand sets the value of the wiFiBand property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterwifiinterfacescanresult/wifiband
func (m_ MTRNetworkCommissioningClusterWiFiInterfaceScanResult) SetWiFiBand(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setWiFiBand:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterwifiinterfacescanresult/channel
func (m_ MTRNetworkCommissioningClusterWiFiInterfaceScanResult) Channel() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("channel"))
	return rv
}


// SetChannel sets the value of the channel property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterwifiinterfacescanresult/channel
func (m_ MTRNetworkCommissioningClusterWiFiInterfaceScanResult) SetChannel(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setChannel:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterwifiinterfacescanresult/ssid
func (m_ MTRNetworkCommissioningClusterWiFiInterfaceScanResult) Ssid() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("ssid"))
	return rv
}


// SetSsid sets the value of the ssid property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterwifiinterfacescanresult/ssid
func (m_ MTRNetworkCommissioningClusterWiFiInterfaceScanResult) SetSsid(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSsid:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterwifiinterfacescanresult/bssid
func (m_ MTRNetworkCommissioningClusterWiFiInterfaceScanResult) Bssid() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("bssid"))
	return rv
}


// SetBssid sets the value of the bssid property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterwifiinterfacescanresult/bssid
func (m_ MTRNetworkCommissioningClusterWiFiInterfaceScanResult) SetBssid(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBssid:"), value)
}



