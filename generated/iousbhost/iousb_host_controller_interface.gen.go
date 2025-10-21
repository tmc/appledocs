// Code generated from Apple documentation for IOUSBHost. DO NOT EDIT.

package iousbhost

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [USBHostControllerInterface] class.
type IUSBHostControllerInterface interface {
	objectivec.IObject
	DescriptionForMessage(message unsafe.Pointer) string
	Destroy()
	EnqueueInterruptsCountExpediteError(interrupts unsafe.Pointer, count uint, expedite bool, error_ unsafe.Pointer) bool
}

//
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostControllerInterface
type USBHostControllerInterface struct {
	objectivec.Object
}

// USBHostControllerInterfaceFrom constructs a [USBHostControllerInterface] from an unsafe.Pointer.
func USBHostControllerInterfaceFrom(ptr unsafe.Pointer) USBHostControllerInterface {
	return USBHostControllerInterface{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (uc _USBHostControllerInterfaceClass) Alloc() USBHostControllerInterface {
	rv := objc.Send[USBHostControllerInterface](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


//
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostControllerInterface/description(for:)
func (u_ USBHostControllerInterface) DescriptionForMessage(message unsafe.Pointer) string {
	rv := objc.Send[string](u_.ID, objc.Sel("descriptionForMessage:"), message)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostControllerInterface/destroy()
func (u_ USBHostControllerInterface) Destroy() {
	objc.Send[objc.ID](u_.ID, objc.Sel("destroy"))
}

//
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostControllerInterface/enqueueInterrupts(_:count:expedite:)
func (u_ USBHostControllerInterface) EnqueueInterruptsCountExpediteError(interrupts unsafe.Pointer, count uint, expedite bool, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("enqueueInterrupts:count:expedite:error:"), interrupts, count, expedite, error_)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostcontrollerinterface/capabilities
func (u_ USBHostControllerInterface) Capabilities() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("capabilities"))
	return rv
}


// SetCapabilities sets the value of the capabilities property.
//
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostcontrollerinterface/capabilities
func (u_ USBHostControllerInterface) SetCapabilities(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setCapabilities:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostcontrollerinterface/controllerstatemachine
func (u_ USBHostControllerInterface) ControllerStateMachine() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("controllerStateMachine"))
	return rv
}


// SetControllerStateMachine sets the value of the controllerStateMachine property.
//
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostcontrollerinterface/controllerstatemachine
func (u_ USBHostControllerInterface) SetControllerStateMachine(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setControllerStateMachine:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostcontrollerinterface/interruptratehz
func (u_ USBHostControllerInterface) InterruptRateHz() int {
	rv := objc.Send[int](u_.ID, objc.Sel("interruptRateHz"))
	return rv
}


// SetInterruptRateHz sets the value of the interruptRateHz property.
//
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostcontrollerinterface/interruptratehz
func (u_ USBHostControllerInterface) SetInterruptRateHz(value int) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setInterruptRateHz:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostcontrollerinterface/queue
func (u_ USBHostControllerInterface) Queue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("queue"))
	return rv
}


// SetQueue sets the value of the queue property.
//
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostcontrollerinterface/queue
func (u_ USBHostControllerInterface) SetQueue(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setQueue:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostcontrollerinterface/uuid
func (u_ USBHostControllerInterface) Uuid() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("uuid"))
	return rv
}


// SetUuid sets the value of the uuid property.
//
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostcontrollerinterface/uuid
func (u_ USBHostControllerInterface) SetUuid(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setUuid:"), value)
}



