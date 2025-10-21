// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AccelerationStructureMotionCurveGeometryDescriptor] class.
var (
	AccelerationStructureMotionCurveGeometryDescriptorClass     _AccelerationStructureMotionCurveGeometryDescriptorClass
	AccelerationStructureMotionCurveGeometryDescriptorClassOnce sync.Once
)

func getAccelerationStructureMotionCurveGeometryDescriptorClass() _AccelerationStructureMotionCurveGeometryDescriptorClass {
	AccelerationStructureMotionCurveGeometryDescriptorClassOnce.Do(func() {
		AccelerationStructureMotionCurveGeometryDescriptorClass = _AccelerationStructureMotionCurveGeometryDescriptorClass{objc.GetClass("MTLAccelerationStructureMotionCurveGeometryDescriptor")}
	})
	return AccelerationStructureMotionCurveGeometryDescriptorClass
}

type _AccelerationStructureMotionCurveGeometryDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [AccelerationStructureMotionCurveGeometryDescriptor] class.
type IAccelerationStructureMotionCurveGeometryDescriptor interface {
	IAccelerationStructureGeometryDescriptor
}

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor
type AccelerationStructureMotionCurveGeometryDescriptor struct {
	AccelerationStructureGeometryDescriptor
}

// AccelerationStructureMotionCurveGeometryDescriptorFrom constructs a [AccelerationStructureMotionCurveGeometryDescriptor] from an unsafe.Pointer.
func AccelerationStructureMotionCurveGeometryDescriptorFrom(ptr unsafe.Pointer) AccelerationStructureMotionCurveGeometryDescriptor {
	return AccelerationStructureMotionCurveGeometryDescriptor{
		AccelerationStructureGeometryDescriptor: AccelerationStructureGeometryDescriptorFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AccelerationStructureMotionCurveGeometryDescriptorClass) Alloc() AccelerationStructureMotionCurveGeometryDescriptor {
	rv := objc.Send[AccelerationStructureMotionCurveGeometryDescriptor](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AccelerationStructureMotionCurveGeometryDescriptorClass) New() AccelerationStructureMotionCurveGeometryDescriptor {
	rv := objc.Send[AccelerationStructureMotionCurveGeometryDescriptor](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) Init() AccelerationStructureMotionCurveGeometryDescriptor {
	rv := objc.Send[AccelerationStructureMotionCurveGeometryDescriptor](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) Autorelease() AccelerationStructureMotionCurveGeometryDescriptor {
	rv := objc.Send[AccelerationStructureMotionCurveGeometryDescriptor](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAccelerationStructureMotionCurveGeometryDescriptor creates a new AccelerationStructureMotionCurveGeometryDescriptor instance.
func NewAccelerationStructureMotionCurveGeometryDescriptor() AccelerationStructureMotionCurveGeometryDescriptor {
	return getAccelerationStructureMotionCurveGeometryDescriptorClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotioncurvegeometrydescriptor/controlpointbuffers
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) ControlPointBuffers() MTLMotionKeyframeData {
	rv := objc.Send[MTLMotionKeyframeData](a_.ID, objc.Sel("controlPointBuffers"))
	return rv
}


// SetControlPointBuffers sets the value of the controlPointBuffers property.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotioncurvegeometrydescriptor/controlpointbuffers
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) SetControlPointBuffers(value IMTLMotionKeyframeData) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setControlPointBuffers:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotioncurvegeometrydescriptor/controlpointcount
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) ControlPointCount() int {
	rv := objc.Send[int](a_.ID, objc.Sel("controlPointCount"))
	return rv
}


// SetControlPointCount sets the value of the controlPointCount property.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotioncurvegeometrydescriptor/controlpointcount
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) SetControlPointCount(value int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setControlPointCount:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotioncurvegeometrydescriptor/controlpointformat
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) ControlPointFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("controlPointFormat"))
	return rv
}


// SetControlPointFormat sets the value of the controlPointFormat property.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotioncurvegeometrydescriptor/controlpointformat
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) SetControlPointFormat(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setControlPointFormat:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotioncurvegeometrydescriptor/controlpointstride
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) ControlPointStride() int {
	rv := objc.Send[int](a_.ID, objc.Sel("controlPointStride"))
	return rv
}


