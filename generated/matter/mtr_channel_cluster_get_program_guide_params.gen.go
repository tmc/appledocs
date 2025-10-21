// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRChannelClusterGetProgramGuideParams] class.
var (
	MTRChannelClusterGetProgramGuideParamsClass     _MTRChannelClusterGetProgramGuideParamsClass
	MTRChannelClusterGetProgramGuideParamsClassOnce sync.Once
)

func getMTRChannelClusterGetProgramGuideParamsClass() _MTRChannelClusterGetProgramGuideParamsClass {
	MTRChannelClusterGetProgramGuideParamsClassOnce.Do(func() {
		MTRChannelClusterGetProgramGuideParamsClass = _MTRChannelClusterGetProgramGuideParamsClass{objc.GetClass("MTRChannelClusterGetProgramGuideParams")}
	})
	return MTRChannelClusterGetProgramGuideParamsClass
}

type _MTRChannelClusterGetProgramGuideParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRChannelClusterGetProgramGuideParams] class.
type IMTRChannelClusterGetProgramGuideParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterGetProgramGuideParams
type MTRChannelClusterGetProgramGuideParams struct {
	objectivec.Object
}

// MTRChannelClusterGetProgramGuideParamsFrom constructs a [MTRChannelClusterGetProgramGuideParams] from an unsafe.Pointer.
func MTRChannelClusterGetProgramGuideParamsFrom(ptr unsafe.Pointer) MTRChannelClusterGetProgramGuideParams {
	return MTRChannelClusterGetProgramGuideParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRChannelClusterGetProgramGuideParamsClass) Alloc() MTRChannelClusterGetProgramGuideParams {
	rv := objc.Send[MTRChannelClusterGetProgramGuideParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRChannelClusterGetProgramGuideParamsClass) New() MTRChannelClusterGetProgramGuideParams {
	rv := objc.Send[MTRChannelClusterGetProgramGuideParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRChannelClusterGetProgramGuideParams) Init() MTRChannelClusterGetProgramGuideParams {
	rv := objc.Send[MTRChannelClusterGetProgramGuideParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRChannelClusterGetProgramGuideParams) Autorelease() MTRChannelClusterGetProgramGuideParams {
	rv := objc.Send[MTRChannelClusterGetProgramGuideParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRChannelClusterGetProgramGuideParams creates a new MTRChannelClusterGetProgramGuideParams instance.
func NewMTRChannelClusterGetProgramGuideParams() MTRChannelClusterGetProgramGuideParams {
	return getMTRChannelClusterGetProgramGuideParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterGetProgramGuideParams/channelList
func (m_ MTRChannelClusterGetProgramGuideParams) ChannelList() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("channelList"))
	return rv
}


// SetChannelList sets the value of the channelList property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterGetProgramGuideParams/channelList
func (m_ MTRChannelClusterGetProgramGuideParams) SetChannelList(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setChannelList:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterGetProgramGuideParams/data
func (m_ MTRChannelClusterGetProgramGuideParams) Data() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("data"))
	return rv
}


// SetData sets the value of the data property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterGetProgramGuideParams/data
func (m_ MTRChannelClusterGetProgramGuideParams) SetData(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setData:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterGetProgramGuideParams/endTime
func (m_ MTRChannelClusterGetProgramGuideParams) EndTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("endTime"))
	return rv
}


// SetEndTime sets the value of the endTime property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterGetProgramGuideParams/endTime
func (m_ MTRChannelClusterGetProgramGuideParams) SetEndTime(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEndTime:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterGetProgramGuideParams/pageToken
func (m_ MTRChannelClusterGetProgramGuideParams) PageToken() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("pageToken"))
	return rv
}


// SetPageToken sets the value of the pageToken property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterGetProgramGuideParams/pageToken
func (m_ MTRChannelClusterGetProgramGuideParams) SetPageToken(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPageToken:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterGetProgramGuideParams/recordingFlag
func (m_ MTRChannelClusterGetProgramGuideParams) RecordingFlag() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("recordingFlag"))
	return rv
}


// SetRecordingFlag sets the value of the recordingFlag property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterGetProgramGuideParams/recordingFlag
func (m_ MTRChannelClusterGetProgramGuideParams) SetRecordingFlag(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRecordingFlag:"), value)
}

// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterGetProgramGuideParams/serverSideProcessingTimeout
func (m_ MTRChannelClusterGetProgramGuideParams) ServerSideProcessingTimeout() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
// Controls how much time, in seconds, we will allow for the server to process the command.

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterGetProgramGuideParams/serverSideProcessingTimeout
func (m_ MTRChannelClusterGetProgramGuideParams) SetServerSideProcessingTimeout(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterGetProgramGuideParams/startTime
func (m_ MTRChannelClusterGetProgramGuideParams) StartTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("startTime"))
	return rv
}


// SetStartTime sets the value of the startTime property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterGetProgramGuideParams/startTime
func (m_ MTRChannelClusterGetProgramGuideParams) SetStartTime(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStartTime:"), value)
}

// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterGetProgramGuideParams/timedInvokeTimeoutMs
func (m_ MTRChannelClusterGetProgramGuideParams) TimedInvokeTimeoutMs() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
// Controls whether the command is a timed command (using Timed Invoke).

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterGetProgramGuideParams/timedInvokeTimeoutMs
func (m_ MTRChannelClusterGetProgramGuideParams) SetTimedInvokeTimeoutMs(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



