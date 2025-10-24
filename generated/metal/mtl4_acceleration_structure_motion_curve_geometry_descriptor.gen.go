// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MTL4AccelerationStructureMotionCurveGeometryDescriptor */


/* debug [class_header]: Header for MTL4AccelerationStructureMotionCurveGeometryDescriptor */
// The class instance for the [MTL4AccelerationStructureMotionCurveGeometryDescriptor] class.
var (
	MTL4AccelerationStructureMotionCurveGeometryDescriptorClass     _MTL4AccelerationStructureMotionCurveGeometryDescriptorClass
	MTL4AccelerationStructureMotionCurveGeometryDescriptorClassOnce sync.Once
)

func getMTL4AccelerationStructureMotionCurveGeometryDescriptorClass() _MTL4AccelerationStructureMotionCurveGeometryDescriptorClass {
	MTL4AccelerationStructureMotionCurveGeometryDescriptorClassOnce.Do(func() {
		MTL4AccelerationStructureMotionCurveGeometryDescriptorClass = _MTL4AccelerationStructureMotionCurveGeometryDescriptorClass{objc.GetClass("MTL4AccelerationStructureMotionCurveGeometryDescriptor")}
	})
	return MTL4AccelerationStructureMotionCurveGeometryDescriptorClass
}

type _MTL4AccelerationStructureMotionCurveGeometryDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTL4AccelerationStructureMotionCurveGeometryDescriptor */
// An interface definition for the [MTL4AccelerationStructureMotionCurveGeometryDescriptor] class.
type IMTL4AccelerationStructureMotionCurveGeometryDescriptor interface {
	IMTL4AccelerationStructureGeometryDescriptor
	
/* debug [class_interface_properties]: Properties for MTL4AccelerationStructureMotionCurveGeometryDescriptor */
	// properties:
	ControlPointBuffers() objc.IObject /* cross-framework: MTL4BufferRange */
	SetControlPointBuffers(value objc.IObject /* cross-framework: MTL4BufferRange */)
	ControlPointCount() uint
	SetControlPointCount(value uint)
	ControlPointFormat() AttributeFormat
	SetControlPointFormat(value AttributeFormat)
	ControlPointStride() uint
	SetControlPointStride(value uint)
	CurveBasis() CurveBasis
	SetCurveBasis(value CurveBasis)
	CurveEndCaps() CurveEndCaps
	SetCurveEndCaps(value CurveEndCaps)
	CurveType() CurveType
	SetCurveType(value CurveType)
	IndexBuffer() objc.IObject /* cross-framework: MTL4BufferRange */
	SetIndexBuffer(value objc.IObject /* cross-framework: MTL4BufferRange */)
	IndexType() IndexType
	SetIndexType(value IndexType)
	RadiusBuffers() objc.IObject /* cross-framework: MTL4BufferRange */
	SetRadiusBuffers(value objc.IObject /* cross-framework: MTL4BufferRange */)
	RadiusFormat() AttributeFormat
	SetRadiusFormat(value AttributeFormat)
	RadiusStride() uint
	SetRadiusStride(value uint)
	SegmentControlPointCount() uint
	SetSegmentControlPointCount(value uint)
	SegmentCount() uint
	SetSegmentCount(value uint)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTL4AccelerationStructureMotionCurveGeometryDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTL4AccelerationStructureMotionCurveGeometryDescriptor */
// Alloc allocates a new instance without initialization.
func (mc _MTL4AccelerationStructureMotionCurveGeometryDescriptorClass) Alloc() MTL4AccelerationStructureMotionCurveGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureMotionCurveGeometryDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTL4AccelerationStructureMotionCurveGeometryDescriptorClass) New() MTL4AccelerationStructureMotionCurveGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureMotionCurveGeometryDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTL4AccelerationStructureMotionCurveGeometryDescriptor) Init() MTL4AccelerationStructureMotionCurveGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureMotionCurveGeometryDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTL4AccelerationStructureMotionCurveGeometryDescriptor) Autorelease() MTL4AccelerationStructureMotionCurveGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureMotionCurveGeometryDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4AccelerationStructureMotionCurveGeometryDescriptor creates a new MTL4AccelerationStructureMotionCurveGeometryDescriptor instance.
func NewMTL4AccelerationStructureMotionCurveGeometryDescriptor() MTL4AccelerationStructureMotionCurveGeometryDescriptor {
	return getMTL4AccelerationStructureMotionCurveGeometryDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTL4AccelerationStructureMotionCurveGeometryDescriptor */
// Describes motion curve geometry, suitable for motion ray tracing.
//
// Use a to mark residency of all buffers this descriptor references when you build this acceleration structure.


// Describes motion curve geometry, suitable for motion ray tracing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionCurveGeometryDescriptor
type MTL4AccelerationStructureMotionCurveGeometryDescriptor struct {
	MTL4AccelerationStructureGeometryDescriptor
}

// MTL4AccelerationStructureMotionCurveGeometryDescriptorFrom constructs a [MTL4AccelerationStructureMotionCurveGeometryDescriptor] from an unsafe.Pointer.
//
// Describes motion curve geometry, suitable for motion ray tracing.
func MTL4AccelerationStructureMotionCurveGeometryDescriptorFrom(ptr unsafe.Pointer) MTL4AccelerationStructureMotionCurveGeometryDescriptor {
	return MTL4AccelerationStructureMotionCurveGeometryDescriptor{
		MTL4AccelerationStructureGeometryDescriptor: MTL4AccelerationStructureGeometryDescriptorFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTL4AccelerationStructureMotionCurveGeometryDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTL4AccelerationStructureMotionCurveGeometryDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTL4AccelerationStructureMotionCurveGeometryDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTL4AccelerationStructureMotionCurveGeometryDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTL4AccelerationStructureMotionCurveGeometryDescriptor */

// Assigns a reference to a buffer where each entry contains a reference to a buffer of control points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionCurveGeometryDescriptor/controlPointBuffers
func (m_ MTL4AccelerationStructureMotionCurveGeometryDescriptor) ControlPointBuffers() objc.IObject /* cross-framework: MTL4BufferRange */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("controlPointBuffers"))
	return rv
}/* debug [instance_properties/getter]: controlPointBuffers */


// Assigns a reference to a buffer where each entry contains a reference to a buffer of control points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionCurveGeometryDescriptor/controlPointBuffers
func (m_ MTL4AccelerationStructureMotionCurveGeometryDescriptor) SetControlPointBuffers(value objc.IObject /* cross-framework: MTL4BufferRange */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setControlPointBuffers:"), value)
}/* debug [instance_properties/setter]: controlPointBuffers */


// Specifies the number of control points in the buffers the control point buffers reference.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionCurveGeometryDescriptor/controlPointCount
func (m_ MTL4AccelerationStructureMotionCurveGeometryDescriptor) ControlPointCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("controlPointCount"))
	return rv
}/* debug [instance_properties/getter]: controlPointCount */


