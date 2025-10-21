// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRLevelControlClusterStopWithOnOffParams] class.
var (
	MTRLevelControlClusterStopWithOnOffParamsClass     _MTRLevelControlClusterStopWithOnOffParamsClass
	MTRLevelControlClusterStopWithOnOffParamsClassOnce sync.Once
)

func getMTRLevelControlClusterStopWithOnOffParamsClass() _MTRLevelControlClusterStopWithOnOffParamsClass {
	MTRLevelControlClusterStopWithOnOffParamsClassOnce.Do(func() {
		MTRLevelControlClusterStopWithOnOffParamsClass = _MTRLevelControlClusterStopWithOnOffParamsClass{objc.GetClass("MTRLevelControlClusterStopWithOnOffParams")}
	})
	return MTRLevelControlClusterStopWithOnOffParamsClass
}

type _MTRLevelControlClusterStopWithOnOffParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRLevelControlClusterStopWithOnOffParams] class.
type IMTRLevelControlClusterStopWithOnOffParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterStopWithOnOffParams
type MTRLevelControlClusterStopWithOnOffParams struct {
	objectivec.Object
}

// MTRLevelControlClusterStopWithOnOffParamsFrom constructs a [MTRLevelControlClusterStopWithOnOffParams] from an unsafe.Pointer.
func MTRLevelControlClusterStopWithOnOffParamsFrom(ptr unsafe.Pointer) MTRLevelControlClusterStopWithOnOffParams {
	return MTRLevelControlClusterStopWithOnOffParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRLevelControlClusterStopWithOnOffParamsClass) Alloc() MTRLevelControlClusterStopWithOnOffParams {
	rv := objc.Send[MTRLevelControlClusterStopWithOnOffParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRLevelControlClusterStopWithOnOffParamsClass) New() MTRLevelControlClusterStopWithOnOffParams {
	rv := objc.Send[MTRLevelControlClusterStopWithOnOffParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRLevelControlClusterStopWithOnOffParams) Init() MTRLevelControlClusterStopWithOnOffParams {
	rv := objc.Send[MTRLevelControlClusterStopWithOnOffParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRLevelControlClusterStopWithOnOffParams) Autorelease() MTRLevelControlClusterStopWithOnOffParams {
	rv := objc.Send[MTRLevelControlClusterStopWithOnOffParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRLevelControlClusterStopWithOnOffParams creates a new MTRLevelControlClusterStopWithOnOffParams instance.
func NewMTRLevelControlClusterStopWithOnOffParams() MTRLevelControlClusterStopWithOnOffParams {
	return getMTRLevelControlClusterStopWithOnOffParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclusterstopwithonoffparams/optionsmask
func (m_ MTRLevelControlClusterStopWithOnOffParams) OptionsMask() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("optionsMask"))
	return rv
}


// SetOptionsMask sets the value of the optionsMask property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclusterstopwithonoffparams/optionsmask
func (m_ MTRLevelControlClusterStopWithOnOffParams) SetOptionsMask(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsMask:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclusterstopwithonoffparams/serversideprocessingtimeout
func (m_ MTRLevelControlClusterStopWithOnOffParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclusterstopwithonoffparams/serversideprocessingtimeout
func (m_ MTRLevelControlClusterStopWithOnOffParams) SetServerSideProcessingTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclusterstopwithonoffparams/optionsoverride
func (m_ MTRLevelControlClusterStopWithOnOffParams) OptionsOverride() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("optionsOverride"))
	return rv
}


// SetOptionsOverride sets the value of the optionsOverride property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclusterstopwithonoffparams/optionsoverride
func (m_ MTRLevelControlClusterStopWithOnOffParams) SetOptionsOverride(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsOverride:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclusterstopwithonoffparams/timedinvoketimeoutms
func (m_ MTRLevelControlClusterStopWithOnOffParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclusterstopwithonoffparams/timedinvoketimeoutms
func (m_ MTRLevelControlClusterStopWithOnOffParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



