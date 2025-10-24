// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLResourceViewPoolDescriptor */


/* debug [class_header]: Header for MTLResourceViewPoolDescriptor */
// The class instance for the [ResourceViewPoolDescriptor] class.
var (
	ResourceViewPoolDescriptorClass     _ResourceViewPoolDescriptorClass
	ResourceViewPoolDescriptorClassOnce sync.Once
)

func getResourceViewPoolDescriptorClass() _ResourceViewPoolDescriptorClass {
	ResourceViewPoolDescriptorClassOnce.Do(func() {
		ResourceViewPoolDescriptorClass = _ResourceViewPoolDescriptorClass{objc.GetClass("MTLResourceViewPoolDescriptor")}
	})
	return ResourceViewPoolDescriptorClass
}

type _ResourceViewPoolDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ResourceViewPoolDescriptor */
// An interface definition for the [ResourceViewPoolDescriptor] class.
type IResourceViewPoolDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ResourceViewPoolDescriptor */
	// properties:
	Label() objc.IObject /* cross-framework: NSString */
	SetLabel(value objc.IObject /* cross-framework: NSString */)
	ResourceViewCount() uint
	SetResourceViewCount(value uint)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ResourceViewPoolDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ResourceViewPoolDescriptor */
// Alloc allocates a new instance without initialization.
func (rc _ResourceViewPoolDescriptorClass) Alloc() ResourceViewPoolDescriptor {
	rv := objc.Send[ResourceViewPoolDescriptor](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _ResourceViewPoolDescriptorClass) New() ResourceViewPoolDescriptor {
	rv := objc.Send[ResourceViewPoolDescriptor](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ResourceViewPoolDescriptor) Init() ResourceViewPoolDescriptor {
	rv := objc.Send[ResourceViewPoolDescriptor](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ResourceViewPoolDescriptor) Autorelease() ResourceViewPoolDescriptor {
	rv := objc.Send[ResourceViewPoolDescriptor](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewResourceViewPoolDescriptor creates a new ResourceViewPoolDescriptor instance.
func NewResourceViewPoolDescriptor() ResourceViewPoolDescriptor {
	return getResourceViewPoolDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ResourceViewPoolDescriptor */
// Provides parameters for creating a resource view pool.


// Provides parameters for creating a resource view pool.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLResourceViewPoolDescriptor
type ResourceViewPoolDescriptor struct {
	objectivec.Object
}

// ResourceViewPoolDescriptorFrom constructs a [ResourceViewPoolDescriptor] from an unsafe.Pointer.
//
// Provides parameters for creating a resource view pool.
func ResourceViewPoolDescriptorFrom(ptr unsafe.Pointer) ResourceViewPoolDescriptor {
	return ResourceViewPoolDescriptor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ResourceViewPoolDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ResourceViewPoolDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ResourceViewPoolDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ResourceViewPoolDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ResourceViewPoolDescriptor */

// Assigns an optional label you to the resource view pool for debugging purposes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLResourceViewPoolDescriptor/label
func (r_ ResourceViewPoolDescriptor) Label() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](r_.ID, objc.Sel("label"))
	return rv
}/* debug [instance_properties/getter]: label */


// Assigns an optional label you to the resource view pool for debugging purposes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLResourceViewPoolDescriptor/label
func (r_ ResourceViewPoolDescriptor) SetLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setLabel:"), value)
}/* debug [instance_properties/setter]: label */


// Configures the number of resource views with which Metal creates the resource view pool.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLResourceViewPoolDescriptor/resourceViewCount
func (r_ ResourceViewPoolDescriptor) ResourceViewCount() uint {
	rv := objc.Send[uint](r_.ID, objc.Sel("resourceViewCount"))
	return rv
}/* debug [instance_properties/getter]: resourceViewCount */


// Configures the number of resource views with which Metal creates the resource view pool.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLResourceViewPoolDescriptor/resourceViewCount
func (r_ ResourceViewPoolDescriptor) SetResourceViewCount(value uint) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setResourceViewCount:"), value)
}/* debug [instance_properties/setter]: resourceViewCount */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLResourceViewPoolDescriptor */



