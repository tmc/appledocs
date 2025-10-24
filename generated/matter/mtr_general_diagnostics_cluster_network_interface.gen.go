// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRGeneralDiagnosticsClusterNetworkInterface] class.
var (
	MTRGeneralDiagnosticsClusterNetworkInterfaceClass     _MTRGeneralDiagnosticsClusterNetworkInterfaceClass
	MTRGeneralDiagnosticsClusterNetworkInterfaceClassOnce sync.Once
)

func getMTRGeneralDiagnosticsClusterNetworkInterfaceClass() _MTRGeneralDiagnosticsClusterNetworkInterfaceClass {
	MTRGeneralDiagnosticsClusterNetworkInterfaceClassOnce.Do(func() {
		MTRGeneralDiagnosticsClusterNetworkInterfaceClass = _MTRGeneralDiagnosticsClusterNetworkInterfaceClass{objc.GetClass("MTRGeneralDiagnosticsClusterNetworkInterface")}
	})
	return MTRGeneralDiagnosticsClusterNetworkInterfaceClass
}

type _MTRGeneralDiagnosticsClusterNetworkInterfaceClass struct {
	class objc.Class
}

// An interface definition for the [MTRGeneralDiagnosticsClusterNetworkInterface] class.
type IMTRGeneralDiagnosticsClusterNetworkInterface interface {
	objectivec.IObject
	// properties:
	HardwareAddress() objc.IObject /* cross-framework: Data */
	SetHardwareAddress(value objc.IObject /* cross-framework: Data */)
	IPv4Addresses() unsafe.Pointer
	SetIPv4Addresses(value unsafe.Pointer)
	IPv6Addresses() unsafe.Pointer
	SetIPv6Addresses(value unsafe.Pointer)
	IsOperational() objc.IObject /* cross-framework: NSNumber */
	SetIsOperational(value objc.IObject /* cross-framework: NSNumber */)
	Name() objc.IObject /* cross-framework: NSString */
	SetName(value objc.IObject /* cross-framework: NSString */)
	OffPremiseServicesReachableIPv4() objc.IObject /* cross-framework: NSNumber */
	SetOffPremiseServicesReachableIPv4(value objc.IObject /* cross-framework: NSNumber */)
	OffPremiseServicesReachableIPv6() objc.IObject /* cross-framework: NSNumber */
	SetOffPremiseServicesReachableIPv6(value objc.IObject /* cross-framework: NSNumber */)
	Type() objc.IObject /* cross-framework: NSNumber */
	SetType(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGeneralDiagnosticsClusterNetworkInterface
type MTRGeneralDiagnosticsClusterNetworkInterface struct {
	objectivec.Object
}

// MTRGeneralDiagnosticsClusterNetworkInterfaceFrom constructs a [MTRGeneralDiagnosticsClusterNetworkInterface] from an unsafe.Pointer.
func MTRGeneralDiagnosticsClusterNetworkInterfaceFrom(ptr unsafe.Pointer) MTRGeneralDiagnosticsClusterNetworkInterface {
	return MTRGeneralDiagnosticsClusterNetworkInterface{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRGeneralDiagnosticsClusterNetworkInterfaceClass) Alloc() MTRGeneralDiagnosticsClusterNetworkInterface {
	rv := objc.Send[MTRGeneralDiagnosticsClusterNetworkInterface](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRGeneralDiagnosticsClusterNetworkInterfaceClass) New() MTRGeneralDiagnosticsClusterNetworkInterface {
	rv := objc.Send[MTRGeneralDiagnosticsClusterNetworkInterface](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRGeneralDiagnosticsClusterNetworkInterface) Init() MTRGeneralDiagnosticsClusterNetworkInterface {
	rv := objc.Send[MTRGeneralDiagnosticsClusterNetworkInterface](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRGeneralDiagnosticsClusterNetworkInterface) Autorelease() MTRGeneralDiagnosticsClusterNetworkInterface {
	rv := objc.Send[MTRGeneralDiagnosticsClusterNetworkInterface](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRGeneralDiagnosticsClusterNetworkInterface creates a new MTRGeneralDiagnosticsClusterNetworkInterface instance.
func NewMTRGeneralDiagnosticsClusterNetworkInterface() MTRGeneralDiagnosticsClusterNetworkInterface {
	return getMTRGeneralDiagnosticsClusterNetworkInterfaceClass().New()
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclusternetworkinterface/hardwareaddress
func (m_ MTRGeneralDiagnosticsClusterNetworkInterface) HardwareAddress() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("hardwareAddress"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclusternetworkinterface/hardwareaddress
func (m_ MTRGeneralDiagnosticsClusterNetworkInterface) SetHardwareAddress(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHardwareAddress:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclusternetworkinterface/ipv4addresses
func (m_ MTRGeneralDiagnosticsClusterNetworkInterface) IPv4Addresses() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("iPv4Addresses"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclusternetworkinterface/ipv4addresses
func (m_ MTRGeneralDiagnosticsClusterNetworkInterface) SetIPv4Addresses(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIPv4Addresses:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclusternetworkinterface/ipv6addresses
func (m_ MTRGeneralDiagnosticsClusterNetworkInterface) IPv6Addresses() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("iPv6Addresses"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclusternetworkinterface/ipv6addresses
func (m_ MTRGeneralDiagnosticsClusterNetworkInterface) SetIPv6Addresses(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIPv6Addresses:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclusternetworkinterface/isoperational
func (m_ MTRGeneralDiagnosticsClusterNetworkInterface) IsOperational() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("isOperational"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclusternetworkinterface/isoperational
func (m_ MTRGeneralDiagnosticsClusterNetworkInterface) SetIsOperational(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsOperational:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclusternetworkinterface/name
func (m_ MTRGeneralDiagnosticsClusterNetworkInterface) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("name"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclusternetworkinterface/name
func (m_ MTRGeneralDiagnosticsClusterNetworkInterface) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setName:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclusternetworkinterface/offpremiseservicesreachableipv4
func (m_ MTRGeneralDiagnosticsClusterNetworkInterface) OffPremiseServicesReachableIPv4() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("offPremiseServicesReachableIPv4"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclusternetworkinterface/offpremiseservicesreachableipv4
func (m_ MTRGeneralDiagnosticsClusterNetworkInterface) SetOffPremiseServicesReachableIPv4(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOffPremiseServicesReachableIPv4:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclusternetworkinterface/offpremiseservicesreachableipv6
func (m_ MTRGeneralDiagnosticsClusterNetworkInterface) OffPremiseServicesReachableIPv6() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("offPremiseServicesReachableIPv6"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclusternetworkinterface/offpremiseservicesreachableipv6
func (m_ MTRGeneralDiagnosticsClusterNetworkInterface) SetOffPremiseServicesReachableIPv6(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOffPremiseServicesReachableIPv6:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclusternetworkinterface/type
func (m_ MTRGeneralDiagnosticsClusterNetworkInterface) Type() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("type"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclusternetworkinterface/type
func (m_ MTRGeneralDiagnosticsClusterNetworkInterface) SetType(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setType:"), value)
}
