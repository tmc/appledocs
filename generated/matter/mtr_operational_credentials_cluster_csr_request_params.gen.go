// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTROperationalCredentialsClusterCSRRequestParams */


/* debug [class_header]: Header for MTROperationalCredentialsClusterCSRRequestParams */
// The class instance for the [MTROperationalCredentialsClusterCSRRequestParams] class.
var (
	MTROperationalCredentialsClusterCSRRequestParamsClass     _MTROperationalCredentialsClusterCSRRequestParamsClass
	MTROperationalCredentialsClusterCSRRequestParamsClassOnce sync.Once
)

func getMTROperationalCredentialsClusterCSRRequestParamsClass() _MTROperationalCredentialsClusterCSRRequestParamsClass {
	MTROperationalCredentialsClusterCSRRequestParamsClassOnce.Do(func() {
		MTROperationalCredentialsClusterCSRRequestParamsClass = _MTROperationalCredentialsClusterCSRRequestParamsClass{objc.GetClass("MTROperationalCredentialsClusterCSRRequestParams")}
	})
	return MTROperationalCredentialsClusterCSRRequestParamsClass
}

type _MTROperationalCredentialsClusterCSRRequestParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTROperationalCredentialsClusterCSRRequestParams */
// An interface definition for the [MTROperationalCredentialsClusterCSRRequestParams] class.
type IMTROperationalCredentialsClusterCSRRequestParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTROperationalCredentialsClusterCSRRequestParams */
	// properties:
	CsrNonce() objc.IObject /* cross-framework: NSData */
	SetCsrNonce(value objc.IObject /* cross-framework: NSData */)
	IsForUpdateNOC() objc.IObject /* cross-framework: NSNumber */
	SetIsForUpdateNOC(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTROperationalCredentialsClusterCSRRequestParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTROperationalCredentialsClusterCSRRequestParams */
// Alloc allocates a new instance without initialization.
func (mc _MTROperationalCredentialsClusterCSRRequestParamsClass) Alloc() MTROperationalCredentialsClusterCSRRequestParams {
	rv := objc.Send[MTROperationalCredentialsClusterCSRRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTROperationalCredentialsClusterCSRRequestParamsClass) New() MTROperationalCredentialsClusterCSRRequestParams {
	rv := objc.Send[MTROperationalCredentialsClusterCSRRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROperationalCredentialsClusterCSRRequestParams) Init() MTROperationalCredentialsClusterCSRRequestParams {
	rv := objc.Send[MTROperationalCredentialsClusterCSRRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROperationalCredentialsClusterCSRRequestParams) Autorelease() MTROperationalCredentialsClusterCSRRequestParams {
	rv := objc.Send[MTROperationalCredentialsClusterCSRRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROperationalCredentialsClusterCSRRequestParams creates a new MTROperationalCredentialsClusterCSRRequestParams instance.
func NewMTROperationalCredentialsClusterCSRRequestParams() MTROperationalCredentialsClusterCSRRequestParams {
	return getMTROperationalCredentialsClusterCSRRequestParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTROperationalCredentialsClusterCSRRequestParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterCSRRequestParams
type MTROperationalCredentialsClusterCSRRequestParams struct {
	objectivec.Object
}

// MTROperationalCredentialsClusterCSRRequestParamsFrom constructs a [MTROperationalCredentialsClusterCSRRequestParams] from an unsafe.Pointer.
func MTROperationalCredentialsClusterCSRRequestParamsFrom(ptr unsafe.Pointer) MTROperationalCredentialsClusterCSRRequestParams {
	return MTROperationalCredentialsClusterCSRRequestParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTROperationalCredentialsClusterCSRRequestParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTROperationalCredentialsClusterCSRRequestParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTROperationalCredentialsClusterCSRRequestParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTROperationalCredentialsClusterCSRRequestParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTROperationalCredentialsClusterCSRRequestParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterCSRRequestParams/csrNonce
func (m_ MTROperationalCredentialsClusterCSRRequestParams) CsrNonce() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("csrNonce"))
	return rv
}/* debug [instance_properties/getter]: csrNonce */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterCSRRequestParams/csrNonce
func (m_ MTROperationalCredentialsClusterCSRRequestParams) SetCsrNonce(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCsrNonce:"), value)
}/* debug [instance_properties/setter]: csrNonce */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterCSRRequestParams/isForUpdateNOC
func (m_ MTROperationalCredentialsClusterCSRRequestParams) IsForUpdateNOC() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("isForUpdateNOC"))
	return rv
}/* debug [instance_properties/getter]: isForUpdateNOC */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterCSRRequestParams/isForUpdateNOC
func (m_ MTROperationalCredentialsClusterCSRRequestParams) SetIsForUpdateNOC(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsForUpdateNOC:"), value)
}/* debug [instance_properties/setter]: isForUpdateNOC */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterCSRRequestParams/serverSideProcessingTimeout
func (m_ MTROperationalCredentialsClusterCSRRequestParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterCSRRequestParams/serverSideProcessingTimeout
func (m_ MTROperationalCredentialsClusterCSRRequestParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterCSRRequestParams/timedInvokeTimeoutMs
func (m_ MTROperationalCredentialsClusterCSRRequestParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterCSRRequestParams/timedInvokeTimeoutMs
func (m_ MTROperationalCredentialsClusterCSRRequestParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTROperationalCredentialsClusterCSRRequestParams */



