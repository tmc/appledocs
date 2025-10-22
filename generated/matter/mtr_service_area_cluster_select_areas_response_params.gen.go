// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRServiceAreaClusterSelectAreasResponseParams] class.
var (
	MTRServiceAreaClusterSelectAreasResponseParamsClass     _MTRServiceAreaClusterSelectAreasResponseParamsClass
	MTRServiceAreaClusterSelectAreasResponseParamsClassOnce sync.Once
)

func getMTRServiceAreaClusterSelectAreasResponseParamsClass() _MTRServiceAreaClusterSelectAreasResponseParamsClass {
	MTRServiceAreaClusterSelectAreasResponseParamsClassOnce.Do(func() {
		MTRServiceAreaClusterSelectAreasResponseParamsClass = _MTRServiceAreaClusterSelectAreasResponseParamsClass{objc.GetClass("MTRServiceAreaClusterSelectAreasResponseParams")}
	})
	return MTRServiceAreaClusterSelectAreasResponseParamsClass
}

type _MTRServiceAreaClusterSelectAreasResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRServiceAreaClusterSelectAreasResponseParams] class.
type IMTRServiceAreaClusterSelectAreasResponseParams interface {
	objectivec.IObject
	Status() foundation.Number
	SetStatus(value foundation.INumber)
	StatusText() string
	SetStatusText(value string)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterSelectAreasResponseParams
type MTRServiceAreaClusterSelectAreasResponseParams struct {
	objectivec.Object
}

// MTRServiceAreaClusterSelectAreasResponseParamsFrom constructs a [MTRServiceAreaClusterSelectAreasResponseParams] from an unsafe.Pointer.
func MTRServiceAreaClusterSelectAreasResponseParamsFrom(ptr unsafe.Pointer) MTRServiceAreaClusterSelectAreasResponseParams {
	return MTRServiceAreaClusterSelectAreasResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRServiceAreaClusterSelectAreasResponseParamsClass) Alloc() MTRServiceAreaClusterSelectAreasResponseParams {
	rv := objc.Send[MTRServiceAreaClusterSelectAreasResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRServiceAreaClusterSelectAreasResponseParamsClass) New() MTRServiceAreaClusterSelectAreasResponseParams {
	rv := objc.Send[MTRServiceAreaClusterSelectAreasResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRServiceAreaClusterSelectAreasResponseParams) Init() MTRServiceAreaClusterSelectAreasResponseParams {
	rv := objc.Send[MTRServiceAreaClusterSelectAreasResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRServiceAreaClusterSelectAreasResponseParams) Autorelease() MTRServiceAreaClusterSelectAreasResponseParams {
	rv := objc.Send[MTRServiceAreaClusterSelectAreasResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRServiceAreaClusterSelectAreasResponseParams creates a new MTRServiceAreaClusterSelectAreasResponseParams instance.
func NewMTRServiceAreaClusterSelectAreasResponseParams() MTRServiceAreaClusterSelectAreasResponseParams {
	return getMTRServiceAreaClusterSelectAreasResponseParamsClass().New()
}




// Initialize an MTRServiceAreaClusterSelectAreasResponseParams with a response-value dictionary of the sort that MTRDeviceResponseHandler would receive.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterSelectAreasResponseParams/init(responseValue:)
func NewMTRServiceAreaClusterSelectAreasResponseParamsWithResponseValueError(responseValue unsafe.Pointer, error_ unsafe.Pointer) MTRServiceAreaClusterSelectAreasResponseParams {
	instance := getMTRServiceAreaClusterSelectAreasResponseParamsClass().Alloc()
	rv := objc.Send[MTRServiceAreaClusterSelectAreasResponseParams](instance.ID, objc.Sel("initWithResponseValue:error:"), responseValue, error_)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterSelectAreasResponseParams/status
func (m_ MTRServiceAreaClusterSelectAreasResponseParams) Status() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("status"))
	return rv
}


// SetStatus sets the value of the status property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterSelectAreasResponseParams/status
func (m_ MTRServiceAreaClusterSelectAreasResponseParams) SetStatus(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatus:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterSelectAreasResponseParams/statusText
func (m_ MTRServiceAreaClusterSelectAreasResponseParams) StatusText() string {
	rv := objc.Send[string](m_.ID, objc.Sel("statusText"))
	return rv
}


// SetStatusText sets the value of the statusText property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterSelectAreasResponseParams/statusText
func (m_ MTRServiceAreaClusterSelectAreasResponseParams) SetStatusText(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatusText:"), objc.String(value))
}


