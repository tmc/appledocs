// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTREnergyEVSEClusterSetTargetsParams] class.
var (
	MTREnergyEVSEClusterSetTargetsParamsClass     _MTREnergyEVSEClusterSetTargetsParamsClass
	MTREnergyEVSEClusterSetTargetsParamsClassOnce sync.Once
)

func getMTREnergyEVSEClusterSetTargetsParamsClass() _MTREnergyEVSEClusterSetTargetsParamsClass {
	MTREnergyEVSEClusterSetTargetsParamsClassOnce.Do(func() {
		MTREnergyEVSEClusterSetTargetsParamsClass = _MTREnergyEVSEClusterSetTargetsParamsClass{objc.GetClass("MTREnergyEVSEClusterSetTargetsParams")}
	})
	return MTREnergyEVSEClusterSetTargetsParamsClass
}

type _MTREnergyEVSEClusterSetTargetsParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTREnergyEVSEClusterSetTargetsParams] class.
type IMTREnergyEVSEClusterSetTargetsParams interface {
	objectivec.IObject
	// properties:
	ChargingTargetSchedules() objc.IObject /* cross-framework: NSArray */
	SetChargingTargetSchedules(value objc.IObject /* cross-framework: NSArray */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterSetTargetsParams
type MTREnergyEVSEClusterSetTargetsParams struct {
	objectivec.Object
}

// MTREnergyEVSEClusterSetTargetsParamsFrom constructs a [MTREnergyEVSEClusterSetTargetsParams] from an unsafe.Pointer.
func MTREnergyEVSEClusterSetTargetsParamsFrom(ptr unsafe.Pointer) MTREnergyEVSEClusterSetTargetsParams {
	return MTREnergyEVSEClusterSetTargetsParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTREnergyEVSEClusterSetTargetsParamsClass) Alloc() MTREnergyEVSEClusterSetTargetsParams {
	rv := objc.Send[MTREnergyEVSEClusterSetTargetsParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTREnergyEVSEClusterSetTargetsParamsClass) New() MTREnergyEVSEClusterSetTargetsParams {
	rv := objc.Send[MTREnergyEVSEClusterSetTargetsParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTREnergyEVSEClusterSetTargetsParams) Init() MTREnergyEVSEClusterSetTargetsParams {
	rv := objc.Send[MTREnergyEVSEClusterSetTargetsParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTREnergyEVSEClusterSetTargetsParams) Autorelease() MTREnergyEVSEClusterSetTargetsParams {
	rv := objc.Send[MTREnergyEVSEClusterSetTargetsParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTREnergyEVSEClusterSetTargetsParams creates a new MTREnergyEVSEClusterSetTargetsParams instance.
func NewMTREnergyEVSEClusterSetTargetsParams() MTREnergyEVSEClusterSetTargetsParams {
	return getMTREnergyEVSEClusterSetTargetsParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterSetTargetsParams/chargingTargetSchedules
func (m_ MTREnergyEVSEClusterSetTargetsParams) ChargingTargetSchedules() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("chargingTargetSchedules"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterSetTargetsParams/chargingTargetSchedules
func (m_ MTREnergyEVSEClusterSetTargetsParams) SetChargingTargetSchedules(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setChargingTargetSchedules:"), value)
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterSetTargetsParams/serverSideProcessingTimeout
func (m_ MTREnergyEVSEClusterSetTargetsParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterSetTargetsParams/serverSideProcessingTimeout
func (m_ MTREnergyEVSEClusterSetTargetsParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterSetTargetsParams/timedInvokeTimeoutMs
func (m_ MTREnergyEVSEClusterSetTargetsParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterSetTargetsParams/timedInvokeTimeoutMs
func (m_ MTREnergyEVSEClusterSetTargetsParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



