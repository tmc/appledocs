// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRWaterHeaterModeClusterChangeToModeResponseParams] class.
var (
	MTRWaterHeaterModeClusterChangeToModeResponseParamsClass     _MTRWaterHeaterModeClusterChangeToModeResponseParamsClass
	MTRWaterHeaterModeClusterChangeToModeResponseParamsClassOnce sync.Once
)

func getMTRWaterHeaterModeClusterChangeToModeResponseParamsClass() _MTRWaterHeaterModeClusterChangeToModeResponseParamsClass {
	MTRWaterHeaterModeClusterChangeToModeResponseParamsClassOnce.Do(func() {
		MTRWaterHeaterModeClusterChangeToModeResponseParamsClass = _MTRWaterHeaterModeClusterChangeToModeResponseParamsClass{objc.GetClass("MTRWaterHeaterModeClusterChangeToModeResponseParams")}
	})
	return MTRWaterHeaterModeClusterChangeToModeResponseParamsClass
}

type _MTRWaterHeaterModeClusterChangeToModeResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRWaterHeaterModeClusterChangeToModeResponseParams] class.
type IMTRWaterHeaterModeClusterChangeToModeResponseParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterModeClusterChangeToModeResponseParams
type MTRWaterHeaterModeClusterChangeToModeResponseParams struct {
	objectivec.Object
}

// MTRWaterHeaterModeClusterChangeToModeResponseParamsFrom constructs a [MTRWaterHeaterModeClusterChangeToModeResponseParams] from an unsafe.Pointer.
func MTRWaterHeaterModeClusterChangeToModeResponseParamsFrom(ptr unsafe.Pointer) MTRWaterHeaterModeClusterChangeToModeResponseParams {
	return MTRWaterHeaterModeClusterChangeToModeResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRWaterHeaterModeClusterChangeToModeResponseParamsClass) Alloc() MTRWaterHeaterModeClusterChangeToModeResponseParams {
	rv := objc.Send[MTRWaterHeaterModeClusterChangeToModeResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRWaterHeaterModeClusterChangeToModeResponseParamsClass) New() MTRWaterHeaterModeClusterChangeToModeResponseParams {
	rv := objc.Send[MTRWaterHeaterModeClusterChangeToModeResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRWaterHeaterModeClusterChangeToModeResponseParams) Init() MTRWaterHeaterModeClusterChangeToModeResponseParams {
	rv := objc.Send[MTRWaterHeaterModeClusterChangeToModeResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRWaterHeaterModeClusterChangeToModeResponseParams) Autorelease() MTRWaterHeaterModeClusterChangeToModeResponseParams {
	rv := objc.Send[MTRWaterHeaterModeClusterChangeToModeResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRWaterHeaterModeClusterChangeToModeResponseParams creates a new MTRWaterHeaterModeClusterChangeToModeResponseParams instance.
func NewMTRWaterHeaterModeClusterChangeToModeResponseParams() MTRWaterHeaterModeClusterChangeToModeResponseParams {
	return getMTRWaterHeaterModeClusterChangeToModeResponseParamsClass().New()
}




// Initialize an MTRWaterHeaterModeClusterChangeToModeResponseParams with a response-value dictionary of the sort that MTRDeviceResponseHandler would receive.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterModeClusterChangeToModeResponseParams/init(responseValue:)
func NewMTRWaterHeaterModeClusterChangeToModeResponseParamsWithResponseValueError(responseValue unsafe.Pointer, error_ unsafe.Pointer) MTRWaterHeaterModeClusterChangeToModeResponseParams {
	instance := getMTRWaterHeaterModeClusterChangeToModeResponseParamsClass().Alloc()
	rv := objc.Send[MTRWaterHeaterModeClusterChangeToModeResponseParams](instance.ID, objc.Sel("initWithResponseValue:error:"), responseValue, error_)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterModeClusterChangeToModeResponseParams/status
func (m_ MTRWaterHeaterModeClusterChangeToModeResponseParams) Status() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("status"))
	return rv
}


// SetStatus sets the value of the status property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterModeClusterChangeToModeResponseParams/status
func (m_ MTRWaterHeaterModeClusterChangeToModeResponseParams) SetStatus(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatus:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterModeClusterChangeToModeResponseParams/statusText
func (m_ MTRWaterHeaterModeClusterChangeToModeResponseParams) StatusText() string {
	rv := objc.Send[string](m_.ID, objc.Sel("statusText"))
	return rv
}


// SetStatusText sets the value of the statusText property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterModeClusterChangeToModeResponseParams/statusText
func (m_ MTRWaterHeaterModeClusterChangeToModeResponseParams) SetStatusText(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatusText:"), objc.String(value))
}


