// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParams] class.
var (
	MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParamsClass     _MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParamsClass
	MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParamsClassOnce sync.Once
)

func getMTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParamsClass() _MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParamsClass {
	MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParamsClassOnce.Do(func() {
		MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParamsClass = _MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParamsClass{objc.GetClass("MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParams")}
	})
	return MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParamsClass
}

type _MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParams] class.
type IMTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParams
type MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParams struct {
	objectivec.Object
}

// MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParamsFrom constructs a [MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParams] from an unsafe.Pointer.
func MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParamsFrom(ptr unsafe.Pointer) MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParams {
	return MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParamsClass) Alloc() MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParams {
	rv := objc.Send[MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParamsClass) New() MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParams {
	rv := objc.Send[MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParams) Init() MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParams {
	rv := objc.Send[MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParams) Autorelease() MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParams {
	rv := objc.Send[MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParams creates a new MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParams instance.
func NewMTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParams() MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParams {
	return getMTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalmeasurementclustergetmeasurementprofileresponsecommandparams/attributeid
func (m_ MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParams) AttributeId() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("attributeId"))
	return rv
}


// SetAttributeId sets the value of the attributeId property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalmeasurementclustergetmeasurementprofileresponsecommandparams/attributeid
func (m_ MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParams) SetAttributeId(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAttributeId:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalmeasurementclustergetmeasurementprofileresponsecommandparams/intervals
func (m_ MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParams) Intervals() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("intervals"))
	return rv
}


// SetIntervals sets the value of the intervals property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalmeasurementclustergetmeasurementprofileresponsecommandparams/intervals
func (m_ MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParams) SetIntervals(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIntervals:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalmeasurementclustergetmeasurementprofileresponsecommandparams/numberofintervalsdelivered
func (m_ MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParams) NumberOfIntervalsDelivered() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("numberOfIntervalsDelivered"))
	return rv
}


// SetNumberOfIntervalsDelivered sets the value of the numberOfIntervalsDelivered property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalmeasurementclustergetmeasurementprofileresponsecommandparams/numberofintervalsdelivered
func (m_ MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParams) SetNumberOfIntervalsDelivered(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNumberOfIntervalsDelivered:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalmeasurementclustergetmeasurementprofileresponsecommandparams/profileintervalperiod
func (m_ MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParams) ProfileIntervalPeriod() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("profileIntervalPeriod"))
	return rv
}


// SetProfileIntervalPeriod sets the value of the profileIntervalPeriod property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalmeasurementclustergetmeasurementprofileresponsecommandparams/profileintervalperiod
func (m_ MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParams) SetProfileIntervalPeriod(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProfileIntervalPeriod:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalmeasurementclustergetmeasurementprofileresponsecommandparams/starttime
func (m_ MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParams) StartTime() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("startTime"))
	return rv
}


// SetStartTime sets the value of the startTime property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalmeasurementclustergetmeasurementprofileresponsecommandparams/starttime
func (m_ MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParams) SetStartTime(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStartTime:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalmeasurementclustergetmeasurementprofileresponsecommandparams/status
func (m_ MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParams) Status() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("status"))
	return rv
}


// SetStatus sets the value of the status property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalmeasurementclustergetmeasurementprofileresponsecommandparams/status
func (m_ MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParams) SetStatus(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatus:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalmeasurementclustergetmeasurementprofileresponsecommandparams/timedinvoketimeoutms
func (m_ MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalmeasurementclustergetmeasurementprofileresponsecommandparams/timedinvoketimeoutms
func (m_ MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



