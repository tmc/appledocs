// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLResourceStatePassDescriptor */


/* debug [class_header]: Header for MTLResourceStatePassDescriptor */
// The class instance for the [ResourceStatePassDescriptor] class.
var (
	ResourceStatePassDescriptorClass     _ResourceStatePassDescriptorClass
	ResourceStatePassDescriptorClassOnce sync.Once
)

func getResourceStatePassDescriptorClass() _ResourceStatePassDescriptorClass {
	ResourceStatePassDescriptorClassOnce.Do(func() {
		ResourceStatePassDescriptorClass = _ResourceStatePassDescriptorClass{objc.GetClass("MTLResourceStatePassDescriptor")}
	})
	return ResourceStatePassDescriptorClass
}

type _ResourceStatePassDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ResourceStatePassDescriptor */
// An interface definition for the [ResourceStatePassDescriptor] class.
type IResourceStatePassDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ResourceStatePassDescriptor */
	// properties:
	SampleBufferAttachments() IMTLResourceStatePassSampleBufferAttachmentDescriptorArray
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ResourceStatePassDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ResourceStatePassDescriptor */
// Alloc allocates a new instance without initialization.
func (rc _ResourceStatePassDescriptorClass) Alloc() ResourceStatePassDescriptor {
	rv := objc.Send[ResourceStatePassDescriptor](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _ResourceStatePassDescriptorClass) New() ResourceStatePassDescriptor {
	rv := objc.Send[ResourceStatePassDescriptor](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ResourceStatePassDescriptor) Init() ResourceStatePassDescriptor {
	rv := objc.Send[ResourceStatePassDescriptor](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ResourceStatePassDescriptor) Autorelease() ResourceStatePassDescriptor {
	rv := objc.Send[ResourceStatePassDescriptor](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewResourceStatePassDescriptor creates a new ResourceStatePassDescriptor instance.
func NewResourceStatePassDescriptor() ResourceStatePassDescriptor {
	return getResourceStatePassDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ResourceStatePassDescriptor */
// A configuration for a resource state pass, used to create a resource state command encoder.


// A configuration for a resource state pass, used to create a resource state command encoder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLResourceStatePassDescriptor
type ResourceStatePassDescriptor struct {
	objectivec.Object
}

// ResourceStatePassDescriptorFrom constructs a [ResourceStatePassDescriptor] from an unsafe.Pointer.
//
// A configuration for a resource state pass, used to create a resource state command encoder.
func ResourceStatePassDescriptorFrom(ptr unsafe.Pointer) ResourceStatePassDescriptor {
	return ResourceStatePassDescriptor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ResourceStatePassDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ResourceStatePassDescriptor */

// Creates a new resource state pass descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLResourceStatePassDescriptor/resourceStatePassDescriptor
func (rc _ResourceStatePassDescriptorClass) ResourceStatePassDescriptor() IResourceStatePassDescriptor {
	rv := objc.Send[ResourceStatePassDescriptor](objc.ID(rc.class), objc.Sel("resourceStatePassDescriptor"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ResourceStatePassDescriptor) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ResourceStatePassDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ResourceStatePassDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ResourceStatePassDescriptor */

// The array of sample buffers that the resource state pass can access.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLResourceStatePassDescriptor/sampleBufferAttachments
func (r_ ResourceStatePassDescriptor) SampleBufferAttachments() IMTLResourceStatePassSampleBufferAttachmentDescriptorArray {
	rv := objc.Send[ResourceStatePassSampleBufferAttachmentDescriptorArray](r_.ID, objc.Sel("sampleBufferAttachments"))
	return rv
}/* debug [instance_properties/getter]: sampleBufferAttachments */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLResourceStatePassDescriptor */



