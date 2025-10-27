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
	

	// properties:
	Label() foundation.foundation.INSString
	SetLabel(value foundation.foundation.INSString)
	ResourceViewCount() uint
	SetResourceViewCount(value uint)


	

	// methods:


}





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

























// Assigns an optional label you to the resource view pool for debugging purposes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLResourceViewPoolDescriptor/label
func (r_ ResourceViewPoolDescriptor) Label() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](r_.ID, objc.Sel("label"))
	return rv
}


// Assigns an optional label you to the resource view pool for debugging purposes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLResourceViewPoolDescriptor/label
func (r_ ResourceViewPoolDescriptor) SetLabel(value foundation.foundation.INSString) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setLabel:"), value)
}


// Configures the number of resource views with which Metal creates the resource view pool.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLResourceViewPoolDescriptor/resourceViewCount
func (r_ ResourceViewPoolDescriptor) ResourceViewCount() uint {
	rv := objc.Send[uint](r_.ID, objc.Sel("resourceViewCount"))
	return rv
}


// Configures the number of resource views with which Metal creates the resource view pool.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLResourceViewPoolDescriptor/resourceViewCount
func (r_ ResourceViewPoolDescriptor) SetResourceViewCount(value uint) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setResourceViewCount:"), value)
}








