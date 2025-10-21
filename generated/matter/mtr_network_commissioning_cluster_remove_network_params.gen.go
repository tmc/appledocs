// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRNetworkCommissioningClusterRemoveNetworkParams] class.
var (
	MTRNetworkCommissioningClusterRemoveNetworkParamsClass     _MTRNetworkCommissioningClusterRemoveNetworkParamsClass
	MTRNetworkCommissioningClusterRemoveNetworkParamsClassOnce sync.Once
)

func getMTRNetworkCommissioningClusterRemoveNetworkParamsClass() _MTRNetworkCommissioningClusterRemoveNetworkParamsClass {
	MTRNetworkCommissioningClusterRemoveNetworkParamsClassOnce.Do(func() {
		MTRNetworkCommissioningClusterRemoveNetworkParamsClass = _MTRNetworkCommissioningClusterRemoveNetworkParamsClass{objc.GetClass("MTRNetworkCommissioningClusterRemoveNetworkParams")}
	})
	return MTRNetworkCommissioningClusterRemoveNetworkParamsClass
}

type _MTRNetworkCommissioningClusterRemoveNetworkParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRNetworkCommissioningClusterRemoveNetworkParams] class.
type IMTRNetworkCommissioningClusterRemoveNetworkParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRNetworkCommissioningClusterRemoveNetworkParams
type MTRNetworkCommissioningClusterRemoveNetworkParams struct {
	objectivec.Object
}

// MTRNetworkCommissioningClusterRemoveNetworkParamsFrom constructs a [MTRNetworkCommissioningClusterRemoveNetworkParams] from an unsafe.Pointer.
func MTRNetworkCommissioningClusterRemoveNetworkParamsFrom(ptr unsafe.Pointer) MTRNetworkCommissioningClusterRemoveNetworkParams {
	return MTRNetworkCommissioningClusterRemoveNetworkParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRNetworkCommissioningClusterRemoveNetworkParamsClass) Alloc() MTRNetworkCommissioningClusterRemoveNetworkParams {
	rv := objc.Send[MTRNetworkCommissioningClusterRemoveNetworkParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRNetworkCommissioningClusterRemoveNetworkParamsClass) New() MTRNetworkCommissioningClusterRemoveNetworkParams {
	rv := objc.Send[MTRNetworkCommissioningClusterRemoveNetworkParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRNetworkCommissioningClusterRemoveNetworkParams) Init() MTRNetworkCommissioningClusterRemoveNetworkParams {
	rv := objc.Send[MTRNetworkCommissioningClusterRemoveNetworkParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRNetworkCommissioningClusterRemoveNetworkParams) Autorelease() MTRNetworkCommissioningClusterRemoveNetworkParams {
	rv := objc.Send[MTRNetworkCommissioningClusterRemoveNetworkParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRNetworkCommissioningClusterRemoveNetworkParams creates a new MTRNetworkCommissioningClusterRemoveNetworkParams instance.
func NewMTRNetworkCommissioningClusterRemoveNetworkParams() MTRNetworkCommissioningClusterRemoveNetworkParams {
	return getMTRNetworkCommissioningClusterRemoveNetworkParamsClass().New()
}




