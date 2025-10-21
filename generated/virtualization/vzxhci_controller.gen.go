// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [VZXHCIController] class.
var (
	VZXHCIControllerClass     _VZXHCIControllerClass
	VZXHCIControllerClassOnce sync.Once
)

func getVZXHCIControllerClass() _VZXHCIControllerClass {
	VZXHCIControllerClassOnce.Do(func() {
		VZXHCIControllerClass = _VZXHCIControllerClass{objc.GetClass("VZXHCIController")}
	})
	return VZXHCIControllerClass
}

type _VZXHCIControllerClass struct {
	class objc.Class
}

// An interface definition for the [VZXHCIController] class.
type IVZXHCIController interface {
	IVZUSBController
}

// A class that represents a USB Extensible Host Controller Interface (XHCI) controller in a VM.
//
// Don’t create objects directly. Instead, you create a object at runtime though the property of the object by populating it with objects.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZXHCIController
type VZXHCIController struct {
	VZUSBController
}

// VZXHCIControllerFrom constructs a [VZXHCIController] from an unsafe.Pointer.
//
// A class that represents a USB Extensible Host Controller Interface (XHCI) controller in a VM.
func VZXHCIControllerFrom(ptr unsafe.Pointer) VZXHCIController {
	return VZXHCIController{
		VZUSBController: VZUSBControllerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (vc _VZXHCIControllerClass) Alloc() VZXHCIController {
	rv := objc.Send[VZXHCIController](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZXHCIControllerClass) New() VZXHCIController {
	rv := objc.Send[VZXHCIController](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZXHCIController) Init() VZXHCIController {
	rv := objc.Send[VZXHCIController](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZXHCIController) Autorelease() VZXHCIController {
	rv := objc.Send[VZXHCIController](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZXHCIController creates a new VZXHCIController instance.
func NewVZXHCIController() VZXHCIController {
	return getVZXHCIControllerClass().New()
}


// The list of configured USB controllers for the VM.
//
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachineconfiguration/usbcontrollers
func (v_ VZXHCIController) UsbControllers() VZUSBControllerConfiguration {
	rv := objc.Send[VZUSBControllerConfiguration](v_.ID, objc.Sel("usbControllers"))
	return rv
}


// SetUsbControllers sets the value of the usbControllers property.
// The list of configured USB controllers for the VM.

//
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachineconfiguration/usbcontrollers
func (v_ VZXHCIController) SetUsbControllers(value IVZUSBControllerConfiguration) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setUsbControllers:"), value)
}



