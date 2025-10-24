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
	// properties:
	Data() objc.IObject /* cross-framework: NSString */
	SetData(value objc.IObject /* cross-framework: NSString */)
	EncodingHint() objc.IObject /* cross-framework: NSString */
	SetEncodingHint(value objc.IObject /* cross-framework: NSString */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentAppObserverClusterContentAppMessageParams/data
func (m_ MTRContentAppObserverClusterContentAppMessageParams) Data() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("data"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentAppObserverClusterContentAppMessageParams/data
func (m_ MTRContentAppObserverClusterContentAppMessageParams) SetData(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setData:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentAppObserverClusterContentAppMessageParams/encodingHint
func (m_ MTRContentAppObserverClusterContentAppMessageParams) EncodingHint() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("encodingHint"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentAppObserverClusterContentAppMessageParams/encodingHint
func (m_ MTRContentAppObserverClusterContentAppMessageParams) SetEncodingHint(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEncodingHint:"), value)
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentAppObserverClusterContentAppMessageParams/serverSideProcessingTimeout
func (m_ MTRContentAppObserverClusterContentAppMessageParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentAppObserverClusterContentAppMessageParams/serverSideProcessingTimeout
func (m_ MTRContentAppObserverClusterContentAppMessageParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentAppObserverClusterContentAppMessageParams/timedInvokeTimeoutMs
func (m_ MTRContentAppObserverClusterContentAppMessageParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentAppObserverClusterContentAppMessageParams/timedInvokeTimeoutMs
func (m_ MTRContentAppObserverClusterContentAppMessageParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



