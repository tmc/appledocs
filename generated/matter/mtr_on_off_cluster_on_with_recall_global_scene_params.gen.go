// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTROnOffClusterOnWithRecallGlobalSceneParams */


/* debug [class_header]: Header for MTROnOffClusterOnWithRecallGlobalSceneParams */
// The class instance for the [MTROnOffClusterOnWithRecallGlobalSceneParams] class.
var (
	MTROnOffClusterOnWithRecallGlobalSceneParamsClass     _MTROnOffClusterOnWithRecallGlobalSceneParamsClass
	MTROnOffClusterOnWithRecallGlobalSceneParamsClassOnce sync.Once
)

func getMTROnOffClusterOnWithRecallGlobalSceneParamsClass() _MTROnOffClusterOnWithRecallGlobalSceneParamsClass {
	MTROnOffClusterOnWithRecallGlobalSceneParamsClassOnce.Do(func() {
		MTROnOffClusterOnWithRecallGlobalSceneParamsClass = _MTROnOffClusterOnWithRecallGlobalSceneParamsClass{objc.GetClass("MTROnOffClusterOnWithRecallGlobalSceneParams")}
	})
	return MTROnOffClusterOnWithRecallGlobalSceneParamsClass
}

type _MTROnOffClusterOnWithRecallGlobalSceneParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTROnOffClusterOnWithRecallGlobalSceneParams */
// An interface definition for the [MTROnOffClusterOnWithRecallGlobalSceneParams] class.
type IMTROnOffClusterOnWithRecallGlobalSceneParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTROnOffClusterOnWithRecallGlobalSceneParams */
	// properties:
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTROnOffClusterOnWithRecallGlobalSceneParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTROnOffClusterOnWithRecallGlobalSceneParams */
// Alloc allocates a new instance without initialization.
func (mc _MTROnOffClusterOnWithRecallGlobalSceneParamsClass) Alloc() MTROnOffClusterOnWithRecallGlobalSceneParams {
	rv := objc.Send[MTROnOffClusterOnWithRecallGlobalSceneParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTROnOffClusterOnWithRecallGlobalSceneParamsClass) New() MTROnOffClusterOnWithRecallGlobalSceneParams {
	rv := objc.Send[MTROnOffClusterOnWithRecallGlobalSceneParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROnOffClusterOnWithRecallGlobalSceneParams) Init() MTROnOffClusterOnWithRecallGlobalSceneParams {
	rv := objc.Send[MTROnOffClusterOnWithRecallGlobalSceneParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROnOffClusterOnWithRecallGlobalSceneParams) Autorelease() MTROnOffClusterOnWithRecallGlobalSceneParams {
	rv := objc.Send[MTROnOffClusterOnWithRecallGlobalSceneParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROnOffClusterOnWithRecallGlobalSceneParams creates a new MTROnOffClusterOnWithRecallGlobalSceneParams instance.
func NewMTROnOffClusterOnWithRecallGlobalSceneParams() MTROnOffClusterOnWithRecallGlobalSceneParams {
	return getMTROnOffClusterOnWithRecallGlobalSceneParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTROnOffClusterOnWithRecallGlobalSceneParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROnOffClusterOnWithRecallGlobalSceneParams
type MTROnOffClusterOnWithRecallGlobalSceneParams struct {
	objectivec.Object
}

// MTROnOffClusterOnWithRecallGlobalSceneParamsFrom constructs a [MTROnOffClusterOnWithRecallGlobalSceneParams] from an unsafe.Pointer.
func MTROnOffClusterOnWithRecallGlobalSceneParamsFrom(ptr unsafe.Pointer) MTROnOffClusterOnWithRecallGlobalSceneParams {
	return MTROnOffClusterOnWithRecallGlobalSceneParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTROnOffClusterOnWithRecallGlobalSceneParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTROnOffClusterOnWithRecallGlobalSceneParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTROnOffClusterOnWithRecallGlobalSceneParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTROnOffClusterOnWithRecallGlobalSceneParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTROnOffClusterOnWithRecallGlobalSceneParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROnOffClusterOnWithRecallGlobalSceneParams/serverSideProcessingTimeout
func (m_ MTROnOffClusterOnWithRecallGlobalSceneParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROnOffClusterOnWithRecallGlobalSceneParams/serverSideProcessingTimeout
func (m_ MTROnOffClusterOnWithRecallGlobalSceneParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROnOffClusterOnWithRecallGlobalSceneParams/timedInvokeTimeoutMs
func (m_ MTROnOffClusterOnWithRecallGlobalSceneParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROnOffClusterOnWithRecallGlobalSceneParams/timedInvokeTimeoutMs
func (m_ MTROnOffClusterOnWithRecallGlobalSceneParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTROnOffClusterOnWithRecallGlobalSceneParams */



