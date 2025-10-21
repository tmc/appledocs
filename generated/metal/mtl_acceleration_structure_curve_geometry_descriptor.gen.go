// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AccelerationStructureCurveGeometryDescriptor] class.
var (
	AccelerationStructureCurveGeometryDescriptorClass     _AccelerationStructureCurveGeometryDescriptorClass
	AccelerationStructureCurveGeometryDescriptorClassOnce sync.Once
)

func getAccelerationStructureCurveGeometryDescriptorClass() _AccelerationStructureCurveGeometryDescriptorClass {
	AccelerationStructureCurveGeometryDescriptorClassOnce.Do(func() {
		AccelerationStructureCurveGeometryDescriptorClass = _AccelerationStructureCurveGeometryDescriptorClass{objc.GetClass("MTLAccelerationStructureCurveGeometryDescriptor")}
	})
	return AccelerationStructureCurveGeometryDescriptorClass
}

type _AccelerationStructureCurveGeometryDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [AccelerationStructureCurveGeometryDescriptor] class.
type IAccelerationStructureCurveGeometryDescriptor interface {
	IAccelerationStructureGeometryDescriptor
}

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor
type AccelerationStructureCurveGeometryDescriptor struct {
	AccelerationStructureGeometryDescriptor
}

// AccelerationStructureCurveGeometryDescriptorFrom constructs a [AccelerationStructureCurveGeometryDescriptor] from an unsafe.Pointer.
func AccelerationStructureCurveGeometryDescriptorFrom(ptr unsafe.Pointer) AccelerationStructureCurveGeometryDescriptor {
	return AccelerationStructureCurveGeometryDescriptor{
		AccelerationStructureGeometryDescriptor: AccelerationStructureGeometryDescriptorFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AccelerationStructureCurveGeometryDescriptorClass) Alloc() AccelerationStructureCurveGeometryDescriptor {
	rv := objc.Send[AccelerationStructureCurveGeometryDescriptor](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AccelerationStructureCurveGeometryDescriptorClass) New() AccelerationStructureCurveGeometryDescriptor {
	rv := objc.Send[AccelerationStructureCurveGeometryDescriptor](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AccelerationStructureCurveGeometryDescriptor) Init() AccelerationStructureCurveGeometryDescriptor {
	rv := objc.Send[AccelerationStructureCurveGeometryDescriptor](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AccelerationStructureCurveGeometryDescriptor) Autorelease() AccelerationStructureCurveGeometryDescriptor {
	rv := objc.Send[AccelerationStructureCurveGeometryDescriptor](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAccelerationStructureCurveGeometryDescriptor creates a new AccelerationStructureCurveGeometryDescriptor instance.
func NewAccelerationStructureCurveGeometryDescriptor() AccelerationStructureCurveGeometryDescriptor {
	return getAccelerationStructureCurveGeometryDescriptorClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructurecurvegeometrydescriptor/controlpointbuffer
func (a_ AccelerationStructureCurveGeometryDescriptor) ControlPointBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("controlPointBuffer"))
	return rv
}


// SetControlPointBuffer sets the value of the controlPointBuffer property.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructurecurvegeometrydescriptor/controlpointbuffer
func (a_ AccelerationStructureCurveGeometryDescriptor) SetControlPointBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setControlPointBuffer:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructurecurvegeometrydescriptor/controlpointbufferoffset
func (a_ AccelerationStructureCurveGeometryDescriptor) ControlPointBufferOffset() int {
	rv := objc.Send[int](a_.ID, objc.Sel("controlPointBufferOffset"))
	return rv
}


// SetControlPointBufferOffset sets the value of the controlPointBufferOffset property.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructurecurvegeometrydescriptor/controlpointbufferoffset
func (a_ AccelerationStructureCurveGeometryDescriptor) SetControlPointBufferOffset(value int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setControlPointBufferOffset:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructurecurvegeometrydescriptor/controlpointcount
func (a_ AccelerationStructureCurveGeometryDescriptor) ControlPointCount() int {
	rv := objc.Send[int](a_.ID, objc.Sel("controlPointCount"))
	return rv
}


// SetControlPointCount sets the value of the controlPointCount property.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructurecurvegeometrydescriptor/controlpointcount
func (a_ AccelerationStructureCurveGeometryDescriptor) SetControlPointCount(value int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setControlPointCount:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructurecurvegeometrydescriptor/controlpointformat
func (a_ AccelerationStructureCurveGeometryDescriptor) ControlPointFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("controlPointFormat"))
	return rv
}


// SetControlPointFormat sets the value of the controlPointFormat property.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructurecurvegeometrydescriptor/controlpointformat
func (a_ AccelerationStructureCurveGeometryDescriptor) SetControlPointFormat(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setControlPointFormat:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructurecurvegeometrydescriptor/controlpointstride
func (a_ AccelerationStructureCurveGeometryDescriptor) ControlPointStride() int {
	rv := objc.Send[int](a_.ID, objc.Sel("controlPointStride"))
	return rv
}


// SetControlPointStride sets the value of the controlPointStride property.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructurecurvegeometrydescriptor/controlpointstride
func (a_ AccelerationStructureCurveGeometryDescriptor) SetControlPointStride(value int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setControlPointStride:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructurecurvegeometrydescriptor/curvebasis
func (a_ AccelerationStructureCurveGeometryDescriptor) CurveBasis() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("curveBasis"))
	return rv
}


// SetCurveBasis sets the value of the curveBasis property.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructurecurvegeometrydescriptor/curvebasis
func (a_ AccelerationStructureCurveGeometryDescriptor) SetCurveBasis(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCurveBasis:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructurecurvegeometrydescriptor/curveendcaps
func (a_ AccelerationStructureCurveGeometryDescriptor) CurveEndCaps() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("curveEndCaps"))
	return rv
}


// SetCurveEndCaps sets the value of the curveEndCaps property.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructurecurvegeometrydescriptor/curveendcaps
func (a_ AccelerationStructureCurveGeometryDescriptor) SetCurveEndCaps(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCurveEndCaps:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructurecurvegeometrydescriptor/curvetype
func (a_ AccelerationStructureCurveGeometryDescriptor) CurveType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("curveType"))
	return rv
}


// SetCurveType sets the value of the curveType property.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructurecurvegeometrydescriptor/curvetype
func (a_ AccelerationStructureCurveGeometryDescriptor) SetCurveType(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCurveType:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructurecurvegeometrydescriptor/indexbuffer
func (a_ AccelerationStructureCurveGeometryDescriptor) IndexBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("indexBuffer"))
	return rv
}


// SetIndexBuffer sets the value of the indexBuffer property.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructurecurvegeometrydescriptor/indexbuffer
func (a_ AccelerationStructureCurveGeometryDescriptor) SetIndexBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIndexBuffer:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructurecurvegeometrydescriptor/indexbufferoffset
func (a_ AccelerationStructureCurveGeometryDescriptor) IndexBufferOffset() int {
	rv := objc.Send[int](a_.ID, objc.Sel("indexBufferOffset"))
	return rv
}


// SetIndexBufferOffset sets the value of the indexBufferOffset property.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructurecurvegeometrydescriptor/indexbufferoffset
func (a_ AccelerationStructureCurveGeometryDescriptor) SetIndexBufferOffset(value int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIndexBufferOffset:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructurecurvegeometrydescriptor/indextype
func (a_ AccelerationStructureCurveGeometryDescriptor) IndexType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("indexType"))
	return rv
}


// SetIndexType sets the value of the indexType property.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructurecurvegeometrydescriptor/indextype
func (a_ AccelerationStructureCurveGeometryDescriptor) SetIndexType(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIndexType:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructurecurvegeometrydescriptor/radiusbuffer
func (a_ AccelerationStructureCurveGeometryDescriptor) RadiusBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("radiusBuffer"))
	return rv
}


