// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTROperationalCredentialsClusterAttestationResponseParams */


/* debug [class_header]: Header for MTROperationalCredentialsClusterAttestationResponseParams */
// The class instance for the [MTROperationalCredentialsClusterAttestationResponseParams] class.
var (
	MTROperationalCredentialsClusterAttestationResponseParamsClass     _MTROperationalCredentialsClusterAttestationResponseParamsClass
	MTROperationalCredentialsClusterAttestationResponseParamsClassOnce sync.Once
)

func getMTROperationalCredentialsClusterAttestationResponseParamsClass() _MTROperationalCredentialsClusterAttestationResponseParamsClass {
	MTROperationalCredentialsClusterAttestationResponseParamsClassOnce.Do(func() {
		MTROperationalCredentialsClusterAttestationResponseParamsClass = _MTROperationalCredentialsClusterAttestationResponseParamsClass{objc.GetClass("MTROperationalCredentialsClusterAttestationResponseParams")}
	})
	return MTROperationalCredentialsClusterAttestationResponseParamsClass
}

type _MTROperationalCredentialsClusterAttestationResponseParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTROperationalCredentialsClusterAttestationResponseParams */
// An interface definition for the [MTROperationalCredentialsClusterAttestationResponseParams] class.
type IMTROperationalCredentialsClusterAttestationResponseParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTROperationalCredentialsClusterAttestationResponseParams */
	// properties:
	AttestationElements() objc.IObject /* cross-framework: NSData */
	SetAttestationElements(value objc.IObject /* cross-framework: NSData */)
	AttestationSignature() objc.IObject /* cross-framework: NSData */
	SetAttestationSignature(value objc.IObject /* cross-framework: NSData */)
	Signature() objc.IObject /* cross-framework: NSData */
	SetSignature(value objc.IObject /* cross-framework: NSData */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTROperationalCredentialsClusterAttestationResponseParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTROperationalCredentialsClusterAttestationResponseParams */
// Alloc allocates a new instance without initialization.
func (mc _MTROperationalCredentialsClusterAttestationResponseParamsClass) Alloc() MTROperationalCredentialsClusterAttestationResponseParams {
	rv := objc.Send[MTROperationalCredentialsClusterAttestationResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTROperationalCredentialsClusterAttestationResponseParamsClass) New() MTROperationalCredentialsClusterAttestationResponseParams {
	rv := objc.Send[MTROperationalCredentialsClusterAttestationResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROperationalCredentialsClusterAttestationResponseParams) Init() MTROperationalCredentialsClusterAttestationResponseParams {
	rv := objc.Send[MTROperationalCredentialsClusterAttestationResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROperationalCredentialsClusterAttestationResponseParams) Autorelease() MTROperationalCredentialsClusterAttestationResponseParams {
	rv := objc.Send[MTROperationalCredentialsClusterAttestationResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROperationalCredentialsClusterAttestationResponseParams creates a new MTROperationalCredentialsClusterAttestationResponseParams instance.
func NewMTROperationalCredentialsClusterAttestationResponseParams() MTROperationalCredentialsClusterAttestationResponseParams {
	return getMTROperationalCredentialsClusterAttestationResponseParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTROperationalCredentialsClusterAttestationResponseParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterAttestationResponseParams
type MTROperationalCredentialsClusterAttestationResponseParams struct {
	objectivec.Object
}

// MTROperationalCredentialsClusterAttestationResponseParamsFrom constructs a [MTROperationalCredentialsClusterAttestationResponseParams] from an unsafe.Pointer.
func MTROperationalCredentialsClusterAttestationResponseParamsFrom(ptr unsafe.Pointer) MTROperationalCredentialsClusterAttestationResponseParams {
	return MTROperationalCredentialsClusterAttestationResponseParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTROperationalCredentialsClusterAttestationResponseParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterAttestationResponseParams/init(responseValue:)
func NewMTROperationalCredentialsClusterAttestationResponseParamsWithResponseValueError(responseValue foundation.IDictionary, error_ unsafe.Pointer) MTROperationalCredentialsClusterAttestationResponseParams {
	instance := getMTROperationalCredentialsClusterAttestationResponseParamsClass().Alloc()
	rv := objc.Send[MTROperationalCredentialsClusterAttestationResponseParams](instance.ID, objc.Sel("initWithResponseValue:error:"), responseValue, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTROperationalCredentialsClusterAttestationResponseParamsWithResponseValueError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTROperationalCredentialsClusterAttestationResponseParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTROperationalCredentialsClusterAttestationResponseParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTROperationalCredentialsClusterAttestationResponseParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTROperationalCredentialsClusterAttestationResponseParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterAttestationResponseParams/attestationElements
func (m_ MTROperationalCredentialsClusterAttestationResponseParams) AttestationElements() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("attestationElements"))
	return rv
}/* debug [instance_properties/getter]: attestationElements */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterAttestationResponseParams/attestationElements
func (m_ MTROperationalCredentialsClusterAttestationResponseParams) SetAttestationElements(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAttestationElements:"), value)
}/* debug [instance_properties/setter]: attestationElements */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterAttestationResponseParams/attestationSignature
func (m_ MTROperationalCredentialsClusterAttestationResponseParams) AttestationSignature() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("attestationSignature"))
	return rv
}/* debug [instance_properties/getter]: attestationSignature */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterAttestationResponseParams/attestationSignature
func (m_ MTROperationalCredentialsClusterAttestationResponseParams) SetAttestationSignature(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAttestationSignature:"), value)
}/* debug [instance_properties/setter]: attestationSignature */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterAttestationResponseParams/signature
func (m_ MTROperationalCredentialsClusterAttestationResponseParams) Signature() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("signature"))
	return rv
}/* debug [instance_properties/getter]: signature */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterAttestationResponseParams/signature
func (m_ MTROperationalCredentialsClusterAttestationResponseParams) SetSignature(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSignature:"), value)
}/* debug [instance_properties/setter]: signature */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterAttestationResponseParams/timedInvokeTimeoutMs
func (m_ MTROperationalCredentialsClusterAttestationResponseParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterAttestationResponseParams/timedInvokeTimeoutMs
func (m_ MTROperationalCredentialsClusterAttestationResponseParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTROperationalCredentialsClusterAttestationResponseParams */


