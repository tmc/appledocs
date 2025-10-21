// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRDishwasherModeClusterChangeToModeResponseParams] class.
var (
	MTRDishwasherModeClusterChangeToModeResponseParamsClass     _MTRDishwasherModeClusterChangeToModeResponseParamsClass
	MTRDishwasherModeClusterChangeToModeResponseParamsClassOnce sync.Once
)

func getMTRDishwasherModeClusterChangeToModeResponseParamsClass() _MTRDishwasherModeClusterChangeToModeResponseParamsClass {
	MTRDishwasherModeClusterChangeToModeResponseParamsClassOnce.Do(func() {
		MTRDishwasherModeClusterChangeToModeResponseParamsClass = _MTRDishwasherModeClusterChangeToModeResponseParamsClass{objc.GetClass("MTRDishwasherModeClusterChangeToModeResponseParams")}
	})
	return MTRDishwasherModeClusterChangeToModeResponseParamsClass
}

type _MTRDishwasherModeClusterChangeToModeResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRDishwasherModeClusterChangeToModeResponseParams] class.
type IMTRDishwasherModeClusterChangeToModeResponseParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherModeClusterChangeToModeResponseParams
type MTRDishwasherModeClusterChangeToModeResponseParams struct {
	objectivec.Object
}

// MTRDishwasherModeClusterChangeToModeResponseParamsFrom constructs a [MTRDishwasherModeClusterChangeToModeResponseParams] from an unsafe.Pointer.
func MTRDishwasherModeClusterChangeToModeResponseParamsFrom(ptr unsafe.Pointer) MTRDishwasherModeClusterChangeToModeResponseParams {
	return MTRDishwasherModeClusterChangeToModeResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDishwasherModeClusterChangeToModeResponseParamsClass) Alloc() MTRDishwasherModeClusterChangeToModeResponseParams {
	rv := objc.Send[MTRDishwasherModeClusterChangeToModeResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDishwasherModeClusterChangeToModeResponseParamsClass) New() MTRDishwasherModeClusterChangeToModeResponseParams {
	rv := objc.Send[MTRDishwasherModeClusterChangeToModeResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDishwasherModeClusterChangeToModeResponseParams) Init() MTRDishwasherModeClusterChangeToModeResponseParams {
	rv := objc.Send[MTRDishwasherModeClusterChangeToModeResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDishwasherModeClusterChangeToModeResponseParams) Autorelease() MTRDishwasherModeClusterChangeToModeResponseParams {
	rv := objc.Send[MTRDishwasherModeClusterChangeToModeResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDishwasherModeClusterChangeToModeResponseParams creates a new MTRDishwasherModeClusterChangeToModeResponseParams instance.
func NewMTRDishwasherModeClusterChangeToModeResponseParams() MTRDishwasherModeClusterChangeToModeResponseParams {
	return getMTRDishwasherModeClusterChangeToModeResponseParamsClass().New()
}




// Initialize an MTRDishwasherModeClusterChangeToModeResponseParams with a response-value dictionary of the sort that MTRDeviceResponseHandler would receive.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherModeClusterChangeToModeResponseParams/init(responseValue:)
func NewMTRDishwasherModeClusterChangeToModeResponseParamsWithResponseValueError(responseValue unsafe.Pointer, error_ unsafe.Pointer) MTRDishwasherModeClusterChangeToModeResponseParams {
	instance := getMTRDishwasherModeClusterChangeToModeResponseParamsClass().Alloc()
	rv := objc.Send[MTRDishwasherModeClusterChangeToModeResponseParams](instance.ID, objc.Sel("initWithResponseValue:error:"), responseValue, error_)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherModeClusterChangeToModeResponseParams/status
func (m_ MTRDishwasherModeClusterChangeToModeResponseParams) Status() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("status"))
	return rv
}


// SetStatus sets the value of the status property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherModeClusterChangeToModeResponseParams/status
func (m_ MTRDishwasherModeClusterChangeToModeResponseParams) SetStatus(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatus:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherModeClusterChangeToModeResponseParams/statusText
func (m_ MTRDishwasherModeClusterChangeToModeResponseParams) StatusText() string {
	rv := objc.Send[string](m_.ID, objc.Sel("statusText"))
	return rv
}


// SetStatusText sets the value of the statusText property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherModeClusterChangeToModeResponseParams/statusText
func (m_ MTRDishwasherModeClusterChangeToModeResponseParams) SetStatusText(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatusText:"), objc.String(value))
}


