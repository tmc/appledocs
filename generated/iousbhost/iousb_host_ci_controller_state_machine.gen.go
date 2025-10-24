// Code generated from Apple documentation for IOUSBHost. DO NOT EDIT.

package iousbhost

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class IOUSBHostCIControllerStateMachine */


/* debug [class_header]: Header for IOUSBHostCIControllerStateMachine */
// The class instance for the [USBHostCIControllerStateMachine] class.
var (
	USBHostCIControllerStateMachineClass     _USBHostCIControllerStateMachineClass
	USBHostCIControllerStateMachineClassOnce sync.Once
)

func getUSBHostCIControllerStateMachineClass() _USBHostCIControllerStateMachineClass {
	USBHostCIControllerStateMachineClassOnce.Do(func() {
		USBHostCIControllerStateMachineClass = _USBHostCIControllerStateMachineClass{objc.GetClass("IOUSBHostCIControllerStateMachine")}
	})
	return USBHostCIControllerStateMachineClass
}

type _USBHostCIControllerStateMachineClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for USBHostCIControllerStateMachine */
// An interface definition for the [USBHostCIControllerStateMachine] class.
type IUSBHostCIControllerStateMachine interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for USBHostCIControllerStateMachine */
	// properties:
	ControllerInterface() IOUSBHostControllerInterface
	ControllerState() USBHostCIControllerState /* not a class type */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for USBHostCIControllerStateMachine */
	// methods:
	EnqueueUpdatedFrameTimestampError(frame uint64, timestamp uint64, error_ unsafe.Pointer) bool
	InspectCommandError(command USBHostCIMessage, error_ unsafe.Pointer) bool
	RespondToCommandStatusError(command USBHostCIMessage, status USBHostCIMessageStatus /* not a class type */, error_ unsafe.Pointer) bool
	RespondToCommandStatusFrameTimestampError(command USBHostCIMessage, status USBHostCIMessageStatus /* not a class type */, frame uint64, timestamp uint64, error_ unsafe.Pointer) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for USBHostCIControllerStateMachine */
// Alloc allocates a new instance without initialization.
func (uc _USBHostCIControllerStateMachineClass) Alloc() USBHostCIControllerStateMachine {
	rv := objc.Send[USBHostCIControllerStateMachine](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (uc _USBHostCIControllerStateMachineClass) New() USBHostCIControllerStateMachine {
	rv := objc.Send[USBHostCIControllerStateMachine](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ USBHostCIControllerStateMachine) Init() USBHostCIControllerStateMachine {
	rv := objc.Send[USBHostCIControllerStateMachine](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ USBHostCIControllerStateMachine) Autorelease() USBHostCIControllerStateMachine {
	rv := objc.Send[USBHostCIControllerStateMachine](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUSBHostCIControllerStateMachine creates a new USBHostCIControllerStateMachine instance.
func NewUSBHostCIControllerStateMachine() USBHostCIControllerStateMachine {
	return getUSBHostCIControllerStateMachineClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for USBHostCIControllerStateMachine */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIControllerStateMachine
type USBHostCIControllerStateMachine struct {
	objectivec.Object
}

// USBHostCIControllerStateMachineFrom constructs a [USBHostCIControllerStateMachine] from an unsafe.Pointer.
func USBHostCIControllerStateMachineFrom(ptr unsafe.Pointer) USBHostCIControllerStateMachine {
	return USBHostCIControllerStateMachine{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for USBHostCIControllerStateMachine */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIControllerStateMachine/initWithInterface:error:
func NewUSBHostCIControllerStateMachineWithInterfaceError(interface_ IOUSBHostControllerInterface, error_ unsafe.Pointer) USBHostCIControllerStateMachine {
	instance := getUSBHostCIControllerStateMachineClass().Alloc()
	rv := objc.Send[USBHostCIControllerStateMachine](instance.ID, objc.Sel("initWithInterface:error:"), interface_, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewUSBHostCIControllerStateMachineWithInterfaceError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for USBHostCIControllerStateMachine */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for USBHostCIControllerStateMachine */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for USBHostCIControllerStateMachine */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIControllerStateMachine/enqueueUpdatedFrame(_:timestamp:)
func (u_ USBHostCIControllerStateMachine) EnqueueUpdatedFrameTimestampError(frame uint64, timestamp uint64, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("enqueueUpdatedFrame:timestamp:error:"), frame, timestamp, error_)
	return rv
}/* debug [instance_methods/method]: EnqueueUpdatedFrameTimestampError */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIControllerStateMachine/inspectCommand(_:)
func (u_ USBHostCIControllerStateMachine) InspectCommandError(command USBHostCIMessage, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("inspectCommand:error:"), command, error_)
	return rv
}/* debug [instance_methods/method]: InspectCommandError */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIControllerStateMachine/respond(toCommand:status:)
func (u_ USBHostCIControllerStateMachine) RespondToCommandStatusError(command USBHostCIMessage, status USBHostCIMessageStatus /* not a class type */, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("respondToCommand:status:error:"), command, status, error_)
	return rv
}/* debug [instance_methods/method]: RespondToCommandStatusError */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIControllerStateMachine/respond(toCommand:status:frame:timestamp:)
func (u_ USBHostCIControllerStateMachine) RespondToCommandStatusFrameTimestampError(command USBHostCIMessage, status USBHostCIMessageStatus /* not a class type */, frame uint64, timestamp uint64, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("respondToCommand:status:frame:timestamp:error:"), command, status, frame, timestamp, error_)
	return rv
}/* debug [instance_methods/method]: RespondToCommandStatusFrameTimestampError */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for USBHostCIControllerStateMachine */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIControllerStateMachine/controllerInterface
func (u_ USBHostCIControllerStateMachine) ControllerInterface() IOUSBHostControllerInterface {
	rv := objc.Send[objc.ID](u_.ID, objc.Sel("controllerInterface"))
	return rv
}/* debug [instance_properties/getter]: controllerInterface */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIControllerStateMachine/controllerState
func (u_ USBHostCIControllerStateMachine) ControllerState() USBHostCIControllerState /* not a class type */ {
	rv := objc.Send[USBHostCIControllerState](u_.ID, objc.Sel("controllerState"))
	return rv
}/* debug [instance_properties/getter]: controllerState */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class IOUSBHostCIControllerStateMachine */


