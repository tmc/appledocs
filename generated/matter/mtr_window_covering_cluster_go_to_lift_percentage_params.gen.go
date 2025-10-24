// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRWindowCoveringClusterGoToLiftPercentageParams] class.
var (
	MTRWindowCoveringClusterGoToLiftPercentageParamsClass     _MTRWindowCoveringClusterGoToLiftPercentageParamsClass
	MTRWindowCoveringClusterGoToLiftPercentageParamsClassOnce sync.Once
)

func getMTRWindowCoveringClusterGoToLiftPercentageParamsClass() _MTRWindowCoveringClusterGoToLiftPercentageParamsClass {
	MTRWindowCoveringClusterGoToLiftPercentageParamsClassOnce.Do(func() {
		MTRWindowCoveringClusterGoToLiftPercentageParamsClass = _MTRWindowCoveringClusterGoToLiftPercentageParamsClass{objc.GetClass("MTRWindowCoveringClusterGoToLiftPercentageParams")}
	})
	return MTRWindowCoveringClusterGoToLiftPercentageParamsClass
}

type _MTRWindowCoveringClusterGoToLiftPercentageParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRWindowCoveringClusterGoToLiftPercentageParams] class.
type IMTRWindowCoveringClusterGoToLiftPercentageParams interface {
	objectivec.IObject
	// properties:
	LiftPercent100thsValue() objc.IObject /* cross-framework: NSNumber */
	SetLiftPercent100thsValue(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWindowCoveringClusterGoToLiftPercentageParams
type MTRWindowCoveringClusterGoToLiftPercentageParams struct {
	objectivec.Object
}

// MTRWindowCoveringClusterGoToLiftPercentageParamsFrom constructs a [MTRWindowCoveringClusterGoToLiftPercentageParams] from an unsafe.Pointer.
func MTRWindowCoveringClusterGoToLiftPercentageParamsFrom(ptr unsafe.Pointer) MTRWindowCoveringClusterGoToLiftPercentageParams {
	return MTRWindowCoveringClusterGoToLiftPercentageParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRWindowCoveringClusterGoToLiftPercentageParamsClass) Alloc() MTRWindowCoveringClusterGoToLiftPercentageParams {
	rv := objc.Send[MTRWindowCoveringClusterGoToLiftPercentageParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRWindowCoveringClusterGoToLiftPercentageParamsClass) New() MTRWindowCoveringClusterGoToLiftPercentageParams {
	rv := objc.Send[MTRWindowCoveringClusterGoToLiftPercentageParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRWindowCoveringClusterGoToLiftPercentageParams) Init() MTRWindowCoveringClusterGoToLiftPercentageParams {
	rv := objc.Send[MTRWindowCoveringClusterGoToLiftPercentageParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRWindowCoveringClusterGoToLiftPercentageParams) Autorelease() MTRWindowCoveringClusterGoToLiftPercentageParams {
	rv := objc.Send[MTRWindowCoveringClusterGoToLiftPercentageParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRWindowCoveringClusterGoToLiftPercentageParams creates a new MTRWindowCoveringClusterGoToLiftPercentageParams instance.
func NewMTRWindowCoveringClusterGoToLiftPercentageParams() MTRWindowCoveringClusterGoToLiftPercentageParams {
	return getMTRWindowCoveringClusterGoToLiftPercentageParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrwindowcoveringclustergotoliftpercentageparams/liftpercent100thsvalue
func (m_ MTRWindowCoveringClusterGoToLiftPercentageParams) LiftPercent100thsValue() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("liftPercent100thsValue"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrwindowcoveringclustergotoliftpercentageparams/liftpercent100thsvalue
func (m_ MTRWindowCoveringClusterGoToLiftPercentageParams) SetLiftPercent100thsValue(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLiftPercent100thsValue:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrwindowcoveringclustergotoliftpercentageparams/serversideprocessingtimeout
func (m_ MTRWindowCoveringClusterGoToLiftPercentageParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrwindowcoveringclustergotoliftpercentageparams/serversideprocessingtimeout
func (m_ MTRWindowCoveringClusterGoToLiftPercentageParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrwindowcoveringclustergotoliftpercentageparams/timedinvoketimeoutms
func (m_ MTRWindowCoveringClusterGoToLiftPercentageParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrwindowcoveringclustergotoliftpercentageparams/timedinvoketimeoutms
func (m_ MTRWindowCoveringClusterGoToLiftPercentageParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



