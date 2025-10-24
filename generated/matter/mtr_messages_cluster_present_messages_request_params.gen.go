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
	// properties:
	Duration() objc.IObject /* cross-framework: NSNumber */
	SetDuration(value objc.IObject /* cross-framework: NSNumber */)
	MessageControl() objc.IObject /* cross-framework: NSNumber */
	SetMessageControl(value objc.IObject /* cross-framework: NSNumber */)
	MessageID() objc.IObject /* cross-framework: NSData */
	SetMessageID(value objc.IObject /* cross-framework: NSData */)
	MessageText() objc.IObject /* cross-framework: NSString */
	SetMessageText(value objc.IObject /* cross-framework: NSString */)
	Priority() objc.IObject /* cross-framework: NSNumber */
	SetPriority(value objc.IObject /* cross-framework: NSNumber */)
	Responses() objc.IObject /* cross-framework: NSArray */
	SetResponses(value objc.IObject /* cross-framework: NSArray */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	StartTime() objc.IObject /* cross-framework: NSNumber */
	SetStartTime(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterPresentMessagesRequestParams/duration
func (m_ MTRMessagesClusterPresentMessagesRequestParams) Duration() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("duration"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterPresentMessagesRequestParams/duration
func (m_ MTRMessagesClusterPresentMessagesRequestParams) SetDuration(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDuration:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterPresentMessagesRequestParams/messageControl
func (m_ MTRMessagesClusterPresentMessagesRequestParams) MessageControl() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("messageControl"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterPresentMessagesRequestParams/messageControl
func (m_ MTRMessagesClusterPresentMessagesRequestParams) SetMessageControl(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMessageControl:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterPresentMessagesRequestParams/messageID
func (m_ MTRMessagesClusterPresentMessagesRequestParams) MessageID() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("messageID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterPresentMessagesRequestParams/messageID
func (m_ MTRMessagesClusterPresentMessagesRequestParams) SetMessageID(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMessageID:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterPresentMessagesRequestParams/messageText
func (m_ MTRMessagesClusterPresentMessagesRequestParams) MessageText() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("messageText"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterPresentMessagesRequestParams/messageText
func (m_ MTRMessagesClusterPresentMessagesRequestParams) SetMessageText(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMessageText:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterPresentMessagesRequestParams/priority
func (m_ MTRMessagesClusterPresentMessagesRequestParams) Priority() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("priority"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterPresentMessagesRequestParams/priority
func (m_ MTRMessagesClusterPresentMessagesRequestParams) SetPriority(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPriority:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterPresentMessagesRequestParams/responses
func (m_ MTRMessagesClusterPresentMessagesRequestParams) Responses() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("responses"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterPresentMessagesRequestParams/responses
func (m_ MTRMessagesClusterPresentMessagesRequestParams) SetResponses(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setResponses:"), value)
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterPresentMessagesRequestParams/serverSideProcessingTimeout
func (m_ MTRMessagesClusterPresentMessagesRequestParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterPresentMessagesRequestParams/serverSideProcessingTimeout
func (m_ MTRMessagesClusterPresentMessagesRequestParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterPresentMessagesRequestParams/startTime
func (m_ MTRMessagesClusterPresentMessagesRequestParams) StartTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("startTime"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterPresentMessagesRequestParams/startTime
func (m_ MTRMessagesClusterPresentMessagesRequestParams) SetStartTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStartTime:"), value)
}


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterPresentMessagesRequestParams/timedInvokeTimeoutMs
func (m_ MTRMessagesClusterPresentMessagesRequestParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterPresentMessagesRequestParams/timedInvokeTimeoutMs
func (m_ MTRMessagesClusterPresentMessagesRequestParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



