// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRNetworkCommissioningClusterThreadInterfaceScanResult] class.
var (
	MTRNetworkCommissioningClusterThreadInterfaceScanResultClass     _MTRNetworkCommissioningClusterThreadInterfaceScanResultClass
	MTRNetworkCommissioningClusterThreadInterfaceScanResultClassOnce sync.Once
)

func getMTRNetworkCommissioningClusterThreadInterfaceScanResultClass() _MTRNetworkCommissioningClusterThreadInterfaceScanResultClass {
	MTRNetworkCommissioningClusterThreadInterfaceScanResultClassOnce.Do(func() {
		MTRNetworkCommissioningClusterThreadInterfaceScanResultClass = _MTRNetworkCommissioningClusterThreadInterfaceScanResultClass{objc.GetClass("MTRNetworkCommissioningClusterThreadInterfaceScanResult")}
	})
	return MTRNetworkCommissioningClusterThreadInterfaceScanResultClass
}

type _MTRNetworkCommissioningClusterThreadInterfaceScanResultClass struct {
	class objc.Class
}

// An interface definition for the [MTRNetworkCommissioningClusterThreadInterfaceScanResult] class.
type IMTRNetworkCommissioningClusterThreadInterfaceScanResult interface {
	IMTRNetworkCommissioningClusterThreadInterfaceScanResultStruct
	// properties:
	Channel() objc.IObject /* cross-framework: NSNumber */
	SetChannel(value objc.IObject /* cross-framework: NSNumber */)
	ExtendedAddress() objc.IObject /* cross-framework: Data */
	SetExtendedAddress(value objc.IObject /* cross-framework: Data */)
	ExtendedPanId() objc.IObject /* cross-framework: NSNumber */
	SetExtendedPanId(value objc.IObject /* cross-framework: NSNumber */)
	Lqi() objc.IObject /* cross-framework: NSNumber */
	SetLqi(value objc.IObject /* cross-framework: NSNumber */)
	NetworkName() objc.IObject /* cross-framework: NSString */
	SetNetworkName(value objc.IObject /* cross-framework: NSString */)
	PanId() objc.IObject /* cross-framework: NSNumber */
	SetPanId(value objc.IObject /* cross-framework: NSNumber */)
	Rssi() objc.IObject /* cross-framework: NSNumber */
	SetRssi(value objc.IObject /* cross-framework: NSNumber */)
	Version() objc.IObject /* cross-framework: NSNumber */
	SetVersion(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRNetworkCommissioningClusterThreadInterfaceScanResult
type MTRNetworkCommissioningClusterThreadInterfaceScanResult struct {
	MTRNetworkCommissioningClusterThreadInterfaceScanResultStruct
}

// MTRNetworkCommissioningClusterThreadInterfaceScanResultFrom constructs a [MTRNetworkCommissioningClusterThreadInterfaceScanResult] from an unsafe.Pointer.
func MTRNetworkCommissioningClusterThreadInterfaceScanResultFrom(ptr unsafe.Pointer) MTRNetworkCommissioningClusterThreadInterfaceScanResult {
	return MTRNetworkCommissioningClusterThreadInterfaceScanResult{
		MTRNetworkCommissioningClusterThreadInterfaceScanResultStruct: MTRNetworkCommissioningClusterThreadInterfaceScanResultStructFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRNetworkCommissioningClusterThreadInterfaceScanResultClass) Alloc() MTRNetworkCommissioningClusterThreadInterfaceScanResult {
	rv := objc.Send[MTRNetworkCommissioningClusterThreadInterfaceScanResult](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRNetworkCommissioningClusterThreadInterfaceScanResultClass) New() MTRNetworkCommissioningClusterThreadInterfaceScanResult {
	rv := objc.Send[MTRNetworkCommissioningClusterThreadInterfaceScanResult](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRNetworkCommissioningClusterThreadInterfaceScanResult) Init() MTRNetworkCommissioningClusterThreadInterfaceScanResult {
	rv := objc.Send[MTRNetworkCommissioningClusterThreadInterfaceScanResult](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRNetworkCommissioningClusterThreadInterfaceScanResult) Autorelease() MTRNetworkCommissioningClusterThreadInterfaceScanResult {
	rv := objc.Send[MTRNetworkCommissioningClusterThreadInterfaceScanResult](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRNetworkCommissioningClusterThreadInterfaceScanResult creates a new MTRNetworkCommissioningClusterThreadInterfaceScanResult instance.
func NewMTRNetworkCommissioningClusterThreadInterfaceScanResult() MTRNetworkCommissioningClusterThreadInterfaceScanResult {
	return getMTRNetworkCommissioningClusterThreadInterfaceScanResultClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterthreadinterfacescanresult/channel
func (m_ MTRNetworkCommissioningClusterThreadInterfaceScanResult) Channel() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("channel"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterthreadinterfacescanresult/channel
func (m_ MTRNetworkCommissioningClusterThreadInterfaceScanResult) SetChannel(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setChannel:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterthreadinterfacescanresult/extendedaddress
func (m_ MTRNetworkCommissioningClusterThreadInterfaceScanResult) ExtendedAddress() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("extendedAddress"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterthreadinterfacescanresult/extendedaddress
func (m_ MTRNetworkCommissioningClusterThreadInterfaceScanResult) SetExtendedAddress(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setExtendedAddress:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterthreadinterfacescanresult/extendedpanid
func (m_ MTRNetworkCommissioningClusterThreadInterfaceScanResult) ExtendedPanId() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("extendedPanId"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterthreadinterfacescanresult/extendedpanid
func (m_ MTRNetworkCommissioningClusterThreadInterfaceScanResult) SetExtendedPanId(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setExtendedPanId:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterthreadinterfacescanresult/lqi
func (m_ MTRNetworkCommissioningClusterThreadInterfaceScanResult) Lqi() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("lqi"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterthreadinterfacescanresult/lqi
func (m_ MTRNetworkCommissioningClusterThreadInterfaceScanResult) SetLqi(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLqi:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterthreadinterfacescanresult/networkname
func (m_ MTRNetworkCommissioningClusterThreadInterfaceScanResult) NetworkName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("networkName"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterthreadinterfacescanresult/networkname
func (m_ MTRNetworkCommissioningClusterThreadInterfaceScanResult) SetNetworkName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNetworkName:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterthreadinterfacescanresult/panid
func (m_ MTRNetworkCommissioningClusterThreadInterfaceScanResult) PanId() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("panId"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterthreadinterfacescanresult/panid
func (m_ MTRNetworkCommissioningClusterThreadInterfaceScanResult) SetPanId(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPanId:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterthreadinterfacescanresult/rssi
func (m_ MTRNetworkCommissioningClusterThreadInterfaceScanResult) Rssi() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("rssi"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterthreadinterfacescanresult/rssi
func (m_ MTRNetworkCommissioningClusterThreadInterfaceScanResult) SetRssi(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRssi:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterthreadinterfacescanresult/version
func (m_ MTRNetworkCommissioningClusterThreadInterfaceScanResult) Version() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("version"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterthreadinterfacescanresult/version
func (m_ MTRNetworkCommissioningClusterThreadInterfaceScanResult) SetVersion(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVersion:"), value)
}



