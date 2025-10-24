// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [VZMacOSInstaller] class.
type IVZMacOSInstaller interface {
	objectivec.IObject
	// properties:
	Progress() objc.IObject /* cross-framework: Progress */
	RestoreImageURL() objc.IObject /* cross-framework: NSURL */
	VirtualMachine() IVZVirtualMachine
	// methods:
	InstallWithCompletionHandler(completionHandler unsafe.Pointer)
}

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

// Alloc allocates a new instance without initialization.
func (vc _VZMacOSInstallerClass) Alloc() VZMacOSInstaller {
	rv := objc.Send[VZMacOSInstaller](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Creates a macOS installer object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacOSInstaller/init(virtualMachine:restoringFromImageAt:)
func NewVZMacOSInstallerWithVirtualMachineRestoreImageURL(virtualMachine IVZVirtualMachine, restoreImageFileURL objc.IObject /* cross-framework: NSURL */) VZMacOSInstaller {
	instance := getVZMacOSInstallerClass().Alloc()
	rv := objc.Send[VZMacOSInstaller](instance.ID, objc.Sel("initWithVirtualMachine:restoreImageURL:"), virtualMachine, restoreImageFileURL)
	rv.Autorelease()
	return rv
}



// Start installing macOS.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacOSInstaller/install()
func (v_ VZMacOSInstaller) InstallWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("installWithCompletionHandler:"), completionHandler)
}


// A progress object that you can use to observe or cancel an installation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacOSInstaller/progress
func (v_ VZMacOSInstaller) Progress() objc.IObject /* cross-framework: Progress */ {
	rv := objc.Send[foundation.Progress](v_.ID, objc.Sel("progress"))
	return rv
}


// The restore image URL used to initialize this installer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacOSInstaller/restoreImageURL
func (v_ VZMacOSInstaller) RestoreImageURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](v_.ID, objc.Sel("restoreImageURL"))
	return rv
}


// The virtual machine used to initialize this installer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacOSInstaller/virtualMachine
func (v_ VZMacOSInstaller) VirtualMachine() IVZVirtualMachine {
	rv := objc.Send[VZVirtualMachine](v_.ID, objc.Sel("virtualMachine"))
	return rv
}


