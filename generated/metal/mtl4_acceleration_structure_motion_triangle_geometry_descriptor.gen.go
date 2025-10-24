// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MTL4AccelerationStructureMotionTriangleGeometryDescriptor */


/* debug [class_header]: Header for MTL4AccelerationStructureMotionTriangleGeometryDescriptor */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTL4AccelerationStructureMotionTriangleGeometryDescriptor */
// An interface definition for the [MTL4AccelerationStructureMotionTriangleGeometryDescriptor] class.
type IMTL4AccelerationStructureMotionTriangleGeometryDescriptor interface {
	IMTL4AccelerationStructureGeometryDescriptor
	
/* debug [class_interface_properties]: Properties for MTL4AccelerationStructureMotionTriangleGeometryDescriptor */
	// properties:
	IndexBuffer() objc.IObject /* cross-framework: MTL4BufferRange */
	SetIndexBuffer(value objc.IObject /* cross-framework: MTL4BufferRange */)
	IndexType() IndexType
	SetIndexType(value IndexType)
	TransformationMatrixBuffer() objc.IObject /* cross-framework: MTL4BufferRange */
	SetTransformationMatrixBuffer(value objc.IObject /* cross-framework: MTL4BufferRange */)
	TransformationMatrixLayout() MatrixLayout
	SetTransformationMatrixLayout(value MatrixLayout)
	TriangleCount() uint
	SetTriangleCount(value uint)
	VertexBuffers() objc.IObject /* cross-framework: MTL4BufferRange */
	SetVertexBuffers(value objc.IObject /* cross-framework: MTL4BufferRange */)
	VertexFormat() AttributeFormat
	SetVertexFormat(value AttributeFormat)
	VertexStride() uint
	SetVertexStride(value uint)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTL4AccelerationStructureMotionTriangleGeometryDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTL4AccelerationStructureMotionTriangleGeometryDescriptor */
// Alloc allocates a new instance without initialization.
func (mc _MTL4AccelerationStructureMotionTriangleGeometryDescriptorClass) Alloc() MTL4AccelerationStructureMotionTriangleGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureMotionTriangleGeometryDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTL4AccelerationStructureMotionTriangleGeometryDescriptor */
// Describes motion triangle geometry, suitable for motion ray tracing.
//
// Use a to mark residency of all buffers this descriptor references when you build this acceleration structure.


// Describes motion triangle geometry, suitable for motion ray tracing.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTL4AccelerationStructureMotionTriangleGeometryDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTL4AccelerationStructureMotionTriangleGeometryDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTL4AccelerationStructureMotionTriangleGeometryDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTL4AccelerationStructureMotionTriangleGeometryDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTL4AccelerationStructureMotionTriangleGeometryDescriptor */

// Assigns an optional index buffer containing references to vertices in the vertex buffers you reference through the vertex buffers property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionTriangleGeometryDescriptor/indexBuffer
func (m_ MTL4AccelerationStructureMotionTriangleGeometryDescriptor) IndexBuffer() objc.IObject /* cross-framework: MTL4BufferRange */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("indexBuffer"))
	return rv
}/* debug [instance_properties/getter]: indexBuffer */


// Assigns an optional index buffer containing references to vertices in the vertex buffers you reference through the vertex buffers property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionTriangleGeometryDescriptor/indexBuffer
func (m_ MTL4AccelerationStructureMotionTriangleGeometryDescriptor) SetIndexBuffer(value objc.IObject /* cross-framework: MTL4BufferRange */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIndexBuffer:"), value)
}/* debug [instance_properties/setter]: indexBuffer */


// Specifies the size of the indices the contains, which is typically either 16 or 32-bits for each index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionTriangleGeometryDescriptor/indexType
func (m_ MTL4AccelerationStructureMotionTriangleGeometryDescriptor) IndexType() IndexType {
	rv := objc.Send[IndexType](m_.ID, objc.Sel("indexType"))
	return rv
}/* debug [instance_properties/getter]: indexType */


// Specifies the size of the indices the contains, which is typically either 16 or 32-bits for each index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionTriangleGeometryDescriptor/indexType
func (m_ MTL4AccelerationStructureMotionTriangleGeometryDescriptor) SetIndexType(value IndexType) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIndexType:"), value)
}/* debug [instance_properties/setter]: indexType */


