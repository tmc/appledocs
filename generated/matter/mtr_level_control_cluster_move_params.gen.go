// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRLevelControlClusterMoveParams] class.
var (
	MTRLevelControlClusterMoveParamsClass     _MTRLevelControlClusterMoveParamsClass
	MTRLevelControlClusterMoveParamsClassOnce sync.Once
)

func getMTRLevelControlClusterMoveParamsClass() _MTRLevelControlClusterMoveParamsClass {
	MTRLevelControlClusterMoveParamsClassOnce.Do(func() {
		MTRLevelControlClusterMoveParamsClass = _MTRLevelControlClusterMoveParamsClass{objc.GetClass("MTRLevelControlClusterMoveParams")}
	})
	return MTRLevelControlClusterMoveParamsClass
}

type _MTRLevelControlClusterMoveParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRLevelControlClusterMoveParams] class.
type IMTRLevelControlClusterMoveParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterMoveParams
type MTRLevelControlClusterMoveParams struct {
	objectivec.Object
}

// MTRLevelControlClusterMoveParamsFrom constructs a [MTRLevelControlClusterMoveParams] from an unsafe.Pointer.
func MTRLevelControlClusterMoveParamsFrom(ptr unsafe.Pointer) MTRLevelControlClusterMoveParams {
	return MTRLevelControlClusterMoveParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRLevelControlClusterMoveParamsClass) Alloc() MTRLevelControlClusterMoveParams {
	rv := objc.Send[MTRLevelControlClusterMoveParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRLevelControlClusterMoveParamsClass) New() MTRLevelControlClusterMoveParams {
	rv := objc.Send[MTRLevelControlClusterMoveParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRLevelControlClusterMoveParams) Init() MTRLevelControlClusterMoveParams {
	rv := objc.Send[MTRLevelControlClusterMoveParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRLevelControlClusterMoveParams) Autorelease() MTRLevelControlClusterMoveParams {
	rv := objc.Send[MTRLevelControlClusterMoveParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRLevelControlClusterMoveParams creates a new MTRLevelControlClusterMoveParams instance.
func NewMTRLevelControlClusterMoveParams() MTRLevelControlClusterMoveParams {
	return getMTRLevelControlClusterMoveParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermoveparams/movemode
func (m_ MTRLevelControlClusterMoveParams) MoveMode() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("moveMode"))
	return rv
}


// SetMoveMode sets the value of the moveMode property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermoveparams/movemode
func (m_ MTRLevelControlClusterMoveParams) SetMoveMode(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMoveMode:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermoveparams/timedinvoketimeoutms
func (m_ MTRLevelControlClusterMoveParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermoveparams/timedinvoketimeoutms
func (m_ MTRLevelControlClusterMoveParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermoveparams/rate
func (m_ MTRLevelControlClusterMoveParams) Rate() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("rate"))
	return rv
}


// SetRate sets the value of the rate property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermoveparams/rate
func (m_ MTRLevelControlClusterMoveParams) SetRate(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRate:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermoveparams/optionsoverride
func (m_ MTRLevelControlClusterMoveParams) OptionsOverride() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("optionsOverride"))
	return rv
}


// SetOptionsOverride sets the value of the optionsOverride property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermoveparams/optionsoverride
func (m_ MTRLevelControlClusterMoveParams) SetOptionsOverride(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsOverride:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermoveparams/serversideprocessingtimeout
func (m_ MTRLevelControlClusterMoveParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermoveparams/serversideprocessingtimeout
func (m_ MTRLevelControlClusterMoveParams) SetServerSideProcessingTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermoveparams/optionsmask
func (m_ MTRLevelControlClusterMoveParams) OptionsMask() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("optionsMask"))
	return rv
}


// SetOptionsMask sets the value of the optionsMask property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermoveparams/optionsmask
func (m_ MTRLevelControlClusterMoveParams) SetOptionsMask(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsMask:"), value)
}



