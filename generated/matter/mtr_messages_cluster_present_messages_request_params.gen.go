// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRMessagesClusterPresentMessagesRequestParams] class.
var (
	MTRMessagesClusterPresentMessagesRequestParamsClass     _MTRMessagesClusterPresentMessagesRequestParamsClass
	MTRMessagesClusterPresentMessagesRequestParamsClassOnce sync.Once
)

func getMTRMessagesClusterPresentMessagesRequestParamsClass() _MTRMessagesClusterPresentMessagesRequestParamsClass {
	MTRMessagesClusterPresentMessagesRequestParamsClassOnce.Do(func() {
		MTRMessagesClusterPresentMessagesRequestParamsClass = _MTRMessagesClusterPresentMessagesRequestParamsClass{objc.GetClass("MTRMessagesClusterPresentMessagesRequestParams")}
	})
	return MTRMessagesClusterPresentMessagesRequestParamsClass
}

type _MTRMessagesClusterPresentMessagesRequestParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRMessagesClusterPresentMessagesRequestParams] class.
type IMTRMessagesClusterPresentMessagesRequestParams interface {
	objectivec.IObject
	Duration() foundation.Number
	SetDuration(value foundation.INumber)
	MessageControl() foundation.Number
	SetMessageControl(value foundation.INumber)
	MessageID() foundation.NSData
	SetMessageID(value foundation.IData)
	MessageText() string
	SetMessageText(value string)
	Priority() foundation.Number
	SetPriority(value foundation.INumber)
	Responses() objc.ID
	SetResponses(value objc.ID)
	ServerSideProcessingTimeout() foundation.Number
	SetServerSideProcessingTimeout(value foundation.INumber)
	StartTime() foundation.Number
	SetStartTime(value foundation.INumber)
	TimedInvokeTimeoutMs() foundation.Number
	SetTimedInvokeTimeoutMs(value foundation.INumber)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterPresentMessagesRequestParams
type MTRMessagesClusterPresentMessagesRequestParams struct {
	objectivec.Object
}

// MTRMessagesClusterPresentMessagesRequestParamsFrom constructs a [MTRMessagesClusterPresentMessagesRequestParams] from an unsafe.Pointer.
func MTRMessagesClusterPresentMessagesRequestParamsFrom(ptr unsafe.Pointer) MTRMessagesClusterPresentMessagesRequestParams {
	return MTRMessagesClusterPresentMessagesRequestParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRMessagesClusterPresentMessagesRequestParamsClass) Alloc() MTRMessagesClusterPresentMessagesRequestParams {
	rv := objc.Send[MTRMessagesClusterPresentMessagesRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRMessagesClusterPresentMessagesRequestParamsClass) New() MTRMessagesClusterPresentMessagesRequestParams {
	rv := objc.Send[MTRMessagesClusterPresentMessagesRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRMessagesClusterPresentMessagesRequestParams) Init() MTRMessagesClusterPresentMessagesRequestParams {
	rv := objc.Send[MTRMessagesClusterPresentMessagesRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRMessagesClusterPresentMessagesRequestParams) Autorelease() MTRMessagesClusterPresentMessagesRequestParams {
	rv := objc.Send[MTRMessagesClusterPresentMessagesRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRMessagesClusterPresentMessagesRequestParams creates a new MTRMessagesClusterPresentMessagesRequestParams instance.
func NewMTRMessagesClusterPresentMessagesRequestParams() MTRMessagesClusterPresentMessagesRequestParams {
	return getMTRMessagesClusterPresentMessagesRequestParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterPresentMessagesRequestParams/duration
func (m_ MTRMessagesClusterPresentMessagesRequestParams) Duration() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("duration"))
	return rv
}


// SetDuration sets the value of the duration property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterPresentMessagesRequestParams/duration
func (m_ MTRMessagesClusterPresentMessagesRequestParams) SetDuration(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDuration:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterPresentMessagesRequestParams/messageControl
func (m_ MTRMessagesClusterPresentMessagesRequestParams) MessageControl() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("messageControl"))
	return rv
}


// SetMessageControl sets the value of the messageControl property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterPresentMessagesRequestParams/messageControl
func (m_ MTRMessagesClusterPresentMessagesRequestParams) SetMessageControl(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMessageControl:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterPresentMessagesRequestParams/messageID
func (m_ MTRMessagesClusterPresentMessagesRequestParams) MessageID() foundation.NSData {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("messageID"))
	return rv
}


// SetMessageID sets the value of the messageID property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterPresentMessagesRequestParams/messageID
func (m_ MTRMessagesClusterPresentMessagesRequestParams) SetMessageID(value foundation.IData) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMessageID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterPresentMessagesRequestParams/messageText
func (m_ MTRMessagesClusterPresentMessagesRequestParams) MessageText() string {
	rv := objc.Send[string](m_.ID, objc.Sel("messageText"))
	return rv
}


// SetMessageText sets the value of the messageText property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterPresentMessagesRequestParams/messageText
func (m_ MTRMessagesClusterPresentMessagesRequestParams) SetMessageText(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMessageText:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterPresentMessagesRequestParams/priority
func (m_ MTRMessagesClusterPresentMessagesRequestParams) Priority() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("priority"))
	return rv
}


// SetPriority sets the value of the priority property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterPresentMessagesRequestParams/priority
func (m_ MTRMessagesClusterPresentMessagesRequestParams) SetPriority(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPriority:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterPresentMessagesRequestParams/responses
func (m_ MTRMessagesClusterPresentMessagesRequestParams) Responses() objc.ID {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("responses"))
	return rv
}


// SetResponses sets the value of the responses property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterPresentMessagesRequestParams/responses
func (m_ MTRMessagesClusterPresentMessagesRequestParams) SetResponses(value objc.ID) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setResponses:"), value)
}

// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterPresentMessagesRequestParams/serverSideProcessingTimeout
func (m_ MTRMessagesClusterPresentMessagesRequestParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
// Controls how much time, in seconds, we will allow for the server to process the command.

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterPresentMessagesRequestParams/serverSideProcessingTimeout
func (m_ MTRMessagesClusterPresentMessagesRequestParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterPresentMessagesRequestParams/startTime
func (m_ MTRMessagesClusterPresentMessagesRequestParams) StartTime() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("startTime"))
	return rv
}


// SetStartTime sets the value of the startTime property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterPresentMessagesRequestParams/startTime
func (m_ MTRMessagesClusterPresentMessagesRequestParams) SetStartTime(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStartTime:"), value)
}

// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterPresentMessagesRequestParams/timedInvokeTimeoutMs
func (m_ MTRMessagesClusterPresentMessagesRequestParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
// Controls whether the command is a timed command (using Timed Invoke).

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterPresentMessagesRequestParams/timedInvokeTimeoutMs
func (m_ MTRMessagesClusterPresentMessagesRequestParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



