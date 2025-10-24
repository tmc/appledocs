// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTROperationalCredentialsClusterAttestationRequestParams */


/* debug [class_header]: Header for MTROperationalCredentialsClusterAttestationRequestParams */
// The class instance for the [MTROperationalCredentialsClusterAttestationRequestParams] class.
var (
	MTROperationalCredentialsClusterAttestationRequestParamsClass     _MTROperationalCredentialsClusterAttestationRequestParamsClass
	MTROperationalCredentialsClusterAttestationRequestParamsClassOnce sync.Once
)

func getMTROperationalCredentialsClusterAttestationRequestParamsClass() _MTROperationalCredentialsClusterAttestationRequestParamsClass {
	MTROperationalCredentialsClusterAttestationRequestParamsClassOnce.Do(func() {
		MTROperationalCredentialsClusterAttestationRequestParamsClass = _MTROperationalCredentialsClusterAttestationRequestParamsClass{objc.GetClass("MTROperationalCredentialsClusterAttestationRequestParams")}
	})
	return MTROperationalCredentialsClusterAttestationRequestParamsClass
}

type _MTROperationalCredentialsClusterAttestationRequestParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTROperationalCredentialsClusterAttestationRequestParams */
// An interface definition for the [MTROperationalCredentialsClusterAttestationRequestParams] class.
type IMTROperationalCredentialsClusterAttestationRequestParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTROperationalCredentialsClusterAttestationRequestParams */
	// properties:
	AttestationNonce() objc.IObject /* cross-framework: NSData */
	SetAttestationNonce(value objc.IObject /* cross-framework: NSData */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTROperationalCredentialsClusterAttestationRequestParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTROperationalCredentialsClusterAttestationRequestParams */
// Alloc allocates a new instance without initialization.
func (mc _MTROperationalCredentialsClusterAttestationRequestParamsClass) Alloc() MTROperationalCredentialsClusterAttestationRequestParams {
	rv := objc.Send[MTROperationalCredentialsClusterAttestationRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTROperationalCredentialsClusterAttestationRequestParamsClass) New() MTROperationalCredentialsClusterAttestationRequestParams {
	rv := objc.Send[MTROperationalCredentialsClusterAttestationRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROperationalCredentialsClusterAttestationRequestParams) Init() MTROperationalCredentialsClusterAttestationRequestParams {
	rv := objc.Send[MTROperationalCredentialsClusterAttestationRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROperationalCredentialsClusterAttestationRequestParams) Autorelease() MTROperationalCredentialsClusterAttestationRequestParams {
	rv := objc.Send[MTROperationalCredentialsClusterAttestationRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROperationalCredentialsClusterAttestationRequestParams creates a new MTROperationalCredentialsClusterAttestationRequestParams instance.
func NewMTROperationalCredentialsClusterAttestationRequestParams() MTROperationalCredentialsClusterAttestationRequestParams {
	return getMTROperationalCredentialsClusterAttestationRequestParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTROperationalCredentialsClusterAttestationRequestParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterAttestationRequestParams
type MTROperationalCredentialsClusterAttestationRequestParams struct {
	objectivec.Object
}

// MTROperationalCredentialsClusterAttestationRequestParamsFrom constructs a [MTROperationalCredentialsClusterAttestationRequestParams] from an unsafe.Pointer.
func MTROperationalCredentialsClusterAttestationRequestParamsFrom(ptr unsafe.Pointer) MTROperationalCredentialsClusterAttestationRequestParams {
	return MTROperationalCredentialsClusterAttestationRequestParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTROperationalCredentialsClusterAttestationRequestParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTROperationalCredentialsClusterAttestationRequestParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTROperationalCredentialsClusterAttestationRequestParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTROperationalCredentialsClusterAttestationRequestParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTROperationalCredentialsClusterAttestationRequestParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterAttestationRequestParams/attestationNonce
func (m_ MTROperationalCredentialsClusterAttestationRequestParams) AttestationNonce() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("attestationNonce"))
	return rv
}/* debug [instance_properties/getter]: attestationNonce */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterAttestationRequestParams/attestationNonce
func (m_ MTROperationalCredentialsClusterAttestationRequestParams) SetAttestationNonce(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAttestationNonce:"), value)
}/* debug [instance_properties/setter]: attestationNonce */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterAttestationRequestParams/serverSideProcessingTimeout
func (m_ MTROperationalCredentialsClusterAttestationRequestParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterAttestationRequestParams/serverSideProcessingTimeout
func (m_ MTROperationalCredentialsClusterAttestationRequestParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterAttestationRequestParams/timedInvokeTimeoutMs
func (m_ MTROperationalCredentialsClusterAttestationRequestParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterAttestationRequestParams/timedInvokeTimeoutMs
func (m_ MTROperationalCredentialsClusterAttestationRequestParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTROperationalCredentialsClusterAttestationRequestParams */



