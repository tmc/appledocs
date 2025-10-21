// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams] class.
var (
	MTRElectricalMeasurementClusterGetMeasurementProfileCommandParamsClass     _MTRElectricalMeasurementClusterGetMeasurementProfileCommandParamsClass
	MTRElectricalMeasurementClusterGetMeasurementProfileCommandParamsClassOnce sync.Once
)

func getMTRElectricalMeasurementClusterGetMeasurementProfileCommandParamsClass() _MTRElectricalMeasurementClusterGetMeasurementProfileCommandParamsClass {
	MTRElectricalMeasurementClusterGetMeasurementProfileCommandParamsClassOnce.Do(func() {
		MTRElectricalMeasurementClusterGetMeasurementProfileCommandParamsClass = _MTRElectricalMeasurementClusterGetMeasurementProfileCommandParamsClass{objc.GetClass("MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams")}
	})
	return MTRElectricalMeasurementClusterGetMeasurementProfileCommandParamsClass
}

type _MTRElectricalMeasurementClusterGetMeasurementProfileCommandParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams] class.
type IMTRElectricalMeasurementClusterGetMeasurementProfileCommandParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams
type MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams struct {
	objectivec.Object
}

// MTRElectricalMeasurementClusterGetMeasurementProfileCommandParamsFrom constructs a [MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams] from an unsafe.Pointer.
func MTRElectricalMeasurementClusterGetMeasurementProfileCommandParamsFrom(ptr unsafe.Pointer) MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams {
	return MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRElectricalMeasurementClusterGetMeasurementProfileCommandParamsClass) Alloc() MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams {
	rv := objc.Send[MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRElectricalMeasurementClusterGetMeasurementProfileCommandParamsClass) New() MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams {
	rv := objc.Send[MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams) Init() MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams {
	rv := objc.Send[MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams) Autorelease() MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams {
	rv := objc.Send[MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRElectricalMeasurementClusterGetMeasurementProfileCommandParams creates a new MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams instance.
func NewMTRElectricalMeasurementClusterGetMeasurementProfileCommandParams() MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams {
	return getMTRElectricalMeasurementClusterGetMeasurementProfileCommandParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalmeasurementclustergetmeasurementprofilecommandparams/attributeid
func (m_ MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams) AttributeId() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("attributeId"))
	return rv
}


// SetAttributeId sets the value of the attributeId property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalmeasurementclustergetmeasurementprofilecommandparams/attributeid
func (m_ MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams) SetAttributeId(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAttributeId:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalmeasurementclustergetmeasurementprofilecommandparams/numberofintervals
func (m_ MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams) NumberOfIntervals() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("numberOfIntervals"))
	return rv
}


// SetNumberOfIntervals sets the value of the numberOfIntervals property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalmeasurementclustergetmeasurementprofilecommandparams/numberofintervals
func (m_ MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams) SetNumberOfIntervals(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNumberOfIntervals:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalmeasurementclustergetmeasurementprofilecommandparams/serversideprocessingtimeout
func (m_ MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalmeasurementclustergetmeasurementprofilecommandparams/serversideprocessingtimeout
func (m_ MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalmeasurementclustergetmeasurementprofilecommandparams/starttime
func (m_ MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams) StartTime() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("startTime"))
	return rv
}


// SetStartTime sets the value of the startTime property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalmeasurementclustergetmeasurementprofilecommandparams/starttime
func (m_ MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams) SetStartTime(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStartTime:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalmeasurementclustergetmeasurementprofilecommandparams/timedinvoketimeoutms
func (m_ MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalmeasurementclustergetmeasurementprofilecommandparams/timedinvoketimeoutms
func (m_ MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



