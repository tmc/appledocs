// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRDeviceEnergyManagementModeClusterChangeToModeResponseParams] class.
var (
	MTRDeviceEnergyManagementModeClusterChangeToModeResponseParamsClass     _MTRDeviceEnergyManagementModeClusterChangeToModeResponseParamsClass
	MTRDeviceEnergyManagementModeClusterChangeToModeResponseParamsClassOnce sync.Once
)

func getMTRDeviceEnergyManagementModeClusterChangeToModeResponseParamsClass() _MTRDeviceEnergyManagementModeClusterChangeToModeResponseParamsClass {
	MTRDeviceEnergyManagementModeClusterChangeToModeResponseParamsClassOnce.Do(func() {
		MTRDeviceEnergyManagementModeClusterChangeToModeResponseParamsClass = _MTRDeviceEnergyManagementModeClusterChangeToModeResponseParamsClass{objc.GetClass("MTRDeviceEnergyManagementModeClusterChangeToModeResponseParams")}
	})
	return MTRDeviceEnergyManagementModeClusterChangeToModeResponseParamsClass
}

type _MTRDeviceEnergyManagementModeClusterChangeToModeResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRDeviceEnergyManagementModeClusterChangeToModeResponseParams] class.
type IMTRDeviceEnergyManagementModeClusterChangeToModeResponseParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementModeClusterChangeToModeResponseParams
type MTRDeviceEnergyManagementModeClusterChangeToModeResponseParams struct {
	objectivec.Object
}

// MTRDeviceEnergyManagementModeClusterChangeToModeResponseParamsFrom constructs a [MTRDeviceEnergyManagementModeClusterChangeToModeResponseParams] from an unsafe.Pointer.
func MTRDeviceEnergyManagementModeClusterChangeToModeResponseParamsFrom(ptr unsafe.Pointer) MTRDeviceEnergyManagementModeClusterChangeToModeResponseParams {
	return MTRDeviceEnergyManagementModeClusterChangeToModeResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDeviceEnergyManagementModeClusterChangeToModeResponseParamsClass) Alloc() MTRDeviceEnergyManagementModeClusterChangeToModeResponseParams {
	rv := objc.Send[MTRDeviceEnergyManagementModeClusterChangeToModeResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDeviceEnergyManagementModeClusterChangeToModeResponseParamsClass) New() MTRDeviceEnergyManagementModeClusterChangeToModeResponseParams {
	rv := objc.Send[MTRDeviceEnergyManagementModeClusterChangeToModeResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDeviceEnergyManagementModeClusterChangeToModeResponseParams) Init() MTRDeviceEnergyManagementModeClusterChangeToModeResponseParams {
	rv := objc.Send[MTRDeviceEnergyManagementModeClusterChangeToModeResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDeviceEnergyManagementModeClusterChangeToModeResponseParams) Autorelease() MTRDeviceEnergyManagementModeClusterChangeToModeResponseParams {
	rv := objc.Send[MTRDeviceEnergyManagementModeClusterChangeToModeResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDeviceEnergyManagementModeClusterChangeToModeResponseParams creates a new MTRDeviceEnergyManagementModeClusterChangeToModeResponseParams instance.
func NewMTRDeviceEnergyManagementModeClusterChangeToModeResponseParams() MTRDeviceEnergyManagementModeClusterChangeToModeResponseParams {
	return getMTRDeviceEnergyManagementModeClusterChangeToModeResponseParamsClass().New()
}


// Initialize an MTRDeviceEnergyManagementModeClusterChangeToModeResponseParams with a response-value dictionary of the sort that MTRDeviceResponseHandler would receive.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementModeClusterChangeToModeResponseParams/init(responseValue:)
func NewMTRDeviceEnergyManagementModeClusterChangeToModeResponseParamsWithResponseValueError(responseValue unsafe.Pointer, error_ unsafe.Pointer) MTRDeviceEnergyManagementModeClusterChangeToModeResponseParams {
	instance := getMTRDeviceEnergyManagementModeClusterChangeToModeResponseParamsClass().Alloc()
	rv := objc.Send[MTRDeviceEnergyManagementModeClusterChangeToModeResponseParams](instance.ID, objc.Sel("initWithResponseValue:error:"), responseValue, error_)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementModeClusterChangeToModeResponseParams/status
func (m_ MTRDeviceEnergyManagementModeClusterChangeToModeResponseParams) Status() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("status"))
	return rv
}


// SetStatus sets the value of the status property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementModeClusterChangeToModeResponseParams/status
func (m_ MTRDeviceEnergyManagementModeClusterChangeToModeResponseParams) SetStatus(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatus:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementModeClusterChangeToModeResponseParams/statusText
func (m_ MTRDeviceEnergyManagementModeClusterChangeToModeResponseParams) StatusText() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("statusText"))
	return rv
}


// SetStatusText sets the value of the statusText property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementModeClusterChangeToModeResponseParams/statusText
func (m_ MTRDeviceEnergyManagementModeClusterChangeToModeResponseParams) SetStatusText(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatusText:"), value)
}

