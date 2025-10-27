// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [RenderPipelineColorAttachmentDescriptorArray] class.
var (
	RenderPipelineColorAttachmentDescriptorArrayClass     _RenderPipelineColorAttachmentDescriptorArrayClass
	RenderPipelineColorAttachmentDescriptorArrayClassOnce sync.Once
)

func getRenderPipelineColorAttachmentDescriptorArrayClass() _RenderPipelineColorAttachmentDescriptorArrayClass {
	RenderPipelineColorAttachmentDescriptorArrayClassOnce.Do(func() {
		RenderPipelineColorAttachmentDescriptorArrayClass = _RenderPipelineColorAttachmentDescriptorArrayClass{objc.GetClass("MTLRenderPipelineColorAttachmentDescriptorArray")}
	})
	return RenderPipelineColorAttachmentDescriptorArrayClass
}

type _RenderPipelineColorAttachmentDescriptorArrayClass struct {
	class objc.Class
}





// An interface definition for the [RenderPipelineColorAttachmentDescriptorArray] class.
type IRenderPipelineColorAttachmentDescriptorArray interface {
	objectivec.IObject
	

	// properties:


	

	// methods:
	SetObjectAtIndexedSubscript(attachment IMTLRenderPipelineColorAttachmentDescriptor, attachmentIndex uint)
	ObjectAtIndexedSubscript(attachmentIndex uint) IRenderPipelineColorAttachmentDescriptor


}





// Alloc allocates a new instance without initialization.
func (rc _RenderPipelineColorAttachmentDescriptorArrayClass) Alloc() RenderPipelineColorAttachmentDescriptorArray {
	rv := objc.Send[RenderPipelineColorAttachmentDescriptorArray](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _RenderPipelineColorAttachmentDescriptorArrayClass) New() RenderPipelineColorAttachmentDescriptorArray {
	rv := objc.Send[RenderPipelineColorAttachmentDescriptorArray](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RenderPipelineColorAttachmentDescriptorArray) Init() RenderPipelineColorAttachmentDescriptorArray {
	rv := objc.Send[RenderPipelineColorAttachmentDescriptorArray](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RenderPipelineColorAttachmentDescriptorArray) Autorelease() RenderPipelineColorAttachmentDescriptorArray {
	rv := objc.Send[RenderPipelineColorAttachmentDescriptorArray](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRenderPipelineColorAttachmentDescriptorArray creates a new RenderPipelineColorAttachmentDescriptorArray instance.
func NewRenderPipelineColorAttachmentDescriptorArray() RenderPipelineColorAttachmentDescriptorArray {
	return getRenderPipelineColorAttachmentDescriptorArrayClass().New()
}





// An array of render pipeline color attachment descriptor objects.


// An array of render pipeline color attachment descriptor objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineColorAttachmentDescriptorArray
type RenderPipelineColorAttachmentDescriptorArray struct {
	objectivec.Object
}

// RenderPipelineColorAttachmentDescriptorArrayFrom constructs a [RenderPipelineColorAttachmentDescriptorArray] from an unsafe.Pointer.
//
// An array of render pipeline color attachment descriptor objects.
func RenderPipelineColorAttachmentDescriptorArrayFrom(ptr unsafe.Pointer) RenderPipelineColorAttachmentDescriptorArray {
	return RenderPipelineColorAttachmentDescriptorArray{objectivec.Object{objc.ID(ptr)}}
}




















// Sets the render pipeline state for a specified color attachment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineColorAttachmentDescriptorArray/setObject:atIndexedSubscript:
func (r_ RenderPipelineColorAttachmentDescriptorArray) SetObjectAtIndexedSubscript(attachment IMTLRenderPipelineColorAttachmentDescriptor, attachmentIndex uint) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setObject:atIndexedSubscript:"), attachment, attachmentIndex)
}


// Returns the render pipeline state for the specified color attachment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineColorAttachmentDescriptorArray/subscript(_:)
func (r_ RenderPipelineColorAttachmentDescriptorArray) ObjectAtIndexedSubscript(attachmentIndex uint) IRenderPipelineColorAttachmentDescriptor {
	rv := objc.Send[RenderPipelineColorAttachmentDescriptor](r_.ID, objc.Sel("objectAtIndexedSubscript:"), attachmentIndex)
	return rv
}













