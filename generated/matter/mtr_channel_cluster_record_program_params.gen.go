// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
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
}

//
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


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterRecordProgramParams/data
func (m_ MTRChannelClusterRecordProgramParams) Data() foundation.NSData {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("data"))
	return rv
}


// SetData sets the value of the data property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterRecordProgramParams/data
func (m_ MTRChannelClusterRecordProgramParams) SetData(value foundation.IData) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setData:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterRecordProgramParams/programIdentifier
func (m_ MTRChannelClusterRecordProgramParams) ProgramIdentifier() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("programIdentifier"))
	return rv
}


// SetProgramIdentifier sets the value of the programIdentifier property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterRecordProgramParams/programIdentifier
func (m_ MTRChannelClusterRecordProgramParams) SetProgramIdentifier(value appkit.string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProgramIdentifier:"), value)
}

// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterRecordProgramParams/serverSideProcessingTimeout
func (m_ MTRChannelClusterRecordProgramParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
// Controls how much time, in seconds, we will allow for the server to process the command.

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterRecordProgramParams/serverSideProcessingTimeout
func (m_ MTRChannelClusterRecordProgramParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterRecordProgramParams/shouldRecordSeries
func (m_ MTRChannelClusterRecordProgramParams) ShouldRecordSeries() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("shouldRecordSeries"))
	return rv
}


// SetShouldRecordSeries sets the value of the shouldRecordSeries property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterRecordProgramParams/shouldRecordSeries
func (m_ MTRChannelClusterRecordProgramParams) SetShouldRecordSeries(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShouldRecordSeries:"), value)
}

// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterRecordProgramParams/timedInvokeTimeoutMs
func (m_ MTRChannelClusterRecordProgramParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
// Controls whether the command is a timed command (using Timed Invoke).

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterRecordProgramParams/timedInvokeTimeoutMs
func (m_ MTRChannelClusterRecordProgramParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



