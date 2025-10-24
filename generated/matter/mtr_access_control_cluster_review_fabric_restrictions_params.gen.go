// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRAccessControlClusterReviewFabricRestrictionsParams */


/* debug [class_header]: Header for MTRAccessControlClusterReviewFabricRestrictionsParams */
// The class instance for the [MTRAccessControlClusterReviewFabricRestrictionsParams] class.
var (
	MTRAccessControlClusterReviewFabricRestrictionsParamsClass     _MTRAccessControlClusterReviewFabricRestrictionsParamsClass
	MTRAccessControlClusterReviewFabricRestrictionsParamsClassOnce sync.Once
)

func getMTRAccessControlClusterReviewFabricRestrictionsParamsClass() _MTRAccessControlClusterReviewFabricRestrictionsParamsClass {
	MTRAccessControlClusterReviewFabricRestrictionsParamsClassOnce.Do(func() {
		MTRAccessControlClusterReviewFabricRestrictionsParamsClass = _MTRAccessControlClusterReviewFabricRestrictionsParamsClass{objc.GetClass("MTRAccessControlClusterReviewFabricRestrictionsParams")}
	})
	return MTRAccessControlClusterReviewFabricRestrictionsParamsClass
}

type _MTRAccessControlClusterReviewFabricRestrictionsParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRAccessControlClusterReviewFabricRestrictionsParams */
// An interface definition for the [MTRAccessControlClusterReviewFabricRestrictionsParams] class.
type IMTRAccessControlClusterReviewFabricRestrictionsParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRAccessControlClusterReviewFabricRestrictionsParams */
	// properties:
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRAccessControlClusterReviewFabricRestrictionsParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRAccessControlClusterReviewFabricRestrictionsParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRAccessControlClusterReviewFabricRestrictionsParamsClass) Alloc() MTRAccessControlClusterReviewFabricRestrictionsParams {
	rv := objc.Send[MTRAccessControlClusterReviewFabricRestrictionsParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRAccessControlClusterReviewFabricRestrictionsParamsClass) New() MTRAccessControlClusterReviewFabricRestrictionsParams {
	rv := objc.Send[MTRAccessControlClusterReviewFabricRestrictionsParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRAccessControlClusterReviewFabricRestrictionsParams) Init() MTRAccessControlClusterReviewFabricRestrictionsParams {
	rv := objc.Send[MTRAccessControlClusterReviewFabricRestrictionsParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRAccessControlClusterReviewFabricRestrictionsParams) Autorelease() MTRAccessControlClusterReviewFabricRestrictionsParams {
	rv := objc.Send[MTRAccessControlClusterReviewFabricRestrictionsParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRAccessControlClusterReviewFabricRestrictionsParams creates a new MTRAccessControlClusterReviewFabricRestrictionsParams instance.
func NewMTRAccessControlClusterReviewFabricRestrictionsParams() MTRAccessControlClusterReviewFabricRestrictionsParams {
	return getMTRAccessControlClusterReviewFabricRestrictionsParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRAccessControlClusterReviewFabricRestrictionsParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccessControlClusterReviewFabricRestrictionsParams
type MTRAccessControlClusterReviewFabricRestrictionsParams struct {
	objectivec.Object
}

// MTRAccessControlClusterReviewFabricRestrictionsParamsFrom constructs a [MTRAccessControlClusterReviewFabricRestrictionsParams] from an unsafe.Pointer.
func MTRAccessControlClusterReviewFabricRestrictionsParamsFrom(ptr unsafe.Pointer) MTRAccessControlClusterReviewFabricRestrictionsParams {
	return MTRAccessControlClusterReviewFabricRestrictionsParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRAccessControlClusterReviewFabricRestrictionsParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRAccessControlClusterReviewFabricRestrictionsParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRAccessControlClusterReviewFabricRestrictionsParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRAccessControlClusterReviewFabricRestrictionsParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRAccessControlClusterReviewFabricRestrictionsParams */

// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccessControlClusterReviewFabricRestrictionsParams/serverSideProcessingTimeout
func (m_ MTRAccessControlClusterReviewFabricRestrictionsParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccessControlClusterReviewFabricRestrictionsParams/serverSideProcessingTimeout
func (m_ MTRAccessControlClusterReviewFabricRestrictionsParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusterreviewfabricrestrictionsparams/timedinvoketimeoutms
func (m_ MTRAccessControlClusterReviewFabricRestrictionsParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusterreviewfabricrestrictionsparams/timedinvoketimeoutms
func (m_ MTRAccessControlClusterReviewFabricRestrictionsParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRAccessControlClusterReviewFabricRestrictionsParams */



