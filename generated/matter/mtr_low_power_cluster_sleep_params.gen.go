// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRLowPowerClusterSleepParams] class.
var (
	MTRLowPowerClusterSleepParamsClass     _MTRLowPowerClusterSleepParamsClass
	MTRLowPowerClusterSleepParamsClassOnce sync.Once
)

func getMTRLowPowerClusterSleepParamsClass() _MTRLowPowerClusterSleepParamsClass {
	MTRLowPowerClusterSleepParamsClassOnce.Do(func() {
		MTRLowPowerClusterSleepParamsClass = _MTRLowPowerClusterSleepParamsClass{objc.GetClass("MTRLowPowerClusterSleepParams")}
	})
	return MTRLowPowerClusterSleepParamsClass
}

type _MTRLowPowerClusterSleepParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRLowPowerClusterSleepParams] class.
type IMTRLowPowerClusterSleepParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLowPowerClusterSleepParams
type MTRLowPowerClusterSleepParams struct {
	objectivec.Object
}

// MTRLowPowerClusterSleepParamsFrom constructs a [MTRLowPowerClusterSleepParams] from an unsafe.Pointer.
func MTRLowPowerClusterSleepParamsFrom(ptr unsafe.Pointer) MTRLowPowerClusterSleepParams {
	return MTRLowPowerClusterSleepParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRLowPowerClusterSleepParamsClass) Alloc() MTRLowPowerClusterSleepParams {
	rv := objc.Send[MTRLowPowerClusterSleepParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRLowPowerClusterSleepParamsClass) New() MTRLowPowerClusterSleepParams {
	rv := objc.Send[MTRLowPowerClusterSleepParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRLowPowerClusterSleepParams) Init() MTRLowPowerClusterSleepParams {
	rv := objc.Send[MTRLowPowerClusterSleepParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRLowPowerClusterSleepParams) Autorelease() MTRLowPowerClusterSleepParams {
	rv := objc.Send[MTRLowPowerClusterSleepParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRLowPowerClusterSleepParams creates a new MTRLowPowerClusterSleepParams instance.
func NewMTRLowPowerClusterSleepParams() MTRLowPowerClusterSleepParams {
	return getMTRLowPowerClusterSleepParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlowpowerclustersleepparams/serversideprocessingtimeout
func (m_ MTRLowPowerClusterSleepParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlowpowerclustersleepparams/serversideprocessingtimeout
func (m_ MTRLowPowerClusterSleepParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlowpowerclustersleepparams/timedinvoketimeoutms
func (m_ MTRLowPowerClusterSleepParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlowpowerclustersleepparams/timedinvoketimeoutms
func (m_ MTRLowPowerClusterSleepParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



