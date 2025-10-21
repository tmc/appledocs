// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRNetworkCommissioningClusterConnectNetworkParams] class.
var (
	MTRNetworkCommissioningClusterConnectNetworkParamsClass     _MTRNetworkCommissioningClusterConnectNetworkParamsClass
	MTRNetworkCommissioningClusterConnectNetworkParamsClassOnce sync.Once
)

func getMTRNetworkCommissioningClusterConnectNetworkParamsClass() _MTRNetworkCommissioningClusterConnectNetworkParamsClass {
	MTRNetworkCommissioningClusterConnectNetworkParamsClassOnce.Do(func() {
		MTRNetworkCommissioningClusterConnectNetworkParamsClass = _MTRNetworkCommissioningClusterConnectNetworkParamsClass{objc.GetClass("MTRNetworkCommissioningClusterConnectNetworkParams")}
	})
	return MTRNetworkCommissioningClusterConnectNetworkParamsClass
}

type _MTRNetworkCommissioningClusterConnectNetworkParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRNetworkCommissioningClusterConnectNetworkParams] class.
type IMTRNetworkCommissioningClusterConnectNetworkParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRNetworkCommissioningClusterConnectNetworkParams
type MTRNetworkCommissioningClusterConnectNetworkParams struct {
	objectivec.Object
}

// MTRNetworkCommissioningClusterConnectNetworkParamsFrom constructs a [MTRNetworkCommissioningClusterConnectNetworkParams] from an unsafe.Pointer.
func MTRNetworkCommissioningClusterConnectNetworkParamsFrom(ptr unsafe.Pointer) MTRNetworkCommissioningClusterConnectNetworkParams {
	return MTRNetworkCommissioningClusterConnectNetworkParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRNetworkCommissioningClusterConnectNetworkParamsClass) Alloc() MTRNetworkCommissioningClusterConnectNetworkParams {
	rv := objc.Send[MTRNetworkCommissioningClusterConnectNetworkParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRNetworkCommissioningClusterConnectNetworkParamsClass) New() MTRNetworkCommissioningClusterConnectNetworkParams {
	rv := objc.Send[MTRNetworkCommissioningClusterConnectNetworkParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRNetworkCommissioningClusterConnectNetworkParams) Init() MTRNetworkCommissioningClusterConnectNetworkParams {
	rv := objc.Send[MTRNetworkCommissioningClusterConnectNetworkParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRNetworkCommissioningClusterConnectNetworkParams) Autorelease() MTRNetworkCommissioningClusterConnectNetworkParams {
	rv := objc.Send[MTRNetworkCommissioningClusterConnectNetworkParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRNetworkCommissioningClusterConnectNetworkParams creates a new MTRNetworkCommissioningClusterConnectNetworkParams instance.
func NewMTRNetworkCommissioningClusterConnectNetworkParams() MTRNetworkCommissioningClusterConnectNetworkParams {
	return getMTRNetworkCommissioningClusterConnectNetworkParamsClass().New()
}




