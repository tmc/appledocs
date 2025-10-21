// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRContentAppObserverClusterContentAppMessageResponseParams] class.
var (
	MTRContentAppObserverClusterContentAppMessageResponseParamsClass     _MTRContentAppObserverClusterContentAppMessageResponseParamsClass
	MTRContentAppObserverClusterContentAppMessageResponseParamsClassOnce sync.Once
)

func getMTRContentAppObserverClusterContentAppMessageResponseParamsClass() _MTRContentAppObserverClusterContentAppMessageResponseParamsClass {
	MTRContentAppObserverClusterContentAppMessageResponseParamsClassOnce.Do(func() {
		MTRContentAppObserverClusterContentAppMessageResponseParamsClass = _MTRContentAppObserverClusterContentAppMessageResponseParamsClass{objc.GetClass("MTRContentAppObserverClusterContentAppMessageResponseParams")}
	})
	return MTRContentAppObserverClusterContentAppMessageResponseParamsClass
}

type _MTRContentAppObserverClusterContentAppMessageResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRContentAppObserverClusterContentAppMessageResponseParams] class.
type IMTRContentAppObserverClusterContentAppMessageResponseParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentAppObserverClusterContentAppMessageResponseParams
type MTRContentAppObserverClusterContentAppMessageResponseParams struct {
	objectivec.Object
}

// MTRContentAppObserverClusterContentAppMessageResponseParamsFrom constructs a [MTRContentAppObserverClusterContentAppMessageResponseParams] from an unsafe.Pointer.
func MTRContentAppObserverClusterContentAppMessageResponseParamsFrom(ptr unsafe.Pointer) MTRContentAppObserverClusterContentAppMessageResponseParams {
	return MTRContentAppObserverClusterContentAppMessageResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRContentAppObserverClusterContentAppMessageResponseParamsClass) Alloc() MTRContentAppObserverClusterContentAppMessageResponseParams {
	rv := objc.Send[MTRContentAppObserverClusterContentAppMessageResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRContentAppObserverClusterContentAppMessageResponseParamsClass) New() MTRContentAppObserverClusterContentAppMessageResponseParams {
	rv := objc.Send[MTRContentAppObserverClusterContentAppMessageResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRContentAppObserverClusterContentAppMessageResponseParams) Init() MTRContentAppObserverClusterContentAppMessageResponseParams {
	rv := objc.Send[MTRContentAppObserverClusterContentAppMessageResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRContentAppObserverClusterContentAppMessageResponseParams) Autorelease() MTRContentAppObserverClusterContentAppMessageResponseParams {
	rv := objc.Send[MTRContentAppObserverClusterContentAppMessageResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRContentAppObserverClusterContentAppMessageResponseParams creates a new MTRContentAppObserverClusterContentAppMessageResponseParams instance.
func NewMTRContentAppObserverClusterContentAppMessageResponseParams() MTRContentAppObserverClusterContentAppMessageResponseParams {
	return getMTRContentAppObserverClusterContentAppMessageResponseParamsClass().New()
}




// Initialize an MTRContentAppObserverClusterContentAppMessageResponseParams with a response-value dictionary of the sort that MTRDeviceResponseHandler would receive.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentAppObserverClusterContentAppMessageResponseParams/init(responseValue:)
func NewMTRContentAppObserverClusterContentAppMessageResponseParamsWithResponseValueError(responseValue unsafe.Pointer, error_ unsafe.Pointer) MTRContentAppObserverClusterContentAppMessageResponseParams {
	instance := getMTRContentAppObserverClusterContentAppMessageResponseParamsClass().Alloc()
	rv := objc.Send[MTRContentAppObserverClusterContentAppMessageResponseParams](instance.ID, objc.Sel("initWithResponseValue:error:"), responseValue, error_)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentAppObserverClusterContentAppMessageResponseParams/data
func (m_ MTRContentAppObserverClusterContentAppMessageResponseParams) Data() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("data"))
	return rv
}


// SetData sets the value of the data property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentAppObserverClusterContentAppMessageResponseParams/data
func (m_ MTRContentAppObserverClusterContentAppMessageResponseParams) SetData(value appkit.string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setData:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentAppObserverClusterContentAppMessageResponseParams/encodingHint
func (m_ MTRContentAppObserverClusterContentAppMessageResponseParams) EncodingHint() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("encodingHint"))
	return rv
}


// SetEncodingHint sets the value of the encodingHint property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentAppObserverClusterContentAppMessageResponseParams/encodingHint
func (m_ MTRContentAppObserverClusterContentAppMessageResponseParams) SetEncodingHint(value appkit.string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEncodingHint:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentAppObserverClusterContentAppMessageResponseParams/status
func (m_ MTRContentAppObserverClusterContentAppMessageResponseParams) Status() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("status"))
	return rv
}


// SetStatus sets the value of the status property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentAppObserverClusterContentAppMessageResponseParams/status
func (m_ MTRContentAppObserverClusterContentAppMessageResponseParams) SetStatus(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatus:"), value)
}


