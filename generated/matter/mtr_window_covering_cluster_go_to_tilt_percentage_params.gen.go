// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRWindowCoveringClusterGoToTiltPercentageParams] class.
var (
	MTRWindowCoveringClusterGoToTiltPercentageParamsClass     _MTRWindowCoveringClusterGoToTiltPercentageParamsClass
	MTRWindowCoveringClusterGoToTiltPercentageParamsClassOnce sync.Once
)

func getMTRWindowCoveringClusterGoToTiltPercentageParamsClass() _MTRWindowCoveringClusterGoToTiltPercentageParamsClass {
	MTRWindowCoveringClusterGoToTiltPercentageParamsClassOnce.Do(func() {
		MTRWindowCoveringClusterGoToTiltPercentageParamsClass = _MTRWindowCoveringClusterGoToTiltPercentageParamsClass{objc.GetClass("MTRWindowCoveringClusterGoToTiltPercentageParams")}
	})
	return MTRWindowCoveringClusterGoToTiltPercentageParamsClass
}

type _MTRWindowCoveringClusterGoToTiltPercentageParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRWindowCoveringClusterGoToTiltPercentageParams] class.
type IMTRWindowCoveringClusterGoToTiltPercentageParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWindowCoveringClusterGoToTiltPercentageParams
type MTRWindowCoveringClusterGoToTiltPercentageParams struct {
	objectivec.Object
}

// MTRWindowCoveringClusterGoToTiltPercentageParamsFrom constructs a [MTRWindowCoveringClusterGoToTiltPercentageParams] from an unsafe.Pointer.
func MTRWindowCoveringClusterGoToTiltPercentageParamsFrom(ptr unsafe.Pointer) MTRWindowCoveringClusterGoToTiltPercentageParams {
	return MTRWindowCoveringClusterGoToTiltPercentageParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRWindowCoveringClusterGoToTiltPercentageParamsClass) Alloc() MTRWindowCoveringClusterGoToTiltPercentageParams {
	rv := objc.Send[MTRWindowCoveringClusterGoToTiltPercentageParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRWindowCoveringClusterGoToTiltPercentageParamsClass) New() MTRWindowCoveringClusterGoToTiltPercentageParams {
	rv := objc.Send[MTRWindowCoveringClusterGoToTiltPercentageParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRWindowCoveringClusterGoToTiltPercentageParams) Init() MTRWindowCoveringClusterGoToTiltPercentageParams {
	rv := objc.Send[MTRWindowCoveringClusterGoToTiltPercentageParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRWindowCoveringClusterGoToTiltPercentageParams) Autorelease() MTRWindowCoveringClusterGoToTiltPercentageParams {
	rv := objc.Send[MTRWindowCoveringClusterGoToTiltPercentageParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRWindowCoveringClusterGoToTiltPercentageParams creates a new MTRWindowCoveringClusterGoToTiltPercentageParams instance.
func NewMTRWindowCoveringClusterGoToTiltPercentageParams() MTRWindowCoveringClusterGoToTiltPercentageParams {
	return getMTRWindowCoveringClusterGoToTiltPercentageParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrwindowcoveringclustergototiltpercentageparams/serversideprocessingtimeout
func (m_ MTRWindowCoveringClusterGoToTiltPercentageParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrwindowcoveringclustergototiltpercentageparams/serversideprocessingtimeout
func (m_ MTRWindowCoveringClusterGoToTiltPercentageParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrwindowcoveringclustergototiltpercentageparams/tiltpercent100thsvalue
func (m_ MTRWindowCoveringClusterGoToTiltPercentageParams) TiltPercent100thsValue() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("tiltPercent100thsValue"))
	return rv
}


// SetTiltPercent100thsValue sets the value of the tiltPercent100thsValue property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrwindowcoveringclustergototiltpercentageparams/tiltpercent100thsvalue
func (m_ MTRWindowCoveringClusterGoToTiltPercentageParams) SetTiltPercent100thsValue(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTiltPercent100thsValue:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrwindowcoveringclustergototiltpercentageparams/timedinvoketimeoutms
func (m_ MTRWindowCoveringClusterGoToTiltPercentageParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrwindowcoveringclustergototiltpercentageparams/timedinvoketimeoutms
func (m_ MTRWindowCoveringClusterGoToTiltPercentageParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



