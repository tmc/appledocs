// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRActionsClusterResumeActionParams */


/* debug [class_header]: Header for MTRActionsClusterResumeActionParams */
// The class instance for the [MTRActionsClusterResumeActionParams] class.
var (
	MTRActionsClusterResumeActionParamsClass     _MTRActionsClusterResumeActionParamsClass
	MTRActionsClusterResumeActionParamsClassOnce sync.Once
)

func getMTRActionsClusterResumeActionParamsClass() _MTRActionsClusterResumeActionParamsClass {
	MTRActionsClusterResumeActionParamsClassOnce.Do(func() {
		MTRActionsClusterResumeActionParamsClass = _MTRActionsClusterResumeActionParamsClass{objc.GetClass("MTRActionsClusterResumeActionParams")}
	})
	return MTRActionsClusterResumeActionParamsClass
}

type _MTRActionsClusterResumeActionParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRActionsClusterResumeActionParams */
// An interface definition for the [MTRActionsClusterResumeActionParams] class.
type IMTRActionsClusterResumeActionParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRActionsClusterResumeActionParams */
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

	
/* debug [class_interface_methods]: Methods for MTRActionsClusterResumeActionParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRActionsClusterResumeActionParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRActionsClusterResumeActionParamsClass) Alloc() MTRActionsClusterResumeActionParams {
	rv := objc.Send[MTRActionsClusterResumeActionParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRActionsClusterResumeActionParamsClass) New() MTRActionsClusterResumeActionParams {
	rv := objc.Send[MTRActionsClusterResumeActionParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRActionsClusterResumeActionParams) Init() MTRActionsClusterResumeActionParams {
	rv := objc.Send[MTRActionsClusterResumeActionParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRActionsClusterResumeActionParams) Autorelease() MTRActionsClusterResumeActionParams {
	rv := objc.Send[MTRActionsClusterResumeActionParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRActionsClusterResumeActionParams creates a new MTRActionsClusterResumeActionParams instance.
func NewMTRActionsClusterResumeActionParams() MTRActionsClusterResumeActionParams {
	return getMTRActionsClusterResumeActionParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRActionsClusterResumeActionParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterResumeActionParams
type MTRActionsClusterResumeActionParams struct {
	objectivec.Object
}

// MTRActionsClusterResumeActionParamsFrom constructs a [MTRActionsClusterResumeActionParams] from an unsafe.Pointer.
func MTRActionsClusterResumeActionParamsFrom(ptr unsafe.Pointer) MTRActionsClusterResumeActionParams {
	return MTRActionsClusterResumeActionParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRActionsClusterResumeActionParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRActionsClusterResumeActionParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRActionsClusterResumeActionParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRActionsClusterResumeActionParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRActionsClusterResumeActionParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterResumeActionParams/actionID
func (m_ MTRActionsClusterResumeActionParams) ActionID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("actionID"))
	return rv
}/* debug [instance_properties/getter]: actionID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterResumeActionParams/actionID
func (m_ MTRActionsClusterResumeActionParams) SetActionID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setActionID:"), value)
}/* debug [instance_properties/setter]: actionID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterResumeActionParams/invokeID
func (m_ MTRActionsClusterResumeActionParams) InvokeID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("invokeID"))
	return rv
}/* debug [instance_properties/getter]: invokeID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterResumeActionParams/invokeID
func (m_ MTRActionsClusterResumeActionParams) SetInvokeID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInvokeID:"), value)
}/* debug [instance_properties/setter]: invokeID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterResumeActionParams/serverSideProcessingTimeout
func (m_ MTRActionsClusterResumeActionParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterResumeActionParams/serverSideProcessingTimeout
func (m_ MTRActionsClusterResumeActionParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterResumeActionParams/timedInvokeTimeoutMs
func (m_ MTRActionsClusterResumeActionParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterResumeActionParams/timedInvokeTimeoutMs
func (m_ MTRActionsClusterResumeActionParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRActionsClusterResumeActionParams */



