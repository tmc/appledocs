// Code generated from Apple documentation for IOUSBHost. DO NOT EDIT.

package iousbhost

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class IOUSBHostCIEndpointStateMachine */


/* debug [class_header]: Header for IOUSBHostCIEndpointStateMachine */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for USBHostCIEndpointStateMachine */
// An interface definition for the [USBHostCIEndpointStateMachine] class.
type IUSBHostCIEndpointStateMachine interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for USBHostCIEndpointStateMachine */
	// properties:
	ControllerInterface() IOUSBHostControllerInterface
	CurrentTransferMessage() IOUSBHostCIMessage
	DeviceAddress() uint
	EndpointAddress() uint
	EndpointState() USBHostCIEndpointState /* not a class type */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for USBHostCIEndpointStateMachine */
	// methods:
	EnqueueTransferCompletionForMessageStatusTransferLengthError(message USBHostCIMessage, status USBHostCIMessageStatus /* not a class type */, transferLength uint, error_ unsafe.Pointer) bool
	InspectCommandError(command USBHostCIMessage, error_ unsafe.Pointer) bool
	ProcessDoorbellError(doorbell USBHostCIDoorbell /* typedef */, error_ unsafe.Pointer) bool
	RespondToCommandStatusError(command USBHostCIMessage, status USBHostCIMessageStatus /* not a class type */, error_ unsafe.Pointer) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for USBHostCIEndpointStateMachine */
// Alloc allocates a new instance without initialization.
func (uc _USBHostCIEndpointStateMachineClass) Alloc() USBHostCIEndpointStateMachine {
	rv := objc.Send[USBHostCIEndpointStateMachine](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for USBHostCIEndpointStateMachine */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIEndpointStateMachine
type USBHostCIEndpointStateMachine struct {
	objectivec.Object
}

// USBHostCIEndpointStateMachineFrom constructs a [USBHostCIEndpointStateMachine] from an unsafe.Pointer.
func USBHostCIEndpointStateMachineFrom(ptr unsafe.Pointer) USBHostCIEndpointStateMachine {
	return USBHostCIEndpointStateMachine{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for USBHostCIEndpointStateMachine */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIEndpointStateMachine/initWithInterface:command:error:
func NewUSBHostCIEndpointStateMachineWithInterfaceCommandError(interface_ IOUSBHostControllerInterface, command USBHostCIMessage, error_ unsafe.Pointer) USBHostCIEndpointStateMachine {
	instance := getUSBHostCIEndpointStateMachineClass().Alloc()
	rv := objc.Send[USBHostCIEndpointStateMachine](instance.ID, objc.Sel("initWithInterface:command:error:"), interface_, command, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewUSBHostCIEndpointStateMachineWithInterfaceCommandError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for USBHostCIEndpointStateMachine */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for USBHostCIEndpointStateMachine */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for USBHostCIEndpointStateMachine */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIEndpointStateMachine/enqueueTransferCompletion(for:status:transferLength:)
func (u_ USBHostCIEndpointStateMachine) EnqueueTransferCompletionForMessageStatusTransferLengthError(message USBHostCIMessage, status USBHostCIMessageStatus /* not a class type */, transferLength uint, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("enqueueTransferCompletionForMessage:status:transferLength:error:"), message, status, transferLength, error_)
	return rv
}/* debug [instance_methods/method]: EnqueueTransferCompletionForMessageStatusTransferLengthError */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIEndpointStateMachine/inspectCommand(_:)
func (u_ USBHostCIEndpointStateMachine) InspectCommandError(command USBHostCIMessage, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("inspectCommand:error:"), command, error_)
	return rv
}/* debug [instance_methods/method]: InspectCommandError */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIEndpointStateMachine/processDoorbell(_:)
func (u_ USBHostCIEndpointStateMachine) ProcessDoorbellError(doorbell USBHostCIDoorbell /* typedef */, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("processDoorbell:error:"), doorbell, error_)
	return rv
}/* debug [instance_methods/method]: ProcessDoorbellError */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIEndpointStateMachine/respond(toCommand:status:)
func (u_ USBHostCIEndpointStateMachine) RespondToCommandStatusError(command USBHostCIMessage, status USBHostCIMessageStatus /* not a class type */, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("respondToCommand:status:error:"), command, status, error_)
	return rv
}/* debug [instance_methods/method]: RespondToCommandStatusError */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for USBHostCIEndpointStateMachine */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIEndpointStateMachine/controllerInterface
func (u_ USBHostCIEndpointStateMachine) ControllerInterface() IOUSBHostControllerInterface {
	rv := objc.Send[objc.ID](u_.ID, objc.Sel("controllerInterface"))
	return rv
}/* debug [instance_properties/getter]: controllerInterface */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIEndpointStateMachine/currentTransferMessage
func (u_ USBHostCIEndpointStateMachine) CurrentTransferMessage() IOUSBHostCIMessage {
	rv := objc.Send[objc.ID](u_.ID, objc.Sel("currentTransferMessage"))
	return rv
}/* debug [instance_properties/getter]: currentTransferMessage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIEndpointStateMachine/deviceAddress
func (u_ USBHostCIEndpointStateMachine) DeviceAddress() uint {
	rv := objc.Send[uint](u_.ID, objc.Sel("deviceAddress"))
	return rv
}/* debug [instance_properties/getter]: deviceAddress */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIEndpointStateMachine/endpointAddress
func (u_ USBHostCIEndpointStateMachine) EndpointAddress() uint {
	rv := objc.Send[uint](u_.ID, objc.Sel("endpointAddress"))
	return rv
}/* debug [instance_properties/getter]: endpointAddress */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIEndpointStateMachine/endpointState
func (u_ USBHostCIEndpointStateMachine) EndpointState() USBHostCIEndpointState /* not a class type */ {
	rv := objc.Send[USBHostCIEndpointState](u_.ID, objc.Sel("endpointState"))
	return rv
}/* debug [instance_properties/getter]: endpointState */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class IOUSBHostCIEndpointStateMachine */


