// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParams] class.
var (
	MTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParamsClass     _MTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParamsClass
	MTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParamsClassOnce sync.Once
)

func getMTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParamsClass() _MTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParamsClass {
	MTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParamsClassOnce.Do(func() {
		MTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParamsClass = _MTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParamsClass{objc.GetClass("MTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParams")}
	})
	return MTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParamsClass
}

type _MTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParams] class.
type IMTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParams
type MTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParams struct {
	objectivec.Object
}

// MTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParamsFrom constructs a [MTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParams] from an unsafe.Pointer.
func MTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParamsFrom(ptr unsafe.Pointer) MTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParams {
	return MTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParamsClass) Alloc() MTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParams {
	rv := objc.Send[MTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParamsClass) New() MTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParams {
	rv := objc.Send[MTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParams) Init() MTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParams {
	rv := objc.Send[MTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParams) Autorelease() MTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParams {
	rv := objc.Send[MTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParams creates a new MTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParams instance.
func NewMTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParams() MTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParams {
	return getMTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParamsClass().New()
}




