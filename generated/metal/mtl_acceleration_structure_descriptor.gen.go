// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLAccelerationStructureDescriptor */


/* debug [class_header]: Header for MTLAccelerationStructureDescriptor */
// The class instance for the [AccelerationStructureDescriptor] class.
var (
	AccelerationStructureDescriptorClass     _AccelerationStructureDescriptorClass
	AccelerationStructureDescriptorClassOnce sync.Once
)

func getAccelerationStructureDescriptorClass() _AccelerationStructureDescriptorClass {
	AccelerationStructureDescriptorClassOnce.Do(func() {
		AccelerationStructureDescriptorClass = _AccelerationStructureDescriptorClass{objc.GetClass("MTLAccelerationStructureDescriptor")}
	})
	return AccelerationStructureDescriptorClass
}

type _AccelerationStructureDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AccelerationStructureDescriptor */
// An interface definition for the [AccelerationStructureDescriptor] class.
type IAccelerationStructureDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AccelerationStructureDescriptor */
	// properties:
	Usage() AccelerationStructureUsage
	SetUsage(value AccelerationStructureUsage)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AccelerationStructureDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AccelerationStructureDescriptor */
// Alloc allocates a new instance without initialization.
func (ac _AccelerationStructureDescriptorClass) Alloc() AccelerationStructureDescriptor {
	rv := objc.Send[AccelerationStructureDescriptor](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AccelerationStructureDescriptorClass) New() AccelerationStructureDescriptor {
	rv := objc.Send[AccelerationStructureDescriptor](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AccelerationStructureDescriptor) Init() AccelerationStructureDescriptor {
	rv := objc.Send[AccelerationStructureDescriptor](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AccelerationStructureDescriptor) Autorelease() AccelerationStructureDescriptor {
	rv := objc.Send[AccelerationStructureDescriptor](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAccelerationStructureDescriptor creates a new AccelerationStructureDescriptor instance.
func NewAccelerationStructureDescriptor() AccelerationStructureDescriptor {
	return getAccelerationStructureDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AccelerationStructureDescriptor */
// A base class for classes that define the configuration for a new acceleration structure.
//
// This is the base class for other acceleration structure descriptors. Don’t use this class directly. Use one of the derived classes instead, as describes.


// A base class for classes that define the configuration for a new acceleration structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureDescriptor
type AccelerationStructureDescriptor struct {
	objectivec.Object
}

// AccelerationStructureDescriptorFrom constructs a [AccelerationStructureDescriptor] from an unsafe.Pointer.
//
// A base class for classes that define the configuration for a new acceleration structure.
func AccelerationStructureDescriptorFrom(ptr unsafe.Pointer) AccelerationStructureDescriptor {
	return AccelerationStructureDescriptor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AccelerationStructureDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AccelerationStructureDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AccelerationStructureDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AccelerationStructureDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AccelerationStructureDescriptor */

// The options that describe how you intend to use the acceleration structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureDescriptor/usage
func (a_ AccelerationStructureDescriptor) Usage() AccelerationStructureUsage {
	rv := objc.Send[AccelerationStructureUsage](a_.ID, objc.Sel("usage"))
	return rv
}/* debug [instance_properties/getter]: usage */


// The options that describe how you intend to use the acceleration structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureDescriptor/usage
func (a_ AccelerationStructureDescriptor) SetUsage(value AccelerationStructureUsage) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setUsage:"), value)
}/* debug [instance_properties/setter]: usage */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLAccelerationStructureDescriptor */



