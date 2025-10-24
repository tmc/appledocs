// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRCommissionerControlClusterCommissionNodeParams */


/* debug [class_header]: Header for MTRCommissionerControlClusterCommissionNodeParams */
// The class instance for the [MTRCommissionerControlClusterCommissionNodeParams] class.
var (
	MTRCommissionerControlClusterCommissionNodeParamsClass     _MTRCommissionerControlClusterCommissionNodeParamsClass
	MTRCommissionerControlClusterCommissionNodeParamsClassOnce sync.Once
)

func getMTRCommissionerControlClusterCommissionNodeParamsClass() _MTRCommissionerControlClusterCommissionNodeParamsClass {
	MTRCommissionerControlClusterCommissionNodeParamsClassOnce.Do(func() {
		MTRCommissionerControlClusterCommissionNodeParamsClass = _MTRCommissionerControlClusterCommissionNodeParamsClass{objc.GetClass("MTRCommissionerControlClusterCommissionNodeParams")}
	})
	return MTRCommissionerControlClusterCommissionNodeParamsClass
}

type _MTRCommissionerControlClusterCommissionNodeParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRCommissionerControlClusterCommissionNodeParams */
// An interface definition for the [MTRCommissionerControlClusterCommissionNodeParams] class.
type IMTRCommissionerControlClusterCommissionNodeParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRCommissionerControlClusterCommissionNodeParams */
	// properties:
	ResponseTimeoutSeconds() objc.IObject /* cross-framework: NSNumber */
	SetResponseTimeoutSeconds(value objc.IObject /* cross-framework: NSNumber */)
	RequestID() objc.IObject /* cross-framework: NSNumber */
	SetRequestID(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRCommissionerControlClusterCommissionNodeParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRCommissionerControlClusterCommissionNodeParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRCommissionerControlClusterCommissionNodeParamsClass) Alloc() MTRCommissionerControlClusterCommissionNodeParams {
	rv := objc.Send[MTRCommissionerControlClusterCommissionNodeParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRCommissionerControlClusterCommissionNodeParamsClass) New() MTRCommissionerControlClusterCommissionNodeParams {
	rv := objc.Send[MTRCommissionerControlClusterCommissionNodeParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRCommissionerControlClusterCommissionNodeParams) Init() MTRCommissionerControlClusterCommissionNodeParams {
	rv := objc.Send[MTRCommissionerControlClusterCommissionNodeParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRCommissionerControlClusterCommissionNodeParams) Autorelease() MTRCommissionerControlClusterCommissionNodeParams {
	rv := objc.Send[MTRCommissionerControlClusterCommissionNodeParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRCommissionerControlClusterCommissionNodeParams creates a new MTRCommissionerControlClusterCommissionNodeParams instance.
func NewMTRCommissionerControlClusterCommissionNodeParams() MTRCommissionerControlClusterCommissionNodeParams {
	return getMTRCommissionerControlClusterCommissionNodeParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRCommissionerControlClusterCommissionNodeParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissionerControlClusterCommissionNodeParams
type MTRCommissionerControlClusterCommissionNodeParams struct {
	objectivec.Object
}

// MTRCommissionerControlClusterCommissionNodeParamsFrom constructs a [MTRCommissionerControlClusterCommissionNodeParams] from an unsafe.Pointer.
func MTRCommissionerControlClusterCommissionNodeParamsFrom(ptr unsafe.Pointer) MTRCommissionerControlClusterCommissionNodeParams {
	return MTRCommissionerControlClusterCommissionNodeParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRCommissionerControlClusterCommissionNodeParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRCommissionerControlClusterCommissionNodeParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRCommissionerControlClusterCommissionNodeParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRCommissionerControlClusterCommissionNodeParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRCommissionerControlClusterCommissionNodeParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissionerControlClusterCommissionNodeParams/responseTimeoutSeconds
func (m_ MTRCommissionerControlClusterCommissionNodeParams) ResponseTimeoutSeconds() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("responseTimeoutSeconds"))
	return rv
}/* debug [instance_properties/getter]: responseTimeoutSeconds */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissionerControlClusterCommissionNodeParams/responseTimeoutSeconds
func (m_ MTRCommissionerControlClusterCommissionNodeParams) SetResponseTimeoutSeconds(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setResponseTimeoutSeconds:"), value)
}/* debug [instance_properties/setter]: responseTimeoutSeconds */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissionercontrolclustercommissionnodeparams/requestid
func (m_ MTRCommissionerControlClusterCommissionNodeParams) RequestID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("requestID"))
	return rv
}/* debug [instance_properties/getter]: requestID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissionercontrolclustercommissionnodeparams/requestid
func (m_ MTRCommissionerControlClusterCommissionNodeParams) SetRequestID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRequestID:"), value)
}/* debug [instance_properties/setter]: requestID */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissionercontrolclustercommissionnodeparams/serversideprocessingtimeout
func (m_ MTRCommissionerControlClusterCommissionNodeParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissionercontrolclustercommissionnodeparams/serversideprocessingtimeout
func (m_ MTRCommissionerControlClusterCommissionNodeParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissionercontrolclustercommissionnodeparams/timedinvoketimeoutms
func (m_ MTRCommissionerControlClusterCommissionNodeParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissionercontrolclustercommissionnodeparams/timedinvoketimeoutms
func (m_ MTRCommissionerControlClusterCommissionNodeParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRCommissionerControlClusterCommissionNodeParams */



