// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRGeneralDiagnosticsClusterTimeSnapshotParams] class.
var (
	MTRGeneralDiagnosticsClusterTimeSnapshotParamsClass     _MTRGeneralDiagnosticsClusterTimeSnapshotParamsClass
	MTRGeneralDiagnosticsClusterTimeSnapshotParamsClassOnce sync.Once
)

func getMTRGeneralDiagnosticsClusterTimeSnapshotParamsClass() _MTRGeneralDiagnosticsClusterTimeSnapshotParamsClass {
	MTRGeneralDiagnosticsClusterTimeSnapshotParamsClassOnce.Do(func() {
		MTRGeneralDiagnosticsClusterTimeSnapshotParamsClass = _MTRGeneralDiagnosticsClusterTimeSnapshotParamsClass{objc.GetClass("MTRGeneralDiagnosticsClusterTimeSnapshotParams")}
	})
	return MTRGeneralDiagnosticsClusterTimeSnapshotParamsClass
}

type _MTRGeneralDiagnosticsClusterTimeSnapshotParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRGeneralDiagnosticsClusterTimeSnapshotParams] class.
type IMTRGeneralDiagnosticsClusterTimeSnapshotParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGeneralDiagnosticsClusterTimeSnapshotParams
type MTRGeneralDiagnosticsClusterTimeSnapshotParams struct {
	objectivec.Object
}

// MTRGeneralDiagnosticsClusterTimeSnapshotParamsFrom constructs a [MTRGeneralDiagnosticsClusterTimeSnapshotParams] from an unsafe.Pointer.
func MTRGeneralDiagnosticsClusterTimeSnapshotParamsFrom(ptr unsafe.Pointer) MTRGeneralDiagnosticsClusterTimeSnapshotParams {
	return MTRGeneralDiagnosticsClusterTimeSnapshotParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRGeneralDiagnosticsClusterTimeSnapshotParamsClass) Alloc() MTRGeneralDiagnosticsClusterTimeSnapshotParams {
	rv := objc.Send[MTRGeneralDiagnosticsClusterTimeSnapshotParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRGeneralDiagnosticsClusterTimeSnapshotParamsClass) New() MTRGeneralDiagnosticsClusterTimeSnapshotParams {
	rv := objc.Send[MTRGeneralDiagnosticsClusterTimeSnapshotParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRGeneralDiagnosticsClusterTimeSnapshotParams) Init() MTRGeneralDiagnosticsClusterTimeSnapshotParams {
	rv := objc.Send[MTRGeneralDiagnosticsClusterTimeSnapshotParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRGeneralDiagnosticsClusterTimeSnapshotParams) Autorelease() MTRGeneralDiagnosticsClusterTimeSnapshotParams {
	rv := objc.Send[MTRGeneralDiagnosticsClusterTimeSnapshotParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRGeneralDiagnosticsClusterTimeSnapshotParams creates a new MTRGeneralDiagnosticsClusterTimeSnapshotParams instance.
func NewMTRGeneralDiagnosticsClusterTimeSnapshotParams() MTRGeneralDiagnosticsClusterTimeSnapshotParams {
	return getMTRGeneralDiagnosticsClusterTimeSnapshotParamsClass().New()
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGeneralDiagnosticsClusterTimeSnapshotParams/serverSideProcessingTimeout
func (m_ MTRGeneralDiagnosticsClusterTimeSnapshotParams) ServerSideProcessingTimeout() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
// Controls how much time, in seconds, we will allow for the server to process the command.

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGeneralDiagnosticsClusterTimeSnapshotParams/serverSideProcessingTimeout
func (m_ MTRGeneralDiagnosticsClusterTimeSnapshotParams) SetServerSideProcessingTimeout(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGeneralDiagnosticsClusterTimeSnapshotParams/timedInvokeTimeoutMs
func (m_ MTRGeneralDiagnosticsClusterTimeSnapshotParams) TimedInvokeTimeoutMs() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
// Controls whether the command is a timed command (using Timed Invoke).

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGeneralDiagnosticsClusterTimeSnapshotParams/timedInvokeTimeoutMs
func (m_ MTRGeneralDiagnosticsClusterTimeSnapshotParams) SetTimedInvokeTimeoutMs(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



