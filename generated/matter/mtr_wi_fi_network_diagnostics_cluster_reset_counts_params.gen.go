// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRWiFiNetworkDiagnosticsClusterResetCountsParams] class.
var (
	MTRWiFiNetworkDiagnosticsClusterResetCountsParamsClass     _MTRWiFiNetworkDiagnosticsClusterResetCountsParamsClass
	MTRWiFiNetworkDiagnosticsClusterResetCountsParamsClassOnce sync.Once
)

func getMTRWiFiNetworkDiagnosticsClusterResetCountsParamsClass() _MTRWiFiNetworkDiagnosticsClusterResetCountsParamsClass {
	MTRWiFiNetworkDiagnosticsClusterResetCountsParamsClassOnce.Do(func() {
		MTRWiFiNetworkDiagnosticsClusterResetCountsParamsClass = _MTRWiFiNetworkDiagnosticsClusterResetCountsParamsClass{objc.GetClass("MTRWiFiNetworkDiagnosticsClusterResetCountsParams")}
	})
	return MTRWiFiNetworkDiagnosticsClusterResetCountsParamsClass
}

type _MTRWiFiNetworkDiagnosticsClusterResetCountsParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRWiFiNetworkDiagnosticsClusterResetCountsParams] class.
type IMTRWiFiNetworkDiagnosticsClusterResetCountsParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWiFiNetworkDiagnosticsClusterResetCountsParams
type MTRWiFiNetworkDiagnosticsClusterResetCountsParams struct {
	objectivec.Object
}

// MTRWiFiNetworkDiagnosticsClusterResetCountsParamsFrom constructs a [MTRWiFiNetworkDiagnosticsClusterResetCountsParams] from an unsafe.Pointer.
func MTRWiFiNetworkDiagnosticsClusterResetCountsParamsFrom(ptr unsafe.Pointer) MTRWiFiNetworkDiagnosticsClusterResetCountsParams {
	return MTRWiFiNetworkDiagnosticsClusterResetCountsParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRWiFiNetworkDiagnosticsClusterResetCountsParamsClass) Alloc() MTRWiFiNetworkDiagnosticsClusterResetCountsParams {
	rv := objc.Send[MTRWiFiNetworkDiagnosticsClusterResetCountsParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRWiFiNetworkDiagnosticsClusterResetCountsParamsClass) New() MTRWiFiNetworkDiagnosticsClusterResetCountsParams {
	rv := objc.Send[MTRWiFiNetworkDiagnosticsClusterResetCountsParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRWiFiNetworkDiagnosticsClusterResetCountsParams) Init() MTRWiFiNetworkDiagnosticsClusterResetCountsParams {
	rv := objc.Send[MTRWiFiNetworkDiagnosticsClusterResetCountsParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRWiFiNetworkDiagnosticsClusterResetCountsParams) Autorelease() MTRWiFiNetworkDiagnosticsClusterResetCountsParams {
	rv := objc.Send[MTRWiFiNetworkDiagnosticsClusterResetCountsParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRWiFiNetworkDiagnosticsClusterResetCountsParams creates a new MTRWiFiNetworkDiagnosticsClusterResetCountsParams instance.
func NewMTRWiFiNetworkDiagnosticsClusterResetCountsParams() MTRWiFiNetworkDiagnosticsClusterResetCountsParams {
	return getMTRWiFiNetworkDiagnosticsClusterResetCountsParamsClass().New()
}




