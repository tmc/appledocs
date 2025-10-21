// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [RenderPassSampleBufferAttachmentDescriptor] class.
var (
	RenderPassSampleBufferAttachmentDescriptorClass     _RenderPassSampleBufferAttachmentDescriptorClass
	RenderPassSampleBufferAttachmentDescriptorClassOnce sync.Once
)

func getRenderPassSampleBufferAttachmentDescriptorClass() _RenderPassSampleBufferAttachmentDescriptorClass {
	RenderPassSampleBufferAttachmentDescriptorClassOnce.Do(func() {
		RenderPassSampleBufferAttachmentDescriptorClass = _RenderPassSampleBufferAttachmentDescriptorClass{objc.GetClass("MTLRenderPassSampleBufferAttachmentDescriptor")}
	})
	return RenderPassSampleBufferAttachmentDescriptorClass
}

type _RenderPassSampleBufferAttachmentDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [RenderPassSampleBufferAttachmentDescriptor] class.
type IRenderPassSampleBufferAttachmentDescriptor interface {
	objectivec.IObject
}

// A description of where to store GPU counter information at the start and end of a render pass.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPassSampleBufferAttachmentDescriptor
type RenderPassSampleBufferAttachmentDescriptor struct {
	objectivec.Object
}

// RenderPassSampleBufferAttachmentDescriptorFrom constructs a [RenderPassSampleBufferAttachmentDescriptor] from an unsafe.Pointer.
//
// A description of where to store GPU counter information at the start and end of a render pass.
func RenderPassSampleBufferAttachmentDescriptorFrom(ptr unsafe.Pointer) RenderPassSampleBufferAttachmentDescriptor {
	return RenderPassSampleBufferAttachmentDescriptor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (rc _RenderPassSampleBufferAttachmentDescriptorClass) Alloc() RenderPassSampleBufferAttachmentDescriptor {
	rv := objc.Send[RenderPassSampleBufferAttachmentDescriptor](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _RenderPassSampleBufferAttachmentDescriptorClass) New() RenderPassSampleBufferAttachmentDescriptor {
	rv := objc.Send[RenderPassSampleBufferAttachmentDescriptor](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RenderPassSampleBufferAttachmentDescriptor) Init() RenderPassSampleBufferAttachmentDescriptor {
	rv := objc.Send[RenderPassSampleBufferAttachmentDescriptor](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RenderPassSampleBufferAttachmentDescriptor) Autorelease() RenderPassSampleBufferAttachmentDescriptor {
	rv := objc.Send[RenderPassSampleBufferAttachmentDescriptor](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRenderPassSampleBufferAttachmentDescriptor creates a new RenderPassSampleBufferAttachmentDescriptor instance.
func NewRenderPassSampleBufferAttachmentDescriptor() RenderPassSampleBufferAttachmentDescriptor {
	return getRenderPassSampleBufferAttachmentDescriptorClass().New()
}


// The index the Metal device object should use to store GPU counters when ending the render pass’s fragment stage.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPassSampleBufferAttachmentDescriptor/endOfFragmentSampleIndex
func (r_ RenderPassSampleBufferAttachmentDescriptor) EndOfFragmentSampleIndex() uint {
	rv := objc.Send[uint](r_.ID, objc.Sel("endOfFragmentSampleIndex"))
	return rv
}


// SetEndOfFragmentSampleIndex sets the value of the endOfFragmentSampleIndex property.
// The index the Metal device object should use to store GPU counters when ending the render pass’s fragment stage.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPassSampleBufferAttachmentDescriptor/endOfFragmentSampleIndex
func (r_ RenderPassSampleBufferAttachmentDescriptor) SetEndOfFragmentSampleIndex(value uint) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setEndOfFragmentSampleIndex:"), value)
}
// The index the Metal device object should use to store GPU counters when ending the render pass’s vertex stage.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPassSampleBufferAttachmentDescriptor/endOfVertexSampleIndex
func (r_ RenderPassSampleBufferAttachmentDescriptor) EndOfVertexSampleIndex() uint {
	rv := objc.Send[uint](r_.ID, objc.Sel("endOfVertexSampleIndex"))
	return rv
}


// SetEndOfVertexSampleIndex sets the value of the endOfVertexSampleIndex property.
// The index the Metal device object should use to store GPU counters when ending the render pass’s vertex stage.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPassSampleBufferAttachmentDescriptor/endOfVertexSampleIndex
func (r_ RenderPassSampleBufferAttachmentDescriptor) SetEndOfVertexSampleIndex(value uint) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setEndOfVertexSampleIndex:"), value)
}
// A specialized memory buffer that the GPU uses to store its counter data during the render pass.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPassSampleBufferAttachmentDescriptor/sampleBuffer
func (r_ RenderPassSampleBufferAttachmentDescriptor) SampleBuffer() objc.ID {
	rv := objc.Send[objc.ID](r_.ID, objc.Sel("sampleBuffer"))
	return rv
}


// SetSampleBuffer sets the value of the sampleBuffer property.
// A specialized memory buffer that the GPU uses to store its counter data during the render pass.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPassSampleBufferAttachmentDescriptor/sampleBuffer
func (r_ RenderPassSampleBufferAttachmentDescriptor) SetSampleBuffer(value objc.ID) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setSampleBuffer:"), value)
}
// The index the Metal device object should use to store GPU counters when starting the render pass’s fragment stage.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPassSampleBufferAttachmentDescriptor/startOfFragmentSampleIndex
func (r_ RenderPassSampleBufferAttachmentDescriptor) StartOfFragmentSampleIndex() uint {
	rv := objc.Send[uint](r_.ID, objc.Sel("startOfFragmentSampleIndex"))
	return rv
}


// SetStartOfFragmentSampleIndex sets the value of the startOfFragmentSampleIndex property.
// The index the Metal device object should use to store GPU counters when starting the render pass’s fragment stage.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPassSampleBufferAttachmentDescriptor/startOfFragmentSampleIndex
func (r_ RenderPassSampleBufferAttachmentDescriptor) SetStartOfFragmentSampleIndex(value uint) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setStartOfFragmentSampleIndex:"), value)
}
// The index the Metal device object should use to store GPU counters when starting the render pass’s vertex stage.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPassSampleBufferAttachmentDescriptor/startOfVertexSampleIndex
func (r_ RenderPassSampleBufferAttachmentDescriptor) StartOfVertexSampleIndex() uint {
	rv := objc.Send[uint](r_.ID, objc.Sel("startOfVertexSampleIndex"))
	return rv
}


// SetStartOfVertexSampleIndex sets the value of the startOfVertexSampleIndex property.
// The index the Metal device object should use to store GPU counters when starting the render pass’s vertex stage.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPassSampleBufferAttachmentDescriptor/startOfVertexSampleIndex
func (r_ RenderPassSampleBufferAttachmentDescriptor) SetStartOfVertexSampleIndex(value uint) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setStartOfVertexSampleIndex:"), value)
}


