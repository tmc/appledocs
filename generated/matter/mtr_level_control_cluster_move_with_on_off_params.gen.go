// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRLevelControlClusterMoveWithOnOffParams] class.
var (
	MTRLevelControlClusterMoveWithOnOffParamsClass     _MTRLevelControlClusterMoveWithOnOffParamsClass
	MTRLevelControlClusterMoveWithOnOffParamsClassOnce sync.Once
)

func getMTRLevelControlClusterMoveWithOnOffParamsClass() _MTRLevelControlClusterMoveWithOnOffParamsClass {
	MTRLevelControlClusterMoveWithOnOffParamsClassOnce.Do(func() {
		MTRLevelControlClusterMoveWithOnOffParamsClass = _MTRLevelControlClusterMoveWithOnOffParamsClass{objc.GetClass("MTRLevelControlClusterMoveWithOnOffParams")}
	})
	return MTRLevelControlClusterMoveWithOnOffParamsClass
}

type _MTRLevelControlClusterMoveWithOnOffParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRLevelControlClusterMoveWithOnOffParams] class.
type IMTRLevelControlClusterMoveWithOnOffParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterMoveWithOnOffParams
type MTRLevelControlClusterMoveWithOnOffParams struct {
	objectivec.Object
}

// MTRLevelControlClusterMoveWithOnOffParamsFrom constructs a [MTRLevelControlClusterMoveWithOnOffParams] from an unsafe.Pointer.
func MTRLevelControlClusterMoveWithOnOffParamsFrom(ptr unsafe.Pointer) MTRLevelControlClusterMoveWithOnOffParams {
	return MTRLevelControlClusterMoveWithOnOffParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRLevelControlClusterMoveWithOnOffParamsClass) Alloc() MTRLevelControlClusterMoveWithOnOffParams {
	rv := objc.Send[MTRLevelControlClusterMoveWithOnOffParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRLevelControlClusterMoveWithOnOffParamsClass) New() MTRLevelControlClusterMoveWithOnOffParams {
	rv := objc.Send[MTRLevelControlClusterMoveWithOnOffParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRLevelControlClusterMoveWithOnOffParams) Init() MTRLevelControlClusterMoveWithOnOffParams {
	rv := objc.Send[MTRLevelControlClusterMoveWithOnOffParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRLevelControlClusterMoveWithOnOffParams) Autorelease() MTRLevelControlClusterMoveWithOnOffParams {
	rv := objc.Send[MTRLevelControlClusterMoveWithOnOffParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRLevelControlClusterMoveWithOnOffParams creates a new MTRLevelControlClusterMoveWithOnOffParams instance.
func NewMTRLevelControlClusterMoveWithOnOffParams() MTRLevelControlClusterMoveWithOnOffParams {
	return getMTRLevelControlClusterMoveWithOnOffParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermovewithonoffparams/optionsoverride
func (m_ MTRLevelControlClusterMoveWithOnOffParams) OptionsOverride() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("optionsOverride"))
	return rv
}


// SetOptionsOverride sets the value of the optionsOverride property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermovewithonoffparams/optionsoverride
func (m_ MTRLevelControlClusterMoveWithOnOffParams) SetOptionsOverride(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsOverride:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermovewithonoffparams/optionsmask
func (m_ MTRLevelControlClusterMoveWithOnOffParams) OptionsMask() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("optionsMask"))
	return rv
}


// SetOptionsMask sets the value of the optionsMask property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermovewithonoffparams/optionsmask
func (m_ MTRLevelControlClusterMoveWithOnOffParams) SetOptionsMask(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsMask:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermovewithonoffparams/rate
func (m_ MTRLevelControlClusterMoveWithOnOffParams) Rate() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("rate"))
	return rv
}


// SetRate sets the value of the rate property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermovewithonoffparams/rate
func (m_ MTRLevelControlClusterMoveWithOnOffParams) SetRate(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRate:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermovewithonoffparams/timedinvoketimeoutms
func (m_ MTRLevelControlClusterMoveWithOnOffParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermovewithonoffparams/timedinvoketimeoutms
func (m_ MTRLevelControlClusterMoveWithOnOffParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermovewithonoffparams/movemode
func (m_ MTRLevelControlClusterMoveWithOnOffParams) MoveMode() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("moveMode"))
	return rv
}


// SetMoveMode sets the value of the moveMode property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermovewithonoffparams/movemode
func (m_ MTRLevelControlClusterMoveWithOnOffParams) SetMoveMode(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMoveMode:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermovewithonoffparams/serversideprocessingtimeout
func (m_ MTRLevelControlClusterMoveWithOnOffParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermovewithonoffparams/serversideprocessingtimeout
func (m_ MTRLevelControlClusterMoveWithOnOffParams) SetServerSideProcessingTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}



