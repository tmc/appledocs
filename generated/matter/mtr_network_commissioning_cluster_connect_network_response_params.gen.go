// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
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


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterconnectnetworkresponseparams/debugtext
func (m_ MTRNetworkCommissioningClusterConnectNetworkResponseParams) DebugText() string {
	rv := objc.Send[string](m_.ID, objc.Sel("debugText"))
	return rv
}


// SetDebugText sets the value of the debugText property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterconnectnetworkresponseparams/debugtext
func (m_ MTRNetworkCommissioningClusterConnectNetworkResponseParams) SetDebugText(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDebugText:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterconnectnetworkresponseparams/errorvalue
func (m_ MTRNetworkCommissioningClusterConnectNetworkResponseParams) ErrorValue() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("errorValue"))
	return rv
}


// SetErrorValue sets the value of the errorValue property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterconnectnetworkresponseparams/errorvalue
func (m_ MTRNetworkCommissioningClusterConnectNetworkResponseParams) SetErrorValue(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setErrorValue:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterconnectnetworkresponseparams/networkingstatus
func (m_ MTRNetworkCommissioningClusterConnectNetworkResponseParams) NetworkingStatus() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("networkingStatus"))
	return rv
}


// SetNetworkingStatus sets the value of the networkingStatus property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterconnectnetworkresponseparams/networkingstatus
func (m_ MTRNetworkCommissioningClusterConnectNetworkResponseParams) SetNetworkingStatus(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNetworkingStatus:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterconnectnetworkresponseparams/timedinvoketimeoutms
func (m_ MTRNetworkCommissioningClusterConnectNetworkResponseParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterconnectnetworkresponseparams/timedinvoketimeoutms
func (m_ MTRNetworkCommissioningClusterConnectNetworkResponseParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



