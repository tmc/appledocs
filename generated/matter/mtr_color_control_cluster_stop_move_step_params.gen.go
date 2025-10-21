// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRColorControlClusterStopMoveStepParams] class.
var (
	MTRColorControlClusterStopMoveStepParamsClass     _MTRColorControlClusterStopMoveStepParamsClass
	MTRColorControlClusterStopMoveStepParamsClassOnce sync.Once
)

func getMTRColorControlClusterStopMoveStepParamsClass() _MTRColorControlClusterStopMoveStepParamsClass {
	MTRColorControlClusterStopMoveStepParamsClassOnce.Do(func() {
		MTRColorControlClusterStopMoveStepParamsClass = _MTRColorControlClusterStopMoveStepParamsClass{objc.GetClass("MTRColorControlClusterStopMoveStepParams")}
	})
	return MTRColorControlClusterStopMoveStepParamsClass
}

type _MTRColorControlClusterStopMoveStepParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRColorControlClusterStopMoveStepParams] class.
type IMTRColorControlClusterStopMoveStepParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterStopMoveStepParams
type MTRColorControlClusterStopMoveStepParams struct {
	objectivec.Object
}

// MTRColorControlClusterStopMoveStepParamsFrom constructs a [MTRColorControlClusterStopMoveStepParams] from an unsafe.Pointer.
func MTRColorControlClusterStopMoveStepParamsFrom(ptr unsafe.Pointer) MTRColorControlClusterStopMoveStepParams {
	return MTRColorControlClusterStopMoveStepParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRColorControlClusterStopMoveStepParamsClass) Alloc() MTRColorControlClusterStopMoveStepParams {
	rv := objc.Send[MTRColorControlClusterStopMoveStepParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRColorControlClusterStopMoveStepParamsClass) New() MTRColorControlClusterStopMoveStepParams {
	rv := objc.Send[MTRColorControlClusterStopMoveStepParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRColorControlClusterStopMoveStepParams) Init() MTRColorControlClusterStopMoveStepParams {
	rv := objc.Send[MTRColorControlClusterStopMoveStepParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRColorControlClusterStopMoveStepParams) Autorelease() MTRColorControlClusterStopMoveStepParams {
	rv := objc.Send[MTRColorControlClusterStopMoveStepParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRColorControlClusterStopMoveStepParams creates a new MTRColorControlClusterStopMoveStepParams instance.
func NewMTRColorControlClusterStopMoveStepParams() MTRColorControlClusterStopMoveStepParams {
	return getMTRColorControlClusterStopMoveStepParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstopmovestepparams/optionsmask
func (m_ MTRColorControlClusterStopMoveStepParams) OptionsMask() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("optionsMask"))
	return rv
}


// SetOptionsMask sets the value of the optionsMask property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstopmovestepparams/optionsmask
func (m_ MTRColorControlClusterStopMoveStepParams) SetOptionsMask(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsMask:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstopmovestepparams/optionsoverride
func (m_ MTRColorControlClusterStopMoveStepParams) OptionsOverride() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("optionsOverride"))
	return rv
}


// SetOptionsOverride sets the value of the optionsOverride property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstopmovestepparams/optionsoverride
func (m_ MTRColorControlClusterStopMoveStepParams) SetOptionsOverride(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsOverride:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstopmovestepparams/serversideprocessingtimeout
func (m_ MTRColorControlClusterStopMoveStepParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstopmovestepparams/serversideprocessingtimeout
func (m_ MTRColorControlClusterStopMoveStepParams) SetServerSideProcessingTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstopmovestepparams/timedinvoketimeoutms
func (m_ MTRColorControlClusterStopMoveStepParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstopmovestepparams/timedinvoketimeoutms
func (m_ MTRColorControlClusterStopMoveStepParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



