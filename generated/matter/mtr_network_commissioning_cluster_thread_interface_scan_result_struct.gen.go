// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRNetworkCommissioningClusterThreadInterfaceScanResultStruct] class.
var (
	MTRNetworkCommissioningClusterThreadInterfaceScanResultStructClass     _MTRNetworkCommissioningClusterThreadInterfaceScanResultStructClass
	MTRNetworkCommissioningClusterThreadInterfaceScanResultStructClassOnce sync.Once
)

func getMTRNetworkCommissioningClusterThreadInterfaceScanResultStructClass() _MTRNetworkCommissioningClusterThreadInterfaceScanResultStructClass {
	MTRNetworkCommissioningClusterThreadInterfaceScanResultStructClassOnce.Do(func() {
		MTRNetworkCommissioningClusterThreadInterfaceScanResultStructClass = _MTRNetworkCommissioningClusterThreadInterfaceScanResultStructClass{objc.GetClass("MTRNetworkCommissioningClusterThreadInterfaceScanResultStruct")}
	})
	return MTRNetworkCommissioningClusterThreadInterfaceScanResultStructClass
}

type _MTRNetworkCommissioningClusterThreadInterfaceScanResultStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRNetworkCommissioningClusterThreadInterfaceScanResultStruct] class.
type IMTRNetworkCommissioningClusterThreadInterfaceScanResultStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRNetworkCommissioningClusterThreadInterfaceScanResultStruct
type MTRNetworkCommissioningClusterThreadInterfaceScanResultStruct struct {
	objectivec.Object
}

// MTRNetworkCommissioningClusterThreadInterfaceScanResultStructFrom constructs a [MTRNetworkCommissioningClusterThreadInterfaceScanResultStruct] from an unsafe.Pointer.
func MTRNetworkCommissioningClusterThreadInterfaceScanResultStructFrom(ptr unsafe.Pointer) MTRNetworkCommissioningClusterThreadInterfaceScanResultStruct {
	return MTRNetworkCommissioningClusterThreadInterfaceScanResultStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRNetworkCommissioningClusterThreadInterfaceScanResultStructClass) Alloc() MTRNetworkCommissioningClusterThreadInterfaceScanResultStruct {
	rv := objc.Send[MTRNetworkCommissioningClusterThreadInterfaceScanResultStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRNetworkCommissioningClusterThreadInterfaceScanResultStructClass) New() MTRNetworkCommissioningClusterThreadInterfaceScanResultStruct {
	rv := objc.Send[MTRNetworkCommissioningClusterThreadInterfaceScanResultStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRNetworkCommissioningClusterThreadInterfaceScanResultStruct) Init() MTRNetworkCommissioningClusterThreadInterfaceScanResultStruct {
	rv := objc.Send[MTRNetworkCommissioningClusterThreadInterfaceScanResultStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRNetworkCommissioningClusterThreadInterfaceScanResultStruct) Autorelease() MTRNetworkCommissioningClusterThreadInterfaceScanResultStruct {
	rv := objc.Send[MTRNetworkCommissioningClusterThreadInterfaceScanResultStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRNetworkCommissioningClusterThreadInterfaceScanResultStruct creates a new MTRNetworkCommissioningClusterThreadInterfaceScanResultStruct instance.
func NewMTRNetworkCommissioningClusterThreadInterfaceScanResultStruct() MTRNetworkCommissioningClusterThreadInterfaceScanResultStruct {
	return getMTRNetworkCommissioningClusterThreadInterfaceScanResultStructClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterthreadinterfacescanresultstruct/channel
func (m_ MTRNetworkCommissioningClusterThreadInterfaceScanResultStruct) Channel() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("channel"))
	return rv
}


// SetChannel sets the value of the channel property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterthreadinterfacescanresultstruct/channel
func (m_ MTRNetworkCommissioningClusterThreadInterfaceScanResultStruct) SetChannel(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setChannel:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterthreadinterfacescanresultstruct/extendedaddress
func (m_ MTRNetworkCommissioningClusterThreadInterfaceScanResultStruct) ExtendedAddress() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("extendedAddress"))
	return rv
}


// SetExtendedAddress sets the value of the extendedAddress property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterthreadinterfacescanresultstruct/extendedaddress
func (m_ MTRNetworkCommissioningClusterThreadInterfaceScanResultStruct) SetExtendedAddress(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setExtendedAddress:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterthreadinterfacescanresultstruct/extendedpanid
func (m_ MTRNetworkCommissioningClusterThreadInterfaceScanResultStruct) ExtendedPanId() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("extendedPanId"))
	return rv
}


// SetExtendedPanId sets the value of the extendedPanId property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterthreadinterfacescanresultstruct/extendedpanid
func (m_ MTRNetworkCommissioningClusterThreadInterfaceScanResultStruct) SetExtendedPanId(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setExtendedPanId:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterthreadinterfacescanresultstruct/lqi
func (m_ MTRNetworkCommissioningClusterThreadInterfaceScanResultStruct) Lqi() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("lqi"))
	return rv
}


// SetLqi sets the value of the lqi property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterthreadinterfacescanresultstruct/lqi
func (m_ MTRNetworkCommissioningClusterThreadInterfaceScanResultStruct) SetLqi(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLqi:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterthreadinterfacescanresultstruct/networkname
func (m_ MTRNetworkCommissioningClusterThreadInterfaceScanResultStruct) NetworkName() string {
	rv := objc.Send[string](m_.ID, objc.Sel("networkName"))
	return rv
}


// SetNetworkName sets the value of the networkName property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterthreadinterfacescanresultstruct/networkname
func (m_ MTRNetworkCommissioningClusterThreadInterfaceScanResultStruct) SetNetworkName(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNetworkName:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterthreadinterfacescanresultstruct/panid
func (m_ MTRNetworkCommissioningClusterThreadInterfaceScanResultStruct) PanId() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("panId"))
	return rv
}


// SetPanId sets the value of the panId property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterthreadinterfacescanresultstruct/panid
func (m_ MTRNetworkCommissioningClusterThreadInterfaceScanResultStruct) SetPanId(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPanId:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterthreadinterfacescanresultstruct/rssi
func (m_ MTRNetworkCommissioningClusterThreadInterfaceScanResultStruct) Rssi() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("rssi"))
	return rv
}


// SetRssi sets the value of the rssi property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterthreadinterfacescanresultstruct/rssi
func (m_ MTRNetworkCommissioningClusterThreadInterfaceScanResultStruct) SetRssi(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRssi:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterthreadinterfacescanresultstruct/version
func (m_ MTRNetworkCommissioningClusterThreadInterfaceScanResultStruct) Version() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("version"))
	return rv
}


// SetVersion sets the value of the version property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterthreadinterfacescanresultstruct/version
func (m_ MTRNetworkCommissioningClusterThreadInterfaceScanResultStruct) SetVersion(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVersion:"), value)
}



