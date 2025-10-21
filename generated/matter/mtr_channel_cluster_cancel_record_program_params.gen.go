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
}

//
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


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterCancelRecordProgramParams/data
func (m_ MTRChannelClusterCancelRecordProgramParams) Data() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("data"))
	return rv
}


// SetData sets the value of the data property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterCancelRecordProgramParams/data
func (m_ MTRChannelClusterCancelRecordProgramParams) SetData(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setData:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterCancelRecordProgramParams/programIdentifier
func (m_ MTRChannelClusterCancelRecordProgramParams) ProgramIdentifier() string {
	rv := objc.Send[string](m_.ID, objc.Sel("programIdentifier"))
	return rv
}


// SetProgramIdentifier sets the value of the programIdentifier property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterCancelRecordProgramParams/programIdentifier
func (m_ MTRChannelClusterCancelRecordProgramParams) SetProgramIdentifier(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProgramIdentifier:"), objc.String(value))
}

// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterCancelRecordProgramParams/serverSideProcessingTimeout
func (m_ MTRChannelClusterCancelRecordProgramParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
// Controls how much time, in seconds, we will allow for the server to process the command.

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterCancelRecordProgramParams/serverSideProcessingTimeout
func (m_ MTRChannelClusterCancelRecordProgramParams) SetServerSideProcessingTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterCancelRecordProgramParams/shouldRecordSeries
func (m_ MTRChannelClusterCancelRecordProgramParams) ShouldRecordSeries() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("shouldRecordSeries"))
	return rv
}


// SetShouldRecordSeries sets the value of the shouldRecordSeries property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterCancelRecordProgramParams/shouldRecordSeries
func (m_ MTRChannelClusterCancelRecordProgramParams) SetShouldRecordSeries(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShouldRecordSeries:"), value)
}

// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterCancelRecordProgramParams/timedInvokeTimeoutMs
func (m_ MTRChannelClusterCancelRecordProgramParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
// Controls whether the command is a timed command (using Timed Invoke).

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterCancelRecordProgramParams/timedInvokeTimeoutMs
func (m_ MTRChannelClusterCancelRecordProgramParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



