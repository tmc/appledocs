// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [VertexDescriptor] class.
var (
	VertexDescriptorClass     _VertexDescriptorClass
	VertexDescriptorClassOnce sync.Once
)

func getVertexDescriptorClass() _VertexDescriptorClass {
	VertexDescriptorClassOnce.Do(func() {
		VertexDescriptorClass = _VertexDescriptorClass{objc.GetClass("MTLVertexDescriptor")}
	})
	return VertexDescriptorClass
}

type _VertexDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [VertexDescriptor] class.
type IVertexDescriptor interface {
	objectivec.IObject
	// properties:
	MTLBufferLayoutStrideDynamic() int
	VertexDescriptor() IMTLVertexDescriptor
	SetVertexDescriptor(value IMTLVertexDescriptor)
	Attributes() objc.IObject /* cross-framework: VertexAttributeDescriptorArray */
	SetAttributes(value objc.IObject /* cross-framework: VertexAttributeDescriptorArray */)
	Layouts() objc.IObject /* cross-framework: VertexBufferLayoutDescriptorArray */
	SetLayouts(value objc.IObject /* cross-framework: VertexBufferLayoutDescriptorArray */)
	// methods:
}

// An object that describes how to organize and map data to a vertex function.
//
// A object is used to configure how vertex data stored in memory is mapped to attributes in a vertex shader. A pipeline state is the state of the graphics rendering pipeline, including shaders, blending, multisampling, and visibility testing. For every pipeline state, there can be only one object. When you configure a object to create this pipeline state, you use a object to establish the vertex layout for the function associated with the pipeline. Create and configure a object, then use this object to set the property of the object.


// An object that describes how to organize and map data to a vertex function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexDescriptor
type VertexDescriptor struct {
	objectivec.Object
}

// VertexDescriptorFrom constructs a [VertexDescriptor] from an unsafe.Pointer.
//
// An object that describes how to organize and map data to a vertex function.
func VertexDescriptorFrom(ptr unsafe.Pointer) VertexDescriptor {
	return VertexDescriptor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (vc _VertexDescriptorClass) Alloc() VertexDescriptor {
	rv := objc.Send[VertexDescriptor](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VertexDescriptorClass) New() VertexDescriptor {
	rv := objc.Send[VertexDescriptor](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VertexDescriptor) Init() VertexDescriptor {
	rv := objc.Send[VertexDescriptor](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VertexDescriptor) Autorelease() VertexDescriptor {
	rv := objc.Send[VertexDescriptor](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVertexDescriptor creates a new VertexDescriptor instance.
func NewVertexDescriptor() VertexDescriptor {
	return getVertexDescriptorClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlbufferlayoutstridedynamic
func (v_ VertexDescriptor) MTLBufferLayoutStrideDynamic() int {
	rv := objc.Send[int](v_.ID, objc.Sel("MTLBufferLayoutStrideDynamic"))
	return rv
}


// The organization of vertex data in an attribute’s argument table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinedescriptor/vertexdescriptor
func (v_ VertexDescriptor) VertexDescriptor() IMTLVertexDescriptor {
	rv := objc.Send[VertexDescriptor](v_.ID, objc.Sel("vertexDescriptor"))
	return rv
}


// The organization of vertex data in an attribute’s argument table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinedescriptor/vertexdescriptor
func (v_ VertexDescriptor) SetVertexDescriptor(value IMTLVertexDescriptor) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setVertexDescriptor:"), value)
}


// An array of state data that describes how vertex attribute data is stored in memory and is mapped to arguments for a vertex shader function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlvertexdescriptor/attributes
func (v_ VertexDescriptor) Attributes() objc.IObject /* cross-framework: VertexAttributeDescriptorArray */ {
	rv := objc.Send[VertexAttributeDescriptorArray](v_.ID, objc.Sel("attributes"))
	return rv
}


// An array of state data that describes how vertex attribute data is stored in memory and is mapped to arguments for a vertex shader function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlvertexdescriptor/attributes
func (v_ VertexDescriptor) SetAttributes(value objc.IObject /* cross-framework: VertexAttributeDescriptorArray */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setAttributes:"), value)
}


// An array of state data that describes how data are fetched by a vertex shader function when rendering primitives.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlvertexdescriptor/layouts
func (v_ VertexDescriptor) Layouts() objc.IObject /* cross-framework: VertexBufferLayoutDescriptorArray */ {
	rv := objc.Send[VertexBufferLayoutDescriptorArray](v_.ID, objc.Sel("layouts"))
	return rv
}


// An array of state data that describes how data are fetched by a vertex shader function when rendering primitives.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlvertexdescriptor/layouts
func (v_ VertexDescriptor) SetLayouts(value objc.IObject /* cross-framework: VertexBufferLayoutDescriptorArray */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setLayouts:"), value)
}




