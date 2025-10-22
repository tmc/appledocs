// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AccelerationStructureMotionTriangleGeometryDescriptor] class.
var (
	AccelerationStructureMotionTriangleGeometryDescriptorClass     _AccelerationStructureMotionTriangleGeometryDescriptorClass
	AccelerationStructureMotionTriangleGeometryDescriptorClassOnce sync.Once
)

func getAccelerationStructureMotionTriangleGeometryDescriptorClass() _AccelerationStructureMotionTriangleGeometryDescriptorClass {
	AccelerationStructureMotionTriangleGeometryDescriptorClassOnce.Do(func() {
		AccelerationStructureMotionTriangleGeometryDescriptorClass = _AccelerationStructureMotionTriangleGeometryDescriptorClass{objc.GetClass("MTLAccelerationStructureMotionTriangleGeometryDescriptor")}
	})
	return AccelerationStructureMotionTriangleGeometryDescriptorClass
}

type _AccelerationStructureMotionTriangleGeometryDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [AccelerationStructureMotionTriangleGeometryDescriptor] class.
type IAccelerationStructureMotionTriangleGeometryDescriptor interface {
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
	VertexBuffers() MTLMotionKeyframeData
	SetVertexBuffers(value IMTLMotionKeyframeData)
	VertexFormat() unsafe.Pointer
	SetVertexFormat(value unsafe.Pointer)
	VertexStride() int
	SetVertexStride(value int)
}

// A description of a list of triangle primitives, as motion keyframe data, to turn into an acceleration structure.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionTriangleGeometryDescriptor
type AccelerationStructureMotionTriangleGeometryDescriptor struct {
	AccelerationStructureGeometryDescriptor
}

// AccelerationStructureMotionTriangleGeometryDescriptorFrom constructs a [AccelerationStructureMotionTriangleGeometryDescriptor] from an unsafe.Pointer.
//
// A description of a list of triangle primitives, as motion keyframe data, to turn into an acceleration structure.
func AccelerationStructureMotionTriangleGeometryDescriptorFrom(ptr unsafe.Pointer) AccelerationStructureMotionTriangleGeometryDescriptor {
	return AccelerationStructureMotionTriangleGeometryDescriptor{
		AccelerationStructureGeometryDescriptor: AccelerationStructureGeometryDescriptorFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AccelerationStructureMotionTriangleGeometryDescriptorClass) Alloc() AccelerationStructureMotionTriangleGeometryDescriptor {
	rv := objc.Send[AccelerationStructureMotionTriangleGeometryDescriptor](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AccelerationStructureMotionTriangleGeometryDescriptorClass) New() AccelerationStructureMotionTriangleGeometryDescriptor {
	rv := objc.Send[AccelerationStructureMotionTriangleGeometryDescriptor](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AccelerationStructureMotionTriangleGeometryDescriptor) Init() AccelerationStructureMotionTriangleGeometryDescriptor {
	rv := objc.Send[AccelerationStructureMotionTriangleGeometryDescriptor](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AccelerationStructureMotionTriangleGeometryDescriptor) Autorelease() AccelerationStructureMotionTriangleGeometryDescriptor {
	rv := objc.Send[AccelerationStructureMotionTriangleGeometryDescriptor](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAccelerationStructureMotionTriangleGeometryDescriptor creates a new AccelerationStructureMotionTriangleGeometryDescriptor instance.
func NewAccelerationStructureMotionTriangleGeometryDescriptor() AccelerationStructureMotionTriangleGeometryDescriptor {
	return getAccelerationStructureMotionTriangleGeometryDescriptorClass().New()
}


// A buffer that contains indices for the vertices that compose the triangle list.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotiontrianglegeometrydescriptor/indexbuffer
func (a_ AccelerationStructureMotionTriangleGeometryDescriptor) IndexBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("indexBuffer"))
	return rv
}


// SetIndexBuffer sets the value of the indexBuffer property.
// A buffer that contains indices for the vertices that compose the triangle list.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotiontrianglegeometrydescriptor/indexbuffer
func (a_ AccelerationStructureMotionTriangleGeometryDescriptor) SetIndexBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIndexBuffer:"), value)
}

// The offset, in bytes, to the first index in the buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotiontrianglegeometrydescriptor/indexbufferoffset
func (a_ AccelerationStructureMotionTriangleGeometryDescriptor) IndexBufferOffset() int {
	rv := objc.Send[int](a_.ID, objc.Sel("indexBufferOffset"))
	return rv
}


// SetIndexBufferOffset sets the value of the indexBufferOffset property.
// The offset, in bytes, to the first index in the buffer.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotiontrianglegeometrydescriptor/indexbufferoffset
func (a_ AccelerationStructureMotionTriangleGeometryDescriptor) SetIndexBufferOffset(value int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIndexBufferOffset:"), value)
}

// The data type of indices in the index buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotiontrianglegeometrydescriptor/indextype
func (a_ AccelerationStructureMotionTriangleGeometryDescriptor) IndexType() IndexType {
	rv := objc.Send[IndexType](a_.ID, objc.Sel("indexType"))
	return rv
}


// SetIndexType sets the value of the indexType property.
// The data type of indices in the index buffer.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotiontrianglegeometrydescriptor/indextype
func (a_ AccelerationStructureMotionTriangleGeometryDescriptor) SetIndexType(value IndexType) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIndexType:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotiontrianglegeometrydescriptor/transformationmatrixbuffer
func (a_ AccelerationStructureMotionTriangleGeometryDescriptor) TransformationMatrixBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("transformationMatrixBuffer"))
	return rv
}


// SetTransformationMatrixBuffer sets the value of the transformationMatrixBuffer property.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotiontrianglegeometrydescriptor/transformationmatrixbuffer
func (a_ AccelerationStructureMotionTriangleGeometryDescriptor) SetTransformationMatrixBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTransformationMatrixBuffer:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotiontrianglegeometrydescriptor/transformationmatrixbufferoffset
func (a_ AccelerationStructureMotionTriangleGeometryDescriptor) TransformationMatrixBufferOffset() int {
	rv := objc.Send[int](a_.ID, objc.Sel("transformationMatrixBufferOffset"))
	return rv
}


// SetTransformationMatrixBufferOffset sets the value of the transformationMatrixBufferOffset property.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotiontrianglegeometrydescriptor/transformationmatrixbufferoffset
func (a_ AccelerationStructureMotionTriangleGeometryDescriptor) SetTransformationMatrixBufferOffset(value int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTransformationMatrixBufferOffset:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotiontrianglegeometrydescriptor/transformationmatrixlayout
func (a_ AccelerationStructureMotionTriangleGeometryDescriptor) TransformationMatrixLayout() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("transformationMatrixLayout"))
	return rv
}


// SetTransformationMatrixLayout sets the value of the transformationMatrixLayout property.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotiontrianglegeometrydescriptor/transformationmatrixlayout
func (a_ AccelerationStructureMotionTriangleGeometryDescriptor) SetTransformationMatrixLayout(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTransformationMatrixLayout:"), value)
}

// The number of triangles in the buffers.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotiontrianglegeometrydescriptor/trianglecount
func (a_ AccelerationStructureMotionTriangleGeometryDescriptor) TriangleCount() int {
	rv := objc.Send[int](a_.ID, objc.Sel("triangleCount"))
	return rv
}


// SetTriangleCount sets the value of the triangleCount property.
// The number of triangles in the buffers.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotiontrianglegeometrydescriptor/trianglecount
func (a_ AccelerationStructureMotionTriangleGeometryDescriptor) SetTriangleCount(value int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTriangleCount:"), value)
}

// An array of motion keyframes, each containing triangle data.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotiontrianglegeometrydescriptor/vertexbuffers
func (a_ AccelerationStructureMotionTriangleGeometryDescriptor) VertexBuffers() MTLMotionKeyframeData {
	rv := objc.Send[MTLMotionKeyframeData](a_.ID, objc.Sel("vertexBuffers"))
	return rv
}


// SetVertexBuffers sets the value of the vertexBuffers property.
// An array of motion keyframes, each containing triangle data.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotiontrianglegeometrydescriptor/vertexbuffers
func (a_ AccelerationStructureMotionTriangleGeometryDescriptor) SetVertexBuffers(value IMTLMotionKeyframeData) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setVertexBuffers:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotiontrianglegeometrydescriptor/vertexformat
func (a_ AccelerationStructureMotionTriangleGeometryDescriptor) VertexFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("vertexFormat"))
	return rv
}


// SetVertexFormat sets the value of the vertexFormat property.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotiontrianglegeometrydescriptor/vertexformat
func (a_ AccelerationStructureMotionTriangleGeometryDescriptor) SetVertexFormat(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setVertexFormat:"), value)
}

// The stride, in bytes, between vertices in each vertex buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotiontrianglegeometrydescriptor/vertexstride
func (a_ AccelerationStructureMotionTriangleGeometryDescriptor) VertexStride() int {
	rv := objc.Send[int](a_.ID, objc.Sel("vertexStride"))
	return rv
}


// SetVertexStride sets the value of the vertexStride property.
// The stride, in bytes, between vertices in each vertex buffer.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotiontrianglegeometrydescriptor/vertexstride
func (a_ AccelerationStructureMotionTriangleGeometryDescriptor) SetVertexStride(value int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setVertexStride:"), value)
}



