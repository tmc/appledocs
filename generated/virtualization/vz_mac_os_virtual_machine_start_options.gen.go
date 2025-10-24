// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class VZMacOSVirtualMachineStartOptions */


/* debug [class_header]: Header for VZMacOSVirtualMachineStartOptions */
// The class instance for the [VZMacOSVirtualMachineStartOptions] class.
var (
	VZMacOSVirtualMachineStartOptionsClass     _VZMacOSVirtualMachineStartOptionsClass
	VZMacOSVirtualMachineStartOptionsClassOnce sync.Once
)

func getVZMacOSVirtualMachineStartOptionsClass() _VZMacOSVirtualMachineStartOptionsClass {
	VZMacOSVirtualMachineStartOptionsClassOnce.Do(func() {
		VZMacOSVirtualMachineStartOptionsClass = _VZMacOSVirtualMachineStartOptionsClass{objc.GetClass("VZMacOSVirtualMachineStartOptions")}
	})
	return VZMacOSVirtualMachineStartOptionsClass
}

type _VZMacOSVirtualMachineStartOptionsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VZMacOSVirtualMachineStartOptions */
// An interface definition for the [VZMacOSVirtualMachineStartOptions] class.
type IVZMacOSVirtualMachineStartOptions interface {
	IVZVirtualMachineStartOptions
	
/* debug [class_interface_properties]: Properties for VZMacOSVirtualMachineStartOptions */
	// properties:
	StartUpFromMacOSRecovery() bool
	SetStartUpFromMacOSRecovery(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VZMacOSVirtualMachineStartOptions */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VZMacOSVirtualMachineStartOptions */
// Alloc allocates a new instance without initialization.
func (vc _VZMacOSVirtualMachineStartOptionsClass) Alloc() VZMacOSVirtualMachineStartOptions {
	rv := objc.Send[VZMacOSVirtualMachineStartOptions](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VZMacOSVirtualMachineStartOptionsClass) New() VZMacOSVirtualMachineStartOptions {
	rv := objc.Send[VZMacOSVirtualMachineStartOptions](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZMacOSVirtualMachineStartOptions) Init() VZMacOSVirtualMachineStartOptions {
	rv := objc.Send[VZMacOSVirtualMachineStartOptions](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZMacOSVirtualMachineStartOptions) Autorelease() VZMacOSVirtualMachineStartOptions {
	rv := objc.Send[VZMacOSVirtualMachineStartOptions](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMacOSVirtualMachineStartOptions creates a new VZMacOSVirtualMachineStartOptions instance.
func NewVZMacOSVirtualMachineStartOptions() VZMacOSVirtualMachineStartOptions {
	return getVZMacOSVirtualMachineStartOptionsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VZMacOSVirtualMachineStartOptions */
// A class that describes start options for macOS VMs.


// A class that describes start options for macOS VMs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacOSVirtualMachineStartOptions
type VZMacOSVirtualMachineStartOptions struct {
	VZVirtualMachineStartOptions
}

// VZMacOSVirtualMachineStartOptionsFrom constructs a [VZMacOSVirtualMachineStartOptions] from an unsafe.Pointer.
//
// A class that describes start options for macOS VMs.
func VZMacOSVirtualMachineStartOptionsFrom(ptr unsafe.Pointer) VZMacOSVirtualMachineStartOptions {
	return VZMacOSVirtualMachineStartOptions{
		VZVirtualMachineStartOptions: VZVirtualMachineStartOptionsFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VZMacOSVirtualMachineStartOptions *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VZMacOSVirtualMachineStartOptions */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VZMacOSVirtualMachineStartOptions */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VZMacOSVirtualMachineStartOptions */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VZMacOSVirtualMachineStartOptions */

// A Boolean value that indicates whether the macOS guest should start in recovery mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacOSVirtualMachineStartOptions/startUpFromMacOSRecovery
func (v_ VZMacOSVirtualMachineStartOptions) StartUpFromMacOSRecovery() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("startUpFromMacOSRecovery"))
	return rv
}/* debug [instance_properties/getter]: startUpFromMacOSRecovery */


// A Boolean value that indicates whether the macOS guest should start in recovery mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacOSVirtualMachineStartOptions/startUpFromMacOSRecovery
func (v_ VZMacOSVirtualMachineStartOptions) SetStartUpFromMacOSRecovery(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setStartUpFromMacOSRecovery:"), value)
}/* debug [instance_properties/setter]: startUpFromMacOSRecovery */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VZMacOSVirtualMachineStartOptions */



