// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRColorControlClusterMoveToColorParams] class.
var (
	MTRColorControlClusterMoveToColorParamsClass     _MTRColorControlClusterMoveToColorParamsClass
	MTRColorControlClusterMoveToColorParamsClassOnce sync.Once
)

func getMTRColorControlClusterMoveToColorParamsClass() _MTRColorControlClusterMoveToColorParamsClass {
	MTRColorControlClusterMoveToColorParamsClassOnce.Do(func() {
		MTRColorControlClusterMoveToColorParamsClass = _MTRColorControlClusterMoveToColorParamsClass{objc.GetClass("MTRColorControlClusterMoveToColorParams")}
	})
	return MTRColorControlClusterMoveToColorParamsClass
}

type _MTRColorControlClusterMoveToColorParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRColorControlClusterMoveToColorParams] class.
type IMTRColorControlClusterMoveToColorParams interface {
	objectivec.IObject
	ColorX() foundation.Number
	SetColorX(value foundation.INumber)
	ColorY() foundation.Number
	SetColorY(value foundation.INumber)
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
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToColorParams
type MTRColorControlClusterMoveToColorParams struct {
	objectivec.Object
}

// MTRColorControlClusterMoveToColorParamsFrom constructs a [MTRColorControlClusterMoveToColorParams] from an unsafe.Pointer.
func MTRColorControlClusterMoveToColorParamsFrom(ptr unsafe.Pointer) MTRColorControlClusterMoveToColorParams {
	return MTRColorControlClusterMoveToColorParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRColorControlClusterMoveToColorParamsClass) Alloc() MTRColorControlClusterMoveToColorParams {
	rv := objc.Send[MTRColorControlClusterMoveToColorParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRColorControlClusterMoveToColorParamsClass) New() MTRColorControlClusterMoveToColorParams {
	rv := objc.Send[MTRColorControlClusterMoveToColorParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRColorControlClusterMoveToColorParams) Init() MTRColorControlClusterMoveToColorParams {
	rv := objc.Send[MTRColorControlClusterMoveToColorParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRColorControlClusterMoveToColorParams) Autorelease() MTRColorControlClusterMoveToColorParams {
	rv := objc.Send[MTRColorControlClusterMoveToColorParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRColorControlClusterMoveToColorParams creates a new MTRColorControlClusterMoveToColorParams instance.
func NewMTRColorControlClusterMoveToColorParams() MTRColorControlClusterMoveToColorParams {
	return getMTRColorControlClusterMoveToColorParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetocolorparams/colorx
func (m_ MTRColorControlClusterMoveToColorParams) ColorX() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("colorX"))
	return rv
}


// SetColorX sets the value of the colorX property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetocolorparams/colorx
func (m_ MTRColorControlClusterMoveToColorParams) SetColorX(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setColorX:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetocolorparams/colory
func (m_ MTRColorControlClusterMoveToColorParams) ColorY() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("colorY"))
	return rv
}


// SetColorY sets the value of the colorY property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetocolorparams/colory
func (m_ MTRColorControlClusterMoveToColorParams) SetColorY(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setColorY:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetocolorparams/optionsmask
func (m_ MTRColorControlClusterMoveToColorParams) OptionsMask() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("optionsMask"))
	return rv
}


// SetOptionsMask sets the value of the optionsMask property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetocolorparams/optionsmask
func (m_ MTRColorControlClusterMoveToColorParams) SetOptionsMask(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsMask:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetocolorparams/optionsoverride
func (m_ MTRColorControlClusterMoveToColorParams) OptionsOverride() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("optionsOverride"))
	return rv
}


// SetOptionsOverride sets the value of the optionsOverride property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetocolorparams/optionsoverride
func (m_ MTRColorControlClusterMoveToColorParams) SetOptionsOverride(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsOverride:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetocolorparams/serversideprocessingtimeout
func (m_ MTRColorControlClusterMoveToColorParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetocolorparams/serversideprocessingtimeout
func (m_ MTRColorControlClusterMoveToColorParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetocolorparams/timedinvoketimeoutms
func (m_ MTRColorControlClusterMoveToColorParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetocolorparams/timedinvoketimeoutms
func (m_ MTRColorControlClusterMoveToColorParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetocolorparams/transitiontime
func (m_ MTRColorControlClusterMoveToColorParams) TransitionTime() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("transitionTime"))
	return rv
}


// SetTransitionTime sets the value of the transitionTime property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetocolorparams/transitiontime
func (m_ MTRColorControlClusterMoveToColorParams) SetTransitionTime(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTransitionTime:"), value)
}



