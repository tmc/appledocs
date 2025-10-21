// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRColorControlClusterMoveColorTemperatureParams] class.
var (
	MTRColorControlClusterMoveColorTemperatureParamsClass     _MTRColorControlClusterMoveColorTemperatureParamsClass
	MTRColorControlClusterMoveColorTemperatureParamsClassOnce sync.Once
)

func getMTRColorControlClusterMoveColorTemperatureParamsClass() _MTRColorControlClusterMoveColorTemperatureParamsClass {
	MTRColorControlClusterMoveColorTemperatureParamsClassOnce.Do(func() {
		MTRColorControlClusterMoveColorTemperatureParamsClass = _MTRColorControlClusterMoveColorTemperatureParamsClass{objc.GetClass("MTRColorControlClusterMoveColorTemperatureParams")}
	})
	return MTRColorControlClusterMoveColorTemperatureParamsClass
}

type _MTRColorControlClusterMoveColorTemperatureParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRColorControlClusterMoveColorTemperatureParams] class.
type IMTRColorControlClusterMoveColorTemperatureParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveColorTemperatureParams
type MTRColorControlClusterMoveColorTemperatureParams struct {
	objectivec.Object
}

// MTRColorControlClusterMoveColorTemperatureParamsFrom constructs a [MTRColorControlClusterMoveColorTemperatureParams] from an unsafe.Pointer.
func MTRColorControlClusterMoveColorTemperatureParamsFrom(ptr unsafe.Pointer) MTRColorControlClusterMoveColorTemperatureParams {
	return MTRColorControlClusterMoveColorTemperatureParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRColorControlClusterMoveColorTemperatureParamsClass) Alloc() MTRColorControlClusterMoveColorTemperatureParams {
	rv := objc.Send[MTRColorControlClusterMoveColorTemperatureParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRColorControlClusterMoveColorTemperatureParamsClass) New() MTRColorControlClusterMoveColorTemperatureParams {
	rv := objc.Send[MTRColorControlClusterMoveColorTemperatureParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRColorControlClusterMoveColorTemperatureParams) Init() MTRColorControlClusterMoveColorTemperatureParams {
	rv := objc.Send[MTRColorControlClusterMoveColorTemperatureParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRColorControlClusterMoveColorTemperatureParams) Autorelease() MTRColorControlClusterMoveColorTemperatureParams {
	rv := objc.Send[MTRColorControlClusterMoveColorTemperatureParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRColorControlClusterMoveColorTemperatureParams creates a new MTRColorControlClusterMoveColorTemperatureParams instance.
func NewMTRColorControlClusterMoveColorTemperatureParams() MTRColorControlClusterMoveColorTemperatureParams {
	return getMTRColorControlClusterMoveColorTemperatureParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovecolortemperatureparams/colortemperaturemaximummireds
func (m_ MTRColorControlClusterMoveColorTemperatureParams) ColorTemperatureMaximumMireds() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("colorTemperatureMaximumMireds"))
	return rv
}


// SetColorTemperatureMaximumMireds sets the value of the colorTemperatureMaximumMireds property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovecolortemperatureparams/colortemperaturemaximummireds
func (m_ MTRColorControlClusterMoveColorTemperatureParams) SetColorTemperatureMaximumMireds(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setColorTemperatureMaximumMireds:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovecolortemperatureparams/colortemperatureminimummireds
func (m_ MTRColorControlClusterMoveColorTemperatureParams) ColorTemperatureMinimumMireds() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("colorTemperatureMinimumMireds"))
	return rv
}


// SetColorTemperatureMinimumMireds sets the value of the colorTemperatureMinimumMireds property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovecolortemperatureparams/colortemperatureminimummireds
func (m_ MTRColorControlClusterMoveColorTemperatureParams) SetColorTemperatureMinimumMireds(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setColorTemperatureMinimumMireds:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovecolortemperatureparams/movemode
func (m_ MTRColorControlClusterMoveColorTemperatureParams) MoveMode() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("moveMode"))
	return rv
}


// SetMoveMode sets the value of the moveMode property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovecolortemperatureparams/movemode
func (m_ MTRColorControlClusterMoveColorTemperatureParams) SetMoveMode(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMoveMode:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovecolortemperatureparams/optionsmask
func (m_ MTRColorControlClusterMoveColorTemperatureParams) OptionsMask() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("optionsMask"))
	return rv
}


// SetOptionsMask sets the value of the optionsMask property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovecolortemperatureparams/optionsmask
func (m_ MTRColorControlClusterMoveColorTemperatureParams) SetOptionsMask(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsMask:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovecolortemperatureparams/optionsoverride
func (m_ MTRColorControlClusterMoveColorTemperatureParams) OptionsOverride() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("optionsOverride"))
	return rv
}


// SetOptionsOverride sets the value of the optionsOverride property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovecolortemperatureparams/optionsoverride
func (m_ MTRColorControlClusterMoveColorTemperatureParams) SetOptionsOverride(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsOverride:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovecolortemperatureparams/rate
func (m_ MTRColorControlClusterMoveColorTemperatureParams) Rate() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("rate"))
	return rv
}


// SetRate sets the value of the rate property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovecolortemperatureparams/rate
func (m_ MTRColorControlClusterMoveColorTemperatureParams) SetRate(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRate:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovecolortemperatureparams/serversideprocessingtimeout
func (m_ MTRColorControlClusterMoveColorTemperatureParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovecolortemperatureparams/serversideprocessingtimeout
func (m_ MTRColorControlClusterMoveColorTemperatureParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovecolortemperatureparams/timedinvoketimeoutms
func (m_ MTRColorControlClusterMoveColorTemperatureParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovecolortemperatureparams/timedinvoketimeoutms
func (m_ MTRColorControlClusterMoveColorTemperatureParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



