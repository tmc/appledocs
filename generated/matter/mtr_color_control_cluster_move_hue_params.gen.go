// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRColorControlClusterMoveHueParams] class.
var (
	MTRColorControlClusterMoveHueParamsClass     _MTRColorControlClusterMoveHueParamsClass
	MTRColorControlClusterMoveHueParamsClassOnce sync.Once
)

func getMTRColorControlClusterMoveHueParamsClass() _MTRColorControlClusterMoveHueParamsClass {
	MTRColorControlClusterMoveHueParamsClassOnce.Do(func() {
		MTRColorControlClusterMoveHueParamsClass = _MTRColorControlClusterMoveHueParamsClass{objc.GetClass("MTRColorControlClusterMoveHueParams")}
	})
	return MTRColorControlClusterMoveHueParamsClass
}

type _MTRColorControlClusterMoveHueParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRColorControlClusterMoveHueParams] class.
type IMTRColorControlClusterMoveHueParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveHueParams
type MTRColorControlClusterMoveHueParams struct {
	objectivec.Object
}

// MTRColorControlClusterMoveHueParamsFrom constructs a [MTRColorControlClusterMoveHueParams] from an unsafe.Pointer.
func MTRColorControlClusterMoveHueParamsFrom(ptr unsafe.Pointer) MTRColorControlClusterMoveHueParams {
	return MTRColorControlClusterMoveHueParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRColorControlClusterMoveHueParamsClass) Alloc() MTRColorControlClusterMoveHueParams {
	rv := objc.Send[MTRColorControlClusterMoveHueParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRColorControlClusterMoveHueParamsClass) New() MTRColorControlClusterMoveHueParams {
	rv := objc.Send[MTRColorControlClusterMoveHueParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRColorControlClusterMoveHueParams) Init() MTRColorControlClusterMoveHueParams {
	rv := objc.Send[MTRColorControlClusterMoveHueParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRColorControlClusterMoveHueParams) Autorelease() MTRColorControlClusterMoveHueParams {
	rv := objc.Send[MTRColorControlClusterMoveHueParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRColorControlClusterMoveHueParams creates a new MTRColorControlClusterMoveHueParams instance.
func NewMTRColorControlClusterMoveHueParams() MTRColorControlClusterMoveHueParams {
	return getMTRColorControlClusterMoveHueParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovehueparams/movemode
func (m_ MTRColorControlClusterMoveHueParams) MoveMode() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("moveMode"))
	return rv
}


// SetMoveMode sets the value of the moveMode property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovehueparams/movemode
func (m_ MTRColorControlClusterMoveHueParams) SetMoveMode(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMoveMode:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovehueparams/optionsmask
func (m_ MTRColorControlClusterMoveHueParams) OptionsMask() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("optionsMask"))
	return rv
}


// SetOptionsMask sets the value of the optionsMask property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovehueparams/optionsmask
func (m_ MTRColorControlClusterMoveHueParams) SetOptionsMask(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsMask:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovehueparams/optionsoverride
func (m_ MTRColorControlClusterMoveHueParams) OptionsOverride() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("optionsOverride"))
	return rv
}


// SetOptionsOverride sets the value of the optionsOverride property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovehueparams/optionsoverride
func (m_ MTRColorControlClusterMoveHueParams) SetOptionsOverride(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsOverride:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovehueparams/rate
func (m_ MTRColorControlClusterMoveHueParams) Rate() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("rate"))
	return rv
}


// SetRate sets the value of the rate property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovehueparams/rate
func (m_ MTRColorControlClusterMoveHueParams) SetRate(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRate:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovehueparams/serversideprocessingtimeout
func (m_ MTRColorControlClusterMoveHueParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovehueparams/serversideprocessingtimeout
func (m_ MTRColorControlClusterMoveHueParams) SetServerSideProcessingTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovehueparams/timedinvoketimeoutms
func (m_ MTRColorControlClusterMoveHueParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovehueparams/timedinvoketimeoutms
func (m_ MTRColorControlClusterMoveHueParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



