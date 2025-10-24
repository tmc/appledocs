// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRBooleanStateConfigurationClusterEnableDisableAlarmParams] class.
var (
	MTRBooleanStateConfigurationClusterEnableDisableAlarmParamsClass     _MTRBooleanStateConfigurationClusterEnableDisableAlarmParamsClass
	MTRBooleanStateConfigurationClusterEnableDisableAlarmParamsClassOnce sync.Once
)

func getMTRBooleanStateConfigurationClusterEnableDisableAlarmParamsClass() _MTRBooleanStateConfigurationClusterEnableDisableAlarmParamsClass {
	MTRBooleanStateConfigurationClusterEnableDisableAlarmParamsClassOnce.Do(func() {
		MTRBooleanStateConfigurationClusterEnableDisableAlarmParamsClass = _MTRBooleanStateConfigurationClusterEnableDisableAlarmParamsClass{objc.GetClass("MTRBooleanStateConfigurationClusterEnableDisableAlarmParams")}
	})
	return MTRBooleanStateConfigurationClusterEnableDisableAlarmParamsClass
}

type _MTRBooleanStateConfigurationClusterEnableDisableAlarmParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRBooleanStateConfigurationClusterEnableDisableAlarmParams] class.
type IMTRBooleanStateConfigurationClusterEnableDisableAlarmParams interface {
	objectivec.IObject
	// properties:
	AlarmsToEnableDisable() objc.IObject /* cross-framework: NSNumber */
	SetAlarmsToEnableDisable(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBooleanStateConfigurationClusterEnableDisableAlarmParams
type MTRBooleanStateConfigurationClusterEnableDisableAlarmParams struct {
	objectivec.Object
}

// MTRBooleanStateConfigurationClusterEnableDisableAlarmParamsFrom constructs a [MTRBooleanStateConfigurationClusterEnableDisableAlarmParams] from an unsafe.Pointer.
func MTRBooleanStateConfigurationClusterEnableDisableAlarmParamsFrom(ptr unsafe.Pointer) MTRBooleanStateConfigurationClusterEnableDisableAlarmParams {
	return MTRBooleanStateConfigurationClusterEnableDisableAlarmParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBooleanStateConfigurationClusterEnableDisableAlarmParamsClass) Alloc() MTRBooleanStateConfigurationClusterEnableDisableAlarmParams {
	rv := objc.Send[MTRBooleanStateConfigurationClusterEnableDisableAlarmParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBooleanStateConfigurationClusterEnableDisableAlarmParamsClass) New() MTRBooleanStateConfigurationClusterEnableDisableAlarmParams {
	rv := objc.Send[MTRBooleanStateConfigurationClusterEnableDisableAlarmParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBooleanStateConfigurationClusterEnableDisableAlarmParams) Init() MTRBooleanStateConfigurationClusterEnableDisableAlarmParams {
	rv := objc.Send[MTRBooleanStateConfigurationClusterEnableDisableAlarmParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBooleanStateConfigurationClusterEnableDisableAlarmParams) Autorelease() MTRBooleanStateConfigurationClusterEnableDisableAlarmParams {
	rv := objc.Send[MTRBooleanStateConfigurationClusterEnableDisableAlarmParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBooleanStateConfigurationClusterEnableDisableAlarmParams creates a new MTRBooleanStateConfigurationClusterEnableDisableAlarmParams instance.
func NewMTRBooleanStateConfigurationClusterEnableDisableAlarmParams() MTRBooleanStateConfigurationClusterEnableDisableAlarmParams {
	return getMTRBooleanStateConfigurationClusterEnableDisableAlarmParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbooleanstateconfigurationclusterenabledisablealarmparams/alarmstoenabledisable
func (m_ MTRBooleanStateConfigurationClusterEnableDisableAlarmParams) AlarmsToEnableDisable() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("alarmsToEnableDisable"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbooleanstateconfigurationclusterenabledisablealarmparams/alarmstoenabledisable
func (m_ MTRBooleanStateConfigurationClusterEnableDisableAlarmParams) SetAlarmsToEnableDisable(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAlarmsToEnableDisable:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbooleanstateconfigurationclusterenabledisablealarmparams/serversideprocessingtimeout
func (m_ MTRBooleanStateConfigurationClusterEnableDisableAlarmParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbooleanstateconfigurationclusterenabledisablealarmparams/serversideprocessingtimeout
func (m_ MTRBooleanStateConfigurationClusterEnableDisableAlarmParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbooleanstateconfigurationclusterenabledisablealarmparams/timedinvoketimeoutms
func (m_ MTRBooleanStateConfigurationClusterEnableDisableAlarmParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbooleanstateconfigurationclusterenabledisablealarmparams/timedinvoketimeoutms
func (m_ MTRBooleanStateConfigurationClusterEnableDisableAlarmParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



