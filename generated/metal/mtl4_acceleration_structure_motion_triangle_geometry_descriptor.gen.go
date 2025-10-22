// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTL4AccelerationStructureMotionTriangleGeometryDescriptor] class.
var (
	MTL4AccelerationStructureMotionTriangleGeometryDescriptorClass     _MTL4AccelerationStructureMotionTriangleGeometryDescriptorClass
	MTL4AccelerationStructureMotionTriangleGeometryDescriptorClassOnce sync.Once
)

func getMTL4AccelerationStructureMotionTriangleGeometryDescriptorClass() _MTL4AccelerationStructureMotionTriangleGeometryDescriptorClass {
	MTL4AccelerationStructureMotionTriangleGeometryDescriptorClassOnce.Do(func() {
		MTL4AccelerationStructureMotionTriangleGeometryDescriptorClass = _MTL4AccelerationStructureMotionTriangleGeometryDescriptorClass{objc.GetClass("MTL4AccelerationStructureMotionTriangleGeometryDescriptor")}
	})
	return MTL4AccelerationStructureMotionTriangleGeometryDescriptorClass
}

type _MTL4AccelerationStructureMotionTriangleGeometryDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [MTL4AccelerationStructureMotionTriangleGeometryDescriptor] class.
type IMTL4AccelerationStructureMotionTriangleGeometryDescriptor interface {
	IMTL4AccelerationStructureGeometryDescriptor
	IndexBuffer() unsafe.Pointer
	SetIndexBuffer(value unsafe.Pointer)
	IndexType() IndexType
	SetIndexType(value IndexType)
	TransformationMatrixBuffer() unsafe.Pointer
	SetTransformationMatrixBuffer(value unsafe.Pointer)
	TransformationMatrixLayout() unsafe.Pointer
	SetTransformationMatrixLayout(value unsafe.Pointer)
	TriangleCount() int
	SetTriangleCount(value int)
	VertexBuffers() unsafe.Pointer
	SetVertexBuffers(value unsafe.Pointer)
	VertexFormat() unsafe.Pointer
	SetVertexFormat(value unsafe.Pointer)
	VertexStride() int
	SetVertexStride(value int)
}

// Describes motion triangle geometry, suitable for motion ray tracing.
//
// Use a to mark residency of all buffers this descriptor references when you build this acceleration structure.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionTriangleGeometryDescriptor
type MTL4AccelerationStructureMotionTriangleGeometryDescriptor struct {
	MTL4AccelerationStructureGeometryDescriptor
}

// MTL4AccelerationStructureMotionTriangleGeometryDescriptorFrom constructs a [MTL4AccelerationStructureMotionTriangleGeometryDescriptor] from an unsafe.Pointer.
//
// Describes motion triangle geometry, suitable for motion ray tracing.
func MTL4AccelerationStructureMotionTriangleGeometryDescriptorFrom(ptr unsafe.Pointer) MTL4AccelerationStructureMotionTriangleGeometryDescriptor {
	return MTL4AccelerationStructureMotionTriangleGeometryDescriptor{
		MTL4AccelerationStructureGeometryDescriptor: MTL4AccelerationStructureGeometryDescriptorFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTL4AccelerationStructureMotionTriangleGeometryDescriptorClass) Alloc() MTL4AccelerationStructureMotionTriangleGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureMotionTriangleGeometryDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTL4AccelerationStructureMotionTriangleGeometryDescriptorClass) New() MTL4AccelerationStructureMotionTriangleGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureMotionTriangleGeometryDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTL4AccelerationStructureMotionTriangleGeometryDescriptor) Init() MTL4AccelerationStructureMotionTriangleGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureMotionTriangleGeometryDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTL4AccelerationStructureMotionTriangleGeometryDescriptor) Autorelease() MTL4AccelerationStructureMotionTriangleGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureMotionTriangleGeometryDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4AccelerationStructureMotionTriangleGeometryDescriptor creates a new MTL4AccelerationStructureMotionTriangleGeometryDescriptor instance.
func NewMTL4AccelerationStructureMotionTriangleGeometryDescriptor() MTL4AccelerationStructureMotionTriangleGeometryDescriptor {
	return getMTL4AccelerationStructureMotionTriangleGeometryDescriptorClass().New()
}


// Assigns an optional index buffer containing references to vertices in the vertex buffers you reference through the
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuremotiontrianglegeometrydescriptor/indexbuffer
func (m_ MTL4AccelerationStructureMotionTriangleGeometryDescriptor) IndexBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("indexBuffer"))
	return rv
}


// SetIndexBuffer sets the value of the indexBuffer property.
// Assigns an optional index buffer containing references to vertices in the vertex buffers you reference through the

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuremotiontrianglegeometrydescriptor/indexbuffer
func (m_ MTL4AccelerationStructureMotionTriangleGeometryDescriptor) SetIndexBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIndexBuffer:"), value)
}

