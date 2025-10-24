// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTL4CommandAllocatorDescriptor */


/* debug [class_header]: Header for MTL4CommandAllocatorDescriptor */
// The class instance for the [MTL4CommandAllocatorDescriptor] class.
var (
	MTL4CommandAllocatorDescriptorClass     _MTL4CommandAllocatorDescriptorClass
	MTL4CommandAllocatorDescriptorClassOnce sync.Once
)

func getMTL4CommandAllocatorDescriptorClass() _MTL4CommandAllocatorDescriptorClass {
	MTL4CommandAllocatorDescriptorClassOnce.Do(func() {
		MTL4CommandAllocatorDescriptorClass = _MTL4CommandAllocatorDescriptorClass{objc.GetClass("MTL4CommandAllocatorDescriptor")}
	})
	return MTL4CommandAllocatorDescriptorClass
}

type _MTL4CommandAllocatorDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTL4CommandAllocatorDescriptor */
// An interface definition for the [MTL4CommandAllocatorDescriptor] class.
type IMTL4CommandAllocatorDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTL4CommandAllocatorDescriptor */
	// properties:
	Label() objc.IObject /* cross-framework: NSString */
	SetLabel(value objc.IObject /* cross-framework: NSString */)
	MTL4CommandQueueErrorDomain() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTL4CommandAllocatorDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTL4CommandAllocatorDescriptor */
// Alloc allocates a new instance without initialization.
func (mc _MTL4CommandAllocatorDescriptorClass) Alloc() MTL4CommandAllocatorDescriptor {
	rv := objc.Send[MTL4CommandAllocatorDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTL4CommandAllocatorDescriptorClass) New() MTL4CommandAllocatorDescriptor {
	rv := objc.Send[MTL4CommandAllocatorDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTL4CommandAllocatorDescriptor) Init() MTL4CommandAllocatorDescriptor {
	rv := objc.Send[MTL4CommandAllocatorDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTL4CommandAllocatorDescriptor) Autorelease() MTL4CommandAllocatorDescriptor {
	rv := objc.Send[MTL4CommandAllocatorDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4CommandAllocatorDescriptor creates a new MTL4CommandAllocatorDescriptor instance.
func NewMTL4CommandAllocatorDescriptor() MTL4CommandAllocatorDescriptor {
	return getMTL4CommandAllocatorDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTL4CommandAllocatorDescriptor */
// Groups together parameters for creating a command allocator.


// Groups together parameters for creating a command allocator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4CommandAllocatorDescriptor
type MTL4CommandAllocatorDescriptor struct {
	objectivec.Object
}

// MTL4CommandAllocatorDescriptorFrom constructs a [MTL4CommandAllocatorDescriptor] from an unsafe.Pointer.
//
// Groups together parameters for creating a command allocator.
func MTL4CommandAllocatorDescriptorFrom(ptr unsafe.Pointer) MTL4CommandAllocatorDescriptor {
	return MTL4CommandAllocatorDescriptor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTL4CommandAllocatorDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTL4CommandAllocatorDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTL4CommandAllocatorDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTL4CommandAllocatorDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTL4CommandAllocatorDescriptor */

// An optional label you can assign to the command allocator to aid debugging.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4CommandAllocatorDescriptor/label
func (m_ MTL4CommandAllocatorDescriptor) Label() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("label"))
	return rv
}/* debug [instance_properties/getter]: label */


// An optional label you can assign to the command allocator to aid debugging.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4CommandAllocatorDescriptor/label
func (m_ MTL4CommandAllocatorDescriptor) SetLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLabel:"), value)
}/* debug [instance_properties/setter]: label */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4commandqueueerrordomain
func (m_ MTL4CommandAllocatorDescriptor) MTL4CommandQueueErrorDomain() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("MTL4CommandQueueErrorDomain"))
	return rv
}/* debug [instance_properties/getter]: MTL4CommandQueueErrorDomain */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTL4CommandAllocatorDescriptor */



