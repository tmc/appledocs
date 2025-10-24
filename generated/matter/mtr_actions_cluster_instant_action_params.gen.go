// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRActionsClusterInstantActionParams */


/* debug [class_header]: Header for MTRActionsClusterInstantActionParams */
// The class instance for the [MTRActionsClusterInstantActionParams] class.
var (
	MTRActionsClusterInstantActionParamsClass     _MTRActionsClusterInstantActionParamsClass
	MTRActionsClusterInstantActionParamsClassOnce sync.Once
)

func getMTRActionsClusterInstantActionParamsClass() _MTRActionsClusterInstantActionParamsClass {
	MTRActionsClusterInstantActionParamsClassOnce.Do(func() {
		MTRActionsClusterInstantActionParamsClass = _MTRActionsClusterInstantActionParamsClass{objc.GetClass("MTRActionsClusterInstantActionParams")}
	})
	return MTRActionsClusterInstantActionParamsClass
}

type _MTRActionsClusterInstantActionParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRActionsClusterInstantActionParams */
// An interface definition for the [MTRActionsClusterInstantActionParams] class.
type IMTRActionsClusterInstantActionParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRActionsClusterInstantActionParams */
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

	
/* debug [class_interface_methods]: Methods for MTRActionsClusterInstantActionParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRActionsClusterInstantActionParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRActionsClusterInstantActionParamsClass) Alloc() MTRActionsClusterInstantActionParams {
	rv := objc.Send[MTRActionsClusterInstantActionParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRActionsClusterInstantActionParamsClass) New() MTRActionsClusterInstantActionParams {
	rv := objc.Send[MTRActionsClusterInstantActionParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRActionsClusterInstantActionParams) Init() MTRActionsClusterInstantActionParams {
	rv := objc.Send[MTRActionsClusterInstantActionParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRActionsClusterInstantActionParams) Autorelease() MTRActionsClusterInstantActionParams {
	rv := objc.Send[MTRActionsClusterInstantActionParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRActionsClusterInstantActionParams creates a new MTRActionsClusterInstantActionParams instance.
func NewMTRActionsClusterInstantActionParams() MTRActionsClusterInstantActionParams {
	return getMTRActionsClusterInstantActionParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRActionsClusterInstantActionParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterInstantActionParams
type MTRActionsClusterInstantActionParams struct {
	objectivec.Object
}

// MTRActionsClusterInstantActionParamsFrom constructs a [MTRActionsClusterInstantActionParams] from an unsafe.Pointer.
func MTRActionsClusterInstantActionParamsFrom(ptr unsafe.Pointer) MTRActionsClusterInstantActionParams {
	return MTRActionsClusterInstantActionParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRActionsClusterInstantActionParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRActionsClusterInstantActionParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRActionsClusterInstantActionParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRActionsClusterInstantActionParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRActionsClusterInstantActionParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterInstantActionParams/actionID
func (m_ MTRActionsClusterInstantActionParams) ActionID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("actionID"))
	return rv
}/* debug [instance_properties/getter]: actionID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterInstantActionParams/actionID
func (m_ MTRActionsClusterInstantActionParams) SetActionID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setActionID:"), value)
}/* debug [instance_properties/setter]: actionID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterInstantActionParams/invokeID
func (m_ MTRActionsClusterInstantActionParams) InvokeID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("invokeID"))
	return rv
}/* debug [instance_properties/getter]: invokeID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterInstantActionParams/invokeID
func (m_ MTRActionsClusterInstantActionParams) SetInvokeID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInvokeID:"), value)
}/* debug [instance_properties/setter]: invokeID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterInstantActionParams/serverSideProcessingTimeout
func (m_ MTRActionsClusterInstantActionParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterInstantActionParams/serverSideProcessingTimeout
func (m_ MTRActionsClusterInstantActionParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterInstantActionParams/timedInvokeTimeoutMs
func (m_ MTRActionsClusterInstantActionParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterInstantActionParams/timedInvokeTimeoutMs
func (m_ MTRActionsClusterInstantActionParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRActionsClusterInstantActionParams */



