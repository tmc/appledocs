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
}

//
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


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterthreadinterfacescanresult/lqi
func (m_ MTRNetworkCommissioningClusterThreadInterfaceScanResult) Lqi() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("lqi"))
	return rv
}


// SetLqi sets the value of the lqi property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterthreadinterfacescanresult/lqi
func (m_ MTRNetworkCommissioningClusterThreadInterfaceScanResult) SetLqi(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLqi:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterthreadinterfacescanresult/extendedaddress
func (m_ MTRNetworkCommissioningClusterThreadInterfaceScanResult) ExtendedAddress() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("extendedAddress"))
	return rv
}


// SetExtendedAddress sets the value of the extendedAddress property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterthreadinterfacescanresult/extendedaddress
func (m_ MTRNetworkCommissioningClusterThreadInterfaceScanResult) SetExtendedAddress(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setExtendedAddress:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterthreadinterfacescanresult/channel
func (m_ MTRNetworkCommissioningClusterThreadInterfaceScanResult) Channel() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("channel"))
	return rv
}


// SetChannel sets the value of the channel property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterthreadinterfacescanresult/channel
func (m_ MTRNetworkCommissioningClusterThreadInterfaceScanResult) SetChannel(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setChannel:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterthreadinterfacescanresult/rssi
func (m_ MTRNetworkCommissioningClusterThreadInterfaceScanResult) Rssi() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("rssi"))
	return rv
}


// SetRssi sets the value of the rssi property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterthreadinterfacescanresult/rssi
func (m_ MTRNetworkCommissioningClusterThreadInterfaceScanResult) SetRssi(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRssi:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterthreadinterfacescanresult/extendedpanid
func (m_ MTRNetworkCommissioningClusterThreadInterfaceScanResult) ExtendedPanId() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("extendedPanId"))
	return rv
}


// SetExtendedPanId sets the value of the extendedPanId property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterthreadinterfacescanresult/extendedpanid
func (m_ MTRNetworkCommissioningClusterThreadInterfaceScanResult) SetExtendedPanId(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setExtendedPanId:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterthreadinterfacescanresult/version
func (m_ MTRNetworkCommissioningClusterThreadInterfaceScanResult) Version() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("version"))
	return rv
}


// SetVersion sets the value of the version property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterthreadinterfacescanresult/version
func (m_ MTRNetworkCommissioningClusterThreadInterfaceScanResult) SetVersion(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVersion:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterthreadinterfacescanresult/panid
func (m_ MTRNetworkCommissioningClusterThreadInterfaceScanResult) PanId() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("panId"))
	return rv
}


// SetPanId sets the value of the panId property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterthreadinterfacescanresult/panid
func (m_ MTRNetworkCommissioningClusterThreadInterfaceScanResult) SetPanId(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPanId:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterthreadinterfacescanresult/networkname
func (m_ MTRNetworkCommissioningClusterThreadInterfaceScanResult) NetworkName() string {
	rv := objc.Send[string](m_.ID, objc.Sel("networkName"))
	return rv
}


// SetNetworkName sets the value of the networkName property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterthreadinterfacescanresult/networkname
func (m_ MTRNetworkCommissioningClusterThreadInterfaceScanResult) SetNetworkName(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNetworkName:"), objc.String(value))
}



