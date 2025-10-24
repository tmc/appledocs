// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRActionsClusterPauseActionParams */


/* debug [class_header]: Header for MTRActionsClusterPauseActionParams */
// The class instance for the [MTRActionsClusterPauseActionParams] class.
var (
	MTRActionsClusterPauseActionParamsClass     _MTRActionsClusterPauseActionParamsClass
	MTRActionsClusterPauseActionParamsClassOnce sync.Once
)

func getMTRActionsClusterPauseActionParamsClass() _MTRActionsClusterPauseActionParamsClass {
	MTRActionsClusterPauseActionParamsClassOnce.Do(func() {
		MTRActionsClusterPauseActionParamsClass = _MTRActionsClusterPauseActionParamsClass{objc.GetClass("MTRActionsClusterPauseActionParams")}
	})
	return MTRActionsClusterPauseActionParamsClass
}

type _MTRActionsClusterPauseActionParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRActionsClusterPauseActionParams */
// An interface definition for the [MTRActionsClusterPauseActionParams] class.
type IMTRActionsClusterPauseActionParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRActionsClusterPauseActionParams */
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

	
/* debug [class_interface_methods]: Methods for MTRActionsClusterPauseActionParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRActionsClusterPauseActionParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRActionsClusterPauseActionParamsClass) Alloc() MTRActionsClusterPauseActionParams {
	rv := objc.Send[MTRActionsClusterPauseActionParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRActionsClusterPauseActionParamsClass) New() MTRActionsClusterPauseActionParams {
	rv := objc.Send[MTRActionsClusterPauseActionParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRActionsClusterPauseActionParams) Init() MTRActionsClusterPauseActionParams {
	rv := objc.Send[MTRActionsClusterPauseActionParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRActionsClusterPauseActionParams) Autorelease() MTRActionsClusterPauseActionParams {
	rv := objc.Send[MTRActionsClusterPauseActionParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRActionsClusterPauseActionParams creates a new MTRActionsClusterPauseActionParams instance.
func NewMTRActionsClusterPauseActionParams() MTRActionsClusterPauseActionParams {
	return getMTRActionsClusterPauseActionParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRActionsClusterPauseActionParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterPauseActionParams
type MTRActionsClusterPauseActionParams struct {
	objectivec.Object
}

// MTRActionsClusterPauseActionParamsFrom constructs a [MTRActionsClusterPauseActionParams] from an unsafe.Pointer.
func MTRActionsClusterPauseActionParamsFrom(ptr unsafe.Pointer) MTRActionsClusterPauseActionParams {
	return MTRActionsClusterPauseActionParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRActionsClusterPauseActionParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRActionsClusterPauseActionParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRActionsClusterPauseActionParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRActionsClusterPauseActionParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRActionsClusterPauseActionParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterPauseActionParams/actionID
func (m_ MTRActionsClusterPauseActionParams) ActionID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("actionID"))
	return rv
}/* debug [instance_properties/getter]: actionID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterPauseActionParams/actionID
func (m_ MTRActionsClusterPauseActionParams) SetActionID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setActionID:"), value)
}/* debug [instance_properties/setter]: actionID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterPauseActionParams/invokeID
func (m_ MTRActionsClusterPauseActionParams) InvokeID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("invokeID"))
	return rv
}/* debug [instance_properties/getter]: invokeID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterPauseActionParams/invokeID
func (m_ MTRActionsClusterPauseActionParams) SetInvokeID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInvokeID:"), value)
}/* debug [instance_properties/setter]: invokeID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterPauseActionParams/serverSideProcessingTimeout
func (m_ MTRActionsClusterPauseActionParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterPauseActionParams/serverSideProcessingTimeout
func (m_ MTRActionsClusterPauseActionParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterPauseActionParams/timedInvokeTimeoutMs
func (m_ MTRActionsClusterPauseActionParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterPauseActionParams/timedInvokeTimeoutMs
func (m_ MTRActionsClusterPauseActionParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRActionsClusterPauseActionParams */



