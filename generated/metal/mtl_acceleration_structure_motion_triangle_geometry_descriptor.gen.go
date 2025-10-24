// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLAccelerationStructureMotionTriangleGeometryDescriptor */


/* debug [class_header]: Header for MTLAccelerationStructureMotionTriangleGeometryDescriptor */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AccelerationStructureMotionTriangleGeometryDescriptor */
// An interface definition for the [AccelerationStructureMotionTriangleGeometryDescriptor] class.
type IAccelerationStructureMotionTriangleGeometryDescriptor interface {
	IAccelerationStructureGeometryDescriptor
	
/* debug [class_interface_properties]: Properties for AccelerationStructureMotionTriangleGeometryDescriptor */
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
	VertexBuffers() []MotionKeyframeData
	SetVertexBuffers(value []MotionKeyframeData)
	VertexFormat() AttributeFormat
	SetVertexFormat(value AttributeFormat)
	VertexStride() uint
	SetVertexStride(value uint)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AccelerationStructureMotionTriangleGeometryDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AccelerationStructureMotionTriangleGeometryDescriptor */
// Alloc allocates a new instance without initialization.
func (ac _AccelerationStructureMotionTriangleGeometryDescriptorClass) Alloc() AccelerationStructureMotionTriangleGeometryDescriptor {
	rv := objc.Send[AccelerationStructureMotionTriangleGeometryDescriptor](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AccelerationStructureMotionTriangleGeometryDescriptor */
// A description of a list of triangle primitives, as motion keyframe data, to turn into an acceleration structure.


// A description of a list of triangle primitives, as motion keyframe data, to turn into an acceleration structure.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AccelerationStructureMotionTriangleGeometryDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AccelerationStructureMotionTriangleGeometryDescriptor */

// Creates a new triangle descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionTriangleGeometryDescriptor/descriptor
func (ac _AccelerationStructureMotionTriangleGeometryDescriptorClass) Descriptor() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ac.class), objc.Sel("descriptor"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=Descriptor) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AccelerationStructureMotionTriangleGeometryDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AccelerationStructureMotionTriangleGeometryDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AccelerationStructureMotionTriangleGeometryDescriptor */

// A buffer that contains indices for the vertices that compose the triangle list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionTriangleGeometryDescriptor/indexBuffer
func (a_ AccelerationStructureMotionTriangleGeometryDescriptor) IndexBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("indexBuffer"))
	return rv
}/* debug [instance_properties/getter]: indexBuffer */


// A buffer that contains indices for the vertices that compose the triangle list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionTriangleGeometryDescriptor/indexBuffer
func (a_ AccelerationStructureMotionTriangleGeometryDescriptor) SetIndexBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIndexBuffer:"), value)
}/* debug [instance_properties/setter]: indexBuffer */


// The offset, in bytes, to the first index in the buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionTriangleGeometryDescriptor/indexBufferOffset
func (a_ AccelerationStructureMotionTriangleGeometryDescriptor) IndexBufferOffset() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("indexBufferOffset"))
	return rv
}/* debug [instance_properties/getter]: indexBufferOffset */


// The offset, in bytes, to the first index in the buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionTriangleGeometryDescriptor/indexBufferOffset
func (a_ AccelerationStructureMotionTriangleGeometryDescriptor) SetIndexBufferOffset(value uint) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIndexBufferOffset:"), value)
}/* debug [instance_properties/setter]: indexBufferOffset */


// The data type of indices in the index buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionTriangleGeometryDescriptor/indexType
func (a_ AccelerationStructureMotionTriangleGeometryDescriptor) IndexType() IndexType {
	rv := objc.Send[IndexType](a_.ID, objc.Sel("indexType"))
	return rv
}/* debug [instance_properties/getter]: indexType */


// The data type of indices in the index buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionTriangleGeometryDescriptor/indexType
func (a_ AccelerationStructureMotionTriangleGeometryDescriptor) SetIndexType(value IndexType) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIndexType:"), value)
}/* debug [instance_properties/setter]: indexType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionTriangleGeometryDescriptor/transformationMatrixBuffer
func (a_ AccelerationStructureMotionTriangleGeometryDescriptor) TransformationMatrixBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("transformationMatrixBuffer"))
	return rv
}/* debug [instance_properties/getter]: transformationMatrixBuffer */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionTriangleGeometryDescriptor/transformationMatrixBuffer
func (a_ AccelerationStructureMotionTriangleGeometryDescriptor) SetTransformationMatrixBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTransformationMatrixBuffer:"), value)
}/* debug [instance_properties/setter]: transformationMatrixBuffer */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionTriangleGeometryDescriptor/transformationMatrixBufferOffset
func (a_ AccelerationStructureMotionTriangleGeometryDescriptor) TransformationMatrixBufferOffset() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("transformationMatrixBufferOffset"))
	return rv
}/* debug [instance_properties/getter]: transformationMatrixBufferOffset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionTriangleGeometryDescriptor/transformationMatrixBufferOffset
func (a_ AccelerationStructureMotionTriangleGeometryDescriptor) SetTransformationMatrixBufferOffset(value uint) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTransformationMatrixBufferOffset:"), value)
}/* debug [instance_properties/setter]: transformationMatrixBufferOffset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionTriangleGeometryDescriptor/transformationMatrixLayout
func (a_ AccelerationStructureMotionTriangleGeometryDescriptor) TransformationMatrixLayout() MatrixLayout {
	rv := objc.Send[MatrixLayout](a_.ID, objc.Sel("transformationMatrixLayout"))
	return rv
}/* debug [instance_properties/getter]: transformationMatrixLayout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionTriangleGeometryDescriptor/transformationMatrixLayout
func (a_ AccelerationStructureMotionTriangleGeometryDescriptor) SetTransformationMatrixLayout(value MatrixLayout) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTransformationMatrixLayout:"), value)
}/* debug [instance_properties/setter]: transformationMatrixLayout */


// The number of triangles in the buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionTriangleGeometryDescriptor/triangleCount
func (a_ AccelerationStructureMotionTriangleGeometryDescriptor) TriangleCount() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("triangleCount"))
	return rv
}/* debug [instance_properties/getter]: triangleCount */


// The number of triangles in the buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionTriangleGeometryDescriptor/triangleCount
func (a_ AccelerationStructureMotionTriangleGeometryDescriptor) SetTriangleCount(value uint) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTriangleCount:"), value)
}/* debug [instance_properties/setter]: triangleCount */


// An array of motion keyframes, each containing triangle data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionTriangleGeometryDescriptor/vertexBuffers
func (a_ AccelerationStructureMotionTriangleGeometryDescriptor) VertexBuffers() []MotionKeyframeData {
	rv := objc.Send[[]MotionKeyframeData](a_.ID, objc.Sel("vertexBuffers"))
	return rv
}/* debug [instance_properties/getter]: vertexBuffers */


// An array of motion keyframes, each containing triangle data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionTriangleGeometryDescriptor/vertexBuffers
func (a_ AccelerationStructureMotionTriangleGeometryDescriptor) SetVertexBuffers(value []MotionKeyframeData) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](a_.ID, objc.Sel("setVertexBuffers:"), nsArray)
}/* debug [instance_properties/setter]: vertexBuffers */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionTriangleGeometryDescriptor/vertexFormat
func (a_ AccelerationStructureMotionTriangleGeometryDescriptor) VertexFormat() AttributeFormat {
	rv := objc.Send[AttributeFormat](a_.ID, objc.Sel("vertexFormat"))
	return rv
}/* debug [instance_properties/getter]: vertexFormat */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionTriangleGeometryDescriptor/vertexFormat
func (a_ AccelerationStructureMotionTriangleGeometryDescriptor) SetVertexFormat(value AttributeFormat) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setVertexFormat:"), value)
}/* debug [instance_properties/setter]: vertexFormat */


// The stride, in bytes, between vertices in each vertex buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionTriangleGeometryDescriptor/vertexStride
func (a_ AccelerationStructureMotionTriangleGeometryDescriptor) VertexStride() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("vertexStride"))
	return rv
}/* debug [instance_properties/getter]: vertexStride */


// The stride, in bytes, between vertices in each vertex buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionTriangleGeometryDescriptor/vertexStride
func (a_ AccelerationStructureMotionTriangleGeometryDescriptor) SetVertexStride(value uint) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setVertexStride:"), value)
}/* debug [instance_properties/setter]: vertexStride */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLAccelerationStructureMotionTriangleGeometryDescriptor */



