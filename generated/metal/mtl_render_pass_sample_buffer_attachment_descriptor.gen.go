// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLRenderPassSampleBufferAttachmentDescriptor */


/* debug [class_header]: Header for MTLRenderPassSampleBufferAttachmentDescriptor */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for RenderPassSampleBufferAttachmentDescriptor */
// An interface definition for the [RenderPassSampleBufferAttachmentDescriptor] class.
type IRenderPassSampleBufferAttachmentDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for RenderPassSampleBufferAttachmentDescriptor */
	// properties:
	EndOfFragmentSampleIndex() uint
	SetEndOfFragmentSampleIndex(value uint)
	EndOfVertexSampleIndex() uint
	SetEndOfVertexSampleIndex(value uint)
	SampleBuffer() unsafe.Pointer
	SetSampleBuffer(value unsafe.Pointer)
	StartOfFragmentSampleIndex() uint
	SetStartOfFragmentSampleIndex(value uint)
	StartOfVertexSampleIndex() uint
	SetStartOfVertexSampleIndex(value uint)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for RenderPassSampleBufferAttachmentDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for RenderPassSampleBufferAttachmentDescriptor */
// Alloc allocates a new instance without initialization.
func (rc _RenderPassSampleBufferAttachmentDescriptorClass) Alloc() RenderPassSampleBufferAttachmentDescriptor {
	rv := objc.Send[RenderPassSampleBufferAttachmentDescriptor](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for RenderPassSampleBufferAttachmentDescriptor */
// A description of where to store GPU counter information at the start and end of a render pass.


// A description of where to store GPU counter information at the start and end of a render pass.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for RenderPassSampleBufferAttachmentDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for RenderPassSampleBufferAttachmentDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for RenderPassSampleBufferAttachmentDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for RenderPassSampleBufferAttachmentDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for RenderPassSampleBufferAttachmentDescriptor */

// The index the Metal device object should use to store GPU counters when ending the render pass’s fragment stage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPassSampleBufferAttachmentDescriptor/endOfFragmentSampleIndex
func (r_ RenderPassSampleBufferAttachmentDescriptor) EndOfFragmentSampleIndex() uint {
	rv := objc.Send[uint](r_.ID, objc.Sel("endOfFragmentSampleIndex"))
	return rv
}/* debug [instance_properties/getter]: endOfFragmentSampleIndex */


// The index the Metal device object should use to store GPU counters when ending the render pass’s fragment stage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPassSampleBufferAttachmentDescriptor/endOfFragmentSampleIndex
func (r_ RenderPassSampleBufferAttachmentDescriptor) SetEndOfFragmentSampleIndex(value uint) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setEndOfFragmentSampleIndex:"), value)
}/* debug [instance_properties/setter]: endOfFragmentSampleIndex */


// The index the Metal device object should use to store GPU counters when ending the render pass’s vertex stage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPassSampleBufferAttachmentDescriptor/endOfVertexSampleIndex
func (r_ RenderPassSampleBufferAttachmentDescriptor) EndOfVertexSampleIndex() uint {
	rv := objc.Send[uint](r_.ID, objc.Sel("endOfVertexSampleIndex"))
	return rv
}/* debug [instance_properties/getter]: endOfVertexSampleIndex */


// The index the Metal device object should use to store GPU counters when ending the render pass’s vertex stage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPassSampleBufferAttachmentDescriptor/endOfVertexSampleIndex
func (r_ RenderPassSampleBufferAttachmentDescriptor) SetEndOfVertexSampleIndex(value uint) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setEndOfVertexSampleIndex:"), value)
}/* debug [instance_properties/setter]: endOfVertexSampleIndex */


// A specialized memory buffer that the GPU uses to store its counter data during the render pass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPassSampleBufferAttachmentDescriptor/sampleBuffer
func (r_ RenderPassSampleBufferAttachmentDescriptor) SampleBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("sampleBuffer"))
	return rv
}/* debug [instance_properties/getter]: sampleBuffer */


// A specialized memory buffer that the GPU uses to store its counter data during the render pass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPassSampleBufferAttachmentDescriptor/sampleBuffer
func (r_ RenderPassSampleBufferAttachmentDescriptor) SetSampleBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setSampleBuffer:"), value)
}/* debug [instance_properties/setter]: sampleBuffer */


// The index the Metal device object should use to store GPU counters when starting the render pass’s fragment stage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPassSampleBufferAttachmentDescriptor/startOfFragmentSampleIndex
func (r_ RenderPassSampleBufferAttachmentDescriptor) StartOfFragmentSampleIndex() uint {
	rv := objc.Send[uint](r_.ID, objc.Sel("startOfFragmentSampleIndex"))
	return rv
}/* debug [instance_properties/getter]: startOfFragmentSampleIndex */


// The index the Metal device object should use to store GPU counters when starting the render pass’s fragment stage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPassSampleBufferAttachmentDescriptor/startOfFragmentSampleIndex
func (r_ RenderPassSampleBufferAttachmentDescriptor) SetStartOfFragmentSampleIndex(value uint) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setStartOfFragmentSampleIndex:"), value)
}/* debug [instance_properties/setter]: startOfFragmentSampleIndex */


// The index the Metal device object should use to store GPU counters when starting the render pass’s vertex stage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPassSampleBufferAttachmentDescriptor/startOfVertexSampleIndex
func (r_ RenderPassSampleBufferAttachmentDescriptor) StartOfVertexSampleIndex() uint {
	rv := objc.Send[uint](r_.ID, objc.Sel("startOfVertexSampleIndex"))
	return rv
}/* debug [instance_properties/getter]: startOfVertexSampleIndex */


// The index the Metal device object should use to store GPU counters when starting the render pass’s vertex stage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPassSampleBufferAttachmentDescriptor/startOfVertexSampleIndex
func (r_ RenderPassSampleBufferAttachmentDescriptor) SetStartOfVertexSampleIndex(value uint) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setStartOfVertexSampleIndex:"), value)
}/* debug [instance_properties/setter]: startOfVertexSampleIndex */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLRenderPassSampleBufferAttachmentDescriptor */



