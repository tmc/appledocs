// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
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
	// properties:
	DebugText() objc.IObject /* cross-framework: NSString */
	SetDebugText(value objc.IObject /* cross-framework: NSString */)
	ErrorValue() objc.IObject /* cross-framework: NSNumber */
	SetErrorValue(value objc.IObject /* cross-framework: NSNumber */)
	NetworkingStatus() objc.IObject /* cross-framework: NSNumber */
	SetNetworkingStatus(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}

// [Full Topic]
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

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterconnectnetworkresponseparams/debugtext
func (m_ MTRNetworkCommissioningClusterConnectNetworkResponseParams) DebugText() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("debugText"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterconnectnetworkresponseparams/debugtext
func (m_ MTRNetworkCommissioningClusterConnectNetworkResponseParams) SetDebugText(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDebugText:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterconnectnetworkresponseparams/errorvalue
func (m_ MTRNetworkCommissioningClusterConnectNetworkResponseParams) ErrorValue() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("errorValue"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterconnectnetworkresponseparams/errorvalue
func (m_ MTRNetworkCommissioningClusterConnectNetworkResponseParams) SetErrorValue(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setErrorValue:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterconnectnetworkresponseparams/networkingstatus
func (m_ MTRNetworkCommissioningClusterConnectNetworkResponseParams) NetworkingStatus() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("networkingStatus"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterconnectnetworkresponseparams/networkingstatus
func (m_ MTRNetworkCommissioningClusterConnectNetworkResponseParams) SetNetworkingStatus(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNetworkingStatus:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterconnectnetworkresponseparams/timedinvoketimeoutms
func (m_ MTRNetworkCommissioningClusterConnectNetworkResponseParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusterconnectnetworkresponseparams/timedinvoketimeoutms
func (m_ MTRNetworkCommissioningClusterConnectNetworkResponseParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}
