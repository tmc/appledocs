// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRColorControlClusterMoveToHueParams] class.
var (
	MTRColorControlClusterMoveToHueParamsClass     _MTRColorControlClusterMoveToHueParamsClass
	MTRColorControlClusterMoveToHueParamsClassOnce sync.Once
)

func getMTRColorControlClusterMoveToHueParamsClass() _MTRColorControlClusterMoveToHueParamsClass {
	MTRColorControlClusterMoveToHueParamsClassOnce.Do(func() {
		MTRColorControlClusterMoveToHueParamsClass = _MTRColorControlClusterMoveToHueParamsClass{objc.GetClass("MTRColorControlClusterMoveToHueParams")}
	})
	return MTRColorControlClusterMoveToHueParamsClass
}

type _MTRColorControlClusterMoveToHueParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRColorControlClusterMoveToHueParams] class.
type IMTRColorControlClusterMoveToHueParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToHueParams
type MTRColorControlClusterMoveToHueParams struct {
	objectivec.Object
}

// MTRColorControlClusterMoveToHueParamsFrom constructs a [MTRColorControlClusterMoveToHueParams] from an unsafe.Pointer.
func MTRColorControlClusterMoveToHueParamsFrom(ptr unsafe.Pointer) MTRColorControlClusterMoveToHueParams {
	return MTRColorControlClusterMoveToHueParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRColorControlClusterMoveToHueParamsClass) Alloc() MTRColorControlClusterMoveToHueParams {
	rv := objc.Send[MTRColorControlClusterMoveToHueParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRColorControlClusterMoveToHueParamsClass) New() MTRColorControlClusterMoveToHueParams {
	rv := objc.Send[MTRColorControlClusterMoveToHueParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRColorControlClusterMoveToHueParams) Init() MTRColorControlClusterMoveToHueParams {
	rv := objc.Send[MTRColorControlClusterMoveToHueParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRColorControlClusterMoveToHueParams) Autorelease() MTRColorControlClusterMoveToHueParams {
	rv := objc.Send[MTRColorControlClusterMoveToHueParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRColorControlClusterMoveToHueParams creates a new MTRColorControlClusterMoveToHueParams instance.
func NewMTRColorControlClusterMoveToHueParams() MTRColorControlClusterMoveToHueParams {
	return getMTRColorControlClusterMoveToHueParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetohueparams/optionsoverride
func (m_ MTRColorControlClusterMoveToHueParams) OptionsOverride() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("optionsOverride"))
	return rv
}


// SetOptionsOverride sets the value of the optionsOverride property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetohueparams/optionsoverride
func (m_ MTRColorControlClusterMoveToHueParams) SetOptionsOverride(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsOverride:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetohueparams/direction
func (m_ MTRColorControlClusterMoveToHueParams) Direction() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("direction"))
	return rv
}


// SetDirection sets the value of the direction property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetohueparams/direction
func (m_ MTRColorControlClusterMoveToHueParams) SetDirection(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDirection:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetohueparams/timedinvoketimeoutms
func (m_ MTRColorControlClusterMoveToHueParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetohueparams/timedinvoketimeoutms
func (m_ MTRColorControlClusterMoveToHueParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetohueparams/transitiontime
func (m_ MTRColorControlClusterMoveToHueParams) TransitionTime() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("transitionTime"))
	return rv
}


// SetTransitionTime sets the value of the transitionTime property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetohueparams/transitiontime
func (m_ MTRColorControlClusterMoveToHueParams) SetTransitionTime(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTransitionTime:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetohueparams/serversideprocessingtimeout
func (m_ MTRColorControlClusterMoveToHueParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetohueparams/serversideprocessingtimeout
func (m_ MTRColorControlClusterMoveToHueParams) SetServerSideProcessingTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetohueparams/hue
func (m_ MTRColorControlClusterMoveToHueParams) Hue() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("hue"))
	return rv
}


// SetHue sets the value of the hue property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetohueparams/hue
func (m_ MTRColorControlClusterMoveToHueParams) SetHue(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHue:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetohueparams/optionsmask
func (m_ MTRColorControlClusterMoveToHueParams) OptionsMask() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("optionsMask"))
	return rv
}


// SetOptionsMask sets the value of the optionsMask property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetohueparams/optionsmask
func (m_ MTRColorControlClusterMoveToHueParams) SetOptionsMask(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsMask:"), value)
}



