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
	// properties:
	AttributeStatus() objc.IObject /* cross-framework: NSArray */
	SetAttributeStatus(value objc.IObject /* cross-framework: NSArray */)
	StatusCode() objc.IObject /* cross-framework: NSNumber */
	SetStatusCode(value objc.IObject /* cross-framework: NSNumber */)
	Timeout() objc.IObject /* cross-framework: NSNumber */
	SetTimeout(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterAtomicResponseParams/init(responseValue:)
func NewMTRThermostatClusterAtomicResponseParamsWithResponseValueError(responseValue foundation.IDictionary, error_ unsafe.Pointer) MTRThermostatClusterAtomicResponseParams {
	instance := getMTRThermostatClusterAtomicResponseParamsClass().Alloc()
	rv := objc.Send[MTRThermostatClusterAtomicResponseParams](instance.ID, objc.Sel("initWithResponseValue:error:"), responseValue, error_)
	rv.Autorelease()
	return rv
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterAtomicResponseParams/attributeStatus
func (m_ MTRThermostatClusterAtomicResponseParams) AttributeStatus() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("attributeStatus"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterAtomicResponseParams/attributeStatus
func (m_ MTRThermostatClusterAtomicResponseParams) SetAttributeStatus(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAttributeStatus:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterAtomicResponseParams/statusCode
func (m_ MTRThermostatClusterAtomicResponseParams) StatusCode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("statusCode"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterAtomicResponseParams/statusCode
func (m_ MTRThermostatClusterAtomicResponseParams) SetStatusCode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatusCode:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterAtomicResponseParams/timeout
func (m_ MTRThermostatClusterAtomicResponseParams) Timeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterAtomicResponseParams/timeout
func (m_ MTRThermostatClusterAtomicResponseParams) SetTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimeout:"), value)
}


