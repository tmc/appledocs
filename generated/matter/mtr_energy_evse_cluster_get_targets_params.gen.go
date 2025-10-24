// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTREnergyEVSEClusterGetTargetsParams */


/* debug [class_header]: Header for MTREnergyEVSEClusterGetTargetsParams */
// The class instance for the [MTREnergyEVSEClusterGetTargetsParams] class.
var (
	MTREnergyEVSEClusterGetTargetsParamsClass     _MTREnergyEVSEClusterGetTargetsParamsClass
	MTREnergyEVSEClusterGetTargetsParamsClassOnce sync.Once
)

func getMTREnergyEVSEClusterGetTargetsParamsClass() _MTREnergyEVSEClusterGetTargetsParamsClass {
	MTREnergyEVSEClusterGetTargetsParamsClassOnce.Do(func() {
		MTREnergyEVSEClusterGetTargetsParamsClass = _MTREnergyEVSEClusterGetTargetsParamsClass{objc.GetClass("MTREnergyEVSEClusterGetTargetsParams")}
	})
	return MTREnergyEVSEClusterGetTargetsParamsClass
}

type _MTREnergyEVSEClusterGetTargetsParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTREnergyEVSEClusterGetTargetsParams */
// An interface definition for the [MTREnergyEVSEClusterGetTargetsParams] class.
type IMTREnergyEVSEClusterGetTargetsParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTREnergyEVSEClusterGetTargetsParams */
	// properties:
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTREnergyEVSEClusterGetTargetsParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTREnergyEVSEClusterGetTargetsParams */
// Alloc allocates a new instance without initialization.
func (mc _MTREnergyEVSEClusterGetTargetsParamsClass) Alloc() MTREnergyEVSEClusterGetTargetsParams {
	rv := objc.Send[MTREnergyEVSEClusterGetTargetsParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTREnergyEVSEClusterGetTargetsParamsClass) New() MTREnergyEVSEClusterGetTargetsParams {
	rv := objc.Send[MTREnergyEVSEClusterGetTargetsParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTREnergyEVSEClusterGetTargetsParams) Init() MTREnergyEVSEClusterGetTargetsParams {
	rv := objc.Send[MTREnergyEVSEClusterGetTargetsParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTREnergyEVSEClusterGetTargetsParams) Autorelease() MTREnergyEVSEClusterGetTargetsParams {
	rv := objc.Send[MTREnergyEVSEClusterGetTargetsParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTREnergyEVSEClusterGetTargetsParams creates a new MTREnergyEVSEClusterGetTargetsParams instance.
func NewMTREnergyEVSEClusterGetTargetsParams() MTREnergyEVSEClusterGetTargetsParams {
	return getMTREnergyEVSEClusterGetTargetsParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTREnergyEVSEClusterGetTargetsParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterGetTargetsParams
type MTREnergyEVSEClusterGetTargetsParams struct {
	objectivec.Object
}

// MTREnergyEVSEClusterGetTargetsParamsFrom constructs a [MTREnergyEVSEClusterGetTargetsParams] from an unsafe.Pointer.
func MTREnergyEVSEClusterGetTargetsParamsFrom(ptr unsafe.Pointer) MTREnergyEVSEClusterGetTargetsParams {
	return MTREnergyEVSEClusterGetTargetsParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTREnergyEVSEClusterGetTargetsParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTREnergyEVSEClusterGetTargetsParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTREnergyEVSEClusterGetTargetsParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTREnergyEVSEClusterGetTargetsParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTREnergyEVSEClusterGetTargetsParams */

// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterGetTargetsParams/serverSideProcessingTimeout
func (m_ MTREnergyEVSEClusterGetTargetsParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterGetTargetsParams/serverSideProcessingTimeout
func (m_ MTREnergyEVSEClusterGetTargetsParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrenergyevseclustergettargetsparams/timedinvoketimeoutms
func (m_ MTREnergyEVSEClusterGetTargetsParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrenergyevseclustergettargetsparams/timedinvoketimeoutms
func (m_ MTREnergyEVSEClusterGetTargetsParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTREnergyEVSEClusterGetTargetsParams */



