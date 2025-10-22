// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRThermostatClusterAtomicResponseParams] class.
var (
	MTRThermostatClusterAtomicResponseParamsClass     _MTRThermostatClusterAtomicResponseParamsClass
	MTRThermostatClusterAtomicResponseParamsClassOnce sync.Once
)

func getMTRThermostatClusterAtomicResponseParamsClass() _MTRThermostatClusterAtomicResponseParamsClass {
	MTRThermostatClusterAtomicResponseParamsClassOnce.Do(func() {
		MTRThermostatClusterAtomicResponseParamsClass = _MTRThermostatClusterAtomicResponseParamsClass{objc.GetClass("MTRThermostatClusterAtomicResponseParams")}
	})
	return MTRThermostatClusterAtomicResponseParamsClass
}

type _MTRThermostatClusterAtomicResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRThermostatClusterAtomicResponseParams] class.
type IMTRThermostatClusterAtomicResponseParams interface {
	objectivec.IObject
	AttributeStatus() objc.ID
	SetAttributeStatus(value objc.ID)
	StatusCode() foundation.Number
	SetStatusCode(value foundation.INumber)
	Timeout() foundation.Number
	SetTimeout(value foundation.INumber)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterAtomicResponseParams
type MTRThermostatClusterAtomicResponseParams struct {
	objectivec.Object
}

// MTRThermostatClusterAtomicResponseParamsFrom constructs a [MTRThermostatClusterAtomicResponseParams] from an unsafe.Pointer.
func MTRThermostatClusterAtomicResponseParamsFrom(ptr unsafe.Pointer) MTRThermostatClusterAtomicResponseParams {
	return MTRThermostatClusterAtomicResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRThermostatClusterAtomicResponseParamsClass) Alloc() MTRThermostatClusterAtomicResponseParams {
	rv := objc.Send[MTRThermostatClusterAtomicResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRThermostatClusterAtomicResponseParamsClass) New() MTRThermostatClusterAtomicResponseParams {
	rv := objc.Send[MTRThermostatClusterAtomicResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRThermostatClusterAtomicResponseParams) Init() MTRThermostatClusterAtomicResponseParams {
	rv := objc.Send[MTRThermostatClusterAtomicResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRThermostatClusterAtomicResponseParams) Autorelease() MTRThermostatClusterAtomicResponseParams {
	rv := objc.Send[MTRThermostatClusterAtomicResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRThermostatClusterAtomicResponseParams creates a new MTRThermostatClusterAtomicResponseParams instance.
func NewMTRThermostatClusterAtomicResponseParams() MTRThermostatClusterAtomicResponseParams {
	return getMTRThermostatClusterAtomicResponseParamsClass().New()
}




// Initialize an MTRThermostatClusterAtomicResponseParams with a response-value dictionary of the sort that MTRDeviceResponseHandler would receive.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterAtomicResponseParams/init(responseValue:)
func NewMTRThermostatClusterAtomicResponseParamsWithResponseValueError(responseValue unsafe.Pointer, error_ unsafe.Pointer) MTRThermostatClusterAtomicResponseParams {
	instance := getMTRThermostatClusterAtomicResponseParamsClass().Alloc()
	rv := objc.Send[MTRThermostatClusterAtomicResponseParams](instance.ID, objc.Sel("initWithResponseValue:error:"), responseValue, error_)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterAtomicResponseParams/attributeStatus
func (m_ MTRThermostatClusterAtomicResponseParams) AttributeStatus() objc.ID {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("attributeStatus"))
	return rv
}


// SetAttributeStatus sets the value of the attributeStatus property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterAtomicResponseParams/attributeStatus
func (m_ MTRThermostatClusterAtomicResponseParams) SetAttributeStatus(value objc.ID) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAttributeStatus:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterAtomicResponseParams/statusCode
func (m_ MTRThermostatClusterAtomicResponseParams) StatusCode() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("statusCode"))
	return rv
}


// SetStatusCode sets the value of the statusCode property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterAtomicResponseParams/statusCode
func (m_ MTRThermostatClusterAtomicResponseParams) SetStatusCode(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatusCode:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterAtomicResponseParams/timeout
func (m_ MTRThermostatClusterAtomicResponseParams) Timeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timeout"))
	return rv
}


// SetTimeout sets the value of the timeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterAtomicResponseParams/timeout
func (m_ MTRThermostatClusterAtomicResponseParams) SetTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimeout:"), value)
}


