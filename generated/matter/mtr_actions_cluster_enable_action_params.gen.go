// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRActionsClusterEnableActionParams */


/* debug [class_header]: Header for MTRActionsClusterEnableActionParams */
// The class instance for the [MTRActionsClusterEnableActionParams] class.
var (
	MTRActionsClusterEnableActionParamsClass     _MTRActionsClusterEnableActionParamsClass
	MTRActionsClusterEnableActionParamsClassOnce sync.Once
)

func getMTRActionsClusterEnableActionParamsClass() _MTRActionsClusterEnableActionParamsClass {
	MTRActionsClusterEnableActionParamsClassOnce.Do(func() {
		MTRActionsClusterEnableActionParamsClass = _MTRActionsClusterEnableActionParamsClass{objc.GetClass("MTRActionsClusterEnableActionParams")}
	})
	return MTRActionsClusterEnableActionParamsClass
}

type _MTRActionsClusterEnableActionParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRActionsClusterEnableActionParams */
// An interface definition for the [MTRActionsClusterEnableActionParams] class.
type IMTRActionsClusterEnableActionParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRActionsClusterEnableActionParams */
	// properties:
	ActionID() objc.IObject /* cross-framework: NSNumber */
	SetActionID(value objc.IObject /* cross-framework: NSNumber */)
	InvokeID() objc.IObject /* cross-framework: NSNumber */
	SetInvokeID(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRActionsClusterEnableActionParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRActionsClusterEnableActionParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRActionsClusterEnableActionParamsClass) Alloc() MTRActionsClusterEnableActionParams {
	rv := objc.Send[MTRActionsClusterEnableActionParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRActionsClusterEnableActionParamsClass) New() MTRActionsClusterEnableActionParams {
	rv := objc.Send[MTRActionsClusterEnableActionParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRActionsClusterEnableActionParams) Init() MTRActionsClusterEnableActionParams {
	rv := objc.Send[MTRActionsClusterEnableActionParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRActionsClusterEnableActionParams) Autorelease() MTRActionsClusterEnableActionParams {
	rv := objc.Send[MTRActionsClusterEnableActionParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRActionsClusterEnableActionParams creates a new MTRActionsClusterEnableActionParams instance.
func NewMTRActionsClusterEnableActionParams() MTRActionsClusterEnableActionParams {
	return getMTRActionsClusterEnableActionParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRActionsClusterEnableActionParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterEnableActionParams
type MTRActionsClusterEnableActionParams struct {
	objectivec.Object
}

// MTRActionsClusterEnableActionParamsFrom constructs a [MTRActionsClusterEnableActionParams] from an unsafe.Pointer.
func MTRActionsClusterEnableActionParamsFrom(ptr unsafe.Pointer) MTRActionsClusterEnableActionParams {
	return MTRActionsClusterEnableActionParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRActionsClusterEnableActionParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRActionsClusterEnableActionParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRActionsClusterEnableActionParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRActionsClusterEnableActionParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRActionsClusterEnableActionParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterEnableActionParams/actionID
func (m_ MTRActionsClusterEnableActionParams) ActionID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("actionID"))
	return rv
}/* debug [instance_properties/getter]: actionID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterEnableActionParams/actionID
func (m_ MTRActionsClusterEnableActionParams) SetActionID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setActionID:"), value)
}/* debug [instance_properties/setter]: actionID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterEnableActionParams/invokeID
func (m_ MTRActionsClusterEnableActionParams) InvokeID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("invokeID"))
	return rv
}/* debug [instance_properties/getter]: invokeID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterEnableActionParams/invokeID
func (m_ MTRActionsClusterEnableActionParams) SetInvokeID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInvokeID:"), value)
}/* debug [instance_properties/setter]: invokeID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterEnableActionParams/serverSideProcessingTimeout
func (m_ MTRActionsClusterEnableActionParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterEnableActionParams/serverSideProcessingTimeout
func (m_ MTRActionsClusterEnableActionParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterEnableActionParams/timedInvokeTimeoutMs
func (m_ MTRActionsClusterEnableActionParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterEnableActionParams/timedInvokeTimeoutMs
func (m_ MTRActionsClusterEnableActionParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRActionsClusterEnableActionParams */



