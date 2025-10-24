// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRThermostatClusterSetWeeklyScheduleParams */


/* debug [class_header]: Header for MTRThermostatClusterSetWeeklyScheduleParams */
// The class instance for the [MTRThermostatClusterSetWeeklyScheduleParams] class.
var (
	MTRThermostatClusterSetWeeklyScheduleParamsClass     _MTRThermostatClusterSetWeeklyScheduleParamsClass
	MTRThermostatClusterSetWeeklyScheduleParamsClassOnce sync.Once
)

func getMTRThermostatClusterSetWeeklyScheduleParamsClass() _MTRThermostatClusterSetWeeklyScheduleParamsClass {
	MTRThermostatClusterSetWeeklyScheduleParamsClassOnce.Do(func() {
		MTRThermostatClusterSetWeeklyScheduleParamsClass = _MTRThermostatClusterSetWeeklyScheduleParamsClass{objc.GetClass("MTRThermostatClusterSetWeeklyScheduleParams")}
	})
	return MTRThermostatClusterSetWeeklyScheduleParamsClass
}

type _MTRThermostatClusterSetWeeklyScheduleParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRThermostatClusterSetWeeklyScheduleParams */
// An interface definition for the [MTRThermostatClusterSetWeeklyScheduleParams] class.
type IMTRThermostatClusterSetWeeklyScheduleParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRThermostatClusterSetWeeklyScheduleParams */
	// properties:
	DayOfWeekForSequence() objc.IObject /* cross-framework: NSNumber */
	SetDayOfWeekForSequence(value objc.IObject /* cross-framework: NSNumber */)
	ModeForSequence() objc.IObject /* cross-framework: NSNumber */
	SetModeForSequence(value objc.IObject /* cross-framework: NSNumber */)
	NumberOfTransitionsForSequence() objc.IObject /* cross-framework: NSNumber */
	SetNumberOfTransitionsForSequence(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	Transitions() objc.IObject /* cross-framework: NSArray */
	SetTransitions(value objc.IObject /* cross-framework: NSArray */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRThermostatClusterSetWeeklyScheduleParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRThermostatClusterSetWeeklyScheduleParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRThermostatClusterSetWeeklyScheduleParamsClass) Alloc() MTRThermostatClusterSetWeeklyScheduleParams {
	rv := objc.Send[MTRThermostatClusterSetWeeklyScheduleParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRThermostatClusterSetWeeklyScheduleParamsClass) New() MTRThermostatClusterSetWeeklyScheduleParams {
	rv := objc.Send[MTRThermostatClusterSetWeeklyScheduleParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRThermostatClusterSetWeeklyScheduleParams) Init() MTRThermostatClusterSetWeeklyScheduleParams {
	rv := objc.Send[MTRThermostatClusterSetWeeklyScheduleParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRThermostatClusterSetWeeklyScheduleParams) Autorelease() MTRThermostatClusterSetWeeklyScheduleParams {
	rv := objc.Send[MTRThermostatClusterSetWeeklyScheduleParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRThermostatClusterSetWeeklyScheduleParams creates a new MTRThermostatClusterSetWeeklyScheduleParams instance.
func NewMTRThermostatClusterSetWeeklyScheduleParams() MTRThermostatClusterSetWeeklyScheduleParams {
	return getMTRThermostatClusterSetWeeklyScheduleParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRThermostatClusterSetWeeklyScheduleParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterSetWeeklyScheduleParams
type MTRThermostatClusterSetWeeklyScheduleParams struct {
	objectivec.Object
}

// MTRThermostatClusterSetWeeklyScheduleParamsFrom constructs a [MTRThermostatClusterSetWeeklyScheduleParams] from an unsafe.Pointer.
func MTRThermostatClusterSetWeeklyScheduleParamsFrom(ptr unsafe.Pointer) MTRThermostatClusterSetWeeklyScheduleParams {
	return MTRThermostatClusterSetWeeklyScheduleParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRThermostatClusterSetWeeklyScheduleParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRThermostatClusterSetWeeklyScheduleParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRThermostatClusterSetWeeklyScheduleParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRThermostatClusterSetWeeklyScheduleParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRThermostatClusterSetWeeklyScheduleParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterSetWeeklyScheduleParams/dayOfWeekForSequence
func (m_ MTRThermostatClusterSetWeeklyScheduleParams) DayOfWeekForSequence() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("dayOfWeekForSequence"))
	return rv
}/* debug [instance_properties/getter]: dayOfWeekForSequence */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterSetWeeklyScheduleParams/dayOfWeekForSequence
func (m_ MTRThermostatClusterSetWeeklyScheduleParams) SetDayOfWeekForSequence(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDayOfWeekForSequence:"), value)
}/* debug [instance_properties/setter]: dayOfWeekForSequence */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterSetWeeklyScheduleParams/modeForSequence
func (m_ MTRThermostatClusterSetWeeklyScheduleParams) ModeForSequence() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("modeForSequence"))
	return rv
}/* debug [instance_properties/getter]: modeForSequence */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterSetWeeklyScheduleParams/modeForSequence
func (m_ MTRThermostatClusterSetWeeklyScheduleParams) SetModeForSequence(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setModeForSequence:"), value)
}/* debug [instance_properties/setter]: modeForSequence */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterSetWeeklyScheduleParams/numberOfTransitionsForSequence
func (m_ MTRThermostatClusterSetWeeklyScheduleParams) NumberOfTransitionsForSequence() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("numberOfTransitionsForSequence"))
	return rv
}/* debug [instance_properties/getter]: numberOfTransitionsForSequence */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterSetWeeklyScheduleParams/numberOfTransitionsForSequence
func (m_ MTRThermostatClusterSetWeeklyScheduleParams) SetNumberOfTransitionsForSequence(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNumberOfTransitionsForSequence:"), value)
}/* debug [instance_properties/setter]: numberOfTransitionsForSequence */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterSetWeeklyScheduleParams/serverSideProcessingTimeout
func (m_ MTRThermostatClusterSetWeeklyScheduleParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterSetWeeklyScheduleParams/serverSideProcessingTimeout
func (m_ MTRThermostatClusterSetWeeklyScheduleParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterSetWeeklyScheduleParams/timedInvokeTimeoutMs
func (m_ MTRThermostatClusterSetWeeklyScheduleParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterSetWeeklyScheduleParams/timedInvokeTimeoutMs
func (m_ MTRThermostatClusterSetWeeklyScheduleParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterSetWeeklyScheduleParams/transitions
func (m_ MTRThermostatClusterSetWeeklyScheduleParams) Transitions() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("transitions"))
	return rv
}/* debug [instance_properties/getter]: transitions */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterSetWeeklyScheduleParams/transitions
func (m_ MTRThermostatClusterSetWeeklyScheduleParams) SetTransitions(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTransitions:"), value)
}/* debug [instance_properties/setter]: transitions */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRThermostatClusterSetWeeklyScheduleParams */



