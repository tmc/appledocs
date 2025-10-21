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
}

// Describes triangle geometry suitable for ray tracing.
//
// Use a to mark residency of all buffers this descriptor references when you build this acceleration structure.
//
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

// Alloc allocates a new instance without initialization.
func (mc _MTL4AccelerationStructureTriangleGeometryDescriptorClass) Alloc() MTL4AccelerationStructureTriangleGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureTriangleGeometryDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Configures the layout for the transformation matrix in the transformation matrix buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuretrianglegeometrydescriptor/transformationmatrixlayout
func (m_ MTL4AccelerationStructureTriangleGeometryDescriptor) TransformationMatrixLayout() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("transformationMatrixLayout"))
	return rv
}


// SetTransformationMatrixLayout sets the value of the transformationMatrixLayout property.
// Configures the layout for the transformation matrix in the transformation matrix buffer.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuretrianglegeometrydescriptor/transformationmatrixlayout
func (m_ MTL4AccelerationStructureTriangleGeometryDescriptor) SetTransformationMatrixLayout(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTransformationMatrixLayout:"), value)
}

// Assigns an optional reference to a buffer containing a
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuretrianglegeometrydescriptor/transformationmatrixbuffer
func (m_ MTL4AccelerationStructureTriangleGeometryDescriptor) TransformationMatrixBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("transformationMatrixBuffer"))
	return rv
}


// SetTransformationMatrixBuffer sets the value of the transformationMatrixBuffer property.
// Assigns an optional reference to a buffer containing a

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuretrianglegeometrydescriptor/transformationmatrixbuffer
func (m_ MTL4AccelerationStructureTriangleGeometryDescriptor) SetTransformationMatrixBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTransformationMatrixBuffer:"), value)
}

// Sets an optional index buffer containing references to vertices in the
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuretrianglegeometrydescriptor/indexbuffer
func (m_ MTL4AccelerationStructureTriangleGeometryDescriptor) IndexBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("indexBuffer"))
	return rv
}


// SetIndexBuffer sets the value of the indexBuffer property.
// Sets an optional index buffer containing references to vertices in the

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuretrianglegeometrydescriptor/indexbuffer
func (m_ MTL4AccelerationStructureTriangleGeometryDescriptor) SetIndexBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIndexBuffer:"), value)
}

// Configures the size of the indices the
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuretrianglegeometrydescriptor/indextype
func (m_ MTL4AccelerationStructureTriangleGeometryDescriptor) IndexType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("indexType"))
	return rv
}


// SetIndexType sets the value of the indexType property.
// Configures the size of the indices the

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuretrianglegeometrydescriptor/indextype
func (m_ MTL4AccelerationStructureTriangleGeometryDescriptor) SetIndexType(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIndexType:"), value)
}

// Declares the number of triangles in this geometry descriptor.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuretrianglegeometrydescriptor/trianglecount
func (m_ MTL4AccelerationStructureTriangleGeometryDescriptor) TriangleCount() int {
	rv := objc.Send[int](m_.ID, objc.Sel("triangleCount"))
	return rv
}


// SetTriangleCount sets the value of the triangleCount property.
// Declares the number of triangles in this geometry descriptor.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuretrianglegeometrydescriptor/trianglecount
func (m_ MTL4AccelerationStructureTriangleGeometryDescriptor) SetTriangleCount(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTriangleCount:"), value)
}

// Describes the format of the vertices in the vertex buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuretrianglegeometrydescriptor/vertexformat
func (m_ MTL4AccelerationStructureTriangleGeometryDescriptor) VertexFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("vertexFormat"))
	return rv
}


// SetVertexFormat sets the value of the vertexFormat property.
// Describes the format of the vertices in the vertex buffer.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuretrianglegeometrydescriptor/vertexformat
func (m_ MTL4AccelerationStructureTriangleGeometryDescriptor) SetVertexFormat(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVertexFormat:"), value)
}

// Associates a vertex buffer containing triangle vertices.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuretrianglegeometrydescriptor/vertexbuffer
func (m_ MTL4AccelerationStructureTriangleGeometryDescriptor) VertexBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("vertexBuffer"))
	return rv
}


// SetVertexBuffer sets the value of the vertexBuffer property.
// Associates a vertex buffer containing triangle vertices.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuretrianglegeometrydescriptor/vertexbuffer
func (m_ MTL4AccelerationStructureTriangleGeometryDescriptor) SetVertexBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVertexBuffer:"), value)
}

// Sets the stride, in bytes, between vertices in the vertex buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuretrianglegeometrydescriptor/vertexstride
func (m_ MTL4AccelerationStructureTriangleGeometryDescriptor) VertexStride() int {
	rv := objc.Send[int](m_.ID, objc.Sel("vertexStride"))
	return rv
}


// SetVertexStride sets the value of the vertexStride property.
// Sets the stride, in bytes, between vertices in the vertex buffer.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuretrianglegeometrydescriptor/vertexstride
func (m_ MTL4AccelerationStructureTriangleGeometryDescriptor) SetVertexStride(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVertexStride:"), value)
}



