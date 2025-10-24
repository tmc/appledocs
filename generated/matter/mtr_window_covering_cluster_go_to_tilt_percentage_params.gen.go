// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRWindowCoveringClusterGoToTiltPercentageParams */


/* debug [class_header]: Header for MTRWindowCoveringClusterGoToTiltPercentageParams */
// The class instance for the [MTRWindowCoveringClusterGoToTiltPercentageParams] class.
var (
	MTRWindowCoveringClusterGoToTiltPercentageParamsClass     _MTRWindowCoveringClusterGoToTiltPercentageParamsClass
	MTRWindowCoveringClusterGoToTiltPercentageParamsClassOnce sync.Once
)

func getMTRWindowCoveringClusterGoToTiltPercentageParamsClass() _MTRWindowCoveringClusterGoToTiltPercentageParamsClass {
	MTRWindowCoveringClusterGoToTiltPercentageParamsClassOnce.Do(func() {
		MTRWindowCoveringClusterGoToTiltPercentageParamsClass = _MTRWindowCoveringClusterGoToTiltPercentageParamsClass{objc.GetClass("MTRWindowCoveringClusterGoToTiltPercentageParams")}
	})
	return MTRWindowCoveringClusterGoToTiltPercentageParamsClass
}

type _MTRWindowCoveringClusterGoToTiltPercentageParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRWindowCoveringClusterGoToTiltPercentageParams */
// An interface definition for the [MTRWindowCoveringClusterGoToTiltPercentageParams] class.
type IMTRWindowCoveringClusterGoToTiltPercentageParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRWindowCoveringClusterGoToTiltPercentageParams */
	// properties:
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TiltPercent100thsValue() objc.IObject /* cross-framework: NSNumber */
	SetTiltPercent100thsValue(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRWindowCoveringClusterGoToTiltPercentageParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRWindowCoveringClusterGoToTiltPercentageParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRWindowCoveringClusterGoToTiltPercentageParamsClass) Alloc() MTRWindowCoveringClusterGoToTiltPercentageParams {
	rv := objc.Send[MTRWindowCoveringClusterGoToTiltPercentageParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRWindowCoveringClusterGoToTiltPercentageParamsClass) New() MTRWindowCoveringClusterGoToTiltPercentageParams {
	rv := objc.Send[MTRWindowCoveringClusterGoToTiltPercentageParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRWindowCoveringClusterGoToTiltPercentageParams) Init() MTRWindowCoveringClusterGoToTiltPercentageParams {
	rv := objc.Send[MTRWindowCoveringClusterGoToTiltPercentageParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRWindowCoveringClusterGoToTiltPercentageParams) Autorelease() MTRWindowCoveringClusterGoToTiltPercentageParams {
	rv := objc.Send[MTRWindowCoveringClusterGoToTiltPercentageParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRWindowCoveringClusterGoToTiltPercentageParams creates a new MTRWindowCoveringClusterGoToTiltPercentageParams instance.
func NewMTRWindowCoveringClusterGoToTiltPercentageParams() MTRWindowCoveringClusterGoToTiltPercentageParams {
	return getMTRWindowCoveringClusterGoToTiltPercentageParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRWindowCoveringClusterGoToTiltPercentageParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWindowCoveringClusterGoToTiltPercentageParams
type MTRWindowCoveringClusterGoToTiltPercentageParams struct {
	objectivec.Object
}

// MTRWindowCoveringClusterGoToTiltPercentageParamsFrom constructs a [MTRWindowCoveringClusterGoToTiltPercentageParams] from an unsafe.Pointer.
func MTRWindowCoveringClusterGoToTiltPercentageParamsFrom(ptr unsafe.Pointer) MTRWindowCoveringClusterGoToTiltPercentageParams {
	return MTRWindowCoveringClusterGoToTiltPercentageParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRWindowCoveringClusterGoToTiltPercentageParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRWindowCoveringClusterGoToTiltPercentageParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRWindowCoveringClusterGoToTiltPercentageParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRWindowCoveringClusterGoToTiltPercentageParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRWindowCoveringClusterGoToTiltPercentageParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWindowCoveringClusterGoToTiltPercentageParams/serverSideProcessingTimeout
func (m_ MTRWindowCoveringClusterGoToTiltPercentageParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWindowCoveringClusterGoToTiltPercentageParams/serverSideProcessingTimeout
func (m_ MTRWindowCoveringClusterGoToTiltPercentageParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWindowCoveringClusterGoToTiltPercentageParams/tiltPercent100thsValue
func (m_ MTRWindowCoveringClusterGoToTiltPercentageParams) TiltPercent100thsValue() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("tiltPercent100thsValue"))
	return rv
}/* debug [instance_properties/getter]: tiltPercent100thsValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWindowCoveringClusterGoToTiltPercentageParams/tiltPercent100thsValue
func (m_ MTRWindowCoveringClusterGoToTiltPercentageParams) SetTiltPercent100thsValue(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTiltPercent100thsValue:"), value)
}/* debug [instance_properties/setter]: tiltPercent100thsValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWindowCoveringClusterGoToTiltPercentageParams/timedInvokeTimeoutMs
func (m_ MTRWindowCoveringClusterGoToTiltPercentageParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWindowCoveringClusterGoToTiltPercentageParams/timedInvokeTimeoutMs
func (m_ MTRWindowCoveringClusterGoToTiltPercentageParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRWindowCoveringClusterGoToTiltPercentageParams */



