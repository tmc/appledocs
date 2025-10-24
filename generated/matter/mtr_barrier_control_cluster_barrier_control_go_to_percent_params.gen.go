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
	// properties:
	PercentOpen() objc.IObject /* cross-framework: NSNumber */
	SetPercentOpen(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbarriercontrolclusterbarriercontrolgotopercentparams/percentopen
func (m_ MTRBarrierControlClusterBarrierControlGoToPercentParams) PercentOpen() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("percentOpen"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbarriercontrolclusterbarriercontrolgotopercentparams/percentopen
func (m_ MTRBarrierControlClusterBarrierControlGoToPercentParams) SetPercentOpen(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPercentOpen:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbarriercontrolclusterbarriercontrolgotopercentparams/serversideprocessingtimeout
func (m_ MTRBarrierControlClusterBarrierControlGoToPercentParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbarriercontrolclusterbarriercontrolgotopercentparams/serversideprocessingtimeout
func (m_ MTRBarrierControlClusterBarrierControlGoToPercentParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbarriercontrolclusterbarriercontrolgotopercentparams/timedinvoketimeoutms
func (m_ MTRBarrierControlClusterBarrierControlGoToPercentParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbarriercontrolclusterbarriercontrolgotopercentparams/timedinvoketimeoutms
func (m_ MTRBarrierControlClusterBarrierControlGoToPercentParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



