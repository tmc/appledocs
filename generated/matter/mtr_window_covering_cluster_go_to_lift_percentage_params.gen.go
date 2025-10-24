// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRWindowCoveringClusterGoToLiftPercentageParams */


/* debug [class_header]: Header for MTRWindowCoveringClusterGoToLiftPercentageParams */
// The class instance for the [MTRWindowCoveringClusterGoToLiftPercentageParams] class.
var (
	MTRWindowCoveringClusterGoToLiftPercentageParamsClass     _MTRWindowCoveringClusterGoToLiftPercentageParamsClass
	MTRWindowCoveringClusterGoToLiftPercentageParamsClassOnce sync.Once
)

func getMTRWindowCoveringClusterGoToLiftPercentageParamsClass() _MTRWindowCoveringClusterGoToLiftPercentageParamsClass {
	MTRWindowCoveringClusterGoToLiftPercentageParamsClassOnce.Do(func() {
		MTRWindowCoveringClusterGoToLiftPercentageParamsClass = _MTRWindowCoveringClusterGoToLiftPercentageParamsClass{objc.GetClass("MTRWindowCoveringClusterGoToLiftPercentageParams")}
	})
	return MTRWindowCoveringClusterGoToLiftPercentageParamsClass
}

type _MTRWindowCoveringClusterGoToLiftPercentageParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRWindowCoveringClusterGoToLiftPercentageParams */
// An interface definition for the [MTRWindowCoveringClusterGoToLiftPercentageParams] class.
type IMTRWindowCoveringClusterGoToLiftPercentageParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRWindowCoveringClusterGoToLiftPercentageParams */
	// properties:
	LiftPercent100thsValue() objc.IObject /* cross-framework: NSNumber */
	SetLiftPercent100thsValue(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRWindowCoveringClusterGoToLiftPercentageParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRWindowCoveringClusterGoToLiftPercentageParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRWindowCoveringClusterGoToLiftPercentageParamsClass) Alloc() MTRWindowCoveringClusterGoToLiftPercentageParams {
	rv := objc.Send[MTRWindowCoveringClusterGoToLiftPercentageParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRWindowCoveringClusterGoToLiftPercentageParamsClass) New() MTRWindowCoveringClusterGoToLiftPercentageParams {
	rv := objc.Send[MTRWindowCoveringClusterGoToLiftPercentageParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRWindowCoveringClusterGoToLiftPercentageParams) Init() MTRWindowCoveringClusterGoToLiftPercentageParams {
	rv := objc.Send[MTRWindowCoveringClusterGoToLiftPercentageParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRWindowCoveringClusterGoToLiftPercentageParams) Autorelease() MTRWindowCoveringClusterGoToLiftPercentageParams {
	rv := objc.Send[MTRWindowCoveringClusterGoToLiftPercentageParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRWindowCoveringClusterGoToLiftPercentageParams creates a new MTRWindowCoveringClusterGoToLiftPercentageParams instance.
func NewMTRWindowCoveringClusterGoToLiftPercentageParams() MTRWindowCoveringClusterGoToLiftPercentageParams {
	return getMTRWindowCoveringClusterGoToLiftPercentageParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRWindowCoveringClusterGoToLiftPercentageParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWindowCoveringClusterGoToLiftPercentageParams
type MTRWindowCoveringClusterGoToLiftPercentageParams struct {
	objectivec.Object
}

// MTRWindowCoveringClusterGoToLiftPercentageParamsFrom constructs a [MTRWindowCoveringClusterGoToLiftPercentageParams] from an unsafe.Pointer.
func MTRWindowCoveringClusterGoToLiftPercentageParamsFrom(ptr unsafe.Pointer) MTRWindowCoveringClusterGoToLiftPercentageParams {
	return MTRWindowCoveringClusterGoToLiftPercentageParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRWindowCoveringClusterGoToLiftPercentageParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRWindowCoveringClusterGoToLiftPercentageParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRWindowCoveringClusterGoToLiftPercentageParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRWindowCoveringClusterGoToLiftPercentageParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRWindowCoveringClusterGoToLiftPercentageParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWindowCoveringClusterGoToLiftPercentageParams/liftPercent100thsValue
func (m_ MTRWindowCoveringClusterGoToLiftPercentageParams) LiftPercent100thsValue() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("liftPercent100thsValue"))
	return rv
}/* debug [instance_properties/getter]: liftPercent100thsValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWindowCoveringClusterGoToLiftPercentageParams/liftPercent100thsValue
func (m_ MTRWindowCoveringClusterGoToLiftPercentageParams) SetLiftPercent100thsValue(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLiftPercent100thsValue:"), value)
}/* debug [instance_properties/setter]: liftPercent100thsValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWindowCoveringClusterGoToLiftPercentageParams/serverSideProcessingTimeout
func (m_ MTRWindowCoveringClusterGoToLiftPercentageParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWindowCoveringClusterGoToLiftPercentageParams/serverSideProcessingTimeout
func (m_ MTRWindowCoveringClusterGoToLiftPercentageParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWindowCoveringClusterGoToLiftPercentageParams/timedInvokeTimeoutMs
func (m_ MTRWindowCoveringClusterGoToLiftPercentageParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWindowCoveringClusterGoToLiftPercentageParams/timedInvokeTimeoutMs
func (m_ MTRWindowCoveringClusterGoToLiftPercentageParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRWindowCoveringClusterGoToLiftPercentageParams */



