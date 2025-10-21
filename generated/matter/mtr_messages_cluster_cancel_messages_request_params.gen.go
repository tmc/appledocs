// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRMessagesClusterCancelMessagesRequestParams] class.
var (
	MTRMessagesClusterCancelMessagesRequestParamsClass     _MTRMessagesClusterCancelMessagesRequestParamsClass
	MTRMessagesClusterCancelMessagesRequestParamsClassOnce sync.Once
)

func getMTRMessagesClusterCancelMessagesRequestParamsClass() _MTRMessagesClusterCancelMessagesRequestParamsClass {
	MTRMessagesClusterCancelMessagesRequestParamsClassOnce.Do(func() {
		MTRMessagesClusterCancelMessagesRequestParamsClass = _MTRMessagesClusterCancelMessagesRequestParamsClass{objc.GetClass("MTRMessagesClusterCancelMessagesRequestParams")}
	})
	return MTRMessagesClusterCancelMessagesRequestParamsClass
}

type _MTRMessagesClusterCancelMessagesRequestParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRMessagesClusterCancelMessagesRequestParams] class.
type IMTRMessagesClusterCancelMessagesRequestParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterCancelMessagesRequestParams
type MTRMessagesClusterCancelMessagesRequestParams struct {
	objectivec.Object
}

// MTRMessagesClusterCancelMessagesRequestParamsFrom constructs a [MTRMessagesClusterCancelMessagesRequestParams] from an unsafe.Pointer.
func MTRMessagesClusterCancelMessagesRequestParamsFrom(ptr unsafe.Pointer) MTRMessagesClusterCancelMessagesRequestParams {
	return MTRMessagesClusterCancelMessagesRequestParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRMessagesClusterCancelMessagesRequestParamsClass) Alloc() MTRMessagesClusterCancelMessagesRequestParams {
	rv := objc.Send[MTRMessagesClusterCancelMessagesRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRMessagesClusterCancelMessagesRequestParamsClass) New() MTRMessagesClusterCancelMessagesRequestParams {
	rv := objc.Send[MTRMessagesClusterCancelMessagesRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRMessagesClusterCancelMessagesRequestParams) Init() MTRMessagesClusterCancelMessagesRequestParams {
	rv := objc.Send[MTRMessagesClusterCancelMessagesRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRMessagesClusterCancelMessagesRequestParams) Autorelease() MTRMessagesClusterCancelMessagesRequestParams {
	rv := objc.Send[MTRMessagesClusterCancelMessagesRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRMessagesClusterCancelMessagesRequestParams creates a new MTRMessagesClusterCancelMessagesRequestParams instance.
func NewMTRMessagesClusterCancelMessagesRequestParams() MTRMessagesClusterCancelMessagesRequestParams {
	return getMTRMessagesClusterCancelMessagesRequestParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterCancelMessagesRequestParams/messageIDs
func (m_ MTRMessagesClusterCancelMessagesRequestParams) MessageIDs() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("messageIDs"))
	return rv
}


// SetMessageIDs sets the value of the messageIDs property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterCancelMessagesRequestParams/messageIDs
func (m_ MTRMessagesClusterCancelMessagesRequestParams) SetMessageIDs(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMessageIDs:"), value)
}

// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterCancelMessagesRequestParams/serverSideProcessingTimeout
func (m_ MTRMessagesClusterCancelMessagesRequestParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
// Controls how much time, in seconds, we will allow for the server to process the command.

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterCancelMessagesRequestParams/serverSideProcessingTimeout
func (m_ MTRMessagesClusterCancelMessagesRequestParams) SetServerSideProcessingTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterCancelMessagesRequestParams/timedInvokeTimeoutMs
func (m_ MTRMessagesClusterCancelMessagesRequestParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
// Controls whether the command is a timed command (using Timed Invoke).

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterCancelMessagesRequestParams/timedInvokeTimeoutMs
func (m_ MTRMessagesClusterCancelMessagesRequestParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



