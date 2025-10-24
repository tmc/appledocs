// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRActionsClusterPauseActionWithDurationParams */


/* debug [class_header]: Header for MTRActionsClusterPauseActionWithDurationParams */
// The class instance for the [MTRActionsClusterPauseActionWithDurationParams] class.
var (
	MTRActionsClusterPauseActionWithDurationParamsClass     _MTRActionsClusterPauseActionWithDurationParamsClass
	MTRActionsClusterPauseActionWithDurationParamsClassOnce sync.Once
)

func getMTRActionsClusterPauseActionWithDurationParamsClass() _MTRActionsClusterPauseActionWithDurationParamsClass {
	MTRActionsClusterPauseActionWithDurationParamsClassOnce.Do(func() {
		MTRActionsClusterPauseActionWithDurationParamsClass = _MTRActionsClusterPauseActionWithDurationParamsClass{objc.GetClass("MTRActionsClusterPauseActionWithDurationParams")}
	})
	return MTRActionsClusterPauseActionWithDurationParamsClass
}

type _MTRActionsClusterPauseActionWithDurationParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRActionsClusterPauseActionWithDurationParams */
// An interface definition for the [MTRActionsClusterPauseActionWithDurationParams] class.
type IMTRActionsClusterPauseActionWithDurationParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRActionsClusterPauseActionWithDurationParams */
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

	
/* debug [class_interface_methods]: Methods for MTRActionsClusterPauseActionWithDurationParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRActionsClusterPauseActionWithDurationParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRActionsClusterPauseActionWithDurationParamsClass) Alloc() MTRActionsClusterPauseActionWithDurationParams {
	rv := objc.Send[MTRActionsClusterPauseActionWithDurationParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRActionsClusterPauseActionWithDurationParamsClass) New() MTRActionsClusterPauseActionWithDurationParams {
	rv := objc.Send[MTRActionsClusterPauseActionWithDurationParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRActionsClusterPauseActionWithDurationParams) Init() MTRActionsClusterPauseActionWithDurationParams {
	rv := objc.Send[MTRActionsClusterPauseActionWithDurationParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRActionsClusterPauseActionWithDurationParams) Autorelease() MTRActionsClusterPauseActionWithDurationParams {
	rv := objc.Send[MTRActionsClusterPauseActionWithDurationParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRActionsClusterPauseActionWithDurationParams creates a new MTRActionsClusterPauseActionWithDurationParams instance.
func NewMTRActionsClusterPauseActionWithDurationParams() MTRActionsClusterPauseActionWithDurationParams {
	return getMTRActionsClusterPauseActionWithDurationParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRActionsClusterPauseActionWithDurationParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterPauseActionWithDurationParams
type MTRActionsClusterPauseActionWithDurationParams struct {
	objectivec.Object
}

// MTRActionsClusterPauseActionWithDurationParamsFrom constructs a [MTRActionsClusterPauseActionWithDurationParams] from an unsafe.Pointer.
func MTRActionsClusterPauseActionWithDurationParamsFrom(ptr unsafe.Pointer) MTRActionsClusterPauseActionWithDurationParams {
	return MTRActionsClusterPauseActionWithDurationParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRActionsClusterPauseActionWithDurationParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRActionsClusterPauseActionWithDurationParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRActionsClusterPauseActionWithDurationParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRActionsClusterPauseActionWithDurationParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRActionsClusterPauseActionWithDurationParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterPauseActionWithDurationParams/actionID
func (m_ MTRActionsClusterPauseActionWithDurationParams) ActionID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("actionID"))
	return rv
}/* debug [instance_properties/getter]: actionID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterPauseActionWithDurationParams/actionID
func (m_ MTRActionsClusterPauseActionWithDurationParams) SetActionID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setActionID:"), value)
}/* debug [instance_properties/setter]: actionID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterPauseActionWithDurationParams/duration
func (m_ MTRActionsClusterPauseActionWithDurationParams) Duration() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("duration"))
	return rv
}/* debug [instance_properties/getter]: duration */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterPauseActionWithDurationParams/duration
func (m_ MTRActionsClusterPauseActionWithDurationParams) SetDuration(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDuration:"), value)
}/* debug [instance_properties/setter]: duration */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterPauseActionWithDurationParams/invokeID
func (m_ MTRActionsClusterPauseActionWithDurationParams) InvokeID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("invokeID"))
	return rv
}/* debug [instance_properties/getter]: invokeID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterPauseActionWithDurationParams/invokeID
func (m_ MTRActionsClusterPauseActionWithDurationParams) SetInvokeID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInvokeID:"), value)
}/* debug [instance_properties/setter]: invokeID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterPauseActionWithDurationParams/serverSideProcessingTimeout
func (m_ MTRActionsClusterPauseActionWithDurationParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterPauseActionWithDurationParams/serverSideProcessingTimeout
func (m_ MTRActionsClusterPauseActionWithDurationParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterPauseActionWithDurationParams/timedInvokeTimeoutMs
func (m_ MTRActionsClusterPauseActionWithDurationParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterPauseActionWithDurationParams/timedInvokeTimeoutMs
func (m_ MTRActionsClusterPauseActionWithDurationParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRActionsClusterPauseActionWithDurationParams */



