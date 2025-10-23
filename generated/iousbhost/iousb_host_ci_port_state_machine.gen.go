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
	// properties:
	Speed() USBHostCIDeviceSpeed /* not a class type */
	Connected() bool /* primitive/slice/pointer. */
	SetConnected(value bool /* primitive/slice/pointer. */)
	ControllerInterface() IOUSBHostControllerInterface /* already interface */
	SetControllerInterface(value IOUSBHostControllerInterface /* already interface */)
	LinkState() USBHostCILinkState /* not a class type */
	SetLinkState(value USBHostCILinkState /* not a class type */)
	Overcurrent() bool /* primitive/slice/pointer. */
	SetOvercurrent(value bool /* primitive/slice/pointer. */)
	PortNumber() int /* primitive/slice/pointer. */
	SetPortNumber(value int /* primitive/slice/pointer. */)
	PortState() USBHostCIPortState /* not a class type */
	SetPortState(value USBHostCIPortState /* not a class type */)
	PortStatus() USBHostCIPortStatus /* not a class type */
	SetPortStatus(value USBHostCIPortStatus /* not a class type */)
	Powered() bool /* primitive/slice/pointer. */
	SetPowered(value bool /* primitive/slice/pointer. */)
	// methods:
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
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIPortStateMachine/speed
func (u_ USBHostCIPortStateMachine) Speed() USBHostCIDeviceSpeed /* not a class type */ {
	rv := objc.Send[USBHostCIDeviceSpeed](u_.ID, objc.Sel("speed"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostciportstatemachine/connected
func (u_ USBHostCIPortStateMachine) Connected() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](u_.ID, objc.Sel("connected"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostciportstatemachine/connected
func (u_ USBHostCIPortStateMachine) SetConnected(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setConnected:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostciportstatemachine/controllerinterface
func (u_ USBHostCIPortStateMachine) ControllerInterface() IOUSBHostControllerInterface /* already interface */ {
	rv := objc.Send[USBHostControllerInterface](u_.ID, objc.Sel("controllerInterface"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostciportstatemachine/controllerinterface
func (u_ USBHostCIPortStateMachine) SetControllerInterface(value IOUSBHostControllerInterface /* already interface */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setControllerInterface:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostciportstatemachine/linkstate
func (u_ USBHostCIPortStateMachine) LinkState() USBHostCILinkState /* not a class type */ {
	rv := objc.Send[USBHostCILinkState](u_.ID, objc.Sel("linkState"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostciportstatemachine/linkstate
func (u_ USBHostCIPortStateMachine) SetLinkState(value USBHostCILinkState /* not a class type */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setLinkState:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostciportstatemachine/overcurrent
func (u_ USBHostCIPortStateMachine) Overcurrent() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](u_.ID, objc.Sel("overcurrent"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostciportstatemachine/overcurrent
func (u_ USBHostCIPortStateMachine) SetOvercurrent(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setOvercurrent:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostciportstatemachine/portnumber
func (u_ USBHostCIPortStateMachine) PortNumber() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](u_.ID, objc.Sel("portNumber"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostciportstatemachine/portnumber
func (u_ USBHostCIPortStateMachine) SetPortNumber(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setPortNumber:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostciportstatemachine/portstate
func (u_ USBHostCIPortStateMachine) PortState() USBHostCIPortState /* not a class type */ {
	rv := objc.Send[USBHostCIPortState](u_.ID, objc.Sel("portState"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostciportstatemachine/portstate
func (u_ USBHostCIPortStateMachine) SetPortState(value USBHostCIPortState /* not a class type */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setPortState:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostciportstatemachine/portstatus
func (u_ USBHostCIPortStateMachine) PortStatus() USBHostCIPortStatus /* not a class type */ {
	rv := objc.Send[USBHostCIPortStatus](u_.ID, objc.Sel("portStatus"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostciportstatemachine/portstatus
func (u_ USBHostCIPortStateMachine) SetPortStatus(value USBHostCIPortStatus /* not a class type */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setPortStatus:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostciportstatemachine/powered
func (u_ USBHostCIPortStateMachine) Powered() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](u_.ID, objc.Sel("powered"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostciportstatemachine/powered
func (u_ USBHostCIPortStateMachine) SetPowered(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setPowered:"), value)
}



