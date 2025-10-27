// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [MTL4AccelerationStructureTriangleGeometryDescriptor] class.
var (
	MTL4AccelerationStructureTriangleGeometryDescriptorClass     _MTL4AccelerationStructureTriangleGeometryDescriptorClass
	MTL4AccelerationStructureTriangleGeometryDescriptorClassOnce sync.Once
)

func getMTL4AccelerationStructureTriangleGeometryDescriptorClass() _MTL4AccelerationStructureTriangleGeometryDescriptorClass {
	MTL4AccelerationStructureTriangleGeometryDescriptorClassOnce.Do(func() {
		MTL4AccelerationStructureTriangleGeometryDescriptorClass = _MTL4AccelerationStructureTriangleGeometryDescriptorClass{objc.GetClass("MTL4AccelerationStructureTriangleGeometryDescriptor")}
	})
	return MTL4AccelerationStructureTriangleGeometryDescriptorClass
}

type _MTL4AccelerationStructureTriangleGeometryDescriptorClass struct {
	class objc.Class
}





// An interface definition for the [MTL4AccelerationStructureTriangleGeometryDescriptor] class.
type IMTL4AccelerationStructureTriangleGeometryDescriptor interface {
	IMTL4AccelerationStructureGeometryDescriptor
	

	// properties:
	IndexBuffer() MTL4BufferRange
	SetIndexBuffer(value MTL4BufferRange)
	IndexType() IndexType
	SetIndexType(value IndexType)
	TransformationMatrixBuffer() MTL4BufferRange
	SetTransformationMatrixBuffer(value MTL4BufferRange)
	TransformationMatrixLayout() MatrixLayout
	SetTransformationMatrixLayout(value MatrixLayout)
	TriangleCount() uint
	SetTriangleCount(value uint)
	VertexBuffer() MTL4BufferRange
	SetVertexBuffer(value MTL4BufferRange)
	VertexFormat() AttributeFormat
	SetVertexFormat(value AttributeFormat)
	VertexStride() uint
	SetVertexStride(value uint)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (mc _MTL4AccelerationStructureTriangleGeometryDescriptorClass) Alloc() MTL4AccelerationStructureTriangleGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureTriangleGeometryDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTL4AccelerationStructureTriangleGeometryDescriptorClass) New() MTL4AccelerationStructureTriangleGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureTriangleGeometryDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTL4AccelerationStructureTriangleGeometryDescriptor) Init() MTL4AccelerationStructureTriangleGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureTriangleGeometryDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTL4AccelerationStructureTriangleGeometryDescriptor) Autorelease() MTL4AccelerationStructureTriangleGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureTriangleGeometryDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4AccelerationStructureTriangleGeometryDescriptor creates a new MTL4AccelerationStructureTriangleGeometryDescriptor instance.
func NewMTL4AccelerationStructureTriangleGeometryDescriptor() MTL4AccelerationStructureTriangleGeometryDescriptor {
	return getMTL4AccelerationStructureTriangleGeometryDescriptorClass().New()
}





// Describes triangle geometry suitable for ray tracing.
//
// Use a to mark residency of all buffers this descriptor references when you build this acceleration structure.


// Describes triangle geometry suitable for ray tracing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureTriangleGeometryDescriptor
type MTL4AccelerationStructureTriangleGeometryDescriptor struct {
	MTL4AccelerationStructureGeometryDescriptor
}

// MTL4AccelerationStructureTriangleGeometryDescriptorFrom constructs a [MTL4AccelerationStructureTriangleGeometryDescriptor] from an unsafe.Pointer.
//
// Describes triangle geometry suitable for ray tracing.
func MTL4AccelerationStructureTriangleGeometryDescriptorFrom(ptr unsafe.Pointer) MTL4AccelerationStructureTriangleGeometryDescriptor {
	return MTL4AccelerationStructureTriangleGeometryDescriptor{
		MTL4AccelerationStructureGeometryDescriptor: MTL4AccelerationStructureGeometryDescriptorFrom(ptr),
	}
}

























// Sets an optional index buffer containing references to vertices in the .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureTriangleGeometryDescriptor/indexBuffer
func (m_ MTL4AccelerationStructureTriangleGeometryDescriptor) IndexBuffer() MTL4BufferRange {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("indexBuffer"))
	return rv
}


