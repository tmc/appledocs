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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterwifiinterfacescanresultstruct/bssid
func (m_ MTRNetworkCommissioningClusterWiFiInterfaceScanResultStruct) Bssid() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("bssid"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterwifiinterfacescanresultstruct/bssid
func (m_ MTRNetworkCommissioningClusterWiFiInterfaceScanResultStruct) SetBssid(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBssid:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterwifiinterfacescanresultstruct/channel
func (m_ MTRNetworkCommissioningClusterWiFiInterfaceScanResultStruct) Channel() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("channel"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterwifiinterfacescanresultstruct/channel
func (m_ MTRNetworkCommissioningClusterWiFiInterfaceScanResultStruct) SetChannel(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setChannel:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterwifiinterfacescanresultstruct/rssi
func (m_ MTRNetworkCommissioningClusterWiFiInterfaceScanResultStruct) Rssi() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("rssi"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterwifiinterfacescanresultstruct/rssi
func (m_ MTRNetworkCommissioningClusterWiFiInterfaceScanResultStruct) SetRssi(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRssi:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterwifiinterfacescanresultstruct/security
func (m_ MTRNetworkCommissioningClusterWiFiInterfaceScanResultStruct) Security() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("security"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterwifiinterfacescanresultstruct/security
func (m_ MTRNetworkCommissioningClusterWiFiInterfaceScanResultStruct) SetSecurity(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSecurity:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterwifiinterfacescanresultstruct/ssid
func (m_ MTRNetworkCommissioningClusterWiFiInterfaceScanResultStruct) Ssid() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("ssid"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterwifiinterfacescanresultstruct/ssid
func (m_ MTRNetworkCommissioningClusterWiFiInterfaceScanResultStruct) SetSsid(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSsid:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterwifiinterfacescanresultstruct/wifiband
func (m_ MTRNetworkCommissioningClusterWiFiInterfaceScanResultStruct) WiFiBand() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("wiFiBand"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterwifiinterfacescanresultstruct/wifiband
func (m_ MTRNetworkCommissioningClusterWiFiInterfaceScanResultStruct) SetWiFiBand(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setWiFiBand:"), value)
}



