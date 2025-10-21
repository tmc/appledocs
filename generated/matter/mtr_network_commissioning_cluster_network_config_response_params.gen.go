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
}

//
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


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusternetworkconfigresponseparams/debugtext
func (m_ MTRNetworkCommissioningClusterNetworkConfigResponseParams) DebugText() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("debugText"))
	return rv
}


// SetDebugText sets the value of the debugText property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusternetworkconfigresponseparams/debugtext
func (m_ MTRNetworkCommissioningClusterNetworkConfigResponseParams) SetDebugText(value appkit.string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDebugText:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusternetworkconfigresponseparams/networkindex
func (m_ MTRNetworkCommissioningClusterNetworkConfigResponseParams) NetworkIndex() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("networkIndex"))
	return rv
}


// SetNetworkIndex sets the value of the networkIndex property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusternetworkconfigresponseparams/networkindex
func (m_ MTRNetworkCommissioningClusterNetworkConfigResponseParams) SetNetworkIndex(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNetworkIndex:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusternetworkconfigresponseparams/networkingstatus
func (m_ MTRNetworkCommissioningClusterNetworkConfigResponseParams) NetworkingStatus() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("networkingStatus"))
	return rv
}


// SetNetworkingStatus sets the value of the networkingStatus property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusternetworkconfigresponseparams/networkingstatus
func (m_ MTRNetworkCommissioningClusterNetworkConfigResponseParams) SetNetworkingStatus(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNetworkingStatus:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusternetworkconfigresponseparams/timedinvoketimeoutms
func (m_ MTRNetworkCommissioningClusterNetworkConfigResponseParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusternetworkconfigresponseparams/timedinvoketimeoutms
func (m_ MTRNetworkCommissioningClusterNetworkConfigResponseParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



