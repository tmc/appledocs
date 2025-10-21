// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRNetworkCommissioningClusterScanNetworksResponseParams] class.
var (
	MTRNetworkCommissioningClusterScanNetworksResponseParamsClass     _MTRNetworkCommissioningClusterScanNetworksResponseParamsClass
	MTRNetworkCommissioningClusterScanNetworksResponseParamsClassOnce sync.Once
)

func getMTRNetworkCommissioningClusterScanNetworksResponseParamsClass() _MTRNetworkCommissioningClusterScanNetworksResponseParamsClass {
	MTRNetworkCommissioningClusterScanNetworksResponseParamsClassOnce.Do(func() {
		MTRNetworkCommissioningClusterScanNetworksResponseParamsClass = _MTRNetworkCommissioningClusterScanNetworksResponseParamsClass{objc.GetClass("MTRNetworkCommissioningClusterScanNetworksResponseParams")}
	})
	return MTRNetworkCommissioningClusterScanNetworksResponseParamsClass
}

type _MTRNetworkCommissioningClusterScanNetworksResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRNetworkCommissioningClusterScanNetworksResponseParams] class.
type IMTRNetworkCommissioningClusterScanNetworksResponseParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRNetworkCommissioningClusterScanNetworksResponseParams
type MTRNetworkCommissioningClusterScanNetworksResponseParams struct {
	objectivec.Object
}

// MTRNetworkCommissioningClusterScanNetworksResponseParamsFrom constructs a [MTRNetworkCommissioningClusterScanNetworksResponseParams] from an unsafe.Pointer.
func MTRNetworkCommissioningClusterScanNetworksResponseParamsFrom(ptr unsafe.Pointer) MTRNetworkCommissioningClusterScanNetworksResponseParams {
	return MTRNetworkCommissioningClusterScanNetworksResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRNetworkCommissioningClusterScanNetworksResponseParamsClass) Alloc() MTRNetworkCommissioningClusterScanNetworksResponseParams {
	rv := objc.Send[MTRNetworkCommissioningClusterScanNetworksResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRNetworkCommissioningClusterScanNetworksResponseParamsClass) New() MTRNetworkCommissioningClusterScanNetworksResponseParams {
	rv := objc.Send[MTRNetworkCommissioningClusterScanNetworksResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRNetworkCommissioningClusterScanNetworksResponseParams) Init() MTRNetworkCommissioningClusterScanNetworksResponseParams {
	rv := objc.Send[MTRNetworkCommissioningClusterScanNetworksResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRNetworkCommissioningClusterScanNetworksResponseParams) Autorelease() MTRNetworkCommissioningClusterScanNetworksResponseParams {
	rv := objc.Send[MTRNetworkCommissioningClusterScanNetworksResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRNetworkCommissioningClusterScanNetworksResponseParams creates a new MTRNetworkCommissioningClusterScanNetworksResponseParams instance.
func NewMTRNetworkCommissioningClusterScanNetworksResponseParams() MTRNetworkCommissioningClusterScanNetworksResponseParams {
	return getMTRNetworkCommissioningClusterScanNetworksResponseParamsClass().New()
}




