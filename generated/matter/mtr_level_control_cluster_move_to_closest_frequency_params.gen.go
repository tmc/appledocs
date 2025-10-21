// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRLevelControlClusterMoveToClosestFrequencyParams] class.
var (
	MTRLevelControlClusterMoveToClosestFrequencyParamsClass     _MTRLevelControlClusterMoveToClosestFrequencyParamsClass
	MTRLevelControlClusterMoveToClosestFrequencyParamsClassOnce sync.Once
)

func getMTRLevelControlClusterMoveToClosestFrequencyParamsClass() _MTRLevelControlClusterMoveToClosestFrequencyParamsClass {
	MTRLevelControlClusterMoveToClosestFrequencyParamsClassOnce.Do(func() {
		MTRLevelControlClusterMoveToClosestFrequencyParamsClass = _MTRLevelControlClusterMoveToClosestFrequencyParamsClass{objc.GetClass("MTRLevelControlClusterMoveToClosestFrequencyParams")}
	})
	return MTRLevelControlClusterMoveToClosestFrequencyParamsClass
}

type _MTRLevelControlClusterMoveToClosestFrequencyParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRLevelControlClusterMoveToClosestFrequencyParams] class.
type IMTRLevelControlClusterMoveToClosestFrequencyParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterMoveToClosestFrequencyParams
type MTRLevelControlClusterMoveToClosestFrequencyParams struct {
	objectivec.Object
}

// MTRLevelControlClusterMoveToClosestFrequencyParamsFrom constructs a [MTRLevelControlClusterMoveToClosestFrequencyParams] from an unsafe.Pointer.
func MTRLevelControlClusterMoveToClosestFrequencyParamsFrom(ptr unsafe.Pointer) MTRLevelControlClusterMoveToClosestFrequencyParams {
	return MTRLevelControlClusterMoveToClosestFrequencyParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRLevelControlClusterMoveToClosestFrequencyParamsClass) Alloc() MTRLevelControlClusterMoveToClosestFrequencyParams {
	rv := objc.Send[MTRLevelControlClusterMoveToClosestFrequencyParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRLevelControlClusterMoveToClosestFrequencyParamsClass) New() MTRLevelControlClusterMoveToClosestFrequencyParams {
	rv := objc.Send[MTRLevelControlClusterMoveToClosestFrequencyParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRLevelControlClusterMoveToClosestFrequencyParams) Init() MTRLevelControlClusterMoveToClosestFrequencyParams {
	rv := objc.Send[MTRLevelControlClusterMoveToClosestFrequencyParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRLevelControlClusterMoveToClosestFrequencyParams) Autorelease() MTRLevelControlClusterMoveToClosestFrequencyParams {
	rv := objc.Send[MTRLevelControlClusterMoveToClosestFrequencyParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRLevelControlClusterMoveToClosestFrequencyParams creates a new MTRLevelControlClusterMoveToClosestFrequencyParams instance.
func NewMTRLevelControlClusterMoveToClosestFrequencyParams() MTRLevelControlClusterMoveToClosestFrequencyParams {
	return getMTRLevelControlClusterMoveToClosestFrequencyParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermovetoclosestfrequencyparams/frequency
func (m_ MTRLevelControlClusterMoveToClosestFrequencyParams) Frequency() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("frequency"))
	return rv
}


// SetFrequency sets the value of the frequency property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermovetoclosestfrequencyparams/frequency
func (m_ MTRLevelControlClusterMoveToClosestFrequencyParams) SetFrequency(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFrequency:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermovetoclosestfrequencyparams/serversideprocessingtimeout
func (m_ MTRLevelControlClusterMoveToClosestFrequencyParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermovetoclosestfrequencyparams/serversideprocessingtimeout
func (m_ MTRLevelControlClusterMoveToClosestFrequencyParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermovetoclosestfrequencyparams/timedinvoketimeoutms
func (m_ MTRLevelControlClusterMoveToClosestFrequencyParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermovetoclosestfrequencyparams/timedinvoketimeoutms
func (m_ MTRLevelControlClusterMoveToClosestFrequencyParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



