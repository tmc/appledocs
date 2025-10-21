// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRColorControlClusterStepHueParams] class.
var (
	MTRColorControlClusterStepHueParamsClass     _MTRColorControlClusterStepHueParamsClass
	MTRColorControlClusterStepHueParamsClassOnce sync.Once
)

func getMTRColorControlClusterStepHueParamsClass() _MTRColorControlClusterStepHueParamsClass {
	MTRColorControlClusterStepHueParamsClassOnce.Do(func() {
		MTRColorControlClusterStepHueParamsClass = _MTRColorControlClusterStepHueParamsClass{objc.GetClass("MTRColorControlClusterStepHueParams")}
	})
	return MTRColorControlClusterStepHueParamsClass
}

type _MTRColorControlClusterStepHueParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRColorControlClusterStepHueParams] class.
type IMTRColorControlClusterStepHueParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterStepHueParams
type MTRColorControlClusterStepHueParams struct {
	objectivec.Object
}

// MTRColorControlClusterStepHueParamsFrom constructs a [MTRColorControlClusterStepHueParams] from an unsafe.Pointer.
func MTRColorControlClusterStepHueParamsFrom(ptr unsafe.Pointer) MTRColorControlClusterStepHueParams {
	return MTRColorControlClusterStepHueParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRColorControlClusterStepHueParamsClass) Alloc() MTRColorControlClusterStepHueParams {
	rv := objc.Send[MTRColorControlClusterStepHueParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRColorControlClusterStepHueParamsClass) New() MTRColorControlClusterStepHueParams {
	rv := objc.Send[MTRColorControlClusterStepHueParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRColorControlClusterStepHueParams) Init() MTRColorControlClusterStepHueParams {
	rv := objc.Send[MTRColorControlClusterStepHueParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRColorControlClusterStepHueParams) Autorelease() MTRColorControlClusterStepHueParams {
	rv := objc.Send[MTRColorControlClusterStepHueParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRColorControlClusterStepHueParams creates a new MTRColorControlClusterStepHueParams instance.
func NewMTRColorControlClusterStepHueParams() MTRColorControlClusterStepHueParams {
	return getMTRColorControlClusterStepHueParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstephueparams/serversideprocessingtimeout
func (m_ MTRColorControlClusterStepHueParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstephueparams/serversideprocessingtimeout
func (m_ MTRColorControlClusterStepHueParams) SetServerSideProcessingTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstephueparams/stepmode
func (m_ MTRColorControlClusterStepHueParams) StepMode() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("stepMode"))
	return rv
}


// SetStepMode sets the value of the stepMode property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstephueparams/stepmode
func (m_ MTRColorControlClusterStepHueParams) SetStepMode(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStepMode:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstephueparams/transitiontime
func (m_ MTRColorControlClusterStepHueParams) TransitionTime() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("transitionTime"))
	return rv
}


// SetTransitionTime sets the value of the transitionTime property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstephueparams/transitiontime
func (m_ MTRColorControlClusterStepHueParams) SetTransitionTime(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTransitionTime:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstephueparams/optionsmask
func (m_ MTRColorControlClusterStepHueParams) OptionsMask() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("optionsMask"))
	return rv
}


// SetOptionsMask sets the value of the optionsMask property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstephueparams/optionsmask
func (m_ MTRColorControlClusterStepHueParams) SetOptionsMask(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsMask:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstephueparams/optionsoverride
func (m_ MTRColorControlClusterStepHueParams) OptionsOverride() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("optionsOverride"))
	return rv
}


// SetOptionsOverride sets the value of the optionsOverride property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstephueparams/optionsoverride
func (m_ MTRColorControlClusterStepHueParams) SetOptionsOverride(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsOverride:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstephueparams/stepsize
func (m_ MTRColorControlClusterStepHueParams) StepSize() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("stepSize"))
	return rv
}


// SetStepSize sets the value of the stepSize property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstephueparams/stepsize
func (m_ MTRColorControlClusterStepHueParams) SetStepSize(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStepSize:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstephueparams/timedinvoketimeoutms
func (m_ MTRColorControlClusterStepHueParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstephueparams/timedinvoketimeoutms
func (m_ MTRColorControlClusterStepHueParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



