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
}

//
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


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusteraddorupdatewifinetworkparams/credentials
func (m_ MTRNetworkCommissioningClusterAddOrUpdateWiFiNetworkParams) Credentials() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("credentials"))
	return rv
}


// SetCredentials sets the value of the credentials property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusteraddorupdatewifinetworkparams/credentials
func (m_ MTRNetworkCommissioningClusterAddOrUpdateWiFiNetworkParams) SetCredentials(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCredentials:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusteraddorupdatewifinetworkparams/ssid
func (m_ MTRNetworkCommissioningClusterAddOrUpdateWiFiNetworkParams) Ssid() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("ssid"))
	return rv
}


// SetSsid sets the value of the ssid property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusteraddorupdatewifinetworkparams/ssid
func (m_ MTRNetworkCommissioningClusterAddOrUpdateWiFiNetworkParams) SetSsid(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSsid:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusteraddorupdatewifinetworkparams/serversideprocessingtimeout
func (m_ MTRNetworkCommissioningClusterAddOrUpdateWiFiNetworkParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusteraddorupdatewifinetworkparams/serversideprocessingtimeout
func (m_ MTRNetworkCommissioningClusterAddOrUpdateWiFiNetworkParams) SetServerSideProcessingTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusteraddorupdatewifinetworkparams/breadcrumb
func (m_ MTRNetworkCommissioningClusterAddOrUpdateWiFiNetworkParams) Breadcrumb() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("breadcrumb"))
	return rv
}


// SetBreadcrumb sets the value of the breadcrumb property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusteraddorupdatewifinetworkparams/breadcrumb
func (m_ MTRNetworkCommissioningClusterAddOrUpdateWiFiNetworkParams) SetBreadcrumb(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBreadcrumb:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusteraddorupdatewifinetworkparams/timedinvoketimeoutms
func (m_ MTRNetworkCommissioningClusterAddOrUpdateWiFiNetworkParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusteraddorupdatewifinetworkparams/timedinvoketimeoutms
func (m_ MTRNetworkCommissioningClusterAddOrUpdateWiFiNetworkParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



