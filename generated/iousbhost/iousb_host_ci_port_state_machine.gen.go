// Code generated from Apple documentation for IOUSBHost. DO NOT EDIT.

package iousbhost

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class IOUSBHostCIPortStateMachine */


/* debug [class_header]: Header for IOUSBHostCIPortStateMachine */
// The class instance for the [USBHostCIPortStateMachine] class.
var (
	USBHostCIPortStateMachineClass     _USBHostCIPortStateMachineClass
	USBHostCIPortStateMachineClassOnce sync.Once
)

func getUSBHostCIPortStateMachineClass() _USBHostCIPortStateMachineClass {
	USBHostCIPortStateMachineClassOnce.Do(func() {
		USBHostCIPortStateMachineClass = _USBHostCIPortStateMachineClass{objc.GetClass("IOUSBHostCIPortStateMachine")}
	})
	return USBHostCIPortStateMachineClass
}

type _USBHostCIPortStateMachineClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for USBHostCIPortStateMachine */
// An interface definition for the [USBHostCIPortStateMachine] class.
type IUSBHostCIPortStateMachine interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for USBHostCIPortStateMachine */
	// properties:
	Connected() bool
	SetConnected(value bool)
	ControllerInterface() IOUSBHostControllerInterface
	LinkState() USBHostCILinkState /* not a class type */
	Overcurrent() bool
	SetOvercurrent(value bool)
	PortNumber() uint
	PortState() USBHostCIPortState /* not a class type */
	PortStatus() USBHostCIPortStatus /* typedef */
	Powered() bool
	SetPowered(value bool)
	Speed() USBHostCIDeviceSpeed /* not a class type */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for USBHostCIPortStateMachine */
	// methods:
	InspectCommandError(command USBHostCIMessage, error_ unsafe.Pointer) bool
	RespondToCommandStatusError(command USBHostCIMessage, status USBHostCIMessageStatus /* not a class type */, error_ unsafe.Pointer) bool
	UpdateLinkStateSpeedInhibitLinkStateChangeError(linkState USBHostCILinkState /* not a class type */, speed USBHostCIDeviceSpeed /* not a class type */, inhibitLinkStateChange bool, error_ unsafe.Pointer) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for USBHostCIPortStateMachine */
// Alloc allocates a new instance without initialization.
func (uc _USBHostCIPortStateMachineClass) Alloc() USBHostCIPortStateMachine {
	rv := objc.Send[USBHostCIPortStateMachine](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (uc _USBHostCIPortStateMachineClass) New() USBHostCIPortStateMachine {
	rv := objc.Send[USBHostCIPortStateMachine](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ USBHostCIPortStateMachine) Init() USBHostCIPortStateMachine {
	rv := objc.Send[USBHostCIPortStateMachine](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ USBHostCIPortStateMachine) Autorelease() USBHostCIPortStateMachine {
	rv := objc.Send[USBHostCIPortStateMachine](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUSBHostCIPortStateMachine creates a new USBHostCIPortStateMachine instance.
func NewUSBHostCIPortStateMachine() USBHostCIPortStateMachine {
	return getUSBHostCIPortStateMachineClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for USBHostCIPortStateMachine */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIPortStateMachine
type USBHostCIPortStateMachine struct {
	objectivec.Object
}

// USBHostCIPortStateMachineFrom constructs a [USBHostCIPortStateMachine] from an unsafe.Pointer.
func USBHostCIPortStateMachineFrom(ptr unsafe.Pointer) USBHostCIPortStateMachine {
	return USBHostCIPortStateMachine{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for USBHostCIPortStateMachine */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIPortStateMachine/initWithInterface:portNumber:error:
func NewUSBHostCIPortStateMachineWithInterfacePortNumberError(interface_ IOUSBHostControllerInterface, portNumber uint, error_ unsafe.Pointer) USBHostCIPortStateMachine {
	instance := getUSBHostCIPortStateMachineClass().Alloc()
	rv := objc.Send[USBHostCIPortStateMachine](instance.ID, objc.Sel("initWithInterface:portNumber:error:"), interface_, portNumber, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewUSBHostCIPortStateMachineWithInterfacePortNumberError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for USBHostCIPortStateMachine */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for USBHostCIPortStateMachine */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for USBHostCIPortStateMachine */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIPortStateMachine/inspectCommand(_:)
func (u_ USBHostCIPortStateMachine) InspectCommandError(command USBHostCIMessage, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("inspectCommand:error:"), command, error_)
	return rv
}/* debug [instance_methods/method]: InspectCommandError */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIPortStateMachine/respond(toCommand:status:)
func (u_ USBHostCIPortStateMachine) RespondToCommandStatusError(command USBHostCIMessage, status USBHostCIMessageStatus /* not a class type */, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("respondToCommand:status:error:"), command, status, error_)
	return rv
}/* debug [instance_methods/method]: RespondToCommandStatusError */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIPortStateMachine/updateLinkState(_:speed:inhibitLinkStateChange:)
func (u_ USBHostCIPortStateMachine) UpdateLinkStateSpeedInhibitLinkStateChangeError(linkState USBHostCILinkState /* not a class type */, speed USBHostCIDeviceSpeed /* not a class type */, inhibitLinkStateChange bool, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("updateLinkState:speed:inhibitLinkStateChange:error:"), linkState, speed, inhibitLinkStateChange, error_)
	return rv
}/* debug [instance_methods/method]: UpdateLinkStateSpeedInhibitLinkStateChangeError */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for USBHostCIPortStateMachine */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIPortStateMachine/connected
func (u_ USBHostCIPortStateMachine) Connected() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("connected"))
	return rv
}/* debug [instance_properties/getter]: connected */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIPortStateMachine/connected
func (u_ USBHostCIPortStateMachine) SetConnected(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setConnected:"), value)
}/* debug [instance_properties/setter]: connected */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIPortStateMachine/controllerInterface
func (u_ USBHostCIPortStateMachine) ControllerInterface() IOUSBHostControllerInterface {
	rv := objc.Send[objc.ID](u_.ID, objc.Sel("controllerInterface"))
	return rv
}/* debug [instance_properties/getter]: controllerInterface */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIPortStateMachine/linkState
func (u_ USBHostCIPortStateMachine) LinkState() USBHostCILinkState /* not a class type */ {
	rv := objc.Send[USBHostCILinkState](u_.ID, objc.Sel("linkState"))
	return rv
}/* debug [instance_properties/getter]: linkState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIPortStateMachine/overcurrent
func (u_ USBHostCIPortStateMachine) Overcurrent() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("overcurrent"))
	return rv
}/* debug [instance_properties/getter]: overcurrent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIPortStateMachine/overcurrent
func (u_ USBHostCIPortStateMachine) SetOvercurrent(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setOvercurrent:"), value)
}/* debug [instance_properties/setter]: overcurrent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIPortStateMachine/portNumber
func (u_ USBHostCIPortStateMachine) PortNumber() uint {
	rv := objc.Send[uint](u_.ID, objc.Sel("portNumber"))
	return rv
}/* debug [instance_properties/getter]: portNumber */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIPortStateMachine/portState
func (u_ USBHostCIPortStateMachine) PortState() USBHostCIPortState /* not a class type */ {
	rv := objc.Send[USBHostCIPortState](u_.ID, objc.Sel("portState"))
	return rv
}/* debug [instance_properties/getter]: portState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIPortStateMachine/portStatus
func (u_ USBHostCIPortStateMachine) PortStatus() USBHostCIPortStatus /* typedef */ {
	rv := objc.Send[uint32](u_.ID, objc.Sel("portStatus"))
	return rv
}/* debug [instance_properties/getter]: portStatus */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIPortStateMachine/powered
func (u_ USBHostCIPortStateMachine) Powered() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("powered"))
	return rv
}/* debug [instance_properties/getter]: powered */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIPortStateMachine/powered
func (u_ USBHostCIPortStateMachine) SetPowered(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setPowered:"), value)
}/* debug [instance_properties/setter]: powered */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIPortStateMachine/speed
func (u_ USBHostCIPortStateMachine) Speed() USBHostCIDeviceSpeed /* not a class type */ {
	rv := objc.Send[USBHostCIDeviceSpeed](u_.ID, objc.Sel("speed"))
	return rv
}/* debug [instance_properties/getter]: speed */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class IOUSBHostCIPortStateMachine */


