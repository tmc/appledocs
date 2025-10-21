// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRGeneralDiagnosticsClusterTimeSnapshotResponseParams] class.
var (
	MTRGeneralDiagnosticsClusterTimeSnapshotResponseParamsClass     _MTRGeneralDiagnosticsClusterTimeSnapshotResponseParamsClass
	MTRGeneralDiagnosticsClusterTimeSnapshotResponseParamsClassOnce sync.Once
)

func getMTRGeneralDiagnosticsClusterTimeSnapshotResponseParamsClass() _MTRGeneralDiagnosticsClusterTimeSnapshotResponseParamsClass {
	MTRGeneralDiagnosticsClusterTimeSnapshotResponseParamsClassOnce.Do(func() {
		MTRGeneralDiagnosticsClusterTimeSnapshotResponseParamsClass = _MTRGeneralDiagnosticsClusterTimeSnapshotResponseParamsClass{objc.GetClass("MTRGeneralDiagnosticsClusterTimeSnapshotResponseParams")}
	})
	return MTRGeneralDiagnosticsClusterTimeSnapshotResponseParamsClass
}

type _MTRGeneralDiagnosticsClusterTimeSnapshotResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRGeneralDiagnosticsClusterTimeSnapshotResponseParams] class.
type IMTRGeneralDiagnosticsClusterTimeSnapshotResponseParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGeneralDiagnosticsClusterTimeSnapshotResponseParams
type MTRGeneralDiagnosticsClusterTimeSnapshotResponseParams struct {
	objectivec.Object
}

// MTRGeneralDiagnosticsClusterTimeSnapshotResponseParamsFrom constructs a [MTRGeneralDiagnosticsClusterTimeSnapshotResponseParams] from an unsafe.Pointer.
func MTRGeneralDiagnosticsClusterTimeSnapshotResponseParamsFrom(ptr unsafe.Pointer) MTRGeneralDiagnosticsClusterTimeSnapshotResponseParams {
	return MTRGeneralDiagnosticsClusterTimeSnapshotResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRGeneralDiagnosticsClusterTimeSnapshotResponseParamsClass) Alloc() MTRGeneralDiagnosticsClusterTimeSnapshotResponseParams {
	rv := objc.Send[MTRGeneralDiagnosticsClusterTimeSnapshotResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRGeneralDiagnosticsClusterTimeSnapshotResponseParamsClass) New() MTRGeneralDiagnosticsClusterTimeSnapshotResponseParams {
	rv := objc.Send[MTRGeneralDiagnosticsClusterTimeSnapshotResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRGeneralDiagnosticsClusterTimeSnapshotResponseParams) Init() MTRGeneralDiagnosticsClusterTimeSnapshotResponseParams {
	rv := objc.Send[MTRGeneralDiagnosticsClusterTimeSnapshotResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRGeneralDiagnosticsClusterTimeSnapshotResponseParams) Autorelease() MTRGeneralDiagnosticsClusterTimeSnapshotResponseParams {
	rv := objc.Send[MTRGeneralDiagnosticsClusterTimeSnapshotResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRGeneralDiagnosticsClusterTimeSnapshotResponseParams creates a new MTRGeneralDiagnosticsClusterTimeSnapshotResponseParams instance.
func NewMTRGeneralDiagnosticsClusterTimeSnapshotResponseParams() MTRGeneralDiagnosticsClusterTimeSnapshotResponseParams {
	return getMTRGeneralDiagnosticsClusterTimeSnapshotResponseParamsClass().New()
}




// Initialize an MTRGeneralDiagnosticsClusterTimeSnapshotResponseParams with a response-value dictionary of the sort that MTRDeviceResponseHandler would receive.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGeneralDiagnosticsClusterTimeSnapshotResponseParams/init(responseValue:)
func NewMTRGeneralDiagnosticsClusterTimeSnapshotResponseParamsWithResponseValueError(responseValue unsafe.Pointer, error_ unsafe.Pointer) MTRGeneralDiagnosticsClusterTimeSnapshotResponseParams {
	instance := getMTRGeneralDiagnosticsClusterTimeSnapshotResponseParamsClass().Alloc()
	rv := objc.Send[MTRGeneralDiagnosticsClusterTimeSnapshotResponseParams](instance.ID, objc.Sel("initWithResponseValue:error:"), responseValue, error_)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGeneralDiagnosticsClusterTimeSnapshotResponseParams/posixTimeMs
func (m_ MTRGeneralDiagnosticsClusterTimeSnapshotResponseParams) PosixTimeMs() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("posixTimeMs"))
	return rv
}


// SetPosixTimeMs sets the value of the posixTimeMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGeneralDiagnosticsClusterTimeSnapshotResponseParams/posixTimeMs
func (m_ MTRGeneralDiagnosticsClusterTimeSnapshotResponseParams) SetPosixTimeMs(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPosixTimeMs:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGeneralDiagnosticsClusterTimeSnapshotResponseParams/systemTimeMs
func (m_ MTRGeneralDiagnosticsClusterTimeSnapshotResponseParams) SystemTimeMs() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("systemTimeMs"))
	return rv
}


// SetSystemTimeMs sets the value of the systemTimeMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGeneralDiagnosticsClusterTimeSnapshotResponseParams/systemTimeMs
func (m_ MTRGeneralDiagnosticsClusterTimeSnapshotResponseParams) SetSystemTimeMs(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSystemTimeMs:"), value)
}


