// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRDishwasherAlarmClusterResetParams] class.
var (
	MTRDishwasherAlarmClusterResetParamsClass     _MTRDishwasherAlarmClusterResetParamsClass
	MTRDishwasherAlarmClusterResetParamsClassOnce sync.Once
)

func getMTRDishwasherAlarmClusterResetParamsClass() _MTRDishwasherAlarmClusterResetParamsClass {
	MTRDishwasherAlarmClusterResetParamsClassOnce.Do(func() {
		MTRDishwasherAlarmClusterResetParamsClass = _MTRDishwasherAlarmClusterResetParamsClass{objc.GetClass("MTRDishwasherAlarmClusterResetParams")}
	})
	return MTRDishwasherAlarmClusterResetParamsClass
}

type _MTRDishwasherAlarmClusterResetParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRDishwasherAlarmClusterResetParams] class.
type IMTRDishwasherAlarmClusterResetParams interface {
	objectivec.IObject
	// properties:
	Alarms() objc.IObject /* cross-framework: NSNumber */
	SetAlarms(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherAlarmClusterResetParams
type MTRDishwasherAlarmClusterResetParams struct {
	objectivec.Object
}

// MTRDishwasherAlarmClusterResetParamsFrom constructs a [MTRDishwasherAlarmClusterResetParams] from an unsafe.Pointer.
func MTRDishwasherAlarmClusterResetParamsFrom(ptr unsafe.Pointer) MTRDishwasherAlarmClusterResetParams {
	return MTRDishwasherAlarmClusterResetParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDishwasherAlarmClusterResetParamsClass) Alloc() MTRDishwasherAlarmClusterResetParams {
	rv := objc.Send[MTRDishwasherAlarmClusterResetParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDishwasherAlarmClusterResetParamsClass) New() MTRDishwasherAlarmClusterResetParams {
	rv := objc.Send[MTRDishwasherAlarmClusterResetParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDishwasherAlarmClusterResetParams) Init() MTRDishwasherAlarmClusterResetParams {
	rv := objc.Send[MTRDishwasherAlarmClusterResetParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDishwasherAlarmClusterResetParams) Autorelease() MTRDishwasherAlarmClusterResetParams {
	rv := objc.Send[MTRDishwasherAlarmClusterResetParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDishwasherAlarmClusterResetParams creates a new MTRDishwasherAlarmClusterResetParams instance.
func NewMTRDishwasherAlarmClusterResetParams() MTRDishwasherAlarmClusterResetParams {
	return getMTRDishwasherAlarmClusterResetParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherAlarmClusterResetParams/alarms
func (m_ MTRDishwasherAlarmClusterResetParams) Alarms() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("alarms"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherAlarmClusterResetParams/alarms
func (m_ MTRDishwasherAlarmClusterResetParams) SetAlarms(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAlarms:"), value)
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherAlarmClusterResetParams/serverSideProcessingTimeout
func (m_ MTRDishwasherAlarmClusterResetParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherAlarmClusterResetParams/serverSideProcessingTimeout
func (m_ MTRDishwasherAlarmClusterResetParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherAlarmClusterResetParams/timedInvokeTimeoutMs
func (m_ MTRDishwasherAlarmClusterResetParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherAlarmClusterResetParams/timedInvokeTimeoutMs
func (m_ MTRDishwasherAlarmClusterResetParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



