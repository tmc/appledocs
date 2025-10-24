// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRDeviceEnergyManagementModeClusterChangeToModeParams] class.
var (
	MTRDeviceEnergyManagementModeClusterChangeToModeParamsClass     _MTRDeviceEnergyManagementModeClusterChangeToModeParamsClass
	MTRDeviceEnergyManagementModeClusterChangeToModeParamsClassOnce sync.Once
)

func getMTRDeviceEnergyManagementModeClusterChangeToModeParamsClass() _MTRDeviceEnergyManagementModeClusterChangeToModeParamsClass {
	MTRDeviceEnergyManagementModeClusterChangeToModeParamsClassOnce.Do(func() {
		MTRDeviceEnergyManagementModeClusterChangeToModeParamsClass = _MTRDeviceEnergyManagementModeClusterChangeToModeParamsClass{objc.GetClass("MTRDeviceEnergyManagementModeClusterChangeToModeParams")}
	})
	return MTRDeviceEnergyManagementModeClusterChangeToModeParamsClass
}

type _MTRDeviceEnergyManagementModeClusterChangeToModeParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRDeviceEnergyManagementModeClusterChangeToModeParams] class.
type IMTRDeviceEnergyManagementModeClusterChangeToModeParams interface {
	objectivec.IObject
	// properties:
	NewMode() objc.IObject /* cross-framework: NSNumber */
	SetNewMode(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementModeClusterChangeToModeParams
type MTRDeviceEnergyManagementModeClusterChangeToModeParams struct {
	objectivec.Object
}

// MTRDeviceEnergyManagementModeClusterChangeToModeParamsFrom constructs a [MTRDeviceEnergyManagementModeClusterChangeToModeParams] from an unsafe.Pointer.
func MTRDeviceEnergyManagementModeClusterChangeToModeParamsFrom(ptr unsafe.Pointer) MTRDeviceEnergyManagementModeClusterChangeToModeParams {
	return MTRDeviceEnergyManagementModeClusterChangeToModeParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDeviceEnergyManagementModeClusterChangeToModeParamsClass) Alloc() MTRDeviceEnergyManagementModeClusterChangeToModeParams {
	rv := objc.Send[MTRDeviceEnergyManagementModeClusterChangeToModeParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDeviceEnergyManagementModeClusterChangeToModeParamsClass) New() MTRDeviceEnergyManagementModeClusterChangeToModeParams {
	rv := objc.Send[MTRDeviceEnergyManagementModeClusterChangeToModeParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDeviceEnergyManagementModeClusterChangeToModeParams) Init() MTRDeviceEnergyManagementModeClusterChangeToModeParams {
	rv := objc.Send[MTRDeviceEnergyManagementModeClusterChangeToModeParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDeviceEnergyManagementModeClusterChangeToModeParams) Autorelease() MTRDeviceEnergyManagementModeClusterChangeToModeParams {
	rv := objc.Send[MTRDeviceEnergyManagementModeClusterChangeToModeParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDeviceEnergyManagementModeClusterChangeToModeParams creates a new MTRDeviceEnergyManagementModeClusterChangeToModeParams instance.
func NewMTRDeviceEnergyManagementModeClusterChangeToModeParams() MTRDeviceEnergyManagementModeClusterChangeToModeParams {
	return getMTRDeviceEnergyManagementModeClusterChangeToModeParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementModeClusterChangeToModeParams/newMode
func (m_ MTRDeviceEnergyManagementModeClusterChangeToModeParams) NewMode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("newMode"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementModeClusterChangeToModeParams/newMode
func (m_ MTRDeviceEnergyManagementModeClusterChangeToModeParams) SetNewMode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNewMode:"), value)
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementModeClusterChangeToModeParams/serverSideProcessingTimeout
func (m_ MTRDeviceEnergyManagementModeClusterChangeToModeParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementModeClusterChangeToModeParams/serverSideProcessingTimeout
func (m_ MTRDeviceEnergyManagementModeClusterChangeToModeParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementModeClusterChangeToModeParams/timedInvokeTimeoutMs
func (m_ MTRDeviceEnergyManagementModeClusterChangeToModeParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementModeClusterChangeToModeParams/timedInvokeTimeoutMs
func (m_ MTRDeviceEnergyManagementModeClusterChangeToModeParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



