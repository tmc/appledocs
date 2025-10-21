// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
}

//
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


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterscannetworksresponseparams/debugtext
func (m_ MTRNetworkCommissioningClusterScanNetworksResponseParams) DebugText() string {
	rv := objc.Send[string](m_.ID, objc.Sel("debugText"))
	return rv
}


// SetDebugText sets the value of the debugText property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterscannetworksresponseparams/debugtext
func (m_ MTRNetworkCommissioningClusterScanNetworksResponseParams) SetDebugText(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDebugText:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterscannetworksresponseparams/networkingstatus
func (m_ MTRNetworkCommissioningClusterScanNetworksResponseParams) NetworkingStatus() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("networkingStatus"))
	return rv
}


// SetNetworkingStatus sets the value of the networkingStatus property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterscannetworksresponseparams/networkingstatus
func (m_ MTRNetworkCommissioningClusterScanNetworksResponseParams) SetNetworkingStatus(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNetworkingStatus:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterscannetworksresponseparams/threadscanresults
func (m_ MTRNetworkCommissioningClusterScanNetworksResponseParams) ThreadScanResults() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("threadScanResults"))
	return rv
}


// SetThreadScanResults sets the value of the threadScanResults property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterscannetworksresponseparams/threadscanresults
func (m_ MTRNetworkCommissioningClusterScanNetworksResponseParams) SetThreadScanResults(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setThreadScanResults:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterscannetworksresponseparams/timedinvoketimeoutms
func (m_ MTRNetworkCommissioningClusterScanNetworksResponseParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterscannetworksresponseparams/timedinvoketimeoutms
func (m_ MTRNetworkCommissioningClusterScanNetworksResponseParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterscannetworksresponseparams/wifiscanresults
func (m_ MTRNetworkCommissioningClusterScanNetworksResponseParams) WiFiScanResults() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("wiFiScanResults"))
	return rv
}


// SetWiFiScanResults sets the value of the wiFiScanResults property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterscannetworksresponseparams/wifiscanresults
func (m_ MTRNetworkCommissioningClusterScanNetworksResponseParams) SetWiFiScanResults(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setWiFiScanResults:"), value)
}



