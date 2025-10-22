// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRLevelControlClusterStepParams] class.
var (
	MTRLevelControlClusterStepParamsClass     _MTRLevelControlClusterStepParamsClass
	MTRLevelControlClusterStepParamsClassOnce sync.Once
)

func getMTRLevelControlClusterStepParamsClass() _MTRLevelControlClusterStepParamsClass {
	MTRLevelControlClusterStepParamsClassOnce.Do(func() {
		MTRLevelControlClusterStepParamsClass = _MTRLevelControlClusterStepParamsClass{objc.GetClass("MTRLevelControlClusterStepParams")}
	})
	return MTRLevelControlClusterStepParamsClass
}

type _MTRLevelControlClusterStepParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRLevelControlClusterStepParams] class.
type IMTRLevelControlClusterStepParams interface {
	objectivec.IObject
	OptionsMask() foundation.Number
	SetOptionsMask(value foundation.INumber)
	OptionsOverride() foundation.Number
	SetOptionsOverride(value foundation.INumber)
	ServerSideProcessingTimeout() foundation.Number
	SetServerSideProcessingTimeout(value foundation.INumber)
	StepMode() foundation.Number
	SetStepMode(value foundation.INumber)
	StepSize() foundation.Number
	SetStepSize(value foundation.INumber)
	TimedInvokeTimeoutMs() foundation.Number
	SetTimedInvokeTimeoutMs(value foundation.INumber)
	TransitionTime() foundation.Number
	SetTransitionTime(value foundation.INumber)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterStepParams
type MTRLevelControlClusterStepParams struct {
	objectivec.Object
}

// MTRLevelControlClusterStepParamsFrom constructs a [MTRLevelControlClusterStepParams] from an unsafe.Pointer.
func MTRLevelControlClusterStepParamsFrom(ptr unsafe.Pointer) MTRLevelControlClusterStepParams {
	return MTRLevelControlClusterStepParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRLevelControlClusterStepParamsClass) Alloc() MTRLevelControlClusterStepParams {
	rv := objc.Send[MTRLevelControlClusterStepParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRLevelControlClusterStepParamsClass) New() MTRLevelControlClusterStepParams {
	rv := objc.Send[MTRLevelControlClusterStepParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRLevelControlClusterStepParams) Init() MTRLevelControlClusterStepParams {
	rv := objc.Send[MTRLevelControlClusterStepParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRLevelControlClusterStepParams) Autorelease() MTRLevelControlClusterStepParams {
	rv := objc.Send[MTRLevelControlClusterStepParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRLevelControlClusterStepParams creates a new MTRLevelControlClusterStepParams instance.
func NewMTRLevelControlClusterStepParams() MTRLevelControlClusterStepParams {
	return getMTRLevelControlClusterStepParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclusterstepparams/optionsmask
func (m_ MTRLevelControlClusterStepParams) OptionsMask() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("optionsMask"))
	return rv
}


// SetOptionsMask sets the value of the optionsMask property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclusterstepparams/optionsmask
func (m_ MTRLevelControlClusterStepParams) SetOptionsMask(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsMask:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclusterstepparams/optionsoverride
func (m_ MTRLevelControlClusterStepParams) OptionsOverride() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("optionsOverride"))
	return rv
}


// SetOptionsOverride sets the value of the optionsOverride property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclusterstepparams/optionsoverride
func (m_ MTRLevelControlClusterStepParams) SetOptionsOverride(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsOverride:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclusterstepparams/serversideprocessingtimeout
func (m_ MTRLevelControlClusterStepParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclusterstepparams/serversideprocessingtimeout
func (m_ MTRLevelControlClusterStepParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclusterstepparams/stepmode
func (m_ MTRLevelControlClusterStepParams) StepMode() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("stepMode"))
	return rv
}


// SetStepMode sets the value of the stepMode property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclusterstepparams/stepmode
func (m_ MTRLevelControlClusterStepParams) SetStepMode(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStepMode:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclusterstepparams/stepsize
func (m_ MTRLevelControlClusterStepParams) StepSize() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("stepSize"))
	return rv
}


// SetStepSize sets the value of the stepSize property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclusterstepparams/stepsize
func (m_ MTRLevelControlClusterStepParams) SetStepSize(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStepSize:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclusterstepparams/timedinvoketimeoutms
func (m_ MTRLevelControlClusterStepParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclusterstepparams/timedinvoketimeoutms
func (m_ MTRLevelControlClusterStepParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclusterstepparams/transitiontime
func (m_ MTRLevelControlClusterStepParams) TransitionTime() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("transitionTime"))
	return rv
}


// SetTransitionTime sets the value of the transitionTime property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclusterstepparams/transitiontime
func (m_ MTRLevelControlClusterStepParams) SetTransitionTime(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTransitionTime:"), value)
}



