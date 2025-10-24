// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTREnergyEVSEClusterClearTargetsParams */


/* debug [class_header]: Header for MTREnergyEVSEClusterClearTargetsParams */
// The class instance for the [MTREnergyEVSEClusterClearTargetsParams] class.
var (
	MTREnergyEVSEClusterClearTargetsParamsClass     _MTREnergyEVSEClusterClearTargetsParamsClass
	MTREnergyEVSEClusterClearTargetsParamsClassOnce sync.Once
)

func getMTREnergyEVSEClusterClearTargetsParamsClass() _MTREnergyEVSEClusterClearTargetsParamsClass {
	MTREnergyEVSEClusterClearTargetsParamsClassOnce.Do(func() {
		MTREnergyEVSEClusterClearTargetsParamsClass = _MTREnergyEVSEClusterClearTargetsParamsClass{objc.GetClass("MTREnergyEVSEClusterClearTargetsParams")}
	})
	return MTREnergyEVSEClusterClearTargetsParamsClass
}

type _MTREnergyEVSEClusterClearTargetsParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTREnergyEVSEClusterClearTargetsParams */
// An interface definition for the [MTREnergyEVSEClusterClearTargetsParams] class.
type IMTREnergyEVSEClusterClearTargetsParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTREnergyEVSEClusterClearTargetsParams */
	// properties:
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTREnergyEVSEClusterClearTargetsParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTREnergyEVSEClusterClearTargetsParams */
// Alloc allocates a new instance without initialization.
func (mc _MTREnergyEVSEClusterClearTargetsParamsClass) Alloc() MTREnergyEVSEClusterClearTargetsParams {
	rv := objc.Send[MTREnergyEVSEClusterClearTargetsParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTREnergyEVSEClusterClearTargetsParamsClass) New() MTREnergyEVSEClusterClearTargetsParams {
	rv := objc.Send[MTREnergyEVSEClusterClearTargetsParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTREnergyEVSEClusterClearTargetsParams) Init() MTREnergyEVSEClusterClearTargetsParams {
	rv := objc.Send[MTREnergyEVSEClusterClearTargetsParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTREnergyEVSEClusterClearTargetsParams) Autorelease() MTREnergyEVSEClusterClearTargetsParams {
	rv := objc.Send[MTREnergyEVSEClusterClearTargetsParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTREnergyEVSEClusterClearTargetsParams creates a new MTREnergyEVSEClusterClearTargetsParams instance.
func NewMTREnergyEVSEClusterClearTargetsParams() MTREnergyEVSEClusterClearTargetsParams {
	return getMTREnergyEVSEClusterClearTargetsParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTREnergyEVSEClusterClearTargetsParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterClearTargetsParams
type MTREnergyEVSEClusterClearTargetsParams struct {
	objectivec.Object
}

// MTREnergyEVSEClusterClearTargetsParamsFrom constructs a [MTREnergyEVSEClusterClearTargetsParams] from an unsafe.Pointer.
func MTREnergyEVSEClusterClearTargetsParamsFrom(ptr unsafe.Pointer) MTREnergyEVSEClusterClearTargetsParams {
	return MTREnergyEVSEClusterClearTargetsParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTREnergyEVSEClusterClearTargetsParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTREnergyEVSEClusterClearTargetsParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTREnergyEVSEClusterClearTargetsParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTREnergyEVSEClusterClearTargetsParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTREnergyEVSEClusterClearTargetsParams */

// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterClearTargetsParams/serverSideProcessingTimeout
func (m_ MTREnergyEVSEClusterClearTargetsParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterClearTargetsParams/serverSideProcessingTimeout
func (m_ MTREnergyEVSEClusterClearTargetsParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterClearTargetsParams/timedInvokeTimeoutMs
func (m_ MTREnergyEVSEClusterClearTargetsParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterClearTargetsParams/timedInvokeTimeoutMs
func (m_ MTREnergyEVSEClusterClearTargetsParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTREnergyEVSEClusterClearTargetsParams */



