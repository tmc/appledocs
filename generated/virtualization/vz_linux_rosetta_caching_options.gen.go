// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VZLinuxRosettaCachingOptions */

/* debug [class_header]: Header for VZLinuxRosettaCachingOptions */
// The class instance for the [VZLinuxRosettaCachingOptions] class.
var (
	VZLinuxRosettaCachingOptionsClass     _VZLinuxRosettaCachingOptionsClass
	VZLinuxRosettaCachingOptionsClassOnce sync.Once
)

func getVZLinuxRosettaCachingOptionsClass() _VZLinuxRosettaCachingOptionsClass {
	VZLinuxRosettaCachingOptionsClassOnce.Do(func() {
		VZLinuxRosettaCachingOptionsClass = _VZLinuxRosettaCachingOptionsClass{objc.GetClass("VZLinuxRosettaCachingOptions")}
	})
	return VZLinuxRosettaCachingOptionsClass
}

type _VZLinuxRosettaCachingOptionsClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for VZLinuxRosettaCachingOptions */
// An interface definition for the [VZLinuxRosettaCachingOptions] class.
type IVZLinuxRosettaCachingOptions interface {
	objectivec.IObject

	/* debug [class_interface_properties]: Properties for VZLinuxRosettaCachingOptions */
	// properties:
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for VZLinuxRosettaCachingOptions */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for VZLinuxRosettaCachingOptions */
// Alloc allocates a new instance without initialization.
func (vc _VZLinuxRosettaCachingOptionsClass) Alloc() VZLinuxRosettaCachingOptions {
	rv := objc.Send[VZLinuxRosettaCachingOptions](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VZLinuxRosettaCachingOptionsClass) New() VZLinuxRosettaCachingOptions {
	rv := objc.Send[VZLinuxRosettaCachingOptions](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZLinuxRosettaCachingOptions) Init() VZLinuxRosettaCachingOptions {
	rv := objc.Send[VZLinuxRosettaCachingOptions](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZLinuxRosettaCachingOptions) Autorelease() VZLinuxRosettaCachingOptions {
	rv := objc.Send[VZLinuxRosettaCachingOptions](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZLinuxRosettaCachingOptions creates a new VZLinuxRosettaCachingOptions instance.
func NewVZLinuxRosettaCachingOptions() VZLinuxRosettaCachingOptions {
	return getVZLinuxRosettaCachingOptionsClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for VZLinuxRosettaCachingOptions */
// An abstract class that defines UNIX socket-based caching options for Rosetta.
//
// define the communication mechanism between the Rosetta daemon and the Rosetta runtime. Don’t instantiate directly. Use one of its subclasses, such as or instead.

// An abstract class that defines UNIX socket-based caching options for Rosetta.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZLinuxRosettaCachingOptions
type VZLinuxRosettaCachingOptions struct {
	objectivec.Object
}

// VZLinuxRosettaCachingOptionsFrom constructs a [VZLinuxRosettaCachingOptions] from an unsafe.Pointer.
//
// An abstract class that defines UNIX socket-based caching options for Rosetta.
func VZLinuxRosettaCachingOptionsFrom(ptr unsafe.Pointer) VZLinuxRosettaCachingOptions {
	return VZLinuxRosettaCachingOptions{objectivec.Object{objc.ID(ptr)}}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for VZLinuxRosettaCachingOptions */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for VZLinuxRosettaCachingOptions */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for VZLinuxRosettaCachingOptions */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for VZLinuxRosettaCachingOptions */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for VZLinuxRosettaCachingOptions */
/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class VZLinuxRosettaCachingOptions */
