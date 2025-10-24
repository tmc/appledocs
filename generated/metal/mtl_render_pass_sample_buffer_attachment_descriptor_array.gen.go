// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLRenderPassSampleBufferAttachmentDescriptorArray */


/* debug [class_header]: Header for MTLRenderPassSampleBufferAttachmentDescriptorArray */
// The class instance for the [RenderPassSampleBufferAttachmentDescriptorArray] class.
var (
	RenderPassSampleBufferAttachmentDescriptorArrayClass     _RenderPassSampleBufferAttachmentDescriptorArrayClass
	RenderPassSampleBufferAttachmentDescriptorArrayClassOnce sync.Once
)

func getRenderPassSampleBufferAttachmentDescriptorArrayClass() _RenderPassSampleBufferAttachmentDescriptorArrayClass {
	RenderPassSampleBufferAttachmentDescriptorArrayClassOnce.Do(func() {
		RenderPassSampleBufferAttachmentDescriptorArrayClass = _RenderPassSampleBufferAttachmentDescriptorArrayClass{objc.GetClass("MTLRenderPassSampleBufferAttachmentDescriptorArray")}
	})
	return RenderPassSampleBufferAttachmentDescriptorArrayClass
}

type _RenderPassSampleBufferAttachmentDescriptorArrayClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for RenderPassSampleBufferAttachmentDescriptorArray */
// An interface definition for the [RenderPassSampleBufferAttachmentDescriptorArray] class.
type IRenderPassSampleBufferAttachmentDescriptorArray interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for RenderPassSampleBufferAttachmentDescriptorArray */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for RenderPassSampleBufferAttachmentDescriptorArray */
	// methods:
	SetObjectAtIndexedSubscript(attachment IMTLRenderPassSampleBufferAttachmentDescriptor, attachmentIndex uint)
	ObjectAtIndexedSubscript(attachmentIndex uint) IRenderPassSampleBufferAttachmentDescriptor
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for RenderPassSampleBufferAttachmentDescriptorArray */
// Alloc allocates a new instance without initialization.
func (rc _RenderPassSampleBufferAttachmentDescriptorArrayClass) Alloc() RenderPassSampleBufferAttachmentDescriptorArray {
	rv := objc.Send[RenderPassSampleBufferAttachmentDescriptorArray](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _RenderPassSampleBufferAttachmentDescriptorArrayClass) New() RenderPassSampleBufferAttachmentDescriptorArray {
	rv := objc.Send[RenderPassSampleBufferAttachmentDescriptorArray](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RenderPassSampleBufferAttachmentDescriptorArray) Init() RenderPassSampleBufferAttachmentDescriptorArray {
	rv := objc.Send[RenderPassSampleBufferAttachmentDescriptorArray](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RenderPassSampleBufferAttachmentDescriptorArray) Autorelease() RenderPassSampleBufferAttachmentDescriptorArray {
	rv := objc.Send[RenderPassSampleBufferAttachmentDescriptorArray](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRenderPassSampleBufferAttachmentDescriptorArray creates a new RenderPassSampleBufferAttachmentDescriptorArray instance.
func NewRenderPassSampleBufferAttachmentDescriptorArray() RenderPassSampleBufferAttachmentDescriptorArray {
	return getRenderPassSampleBufferAttachmentDescriptorArrayClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for RenderPassSampleBufferAttachmentDescriptorArray */
// An array of sample buffer attachments for a render pass.


// An array of sample buffer attachments for a render pass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPassSampleBufferAttachmentDescriptorArray
type RenderPassSampleBufferAttachmentDescriptorArray struct {
	objectivec.Object
}

// RenderPassSampleBufferAttachmentDescriptorArrayFrom constructs a [RenderPassSampleBufferAttachmentDescriptorArray] from an unsafe.Pointer.
//
// An array of sample buffer attachments for a render pass.
func RenderPassSampleBufferAttachmentDescriptorArrayFrom(ptr unsafe.Pointer) RenderPassSampleBufferAttachmentDescriptorArray {
	return RenderPassSampleBufferAttachmentDescriptorArray{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for RenderPassSampleBufferAttachmentDescriptorArray *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for RenderPassSampleBufferAttachmentDescriptorArray */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for RenderPassSampleBufferAttachmentDescriptorArray */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for RenderPassSampleBufferAttachmentDescriptorArray */

// Sets the descriptor object for the specified sample buffer attachment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPassSampleBufferAttachmentDescriptorArray/setObject:atIndexedSubscript:
func (r_ RenderPassSampleBufferAttachmentDescriptorArray) SetObjectAtIndexedSubscript(attachment IMTLRenderPassSampleBufferAttachmentDescriptor, attachmentIndex uint) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setObject:atIndexedSubscript:"), attachment, attachmentIndex)
}/* debug [instance_methods/method]: SetObjectAtIndexedSubscript */


// Returns the descriptor object for the specified sample buffer attachment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPassSampleBufferAttachmentDescriptorArray/subscript(_:)
func (r_ RenderPassSampleBufferAttachmentDescriptorArray) ObjectAtIndexedSubscript(attachmentIndex uint) IRenderPassSampleBufferAttachmentDescriptor {
	rv := objc.Send[RenderPassSampleBufferAttachmentDescriptor](r_.ID, objc.Sel("objectAtIndexedSubscript:"), attachmentIndex)
	return rv
}/* debug [instance_methods/method]: ObjectAtIndexedSubscript */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for RenderPassSampleBufferAttachmentDescriptorArray */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLRenderPassSampleBufferAttachmentDescriptorArray */



