// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
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
}

//
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


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterSetDSTOffsetParams/dstOffset
func (m_ MTRTimeSynchronizationClusterSetDSTOffsetParams) DstOffset() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("dstOffset"))
	return rv
}


// SetDstOffset sets the value of the dstOffset property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterSetDSTOffsetParams/dstOffset
func (m_ MTRTimeSynchronizationClusterSetDSTOffsetParams) SetDstOffset(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDstOffset:"), value)
}
// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterSetDSTOffsetParams/serverSideProcessingTimeout
func (m_ MTRTimeSynchronizationClusterSetDSTOffsetParams) ServerSideProcessingTimeout() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
// Controls how much time, in seconds, we will allow for the server to process the command.

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterSetDSTOffsetParams/serverSideProcessingTimeout
func (m_ MTRTimeSynchronizationClusterSetDSTOffsetParams) SetServerSideProcessingTimeout(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}
// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterSetDSTOffsetParams/timedInvokeTimeoutMs
func (m_ MTRTimeSynchronizationClusterSetDSTOffsetParams) TimedInvokeTimeoutMs() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
// Controls whether the command is a timed command (using Timed Invoke).

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterSetDSTOffsetParams/timedInvokeTimeoutMs
func (m_ MTRTimeSynchronizationClusterSetDSTOffsetParams) SetTimedInvokeTimeoutMs(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}


