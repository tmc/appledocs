// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRCommissioningParameters] class.
var (
	MTRCommissioningParametersClass     _MTRCommissioningParametersClass
	MTRCommissioningParametersClassOnce sync.Once
)

func getMTRCommissioningParametersClass() _MTRCommissioningParametersClass {
	MTRCommissioningParametersClassOnce.Do(func() {
		MTRCommissioningParametersClass = _MTRCommissioningParametersClass{objc.GetClass("MTRCommissioningParameters")}
	})
	return MTRCommissioningParametersClass
}

type _MTRCommissioningParametersClass struct {
	class objc.Class
}

// An interface definition for the [MTRCommissioningParameters] class.
type IMTRCommissioningParameters interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissioningParameters
type MTRCommissioningParameters struct {
	objectivec.Object
}

// MTRCommissioningParametersFrom constructs a [MTRCommissioningParameters] from an unsafe.Pointer.
func MTRCommissioningParametersFrom(ptr unsafe.Pointer) MTRCommissioningParameters {
	return MTRCommissioningParameters{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRCommissioningParametersClass) Alloc() MTRCommissioningParameters {
	rv := objc.Send[MTRCommissioningParameters](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRCommissioningParametersClass) New() MTRCommissioningParameters {
	rv := objc.Send[MTRCommissioningParameters](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRCommissioningParameters) Init() MTRCommissioningParameters {
	rv := objc.Send[MTRCommissioningParameters](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRCommissioningParameters) Autorelease() MTRCommissioningParameters {
	rv := objc.Send[MTRCommissioningParameters](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRCommissioningParameters creates a new MTRCommissioningParameters instance.
func NewMTRCommissioningParameters() MTRCommissioningParameters {
	return getMTRCommissioningParametersClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissioningparameters/attestationnonce
func (m_ MTRCommissioningParameters) AttestationNonce() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("attestationNonce"))
	return rv
}


// SetAttestationNonce sets the value of the attestationNonce property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissioningparameters/attestationnonce
func (m_ MTRCommissioningParameters) SetAttestationNonce(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAttestationNonce:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissioningparameters/countrycode
func (m_ MTRCommissioningParameters) CountryCode() string {
	rv := objc.Send[string](m_.ID, objc.Sel("countryCode"))
	return rv
}


// SetCountryCode sets the value of the countryCode property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissioningparameters/countrycode
func (m_ MTRCommissioningParameters) SetCountryCode(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCountryCode:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissioningparameters/csrnonce-8gx94
func (m_ MTRCommissioningParameters) CsrNonce() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("csrNonce"))
	return rv
}


// SetCsrNonce sets the value of the csrNonce property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissioningparameters/csrnonce-8gx94
func (m_ MTRCommissioningParameters) SetCsrNonce(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCsrNonce:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissioningparameters/deviceattestationdelegate
func (m_ MTRCommissioningParameters) DeviceAttestationDelegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("deviceAttestationDelegate"))
	return rv
}


// SetDeviceAttestationDelegate sets the value of the deviceAttestationDelegate property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissioningparameters/deviceattestationdelegate
func (m_ MTRCommissioningParameters) SetDeviceAttestationDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDeviceAttestationDelegate:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissioningparameters/failsafeexpirytimeoutsecs
func (m_ MTRCommissioningParameters) FailSafeExpiryTimeoutSecs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("failSafeExpiryTimeoutSecs"))
	return rv
}


// SetFailSafeExpiryTimeoutSecs sets the value of the failSafeExpiryTimeoutSecs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissioningparameters/failsafeexpirytimeoutsecs
func (m_ MTRCommissioningParameters) SetFailSafeExpiryTimeoutSecs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFailSafeExpiryTimeoutSecs:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissioningparameters/failsafetimeout
func (m_ MTRCommissioningParameters) FailSafeTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("failSafeTimeout"))
	return rv
}


// SetFailSafeTimeout sets the value of the failSafeTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissioningparameters/failsafetimeout
func (m_ MTRCommissioningParameters) SetFailSafeTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFailSafeTimeout:"), value)
}

// Read device type information from all endpoints during commissioning.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissioningparameters/readendpointinformation
func (m_ MTRCommissioningParameters) ReadEndpointInformation() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("readEndpointInformation"))
	return rv
}


// SetReadEndpointInformation sets the value of the readEndpointInformation property.
// Read device type information from all endpoints during commissioning.

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissioningparameters/readendpointinformation
func (m_ MTRCommissioningParameters) SetReadEndpointInformation(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setReadEndpointInformation:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissioningparameters/skipcommissioningcomplete
func (m_ MTRCommissioningParameters) SkipCommissioningComplete() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("skipCommissioningComplete"))
	return rv
}


// SetSkipCommissioningComplete sets the value of the skipCommissioningComplete property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissioningparameters/skipcommissioningcomplete
func (m_ MTRCommissioningParameters) SetSkipCommissioningComplete(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSkipCommissioningComplete:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissioningparameters/threadoperationaldataset
func (m_ MTRCommissioningParameters) ThreadOperationalDataset() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("threadOperationalDataset"))
	return rv
}


// SetThreadOperationalDataset sets the value of the threadOperationalDataset property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissioningparameters/threadoperationaldataset
func (m_ MTRCommissioningParameters) SetThreadOperationalDataset(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setThreadOperationalDataset:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissioningparameters/wificredentials
func (m_ MTRCommissioningParameters) WifiCredentials() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("wifiCredentials"))
	return rv
}


// SetWifiCredentials sets the value of the wifiCredentials property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissioningparameters/wificredentials
func (m_ MTRCommissioningParameters) SetWifiCredentials(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setWifiCredentials:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissioningparameters/wifissid
func (m_ MTRCommissioningParameters) WifiSSID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("wifiSSID"))
	return rv
}


// SetWifiSSID sets the value of the wifiSSID property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissioningparameters/wifissid
func (m_ MTRCommissioningParameters) SetWifiSSID(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setWifiSSID:"), value)
}



