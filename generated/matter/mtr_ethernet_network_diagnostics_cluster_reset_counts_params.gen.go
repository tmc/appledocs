// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTREthernetNetworkDiagnosticsClusterResetCountsParams] class.
var (
	MTREthernetNetworkDiagnosticsClusterResetCountsParamsClass     _MTREthernetNetworkDiagnosticsClusterResetCountsParamsClass
	MTREthernetNetworkDiagnosticsClusterResetCountsParamsClassOnce sync.Once
)

func getMTREthernetNetworkDiagnosticsClusterResetCountsParamsClass() _MTREthernetNetworkDiagnosticsClusterResetCountsParamsClass {
	MTREthernetNetworkDiagnosticsClusterResetCountsParamsClassOnce.Do(func() {
		MTREthernetNetworkDiagnosticsClusterResetCountsParamsClass = _MTREthernetNetworkDiagnosticsClusterResetCountsParamsClass{objc.GetClass("MTREthernetNetworkDiagnosticsClusterResetCountsParams")}
	})
	return MTREthernetNetworkDiagnosticsClusterResetCountsParamsClass
}

type _MTREthernetNetworkDiagnosticsClusterResetCountsParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTREthernetNetworkDiagnosticsClusterResetCountsParams] class.
type IMTREthernetNetworkDiagnosticsClusterResetCountsParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREthernetNetworkDiagnosticsClusterResetCountsParams
type MTREthernetNetworkDiagnosticsClusterResetCountsParams struct {
	objectivec.Object
}

// MTREthernetNetworkDiagnosticsClusterResetCountsParamsFrom constructs a [MTREthernetNetworkDiagnosticsClusterResetCountsParams] from an unsafe.Pointer.
func MTREthernetNetworkDiagnosticsClusterResetCountsParamsFrom(ptr unsafe.Pointer) MTREthernetNetworkDiagnosticsClusterResetCountsParams {
	return MTREthernetNetworkDiagnosticsClusterResetCountsParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTREthernetNetworkDiagnosticsClusterResetCountsParamsClass) Alloc() MTREthernetNetworkDiagnosticsClusterResetCountsParams {
	rv := objc.Send[MTREthernetNetworkDiagnosticsClusterResetCountsParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTREthernetNetworkDiagnosticsClusterResetCountsParamsClass) New() MTREthernetNetworkDiagnosticsClusterResetCountsParams {
	rv := objc.Send[MTREthernetNetworkDiagnosticsClusterResetCountsParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTREthernetNetworkDiagnosticsClusterResetCountsParams) Init() MTREthernetNetworkDiagnosticsClusterResetCountsParams {
	rv := objc.Send[MTREthernetNetworkDiagnosticsClusterResetCountsParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTREthernetNetworkDiagnosticsClusterResetCountsParams) Autorelease() MTREthernetNetworkDiagnosticsClusterResetCountsParams {
	rv := objc.Send[MTREthernetNetworkDiagnosticsClusterResetCountsParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTREthernetNetworkDiagnosticsClusterResetCountsParams creates a new MTREthernetNetworkDiagnosticsClusterResetCountsParams instance.
func NewMTREthernetNetworkDiagnosticsClusterResetCountsParams() MTREthernetNetworkDiagnosticsClusterResetCountsParams {
	return getMTREthernetNetworkDiagnosticsClusterResetCountsParamsClass().New()
}




