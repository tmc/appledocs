// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRTimeSynchronizationClusterSetDSTOffsetParams] class.
var (
	MTRTimeSynchronizationClusterSetDSTOffsetParamsClass     _MTRTimeSynchronizationClusterSetDSTOffsetParamsClass
	MTRTimeSynchronizationClusterSetDSTOffsetParamsClassOnce sync.Once
)

func getMTRTimeSynchronizationClusterSetDSTOffsetParamsClass() _MTRTimeSynchronizationClusterSetDSTOffsetParamsClass {
	MTRTimeSynchronizationClusterSetDSTOffsetParamsClassOnce.Do(func() {
		MTRTimeSynchronizationClusterSetDSTOffsetParamsClass = _MTRTimeSynchronizationClusterSetDSTOffsetParamsClass{objc.GetClass("MTRTimeSynchronizationClusterSetDSTOffsetParams")}
	})
	return MTRTimeSynchronizationClusterSetDSTOffsetParamsClass
}

type _MTRTimeSynchronizationClusterSetDSTOffsetParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRTimeSynchronizationClusterSetDSTOffsetParams] class.
type IMTRTimeSynchronizationClusterSetDSTOffsetParams interface {
	objectivec.IObject
	// properties:
	DstOffset() objc.IObject /* cross-framework: NSArray */
	SetDstOffset(value objc.IObject /* cross-framework: NSArray */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterSetDSTOffsetParams
type MTRTimeSynchronizationClusterSetDSTOffsetParams struct {
	objectivec.Object
}

// MTRTimeSynchronizationClusterSetDSTOffsetParamsFrom constructs a [MTRTimeSynchronizationClusterSetDSTOffsetParams] from an unsafe.Pointer.
func MTRTimeSynchronizationClusterSetDSTOffsetParamsFrom(ptr unsafe.Pointer) MTRTimeSynchronizationClusterSetDSTOffsetParams {
	return MTRTimeSynchronizationClusterSetDSTOffsetParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTimeSynchronizationClusterSetDSTOffsetParamsClass) Alloc() MTRTimeSynchronizationClusterSetDSTOffsetParams {
	rv := objc.Send[MTRTimeSynchronizationClusterSetDSTOffsetParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTimeSynchronizationClusterSetDSTOffsetParamsClass) New() MTRTimeSynchronizationClusterSetDSTOffsetParams {
	rv := objc.Send[MTRTimeSynchronizationClusterSetDSTOffsetParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTimeSynchronizationClusterSetDSTOffsetParams) Init() MTRTimeSynchronizationClusterSetDSTOffsetParams {
	rv := objc.Send[MTRTimeSynchronizationClusterSetDSTOffsetParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTimeSynchronizationClusterSetDSTOffsetParams) Autorelease() MTRTimeSynchronizationClusterSetDSTOffsetParams {
	rv := objc.Send[MTRTimeSynchronizationClusterSetDSTOffsetParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTimeSynchronizationClusterSetDSTOffsetParams creates a new MTRTimeSynchronizationClusterSetDSTOffsetParams instance.
func NewMTRTimeSynchronizationClusterSetDSTOffsetParams() MTRTimeSynchronizationClusterSetDSTOffsetParams {
	return getMTRTimeSynchronizationClusterSetDSTOffsetParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterSetDSTOffsetParams/dstOffset
func (m_ MTRTimeSynchronizationClusterSetDSTOffsetParams) DstOffset() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("dstOffset"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterSetDSTOffsetParams/dstOffset
func (m_ MTRTimeSynchronizationClusterSetDSTOffsetParams) SetDstOffset(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDstOffset:"), value)
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterSetDSTOffsetParams/serverSideProcessingTimeout
func (m_ MTRTimeSynchronizationClusterSetDSTOffsetParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterSetDSTOffsetParams/serverSideProcessingTimeout
func (m_ MTRTimeSynchronizationClusterSetDSTOffsetParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterSetDSTOffsetParams/timedInvokeTimeoutMs
func (m_ MTRTimeSynchronizationClusterSetDSTOffsetParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterSetDSTOffsetParams/timedInvokeTimeoutMs
func (m_ MTRTimeSynchronizationClusterSetDSTOffsetParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



