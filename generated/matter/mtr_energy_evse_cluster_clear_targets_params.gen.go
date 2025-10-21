// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTREnergyEVSEClusterClearTargetsParams] class.
var (
	MTREnergyEVSEClusterClearTargetsParamsClass     _MTREnergyEVSEClusterClearTargetsParamsClass
	MTREnergyEVSEClusterClearTargetsParamsClassOnce sync.Once
)

func getMTREnergyEVSEClusterClearTargetsParamsClass() _MTREnergyEVSEClusterClearTargetsParamsClass {
	MTREnergyEVSEClusterClearTargetsParamsClassOnce.Do(func() {
		MTREnergyEVSEClusterClearTargetsParamsClass = _MTREnergyEVSEClusterClearTargetsParamsClass{objc.GetClass("MTREnergyEVSEClusterClearTargetsParams")}
	})
	return MTREnergyEVSEClusterClearTargetsParamsClass
}

type _MTREnergyEVSEClusterClearTargetsParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTREnergyEVSEClusterClearTargetsParams] class.
type IMTREnergyEVSEClusterClearTargetsParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterClearTargetsParams
type MTREnergyEVSEClusterClearTargetsParams struct {
	objectivec.Object
}

// MTREnergyEVSEClusterClearTargetsParamsFrom constructs a [MTREnergyEVSEClusterClearTargetsParams] from an unsafe.Pointer.
func MTREnergyEVSEClusterClearTargetsParamsFrom(ptr unsafe.Pointer) MTREnergyEVSEClusterClearTargetsParams {
	return MTREnergyEVSEClusterClearTargetsParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTREnergyEVSEClusterClearTargetsParamsClass) Alloc() MTREnergyEVSEClusterClearTargetsParams {
	rv := objc.Send[MTREnergyEVSEClusterClearTargetsParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTREnergyEVSEClusterClearTargetsParamsClass) New() MTREnergyEVSEClusterClearTargetsParams {
	rv := objc.Send[MTREnergyEVSEClusterClearTargetsParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTREnergyEVSEClusterClearTargetsParams) Init() MTREnergyEVSEClusterClearTargetsParams {
	rv := objc.Send[MTREnergyEVSEClusterClearTargetsParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTREnergyEVSEClusterClearTargetsParams) Autorelease() MTREnergyEVSEClusterClearTargetsParams {
	rv := objc.Send[MTREnergyEVSEClusterClearTargetsParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTREnergyEVSEClusterClearTargetsParams creates a new MTREnergyEVSEClusterClearTargetsParams instance.
func NewMTREnergyEVSEClusterClearTargetsParams() MTREnergyEVSEClusterClearTargetsParams {
	return getMTREnergyEVSEClusterClearTargetsParamsClass().New()
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterClearTargetsParams/serverSideProcessingTimeout
func (m_ MTREnergyEVSEClusterClearTargetsParams) ServerSideProcessingTimeout() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
// Controls how much time, in seconds, we will allow for the server to process the command.

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterClearTargetsParams/serverSideProcessingTimeout
func (m_ MTREnergyEVSEClusterClearTargetsParams) SetServerSideProcessingTimeout(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}
// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterClearTargetsParams/timedInvokeTimeoutMs
func (m_ MTREnergyEVSEClusterClearTargetsParams) TimedInvokeTimeoutMs() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
// Controls whether the command is a timed command (using Timed Invoke).

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterClearTargetsParams/timedInvokeTimeoutMs
func (m_ MTREnergyEVSEClusterClearTargetsParams) SetTimedInvokeTimeoutMs(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}


