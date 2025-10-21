// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [IndirectCommandBufferDescriptor] class.
var (
	IndirectCommandBufferDescriptorClass     _IndirectCommandBufferDescriptorClass
	IndirectCommandBufferDescriptorClassOnce sync.Once
)

func getIndirectCommandBufferDescriptorClass() _IndirectCommandBufferDescriptorClass {
	IndirectCommandBufferDescriptorClassOnce.Do(func() {
		IndirectCommandBufferDescriptorClass = _IndirectCommandBufferDescriptorClass{objc.GetClass("MTLIndirectCommandBufferDescriptor")}
	})
	return IndirectCommandBufferDescriptorClass
}

type _IndirectCommandBufferDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [IndirectCommandBufferDescriptor] class.
type IIndirectCommandBufferDescriptor interface {
	objectivec.IObject
}

// A configuration you create to customize an indirect command buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBufferDescriptor
type IndirectCommandBufferDescriptor struct {
	objectivec.Object
}

// IndirectCommandBufferDescriptorFrom constructs a [IndirectCommandBufferDescriptor] from an unsafe.Pointer.
//
// A configuration you create to customize an indirect command buffer.
func IndirectCommandBufferDescriptorFrom(ptr unsafe.Pointer) IndirectCommandBufferDescriptor {
	return IndirectCommandBufferDescriptor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _IndirectCommandBufferDescriptorClass) Alloc() IndirectCommandBufferDescriptor {
	rv := objc.Send[IndirectCommandBufferDescriptor](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _IndirectCommandBufferDescriptorClass) New() IndirectCommandBufferDescriptor {
	rv := objc.Send[IndirectCommandBufferDescriptor](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ IndirectCommandBufferDescriptor) Init() IndirectCommandBufferDescriptor {
	rv := objc.Send[IndirectCommandBufferDescriptor](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ IndirectCommandBufferDescriptor) Autorelease() IndirectCommandBufferDescriptor {
	rv := objc.Send[IndirectCommandBufferDescriptor](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewIndirectCommandBufferDescriptor creates a new IndirectCommandBufferDescriptor instance.
func NewIndirectCommandBufferDescriptor() IndirectCommandBufferDescriptor {
	return getIndirectCommandBufferDescriptorClass().New()
}


// The set of command types that you can encode into the indirect command buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBufferDescriptor/commandTypes
func (i_ IndirectCommandBufferDescriptor) CommandTypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("commandTypes"))
	return rv
}


// SetCommandTypes sets the value of the commandTypes property.
// The set of command types that you can encode into the indirect command buffer.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBufferDescriptor/commandTypes
func (i_ IndirectCommandBufferDescriptor) SetCommandTypes(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setCommandTypes:"), value)
}
// A Boolean value that determines where commands in the indirect command buffer get their buffer arguments from when you execute them.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBufferDescriptor/inheritBuffers
func (i_ IndirectCommandBufferDescriptor) InheritBuffers() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("inheritBuffers"))
	return rv
}


// SetInheritBuffers sets the value of the inheritBuffers property.
// A Boolean value that determines where commands in the indirect command buffer get their buffer arguments from when you execute them.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBufferDescriptor/inheritBuffers
func (i_ IndirectCommandBufferDescriptor) SetInheritBuffers(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setInheritBuffers:"), value)
}
// Configures whether the indirect command buffer inherits the cull mode from the encoder.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBufferDescriptor/inheritCullMode
func (i_ IndirectCommandBufferDescriptor) InheritCullMode() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("inheritCullMode"))
	return rv
}


// SetInheritCullMode sets the value of the inheritCullMode property.
// Configures whether the indirect command buffer inherits the cull mode from the encoder.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBufferDescriptor/inheritCullMode
func (i_ IndirectCommandBufferDescriptor) SetInheritCullMode(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setInheritCullMode:"), value)
}
// Configures whether the indirect command buffer inherits the depth bias from the encoder.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBufferDescriptor/inheritDepthBias
func (i_ IndirectCommandBufferDescriptor) InheritDepthBias() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("inheritDepthBias"))
	return rv
}


// SetInheritDepthBias sets the value of the inheritDepthBias property.
// Configures whether the indirect command buffer inherits the depth bias from the encoder.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBufferDescriptor/inheritDepthBias
func (i_ IndirectCommandBufferDescriptor) SetInheritDepthBias(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setInheritDepthBias:"), value)
}
// Configures whether the indirect command buffer inherits the depth clip mode from the encoder.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBufferDescriptor/inheritDepthClipMode
func (i_ IndirectCommandBufferDescriptor) InheritDepthClipMode() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("inheritDepthClipMode"))
	return rv
}


// SetInheritDepthClipMode sets the value of the inheritDepthClipMode property.
// Configures whether the indirect command buffer inherits the depth clip mode from the encoder.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBufferDescriptor/inheritDepthClipMode
func (i_ IndirectCommandBufferDescriptor) SetInheritDepthClipMode(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setInheritDepthClipMode:"), value)
}
// Configures whether the indirect command buffer inherits the depth stencil state from the encoder.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBufferDescriptor/inheritDepthStencilState
func (i_ IndirectCommandBufferDescriptor) InheritDepthStencilState() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("inheritDepthStencilState"))
	return rv
}


// SetInheritDepthStencilState sets the value of the inheritDepthStencilState property.
// Configures whether the indirect command buffer inherits the depth stencil state from the encoder.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBufferDescriptor/inheritDepthStencilState
func (i_ IndirectCommandBufferDescriptor) SetInheritDepthStencilState(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setInheritDepthStencilState:"), value)
}
// Configures whether the indirect command buffer inherits the front facing winding from the encoder.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBufferDescriptor/inheritFrontFacingWinding
func (i_ IndirectCommandBufferDescriptor) InheritFrontFacingWinding() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("inheritFrontFacingWinding"))
	return rv
}


// SetInheritFrontFacingWinding sets the value of the inheritFrontFacingWinding property.
// Configures whether the indirect command buffer inherits the front facing winding from the encoder.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBufferDescriptor/inheritFrontFacingWinding
func (i_ IndirectCommandBufferDescriptor) SetInheritFrontFacingWinding(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setInheritFrontFacingWinding:"), value)
}
// A Boolean value that determines where commands in the indirect command buffer get their pipeline state from when you execute them.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBufferDescriptor/inheritPipelineState
func (i_ IndirectCommandBufferDescriptor) InheritPipelineState() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("inheritPipelineState"))
	return rv
}


// SetInheritPipelineState sets the value of the inheritPipelineState property.
// A Boolean value that determines where commands in the indirect command buffer get their pipeline state from when you execute them.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBufferDescriptor/inheritPipelineState
func (i_ IndirectCommandBufferDescriptor) SetInheritPipelineState(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setInheritPipelineState:"), value)
}
// Configures whether the indirect command buffer inherits the triangle fill mode from the encoder.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBufferDescriptor/inheritTriangleFillMode
func (i_ IndirectCommandBufferDescriptor) InheritTriangleFillMode() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("inheritTriangleFillMode"))
	return rv
}


// SetInheritTriangleFillMode sets the value of the inheritTriangleFillMode property.
// Configures whether the indirect command buffer inherits the triangle fill mode from the encoder.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBufferDescriptor/inheritTriangleFillMode
func (i_ IndirectCommandBufferDescriptor) SetInheritTriangleFillMode(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setInheritTriangleFillMode:"), value)
}
// The maximum number of buffers that you can set per command for the fragment stage.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBufferDescriptor/maxFragmentBufferBindCount
func (i_ IndirectCommandBufferDescriptor) MaxFragmentBufferBindCount() uint {
	rv := objc.Send[uint](i_.ID, objc.Sel("maxFragmentBufferBindCount"))
	return rv
}


// SetMaxFragmentBufferBindCount sets the value of the maxFragmentBufferBindCount property.
// The maximum number of buffers that you can set per command for the fragment stage.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBufferDescriptor/maxFragmentBufferBindCount
func (i_ IndirectCommandBufferDescriptor) SetMaxFragmentBufferBindCount(value uint) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMaxFragmentBufferBindCount:"), value)
}
// The maximum number of buffers that you can set per command for the compute kernel.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBufferDescriptor/maxKernelBufferBindCount
func (i_ IndirectCommandBufferDescriptor) MaxKernelBufferBindCount() uint {
	rv := objc.Send[uint](i_.ID, objc.Sel("maxKernelBufferBindCount"))
	return rv
}


// SetMaxKernelBufferBindCount sets the value of the maxKernelBufferBindCount property.
// The maximum number of buffers that you can set per command for the compute kernel.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBufferDescriptor/maxKernelBufferBindCount
func (i_ IndirectCommandBufferDescriptor) SetMaxKernelBufferBindCount(value uint) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMaxKernelBufferBindCount:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBufferDescriptor/maxKernelThreadgroupMemoryBindCount
func (i_ IndirectCommandBufferDescriptor) MaxKernelThreadgroupMemoryBindCount() uint {
	rv := objc.Send[uint](i_.ID, objc.Sel("maxKernelThreadgroupMemoryBindCount"))
	return rv
}


// SetMaxKernelThreadgroupMemoryBindCount sets the value of the maxKernelThreadgroupMemoryBindCount property.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBufferDescriptor/maxKernelThreadgroupMemoryBindCount
func (i_ IndirectCommandBufferDescriptor) SetMaxKernelThreadgroupMemoryBindCount(value uint) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMaxKernelThreadgroupMemoryBindCount:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBufferDescriptor/maxMeshBufferBindCount
func (i_ IndirectCommandBufferDescriptor) MaxMeshBufferBindCount() uint {
	rv := objc.Send[uint](i_.ID, objc.Sel("maxMeshBufferBindCount"))
	return rv
}


// SetMaxMeshBufferBindCount sets the value of the maxMeshBufferBindCount property.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBufferDescriptor/maxMeshBufferBindCount
func (i_ IndirectCommandBufferDescriptor) SetMaxMeshBufferBindCount(value uint) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMaxMeshBufferBindCount:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBufferDescriptor/maxObjectBufferBindCount
func (i_ IndirectCommandBufferDescriptor) MaxObjectBufferBindCount() uint {
	rv := objc.Send[uint](i_.ID, objc.Sel("maxObjectBufferBindCount"))
	return rv
}


// SetMaxObjectBufferBindCount sets the value of the maxObjectBufferBindCount property.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBufferDescriptor/maxObjectBufferBindCount
func (i_ IndirectCommandBufferDescriptor) SetMaxObjectBufferBindCount(value uint) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMaxObjectBufferBindCount:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBufferDescriptor/maxObjectThreadgroupMemoryBindCount
func (i_ IndirectCommandBufferDescriptor) MaxObjectThreadgroupMemoryBindCount() uint {
	rv := objc.Send[uint](i_.ID, objc.Sel("maxObjectThreadgroupMemoryBindCount"))
	return rv
}


// SetMaxObjectThreadgroupMemoryBindCount sets the value of the maxObjectThreadgroupMemoryBindCount property.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBufferDescriptor/maxObjectThreadgroupMemoryBindCount
func (i_ IndirectCommandBufferDescriptor) SetMaxObjectThreadgroupMemoryBindCount(value uint) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMaxObjectThreadgroupMemoryBindCount:"), value)
}
// The maximum number of buffers that you can set per command for the vertex stage.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBufferDescriptor/maxVertexBufferBindCount
func (i_ IndirectCommandBufferDescriptor) MaxVertexBufferBindCount() uint {
	rv := objc.Send[uint](i_.ID, objc.Sel("maxVertexBufferBindCount"))
	return rv
}


// SetMaxVertexBufferBindCount sets the value of the maxVertexBufferBindCount property.
// The maximum number of buffers that you can set per command for the vertex stage.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBufferDescriptor/maxVertexBufferBindCount
func (i_ IndirectCommandBufferDescriptor) SetMaxVertexBufferBindCount(value uint) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMaxVertexBufferBindCount:"), value)
}
// Specifies if the indirect command buffer should support color attachment mapping.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBufferDescriptor/supportColorAttachmentMapping
func (i_ IndirectCommandBufferDescriptor) SupportColorAttachmentMapping() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("supportColorAttachmentMapping"))
	return rv
}


// SetSupportColorAttachmentMapping sets the value of the supportColorAttachmentMapping property.
// Specifies if the indirect command buffer should support color attachment mapping.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBufferDescriptor/supportColorAttachmentMapping
func (i_ IndirectCommandBufferDescriptor) SetSupportColorAttachmentMapping(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSupportColorAttachmentMapping:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBufferDescriptor/supportDynamicAttributeStride
func (i_ IndirectCommandBufferDescriptor) SupportDynamicAttributeStride() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("supportDynamicAttributeStride"))
	return rv
}


// SetSupportDynamicAttributeStride sets the value of the supportDynamicAttributeStride property.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBufferDescriptor/supportDynamicAttributeStride
func (i_ IndirectCommandBufferDescriptor) SetSupportDynamicAttributeStride(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSupportDynamicAttributeStride:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBufferDescriptor/supportRayTracing
func (i_ IndirectCommandBufferDescriptor) SupportRayTracing() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("supportRayTracing"))
	return rv
}


// SetSupportRayTracing sets the value of the supportRayTracing property.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBufferDescriptor/supportRayTracing
func (i_ IndirectCommandBufferDescriptor) SetSupportRayTracing(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSupportRayTracing:"), value)
}


