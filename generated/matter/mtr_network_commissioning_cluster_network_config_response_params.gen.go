// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRNetworkCommissioningClusterNetworkConfigResponseParams] class.
var (
	MTRNetworkCommissioningClusterNetworkConfigResponseParamsClass     _MTRNetworkCommissioningClusterNetworkConfigResponseParamsClass
	MTRNetworkCommissioningClusterNetworkConfigResponseParamsClassOnce sync.Once
)

func getMTRNetworkCommissioningClusterNetworkConfigResponseParamsClass() _MTRNetworkCommissioningClusterNetworkConfigResponseParamsClass {
	MTRNetworkCommissioningClusterNetworkConfigResponseParamsClassOnce.Do(func() {
		MTRNetworkCommissioningClusterNetworkConfigResponseParamsClass = _MTRNetworkCommissioningClusterNetworkConfigResponseParamsClass{objc.GetClass("MTRNetworkCommissioningClusterNetworkConfigResponseParams")}
	})
	return MTRNetworkCommissioningClusterNetworkConfigResponseParamsClass
}

type _MTRNetworkCommissioningClusterNetworkConfigResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRNetworkCommissioningClusterNetworkConfigResponseParams] class.
type IMTRNetworkCommissioningClusterNetworkConfigResponseParams interface {
	objectivec.IObject
	// properties:
	DebugText() objc.IObject /* cross-framework: NSString */
	SetDebugText(value objc.IObject /* cross-framework: NSString */)
	NetworkIndex() objc.IObject /* cross-framework: NSNumber */
	SetNetworkIndex(value objc.IObject /* cross-framework: NSNumber */)
	NetworkingStatus() objc.IObject /* cross-framework: NSNumber */
	SetNetworkingStatus(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRNetworkCommissioningClusterNetworkConfigResponseParams
type MTRNetworkCommissioningClusterNetworkConfigResponseParams struct {
	objectivec.Object
}

// MTRNetworkCommissioningClusterNetworkConfigResponseParamsFrom constructs a [MTRNetworkCommissioningClusterNetworkConfigResponseParams] from an unsafe.Pointer.
func MTRNetworkCommissioningClusterNetworkConfigResponseParamsFrom(ptr unsafe.Pointer) MTRNetworkCommissioningClusterNetworkConfigResponseParams {
	return MTRNetworkCommissioningClusterNetworkConfigResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRNetworkCommissioningClusterNetworkConfigResponseParamsClass) Alloc() MTRNetworkCommissioningClusterNetworkConfigResponseParams {
	rv := objc.Send[MTRNetworkCommissioningClusterNetworkConfigResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRNetworkCommissioningClusterNetworkConfigResponseParamsClass) New() MTRNetworkCommissioningClusterNetworkConfigResponseParams {
	rv := objc.Send[MTRNetworkCommissioningClusterNetworkConfigResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRNetworkCommissioningClusterNetworkConfigResponseParams) Init() MTRNetworkCommissioningClusterNetworkConfigResponseParams {
	rv := objc.Send[MTRNetworkCommissioningClusterNetworkConfigResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRNetworkCommissioningClusterNetworkConfigResponseParams) Autorelease() MTRNetworkCommissioningClusterNetworkConfigResponseParams {
	rv := objc.Send[MTRNetworkCommissioningClusterNetworkConfigResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRNetworkCommissioningClusterNetworkConfigResponseParams creates a new MTRNetworkCommissioningClusterNetworkConfigResponseParams instance.
func NewMTRNetworkCommissioningClusterNetworkConfigResponseParams() MTRNetworkCommissioningClusterNetworkConfigResponseParams {
	return getMTRNetworkCommissioningClusterNetworkConfigResponseParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusternetworkconfigresponseparams/debugtext
func (m_ MTRNetworkCommissioningClusterNetworkConfigResponseParams) DebugText() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("debugText"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusternetworkconfigresponseparams/debugtext
func (m_ MTRNetworkCommissioningClusterNetworkConfigResponseParams) SetDebugText(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDebugText:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusternetworkconfigresponseparams/networkindex
func (m_ MTRNetworkCommissioningClusterNetworkConfigResponseParams) NetworkIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("networkIndex"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusternetworkconfigresponseparams/networkindex
func (m_ MTRNetworkCommissioningClusterNetworkConfigResponseParams) SetNetworkIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNetworkIndex:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusternetworkconfigresponseparams/networkingstatus
func (m_ MTRNetworkCommissioningClusterNetworkConfigResponseParams) NetworkingStatus() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("networkingStatus"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusternetworkconfigresponseparams/networkingstatus
func (m_ MTRNetworkCommissioningClusterNetworkConfigResponseParams) SetNetworkingStatus(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNetworkingStatus:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusternetworkconfigresponseparams/timedinvoketimeoutms
func (m_ MTRNetworkCommissioningClusterNetworkConfigResponseParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusternetworkconfigresponseparams/timedinvoketimeoutms
func (m_ MTRNetworkCommissioningClusterNetworkConfigResponseParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



