// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRBarrierControlClusterBarrierControlGoToPercentParams] class.
var (
	MTRBarrierControlClusterBarrierControlGoToPercentParamsClass     _MTRBarrierControlClusterBarrierControlGoToPercentParamsClass
	MTRBarrierControlClusterBarrierControlGoToPercentParamsClassOnce sync.Once
)

func getMTRBarrierControlClusterBarrierControlGoToPercentParamsClass() _MTRBarrierControlClusterBarrierControlGoToPercentParamsClass {
	MTRBarrierControlClusterBarrierControlGoToPercentParamsClassOnce.Do(func() {
		MTRBarrierControlClusterBarrierControlGoToPercentParamsClass = _MTRBarrierControlClusterBarrierControlGoToPercentParamsClass{objc.GetClass("MTRBarrierControlClusterBarrierControlGoToPercentParams")}
	})
	return MTRBarrierControlClusterBarrierControlGoToPercentParamsClass
}

type _MTRBarrierControlClusterBarrierControlGoToPercentParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRBarrierControlClusterBarrierControlGoToPercentParams] class.
type IMTRBarrierControlClusterBarrierControlGoToPercentParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBarrierControlClusterBarrierControlGoToPercentParams
type MTRBarrierControlClusterBarrierControlGoToPercentParams struct {
	objectivec.Object
}

// MTRBarrierControlClusterBarrierControlGoToPercentParamsFrom constructs a [MTRBarrierControlClusterBarrierControlGoToPercentParams] from an unsafe.Pointer.
func MTRBarrierControlClusterBarrierControlGoToPercentParamsFrom(ptr unsafe.Pointer) MTRBarrierControlClusterBarrierControlGoToPercentParams {
	return MTRBarrierControlClusterBarrierControlGoToPercentParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBarrierControlClusterBarrierControlGoToPercentParamsClass) Alloc() MTRBarrierControlClusterBarrierControlGoToPercentParams {
	rv := objc.Send[MTRBarrierControlClusterBarrierControlGoToPercentParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBarrierControlClusterBarrierControlGoToPercentParamsClass) New() MTRBarrierControlClusterBarrierControlGoToPercentParams {
	rv := objc.Send[MTRBarrierControlClusterBarrierControlGoToPercentParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBarrierControlClusterBarrierControlGoToPercentParams) Init() MTRBarrierControlClusterBarrierControlGoToPercentParams {
	rv := objc.Send[MTRBarrierControlClusterBarrierControlGoToPercentParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBarrierControlClusterBarrierControlGoToPercentParams) Autorelease() MTRBarrierControlClusterBarrierControlGoToPercentParams {
	rv := objc.Send[MTRBarrierControlClusterBarrierControlGoToPercentParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBarrierControlClusterBarrierControlGoToPercentParams creates a new MTRBarrierControlClusterBarrierControlGoToPercentParams instance.
func NewMTRBarrierControlClusterBarrierControlGoToPercentParams() MTRBarrierControlClusterBarrierControlGoToPercentParams {
	return getMTRBarrierControlClusterBarrierControlGoToPercentParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbarriercontrolclusterbarriercontrolgotopercentparams/percentopen
func (m_ MTRBarrierControlClusterBarrierControlGoToPercentParams) PercentOpen() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("percentOpen"))
	return rv
}


// SetPercentOpen sets the value of the percentOpen property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbarriercontrolclusterbarriercontrolgotopercentparams/percentopen
func (m_ MTRBarrierControlClusterBarrierControlGoToPercentParams) SetPercentOpen(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPercentOpen:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbarriercontrolclusterbarriercontrolgotopercentparams/timedinvoketimeoutms
func (m_ MTRBarrierControlClusterBarrierControlGoToPercentParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbarriercontrolclusterbarriercontrolgotopercentparams/timedinvoketimeoutms
func (m_ MTRBarrierControlClusterBarrierControlGoToPercentParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbarriercontrolclusterbarriercontrolgotopercentparams/serversideprocessingtimeout
func (m_ MTRBarrierControlClusterBarrierControlGoToPercentParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbarriercontrolclusterbarriercontrolgotopercentparams/serversideprocessingtimeout
func (m_ MTRBarrierControlClusterBarrierControlGoToPercentParams) SetServerSideProcessingTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}



