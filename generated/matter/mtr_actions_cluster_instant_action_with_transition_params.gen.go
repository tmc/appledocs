// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRActionsClusterInstantActionWithTransitionParams */


/* debug [class_header]: Header for MTRActionsClusterInstantActionWithTransitionParams */
// The class instance for the [MTRActionsClusterInstantActionWithTransitionParams] class.
var (
	MTRActionsClusterInstantActionWithTransitionParamsClass     _MTRActionsClusterInstantActionWithTransitionParamsClass
	MTRActionsClusterInstantActionWithTransitionParamsClassOnce sync.Once
)

func getMTRActionsClusterInstantActionWithTransitionParamsClass() _MTRActionsClusterInstantActionWithTransitionParamsClass {
	MTRActionsClusterInstantActionWithTransitionParamsClassOnce.Do(func() {
		MTRActionsClusterInstantActionWithTransitionParamsClass = _MTRActionsClusterInstantActionWithTransitionParamsClass{objc.GetClass("MTRActionsClusterInstantActionWithTransitionParams")}
	})
	return MTRActionsClusterInstantActionWithTransitionParamsClass
}

type _MTRActionsClusterInstantActionWithTransitionParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRActionsClusterInstantActionWithTransitionParams */
// An interface definition for the [MTRActionsClusterInstantActionWithTransitionParams] class.
type IMTRActionsClusterInstantActionWithTransitionParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRActionsClusterInstantActionWithTransitionParams */
	// properties:
	ActionID() objc.IObject /* cross-framework: NSNumber */
	SetActionID(value objc.IObject /* cross-framework: NSNumber */)
	InvokeID() objc.IObject /* cross-framework: NSNumber */
	SetInvokeID(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	TransitionTime() objc.IObject /* cross-framework: NSNumber */
	SetTransitionTime(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRActionsClusterInstantActionWithTransitionParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRActionsClusterInstantActionWithTransitionParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRActionsClusterInstantActionWithTransitionParamsClass) Alloc() MTRActionsClusterInstantActionWithTransitionParams {
	rv := objc.Send[MTRActionsClusterInstantActionWithTransitionParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRActionsClusterInstantActionWithTransitionParamsClass) New() MTRActionsClusterInstantActionWithTransitionParams {
	rv := objc.Send[MTRActionsClusterInstantActionWithTransitionParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRActionsClusterInstantActionWithTransitionParams) Init() MTRActionsClusterInstantActionWithTransitionParams {
	rv := objc.Send[MTRActionsClusterInstantActionWithTransitionParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRActionsClusterInstantActionWithTransitionParams) Autorelease() MTRActionsClusterInstantActionWithTransitionParams {
	rv := objc.Send[MTRActionsClusterInstantActionWithTransitionParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRActionsClusterInstantActionWithTransitionParams creates a new MTRActionsClusterInstantActionWithTransitionParams instance.
func NewMTRActionsClusterInstantActionWithTransitionParams() MTRActionsClusterInstantActionWithTransitionParams {
	return getMTRActionsClusterInstantActionWithTransitionParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRActionsClusterInstantActionWithTransitionParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterInstantActionWithTransitionParams
type MTRActionsClusterInstantActionWithTransitionParams struct {
	objectivec.Object
}

// MTRActionsClusterInstantActionWithTransitionParamsFrom constructs a [MTRActionsClusterInstantActionWithTransitionParams] from an unsafe.Pointer.
func MTRActionsClusterInstantActionWithTransitionParamsFrom(ptr unsafe.Pointer) MTRActionsClusterInstantActionWithTransitionParams {
	return MTRActionsClusterInstantActionWithTransitionParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRActionsClusterInstantActionWithTransitionParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRActionsClusterInstantActionWithTransitionParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRActionsClusterInstantActionWithTransitionParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRActionsClusterInstantActionWithTransitionParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRActionsClusterInstantActionWithTransitionParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterInstantActionWithTransitionParams/actionID
func (m_ MTRActionsClusterInstantActionWithTransitionParams) ActionID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("actionID"))
	return rv
}/* debug [instance_properties/getter]: actionID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterInstantActionWithTransitionParams/actionID
func (m_ MTRActionsClusterInstantActionWithTransitionParams) SetActionID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setActionID:"), value)
}/* debug [instance_properties/setter]: actionID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterInstantActionWithTransitionParams/invokeID
func (m_ MTRActionsClusterInstantActionWithTransitionParams) InvokeID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("invokeID"))
	return rv
}/* debug [instance_properties/getter]: invokeID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterInstantActionWithTransitionParams/invokeID
func (m_ MTRActionsClusterInstantActionWithTransitionParams) SetInvokeID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInvokeID:"), value)
}/* debug [instance_properties/setter]: invokeID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterInstantActionWithTransitionParams/serverSideProcessingTimeout
func (m_ MTRActionsClusterInstantActionWithTransitionParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterInstantActionWithTransitionParams/serverSideProcessingTimeout
func (m_ MTRActionsClusterInstantActionWithTransitionParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterInstantActionWithTransitionParams/timedInvokeTimeoutMs
func (m_ MTRActionsClusterInstantActionWithTransitionParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterInstantActionWithTransitionParams/timedInvokeTimeoutMs
func (m_ MTRActionsClusterInstantActionWithTransitionParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterInstantActionWithTransitionParams/transitionTime
func (m_ MTRActionsClusterInstantActionWithTransitionParams) TransitionTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("transitionTime"))
	return rv
}/* debug [instance_properties/getter]: transitionTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterInstantActionWithTransitionParams/transitionTime
func (m_ MTRActionsClusterInstantActionWithTransitionParams) SetTransitionTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTransitionTime:"), value)
}/* debug [instance_properties/setter]: transitionTime */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRActionsClusterInstantActionWithTransitionParams */



