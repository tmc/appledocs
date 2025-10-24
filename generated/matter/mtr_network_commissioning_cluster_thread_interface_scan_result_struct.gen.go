// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
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

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterthreadinterfacescanresultstruct/channel
func (m_ MTRNetworkCommissioningClusterThreadInterfaceScanResultStruct) Channel() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("channel"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterthreadinterfacescanresultstruct/channel
func (m_ MTRNetworkCommissioningClusterThreadInterfaceScanResultStruct) SetChannel(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setChannel:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterthreadinterfacescanresultstruct/extendedaddress
func (m_ MTRNetworkCommissioningClusterThreadInterfaceScanResultStruct) ExtendedAddress() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("extendedAddress"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterthreadinterfacescanresultstruct/extendedaddress
func (m_ MTRNetworkCommissioningClusterThreadInterfaceScanResultStruct) SetExtendedAddress(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setExtendedAddress:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterthreadinterfacescanresultstruct/extendedpanid
func (m_ MTRNetworkCommissioningClusterThreadInterfaceScanResultStruct) ExtendedPanId() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("extendedPanId"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterthreadinterfacescanresultstruct/extendedpanid
func (m_ MTRNetworkCommissioningClusterThreadInterfaceScanResultStruct) SetExtendedPanId(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setExtendedPanId:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterthreadinterfacescanresultstruct/lqi
func (m_ MTRNetworkCommissioningClusterThreadInterfaceScanResultStruct) Lqi() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("lqi"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterthreadinterfacescanresultstruct/lqi
func (m_ MTRNetworkCommissioningClusterThreadInterfaceScanResultStruct) SetLqi(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLqi:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterthreadinterfacescanresultstruct/networkname
func (m_ MTRNetworkCommissioningClusterThreadInterfaceScanResultStruct) NetworkName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("networkName"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterthreadinterfacescanresultstruct/networkname
func (m_ MTRNetworkCommissioningClusterThreadInterfaceScanResultStruct) SetNetworkName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNetworkName:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterthreadinterfacescanresultstruct/panid
func (m_ MTRNetworkCommissioningClusterThreadInterfaceScanResultStruct) PanId() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("panId"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterthreadinterfacescanresultstruct/panid
func (m_ MTRNetworkCommissioningClusterThreadInterfaceScanResultStruct) SetPanId(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPanId:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterthreadinterfacescanresultstruct/rssi
func (m_ MTRNetworkCommissioningClusterThreadInterfaceScanResultStruct) Rssi() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("rssi"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterthreadinterfacescanresultstruct/rssi
func (m_ MTRNetworkCommissioningClusterThreadInterfaceScanResultStruct) SetRssi(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRssi:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterthreadinterfacescanresultstruct/version
func (m_ MTRNetworkCommissioningClusterThreadInterfaceScanResultStruct) Version() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("version"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterthreadinterfacescanresultstruct/version
func (m_ MTRNetworkCommissioningClusterThreadInterfaceScanResultStruct) SetVersion(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVersion:"), value)
}
