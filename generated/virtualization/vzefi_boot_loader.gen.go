// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VZEFIBootLoader */


/* debug [class_header]: Header for VZEFIBootLoader */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VZEFIBootLoader */
// An interface definition for the [VZEFIBootLoader] class.
type IVZEFIBootLoader interface {
	IVZBootLoader
	
/* debug [class_interface_properties]: Properties for VZEFIBootLoader */
	// properties:
	VariableStore() IVZEFIVariableStore
	SetVariableStore(value IVZEFIVariableStore)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VZEFIBootLoader */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VZEFIBootLoader */
// Alloc allocates a new instance without initialization.
func (vc _VZEFIBootLoaderClass) Alloc() VZEFIBootLoader {
	rv := objc.Send[VZEFIBootLoader](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VZEFIBootLoader */
// The boot loader configuration the system uses to boot guest-operating systems that expect an Extensible Firmware Interface (EFI) ROM.


// The boot loader configuration the system uses to boot guest-operating systems that expect an Extensible Firmware Interface (EFI) ROM.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VZEFIBootLoader */
/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VZEFIBootLoader */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VZEFIBootLoader */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VZEFIBootLoader */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VZEFIBootLoader */

// The boot loader’s EFI variable store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZEFIBootLoader/variableStore
func (v_ VZEFIBootLoader) VariableStore() IVZEFIVariableStore {
	rv := objc.Send[VZEFIVariableStore](v_.ID, objc.Sel("variableStore"))
	return rv
}/* debug [instance_properties/getter]: variableStore */


// The boot loader’s EFI variable store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZEFIBootLoader/variableStore
func (v_ VZEFIBootLoader) SetVariableStore(value IVZEFIVariableStore) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setVariableStore:"), value)
}/* debug [instance_properties/setter]: variableStore */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VZEFIBootLoader */


