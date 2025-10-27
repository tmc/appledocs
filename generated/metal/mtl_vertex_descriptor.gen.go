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
	Attributes() IMTLVertexAttributeDescriptorArray
	Layouts() IMTLVertexBufferLayoutDescriptorArray
	MTLBufferLayoutStrideDynamic() int
	VertexDescriptor() IMTLVertexDescriptor
	SetVertexDescriptor(value IMTLVertexDescriptor)


	

	// methods:
	Reset()


}





// Alloc allocates a new instance without initialization.
func (vc _VertexDescriptorClass) Alloc() VertexDescriptor {
	rv := objc.Send[VertexDescriptor](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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





// An instance that describes how to organize and map data to a vertex function.
//
// An instance is used to configure how vertex data stored in memory is mapped to attributes in a vertex shader. A pipeline state is the state of the graphics rendering pipeline, including shaders, blending, multisampling, and visibility testing. For every pipeline state, there can be only one instance. When you configure an instance to create this pipeline state, you use an instance to establish the vertex layout for the function associated with the pipeline. Create and configure an instance, then use this instance to set the property of the instance.


// An instance that describes how to organize and map data to a vertex function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexDescriptor
type VertexDescriptor struct {
	objectivec.Object
}

// VertexDescriptorFrom constructs a [VertexDescriptor] from an unsafe.Pointer.
//
// An instance that describes how to organize and map data to a vertex function.
func VertexDescriptorFrom(ptr unsafe.Pointer) VertexDescriptor {
	return VertexDescriptor{objectivec.Object{objc.ID(ptr)}}
}










// Creates and returns a new vertex descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexDescriptor/vertexDescriptor
func (vc _VertexDescriptorClass) VertexDescriptor() IVertexDescriptor {
	rv := objc.Send[VertexDescriptor](objc.ID(vc.class), objc.Sel("vertexDescriptor"))
	return rv
}












// Resets the default state for the vertex descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexDescriptor/reset()
func (v_ VertexDescriptor) Reset() {
	objc.Send[objc.ID](v_.ID, objc.Sel("reset"))
}







// An array of state data that describes how vertex attribute data is stored in memory and is mapped to arguments for a vertex shader function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexDescriptor/attributes
func (v_ VertexDescriptor) Attributes() IMTLVertexAttributeDescriptorArray {
	rv := objc.Send[VertexAttributeDescriptorArray](v_.ID, objc.Sel("attributes"))
	return rv
}


// An array of state data that describes how data are fetched by a vertex shader function when rendering primitives.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexDescriptor/layouts
func (v_ VertexDescriptor) Layouts() IMTLVertexBufferLayoutDescriptorArray {
	rv := objc.Send[VertexBufferLayoutDescriptorArray](v_.ID, objc.Sel("layouts"))
	return rv
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