// Specifies the size of the indices the
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuremotiontrianglegeometrydescriptor/indextype
func (m_ MTL4AccelerationStructureMotionTriangleGeometryDescriptor) IndexType() IndexType {
	rv := objc.Send[IndexType](m_.ID, objc.Sel("indexType"))
	return rv
}


// SetIndexType sets the value of the indexType property.
// Specifies the size of the indices the

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuremotiontrianglegeometrydescriptor/indextype
func (m_ MTL4AccelerationStructureMotionTriangleGeometryDescriptor) SetIndexType(value IndexType) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIndexType:"), value)
}

// Assings an optional reference to a buffer containing a
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuremotiontrianglegeometrydescriptor/transformationmatrixbuffer
func (m_ MTL4AccelerationStructureMotionTriangleGeometryDescriptor) TransformationMatrixBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("transformationMatrixBuffer"))
	return rv
}


// SetTransformationMatrixBuffer sets the value of the transformationMatrixBuffer property.
// Assings an optional reference to a buffer containing a

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuremotiontrianglegeometrydescriptor/transformationmatrixbuffer
func (m_ MTL4AccelerationStructureMotionTriangleGeometryDescriptor) SetTransformationMatrixBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTransformationMatrixBuffer:"), value)
}

// Configures the layout for the transformation matrix in the transformation matrix buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuremotiontrianglegeometrydescriptor/transformationmatrixlayout
func (m_ MTL4AccelerationStructureMotionTriangleGeometryDescriptor) TransformationMatrixLayout() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("transformationMatrixLayout"))
	return rv
}


// SetTransformationMatrixLayout sets the value of the transformationMatrixLayout property.
// Configures the layout for the transformation matrix in the transformation matrix buffer.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuremotiontrianglegeometrydescriptor/transformationmatrixlayout
func (m_ MTL4AccelerationStructureMotionTriangleGeometryDescriptor) SetTransformationMatrixLayout(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTransformationMatrixLayout:"), value)
}

// Declares the number of triangles in the vertex buffers that the buffer in the vertex buffers property references.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuremotiontrianglegeometrydescriptor/trianglecount
func (m_ MTL4AccelerationStructureMotionTriangleGeometryDescriptor) TriangleCount() int {
	rv := objc.Send[int](m_.ID, objc.Sel("triangleCount"))
	return rv
}


// SetTriangleCount sets the value of the triangleCount property.
// Declares the number of triangles in the vertex buffers that the buffer in the vertex buffers property references.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuremotiontrianglegeometrydescriptor/trianglecount
func (m_ MTL4AccelerationStructureMotionTriangleGeometryDescriptor) SetTriangleCount(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTriangleCount:"), value)
}

// Assigns a buffer where each entry contains a reference to a vertex buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuremotiontrianglegeometrydescriptor/vertexbuffers
func (m_ MTL4AccelerationStructureMotionTriangleGeometryDescriptor) VertexBuffers() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("vertexBuffers"))
	return rv
}


// SetVertexBuffers sets the value of the vertexBuffers property.
// Assigns a buffer where each entry contains a reference to a vertex buffer.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuremotiontrianglegeometrydescriptor/vertexbuffers
func (m_ MTL4AccelerationStructureMotionTriangleGeometryDescriptor) SetVertexBuffers(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVertexBuffers:"), value)
}

// Defines the format of the vertices in the vertex buffers.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuremotiontrianglegeometrydescriptor/vertexformat
func (m_ MTL4AccelerationStructureMotionTriangleGeometryDescriptor) VertexFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("vertexFormat"))
	return rv
}


// SetVertexFormat sets the value of the vertexFormat property.
// Defines the format of the vertices in the vertex buffers.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuremotiontrianglegeometrydescriptor/vertexformat
func (m_ MTL4AccelerationStructureMotionTriangleGeometryDescriptor) SetVertexFormat(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVertexFormat:"), value)
}

// Sets the stride, in bytes, between vertices in all the vertex buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuremotiontrianglegeometrydescriptor/vertexstride
func (m_ MTL4AccelerationStructureMotionTriangleGeometryDescriptor) VertexStride() int {
	rv := objc.Send[int](m_.ID, objc.Sel("vertexStride"))
	return rv
}


// SetVertexStride sets the value of the vertexStride property.
// Sets the stride, in bytes, between vertices in all the vertex buffer.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuremotiontrianglegeometrydescriptor/vertexstride
func (m_ MTL4AccelerationStructureMotionTriangleGeometryDescriptor) SetVertexStride(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVertexStride:"), value)
}



