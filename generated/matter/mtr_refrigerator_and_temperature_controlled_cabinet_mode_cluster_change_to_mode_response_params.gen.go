// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams] class.
var (
	MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParamsClass     _MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParamsClass
	MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParamsClassOnce sync.Once
)

func getMTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParamsClass() _MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParamsClass {
	MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParamsClassOnce.Do(func() {
		MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParamsClass = _MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParamsClass{objc.GetClass("MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams")}
	})
	return MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParamsClass
}

type _MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams] class.
type IMTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams interface {
	objectivec.IObject
	// properties:
	Status() objc.IObject /* cross-framework: NSNumber */
	SetStatus(value objc.IObject /* cross-framework: NSNumber */)
	StatusText() objc.IObject /* cross-framework: NSString */
	SetStatusText(value objc.IObject /* cross-framework: NSString */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams
type MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams struct {
	objectivec.Object
}

// MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParamsFrom constructs a [MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams] from an unsafe.Pointer.
func MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParamsFrom(ptr unsafe.Pointer) MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams {
	return MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParamsClass) Alloc() MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams {
	rv := objc.Send[MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParamsClass) New() MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams {
	rv := objc.Send[MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams) Init() MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams {
	rv := objc.Send[MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams) Autorelease() MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams {
	rv := objc.Send[MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams creates a new MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams instance.
func NewMTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams() MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams {
	return getMTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParamsClass().New()
}



// Initialize an MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams with a response-value dictionary of the sort that MTRDeviceResponseHandler would receive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams/init(responseValue:)
func NewMTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParamsWithResponseValueError(responseValue foundation.IDictionary, error_ unsafe.Pointer) MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams {
	instance := getMTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParamsClass().Alloc()
	rv := objc.Send[MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams](instance.ID, objc.Sel("initWithResponseValue:error:"), responseValue, error_)
	rv.Autorelease()
	return rv
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams/status
func (m_ MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams) Status() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("status"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams/status
func (m_ MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams) SetStatus(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatus:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams/statusText
func (m_ MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams) StatusText() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("statusText"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams/statusText
func (m_ MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeResponseParams) SetStatusText(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatusText:"), value)
}


