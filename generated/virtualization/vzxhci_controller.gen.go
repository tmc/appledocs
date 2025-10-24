// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class VZXHCIController */


/* debug [class_header]: Header for VZXHCIController */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VZXHCIController */
// An interface definition for the [VZXHCIController] class.
type IVZXHCIController interface {
	IVZUSBController
	
/* debug [class_interface_properties]: Properties for VZXHCIController */
	// properties:
	UsbControllers() IVZUSBControllerConfiguration
	SetUsbControllers(value IVZUSBControllerConfiguration)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VZXHCIController */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VZXHCIController */
// Alloc allocates a new instance without initialization.
func (vc _VZXHCIControllerClass) Alloc() VZXHCIController {
	rv := objc.Send[VZXHCIController](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VZXHCIController */
// A class that represents a USB Extensible Host Controller Interface (XHCI) controller in a VM.
//
// Don’t create objects directly. Instead, you create a object at runtime though the property of the object by populating it with objects.


// A class that represents a USB Extensible Host Controller Interface (XHCI) controller in a VM.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VZXHCIController *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VZXHCIController */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VZXHCIController */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VZXHCIController */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VZXHCIController */

// The list of configured USB controllers for the VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachineconfiguration/usbcontrollers
func (v_ VZXHCIController) UsbControllers() IVZUSBControllerConfiguration {
	rv := objc.Send[VZUSBControllerConfiguration](v_.ID, objc.Sel("usbControllers"))
	return rv
}/* debug [instance_properties/getter]: usbControllers */


// The list of configured USB controllers for the VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachineconfiguration/usbcontrollers
func (v_ VZXHCIController) SetUsbControllers(value IVZUSBControllerConfiguration) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setUsbControllers:"), value)
}/* debug [instance_properties/setter]: usbControllers */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VZXHCIController */



