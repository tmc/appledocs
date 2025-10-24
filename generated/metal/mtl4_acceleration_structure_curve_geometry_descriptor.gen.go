// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MTL4AccelerationStructureCurveGeometryDescriptor */


/* debug [class_header]: Header for MTL4AccelerationStructureCurveGeometryDescriptor */
// The class instance for the [MTL4AccelerationStructureCurveGeometryDescriptor] class.
var (
	MTL4AccelerationStructureCurveGeometryDescriptorClass     _MTL4AccelerationStructureCurveGeometryDescriptorClass
	MTL4AccelerationStructureCurveGeometryDescriptorClassOnce sync.Once
)

func getMTL4AccelerationStructureCurveGeometryDescriptorClass() _MTL4AccelerationStructureCurveGeometryDescriptorClass {
	MTL4AccelerationStructureCurveGeometryDescriptorClassOnce.Do(func() {
		MTL4AccelerationStructureCurveGeometryDescriptorClass = _MTL4AccelerationStructureCurveGeometryDescriptorClass{objc.GetClass("MTL4AccelerationStructureCurveGeometryDescriptor")}
	})
	return MTL4AccelerationStructureCurveGeometryDescriptorClass
}

type _MTL4AccelerationStructureCurveGeometryDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTL4AccelerationStructureCurveGeometryDescriptor */
// An interface definition for the [MTL4AccelerationStructureCurveGeometryDescriptor] class.
type IMTL4AccelerationStructureCurveGeometryDescriptor interface {
	IMTL4AccelerationStructureGeometryDescriptor
	
/* debug [class_interface_properties]: Properties for MTL4AccelerationStructureCurveGeometryDescriptor */
	// properties:
	ControlPointBuffer() objc.IObject /* cross-framework: MTL4BufferRange */
	SetControlPointBuffer(value objc.IObject /* cross-framework: MTL4BufferRange */)
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
	RadiusBuffer() objc.IObject /* cross-framework: MTL4BufferRange */
	SetRadiusBuffer(value objc.IObject /* cross-framework: MTL4BufferRange */)
	RadiusFormat() AttributeFormat
	SetRadiusFormat(value AttributeFormat)
	RadiusStride() uint
	SetRadiusStride(value uint)
	SegmentControlPointCount() uint
	SetSegmentControlPointCount(value uint)
	SegmentCount() uint
	SetSegmentCount(value uint)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTL4AccelerationStructureCurveGeometryDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTL4AccelerationStructureCurveGeometryDescriptor */
// Alloc allocates a new instance without initialization.
func (mc _MTL4AccelerationStructureCurveGeometryDescriptorClass) Alloc() MTL4AccelerationStructureCurveGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureCurveGeometryDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTL4AccelerationStructureCurveGeometryDescriptorClass) New() MTL4AccelerationStructureCurveGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureCurveGeometryDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) Init() MTL4AccelerationStructureCurveGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureCurveGeometryDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) Autorelease() MTL4AccelerationStructureCurveGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureCurveGeometryDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4AccelerationStructureCurveGeometryDescriptor creates a new MTL4AccelerationStructureCurveGeometryDescriptor instance.
func NewMTL4AccelerationStructureCurveGeometryDescriptor() MTL4AccelerationStructureCurveGeometryDescriptor {
	return getMTL4AccelerationStructureCurveGeometryDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTL4AccelerationStructureCurveGeometryDescriptor */
// Describes curve geometry suitable for ray tracing.
//
// Use a to mark residency of all buffers this descriptor references when you build this acceleration structure.


// Describes curve geometry suitable for ray tracing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureCurveGeometryDescriptor
type MTL4AccelerationStructureCurveGeometryDescriptor struct {
	MTL4AccelerationStructureGeometryDescriptor
}

// MTL4AccelerationStructureCurveGeometryDescriptorFrom constructs a [MTL4AccelerationStructureCurveGeometryDescriptor] from an unsafe.Pointer.
//
// Describes curve geometry suitable for ray tracing.
func MTL4AccelerationStructureCurveGeometryDescriptorFrom(ptr unsafe.Pointer) MTL4AccelerationStructureCurveGeometryDescriptor {
	return MTL4AccelerationStructureCurveGeometryDescriptor{
		MTL4AccelerationStructureGeometryDescriptor: MTL4AccelerationStructureGeometryDescriptorFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTL4AccelerationStructureCurveGeometryDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTL4AccelerationStructureCurveGeometryDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTL4AccelerationStructureCurveGeometryDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTL4AccelerationStructureCurveGeometryDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTL4AccelerationStructureCurveGeometryDescriptor */

// References a buffer containing curve control points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureCurveGeometryDescriptor/controlPointBuffer
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) ControlPointBuffer() objc.IObject /* cross-framework: MTL4BufferRange */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("controlPointBuffer"))
	return rv
}/* debug [instance_properties/getter]: controlPointBuffer */


// References a buffer containing curve control points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureCurveGeometryDescriptor/controlPointBuffer
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) SetControlPointBuffer(value objc.IObject /* cross-framework: MTL4BufferRange */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setControlPointBuffer:"), value)
}/* debug [instance_properties/setter]: controlPointBuffer */


// Declares the number of control points in the control point buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureCurveGeometryDescriptor/controlPointCount
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) ControlPointCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("controlPointCount"))
	return rv
}/* debug [instance_properties/getter]: controlPointCount */


// Declares the number of control points in the control point buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureCurveGeometryDescriptor/controlPointCount
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) SetControlPointCount(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setControlPointCount:"), value)
}/* debug [instance_properties/setter]: controlPointCount */


// Declares the format of the control points the control point buffer references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureCurveGeometryDescriptor/controlPointFormat
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) ControlPointFormat() AttributeFormat {
	rv := objc.Send[AttributeFormat](m_.ID, objc.Sel("controlPointFormat"))
	return rv
}/* debug [instance_properties/getter]: controlPointFormat */


// Declares the format of the control points the control point buffer references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureCurveGeometryDescriptor/controlPointFormat
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) SetControlPointFormat(value AttributeFormat) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setControlPointFormat:"), value)
}/* debug [instance_properties/setter]: controlPointFormat */


// Sets the stride, in bytes, between control points in the control point buffer the control point buffer references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureCurveGeometryDescriptor/controlPointStride
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) ControlPointStride() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("controlPointStride"))
	return rv
}/* debug [instance_properties/getter]: controlPointStride */


// Sets the stride, in bytes, between control points in the control point buffer the control point buffer references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureCurveGeometryDescriptor/controlPointStride
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) SetControlPointStride(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setControlPointStride:"), value)
}/* debug [instance_properties/setter]: controlPointStride */


// Controls the curve basis function, determining how Metal interpolates the control points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureCurveGeometryDescriptor/curveBasis
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) CurveBasis() CurveBasis {
	rv := objc.Send[CurveBasis](m_.ID, objc.Sel("curveBasis"))
	return rv
}/* debug [instance_properties/getter]: curveBasis */


// Controls the curve basis function, determining how Metal interpolates the control points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureCurveGeometryDescriptor/curveBasis
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) SetCurveBasis(value CurveBasis) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCurveBasis:"), value)
}/* debug [instance_properties/setter]: curveBasis */


// Sets the type of curve end caps.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureCurveGeometryDescriptor/curveEndCaps
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) CurveEndCaps() CurveEndCaps {
	rv := objc.Send[CurveEndCaps](m_.ID, objc.Sel("curveEndCaps"))
	return rv
}/* debug [instance_properties/getter]: curveEndCaps */


// Sets the type of curve end caps.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureCurveGeometryDescriptor/curveEndCaps
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) SetCurveEndCaps(value CurveEndCaps) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCurveEndCaps:"), value)
}/* debug [instance_properties/setter]: curveEndCaps */


// Controls the curve type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureCurveGeometryDescriptor/curveType
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) CurveType() CurveType {
	rv := objc.Send[CurveType](m_.ID, objc.Sel("curveType"))
	return rv
}/* debug [instance_properties/getter]: curveType */


