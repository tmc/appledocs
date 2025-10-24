// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRThermostatClusterGetWeeklyScheduleResponseParams */


/* debug [class_header]: Header for MTRThermostatClusterGetWeeklyScheduleResponseParams */
// The class instance for the [MTRThermostatClusterGetWeeklyScheduleResponseParams] class.
var (
	MTRThermostatClusterGetWeeklyScheduleResponseParamsClass     _MTRThermostatClusterGetWeeklyScheduleResponseParamsClass
	MTRThermostatClusterGetWeeklyScheduleResponseParamsClassOnce sync.Once
)

func getMTRThermostatClusterGetWeeklyScheduleResponseParamsClass() _MTRThermostatClusterGetWeeklyScheduleResponseParamsClass {
	MTRThermostatClusterGetWeeklyScheduleResponseParamsClassOnce.Do(func() {
		MTRThermostatClusterGetWeeklyScheduleResponseParamsClass = _MTRThermostatClusterGetWeeklyScheduleResponseParamsClass{objc.GetClass("MTRThermostatClusterGetWeeklyScheduleResponseParams")}
	})
	return MTRThermostatClusterGetWeeklyScheduleResponseParamsClass
}

type _MTRThermostatClusterGetWeeklyScheduleResponseParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRThermostatClusterGetWeeklyScheduleResponseParams */
// An interface definition for the [MTRThermostatClusterGetWeeklyScheduleResponseParams] class.
type IMTRThermostatClusterGetWeeklyScheduleResponseParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRThermostatClusterGetWeeklyScheduleResponseParams */
	// properties:
	DayOfWeekForSequence() objc.IObject /* cross-framework: NSNumber */
	SetDayOfWeekForSequence(value objc.IObject /* cross-framework: NSNumber */)
	ModeForSequence() objc.IObject /* cross-framework: NSNumber */
	SetModeForSequence(value objc.IObject /* cross-framework: NSNumber */)
	NumberOfTransitionsForSequence() objc.IObject /* cross-framework: NSNumber */
	SetNumberOfTransitionsForSequence(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	Transitions() objc.IObject /* cross-framework: NSArray */
	SetTransitions(value objc.IObject /* cross-framework: NSArray */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRThermostatClusterGetWeeklyScheduleResponseParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRThermostatClusterGetWeeklyScheduleResponseParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRThermostatClusterGetWeeklyScheduleResponseParamsClass) Alloc() MTRThermostatClusterGetWeeklyScheduleResponseParams {
	rv := objc.Send[MTRThermostatClusterGetWeeklyScheduleResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRThermostatClusterGetWeeklyScheduleResponseParamsClass) New() MTRThermostatClusterGetWeeklyScheduleResponseParams {
	rv := objc.Send[MTRThermostatClusterGetWeeklyScheduleResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRThermostatClusterGetWeeklyScheduleResponseParams) Init() MTRThermostatClusterGetWeeklyScheduleResponseParams {
	rv := objc.Send[MTRThermostatClusterGetWeeklyScheduleResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRThermostatClusterGetWeeklyScheduleResponseParams) Autorelease() MTRThermostatClusterGetWeeklyScheduleResponseParams {
	rv := objc.Send[MTRThermostatClusterGetWeeklyScheduleResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRThermostatClusterGetWeeklyScheduleResponseParams creates a new MTRThermostatClusterGetWeeklyScheduleResponseParams instance.
func NewMTRThermostatClusterGetWeeklyScheduleResponseParams() MTRThermostatClusterGetWeeklyScheduleResponseParams {
	return getMTRThermostatClusterGetWeeklyScheduleResponseParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRThermostatClusterGetWeeklyScheduleResponseParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterGetWeeklyScheduleResponseParams
type MTRThermostatClusterGetWeeklyScheduleResponseParams struct {
	objectivec.Object
}

// MTRThermostatClusterGetWeeklyScheduleResponseParamsFrom constructs a [MTRThermostatClusterGetWeeklyScheduleResponseParams] from an unsafe.Pointer.
func MTRThermostatClusterGetWeeklyScheduleResponseParamsFrom(ptr unsafe.Pointer) MTRThermostatClusterGetWeeklyScheduleResponseParams {
	return MTRThermostatClusterGetWeeklyScheduleResponseParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRThermostatClusterGetWeeklyScheduleResponseParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterGetWeeklyScheduleResponseParams/init(responseValue:)
func NewMTRThermostatClusterGetWeeklyScheduleResponseParamsWithResponseValueError(responseValue foundation.IDictionary, error_ unsafe.Pointer) MTRThermostatClusterGetWeeklyScheduleResponseParams {
	instance := getMTRThermostatClusterGetWeeklyScheduleResponseParamsClass().Alloc()
	rv := objc.Send[MTRThermostatClusterGetWeeklyScheduleResponseParams](instance.ID, objc.Sel("initWithResponseValue:error:"), responseValue, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRThermostatClusterGetWeeklyScheduleResponseParamsWithResponseValueError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRThermostatClusterGetWeeklyScheduleResponseParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRThermostatClusterGetWeeklyScheduleResponseParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRThermostatClusterGetWeeklyScheduleResponseParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRThermostatClusterGetWeeklyScheduleResponseParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterGetWeeklyScheduleResponseParams/dayOfWeekForSequence
func (m_ MTRThermostatClusterGetWeeklyScheduleResponseParams) DayOfWeekForSequence() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("dayOfWeekForSequence"))
	return rv
}/* debug [instance_properties/getter]: dayOfWeekForSequence */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterGetWeeklyScheduleResponseParams/dayOfWeekForSequence
func (m_ MTRThermostatClusterGetWeeklyScheduleResponseParams) SetDayOfWeekForSequence(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDayOfWeekForSequence:"), value)
}/* debug [instance_properties/setter]: dayOfWeekForSequence */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterGetWeeklyScheduleResponseParams/modeForSequence
func (m_ MTRThermostatClusterGetWeeklyScheduleResponseParams) ModeForSequence() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("modeForSequence"))
	return rv
}/* debug [instance_properties/getter]: modeForSequence */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterGetWeeklyScheduleResponseParams/modeForSequence
func (m_ MTRThermostatClusterGetWeeklyScheduleResponseParams) SetModeForSequence(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setModeForSequence:"), value)
}/* debug [instance_properties/setter]: modeForSequence */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterGetWeeklyScheduleResponseParams/numberOfTransitionsForSequence
func (m_ MTRThermostatClusterGetWeeklyScheduleResponseParams) NumberOfTransitionsForSequence() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("numberOfTransitionsForSequence"))
	return rv
}/* debug [instance_properties/getter]: numberOfTransitionsForSequence */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterGetWeeklyScheduleResponseParams/numberOfTransitionsForSequence
func (m_ MTRThermostatClusterGetWeeklyScheduleResponseParams) SetNumberOfTransitionsForSequence(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNumberOfTransitionsForSequence:"), value)
}/* debug [instance_properties/setter]: numberOfTransitionsForSequence */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterGetWeeklyScheduleResponseParams/timedInvokeTimeoutMs
func (m_ MTRThermostatClusterGetWeeklyScheduleResponseParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterGetWeeklyScheduleResponseParams/timedInvokeTimeoutMs
func (m_ MTRThermostatClusterGetWeeklyScheduleResponseParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterGetWeeklyScheduleResponseParams/transitions
func (m_ MTRThermostatClusterGetWeeklyScheduleResponseParams) Transitions() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("transitions"))
	return rv
}/* debug [instance_properties/getter]: transitions */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterGetWeeklyScheduleResponseParams/transitions
func (m_ MTRThermostatClusterGetWeeklyScheduleResponseParams) SetTransitions(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTransitions:"), value)
}/* debug [instance_properties/setter]: transitions */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRThermostatClusterGetWeeklyScheduleResponseParams */


