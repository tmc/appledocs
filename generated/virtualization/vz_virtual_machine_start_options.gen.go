// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VZVirtualMachineStartOptions */

/* debug [class_header]: Header for VZVirtualMachineStartOptions */
// The class instance for the [VZVirtualMachineStartOptions] class.
var (
	VZVirtualMachineStartOptionsClass     _VZVirtualMachineStartOptionsClass
	VZVirtualMachineStartOptionsClassOnce sync.Once
)

func getVZVirtualMachineStartOptionsClass() _VZVirtualMachineStartOptionsClass {
	VZVirtualMachineStartOptionsClassOnce.Do(func() {
		VZVirtualMachineStartOptionsClass = _VZVirtualMachineStartOptionsClass{objc.GetClass("VZVirtualMachineStartOptions")}
	})
	return VZVirtualMachineStartOptionsClass
}

type _VZVirtualMachineStartOptionsClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for VZVirtualMachineStartOptions */
// An interface definition for the [VZVirtualMachineStartOptions] class.
type IVZVirtualMachineStartOptions interface {
	objectivec.IObject

	/* debug [class_interface_properties]: Properties for VZVirtualMachineStartOptions */
	// properties:
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for VZVirtualMachineStartOptions */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for VZVirtualMachineStartOptions */
// Alloc allocates a new instance without initialization.
func (vc _VZVirtualMachineStartOptionsClass) Alloc() VZVirtualMachineStartOptions {
	rv := objc.Send[VZVirtualMachineStartOptions](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VZVirtualMachineStartOptionsClass) New() VZVirtualMachineStartOptions {
	rv := objc.Send[VZVirtualMachineStartOptions](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZVirtualMachineStartOptions) Init() VZVirtualMachineStartOptions {
	rv := objc.Send[VZVirtualMachineStartOptions](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZVirtualMachineStartOptions) Autorelease() VZVirtualMachineStartOptions {
	rv := objc.Send[VZVirtualMachineStartOptions](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtualMachineStartOptions creates a new VZVirtualMachineStartOptions instance.
func NewVZVirtualMachineStartOptions() VZVirtualMachineStartOptions {
	return getVZVirtualMachineStartOptionsClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for VZVirtualMachineStartOptions */
// The abstract class for VM start options.

// The abstract class for VM start options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineStartOptions
type VZVirtualMachineStartOptions struct {
	objectivec.Object
}

// VZVirtualMachineStartOptionsFrom constructs a [VZVirtualMachineStartOptions] from an unsafe.Pointer.
//
// The abstract class for VM start options.
func VZVirtualMachineStartOptionsFrom(ptr unsafe.Pointer) VZVirtualMachineStartOptions {
	return VZVirtualMachineStartOptions{objectivec.Object{objc.ID(ptr)}}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for VZVirtualMachineStartOptions */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for VZVirtualMachineStartOptions */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for VZVirtualMachineStartOptions */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for VZVirtualMachineStartOptions */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for VZVirtualMachineStartOptions */
/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class VZVirtualMachineStartOptions */
