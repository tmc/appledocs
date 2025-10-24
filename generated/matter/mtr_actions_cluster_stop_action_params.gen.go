// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRActionsClusterStopActionParams */


/* debug [class_header]: Header for MTRActionsClusterStopActionParams */
// The class instance for the [MTRActionsClusterStopActionParams] class.
var (
	MTRActionsClusterStopActionParamsClass     _MTRActionsClusterStopActionParamsClass
	MTRActionsClusterStopActionParamsClassOnce sync.Once
)

func getMTRActionsClusterStopActionParamsClass() _MTRActionsClusterStopActionParamsClass {
	MTRActionsClusterStopActionParamsClassOnce.Do(func() {
		MTRActionsClusterStopActionParamsClass = _MTRActionsClusterStopActionParamsClass{objc.GetClass("MTRActionsClusterStopActionParams")}
	})
	return MTRActionsClusterStopActionParamsClass
}

type _MTRActionsClusterStopActionParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRActionsClusterStopActionParams */
// An interface definition for the [MTRActionsClusterStopActionParams] class.
type IMTRActionsClusterStopActionParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRActionsClusterStopActionParams */
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

	
/* debug [class_interface_methods]: Methods for MTRActionsClusterStopActionParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRActionsClusterStopActionParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRActionsClusterStopActionParamsClass) Alloc() MTRActionsClusterStopActionParams {
	rv := objc.Send[MTRActionsClusterStopActionParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRActionsClusterStopActionParamsClass) New() MTRActionsClusterStopActionParams {
	rv := objc.Send[MTRActionsClusterStopActionParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRActionsClusterStopActionParams) Init() MTRActionsClusterStopActionParams {
	rv := objc.Send[MTRActionsClusterStopActionParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRActionsClusterStopActionParams) Autorelease() MTRActionsClusterStopActionParams {
	rv := objc.Send[MTRActionsClusterStopActionParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRActionsClusterStopActionParams creates a new MTRActionsClusterStopActionParams instance.
func NewMTRActionsClusterStopActionParams() MTRActionsClusterStopActionParams {
	return getMTRActionsClusterStopActionParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRActionsClusterStopActionParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterStopActionParams
type MTRActionsClusterStopActionParams struct {
	objectivec.Object
}

// MTRActionsClusterStopActionParamsFrom constructs a [MTRActionsClusterStopActionParams] from an unsafe.Pointer.
func MTRActionsClusterStopActionParamsFrom(ptr unsafe.Pointer) MTRActionsClusterStopActionParams {
	return MTRActionsClusterStopActionParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRActionsClusterStopActionParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRActionsClusterStopActionParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRActionsClusterStopActionParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRActionsClusterStopActionParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRActionsClusterStopActionParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterStopActionParams/actionID
func (m_ MTRActionsClusterStopActionParams) ActionID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("actionID"))
	return rv
}/* debug [instance_properties/getter]: actionID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterStopActionParams/actionID
func (m_ MTRActionsClusterStopActionParams) SetActionID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setActionID:"), value)
}/* debug [instance_properties/setter]: actionID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterStopActionParams/invokeID
func (m_ MTRActionsClusterStopActionParams) InvokeID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("invokeID"))
	return rv
}/* debug [instance_properties/getter]: invokeID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterStopActionParams/invokeID
func (m_ MTRActionsClusterStopActionParams) SetInvokeID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInvokeID:"), value)
}/* debug [instance_properties/setter]: invokeID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterStopActionParams/serverSideProcessingTimeout
func (m_ MTRActionsClusterStopActionParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterStopActionParams/serverSideProcessingTimeout
func (m_ MTRActionsClusterStopActionParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterStopActionParams/timedInvokeTimeoutMs
func (m_ MTRActionsClusterStopActionParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterStopActionParams/timedInvokeTimeoutMs
func (m_ MTRActionsClusterStopActionParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRActionsClusterStopActionParams */



