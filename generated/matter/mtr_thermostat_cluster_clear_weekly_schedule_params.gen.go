// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRThermostatClusterClearWeeklyScheduleParams */


/* debug [class_header]: Header for MTRThermostatClusterClearWeeklyScheduleParams */
// The class instance for the [MTRThermostatClusterClearWeeklyScheduleParams] class.
var (
	MTRThermostatClusterClearWeeklyScheduleParamsClass     _MTRThermostatClusterClearWeeklyScheduleParamsClass
	MTRThermostatClusterClearWeeklyScheduleParamsClassOnce sync.Once
)

func getMTRThermostatClusterClearWeeklyScheduleParamsClass() _MTRThermostatClusterClearWeeklyScheduleParamsClass {
	MTRThermostatClusterClearWeeklyScheduleParamsClassOnce.Do(func() {
		MTRThermostatClusterClearWeeklyScheduleParamsClass = _MTRThermostatClusterClearWeeklyScheduleParamsClass{objc.GetClass("MTRThermostatClusterClearWeeklyScheduleParams")}
	})
	return MTRThermostatClusterClearWeeklyScheduleParamsClass
}

type _MTRThermostatClusterClearWeeklyScheduleParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRThermostatClusterClearWeeklyScheduleParams */
// An interface definition for the [MTRThermostatClusterClearWeeklyScheduleParams] class.
type IMTRThermostatClusterClearWeeklyScheduleParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRThermostatClusterClearWeeklyScheduleParams */
	// properties:
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRThermostatClusterClearWeeklyScheduleParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRThermostatClusterClearWeeklyScheduleParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRThermostatClusterClearWeeklyScheduleParamsClass) Alloc() MTRThermostatClusterClearWeeklyScheduleParams {
	rv := objc.Send[MTRThermostatClusterClearWeeklyScheduleParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRThermostatClusterClearWeeklyScheduleParamsClass) New() MTRThermostatClusterClearWeeklyScheduleParams {
	rv := objc.Send[MTRThermostatClusterClearWeeklyScheduleParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRThermostatClusterClearWeeklyScheduleParams) Init() MTRThermostatClusterClearWeeklyScheduleParams {
	rv := objc.Send[MTRThermostatClusterClearWeeklyScheduleParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRThermostatClusterClearWeeklyScheduleParams) Autorelease() MTRThermostatClusterClearWeeklyScheduleParams {
	rv := objc.Send[MTRThermostatClusterClearWeeklyScheduleParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRThermostatClusterClearWeeklyScheduleParams creates a new MTRThermostatClusterClearWeeklyScheduleParams instance.
func NewMTRThermostatClusterClearWeeklyScheduleParams() MTRThermostatClusterClearWeeklyScheduleParams {
	return getMTRThermostatClusterClearWeeklyScheduleParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRThermostatClusterClearWeeklyScheduleParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterClearWeeklyScheduleParams
type MTRThermostatClusterClearWeeklyScheduleParams struct {
	objectivec.Object
}

// MTRThermostatClusterClearWeeklyScheduleParamsFrom constructs a [MTRThermostatClusterClearWeeklyScheduleParams] from an unsafe.Pointer.
func MTRThermostatClusterClearWeeklyScheduleParamsFrom(ptr unsafe.Pointer) MTRThermostatClusterClearWeeklyScheduleParams {
	return MTRThermostatClusterClearWeeklyScheduleParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRThermostatClusterClearWeeklyScheduleParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRThermostatClusterClearWeeklyScheduleParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRThermostatClusterClearWeeklyScheduleParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRThermostatClusterClearWeeklyScheduleParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRThermostatClusterClearWeeklyScheduleParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterClearWeeklyScheduleParams/serverSideProcessingTimeout
func (m_ MTRThermostatClusterClearWeeklyScheduleParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterClearWeeklyScheduleParams/serverSideProcessingTimeout
func (m_ MTRThermostatClusterClearWeeklyScheduleParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterClearWeeklyScheduleParams/timedInvokeTimeoutMs
func (m_ MTRThermostatClusterClearWeeklyScheduleParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterClearWeeklyScheduleParams/timedInvokeTimeoutMs
func (m_ MTRThermostatClusterClearWeeklyScheduleParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRThermostatClusterClearWeeklyScheduleParams */



