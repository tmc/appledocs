// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRActionsClusterEnableActionWithDurationParams */


/* debug [class_header]: Header for MTRActionsClusterEnableActionWithDurationParams */
// The class instance for the [MTRActionsClusterEnableActionWithDurationParams] class.
var (
	MTRActionsClusterEnableActionWithDurationParamsClass     _MTRActionsClusterEnableActionWithDurationParamsClass
	MTRActionsClusterEnableActionWithDurationParamsClassOnce sync.Once
)

func getMTRActionsClusterEnableActionWithDurationParamsClass() _MTRActionsClusterEnableActionWithDurationParamsClass {
	MTRActionsClusterEnableActionWithDurationParamsClassOnce.Do(func() {
		MTRActionsClusterEnableActionWithDurationParamsClass = _MTRActionsClusterEnableActionWithDurationParamsClass{objc.GetClass("MTRActionsClusterEnableActionWithDurationParams")}
	})
	return MTRActionsClusterEnableActionWithDurationParamsClass
}

type _MTRActionsClusterEnableActionWithDurationParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRActionsClusterEnableActionWithDurationParams */
// An interface definition for the [MTRActionsClusterEnableActionWithDurationParams] class.
type IMTRActionsClusterEnableActionWithDurationParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRActionsClusterEnableActionWithDurationParams */
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

	
/* debug [class_interface_methods]: Methods for MTRActionsClusterEnableActionWithDurationParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRActionsClusterEnableActionWithDurationParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRActionsClusterEnableActionWithDurationParamsClass) Alloc() MTRActionsClusterEnableActionWithDurationParams {
	rv := objc.Send[MTRActionsClusterEnableActionWithDurationParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRActionsClusterEnableActionWithDurationParamsClass) New() MTRActionsClusterEnableActionWithDurationParams {
	rv := objc.Send[MTRActionsClusterEnableActionWithDurationParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRActionsClusterEnableActionWithDurationParams) Init() MTRActionsClusterEnableActionWithDurationParams {
	rv := objc.Send[MTRActionsClusterEnableActionWithDurationParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRActionsClusterEnableActionWithDurationParams) Autorelease() MTRActionsClusterEnableActionWithDurationParams {
	rv := objc.Send[MTRActionsClusterEnableActionWithDurationParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRActionsClusterEnableActionWithDurationParams creates a new MTRActionsClusterEnableActionWithDurationParams instance.
func NewMTRActionsClusterEnableActionWithDurationParams() MTRActionsClusterEnableActionWithDurationParams {
	return getMTRActionsClusterEnableActionWithDurationParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRActionsClusterEnableActionWithDurationParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterEnableActionWithDurationParams
type MTRActionsClusterEnableActionWithDurationParams struct {
	objectivec.Object
}

// MTRActionsClusterEnableActionWithDurationParamsFrom constructs a [MTRActionsClusterEnableActionWithDurationParams] from an unsafe.Pointer.
func MTRActionsClusterEnableActionWithDurationParamsFrom(ptr unsafe.Pointer) MTRActionsClusterEnableActionWithDurationParams {
	return MTRActionsClusterEnableActionWithDurationParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRActionsClusterEnableActionWithDurationParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRActionsClusterEnableActionWithDurationParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRActionsClusterEnableActionWithDurationParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRActionsClusterEnableActionWithDurationParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRActionsClusterEnableActionWithDurationParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterEnableActionWithDurationParams/actionID
func (m_ MTRActionsClusterEnableActionWithDurationParams) ActionID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("actionID"))
	return rv
}/* debug [instance_properties/getter]: actionID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterEnableActionWithDurationParams/actionID
func (m_ MTRActionsClusterEnableActionWithDurationParams) SetActionID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setActionID:"), value)
}/* debug [instance_properties/setter]: actionID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterEnableActionWithDurationParams/duration
func (m_ MTRActionsClusterEnableActionWithDurationParams) Duration() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("duration"))
	return rv
}/* debug [instance_properties/getter]: duration */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterEnableActionWithDurationParams/duration
func (m_ MTRActionsClusterEnableActionWithDurationParams) SetDuration(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDuration:"), value)
}/* debug [instance_properties/setter]: duration */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterEnableActionWithDurationParams/invokeID
func (m_ MTRActionsClusterEnableActionWithDurationParams) InvokeID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("invokeID"))
	return rv
}/* debug [instance_properties/getter]: invokeID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterEnableActionWithDurationParams/invokeID
func (m_ MTRActionsClusterEnableActionWithDurationParams) SetInvokeID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInvokeID:"), value)
}/* debug [instance_properties/setter]: invokeID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterEnableActionWithDurationParams/serverSideProcessingTimeout
func (m_ MTRActionsClusterEnableActionWithDurationParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterEnableActionWithDurationParams/serverSideProcessingTimeout
func (m_ MTRActionsClusterEnableActionWithDurationParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterEnableActionWithDurationParams/timedInvokeTimeoutMs
func (m_ MTRActionsClusterEnableActionWithDurationParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterEnableActionWithDurationParams/timedInvokeTimeoutMs
func (m_ MTRActionsClusterEnableActionWithDurationParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRActionsClusterEnableActionWithDurationParams */



