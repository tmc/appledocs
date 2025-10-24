// Code generated from Apple documentation for IOUSBHost. DO NOT EDIT.

package iousbhost

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class IOUSBHostCIDeviceStateMachine */


/* debug [class_header]: Header for IOUSBHostCIDeviceStateMachine */
// The class instance for the [USBHostCIDeviceStateMachine] class.
var (
	USBHostCIDeviceStateMachineClass     _USBHostCIDeviceStateMachineClass
	USBHostCIDeviceStateMachineClassOnce sync.Once
)

func getUSBHostCIDeviceStateMachineClass() _USBHostCIDeviceStateMachineClass {
	USBHostCIDeviceStateMachineClassOnce.Do(func() {
		USBHostCIDeviceStateMachineClass = _USBHostCIDeviceStateMachineClass{objc.GetClass("IOUSBHostCIDeviceStateMachine")}
	})
	return USBHostCIDeviceStateMachineClass
}

type _USBHostCIDeviceStateMachineClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for USBHostCIDeviceStateMachine */
// An interface definition for the [USBHostCIDeviceStateMachine] class.
type IUSBHostCIDeviceStateMachine interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for USBHostCIDeviceStateMachine */
	// properties:
	CompleteRoute() uint
	ControllerInterface() IOUSBHostControllerInterface
	DeviceAddress() uint
	DeviceState() USBHostCIDeviceState /* not a class type */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for USBHostCIDeviceStateMachine */
	// methods:
	InspectCommandError(command USBHostCIMessage, error_ unsafe.Pointer) bool
	RespondToCommandStatusError(command USBHostCIMessage, status USBHostCIMessageStatus /* not a class type */, error_ unsafe.Pointer) bool
	RespondToCommandStatusDeviceAddressError(command USBHostCIMessage, status USBHostCIMessageStatus /* not a class type */, deviceAddress uint, error_ unsafe.Pointer) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for USBHostCIDeviceStateMachine */
// Alloc allocates a new instance without initialization.
func (uc _USBHostCIDeviceStateMachineClass) Alloc() USBHostCIDeviceStateMachine {
	rv := objc.Send[USBHostCIDeviceStateMachine](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (uc _USBHostCIDeviceStateMachineClass) New() USBHostCIDeviceStateMachine {
	rv := objc.Send[USBHostCIDeviceStateMachine](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ USBHostCIDeviceStateMachine) Init() USBHostCIDeviceStateMachine {
	rv := objc.Send[USBHostCIDeviceStateMachine](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ USBHostCIDeviceStateMachine) Autorelease() USBHostCIDeviceStateMachine {
	rv := objc.Send[USBHostCIDeviceStateMachine](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUSBHostCIDeviceStateMachine creates a new USBHostCIDeviceStateMachine instance.
func NewUSBHostCIDeviceStateMachine() USBHostCIDeviceStateMachine {
	return getUSBHostCIDeviceStateMachineClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for USBHostCIDeviceStateMachine */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIDeviceStateMachine
type USBHostCIDeviceStateMachine struct {
	objectivec.Object
}

// USBHostCIDeviceStateMachineFrom constructs a [USBHostCIDeviceStateMachine] from an unsafe.Pointer.
func USBHostCIDeviceStateMachineFrom(ptr unsafe.Pointer) USBHostCIDeviceStateMachine {
	return USBHostCIDeviceStateMachine{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for USBHostCIDeviceStateMachine */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIDeviceStateMachine/initWithInterface:command:error:
func NewUSBHostCIDeviceStateMachineWithInterfaceCommandError(interface_ IOUSBHostControllerInterface, command USBHostCIMessage, error_ unsafe.Pointer) USBHostCIDeviceStateMachine {
	instance := getUSBHostCIDeviceStateMachineClass().Alloc()
	rv := objc.Send[USBHostCIDeviceStateMachine](instance.ID, objc.Sel("initWithInterface:command:error:"), interface_, command, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewUSBHostCIDeviceStateMachineWithInterfaceCommandError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for USBHostCIDeviceStateMachine */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for USBHostCIDeviceStateMachine */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for USBHostCIDeviceStateMachine */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIDeviceStateMachine/inspectCommand(_:)
func (u_ USBHostCIDeviceStateMachine) InspectCommandError(command USBHostCIMessage, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("inspectCommand:error:"), command, error_)
	return rv
}/* debug [instance_methods/method]: InspectCommandError */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIDeviceStateMachine/respond(toCommand:status:)
func (u_ USBHostCIDeviceStateMachine) RespondToCommandStatusError(command USBHostCIMessage, status USBHostCIMessageStatus /* not a class type */, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("respondToCommand:status:error:"), command, status, error_)
	return rv
}/* debug [instance_methods/method]: RespondToCommandStatusError */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIDeviceStateMachine/respond(toCommand:status:deviceAddress:)
func (u_ USBHostCIDeviceStateMachine) RespondToCommandStatusDeviceAddressError(command USBHostCIMessage, status USBHostCIMessageStatus /* not a class type */, deviceAddress uint, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("respondToCommand:status:deviceAddress:error:"), command, status, deviceAddress, error_)
	return rv
}/* debug [instance_methods/method]: RespondToCommandStatusDeviceAddressError */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for USBHostCIDeviceStateMachine */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIDeviceStateMachine/completeRoute
func (u_ USBHostCIDeviceStateMachine) CompleteRoute() uint {
	rv := objc.Send[uint](u_.ID, objc.Sel("completeRoute"))
	return rv
}/* debug [instance_properties/getter]: completeRoute */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIDeviceStateMachine/controllerInterface
func (u_ USBHostCIDeviceStateMachine) ControllerInterface() IOUSBHostControllerInterface {
	rv := objc.Send[objc.ID](u_.ID, objc.Sel("controllerInterface"))
	return rv
}/* debug [instance_properties/getter]: controllerInterface */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIDeviceStateMachine/deviceAddress
func (u_ USBHostCIDeviceStateMachine) DeviceAddress() uint {
	rv := objc.Send[uint](u_.ID, objc.Sel("deviceAddress"))
	return rv
}/* debug [instance_properties/getter]: deviceAddress */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIDeviceStateMachine/deviceState
func (u_ USBHostCIDeviceStateMachine) DeviceState() USBHostCIDeviceState /* not a class type */ {
	rv := objc.Send[USBHostCIDeviceState](u_.ID, objc.Sel("deviceState"))
	return rv
}/* debug [instance_properties/getter]: deviceState */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class IOUSBHostCIDeviceStateMachine */


