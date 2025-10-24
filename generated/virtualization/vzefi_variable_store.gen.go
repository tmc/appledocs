// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VZEFIVariableStore */

/* debug [class_header]: Header for VZEFIVariableStore */
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

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for VZEFIVariableStore */
// An interface definition for the [VZEFIVariableStore] class.
type IVZEFIVariableStore interface {
	objectivec.IObject

	/* debug [class_interface_properties]: Properties for VZEFIVariableStore */
	// properties:
	URL() objc.IObject /* cross-framework: NSURL */
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for VZEFIVariableStore */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for VZEFIVariableStore */
// Alloc allocates a new instance without initialization.
func (vc _VZEFIVariableStoreClass) Alloc() VZEFIVariableStore {
	rv := objc.Send[VZEFIVariableStore](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for VZEFIVariableStore */
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

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for VZEFIVariableStore */

// Creates a new EFI variable store at specified the URL on the filesystem, initialization options, and error-return variable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZEFIVariableStore/init(creatingVariableStoreAt:options:)
func NewVZEFIVariableStoreCreatingVariableStoreAtURLOptionsError(URL objc.IObject /* cross-framework: NSURL */, options VZEFIVariableStoreInitializationOptions, error_ unsafe.Pointer) VZEFIVariableStore {
	instance := getVZEFIVariableStoreClass().Alloc()
	rv := objc.Send[VZEFIVariableStore](instance.ID, objc.Sel("initCreatingVariableStoreAtURL:options:error:"), URL, options, error_)
	rv.Autorelease()
	return rv
} /* debug [class_init_methods/constructor]: NewVZEFIVariableStoreCreatingVariableStoreAtURLOptionsError */

// Initialize the variable store from the URL of an existing file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZEFIVariableStore/init(url:)
func NewVZEFIVariableStoreWithURL(URL objc.IObject /* cross-framework: NSURL */) VZEFIVariableStore {
	instance := getVZEFIVariableStoreClass().Alloc()
	rv := objc.Send[VZEFIVariableStore](instance.ID, objc.Sel("initWithURL:"), URL)
	rv.Autorelease()
	return rv
} /* debug [class_init_methods/constructor]: NewVZEFIVariableStoreWithURL */

/* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for VZEFIVariableStore */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for VZEFIVariableStore */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for VZEFIVariableStore */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for VZEFIVariableStore */

// The URL of the variable store on the local file system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZEFIVariableStore/url
func (v_ VZEFIVariableStore) URL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](v_.ID, objc.Sel("URL"))
	return rv
} /* debug [instance_properties/getter]: URL */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class VZEFIVariableStore */
