// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	// properties:
	ChannelList() objc.IObject /* cross-framework: NSArray */
	SetChannelList(value objc.IObject /* cross-framework: NSArray */)
	Data() objc.IObject /* cross-framework: NSData */
	SetData(value objc.IObject /* cross-framework: NSData */)
	EndTime() objc.IObject /* cross-framework: NSNumber */
	SetEndTime(value objc.IObject /* cross-framework: NSNumber */)
	PageToken() IMTRChannelClusterPageTokenStruct
	SetPageToken(value IMTRChannelClusterPageTokenStruct)
	RecordingFlag() objc.IObject /* cross-framework: NSNumber */
	SetRecordingFlag(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	StartTime() objc.IObject /* cross-framework: NSNumber */
	SetStartTime(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterGetProgramGuideParams/channelList
func (m_ MTRChannelClusterGetProgramGuideParams) ChannelList() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("channelList"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterGetProgramGuideParams/channelList
func (m_ MTRChannelClusterGetProgramGuideParams) SetChannelList(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setChannelList:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterGetProgramGuideParams/data
func (m_ MTRChannelClusterGetProgramGuideParams) Data() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("data"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterGetProgramGuideParams/data
func (m_ MTRChannelClusterGetProgramGuideParams) SetData(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setData:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterGetProgramGuideParams/endTime
func (m_ MTRChannelClusterGetProgramGuideParams) EndTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("endTime"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterGetProgramGuideParams/endTime
func (m_ MTRChannelClusterGetProgramGuideParams) SetEndTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEndTime:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterGetProgramGuideParams/pageToken
func (m_ MTRChannelClusterGetProgramGuideParams) PageToken() IMTRChannelClusterPageTokenStruct {
	rv := objc.Send[MTRChannelClusterPageTokenStruct](m_.ID, objc.Sel("pageToken"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterGetProgramGuideParams/pageToken
func (m_ MTRChannelClusterGetProgramGuideParams) SetPageToken(value IMTRChannelClusterPageTokenStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPageToken:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterGetProgramGuideParams/recordingFlag
func (m_ MTRChannelClusterGetProgramGuideParams) RecordingFlag() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("recordingFlag"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterGetProgramGuideParams/recordingFlag
func (m_ MTRChannelClusterGetProgramGuideParams) SetRecordingFlag(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRecordingFlag:"), value)
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterGetProgramGuideParams/serverSideProcessingTimeout
func (m_ MTRChannelClusterGetProgramGuideParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterGetProgramGuideParams/serverSideProcessingTimeout
func (m_ MTRChannelClusterGetProgramGuideParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterGetProgramGuideParams/startTime
func (m_ MTRChannelClusterGetProgramGuideParams) StartTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("startTime"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterGetProgramGuideParams/startTime
func (m_ MTRChannelClusterGetProgramGuideParams) SetStartTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStartTime:"), value)
}


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterGetProgramGuideParams/timedInvokeTimeoutMs
func (m_ MTRChannelClusterGetProgramGuideParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterGetProgramGuideParams/timedInvokeTimeoutMs
func (m_ MTRChannelClusterGetProgramGuideParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



