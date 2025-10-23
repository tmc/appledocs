// Code generated from Apple documentation for IOUSBHost. DO NOT EDIT.

package iousbhost

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [USBHostCIPortStateMachine] class.
type IUSBHostCIPortStateMachine interface {
	objectivec.IObject
	RespondToCommandStatusError(command unsafe.Pointer, status unsafe.Pointer, error_ unsafe.Pointer) bool
	UpdateLinkStateSpeedInhibitLinkStateChangeError(linkState unsafe.Pointer, speed unsafe.Pointer, inhibitLinkStateChange bool, error_ unsafe.Pointer) bool
	Overcurrent() bool
	SetOvercurrent(value bool)
	PortStatus() unsafe.Pointer
	Powered() bool
	SetPowered(value bool)
	Speed() unsafe.Pointer
	Connected() bool
	SetConnected(value bool)
	ControllerInterface() IOUSBHostControllerInterface
	SetControllerInterface(value IOUSBHostControllerInterface)
	LinkState() unsafe.Pointer
	SetLinkState(value unsafe.Pointer)
	PortNumber() int
	SetPortNumber(value int)
	PortState() unsafe.Pointer
	SetPortState(value unsafe.Pointer)
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIPortStateMachine
type USBHostCIPortStateMachine struct {
	objectivec.Object
}

// USBHostCIPortStateMachineFrom constructs a [USBHostCIPortStateMachine] from an unsafe.Pointer.
func USBHostCIPortStateMachineFrom(ptr unsafe.Pointer) USBHostCIPortStateMachine {
	return USBHostCIPortStateMachine{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (uc _USBHostCIPortStateMachineClass) Alloc() USBHostCIPortStateMachine {
	rv := objc.Send[USBHostCIPortStateMachine](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIPortStateMachine/respond(toCommand:status:)
func (u_ USBHostCIPortStateMachine) RespondToCommandStatusError(command unsafe.Pointer, status unsafe.Pointer, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("respondToCommand:status:error:"), command, status, error_)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIPortStateMachine/updateLinkState(_:speed:inhibitLinkStateChange:)
func (u_ USBHostCIPortStateMachine) UpdateLinkStateSpeedInhibitLinkStateChangeError(linkState unsafe.Pointer, speed unsafe.Pointer, inhibitLinkStateChange bool, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("updateLinkState:speed:inhibitLinkStateChange:error:"), linkState, speed, inhibitLinkStateChange, error_)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIPortStateMachine/overcurrent
func (u_ USBHostCIPortStateMachine) Overcurrent() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("overcurrent"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIPortStateMachine/overcurrent
func (u_ USBHostCIPortStateMachine) SetOvercurrent(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setOvercurrent:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIPortStateMachine/portStatus
func (u_ USBHostCIPortStateMachine) PortStatus() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("portStatus"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIPortStateMachine/powered
func (u_ USBHostCIPortStateMachine) Powered() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("powered"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIPortStateMachine/powered
func (u_ USBHostCIPortStateMachine) SetPowered(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setPowered:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIPortStateMachine/speed
func (u_ USBHostCIPortStateMachine) Speed() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("speed"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostciportstatemachine/connected
func (u_ USBHostCIPortStateMachine) Connected() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("connected"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostciportstatemachine/connected
func (u_ USBHostCIPortStateMachine) SetConnected(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setConnected:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostciportstatemachine/controllerinterface
func (u_ USBHostCIPortStateMachine) ControllerInterface() IOUSBHostControllerInterface {
	rv := objc.Send[IOUSBHostControllerInterface](u_.ID, objc.Sel("controllerInterface"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostciportstatemachine/controllerinterface
func (u_ USBHostCIPortStateMachine) SetControllerInterface(value IOUSBHostControllerInterface) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setControllerInterface:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostciportstatemachine/linkstate
func (u_ USBHostCIPortStateMachine) LinkState() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("linkState"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostciportstatemachine/linkstate
func (u_ USBHostCIPortStateMachine) SetLinkState(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setLinkState:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostciportstatemachine/portnumber
func (u_ USBHostCIPortStateMachine) PortNumber() int {
	rv := objc.Send[int](u_.ID, objc.Sel("portNumber"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostciportstatemachine/portnumber
func (u_ USBHostCIPortStateMachine) SetPortNumber(value int) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setPortNumber:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostciportstatemachine/portstate
func (u_ USBHostCIPortStateMachine) PortState() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("portState"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostciportstatemachine/portstate
func (u_ USBHostCIPortStateMachine) SetPortState(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setPortState:"), value)
}



