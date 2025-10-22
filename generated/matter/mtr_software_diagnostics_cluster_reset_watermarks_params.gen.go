// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRSoftwareDiagnosticsClusterResetWatermarksParams] class.
var (
	MTRSoftwareDiagnosticsClusterResetWatermarksParamsClass     _MTRSoftwareDiagnosticsClusterResetWatermarksParamsClass
	MTRSoftwareDiagnosticsClusterResetWatermarksParamsClassOnce sync.Once
)

func getMTRSoftwareDiagnosticsClusterResetWatermarksParamsClass() _MTRSoftwareDiagnosticsClusterResetWatermarksParamsClass {
	MTRSoftwareDiagnosticsClusterResetWatermarksParamsClassOnce.Do(func() {
		MTRSoftwareDiagnosticsClusterResetWatermarksParamsClass = _MTRSoftwareDiagnosticsClusterResetWatermarksParamsClass{objc.GetClass("MTRSoftwareDiagnosticsClusterResetWatermarksParams")}
	})
	return MTRSoftwareDiagnosticsClusterResetWatermarksParamsClass
}

type _MTRSoftwareDiagnosticsClusterResetWatermarksParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRSoftwareDiagnosticsClusterResetWatermarksParams] class.
type IMTRSoftwareDiagnosticsClusterResetWatermarksParams interface {
	objectivec.IObject
	ServerSideProcessingTimeout() foundation.Number
	SetServerSideProcessingTimeout(value foundation.INumber)
	TimedInvokeTimeoutMs() foundation.Number
	SetTimedInvokeTimeoutMs(value foundation.INumber)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSoftwareDiagnosticsClusterResetWatermarksParams
type MTRSoftwareDiagnosticsClusterResetWatermarksParams struct {
	objectivec.Object
}

// MTRSoftwareDiagnosticsClusterResetWatermarksParamsFrom constructs a [MTRSoftwareDiagnosticsClusterResetWatermarksParams] from an unsafe.Pointer.
func MTRSoftwareDiagnosticsClusterResetWatermarksParamsFrom(ptr unsafe.Pointer) MTRSoftwareDiagnosticsClusterResetWatermarksParams {
	return MTRSoftwareDiagnosticsClusterResetWatermarksParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRSoftwareDiagnosticsClusterResetWatermarksParamsClass) Alloc() MTRSoftwareDiagnosticsClusterResetWatermarksParams {
	rv := objc.Send[MTRSoftwareDiagnosticsClusterResetWatermarksParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRSoftwareDiagnosticsClusterResetWatermarksParamsClass) New() MTRSoftwareDiagnosticsClusterResetWatermarksParams {
	rv := objc.Send[MTRSoftwareDiagnosticsClusterResetWatermarksParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRSoftwareDiagnosticsClusterResetWatermarksParams) Init() MTRSoftwareDiagnosticsClusterResetWatermarksParams {
	rv := objc.Send[MTRSoftwareDiagnosticsClusterResetWatermarksParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRSoftwareDiagnosticsClusterResetWatermarksParams) Autorelease() MTRSoftwareDiagnosticsClusterResetWatermarksParams {
	rv := objc.Send[MTRSoftwareDiagnosticsClusterResetWatermarksParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRSoftwareDiagnosticsClusterResetWatermarksParams creates a new MTRSoftwareDiagnosticsClusterResetWatermarksParams instance.
func NewMTRSoftwareDiagnosticsClusterResetWatermarksParams() MTRSoftwareDiagnosticsClusterResetWatermarksParams {
	return getMTRSoftwareDiagnosticsClusterResetWatermarksParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsoftwarediagnosticsclusterresetwatermarksparams/serversideprocessingtimeout
func (m_ MTRSoftwareDiagnosticsClusterResetWatermarksParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsoftwarediagnosticsclusterresetwatermarksparams/serversideprocessingtimeout
func (m_ MTRSoftwareDiagnosticsClusterResetWatermarksParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsoftwarediagnosticsclusterresetwatermarksparams/timedinvoketimeoutms
func (m_ MTRSoftwareDiagnosticsClusterResetWatermarksParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsoftwarediagnosticsclusterresetwatermarksparams/timedinvoketimeoutms
func (m_ MTRSoftwareDiagnosticsClusterResetWatermarksParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