// SetRadiusBuffer sets the value of the radiusBuffer property.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructurecurvegeometrydescriptor/radiusbuffer
func (a_ AccelerationStructureCurveGeometryDescriptor) SetRadiusBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRadiusBuffer:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructurecurvegeometrydescriptor/radiusbufferoffset
func (a_ AccelerationStructureCurveGeometryDescriptor) RadiusBufferOffset() int {
	rv := objc.Send[int](a_.ID, objc.Sel("radiusBufferOffset"))
	return rv
}


// SetRadiusBufferOffset sets the value of the radiusBufferOffset property.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructurecurvegeometrydescriptor/radiusbufferoffset
func (a_ AccelerationStructureCurveGeometryDescriptor) SetRadiusBufferOffset(value int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRadiusBufferOffset:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructurecurvegeometrydescriptor/radiusformat
func (a_ AccelerationStructureCurveGeometryDescriptor) RadiusFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("radiusFormat"))
	return rv
}


// SetRadiusFormat sets the value of the radiusFormat property.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructurecurvegeometrydescriptor/radiusformat
func (a_ AccelerationStructureCurveGeometryDescriptor) SetRadiusFormat(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRadiusFormat:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructurecurvegeometrydescriptor/radiusstride
func (a_ AccelerationStructureCurveGeometryDescriptor) RadiusStride() int {
	rv := objc.Send[int](a_.ID, objc.Sel("radiusStride"))
	return rv
}


// SetRadiusStride sets the value of the radiusStride property.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructurecurvegeometrydescriptor/radiusstride
func (a_ AccelerationStructureCurveGeometryDescriptor) SetRadiusStride(value int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRadiusStride:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructurecurvegeometrydescriptor/segmentcontrolpointcount
func (a_ AccelerationStructureCurveGeometryDescriptor) SegmentControlPointCount() int {
	rv := objc.Send[int](a_.ID, objc.Sel("segmentControlPointCount"))
	return rv
}


// SetSegmentControlPointCount sets the value of the segmentControlPointCount property.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructurecurvegeometrydescriptor/segmentcontrolpointcount
func (a_ AccelerationStructureCurveGeometryDescriptor) SetSegmentControlPointCount(value int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSegmentControlPointCount:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructurecurvegeometrydescriptor/segmentcount
func (a_ AccelerationStructureCurveGeometryDescriptor) SegmentCount() int {
	rv := objc.Send[int](a_.ID, objc.Sel("segmentCount"))
	return rv
}


// SetSegmentCount sets the value of the segmentCount property.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructurecurvegeometrydescriptor/segmentcount
func (a_ AccelerationStructureCurveGeometryDescriptor) SetSegmentCount(value int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSegmentCount:"), value)
}



