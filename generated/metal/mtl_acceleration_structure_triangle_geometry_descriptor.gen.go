// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLAccelerationStructureTriangleGeometryDescriptor */


/* debug [class_header]: Header for MTLAccelerationStructureTriangleGeometryDescriptor */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AccelerationStructureTriangleGeometryDescriptor */
// An interface definition for the [AccelerationStructureTriangleGeometryDescriptor] class.
type IAccelerationStructureTriangleGeometryDescriptor interface {
	IAccelerationStructureGeometryDescriptor
	
/* debug [class_interface_properties]: Properties for AccelerationStructureTriangleGeometryDescriptor */
	// properties:
	IndexBuffer() unsafe.Pointer
	SetIndexBuffer(value unsafe.Pointer)
	IndexBufferOffset() uint
	SetIndexBufferOffset(value uint)
	IndexType() IndexType
	SetIndexType(value IndexType)
	TransformationMatrixBuffer() unsafe.Pointer
	SetTransformationMatrixBuffer(value unsafe.Pointer)
	TransformationMatrixBufferOffset() uint
	SetTransformationMatrixBufferOffset(value uint)
	TransformationMatrixLayout() MatrixLayout
	SetTransformationMatrixLayout(value MatrixLayout)
	TriangleCount() uint
	SetTriangleCount(value uint)
	VertexBuffer() unsafe.Pointer
	SetVertexBuffer(value unsafe.Pointer)
	VertexBufferOffset() uint
	SetVertexBufferOffset(value uint)
	VertexFormat() AttributeFormat
	SetVertexFormat(value AttributeFormat)
	VertexStride() uint
	SetVertexStride(value uint)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AccelerationStructureTriangleGeometryDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AccelerationStructureTriangleGeometryDescriptor */
// Alloc allocates a new instance without initialization.
func (ac _AccelerationStructureTriangleGeometryDescriptorClass) Alloc() AccelerationStructureTriangleGeometryDescriptor {
	rv := objc.Send[AccelerationStructureTriangleGeometryDescriptor](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AccelerationStructureTriangleGeometryDescriptor */
// A description of a list of triangle primitives to turn into an acceleration structure.


// A description of a list of triangle primitives to turn into an acceleration structure.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AccelerationStructureTriangleGeometryDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AccelerationStructureTriangleGeometryDescriptor */

// Creates a new triangle descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureTriangleGeometryDescriptor/descriptor
func (ac _AccelerationStructureTriangleGeometryDescriptorClass) Descriptor() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ac.class), objc.Sel("descriptor"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=Descriptor) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AccelerationStructureTriangleGeometryDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AccelerationStructureTriangleGeometryDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AccelerationStructureTriangleGeometryDescriptor */

// A buffer that contains indices for the vertices that compose the triangle list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureTriangleGeometryDescriptor/indexBuffer
func (a_ AccelerationStructureTriangleGeometryDescriptor) IndexBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("indexBuffer"))
	return rv
}/* debug [instance_properties/getter]: indexBuffer */


// A buffer that contains indices for the vertices that compose the triangle list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureTriangleGeometryDescriptor/indexBuffer
func (a_ AccelerationStructureTriangleGeometryDescriptor) SetIndexBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIndexBuffer:"), value)
}/* debug [instance_properties/setter]: indexBuffer */


// The offset, in bytes, to the first index in the buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureTriangleGeometryDescriptor/indexBufferOffset
func (a_ AccelerationStructureTriangleGeometryDescriptor) IndexBufferOffset() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("indexBufferOffset"))
	return rv
}/* debug [instance_properties/getter]: indexBufferOffset */


// The offset, in bytes, to the first index in the buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureTriangleGeometryDescriptor/indexBufferOffset
func (a_ AccelerationStructureTriangleGeometryDescriptor) SetIndexBufferOffset(value uint) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIndexBufferOffset:"), value)
}/* debug [instance_properties/setter]: indexBufferOffset */


// The data type of indices in the index buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureTriangleGeometryDescriptor/indexType
func (a_ AccelerationStructureTriangleGeometryDescriptor) IndexType() IndexType {
	rv := objc.Send[IndexType](a_.ID, objc.Sel("indexType"))
	return rv
}/* debug [instance_properties/getter]: indexType */


