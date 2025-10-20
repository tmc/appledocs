// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [VZEFIBootLoader] class.
var (
	VZEFIBootLoaderClass     _VZEFIBootLoaderClass
	VZEFIBootLoaderClassOnce sync.Once
)

func getVZEFIBootLoaderClass() _VZEFIBootLoaderClass {
	VZEFIBootLoaderClassOnce.Do(func() {
		VZEFIBootLoaderClass = _VZEFIBootLoaderClass{objc.GetClass("VZEFIBootLoader")}
	})
	return VZEFIBootLoaderClass
}

type _VZEFIBootLoaderClass struct {
	class objc.Class
}

// An interface definition for the [VZEFIBootLoader] class.
type IVZEFIBootLoader interface {
	IVZBootLoader
}

// The boot loader configuration the system uses to boot guest-operating systems that expect an Extensible Firmware Interface (EFI) ROM.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZEFIBootLoader
type VZEFIBootLoader struct {
	VZBootLoader
}

// VZEFIBootLoaderFrom constructs a [VZEFIBootLoader] from an unsafe.Pointer.
//
// The boot loader configuration the system uses to boot guest-operating systems that expect an Extensible Firmware Interface (EFI) ROM.
func VZEFIBootLoaderFrom(ptr unsafe.Pointer) VZEFIBootLoader {
	return VZEFIBootLoader{
		VZBootLoader: VZBootLoaderFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (vc _VZEFIBootLoaderClass) Alloc() VZEFIBootLoader {
	rv := objc.Send[VZEFIBootLoader](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZEFIBootLoaderClass) New() VZEFIBootLoader {
	rv := objc.Send[VZEFIBootLoader](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZEFIBootLoader) Init() VZEFIBootLoader {
	rv := objc.Send[VZEFIBootLoader](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZEFIBootLoader) Autorelease() VZEFIBootLoader {
	rv := objc.Send[VZEFIBootLoader](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZEFIBootLoader creates a new VZEFIBootLoader instance.
func NewVZEFIBootLoader() VZEFIBootLoader {
	return getVZEFIBootLoaderClass().New()
}



// The boot loader’s EFI variable store.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZEFIBootLoader/variableStore
func (v_ VZEFIBootLoader) VariableStore() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("variableStore"))
	return rv
}


// SetVariableStore sets the value of the variableStore property.
// The boot loader’s EFI variable store.

//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZEFIBootLoader/variableStore
func (v_ VZEFIBootLoader) SetVariableStore(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setVariableStore:"), value)
}

