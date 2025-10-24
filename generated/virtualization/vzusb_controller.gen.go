// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VZUSBController */

/* debug [class_header]: Header for VZUSBController */
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

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for VZUSBController */
// An interface definition for the [VZUSBController] class.
type IVZUSBController interface {
	objectivec.IObject

	/* debug [class_interface_properties]: Properties for VZUSBController */
	// properties:
	UsbDevices() []objc.ID
	UsbControllers() IVZUSBController
	SetUsbControllers(value IVZUSBController)
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for VZUSBController */
	// methods:
	AttachDeviceCompletionHandler(device unsafe.Pointer, completionHandler unsafe.Pointer)
	DetachDeviceCompletionHandler(device unsafe.Pointer, completionHandler unsafe.Pointer)
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for VZUSBController */
// Alloc allocates a new instance without initialization.
func (vc _VZUSBControllerClass) Alloc() VZUSBController {
	rv := objc.Send[VZUSBController](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for VZUSBController */
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

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for VZUSBController */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for VZUSBController */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for VZUSBController */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for VZUSBController */

// Attaches a USB device to the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZUSBController/attach(device:completionHandler:)
func (v_ VZUSBController) AttachDeviceCompletionHandler(device unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("attachDevice:completionHandler:"), device, completionHandler)
} /* debug [instance_methods/method]: AttachDeviceCompletionHandler */

// Detaches a USB device from the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZUSBController/detach(device:completionHandler:)
func (v_ VZUSBController) DetachDeviceCompletionHandler(device unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("detachDevice:completionHandler:"), device, completionHandler)
} /* debug [instance_methods/method]: DetachDeviceCompletionHandler */

/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for VZUSBController */

// The list of attached USB devices for the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZUSBController/usbDevices
func (v_ VZUSBController) UsbDevices() []objc.ID {
	rv := objc.Send[[]objc.ID](v_.ID, objc.Sel("usbDevices"))
	return rv
} /* debug [instance_properties/getter]: usbDevices */

// The list of runtime USB controller objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachine/usbcontrollers
func (v_ VZUSBController) UsbControllers() IVZUSBController {
	rv := objc.Send[VZUSBController](v_.ID, objc.Sel("usbControllers"))
	return rv
} /* debug [instance_properties/getter]: usbControllers */

// The list of runtime USB controller objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachine/usbcontrollers
func (v_ VZUSBController) SetUsbControllers(value IVZUSBController) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setUsbControllers:"), value)
} /* debug [instance_properties/setter]: usbControllers */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class VZUSBController */