// The data type of indices in the index buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureTriangleGeometryDescriptor/indexType
func (a_ AccelerationStructureTriangleGeometryDescriptor) SetIndexType(value IndexType) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIndexType:"), value)
}/* debug [instance_properties/setter]: indexType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureTriangleGeometryDescriptor/transformationMatrixBuffer
func (a_ AccelerationStructureTriangleGeometryDescriptor) TransformationMatrixBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("transformationMatrixBuffer"))
	return rv
}/* debug [instance_properties/getter]: transformationMatrixBuffer */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureTriangleGeometryDescriptor/transformationMatrixBuffer
func (a_ AccelerationStructureTriangleGeometryDescriptor) SetTransformationMatrixBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTransformationMatrixBuffer:"), value)
}/* debug [instance_properties/setter]: transformationMatrixBuffer */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureTriangleGeometryDescriptor/transformationMatrixBufferOffset
func (a_ AccelerationStructureTriangleGeometryDescriptor) TransformationMatrixBufferOffset() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("transformationMatrixBufferOffset"))
	return rv
}/* debug [instance_properties/getter]: transformationMatrixBufferOffset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureTriangleGeometryDescriptor/transformationMatrixBufferOffset
func (a_ AccelerationStructureTriangleGeometryDescriptor) SetTransformationMatrixBufferOffset(value uint) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTransformationMatrixBufferOffset:"), value)
}/* debug [instance_properties/setter]: transformationMatrixBufferOffset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureTriangleGeometryDescriptor/transformationMatrixLayout
func (a_ AccelerationStructureTriangleGeometryDescriptor) TransformationMatrixLayout() MatrixLayout {
	rv := objc.Send[MatrixLayout](a_.ID, objc.Sel("transformationMatrixLayout"))
	return rv
}/* debug [instance_properties/getter]: transformationMatrixLayout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureTriangleGeometryDescriptor/transformationMatrixLayout
func (a_ AccelerationStructureTriangleGeometryDescriptor) SetTransformationMatrixLayout(value MatrixLayout) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTransformationMatrixLayout:"), value)
}/* debug [instance_properties/setter]: transformationMatrixLayout */


// The number of triangles in the buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureTriangleGeometryDescriptor/triangleCount
func (a_ AccelerationStructureTriangleGeometryDescriptor) TriangleCount() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("triangleCount"))
	return rv
}/* debug [instance_properties/getter]: triangleCount */


// The number of triangles in the buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureTriangleGeometryDescriptor/triangleCount
func (a_ AccelerationStructureTriangleGeometryDescriptor) SetTriangleCount(value uint) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTriangleCount:"), value)
}/* debug [instance_properties/setter]: triangleCount */


// A buffer that contains vertex data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureTriangleGeometryDescriptor/vertexBuffer
func (a_ AccelerationStructureTriangleGeometryDescriptor) VertexBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("vertexBuffer"))
	return rv
}/* debug [instance_properties/getter]: vertexBuffer */


// A buffer that contains vertex data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureTriangleGeometryDescriptor/vertexBuffer
func (a_ AccelerationStructureTriangleGeometryDescriptor) SetVertexBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setVertexBuffer:"), value)
}/* debug [instance_properties/setter]: vertexBuffer */


// The offset, in bytes, for the first vertex in the vertex buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureTriangleGeometryDescriptor/vertexBufferOffset
func (a_ AccelerationStructureTriangleGeometryDescriptor) VertexBufferOffset() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("vertexBufferOffset"))
	return rv
}/* debug [instance_properties/getter]: vertexBufferOffset */


// The offset, in bytes, for the first vertex in the vertex buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureTriangleGeometryDescriptor/vertexBufferOffset
func (a_ AccelerationStructureTriangleGeometryDescriptor) SetVertexBufferOffset(value uint) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setVertexBufferOffset:"), value)
}/* debug [instance_properties/setter]: vertexBufferOffset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureTriangleGeometryDescriptor/vertexFormat
func (a_ AccelerationStructureTriangleGeometryDescriptor) VertexFormat() AttributeFormat {
	rv := objc.Send[AttributeFormat](a_.ID, objc.Sel("vertexFormat"))
	return rv
}/* debug [instance_properties/getter]: vertexFormat */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureTriangleGeometryDescriptor/vertexFormat
func (a_ AccelerationStructureTriangleGeometryDescriptor) SetVertexFormat(value AttributeFormat) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setVertexFormat:"), value)
}/* debug [instance_properties/setter]: vertexFormat */


// The stride, in bytes, between vertices in the vertex buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureTriangleGeometryDescriptor/vertexStride
func (a_ AccelerationStructureTriangleGeometryDescriptor) VertexStride() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("vertexStride"))
	return rv
}/* debug [instance_properties/getter]: vertexStride */


// The stride, in bytes, between vertices in the vertex buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureTriangleGeometryDescriptor/vertexStride
func (a_ AccelerationStructureTriangleGeometryDescriptor) SetVertexStride(value uint) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setVertexStride:"), value)
}/* debug [instance_properties/setter]: vertexStride */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLAccelerationStructureTriangleGeometryDescriptor */



