// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRLevelControlClusterStepWithOnOffParams] class.
var (
	MTRLevelControlClusterStepWithOnOffParamsClass     _MTRLevelControlClusterStepWithOnOffParamsClass
	MTRLevelControlClusterStepWithOnOffParamsClassOnce sync.Once
)

func getMTRLevelControlClusterStepWithOnOffParamsClass() _MTRLevelControlClusterStepWithOnOffParamsClass {
	MTRLevelControlClusterStepWithOnOffParamsClassOnce.Do(func() {
		MTRLevelControlClusterStepWithOnOffParamsClass = _MTRLevelControlClusterStepWithOnOffParamsClass{objc.GetClass("MTRLevelControlClusterStepWithOnOffParams")}
	})
	return MTRLevelControlClusterStepWithOnOffParamsClass
}

type _MTRLevelControlClusterStepWithOnOffParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRLevelControlClusterStepWithOnOffParams] class.
type IMTRLevelControlClusterStepWithOnOffParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterStepWithOnOffParams
type MTRLevelControlClusterStepWithOnOffParams struct {
	objectivec.Object
}

// MTRLevelControlClusterStepWithOnOffParamsFrom constructs a [MTRLevelControlClusterStepWithOnOffParams] from an unsafe.Pointer.
func MTRLevelControlClusterStepWithOnOffParamsFrom(ptr unsafe.Pointer) MTRLevelControlClusterStepWithOnOffParams {
	return MTRLevelControlClusterStepWithOnOffParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRLevelControlClusterStepWithOnOffParamsClass) Alloc() MTRLevelControlClusterStepWithOnOffParams {
	rv := objc.Send[MTRLevelControlClusterStepWithOnOffParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRLevelControlClusterStepWithOnOffParamsClass) New() MTRLevelControlClusterStepWithOnOffParams {
	rv := objc.Send[MTRLevelControlClusterStepWithOnOffParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRLevelControlClusterStepWithOnOffParams) Init() MTRLevelControlClusterStepWithOnOffParams {
	rv := objc.Send[MTRLevelControlClusterStepWithOnOffParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRLevelControlClusterStepWithOnOffParams) Autorelease() MTRLevelControlClusterStepWithOnOffParams {
	rv := objc.Send[MTRLevelControlClusterStepWithOnOffParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRLevelControlClusterStepWithOnOffParams creates a new MTRLevelControlClusterStepWithOnOffParams instance.
func NewMTRLevelControlClusterStepWithOnOffParams() MTRLevelControlClusterStepWithOnOffParams {
	return getMTRLevelControlClusterStepWithOnOffParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclusterstepwithonoffparams/optionsmask
func (m_ MTRLevelControlClusterStepWithOnOffParams) OptionsMask() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("optionsMask"))
	return rv
}


// SetOptionsMask sets the value of the optionsMask property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclusterstepwithonoffparams/optionsmask
func (m_ MTRLevelControlClusterStepWithOnOffParams) SetOptionsMask(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsMask:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclusterstepwithonoffparams/optionsoverride
func (m_ MTRLevelControlClusterStepWithOnOffParams) OptionsOverride() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("optionsOverride"))
	return rv
}


// SetOptionsOverride sets the value of the optionsOverride property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclusterstepwithonoffparams/optionsoverride
func (m_ MTRLevelControlClusterStepWithOnOffParams) SetOptionsOverride(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsOverride:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclusterstepwithonoffparams/serversideprocessingtimeout
func (m_ MTRLevelControlClusterStepWithOnOffParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclusterstepwithonoffparams/serversideprocessingtimeout
func (m_ MTRLevelControlClusterStepWithOnOffParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclusterstepwithonoffparams/stepmode
func (m_ MTRLevelControlClusterStepWithOnOffParams) StepMode() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("stepMode"))
	return rv
}


// SetStepMode sets the value of the stepMode property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclusterstepwithonoffparams/stepmode
func (m_ MTRLevelControlClusterStepWithOnOffParams) SetStepMode(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStepMode:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclusterstepwithonoffparams/stepsize
func (m_ MTRLevelControlClusterStepWithOnOffParams) StepSize() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("stepSize"))
	return rv
}


// SetStepSize sets the value of the stepSize property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclusterstepwithonoffparams/stepsize
func (m_ MTRLevelControlClusterStepWithOnOffParams) SetStepSize(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStepSize:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclusterstepwithonoffparams/timedinvoketimeoutms
func (m_ MTRLevelControlClusterStepWithOnOffParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclusterstepwithonoffparams/timedinvoketimeoutms
func (m_ MTRLevelControlClusterStepWithOnOffParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclusterstepwithonoffparams/transitiontime
func (m_ MTRLevelControlClusterStepWithOnOffParams) TransitionTime() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("transitionTime"))
	return rv
}


// SetTransitionTime sets the value of the transitionTime property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclusterstepwithonoffparams/transitiontime
func (m_ MTRLevelControlClusterStepWithOnOffParams) SetTransitionTime(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTransitionTime:"), value)
}



