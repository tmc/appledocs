// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [MTRThermostatClusterGetWeeklyScheduleParams] class.
type IMTRThermostatClusterGetWeeklyScheduleParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterGetWeeklyScheduleParams
type MTRThermostatClusterGetWeeklyScheduleParams struct {
	objectivec.Object
}

// MTRThermostatClusterGetWeeklyScheduleParamsFrom constructs a [MTRThermostatClusterGetWeeklyScheduleParams] from an unsafe.Pointer.
func MTRThermostatClusterGetWeeklyScheduleParamsFrom(ptr unsafe.Pointer) MTRThermostatClusterGetWeeklyScheduleParams {
	return MTRThermostatClusterGetWeeklyScheduleParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRThermostatClusterGetWeeklyScheduleParamsClass) Alloc() MTRThermostatClusterGetWeeklyScheduleParams {
	rv := objc.Send[MTRThermostatClusterGetWeeklyScheduleParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclustergetweeklyscheduleparams/timedinvoketimeoutms
func (m_ MTRThermostatClusterGetWeeklyScheduleParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclustergetweeklyscheduleparams/timedinvoketimeoutms
func (m_ MTRThermostatClusterGetWeeklyScheduleParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclustergetweeklyscheduleparams/daystoreturn
func (m_ MTRThermostatClusterGetWeeklyScheduleParams) DaysToReturn() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("daysToReturn"))
	return rv
}


// SetDaysToReturn sets the value of the daysToReturn property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclustergetweeklyscheduleparams/daystoreturn
func (m_ MTRThermostatClusterGetWeeklyScheduleParams) SetDaysToReturn(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDaysToReturn:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclustergetweeklyscheduleparams/serversideprocessingtimeout
func (m_ MTRThermostatClusterGetWeeklyScheduleParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclustergetweeklyscheduleparams/serversideprocessingtimeout
func (m_ MTRThermostatClusterGetWeeklyScheduleParams) SetServerSideProcessingTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclustergetweeklyscheduleparams/modetoreturn
func (m_ MTRThermostatClusterGetWeeklyScheduleParams) ModeToReturn() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("modeToReturn"))
	return rv
}


// SetModeToReturn sets the value of the modeToReturn property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclustergetweeklyscheduleparams/modetoreturn
func (m_ MTRThermostatClusterGetWeeklyScheduleParams) SetModeToReturn(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setModeToReturn:"), value)
}



