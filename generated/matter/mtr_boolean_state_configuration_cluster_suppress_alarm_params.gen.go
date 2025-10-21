// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRBooleanStateConfigurationClusterSuppressAlarmParams] class.
var (
	MTRBooleanStateConfigurationClusterSuppressAlarmParamsClass     _MTRBooleanStateConfigurationClusterSuppressAlarmParamsClass
	MTRBooleanStateConfigurationClusterSuppressAlarmParamsClassOnce sync.Once
)

func getMTRBooleanStateConfigurationClusterSuppressAlarmParamsClass() _MTRBooleanStateConfigurationClusterSuppressAlarmParamsClass {
	MTRBooleanStateConfigurationClusterSuppressAlarmParamsClassOnce.Do(func() {
		MTRBooleanStateConfigurationClusterSuppressAlarmParamsClass = _MTRBooleanStateConfigurationClusterSuppressAlarmParamsClass{objc.GetClass("MTRBooleanStateConfigurationClusterSuppressAlarmParams")}
	})
	return MTRBooleanStateConfigurationClusterSuppressAlarmParamsClass
}

type _MTRBooleanStateConfigurationClusterSuppressAlarmParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRBooleanStateConfigurationClusterSuppressAlarmParams] class.
type IMTRBooleanStateConfigurationClusterSuppressAlarmParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBooleanStateConfigurationClusterSuppressAlarmParams
type MTRBooleanStateConfigurationClusterSuppressAlarmParams struct {
	objectivec.Object
}

// MTRBooleanStateConfigurationClusterSuppressAlarmParamsFrom constructs a [MTRBooleanStateConfigurationClusterSuppressAlarmParams] from an unsafe.Pointer.
func MTRBooleanStateConfigurationClusterSuppressAlarmParamsFrom(ptr unsafe.Pointer) MTRBooleanStateConfigurationClusterSuppressAlarmParams {
	return MTRBooleanStateConfigurationClusterSuppressAlarmParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBooleanStateConfigurationClusterSuppressAlarmParamsClass) Alloc() MTRBooleanStateConfigurationClusterSuppressAlarmParams {
	rv := objc.Send[MTRBooleanStateConfigurationClusterSuppressAlarmParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBooleanStateConfigurationClusterSuppressAlarmParamsClass) New() MTRBooleanStateConfigurationClusterSuppressAlarmParams {
	rv := objc.Send[MTRBooleanStateConfigurationClusterSuppressAlarmParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBooleanStateConfigurationClusterSuppressAlarmParams) Init() MTRBooleanStateConfigurationClusterSuppressAlarmParams {
	rv := objc.Send[MTRBooleanStateConfigurationClusterSuppressAlarmParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBooleanStateConfigurationClusterSuppressAlarmParams) Autorelease() MTRBooleanStateConfigurationClusterSuppressAlarmParams {
	rv := objc.Send[MTRBooleanStateConfigurationClusterSuppressAlarmParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBooleanStateConfigurationClusterSuppressAlarmParams creates a new MTRBooleanStateConfigurationClusterSuppressAlarmParams instance.
func NewMTRBooleanStateConfigurationClusterSuppressAlarmParams() MTRBooleanStateConfigurationClusterSuppressAlarmParams {
	return getMTRBooleanStateConfigurationClusterSuppressAlarmParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbooleanstateconfigurationclustersuppressalarmparams/alarmstosuppress
func (m_ MTRBooleanStateConfigurationClusterSuppressAlarmParams) AlarmsToSuppress() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("alarmsToSuppress"))
	return rv
}


// SetAlarmsToSuppress sets the value of the alarmsToSuppress property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbooleanstateconfigurationclustersuppressalarmparams/alarmstosuppress
func (m_ MTRBooleanStateConfigurationClusterSuppressAlarmParams) SetAlarmsToSuppress(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAlarmsToSuppress:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbooleanstateconfigurationclustersuppressalarmparams/serversideprocessingtimeout
func (m_ MTRBooleanStateConfigurationClusterSuppressAlarmParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbooleanstateconfigurationclustersuppressalarmparams/serversideprocessingtimeout
func (m_ MTRBooleanStateConfigurationClusterSuppressAlarmParams) SetServerSideProcessingTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbooleanstateconfigurationclustersuppressalarmparams/timedinvoketimeoutms
func (m_ MTRBooleanStateConfigurationClusterSuppressAlarmParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbooleanstateconfigurationclustersuppressalarmparams/timedinvoketimeoutms
func (m_ MTRBooleanStateConfigurationClusterSuppressAlarmParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



