// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [MTRThermostatClusterSetWeeklyScheduleParams] class.
type IMTRThermostatClusterSetWeeklyScheduleParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterSetWeeklyScheduleParams
type MTRThermostatClusterSetWeeklyScheduleParams struct {
	objectivec.Object
}

// MTRThermostatClusterSetWeeklyScheduleParamsFrom constructs a [MTRThermostatClusterSetWeeklyScheduleParams] from an unsafe.Pointer.
func MTRThermostatClusterSetWeeklyScheduleParamsFrom(ptr unsafe.Pointer) MTRThermostatClusterSetWeeklyScheduleParams {
	return MTRThermostatClusterSetWeeklyScheduleParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRThermostatClusterSetWeeklyScheduleParamsClass) Alloc() MTRThermostatClusterSetWeeklyScheduleParams {
	rv := objc.Send[MTRThermostatClusterSetWeeklyScheduleParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclustersetweeklyscheduleparams/dayofweekforsequence
func (m_ MTRThermostatClusterSetWeeklyScheduleParams) DayOfWeekForSequence() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("dayOfWeekForSequence"))
	return rv
}


// SetDayOfWeekForSequence sets the value of the dayOfWeekForSequence property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclustersetweeklyscheduleparams/dayofweekforsequence
func (m_ MTRThermostatClusterSetWeeklyScheduleParams) SetDayOfWeekForSequence(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDayOfWeekForSequence:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclustersetweeklyscheduleparams/modeforsequence
func (m_ MTRThermostatClusterSetWeeklyScheduleParams) ModeForSequence() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("modeForSequence"))
	return rv
}


// SetModeForSequence sets the value of the modeForSequence property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclustersetweeklyscheduleparams/modeforsequence
func (m_ MTRThermostatClusterSetWeeklyScheduleParams) SetModeForSequence(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setModeForSequence:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclustersetweeklyscheduleparams/numberoftransitionsforsequence
func (m_ MTRThermostatClusterSetWeeklyScheduleParams) NumberOfTransitionsForSequence() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("numberOfTransitionsForSequence"))
	return rv
}


// SetNumberOfTransitionsForSequence sets the value of the numberOfTransitionsForSequence property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclustersetweeklyscheduleparams/numberoftransitionsforsequence
func (m_ MTRThermostatClusterSetWeeklyScheduleParams) SetNumberOfTransitionsForSequence(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNumberOfTransitionsForSequence:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclustersetweeklyscheduleparams/serversideprocessingtimeout
func (m_ MTRThermostatClusterSetWeeklyScheduleParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclustersetweeklyscheduleparams/serversideprocessingtimeout
func (m_ MTRThermostatClusterSetWeeklyScheduleParams) SetServerSideProcessingTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclustersetweeklyscheduleparams/timedinvoketimeoutms
func (m_ MTRThermostatClusterSetWeeklyScheduleParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclustersetweeklyscheduleparams/timedinvoketimeoutms
func (m_ MTRThermostatClusterSetWeeklyScheduleParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclustersetweeklyscheduleparams/transitions
func (m_ MTRThermostatClusterSetWeeklyScheduleParams) Transitions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("transitions"))
	return rv
}


// SetTransitions sets the value of the transitions property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclustersetweeklyscheduleparams/transitions
func (m_ MTRThermostatClusterSetWeeklyScheduleParams) SetTransitions(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTransitions:"), value)
}



