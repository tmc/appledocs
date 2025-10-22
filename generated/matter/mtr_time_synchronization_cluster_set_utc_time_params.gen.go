// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRTimeSynchronizationClusterSetUtcTimeParams] class.
var (
	MTRTimeSynchronizationClusterSetUtcTimeParamsClass     _MTRTimeSynchronizationClusterSetUtcTimeParamsClass
	MTRTimeSynchronizationClusterSetUtcTimeParamsClassOnce sync.Once
)

func getMTRTimeSynchronizationClusterSetUtcTimeParamsClass() _MTRTimeSynchronizationClusterSetUtcTimeParamsClass {
	MTRTimeSynchronizationClusterSetUtcTimeParamsClassOnce.Do(func() {
		MTRTimeSynchronizationClusterSetUtcTimeParamsClass = _MTRTimeSynchronizationClusterSetUtcTimeParamsClass{objc.GetClass("MTRTimeSynchronizationClusterSetUtcTimeParams")}
	})
	return MTRTimeSynchronizationClusterSetUtcTimeParamsClass
}

type _MTRTimeSynchronizationClusterSetUtcTimeParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRTimeSynchronizationClusterSetUtcTimeParams] class.
type IMTRTimeSynchronizationClusterSetUtcTimeParams interface {
	IMTRTimeSynchronizationClusterSetUTCTimeParams
	Granularity() foundation.Number
	SetGranularity(value foundation.INumber)
	ServerSideProcessingTimeout() foundation.Number
	SetServerSideProcessingTimeout(value foundation.INumber)
	TimeSource() foundation.Number
	SetTimeSource(value foundation.INumber)
	TimedInvokeTimeoutMs() foundation.Number
	SetTimedInvokeTimeoutMs(value foundation.INumber)
	UtcTime() foundation.Number
	SetUtcTime(value foundation.INumber)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterSetUtcTimeParams-2ms2i
type MTRTimeSynchronizationClusterSetUtcTimeParams struct {
	MTRTimeSynchronizationClusterSetUTCTimeParams
}

// MTRTimeSynchronizationClusterSetUtcTimeParamsFrom constructs a [MTRTimeSynchronizationClusterSetUtcTimeParams] from an unsafe.Pointer.
func MTRTimeSynchronizationClusterSetUtcTimeParamsFrom(ptr unsafe.Pointer) MTRTimeSynchronizationClusterSetUtcTimeParams {
	return MTRTimeSynchronizationClusterSetUtcTimeParams{
		MTRTimeSynchronizationClusterSetUTCTimeParams: MTRTimeSynchronizationClusterSetUTCTimeParamsFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTimeSynchronizationClusterSetUtcTimeParamsClass) Alloc() MTRTimeSynchronizationClusterSetUtcTimeParams {
	rv := objc.Send[MTRTimeSynchronizationClusterSetUtcTimeParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTimeSynchronizationClusterSetUtcTimeParamsClass) New() MTRTimeSynchronizationClusterSetUtcTimeParams {
	rv := objc.Send[MTRTimeSynchronizationClusterSetUtcTimeParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTimeSynchronizationClusterSetUtcTimeParams) Init() MTRTimeSynchronizationClusterSetUtcTimeParams {
	rv := objc.Send[MTRTimeSynchronizationClusterSetUtcTimeParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTimeSynchronizationClusterSetUtcTimeParams) Autorelease() MTRTimeSynchronizationClusterSetUtcTimeParams {
	rv := objc.Send[MTRTimeSynchronizationClusterSetUtcTimeParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTimeSynchronizationClusterSetUtcTimeParams creates a new MTRTimeSynchronizationClusterSetUtcTimeParams instance.
func NewMTRTimeSynchronizationClusterSetUtcTimeParams() MTRTimeSynchronizationClusterSetUtcTimeParams {
	return getMTRTimeSynchronizationClusterSetUtcTimeParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtimesynchronizationclustersetutctimeparams-2ms2i/granularity
func (m_ MTRTimeSynchronizationClusterSetUtcTimeParams) Granularity() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("granularity"))
	return rv
}


// SetGranularity sets the value of the granularity property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtimesynchronizationclustersetutctimeparams-2ms2i/granularity
func (m_ MTRTimeSynchronizationClusterSetUtcTimeParams) SetGranularity(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGranularity:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtimesynchronizationclustersetutctimeparams-2ms2i/serversideprocessingtimeout
func (m_ MTRTimeSynchronizationClusterSetUtcTimeParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtimesynchronizationclustersetutctimeparams-2ms2i/serversideprocessingtimeout
func (m_ MTRTimeSynchronizationClusterSetUtcTimeParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtimesynchronizationclustersetutctimeparams-2ms2i/timesource
func (m_ MTRTimeSynchronizationClusterSetUtcTimeParams) TimeSource() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timeSource"))
	return rv
}


// SetTimeSource sets the value of the timeSource property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtimesynchronizationclustersetutctimeparams-2ms2i/timesource
func (m_ MTRTimeSynchronizationClusterSetUtcTimeParams) SetTimeSource(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimeSource:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtimesynchronizationclustersetutctimeparams-2ms2i/timedinvoketimeoutms
func (m_ MTRTimeSynchronizationClusterSetUtcTimeParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtimesynchronizationclustersetutctimeparams-2ms2i/timedinvoketimeoutms
func (m_ MTRTimeSynchronizationClusterSetUtcTimeParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtimesynchronizationclustersetutctimeparams-2ms2i/utctime
func (m_ MTRTimeSynchronizationClusterSetUtcTimeParams) UtcTime() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("utcTime"))
	return rv
}


// SetUtcTime sets the value of the utcTime property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtimesynchronizationclustersetutctimeparams-2ms2i/utctime
func (m_ MTRTimeSynchronizationClusterSetUtcTimeParams) SetUtcTime(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUtcTime:"), value)
}



