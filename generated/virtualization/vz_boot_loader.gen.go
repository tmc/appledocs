// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VZBootLoader */

/* debug [class_header]: Header for VZBootLoader */
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

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for VZBootLoader */
// An interface definition for the [VZBootLoader] class.
type IVZBootLoader interface {
	objectivec.IObject

	/* debug [class_interface_properties]: Properties for VZBootLoader */
	// properties:
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for VZBootLoader */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for VZBootLoader */
// Alloc allocates a new instance without initialization.
func (vc _VZBootLoaderClass) Alloc() VZBootLoader {
	rv := objc.Send[VZBootLoader](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for VZBootLoader */
// The base class that defines the management of the initial process of the guest system.
//
// The abstract class defines the common behaviors for booting a guest operating system into a VM. Don’t create instances of this class directly. Instead, instantiate the subclass that corresponds to the type of operating system you plan to load. For example, to create a boot loader object for a Linux kernel, create a object; to create a boot loader object for installation using an ISO image create a . For a macOS system create .

// The base class that defines the management of the initial process of the guest system.
//
// [Full Topic]
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

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for VZBootLoader */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for VZBootLoader */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for VZBootLoader */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for VZBootLoader */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for VZBootLoader */
/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class VZBootLoader */
