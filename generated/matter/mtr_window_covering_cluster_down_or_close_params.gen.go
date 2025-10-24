// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRWindowCoveringClusterDownOrCloseParams */


/* debug [class_header]: Header for MTRWindowCoveringClusterDownOrCloseParams */
// The class instance for the [MTRWindowCoveringClusterDownOrCloseParams] class.
var (
	MTRWindowCoveringClusterDownOrCloseParamsClass     _MTRWindowCoveringClusterDownOrCloseParamsClass
	MTRWindowCoveringClusterDownOrCloseParamsClassOnce sync.Once
)

func getMTRWindowCoveringClusterDownOrCloseParamsClass() _MTRWindowCoveringClusterDownOrCloseParamsClass {
	MTRWindowCoveringClusterDownOrCloseParamsClassOnce.Do(func() {
		MTRWindowCoveringClusterDownOrCloseParamsClass = _MTRWindowCoveringClusterDownOrCloseParamsClass{objc.GetClass("MTRWindowCoveringClusterDownOrCloseParams")}
	})
	return MTRWindowCoveringClusterDownOrCloseParamsClass
}

type _MTRWindowCoveringClusterDownOrCloseParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRWindowCoveringClusterDownOrCloseParams */
// An interface definition for the [MTRWindowCoveringClusterDownOrCloseParams] class.
type IMTRWindowCoveringClusterDownOrCloseParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRWindowCoveringClusterDownOrCloseParams */
	// properties:
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRWindowCoveringClusterDownOrCloseParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRWindowCoveringClusterDownOrCloseParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRWindowCoveringClusterDownOrCloseParamsClass) Alloc() MTRWindowCoveringClusterDownOrCloseParams {
	rv := objc.Send[MTRWindowCoveringClusterDownOrCloseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRWindowCoveringClusterDownOrCloseParamsClass) New() MTRWindowCoveringClusterDownOrCloseParams {
	rv := objc.Send[MTRWindowCoveringClusterDownOrCloseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRWindowCoveringClusterDownOrCloseParams) Init() MTRWindowCoveringClusterDownOrCloseParams {
	rv := objc.Send[MTRWindowCoveringClusterDownOrCloseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRWindowCoveringClusterDownOrCloseParams) Autorelease() MTRWindowCoveringClusterDownOrCloseParams {
	rv := objc.Send[MTRWindowCoveringClusterDownOrCloseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRWindowCoveringClusterDownOrCloseParams creates a new MTRWindowCoveringClusterDownOrCloseParams instance.
func NewMTRWindowCoveringClusterDownOrCloseParams() MTRWindowCoveringClusterDownOrCloseParams {
	return getMTRWindowCoveringClusterDownOrCloseParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRWindowCoveringClusterDownOrCloseParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWindowCoveringClusterDownOrCloseParams
type MTRWindowCoveringClusterDownOrCloseParams struct {
	objectivec.Object
}

// MTRWindowCoveringClusterDownOrCloseParamsFrom constructs a [MTRWindowCoveringClusterDownOrCloseParams] from an unsafe.Pointer.
func MTRWindowCoveringClusterDownOrCloseParamsFrom(ptr unsafe.Pointer) MTRWindowCoveringClusterDownOrCloseParams {
	return MTRWindowCoveringClusterDownOrCloseParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRWindowCoveringClusterDownOrCloseParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRWindowCoveringClusterDownOrCloseParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRWindowCoveringClusterDownOrCloseParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRWindowCoveringClusterDownOrCloseParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRWindowCoveringClusterDownOrCloseParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWindowCoveringClusterDownOrCloseParams/serverSideProcessingTimeout
func (m_ MTRWindowCoveringClusterDownOrCloseParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWindowCoveringClusterDownOrCloseParams/serverSideProcessingTimeout
func (m_ MTRWindowCoveringClusterDownOrCloseParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWindowCoveringClusterDownOrCloseParams/timedInvokeTimeoutMs
func (m_ MTRWindowCoveringClusterDownOrCloseParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWindowCoveringClusterDownOrCloseParams/timedInvokeTimeoutMs
func (m_ MTRWindowCoveringClusterDownOrCloseParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRWindowCoveringClusterDownOrCloseParams */



