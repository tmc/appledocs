// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
}

//
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


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclusternetworkinterface/hardwareaddress
func (m_ MTRGeneralDiagnosticsClusterNetworkInterface) HardwareAddress() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("hardwareAddress"))
	return rv
}


// SetHardwareAddress sets the value of the hardwareAddress property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclusternetworkinterface/hardwareaddress
func (m_ MTRGeneralDiagnosticsClusterNetworkInterface) SetHardwareAddress(value foundation.IData) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHardwareAddress:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclusternetworkinterface/ipv4addresses
func (m_ MTRGeneralDiagnosticsClusterNetworkInterface) IPv4Addresses() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("iPv4Addresses"))
	return rv
}


// SetIPv4Addresses sets the value of the iPv4Addresses property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclusternetworkinterface/ipv4addresses
func (m_ MTRGeneralDiagnosticsClusterNetworkInterface) SetIPv4Addresses(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIPv4Addresses:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclusternetworkinterface/ipv6addresses
func (m_ MTRGeneralDiagnosticsClusterNetworkInterface) IPv6Addresses() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("iPv6Addresses"))
	return rv
}


// SetIPv6Addresses sets the value of the iPv6Addresses property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclusternetworkinterface/ipv6addresses
func (m_ MTRGeneralDiagnosticsClusterNetworkInterface) SetIPv6Addresses(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIPv6Addresses:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclusternetworkinterface/isoperational
func (m_ MTRGeneralDiagnosticsClusterNetworkInterface) IsOperational() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("isOperational"))
	return rv
}


// SetIsOperational sets the value of the isOperational property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclusternetworkinterface/isoperational
func (m_ MTRGeneralDiagnosticsClusterNetworkInterface) SetIsOperational(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsOperational:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclusternetworkinterface/name
func (m_ MTRGeneralDiagnosticsClusterNetworkInterface) Name() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("name"))
	return rv
}


// SetName sets the value of the name property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclusternetworkinterface/name
func (m_ MTRGeneralDiagnosticsClusterNetworkInterface) SetName(value appkit.string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setName:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclusternetworkinterface/offpremiseservicesreachableipv4
func (m_ MTRGeneralDiagnosticsClusterNetworkInterface) OffPremiseServicesReachableIPv4() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("offPremiseServicesReachableIPv4"))
	return rv
}


// SetOffPremiseServicesReachableIPv4 sets the value of the offPremiseServicesReachableIPv4 property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclusternetworkinterface/offpremiseservicesreachableipv4
func (m_ MTRGeneralDiagnosticsClusterNetworkInterface) SetOffPremiseServicesReachableIPv4(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOffPremiseServicesReachableIPv4:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclusternetworkinterface/offpremiseservicesreachableipv6
func (m_ MTRGeneralDiagnosticsClusterNetworkInterface) OffPremiseServicesReachableIPv6() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("offPremiseServicesReachableIPv6"))
	return rv
}


// SetOffPremiseServicesReachableIPv6 sets the value of the offPremiseServicesReachableIPv6 property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclusternetworkinterface/offpremiseservicesreachableipv6
func (m_ MTRGeneralDiagnosticsClusterNetworkInterface) SetOffPremiseServicesReachableIPv6(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOffPremiseServicesReachableIPv6:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclusternetworkinterface/type
func (m_ MTRGeneralDiagnosticsClusterNetworkInterface) Type() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("type"))
	return rv
}


// SetType sets the value of the type property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclusternetworkinterface/type
func (m_ MTRGeneralDiagnosticsClusterNetworkInterface) SetType(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setType:"), value)
}



