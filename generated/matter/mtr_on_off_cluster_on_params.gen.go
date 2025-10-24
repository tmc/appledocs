// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTROnOffClusterOnParams */


/* debug [class_header]: Header for MTROnOffClusterOnParams */
// The class instance for the [MTROnOffClusterOnParams] class.
var (
	MTROnOffClusterOnParamsClass     _MTROnOffClusterOnParamsClass
	MTROnOffClusterOnParamsClassOnce sync.Once
)

func getMTROnOffClusterOnParamsClass() _MTROnOffClusterOnParamsClass {
	MTROnOffClusterOnParamsClassOnce.Do(func() {
		MTROnOffClusterOnParamsClass = _MTROnOffClusterOnParamsClass{objc.GetClass("MTROnOffClusterOnParams")}
	})
	return MTROnOffClusterOnParamsClass
}

type _MTROnOffClusterOnParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTROnOffClusterOnParams */
// An interface definition for the [MTROnOffClusterOnParams] class.
type IMTROnOffClusterOnParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTROnOffClusterOnParams */
	// properties:
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTROnOffClusterOnParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTROnOffClusterOnParams */
// Alloc allocates a new instance without initialization.
func (mc _MTROnOffClusterOnParamsClass) Alloc() MTROnOffClusterOnParams {
	rv := objc.Send[MTROnOffClusterOnParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTROnOffClusterOnParamsClass) New() MTROnOffClusterOnParams {
	rv := objc.Send[MTROnOffClusterOnParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROnOffClusterOnParams) Init() MTROnOffClusterOnParams {
	rv := objc.Send[MTROnOffClusterOnParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROnOffClusterOnParams) Autorelease() MTROnOffClusterOnParams {
	rv := objc.Send[MTROnOffClusterOnParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROnOffClusterOnParams creates a new MTROnOffClusterOnParams instance.
func NewMTROnOffClusterOnParams() MTROnOffClusterOnParams {
	return getMTROnOffClusterOnParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTROnOffClusterOnParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROnOffClusterOnParams
type MTROnOffClusterOnParams struct {
	objectivec.Object
}

// MTROnOffClusterOnParamsFrom constructs a [MTROnOffClusterOnParams] from an unsafe.Pointer.
func MTROnOffClusterOnParamsFrom(ptr unsafe.Pointer) MTROnOffClusterOnParams {
	return MTROnOffClusterOnParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTROnOffClusterOnParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTROnOffClusterOnParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTROnOffClusterOnParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTROnOffClusterOnParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTROnOffClusterOnParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROnOffClusterOnParams/serverSideProcessingTimeout
func (m_ MTROnOffClusterOnParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROnOffClusterOnParams/serverSideProcessingTimeout
func (m_ MTROnOffClusterOnParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROnOffClusterOnParams/timedInvokeTimeoutMs
func (m_ MTROnOffClusterOnParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROnOffClusterOnParams/timedInvokeTimeoutMs
func (m_ MTROnOffClusterOnParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTROnOffClusterOnParams */



