// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	// properties:
	Url() objc.IObject /* cross-framework: URL */
	SetUrl(value objc.IObject /* cross-framework: URL */)
	// methods:
}

// An object that represents the Extensible Firmware Interface (EFI) variable store that contains NVRAM variables the EFI exposes.


// An object that represents the Extensible Firmware Interface (EFI) variable store that contains NVRAM variables the EFI exposes.
//
// [Full Topic]
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



// The URL of the variable store on the local file system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzefivariablestore/url
func (v_ VZEFIVariableStore) Url() objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[foundation.URL](v_.ID, objc.Sel("url"))
	return rv
}


// The URL of the variable store on the local file system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzefivariablestore/url
func (v_ VZEFIVariableStore) SetUrl(value objc.IObject /* cross-framework: URL */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setUrl:"), value)
}



