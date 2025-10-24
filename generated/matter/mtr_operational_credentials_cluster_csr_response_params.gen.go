// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTROperationalCredentialsClusterCSRResponseParams */


/* debug [class_header]: Header for MTROperationalCredentialsClusterCSRResponseParams */
// The class instance for the [MTROperationalCredentialsClusterCSRResponseParams] class.
var (
	MTROperationalCredentialsClusterCSRResponseParamsClass     _MTROperationalCredentialsClusterCSRResponseParamsClass
	MTROperationalCredentialsClusterCSRResponseParamsClassOnce sync.Once
)

func getMTROperationalCredentialsClusterCSRResponseParamsClass() _MTROperationalCredentialsClusterCSRResponseParamsClass {
	MTROperationalCredentialsClusterCSRResponseParamsClassOnce.Do(func() {
		MTROperationalCredentialsClusterCSRResponseParamsClass = _MTROperationalCredentialsClusterCSRResponseParamsClass{objc.GetClass("MTROperationalCredentialsClusterCSRResponseParams")}
	})
	return MTROperationalCredentialsClusterCSRResponseParamsClass
}

type _MTROperationalCredentialsClusterCSRResponseParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTROperationalCredentialsClusterCSRResponseParams */
// An interface definition for the [MTROperationalCredentialsClusterCSRResponseParams] class.
type IMTROperationalCredentialsClusterCSRResponseParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTROperationalCredentialsClusterCSRResponseParams */
	// properties:
	AttestationSignature() objc.IObject /* cross-framework: NSData */
	SetAttestationSignature(value objc.IObject /* cross-framework: NSData */)
	NocsrElements() objc.IObject /* cross-framework: NSData */
	SetNocsrElements(value objc.IObject /* cross-framework: NSData */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTROperationalCredentialsClusterCSRResponseParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTROperationalCredentialsClusterCSRResponseParams */
// Alloc allocates a new instance without initialization.
func (mc _MTROperationalCredentialsClusterCSRResponseParamsClass) Alloc() MTROperationalCredentialsClusterCSRResponseParams {
	rv := objc.Send[MTROperationalCredentialsClusterCSRResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTROperationalCredentialsClusterCSRResponseParamsClass) New() MTROperationalCredentialsClusterCSRResponseParams {
	rv := objc.Send[MTROperationalCredentialsClusterCSRResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROperationalCredentialsClusterCSRResponseParams) Init() MTROperationalCredentialsClusterCSRResponseParams {
	rv := objc.Send[MTROperationalCredentialsClusterCSRResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROperationalCredentialsClusterCSRResponseParams) Autorelease() MTROperationalCredentialsClusterCSRResponseParams {
	rv := objc.Send[MTROperationalCredentialsClusterCSRResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROperationalCredentialsClusterCSRResponseParams creates a new MTROperationalCredentialsClusterCSRResponseParams instance.
func NewMTROperationalCredentialsClusterCSRResponseParams() MTROperationalCredentialsClusterCSRResponseParams {
	return getMTROperationalCredentialsClusterCSRResponseParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTROperationalCredentialsClusterCSRResponseParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterCSRResponseParams
type MTROperationalCredentialsClusterCSRResponseParams struct {
	objectivec.Object
}

// MTROperationalCredentialsClusterCSRResponseParamsFrom constructs a [MTROperationalCredentialsClusterCSRResponseParams] from an unsafe.Pointer.
func MTROperationalCredentialsClusterCSRResponseParamsFrom(ptr unsafe.Pointer) MTROperationalCredentialsClusterCSRResponseParams {
	return MTROperationalCredentialsClusterCSRResponseParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTROperationalCredentialsClusterCSRResponseParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterCSRResponseParams/init(responseValue:)
func NewMTROperationalCredentialsClusterCSRResponseParamsWithResponseValueError(responseValue foundation.IDictionary, error_ unsafe.Pointer) MTROperationalCredentialsClusterCSRResponseParams {
	instance := getMTROperationalCredentialsClusterCSRResponseParamsClass().Alloc()
	rv := objc.Send[MTROperationalCredentialsClusterCSRResponseParams](instance.ID, objc.Sel("initWithResponseValue:error:"), responseValue, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTROperationalCredentialsClusterCSRResponseParamsWithResponseValueError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTROperationalCredentialsClusterCSRResponseParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTROperationalCredentialsClusterCSRResponseParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTROperationalCredentialsClusterCSRResponseParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTROperationalCredentialsClusterCSRResponseParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterCSRResponseParams/attestationSignature
func (m_ MTROperationalCredentialsClusterCSRResponseParams) AttestationSignature() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("attestationSignature"))
	return rv
}/* debug [instance_properties/getter]: attestationSignature */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterCSRResponseParams/attestationSignature
func (m_ MTROperationalCredentialsClusterCSRResponseParams) SetAttestationSignature(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAttestationSignature:"), value)
}/* debug [instance_properties/setter]: attestationSignature */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterCSRResponseParams/nocsrElements
func (m_ MTROperationalCredentialsClusterCSRResponseParams) NocsrElements() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("nocsrElements"))
	return rv
}/* debug [instance_properties/getter]: nocsrElements */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterCSRResponseParams/nocsrElements
func (m_ MTROperationalCredentialsClusterCSRResponseParams) SetNocsrElements(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNocsrElements:"), value)
}/* debug [instance_properties/setter]: nocsrElements */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterCSRResponseParams/timedInvokeTimeoutMs
func (m_ MTROperationalCredentialsClusterCSRResponseParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterCSRResponseParams/timedInvokeTimeoutMs
func (m_ MTROperationalCredentialsClusterCSRResponseParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTROperationalCredentialsClusterCSRResponseParams */


