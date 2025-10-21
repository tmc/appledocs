// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRColorControlClusterMoveToSaturationParams] class.
var (
	MTRColorControlClusterMoveToSaturationParamsClass     _MTRColorControlClusterMoveToSaturationParamsClass
	MTRColorControlClusterMoveToSaturationParamsClassOnce sync.Once
)

func getMTRColorControlClusterMoveToSaturationParamsClass() _MTRColorControlClusterMoveToSaturationParamsClass {
	MTRColorControlClusterMoveToSaturationParamsClassOnce.Do(func() {
		MTRColorControlClusterMoveToSaturationParamsClass = _MTRColorControlClusterMoveToSaturationParamsClass{objc.GetClass("MTRColorControlClusterMoveToSaturationParams")}
	})
	return MTRColorControlClusterMoveToSaturationParamsClass
}

type _MTRColorControlClusterMoveToSaturationParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRColorControlClusterMoveToSaturationParams] class.
type IMTRColorControlClusterMoveToSaturationParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToSaturationParams
type MTRColorControlClusterMoveToSaturationParams struct {
	objectivec.Object
}

// MTRColorControlClusterMoveToSaturationParamsFrom constructs a [MTRColorControlClusterMoveToSaturationParams] from an unsafe.Pointer.
func MTRColorControlClusterMoveToSaturationParamsFrom(ptr unsafe.Pointer) MTRColorControlClusterMoveToSaturationParams {
	return MTRColorControlClusterMoveToSaturationParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRColorControlClusterMoveToSaturationParamsClass) Alloc() MTRColorControlClusterMoveToSaturationParams {
	rv := objc.Send[MTRColorControlClusterMoveToSaturationParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRColorControlClusterMoveToSaturationParamsClass) New() MTRColorControlClusterMoveToSaturationParams {
	rv := objc.Send[MTRColorControlClusterMoveToSaturationParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRColorControlClusterMoveToSaturationParams) Init() MTRColorControlClusterMoveToSaturationParams {
	rv := objc.Send[MTRColorControlClusterMoveToSaturationParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRColorControlClusterMoveToSaturationParams) Autorelease() MTRColorControlClusterMoveToSaturationParams {
	rv := objc.Send[MTRColorControlClusterMoveToSaturationParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRColorControlClusterMoveToSaturationParams creates a new MTRColorControlClusterMoveToSaturationParams instance.
func NewMTRColorControlClusterMoveToSaturationParams() MTRColorControlClusterMoveToSaturationParams {
	return getMTRColorControlClusterMoveToSaturationParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetosaturationparams/optionsmask
func (m_ MTRColorControlClusterMoveToSaturationParams) OptionsMask() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("optionsMask"))
	return rv
}


// SetOptionsMask sets the value of the optionsMask property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetosaturationparams/optionsmask
func (m_ MTRColorControlClusterMoveToSaturationParams) SetOptionsMask(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsMask:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetosaturationparams/optionsoverride
func (m_ MTRColorControlClusterMoveToSaturationParams) OptionsOverride() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("optionsOverride"))
	return rv
}


// SetOptionsOverride sets the value of the optionsOverride property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetosaturationparams/optionsoverride
func (m_ MTRColorControlClusterMoveToSaturationParams) SetOptionsOverride(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsOverride:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetosaturationparams/saturation
func (m_ MTRColorControlClusterMoveToSaturationParams) Saturation() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("saturation"))
	return rv
}


// SetSaturation sets the value of the saturation property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetosaturationparams/saturation
func (m_ MTRColorControlClusterMoveToSaturationParams) SetSaturation(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSaturation:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetosaturationparams/serversideprocessingtimeout
func (m_ MTRColorControlClusterMoveToSaturationParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetosaturationparams/serversideprocessingtimeout
func (m_ MTRColorControlClusterMoveToSaturationParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetosaturationparams/timedinvoketimeoutms
func (m_ MTRColorControlClusterMoveToSaturationParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetosaturationparams/timedinvoketimeoutms
func (m_ MTRColorControlClusterMoveToSaturationParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetosaturationparams/transitiontime
func (m_ MTRColorControlClusterMoveToSaturationParams) TransitionTime() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("transitionTime"))
	return rv
}


// SetTransitionTime sets the value of the transitionTime property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetosaturationparams/transitiontime
func (m_ MTRColorControlClusterMoveToSaturationParams) SetTransitionTime(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTransitionTime:"), value)
}



