// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [MTRThermostatClusterGetWeeklyScheduleResponseParams] class.
type IMTRThermostatClusterGetWeeklyScheduleResponseParams interface {
	objectivec.IObject
	// properties:
	DayOfWeekForSequence() objc.IObject /* cross-framework: NSNumber */
	SetDayOfWeekForSequence(value objc.IObject /* cross-framework: NSNumber */)
	ModeForSequence() objc.IObject /* cross-framework: NSNumber */
	SetModeForSequence(value objc.IObject /* cross-framework: NSNumber */)
	NumberOfTransitionsForSequence() objc.IObject /* cross-framework: NSNumber */
	SetNumberOfTransitionsForSequence(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	Transitions() unsafe.Pointer
	SetTransitions(value unsafe.Pointer)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterGetWeeklyScheduleResponseParams
type MTRThermostatClusterGetWeeklyScheduleResponseParams struct {
	objectivec.Object
}

// MTRThermostatClusterGetWeeklyScheduleResponseParamsFrom constructs a [MTRThermostatClusterGetWeeklyScheduleResponseParams] from an unsafe.Pointer.
func MTRThermostatClusterGetWeeklyScheduleResponseParamsFrom(ptr unsafe.Pointer) MTRThermostatClusterGetWeeklyScheduleResponseParams {
	return MTRThermostatClusterGetWeeklyScheduleResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRThermostatClusterGetWeeklyScheduleResponseParamsClass) Alloc() MTRThermostatClusterGetWeeklyScheduleResponseParams {
	rv := objc.Send[MTRThermostatClusterGetWeeklyScheduleResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclustergetweeklyscheduleresponseparams/dayofweekforsequence
func (m_ MTRThermostatClusterGetWeeklyScheduleResponseParams) DayOfWeekForSequence() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("dayOfWeekForSequence"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclustergetweeklyscheduleresponseparams/dayofweekforsequence
func (m_ MTRThermostatClusterGetWeeklyScheduleResponseParams) SetDayOfWeekForSequence(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDayOfWeekForSequence:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclustergetweeklyscheduleresponseparams/modeforsequence
func (m_ MTRThermostatClusterGetWeeklyScheduleResponseParams) ModeForSequence() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("modeForSequence"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclustergetweeklyscheduleresponseparams/modeforsequence
func (m_ MTRThermostatClusterGetWeeklyScheduleResponseParams) SetModeForSequence(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setModeForSequence:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclustergetweeklyscheduleresponseparams/numberoftransitionsforsequence
func (m_ MTRThermostatClusterGetWeeklyScheduleResponseParams) NumberOfTransitionsForSequence() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("numberOfTransitionsForSequence"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclustergetweeklyscheduleresponseparams/numberoftransitionsforsequence
func (m_ MTRThermostatClusterGetWeeklyScheduleResponseParams) SetNumberOfTransitionsForSequence(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNumberOfTransitionsForSequence:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclustergetweeklyscheduleresponseparams/timedinvoketimeoutms
func (m_ MTRThermostatClusterGetWeeklyScheduleResponseParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclustergetweeklyscheduleresponseparams/timedinvoketimeoutms
func (m_ MTRThermostatClusterGetWeeklyScheduleResponseParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclustergetweeklyscheduleresponseparams/transitions
func (m_ MTRThermostatClusterGetWeeklyScheduleResponseParams) Transitions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("transitions"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclustergetweeklyscheduleresponseparams/transitions
func (m_ MTRThermostatClusterGetWeeklyScheduleResponseParams) SetTransitions(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTransitions:"), value)
}



