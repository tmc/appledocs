// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRNetworkCommissioningClusterReorderNetworkParams] class.
var (
	MTRNetworkCommissioningClusterReorderNetworkParamsClass     _MTRNetworkCommissioningClusterReorderNetworkParamsClass
	MTRNetworkCommissioningClusterReorderNetworkParamsClassOnce sync.Once
)

func getMTRNetworkCommissioningClusterReorderNetworkParamsClass() _MTRNetworkCommissioningClusterReorderNetworkParamsClass {
	MTRNetworkCommissioningClusterReorderNetworkParamsClassOnce.Do(func() {
		MTRNetworkCommissioningClusterReorderNetworkParamsClass = _MTRNetworkCommissioningClusterReorderNetworkParamsClass{objc.GetClass("MTRNetworkCommissioningClusterReorderNetworkParams")}
	})
	return MTRNetworkCommissioningClusterReorderNetworkParamsClass
}

type _MTRNetworkCommissioningClusterReorderNetworkParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRNetworkCommissioningClusterReorderNetworkParams] class.
type IMTRNetworkCommissioningClusterReorderNetworkParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRNetworkCommissioningClusterReorderNetworkParams
type MTRNetworkCommissioningClusterReorderNetworkParams struct {
	objectivec.Object
}

// MTRNetworkCommissioningClusterReorderNetworkParamsFrom constructs a [MTRNetworkCommissioningClusterReorderNetworkParams] from an unsafe.Pointer.
func MTRNetworkCommissioningClusterReorderNetworkParamsFrom(ptr unsafe.Pointer) MTRNetworkCommissioningClusterReorderNetworkParams {
	return MTRNetworkCommissioningClusterReorderNetworkParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRNetworkCommissioningClusterReorderNetworkParamsClass) Alloc() MTRNetworkCommissioningClusterReorderNetworkParams {
	rv := objc.Send[MTRNetworkCommissioningClusterReorderNetworkParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRNetworkCommissioningClusterReorderNetworkParamsClass) New() MTRNetworkCommissioningClusterReorderNetworkParams {
	rv := objc.Send[MTRNetworkCommissioningClusterReorderNetworkParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRNetworkCommissioningClusterReorderNetworkParams) Init() MTRNetworkCommissioningClusterReorderNetworkParams {
	rv := objc.Send[MTRNetworkCommissioningClusterReorderNetworkParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRNetworkCommissioningClusterReorderNetworkParams) Autorelease() MTRNetworkCommissioningClusterReorderNetworkParams {
	rv := objc.Send[MTRNetworkCommissioningClusterReorderNetworkParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRNetworkCommissioningClusterReorderNetworkParams creates a new MTRNetworkCommissioningClusterReorderNetworkParams instance.
func NewMTRNetworkCommissioningClusterReorderNetworkParams() MTRNetworkCommissioningClusterReorderNetworkParams {
	return getMTRNetworkCommissioningClusterReorderNetworkParamsClass().New()
}




