// Code generated from Apple documentation for IOUSBHost. DO NOT EDIT.

package iousbhost

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [USBHostCIDeviceStateMachine] class.
type IUSBHostCIDeviceStateMachine interface {
	objectivec.IObject
	InspectCommandError(command unsafe.Pointer, error_ unsafe.Pointer) bool
	RespondToCommandStatusError(command unsafe.Pointer, status unsafe.Pointer, error_ unsafe.Pointer) bool
}

//
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIDeviceStateMachine
type USBHostCIDeviceStateMachine struct {
	objectivec.Object
}

// USBHostCIDeviceStateMachineFrom constructs a [USBHostCIDeviceStateMachine] from an unsafe.Pointer.
func USBHostCIDeviceStateMachineFrom(ptr unsafe.Pointer) USBHostCIDeviceStateMachine {
	return USBHostCIDeviceStateMachine{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (uc _USBHostCIDeviceStateMachineClass) Alloc() USBHostCIDeviceStateMachine {
	rv := objc.Send[USBHostCIDeviceStateMachine](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


//
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIDeviceStateMachine/inspectCommand(_:)
func (u_ USBHostCIDeviceStateMachine) InspectCommandError(command unsafe.Pointer, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("inspectCommand:error:"), command, error_)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIDeviceStateMachine/respond(toCommand:status:)
func (u_ USBHostCIDeviceStateMachine) RespondToCommandStatusError(command unsafe.Pointer, status unsafe.Pointer, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("respondToCommand:status:error:"), command, status, error_)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIDeviceStateMachine/controllerInterface
func (u_ USBHostCIDeviceStateMachine) ControllerInterface() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("controllerInterface"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIDeviceStateMachine/deviceAddress
func (u_ USBHostCIDeviceStateMachine) DeviceAddress() uint {
	rv := objc.Send[uint](u_.ID, objc.Sel("deviceAddress"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIDeviceStateMachine/deviceState
func (u_ USBHostCIDeviceStateMachine) DeviceState() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("deviceState"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostcidevicestatemachine/completeroute
func (u_ USBHostCIDeviceStateMachine) CompleteRoute() int {
	rv := objc.Send[int](u_.ID, objc.Sel("completeRoute"))
	return rv
}


// SetCompleteRoute sets the value of the completeRoute property.
//
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostcidevicestatemachine/completeroute
func (u_ USBHostCIDeviceStateMachine) SetCompleteRoute(value int) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setCompleteRoute:"), value)
}



