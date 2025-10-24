// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	// properties:
	Breadcrumb() objc.IObject /* cross-framework: NSNumber */
	SetBreadcrumb(value objc.IObject /* cross-framework: NSNumber */)
	NetworkID() objc.IObject /* cross-framework: Data */
	SetNetworkID(value objc.IObject /* cross-framework: Data */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterconnectnetworkparams/breadcrumb
func (m_ MTRNetworkCommissioningClusterConnectNetworkParams) Breadcrumb() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("breadcrumb"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterconnectnetworkparams/breadcrumb
func (m_ MTRNetworkCommissioningClusterConnectNetworkParams) SetBreadcrumb(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBreadcrumb:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterconnectnetworkparams/networkid
func (m_ MTRNetworkCommissioningClusterConnectNetworkParams) NetworkID() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("networkID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterconnectnetworkparams/networkid
func (m_ MTRNetworkCommissioningClusterConnectNetworkParams) SetNetworkID(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNetworkID:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterconnectnetworkparams/serversideprocessingtimeout
func (m_ MTRNetworkCommissioningClusterConnectNetworkParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterconnectnetworkparams/serversideprocessingtimeout
func (m_ MTRNetworkCommissioningClusterConnectNetworkParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterconnectnetworkparams/timedinvoketimeoutms
func (m_ MTRNetworkCommissioningClusterConnectNetworkParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterconnectnetworkparams/timedinvoketimeoutms
func (m_ MTRNetworkCommissioningClusterConnectNetworkParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



