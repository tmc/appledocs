// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
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
	// properties:
	Breadcrumb() objc.IObject /* cross-framework: NSNumber */
	SetBreadcrumb(value objc.IObject /* cross-framework: NSNumber */)
	NetworkID() objc.IObject /* cross-framework: Data */
	SetNetworkID(value objc.IObject /* cross-framework: Data */)
	NetworkIndex() objc.IObject /* cross-framework: NSNumber */
	SetNetworkIndex(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}

// [Full Topic]
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

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterreordernetworkparams/breadcrumb
func (m_ MTRNetworkCommissioningClusterReorderNetworkParams) Breadcrumb() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("breadcrumb"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterreordernetworkparams/breadcrumb
func (m_ MTRNetworkCommissioningClusterReorderNetworkParams) SetBreadcrumb(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBreadcrumb:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterreordernetworkparams/networkid
func (m_ MTRNetworkCommissioningClusterReorderNetworkParams) NetworkID() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("networkID"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterreordernetworkparams/networkid
func (m_ MTRNetworkCommissioningClusterReorderNetworkParams) SetNetworkID(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNetworkID:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterreordernetworkparams/networkindex
func (m_ MTRNetworkCommissioningClusterReorderNetworkParams) NetworkIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("networkIndex"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterreordernetworkparams/networkindex
func (m_ MTRNetworkCommissioningClusterReorderNetworkParams) SetNetworkIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNetworkIndex:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterreordernetworkparams/serversideprocessingtimeout
func (m_ MTRNetworkCommissioningClusterReorderNetworkParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterreordernetworkparams/serversideprocessingtimeout
func (m_ MTRNetworkCommissioningClusterReorderNetworkParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterreordernetworkparams/timedinvoketimeoutms
func (m_ MTRNetworkCommissioningClusterReorderNetworkParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterreordernetworkparams/timedinvoketimeoutms
func (m_ MTRNetworkCommissioningClusterReorderNetworkParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}
