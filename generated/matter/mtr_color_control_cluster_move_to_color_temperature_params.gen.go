// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRColorControlClusterMoveToColorTemperatureParams] class.
var (
	MTRColorControlClusterMoveToColorTemperatureParamsClass     _MTRColorControlClusterMoveToColorTemperatureParamsClass
	MTRColorControlClusterMoveToColorTemperatureParamsClassOnce sync.Once
)

func getMTRColorControlClusterMoveToColorTemperatureParamsClass() _MTRColorControlClusterMoveToColorTemperatureParamsClass {
	MTRColorControlClusterMoveToColorTemperatureParamsClassOnce.Do(func() {
		MTRColorControlClusterMoveToColorTemperatureParamsClass = _MTRColorControlClusterMoveToColorTemperatureParamsClass{objc.GetClass("MTRColorControlClusterMoveToColorTemperatureParams")}
	})
	return MTRColorControlClusterMoveToColorTemperatureParamsClass
}

type _MTRColorControlClusterMoveToColorTemperatureParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRColorControlClusterMoveToColorTemperatureParams] class.
type IMTRColorControlClusterMoveToColorTemperatureParams interface {
	objectivec.IObject
	ColorTemperature() foundation.Number
	SetColorTemperature(value foundation.INumber)
	ColorTemperatureMireds() foundation.Number
	SetColorTemperatureMireds(value foundation.INumber)
	OptionsMask() foundation.Number
	SetOptionsMask(value foundation.INumber)
	OptionsOverride() foundation.Number
	SetOptionsOverride(value foundation.INumber)
	ServerSideProcessingTimeout() foundation.Number
	SetServerSideProcessingTimeout(value foundation.INumber)
	TimedInvokeTimeoutMs() foundation.Number
	SetTimedInvokeTimeoutMs(value foundation.INumber)
	TransitionTime() foundation.Number
	SetTransitionTime(value foundation.INumber)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToColorTemperatureParams
type MTRColorControlClusterMoveToColorTemperatureParams struct {
	objectivec.Object
}

// MTRColorControlClusterMoveToColorTemperatureParamsFrom constructs a [MTRColorControlClusterMoveToColorTemperatureParams] from an unsafe.Pointer.
func MTRColorControlClusterMoveToColorTemperatureParamsFrom(ptr unsafe.Pointer) MTRColorControlClusterMoveToColorTemperatureParams {
	return MTRColorControlClusterMoveToColorTemperatureParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRColorControlClusterMoveToColorTemperatureParamsClass) Alloc() MTRColorControlClusterMoveToColorTemperatureParams {
	rv := objc.Send[MTRColorControlClusterMoveToColorTemperatureParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRColorControlClusterMoveToColorTemperatureParamsClass) New() MTRColorControlClusterMoveToColorTemperatureParams {
	rv := objc.Send[MTRColorControlClusterMoveToColorTemperatureParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRColorControlClusterMoveToColorTemperatureParams) Init() MTRColorControlClusterMoveToColorTemperatureParams {
	rv := objc.Send[MTRColorControlClusterMoveToColorTemperatureParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRColorControlClusterMoveToColorTemperatureParams) Autorelease() MTRColorControlClusterMoveToColorTemperatureParams {
	rv := objc.Send[MTRColorControlClusterMoveToColorTemperatureParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRColorControlClusterMoveToColorTemperatureParams creates a new MTRColorControlClusterMoveToColorTemperatureParams instance.
func NewMTRColorControlClusterMoveToColorTemperatureParams() MTRColorControlClusterMoveToColorTemperatureParams {
	return getMTRColorControlClusterMoveToColorTemperatureParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetocolortemperatureparams/colortemperature
func (m_ MTRColorControlClusterMoveToColorTemperatureParams) ColorTemperature() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("colorTemperature"))
	return rv
}


// SetColorTemperature sets the value of the colorTemperature property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetocolortemperatureparams/colortemperature
func (m_ MTRColorControlClusterMoveToColorTemperatureParams) SetColorTemperature(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setColorTemperature:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetocolortemperatureparams/colortemperaturemireds
func (m_ MTRColorControlClusterMoveToColorTemperatureParams) ColorTemperatureMireds() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("colorTemperatureMireds"))
	return rv
}


// SetColorTemperatureMireds sets the value of the colorTemperatureMireds property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetocolortemperatureparams/colortemperaturemireds
func (m_ MTRColorControlClusterMoveToColorTemperatureParams) SetColorTemperatureMireds(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setColorTemperatureMireds:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetocolortemperatureparams/optionsmask
func (m_ MTRColorControlClusterMoveToColorTemperatureParams) OptionsMask() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("optionsMask"))
	return rv
}


// SetOptionsMask sets the value of the optionsMask property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetocolortemperatureparams/optionsmask
func (m_ MTRColorControlClusterMoveToColorTemperatureParams) SetOptionsMask(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsMask:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetocolortemperatureparams/optionsoverride
func (m_ MTRColorControlClusterMoveToColorTemperatureParams) OptionsOverride() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("optionsOverride"))
	return rv
}


// SetOptionsOverride sets the value of the optionsOverride property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetocolortemperatureparams/optionsoverride
func (m_ MTRColorControlClusterMoveToColorTemperatureParams) SetOptionsOverride(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsOverride:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetocolortemperatureparams/serversideprocessingtimeout
func (m_ MTRColorControlClusterMoveToColorTemperatureParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetocolortemperatureparams/serversideprocessingtimeout
func (m_ MTRColorControlClusterMoveToColorTemperatureParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetocolortemperatureparams/timedinvoketimeoutms
func (m_ MTRColorControlClusterMoveToColorTemperatureParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetocolortemperatureparams/timedinvoketimeoutms
func (m_ MTRColorControlClusterMoveToColorTemperatureParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetocolortemperatureparams/transitiontime
func (m_ MTRColorControlClusterMoveToColorTemperatureParams) TransitionTime() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("transitionTime"))
	return rv
}


// SetTransitionTime sets the value of the transitionTime property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetocolortemperatureparams/transitiontime
func (m_ MTRColorControlClusterMoveToColorTemperatureParams) SetTransitionTime(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTransitionTime:"), value)
}