// Specifies the number of control points in the buffers the control point buffers reference.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionCurveGeometryDescriptor/controlPointCount
func (m_ MTL4AccelerationStructureMotionCurveGeometryDescriptor) SetControlPointCount(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setControlPointCount:"), value)
}/* debug [instance_properties/setter]: controlPointCount */


// Declares the format of the control points in the buffers that the control point buffers reference.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionCurveGeometryDescriptor/controlPointFormat
func (m_ MTL4AccelerationStructureMotionCurveGeometryDescriptor) ControlPointFormat() AttributeFormat {
	rv := objc.Send[AttributeFormat](m_.ID, objc.Sel("controlPointFormat"))
	return rv
}/* debug [instance_properties/getter]: controlPointFormat */


// Declares the format of the control points in the buffers that the control point buffers reference.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionCurveGeometryDescriptor/controlPointFormat
func (m_ MTL4AccelerationStructureMotionCurveGeometryDescriptor) SetControlPointFormat(value AttributeFormat) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setControlPointFormat:"), value)
}/* debug [instance_properties/setter]: controlPointFormat */


// Sets the stride, in bytes, between control points in the control point buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionCurveGeometryDescriptor/controlPointStride
func (m_ MTL4AccelerationStructureMotionCurveGeometryDescriptor) ControlPointStride() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("controlPointStride"))
	return rv
}/* debug [instance_properties/getter]: controlPointStride */


