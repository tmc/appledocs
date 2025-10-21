// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRNetworkCommissioningClusterScanNetworksParams] class.
var (
	MTRNetworkCommissioningClusterScanNetworksParamsClass     _MTRNetworkCommissioningClusterScanNetworksParamsClass
	MTRNetworkCommissioningClusterScanNetworksParamsClassOnce sync.Once
)

func getMTRNetworkCommissioningClusterScanNetworksParamsClass() _MTRNetworkCommissioningClusterScanNetworksParamsClass {
	MTRNetworkCommissioningClusterScanNetworksParamsClassOnce.Do(func() {
		MTRNetworkCommissioningClusterScanNetworksParamsClass = _MTRNetworkCommissioningClusterScanNetworksParamsClass{objc.GetClass("MTRNetworkCommissioningClusterScanNetworksParams")}
	})
	return MTRNetworkCommissioningClusterScanNetworksParamsClass
}

type _MTRNetworkCommissioningClusterScanNetworksParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRNetworkCommissioningClusterScanNetworksParams] class.
type IMTRNetworkCommissioningClusterScanNetworksParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRNetworkCommissioningClusterScanNetworksParams
type MTRNetworkCommissioningClusterScanNetworksParams struct {
	objectivec.Object
}

// MTRNetworkCommissioningClusterScanNetworksParamsFrom constructs a [MTRNetworkCommissioningClusterScanNetworksParams] from an unsafe.Pointer.
func MTRNetworkCommissioningClusterScanNetworksParamsFrom(ptr unsafe.Pointer) MTRNetworkCommissioningClusterScanNetworksParams {
	return MTRNetworkCommissioningClusterScanNetworksParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRNetworkCommissioningClusterScanNetworksParamsClass) Alloc() MTRNetworkCommissioningClusterScanNetworksParams {
	rv := objc.Send[MTRNetworkCommissioningClusterScanNetworksParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRNetworkCommissioningClusterScanNetworksParamsClass) New() MTRNetworkCommissioningClusterScanNetworksParams {
	rv := objc.Send[MTRNetworkCommissioningClusterScanNetworksParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRNetworkCommissioningClusterScanNetworksParams) Init() MTRNetworkCommissioningClusterScanNetworksParams {
	rv := objc.Send[MTRNetworkCommissioningClusterScanNetworksParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRNetworkCommissioningClusterScanNetworksParams) Autorelease() MTRNetworkCommissioningClusterScanNetworksParams {
	rv := objc.Send[MTRNetworkCommissioningClusterScanNetworksParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRNetworkCommissioningClusterScanNetworksParams creates a new MTRNetworkCommissioningClusterScanNetworksParams instance.
func NewMTRNetworkCommissioningClusterScanNetworksParams() MTRNetworkCommissioningClusterScanNetworksParams {
	return getMTRNetworkCommissioningClusterScanNetworksParamsClass().New()
}




