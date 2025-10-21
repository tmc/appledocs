// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [ResourceStatePassDescriptor] class.
type IResourceStatePassDescriptor interface {
	objectivec.IObject
}

// A configuration for a resource state pass, used to create a resource state command encoder.
//
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

// Alloc allocates a new instance without initialization.
func (rc _ResourceStatePassDescriptorClass) Alloc() ResourceStatePassDescriptor {
	rv := objc.Send[ResourceStatePassDescriptor](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Creates a new resource state pass descriptor.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLResourceStatePassDescriptor/resourceStatePassDescriptor
func (rc _ResourceStatePassDescriptorClass) ResourceStatePassDescriptor() ResourceStatePassDescriptor {
	rv := objc.Send[ResourceStatePassDescriptor](objc.ID(rc.class), objc.Sel("resourceStatePassDescriptor"))
	return rv
}

// The array of sample buffers that the resource state pass can access.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlresourcestatepassdescriptor/samplebufferattachments
func (r_ ResourceStatePassDescriptor) SampleBufferAttachments() MTLResourceStatePassSampleBufferAttachmentDescriptorArray {
	rv := objc.Send[MTLResourceStatePassSampleBufferAttachmentDescriptorArray](r_.ID, objc.Sel("sampleBufferAttachments"))
	return rv
}


// SetSampleBufferAttachments sets the value of the sampleBufferAttachments property.
// The array of sample buffers that the resource state pass can access.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlresourcestatepassdescriptor/samplebufferattachments
func (r_ ResourceStatePassDescriptor) SetSampleBufferAttachments(value IMTLResourceStatePassSampleBufferAttachmentDescriptorArray) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setSampleBufferAttachments:"), value)
}



