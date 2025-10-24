// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTROnOffClusterToggleParams */


/* debug [class_header]: Header for MTROnOffClusterToggleParams */
// The class instance for the [MTROnOffClusterToggleParams] class.
var (
	MTROnOffClusterToggleParamsClass     _MTROnOffClusterToggleParamsClass
	MTROnOffClusterToggleParamsClassOnce sync.Once
)

func getMTROnOffClusterToggleParamsClass() _MTROnOffClusterToggleParamsClass {
	MTROnOffClusterToggleParamsClassOnce.Do(func() {
		MTROnOffClusterToggleParamsClass = _MTROnOffClusterToggleParamsClass{objc.GetClass("MTROnOffClusterToggleParams")}
	})
	return MTROnOffClusterToggleParamsClass
}

type _MTROnOffClusterToggleParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTROnOffClusterToggleParams */
// An interface definition for the [MTROnOffClusterToggleParams] class.
type IMTROnOffClusterToggleParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTROnOffClusterToggleParams */
	// properties:
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTROnOffClusterToggleParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTROnOffClusterToggleParams */
// Alloc allocates a new instance without initialization.
func (mc _MTROnOffClusterToggleParamsClass) Alloc() MTROnOffClusterToggleParams {
	rv := objc.Send[MTROnOffClusterToggleParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTROnOffClusterToggleParamsClass) New() MTROnOffClusterToggleParams {
	rv := objc.Send[MTROnOffClusterToggleParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROnOffClusterToggleParams) Init() MTROnOffClusterToggleParams {
	rv := objc.Send[MTROnOffClusterToggleParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROnOffClusterToggleParams) Autorelease() MTROnOffClusterToggleParams {
	rv := objc.Send[MTROnOffClusterToggleParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROnOffClusterToggleParams creates a new MTROnOffClusterToggleParams instance.
func NewMTROnOffClusterToggleParams() MTROnOffClusterToggleParams {
	return getMTROnOffClusterToggleParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTROnOffClusterToggleParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROnOffClusterToggleParams
type MTROnOffClusterToggleParams struct {
	objectivec.Object
}

// MTROnOffClusterToggleParamsFrom constructs a [MTROnOffClusterToggleParams] from an unsafe.Pointer.
func MTROnOffClusterToggleParamsFrom(ptr unsafe.Pointer) MTROnOffClusterToggleParams {
	return MTROnOffClusterToggleParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTROnOffClusterToggleParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTROnOffClusterToggleParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTROnOffClusterToggleParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTROnOffClusterToggleParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTROnOffClusterToggleParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROnOffClusterToggleParams/serverSideProcessingTimeout
func (m_ MTROnOffClusterToggleParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROnOffClusterToggleParams/serverSideProcessingTimeout
func (m_ MTROnOffClusterToggleParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROnOffClusterToggleParams/timedInvokeTimeoutMs
func (m_ MTROnOffClusterToggleParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROnOffClusterToggleParams/timedInvokeTimeoutMs
func (m_ MTROnOffClusterToggleParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTROnOffClusterToggleParams */



