// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRGeneralDiagnosticsClusterPayloadTestResponseParams] class.
var (
	MTRGeneralDiagnosticsClusterPayloadTestResponseParamsClass     _MTRGeneralDiagnosticsClusterPayloadTestResponseParamsClass
	MTRGeneralDiagnosticsClusterPayloadTestResponseParamsClassOnce sync.Once
)

func getMTRGeneralDiagnosticsClusterPayloadTestResponseParamsClass() _MTRGeneralDiagnosticsClusterPayloadTestResponseParamsClass {
	MTRGeneralDiagnosticsClusterPayloadTestResponseParamsClassOnce.Do(func() {
		MTRGeneralDiagnosticsClusterPayloadTestResponseParamsClass = _MTRGeneralDiagnosticsClusterPayloadTestResponseParamsClass{objc.GetClass("MTRGeneralDiagnosticsClusterPayloadTestResponseParams")}
	})
	return MTRGeneralDiagnosticsClusterPayloadTestResponseParamsClass
}

type _MTRGeneralDiagnosticsClusterPayloadTestResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRGeneralDiagnosticsClusterPayloadTestResponseParams] class.
type IMTRGeneralDiagnosticsClusterPayloadTestResponseParams interface {
	objectivec.IObject
	Payload() foundation.NSData
	SetPayload(value foundation.IData)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGeneralDiagnosticsClusterPayloadTestResponseParams
type MTRGeneralDiagnosticsClusterPayloadTestResponseParams struct {
	objectivec.Object
}

// MTRGeneralDiagnosticsClusterPayloadTestResponseParamsFrom constructs a [MTRGeneralDiagnosticsClusterPayloadTestResponseParams] from an unsafe.Pointer.
func MTRGeneralDiagnosticsClusterPayloadTestResponseParamsFrom(ptr unsafe.Pointer) MTRGeneralDiagnosticsClusterPayloadTestResponseParams {
	return MTRGeneralDiagnosticsClusterPayloadTestResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRGeneralDiagnosticsClusterPayloadTestResponseParamsClass) Alloc() MTRGeneralDiagnosticsClusterPayloadTestResponseParams {
	rv := objc.Send[MTRGeneralDiagnosticsClusterPayloadTestResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRGeneralDiagnosticsClusterPayloadTestResponseParamsClass) New() MTRGeneralDiagnosticsClusterPayloadTestResponseParams {
	rv := objc.Send[MTRGeneralDiagnosticsClusterPayloadTestResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRGeneralDiagnosticsClusterPayloadTestResponseParams) Init() MTRGeneralDiagnosticsClusterPayloadTestResponseParams {
	rv := objc.Send[MTRGeneralDiagnosticsClusterPayloadTestResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRGeneralDiagnosticsClusterPayloadTestResponseParams) Autorelease() MTRGeneralDiagnosticsClusterPayloadTestResponseParams {
	rv := objc.Send[MTRGeneralDiagnosticsClusterPayloadTestResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRGeneralDiagnosticsClusterPayloadTestResponseParams creates a new MTRGeneralDiagnosticsClusterPayloadTestResponseParams instance.
func NewMTRGeneralDiagnosticsClusterPayloadTestResponseParams() MTRGeneralDiagnosticsClusterPayloadTestResponseParams {
	return getMTRGeneralDiagnosticsClusterPayloadTestResponseParamsClass().New()
}




// Initialize an MTRGeneralDiagnosticsClusterPayloadTestResponseParams with a response-value dictionary of the sort that MTRDeviceResponseHandler would receive.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGeneralDiagnosticsClusterPayloadTestResponseParams/init(responseValue:)
func NewMTRGeneralDiagnosticsClusterPayloadTestResponseParamsWithResponseValueError(responseValue unsafe.Pointer, error_ unsafe.Pointer) MTRGeneralDiagnosticsClusterPayloadTestResponseParams {
	instance := getMTRGeneralDiagnosticsClusterPayloadTestResponseParamsClass().Alloc()
	rv := objc.Send[MTRGeneralDiagnosticsClusterPayloadTestResponseParams](instance.ID, objc.Sel("initWithResponseValue:error:"), responseValue, error_)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGeneralDiagnosticsClusterPayloadTestResponseParams/payload
func (m_ MTRGeneralDiagnosticsClusterPayloadTestResponseParams) Payload() foundation.NSData {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("payload"))
	return rv
}


// SetPayload sets the value of the payload property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGeneralDiagnosticsClusterPayloadTestResponseParams/payload
func (m_ MTRGeneralDiagnosticsClusterPayloadTestResponseParams) SetPayload(value foundation.IData) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPayload:"), value)
}


