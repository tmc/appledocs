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
}

// Describes curve geometry suitable for ray tracing.
//
// Use a to mark residency of all buffers this descriptor references when you build this acceleration structure.
//
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

// Alloc allocates a new instance without initialization.
func (mc _MTL4AccelerationStructureCurveGeometryDescriptorClass) Alloc() MTL4AccelerationStructureCurveGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureCurveGeometryDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// References a buffer containing curve control points.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructurecurvegeometrydescriptor/controlpointbuffer
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) ControlPointBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("controlPointBuffer"))
	return rv
}


// SetControlPointBuffer sets the value of the controlPointBuffer property.
// References a buffer containing curve control points.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructurecurvegeometrydescriptor/controlpointbuffer
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) SetControlPointBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setControlPointBuffer:"), value)
}

// Declares the number of control points in the control point buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructurecurvegeometrydescriptor/controlpointcount
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) ControlPointCount() int {
	rv := objc.Send[int](m_.ID, objc.Sel("controlPointCount"))
	return rv
}


// SetControlPointCount sets the value of the controlPointCount property.
// Declares the number of control points in the control point buffer.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructurecurvegeometrydescriptor/controlpointcount
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) SetControlPointCount(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setControlPointCount:"), value)
}

// Declares the format of the control points the control point buffer references.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructurecurvegeometrydescriptor/controlpointformat
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) ControlPointFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("controlPointFormat"))
	return rv
}


// SetControlPointFormat sets the value of the controlPointFormat property.
// Declares the format of the control points the control point buffer references.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructurecurvegeometrydescriptor/controlpointformat
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) SetControlPointFormat(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setControlPointFormat:"), value)
}

// Sets the stride, in bytes, between control points in the control point buffer the control point buffer references.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructurecurvegeometrydescriptor/controlpointstride
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) ControlPointStride() int {
	rv := objc.Send[int](m_.ID, objc.Sel("controlPointStride"))
	return rv
}


// SetControlPointStride sets the value of the controlPointStride property.
// Sets the stride, in bytes, between control points in the control point buffer the control point buffer references.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructurecurvegeometrydescriptor/controlpointstride
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) SetControlPointStride(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setControlPointStride:"), value)
}

// Controls the curve basis function, determining how Metal interpolates the control points.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructurecurvegeometrydescriptor/curvebasis
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) CurveBasis() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("curveBasis"))
	return rv
}


// SetCurveBasis sets the value of the curveBasis property.
// Controls the curve basis function, determining how Metal interpolates the control points.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructurecurvegeometrydescriptor/curvebasis
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) SetCurveBasis(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCurveBasis:"), value)
}

// Sets the type of curve end caps.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructurecurvegeometrydescriptor/curveendcaps
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) CurveEndCaps() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("curveEndCaps"))
	return rv
}


// SetCurveEndCaps sets the value of the curveEndCaps property.
// Sets the type of curve end caps.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructurecurvegeometrydescriptor/curveendcaps
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) SetCurveEndCaps(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCurveEndCaps:"), value)
}

// Controls the curve type.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructurecurvegeometrydescriptor/curvetype
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) CurveType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("curveType"))
	return rv
}


// SetCurveType sets the value of the curveType property.
// Controls the curve type.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructurecurvegeometrydescriptor/curvetype
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) SetCurveType(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCurveType:"), value)
}

// Assigns an optional index buffer containing references to control points in the control point buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructurecurvegeometrydescriptor/indexbuffer
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) IndexBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("indexBuffer"))
	return rv
}


// SetIndexBuffer sets the value of the indexBuffer property.
// Assigns an optional index buffer containing references to control points in the control point buffer.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructurecurvegeometrydescriptor/indexbuffer
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) SetIndexBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIndexBuffer:"), value)
}

// Specifies the size of the indices the
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructurecurvegeometrydescriptor/indextype
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) IndexType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("indexType"))
	return rv
}


// SetIndexType sets the value of the indexType property.
// Specifies the size of the indices the

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructurecurvegeometrydescriptor/indextype
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) SetIndexType(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIndexType:"), value)
}

// Assigns a reference to a buffer containing the curve radius for each control point.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructurecurvegeometrydescriptor/radiusbuffer
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) RadiusBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("radiusBuffer"))
	return rv
}


// SetRadiusBuffer sets the value of the radiusBuffer property.
// Assigns a reference to a buffer containing the curve radius for each control point.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructurecurvegeometrydescriptor/radiusbuffer
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) SetRadiusBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRadiusBuffer:"), value)
}

// Declares the format of the radii in the radius buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructurecurvegeometrydescriptor/radiusformat
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) RadiusFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("radiusFormat"))
	return rv
}


// SetRadiusFormat sets the value of the radiusFormat property.
// Declares the format of the radii in the radius buffer.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructurecurvegeometrydescriptor/radiusformat
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) SetRadiusFormat(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRadiusFormat:"), value)
}

// Configures the stride, in bytes, between radii in the radius buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructurecurvegeometrydescriptor/radiusstride
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) RadiusStride() int {
	rv := objc.Send[int](m_.ID, objc.Sel("radiusStride"))
	return rv
}


// SetRadiusStride sets the value of the radiusStride property.
// Configures the stride, in bytes, between radii in the radius buffer.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructurecurvegeometrydescriptor/radiusstride
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) SetRadiusStride(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRadiusStride:"), value)
}

// Declares the number of control points per curve segment.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructurecurvegeometrydescriptor/segmentcontrolpointcount
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) SegmentControlPointCount() int {
	rv := objc.Send[int](m_.ID, objc.Sel("segmentControlPointCount"))
	return rv
}


// SetSegmentControlPointCount sets the value of the segmentControlPointCount property.
// Declares the number of control points per curve segment.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructurecurvegeometrydescriptor/segmentcontrolpointcount
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) SetSegmentControlPointCount(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSegmentControlPointCount:"), value)
}

// Declares the number of curve segments.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructurecurvegeometrydescriptor/segmentcount
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) SegmentCount() int {
	rv := objc.Send[int](m_.ID, objc.Sel("segmentCount"))
	return rv
}


// SetSegmentCount sets the value of the segmentCount property.
// Declares the number of curve segments.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructurecurvegeometrydescriptor/segmentcount
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) SetSegmentCount(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSegmentCount:"), value)
}



