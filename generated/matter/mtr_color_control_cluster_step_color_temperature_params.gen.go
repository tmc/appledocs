// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRColorControlClusterStepColorTemperatureParams] class.
var (
	MTRColorControlClusterStepColorTemperatureParamsClass     _MTRColorControlClusterStepColorTemperatureParamsClass
	MTRColorControlClusterStepColorTemperatureParamsClassOnce sync.Once
)

func getMTRColorControlClusterStepColorTemperatureParamsClass() _MTRColorControlClusterStepColorTemperatureParamsClass {
	MTRColorControlClusterStepColorTemperatureParamsClassOnce.Do(func() {
		MTRColorControlClusterStepColorTemperatureParamsClass = _MTRColorControlClusterStepColorTemperatureParamsClass{objc.GetClass("MTRColorControlClusterStepColorTemperatureParams")}
	})
	return MTRColorControlClusterStepColorTemperatureParamsClass
}

type _MTRColorControlClusterStepColorTemperatureParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRColorControlClusterStepColorTemperatureParams] class.
type IMTRColorControlClusterStepColorTemperatureParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterStepColorTemperatureParams
type MTRColorControlClusterStepColorTemperatureParams struct {
	objectivec.Object
}

// MTRColorControlClusterStepColorTemperatureParamsFrom constructs a [MTRColorControlClusterStepColorTemperatureParams] from an unsafe.Pointer.
func MTRColorControlClusterStepColorTemperatureParamsFrom(ptr unsafe.Pointer) MTRColorControlClusterStepColorTemperatureParams {
	return MTRColorControlClusterStepColorTemperatureParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRColorControlClusterStepColorTemperatureParamsClass) Alloc() MTRColorControlClusterStepColorTemperatureParams {
	rv := objc.Send[MTRColorControlClusterStepColorTemperatureParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRColorControlClusterStepColorTemperatureParamsClass) New() MTRColorControlClusterStepColorTemperatureParams {
	rv := objc.Send[MTRColorControlClusterStepColorTemperatureParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRColorControlClusterStepColorTemperatureParams) Init() MTRColorControlClusterStepColorTemperatureParams {
	rv := objc.Send[MTRColorControlClusterStepColorTemperatureParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRColorControlClusterStepColorTemperatureParams) Autorelease() MTRColorControlClusterStepColorTemperatureParams {
	rv := objc.Send[MTRColorControlClusterStepColorTemperatureParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRColorControlClusterStepColorTemperatureParams creates a new MTRColorControlClusterStepColorTemperatureParams instance.
func NewMTRColorControlClusterStepColorTemperatureParams() MTRColorControlClusterStepColorTemperatureParams {
	return getMTRColorControlClusterStepColorTemperatureParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepcolortemperatureparams/colortemperaturemaximummireds
func (m_ MTRColorControlClusterStepColorTemperatureParams) ColorTemperatureMaximumMireds() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("colorTemperatureMaximumMireds"))
	return rv
}


// SetColorTemperatureMaximumMireds sets the value of the colorTemperatureMaximumMireds property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepcolortemperatureparams/colortemperaturemaximummireds
func (m_ MTRColorControlClusterStepColorTemperatureParams) SetColorTemperatureMaximumMireds(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setColorTemperatureMaximumMireds:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepcolortemperatureparams/colortemperatureminimummireds
func (m_ MTRColorControlClusterStepColorTemperatureParams) ColorTemperatureMinimumMireds() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("colorTemperatureMinimumMireds"))
	return rv
}


// SetColorTemperatureMinimumMireds sets the value of the colorTemperatureMinimumMireds property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepcolortemperatureparams/colortemperatureminimummireds
func (m_ MTRColorControlClusterStepColorTemperatureParams) SetColorTemperatureMinimumMireds(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setColorTemperatureMinimumMireds:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepcolortemperatureparams/optionsmask
func (m_ MTRColorControlClusterStepColorTemperatureParams) OptionsMask() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("optionsMask"))
	return rv
}


// SetOptionsMask sets the value of the optionsMask property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepcolortemperatureparams/optionsmask
func (m_ MTRColorControlClusterStepColorTemperatureParams) SetOptionsMask(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsMask:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepcolortemperatureparams/optionsoverride
func (m_ MTRColorControlClusterStepColorTemperatureParams) OptionsOverride() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("optionsOverride"))
	return rv
}


// SetOptionsOverride sets the value of the optionsOverride property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepcolortemperatureparams/optionsoverride
func (m_ MTRColorControlClusterStepColorTemperatureParams) SetOptionsOverride(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsOverride:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepcolortemperatureparams/serversideprocessingtimeout
func (m_ MTRColorControlClusterStepColorTemperatureParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepcolortemperatureparams/serversideprocessingtimeout
func (m_ MTRColorControlClusterStepColorTemperatureParams) SetServerSideProcessingTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepcolortemperatureparams/stepmode
func (m_ MTRColorControlClusterStepColorTemperatureParams) StepMode() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("stepMode"))
	return rv
}


// SetStepMode sets the value of the stepMode property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepcolortemperatureparams/stepmode
func (m_ MTRColorControlClusterStepColorTemperatureParams) SetStepMode(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStepMode:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepcolortemperatureparams/stepsize
func (m_ MTRColorControlClusterStepColorTemperatureParams) StepSize() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("stepSize"))
	return rv
}


// SetStepSize sets the value of the stepSize property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepcolortemperatureparams/stepsize
func (m_ MTRColorControlClusterStepColorTemperatureParams) SetStepSize(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStepSize:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepcolortemperatureparams/timedinvoketimeoutms
func (m_ MTRColorControlClusterStepColorTemperatureParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepcolortemperatureparams/timedinvoketimeoutms
func (m_ MTRColorControlClusterStepColorTemperatureParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepcolortemperatureparams/transitiontime
func (m_ MTRColorControlClusterStepColorTemperatureParams) TransitionTime() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("transitionTime"))
	return rv
}


// SetTransitionTime sets the value of the transitionTime property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepcolortemperatureparams/transitiontime
func (m_ MTRColorControlClusterStepColorTemperatureParams) SetTransitionTime(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTransitionTime:"), value)
}



