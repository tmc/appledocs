// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTREnergyEVSEClusterDisableParams */


/* debug [class_header]: Header for MTREnergyEVSEClusterDisableParams */
// The class instance for the [MTREnergyEVSEClusterDisableParams] class.
var (
	MTREnergyEVSEClusterDisableParamsClass     _MTREnergyEVSEClusterDisableParamsClass
	MTREnergyEVSEClusterDisableParamsClassOnce sync.Once
)

func getMTREnergyEVSEClusterDisableParamsClass() _MTREnergyEVSEClusterDisableParamsClass {
	MTREnergyEVSEClusterDisableParamsClassOnce.Do(func() {
		MTREnergyEVSEClusterDisableParamsClass = _MTREnergyEVSEClusterDisableParamsClass{objc.GetClass("MTREnergyEVSEClusterDisableParams")}
	})
	return MTREnergyEVSEClusterDisableParamsClass
}

type _MTREnergyEVSEClusterDisableParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTREnergyEVSEClusterDisableParams */
// An interface definition for the [MTREnergyEVSEClusterDisableParams] class.
type IMTREnergyEVSEClusterDisableParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTREnergyEVSEClusterDisableParams */
	// properties:
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTREnergyEVSEClusterDisableParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTREnergyEVSEClusterDisableParams */
// Alloc allocates a new instance without initialization.
func (mc _MTREnergyEVSEClusterDisableParamsClass) Alloc() MTREnergyEVSEClusterDisableParams {
	rv := objc.Send[MTREnergyEVSEClusterDisableParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTREnergyEVSEClusterDisableParamsClass) New() MTREnergyEVSEClusterDisableParams {
	rv := objc.Send[MTREnergyEVSEClusterDisableParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTREnergyEVSEClusterDisableParams) Init() MTREnergyEVSEClusterDisableParams {
	rv := objc.Send[MTREnergyEVSEClusterDisableParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTREnergyEVSEClusterDisableParams) Autorelease() MTREnergyEVSEClusterDisableParams {
	rv := objc.Send[MTREnergyEVSEClusterDisableParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTREnergyEVSEClusterDisableParams creates a new MTREnergyEVSEClusterDisableParams instance.
func NewMTREnergyEVSEClusterDisableParams() MTREnergyEVSEClusterDisableParams {
	return getMTREnergyEVSEClusterDisableParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTREnergyEVSEClusterDisableParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterDisableParams
type MTREnergyEVSEClusterDisableParams struct {
	objectivec.Object
}

// MTREnergyEVSEClusterDisableParamsFrom constructs a [MTREnergyEVSEClusterDisableParams] from an unsafe.Pointer.
func MTREnergyEVSEClusterDisableParamsFrom(ptr unsafe.Pointer) MTREnergyEVSEClusterDisableParams {
	return MTREnergyEVSEClusterDisableParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTREnergyEVSEClusterDisableParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTREnergyEVSEClusterDisableParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTREnergyEVSEClusterDisableParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTREnergyEVSEClusterDisableParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTREnergyEVSEClusterDisableParams */

// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterDisableParams/serverSideProcessingTimeout
func (m_ MTREnergyEVSEClusterDisableParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterDisableParams/serverSideProcessingTimeout
func (m_ MTREnergyEVSEClusterDisableParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrenergyevseclusterdisableparams/timedinvoketimeoutms
func (m_ MTREnergyEVSEClusterDisableParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrenergyevseclusterdisableparams/timedinvoketimeoutms
func (m_ MTREnergyEVSEClusterDisableParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTREnergyEVSEClusterDisableParams */



