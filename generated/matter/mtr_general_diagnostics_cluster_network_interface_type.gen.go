// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRGeneralDiagnosticsClusterNetworkInterfaceType] class.
var (
	MTRGeneralDiagnosticsClusterNetworkInterfaceTypeClass     _MTRGeneralDiagnosticsClusterNetworkInterfaceTypeClass
	MTRGeneralDiagnosticsClusterNetworkInterfaceTypeClassOnce sync.Once
)

func getMTRGeneralDiagnosticsClusterNetworkInterfaceTypeClass() _MTRGeneralDiagnosticsClusterNetworkInterfaceTypeClass {
	MTRGeneralDiagnosticsClusterNetworkInterfaceTypeClassOnce.Do(func() {
		MTRGeneralDiagnosticsClusterNetworkInterfaceTypeClass = _MTRGeneralDiagnosticsClusterNetworkInterfaceTypeClass{objc.GetClass("MTRGeneralDiagnosticsClusterNetworkInterfaceType")}
	})
	return MTRGeneralDiagnosticsClusterNetworkInterfaceTypeClass
}

type _MTRGeneralDiagnosticsClusterNetworkInterfaceTypeClass struct {
	class objc.Class
}

// An interface definition for the [MTRGeneralDiagnosticsClusterNetworkInterfaceType] class.
type IMTRGeneralDiagnosticsClusterNetworkInterfaceType interface {
	IMTRGeneralDiagnosticsClusterNetworkInterface
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGeneralDiagnosticsClusterNetworkInterfaceType
type MTRGeneralDiagnosticsClusterNetworkInterfaceType struct {
	MTRGeneralDiagnosticsClusterNetworkInterface
}

// MTRGeneralDiagnosticsClusterNetworkInterfaceTypeFrom constructs a [MTRGeneralDiagnosticsClusterNetworkInterfaceType] from an unsafe.Pointer.
func MTRGeneralDiagnosticsClusterNetworkInterfaceTypeFrom(ptr unsafe.Pointer) MTRGeneralDiagnosticsClusterNetworkInterfaceType {
	return MTRGeneralDiagnosticsClusterNetworkInterfaceType{
		MTRGeneralDiagnosticsClusterNetworkInterface: MTRGeneralDiagnosticsClusterNetworkInterfaceFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRGeneralDiagnosticsClusterNetworkInterfaceTypeClass) Alloc() MTRGeneralDiagnosticsClusterNetworkInterfaceType {
	rv := objc.Send[MTRGeneralDiagnosticsClusterNetworkInterfaceType](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRGeneralDiagnosticsClusterNetworkInterfaceTypeClass) New() MTRGeneralDiagnosticsClusterNetworkInterfaceType {
	rv := objc.Send[MTRGeneralDiagnosticsClusterNetworkInterfaceType](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRGeneralDiagnosticsClusterNetworkInterfaceType) Init() MTRGeneralDiagnosticsClusterNetworkInterfaceType {
	rv := objc.Send[MTRGeneralDiagnosticsClusterNetworkInterfaceType](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRGeneralDiagnosticsClusterNetworkInterfaceType) Autorelease() MTRGeneralDiagnosticsClusterNetworkInterfaceType {
	rv := objc.Send[MTRGeneralDiagnosticsClusterNetworkInterfaceType](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRGeneralDiagnosticsClusterNetworkInterfaceType creates a new MTRGeneralDiagnosticsClusterNetworkInterfaceType instance.
func NewMTRGeneralDiagnosticsClusterNetworkInterfaceType() MTRGeneralDiagnosticsClusterNetworkInterfaceType {
	return getMTRGeneralDiagnosticsClusterNetworkInterfaceTypeClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclusternetworkinterfacetype/hardwareaddress
func (m_ MTRGeneralDiagnosticsClusterNetworkInterfaceType) HardwareAddress() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("hardwareAddress"))
	return rv
}


// SetHardwareAddress sets the value of the hardwareAddress property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclusternetworkinterfacetype/hardwareaddress
func (m_ MTRGeneralDiagnosticsClusterNetworkInterfaceType) SetHardwareAddress(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHardwareAddress:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclusternetworkinterfacetype/ipv4addresses
func (m_ MTRGeneralDiagnosticsClusterNetworkInterfaceType) IPv4Addresses() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("iPv4Addresses"))
	return rv
}


// SetIPv4Addresses sets the value of the iPv4Addresses property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclusternetworkinterfacetype/ipv4addresses
func (m_ MTRGeneralDiagnosticsClusterNetworkInterfaceType) SetIPv4Addresses(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIPv4Addresses:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclusternetworkinterfacetype/ipv6addresses
func (m_ MTRGeneralDiagnosticsClusterNetworkInterfaceType) IPv6Addresses() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("iPv6Addresses"))
	return rv
}


// SetIPv6Addresses sets the value of the iPv6Addresses property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclusternetworkinterfacetype/ipv6addresses
func (m_ MTRGeneralDiagnosticsClusterNetworkInterfaceType) SetIPv6Addresses(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIPv6Addresses:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclusternetworkinterfacetype/isoperational
func (m_ MTRGeneralDiagnosticsClusterNetworkInterfaceType) IsOperational() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("isOperational"))
	return rv
}


// SetIsOperational sets the value of the isOperational property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclusternetworkinterfacetype/isoperational
func (m_ MTRGeneralDiagnosticsClusterNetworkInterfaceType) SetIsOperational(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsOperational:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclusternetworkinterfacetype/name
func (m_ MTRGeneralDiagnosticsClusterNetworkInterfaceType) Name() string {
	rv := objc.Send[string](m_.ID, objc.Sel("name"))
	return rv
}


// SetName sets the value of the name property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclusternetworkinterfacetype/name
func (m_ MTRGeneralDiagnosticsClusterNetworkInterfaceType) SetName(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setName:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclusternetworkinterfacetype/offpremiseservicesreachableipv4
func (m_ MTRGeneralDiagnosticsClusterNetworkInterfaceType) OffPremiseServicesReachableIPv4() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("offPremiseServicesReachableIPv4"))
	return rv
}


// SetOffPremiseServicesReachableIPv4 sets the value of the offPremiseServicesReachableIPv4 property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclusternetworkinterfacetype/offpremiseservicesreachableipv4
func (m_ MTRGeneralDiagnosticsClusterNetworkInterfaceType) SetOffPremiseServicesReachableIPv4(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOffPremiseServicesReachableIPv4:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclusternetworkinterfacetype/offpremiseservicesreachableipv6
func (m_ MTRGeneralDiagnosticsClusterNetworkInterfaceType) OffPremiseServicesReachableIPv6() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("offPremiseServicesReachableIPv6"))
	return rv
}


// SetOffPremiseServicesReachableIPv6 sets the value of the offPremiseServicesReachableIPv6 property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclusternetworkinterfacetype/offpremiseservicesreachableipv6
func (m_ MTRGeneralDiagnosticsClusterNetworkInterfaceType) SetOffPremiseServicesReachableIPv6(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOffPremiseServicesReachableIPv6:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclusternetworkinterfacetype/type
func (m_ MTRGeneralDiagnosticsClusterNetworkInterfaceType) Type() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("type"))
	return rv
}


// SetType sets the value of the type property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclusternetworkinterfacetype/type
func (m_ MTRGeneralDiagnosticsClusterNetworkInterfaceType) SetType(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setType:"), value)
}



