// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTROvenModeClusterChangeToModeResponseParams] class.
var (
	MTROvenModeClusterChangeToModeResponseParamsClass     _MTROvenModeClusterChangeToModeResponseParamsClass
	MTROvenModeClusterChangeToModeResponseParamsClassOnce sync.Once
)

func getMTROvenModeClusterChangeToModeResponseParamsClass() _MTROvenModeClusterChangeToModeResponseParamsClass {
	MTROvenModeClusterChangeToModeResponseParamsClassOnce.Do(func() {
		MTROvenModeClusterChangeToModeResponseParamsClass = _MTROvenModeClusterChangeToModeResponseParamsClass{objc.GetClass("MTROvenModeClusterChangeToModeResponseParams")}
	})
	return MTROvenModeClusterChangeToModeResponseParamsClass
}

type _MTROvenModeClusterChangeToModeResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTROvenModeClusterChangeToModeResponseParams] class.
type IMTROvenModeClusterChangeToModeResponseParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenModeClusterChangeToModeResponseParams
type MTROvenModeClusterChangeToModeResponseParams struct {
	objectivec.Object
}

// MTROvenModeClusterChangeToModeResponseParamsFrom constructs a [MTROvenModeClusterChangeToModeResponseParams] from an unsafe.Pointer.
func MTROvenModeClusterChangeToModeResponseParamsFrom(ptr unsafe.Pointer) MTROvenModeClusterChangeToModeResponseParams {
	return MTROvenModeClusterChangeToModeResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROvenModeClusterChangeToModeResponseParamsClass) Alloc() MTROvenModeClusterChangeToModeResponseParams {
	rv := objc.Send[MTROvenModeClusterChangeToModeResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROvenModeClusterChangeToModeResponseParamsClass) New() MTROvenModeClusterChangeToModeResponseParams {
	rv := objc.Send[MTROvenModeClusterChangeToModeResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROvenModeClusterChangeToModeResponseParams) Init() MTROvenModeClusterChangeToModeResponseParams {
	rv := objc.Send[MTROvenModeClusterChangeToModeResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROvenModeClusterChangeToModeResponseParams) Autorelease() MTROvenModeClusterChangeToModeResponseParams {
	rv := objc.Send[MTROvenModeClusterChangeToModeResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROvenModeClusterChangeToModeResponseParams creates a new MTROvenModeClusterChangeToModeResponseParams instance.
func NewMTROvenModeClusterChangeToModeResponseParams() MTROvenModeClusterChangeToModeResponseParams {
	return getMTROvenModeClusterChangeToModeResponseParamsClass().New()
}


// Initialize an MTROvenModeClusterChangeToModeResponseParams with a response-value dictionary of the sort that MTRDeviceResponseHandler would receive.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenModeClusterChangeToModeResponseParams/init(responseValue:)
func NewMTROvenModeClusterChangeToModeResponseParamsWithResponseValueError(responseValue unsafe.Pointer, error_ unsafe.Pointer) MTROvenModeClusterChangeToModeResponseParams {
	instance := getMTROvenModeClusterChangeToModeResponseParamsClass().Alloc()
	rv := objc.Send[MTROvenModeClusterChangeToModeResponseParams](instance.ID, objc.Sel("initWithResponseValue:error:"), responseValue, error_)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenModeClusterChangeToModeResponseParams/status
func (m_ MTROvenModeClusterChangeToModeResponseParams) Status() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("status"))
	return rv
}


// SetStatus sets the value of the status property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenModeClusterChangeToModeResponseParams/status
func (m_ MTROvenModeClusterChangeToModeResponseParams) SetStatus(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatus:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenModeClusterChangeToModeResponseParams/statusText
func (m_ MTROvenModeClusterChangeToModeResponseParams) StatusText() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("statusText"))
	return rv
}


// SetStatusText sets the value of the statusText property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenModeClusterChangeToModeResponseParams/statusText
func (m_ MTROvenModeClusterChangeToModeResponseParams) SetStatusText(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatusText:"), value)
}