// Assings an optional reference to a buffer containing a transformation matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionTriangleGeometryDescriptor/transformationMatrixBuffer
func (m_ MTL4AccelerationStructureMotionTriangleGeometryDescriptor) TransformationMatrixBuffer() objc.IObject /* cross-framework: MTL4BufferRange */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("transformationMatrixBuffer"))
	return rv
}/* debug [instance_properties/getter]: transformationMatrixBuffer */


// Assings an optional reference to a buffer containing a transformation matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionTriangleGeometryDescriptor/transformationMatrixBuffer
func (m_ MTL4AccelerationStructureMotionTriangleGeometryDescriptor) SetTransformationMatrixBuffer(value objc.IObject /* cross-framework: MTL4BufferRange */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTransformationMatrixBuffer:"), value)
}/* debug [instance_properties/setter]: transformationMatrixBuffer */


// Configures the layout for the transformation matrix in the transformation matrix buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionTriangleGeometryDescriptor/transformationMatrixLayout
func (m_ MTL4AccelerationStructureMotionTriangleGeometryDescriptor) TransformationMatrixLayout() MatrixLayout {
	rv := objc.Send[MatrixLayout](m_.ID, objc.Sel("transformationMatrixLayout"))
	return rv
}/* debug [instance_properties/getter]: transformationMatrixLayout */


// Configures the layout for the transformation matrix in the transformation matrix buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionTriangleGeometryDescriptor/transformationMatrixLayout
func (m_ MTL4AccelerationStructureMotionTriangleGeometryDescriptor) SetTransformationMatrixLayout(value MatrixLayout) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTransformationMatrixLayout:"), value)
}/* debug [instance_properties/setter]: transformationMatrixLayout */


// Declares the number of triangles in the vertex buffers that the buffer in the vertex buffers property references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionTriangleGeometryDescriptor/triangleCount
func (m_ MTL4AccelerationStructureMotionTriangleGeometryDescriptor) TriangleCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("triangleCount"))
	return rv
}/* debug [instance_properties/getter]: triangleCount */


// Declares the number of triangles in the vertex buffers that the buffer in the vertex buffers property references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionTriangleGeometryDescriptor/triangleCount
func (m_ MTL4AccelerationStructureMotionTriangleGeometryDescriptor) SetTriangleCount(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTriangleCount:"), value)
}/* debug [instance_properties/setter]: triangleCount */


// Assigns a buffer where each entry contains a reference to a vertex buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionTriangleGeometryDescriptor/vertexBuffers
func (m_ MTL4AccelerationStructureMotionTriangleGeometryDescriptor) VertexBuffers() objc.IObject /* cross-framework: MTL4BufferRange */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("vertexBuffers"))
	return rv
}/* debug [instance_properties/getter]: vertexBuffers */


// Assigns a buffer where each entry contains a reference to a vertex buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionTriangleGeometryDescriptor/vertexBuffers
func (m_ MTL4AccelerationStructureMotionTriangleGeometryDescriptor) SetVertexBuffers(value objc.IObject /* cross-framework: MTL4BufferRange */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVertexBuffers:"), value)
}/* debug [instance_properties/setter]: vertexBuffers */


// Defines the format of the vertices in the vertex buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionTriangleGeometryDescriptor/vertexFormat
func (m_ MTL4AccelerationStructureMotionTriangleGeometryDescriptor) VertexFormat() AttributeFormat {
	rv := objc.Send[AttributeFormat](m_.ID, objc.Sel("vertexFormat"))
	return rv
}/* debug [instance_properties/getter]: vertexFormat */


// Defines the format of the vertices in the vertex buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionTriangleGeometryDescriptor/vertexFormat
func (m_ MTL4AccelerationStructureMotionTriangleGeometryDescriptor) SetVertexFormat(value AttributeFormat) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVertexFormat:"), value)
}/* debug [instance_properties/setter]: vertexFormat */


// Sets the stride, in bytes, between vertices in all the vertex buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionTriangleGeometryDescriptor/vertexStride
func (m_ MTL4AccelerationStructureMotionTriangleGeometryDescriptor) VertexStride() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("vertexStride"))
	return rv
}/* debug [instance_properties/getter]: vertexStride */


// Sets the stride, in bytes, between vertices in all the vertex buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionTriangleGeometryDescriptor/vertexStride
func (m_ MTL4AccelerationStructureMotionTriangleGeometryDescriptor) SetVertexStride(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVertexStride:"), value)
}/* debug [instance_properties/setter]: vertexStride */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTL4AccelerationStructureMotionTriangleGeometryDescriptor */



