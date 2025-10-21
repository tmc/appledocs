// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRDishwasherAlarmClusterModifyEnabledAlarmsParams] class.
var (
	MTRDishwasherAlarmClusterModifyEnabledAlarmsParamsClass     _MTRDishwasherAlarmClusterModifyEnabledAlarmsParamsClass
	MTRDishwasherAlarmClusterModifyEnabledAlarmsParamsClassOnce sync.Once
)

func getMTRDishwasherAlarmClusterModifyEnabledAlarmsParamsClass() _MTRDishwasherAlarmClusterModifyEnabledAlarmsParamsClass {
	MTRDishwasherAlarmClusterModifyEnabledAlarmsParamsClassOnce.Do(func() {
		MTRDishwasherAlarmClusterModifyEnabledAlarmsParamsClass = _MTRDishwasherAlarmClusterModifyEnabledAlarmsParamsClass{objc.GetClass("MTRDishwasherAlarmClusterModifyEnabledAlarmsParams")}
	})
	return MTRDishwasherAlarmClusterModifyEnabledAlarmsParamsClass
}

type _MTRDishwasherAlarmClusterModifyEnabledAlarmsParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRDishwasherAlarmClusterModifyEnabledAlarmsParams] class.
type IMTRDishwasherAlarmClusterModifyEnabledAlarmsParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherAlarmClusterModifyEnabledAlarmsParams
type MTRDishwasherAlarmClusterModifyEnabledAlarmsParams struct {
	objectivec.Object
}

// MTRDishwasherAlarmClusterModifyEnabledAlarmsParamsFrom constructs a [MTRDishwasherAlarmClusterModifyEnabledAlarmsParams] from an unsafe.Pointer.
func MTRDishwasherAlarmClusterModifyEnabledAlarmsParamsFrom(ptr unsafe.Pointer) MTRDishwasherAlarmClusterModifyEnabledAlarmsParams {
	return MTRDishwasherAlarmClusterModifyEnabledAlarmsParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDishwasherAlarmClusterModifyEnabledAlarmsParamsClass) Alloc() MTRDishwasherAlarmClusterModifyEnabledAlarmsParams {
	rv := objc.Send[MTRDishwasherAlarmClusterModifyEnabledAlarmsParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDishwasherAlarmClusterModifyEnabledAlarmsParamsClass) New() MTRDishwasherAlarmClusterModifyEnabledAlarmsParams {
	rv := objc.Send[MTRDishwasherAlarmClusterModifyEnabledAlarmsParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDishwasherAlarmClusterModifyEnabledAlarmsParams) Init() MTRDishwasherAlarmClusterModifyEnabledAlarmsParams {
	rv := objc.Send[MTRDishwasherAlarmClusterModifyEnabledAlarmsParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDishwasherAlarmClusterModifyEnabledAlarmsParams) Autorelease() MTRDishwasherAlarmClusterModifyEnabledAlarmsParams {
	rv := objc.Send[MTRDishwasherAlarmClusterModifyEnabledAlarmsParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDishwasherAlarmClusterModifyEnabledAlarmsParams creates a new MTRDishwasherAlarmClusterModifyEnabledAlarmsParams instance.
func NewMTRDishwasherAlarmClusterModifyEnabledAlarmsParams() MTRDishwasherAlarmClusterModifyEnabledAlarmsParams {
	return getMTRDishwasherAlarmClusterModifyEnabledAlarmsParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherAlarmClusterModifyEnabledAlarmsParams/mask
func (m_ MTRDishwasherAlarmClusterModifyEnabledAlarmsParams) Mask() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("mask"))
	return rv
}


// SetMask sets the value of the mask property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherAlarmClusterModifyEnabledAlarmsParams/mask
func (m_ MTRDishwasherAlarmClusterModifyEnabledAlarmsParams) SetMask(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMask:"), value)
}

// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherAlarmClusterModifyEnabledAlarmsParams/serverSideProcessingTimeout
func (m_ MTRDishwasherAlarmClusterModifyEnabledAlarmsParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
// Controls how much time, in seconds, we will allow for the server to process the command.

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherAlarmClusterModifyEnabledAlarmsParams/serverSideProcessingTimeout
func (m_ MTRDishwasherAlarmClusterModifyEnabledAlarmsParams) SetServerSideProcessingTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherAlarmClusterModifyEnabledAlarmsParams/timedInvokeTimeoutMs
func (m_ MTRDishwasherAlarmClusterModifyEnabledAlarmsParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
// Controls whether the command is a timed command (using Timed Invoke).

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherAlarmClusterModifyEnabledAlarmsParams/timedInvokeTimeoutMs
func (m_ MTRDishwasherAlarmClusterModifyEnabledAlarmsParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



