// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLResourceStatePassSampleBufferAttachmentDescriptorArray */


/* debug [class_header]: Header for MTLResourceStatePassSampleBufferAttachmentDescriptorArray */
// The class instance for the [ResourceStatePassSampleBufferAttachmentDescriptorArray] class.
var (
	ResourceStatePassSampleBufferAttachmentDescriptorArrayClass     _ResourceStatePassSampleBufferAttachmentDescriptorArrayClass
	ResourceStatePassSampleBufferAttachmentDescriptorArrayClassOnce sync.Once
)

func getResourceStatePassSampleBufferAttachmentDescriptorArrayClass() _ResourceStatePassSampleBufferAttachmentDescriptorArrayClass {
	ResourceStatePassSampleBufferAttachmentDescriptorArrayClassOnce.Do(func() {
		ResourceStatePassSampleBufferAttachmentDescriptorArrayClass = _ResourceStatePassSampleBufferAttachmentDescriptorArrayClass{objc.GetClass("MTLResourceStatePassSampleBufferAttachmentDescriptorArray")}
	})
	return ResourceStatePassSampleBufferAttachmentDescriptorArrayClass
}

type _ResourceStatePassSampleBufferAttachmentDescriptorArrayClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ResourceStatePassSampleBufferAttachmentDescriptorArray */
// An interface definition for the [ResourceStatePassSampleBufferAttachmentDescriptorArray] class.
type IResourceStatePassSampleBufferAttachmentDescriptorArray interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ResourceStatePassSampleBufferAttachmentDescriptorArray */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ResourceStatePassSampleBufferAttachmentDescriptorArray */
	// methods:
	SetObjectAtIndexedSubscript(attachment IMTLResourceStatePassSampleBufferAttachmentDescriptor, attachmentIndex uint)
	ObjectAtIndexedSubscript(attachmentIndex uint) IResourceStatePassSampleBufferAttachmentDescriptor
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ResourceStatePassSampleBufferAttachmentDescriptorArray */
// Alloc allocates a new instance without initialization.
func (rc _ResourceStatePassSampleBufferAttachmentDescriptorArrayClass) Alloc() ResourceStatePassSampleBufferAttachmentDescriptorArray {
	rv := objc.Send[ResourceStatePassSampleBufferAttachmentDescriptorArray](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _ResourceStatePassSampleBufferAttachmentDescriptorArrayClass) New() ResourceStatePassSampleBufferAttachmentDescriptorArray {
	rv := objc.Send[ResourceStatePassSampleBufferAttachmentDescriptorArray](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ResourceStatePassSampleBufferAttachmentDescriptorArray) Init() ResourceStatePassSampleBufferAttachmentDescriptorArray {
	rv := objc.Send[ResourceStatePassSampleBufferAttachmentDescriptorArray](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ResourceStatePassSampleBufferAttachmentDescriptorArray) Autorelease() ResourceStatePassSampleBufferAttachmentDescriptorArray {
	rv := objc.Send[ResourceStatePassSampleBufferAttachmentDescriptorArray](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewResourceStatePassSampleBufferAttachmentDescriptorArray creates a new ResourceStatePassSampleBufferAttachmentDescriptorArray instance.
func NewResourceStatePassSampleBufferAttachmentDescriptorArray() ResourceStatePassSampleBufferAttachmentDescriptorArray {
	return getResourceStatePassSampleBufferAttachmentDescriptorArrayClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ResourceStatePassSampleBufferAttachmentDescriptorArray */
// An array of sample buffer attachments for a resource state pass.


// An array of sample buffer attachments for a resource state pass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLResourceStatePassSampleBufferAttachmentDescriptorArray
type ResourceStatePassSampleBufferAttachmentDescriptorArray struct {
	objectivec.Object
}

// ResourceStatePassSampleBufferAttachmentDescriptorArrayFrom constructs a [ResourceStatePassSampleBufferAttachmentDescriptorArray] from an unsafe.Pointer.
//
// An array of sample buffer attachments for a resource state pass.
func ResourceStatePassSampleBufferAttachmentDescriptorArrayFrom(ptr unsafe.Pointer) ResourceStatePassSampleBufferAttachmentDescriptorArray {
	return ResourceStatePassSampleBufferAttachmentDescriptorArray{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ResourceStatePassSampleBufferAttachmentDescriptorArray *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ResourceStatePassSampleBufferAttachmentDescriptorArray */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ResourceStatePassSampleBufferAttachmentDescriptorArray */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ResourceStatePassSampleBufferAttachmentDescriptorArray */

// Sets the descriptor object for the specified sample buffer attachment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLResourceStatePassSampleBufferAttachmentDescriptorArray/setObject:atIndexedSubscript:
func (r_ ResourceStatePassSampleBufferAttachmentDescriptorArray) SetObjectAtIndexedSubscript(attachment IMTLResourceStatePassSampleBufferAttachmentDescriptor, attachmentIndex uint) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setObject:atIndexedSubscript:"), attachment, attachmentIndex)
}/* debug [instance_methods/method]: SetObjectAtIndexedSubscript */


// Returns the descriptor object for the specified sample buffer attachment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLResourceStatePassSampleBufferAttachmentDescriptorArray/subscript(_:)
func (r_ ResourceStatePassSampleBufferAttachmentDescriptorArray) ObjectAtIndexedSubscript(attachmentIndex uint) IResourceStatePassSampleBufferAttachmentDescriptor {
	rv := objc.Send[ResourceStatePassSampleBufferAttachmentDescriptor](r_.ID, objc.Sel("objectAtIndexedSubscript:"), attachmentIndex)
	return rv
}/* debug [instance_methods/method]: ObjectAtIndexedSubscript */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ResourceStatePassSampleBufferAttachmentDescriptorArray */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLResourceStatePassSampleBufferAttachmentDescriptorArray */



