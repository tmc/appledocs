// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRThreadNetworkDiagnosticsClusterResetCountsParams] class.
var (
	MTRThreadNetworkDiagnosticsClusterResetCountsParamsClass     _MTRThreadNetworkDiagnosticsClusterResetCountsParamsClass
	MTRThreadNetworkDiagnosticsClusterResetCountsParamsClassOnce sync.Once
)

func getMTRThreadNetworkDiagnosticsClusterResetCountsParamsClass() _MTRThreadNetworkDiagnosticsClusterResetCountsParamsClass {
	MTRThreadNetworkDiagnosticsClusterResetCountsParamsClassOnce.Do(func() {
		MTRThreadNetworkDiagnosticsClusterResetCountsParamsClass = _MTRThreadNetworkDiagnosticsClusterResetCountsParamsClass{objc.GetClass("MTRThreadNetworkDiagnosticsClusterResetCountsParams")}
	})
	return MTRThreadNetworkDiagnosticsClusterResetCountsParamsClass
}

type _MTRThreadNetworkDiagnosticsClusterResetCountsParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRThreadNetworkDiagnosticsClusterResetCountsParams] class.
type IMTRThreadNetworkDiagnosticsClusterResetCountsParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterResetCountsParams
type MTRThreadNetworkDiagnosticsClusterResetCountsParams struct {
	objectivec.Object
}

// MTRThreadNetworkDiagnosticsClusterResetCountsParamsFrom constructs a [MTRThreadNetworkDiagnosticsClusterResetCountsParams] from an unsafe.Pointer.
func MTRThreadNetworkDiagnosticsClusterResetCountsParamsFrom(ptr unsafe.Pointer) MTRThreadNetworkDiagnosticsClusterResetCountsParams {
	return MTRThreadNetworkDiagnosticsClusterResetCountsParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRThreadNetworkDiagnosticsClusterResetCountsParamsClass) Alloc() MTRThreadNetworkDiagnosticsClusterResetCountsParams {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterResetCountsParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRThreadNetworkDiagnosticsClusterResetCountsParamsClass) New() MTRThreadNetworkDiagnosticsClusterResetCountsParams {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterResetCountsParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRThreadNetworkDiagnosticsClusterResetCountsParams) Init() MTRThreadNetworkDiagnosticsClusterResetCountsParams {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterResetCountsParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRThreadNetworkDiagnosticsClusterResetCountsParams) Autorelease() MTRThreadNetworkDiagnosticsClusterResetCountsParams {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterResetCountsParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRThreadNetworkDiagnosticsClusterResetCountsParams creates a new MTRThreadNetworkDiagnosticsClusterResetCountsParams instance.
func NewMTRThreadNetworkDiagnosticsClusterResetCountsParams() MTRThreadNetworkDiagnosticsClusterResetCountsParams {
	return getMTRThreadNetworkDiagnosticsClusterResetCountsParamsClass().New()
}




