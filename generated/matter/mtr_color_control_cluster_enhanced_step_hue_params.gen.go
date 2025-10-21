// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRColorControlClusterEnhancedStepHueParams] class.
var (
	MTRColorControlClusterEnhancedStepHueParamsClass     _MTRColorControlClusterEnhancedStepHueParamsClass
	MTRColorControlClusterEnhancedStepHueParamsClassOnce sync.Once
)

func getMTRColorControlClusterEnhancedStepHueParamsClass() _MTRColorControlClusterEnhancedStepHueParamsClass {
	MTRColorControlClusterEnhancedStepHueParamsClassOnce.Do(func() {
		MTRColorControlClusterEnhancedStepHueParamsClass = _MTRColorControlClusterEnhancedStepHueParamsClass{objc.GetClass("MTRColorControlClusterEnhancedStepHueParams")}
	})
	return MTRColorControlClusterEnhancedStepHueParamsClass
}

type _MTRColorControlClusterEnhancedStepHueParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRColorControlClusterEnhancedStepHueParams] class.
type IMTRColorControlClusterEnhancedStepHueParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterEnhancedStepHueParams
type MTRColorControlClusterEnhancedStepHueParams struct {
	objectivec.Object
}

// MTRColorControlClusterEnhancedStepHueParamsFrom constructs a [MTRColorControlClusterEnhancedStepHueParams] from an unsafe.Pointer.
func MTRColorControlClusterEnhancedStepHueParamsFrom(ptr unsafe.Pointer) MTRColorControlClusterEnhancedStepHueParams {
	return MTRColorControlClusterEnhancedStepHueParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRColorControlClusterEnhancedStepHueParamsClass) Alloc() MTRColorControlClusterEnhancedStepHueParams {
	rv := objc.Send[MTRColorControlClusterEnhancedStepHueParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRColorControlClusterEnhancedStepHueParamsClass) New() MTRColorControlClusterEnhancedStepHueParams {
	rv := objc.Send[MTRColorControlClusterEnhancedStepHueParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRColorControlClusterEnhancedStepHueParams) Init() MTRColorControlClusterEnhancedStepHueParams {
	rv := objc.Send[MTRColorControlClusterEnhancedStepHueParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRColorControlClusterEnhancedStepHueParams) Autorelease() MTRColorControlClusterEnhancedStepHueParams {
	rv := objc.Send[MTRColorControlClusterEnhancedStepHueParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRColorControlClusterEnhancedStepHueParams creates a new MTRColorControlClusterEnhancedStepHueParams instance.
func NewMTRColorControlClusterEnhancedStepHueParams() MTRColorControlClusterEnhancedStepHueParams {
	return getMTRColorControlClusterEnhancedStepHueParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedstephueparams/stepmode
func (m_ MTRColorControlClusterEnhancedStepHueParams) StepMode() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("stepMode"))
	return rv
}


// SetStepMode sets the value of the stepMode property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedstephueparams/stepmode
func (m_ MTRColorControlClusterEnhancedStepHueParams) SetStepMode(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStepMode:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedstephueparams/optionsmask
func (m_ MTRColorControlClusterEnhancedStepHueParams) OptionsMask() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("optionsMask"))
	return rv
}


// SetOptionsMask sets the value of the optionsMask property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedstephueparams/optionsmask
func (m_ MTRColorControlClusterEnhancedStepHueParams) SetOptionsMask(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsMask:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedstephueparams/transitiontime
func (m_ MTRColorControlClusterEnhancedStepHueParams) TransitionTime() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("transitionTime"))
	return rv
}


// SetTransitionTime sets the value of the transitionTime property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedstephueparams/transitiontime
func (m_ MTRColorControlClusterEnhancedStepHueParams) SetTransitionTime(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTransitionTime:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedstephueparams/serversideprocessingtimeout
func (m_ MTRColorControlClusterEnhancedStepHueParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedstephueparams/serversideprocessingtimeout
func (m_ MTRColorControlClusterEnhancedStepHueParams) SetServerSideProcessingTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedstephueparams/optionsoverride
func (m_ MTRColorControlClusterEnhancedStepHueParams) OptionsOverride() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("optionsOverride"))
	return rv
}


// SetOptionsOverride sets the value of the optionsOverride property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedstephueparams/optionsoverride
func (m_ MTRColorControlClusterEnhancedStepHueParams) SetOptionsOverride(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsOverride:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedstephueparams/stepsize
func (m_ MTRColorControlClusterEnhancedStepHueParams) StepSize() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("stepSize"))
	return rv
}


// SetStepSize sets the value of the stepSize property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedstephueparams/stepsize
func (m_ MTRColorControlClusterEnhancedStepHueParams) SetStepSize(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStepSize:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedstephueparams/timedinvoketimeoutms
func (m_ MTRColorControlClusterEnhancedStepHueParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedstephueparams/timedinvoketimeoutms
func (m_ MTRColorControlClusterEnhancedStepHueParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



