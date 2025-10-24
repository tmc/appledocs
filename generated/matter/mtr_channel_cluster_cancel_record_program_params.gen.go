// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRChannelClusterCancelRecordProgramParams] class.
var (
	MTRChannelClusterCancelRecordProgramParamsClass     _MTRChannelClusterCancelRecordProgramParamsClass
	MTRChannelClusterCancelRecordProgramParamsClassOnce sync.Once
)

func getMTRChannelClusterCancelRecordProgramParamsClass() _MTRChannelClusterCancelRecordProgramParamsClass {
	MTRChannelClusterCancelRecordProgramParamsClassOnce.Do(func() {
		MTRChannelClusterCancelRecordProgramParamsClass = _MTRChannelClusterCancelRecordProgramParamsClass{objc.GetClass("MTRChannelClusterCancelRecordProgramParams")}
	})
	return MTRChannelClusterCancelRecordProgramParamsClass
}

type _MTRChannelClusterCancelRecordProgramParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRChannelClusterCancelRecordProgramParams] class.
type IMTRChannelClusterCancelRecordProgramParams interface {
	objectivec.IObject
	// properties:
	Data() objc.IObject /* cross-framework: NSData */
	SetData(value objc.IObject /* cross-framework: NSData */)
	ProgramIdentifier() objc.IObject /* cross-framework: NSString */
	SetProgramIdentifier(value objc.IObject /* cross-framework: NSString */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	ShouldRecordSeries() objc.IObject /* cross-framework: NSNumber */
	SetShouldRecordSeries(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterCancelRecordProgramParams
type MTRChannelClusterCancelRecordProgramParams struct {
	objectivec.Object
}

// MTRChannelClusterCancelRecordProgramParamsFrom constructs a [MTRChannelClusterCancelRecordProgramParams] from an unsafe.Pointer.
func MTRChannelClusterCancelRecordProgramParamsFrom(ptr unsafe.Pointer) MTRChannelClusterCancelRecordProgramParams {
	return MTRChannelClusterCancelRecordProgramParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRChannelClusterCancelRecordProgramParamsClass) Alloc() MTRChannelClusterCancelRecordProgramParams {
	rv := objc.Send[MTRChannelClusterCancelRecordProgramParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRChannelClusterCancelRecordProgramParamsClass) New() MTRChannelClusterCancelRecordProgramParams {
	rv := objc.Send[MTRChannelClusterCancelRecordProgramParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRChannelClusterCancelRecordProgramParams) Init() MTRChannelClusterCancelRecordProgramParams {
	rv := objc.Send[MTRChannelClusterCancelRecordProgramParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRChannelClusterCancelRecordProgramParams) Autorelease() MTRChannelClusterCancelRecordProgramParams {
	rv := objc.Send[MTRChannelClusterCancelRecordProgramParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRChannelClusterCancelRecordProgramParams creates a new MTRChannelClusterCancelRecordProgramParams instance.
func NewMTRChannelClusterCancelRecordProgramParams() MTRChannelClusterCancelRecordProgramParams {
	return getMTRChannelClusterCancelRecordProgramParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterCancelRecordProgramParams/data
func (m_ MTRChannelClusterCancelRecordProgramParams) Data() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("data"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterCancelRecordProgramParams/data
func (m_ MTRChannelClusterCancelRecordProgramParams) SetData(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setData:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterCancelRecordProgramParams/programIdentifier
func (m_ MTRChannelClusterCancelRecordProgramParams) ProgramIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("programIdentifier"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterCancelRecordProgramParams/programIdentifier
func (m_ MTRChannelClusterCancelRecordProgramParams) SetProgramIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProgramIdentifier:"), value)
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterCancelRecordProgramParams/serverSideProcessingTimeout
func (m_ MTRChannelClusterCancelRecordProgramParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterCancelRecordProgramParams/serverSideProcessingTimeout
func (m_ MTRChannelClusterCancelRecordProgramParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterCancelRecordProgramParams/shouldRecordSeries
func (m_ MTRChannelClusterCancelRecordProgramParams) ShouldRecordSeries() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("shouldRecordSeries"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterCancelRecordProgramParams/shouldRecordSeries
func (m_ MTRChannelClusterCancelRecordProgramParams) SetShouldRecordSeries(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShouldRecordSeries:"), value)
}


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterCancelRecordProgramParams/timedInvokeTimeoutMs
func (m_ MTRChannelClusterCancelRecordProgramParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterCancelRecordProgramParams/timedInvokeTimeoutMs
func (m_ MTRChannelClusterCancelRecordProgramParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



