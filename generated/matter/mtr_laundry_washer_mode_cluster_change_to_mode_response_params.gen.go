// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRLaundryWasherModeClusterChangeToModeResponseParams] class.
var (
	MTRLaundryWasherModeClusterChangeToModeResponseParamsClass     _MTRLaundryWasherModeClusterChangeToModeResponseParamsClass
	MTRLaundryWasherModeClusterChangeToModeResponseParamsClassOnce sync.Once
)

func getMTRLaundryWasherModeClusterChangeToModeResponseParamsClass() _MTRLaundryWasherModeClusterChangeToModeResponseParamsClass {
	MTRLaundryWasherModeClusterChangeToModeResponseParamsClassOnce.Do(func() {
		MTRLaundryWasherModeClusterChangeToModeResponseParamsClass = _MTRLaundryWasherModeClusterChangeToModeResponseParamsClass{objc.GetClass("MTRLaundryWasherModeClusterChangeToModeResponseParams")}
	})
	return MTRLaundryWasherModeClusterChangeToModeResponseParamsClass
}

type _MTRLaundryWasherModeClusterChangeToModeResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRLaundryWasherModeClusterChangeToModeResponseParams] class.
type IMTRLaundryWasherModeClusterChangeToModeResponseParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLaundryWasherModeClusterChangeToModeResponseParams
type MTRLaundryWasherModeClusterChangeToModeResponseParams struct {
	objectivec.Object
}

// MTRLaundryWasherModeClusterChangeToModeResponseParamsFrom constructs a [MTRLaundryWasherModeClusterChangeToModeResponseParams] from an unsafe.Pointer.
func MTRLaundryWasherModeClusterChangeToModeResponseParamsFrom(ptr unsafe.Pointer) MTRLaundryWasherModeClusterChangeToModeResponseParams {
	return MTRLaundryWasherModeClusterChangeToModeResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRLaundryWasherModeClusterChangeToModeResponseParamsClass) Alloc() MTRLaundryWasherModeClusterChangeToModeResponseParams {
	rv := objc.Send[MTRLaundryWasherModeClusterChangeToModeResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRLaundryWasherModeClusterChangeToModeResponseParamsClass) New() MTRLaundryWasherModeClusterChangeToModeResponseParams {
	rv := objc.Send[MTRLaundryWasherModeClusterChangeToModeResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRLaundryWasherModeClusterChangeToModeResponseParams) Init() MTRLaundryWasherModeClusterChangeToModeResponseParams {
	rv := objc.Send[MTRLaundryWasherModeClusterChangeToModeResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRLaundryWasherModeClusterChangeToModeResponseParams) Autorelease() MTRLaundryWasherModeClusterChangeToModeResponseParams {
	rv := objc.Send[MTRLaundryWasherModeClusterChangeToModeResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRLaundryWasherModeClusterChangeToModeResponseParams creates a new MTRLaundryWasherModeClusterChangeToModeResponseParams instance.
func NewMTRLaundryWasherModeClusterChangeToModeResponseParams() MTRLaundryWasherModeClusterChangeToModeResponseParams {
	return getMTRLaundryWasherModeClusterChangeToModeResponseParamsClass().New()
}




// Initialize an MTRLaundryWasherModeClusterChangeToModeResponseParams with a response-value dictionary of the sort that MTRDeviceResponseHandler would receive.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLaundryWasherModeClusterChangeToModeResponseParams/init(responseValue:)
func NewMTRLaundryWasherModeClusterChangeToModeResponseParamsWithResponseValueError(responseValue unsafe.Pointer, error_ unsafe.Pointer) MTRLaundryWasherModeClusterChangeToModeResponseParams {
	instance := getMTRLaundryWasherModeClusterChangeToModeResponseParamsClass().Alloc()
	rv := objc.Send[MTRLaundryWasherModeClusterChangeToModeResponseParams](instance.ID, objc.Sel("initWithResponseValue:error:"), responseValue, error_)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLaundryWasherModeClusterChangeToModeResponseParams/status
func (m_ MTRLaundryWasherModeClusterChangeToModeResponseParams) Status() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("status"))
	return rv
}


// SetStatus sets the value of the status property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLaundryWasherModeClusterChangeToModeResponseParams/status
func (m_ MTRLaundryWasherModeClusterChangeToModeResponseParams) SetStatus(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatus:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLaundryWasherModeClusterChangeToModeResponseParams/statusText
func (m_ MTRLaundryWasherModeClusterChangeToModeResponseParams) StatusText() string {
	rv := objc.Send[string](m_.ID, objc.Sel("statusText"))
	return rv
}


// SetStatusText sets the value of the statusText property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLaundryWasherModeClusterChangeToModeResponseParams/statusText
func (m_ MTRLaundryWasherModeClusterChangeToModeResponseParams) SetStatusText(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatusText:"), objc.String(value))
}


