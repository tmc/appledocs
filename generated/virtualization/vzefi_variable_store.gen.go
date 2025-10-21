// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [VZEFIVariableStore] class.
var (
	VZEFIVariableStoreClass     _VZEFIVariableStoreClass
	VZEFIVariableStoreClassOnce sync.Once
)

func getVZEFIVariableStoreClass() _VZEFIVariableStoreClass {
	VZEFIVariableStoreClassOnce.Do(func() {
		VZEFIVariableStoreClass = _VZEFIVariableStoreClass{objc.GetClass("VZEFIVariableStore")}
	})
	return VZEFIVariableStoreClass
}

type _VZEFIVariableStoreClass struct {
	class objc.Class
}

// An interface definition for the [VZEFIVariableStore] class.
type IVZEFIVariableStore interface {
	objectivec.IObject
}

// An object that represents the Extensible Firmware Interface (EFI) variable store that contains NVRAM variables the EFI exposes.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZEFIVariableStore
type VZEFIVariableStore struct {
	objectivec.Object
}

// VZEFIVariableStoreFrom constructs a [VZEFIVariableStore] from an unsafe.Pointer.
//
// An object that represents the Extensible Firmware Interface (EFI) variable store that contains NVRAM variables the EFI exposes.
func VZEFIVariableStoreFrom(ptr unsafe.Pointer) VZEFIVariableStore {
	return VZEFIVariableStore{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (vc _VZEFIVariableStoreClass) Alloc() VZEFIVariableStore {
	rv := objc.Send[VZEFIVariableStore](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZEFIVariableStoreClass) New() VZEFIVariableStore {
	rv := objc.Send[VZEFIVariableStore](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZEFIVariableStore) Init() VZEFIVariableStore {
	rv := objc.Send[VZEFIVariableStore](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZEFIVariableStore) Autorelease() VZEFIVariableStore {
	rv := objc.Send[VZEFIVariableStore](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZEFIVariableStore creates a new VZEFIVariableStore instance.
func NewVZEFIVariableStore() VZEFIVariableStore {
	return getVZEFIVariableStoreClass().New()
}




// Creates a new EFI variable store at specified the URL on the filesystem, initialization options, and error-return variable.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZEFIVariableStore/init(creatingVariableStoreAt:options:)
func NewVZEFIVariableStoreCreatingVariableStoreAtURLOptionsError(URL unsafe.Pointer, options unsafe.Pointer, error_ unsafe.Pointer) VZEFIVariableStore {
	instance := getVZEFIVariableStoreClass().Alloc()
	rv := objc.Send[VZEFIVariableStore](instance.ID, objc.Sel("initCreatingVariableStoreAtURL:options:error:"), URL, options, error_)
	rv.Autorelease()
	return rv
}



// Initialize the variable store from the URL of an existing file.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZEFIVariableStore/init(url:)
func NewVZEFIVariableStoreWithURL(URL unsafe.Pointer) VZEFIVariableStore {
	instance := getVZEFIVariableStoreClass().Alloc()
	rv := objc.Send[VZEFIVariableStore](instance.ID, objc.Sel("initWithURL:"), URL)
	rv.Autorelease()
	return rv
}


// The URL of the variable store on the local file system.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZEFIVariableStore/url
func (v_ VZEFIVariableStore) URL() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("URL"))
	return rv
}