// Sets the stride, in bytes, between control points in the control point buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionCurveGeometryDescriptor/controlPointStride
func (m_ MTL4AccelerationStructureMotionCurveGeometryDescriptor) SetControlPointStride(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setControlPointStride:"), value)
}/* debug [instance_properties/setter]: controlPointStride */


// Sets the curve basis function, determining how Metal interpolates the control points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionCurveGeometryDescriptor/curveBasis
func (m_ MTL4AccelerationStructureMotionCurveGeometryDescriptor) CurveBasis() CurveBasis {
	rv := objc.Send[CurveBasis](m_.ID, objc.Sel("curveBasis"))
	return rv
}/* debug [instance_properties/getter]: curveBasis */


// Sets the curve basis function, determining how Metal interpolates the control points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionCurveGeometryDescriptor/curveBasis
func (m_ MTL4AccelerationStructureMotionCurveGeometryDescriptor) SetCurveBasis(value CurveBasis) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCurveBasis:"), value)
}/* debug [instance_properties/setter]: curveBasis */


// Configures the type of curve end caps.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionCurveGeometryDescriptor/curveEndCaps
func (m_ MTL4AccelerationStructureMotionCurveGeometryDescriptor) CurveEndCaps() CurveEndCaps {
	rv := objc.Send[CurveEndCaps](m_.ID, objc.Sel("curveEndCaps"))
	return rv
}/* debug [instance_properties/getter]: curveEndCaps */


// Configures the type of curve end caps.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionCurveGeometryDescriptor/curveEndCaps
func (m_ MTL4AccelerationStructureMotionCurveGeometryDescriptor) SetCurveEndCaps(value CurveEndCaps) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCurveEndCaps:"), value)
}/* debug [instance_properties/setter]: curveEndCaps */


// Controls the curve type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionCurveGeometryDescriptor/curveType
func (m_ MTL4AccelerationStructureMotionCurveGeometryDescriptor) CurveType() CurveType {
	rv := objc.Send[CurveType](m_.ID, objc.Sel("curveType"))
	return rv
}/* debug [instance_properties/getter]: curveType */


// Controls the curve type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionCurveGeometryDescriptor/curveType
func (m_ MTL4AccelerationStructureMotionCurveGeometryDescriptor) SetCurveType(value CurveType) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCurveType:"), value)
}/* debug [instance_properties/setter]: curveType */


// Assigns an optional index buffer containing references to control points in the control point buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionCurveGeometryDescriptor/indexBuffer
func (m_ MTL4AccelerationStructureMotionCurveGeometryDescriptor) IndexBuffer() objc.IObject /* cross-framework: MTL4BufferRange */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("indexBuffer"))
	return rv
}/* debug [instance_properties/getter]: indexBuffer */


// Assigns an optional index buffer containing references to control points in the control point buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionCurveGeometryDescriptor/indexBuffer
func (m_ MTL4AccelerationStructureMotionCurveGeometryDescriptor) SetIndexBuffer(value objc.IObject /* cross-framework: MTL4BufferRange */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIndexBuffer:"), value)
}/* debug [instance_properties/setter]: indexBuffer */


// Configures the size of the indices the contains, which is typically either 16 or 32-bits for each index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionCurveGeometryDescriptor/indexType
func (m_ MTL4AccelerationStructureMotionCurveGeometryDescriptor) IndexType() IndexType {
	rv := objc.Send[IndexType](m_.ID, objc.Sel("indexType"))
	return rv
}/* debug [instance_properties/getter]: indexType */


