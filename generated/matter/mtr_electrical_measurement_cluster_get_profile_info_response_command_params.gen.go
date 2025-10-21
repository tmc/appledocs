// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRElectricalMeasurementClusterGetProfileInfoResponseCommandParams] class.
var (
	MTRElectricalMeasurementClusterGetProfileInfoResponseCommandParamsClass     _MTRElectricalMeasurementClusterGetProfileInfoResponseCommandParamsClass
	MTRElectricalMeasurementClusterGetProfileInfoResponseCommandParamsClassOnce sync.Once
)

func getMTRElectricalMeasurementClusterGetProfileInfoResponseCommandParamsClass() _MTRElectricalMeasurementClusterGetProfileInfoResponseCommandParamsClass {
	MTRElectricalMeasurementClusterGetProfileInfoResponseCommandParamsClassOnce.Do(func() {
		MTRElectricalMeasurementClusterGetProfileInfoResponseCommandParamsClass = _MTRElectricalMeasurementClusterGetProfileInfoResponseCommandParamsClass{objc.GetClass("MTRElectricalMeasurementClusterGetProfileInfoResponseCommandParams")}
	})
	return MTRElectricalMeasurementClusterGetProfileInfoResponseCommandParamsClass
}

type _MTRElectricalMeasurementClusterGetProfileInfoResponseCommandParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRElectricalMeasurementClusterGetProfileInfoResponseCommandParams] class.
type IMTRElectricalMeasurementClusterGetProfileInfoResponseCommandParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRElectricalMeasurementClusterGetProfileInfoResponseCommandParams
type MTRElectricalMeasurementClusterGetProfileInfoResponseCommandParams struct {
	objectivec.Object
}

// MTRElectricalMeasurementClusterGetProfileInfoResponseCommandParamsFrom constructs a [MTRElectricalMeasurementClusterGetProfileInfoResponseCommandParams] from an unsafe.Pointer.
func MTRElectricalMeasurementClusterGetProfileInfoResponseCommandParamsFrom(ptr unsafe.Pointer) MTRElectricalMeasurementClusterGetProfileInfoResponseCommandParams {
	return MTRElectricalMeasurementClusterGetProfileInfoResponseCommandParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRElectricalMeasurementClusterGetProfileInfoResponseCommandParamsClass) Alloc() MTRElectricalMeasurementClusterGetProfileInfoResponseCommandParams {
	rv := objc.Send[MTRElectricalMeasurementClusterGetProfileInfoResponseCommandParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRElectricalMeasurementClusterGetProfileInfoResponseCommandParamsClass) New() MTRElectricalMeasurementClusterGetProfileInfoResponseCommandParams {
	rv := objc.Send[MTRElectricalMeasurementClusterGetProfileInfoResponseCommandParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRElectricalMeasurementClusterGetProfileInfoResponseCommandParams) Init() MTRElectricalMeasurementClusterGetProfileInfoResponseCommandParams {
	rv := objc.Send[MTRElectricalMeasurementClusterGetProfileInfoResponseCommandParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRElectricalMeasurementClusterGetProfileInfoResponseCommandParams) Autorelease() MTRElectricalMeasurementClusterGetProfileInfoResponseCommandParams {
	rv := objc.Send[MTRElectricalMeasurementClusterGetProfileInfoResponseCommandParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRElectricalMeasurementClusterGetProfileInfoResponseCommandParams creates a new MTRElectricalMeasurementClusterGetProfileInfoResponseCommandParams instance.
func NewMTRElectricalMeasurementClusterGetProfileInfoResponseCommandParams() MTRElectricalMeasurementClusterGetProfileInfoResponseCommandParams {
	return getMTRElectricalMeasurementClusterGetProfileInfoResponseCommandParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalmeasurementclustergetprofileinforesponsecommandparams/listofattributes
func (m_ MTRElectricalMeasurementClusterGetProfileInfoResponseCommandParams) ListOfAttributes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("listOfAttributes"))
	return rv
}


// SetListOfAttributes sets the value of the listOfAttributes property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalmeasurementclustergetprofileinforesponsecommandparams/listofattributes
func (m_ MTRElectricalMeasurementClusterGetProfileInfoResponseCommandParams) SetListOfAttributes(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setListOfAttributes:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalmeasurementclustergetprofileinforesponsecommandparams/maxnumberofintervals
func (m_ MTRElectricalMeasurementClusterGetProfileInfoResponseCommandParams) MaxNumberOfIntervals() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("maxNumberOfIntervals"))
	return rv
}


// SetMaxNumberOfIntervals sets the value of the maxNumberOfIntervals property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalmeasurementclustergetprofileinforesponsecommandparams/maxnumberofintervals
func (m_ MTRElectricalMeasurementClusterGetProfileInfoResponseCommandParams) SetMaxNumberOfIntervals(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxNumberOfIntervals:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalmeasurementclustergetprofileinforesponsecommandparams/profilecount
func (m_ MTRElectricalMeasurementClusterGetProfileInfoResponseCommandParams) ProfileCount() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("profileCount"))
	return rv
}


// SetProfileCount sets the value of the profileCount property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalmeasurementclustergetprofileinforesponsecommandparams/profilecount
func (m_ MTRElectricalMeasurementClusterGetProfileInfoResponseCommandParams) SetProfileCount(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProfileCount:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalmeasurementclustergetprofileinforesponsecommandparams/profileintervalperiod
func (m_ MTRElectricalMeasurementClusterGetProfileInfoResponseCommandParams) ProfileIntervalPeriod() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("profileIntervalPeriod"))
	return rv
}


// SetProfileIntervalPeriod sets the value of the profileIntervalPeriod property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalmeasurementclustergetprofileinforesponsecommandparams/profileintervalperiod
func (m_ MTRElectricalMeasurementClusterGetProfileInfoResponseCommandParams) SetProfileIntervalPeriod(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProfileIntervalPeriod:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalmeasurementclustergetprofileinforesponsecommandparams/timedinvoketimeoutms
func (m_ MTRElectricalMeasurementClusterGetProfileInfoResponseCommandParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalmeasurementclustergetprofileinforesponsecommandparams/timedinvoketimeoutms
func (m_ MTRElectricalMeasurementClusterGetProfileInfoResponseCommandParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



