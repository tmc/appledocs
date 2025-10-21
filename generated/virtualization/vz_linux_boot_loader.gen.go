// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [VZLinuxBootLoader] class.
type IVZLinuxBootLoader interface {
	IVZBootLoader
}

// An object that loads and configures a Linux kernel as the guest system of your VM.
//
// Create and configure a object during the initial configuration of your VM. Use this object to specify the location of the Linux kernel that serves as the guest operating system. You can also specify additional information to use during the boot process, such as command-line parameters to pass to the kernel. Assign the object to the property of your object.  A configuration with   is only valid if used with  .
//
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

// Alloc allocates a new instance without initialization.
func (vc _VZLinuxBootLoaderClass) Alloc() VZLinuxBootLoader {
	rv := objc.Send[VZLinuxBootLoader](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Creates a boot loader that launches the Linux kernel at the specified URL.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZLinuxBootLoader/init(kernelURL:)
func NewVZLinuxBootLoaderWithKernelURL(kernelURL unsafe.Pointer) VZLinuxBootLoader {
	instance := getVZLinuxBootLoaderClass().Alloc()
	rv := objc.Send[VZLinuxBootLoader](instance.ID, objc.Sel("initWithKernelURL:"), kernelURL)
	rv.Autorelease()
	return rv
}


// The command-line parameters to pass to the Linux kernel at boot time.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZLinuxBootLoader/commandLine
func (v_ VZLinuxBootLoader) CommandLine() string {
	rv := objc.Send[string](v_.ID, objc.Sel("commandLine"))
	return rv
}


// SetCommandLine sets the value of the commandLine property.
// The command-line parameters to pass to the Linux kernel at boot time.

//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZLinuxBootLoader/commandLine
func (v_ VZLinuxBootLoader) SetCommandLine(value string) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setCommandLine:"), objc.String(value))
}
// The location of an optional RAM disk, which the boot loader maps into memory before it boots the Linux kernel.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZLinuxBootLoader/initialRamdiskURL
func (v_ VZLinuxBootLoader) InitialRamdiskURL() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("initialRamdiskURL"))
	return rv
}


// SetInitialRamdiskURL sets the value of the initialRamdiskURL property.
// The location of an optional RAM disk, which the boot loader maps into memory before it boots the Linux kernel.

//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZLinuxBootLoader/initialRamdiskURL
func (v_ VZLinuxBootLoader) SetInitialRamdiskURL(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setInitialRamdiskURL:"), value)
}
// The URL of the Linux kernel file.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZLinuxBootLoader/kernelURL
func (v_ VZLinuxBootLoader) KernelURL() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("kernelURL"))
	return rv
}


// SetKernelURL sets the value of the kernelURL property.
// The URL of the Linux kernel file.

//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZLinuxBootLoader/kernelURL
func (v_ VZLinuxBootLoader) SetKernelURL(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setKernelURL:"), value)
}

