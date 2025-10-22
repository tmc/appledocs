// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRKeypadInputClusterSendKeyParams] class.
var (
	MTRKeypadInputClusterSendKeyParamsClass     _MTRKeypadInputClusterSendKeyParamsClass
	MTRKeypadInputClusterSendKeyParamsClassOnce sync.Once
)

func getMTRKeypadInputClusterSendKeyParamsClass() _MTRKeypadInputClusterSendKeyParamsClass {
	MTRKeypadInputClusterSendKeyParamsClassOnce.Do(func() {
		MTRKeypadInputClusterSendKeyParamsClass = _MTRKeypadInputClusterSendKeyParamsClass{objc.GetClass("MTRKeypadInputClusterSendKeyParams")}
	})
	return MTRKeypadInputClusterSendKeyParamsClass
}

type _MTRKeypadInputClusterSendKeyParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRKeypadInputClusterSendKeyParams] class.
type IMTRKeypadInputClusterSendKeyParams interface {
	objectivec.IObject
	KeyCode() foundation.Number
	SetKeyCode(value foundation.INumber)
	ServerSideProcessingTimeout() foundation.Number
	SetServerSideProcessingTimeout(value foundation.INumber)
	TimedInvokeTimeoutMs() foundation.Number
	SetTimedInvokeTimeoutMs(value foundation.INumber)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRKeypadInputClusterSendKeyParams
type MTRKeypadInputClusterSendKeyParams struct {
	objectivec.Object
}

// MTRKeypadInputClusterSendKeyParamsFrom constructs a [MTRKeypadInputClusterSendKeyParams] from an unsafe.Pointer.
func MTRKeypadInputClusterSendKeyParamsFrom(ptr unsafe.Pointer) MTRKeypadInputClusterSendKeyParams {
	return MTRKeypadInputClusterSendKeyParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRKeypadInputClusterSendKeyParamsClass) Alloc() MTRKeypadInputClusterSendKeyParams {
	rv := objc.Send[MTRKeypadInputClusterSendKeyParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRKeypadInputClusterSendKeyParamsClass) New() MTRKeypadInputClusterSendKeyParams {
	rv := objc.Send[MTRKeypadInputClusterSendKeyParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRKeypadInputClusterSendKeyParams) Init() MTRKeypadInputClusterSendKeyParams {
	rv := objc.Send[MTRKeypadInputClusterSendKeyParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRKeypadInputClusterSendKeyParams) Autorelease() MTRKeypadInputClusterSendKeyParams {
	rv := objc.Send[MTRKeypadInputClusterSendKeyParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRKeypadInputClusterSendKeyParams creates a new MTRKeypadInputClusterSendKeyParams instance.
func NewMTRKeypadInputClusterSendKeyParams() MTRKeypadInputClusterSendKeyParams {
	return getMTRKeypadInputClusterSendKeyParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrkeypadinputclustersendkeyparams/keycode
func (m_ MTRKeypadInputClusterSendKeyParams) KeyCode() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("keyCode"))
	return rv
}


// SetKeyCode sets the value of the keyCode property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrkeypadinputclustersendkeyparams/keycode
func (m_ MTRKeypadInputClusterSendKeyParams) SetKeyCode(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setKeyCode:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrkeypadinputclustersendkeyparams/serversideprocessingtimeout
func (m_ MTRKeypadInputClusterSendKeyParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrkeypadinputclustersendkeyparams/serversideprocessingtimeout
func (m_ MTRKeypadInputClusterSendKeyParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrkeypadinputclustersendkeyparams/timedinvoketimeoutms
func (m_ MTRKeypadInputClusterSendKeyParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrkeypadinputclustersendkeyparams/timedinvoketimeoutms
func (m_ MTRKeypadInputClusterSendKeyParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



