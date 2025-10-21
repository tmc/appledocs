// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRNetworkCommissioningClusterConnectNetworkResponseParams] class.
var (
	MTRNetworkCommissioningClusterConnectNetworkResponseParamsClass     _MTRNetworkCommissioningClusterConnectNetworkResponseParamsClass
	MTRNetworkCommissioningClusterConnectNetworkResponseParamsClassOnce sync.Once
)

func getMTRNetworkCommissioningClusterConnectNetworkResponseParamsClass() _MTRNetworkCommissioningClusterConnectNetworkResponseParamsClass {
	MTRNetworkCommissioningClusterConnectNetworkResponseParamsClassOnce.Do(func() {
		MTRNetworkCommissioningClusterConnectNetworkResponseParamsClass = _MTRNetworkCommissioningClusterConnectNetworkResponseParamsClass{objc.GetClass("MTRNetworkCommissioningClusterConnectNetworkResponseParams")}
	})
	return MTRNetworkCommissioningClusterConnectNetworkResponseParamsClass
}

type _MTRNetworkCommissioningClusterConnectNetworkResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRNetworkCommissioningClusterConnectNetworkResponseParams] class.
type IMTRNetworkCommissioningClusterConnectNetworkResponseParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRNetworkCommissioningClusterConnectNetworkResponseParams
type MTRNetworkCommissioningClusterConnectNetworkResponseParams struct {
	objectivec.Object
}

// MTRNetworkCommissioningClusterConnectNetworkResponseParamsFrom constructs a [MTRNetworkCommissioningClusterConnectNetworkResponseParams] from an unsafe.Pointer.
func MTRNetworkCommissioningClusterConnectNetworkResponseParamsFrom(ptr unsafe.Pointer) MTRNetworkCommissioningClusterConnectNetworkResponseParams {
	return MTRNetworkCommissioningClusterConnectNetworkResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRNetworkCommissioningClusterConnectNetworkResponseParamsClass) Alloc() MTRNetworkCommissioningClusterConnectNetworkResponseParams {
	rv := objc.Send[MTRNetworkCommissioningClusterConnectNetworkResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRNetworkCommissioningClusterConnectNetworkResponseParamsClass) New() MTRNetworkCommissioningClusterConnectNetworkResponseParams {
	rv := objc.Send[MTRNetworkCommissioningClusterConnectNetworkResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRNetworkCommissioningClusterConnectNetworkResponseParams) Init() MTRNetworkCommissioningClusterConnectNetworkResponseParams {
	rv := objc.Send[MTRNetworkCommissioningClusterConnectNetworkResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRNetworkCommissioningClusterConnectNetworkResponseParams) Autorelease() MTRNetworkCommissioningClusterConnectNetworkResponseParams {
	rv := objc.Send[MTRNetworkCommissioningClusterConnectNetworkResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRNetworkCommissioningClusterConnectNetworkResponseParams creates a new MTRNetworkCommissioningClusterConnectNetworkResponseParams instance.
func NewMTRNetworkCommissioningClusterConnectNetworkResponseParams() MTRNetworkCommissioningClusterConnectNetworkResponseParams {
	return getMTRNetworkCommissioningClusterConnectNetworkResponseParamsClass().New()
}




