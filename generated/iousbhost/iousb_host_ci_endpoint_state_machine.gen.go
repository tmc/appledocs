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
	// properties:
	ControllerInterface() IOUSBHostControllerInterface /* already interface */
	SetControllerInterface(value IOUSBHostControllerInterface /* already interface */)
	CurrentTransferMessage() USBHostCIMessage /* not a class type */
	SetCurrentTransferMessage(value USBHostCIMessage /* not a class type */)
	DeviceAddress() int /* primitive/slice/pointer. */
	SetDeviceAddress(value int /* primitive/slice/pointer. */)
	EndpointAddress() int /* primitive/slice/pointer. */
	SetEndpointAddress(value int /* primitive/slice/pointer. */)
	EndpointState() USBHostCIEndpointState /* not a class type */
	SetEndpointState(value USBHostCIEndpointState /* not a class type */)
	// methods:
	InspectCommandError(command USBHostCIMessage /* not a class type */, error_ unsafe.Pointer) bool /* primitive/slice/pointer. */
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIEndpointStateMachine/inspectCommand(_:)
func (u_ USBHostCIEndpointStateMachine) InspectCommandError(command USBHostCIMessage /* not a class type */, error_ unsafe.Pointer) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](u_.ID, objc.Sel("inspectCommand:error:"), command, error_)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostciendpointstatemachine/controllerinterface
func (u_ USBHostCIEndpointStateMachine) ControllerInterface() IOUSBHostControllerInterface /* already interface */ {
	rv := objc.Send[USBHostControllerInterface](u_.ID, objc.Sel("controllerInterface"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostciendpointstatemachine/controllerinterface
func (u_ USBHostCIEndpointStateMachine) SetControllerInterface(value IOUSBHostControllerInterface /* already interface */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setControllerInterface:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostciendpointstatemachine/currenttransfermessage
func (u_ USBHostCIEndpointStateMachine) CurrentTransferMessage() USBHostCIMessage /* not a class type */ {
	rv := objc.Send[USBHostCIMessage](u_.ID, objc.Sel("currentTransferMessage"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostciendpointstatemachine/currenttransfermessage
func (u_ USBHostCIEndpointStateMachine) SetCurrentTransferMessage(value USBHostCIMessage /* not a class type */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setCurrentTransferMessage:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostciendpointstatemachine/deviceaddress
func (u_ USBHostCIEndpointStateMachine) DeviceAddress() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](u_.ID, objc.Sel("deviceAddress"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostciendpointstatemachine/deviceaddress
func (u_ USBHostCIEndpointStateMachine) SetDeviceAddress(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setDeviceAddress:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostciendpointstatemachine/endpointaddress
func (u_ USBHostCIEndpointStateMachine) EndpointAddress() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](u_.ID, objc.Sel("endpointAddress"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostciendpointstatemachine/endpointaddress
func (u_ USBHostCIEndpointStateMachine) SetEndpointAddress(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setEndpointAddress:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostciendpointstatemachine/endpointstate
func (u_ USBHostCIEndpointStateMachine) EndpointState() USBHostCIEndpointState /* not a class type */ {
	rv := objc.Send[USBHostCIEndpointState](u_.ID, objc.Sel("endpointState"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostciendpointstatemachine/endpointstate
func (u_ USBHostCIEndpointStateMachine) SetEndpointState(value USBHostCIEndpointState /* not a class type */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setEndpointState:"), value)
}



