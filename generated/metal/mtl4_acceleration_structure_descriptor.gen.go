// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MTL4AccelerationStructureDescriptor */


/* debug [class_header]: Header for MTL4AccelerationStructureDescriptor */
// The class instance for the [MTL4AccelerationStructureDescriptor] class.
var (
	MTL4AccelerationStructureDescriptorClass     _MTL4AccelerationStructureDescriptorClass
	MTL4AccelerationStructureDescriptorClassOnce sync.Once
)

func getMTL4AccelerationStructureDescriptorClass() _MTL4AccelerationStructureDescriptorClass {
	MTL4AccelerationStructureDescriptorClassOnce.Do(func() {
		MTL4AccelerationStructureDescriptorClass = _MTL4AccelerationStructureDescriptorClass{objc.GetClass("MTL4AccelerationStructureDescriptor")}
	})
	return MTL4AccelerationStructureDescriptorClass
}

type _MTL4AccelerationStructureDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTL4AccelerationStructureDescriptor */
// An interface definition for the [MTL4AccelerationStructureDescriptor] class.
type IMTL4AccelerationStructureDescriptor interface {
	IAccelerationStructureDescriptor
	
/* debug [class_interface_properties]: Properties for MTL4AccelerationStructureDescriptor */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTL4AccelerationStructureDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTL4AccelerationStructureDescriptor */
// Alloc allocates a new instance without initialization.
func (mc _MTL4AccelerationStructureDescriptorClass) Alloc() MTL4AccelerationStructureDescriptor {
	rv := objc.Send[MTL4AccelerationStructureDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTL4AccelerationStructureDescriptorClass) New() MTL4AccelerationStructureDescriptor {
	rv := objc.Send[MTL4AccelerationStructureDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTL4AccelerationStructureDescriptor) Init() MTL4AccelerationStructureDescriptor {
	rv := objc.Send[MTL4AccelerationStructureDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTL4AccelerationStructureDescriptor) Autorelease() MTL4AccelerationStructureDescriptor {
	rv := objc.Send[MTL4AccelerationStructureDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4AccelerationStructureDescriptor creates a new MTL4AccelerationStructureDescriptor instance.
func NewMTL4AccelerationStructureDescriptor() MTL4AccelerationStructureDescriptor {
	return getMTL4AccelerationStructureDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTL4AccelerationStructureDescriptor */
// Base class for Metal 4 acceleration structure descriptors.
//
// Don’t use this class directly. Use one of its subclasses instead.


// Base class for Metal 4 acceleration structure descriptors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureDescriptor
type MTL4AccelerationStructureDescriptor struct {
	AccelerationStructureDescriptor
}

// MTL4AccelerationStructureDescriptorFrom constructs a [MTL4AccelerationStructureDescriptor] from an unsafe.Pointer.
//
// Base class for Metal 4 acceleration structure descriptors.
func MTL4AccelerationStructureDescriptorFrom(ptr unsafe.Pointer) MTL4AccelerationStructureDescriptor {
	return MTL4AccelerationStructureDescriptor{
		AccelerationStructureDescriptor: AccelerationStructureDescriptorFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTL4AccelerationStructureDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTL4AccelerationStructureDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTL4AccelerationStructureDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTL4AccelerationStructureDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTL4AccelerationStructureDescriptor */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTL4AccelerationStructureDescriptor */



