// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
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
	// properties:
	Breadcrumb() objc.IObject /* cross-framework: NSNumber */
	SetBreadcrumb(value objc.IObject /* cross-framework: NSNumber */)
	Credentials() objc.IObject /* cross-framework: Data */
	SetCredentials(value objc.IObject /* cross-framework: Data */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	Ssid() objc.IObject /* cross-framework: Data */
	SetSsid(value objc.IObject /* cross-framework: Data */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusteraddorupdatewifinetworkparams/breadcrumb
func (m_ MTRNetworkCommissioningClusterAddOrUpdateWiFiNetworkParams) Breadcrumb() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("breadcrumb"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusteraddorupdatewifinetworkparams/breadcrumb
func (m_ MTRNetworkCommissioningClusterAddOrUpdateWiFiNetworkParams) SetBreadcrumb(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBreadcrumb:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusteraddorupdatewifinetworkparams/credentials
func (m_ MTRNetworkCommissioningClusterAddOrUpdateWiFiNetworkParams) Credentials() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("credentials"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusteraddorupdatewifinetworkparams/credentials
func (m_ MTRNetworkCommissioningClusterAddOrUpdateWiFiNetworkParams) SetCredentials(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCredentials:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusteraddorupdatewifinetworkparams/serversideprocessingtimeout
func (m_ MTRNetworkCommissioningClusterAddOrUpdateWiFiNetworkParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusteraddorupdatewifinetworkparams/serversideprocessingtimeout
func (m_ MTRNetworkCommissioningClusterAddOrUpdateWiFiNetworkParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusteraddorupdatewifinetworkparams/ssid
func (m_ MTRNetworkCommissioningClusterAddOrUpdateWiFiNetworkParams) Ssid() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("ssid"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusteraddorupdatewifinetworkparams/ssid
func (m_ MTRNetworkCommissioningClusterAddOrUpdateWiFiNetworkParams) SetSsid(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSsid:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusteraddorupdatewifinetworkparams/timedinvoketimeoutms
func (m_ MTRNetworkCommissioningClusterAddOrUpdateWiFiNetworkParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusteraddorupdatewifinetworkparams/timedinvoketimeoutms
func (m_ MTRNetworkCommissioningClusterAddOrUpdateWiFiNetworkParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



