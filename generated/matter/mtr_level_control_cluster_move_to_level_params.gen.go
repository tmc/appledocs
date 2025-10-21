// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRLevelControlClusterMoveToLevelParams] class.
var (
	MTRLevelControlClusterMoveToLevelParamsClass     _MTRLevelControlClusterMoveToLevelParamsClass
	MTRLevelControlClusterMoveToLevelParamsClassOnce sync.Once
)

func getMTRLevelControlClusterMoveToLevelParamsClass() _MTRLevelControlClusterMoveToLevelParamsClass {
	MTRLevelControlClusterMoveToLevelParamsClassOnce.Do(func() {
		MTRLevelControlClusterMoveToLevelParamsClass = _MTRLevelControlClusterMoveToLevelParamsClass{objc.GetClass("MTRLevelControlClusterMoveToLevelParams")}
	})
	return MTRLevelControlClusterMoveToLevelParamsClass
}

type _MTRLevelControlClusterMoveToLevelParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRLevelControlClusterMoveToLevelParams] class.
type IMTRLevelControlClusterMoveToLevelParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterMoveToLevelParams
type MTRLevelControlClusterMoveToLevelParams struct {
	objectivec.Object
}

// MTRLevelControlClusterMoveToLevelParamsFrom constructs a [MTRLevelControlClusterMoveToLevelParams] from an unsafe.Pointer.
func MTRLevelControlClusterMoveToLevelParamsFrom(ptr unsafe.Pointer) MTRLevelControlClusterMoveToLevelParams {
	return MTRLevelControlClusterMoveToLevelParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRLevelControlClusterMoveToLevelParamsClass) Alloc() MTRLevelControlClusterMoveToLevelParams {
	rv := objc.Send[MTRLevelControlClusterMoveToLevelParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRLevelControlClusterMoveToLevelParamsClass) New() MTRLevelControlClusterMoveToLevelParams {
	rv := objc.Send[MTRLevelControlClusterMoveToLevelParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRLevelControlClusterMoveToLevelParams) Init() MTRLevelControlClusterMoveToLevelParams {
	rv := objc.Send[MTRLevelControlClusterMoveToLevelParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRLevelControlClusterMoveToLevelParams) Autorelease() MTRLevelControlClusterMoveToLevelParams {
	rv := objc.Send[MTRLevelControlClusterMoveToLevelParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRLevelControlClusterMoveToLevelParams creates a new MTRLevelControlClusterMoveToLevelParams instance.
func NewMTRLevelControlClusterMoveToLevelParams() MTRLevelControlClusterMoveToLevelParams {
	return getMTRLevelControlClusterMoveToLevelParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermovetolevelparams/level
func (m_ MTRLevelControlClusterMoveToLevelParams) Level() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("level"))
	return rv
}


// SetLevel sets the value of the level property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermovetolevelparams/level
func (m_ MTRLevelControlClusterMoveToLevelParams) SetLevel(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLevel:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermovetolevelparams/optionsmask
func (m_ MTRLevelControlClusterMoveToLevelParams) OptionsMask() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("optionsMask"))
	return rv
}


// SetOptionsMask sets the value of the optionsMask property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermovetolevelparams/optionsmask
func (m_ MTRLevelControlClusterMoveToLevelParams) SetOptionsMask(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsMask:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermovetolevelparams/optionsoverride
func (m_ MTRLevelControlClusterMoveToLevelParams) OptionsOverride() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("optionsOverride"))
	return rv
}


// SetOptionsOverride sets the value of the optionsOverride property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermovetolevelparams/optionsoverride
func (m_ MTRLevelControlClusterMoveToLevelParams) SetOptionsOverride(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsOverride:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermovetolevelparams/serversideprocessingtimeout
func (m_ MTRLevelControlClusterMoveToLevelParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermovetolevelparams/serversideprocessingtimeout
func (m_ MTRLevelControlClusterMoveToLevelParams) SetServerSideProcessingTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermovetolevelparams/timedinvoketimeoutms
func (m_ MTRLevelControlClusterMoveToLevelParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermovetolevelparams/timedinvoketimeoutms
func (m_ MTRLevelControlClusterMoveToLevelParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermovetolevelparams/transitiontime
func (m_ MTRLevelControlClusterMoveToLevelParams) TransitionTime() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("transitionTime"))
	return rv
}


// SetTransitionTime sets the value of the transitionTime property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermovetolevelparams/transitiontime
func (m_ MTRLevelControlClusterMoveToLevelParams) SetTransitionTime(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTransitionTime:"), value)
}



