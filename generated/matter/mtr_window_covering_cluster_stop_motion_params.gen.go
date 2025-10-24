// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRWindowCoveringClusterStopMotionParams */


/* debug [class_header]: Header for MTRWindowCoveringClusterStopMotionParams */
// The class instance for the [MTRWindowCoveringClusterStopMotionParams] class.
var (
	MTRWindowCoveringClusterStopMotionParamsClass     _MTRWindowCoveringClusterStopMotionParamsClass
	MTRWindowCoveringClusterStopMotionParamsClassOnce sync.Once
)

func getMTRWindowCoveringClusterStopMotionParamsClass() _MTRWindowCoveringClusterStopMotionParamsClass {
	MTRWindowCoveringClusterStopMotionParamsClassOnce.Do(func() {
		MTRWindowCoveringClusterStopMotionParamsClass = _MTRWindowCoveringClusterStopMotionParamsClass{objc.GetClass("MTRWindowCoveringClusterStopMotionParams")}
	})
	return MTRWindowCoveringClusterStopMotionParamsClass
}

type _MTRWindowCoveringClusterStopMotionParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRWindowCoveringClusterStopMotionParams */
// An interface definition for the [MTRWindowCoveringClusterStopMotionParams] class.
type IMTRWindowCoveringClusterStopMotionParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRWindowCoveringClusterStopMotionParams */
	// properties:
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRWindowCoveringClusterStopMotionParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRWindowCoveringClusterStopMotionParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRWindowCoveringClusterStopMotionParamsClass) Alloc() MTRWindowCoveringClusterStopMotionParams {
	rv := objc.Send[MTRWindowCoveringClusterStopMotionParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRWindowCoveringClusterStopMotionParamsClass) New() MTRWindowCoveringClusterStopMotionParams {
	rv := objc.Send[MTRWindowCoveringClusterStopMotionParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRWindowCoveringClusterStopMotionParams) Init() MTRWindowCoveringClusterStopMotionParams {
	rv := objc.Send[MTRWindowCoveringClusterStopMotionParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRWindowCoveringClusterStopMotionParams) Autorelease() MTRWindowCoveringClusterStopMotionParams {
	rv := objc.Send[MTRWindowCoveringClusterStopMotionParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRWindowCoveringClusterStopMotionParams creates a new MTRWindowCoveringClusterStopMotionParams instance.
func NewMTRWindowCoveringClusterStopMotionParams() MTRWindowCoveringClusterStopMotionParams {
	return getMTRWindowCoveringClusterStopMotionParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRWindowCoveringClusterStopMotionParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWindowCoveringClusterStopMotionParams
type MTRWindowCoveringClusterStopMotionParams struct {
	objectivec.Object
}

// MTRWindowCoveringClusterStopMotionParamsFrom constructs a [MTRWindowCoveringClusterStopMotionParams] from an unsafe.Pointer.
func MTRWindowCoveringClusterStopMotionParamsFrom(ptr unsafe.Pointer) MTRWindowCoveringClusterStopMotionParams {
	return MTRWindowCoveringClusterStopMotionParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRWindowCoveringClusterStopMotionParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRWindowCoveringClusterStopMotionParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRWindowCoveringClusterStopMotionParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRWindowCoveringClusterStopMotionParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRWindowCoveringClusterStopMotionParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWindowCoveringClusterStopMotionParams/serverSideProcessingTimeout
func (m_ MTRWindowCoveringClusterStopMotionParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWindowCoveringClusterStopMotionParams/serverSideProcessingTimeout
func (m_ MTRWindowCoveringClusterStopMotionParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWindowCoveringClusterStopMotionParams/timedInvokeTimeoutMs
func (m_ MTRWindowCoveringClusterStopMotionParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWindowCoveringClusterStopMotionParams/timedInvokeTimeoutMs
func (m_ MTRWindowCoveringClusterStopMotionParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRWindowCoveringClusterStopMotionParams */



