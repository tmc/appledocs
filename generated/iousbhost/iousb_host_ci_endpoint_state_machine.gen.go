// Code generated from Apple documentation for IOUSBHost. DO NOT EDIT.

package iousbhost

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [USBHostCIEndpointStateMachine] class.
var (
	USBHostCIEndpointStateMachineClass     _USBHostCIEndpointStateMachineClass
	USBHostCIEndpointStateMachineClassOnce sync.Once
)

func getUSBHostCIEndpointStateMachineClass() _USBHostCIEndpointStateMachineClass {
	USBHostCIEndpointStateMachineClassOnce.Do(func() {
		USBHostCIEndpointStateMachineClass = _USBHostCIEndpointStateMachineClass{objc.GetClass("IOUSBHostCIEndpointStateMachine")}
	})
	return USBHostCIEndpointStateMachineClass
}

type _USBHostCIEndpointStateMachineClass struct {
	class objc.Class
}

// An interface definition for the [USBHostCIEndpointStateMachine] class.
type IUSBHostCIEndpointStateMachine interface {
	objectivec.IObject
	InspectCommandError(command unsafe.Pointer, error_ unsafe.Pointer) bool
	ProcessDoorbellError(doorbell IUSBHostCIDoorbell, error_ unsafe.Pointer) bool
	ControllerInterface() IOUSBHostControllerInterface
	CurrentTransferMessage() unsafe.Pointer
	SetCurrentTransferMessage(value unsafe.Pointer)
	DeviceAddress() int
	SetDeviceAddress(value int)
	EndpointAddress() int
	SetEndpointAddress(value int)
	EndpointState() unsafe.Pointer
	SetEndpointState(value unsafe.Pointer)
}

//
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIEndpointStateMachine
type USBHostCIEndpointStateMachine struct {
	objectivec.Object
}

// USBHostCIEndpointStateMachineFrom constructs a [USBHostCIEndpointStateMachine] from an unsafe.Pointer.
func USBHostCIEndpointStateMachineFrom(ptr unsafe.Pointer) USBHostCIEndpointStateMachine {
	return USBHostCIEndpointStateMachine{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (uc _USBHostCIEndpointStateMachineClass) Alloc() USBHostCIEndpointStateMachine {
	rv := objc.Send[USBHostCIEndpointStateMachine](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _USBHostCIEndpointStateMachineClass) New() USBHostCIEndpointStateMachine {
	rv := objc.Send[USBHostCIEndpointStateMachine](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ USBHostCIEndpointStateMachine) Init() USBHostCIEndpointStateMachine {
	rv := objc.Send[USBHostCIEndpointStateMachine](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ USBHostCIEndpointStateMachine) Autorelease() USBHostCIEndpointStateMachine {
	rv := objc.Send[USBHostCIEndpointStateMachine](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUSBHostCIEndpointStateMachine creates a new USBHostCIEndpointStateMachine instance.
func NewUSBHostCIEndpointStateMachine() USBHostCIEndpointStateMachine {
	return getUSBHostCIEndpointStateMachineClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIEndpointStateMachine/inspectCommand(_:)
func (u_ USBHostCIEndpointStateMachine) InspectCommandError(command unsafe.Pointer, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("inspectCommand:error:"), command, error_)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIEndpointStateMachine/processDoorbell(_:)
func (u_ USBHostCIEndpointStateMachine) ProcessDoorbellError(doorbell IUSBHostCIDoorbell, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("processDoorbell:error:"), doorbell, error_)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIEndpointStateMachine/controllerInterface
func (u_ USBHostCIEndpointStateMachine) ControllerInterface() IOUSBHostControllerInterface {
	rv := objc.Send[IOUSBHostControllerInterface](u_.ID, objc.Sel("controllerInterface"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostciendpointstatemachine/currenttransfermessage
func (u_ USBHostCIEndpointStateMachine) CurrentTransferMessage() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("currentTransferMessage"))
	return rv
}


// SetCurrentTransferMessage sets the value of the currentTransferMessage property.
//
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostciendpointstatemachine/currenttransfermessage
func (u_ USBHostCIEndpointStateMachine) SetCurrentTransferMessage(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setCurrentTransferMessage:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostciendpointstatemachine/deviceaddress
func (u_ USBHostCIEndpointStateMachine) DeviceAddress() int {
	rv := objc.Send[int](u_.ID, objc.Sel("deviceAddress"))
	return rv
}


// SetDeviceAddress sets the value of the deviceAddress property.
//
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostciendpointstatemachine/deviceaddress
func (u_ USBHostCIEndpointStateMachine) SetDeviceAddress(value int) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setDeviceAddress:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostciendpointstatemachine/endpointaddress
func (u_ USBHostCIEndpointStateMachine) EndpointAddress() int {
	rv := objc.Send[int](u_.ID, objc.Sel("endpointAddress"))
	return rv
}


// SetEndpointAddress sets the value of the endpointAddress property.
//
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostciendpointstatemachine/endpointaddress
func (u_ USBHostCIEndpointStateMachine) SetEndpointAddress(value int) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setEndpointAddress:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostciendpointstatemachine/endpointstate
func (u_ USBHostCIEndpointStateMachine) EndpointState() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("endpointState"))
	return rv
}


// SetEndpointState sets the value of the endpointState property.
//
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostciendpointstatemachine/endpointstate
func (u_ USBHostCIEndpointStateMachine) SetEndpointState(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setEndpointState:"), value)
}



