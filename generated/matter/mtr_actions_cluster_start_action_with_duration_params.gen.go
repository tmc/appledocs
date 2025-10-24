// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRActionsClusterStartActionWithDurationParams */


/* debug [class_header]: Header for MTRActionsClusterStartActionWithDurationParams */
// The class instance for the [MTRActionsClusterStartActionWithDurationParams] class.
var (
	MTRActionsClusterStartActionWithDurationParamsClass     _MTRActionsClusterStartActionWithDurationParamsClass
	MTRActionsClusterStartActionWithDurationParamsClassOnce sync.Once
)

func getMTRActionsClusterStartActionWithDurationParamsClass() _MTRActionsClusterStartActionWithDurationParamsClass {
	MTRActionsClusterStartActionWithDurationParamsClassOnce.Do(func() {
		MTRActionsClusterStartActionWithDurationParamsClass = _MTRActionsClusterStartActionWithDurationParamsClass{objc.GetClass("MTRActionsClusterStartActionWithDurationParams")}
	})
	return MTRActionsClusterStartActionWithDurationParamsClass
}

type _MTRActionsClusterStartActionWithDurationParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRActionsClusterStartActionWithDurationParams */
// An interface definition for the [MTRActionsClusterStartActionWithDurationParams] class.
type IMTRActionsClusterStartActionWithDurationParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRActionsClusterStartActionWithDurationParams */
	// properties:
	ActionID() objc.IObject /* cross-framework: NSNumber */
	SetActionID(value objc.IObject /* cross-framework: NSNumber */)
	Duration() objc.IObject /* cross-framework: NSNumber */
	SetDuration(value objc.IObject /* cross-framework: NSNumber */)
	InvokeID() objc.IObject /* cross-framework: NSNumber */
	SetInvokeID(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRActionsClusterStartActionWithDurationParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRActionsClusterStartActionWithDurationParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRActionsClusterStartActionWithDurationParamsClass) Alloc() MTRActionsClusterStartActionWithDurationParams {
	rv := objc.Send[MTRActionsClusterStartActionWithDurationParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRActionsClusterStartActionWithDurationParamsClass) New() MTRActionsClusterStartActionWithDurationParams {
	rv := objc.Send[MTRActionsClusterStartActionWithDurationParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRActionsClusterStartActionWithDurationParams) Init() MTRActionsClusterStartActionWithDurationParams {
	rv := objc.Send[MTRActionsClusterStartActionWithDurationParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRActionsClusterStartActionWithDurationParams) Autorelease() MTRActionsClusterStartActionWithDurationParams {
	rv := objc.Send[MTRActionsClusterStartActionWithDurationParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRActionsClusterStartActionWithDurationParams creates a new MTRActionsClusterStartActionWithDurationParams instance.
func NewMTRActionsClusterStartActionWithDurationParams() MTRActionsClusterStartActionWithDurationParams {
	return getMTRActionsClusterStartActionWithDurationParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRActionsClusterStartActionWithDurationParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterStartActionWithDurationParams
type MTRActionsClusterStartActionWithDurationParams struct {
	objectivec.Object
}

// MTRActionsClusterStartActionWithDurationParamsFrom constructs a [MTRActionsClusterStartActionWithDurationParams] from an unsafe.Pointer.
func MTRActionsClusterStartActionWithDurationParamsFrom(ptr unsafe.Pointer) MTRActionsClusterStartActionWithDurationParams {
	return MTRActionsClusterStartActionWithDurationParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRActionsClusterStartActionWithDurationParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRActionsClusterStartActionWithDurationParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRActionsClusterStartActionWithDurationParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRActionsClusterStartActionWithDurationParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRActionsClusterStartActionWithDurationParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterStartActionWithDurationParams/actionID
func (m_ MTRActionsClusterStartActionWithDurationParams) ActionID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("actionID"))
	return rv
}/* debug [instance_properties/getter]: actionID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterStartActionWithDurationParams/actionID
func (m_ MTRActionsClusterStartActionWithDurationParams) SetActionID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setActionID:"), value)
}/* debug [instance_properties/setter]: actionID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterStartActionWithDurationParams/duration
func (m_ MTRActionsClusterStartActionWithDurationParams) Duration() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("duration"))
	return rv
}/* debug [instance_properties/getter]: duration */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterStartActionWithDurationParams/duration
func (m_ MTRActionsClusterStartActionWithDurationParams) SetDuration(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDuration:"), value)
}/* debug [instance_properties/setter]: duration */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterStartActionWithDurationParams/invokeID
func (m_ MTRActionsClusterStartActionWithDurationParams) InvokeID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("invokeID"))
	return rv
}/* debug [instance_properties/getter]: invokeID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterStartActionWithDurationParams/invokeID
func (m_ MTRActionsClusterStartActionWithDurationParams) SetInvokeID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInvokeID:"), value)
}/* debug [instance_properties/setter]: invokeID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterStartActionWithDurationParams/serverSideProcessingTimeout
func (m_ MTRActionsClusterStartActionWithDurationParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterStartActionWithDurationParams/serverSideProcessingTimeout
func (m_ MTRActionsClusterStartActionWithDurationParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterStartActionWithDurationParams/timedInvokeTimeoutMs
func (m_ MTRActionsClusterStartActionWithDurationParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterStartActionWithDurationParams/timedInvokeTimeoutMs
func (m_ MTRActionsClusterStartActionWithDurationParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRActionsClusterStartActionWithDurationParams */



