// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTL4RenderPipelineDynamicLinkingDescriptor */


/* debug [class_header]: Header for MTL4RenderPipelineDynamicLinkingDescriptor */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTL4RenderPipelineDynamicLinkingDescriptor */
// An interface definition for the [MTL4RenderPipelineDynamicLinkingDescriptor] class.
type IMTL4RenderPipelineDynamicLinkingDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTL4RenderPipelineDynamicLinkingDescriptor */
	// properties:
	FragmentLinkingDescriptor() IMTL4PipelineStageDynamicLinkingDescriptor
	MeshLinkingDescriptor() IMTL4PipelineStageDynamicLinkingDescriptor
	ObjectLinkingDescriptor() IMTL4PipelineStageDynamicLinkingDescriptor
	TileLinkingDescriptor() IMTL4PipelineStageDynamicLinkingDescriptor
	VertexLinkingDescriptor() IMTL4PipelineStageDynamicLinkingDescriptor
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTL4RenderPipelineDynamicLinkingDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTL4RenderPipelineDynamicLinkingDescriptor */
// Alloc allocates a new instance without initialization.
func (mc _MTL4RenderPipelineDynamicLinkingDescriptorClass) Alloc() MTL4RenderPipelineDynamicLinkingDescriptor {
	rv := objc.Send[MTL4RenderPipelineDynamicLinkingDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTL4RenderPipelineDynamicLinkingDescriptor */
// Groups together properties that provide linking properties for render pipelines.


// Groups together properties that provide linking properties for render pipelines.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTL4RenderPipelineDynamicLinkingDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTL4RenderPipelineDynamicLinkingDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTL4RenderPipelineDynamicLinkingDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTL4RenderPipelineDynamicLinkingDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTL4RenderPipelineDynamicLinkingDescriptor */

// Controls properties for linking the fragment stage of the render pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineDynamicLinkingDescriptor/fragmentLinkingDescriptor
func (m_ MTL4RenderPipelineDynamicLinkingDescriptor) FragmentLinkingDescriptor() IMTL4PipelineStageDynamicLinkingDescriptor {
	rv := objc.Send[MTL4PipelineStageDynamicLinkingDescriptor](m_.ID, objc.Sel("fragmentLinkingDescriptor"))
	return rv
}/* debug [instance_properties/getter]: fragmentLinkingDescriptor */


// Controls properties for linking the mesh stage of the render pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineDynamicLinkingDescriptor/meshLinkingDescriptor
func (m_ MTL4RenderPipelineDynamicLinkingDescriptor) MeshLinkingDescriptor() IMTL4PipelineStageDynamicLinkingDescriptor {
	rv := objc.Send[MTL4PipelineStageDynamicLinkingDescriptor](m_.ID, objc.Sel("meshLinkingDescriptor"))
	return rv
}/* debug [instance_properties/getter]: meshLinkingDescriptor */


// Controls properties for link the object stage of the render pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineDynamicLinkingDescriptor/objectLinkingDescriptor
func (m_ MTL4RenderPipelineDynamicLinkingDescriptor) ObjectLinkingDescriptor() IMTL4PipelineStageDynamicLinkingDescriptor {
	rv := objc.Send[MTL4PipelineStageDynamicLinkingDescriptor](m_.ID, objc.Sel("objectLinkingDescriptor"))
	return rv
}/* debug [instance_properties/getter]: objectLinkingDescriptor */


// Controls properties for linking the tile stage of the render pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineDynamicLinkingDescriptor/tileLinkingDescriptor
func (m_ MTL4RenderPipelineDynamicLinkingDescriptor) TileLinkingDescriptor() IMTL4PipelineStageDynamicLinkingDescriptor {
	rv := objc.Send[MTL4PipelineStageDynamicLinkingDescriptor](m_.ID, objc.Sel("tileLinkingDescriptor"))
	return rv
}/* debug [instance_properties/getter]: tileLinkingDescriptor */


// Controls properties for linking the vertex stage of the render pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineDynamicLinkingDescriptor/vertexLinkingDescriptor
func (m_ MTL4RenderPipelineDynamicLinkingDescriptor) VertexLinkingDescriptor() IMTL4PipelineStageDynamicLinkingDescriptor {
	rv := objc.Send[MTL4PipelineStageDynamicLinkingDescriptor](m_.ID, objc.Sel("vertexLinkingDescriptor"))
	return rv
}/* debug [instance_properties/getter]: vertexLinkingDescriptor */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTL4RenderPipelineDynamicLinkingDescriptor */



