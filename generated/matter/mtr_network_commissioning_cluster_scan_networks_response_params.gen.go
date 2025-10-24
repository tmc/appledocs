// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRNetworkCommissioningClusterScanNetworksResponseParams] class.
var (
	MTRNetworkCommissioningClusterScanNetworksResponseParamsClass     _MTRNetworkCommissioningClusterScanNetworksResponseParamsClass
	MTRNetworkCommissioningClusterScanNetworksResponseParamsClassOnce sync.Once
)

func getMTRNetworkCommissioningClusterScanNetworksResponseParamsClass() _MTRNetworkCommissioningClusterScanNetworksResponseParamsClass {
	MTRNetworkCommissioningClusterScanNetworksResponseParamsClassOnce.Do(func() {
		MTRNetworkCommissioningClusterScanNetworksResponseParamsClass = _MTRNetworkCommissioningClusterScanNetworksResponseParamsClass{objc.GetClass("MTRNetworkCommissioningClusterScanNetworksResponseParams")}
	})
	return MTRNetworkCommissioningClusterScanNetworksResponseParamsClass
}

type _MTRNetworkCommissioningClusterScanNetworksResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRNetworkCommissioningClusterScanNetworksResponseParams] class.
type IMTRNetworkCommissioningClusterScanNetworksResponseParams interface {
	objectivec.IObject
	// properties:
	DebugText() objc.IObject /* cross-framework: NSString */
	SetDebugText(value objc.IObject /* cross-framework: NSString */)
	NetworkingStatus() objc.IObject /* cross-framework: NSNumber */
	SetNetworkingStatus(value objc.IObject /* cross-framework: NSNumber */)
	ThreadScanResults() unsafe.Pointer
	SetThreadScanResults(value unsafe.Pointer)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	WiFiScanResults() unsafe.Pointer
	SetWiFiScanResults(value unsafe.Pointer)
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRNetworkCommissioningClusterScanNetworksResponseParams
type MTRNetworkCommissioningClusterScanNetworksResponseParams struct {
	objectivec.Object
}

// MTRNetworkCommissioningClusterScanNetworksResponseParamsFrom constructs a [MTRNetworkCommissioningClusterScanNetworksResponseParams] from an unsafe.Pointer.
func MTRNetworkCommissioningClusterScanNetworksResponseParamsFrom(ptr unsafe.Pointer) MTRNetworkCommissioningClusterScanNetworksResponseParams {
	return MTRNetworkCommissioningClusterScanNetworksResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRNetworkCommissioningClusterScanNetworksResponseParamsClass) Alloc() MTRNetworkCommissioningClusterScanNetworksResponseParams {
	rv := objc.Send[MTRNetworkCommissioningClusterScanNetworksResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRNetworkCommissioningClusterScanNetworksResponseParamsClass) New() MTRNetworkCommissioningClusterScanNetworksResponseParams {
	rv := objc.Send[MTRNetworkCommissioningClusterScanNetworksResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRNetworkCommissioningClusterScanNetworksResponseParams) Init() MTRNetworkCommissioningClusterScanNetworksResponseParams {
	rv := objc.Send[MTRNetworkCommissioningClusterScanNetworksResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRNetworkCommissioningClusterScanNetworksResponseParams) Autorelease() MTRNetworkCommissioningClusterScanNetworksResponseParams {
	rv := objc.Send[MTRNetworkCommissioningClusterScanNetworksResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRNetworkCommissioningClusterScanNetworksResponseParams creates a new MTRNetworkCommissioningClusterScanNetworksResponseParams instance.
func NewMTRNetworkCommissioningClusterScanNetworksResponseParams() MTRNetworkCommissioningClusterScanNetworksResponseParams {
	return getMTRNetworkCommissioningClusterScanNetworksResponseParamsClass().New()
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterscannetworksresponseparams/debugtext
func (m_ MTRNetworkCommissioningClusterScanNetworksResponseParams) DebugText() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("debugText"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterscannetworksresponseparams/debugtext
func (m_ MTRNetworkCommissioningClusterScanNetworksResponseParams) SetDebugText(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDebugText:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterscannetworksresponseparams/networkingstatus
func (m_ MTRNetworkCommissioningClusterScanNetworksResponseParams) NetworkingStatus() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("networkingStatus"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterscannetworksresponseparams/networkingstatus
func (m_ MTRNetworkCommissioningClusterScanNetworksResponseParams) SetNetworkingStatus(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNetworkingStatus:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterscannetworksresponseparams/threadscanresults
func (m_ MTRNetworkCommissioningClusterScanNetworksResponseParams) ThreadScanResults() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("threadScanResults"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterscannetworksresponseparams/threadscanresults
func (m_ MTRNetworkCommissioningClusterScanNetworksResponseParams) SetThreadScanResults(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setThreadScanResults:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterscannetworksresponseparams/timedinvoketimeoutms
func (m_ MTRNetworkCommissioningClusterScanNetworksResponseParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterscannetworksresponseparams/timedinvoketimeoutms
func (m_ MTRNetworkCommissioningClusterScanNetworksResponseParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterscannetworksresponseparams/wifiscanresults
func (m_ MTRNetworkCommissioningClusterScanNetworksResponseParams) WiFiScanResults() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("wiFiScanResults"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterscannetworksresponseparams/wifiscanresults
func (m_ MTRNetworkCommissioningClusterScanNetworksResponseParams) SetWiFiScanResults(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setWiFiScanResults:"), value)
}
