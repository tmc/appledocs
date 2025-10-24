// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRCommissioningParameters */


/* debug [class_header]: Header for MTRCommissioningParameters */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRCommissioningParameters */
// An interface definition for the [MTRCommissioningParameters] class.
type IMTRCommissioningParameters interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRCommissioningParameters */
	// properties:
	AttestationNonce() objc.IObject /* cross-framework: NSData */
	SetAttestationNonce(value objc.IObject /* cross-framework: NSData */)
	CountryCode() objc.IObject /* cross-framework: NSString */
	SetCountryCode(value objc.IObject /* cross-framework: NSString */)
	CsrNonce() objc.IObject /* cross-framework: NSData */
	SetCsrNonce(value objc.IObject /* cross-framework: NSData */)
	CSRNonce() objc.IObject /* cross-framework: NSData */
	SetCSRNonce(value objc.IObject /* cross-framework: NSData */)
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
	ThreadOperationalDataset() objc.IObject /* cross-framework: NSData */
	SetThreadOperationalDataset(value objc.IObject /* cross-framework: NSData */)
	WifiCredentials() objc.IObject /* cross-framework: NSData */
	SetWifiCredentials(value objc.IObject /* cross-framework: NSData */)
	WifiSSID() objc.IObject /* cross-framework: NSData */
	SetWifiSSID(value objc.IObject /* cross-framework: NSData */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRCommissioningParameters */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRCommissioningParameters */
// Alloc allocates a new instance without initialization.
func (mc _MTRCommissioningParametersClass) Alloc() MTRCommissioningParameters {
	rv := objc.Send[MTRCommissioningParameters](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRCommissioningParameters */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissioningParameters
type MTRCommissioningParameters struct {
	objectivec.Object
}

// MTRCommissioningParametersFrom constructs a [MTRCommissioningParameters] from an unsafe.Pointer.
func MTRCommissioningParametersFrom(ptr unsafe.Pointer) MTRCommissioningParameters {
	return MTRCommissioningParameters{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRCommissioningParameters *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRCommissioningParameters */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRCommissioningParameters */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRCommissioningParameters */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRCommissioningParameters */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissioningParameters/attestationNonce
func (m_ MTRCommissioningParameters) AttestationNonce() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("attestationNonce"))
	return rv
}/* debug [instance_properties/getter]: attestationNonce */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissioningParameters/attestationNonce
func (m_ MTRCommissioningParameters) SetAttestationNonce(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAttestationNonce:"), value)
}/* debug [instance_properties/setter]: attestationNonce */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissioningParameters/countryCode
func (m_ MTRCommissioningParameters) CountryCode() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("countryCode"))
	return rv
}/* debug [instance_properties/getter]: countryCode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissioningParameters/countryCode
func (m_ MTRCommissioningParameters) SetCountryCode(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCountryCode:"), value)
}/* debug [instance_properties/setter]: countryCode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissioningParameters/csrNonce-8gx94
func (m_ MTRCommissioningParameters) CsrNonce() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("csrNonce"))
	return rv
}/* debug [instance_properties/getter]: csrNonce */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissioningParameters/csrNonce-8gx94
func (m_ MTRCommissioningParameters) SetCsrNonce(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCsrNonce:"), value)
}/* debug [instance_properties/setter]: csrNonce */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissioningParameters/csrNonce-9eaxq
func (m_ MTRCommissioningParameters) CSRNonce() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("CSRNonce"))
	return rv
}/* debug [instance_properties/getter]: CSRNonce */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissioningParameters/csrNonce-9eaxq
func (m_ MTRCommissioningParameters) SetCSRNonce(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCSRNonce:"), value)
}/* debug [instance_properties/setter]: CSRNonce */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissioningParameters/deviceAttestationDelegate
func (m_ MTRCommissioningParameters) DeviceAttestationDelegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("deviceAttestationDelegate"))
	return rv
}/* debug [instance_properties/getter]: deviceAttestationDelegate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissioningParameters/deviceAttestationDelegate
func (m_ MTRCommissioningParameters) SetDeviceAttestationDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDeviceAttestationDelegate:"), value)
}/* debug [instance_properties/setter]: deviceAttestationDelegate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissioningParameters/failSafeExpiryTimeoutSecs
func (m_ MTRCommissioningParameters) FailSafeExpiryTimeoutSecs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("failSafeExpiryTimeoutSecs"))
	return rv
}/* debug [instance_properties/getter]: failSafeExpiryTimeoutSecs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissioningParameters/failSafeExpiryTimeoutSecs
func (m_ MTRCommissioningParameters) SetFailSafeExpiryTimeoutSecs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFailSafeExpiryTimeoutSecs:"), value)
}/* debug [instance_properties/setter]: failSafeExpiryTimeoutSecs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissioningParameters/failSafeTimeout
func (m_ MTRCommissioningParameters) FailSafeTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("failSafeTimeout"))
	return rv
}/* debug [instance_properties/getter]: failSafeTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissioningParameters/failSafeTimeout
func (m_ MTRCommissioningParameters) SetFailSafeTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFailSafeTimeout:"), value)
}/* debug [instance_properties/setter]: failSafeTimeout */


// Read device type information from all endpoints during commissioning. Defaults to NO.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissioningParameters/readEndpointInformation
func (m_ MTRCommissioningParameters) ReadEndpointInformation() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("readEndpointInformation"))
	return rv
}/* debug [instance_properties/getter]: readEndpointInformation */


// Read device type information from all endpoints during commissioning. Defaults to NO.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissioningParameters/readEndpointInformation
func (m_ MTRCommissioningParameters) SetReadEndpointInformation(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setReadEndpointInformation:"), value)
}/* debug [instance_properties/setter]: readEndpointInformation */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissioningParameters/skipCommissioningComplete
func (m_ MTRCommissioningParameters) SkipCommissioningComplete() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("skipCommissioningComplete"))
	return rv
}/* debug [instance_properties/getter]: skipCommissioningComplete */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissioningParameters/skipCommissioningComplete
func (m_ MTRCommissioningParameters) SetSkipCommissioningComplete(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSkipCommissioningComplete:"), value)
}/* debug [instance_properties/setter]: skipCommissioningComplete */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissioningParameters/threadOperationalDataset
func (m_ MTRCommissioningParameters) ThreadOperationalDataset() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("threadOperationalDataset"))
	return rv
}/* debug [instance_properties/getter]: threadOperationalDataset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissioningParameters/threadOperationalDataset
func (m_ MTRCommissioningParameters) SetThreadOperationalDataset(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setThreadOperationalDataset:"), value)
}/* debug [instance_properties/setter]: threadOperationalDataset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissioningParameters/wifiCredentials
func (m_ MTRCommissioningParameters) WifiCredentials() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("wifiCredentials"))
	return rv
}/* debug [instance_properties/getter]: wifiCredentials */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissioningParameters/wifiCredentials
func (m_ MTRCommissioningParameters) SetWifiCredentials(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setWifiCredentials:"), value)
}/* debug [instance_properties/setter]: wifiCredentials */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissioningParameters/wifiSSID
func (m_ MTRCommissioningParameters) WifiSSID() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("wifiSSID"))
	return rv
}/* debug [instance_properties/getter]: wifiSSID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissioningParameters/wifiSSID
func (m_ MTRCommissioningParameters) SetWifiSSID(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setWifiSSID:"), value)
}/* debug [instance_properties/setter]: wifiSSID */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRCommissioningParameters */



