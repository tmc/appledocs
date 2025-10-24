// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams] class.
var (
	MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParamsClass     _MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParamsClass
	MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParamsClassOnce sync.Once
)

func getMTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParamsClass() _MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParamsClass {
	MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParamsClassOnce.Do(func() {
		MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParamsClass = _MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParamsClass{objc.GetClass("MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams")}
	})
	return MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParamsClass
}

type _MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams] class.
type IMTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams interface {
	objectivec.IObject
	// properties:
	Cause() objc.IObject /* cross-framework: NSNumber */
	SetCause(value objc.IObject /* cross-framework: NSNumber */)
	Constraints() objc.IObject /* cross-framework: NSArray */
	SetConstraints(value objc.IObject /* cross-framework: NSArray */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams
type MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams struct {
	objectivec.Object
}

// MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParamsFrom constructs a [MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams] from an unsafe.Pointer.
func MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParamsFrom(ptr unsafe.Pointer) MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams {
	return MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParamsClass) Alloc() MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams {
	rv := objc.Send[MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParamsClass) New() MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams {
	rv := objc.Send[MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams) Init() MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams {
	rv := objc.Send[MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams) Autorelease() MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams {
	rv := objc.Send[MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams creates a new MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams instance.
func NewMTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams() MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams {
	return getMTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams/cause
func (m_ MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams) Cause() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("cause"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams/cause
func (m_ MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams) SetCause(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCause:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams/constraints
func (m_ MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams) Constraints() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("constraints"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams/constraints
func (m_ MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams) SetConstraints(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setConstraints:"), value)
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams/serverSideProcessingTimeout
func (m_ MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams/serverSideProcessingTimeout
func (m_ MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams/timedInvokeTimeoutMs
func (m_ MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams/timedInvokeTimeoutMs
func (m_ MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



