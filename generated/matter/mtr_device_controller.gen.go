// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRDeviceController] class.
var (
	MTRDeviceControllerClass     _MTRDeviceControllerClass
	MTRDeviceControllerClassOnce sync.Once
)

func getMTRDeviceControllerClass() _MTRDeviceControllerClass {
	MTRDeviceControllerClassOnce.Do(func() {
		MTRDeviceControllerClass = _MTRDeviceControllerClass{objc.GetClass("MTRDeviceController")}
	})
	return MTRDeviceControllerClass
}

type _MTRDeviceControllerClass struct {
	class objc.Class
}

// An interface definition for the [MTRDeviceController] class.
type IMTRDeviceController interface {
	objectivec.IObject
	CommissionNodeWithIDCommissioningParamsError(nodeID foundation.INumber, commissioningParams IMTRCommissioningParameters, error_ unsafe.Pointer) bool
	SetDeviceControllerDelegateQueue(delegate objectivec.IObject, queue unsafe.Pointer)
	SetupCommissioningSessionWithPayloadNewNodeIDError(payload IMTRSetupPayload, newNodeID foundation.INumber, error_ unsafe.Pointer) bool
	ControllerNodeID() foundation.Number
	SetControllerNodeID(value foundation.INumber)
	ControllerNodeId() foundation.Number
	SetControllerNodeId(value foundation.INumber)
	Devices() MTRDevice
	SetDevices(value IMTRDevice)
	IsRunning() bool
	SetIsRunning(value bool)
	IsSuspended() bool
	SetIsSuspended(value bool)
	NodesWithStoredData() foundation.Number
	SetNodesWithStoredData(value foundation.INumber)
	UniqueIdentifier() foundation.UUID
	SetUniqueIdentifier(value foundation.IUUID)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceController
type MTRDeviceController struct {
	objectivec.Object
}

// MTRDeviceControllerFrom constructs a [MTRDeviceController] from an unsafe.Pointer.
func MTRDeviceControllerFrom(ptr unsafe.Pointer) MTRDeviceController {
	return MTRDeviceController{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDeviceControllerClass) Alloc() MTRDeviceController {
	rv := objc.Send[MTRDeviceController](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDeviceControllerClass) New() MTRDeviceController {
	rv := objc.Send[MTRDeviceController](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDeviceController) Init() MTRDeviceController {
	rv := objc.Send[MTRDeviceController](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDeviceController) Autorelease() MTRDeviceController {
	rv := objc.Send[MTRDeviceController](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDeviceController creates a new MTRDeviceController instance.
func NewMTRDeviceController() MTRDeviceController {
	return getMTRDeviceControllerClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceController/commissionNode(withID:commissioningParams:)
func (m_ MTRDeviceController) CommissionNodeWithIDCommissioningParamsError(nodeID foundation.INumber, commissioningParams IMTRCommissioningParameters, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("commissionNodeWithID:commissioningParams:error:"), nodeID, commissioningParams, error_)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceController/setDeviceControllerDelegate(_:queue:)
func (m_ MTRDeviceController) SetDeviceControllerDelegateQueue(delegate objectivec.IObject, queue unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDeviceControllerDelegate:queue:"), delegate, queue)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceController/setupCommissioningSession(with:newNodeID:)
func (m_ MTRDeviceController) SetupCommissioningSessionWithPayloadNewNodeIDError(payload IMTRSetupPayload, newNodeID foundation.INumber, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("setupCommissioningSessionWithPayload:newNodeID:error:"), payload, newNodeID, error_)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontroller/controllernodeid-6a04u
func (m_ MTRDeviceController) ControllerNodeID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("controllerNodeID"))
	return rv
}


// SetControllerNodeID sets the value of the controllerNodeID property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontroller/controllernodeid-6a04u
func (m_ MTRDeviceController) SetControllerNodeID(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setControllerNodeID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontroller/controllernodeid-6a03y
func (m_ MTRDeviceController) ControllerNodeId() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("controllerNodeId"))
	return rv
}


// SetControllerNodeId sets the value of the controllerNodeId property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontroller/controllernodeid-6a03y
func (m_ MTRDeviceController) SetControllerNodeId(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setControllerNodeId:"), value)
}

// Returns the list of MTRDevice instances that this controller has loaded
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontroller/devices
func (m_ MTRDeviceController) Devices() MTRDevice {
	rv := objc.Send[MTRDevice](m_.ID, objc.Sel("devices"))
	return rv
}


// SetDevices sets the value of the devices property.
// Returns the list of MTRDevice instances that this controller has loaded

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontroller/devices
func (m_ MTRDeviceController) SetDevices(value IMTRDevice) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDevices:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontroller/isrunning
func (m_ MTRDeviceController) IsRunning() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isRunning"))
	return rv
}


// SetIsRunning sets the value of the isRunning property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontroller/isrunning
func (m_ MTRDeviceController) SetIsRunning(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsRunning:"), value)
}

// If true, the controller has been suspended via
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontroller/issuspended
func (m_ MTRDeviceController) IsSuspended() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isSuspended"))
	return rv
}


// SetIsSuspended sets the value of the isSuspended property.
// If true, the controller has been suspended via

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontroller/issuspended
func (m_ MTRDeviceController) SetIsSuspended(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsSuspended:"), value)
}

// Returns the list of node IDs for which this controller has stored
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontroller/nodeswithstoreddata
func (m_ MTRDeviceController) NodesWithStoredData() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("nodesWithStoredData"))
	return rv
}


// SetNodesWithStoredData sets the value of the nodesWithStoredData property.
// Returns the list of node IDs for which this controller has stored

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontroller/nodeswithstoreddata
func (m_ MTRDeviceController) SetNodesWithStoredData(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNodesWithStoredData:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontroller/uniqueidentifier
func (m_ MTRDeviceController) UniqueIdentifier() foundation.UUID {
	rv := objc.Send[foundation.UUID](m_.ID, objc.Sel("uniqueIdentifier"))
	return rv
}


// SetUniqueIdentifier sets the value of the uniqueIdentifier property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontroller/uniqueidentifier
func (m_ MTRDeviceController) SetUniqueIdentifier(value foundation.IUUID) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUniqueIdentifier:"), value)
}



