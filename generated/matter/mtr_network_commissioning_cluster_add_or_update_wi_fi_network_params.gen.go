// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRNetworkCommissioningClusterAddOrUpdateWiFiNetworkParams] class.
var (
	MTRNetworkCommissioningClusterAddOrUpdateWiFiNetworkParamsClass     _MTRNetworkCommissioningClusterAddOrUpdateWiFiNetworkParamsClass
	MTRNetworkCommissioningClusterAddOrUpdateWiFiNetworkParamsClassOnce sync.Once
)

func getMTRNetworkCommissioningClusterAddOrUpdateWiFiNetworkParamsClass() _MTRNetworkCommissioningClusterAddOrUpdateWiFiNetworkParamsClass {
	MTRNetworkCommissioningClusterAddOrUpdateWiFiNetworkParamsClassOnce.Do(func() {
		MTRNetworkCommissioningClusterAddOrUpdateWiFiNetworkParamsClass = _MTRNetworkCommissioningClusterAddOrUpdateWiFiNetworkParamsClass{objc.GetClass("MTRNetworkCommissioningClusterAddOrUpdateWiFiNetworkParams")}
	})
	return MTRNetworkCommissioningClusterAddOrUpdateWiFiNetworkParamsClass
}

type _MTRNetworkCommissioningClusterAddOrUpdateWiFiNetworkParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRNetworkCommissioningClusterAddOrUpdateWiFiNetworkParams] class.
type IMTRNetworkCommissioningClusterAddOrUpdateWiFiNetworkParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRNetworkCommissioningClusterAddOrUpdateWiFiNetworkParams
type MTRNetworkCommissioningClusterAddOrUpdateWiFiNetworkParams struct {
	objectivec.Object
}

// MTRNetworkCommissioningClusterAddOrUpdateWiFiNetworkParamsFrom constructs a [MTRNetworkCommissioningClusterAddOrUpdateWiFiNetworkParams] from an unsafe.Pointer.
func MTRNetworkCommissioningClusterAddOrUpdateWiFiNetworkParamsFrom(ptr unsafe.Pointer) MTRNetworkCommissioningClusterAddOrUpdateWiFiNetworkParams {
	return MTRNetworkCommissioningClusterAddOrUpdateWiFiNetworkParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRNetworkCommissioningClusterAddOrUpdateWiFiNetworkParamsClass) Alloc() MTRNetworkCommissioningClusterAddOrUpdateWiFiNetworkParams {
	rv := objc.Send[MTRNetworkCommissioningClusterAddOrUpdateWiFiNetworkParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRNetworkCommissioningClusterAddOrUpdateWiFiNetworkParamsClass) New() MTRNetworkCommissioningClusterAddOrUpdateWiFiNetworkParams {
	rv := objc.Send[MTRNetworkCommissioningClusterAddOrUpdateWiFiNetworkParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRNetworkCommissioningClusterAddOrUpdateWiFiNetworkParams) Init() MTRNetworkCommissioningClusterAddOrUpdateWiFiNetworkParams {
	rv := objc.Send[MTRNetworkCommissioningClusterAddOrUpdateWiFiNetworkParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRNetworkCommissioningClusterAddOrUpdateWiFiNetworkParams) Autorelease() MTRNetworkCommissioningClusterAddOrUpdateWiFiNetworkParams {
	rv := objc.Send[MTRNetworkCommissioningClusterAddOrUpdateWiFiNetworkParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRNetworkCommissioningClusterAddOrUpdateWiFiNetworkParams creates a new MTRNetworkCommissioningClusterAddOrUpdateWiFiNetworkParams instance.
func NewMTRNetworkCommissioningClusterAddOrUpdateWiFiNetworkParams() MTRNetworkCommissioningClusterAddOrUpdateWiFiNetworkParams {
	return getMTRNetworkCommissioningClusterAddOrUpdateWiFiNetworkParamsClass().New()
}