// Configures the size of the indices the contains, which is typically either 16 or 32-bits for each index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionCurveGeometryDescriptor/indexType
func (m_ MTL4AccelerationStructureMotionCurveGeometryDescriptor) SetIndexType(value IndexType) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIndexType:"), value)
}/* debug [instance_properties/setter]: indexType */


// Assigns a reference to a buffer containing, in turn, references to curve radii buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionCurveGeometryDescriptor/radiusBuffers
func (m_ MTL4AccelerationStructureMotionCurveGeometryDescriptor) RadiusBuffers() objc.IObject /* cross-framework: MTL4BufferRange */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("radiusBuffers"))
	return rv
}/* debug [instance_properties/getter]: radiusBuffers */


// Assigns a reference to a buffer containing, in turn, references to curve radii buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionCurveGeometryDescriptor/radiusBuffers
func (m_ MTL4AccelerationStructureMotionCurveGeometryDescriptor) SetRadiusBuffers(value objc.IObject /* cross-framework: MTL4BufferRange */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRadiusBuffers:"), value)
}/* debug [instance_properties/setter]: radiusBuffers */


// Sets the format of the radii in the radius buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionCurveGeometryDescriptor/radiusFormat
func (m_ MTL4AccelerationStructureMotionCurveGeometryDescriptor) RadiusFormat() AttributeFormat {
	rv := objc.Send[AttributeFormat](m_.ID, objc.Sel("radiusFormat"))
	return rv
}/* debug [instance_properties/getter]: radiusFormat */


// Sets the format of the radii in the radius buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionCurveGeometryDescriptor/radiusFormat
func (m_ MTL4AccelerationStructureMotionCurveGeometryDescriptor) SetRadiusFormat(value AttributeFormat) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRadiusFormat:"), value)
}/* debug [instance_properties/setter]: radiusFormat */


// Sets the stride, in bytes, between radii in the radius buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionCurveGeometryDescriptor/radiusStride
func (m_ MTL4AccelerationStructureMotionCurveGeometryDescriptor) RadiusStride() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("radiusStride"))
	return rv
}/* debug [instance_properties/getter]: radiusStride */


// Sets the stride, in bytes, between radii in the radius buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionCurveGeometryDescriptor/radiusStride
func (m_ MTL4AccelerationStructureMotionCurveGeometryDescriptor) SetRadiusStride(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRadiusStride:"), value)
}/* debug [instance_properties/setter]: radiusStride */


// Controls the number of control points per curve segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionCurveGeometryDescriptor/segmentControlPointCount
func (m_ MTL4AccelerationStructureMotionCurveGeometryDescriptor) SegmentControlPointCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("segmentControlPointCount"))
	return rv
}/* debug [instance_properties/getter]: segmentControlPointCount */


// Controls the number of control points per curve segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionCurveGeometryDescriptor/segmentControlPointCount
func (m_ MTL4AccelerationStructureMotionCurveGeometryDescriptor) SetSegmentControlPointCount(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSegmentControlPointCount:"), value)
}/* debug [instance_properties/setter]: segmentControlPointCount */


// Declares the number of curve segments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionCurveGeometryDescriptor/segmentCount
func (m_ MTL4AccelerationStructureMotionCurveGeometryDescriptor) SegmentCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("segmentCount"))
	return rv
}/* debug [instance_properties/getter]: segmentCount */


// Declares the number of curve segments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionCurveGeometryDescriptor/segmentCount
func (m_ MTL4AccelerationStructureMotionCurveGeometryDescriptor) SetSegmentCount(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSegmentCount:"), value)
}/* debug [instance_properties/setter]: segmentCount */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTL4AccelerationStructureMotionCurveGeometryDescriptor */



