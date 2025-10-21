// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [VZBootLoader] class.
var (
	VZBootLoaderClass     _VZBootLoaderClass
	VZBootLoaderClassOnce sync.Once
)

func getVZBootLoaderClass() _VZBootLoaderClass {
	VZBootLoaderClassOnce.Do(func() {
		VZBootLoaderClass = _VZBootLoaderClass{objc.GetClass("VZBootLoader")}
	})
	return VZBootLoaderClass
}

type _VZBootLoaderClass struct {
	class objc.Class
}

// An interface definition for the [VZBootLoader] class.
type IVZBootLoader interface {
	objectivec.IObject
}

// The base class that defines the management of the initial process of the guest system.
//
// The abstract class defines the common behaviors for booting a guest operating system into a VM. Don’t create instances of this class directly. Instead, instantiate the subclass that corresponds to the type of operating system you plan to load. For example, to create a boot loader object for a Linux kernel, create a object; to create a boot loader object for installation using an ISO image create a . For a macOS system create .
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZBootLoader
type VZBootLoader struct {
	objectivec.Object
}

// VZBootLoaderFrom constructs a [VZBootLoader] from an unsafe.Pointer.
//
// The base class that defines the management of the initial process of the guest system.
func VZBootLoaderFrom(ptr unsafe.Pointer) VZBootLoader {
	return VZBootLoader{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (vc _VZBootLoaderClass) Alloc() VZBootLoader {
	rv := objc.Send[VZBootLoader](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZBootLoaderClass) New() VZBootLoader {
	rv := objc.Send[VZBootLoader](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZBootLoader) Init() VZBootLoader {
	rv := objc.Send[VZBootLoader](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZBootLoader) Autorelease() VZBootLoader {
	rv := objc.Send[VZBootLoader](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZBootLoader creates a new VZBootLoader instance.
func NewVZBootLoader() VZBootLoader {
	return getVZBootLoaderClass().New()
}




