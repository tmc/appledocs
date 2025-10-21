// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [MTL4AccelerationStructureMotionCurveGeometryDescriptor] class.
type IMTL4AccelerationStructureMotionCurveGeometryDescriptor interface {
	IMTL4AccelerationStructureGeometryDescriptor
}

// Describes motion curve geometry, suitable for motion ray tracing.
//
// Use a to mark residency of all buffers this descriptor references when you build this acceleration structure.
//
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

// Alloc allocates a new instance without initialization.
func (mc _MTL4AccelerationStructureMotionCurveGeometryDescriptorClass) Alloc() MTL4AccelerationStructureMotionCurveGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureMotionCurveGeometryDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Assigns an optional index buffer containing references to control points in the control point buffers.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuremotioncurvegeometrydescriptor/indexbuffer
func (m_ MTL4AccelerationStructureMotionCurveGeometryDescriptor) IndexBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("indexBuffer"))
	return rv
}


// SetIndexBuffer sets the value of the indexBuffer property.
// Assigns an optional index buffer containing references to control points in the control point buffers.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuremotioncurvegeometrydescriptor/indexbuffer
func (m_ MTL4AccelerationStructureMotionCurveGeometryDescriptor) SetIndexBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIndexBuffer:"), value)
}

// Controls the number of control points per curve segment.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuremotioncurvegeometrydescriptor/segmentcontrolpointcount
func (m_ MTL4AccelerationStructureMotionCurveGeometryDescriptor) SegmentControlPointCount() int {
	rv := objc.Send[int](m_.ID, objc.Sel("segmentControlPointCount"))
	return rv
}


// SetSegmentControlPointCount sets the value of the segmentControlPointCount property.
// Controls the number of control points per curve segment.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuremotioncurvegeometrydescriptor/segmentcontrolpointcount
func (m_ MTL4AccelerationStructureMotionCurveGeometryDescriptor) SetSegmentControlPointCount(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSegmentControlPointCount:"), value)
}

// Sets the stride, in bytes, between control points in the control point buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuremotioncurvegeometrydescriptor/controlpointstride
func (m_ MTL4AccelerationStructureMotionCurveGeometryDescriptor) ControlPointStride() int {
	rv := objc.Send[int](m_.ID, objc.Sel("controlPointStride"))
	return rv
}


// SetControlPointStride sets the value of the controlPointStride property.
// Sets the stride, in bytes, between control points in the control point buffer.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuremotioncurvegeometrydescriptor/controlpointstride
func (m_ MTL4AccelerationStructureMotionCurveGeometryDescriptor) SetControlPointStride(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setControlPointStride:"), value)
}

// Declares the number of curve segments.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuremotioncurvegeometrydescriptor/segmentcount
func (m_ MTL4AccelerationStructureMotionCurveGeometryDescriptor) SegmentCount() int {
	rv := objc.Send[int](m_.ID, objc.Sel("segmentCount"))
	return rv
}


// SetSegmentCount sets the value of the segmentCount property.
// Declares the number of curve segments.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuremotioncurvegeometrydescriptor/segmentcount
func (m_ MTL4AccelerationStructureMotionCurveGeometryDescriptor) SetSegmentCount(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSegmentCount:"), value)
}

// Controls the curve type.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuremotioncurvegeometrydescriptor/curvetype
func (m_ MTL4AccelerationStructureMotionCurveGeometryDescriptor) CurveType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("curveType"))
	return rv
}


// SetCurveType sets the value of the curveType property.
// Controls the curve type.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuremotioncurvegeometrydescriptor/curvetype
func (m_ MTL4AccelerationStructureMotionCurveGeometryDescriptor) SetCurveType(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCurveType:"), value)
}

// Sets the stride, in bytes, between radii in the radius buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuremotioncurvegeometrydescriptor/radiusstride
func (m_ MTL4AccelerationStructureMotionCurveGeometryDescriptor) RadiusStride() int {
	rv := objc.Send[int](m_.ID, objc.Sel("radiusStride"))
	return rv
}


// SetRadiusStride sets the value of the radiusStride property.
// Sets the stride, in bytes, between radii in the radius buffer.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuremotioncurvegeometrydescriptor/radiusstride
func (m_ MTL4AccelerationStructureMotionCurveGeometryDescriptor) SetRadiusStride(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRadiusStride:"), value)
}

// Configures the size of the indices the
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuremotioncurvegeometrydescriptor/indextype
func (m_ MTL4AccelerationStructureMotionCurveGeometryDescriptor) IndexType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("indexType"))
	return rv
}


// SetIndexType sets the value of the indexType property.
// Configures the size of the indices the

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuremotioncurvegeometrydescriptor/indextype
func (m_ MTL4AccelerationStructureMotionCurveGeometryDescriptor) SetIndexType(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIndexType:"), value)
}

// Declares the format of the control points in the buffers that the control point buffers reference.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuremotioncurvegeometrydescriptor/controlpointformat
func (m_ MTL4AccelerationStructureMotionCurveGeometryDescriptor) ControlPointFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("controlPointFormat"))
	return rv
}


// SetControlPointFormat sets the value of the controlPointFormat property.
// Declares the format of the control points in the buffers that the control point buffers reference.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuremotioncurvegeometrydescriptor/controlpointformat
func (m_ MTL4AccelerationStructureMotionCurveGeometryDescriptor) SetControlPointFormat(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setControlPointFormat:"), value)
}

// Assigns a reference to a buffer containing, in turn, references to curve radii buffers.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuremotioncurvegeometrydescriptor/radiusbuffers
func (m_ MTL4AccelerationStructureMotionCurveGeometryDescriptor) RadiusBuffers() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("radiusBuffers"))
	return rv
}


// SetRadiusBuffers sets the value of the radiusBuffers property.
// Assigns a reference to a buffer containing, in turn, references to curve radii buffers.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuremotioncurvegeometrydescriptor/radiusbuffers
func (m_ MTL4AccelerationStructureMotionCurveGeometryDescriptor) SetRadiusBuffers(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRadiusBuffers:"), value)
}

// Specifies the number of control points in the buffers the control point buffers reference.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuremotioncurvegeometrydescriptor/controlpointcount
func (m_ MTL4AccelerationStructureMotionCurveGeometryDescriptor) ControlPointCount() int {
	rv := objc.Send[int](m_.ID, objc.Sel("controlPointCount"))
	return rv
}


// SetControlPointCount sets the value of the controlPointCount property.
// Specifies the number of control points in the buffers the control point buffers reference.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuremotioncurvegeometrydescriptor/controlpointcount
func (m_ MTL4AccelerationStructureMotionCurveGeometryDescriptor) SetControlPointCount(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setControlPointCount:"), value)
}

// Sets the format of the radii in the radius buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuremotioncurvegeometrydescriptor/radiusformat
func (m_ MTL4AccelerationStructureMotionCurveGeometryDescriptor) RadiusFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("radiusFormat"))
	return rv
}


// SetRadiusFormat sets the value of the radiusFormat property.
// Sets the format of the radii in the radius buffer.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuremotioncurvegeometrydescriptor/radiusformat
func (m_ MTL4AccelerationStructureMotionCurveGeometryDescriptor) SetRadiusFormat(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRadiusFormat:"), value)
}

// Sets the curve basis function, determining how Metal interpolates the control points.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuremotioncurvegeometrydescriptor/curvebasis
func (m_ MTL4AccelerationStructureMotionCurveGeometryDescriptor) CurveBasis() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("curveBasis"))
	return rv
}


// SetCurveBasis sets the value of the curveBasis property.
// Sets the curve basis function, determining how Metal interpolates the control points.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuremotioncurvegeometrydescriptor/curvebasis
func (m_ MTL4AccelerationStructureMotionCurveGeometryDescriptor) SetCurveBasis(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCurveBasis:"), value)
}

// Assigns a reference to a buffer where each entry contains a reference to a buffer of control points.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuremotioncurvegeometrydescriptor/controlpointbuffers
func (m_ MTL4AccelerationStructureMotionCurveGeometryDescriptor) ControlPointBuffers() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("controlPointBuffers"))
	return rv
}


// SetControlPointBuffers sets the value of the controlPointBuffers property.
// Assigns a reference to a buffer where each entry contains a reference to a buffer of control points.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuremotioncurvegeometrydescriptor/controlpointbuffers
func (m_ MTL4AccelerationStructureMotionCurveGeometryDescriptor) SetControlPointBuffers(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setControlPointBuffers:"), value)
}

// Configures the type of curve end caps.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuremotioncurvegeometrydescriptor/curveendcaps
func (m_ MTL4AccelerationStructureMotionCurveGeometryDescriptor) CurveEndCaps() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("curveEndCaps"))
	return rv
}


// SetCurveEndCaps sets the value of the curveEndCaps property.
// Configures the type of curve end caps.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuremotioncurvegeometrydescriptor/curveendcaps
func (m_ MTL4AccelerationStructureMotionCurveGeometryDescriptor) SetCurveEndCaps(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCurveEndCaps:"), value)
}



