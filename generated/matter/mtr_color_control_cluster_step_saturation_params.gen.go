// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRColorControlClusterStepSaturationParams] class.
var (
	MTRColorControlClusterStepSaturationParamsClass     _MTRColorControlClusterStepSaturationParamsClass
	MTRColorControlClusterStepSaturationParamsClassOnce sync.Once
)

func getMTRColorControlClusterStepSaturationParamsClass() _MTRColorControlClusterStepSaturationParamsClass {
	MTRColorControlClusterStepSaturationParamsClassOnce.Do(func() {
		MTRColorControlClusterStepSaturationParamsClass = _MTRColorControlClusterStepSaturationParamsClass{objc.GetClass("MTRColorControlClusterStepSaturationParams")}
	})
	return MTRColorControlClusterStepSaturationParamsClass
}

type _MTRColorControlClusterStepSaturationParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRColorControlClusterStepSaturationParams] class.
type IMTRColorControlClusterStepSaturationParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterStepSaturationParams
type MTRColorControlClusterStepSaturationParams struct {
	objectivec.Object
}

// MTRColorControlClusterStepSaturationParamsFrom constructs a [MTRColorControlClusterStepSaturationParams] from an unsafe.Pointer.
func MTRColorControlClusterStepSaturationParamsFrom(ptr unsafe.Pointer) MTRColorControlClusterStepSaturationParams {
	return MTRColorControlClusterStepSaturationParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRColorControlClusterStepSaturationParamsClass) Alloc() MTRColorControlClusterStepSaturationParams {
	rv := objc.Send[MTRColorControlClusterStepSaturationParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRColorControlClusterStepSaturationParamsClass) New() MTRColorControlClusterStepSaturationParams {
	rv := objc.Send[MTRColorControlClusterStepSaturationParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRColorControlClusterStepSaturationParams) Init() MTRColorControlClusterStepSaturationParams {
	rv := objc.Send[MTRColorControlClusterStepSaturationParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRColorControlClusterStepSaturationParams) Autorelease() MTRColorControlClusterStepSaturationParams {
	rv := objc.Send[MTRColorControlClusterStepSaturationParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRColorControlClusterStepSaturationParams creates a new MTRColorControlClusterStepSaturationParams instance.
func NewMTRColorControlClusterStepSaturationParams() MTRColorControlClusterStepSaturationParams {
	return getMTRColorControlClusterStepSaturationParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepsaturationparams/serversideprocessingtimeout
func (m_ MTRColorControlClusterStepSaturationParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepsaturationparams/serversideprocessingtimeout
func (m_ MTRColorControlClusterStepSaturationParams) SetServerSideProcessingTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepsaturationparams/timedinvoketimeoutms
func (m_ MTRColorControlClusterStepSaturationParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepsaturationparams/timedinvoketimeoutms
func (m_ MTRColorControlClusterStepSaturationParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepsaturationparams/optionsoverride
func (m_ MTRColorControlClusterStepSaturationParams) OptionsOverride() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("optionsOverride"))
	return rv
}


// SetOptionsOverride sets the value of the optionsOverride property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepsaturationparams/optionsoverride
func (m_ MTRColorControlClusterStepSaturationParams) SetOptionsOverride(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsOverride:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepsaturationparams/stepsize
func (m_ MTRColorControlClusterStepSaturationParams) StepSize() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("stepSize"))
	return rv
}


// SetStepSize sets the value of the stepSize property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepsaturationparams/stepsize
func (m_ MTRColorControlClusterStepSaturationParams) SetStepSize(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStepSize:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepsaturationparams/transitiontime
func (m_ MTRColorControlClusterStepSaturationParams) TransitionTime() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("transitionTime"))
	return rv
}


// SetTransitionTime sets the value of the transitionTime property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepsaturationparams/transitiontime
func (m_ MTRColorControlClusterStepSaturationParams) SetTransitionTime(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTransitionTime:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepsaturationparams/stepmode
func (m_ MTRColorControlClusterStepSaturationParams) StepMode() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("stepMode"))
	return rv
}


// SetStepMode sets the value of the stepMode property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepsaturationparams/stepmode
func (m_ MTRColorControlClusterStepSaturationParams) SetStepMode(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStepMode:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepsaturationparams/optionsmask
func (m_ MTRColorControlClusterStepSaturationParams) OptionsMask() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("optionsMask"))
	return rv
}


// SetOptionsMask sets the value of the optionsMask property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepsaturationparams/optionsmask
func (m_ MTRColorControlClusterStepSaturationParams) SetOptionsMask(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsMask:"), value)
}