// Controls the curve type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureCurveGeometryDescriptor/curveType
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) SetCurveType(value CurveType) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCurveType:"), value)
}/* debug [instance_properties/setter]: curveType */


// Assigns an optional index buffer containing references to control points in the control point buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureCurveGeometryDescriptor/indexBuffer
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) IndexBuffer() objc.IObject /* cross-framework: MTL4BufferRange */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("indexBuffer"))
	return rv
}/* debug [instance_properties/getter]: indexBuffer */


// Assigns an optional index buffer containing references to control points in the control point buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureCurveGeometryDescriptor/indexBuffer
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) SetIndexBuffer(value objc.IObject /* cross-framework: MTL4BufferRange */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIndexBuffer:"), value)
}/* debug [instance_properties/setter]: indexBuffer */


// Specifies the size of the indices the contains, which is typically either 16 or 32-bits for each index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureCurveGeometryDescriptor/indexType
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) IndexType() IndexType {
	rv := objc.Send[IndexType](m_.ID, objc.Sel("indexType"))
	return rv
}/* debug [instance_properties/getter]: indexType */


// Specifies the size of the indices the contains, which is typically either 16 or 32-bits for each index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureCurveGeometryDescriptor/indexType
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) SetIndexType(value IndexType) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIndexType:"), value)
}/* debug [instance_properties/setter]: indexType */


// Assigns a reference to a buffer containing the curve radius for each control point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureCurveGeometryDescriptor/radiusBuffer
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) RadiusBuffer() objc.IObject /* cross-framework: MTL4BufferRange */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("radiusBuffer"))
	return rv
}/* debug [instance_properties/getter]: radiusBuffer */


// Assigns a reference to a buffer containing the curve radius for each control point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureCurveGeometryDescriptor/radiusBuffer
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) SetRadiusBuffer(value objc.IObject /* cross-framework: MTL4BufferRange */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRadiusBuffer:"), value)
}/* debug [instance_properties/setter]: radiusBuffer */


// Declares the format of the radii in the radius buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureCurveGeometryDescriptor/radiusFormat
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) RadiusFormat() AttributeFormat {
	rv := objc.Send[AttributeFormat](m_.ID, objc.Sel("radiusFormat"))
	return rv
}/* debug [instance_properties/getter]: radiusFormat */


// Declares the format of the radii in the radius buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureCurveGeometryDescriptor/radiusFormat
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) SetRadiusFormat(value AttributeFormat) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRadiusFormat:"), value)
}/* debug [instance_properties/setter]: radiusFormat */


// Configures the stride, in bytes, between radii in the radius buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureCurveGeometryDescriptor/radiusStride
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) RadiusStride() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("radiusStride"))
	return rv
}/* debug [instance_properties/getter]: radiusStride */


// Configures the stride, in bytes, between radii in the radius buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureCurveGeometryDescriptor/radiusStride
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) SetRadiusStride(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRadiusStride:"), value)
}/* debug [instance_properties/setter]: radiusStride */


// Declares the number of control points per curve segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureCurveGeometryDescriptor/segmentControlPointCount
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) SegmentControlPointCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("segmentControlPointCount"))
	return rv
}/* debug [instance_properties/getter]: segmentControlPointCount */


// Declares the number of control points per curve segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureCurveGeometryDescriptor/segmentControlPointCount
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) SetSegmentControlPointCount(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSegmentControlPointCount:"), value)
}/* debug [instance_properties/setter]: segmentControlPointCount */


// Declares the number of curve segments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureCurveGeometryDescriptor/segmentCount
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) SegmentCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("segmentCount"))
	return rv
}/* debug [instance_properties/getter]: segmentCount */


// Declares the number of curve segments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureCurveGeometryDescriptor/segmentCount
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) SetSegmentCount(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSegmentCount:"), value)
}/* debug [instance_properties/setter]: segmentCount */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTL4AccelerationStructureCurveGeometryDescriptor */



