// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTREnergyEVSEModeClusterChangeToModeResponseParams] class.
var (
	MTREnergyEVSEModeClusterChangeToModeResponseParamsClass     _MTREnergyEVSEModeClusterChangeToModeResponseParamsClass
	MTREnergyEVSEModeClusterChangeToModeResponseParamsClassOnce sync.Once
)

func getMTREnergyEVSEModeClusterChangeToModeResponseParamsClass() _MTREnergyEVSEModeClusterChangeToModeResponseParamsClass {
	MTREnergyEVSEModeClusterChangeToModeResponseParamsClassOnce.Do(func() {
		MTREnergyEVSEModeClusterChangeToModeResponseParamsClass = _MTREnergyEVSEModeClusterChangeToModeResponseParamsClass{objc.GetClass("MTREnergyEVSEModeClusterChangeToModeResponseParams")}
	})
	return MTREnergyEVSEModeClusterChangeToModeResponseParamsClass
}

type _MTREnergyEVSEModeClusterChangeToModeResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTREnergyEVSEModeClusterChangeToModeResponseParams] class.
type IMTREnergyEVSEModeClusterChangeToModeResponseParams interface {
	objectivec.IObject
	Status() foundation.Number
	SetStatus(value foundation.INumber)
	StatusText() string
	SetStatusText(value string)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEModeClusterChangeToModeResponseParams
type MTREnergyEVSEModeClusterChangeToModeResponseParams struct {
	objectivec.Object
}

// MTREnergyEVSEModeClusterChangeToModeResponseParamsFrom constructs a [MTREnergyEVSEModeClusterChangeToModeResponseParams] from an unsafe.Pointer.
func MTREnergyEVSEModeClusterChangeToModeResponseParamsFrom(ptr unsafe.Pointer) MTREnergyEVSEModeClusterChangeToModeResponseParams {
	return MTREnergyEVSEModeClusterChangeToModeResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTREnergyEVSEModeClusterChangeToModeResponseParamsClass) Alloc() MTREnergyEVSEModeClusterChangeToModeResponseParams {
	rv := objc.Send[MTREnergyEVSEModeClusterChangeToModeResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTREnergyEVSEModeClusterChangeToModeResponseParamsClass) New() MTREnergyEVSEModeClusterChangeToModeResponseParams {
	rv := objc.Send[MTREnergyEVSEModeClusterChangeToModeResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTREnergyEVSEModeClusterChangeToModeResponseParams) Init() MTREnergyEVSEModeClusterChangeToModeResponseParams {
	rv := objc.Send[MTREnergyEVSEModeClusterChangeToModeResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTREnergyEVSEModeClusterChangeToModeResponseParams) Autorelease() MTREnergyEVSEModeClusterChangeToModeResponseParams {
	rv := objc.Send[MTREnergyEVSEModeClusterChangeToModeResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTREnergyEVSEModeClusterChangeToModeResponseParams creates a new MTREnergyEVSEModeClusterChangeToModeResponseParams instance.
func NewMTREnergyEVSEModeClusterChangeToModeResponseParams() MTREnergyEVSEModeClusterChangeToModeResponseParams {
	return getMTREnergyEVSEModeClusterChangeToModeResponseParamsClass().New()
}




// Initialize an MTREnergyEVSEModeClusterChangeToModeResponseParams with a response-value dictionary of the sort that MTRDeviceResponseHandler would receive.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEModeClusterChangeToModeResponseParams/init(responseValue:)
func NewMTREnergyEVSEModeClusterChangeToModeResponseParamsWithResponseValueError(responseValue unsafe.Pointer, error_ unsafe.Pointer) MTREnergyEVSEModeClusterChangeToModeResponseParams {
	instance := getMTREnergyEVSEModeClusterChangeToModeResponseParamsClass().Alloc()
	rv := objc.Send[MTREnergyEVSEModeClusterChangeToModeResponseParams](instance.ID, objc.Sel("initWithResponseValue:error:"), responseValue, error_)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEModeClusterChangeToModeResponseParams/status
func (m_ MTREnergyEVSEModeClusterChangeToModeResponseParams) Status() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("status"))
	return rv
}


// SetStatus sets the value of the status property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEModeClusterChangeToModeResponseParams/status
func (m_ MTREnergyEVSEModeClusterChangeToModeResponseParams) SetStatus(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatus:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEModeClusterChangeToModeResponseParams/statusText
func (m_ MTREnergyEVSEModeClusterChangeToModeResponseParams) StatusText() string {
	rv := objc.Send[string](m_.ID, objc.Sel("statusText"))
	return rv
}


// SetStatusText sets the value of the statusText property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEModeClusterChangeToModeResponseParams/statusText
func (m_ MTREnergyEVSEModeClusterChangeToModeResponseParams) SetStatusText(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatusText:"), objc.String(value))
}


