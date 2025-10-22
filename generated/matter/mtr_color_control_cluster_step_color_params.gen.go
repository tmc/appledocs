// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRColorControlClusterStepColorParams] class.
var (
	MTRColorControlClusterStepColorParamsClass     _MTRColorControlClusterStepColorParamsClass
	MTRColorControlClusterStepColorParamsClassOnce sync.Once
)

func getMTRColorControlClusterStepColorParamsClass() _MTRColorControlClusterStepColorParamsClass {
	MTRColorControlClusterStepColorParamsClassOnce.Do(func() {
		MTRColorControlClusterStepColorParamsClass = _MTRColorControlClusterStepColorParamsClass{objc.GetClass("MTRColorControlClusterStepColorParams")}
	})
	return MTRColorControlClusterStepColorParamsClass
}

type _MTRColorControlClusterStepColorParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRColorControlClusterStepColorParams] class.
type IMTRColorControlClusterStepColorParams interface {
	objectivec.IObject
	OptionsMask() foundation.Number
	SetOptionsMask(value foundation.INumber)
	OptionsOverride() foundation.Number
	SetOptionsOverride(value foundation.INumber)
	ServerSideProcessingTimeout() foundation.Number
	SetServerSideProcessingTimeout(value foundation.INumber)
	StepX() foundation.Number
	SetStepX(value foundation.INumber)
	StepY() foundation.Number
	SetStepY(value foundation.INumber)
	TimedInvokeTimeoutMs() foundation.Number
	SetTimedInvokeTimeoutMs(value foundation.INumber)
	TransitionTime() foundation.Number
	SetTransitionTime(value foundation.INumber)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterStepColorParams
type MTRColorControlClusterStepColorParams struct {
	objectivec.Object
}

// MTRColorControlClusterStepColorParamsFrom constructs a [MTRColorControlClusterStepColorParams] from an unsafe.Pointer.
func MTRColorControlClusterStepColorParamsFrom(ptr unsafe.Pointer) MTRColorControlClusterStepColorParams {
	return MTRColorControlClusterStepColorParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRColorControlClusterStepColorParamsClass) Alloc() MTRColorControlClusterStepColorParams {
	rv := objc.Send[MTRColorControlClusterStepColorParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRColorControlClusterStepColorParamsClass) New() MTRColorControlClusterStepColorParams {
	rv := objc.Send[MTRColorControlClusterStepColorParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRColorControlClusterStepColorParams) Init() MTRColorControlClusterStepColorParams {
	rv := objc.Send[MTRColorControlClusterStepColorParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRColorControlClusterStepColorParams) Autorelease() MTRColorControlClusterStepColorParams {
	rv := objc.Send[MTRColorControlClusterStepColorParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRColorControlClusterStepColorParams creates a new MTRColorControlClusterStepColorParams instance.
func NewMTRColorControlClusterStepColorParams() MTRColorControlClusterStepColorParams {
	return getMTRColorControlClusterStepColorParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepcolorparams/optionsmask
func (m_ MTRColorControlClusterStepColorParams) OptionsMask() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("optionsMask"))
	return rv
}


// SetOptionsMask sets the value of the optionsMask property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepcolorparams/optionsmask
func (m_ MTRColorControlClusterStepColorParams) SetOptionsMask(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsMask:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepcolorparams/optionsoverride
func (m_ MTRColorControlClusterStepColorParams) OptionsOverride() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("optionsOverride"))
	return rv
}


// SetOptionsOverride sets the value of the optionsOverride property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepcolorparams/optionsoverride
func (m_ MTRColorControlClusterStepColorParams) SetOptionsOverride(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsOverride:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepcolorparams/serversideprocessingtimeout
func (m_ MTRColorControlClusterStepColorParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepcolorparams/serversideprocessingtimeout
func (m_ MTRColorControlClusterStepColorParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepcolorparams/stepx
func (m_ MTRColorControlClusterStepColorParams) StepX() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("stepX"))
	return rv
}


// SetStepX sets the value of the stepX property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepcolorparams/stepx
func (m_ MTRColorControlClusterStepColorParams) SetStepX(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStepX:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepcolorparams/stepy
func (m_ MTRColorControlClusterStepColorParams) StepY() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("stepY"))
	return rv
}


// SetStepY sets the value of the stepY property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepcolorparams/stepy
func (m_ MTRColorControlClusterStepColorParams) SetStepY(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStepY:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepcolorparams/timedinvoketimeoutms
func (m_ MTRColorControlClusterStepColorParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepcolorparams/timedinvoketimeoutms
func (m_ MTRColorControlClusterStepColorParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepcolorparams/transitiontime
func (m_ MTRColorControlClusterStepColorParams) TransitionTime() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("transitionTime"))
	return rv
}


// SetTransitionTime sets the value of the transitionTime property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepcolorparams/transitiontime
func (m_ MTRColorControlClusterStepColorParams) SetTransitionTime(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTransitionTime:"), value)
}



