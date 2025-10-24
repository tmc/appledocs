// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRWindowCoveringClusterGoToLiftValueParams */


/* debug [class_header]: Header for MTRWindowCoveringClusterGoToLiftValueParams */
// The class instance for the [MTRWindowCoveringClusterGoToLiftValueParams] class.
var (
	MTRWindowCoveringClusterGoToLiftValueParamsClass     _MTRWindowCoveringClusterGoToLiftValueParamsClass
	MTRWindowCoveringClusterGoToLiftValueParamsClassOnce sync.Once
)

func getMTRWindowCoveringClusterGoToLiftValueParamsClass() _MTRWindowCoveringClusterGoToLiftValueParamsClass {
	MTRWindowCoveringClusterGoToLiftValueParamsClassOnce.Do(func() {
		MTRWindowCoveringClusterGoToLiftValueParamsClass = _MTRWindowCoveringClusterGoToLiftValueParamsClass{objc.GetClass("MTRWindowCoveringClusterGoToLiftValueParams")}
	})
	return MTRWindowCoveringClusterGoToLiftValueParamsClass
}

type _MTRWindowCoveringClusterGoToLiftValueParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRWindowCoveringClusterGoToLiftValueParams */
// An interface definition for the [MTRWindowCoveringClusterGoToLiftValueParams] class.
type IMTRWindowCoveringClusterGoToLiftValueParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRWindowCoveringClusterGoToLiftValueParams */
	// properties:
	LiftValue() objc.IObject /* cross-framework: NSNumber */
	SetLiftValue(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRWindowCoveringClusterGoToLiftValueParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRWindowCoveringClusterGoToLiftValueParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRWindowCoveringClusterGoToLiftValueParamsClass) Alloc() MTRWindowCoveringClusterGoToLiftValueParams {
	rv := objc.Send[MTRWindowCoveringClusterGoToLiftValueParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRWindowCoveringClusterGoToLiftValueParamsClass) New() MTRWindowCoveringClusterGoToLiftValueParams {
	rv := objc.Send[MTRWindowCoveringClusterGoToLiftValueParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRWindowCoveringClusterGoToLiftValueParams) Init() MTRWindowCoveringClusterGoToLiftValueParams {
	rv := objc.Send[MTRWindowCoveringClusterGoToLiftValueParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRWindowCoveringClusterGoToLiftValueParams) Autorelease() MTRWindowCoveringClusterGoToLiftValueParams {
	rv := objc.Send[MTRWindowCoveringClusterGoToLiftValueParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRWindowCoveringClusterGoToLiftValueParams creates a new MTRWindowCoveringClusterGoToLiftValueParams instance.
func NewMTRWindowCoveringClusterGoToLiftValueParams() MTRWindowCoveringClusterGoToLiftValueParams {
	return getMTRWindowCoveringClusterGoToLiftValueParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRWindowCoveringClusterGoToLiftValueParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWindowCoveringClusterGoToLiftValueParams
type MTRWindowCoveringClusterGoToLiftValueParams struct {
	objectivec.Object
}

// MTRWindowCoveringClusterGoToLiftValueParamsFrom constructs a [MTRWindowCoveringClusterGoToLiftValueParams] from an unsafe.Pointer.
func MTRWindowCoveringClusterGoToLiftValueParamsFrom(ptr unsafe.Pointer) MTRWindowCoveringClusterGoToLiftValueParams {
	return MTRWindowCoveringClusterGoToLiftValueParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRWindowCoveringClusterGoToLiftValueParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRWindowCoveringClusterGoToLiftValueParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRWindowCoveringClusterGoToLiftValueParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRWindowCoveringClusterGoToLiftValueParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRWindowCoveringClusterGoToLiftValueParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWindowCoveringClusterGoToLiftValueParams/liftValue
func (m_ MTRWindowCoveringClusterGoToLiftValueParams) LiftValue() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("liftValue"))
	return rv
}/* debug [instance_properties/getter]: liftValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWindowCoveringClusterGoToLiftValueParams/liftValue
func (m_ MTRWindowCoveringClusterGoToLiftValueParams) SetLiftValue(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLiftValue:"), value)
}/* debug [instance_properties/setter]: liftValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWindowCoveringClusterGoToLiftValueParams/serverSideProcessingTimeout
func (m_ MTRWindowCoveringClusterGoToLiftValueParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWindowCoveringClusterGoToLiftValueParams/serverSideProcessingTimeout
func (m_ MTRWindowCoveringClusterGoToLiftValueParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWindowCoveringClusterGoToLiftValueParams/timedInvokeTimeoutMs
func (m_ MTRWindowCoveringClusterGoToLiftValueParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWindowCoveringClusterGoToLiftValueParams/timedInvokeTimeoutMs
func (m_ MTRWindowCoveringClusterGoToLiftValueParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRWindowCoveringClusterGoToLiftValueParams */



