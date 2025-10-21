// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRColorControlClusterColorLoopSetParams] class.
var (
	MTRColorControlClusterColorLoopSetParamsClass     _MTRColorControlClusterColorLoopSetParamsClass
	MTRColorControlClusterColorLoopSetParamsClassOnce sync.Once
)

func getMTRColorControlClusterColorLoopSetParamsClass() _MTRColorControlClusterColorLoopSetParamsClass {
	MTRColorControlClusterColorLoopSetParamsClassOnce.Do(func() {
		MTRColorControlClusterColorLoopSetParamsClass = _MTRColorControlClusterColorLoopSetParamsClass{objc.GetClass("MTRColorControlClusterColorLoopSetParams")}
	})
	return MTRColorControlClusterColorLoopSetParamsClass
}

type _MTRColorControlClusterColorLoopSetParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRColorControlClusterColorLoopSetParams] class.
type IMTRColorControlClusterColorLoopSetParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterColorLoopSetParams
type MTRColorControlClusterColorLoopSetParams struct {
	objectivec.Object
}

// MTRColorControlClusterColorLoopSetParamsFrom constructs a [MTRColorControlClusterColorLoopSetParams] from an unsafe.Pointer.
func MTRColorControlClusterColorLoopSetParamsFrom(ptr unsafe.Pointer) MTRColorControlClusterColorLoopSetParams {
	return MTRColorControlClusterColorLoopSetParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRColorControlClusterColorLoopSetParamsClass) Alloc() MTRColorControlClusterColorLoopSetParams {
	rv := objc.Send[MTRColorControlClusterColorLoopSetParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRColorControlClusterColorLoopSetParamsClass) New() MTRColorControlClusterColorLoopSetParams {
	rv := objc.Send[MTRColorControlClusterColorLoopSetParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRColorControlClusterColorLoopSetParams) Init() MTRColorControlClusterColorLoopSetParams {
	rv := objc.Send[MTRColorControlClusterColorLoopSetParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRColorControlClusterColorLoopSetParams) Autorelease() MTRColorControlClusterColorLoopSetParams {
	rv := objc.Send[MTRColorControlClusterColorLoopSetParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRColorControlClusterColorLoopSetParams creates a new MTRColorControlClusterColorLoopSetParams instance.
func NewMTRColorControlClusterColorLoopSetParams() MTRColorControlClusterColorLoopSetParams {
	return getMTRColorControlClusterColorLoopSetParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustercolorloopsetparams/optionsmask
func (m_ MTRColorControlClusterColorLoopSetParams) OptionsMask() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("optionsMask"))
	return rv
}


// SetOptionsMask sets the value of the optionsMask property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustercolorloopsetparams/optionsmask
func (m_ MTRColorControlClusterColorLoopSetParams) SetOptionsMask(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsMask:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustercolorloopsetparams/serversideprocessingtimeout
func (m_ MTRColorControlClusterColorLoopSetParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustercolorloopsetparams/serversideprocessingtimeout
func (m_ MTRColorControlClusterColorLoopSetParams) SetServerSideProcessingTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustercolorloopsetparams/time
func (m_ MTRColorControlClusterColorLoopSetParams) Time() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("time"))
	return rv
}


// SetTime sets the value of the time property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustercolorloopsetparams/time
func (m_ MTRColorControlClusterColorLoopSetParams) SetTime(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTime:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustercolorloopsetparams/optionsoverride
func (m_ MTRColorControlClusterColorLoopSetParams) OptionsOverride() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("optionsOverride"))
	return rv
}


// SetOptionsOverride sets the value of the optionsOverride property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustercolorloopsetparams/optionsoverride
func (m_ MTRColorControlClusterColorLoopSetParams) SetOptionsOverride(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsOverride:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustercolorloopsetparams/action
func (m_ MTRColorControlClusterColorLoopSetParams) Action() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("action"))
	return rv
}


// SetAction sets the value of the action property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustercolorloopsetparams/action
func (m_ MTRColorControlClusterColorLoopSetParams) SetAction(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAction:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustercolorloopsetparams/direction
func (m_ MTRColorControlClusterColorLoopSetParams) Direction() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("direction"))
	return rv
}


// SetDirection sets the value of the direction property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustercolorloopsetparams/direction
func (m_ MTRColorControlClusterColorLoopSetParams) SetDirection(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDirection:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustercolorloopsetparams/timedinvoketimeoutms
func (m_ MTRColorControlClusterColorLoopSetParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustercolorloopsetparams/timedinvoketimeoutms
func (m_ MTRColorControlClusterColorLoopSetParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustercolorloopsetparams/updateflags
func (m_ MTRColorControlClusterColorLoopSetParams) UpdateFlags() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("updateFlags"))
	return rv
}


// SetUpdateFlags sets the value of the updateFlags property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustercolorloopsetparams/updateflags
func (m_ MTRColorControlClusterColorLoopSetParams) SetUpdateFlags(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUpdateFlags:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustercolorloopsetparams/starthue
func (m_ MTRColorControlClusterColorLoopSetParams) StartHue() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("startHue"))
	return rv
}


// SetStartHue sets the value of the startHue property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustercolorloopsetparams/starthue
func (m_ MTRColorControlClusterColorLoopSetParams) SetStartHue(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStartHue:"), value)
}



