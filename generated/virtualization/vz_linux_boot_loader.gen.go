// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class VZLinuxBootLoader */

/* debug [class_header]: Header for VZLinuxBootLoader */
// The class instance for the [VZLinuxBootLoader] class.
var (
	VZLinuxBootLoaderClass     _VZLinuxBootLoaderClass
	VZLinuxBootLoaderClassOnce sync.Once
)

func getVZLinuxBootLoaderClass() _VZLinuxBootLoaderClass {
	VZLinuxBootLoaderClassOnce.Do(func() {
		VZLinuxBootLoaderClass = _VZLinuxBootLoaderClass{objc.GetClass("VZLinuxBootLoader")}
	})
	return VZLinuxBootLoaderClass
}

type _VZLinuxBootLoaderClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for VZLinuxBootLoader */
// An interface definition for the [VZLinuxBootLoader] class.
type IVZLinuxBootLoader interface {
	IVZBootLoader

	/* debug [class_interface_properties]: Properties for VZLinuxBootLoader */
	// properties:
	CommandLine() objc.IObject /* cross-framework: NSString */
	SetCommandLine(value objc.IObject /* cross-framework: NSString */)
	InitialRamdiskURL() objc.IObject /* cross-framework: NSURL */
	SetInitialRamdiskURL(value objc.IObject /* cross-framework: NSURL */)
	KernelURL() objc.IObject /* cross-framework: NSURL */
	SetKernelURL(value objc.IObject /* cross-framework: NSURL */)
	BootLoader() IVZBootLoader
	SetBootLoader(value IVZBootLoader)
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for VZLinuxBootLoader */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for VZLinuxBootLoader */
// Alloc allocates a new instance without initialization.
func (vc _VZLinuxBootLoaderClass) Alloc() VZLinuxBootLoader {
	rv := objc.Send[VZLinuxBootLoader](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VZLinuxBootLoaderClass) New() VZLinuxBootLoader {
	rv := objc.Send[VZLinuxBootLoader](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZLinuxBootLoader) Init() VZLinuxBootLoader {
	rv := objc.Send[VZLinuxBootLoader](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZLinuxBootLoader) Autorelease() VZLinuxBootLoader {
	rv := objc.Send[VZLinuxBootLoader](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZLinuxBootLoader creates a new VZLinuxBootLoader instance.
func NewVZLinuxBootLoader() VZLinuxBootLoader {
	return getVZLinuxBootLoaderClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for VZLinuxBootLoader */
// An object that loads and configures a Linux kernel as the guest system of your VM.
//
// Create and configure a object during the initial configuration of your VM. Use this object to specify the location of the Linux kernel that serves as the guest operating system. You can also specify additional information to use during the boot process, such as command-line parameters to pass to the kernel. Assign the object to the property of your object.  A configuration with   is only valid if used with  .

// An object that loads and configures a Linux kernel as the guest system of your VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZLinuxBootLoader
type VZLinuxBootLoader struct {
	VZBootLoader
}

// VZLinuxBootLoaderFrom constructs a [VZLinuxBootLoader] from an unsafe.Pointer.
//
// An object that loads and configures a Linux kernel as the guest system of your VM.
func VZLinuxBootLoaderFrom(ptr unsafe.Pointer) VZLinuxBootLoader {
	return VZLinuxBootLoader{
		VZBootLoader: VZBootLoaderFrom(ptr),
	}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for VZLinuxBootLoader */

// Creates a boot loader that launches the Linux kernel at the specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZLinuxBootLoader/init(kernelURL:)
func NewVZLinuxBootLoaderWithKernelURL(kernelURL objc.IObject /* cross-framework: NSURL */) VZLinuxBootLoader {
	instance := getVZLinuxBootLoaderClass().Alloc()
	rv := objc.Send[VZLinuxBootLoader](instance.ID, objc.Sel("initWithKernelURL:"), kernelURL)
	rv.Autorelease()
	return rv
} /* debug [class_init_methods/constructor]: NewVZLinuxBootLoaderWithKernelURL */

/* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for VZLinuxBootLoader */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for VZLinuxBootLoader */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for VZLinuxBootLoader */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for VZLinuxBootLoader */

// The command-line parameters to pass to the Linux kernel at boot time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZLinuxBootLoader/commandLine
func (v_ VZLinuxBootLoader) CommandLine() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](v_.ID, objc.Sel("commandLine"))
	return rv
} /* debug [instance_properties/getter]: commandLine */

// The command-line parameters to pass to the Linux kernel at boot time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZLinuxBootLoader/commandLine
func (v_ VZLinuxBootLoader) SetCommandLine(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setCommandLine:"), value)
} /* debug [instance_properties/setter]: commandLine */

// The location of an optional RAM disk, which the boot loader maps into memory before it boots the Linux kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZLinuxBootLoader/initialRamdiskURL
func (v_ VZLinuxBootLoader) InitialRamdiskURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](v_.ID, objc.Sel("initialRamdiskURL"))
	return rv
} /* debug [instance_properties/getter]: initialRamdiskURL */

// The location of an optional RAM disk, which the boot loader maps into memory before it boots the Linux kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZLinuxBootLoader/initialRamdiskURL
func (v_ VZLinuxBootLoader) SetInitialRamdiskURL(value objc.IObject /* cross-framework: NSURL */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setInitialRamdiskURL:"), value)
} /* debug [instance_properties/setter]: initialRamdiskURL */

// The URL of the Linux kernel file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZLinuxBootLoader/kernelURL
func (v_ VZLinuxBootLoader) KernelURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](v_.ID, objc.Sel("kernelURL"))
	return rv
} /* debug [instance_properties/getter]: kernelURL */

// The URL of the Linux kernel file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZLinuxBootLoader/kernelURL
func (v_ VZLinuxBootLoader) SetKernelURL(value objc.IObject /* cross-framework: NSURL */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setKernelURL:"), value)
} /* debug [instance_properties/setter]: kernelURL */

// The guest system to boot when the VM starts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachineconfiguration/bootloader
func (v_ VZLinuxBootLoader) BootLoader() IVZBootLoader {
	rv := objc.Send[VZBootLoader](v_.ID, objc.Sel("bootLoader"))
	return rv
} /* debug [instance_properties/getter]: bootLoader */

// The guest system to boot when the VM starts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachineconfiguration/bootloader
func (v_ VZLinuxBootLoader) SetBootLoader(value IVZBootLoader) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setBootLoader:"), value)
} /* debug [instance_properties/setter]: bootLoader */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class VZLinuxBootLoader */
