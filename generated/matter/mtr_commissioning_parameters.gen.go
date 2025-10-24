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
	// properties:
	AttestationNonce() objc.IObject /* cross-framework: Data */
	SetAttestationNonce(value objc.IObject /* cross-framework: Data */)
	CountryCode() objc.IObject /* cross-framework: NSString */
	SetCountryCode(value objc.IObject /* cross-framework: NSString */)
	CsrNonce() objc.IObject /* cross-framework: Data */
	SetCsrNonce(value objc.IObject /* cross-framework: Data */)
	DeviceAttestationDelegate() unsafe.Pointer
	SetDeviceAttestationDelegate(value unsafe.Pointer)
	FailSafeExpiryTimeoutSecs() objc.IObject /* cross-framework: NSNumber */
	SetFailSafeExpiryTimeoutSecs(value objc.IObject /* cross-framework: NSNumber */)
	FailSafeTimeout() objc.IObject /* cross-framework: NSNumber */
	SetFailSafeTimeout(value objc.IObject /* cross-framework: NSNumber */)
	ReadEndpointInformation() bool
	SetReadEndpointInformation(value bool)
	SkipCommissioningComplete() bool
	SetSkipCommissioningComplete(value bool)
	ThreadOperationalDataset() objc.IObject /* cross-framework: Data */
	SetThreadOperationalDataset(value objc.IObject /* cross-framework: Data */)
	WifiCredentials() objc.IObject /* cross-framework: Data */
	SetWifiCredentials(value objc.IObject /* cross-framework: Data */)
	WifiSSID() objc.IObject /* cross-framework: Data */
	SetWifiSSID(value objc.IObject /* cross-framework: Data */)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissioningparameters/attestationnonce
func (m_ MTRCommissioningParameters) AttestationNonce() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("attestationNonce"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissioningparameters/attestationnonce
func (m_ MTRCommissioningParameters) SetAttestationNonce(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAttestationNonce:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissioningparameters/countrycode
func (m_ MTRCommissioningParameters) CountryCode() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("countryCode"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissioningparameters/countrycode
func (m_ MTRCommissioningParameters) SetCountryCode(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCountryCode:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissioningparameters/csrnonce-8gx94
func (m_ MTRCommissioningParameters) CsrNonce() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("csrNonce"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissioningparameters/csrnonce-8gx94
func (m_ MTRCommissioningParameters) SetCsrNonce(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCsrNonce:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissioningparameters/deviceattestationdelegate
func (m_ MTRCommissioningParameters) DeviceAttestationDelegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("deviceAttestationDelegate"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissioningparameters/deviceattestationdelegate
func (m_ MTRCommissioningParameters) SetDeviceAttestationDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDeviceAttestationDelegate:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissioningparameters/failsafeexpirytimeoutsecs
func (m_ MTRCommissioningParameters) FailSafeExpiryTimeoutSecs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("failSafeExpiryTimeoutSecs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissioningparameters/failsafeexpirytimeoutsecs
func (m_ MTRCommissioningParameters) SetFailSafeExpiryTimeoutSecs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFailSafeExpiryTimeoutSecs:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissioningparameters/failsafetimeout
func (m_ MTRCommissioningParameters) FailSafeTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("failSafeTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissioningparameters/failsafetimeout
func (m_ MTRCommissioningParameters) SetFailSafeTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFailSafeTimeout:"), value)
}


// Read device type information from all endpoints during commissioning.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissioningparameters/readendpointinformation
func (m_ MTRCommissioningParameters) ReadEndpointInformation() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("readEndpointInformation"))
	return rv
}


// Read device type information from all endpoints during commissioning.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissioningparameters/readendpointinformation
func (m_ MTRCommissioningParameters) SetReadEndpointInformation(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setReadEndpointInformation:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissioningparameters/skipcommissioningcomplete
func (m_ MTRCommissioningParameters) SkipCommissioningComplete() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("skipCommissioningComplete"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissioningparameters/skipcommissioningcomplete
func (m_ MTRCommissioningParameters) SetSkipCommissioningComplete(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSkipCommissioningComplete:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissioningparameters/threadoperationaldataset
func (m_ MTRCommissioningParameters) ThreadOperationalDataset() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("threadOperationalDataset"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissioningparameters/threadoperationaldataset
func (m_ MTRCommissioningParameters) SetThreadOperationalDataset(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setThreadOperationalDataset:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissioningparameters/wificredentials
func (m_ MTRCommissioningParameters) WifiCredentials() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("wifiCredentials"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissioningparameters/wificredentials
func (m_ MTRCommissioningParameters) SetWifiCredentials(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setWifiCredentials:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissioningparameters/wifissid
func (m_ MTRCommissioningParameters) WifiSSID() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("wifiSSID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissioningparameters/wifissid
func (m_ MTRCommissioningParameters) SetWifiSSID(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setWifiSSID:"), value)
}



