// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [ResourceViewPoolDescriptor] class.
type IResourceViewPoolDescriptor interface {
	objectivec.IObject
	Label() string
	SetLabel(value string)
	ResourceViewCount() int
	SetResourceViewCount(value int)
}

// Provides parameters for creating a resource view pool.
//
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

// Alloc allocates a new instance without initialization.
func (rc _ResourceViewPoolDescriptorClass) Alloc() ResourceViewPoolDescriptor {
	rv := objc.Send[ResourceViewPoolDescriptor](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Assigns an optional label you to the resource view pool for debugging purposes.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLResourceViewPoolDescriptor/label
func (r_ ResourceViewPoolDescriptor) Label() string {
	rv := objc.Send[string](r_.ID, objc.Sel("label"))
	return rv
}


// SetLabel sets the value of the label property.
// Assigns an optional label you to the resource view pool for debugging purposes.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLResourceViewPoolDescriptor/label
func (r_ ResourceViewPoolDescriptor) SetLabel(value string) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setLabel:"), objc.String(value))
}

// Configures the number of resource views with which Metal creates the resource view pool.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlresourceviewpooldescriptor/resourceviewcount
func (r_ ResourceViewPoolDescriptor) ResourceViewCount() int {
	rv := objc.Send[int](r_.ID, objc.Sel("resourceViewCount"))
	return rv
}


// SetResourceViewCount sets the value of the resourceViewCount property.
// Configures the number of resource views with which Metal creates the resource view pool.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlresourceviewpooldescriptor/resourceviewcount
func (r_ ResourceViewPoolDescriptor) SetResourceViewCount(value int) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setResourceViewCount:"), value)
}



