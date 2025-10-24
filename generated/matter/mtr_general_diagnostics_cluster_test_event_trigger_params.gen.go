// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRGeneralDiagnosticsClusterTestEventTriggerParams] class.
var (
	MTRGeneralDiagnosticsClusterTestEventTriggerParamsClass     _MTRGeneralDiagnosticsClusterTestEventTriggerParamsClass
	MTRGeneralDiagnosticsClusterTestEventTriggerParamsClassOnce sync.Once
)

func getMTRGeneralDiagnosticsClusterTestEventTriggerParamsClass() _MTRGeneralDiagnosticsClusterTestEventTriggerParamsClass {
	MTRGeneralDiagnosticsClusterTestEventTriggerParamsClassOnce.Do(func() {
		MTRGeneralDiagnosticsClusterTestEventTriggerParamsClass = _MTRGeneralDiagnosticsClusterTestEventTriggerParamsClass{objc.GetClass("MTRGeneralDiagnosticsClusterTestEventTriggerParams")}
	})
	return MTRGeneralDiagnosticsClusterTestEventTriggerParamsClass
}

type _MTRGeneralDiagnosticsClusterTestEventTriggerParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRGeneralDiagnosticsClusterTestEventTriggerParams] class.
type IMTRGeneralDiagnosticsClusterTestEventTriggerParams interface {
	objectivec.IObject
	// properties:
	EnableKey() objc.IObject /* cross-framework: Data */
	SetEnableKey(value objc.IObject /* cross-framework: Data */)
	EventTrigger() objc.IObject /* cross-framework: NSNumber */
	SetEventTrigger(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGeneralDiagnosticsClusterTestEventTriggerParams
type MTRGeneralDiagnosticsClusterTestEventTriggerParams struct {
	objectivec.Object
}

// MTRGeneralDiagnosticsClusterTestEventTriggerParamsFrom constructs a [MTRGeneralDiagnosticsClusterTestEventTriggerParams] from an unsafe.Pointer.
func MTRGeneralDiagnosticsClusterTestEventTriggerParamsFrom(ptr unsafe.Pointer) MTRGeneralDiagnosticsClusterTestEventTriggerParams {
	return MTRGeneralDiagnosticsClusterTestEventTriggerParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRGeneralDiagnosticsClusterTestEventTriggerParamsClass) Alloc() MTRGeneralDiagnosticsClusterTestEventTriggerParams {
	rv := objc.Send[MTRGeneralDiagnosticsClusterTestEventTriggerParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRGeneralDiagnosticsClusterTestEventTriggerParamsClass) New() MTRGeneralDiagnosticsClusterTestEventTriggerParams {
	rv := objc.Send[MTRGeneralDiagnosticsClusterTestEventTriggerParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRGeneralDiagnosticsClusterTestEventTriggerParams) Init() MTRGeneralDiagnosticsClusterTestEventTriggerParams {
	rv := objc.Send[MTRGeneralDiagnosticsClusterTestEventTriggerParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRGeneralDiagnosticsClusterTestEventTriggerParams) Autorelease() MTRGeneralDiagnosticsClusterTestEventTriggerParams {
	rv := objc.Send[MTRGeneralDiagnosticsClusterTestEventTriggerParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRGeneralDiagnosticsClusterTestEventTriggerParams creates a new MTRGeneralDiagnosticsClusterTestEventTriggerParams instance.
func NewMTRGeneralDiagnosticsClusterTestEventTriggerParams() MTRGeneralDiagnosticsClusterTestEventTriggerParams {
	return getMTRGeneralDiagnosticsClusterTestEventTriggerParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclustertesteventtriggerparams/enablekey
func (m_ MTRGeneralDiagnosticsClusterTestEventTriggerParams) EnableKey() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("enableKey"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclustertesteventtriggerparams/enablekey
func (m_ MTRGeneralDiagnosticsClusterTestEventTriggerParams) SetEnableKey(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEnableKey:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclustertesteventtriggerparams/eventtrigger
func (m_ MTRGeneralDiagnosticsClusterTestEventTriggerParams) EventTrigger() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("eventTrigger"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclustertesteventtriggerparams/eventtrigger
func (m_ MTRGeneralDiagnosticsClusterTestEventTriggerParams) SetEventTrigger(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEventTrigger:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclustertesteventtriggerparams/serversideprocessingtimeout
func (m_ MTRGeneralDiagnosticsClusterTestEventTriggerParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclustertesteventtriggerparams/serversideprocessingtimeout
func (m_ MTRGeneralDiagnosticsClusterTestEventTriggerParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclustertesteventtriggerparams/timedinvoketimeoutms
func (m_ MTRGeneralDiagnosticsClusterTestEventTriggerParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclustertesteventtriggerparams/timedinvoketimeoutms
func (m_ MTRGeneralDiagnosticsClusterTestEventTriggerParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



