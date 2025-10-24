// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRNetworkCommissioningClusterScanNetworksParams] class.
var (
	MTRNetworkCommissioningClusterScanNetworksParamsClass     _MTRNetworkCommissioningClusterScanNetworksParamsClass
	MTRNetworkCommissioningClusterScanNetworksParamsClassOnce sync.Once
)

func getMTRNetworkCommissioningClusterScanNetworksParamsClass() _MTRNetworkCommissioningClusterScanNetworksParamsClass {
	MTRNetworkCommissioningClusterScanNetworksParamsClassOnce.Do(func() {
		MTRNetworkCommissioningClusterScanNetworksParamsClass = _MTRNetworkCommissioningClusterScanNetworksParamsClass{objc.GetClass("MTRNetworkCommissioningClusterScanNetworksParams")}
	})
	return MTRNetworkCommissioningClusterScanNetworksParamsClass
}

type _MTRNetworkCommissioningClusterScanNetworksParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRNetworkCommissioningClusterScanNetworksParams] class.
type IMTRNetworkCommissioningClusterScanNetworksParams interface {
	objectivec.IObject
	// properties:
	Breadcrumb() objc.IObject /* cross-framework: NSNumber */
	SetBreadcrumb(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	Ssid() objc.IObject /* cross-framework: Data */
	SetSsid(value objc.IObject /* cross-framework: Data */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRNetworkCommissioningClusterScanNetworksParams
type MTRNetworkCommissioningClusterScanNetworksParams struct {
	objectivec.Object
}

// MTRNetworkCommissioningClusterScanNetworksParamsFrom constructs a [MTRNetworkCommissioningClusterScanNetworksParams] from an unsafe.Pointer.
func MTRNetworkCommissioningClusterScanNetworksParamsFrom(ptr unsafe.Pointer) MTRNetworkCommissioningClusterScanNetworksParams {
	return MTRNetworkCommissioningClusterScanNetworksParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRNetworkCommissioningClusterScanNetworksParamsClass) Alloc() MTRNetworkCommissioningClusterScanNetworksParams {
	rv := objc.Send[MTRNetworkCommissioningClusterScanNetworksParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRNetworkCommissioningClusterScanNetworksParamsClass) New() MTRNetworkCommissioningClusterScanNetworksParams {
	rv := objc.Send[MTRNetworkCommissioningClusterScanNetworksParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRNetworkCommissioningClusterScanNetworksParams) Init() MTRNetworkCommissioningClusterScanNetworksParams {
	rv := objc.Send[MTRNetworkCommissioningClusterScanNetworksParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRNetworkCommissioningClusterScanNetworksParams) Autorelease() MTRNetworkCommissioningClusterScanNetworksParams {
	rv := objc.Send[MTRNetworkCommissioningClusterScanNetworksParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRNetworkCommissioningClusterScanNetworksParams creates a new MTRNetworkCommissioningClusterScanNetworksParams instance.
func NewMTRNetworkCommissioningClusterScanNetworksParams() MTRNetworkCommissioningClusterScanNetworksParams {
	return getMTRNetworkCommissioningClusterScanNetworksParamsClass().New()
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterscannetworksparams/breadcrumb
func (m_ MTRNetworkCommissioningClusterScanNetworksParams) Breadcrumb() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("breadcrumb"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterscannetworksparams/breadcrumb
func (m_ MTRNetworkCommissioningClusterScanNetworksParams) SetBreadcrumb(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBreadcrumb:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterscannetworksparams/serversideprocessingtimeout
func (m_ MTRNetworkCommissioningClusterScanNetworksParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterscannetworksparams/serversideprocessingtimeout
func (m_ MTRNetworkCommissioningClusterScanNetworksParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterscannetworksparams/ssid
func (m_ MTRNetworkCommissioningClusterScanNetworksParams) Ssid() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("ssid"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterscannetworksparams/ssid
func (m_ MTRNetworkCommissioningClusterScanNetworksParams) SetSsid(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSsid:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterscannetworksparams/timedinvoketimeoutms
func (m_ MTRNetworkCommissioningClusterScanNetworksParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterscannetworksparams/timedinvoketimeoutms
func (m_ MTRNetworkCommissioningClusterScanNetworksParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}
