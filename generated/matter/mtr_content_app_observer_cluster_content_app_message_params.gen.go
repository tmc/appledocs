// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRContentAppObserverClusterContentAppMessageParams] class.
var (
	MTRContentAppObserverClusterContentAppMessageParamsClass     _MTRContentAppObserverClusterContentAppMessageParamsClass
	MTRContentAppObserverClusterContentAppMessageParamsClassOnce sync.Once
)

func getMTRContentAppObserverClusterContentAppMessageParamsClass() _MTRContentAppObserverClusterContentAppMessageParamsClass {
	MTRContentAppObserverClusterContentAppMessageParamsClassOnce.Do(func() {
		MTRContentAppObserverClusterContentAppMessageParamsClass = _MTRContentAppObserverClusterContentAppMessageParamsClass{objc.GetClass("MTRContentAppObserverClusterContentAppMessageParams")}
	})
	return MTRContentAppObserverClusterContentAppMessageParamsClass
}

type _MTRContentAppObserverClusterContentAppMessageParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRContentAppObserverClusterContentAppMessageParams] class.
type IMTRContentAppObserverClusterContentAppMessageParams interface {
	objectivec.IObject
	Data() string
	SetData(value string)
	EncodingHint() string
	SetEncodingHint(value string)
	ServerSideProcessingTimeout() foundation.Number
	SetServerSideProcessingTimeout(value foundation.INumber)
	TimedInvokeTimeoutMs() foundation.Number
	SetTimedInvokeTimeoutMs(value foundation.INumber)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentAppObserverClusterContentAppMessageParams
type MTRContentAppObserverClusterContentAppMessageParams struct {
	objectivec.Object
}

// MTRContentAppObserverClusterContentAppMessageParamsFrom constructs a [MTRContentAppObserverClusterContentAppMessageParams] from an unsafe.Pointer.
func MTRContentAppObserverClusterContentAppMessageParamsFrom(ptr unsafe.Pointer) MTRContentAppObserverClusterContentAppMessageParams {
	return MTRContentAppObserverClusterContentAppMessageParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRContentAppObserverClusterContentAppMessageParamsClass) Alloc() MTRContentAppObserverClusterContentAppMessageParams {
	rv := objc.Send[MTRContentAppObserverClusterContentAppMessageParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRContentAppObserverClusterContentAppMessageParamsClass) New() MTRContentAppObserverClusterContentAppMessageParams {
	rv := objc.Send[MTRContentAppObserverClusterContentAppMessageParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRContentAppObserverClusterContentAppMessageParams) Init() MTRContentAppObserverClusterContentAppMessageParams {
	rv := objc.Send[MTRContentAppObserverClusterContentAppMessageParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRContentAppObserverClusterContentAppMessageParams) Autorelease() MTRContentAppObserverClusterContentAppMessageParams {
	rv := objc.Send[MTRContentAppObserverClusterContentAppMessageParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRContentAppObserverClusterContentAppMessageParams creates a new MTRContentAppObserverClusterContentAppMessageParams instance.
func NewMTRContentAppObserverClusterContentAppMessageParams() MTRContentAppObserverClusterContentAppMessageParams {
	return getMTRContentAppObserverClusterContentAppMessageParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentAppObserverClusterContentAppMessageParams/data
func (m_ MTRContentAppObserverClusterContentAppMessageParams) Data() string {
	rv := objc.Send[string](m_.ID, objc.Sel("data"))
	return rv
}


// SetData sets the value of the data property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentAppObserverClusterContentAppMessageParams/data
func (m_ MTRContentAppObserverClusterContentAppMessageParams) SetData(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setData:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentAppObserverClusterContentAppMessageParams/encodingHint
func (m_ MTRContentAppObserverClusterContentAppMessageParams) EncodingHint() string {
	rv := objc.Send[string](m_.ID, objc.Sel("encodingHint"))
	return rv
}


// SetEncodingHint sets the value of the encodingHint property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentAppObserverClusterContentAppMessageParams/encodingHint
func (m_ MTRContentAppObserverClusterContentAppMessageParams) SetEncodingHint(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEncodingHint:"), objc.String(value))
}

// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentAppObserverClusterContentAppMessageParams/serverSideProcessingTimeout
func (m_ MTRContentAppObserverClusterContentAppMessageParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
// Controls how much time, in seconds, we will allow for the server to process the command.

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentAppObserverClusterContentAppMessageParams/serverSideProcessingTimeout
func (m_ MTRContentAppObserverClusterContentAppMessageParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentAppObserverClusterContentAppMessageParams/timedInvokeTimeoutMs
func (m_ MTRContentAppObserverClusterContentAppMessageParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
// Controls whether the command is a timed command (using Timed Invoke).

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentAppObserverClusterContentAppMessageParams/timedInvokeTimeoutMs
func (m_ MTRContentAppObserverClusterContentAppMessageParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



