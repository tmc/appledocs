// Code generated from Apple documentation for IOUSBHost. DO NOT EDIT.

package iousbhost

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [USBHostCIControllerStateMachine] class.
type IUSBHostCIControllerStateMachine interface {
	objectivec.IObject
	EnqueueUpdatedFrameTimestampError(frame uint64, timestamp uint64, error_ unsafe.Pointer) bool
	RespondToCommandStatusFrameTimestampError(command unsafe.Pointer, status unsafe.Pointer, frame uint64, timestamp uint64, error_ unsafe.Pointer) bool
}

//
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIControllerStateMachine
type USBHostCIControllerStateMachine struct {
	objectivec.Object
}

// USBHostCIControllerStateMachineFrom constructs a [USBHostCIControllerStateMachine] from an unsafe.Pointer.
func USBHostCIControllerStateMachineFrom(ptr unsafe.Pointer) USBHostCIControllerStateMachine {
	return USBHostCIControllerStateMachine{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (uc _USBHostCIControllerStateMachineClass) Alloc() USBHostCIControllerStateMachine {
	rv := objc.Send[USBHostCIControllerStateMachine](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


//
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIControllerStateMachine/initWithInterface:error:
func NewUSBHostCIControllerStateMachineWithInterfaceError(interface_ IOUSBHostControllerInterface, error_ unsafe.Pointer) USBHostCIControllerStateMachine {
	instance := getUSBHostCIControllerStateMachineClass().Alloc()
	rv := objc.Send[USBHostCIControllerStateMachine](instance.ID, objc.Sel("initWithInterface:error:"), interface_, error_)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIControllerStateMachine/enqueueUpdatedFrame(_:timestamp:)
func (u_ USBHostCIControllerStateMachine) EnqueueUpdatedFrameTimestampError(frame uint64, timestamp uint64, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("enqueueUpdatedFrame:timestamp:error:"), frame, timestamp, error_)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIControllerStateMachine/respond(toCommand:status:frame:timestamp:)
func (u_ USBHostCIControllerStateMachine) RespondToCommandStatusFrameTimestampError(command unsafe.Pointer, status unsafe.Pointer, frame uint64, timestamp uint64, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("respondToCommand:status:frame:timestamp:error:"), command, status, frame, timestamp, error_)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIControllerStateMachine/controllerInterface
func (u_ USBHostCIControllerStateMachine) ControllerInterface() IOUSBHostControllerInterface {
	rv := objc.Send[IOUSBHostControllerInterface](u_.ID, objc.Sel("controllerInterface"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIControllerStateMachine/controllerState
func (u_ USBHostCIControllerStateMachine) ControllerState() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("controllerState"))
	return rv
}


