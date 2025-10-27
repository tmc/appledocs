// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





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





// An interface definition for the [MTL4AccelerationStructureCurveGeometryDescriptor] class.
type IMTL4AccelerationStructureCurveGeometryDescriptor interface {
	IMTL4AccelerationStructureGeometryDescriptor
	

	// properties:
	ControlPointBuffer() MTL4BufferRange
	SetControlPointBuffer(value MTL4BufferRange)
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
	IndexBuffer() MTL4BufferRange
	SetIndexBuffer(value MTL4BufferRange)
	IndexType() IndexType
	SetIndexType(value IndexType)
	RadiusBuffer() MTL4BufferRange
	SetRadiusBuffer(value MTL4BufferRange)
	RadiusFormat() AttributeFormat
	SetRadiusFormat(value AttributeFormat)
	RadiusStride() uint
	SetRadiusStride(value uint)
	SegmentControlPointCount() uint
	SetSegmentControlPointCount(value uint)
	SegmentCount() uint
	SetSegmentCount(value uint)


	

	// methods:


}





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

























// References a buffer containing curve control points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureCurveGeometryDescriptor/controlPointBuffer
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) ControlPointBuffer() MTL4BufferRange {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("controlPointBuffer"))
	return rv
}


// References a buffer containing curve control points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureCurveGeometryDescriptor/controlPointBuffer
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) SetControlPointBuffer(value MTL4BufferRange) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setControlPointBuffer:"), value)
}


// Declares the number of control points in the control point buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureCurveGeometryDescriptor/controlPointCount
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) ControlPointCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("controlPointCount"))
	return rv
}


// Declares the number of control points in the control point buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureCurveGeometryDescriptor/controlPointCount
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) SetControlPointCount(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setControlPointCount:"), value)
}


// Declares the format of the control points the control point buffer references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureCurveGeometryDescriptor/controlPointFormat
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) ControlPointFormat() AttributeFormat {
	rv := objc.Send[AttributeFormat](m_.ID, objc.Sel("controlPointFormat"))
	return rv
}


// Declares the format of the control points the control point buffer references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureCurveGeometryDescriptor/controlPointFormat
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) SetControlPointFormat(value AttributeFormat) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setControlPointFormat:"), value)
}


// Sets the stride, in bytes, between control points in the control point buffer the control point buffer references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureCurveGeometryDescriptor/controlPointStride
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) ControlPointStride() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("controlPointStride"))
	return rv
}


// Sets the stride, in bytes, between control points in the control point buffer the control point buffer references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureCurveGeometryDescriptor/controlPointStride
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) SetControlPointStride(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setControlPointStride:"), value)
}


// Controls the curve basis function, determining how Metal interpolates the control points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureCurveGeometryDescriptor/curveBasis
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) CurveBasis() CurveBasis {
	rv := objc.Send[CurveBasis](m_.ID, objc.Sel("curveBasis"))
	return rv
}


// Controls the curve basis function, determining how Metal interpolates the control points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureCurveGeometryDescriptor/curveBasis
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) SetCurveBasis(value CurveBasis) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCurveBasis:"), value)
}


// Sets the type of curve end caps.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureCurveGeometryDescriptor/curveEndCaps
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) CurveEndCaps() CurveEndCaps {
	rv := objc.Send[CurveEndCaps](m_.ID, objc.Sel("curveEndCaps"))
	return rv
}


// Sets the type of curve end caps.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureCurveGeometryDescriptor/curveEndCaps
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) SetCurveEndCaps(value CurveEndCaps) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCurveEndCaps:"), value)
}


// Controls the curve type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureCurveGeometryDescriptor/curveType
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) CurveType() CurveType {
	rv := objc.Send[CurveType](m_.ID, objc.Sel("curveType"))
	return rv
}


// Controls the curve type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureCurveGeometryDescriptor/curveType
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) SetCurveType(value CurveType) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCurveType:"), value)
}


// Assigns an optional index buffer containing references to control points in the control point buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureCurveGeometryDescriptor/indexBuffer
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) IndexBuffer() MTL4BufferRange {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("indexBuffer"))
	return rv
}


// Assigns an optional index buffer containing references to control points in the control point buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureCurveGeometryDescriptor/indexBuffer
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) SetIndexBuffer(value MTL4BufferRange) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIndexBuffer:"), value)
}


// Specifies the size of the indices the contains, which is typically either 16 or 32-bits for each index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureCurveGeometryDescriptor/indexType
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) IndexType() IndexType {
	rv := objc.Send[IndexType](m_.ID, objc.Sel("indexType"))
	return rv
}


// Specifies the size of the indices the contains, which is typically either 16 or 32-bits for each index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureCurveGeometryDescriptor/indexType
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) SetIndexType(value IndexType) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIndexType:"), value)
}


// Assigns a reference to a buffer containing the curve radius for each control point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureCurveGeometryDescriptor/radiusBuffer
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) RadiusBuffer() MTL4BufferRange {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("radiusBuffer"))
	return rv
}


// Assigns a reference to a buffer containing the curve radius for each control point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureCurveGeometryDescriptor/radiusBuffer
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) SetRadiusBuffer(value MTL4BufferRange) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRadiusBuffer:"), value)
}


// Declares the format of the radii in the radius buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureCurveGeometryDescriptor/radiusFormat
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) RadiusFormat() AttributeFormat {
	rv := objc.Send[AttributeFormat](m_.ID, objc.Sel("radiusFormat"))
	return rv
}


// Declares the format of the radii in the radius buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureCurveGeometryDescriptor/radiusFormat
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) SetRadiusFormat(value AttributeFormat) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRadiusFormat:"), value)
}


// Configures the stride, in bytes, between radii in the radius buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureCurveGeometryDescriptor/radiusStride
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) RadiusStride() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("radiusStride"))
	return rv
}


// Configures the stride, in bytes, between radii in the radius buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureCurveGeometryDescriptor/radiusStride
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) SetRadiusStride(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRadiusStride:"), value)
}


// Declares the number of control points per curve segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureCurveGeometryDescriptor/segmentControlPointCount
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) SegmentControlPointCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("segmentControlPointCount"))
	return rv
}


// Declares the number of control points per curve segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureCurveGeometryDescriptor/segmentControlPointCount
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) SetSegmentControlPointCount(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSegmentControlPointCount:"), value)
}


// Declares the number of curve segments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureCurveGeometryDescriptor/segmentCount
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) SegmentCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("segmentCount"))
	return rv
}


// Declares the number of curve segments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureCurveGeometryDescriptor/segmentCount
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) SetSegmentCount(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSegmentCount:"), value)
}








