// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [VZUSBController] class.
var (
	VZUSBControllerClass     _VZUSBControllerClass
	VZUSBControllerClassOnce sync.Once
)

func getVZUSBControllerClass() _VZUSBControllerClass {
	VZUSBControllerClassOnce.Do(func() {
		VZUSBControllerClass = _VZUSBControllerClass{objc.GetClass("VZUSBController")}
	})
	return VZUSBControllerClass
}

type _VZUSBControllerClass struct {
	class objc.Class
}

// An interface definition for the [VZUSBController] class.
type IVZUSBController interface {
	objectivec.IObject
	// properties:
	UsbDevices() []objc.ID
	UsbControllers() IVZUSBController
	SetUsbControllers(value IVZUSBController)
	// methods:
	AttachDeviceCompletionHandler(device objectivec.IObject, completionHandler unsafe.Pointer)
	DetachDeviceCompletionHandler(device objectivec.IObject, completionHandler unsafe.Pointer)
}

// A class that represents a USB controller in a VM.
//
// Don’t create a directly. You need to first configure USB controllers on a through a subclass of . When you create a from the configuration, the USB controllers are available through the property. The concrete type of a corresponds to the type the configuration uses. For example, a leads to a device of type .


// A class that represents a USB controller in a VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZUSBController
type VZUSBController struct {
	objectivec.Object
}

// VZUSBControllerFrom constructs a [VZUSBController] from an unsafe.Pointer.
//
// A class that represents a USB controller in a VM.
func VZUSBControllerFrom(ptr unsafe.Pointer) VZUSBController {
	return VZUSBController{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (vc _VZUSBControllerClass) Alloc() VZUSBController {
	rv := objc.Send[VZUSBController](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZUSBControllerClass) New() VZUSBController {
	rv := objc.Send[VZUSBController](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZUSBController) Init() VZUSBController {
	rv := objc.Send[VZUSBController](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZUSBController) Autorelease() VZUSBController {
	rv := objc.Send[VZUSBController](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZUSBController creates a new VZUSBController instance.
func NewVZUSBController() VZUSBController {
	return getVZUSBControllerClass().New()
}



// Attaches a USB device to the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZUSBController/attach(device:completionHandler:)
func (v_ VZUSBController) AttachDeviceCompletionHandler(device objectivec.IObject, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("attachDevice:completionHandler:"), device, completionHandler)
}


// Detaches a USB device from the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZUSBController/detach(device:completionHandler:)
func (v_ VZUSBController) DetachDeviceCompletionHandler(device objectivec.IObject, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("detachDevice:completionHandler:"), device, completionHandler)
}


// The list of attached USB devices for the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZUSBController/usbDevices
func (v_ VZUSBController) UsbDevices() []objc.ID {
	rv := objc.Send[[]objc.ID](v_.ID, objc.Sel("usbDevices"))
	return rv
}


// The list of runtime USB controller objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachine/usbcontrollers
func (v_ VZUSBController) UsbControllers() IVZUSBController {
	rv := objc.Send[VZUSBController](v_.ID, objc.Sel("usbControllers"))
	return rv
}


// The list of runtime USB controller objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachine/usbcontrollers
func (v_ VZUSBController) SetUsbControllers(value IVZUSBController) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setUsbControllers:"), value)
}



