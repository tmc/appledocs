// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRColorControlClusterEnhancedMoveHueParams] class.
var (
	MTRColorControlClusterEnhancedMoveHueParamsClass     _MTRColorControlClusterEnhancedMoveHueParamsClass
	MTRColorControlClusterEnhancedMoveHueParamsClassOnce sync.Once
)

func getMTRColorControlClusterEnhancedMoveHueParamsClass() _MTRColorControlClusterEnhancedMoveHueParamsClass {
	MTRColorControlClusterEnhancedMoveHueParamsClassOnce.Do(func() {
		MTRColorControlClusterEnhancedMoveHueParamsClass = _MTRColorControlClusterEnhancedMoveHueParamsClass{objc.GetClass("MTRColorControlClusterEnhancedMoveHueParams")}
	})
	return MTRColorControlClusterEnhancedMoveHueParamsClass
}

type _MTRColorControlClusterEnhancedMoveHueParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRColorControlClusterEnhancedMoveHueParams] class.
type IMTRColorControlClusterEnhancedMoveHueParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterEnhancedMoveHueParams
type MTRColorControlClusterEnhancedMoveHueParams struct {
	objectivec.Object
}

// MTRColorControlClusterEnhancedMoveHueParamsFrom constructs a [MTRColorControlClusterEnhancedMoveHueParams] from an unsafe.Pointer.
func MTRColorControlClusterEnhancedMoveHueParamsFrom(ptr unsafe.Pointer) MTRColorControlClusterEnhancedMoveHueParams {
	return MTRColorControlClusterEnhancedMoveHueParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRColorControlClusterEnhancedMoveHueParamsClass) Alloc() MTRColorControlClusterEnhancedMoveHueParams {
	rv := objc.Send[MTRColorControlClusterEnhancedMoveHueParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRColorControlClusterEnhancedMoveHueParamsClass) New() MTRColorControlClusterEnhancedMoveHueParams {
	rv := objc.Send[MTRColorControlClusterEnhancedMoveHueParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRColorControlClusterEnhancedMoveHueParams) Init() MTRColorControlClusterEnhancedMoveHueParams {
	rv := objc.Send[MTRColorControlClusterEnhancedMoveHueParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRColorControlClusterEnhancedMoveHueParams) Autorelease() MTRColorControlClusterEnhancedMoveHueParams {
	rv := objc.Send[MTRColorControlClusterEnhancedMoveHueParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRColorControlClusterEnhancedMoveHueParams creates a new MTRColorControlClusterEnhancedMoveHueParams instance.
func NewMTRColorControlClusterEnhancedMoveHueParams() MTRColorControlClusterEnhancedMoveHueParams {
	return getMTRColorControlClusterEnhancedMoveHueParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedmovehueparams/movemode
func (m_ MTRColorControlClusterEnhancedMoveHueParams) MoveMode() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("moveMode"))
	return rv
}


// SetMoveMode sets the value of the moveMode property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedmovehueparams/movemode
func (m_ MTRColorControlClusterEnhancedMoveHueParams) SetMoveMode(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMoveMode:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedmovehueparams/optionsmask
func (m_ MTRColorControlClusterEnhancedMoveHueParams) OptionsMask() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("optionsMask"))
	return rv
}


// SetOptionsMask sets the value of the optionsMask property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedmovehueparams/optionsmask
func (m_ MTRColorControlClusterEnhancedMoveHueParams) SetOptionsMask(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsMask:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedmovehueparams/optionsoverride
func (m_ MTRColorControlClusterEnhancedMoveHueParams) OptionsOverride() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("optionsOverride"))
	return rv
}


// SetOptionsOverride sets the value of the optionsOverride property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedmovehueparams/optionsoverride
func (m_ MTRColorControlClusterEnhancedMoveHueParams) SetOptionsOverride(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsOverride:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedmovehueparams/rate
func (m_ MTRColorControlClusterEnhancedMoveHueParams) Rate() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("rate"))
	return rv
}


// SetRate sets the value of the rate property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedmovehueparams/rate
func (m_ MTRColorControlClusterEnhancedMoveHueParams) SetRate(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRate:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedmovehueparams/serversideprocessingtimeout
func (m_ MTRColorControlClusterEnhancedMoveHueParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedmovehueparams/serversideprocessingtimeout
func (m_ MTRColorControlClusterEnhancedMoveHueParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedmovehueparams/timedinvoketimeoutms
func (m_ MTRColorControlClusterEnhancedMoveHueParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedmovehueparams/timedinvoketimeoutms
func (m_ MTRColorControlClusterEnhancedMoveHueParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



