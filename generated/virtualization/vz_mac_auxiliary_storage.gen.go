// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [VZMacAuxiliaryStorage] class.
var (
	VZMacAuxiliaryStorageClass     _VZMacAuxiliaryStorageClass
	VZMacAuxiliaryStorageClassOnce sync.Once
)

func getVZMacAuxiliaryStorageClass() _VZMacAuxiliaryStorageClass {
	VZMacAuxiliaryStorageClassOnce.Do(func() {
		VZMacAuxiliaryStorageClass = _VZMacAuxiliaryStorageClass{objc.GetClass("VZMacAuxiliaryStorage")}
	})
	return VZMacAuxiliaryStorageClass
}

type _VZMacAuxiliaryStorageClass struct {
	class objc.Class
}

// An interface definition for the [VZMacAuxiliaryStorage] class.
type IVZMacAuxiliaryStorage interface {
	objectivec.IObject
}

// An object that contains information the boot loader needs for booting macOS as a guest operating system.
//
// The Mac auxiliary storage contains data used by the boot loader and the guest operating system. It’s necessary to boot a macOS guest OS. When creating a new VM, use to create a default initialized auxiliary storage. The hardware model you use when creating the new auxiliary storage depends on the restore image that you’ll use for installation. From the restore image, use to get a supported configuration. A configuration has a associated with it. After initializing the new auxiliary storage, set it on . . The hardware model in . must be identical to the one used to create the empty auxiliary storage., otherwise the behavior isn’t defined. When installing macOS, the lays out data on the auxiliary storage. After installation, the macOS guest uses the auxiliary storage for every subsequent boot. When moving or performing a backup of a VM, you must move or copy the file containing the auxiliary storage along with the main disk image. To boot a VM created with , use to set up the auxiliary storage from the existing file used during installation. When using an existing file, the hardware model of the . must match the hardware model used when creating the original file.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacAuxiliaryStorage
type VZMacAuxiliaryStorage struct {
	objectivec.Object
}

// VZMacAuxiliaryStorageFrom constructs a [VZMacAuxiliaryStorage] from an unsafe.Pointer.
//
// An object that contains information the boot loader needs for booting macOS as a guest operating system.
func VZMacAuxiliaryStorageFrom(ptr unsafe.Pointer) VZMacAuxiliaryStorage {
	return VZMacAuxiliaryStorage{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (vc _VZMacAuxiliaryStorageClass) Alloc() VZMacAuxiliaryStorage {
	rv := objc.Send[VZMacAuxiliaryStorage](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZMacAuxiliaryStorageClass) New() VZMacAuxiliaryStorage {
	rv := objc.Send[VZMacAuxiliaryStorage](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZMacAuxiliaryStorage) Init() VZMacAuxiliaryStorage {
	rv := objc.Send[VZMacAuxiliaryStorage](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZMacAuxiliaryStorage) Autorelease() VZMacAuxiliaryStorage {
	rv := objc.Send[VZMacAuxiliaryStorage](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMacAuxiliaryStorage creates a new VZMacAuxiliaryStorage instance.
func NewVZMacAuxiliaryStorage() VZMacAuxiliaryStorage {
	return getVZMacAuxiliaryStorageClass().New()
}


// Initializes an auxiliary storage object with data from the location at the URL you provide.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacAuxiliaryStorage/init(contentsOfURL:)
func NewVZMacAuxiliaryStorageWithContentsOfURL(URL unsafe.Pointer) VZMacAuxiliaryStorage {
	instance := getVZMacAuxiliaryStorageClass().Alloc()
	rv := objc.Send[VZMacAuxiliaryStorage](instance.ID, objc.Sel("initWithContentsOfURL:"), URL)
	rv.Autorelease()
	return rv
}

// Creates an initialized Mac auxiliary storage instance that describes a specific hardware model at a URL you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacAuxiliaryStorage/init(creatingStorageAt:hardwareModel:options:)
func NewVZMacAuxiliaryStorageCreatingStorageAtURLHardwareModelOptionsError(URL unsafe.Pointer, hardwareModel unsafe.Pointer, options unsafe.Pointer, error_ unsafe.Pointer) VZMacAuxiliaryStorage {
	instance := getVZMacAuxiliaryStorageClass().Alloc()
	rv := objc.Send[VZMacAuxiliaryStorage](instance.ID, objc.Sel("initCreatingStorageAtURL:hardwareModel:options:error:"), URL, hardwareModel, options, error_)
	rv.Autorelease()
	return rv
}

// Initializes an auxiliary storage object with data from the location at the URL you provide.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacAuxiliaryStorage/init(url:)
func NewVZMacAuxiliaryStorageWithURL(URL unsafe.Pointer) VZMacAuxiliaryStorage {
	instance := getVZMacAuxiliaryStorageClass().Alloc()
	rv := objc.Send[VZMacAuxiliaryStorage](instance.ID, objc.Sel("initWithURL:"), URL)
	rv.Autorelease()
	return rv
}


// The URL of the auxiliary storage on the local file system.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacAuxiliaryStorage/url
func (v_ VZMacAuxiliaryStorage) URL() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("URL"))
	return rv
}


