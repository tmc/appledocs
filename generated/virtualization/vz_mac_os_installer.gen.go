// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VZMacOSInstaller */

/* debug [class_header]: Header for VZMacOSInstaller */
// The class instance for the [VZMacOSInstaller] class.
var (
	VZMacOSInstallerClass     _VZMacOSInstallerClass
	VZMacOSInstallerClassOnce sync.Once
)

func getVZMacOSInstallerClass() _VZMacOSInstallerClass {
	VZMacOSInstallerClassOnce.Do(func() {
		VZMacOSInstallerClass = _VZMacOSInstallerClass{objc.GetClass("VZMacOSInstaller")}
	})
	return VZMacOSInstallerClass
}

type _VZMacOSInstallerClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for VZMacOSInstaller */
// An interface definition for the [VZMacOSInstaller] class.
type IVZMacOSInstaller interface {
	objectivec.IObject

	/* debug [class_interface_properties]: Properties for VZMacOSInstaller */
	// properties:
	Progress() foundation.Progress
	RestoreImageURL() objc.IObject /* cross-framework: NSURL */
	VirtualMachine() IVZVirtualMachine
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for VZMacOSInstaller */
	// methods:
	InstallWithCompletionHandler(completionHandler unsafe.Pointer)
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for VZMacOSInstaller */
// Alloc allocates a new instance without initialization.
func (vc _VZMacOSInstallerClass) Alloc() VZMacOSInstaller {
	rv := objc.Send[VZMacOSInstaller](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VZMacOSInstallerClass) New() VZMacOSInstaller {
	rv := objc.Send[VZMacOSInstaller](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZMacOSInstaller) Init() VZMacOSInstaller {
	rv := objc.Send[VZMacOSInstaller](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZMacOSInstaller) Autorelease() VZMacOSInstaller {
	rv := objc.Send[VZMacOSInstaller](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMacOSInstaller creates a new VZMacOSInstaller instance.
func NewVZMacOSInstaller() VZMacOSInstaller {
	return getVZMacOSInstallerClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for VZMacOSInstaller */
// An object you use to install macOS on the specified virtual machine.
//
// Initialize a object with a and a file URL that refers to a macOS restore image. The following code example shows how to use a

// An object you use to install macOS on the specified virtual machine.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacOSInstaller
type VZMacOSInstaller struct {
	objectivec.Object
}

// VZMacOSInstallerFrom constructs a [VZMacOSInstaller] from an unsafe.Pointer.
//
// An object you use to install macOS on the specified virtual machine.
func VZMacOSInstallerFrom(ptr unsafe.Pointer) VZMacOSInstaller {
	return VZMacOSInstaller{objectivec.Object{objc.ID(ptr)}}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for VZMacOSInstaller */

// Creates a macOS installer object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacOSInstaller/init(virtualMachine:restoringFromImageAt:)
func NewVZMacOSInstallerWithVirtualMachineRestoreImageURL(virtualMachine IVZVirtualMachine, restoreImageFileURL objc.IObject /* cross-framework: NSURL */) VZMacOSInstaller {
	instance := getVZMacOSInstallerClass().Alloc()
	rv := objc.Send[VZMacOSInstaller](instance.ID, objc.Sel("initWithVirtualMachine:restoreImageURL:"), virtualMachine, restoreImageFileURL)
	rv.Autorelease()
	return rv
} /* debug [class_init_methods/constructor]: NewVZMacOSInstallerWithVirtualMachineRestoreImageURL */

/* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for VZMacOSInstaller */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for VZMacOSInstaller */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for VZMacOSInstaller */

// Start installing macOS.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacOSInstaller/install()
func (v_ VZMacOSInstaller) InstallWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("installWithCompletionHandler:"), completionHandler)
} /* debug [instance_methods/method]: InstallWithCompletionHandler */

/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for VZMacOSInstaller */

// A progress object that you can use to observe or cancel an installation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacOSInstaller/progress
func (v_ VZMacOSInstaller) Progress() foundation.Progress {
	rv := objc.Send[foundation.Progress](v_.ID, objc.Sel("progress"))
	return rv
} /* debug [instance_properties/getter]: progress */

// The restore image URL used to initialize this installer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacOSInstaller/restoreImageURL
func (v_ VZMacOSInstaller) RestoreImageURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](v_.ID, objc.Sel("restoreImageURL"))
	return rv
} /* debug [instance_properties/getter]: restoreImageURL */

// The virtual machine used to initialize this installer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacOSInstaller/virtualMachine
func (v_ VZMacOSInstaller) VirtualMachine() IVZVirtualMachine {
	rv := objc.Send[VZVirtualMachine](v_.ID, objc.Sel("virtualMachine"))
	return rv
} /* debug [instance_properties/getter]: virtualMachine */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class VZMacOSInstaller */
