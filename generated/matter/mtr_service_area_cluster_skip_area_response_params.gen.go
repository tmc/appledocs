// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRServiceAreaClusterSkipAreaResponseParams] class.
var (
	MTRServiceAreaClusterSkipAreaResponseParamsClass     _MTRServiceAreaClusterSkipAreaResponseParamsClass
	MTRServiceAreaClusterSkipAreaResponseParamsClassOnce sync.Once
)

func getMTRServiceAreaClusterSkipAreaResponseParamsClass() _MTRServiceAreaClusterSkipAreaResponseParamsClass {
	MTRServiceAreaClusterSkipAreaResponseParamsClassOnce.Do(func() {
		MTRServiceAreaClusterSkipAreaResponseParamsClass = _MTRServiceAreaClusterSkipAreaResponseParamsClass{objc.GetClass("MTRServiceAreaClusterSkipAreaResponseParams")}
	})
	return MTRServiceAreaClusterSkipAreaResponseParamsClass
}

type _MTRServiceAreaClusterSkipAreaResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRServiceAreaClusterSkipAreaResponseParams] class.
type IMTRServiceAreaClusterSkipAreaResponseParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterSkipAreaResponseParams
type MTRServiceAreaClusterSkipAreaResponseParams struct {
	objectivec.Object
}

// MTRServiceAreaClusterSkipAreaResponseParamsFrom constructs a [MTRServiceAreaClusterSkipAreaResponseParams] from an unsafe.Pointer.
func MTRServiceAreaClusterSkipAreaResponseParamsFrom(ptr unsafe.Pointer) MTRServiceAreaClusterSkipAreaResponseParams {
	return MTRServiceAreaClusterSkipAreaResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRServiceAreaClusterSkipAreaResponseParamsClass) Alloc() MTRServiceAreaClusterSkipAreaResponseParams {
	rv := objc.Send[MTRServiceAreaClusterSkipAreaResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRServiceAreaClusterSkipAreaResponseParamsClass) New() MTRServiceAreaClusterSkipAreaResponseParams {
	rv := objc.Send[MTRServiceAreaClusterSkipAreaResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRServiceAreaClusterSkipAreaResponseParams) Init() MTRServiceAreaClusterSkipAreaResponseParams {
	rv := objc.Send[MTRServiceAreaClusterSkipAreaResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRServiceAreaClusterSkipAreaResponseParams) Autorelease() MTRServiceAreaClusterSkipAreaResponseParams {
	rv := objc.Send[MTRServiceAreaClusterSkipAreaResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRServiceAreaClusterSkipAreaResponseParams creates a new MTRServiceAreaClusterSkipAreaResponseParams instance.
func NewMTRServiceAreaClusterSkipAreaResponseParams() MTRServiceAreaClusterSkipAreaResponseParams {
	return getMTRServiceAreaClusterSkipAreaResponseParamsClass().New()
}




// Initialize an MTRServiceAreaClusterSkipAreaResponseParams with a response-value dictionary of the sort that MTRDeviceResponseHandler would receive.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterSkipAreaResponseParams/init(responseValue:)
func NewMTRServiceAreaClusterSkipAreaResponseParamsWithResponseValueError(responseValue unsafe.Pointer, error_ unsafe.Pointer) MTRServiceAreaClusterSkipAreaResponseParams {
	instance := getMTRServiceAreaClusterSkipAreaResponseParamsClass().Alloc()
	rv := objc.Send[MTRServiceAreaClusterSkipAreaResponseParams](instance.ID, objc.Sel("initWithResponseValue:error:"), responseValue, error_)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterSkipAreaResponseParams/status
func (m_ MTRServiceAreaClusterSkipAreaResponseParams) Status() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("status"))
	return rv
}


// SetStatus sets the value of the status property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterSkipAreaResponseParams/status
func (m_ MTRServiceAreaClusterSkipAreaResponseParams) SetStatus(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatus:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterSkipAreaResponseParams/statusText
func (m_ MTRServiceAreaClusterSkipAreaResponseParams) StatusText() string {
	rv := objc.Send[string](m_.ID, objc.Sel("statusText"))
	return rv
}


// SetStatusText sets the value of the statusText property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterSkipAreaResponseParams/statusText
func (m_ MTRServiceAreaClusterSkipAreaResponseParams) SetStatusText(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatusText:"), objc.String(value))
}


