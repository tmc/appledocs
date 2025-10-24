// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRActionsClusterDisableActionParams */


/* debug [class_header]: Header for MTRActionsClusterDisableActionParams */
// The class instance for the [MTRActionsClusterDisableActionParams] class.
var (
	MTRActionsClusterDisableActionParamsClass     _MTRActionsClusterDisableActionParamsClass
	MTRActionsClusterDisableActionParamsClassOnce sync.Once
)

func getMTRActionsClusterDisableActionParamsClass() _MTRActionsClusterDisableActionParamsClass {
	MTRActionsClusterDisableActionParamsClassOnce.Do(func() {
		MTRActionsClusterDisableActionParamsClass = _MTRActionsClusterDisableActionParamsClass{objc.GetClass("MTRActionsClusterDisableActionParams")}
	})
	return MTRActionsClusterDisableActionParamsClass
}

type _MTRActionsClusterDisableActionParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRActionsClusterDisableActionParams */
// An interface definition for the [MTRActionsClusterDisableActionParams] class.
type IMTRActionsClusterDisableActionParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRActionsClusterDisableActionParams */
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

	
/* debug [class_interface_methods]: Methods for MTRActionsClusterDisableActionParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRActionsClusterDisableActionParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRActionsClusterDisableActionParamsClass) Alloc() MTRActionsClusterDisableActionParams {
	rv := objc.Send[MTRActionsClusterDisableActionParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRActionsClusterDisableActionParamsClass) New() MTRActionsClusterDisableActionParams {
	rv := objc.Send[MTRActionsClusterDisableActionParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRActionsClusterDisableActionParams) Init() MTRActionsClusterDisableActionParams {
	rv := objc.Send[MTRActionsClusterDisableActionParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRActionsClusterDisableActionParams) Autorelease() MTRActionsClusterDisableActionParams {
	rv := objc.Send[MTRActionsClusterDisableActionParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRActionsClusterDisableActionParams creates a new MTRActionsClusterDisableActionParams instance.
func NewMTRActionsClusterDisableActionParams() MTRActionsClusterDisableActionParams {
	return getMTRActionsClusterDisableActionParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRActionsClusterDisableActionParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterDisableActionParams
type MTRActionsClusterDisableActionParams struct {
	objectivec.Object
}

// MTRActionsClusterDisableActionParamsFrom constructs a [MTRActionsClusterDisableActionParams] from an unsafe.Pointer.
func MTRActionsClusterDisableActionParamsFrom(ptr unsafe.Pointer) MTRActionsClusterDisableActionParams {
	return MTRActionsClusterDisableActionParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRActionsClusterDisableActionParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRActionsClusterDisableActionParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRActionsClusterDisableActionParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRActionsClusterDisableActionParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRActionsClusterDisableActionParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterDisableActionParams/actionID
func (m_ MTRActionsClusterDisableActionParams) ActionID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("actionID"))
	return rv
}/* debug [instance_properties/getter]: actionID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterDisableActionParams/actionID
func (m_ MTRActionsClusterDisableActionParams) SetActionID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setActionID:"), value)
}/* debug [instance_properties/setter]: actionID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterDisableActionParams/invokeID
func (m_ MTRActionsClusterDisableActionParams) InvokeID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("invokeID"))
	return rv
}/* debug [instance_properties/getter]: invokeID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterDisableActionParams/invokeID
func (m_ MTRActionsClusterDisableActionParams) SetInvokeID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInvokeID:"), value)
}/* debug [instance_properties/setter]: invokeID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterDisableActionParams/serverSideProcessingTimeout
func (m_ MTRActionsClusterDisableActionParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterDisableActionParams/serverSideProcessingTimeout
func (m_ MTRActionsClusterDisableActionParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterDisableActionParams/timedInvokeTimeoutMs
func (m_ MTRActionsClusterDisableActionParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterDisableActionParams/timedInvokeTimeoutMs
func (m_ MTRActionsClusterDisableActionParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRActionsClusterDisableActionParams */



