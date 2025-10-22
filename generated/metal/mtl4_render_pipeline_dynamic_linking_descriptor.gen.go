// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTL4RenderPipelineDynamicLinkingDescriptor] class.
var (
	MTL4RenderPipelineDynamicLinkingDescriptorClass     _MTL4RenderPipelineDynamicLinkingDescriptorClass
	MTL4RenderPipelineDynamicLinkingDescriptorClassOnce sync.Once
)

func getMTL4RenderPipelineDynamicLinkingDescriptorClass() _MTL4RenderPipelineDynamicLinkingDescriptorClass {
	MTL4RenderPipelineDynamicLinkingDescriptorClassOnce.Do(func() {
		MTL4RenderPipelineDynamicLinkingDescriptorClass = _MTL4RenderPipelineDynamicLinkingDescriptorClass{objc.GetClass("MTL4RenderPipelineDynamicLinkingDescriptor")}
	})
	return MTL4RenderPipelineDynamicLinkingDescriptorClass
}

type _MTL4RenderPipelineDynamicLinkingDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [MTL4RenderPipelineDynamicLinkingDescriptor] class.
type IMTL4RenderPipelineDynamicLinkingDescriptor interface {
	objectivec.IObject
	FragmentLinkingDescriptor() unsafe.Pointer
	MeshLinkingDescriptor() unsafe.Pointer
	ObjectLinkingDescriptor() unsafe.Pointer
	TileLinkingDescriptor() unsafe.Pointer
	VertexLinkingDescriptor() unsafe.Pointer
}

// Groups together properties that provide linking properties for render pipelines.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineDynamicLinkingDescriptor
type MTL4RenderPipelineDynamicLinkingDescriptor struct {
	objectivec.Object
}

// MTL4RenderPipelineDynamicLinkingDescriptorFrom constructs a [MTL4RenderPipelineDynamicLinkingDescriptor] from an unsafe.Pointer.
//
// Groups together properties that provide linking properties for render pipelines.
func MTL4RenderPipelineDynamicLinkingDescriptorFrom(ptr unsafe.Pointer) MTL4RenderPipelineDynamicLinkingDescriptor {
	return MTL4RenderPipelineDynamicLinkingDescriptor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTL4RenderPipelineDynamicLinkingDescriptorClass) Alloc() MTL4RenderPipelineDynamicLinkingDescriptor {
	rv := objc.Send[MTL4RenderPipelineDynamicLinkingDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTL4RenderPipelineDynamicLinkingDescriptorClass) New() MTL4RenderPipelineDynamicLinkingDescriptor {
	rv := objc.Send[MTL4RenderPipelineDynamicLinkingDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTL4RenderPipelineDynamicLinkingDescriptor) Init() MTL4RenderPipelineDynamicLinkingDescriptor {
	rv := objc.Send[MTL4RenderPipelineDynamicLinkingDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTL4RenderPipelineDynamicLinkingDescriptor) Autorelease() MTL4RenderPipelineDynamicLinkingDescriptor {
	rv := objc.Send[MTL4RenderPipelineDynamicLinkingDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4RenderPipelineDynamicLinkingDescriptor creates a new MTL4RenderPipelineDynamicLinkingDescriptor instance.
func NewMTL4RenderPipelineDynamicLinkingDescriptor() MTL4RenderPipelineDynamicLinkingDescriptor {
	return getMTL4RenderPipelineDynamicLinkingDescriptorClass().New()
}


// Controls properties for linking the fragment stage of the render pipeline.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineDynamicLinkingDescriptor/fragmentLinkingDescriptor
func (m_ MTL4RenderPipelineDynamicLinkingDescriptor) FragmentLinkingDescriptor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("fragmentLinkingDescriptor"))
	return rv
}

// Controls properties for linking the mesh stage of the render pipeline.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineDynamicLinkingDescriptor/meshLinkingDescriptor
func (m_ MTL4RenderPipelineDynamicLinkingDescriptor) MeshLinkingDescriptor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("meshLinkingDescriptor"))
	return rv
}

// Controls properties for link the object stage of the render pipeline.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineDynamicLinkingDescriptor/objectLinkingDescriptor
func (m_ MTL4RenderPipelineDynamicLinkingDescriptor) ObjectLinkingDescriptor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("objectLinkingDescriptor"))
	return rv
}

// Controls properties for linking the tile stage of the render pipeline.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineDynamicLinkingDescriptor/tileLinkingDescriptor
func (m_ MTL4RenderPipelineDynamicLinkingDescriptor) TileLinkingDescriptor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("tileLinkingDescriptor"))
	return rv
}

// Controls properties for linking the vertex stage of the render pipeline.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineDynamicLinkingDescriptor/vertexLinkingDescriptor
func (m_ MTL4RenderPipelineDynamicLinkingDescriptor) VertexLinkingDescriptor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("vertexLinkingDescriptor"))
	return rv
}



