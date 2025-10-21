// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
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


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterreordernetworkparams/serversideprocessingtimeout
func (m_ MTRNetworkCommissioningClusterReorderNetworkParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterreordernetworkparams/serversideprocessingtimeout
func (m_ MTRNetworkCommissioningClusterReorderNetworkParams) SetServerSideProcessingTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterreordernetworkparams/networkindex
func (m_ MTRNetworkCommissioningClusterReorderNetworkParams) NetworkIndex() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("networkIndex"))
	return rv
}


// SetNetworkIndex sets the value of the networkIndex property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterreordernetworkparams/networkindex
func (m_ MTRNetworkCommissioningClusterReorderNetworkParams) SetNetworkIndex(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNetworkIndex:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterreordernetworkparams/networkid
func (m_ MTRNetworkCommissioningClusterReorderNetworkParams) NetworkID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("networkID"))
	return rv
}


// SetNetworkID sets the value of the networkID property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterreordernetworkparams/networkid
func (m_ MTRNetworkCommissioningClusterReorderNetworkParams) SetNetworkID(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNetworkID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterreordernetworkparams/breadcrumb
func (m_ MTRNetworkCommissioningClusterReorderNetworkParams) Breadcrumb() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("breadcrumb"))
	return rv
}


// SetBreadcrumb sets the value of the breadcrumb property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterreordernetworkparams/breadcrumb
func (m_ MTRNetworkCommissioningClusterReorderNetworkParams) SetBreadcrumb(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBreadcrumb:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterreordernetworkparams/timedinvoketimeoutms
func (m_ MTRNetworkCommissioningClusterReorderNetworkParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterreordernetworkparams/timedinvoketimeoutms
func (m_ MTRNetworkCommissioningClusterReorderNetworkParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



