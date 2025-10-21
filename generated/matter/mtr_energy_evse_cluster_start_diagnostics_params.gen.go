// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTREnergyEVSEClusterStartDiagnosticsParams] class.
var (
	MTREnergyEVSEClusterStartDiagnosticsParamsClass     _MTREnergyEVSEClusterStartDiagnosticsParamsClass
	MTREnergyEVSEClusterStartDiagnosticsParamsClassOnce sync.Once
)

func getMTREnergyEVSEClusterStartDiagnosticsParamsClass() _MTREnergyEVSEClusterStartDiagnosticsParamsClass {
	MTREnergyEVSEClusterStartDiagnosticsParamsClassOnce.Do(func() {
		MTREnergyEVSEClusterStartDiagnosticsParamsClass = _MTREnergyEVSEClusterStartDiagnosticsParamsClass{objc.GetClass("MTREnergyEVSEClusterStartDiagnosticsParams")}
	})
	return MTREnergyEVSEClusterStartDiagnosticsParamsClass
}

type _MTREnergyEVSEClusterStartDiagnosticsParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTREnergyEVSEClusterStartDiagnosticsParams] class.
type IMTREnergyEVSEClusterStartDiagnosticsParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterStartDiagnosticsParams
type MTREnergyEVSEClusterStartDiagnosticsParams struct {
	objectivec.Object
}

// MTREnergyEVSEClusterStartDiagnosticsParamsFrom constructs a [MTREnergyEVSEClusterStartDiagnosticsParams] from an unsafe.Pointer.
func MTREnergyEVSEClusterStartDiagnosticsParamsFrom(ptr unsafe.Pointer) MTREnergyEVSEClusterStartDiagnosticsParams {
	return MTREnergyEVSEClusterStartDiagnosticsParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTREnergyEVSEClusterStartDiagnosticsParamsClass) Alloc() MTREnergyEVSEClusterStartDiagnosticsParams {
	rv := objc.Send[MTREnergyEVSEClusterStartDiagnosticsParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTREnergyEVSEClusterStartDiagnosticsParamsClass) New() MTREnergyEVSEClusterStartDiagnosticsParams {
	rv := objc.Send[MTREnergyEVSEClusterStartDiagnosticsParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTREnergyEVSEClusterStartDiagnosticsParams) Init() MTREnergyEVSEClusterStartDiagnosticsParams {
	rv := objc.Send[MTREnergyEVSEClusterStartDiagnosticsParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTREnergyEVSEClusterStartDiagnosticsParams) Autorelease() MTREnergyEVSEClusterStartDiagnosticsParams {
	rv := objc.Send[MTREnergyEVSEClusterStartDiagnosticsParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTREnergyEVSEClusterStartDiagnosticsParams creates a new MTREnergyEVSEClusterStartDiagnosticsParams instance.
func NewMTREnergyEVSEClusterStartDiagnosticsParams() MTREnergyEVSEClusterStartDiagnosticsParams {
	return getMTREnergyEVSEClusterStartDiagnosticsParamsClass().New()
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterStartDiagnosticsParams/serverSideProcessingTimeout
func (m_ MTREnergyEVSEClusterStartDiagnosticsParams) ServerSideProcessingTimeout() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
// Controls how much time, in seconds, we will allow for the server to process the command.

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterStartDiagnosticsParams/serverSideProcessingTimeout
func (m_ MTREnergyEVSEClusterStartDiagnosticsParams) SetServerSideProcessingTimeout(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterStartDiagnosticsParams/timedInvokeTimeoutMs
func (m_ MTREnergyEVSEClusterStartDiagnosticsParams) TimedInvokeTimeoutMs() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
// Controls whether the command is a timed command (using Timed Invoke).

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterStartDiagnosticsParams/timedInvokeTimeoutMs
func (m_ MTREnergyEVSEClusterStartDiagnosticsParams) SetTimedInvokeTimeoutMs(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