// SetControlPointStride sets the value of the controlPointStride property.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotioncurvegeometrydescriptor/controlpointstride
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) SetControlPointStride(value int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setControlPointStride:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotioncurvegeometrydescriptor/curvebasis
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) CurveBasis() CurveBasis {
	rv := objc.Send[CurveBasis](a_.ID, objc.Sel("curveBasis"))
	return rv
}


// SetCurveBasis sets the value of the curveBasis property.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotioncurvegeometrydescriptor/curvebasis
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) SetCurveBasis(value ICurveBasis) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCurveBasis:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotioncurvegeometrydescriptor/curveendcaps
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) CurveEndCaps() CurveEndCaps {
	rv := objc.Send[CurveEndCaps](a_.ID, objc.Sel("curveEndCaps"))
	return rv
}


// SetCurveEndCaps sets the value of the curveEndCaps property.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotioncurvegeometrydescriptor/curveendcaps
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) SetCurveEndCaps(value ICurveEndCaps) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCurveEndCaps:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotioncurvegeometrydescriptor/curvetype
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) CurveType() CurveType {
	rv := objc.Send[CurveType](a_.ID, objc.Sel("curveType"))
	return rv
}


// SetCurveType sets the value of the curveType property.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotioncurvegeometrydescriptor/curvetype
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) SetCurveType(value CurveType) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCurveType:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotioncurvegeometrydescriptor/indexbuffer
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) IndexBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("indexBuffer"))
	return rv
}


// SetIndexBuffer sets the value of the indexBuffer property.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotioncurvegeometrydescriptor/indexbuffer
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) SetIndexBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIndexBuffer:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotioncurvegeometrydescriptor/indexbufferoffset
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) IndexBufferOffset() int {
	rv := objc.Send[int](a_.ID, objc.Sel("indexBufferOffset"))
	return rv
}


// SetIndexBufferOffset sets the value of the indexBufferOffset property.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotioncurvegeometrydescriptor/indexbufferoffset
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) SetIndexBufferOffset(value int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIndexBufferOffset:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotioncurvegeometrydescriptor/indextype
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) IndexType() IndexType {
	rv := objc.Send[IndexType](a_.ID, objc.Sel("indexType"))
	return rv
}


// SetIndexType sets the value of the indexType property.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotioncurvegeometrydescriptor/indextype
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) SetIndexType(value IndexType) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIndexType:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotioncurvegeometrydescriptor/radiusbuffers
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) RadiusBuffers() MTLMotionKeyframeData {
	rv := objc.Send[MTLMotionKeyframeData](a_.ID, objc.Sel("radiusBuffers"))
	return rv
}


// SetRadiusBuffers sets the value of the radiusBuffers property.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotioncurvegeometrydescriptor/radiusbuffers
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) SetRadiusBuffers(value IMTLMotionKeyframeData) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRadiusBuffers:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotioncurvegeometrydescriptor/radiusformat
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) RadiusFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("radiusFormat"))
	return rv
}


// SetRadiusFormat sets the value of the radiusFormat property.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotioncurvegeometrydescriptor/radiusformat
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) SetRadiusFormat(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRadiusFormat:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotioncurvegeometrydescriptor/radiusstride
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) RadiusStride() int {
	rv := objc.Send[int](a_.ID, objc.Sel("radiusStride"))
	return rv
}


// SetRadiusStride sets the value of the radiusStride property.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotioncurvegeometrydescriptor/radiusstride
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) SetRadiusStride(value int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRadiusStride:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotioncurvegeometrydescriptor/segmentcontrolpointcount
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) SegmentControlPointCount() int {
	rv := objc.Send[int](a_.ID, objc.Sel("segmentControlPointCount"))
	return rv
}


// SetSegmentControlPointCount sets the value of the segmentControlPointCount property.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotioncurvegeometrydescriptor/segmentcontrolpointcount
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) SetSegmentControlPointCount(value int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSegmentControlPointCount:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotioncurvegeometrydescriptor/segmentcount
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) SegmentCount() int {
	rv := objc.Send[int](a_.ID, objc.Sel("segmentCount"))
	return rv
}


// SetSegmentCount sets the value of the segmentCount property.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotioncurvegeometrydescriptor/segmentcount
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) SetSegmentCount(value int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSegmentCount:"), value)
}



