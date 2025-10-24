// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRChannelClusterGetProgramGuideParams */


/* debug [class_header]: Header for MTRChannelClusterGetProgramGuideParams */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRChannelClusterGetProgramGuideParams */
// An interface definition for the [MTRChannelClusterGetProgramGuideParams] class.
type IMTRChannelClusterGetProgramGuideParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRChannelClusterGetProgramGuideParams */
	// properties:
	ChannelList() objc.IObject /* cross-framework: NSArray */
	SetChannelList(value objc.IObject /* cross-framework: NSArray */)
	Data() foundation.Data
	SetData(value foundation.Data)
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
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRChannelClusterGetProgramGuideParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRChannelClusterGetProgramGuideParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRChannelClusterGetProgramGuideParamsClass) Alloc() MTRChannelClusterGetProgramGuideParams {
	rv := objc.Send[MTRChannelClusterGetProgramGuideParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRChannelClusterGetProgramGuideParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterGetProgramGuideParams
type MTRChannelClusterGetProgramGuideParams struct {
	objectivec.Object
}

// MTRChannelClusterGetProgramGuideParamsFrom constructs a [MTRChannelClusterGetProgramGuideParams] from an unsafe.Pointer.
func MTRChannelClusterGetProgramGuideParamsFrom(ptr unsafe.Pointer) MTRChannelClusterGetProgramGuideParams {
	return MTRChannelClusterGetProgramGuideParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRChannelClusterGetProgramGuideParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRChannelClusterGetProgramGuideParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRChannelClusterGetProgramGuideParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRChannelClusterGetProgramGuideParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRChannelClusterGetProgramGuideParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterGetProgramGuideParams/channelList
func (m_ MTRChannelClusterGetProgramGuideParams) ChannelList() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("channelList"))
	return rv
}/* debug [instance_properties/getter]: channelList */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterGetProgramGuideParams/channelList
func (m_ MTRChannelClusterGetProgramGuideParams) SetChannelList(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setChannelList:"), value)
}/* debug [instance_properties/setter]: channelList */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclustergetprogramguideparams/data
func (m_ MTRChannelClusterGetProgramGuideParams) Data() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("data"))
	return rv
}/* debug [instance_properties/getter]: data */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclustergetprogramguideparams/data
func (m_ MTRChannelClusterGetProgramGuideParams) SetData(value foundation.Data) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setData:"), value)
}/* debug [instance_properties/setter]: data */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclustergetprogramguideparams/endtime
func (m_ MTRChannelClusterGetProgramGuideParams) EndTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("endTime"))
	return rv
}/* debug [instance_properties/getter]: endTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclustergetprogramguideparams/endtime
func (m_ MTRChannelClusterGetProgramGuideParams) SetEndTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEndTime:"), value)
}/* debug [instance_properties/setter]: endTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclustergetprogramguideparams/pagetoken
func (m_ MTRChannelClusterGetProgramGuideParams) PageToken() IMTRChannelClusterPageTokenStruct {
	rv := objc.Send[MTRChannelClusterPageTokenStruct](m_.ID, objc.Sel("pageToken"))
	return rv
}/* debug [instance_properties/getter]: pageToken */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclustergetprogramguideparams/pagetoken
func (m_ MTRChannelClusterGetProgramGuideParams) SetPageToken(value IMTRChannelClusterPageTokenStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPageToken:"), value)
}/* debug [instance_properties/setter]: pageToken */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclustergetprogramguideparams/recordingflag
func (m_ MTRChannelClusterGetProgramGuideParams) RecordingFlag() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("recordingFlag"))
	return rv
}/* debug [instance_properties/getter]: recordingFlag */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclustergetprogramguideparams/recordingflag
func (m_ MTRChannelClusterGetProgramGuideParams) SetRecordingFlag(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRecordingFlag:"), value)
}/* debug [instance_properties/setter]: recordingFlag */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclustergetprogramguideparams/serversideprocessingtimeout
func (m_ MTRChannelClusterGetProgramGuideParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclustergetprogramguideparams/serversideprocessingtimeout
func (m_ MTRChannelClusterGetProgramGuideParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclustergetprogramguideparams/starttime
func (m_ MTRChannelClusterGetProgramGuideParams) StartTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("startTime"))
	return rv
}/* debug [instance_properties/getter]: startTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclustergetprogramguideparams/starttime
func (m_ MTRChannelClusterGetProgramGuideParams) SetStartTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStartTime:"), value)
}/* debug [instance_properties/setter]: startTime */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclustergetprogramguideparams/timedinvoketimeoutms
func (m_ MTRChannelClusterGetProgramGuideParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclustergetprogramguideparams/timedinvoketimeoutms
func (m_ MTRChannelClusterGetProgramGuideParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRChannelClusterGetProgramGuideParams */



