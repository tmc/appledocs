// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTREnergyEVSEClusterGetTargetsParams] class.
var (
	MTREnergyEVSEClusterGetTargetsParamsClass     _MTREnergyEVSEClusterGetTargetsParamsClass
	MTREnergyEVSEClusterGetTargetsParamsClassOnce sync.Once
)

func getMTREnergyEVSEClusterGetTargetsParamsClass() _MTREnergyEVSEClusterGetTargetsParamsClass {
	MTREnergyEVSEClusterGetTargetsParamsClassOnce.Do(func() {
		MTREnergyEVSEClusterGetTargetsParamsClass = _MTREnergyEVSEClusterGetTargetsParamsClass{objc.GetClass("MTREnergyEVSEClusterGetTargetsParams")}
	})
	return MTREnergyEVSEClusterGetTargetsParamsClass
}

type _MTREnergyEVSEClusterGetTargetsParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTREnergyEVSEClusterGetTargetsParams] class.
type IMTREnergyEVSEClusterGetTargetsParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterGetTargetsParams
type MTREnergyEVSEClusterGetTargetsParams struct {
	objectivec.Object
}

// MTREnergyEVSEClusterGetTargetsParamsFrom constructs a [MTREnergyEVSEClusterGetTargetsParams] from an unsafe.Pointer.
func MTREnergyEVSEClusterGetTargetsParamsFrom(ptr unsafe.Pointer) MTREnergyEVSEClusterGetTargetsParams {
	return MTREnergyEVSEClusterGetTargetsParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTREnergyEVSEClusterGetTargetsParamsClass) Alloc() MTREnergyEVSEClusterGetTargetsParams {
	rv := objc.Send[MTREnergyEVSEClusterGetTargetsParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTREnergyEVSEClusterGetTargetsParamsClass) New() MTREnergyEVSEClusterGetTargetsParams {
	rv := objc.Send[MTREnergyEVSEClusterGetTargetsParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTREnergyEVSEClusterGetTargetsParams) Init() MTREnergyEVSEClusterGetTargetsParams {
	rv := objc.Send[MTREnergyEVSEClusterGetTargetsParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTREnergyEVSEClusterGetTargetsParams) Autorelease() MTREnergyEVSEClusterGetTargetsParams {
	rv := objc.Send[MTREnergyEVSEClusterGetTargetsParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTREnergyEVSEClusterGetTargetsParams creates a new MTREnergyEVSEClusterGetTargetsParams instance.
func NewMTREnergyEVSEClusterGetTargetsParams() MTREnergyEVSEClusterGetTargetsParams {
	return getMTREnergyEVSEClusterGetTargetsParamsClass().New()
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterGetTargetsParams/serverSideProcessingTimeout
func (m_ MTREnergyEVSEClusterGetTargetsParams) ServerSideProcessingTimeout() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
// Controls how much time, in seconds, we will allow for the server to process the command.

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterGetTargetsParams/serverSideProcessingTimeout
func (m_ MTREnergyEVSEClusterGetTargetsParams) SetServerSideProcessingTimeout(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}
// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterGetTargetsParams/timedInvokeTimeoutMs
func (m_ MTREnergyEVSEClusterGetTargetsParams) TimedInvokeTimeoutMs() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
// Controls whether the command is a timed command (using Timed Invoke).

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterGetTargetsParams/timedInvokeTimeoutMs
func (m_ MTREnergyEVSEClusterGetTargetsParams) SetTimedInvokeTimeoutMs(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}


