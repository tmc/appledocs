// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTL4RenderPipelineColorAttachmentDescriptorArray] class.
var (
	MTL4RenderPipelineColorAttachmentDescriptorArrayClass     _MTL4RenderPipelineColorAttachmentDescriptorArrayClass
	MTL4RenderPipelineColorAttachmentDescriptorArrayClassOnce sync.Once
)

func getMTL4RenderPipelineColorAttachmentDescriptorArrayClass() _MTL4RenderPipelineColorAttachmentDescriptorArrayClass {
	MTL4RenderPipelineColorAttachmentDescriptorArrayClassOnce.Do(func() {
		MTL4RenderPipelineColorAttachmentDescriptorArrayClass = _MTL4RenderPipelineColorAttachmentDescriptorArrayClass{objc.GetClass("MTL4RenderPipelineColorAttachmentDescriptorArray")}
	})
	return MTL4RenderPipelineColorAttachmentDescriptorArrayClass
}

type _MTL4RenderPipelineColorAttachmentDescriptorArrayClass struct {
	class objc.Class
}

// An interface definition for the [MTL4RenderPipelineColorAttachmentDescriptorArray] class.
type IMTL4RenderPipelineColorAttachmentDescriptorArray interface {
	objectivec.IObject
	Reset()
	SetObjectAtIndexedSubscript(attachment unsafe.Pointer, attachmentIndex uint)
	ObjectAtIndexedSubscript(attachmentIndex uint) unsafe.Pointer
}

// An array of color attachment descriptions for a render pipeline.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptorArray
type MTL4RenderPipelineColorAttachmentDescriptorArray struct {
	objectivec.Object
}

// MTL4RenderPipelineColorAttachmentDescriptorArrayFrom constructs a [MTL4RenderPipelineColorAttachmentDescriptorArray] from an unsafe.Pointer.
//
// An array of color attachment descriptions for a render pipeline.
func MTL4RenderPipelineColorAttachmentDescriptorArrayFrom(ptr unsafe.Pointer) MTL4RenderPipelineColorAttachmentDescriptorArray {
	return MTL4RenderPipelineColorAttachmentDescriptorArray{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTL4RenderPipelineColorAttachmentDescriptorArrayClass) Alloc() MTL4RenderPipelineColorAttachmentDescriptorArray {
	rv := objc.Send[MTL4RenderPipelineColorAttachmentDescriptorArray](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTL4RenderPipelineColorAttachmentDescriptorArrayClass) New() MTL4RenderPipelineColorAttachmentDescriptorArray {
	rv := objc.Send[MTL4RenderPipelineColorAttachmentDescriptorArray](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTL4RenderPipelineColorAttachmentDescriptorArray) Init() MTL4RenderPipelineColorAttachmentDescriptorArray {
	rv := objc.Send[MTL4RenderPipelineColorAttachmentDescriptorArray](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTL4RenderPipelineColorAttachmentDescriptorArray) Autorelease() MTL4RenderPipelineColorAttachmentDescriptorArray {
	rv := objc.Send[MTL4RenderPipelineColorAttachmentDescriptorArray](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4RenderPipelineColorAttachmentDescriptorArray creates a new MTL4RenderPipelineColorAttachmentDescriptorArray instance.
func NewMTL4RenderPipelineColorAttachmentDescriptorArray() MTL4RenderPipelineColorAttachmentDescriptorArray {
	return getMTL4RenderPipelineColorAttachmentDescriptorArrayClass().New()
}


// Resets the elements of the descriptor array
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptorArray/reset()
func (m_ MTL4RenderPipelineColorAttachmentDescriptorArray) Reset() {
	objc.Send[objc.ID](m_.ID, objc.Sel("reset"))
}

// Sets an attachment at an index.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptorArray/setObject:atIndexedSubscript:
func (m_ MTL4RenderPipelineColorAttachmentDescriptorArray) SetObjectAtIndexedSubscript(attachment unsafe.Pointer, attachmentIndex uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setObject:atIndexedSubscript:"), attachment, attachmentIndex)
}

// Accesses a color attachment at a specific index.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptorArray/subscript(_:)
func (m_ MTL4RenderPipelineColorAttachmentDescriptorArray) ObjectAtIndexedSubscript(attachmentIndex uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("objectAtIndexedSubscript:"), attachmentIndex)
	return rv
}



