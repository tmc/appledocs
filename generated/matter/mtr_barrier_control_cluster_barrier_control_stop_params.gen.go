// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRBarrierControlClusterBarrierControlStopParams] class.
var (
	MTRBarrierControlClusterBarrierControlStopParamsClass     _MTRBarrierControlClusterBarrierControlStopParamsClass
	MTRBarrierControlClusterBarrierControlStopParamsClassOnce sync.Once
)

func getMTRBarrierControlClusterBarrierControlStopParamsClass() _MTRBarrierControlClusterBarrierControlStopParamsClass {
	MTRBarrierControlClusterBarrierControlStopParamsClassOnce.Do(func() {
		MTRBarrierControlClusterBarrierControlStopParamsClass = _MTRBarrierControlClusterBarrierControlStopParamsClass{objc.GetClass("MTRBarrierControlClusterBarrierControlStopParams")}
	})
	return MTRBarrierControlClusterBarrierControlStopParamsClass
}

type _MTRBarrierControlClusterBarrierControlStopParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRBarrierControlClusterBarrierControlStopParams] class.
type IMTRBarrierControlClusterBarrierControlStopParams interface {
	objectivec.IObject
	ServerSideProcessingTimeout() foundation.Number
	SetServerSideProcessingTimeout(value foundation.INumber)
	TimedInvokeTimeoutMs() foundation.Number
	SetTimedInvokeTimeoutMs(value foundation.INumber)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBarrierControlClusterBarrierControlStopParams
type MTRBarrierControlClusterBarrierControlStopParams struct {
	objectivec.Object
}

// MTRBarrierControlClusterBarrierControlStopParamsFrom constructs a [MTRBarrierControlClusterBarrierControlStopParams] from an unsafe.Pointer.
func MTRBarrierControlClusterBarrierControlStopParamsFrom(ptr unsafe.Pointer) MTRBarrierControlClusterBarrierControlStopParams {
	return MTRBarrierControlClusterBarrierControlStopParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBarrierControlClusterBarrierControlStopParamsClass) Alloc() MTRBarrierControlClusterBarrierControlStopParams {
	rv := objc.Send[MTRBarrierControlClusterBarrierControlStopParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBarrierControlClusterBarrierControlStopParamsClass) New() MTRBarrierControlClusterBarrierControlStopParams {
	rv := objc.Send[MTRBarrierControlClusterBarrierControlStopParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBarrierControlClusterBarrierControlStopParams) Init() MTRBarrierControlClusterBarrierControlStopParams {
	rv := objc.Send[MTRBarrierControlClusterBarrierControlStopParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBarrierControlClusterBarrierControlStopParams) Autorelease() MTRBarrierControlClusterBarrierControlStopParams {
	rv := objc.Send[MTRBarrierControlClusterBarrierControlStopParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBarrierControlClusterBarrierControlStopParams creates a new MTRBarrierControlClusterBarrierControlStopParams instance.
func NewMTRBarrierControlClusterBarrierControlStopParams() MTRBarrierControlClusterBarrierControlStopParams {
	return getMTRBarrierControlClusterBarrierControlStopParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbarriercontrolclusterbarriercontrolstopparams/serversideprocessingtimeout
func (m_ MTRBarrierControlClusterBarrierControlStopParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbarriercontrolclusterbarriercontrolstopparams/serversideprocessingtimeout
func (m_ MTRBarrierControlClusterBarrierControlStopParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbarriercontrolclusterbarriercontrolstopparams/timedinvoketimeoutms
func (m_ MTRBarrierControlClusterBarrierControlStopParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbarriercontrolclusterbarriercontrolstopparams/timedinvoketimeoutms
func (m_ MTRBarrierControlClusterBarrierControlStopParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



