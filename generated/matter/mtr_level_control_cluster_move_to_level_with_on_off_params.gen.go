// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRLevelControlClusterMoveToLevelWithOnOffParams] class.
var (
	MTRLevelControlClusterMoveToLevelWithOnOffParamsClass     _MTRLevelControlClusterMoveToLevelWithOnOffParamsClass
	MTRLevelControlClusterMoveToLevelWithOnOffParamsClassOnce sync.Once
)

func getMTRLevelControlClusterMoveToLevelWithOnOffParamsClass() _MTRLevelControlClusterMoveToLevelWithOnOffParamsClass {
	MTRLevelControlClusterMoveToLevelWithOnOffParamsClassOnce.Do(func() {
		MTRLevelControlClusterMoveToLevelWithOnOffParamsClass = _MTRLevelControlClusterMoveToLevelWithOnOffParamsClass{objc.GetClass("MTRLevelControlClusterMoveToLevelWithOnOffParams")}
	})
	return MTRLevelControlClusterMoveToLevelWithOnOffParamsClass
}

type _MTRLevelControlClusterMoveToLevelWithOnOffParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRLevelControlClusterMoveToLevelWithOnOffParams] class.
type IMTRLevelControlClusterMoveToLevelWithOnOffParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterMoveToLevelWithOnOffParams
type MTRLevelControlClusterMoveToLevelWithOnOffParams struct {
	objectivec.Object
}

// MTRLevelControlClusterMoveToLevelWithOnOffParamsFrom constructs a [MTRLevelControlClusterMoveToLevelWithOnOffParams] from an unsafe.Pointer.
func MTRLevelControlClusterMoveToLevelWithOnOffParamsFrom(ptr unsafe.Pointer) MTRLevelControlClusterMoveToLevelWithOnOffParams {
	return MTRLevelControlClusterMoveToLevelWithOnOffParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRLevelControlClusterMoveToLevelWithOnOffParamsClass) Alloc() MTRLevelControlClusterMoveToLevelWithOnOffParams {
	rv := objc.Send[MTRLevelControlClusterMoveToLevelWithOnOffParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRLevelControlClusterMoveToLevelWithOnOffParamsClass) New() MTRLevelControlClusterMoveToLevelWithOnOffParams {
	rv := objc.Send[MTRLevelControlClusterMoveToLevelWithOnOffParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRLevelControlClusterMoveToLevelWithOnOffParams) Init() MTRLevelControlClusterMoveToLevelWithOnOffParams {
	rv := objc.Send[MTRLevelControlClusterMoveToLevelWithOnOffParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRLevelControlClusterMoveToLevelWithOnOffParams) Autorelease() MTRLevelControlClusterMoveToLevelWithOnOffParams {
	rv := objc.Send[MTRLevelControlClusterMoveToLevelWithOnOffParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRLevelControlClusterMoveToLevelWithOnOffParams creates a new MTRLevelControlClusterMoveToLevelWithOnOffParams instance.
func NewMTRLevelControlClusterMoveToLevelWithOnOffParams() MTRLevelControlClusterMoveToLevelWithOnOffParams {
	return getMTRLevelControlClusterMoveToLevelWithOnOffParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermovetolevelwithonoffparams/level
func (m_ MTRLevelControlClusterMoveToLevelWithOnOffParams) Level() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("level"))
	return rv
}


// SetLevel sets the value of the level property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermovetolevelwithonoffparams/level
func (m_ MTRLevelControlClusterMoveToLevelWithOnOffParams) SetLevel(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLevel:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermovetolevelwithonoffparams/optionsmask
func (m_ MTRLevelControlClusterMoveToLevelWithOnOffParams) OptionsMask() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("optionsMask"))
	return rv
}


// SetOptionsMask sets the value of the optionsMask property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermovetolevelwithonoffparams/optionsmask
func (m_ MTRLevelControlClusterMoveToLevelWithOnOffParams) SetOptionsMask(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsMask:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermovetolevelwithonoffparams/optionsoverride
func (m_ MTRLevelControlClusterMoveToLevelWithOnOffParams) OptionsOverride() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("optionsOverride"))
	return rv
}


// SetOptionsOverride sets the value of the optionsOverride property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermovetolevelwithonoffparams/optionsoverride
func (m_ MTRLevelControlClusterMoveToLevelWithOnOffParams) SetOptionsOverride(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsOverride:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermovetolevelwithonoffparams/serversideprocessingtimeout
func (m_ MTRLevelControlClusterMoveToLevelWithOnOffParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermovetolevelwithonoffparams/serversideprocessingtimeout
func (m_ MTRLevelControlClusterMoveToLevelWithOnOffParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermovetolevelwithonoffparams/timedinvoketimeoutms
func (m_ MTRLevelControlClusterMoveToLevelWithOnOffParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermovetolevelwithonoffparams/timedinvoketimeoutms
func (m_ MTRLevelControlClusterMoveToLevelWithOnOffParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermovetolevelwithonoffparams/transitiontime
func (m_ MTRLevelControlClusterMoveToLevelWithOnOffParams) TransitionTime() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("transitionTime"))
	return rv
}


// SetTransitionTime sets the value of the transitionTime property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermovetolevelwithonoffparams/transitiontime
func (m_ MTRLevelControlClusterMoveToLevelWithOnOffParams) SetTransitionTime(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTransitionTime:"), value)
}



