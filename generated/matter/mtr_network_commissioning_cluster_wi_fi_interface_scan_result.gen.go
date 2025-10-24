// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
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
	// properties:
	Bssid() objc.IObject /* cross-framework: Data */
	SetBssid(value objc.IObject /* cross-framework: Data */)
	Channel() objc.IObject /* cross-framework: NSNumber */
	SetChannel(value objc.IObject /* cross-framework: NSNumber */)
	Rssi() objc.IObject /* cross-framework: NSNumber */
	SetRssi(value objc.IObject /* cross-framework: NSNumber */)
	Security() objc.IObject /* cross-framework: NSNumber */
	SetSecurity(value objc.IObject /* cross-framework: NSNumber */)
	Ssid() objc.IObject /* cross-framework: Data */
	SetSsid(value objc.IObject /* cross-framework: Data */)
	WiFiBand() objc.IObject /* cross-framework: NSNumber */
	SetWiFiBand(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}

// [Full Topic]
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

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterwifiinterfacescanresult/bssid
func (m_ MTRNetworkCommissioningClusterWiFiInterfaceScanResult) Bssid() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("bssid"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterwifiinterfacescanresult/bssid
func (m_ MTRNetworkCommissioningClusterWiFiInterfaceScanResult) SetBssid(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBssid:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterwifiinterfacescanresult/channel
func (m_ MTRNetworkCommissioningClusterWiFiInterfaceScanResult) Channel() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("channel"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterwifiinterfacescanresult/channel
func (m_ MTRNetworkCommissioningClusterWiFiInterfaceScanResult) SetChannel(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setChannel:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterwifiinterfacescanresult/rssi
func (m_ MTRNetworkCommissioningClusterWiFiInterfaceScanResult) Rssi() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("rssi"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterwifiinterfacescanresult/rssi
func (m_ MTRNetworkCommissioningClusterWiFiInterfaceScanResult) SetRssi(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRssi:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterwifiinterfacescanresult/security
func (m_ MTRNetworkCommissioningClusterWiFiInterfaceScanResult) Security() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("security"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterwifiinterfacescanresult/security
func (m_ MTRNetworkCommissioningClusterWiFiInterfaceScanResult) SetSecurity(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSecurity:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterwifiinterfacescanresult/ssid
func (m_ MTRNetworkCommissioningClusterWiFiInterfaceScanResult) Ssid() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("ssid"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterwifiinterfacescanresult/ssid
func (m_ MTRNetworkCommissioningClusterWiFiInterfaceScanResult) SetSsid(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSsid:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterwifiinterfacescanresult/wifiband
func (m_ MTRNetworkCommissioningClusterWiFiInterfaceScanResult) WiFiBand() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("wiFiBand"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterwifiinterfacescanresult/wifiband
func (m_ MTRNetworkCommissioningClusterWiFiInterfaceScanResult) SetWiFiBand(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setWiFiBand:"), value)
}
