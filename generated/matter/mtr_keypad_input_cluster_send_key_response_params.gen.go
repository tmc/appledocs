// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRKeypadInputClusterSendKeyResponseParams] class.
var (
	MTRKeypadInputClusterSendKeyResponseParamsClass     _MTRKeypadInputClusterSendKeyResponseParamsClass
	MTRKeypadInputClusterSendKeyResponseParamsClassOnce sync.Once
)

func getMTRKeypadInputClusterSendKeyResponseParamsClass() _MTRKeypadInputClusterSendKeyResponseParamsClass {
	MTRKeypadInputClusterSendKeyResponseParamsClassOnce.Do(func() {
		MTRKeypadInputClusterSendKeyResponseParamsClass = _MTRKeypadInputClusterSendKeyResponseParamsClass{objc.GetClass("MTRKeypadInputClusterSendKeyResponseParams")}
	})
	return MTRKeypadInputClusterSendKeyResponseParamsClass
}

type _MTRKeypadInputClusterSendKeyResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRKeypadInputClusterSendKeyResponseParams] class.
type IMTRKeypadInputClusterSendKeyResponseParams interface {
	objectivec.IObject
	Status() foundation.Number
	SetStatus(value foundation.INumber)
	TimedInvokeTimeoutMs() foundation.Number
	SetTimedInvokeTimeoutMs(value foundation.INumber)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRKeypadInputClusterSendKeyResponseParams
type MTRKeypadInputClusterSendKeyResponseParams struct {
	objectivec.Object
}

// MTRKeypadInputClusterSendKeyResponseParamsFrom constructs a [MTRKeypadInputClusterSendKeyResponseParams] from an unsafe.Pointer.
func MTRKeypadInputClusterSendKeyResponseParamsFrom(ptr unsafe.Pointer) MTRKeypadInputClusterSendKeyResponseParams {
	return MTRKeypadInputClusterSendKeyResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRKeypadInputClusterSendKeyResponseParamsClass) Alloc() MTRKeypadInputClusterSendKeyResponseParams {
	rv := objc.Send[MTRKeypadInputClusterSendKeyResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRKeypadInputClusterSendKeyResponseParamsClass) New() MTRKeypadInputClusterSendKeyResponseParams {
	rv := objc.Send[MTRKeypadInputClusterSendKeyResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRKeypadInputClusterSendKeyResponseParams) Init() MTRKeypadInputClusterSendKeyResponseParams {
	rv := objc.Send[MTRKeypadInputClusterSendKeyResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRKeypadInputClusterSendKeyResponseParams) Autorelease() MTRKeypadInputClusterSendKeyResponseParams {
	rv := objc.Send[MTRKeypadInputClusterSendKeyResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRKeypadInputClusterSendKeyResponseParams creates a new MTRKeypadInputClusterSendKeyResponseParams instance.
func NewMTRKeypadInputClusterSendKeyResponseParams() MTRKeypadInputClusterSendKeyResponseParams {
	return getMTRKeypadInputClusterSendKeyResponseParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrkeypadinputclustersendkeyresponseparams/status
func (m_ MTRKeypadInputClusterSendKeyResponseParams) Status() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("status"))
	return rv
}


// SetStatus sets the value of the status property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrkeypadinputclustersendkeyresponseparams/status
func (m_ MTRKeypadInputClusterSendKeyResponseParams) SetStatus(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatus:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrkeypadinputclustersendkeyresponseparams/timedinvoketimeoutms
func (m_ MTRKeypadInputClusterSendKeyResponseParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrkeypadinputclustersendkeyresponseparams/timedinvoketimeoutms
func (m_ MTRKeypadInputClusterSendKeyResponseParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



