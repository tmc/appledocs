// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [VZMacOSBootLoader] class.
var (
	VZMacOSBootLoaderClass     _VZMacOSBootLoaderClass
	VZMacOSBootLoaderClassOnce sync.Once
)

func getVZMacOSBootLoaderClass() _VZMacOSBootLoaderClass {
	VZMacOSBootLoaderClassOnce.Do(func() {
		VZMacOSBootLoaderClass = _VZMacOSBootLoaderClass{objc.GetClass("VZMacOSBootLoader")}
	})
	return VZMacOSBootLoaderClass
}

type _VZMacOSBootLoaderClass struct {
	class objc.Class
}

// An interface definition for the [VZMacOSBootLoader] class.
type IVZMacOSBootLoader interface {
	IVZBootLoader
	Platform() IVZPlatformConfiguration
	SetPlatform(value IVZPlatformConfiguration)
}

// An object that loads and configures a boot loader for running macOS on Apple silicon as a guest system of your VM.
//
// You must use a in conjunction with the macOS boot loader. It’s invalid to use it with any other platform configuration.


// An object that loads and configures a boot loader for running macOS on Apple silicon as a guest system of your VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacOSBootLoader
type VZMacOSBootLoader struct {
	VZBootLoader
}

// VZMacOSBootLoaderFrom constructs a [VZMacOSBootLoader] from an unsafe.Pointer.
//
// An object that loads and configures a boot loader for running macOS on Apple silicon as a guest system of your VM.
func VZMacOSBootLoaderFrom(ptr unsafe.Pointer) VZMacOSBootLoader {
	return VZMacOSBootLoader{
		VZBootLoader: VZBootLoaderFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (vc _VZMacOSBootLoaderClass) Alloc() VZMacOSBootLoader {
	rv := objc.Send[VZMacOSBootLoader](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZMacOSBootLoaderClass) New() VZMacOSBootLoader {
	rv := objc.Send[VZMacOSBootLoader](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZMacOSBootLoader) Init() VZMacOSBootLoader {
	rv := objc.Send[VZMacOSBootLoader](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZMacOSBootLoader) Autorelease() VZMacOSBootLoader {
	rv := objc.Send[VZMacOSBootLoader](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMacOSBootLoader creates a new VZMacOSBootLoader instance.
func NewVZMacOSBootLoader() VZMacOSBootLoader {
	return getVZMacOSBootLoaderClass().New()
}



// The hardware platform to use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachineconfiguration/platform
func (v_ VZMacOSBootLoader) Platform() IVZPlatformConfiguration {
	rv := objc.Send[VZPlatformConfiguration](v_.ID, objc.Sel("platform"))
	return rv
}


// The hardware platform to use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachineconfiguration/platform
func (v_ VZMacOSBootLoader) SetPlatform(value IVZPlatformConfiguration) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setPlatform:"), value)
}



