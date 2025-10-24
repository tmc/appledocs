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
	// properties:
	Granularity() objc.IObject /* cross-framework: NSNumber */
	SetGranularity(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimeSource() objc.IObject /* cross-framework: NSNumber */
	SetTimeSource(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	UtcTime() objc.IObject /* cross-framework: NSNumber */
	SetUtcTime(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtimesynchronizationclustersetutctimeparams-2ms2i/granularity
func (m_ MTRTimeSynchronizationClusterSetUtcTimeParams) Granularity() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("granularity"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtimesynchronizationclustersetutctimeparams-2ms2i/granularity
func (m_ MTRTimeSynchronizationClusterSetUtcTimeParams) SetGranularity(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGranularity:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtimesynchronizationclustersetutctimeparams-2ms2i/serversideprocessingtimeout
func (m_ MTRTimeSynchronizationClusterSetUtcTimeParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtimesynchronizationclustersetutctimeparams-2ms2i/serversideprocessingtimeout
func (m_ MTRTimeSynchronizationClusterSetUtcTimeParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtimesynchronizationclustersetutctimeparams-2ms2i/timesource
func (m_ MTRTimeSynchronizationClusterSetUtcTimeParams) TimeSource() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timeSource"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtimesynchronizationclustersetutctimeparams-2ms2i/timesource
func (m_ MTRTimeSynchronizationClusterSetUtcTimeParams) SetTimeSource(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimeSource:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtimesynchronizationclustersetutctimeparams-2ms2i/timedinvoketimeoutms
func (m_ MTRTimeSynchronizationClusterSetUtcTimeParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtimesynchronizationclustersetutctimeparams-2ms2i/timedinvoketimeoutms
func (m_ MTRTimeSynchronizationClusterSetUtcTimeParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtimesynchronizationclustersetutctimeparams-2ms2i/utctime
func (m_ MTRTimeSynchronizationClusterSetUtcTimeParams) UtcTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("utcTime"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtimesynchronizationclustersetutctimeparams-2ms2i/utctime
func (m_ MTRTimeSynchronizationClusterSetUtcTimeParams) SetUtcTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUtcTime:"), value)
}



