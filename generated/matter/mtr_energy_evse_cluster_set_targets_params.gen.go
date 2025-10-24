// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTREnergyEVSEClusterSetTargetsParams */


/* debug [class_header]: Header for MTREnergyEVSEClusterSetTargetsParams */
// The class instance for the [MTREnergyEVSEClusterSetTargetsParams] class.
var (
	MTREnergyEVSEClusterSetTargetsParamsClass     _MTREnergyEVSEClusterSetTargetsParamsClass
	MTREnergyEVSEClusterSetTargetsParamsClassOnce sync.Once
)

func getMTREnergyEVSEClusterSetTargetsParamsClass() _MTREnergyEVSEClusterSetTargetsParamsClass {
	MTREnergyEVSEClusterSetTargetsParamsClassOnce.Do(func() {
		MTREnergyEVSEClusterSetTargetsParamsClass = _MTREnergyEVSEClusterSetTargetsParamsClass{objc.GetClass("MTREnergyEVSEClusterSetTargetsParams")}
	})
	return MTREnergyEVSEClusterSetTargetsParamsClass
}

type _MTREnergyEVSEClusterSetTargetsParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTREnergyEVSEClusterSetTargetsParams */
// An interface definition for the [MTREnergyEVSEClusterSetTargetsParams] class.
type IMTREnergyEVSEClusterSetTargetsParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTREnergyEVSEClusterSetTargetsParams */
	// properties:
	ChargingTargetSchedules() objc.IObject /* cross-framework: NSArray */
	SetChargingTargetSchedules(value objc.IObject /* cross-framework: NSArray */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTREnergyEVSEClusterSetTargetsParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTREnergyEVSEClusterSetTargetsParams */
// Alloc allocates a new instance without initialization.
func (mc _MTREnergyEVSEClusterSetTargetsParamsClass) Alloc() MTREnergyEVSEClusterSetTargetsParams {
	rv := objc.Send[MTREnergyEVSEClusterSetTargetsParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTREnergyEVSEClusterSetTargetsParamsClass) New() MTREnergyEVSEClusterSetTargetsParams {
	rv := objc.Send[MTREnergyEVSEClusterSetTargetsParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTREnergyEVSEClusterSetTargetsParams) Init() MTREnergyEVSEClusterSetTargetsParams {
	rv := objc.Send[MTREnergyEVSEClusterSetTargetsParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTREnergyEVSEClusterSetTargetsParams) Autorelease() MTREnergyEVSEClusterSetTargetsParams {
	rv := objc.Send[MTREnergyEVSEClusterSetTargetsParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTREnergyEVSEClusterSetTargetsParams creates a new MTREnergyEVSEClusterSetTargetsParams instance.
func NewMTREnergyEVSEClusterSetTargetsParams() MTREnergyEVSEClusterSetTargetsParams {
	return getMTREnergyEVSEClusterSetTargetsParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTREnergyEVSEClusterSetTargetsParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterSetTargetsParams
type MTREnergyEVSEClusterSetTargetsParams struct {
	objectivec.Object
}

// MTREnergyEVSEClusterSetTargetsParamsFrom constructs a [MTREnergyEVSEClusterSetTargetsParams] from an unsafe.Pointer.
func MTREnergyEVSEClusterSetTargetsParamsFrom(ptr unsafe.Pointer) MTREnergyEVSEClusterSetTargetsParams {
	return MTREnergyEVSEClusterSetTargetsParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTREnergyEVSEClusterSetTargetsParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTREnergyEVSEClusterSetTargetsParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTREnergyEVSEClusterSetTargetsParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTREnergyEVSEClusterSetTargetsParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTREnergyEVSEClusterSetTargetsParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterSetTargetsParams/chargingTargetSchedules
func (m_ MTREnergyEVSEClusterSetTargetsParams) ChargingTargetSchedules() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("chargingTargetSchedules"))
	return rv
}/* debug [instance_properties/getter]: chargingTargetSchedules */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterSetTargetsParams/chargingTargetSchedules
func (m_ MTREnergyEVSEClusterSetTargetsParams) SetChargingTargetSchedules(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setChargingTargetSchedules:"), value)
}/* debug [instance_properties/setter]: chargingTargetSchedules */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrenergyevseclustersettargetsparams/serversideprocessingtimeout
func (m_ MTREnergyEVSEClusterSetTargetsParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrenergyevseclustersettargetsparams/serversideprocessingtimeout
func (m_ MTREnergyEVSEClusterSetTargetsParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrenergyevseclustersettargetsparams/timedinvoketimeoutms
func (m_ MTREnergyEVSEClusterSetTargetsParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrenergyevseclustersettargetsparams/timedinvoketimeoutms
func (m_ MTREnergyEVSEClusterSetTargetsParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTREnergyEVSEClusterSetTargetsParams */



