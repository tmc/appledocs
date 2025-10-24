// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLResourceStatePassSampleBufferAttachmentDescriptor */


/* debug [class_header]: Header for MTLResourceStatePassSampleBufferAttachmentDescriptor */
// The class instance for the [ResourceStatePassSampleBufferAttachmentDescriptor] class.
var (
	ResourceStatePassSampleBufferAttachmentDescriptorClass     _ResourceStatePassSampleBufferAttachmentDescriptorClass
	ResourceStatePassSampleBufferAttachmentDescriptorClassOnce sync.Once
)

func getResourceStatePassSampleBufferAttachmentDescriptorClass() _ResourceStatePassSampleBufferAttachmentDescriptorClass {
	ResourceStatePassSampleBufferAttachmentDescriptorClassOnce.Do(func() {
		ResourceStatePassSampleBufferAttachmentDescriptorClass = _ResourceStatePassSampleBufferAttachmentDescriptorClass{objc.GetClass("MTLResourceStatePassSampleBufferAttachmentDescriptor")}
	})
	return ResourceStatePassSampleBufferAttachmentDescriptorClass
}

type _ResourceStatePassSampleBufferAttachmentDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ResourceStatePassSampleBufferAttachmentDescriptor */
// An interface definition for the [ResourceStatePassSampleBufferAttachmentDescriptor] class.
type IResourceStatePassSampleBufferAttachmentDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ResourceStatePassSampleBufferAttachmentDescriptor */
	// properties:
	EndOfEncoderSampleIndex() uint
	SetEndOfEncoderSampleIndex(value uint)
	SampleBuffer() unsafe.Pointer
	SetSampleBuffer(value unsafe.Pointer)
	StartOfEncoderSampleIndex() uint
	SetStartOfEncoderSampleIndex(value uint)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ResourceStatePassSampleBufferAttachmentDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ResourceStatePassSampleBufferAttachmentDescriptor */
// Alloc allocates a new instance without initialization.
func (rc _ResourceStatePassSampleBufferAttachmentDescriptorClass) Alloc() ResourceStatePassSampleBufferAttachmentDescriptor {
	rv := objc.Send[ResourceStatePassSampleBufferAttachmentDescriptor](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _ResourceStatePassSampleBufferAttachmentDescriptorClass) New() ResourceStatePassSampleBufferAttachmentDescriptor {
	rv := objc.Send[ResourceStatePassSampleBufferAttachmentDescriptor](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ResourceStatePassSampleBufferAttachmentDescriptor) Init() ResourceStatePassSampleBufferAttachmentDescriptor {
	rv := objc.Send[ResourceStatePassSampleBufferAttachmentDescriptor](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ResourceStatePassSampleBufferAttachmentDescriptor) Autorelease() ResourceStatePassSampleBufferAttachmentDescriptor {
	rv := objc.Send[ResourceStatePassSampleBufferAttachmentDescriptor](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewResourceStatePassSampleBufferAttachmentDescriptor creates a new ResourceStatePassSampleBufferAttachmentDescriptor instance.
func NewResourceStatePassSampleBufferAttachmentDescriptor() ResourceStatePassSampleBufferAttachmentDescriptor {
	return getResourceStatePassSampleBufferAttachmentDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ResourceStatePassSampleBufferAttachmentDescriptor */
// A description of where to store GPU counter information at the start and end of a resource state pass.


// A description of where to store GPU counter information at the start and end of a resource state pass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLResourceStatePassSampleBufferAttachmentDescriptor
type ResourceStatePassSampleBufferAttachmentDescriptor struct {
	objectivec.Object
}

// ResourceStatePassSampleBufferAttachmentDescriptorFrom constructs a [ResourceStatePassSampleBufferAttachmentDescriptor] from an unsafe.Pointer.
//
// A description of where to store GPU counter information at the start and end of a resource state pass.
func ResourceStatePassSampleBufferAttachmentDescriptorFrom(ptr unsafe.Pointer) ResourceStatePassSampleBufferAttachmentDescriptor {
	return ResourceStatePassSampleBufferAttachmentDescriptor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ResourceStatePassSampleBufferAttachmentDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ResourceStatePassSampleBufferAttachmentDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ResourceStatePassSampleBufferAttachmentDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ResourceStatePassSampleBufferAttachmentDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ResourceStatePassSampleBufferAttachmentDescriptor */

// The index the Metal device object should use to store GPU counters when ending the resource state pass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLResourceStatePassSampleBufferAttachmentDescriptor/endOfEncoderSampleIndex
func (r_ ResourceStatePassSampleBufferAttachmentDescriptor) EndOfEncoderSampleIndex() uint {
	rv := objc.Send[uint](r_.ID, objc.Sel("endOfEncoderSampleIndex"))
	return rv
}/* debug [instance_properties/getter]: endOfEncoderSampleIndex */


// The index the Metal device object should use to store GPU counters when ending the resource state pass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLResourceStatePassSampleBufferAttachmentDescriptor/endOfEncoderSampleIndex
func (r_ ResourceStatePassSampleBufferAttachmentDescriptor) SetEndOfEncoderSampleIndex(value uint) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setEndOfEncoderSampleIndex:"), value)
}/* debug [instance_properties/setter]: endOfEncoderSampleIndex */


// A specialized memory buffer that the GPU uses to store its counter data during the resource state pass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLResourceStatePassSampleBufferAttachmentDescriptor/sampleBuffer
func (r_ ResourceStatePassSampleBufferAttachmentDescriptor) SampleBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("sampleBuffer"))
	return rv
}/* debug [instance_properties/getter]: sampleBuffer */


// A specialized memory buffer that the GPU uses to store its counter data during the resource state pass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLResourceStatePassSampleBufferAttachmentDescriptor/sampleBuffer
func (r_ ResourceStatePassSampleBufferAttachmentDescriptor) SetSampleBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setSampleBuffer:"), value)
}/* debug [instance_properties/setter]: sampleBuffer */


// The index the Metal device object should use to store GPU counters when starting the resource state pass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLResourceStatePassSampleBufferAttachmentDescriptor/startOfEncoderSampleIndex
func (r_ ResourceStatePassSampleBufferAttachmentDescriptor) StartOfEncoderSampleIndex() uint {
	rv := objc.Send[uint](r_.ID, objc.Sel("startOfEncoderSampleIndex"))
	return rv
}/* debug [instance_properties/getter]: startOfEncoderSampleIndex */


// The index the Metal device object should use to store GPU counters when starting the resource state pass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLResourceStatePassSampleBufferAttachmentDescriptor/startOfEncoderSampleIndex
func (r_ ResourceStatePassSampleBufferAttachmentDescriptor) SetStartOfEncoderSampleIndex(value uint) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setStartOfEncoderSampleIndex:"), value)
}/* debug [instance_properties/setter]: startOfEncoderSampleIndex */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLResourceStatePassSampleBufferAttachmentDescriptor */



