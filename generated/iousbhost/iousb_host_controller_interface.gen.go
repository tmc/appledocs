// Code generated from Apple documentation for IOUSBHost. DO NOT EDIT.

package iousbhost

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	// properties:
	Capabilities() USBHostCIMessage /* not a class type */
	SetCapabilities(value USBHostCIMessage /* not a class type */)
	ControllerStateMachine() IOUSBHostCIControllerStateMachine /* already interface */
	SetControllerStateMachine(value IOUSBHostCIControllerStateMachine /* already interface */)
	InterruptRateHz() int /* primitive/slice/pointer. */
	SetInterruptRateHz(value int /* primitive/slice/pointer. */)
	Queue() unsafe.Pointer
	SetQueue(value unsafe.Pointer)
	Uuid() foundation.objc.IObject /* cross-framework: UUID */
	SetUuid(value foundation.objc.IObject /* cross-framework: UUID */)
	// methods:
	DescriptionForMessage(message USBHostCIMessage /* not a class type */) objc.IObject /* cross-framework: String */
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostControllerInterface/description(for:)
func (u_ USBHostControllerInterface) DescriptionForMessage(message USBHostCIMessage /* not a class type */) objc.IObject /* cross-framework: String */ {
	rv := objc.Send[String](u_.ID, objc.Sel("descriptionForMessage:"), message)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostcontrollerinterface/capabilities
func (u_ USBHostControllerInterface) Capabilities() USBHostCIMessage /* not a class type */ {
	rv := objc.Send[USBHostCIMessage](u_.ID, objc.Sel("capabilities"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostcontrollerinterface/capabilities
func (u_ USBHostControllerInterface) SetCapabilities(value USBHostCIMessage /* not a class type */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setCapabilities:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostcontrollerinterface/controllerstatemachine
func (u_ USBHostControllerInterface) ControllerStateMachine() IOUSBHostCIControllerStateMachine /* already interface */ {
	rv := objc.Send[USBHostCIControllerStateMachine](u_.ID, objc.Sel("controllerStateMachine"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostcontrollerinterface/controllerstatemachine
func (u_ USBHostControllerInterface) SetControllerStateMachine(value IOUSBHostCIControllerStateMachine /* already interface */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setControllerStateMachine:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostcontrollerinterface/interruptratehz
func (u_ USBHostControllerInterface) InterruptRateHz() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](u_.ID, objc.Sel("interruptRateHz"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostcontrollerinterface/interruptratehz
func (u_ USBHostControllerInterface) SetInterruptRateHz(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setInterruptRateHz:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostcontrollerinterface/queue
func (u_ USBHostControllerInterface) Queue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("queue"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostcontrollerinterface/queue
func (u_ USBHostControllerInterface) SetQueue(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setQueue:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostcontrollerinterface/uuid
func (u_ USBHostControllerInterface) Uuid() foundation.objc.IObject /* cross-framework: UUID */ {
	rv := objc.Send[foundation.UUID](u_.ID, objc.Sel("uuid"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostcontrollerinterface/uuid
func (u_ USBHostControllerInterface) SetUuid(value foundation.objc.IObject /* cross-framework: UUID */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setUuid:"), value)
}



