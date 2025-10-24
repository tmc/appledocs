// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRActionsClusterDisableActionWithDurationParams */


/* debug [class_header]: Header for MTRActionsClusterDisableActionWithDurationParams */
// The class instance for the [MTRActionsClusterDisableActionWithDurationParams] class.
var (
	MTRActionsClusterDisableActionWithDurationParamsClass     _MTRActionsClusterDisableActionWithDurationParamsClass
	MTRActionsClusterDisableActionWithDurationParamsClassOnce sync.Once
)

func getMTRActionsClusterDisableActionWithDurationParamsClass() _MTRActionsClusterDisableActionWithDurationParamsClass {
	MTRActionsClusterDisableActionWithDurationParamsClassOnce.Do(func() {
		MTRActionsClusterDisableActionWithDurationParamsClass = _MTRActionsClusterDisableActionWithDurationParamsClass{objc.GetClass("MTRActionsClusterDisableActionWithDurationParams")}
	})
	return MTRActionsClusterDisableActionWithDurationParamsClass
}

type _MTRActionsClusterDisableActionWithDurationParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRActionsClusterDisableActionWithDurationParams */
// An interface definition for the [MTRActionsClusterDisableActionWithDurationParams] class.
type IMTRActionsClusterDisableActionWithDurationParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRActionsClusterDisableActionWithDurationParams */
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

	
/* debug [class_interface_methods]: Methods for MTRActionsClusterDisableActionWithDurationParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRActionsClusterDisableActionWithDurationParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRActionsClusterDisableActionWithDurationParamsClass) Alloc() MTRActionsClusterDisableActionWithDurationParams {
	rv := objc.Send[MTRActionsClusterDisableActionWithDurationParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRActionsClusterDisableActionWithDurationParamsClass) New() MTRActionsClusterDisableActionWithDurationParams {
	rv := objc.Send[MTRActionsClusterDisableActionWithDurationParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRActionsClusterDisableActionWithDurationParams) Init() MTRActionsClusterDisableActionWithDurationParams {
	rv := objc.Send[MTRActionsClusterDisableActionWithDurationParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRActionsClusterDisableActionWithDurationParams) Autorelease() MTRActionsClusterDisableActionWithDurationParams {
	rv := objc.Send[MTRActionsClusterDisableActionWithDurationParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRActionsClusterDisableActionWithDurationParams creates a new MTRActionsClusterDisableActionWithDurationParams instance.
func NewMTRActionsClusterDisableActionWithDurationParams() MTRActionsClusterDisableActionWithDurationParams {
	return getMTRActionsClusterDisableActionWithDurationParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRActionsClusterDisableActionWithDurationParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterDisableActionWithDurationParams
type MTRActionsClusterDisableActionWithDurationParams struct {
	objectivec.Object
}

// MTRActionsClusterDisableActionWithDurationParamsFrom constructs a [MTRActionsClusterDisableActionWithDurationParams] from an unsafe.Pointer.
func MTRActionsClusterDisableActionWithDurationParamsFrom(ptr unsafe.Pointer) MTRActionsClusterDisableActionWithDurationParams {
	return MTRActionsClusterDisableActionWithDurationParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRActionsClusterDisableActionWithDurationParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRActionsClusterDisableActionWithDurationParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRActionsClusterDisableActionWithDurationParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRActionsClusterDisableActionWithDurationParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRActionsClusterDisableActionWithDurationParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterDisableActionWithDurationParams/actionID
func (m_ MTRActionsClusterDisableActionWithDurationParams) ActionID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("actionID"))
	return rv
}/* debug [instance_properties/getter]: actionID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterDisableActionWithDurationParams/actionID
func (m_ MTRActionsClusterDisableActionWithDurationParams) SetActionID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setActionID:"), value)
}/* debug [instance_properties/setter]: actionID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterDisableActionWithDurationParams/duration
func (m_ MTRActionsClusterDisableActionWithDurationParams) Duration() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("duration"))
	return rv
}/* debug [instance_properties/getter]: duration */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterDisableActionWithDurationParams/duration
func (m_ MTRActionsClusterDisableActionWithDurationParams) SetDuration(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDuration:"), value)
}/* debug [instance_properties/setter]: duration */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterDisableActionWithDurationParams/invokeID
func (m_ MTRActionsClusterDisableActionWithDurationParams) InvokeID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("invokeID"))
	return rv
}/* debug [instance_properties/getter]: invokeID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterDisableActionWithDurationParams/invokeID
func (m_ MTRActionsClusterDisableActionWithDurationParams) SetInvokeID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInvokeID:"), value)
}/* debug [instance_properties/setter]: invokeID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterDisableActionWithDurationParams/serverSideProcessingTimeout
func (m_ MTRActionsClusterDisableActionWithDurationParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterDisableActionWithDurationParams/serverSideProcessingTimeout
func (m_ MTRActionsClusterDisableActionWithDurationParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterDisableActionWithDurationParams/timedInvokeTimeoutMs
func (m_ MTRActionsClusterDisableActionWithDurationParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterDisableActionWithDurationParams/timedInvokeTimeoutMs
func (m_ MTRActionsClusterDisableActionWithDurationParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRActionsClusterDisableActionWithDurationParams */



