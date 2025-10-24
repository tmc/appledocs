// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRTimeSynchronizationClusterSetTimeZoneParams] class.
var (
	MTRTimeSynchronizationClusterSetTimeZoneParamsClass     _MTRTimeSynchronizationClusterSetTimeZoneParamsClass
	MTRTimeSynchronizationClusterSetTimeZoneParamsClassOnce sync.Once
)

func getMTRTimeSynchronizationClusterSetTimeZoneParamsClass() _MTRTimeSynchronizationClusterSetTimeZoneParamsClass {
	MTRTimeSynchronizationClusterSetTimeZoneParamsClassOnce.Do(func() {
		MTRTimeSynchronizationClusterSetTimeZoneParamsClass = _MTRTimeSynchronizationClusterSetTimeZoneParamsClass{objc.GetClass("MTRTimeSynchronizationClusterSetTimeZoneParams")}
	})
	return MTRTimeSynchronizationClusterSetTimeZoneParamsClass
}

type _MTRTimeSynchronizationClusterSetTimeZoneParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRTimeSynchronizationClusterSetTimeZoneParams] class.
type IMTRTimeSynchronizationClusterSetTimeZoneParams interface {
	objectivec.IObject
	// properties:
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimeZone() objc.IObject /* cross-framework: NSArray */
	SetTimeZone(value objc.IObject /* cross-framework: NSArray */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterSetTimeZoneParams
type MTRTimeSynchronizationClusterSetTimeZoneParams struct {
	objectivec.Object
}

// MTRTimeSynchronizationClusterSetTimeZoneParamsFrom constructs a [MTRTimeSynchronizationClusterSetTimeZoneParams] from an unsafe.Pointer.
func MTRTimeSynchronizationClusterSetTimeZoneParamsFrom(ptr unsafe.Pointer) MTRTimeSynchronizationClusterSetTimeZoneParams {
	return MTRTimeSynchronizationClusterSetTimeZoneParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTimeSynchronizationClusterSetTimeZoneParamsClass) Alloc() MTRTimeSynchronizationClusterSetTimeZoneParams {
	rv := objc.Send[MTRTimeSynchronizationClusterSetTimeZoneParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTimeSynchronizationClusterSetTimeZoneParamsClass) New() MTRTimeSynchronizationClusterSetTimeZoneParams {
	rv := objc.Send[MTRTimeSynchronizationClusterSetTimeZoneParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTimeSynchronizationClusterSetTimeZoneParams) Init() MTRTimeSynchronizationClusterSetTimeZoneParams {
	rv := objc.Send[MTRTimeSynchronizationClusterSetTimeZoneParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTimeSynchronizationClusterSetTimeZoneParams) Autorelease() MTRTimeSynchronizationClusterSetTimeZoneParams {
	rv := objc.Send[MTRTimeSynchronizationClusterSetTimeZoneParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTimeSynchronizationClusterSetTimeZoneParams creates a new MTRTimeSynchronizationClusterSetTimeZoneParams instance.
func NewMTRTimeSynchronizationClusterSetTimeZoneParams() MTRTimeSynchronizationClusterSetTimeZoneParams {
	return getMTRTimeSynchronizationClusterSetTimeZoneParamsClass().New()
}



// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterSetTimeZoneParams/serverSideProcessingTimeout
func (m_ MTRTimeSynchronizationClusterSetTimeZoneParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterSetTimeZoneParams/serverSideProcessingTimeout
func (m_ MTRTimeSynchronizationClusterSetTimeZoneParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterSetTimeZoneParams/timeZone
func (m_ MTRTimeSynchronizationClusterSetTimeZoneParams) TimeZone() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("timeZone"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterSetTimeZoneParams/timeZone
func (m_ MTRTimeSynchronizationClusterSetTimeZoneParams) SetTimeZone(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimeZone:"), value)
}


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterSetTimeZoneParams/timedInvokeTimeoutMs
func (m_ MTRTimeSynchronizationClusterSetTimeZoneParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterSetTimeZoneParams/timedInvokeTimeoutMs
func (m_ MTRTimeSynchronizationClusterSetTimeZoneParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



