// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRDeviceEnergyManagementClusterModifyForecastRequestParams] class.
var (
	MTRDeviceEnergyManagementClusterModifyForecastRequestParamsClass     _MTRDeviceEnergyManagementClusterModifyForecastRequestParamsClass
	MTRDeviceEnergyManagementClusterModifyForecastRequestParamsClassOnce sync.Once
)

func getMTRDeviceEnergyManagementClusterModifyForecastRequestParamsClass() _MTRDeviceEnergyManagementClusterModifyForecastRequestParamsClass {
	MTRDeviceEnergyManagementClusterModifyForecastRequestParamsClassOnce.Do(func() {
		MTRDeviceEnergyManagementClusterModifyForecastRequestParamsClass = _MTRDeviceEnergyManagementClusterModifyForecastRequestParamsClass{objc.GetClass("MTRDeviceEnergyManagementClusterModifyForecastRequestParams")}
	})
	return MTRDeviceEnergyManagementClusterModifyForecastRequestParamsClass
}

type _MTRDeviceEnergyManagementClusterModifyForecastRequestParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRDeviceEnergyManagementClusterModifyForecastRequestParams] class.
type IMTRDeviceEnergyManagementClusterModifyForecastRequestParams interface {
	objectivec.IObject
	// properties:
	Cause() objc.IObject /* cross-framework: NSNumber */
	SetCause(value objc.IObject /* cross-framework: NSNumber */)
	ForecastID() objc.IObject /* cross-framework: NSNumber */
	SetForecastID(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	SlotAdjustments() objc.IObject /* cross-framework: NSArray */
	SetSlotAdjustments(value objc.IObject /* cross-framework: NSArray */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterModifyForecastRequestParams
type MTRDeviceEnergyManagementClusterModifyForecastRequestParams struct {
	objectivec.Object
}

// MTRDeviceEnergyManagementClusterModifyForecastRequestParamsFrom constructs a [MTRDeviceEnergyManagementClusterModifyForecastRequestParams] from an unsafe.Pointer.
func MTRDeviceEnergyManagementClusterModifyForecastRequestParamsFrom(ptr unsafe.Pointer) MTRDeviceEnergyManagementClusterModifyForecastRequestParams {
	return MTRDeviceEnergyManagementClusterModifyForecastRequestParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDeviceEnergyManagementClusterModifyForecastRequestParamsClass) Alloc() MTRDeviceEnergyManagementClusterModifyForecastRequestParams {
	rv := objc.Send[MTRDeviceEnergyManagementClusterModifyForecastRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDeviceEnergyManagementClusterModifyForecastRequestParamsClass) New() MTRDeviceEnergyManagementClusterModifyForecastRequestParams {
	rv := objc.Send[MTRDeviceEnergyManagementClusterModifyForecastRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDeviceEnergyManagementClusterModifyForecastRequestParams) Init() MTRDeviceEnergyManagementClusterModifyForecastRequestParams {
	rv := objc.Send[MTRDeviceEnergyManagementClusterModifyForecastRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDeviceEnergyManagementClusterModifyForecastRequestParams) Autorelease() MTRDeviceEnergyManagementClusterModifyForecastRequestParams {
	rv := objc.Send[MTRDeviceEnergyManagementClusterModifyForecastRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDeviceEnergyManagementClusterModifyForecastRequestParams creates a new MTRDeviceEnergyManagementClusterModifyForecastRequestParams instance.
func NewMTRDeviceEnergyManagementClusterModifyForecastRequestParams() MTRDeviceEnergyManagementClusterModifyForecastRequestParams {
	return getMTRDeviceEnergyManagementClusterModifyForecastRequestParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterModifyForecastRequestParams/cause
func (m_ MTRDeviceEnergyManagementClusterModifyForecastRequestParams) Cause() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("cause"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterModifyForecastRequestParams/cause
func (m_ MTRDeviceEnergyManagementClusterModifyForecastRequestParams) SetCause(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCause:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterModifyForecastRequestParams/forecastID
func (m_ MTRDeviceEnergyManagementClusterModifyForecastRequestParams) ForecastID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("forecastID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterModifyForecastRequestParams/forecastID
func (m_ MTRDeviceEnergyManagementClusterModifyForecastRequestParams) SetForecastID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setForecastID:"), value)
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterModifyForecastRequestParams/serverSideProcessingTimeout
func (m_ MTRDeviceEnergyManagementClusterModifyForecastRequestParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterModifyForecastRequestParams/serverSideProcessingTimeout
func (m_ MTRDeviceEnergyManagementClusterModifyForecastRequestParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterModifyForecastRequestParams/slotAdjustments
func (m_ MTRDeviceEnergyManagementClusterModifyForecastRequestParams) SlotAdjustments() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("slotAdjustments"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterModifyForecastRequestParams/slotAdjustments
func (m_ MTRDeviceEnergyManagementClusterModifyForecastRequestParams) SetSlotAdjustments(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSlotAdjustments:"), value)
}


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterModifyForecastRequestParams/timedInvokeTimeoutMs
func (m_ MTRDeviceEnergyManagementClusterModifyForecastRequestParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterModifyForecastRequestParams/timedInvokeTimeoutMs
func (m_ MTRDeviceEnergyManagementClusterModifyForecastRequestParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