// Sets an optional index buffer containing references to vertices in the .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureTriangleGeometryDescriptor/indexBuffer
func (m_ MTL4AccelerationStructureTriangleGeometryDescriptor) SetIndexBuffer(value MTL4BufferRange) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIndexBuffer:"), value)
}


// Configures the size of the indices the contains, which is typically either 16 or 32-bits for each index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureTriangleGeometryDescriptor/indexType
func (m_ MTL4AccelerationStructureTriangleGeometryDescriptor) IndexType() IndexType {
	rv := objc.Send[IndexType](m_.ID, objc.Sel("indexType"))
	return rv
}


// Configures the size of the indices the contains, which is typically either 16 or 32-bits for each index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureTriangleGeometryDescriptor/indexType
func (m_ MTL4AccelerationStructureTriangleGeometryDescriptor) SetIndexType(value IndexType) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIndexType:"), value)
}


// Assigns an optional reference to a buffer containing a transformation matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureTriangleGeometryDescriptor/transformationMatrixBuffer
func (m_ MTL4AccelerationStructureTriangleGeometryDescriptor) TransformationMatrixBuffer() MTL4BufferRange {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("transformationMatrixBuffer"))
	return rv
}


// Assigns an optional reference to a buffer containing a transformation matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureTriangleGeometryDescriptor/transformationMatrixBuffer
func (m_ MTL4AccelerationStructureTriangleGeometryDescriptor) SetTransformationMatrixBuffer(value MTL4BufferRange) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTransformationMatrixBuffer:"), value)
}


// Configures the layout for the transformation matrix in the transformation matrix buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureTriangleGeometryDescriptor/transformationMatrixLayout
func (m_ MTL4AccelerationStructureTriangleGeometryDescriptor) TransformationMatrixLayout() MatrixLayout {
	rv := objc.Send[MatrixLayout](m_.ID, objc.Sel("transformationMatrixLayout"))
	return rv
}


// Configures the layout for the transformation matrix in the transformation matrix buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureTriangleGeometryDescriptor/transformationMatrixLayout
func (m_ MTL4AccelerationStructureTriangleGeometryDescriptor) SetTransformationMatrixLayout(value MatrixLayout) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTransformationMatrixLayout:"), value)
}


// Declares the number of triangles in this geometry descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureTriangleGeometryDescriptor/triangleCount
func (m_ MTL4AccelerationStructureTriangleGeometryDescriptor) TriangleCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("triangleCount"))
	return rv
}


// Declares the number of triangles in this geometry descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureTriangleGeometryDescriptor/triangleCount
func (m_ MTL4AccelerationStructureTriangleGeometryDescriptor) SetTriangleCount(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTriangleCount:"), value)
}


// Associates a vertex buffer containing triangle vertices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureTriangleGeometryDescriptor/vertexBuffer
func (m_ MTL4AccelerationStructureTriangleGeometryDescriptor) VertexBuffer() MTL4BufferRange {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("vertexBuffer"))
	return rv
}


// Associates a vertex buffer containing triangle vertices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureTriangleGeometryDescriptor/vertexBuffer
func (m_ MTL4AccelerationStructureTriangleGeometryDescriptor) SetVertexBuffer(value MTL4BufferRange) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVertexBuffer:"), value)
}


// Describes the format of the vertices in the vertex buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureTriangleGeometryDescriptor/vertexFormat
func (m_ MTL4AccelerationStructureTriangleGeometryDescriptor) VertexFormat() AttributeFormat {
	rv := objc.Send[AttributeFormat](m_.ID, objc.Sel("vertexFormat"))
	return rv
}


// Describes the format of the vertices in the vertex buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureTriangleGeometryDescriptor/vertexFormat
func (m_ MTL4AccelerationStructureTriangleGeometryDescriptor) SetVertexFormat(value AttributeFormat) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVertexFormat:"), value)
}


// Sets the stride, in bytes, between vertices in the vertex buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureTriangleGeometryDescriptor/vertexStride
func (m_ MTL4AccelerationStructureTriangleGeometryDescriptor) VertexStride() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("vertexStride"))
	return rv
}


// Sets the stride, in bytes, between vertices in the vertex buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureTriangleGeometryDescriptor/vertexStride
func (m_ MTL4AccelerationStructureTriangleGeometryDescriptor) SetVertexStride(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVertexStride:"), value)
}








