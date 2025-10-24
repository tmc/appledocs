// Code generated from Apple documentation for IOUSBHost. DO NOT EDIT.

package iousbhost

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class IOUSBHostControllerInterface */


/* debug [class_header]: Header for IOUSBHostControllerInterface */
// The class instance for the [USBHostControllerInterface] class.
var (
	USBHostControllerInterfaceClass     _USBHostControllerInterfaceClass
	USBHostControllerInterfaceClassOnce sync.Once
)

func getUSBHostControllerInterfaceClass() _USBHostControllerInterfaceClass {
	USBHostControllerInterfaceClassOnce.Do(func() {
		USBHostControllerInterfaceClass = _USBHostControllerInterfaceClass{objc.GetClass("IOUSBHostControllerInterface")}
	})
	return USBHostControllerInterfaceClass
}

type _USBHostControllerInterfaceClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for USBHostControllerInterface */
// An interface definition for the [USBHostControllerInterface] class.
type IUSBHostControllerInterface interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for USBHostControllerInterface */
	// properties:
	Capabilities() IOUSBHostCIMessage
	ControllerStateMachine() IOUSBHostCIControllerStateMachine
	InterruptRateHz() uint
	SetInterruptRateHz(value uint)
	Queue() unsafe.Pointer
	Uuid() foundation.UUID
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for USBHostControllerInterface */
	// methods:
	CapabilitiesForPort(port uint) IOUSBHostCIMessage
	DescriptionForMessage(message USBHostCIMessage) foundation.String
	Destroy()
	EnqueueInterruptError(interrupt USBHostCIMessage, error_ unsafe.Pointer) bool
	EnqueueInterruptExpediteError(interrupt USBHostCIMessage, expedite bool, error_ unsafe.Pointer) bool
	EnqueueInterruptsCountError(interrupts USBHostCIMessage, count uint, error_ unsafe.Pointer) bool
	EnqueueInterruptsCountExpediteError(interrupts USBHostCIMessage, count uint, expedite bool, error_ unsafe.Pointer) bool
	GetPortStateMachineForCommandError(command USBHostCIMessage, error_ unsafe.Pointer) IUSBHostCIPortStateMachine
	GetPortStateMachineForPortError(port uint, error_ unsafe.Pointer) IUSBHostCIPortStateMachine
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for USBHostControllerInterface */
// Alloc allocates a new instance without initialization.
func (uc _USBHostControllerInterfaceClass) Alloc() USBHostControllerInterface {
	rv := objc.Send[USBHostControllerInterface](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (uc _USBHostControllerInterfaceClass) New() USBHostControllerInterface {
	rv := objc.Send[USBHostControllerInterface](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ USBHostControllerInterface) Init() USBHostControllerInterface {
	rv := objc.Send[USBHostControllerInterface](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ USBHostControllerInterface) Autorelease() USBHostControllerInterface {
	rv := objc.Send[USBHostControllerInterface](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUSBHostControllerInterface creates a new USBHostControllerInterface instance.
func NewUSBHostControllerInterface() USBHostControllerInterface {
	return getUSBHostControllerInterfaceClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for USBHostControllerInterface */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostControllerInterface
type USBHostControllerInterface struct {
	objectivec.Object
}

// USBHostControllerInterfaceFrom constructs a [USBHostControllerInterface] from an unsafe.Pointer.
func USBHostControllerInterfaceFrom(ptr unsafe.Pointer) USBHostControllerInterface {
	return USBHostControllerInterface{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for USBHostControllerInterface */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostControllerInterface/initWithCapabilities:queue:interruptRateHz:error:commandHandler:doorbellHandler:interestHandler:
func NewUSBHostControllerInterfaceWithCapabilitiesQueueInterruptRateHzErrorCommandHandlerDoorbellHandlerInterestHandler(capabilities objc.IObject /* cross-framework: NSData */, queue unsafe.Pointer, interruptRateHz uint, error_ unsafe.Pointer, commandHandler USBHostControllerInterfaceCommandHandler /* not a class type */, doorbellHandler USBHostControllerInterfaceDoorbellHandler /* not a class type */, interestHandler ServiceInterestCallback /* not a class type */) USBHostControllerInterface {
	instance := getUSBHostControllerInterfaceClass().Alloc()
	rv := objc.Send[USBHostControllerInterface](instance.ID, objc.Sel("initWithCapabilities:queue:interruptRateHz:error:commandHandler:doorbellHandler:interestHandler:"), capabilities, queue, interruptRateHz, error_, commandHandler, doorbellHandler, interestHandler)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewUSBHostControllerInterfaceWithCapabilitiesQueueInterruptRateHzErrorCommandHandlerDoorbellHandlerInterestHandler */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for USBHostControllerInterface */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for USBHostControllerInterface */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for USBHostControllerInterface */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostControllerInterface/capabilities(forPort:)
func (u_ USBHostControllerInterface) CapabilitiesForPort(port uint) IOUSBHostCIMessage {
	rv := objc.Send[objc.ID](u_.ID, objc.Sel("capabilitiesForPort:"), port)
	return rv
}/* debug [instance_methods/method]: CapabilitiesForPort */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostControllerInterface/description(for:)
func (u_ USBHostControllerInterface) DescriptionForMessage(message USBHostCIMessage) foundation.String {
	rv := objc.Send[foundation.String](u_.ID, objc.Sel("descriptionForMessage:"), message)
	return rv
}/* debug [instance_methods/method]: DescriptionForMessage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostControllerInterface/destroy()
func (u_ USBHostControllerInterface) Destroy() {
	objc.Send[objc.ID](u_.ID, objc.Sel("destroy"))
}/* debug [instance_methods/method]: Destroy */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostControllerInterface/enqueueInterrupt(_:)
func (u_ USBHostControllerInterface) EnqueueInterruptError(interrupt USBHostCIMessage, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("enqueueInterrupt:error:"), interrupt, error_)
	return rv
}/* debug [instance_methods/method]: EnqueueInterruptError */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostControllerInterface/enqueueInterrupt(_:expedite:)
func (u_ USBHostControllerInterface) EnqueueInterruptExpediteError(interrupt USBHostCIMessage, expedite bool, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("enqueueInterrupt:expedite:error:"), interrupt, expedite, error_)
	return rv
}/* debug [instance_methods/method]: EnqueueInterruptExpediteError */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostControllerInterface/enqueueInterrupts(_:count:)
func (u_ USBHostControllerInterface) EnqueueInterruptsCountError(interrupts USBHostCIMessage, count uint, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("enqueueInterrupts:count:error:"), interrupts, count, error_)
	return rv
}/* debug [instance_methods/method]: EnqueueInterruptsCountError */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostControllerInterface/enqueueInterrupts(_:count:expedite:)
func (u_ USBHostControllerInterface) EnqueueInterruptsCountExpediteError(interrupts USBHostCIMessage, count uint, expedite bool, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("enqueueInterrupts:count:expedite:error:"), interrupts, count, expedite, error_)
	return rv
}/* debug [instance_methods/method]: EnqueueInterruptsCountExpediteError */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostControllerInterface/getPortStateMachine(forCommand:error:)
func (u_ USBHostControllerInterface) GetPortStateMachineForCommandError(command USBHostCIMessage, error_ unsafe.Pointer) IUSBHostCIPortStateMachine {
	rv := objc.Send[USBHostCIPortStateMachine](u_.ID, objc.Sel("getPortStateMachineForCommand:error:"), command, error_)
	return rv
}/* debug [instance_methods/method]: GetPortStateMachineForCommandError */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostControllerInterface/getPortStateMachine(forPort:error:)
func (u_ USBHostControllerInterface) GetPortStateMachineForPortError(port uint, error_ unsafe.Pointer) IUSBHostCIPortStateMachine {
	rv := objc.Send[USBHostCIPortStateMachine](u_.ID, objc.Sel("getPortStateMachineForPort:error:"), port, error_)
	return rv
}/* debug [instance_methods/method]: GetPortStateMachineForPortError */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for USBHostControllerInterface */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostControllerInterface/capabilities
func (u_ USBHostControllerInterface) Capabilities() IOUSBHostCIMessage {
	rv := objc.Send[objc.ID](u_.ID, objc.Sel("capabilities"))
	return rv
}/* debug [instance_properties/getter]: capabilities */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostControllerInterface/controllerStateMachine
func (u_ USBHostControllerInterface) ControllerStateMachine() IOUSBHostCIControllerStateMachine {
	rv := objc.Send[objc.ID](u_.ID, objc.Sel("controllerStateMachine"))
	return rv
}/* debug [instance_properties/getter]: controllerStateMachine */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostControllerInterface/interruptRateHz
func (u_ USBHostControllerInterface) InterruptRateHz() uint {
	rv := objc.Send[uint](u_.ID, objc.Sel("interruptRateHz"))
	return rv
}/* debug [instance_properties/getter]: interruptRateHz */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostControllerInterface/interruptRateHz
func (u_ USBHostControllerInterface) SetInterruptRateHz(value uint) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setInterruptRateHz:"), value)
}/* debug [instance_properties/setter]: interruptRateHz */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostControllerInterface/queue
func (u_ USBHostControllerInterface) Queue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("queue"))
	return rv
}/* debug [instance_properties/getter]: queue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostControllerInterface/uuid
func (u_ USBHostControllerInterface) Uuid() foundation.UUID {
	rv := objc.Send[foundation.UUID](u_.ID, objc.Sel("uuid"))
	return rv
}/* debug [instance_properties/getter]: uuid */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class IOUSBHostControllerInterface */


