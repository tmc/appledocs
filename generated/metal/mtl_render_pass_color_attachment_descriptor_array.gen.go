// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [RenderPassColorAttachmentDescriptorArray] class.
var (
	RenderPassColorAttachmentDescriptorArrayClass     _RenderPassColorAttachmentDescriptorArrayClass
	RenderPassColorAttachmentDescriptorArrayClassOnce sync.Once
)

func getRenderPassColorAttachmentDescriptorArrayClass() _RenderPassColorAttachmentDescriptorArrayClass {
	RenderPassColorAttachmentDescriptorArrayClassOnce.Do(func() {
		RenderPassColorAttachmentDescriptorArrayClass = _RenderPassColorAttachmentDescriptorArrayClass{objc.GetClass("MTLRenderPassColorAttachmentDescriptorArray")}
	})
	return RenderPassColorAttachmentDescriptorArrayClass
}

type _RenderPassColorAttachmentDescriptorArrayClass struct {
	class objc.Class
}

// An interface definition for the [RenderPassColorAttachmentDescriptorArray] class.
type IRenderPassColorAttachmentDescriptorArray interface {
	objectivec.IObject
	SetObjectAtIndexedSubscript(attachment unsafe.Pointer, attachmentIndex uint)
	ObjectAtIndexedSubscript(attachmentIndex uint) unsafe.Pointer
}

// An array of render pass color attachment descriptor objects.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPassColorAttachmentDescriptorArray
type RenderPassColorAttachmentDescriptorArray struct {
	objectivec.Object
}

// RenderPassColorAttachmentDescriptorArrayFrom constructs a [RenderPassColorAttachmentDescriptorArray] from an unsafe.Pointer.
//
// An array of render pass color attachment descriptor objects.
func RenderPassColorAttachmentDescriptorArrayFrom(ptr unsafe.Pointer) RenderPassColorAttachmentDescriptorArray {
	return RenderPassColorAttachmentDescriptorArray{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (rc _RenderPassColorAttachmentDescriptorArrayClass) Alloc() RenderPassColorAttachmentDescriptorArray {
	rv := objc.Send[RenderPassColorAttachmentDescriptorArray](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _RenderPassColorAttachmentDescriptorArrayClass) New() RenderPassColorAttachmentDescriptorArray {
	rv := objc.Send[RenderPassColorAttachmentDescriptorArray](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RenderPassColorAttachmentDescriptorArray) Init() RenderPassColorAttachmentDescriptorArray {
	rv := objc.Send[RenderPassColorAttachmentDescriptorArray](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RenderPassColorAttachmentDescriptorArray) Autorelease() RenderPassColorAttachmentDescriptorArray {
	rv := objc.Send[RenderPassColorAttachmentDescriptorArray](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRenderPassColorAttachmentDescriptorArray creates a new RenderPassColorAttachmentDescriptorArray instance.
func NewRenderPassColorAttachmentDescriptorArray() RenderPassColorAttachmentDescriptorArray {
	return getRenderPassColorAttachmentDescriptorArrayClass().New()
}


// Sets the descriptor for the specified color attachment.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPassColorAttachmentDescriptorArray/setObject:atIndexedSubscript:
func (r_ RenderPassColorAttachmentDescriptorArray) SetObjectAtIndexedSubscript(attachment unsafe.Pointer, attachmentIndex uint) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setObject:atIndexedSubscript:"), attachment, attachmentIndex)
}

// Returns the descriptor object for the specified color attachment.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPassColorAttachmentDescriptorArray/subscript(_:)
func (r_ RenderPassColorAttachmentDescriptorArray) ObjectAtIndexedSubscript(attachmentIndex uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("objectAtIndexedSubscript:"), attachmentIndex)
	return rv
}



