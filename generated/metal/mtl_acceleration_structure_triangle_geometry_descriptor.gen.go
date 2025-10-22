// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AccelerationStructureTriangleGeometryDescriptor] class.
var (
	AccelerationStructureTriangleGeometryDescriptorClass     _AccelerationStructureTriangleGeometryDescriptorClass
	AccelerationStructureTriangleGeometryDescriptorClassOnce sync.Once
)

func getAccelerationStructureTriangleGeometryDescriptorClass() _AccelerationStructureTriangleGeometryDescriptorClass {
	AccelerationStructureTriangleGeometryDescriptorClassOnce.Do(func() {
		AccelerationStructureTriangleGeometryDescriptorClass = _AccelerationStructureTriangleGeometryDescriptorClass{objc.GetClass("MTLAccelerationStructureTriangleGeometryDescriptor")}
	})
	return AccelerationStructureTriangleGeometryDescriptorClass
}

type _AccelerationStructureTriangleGeometryDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [AccelerationStructureTriangleGeometryDescriptor] class.
type IAccelerationStructureTriangleGeometryDescriptor interface {
	IAccelerationStructureGeometryDescriptor
	IndexBuffer() unsafe.Pointer
	SetIndexBuffer(value unsafe.Pointer)
	IndexBufferOffset() int
	SetIndexBufferOffset(value int)
	IndexType() IndexType
	SetIndexType(value IndexType)
	TransformationMatrixBuffer() unsafe.Pointer
	SetTransformationMatrixBuffer(value unsafe.Pointer)
	TransformationMatrixBufferOffset() int
	SetTransformationMatrixBufferOffset(value int)
	TransformationMatrixLayout() unsafe.Pointer
	SetTransformationMatrixLayout(value unsafe.Pointer)
	TriangleCount() int
	SetTriangleCount(value int)
	VertexBuffer() unsafe.Pointer
	SetVertexBuffer(value unsafe.Pointer)
	VertexBufferOffset() int
	SetVertexBufferOffset(value int)
	VertexFormat() unsafe.Pointer
	SetVertexFormat(value unsafe.Pointer)
	VertexStride() int
	SetVertexStride(value int)
}

// A description of a list of triangle primitives to turn into an acceleration structure.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureTriangleGeometryDescriptor
type AccelerationStructureTriangleGeometryDescriptor struct {
	AccelerationStructureGeometryDescriptor
}

// AccelerationStructureTriangleGeometryDescriptorFrom constructs a [AccelerationStructureTriangleGeometryDescriptor] from an unsafe.Pointer.
//
// A description of a list of triangle primitives to turn into an acceleration structure.
func AccelerationStructureTriangleGeometryDescriptorFrom(ptr unsafe.Pointer) AccelerationStructureTriangleGeometryDescriptor {
	return AccelerationStructureTriangleGeometryDescriptor{
		AccelerationStructureGeometryDescriptor: AccelerationStructureGeometryDescriptorFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AccelerationStructureTriangleGeometryDescriptorClass) Alloc() AccelerationStructureTriangleGeometryDescriptor {
	rv := objc.Send[AccelerationStructureTriangleGeometryDescriptor](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AccelerationStructureTriangleGeometryDescriptorClass) New() AccelerationStructureTriangleGeometryDescriptor {
	rv := objc.Send[AccelerationStructureTriangleGeometryDescriptor](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AccelerationStructureTriangleGeometryDescriptor) Init() AccelerationStructureTriangleGeometryDescriptor {
	rv := objc.Send[AccelerationStructureTriangleGeometryDescriptor](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AccelerationStructureTriangleGeometryDescriptor) Autorelease() AccelerationStructureTriangleGeometryDescriptor {
	rv := objc.Send[AccelerationStructureTriangleGeometryDescriptor](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAccelerationStructureTriangleGeometryDescriptor creates a new AccelerationStructureTriangleGeometryDescriptor instance.
func NewAccelerationStructureTriangleGeometryDescriptor() AccelerationStructureTriangleGeometryDescriptor {
	return getAccelerationStructureTriangleGeometryDescriptorClass().New()
}


// A buffer that contains indices for the vertices that compose the triangle list.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuretrianglegeometrydescriptor/indexbuffer
func (a_ AccelerationStructureTriangleGeometryDescriptor) IndexBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("indexBuffer"))
	return rv
}


// SetIndexBuffer sets the value of the indexBuffer property.
// A buffer that contains indices for the vertices that compose the triangle list.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuretrianglegeometrydescriptor/indexbuffer
func (a_ AccelerationStructureTriangleGeometryDescriptor) SetIndexBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIndexBuffer:"), value)
}

// The offset, in bytes, to the first index in the buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuretrianglegeometrydescriptor/indexbufferoffset
func (a_ AccelerationStructureTriangleGeometryDescriptor) IndexBufferOffset() int {
	rv := objc.Send[int](a_.ID, objc.Sel("indexBufferOffset"))
	return rv
}


// SetIndexBufferOffset sets the value of the indexBufferOffset property.
// The offset, in bytes, to the first index in the buffer.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuretrianglegeometrydescriptor/indexbufferoffset
func (a_ AccelerationStructureTriangleGeometryDescriptor) SetIndexBufferOffset(value int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIndexBufferOffset:"), value)
}

// The data type of indices in the index buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuretrianglegeometrydescriptor/indextype
func (a_ AccelerationStructureTriangleGeometryDescriptor) IndexType() IndexType {
	rv := objc.Send[IndexType](a_.ID, objc.Sel("indexType"))
	return rv
}


// SetIndexType sets the value of the indexType property.
// The data type of indices in the index buffer.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuretrianglegeometrydescriptor/indextype
func (a_ AccelerationStructureTriangleGeometryDescriptor) SetIndexType(value IndexType) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIndexType:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuretrianglegeometrydescriptor/transformationmatrixbuffer
func (a_ AccelerationStructureTriangleGeometryDescriptor) TransformationMatrixBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("transformationMatrixBuffer"))
	return rv
}


// SetTransformationMatrixBuffer sets the value of the transformationMatrixBuffer property.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuretrianglegeometrydescriptor/transformationmatrixbuffer
func (a_ AccelerationStructureTriangleGeometryDescriptor) SetTransformationMatrixBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTransformationMatrixBuffer:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuretrianglegeometrydescriptor/transformationmatrixbufferoffset
func (a_ AccelerationStructureTriangleGeometryDescriptor) TransformationMatrixBufferOffset() int {
	rv := objc.Send[int](a_.ID, objc.Sel("transformationMatrixBufferOffset"))
	return rv
}


// SetTransformationMatrixBufferOffset sets the value of the transformationMatrixBufferOffset property.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuretrianglegeometrydescriptor/transformationmatrixbufferoffset
func (a_ AccelerationStructureTriangleGeometryDescriptor) SetTransformationMatrixBufferOffset(value int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTransformationMatrixBufferOffset:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuretrianglegeometrydescriptor/transformationmatrixlayout
func (a_ AccelerationStructureTriangleGeometryDescriptor) TransformationMatrixLayout() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("transformationMatrixLayout"))
	return rv
}


// SetTransformationMatrixLayout sets the value of the transformationMatrixLayout property.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuretrianglegeometrydescriptor/transformationmatrixlayout
func (a_ AccelerationStructureTriangleGeometryDescriptor) SetTransformationMatrixLayout(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTransformationMatrixLayout:"), value)
}

// The number of triangles in the buffers.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuretrianglegeometrydescriptor/trianglecount
func (a_ AccelerationStructureTriangleGeometryDescriptor) TriangleCount() int {
	rv := objc.Send[int](a_.ID, objc.Sel("triangleCount"))
	return rv
}


// SetTriangleCount sets the value of the triangleCount property.
// The number of triangles in the buffers.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuretrianglegeometrydescriptor/trianglecount
func (a_ AccelerationStructureTriangleGeometryDescriptor) SetTriangleCount(value int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTriangleCount:"), value)
}

// A buffer that contains vertex data.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuretrianglegeometrydescriptor/vertexbuffer
func (a_ AccelerationStructureTriangleGeometryDescriptor) VertexBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("vertexBuffer"))
	return rv
}


// SetVertexBuffer sets the value of the vertexBuffer property.
// A buffer that contains vertex data.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuretrianglegeometrydescriptor/vertexbuffer
func (a_ AccelerationStructureTriangleGeometryDescriptor) SetVertexBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setVertexBuffer:"), value)
}

// The offset, in bytes, for the first vertex in the vertex buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuretrianglegeometrydescriptor/vertexbufferoffset
func (a_ AccelerationStructureTriangleGeometryDescriptor) VertexBufferOffset() int {
	rv := objc.Send[int](a_.ID, objc.Sel("vertexBufferOffset"))
	return rv
}


// SetVertexBufferOffset sets the value of the vertexBufferOffset property.
// The offset, in bytes, for the first vertex in the vertex buffer.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuretrianglegeometrydescriptor/vertexbufferoffset
func (a_ AccelerationStructureTriangleGeometryDescriptor) SetVertexBufferOffset(value int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setVertexBufferOffset:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuretrianglegeometrydescriptor/vertexformat
func (a_ AccelerationStructureTriangleGeometryDescriptor) VertexFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("vertexFormat"))
	return rv
}


// SetVertexFormat sets the value of the vertexFormat property.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuretrianglegeometrydescriptor/vertexformat
func (a_ AccelerationStructureTriangleGeometryDescriptor) SetVertexFormat(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setVertexFormat:"), value)
}

// The stride, in bytes, between vertices in the vertex buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuretrianglegeometrydescriptor/vertexstride
func (a_ AccelerationStructureTriangleGeometryDescriptor) VertexStride() int {
	rv := objc.Send[int](a_.ID, objc.Sel("vertexStride"))
	return rv
}


// SetVertexStride sets the value of the vertexStride property.
// The stride, in bytes, between vertices in the vertex buffer.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuretrianglegeometrydescriptor/vertexstride
func (a_ AccelerationStructureTriangleGeometryDescriptor) SetVertexStride(value int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setVertexStride:"), value)
}



