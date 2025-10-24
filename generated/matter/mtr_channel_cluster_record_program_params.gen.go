// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRChannelClusterRecordProgramParams] class.
var (
	MTRChannelClusterRecordProgramParamsClass     _MTRChannelClusterRecordProgramParamsClass
	MTRChannelClusterRecordProgramParamsClassOnce sync.Once
)

func getMTRChannelClusterRecordProgramParamsClass() _MTRChannelClusterRecordProgramParamsClass {
	MTRChannelClusterRecordProgramParamsClassOnce.Do(func() {
		MTRChannelClusterRecordProgramParamsClass = _MTRChannelClusterRecordProgramParamsClass{objc.GetClass("MTRChannelClusterRecordProgramParams")}
	})
	return MTRChannelClusterRecordProgramParamsClass
}

type _MTRChannelClusterRecordProgramParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRChannelClusterRecordProgramParams] class.
type IMTRChannelClusterRecordProgramParams interface {
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
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterRecordProgramParams
type MTRChannelClusterRecordProgramParams struct {
	objectivec.Object
}

// MTRChannelClusterRecordProgramParamsFrom constructs a [MTRChannelClusterRecordProgramParams] from an unsafe.Pointer.
func MTRChannelClusterRecordProgramParamsFrom(ptr unsafe.Pointer) MTRChannelClusterRecordProgramParams {
	return MTRChannelClusterRecordProgramParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRChannelClusterRecordProgramParamsClass) Alloc() MTRChannelClusterRecordProgramParams {
	rv := objc.Send[MTRChannelClusterRecordProgramParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRChannelClusterRecordProgramParamsClass) New() MTRChannelClusterRecordProgramParams {
	rv := objc.Send[MTRChannelClusterRecordProgramParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRChannelClusterRecordProgramParams) Init() MTRChannelClusterRecordProgramParams {
	rv := objc.Send[MTRChannelClusterRecordProgramParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRChannelClusterRecordProgramParams) Autorelease() MTRChannelClusterRecordProgramParams {
	rv := objc.Send[MTRChannelClusterRecordProgramParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRChannelClusterRecordProgramParams creates a new MTRChannelClusterRecordProgramParams instance.
func NewMTRChannelClusterRecordProgramParams() MTRChannelClusterRecordProgramParams {
	return getMTRChannelClusterRecordProgramParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterRecordProgramParams/data
func (m_ MTRChannelClusterRecordProgramParams) Data() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("data"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterRecordProgramParams/data
func (m_ MTRChannelClusterRecordProgramParams) SetData(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setData:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterRecordProgramParams/programIdentifier
func (m_ MTRChannelClusterRecordProgramParams) ProgramIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("programIdentifier"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterRecordProgramParams/programIdentifier
func (m_ MTRChannelClusterRecordProgramParams) SetProgramIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProgramIdentifier:"), value)
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterRecordProgramParams/serverSideProcessingTimeout
func (m_ MTRChannelClusterRecordProgramParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterRecordProgramParams/serverSideProcessingTimeout
func (m_ MTRChannelClusterRecordProgramParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterRecordProgramParams/shouldRecordSeries
func (m_ MTRChannelClusterRecordProgramParams) ShouldRecordSeries() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("shouldRecordSeries"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterRecordProgramParams/shouldRecordSeries
func (m_ MTRChannelClusterRecordProgramParams) SetShouldRecordSeries(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShouldRecordSeries:"), value)
}


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterRecordProgramParams/timedInvokeTimeoutMs
func (m_ MTRChannelClusterRecordProgramParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterRecordProgramParams/timedInvokeTimeoutMs
func (m_ MTRChannelClusterRecordProgramParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



