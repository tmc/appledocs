// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRColorControlClusterMoveSaturationParams] class.
var (
	MTRColorControlClusterMoveSaturationParamsClass     _MTRColorControlClusterMoveSaturationParamsClass
	MTRColorControlClusterMoveSaturationParamsClassOnce sync.Once
)

func getMTRColorControlClusterMoveSaturationParamsClass() _MTRColorControlClusterMoveSaturationParamsClass {
	MTRColorControlClusterMoveSaturationParamsClassOnce.Do(func() {
		MTRColorControlClusterMoveSaturationParamsClass = _MTRColorControlClusterMoveSaturationParamsClass{objc.GetClass("MTRColorControlClusterMoveSaturationParams")}
	})
	return MTRColorControlClusterMoveSaturationParamsClass
}

type _MTRColorControlClusterMoveSaturationParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRColorControlClusterMoveSaturationParams] class.
type IMTRColorControlClusterMoveSaturationParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveSaturationParams
type MTRColorControlClusterMoveSaturationParams struct {
	objectivec.Object
}

// MTRColorControlClusterMoveSaturationParamsFrom constructs a [MTRColorControlClusterMoveSaturationParams] from an unsafe.Pointer.
func MTRColorControlClusterMoveSaturationParamsFrom(ptr unsafe.Pointer) MTRColorControlClusterMoveSaturationParams {
	return MTRColorControlClusterMoveSaturationParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRColorControlClusterMoveSaturationParamsClass) Alloc() MTRColorControlClusterMoveSaturationParams {
	rv := objc.Send[MTRColorControlClusterMoveSaturationParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRColorControlClusterMoveSaturationParamsClass) New() MTRColorControlClusterMoveSaturationParams {
	rv := objc.Send[MTRColorControlClusterMoveSaturationParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRColorControlClusterMoveSaturationParams) Init() MTRColorControlClusterMoveSaturationParams {
	rv := objc.Send[MTRColorControlClusterMoveSaturationParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRColorControlClusterMoveSaturationParams) Autorelease() MTRColorControlClusterMoveSaturationParams {
	rv := objc.Send[MTRColorControlClusterMoveSaturationParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRColorControlClusterMoveSaturationParams creates a new MTRColorControlClusterMoveSaturationParams instance.
func NewMTRColorControlClusterMoveSaturationParams() MTRColorControlClusterMoveSaturationParams {
	return getMTRColorControlClusterMoveSaturationParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovesaturationparams/rate
func (m_ MTRColorControlClusterMoveSaturationParams) Rate() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("rate"))
	return rv
}


// SetRate sets the value of the rate property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovesaturationparams/rate
func (m_ MTRColorControlClusterMoveSaturationParams) SetRate(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRate:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovesaturationparams/movemode
func (m_ MTRColorControlClusterMoveSaturationParams) MoveMode() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("moveMode"))
	return rv
}


// SetMoveMode sets the value of the moveMode property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovesaturationparams/movemode
func (m_ MTRColorControlClusterMoveSaturationParams) SetMoveMode(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMoveMode:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovesaturationparams/timedinvoketimeoutms
func (m_ MTRColorControlClusterMoveSaturationParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovesaturationparams/timedinvoketimeoutms
func (m_ MTRColorControlClusterMoveSaturationParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovesaturationparams/optionsoverride
func (m_ MTRColorControlClusterMoveSaturationParams) OptionsOverride() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("optionsOverride"))
	return rv
}


// SetOptionsOverride sets the value of the optionsOverride property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovesaturationparams/optionsoverride
func (m_ MTRColorControlClusterMoveSaturationParams) SetOptionsOverride(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsOverride:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovesaturationparams/optionsmask
func (m_ MTRColorControlClusterMoveSaturationParams) OptionsMask() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("optionsMask"))
	return rv
}


// SetOptionsMask sets the value of the optionsMask property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovesaturationparams/optionsmask
func (m_ MTRColorControlClusterMoveSaturationParams) SetOptionsMask(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsMask:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovesaturationparams/serversideprocessingtimeout
func (m_ MTRColorControlClusterMoveSaturationParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovesaturationparams/serversideprocessingtimeout
func (m_ MTRColorControlClusterMoveSaturationParams) SetServerSideProcessingTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}



