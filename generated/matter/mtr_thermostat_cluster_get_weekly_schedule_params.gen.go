// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRThermostatClusterGetWeeklyScheduleParams */


/* debug [class_header]: Header for MTRThermostatClusterGetWeeklyScheduleParams */
// The class instance for the [MTRThermostatClusterGetWeeklyScheduleParams] class.
var (
	MTRThermostatClusterGetWeeklyScheduleParamsClass     _MTRThermostatClusterGetWeeklyScheduleParamsClass
	MTRThermostatClusterGetWeeklyScheduleParamsClassOnce sync.Once
)

func getMTRThermostatClusterGetWeeklyScheduleParamsClass() _MTRThermostatClusterGetWeeklyScheduleParamsClass {
	MTRThermostatClusterGetWeeklyScheduleParamsClassOnce.Do(func() {
		MTRThermostatClusterGetWeeklyScheduleParamsClass = _MTRThermostatClusterGetWeeklyScheduleParamsClass{objc.GetClass("MTRThermostatClusterGetWeeklyScheduleParams")}
	})
	return MTRThermostatClusterGetWeeklyScheduleParamsClass
}

type _MTRThermostatClusterGetWeeklyScheduleParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRThermostatClusterGetWeeklyScheduleParams */
// An interface definition for the [MTRThermostatClusterGetWeeklyScheduleParams] class.
type IMTRThermostatClusterGetWeeklyScheduleParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRThermostatClusterGetWeeklyScheduleParams */
	// properties:
	DaysToReturn() objc.IObject /* cross-framework: NSNumber */
	SetDaysToReturn(value objc.IObject /* cross-framework: NSNumber */)
	ModeToReturn() objc.IObject /* cross-framework: NSNumber */
	SetModeToReturn(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRThermostatClusterGetWeeklyScheduleParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRThermostatClusterGetWeeklyScheduleParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRThermostatClusterGetWeeklyScheduleParamsClass) Alloc() MTRThermostatClusterGetWeeklyScheduleParams {
	rv := objc.Send[MTRThermostatClusterGetWeeklyScheduleParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRThermostatClusterGetWeeklyScheduleParamsClass) New() MTRThermostatClusterGetWeeklyScheduleParams {
	rv := objc.Send[MTRThermostatClusterGetWeeklyScheduleParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRThermostatClusterGetWeeklyScheduleParams) Init() MTRThermostatClusterGetWeeklyScheduleParams {
	rv := objc.Send[MTRThermostatClusterGetWeeklyScheduleParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRThermostatClusterGetWeeklyScheduleParams) Autorelease() MTRThermostatClusterGetWeeklyScheduleParams {
	rv := objc.Send[MTRThermostatClusterGetWeeklyScheduleParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRThermostatClusterGetWeeklyScheduleParams creates a new MTRThermostatClusterGetWeeklyScheduleParams instance.
func NewMTRThermostatClusterGetWeeklyScheduleParams() MTRThermostatClusterGetWeeklyScheduleParams {
	return getMTRThermostatClusterGetWeeklyScheduleParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRThermostatClusterGetWeeklyScheduleParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterGetWeeklyScheduleParams
type MTRThermostatClusterGetWeeklyScheduleParams struct {
	objectivec.Object
}

// MTRThermostatClusterGetWeeklyScheduleParamsFrom constructs a [MTRThermostatClusterGetWeeklyScheduleParams] from an unsafe.Pointer.
func MTRThermostatClusterGetWeeklyScheduleParamsFrom(ptr unsafe.Pointer) MTRThermostatClusterGetWeeklyScheduleParams {
	return MTRThermostatClusterGetWeeklyScheduleParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRThermostatClusterGetWeeklyScheduleParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRThermostatClusterGetWeeklyScheduleParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRThermostatClusterGetWeeklyScheduleParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRThermostatClusterGetWeeklyScheduleParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRThermostatClusterGetWeeklyScheduleParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterGetWeeklyScheduleParams/daysToReturn
func (m_ MTRThermostatClusterGetWeeklyScheduleParams) DaysToReturn() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("daysToReturn"))
	return rv
}/* debug [instance_properties/getter]: daysToReturn */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterGetWeeklyScheduleParams/daysToReturn
func (m_ MTRThermostatClusterGetWeeklyScheduleParams) SetDaysToReturn(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDaysToReturn:"), value)
}/* debug [instance_properties/setter]: daysToReturn */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterGetWeeklyScheduleParams/modeToReturn
func (m_ MTRThermostatClusterGetWeeklyScheduleParams) ModeToReturn() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("modeToReturn"))
	return rv
}/* debug [instance_properties/getter]: modeToReturn */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterGetWeeklyScheduleParams/modeToReturn
func (m_ MTRThermostatClusterGetWeeklyScheduleParams) SetModeToReturn(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setModeToReturn:"), value)
}/* debug [instance_properties/setter]: modeToReturn */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterGetWeeklyScheduleParams/serverSideProcessingTimeout
func (m_ MTRThermostatClusterGetWeeklyScheduleParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterGetWeeklyScheduleParams/serverSideProcessingTimeout
func (m_ MTRThermostatClusterGetWeeklyScheduleParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterGetWeeklyScheduleParams/timedInvokeTimeoutMs
func (m_ MTRThermostatClusterGetWeeklyScheduleParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterGetWeeklyScheduleParams/timedInvokeTimeoutMs
func (m_ MTRThermostatClusterGetWeeklyScheduleParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRThermostatClusterGetWeeklyScheduleParams */



