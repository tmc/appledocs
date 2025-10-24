// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRWindowCoveringClusterUpOrOpenParams */


/* debug [class_header]: Header for MTRWindowCoveringClusterUpOrOpenParams */
// The class instance for the [MTRWindowCoveringClusterUpOrOpenParams] class.
var (
	MTRWindowCoveringClusterUpOrOpenParamsClass     _MTRWindowCoveringClusterUpOrOpenParamsClass
	MTRWindowCoveringClusterUpOrOpenParamsClassOnce sync.Once
)

func getMTRWindowCoveringClusterUpOrOpenParamsClass() _MTRWindowCoveringClusterUpOrOpenParamsClass {
	MTRWindowCoveringClusterUpOrOpenParamsClassOnce.Do(func() {
		MTRWindowCoveringClusterUpOrOpenParamsClass = _MTRWindowCoveringClusterUpOrOpenParamsClass{objc.GetClass("MTRWindowCoveringClusterUpOrOpenParams")}
	})
	return MTRWindowCoveringClusterUpOrOpenParamsClass
}

type _MTRWindowCoveringClusterUpOrOpenParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRWindowCoveringClusterUpOrOpenParams */
// An interface definition for the [MTRWindowCoveringClusterUpOrOpenParams] class.
type IMTRWindowCoveringClusterUpOrOpenParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRWindowCoveringClusterUpOrOpenParams */
	// properties:
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRWindowCoveringClusterUpOrOpenParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRWindowCoveringClusterUpOrOpenParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRWindowCoveringClusterUpOrOpenParamsClass) Alloc() MTRWindowCoveringClusterUpOrOpenParams {
	rv := objc.Send[MTRWindowCoveringClusterUpOrOpenParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRWindowCoveringClusterUpOrOpenParamsClass) New() MTRWindowCoveringClusterUpOrOpenParams {
	rv := objc.Send[MTRWindowCoveringClusterUpOrOpenParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRWindowCoveringClusterUpOrOpenParams) Init() MTRWindowCoveringClusterUpOrOpenParams {
	rv := objc.Send[MTRWindowCoveringClusterUpOrOpenParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRWindowCoveringClusterUpOrOpenParams) Autorelease() MTRWindowCoveringClusterUpOrOpenParams {
	rv := objc.Send[MTRWindowCoveringClusterUpOrOpenParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRWindowCoveringClusterUpOrOpenParams creates a new MTRWindowCoveringClusterUpOrOpenParams instance.
func NewMTRWindowCoveringClusterUpOrOpenParams() MTRWindowCoveringClusterUpOrOpenParams {
	return getMTRWindowCoveringClusterUpOrOpenParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRWindowCoveringClusterUpOrOpenParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWindowCoveringClusterUpOrOpenParams
type MTRWindowCoveringClusterUpOrOpenParams struct {
	objectivec.Object
}

// MTRWindowCoveringClusterUpOrOpenParamsFrom constructs a [MTRWindowCoveringClusterUpOrOpenParams] from an unsafe.Pointer.
func MTRWindowCoveringClusterUpOrOpenParamsFrom(ptr unsafe.Pointer) MTRWindowCoveringClusterUpOrOpenParams {
	return MTRWindowCoveringClusterUpOrOpenParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRWindowCoveringClusterUpOrOpenParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRWindowCoveringClusterUpOrOpenParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRWindowCoveringClusterUpOrOpenParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRWindowCoveringClusterUpOrOpenParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRWindowCoveringClusterUpOrOpenParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWindowCoveringClusterUpOrOpenParams/serverSideProcessingTimeout
func (m_ MTRWindowCoveringClusterUpOrOpenParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWindowCoveringClusterUpOrOpenParams/serverSideProcessingTimeout
func (m_ MTRWindowCoveringClusterUpOrOpenParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWindowCoveringClusterUpOrOpenParams/timedInvokeTimeoutMs
func (m_ MTRWindowCoveringClusterUpOrOpenParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWindowCoveringClusterUpOrOpenParams/timedInvokeTimeoutMs
func (m_ MTRWindowCoveringClusterUpOrOpenParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRWindowCoveringClusterUpOrOpenParams */



