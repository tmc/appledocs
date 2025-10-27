// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
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
	

	// properties:
	ControlPointBuffer() unsafe.Pointer
	SetControlPointBuffer(value unsafe.Pointer)
	ControlPointBufferOffset() uint
	SetControlPointBufferOffset(value uint)
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
	IndexBuffer() unsafe.Pointer
	SetIndexBuffer(value unsafe.Pointer)
	IndexBufferOffset() uint
	SetIndexBufferOffset(value uint)
	IndexType() IndexType
	SetIndexType(value IndexType)
	RadiusBuffer() unsafe.Pointer
	SetRadiusBuffer(value unsafe.Pointer)
	RadiusBufferOffset() uint
	SetRadiusBufferOffset(value uint)
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
func (ac _AccelerationStructureCurveGeometryDescriptorClass) Alloc() AccelerationStructureCurveGeometryDescriptor {
	rv := objc.Send[AccelerationStructureCurveGeometryDescriptor](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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







// [Full Topic]
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










// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/descriptor
func (ac _AccelerationStructureCurveGeometryDescriptorClass) Descriptor() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ac.class), objc.Sel("descriptor"))
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/controlPointBuffer
func (a_ AccelerationStructureCurveGeometryDescriptor) ControlPointBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("controlPointBuffer"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/controlPointBuffer
func (a_ AccelerationStructureCurveGeometryDescriptor) SetControlPointBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setControlPointBuffer:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/controlPointBufferOffset
func (a_ AccelerationStructureCurveGeometryDescriptor) ControlPointBufferOffset() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("controlPointBufferOffset"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/controlPointBufferOffset
func (a_ AccelerationStructureCurveGeometryDescriptor) SetControlPointBufferOffset(value uint) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setControlPointBufferOffset:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/controlPointCount
func (a_ AccelerationStructureCurveGeometryDescriptor) ControlPointCount() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("controlPointCount"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/controlPointCount
func (a_ AccelerationStructureCurveGeometryDescriptor) SetControlPointCount(value uint) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setControlPointCount:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/controlPointFormat
func (a_ AccelerationStructureCurveGeometryDescriptor) ControlPointFormat() AttributeFormat {
	rv := objc.Send[AttributeFormat](a_.ID, objc.Sel("controlPointFormat"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/controlPointFormat
func (a_ AccelerationStructureCurveGeometryDescriptor) SetControlPointFormat(value AttributeFormat) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setControlPointFormat:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/controlPointStride
func (a_ AccelerationStructureCurveGeometryDescriptor) ControlPointStride() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("controlPointStride"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/controlPointStride
func (a_ AccelerationStructureCurveGeometryDescriptor) SetControlPointStride(value uint) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setControlPointStride:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/curveBasis
func (a_ AccelerationStructureCurveGeometryDescriptor) CurveBasis() CurveBasis {
	rv := objc.Send[CurveBasis](a_.ID, objc.Sel("curveBasis"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/curveBasis
func (a_ AccelerationStructureCurveGeometryDescriptor) SetCurveBasis(value CurveBasis) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCurveBasis:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/curveEndCaps
func (a_ AccelerationStructureCurveGeometryDescriptor) CurveEndCaps() CurveEndCaps {
	rv := objc.Send[CurveEndCaps](a_.ID, objc.Sel("curveEndCaps"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/curveEndCaps
func (a_ AccelerationStructureCurveGeometryDescriptor) SetCurveEndCaps(value CurveEndCaps) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCurveEndCaps:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/curveType
func (a_ AccelerationStructureCurveGeometryDescriptor) CurveType() CurveType {
	rv := objc.Send[CurveType](a_.ID, objc.Sel("curveType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/curveType
func (a_ AccelerationStructureCurveGeometryDescriptor) SetCurveType(value CurveType) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCurveType:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/indexBuffer
func (a_ AccelerationStructureCurveGeometryDescriptor) IndexBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("indexBuffer"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/indexBuffer
func (a_ AccelerationStructureCurveGeometryDescriptor) SetIndexBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIndexBuffer:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/indexBufferOffset
func (a_ AccelerationStructureCurveGeometryDescriptor) IndexBufferOffset() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("indexBufferOffset"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/indexBufferOffset
func (a_ AccelerationStructureCurveGeometryDescriptor) SetIndexBufferOffset(value uint) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIndexBufferOffset:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/indexType
func (a_ AccelerationStructureCurveGeometryDescriptor) IndexType() IndexType {
	rv := objc.Send[IndexType](a_.ID, objc.Sel("indexType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/indexType
func (a_ AccelerationStructureCurveGeometryDescriptor) SetIndexType(value IndexType) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIndexType:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/radiusBuffer
func (a_ AccelerationStructureCurveGeometryDescriptor) RadiusBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("radiusBuffer"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/radiusBuffer
func (a_ AccelerationStructureCurveGeometryDescriptor) SetRadiusBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRadiusBuffer:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/radiusBufferOffset
func (a_ AccelerationStructureCurveGeometryDescriptor) RadiusBufferOffset() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("radiusBufferOffset"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/radiusBufferOffset
func (a_ AccelerationStructureCurveGeometryDescriptor) SetRadiusBufferOffset(value uint) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRadiusBufferOffset:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/radiusFormat
func (a_ AccelerationStructureCurveGeometryDescriptor) RadiusFormat() AttributeFormat {
	rv := objc.Send[AttributeFormat](a_.ID, objc.Sel("radiusFormat"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/radiusFormat
func (a_ AccelerationStructureCurveGeometryDescriptor) SetRadiusFormat(value AttributeFormat) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRadiusFormat:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/radiusStride
func (a_ AccelerationStructureCurveGeometryDescriptor) RadiusStride() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("radiusStride"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/radiusStride
func (a_ AccelerationStructureCurveGeometryDescriptor) SetRadiusStride(value uint) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRadiusStride:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/segmentControlPointCount
func (a_ AccelerationStructureCurveGeometryDescriptor) SegmentControlPointCount() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("segmentControlPointCount"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/segmentControlPointCount
func (a_ AccelerationStructureCurveGeometryDescriptor) SetSegmentControlPointCount(value uint) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSegmentControlPointCount:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/segmentCount
func (a_ AccelerationStructureCurveGeometryDescriptor) SegmentCount() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("segmentCount"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/segmentCount
func (a_ AccelerationStructureCurveGeometryDescriptor) SetSegmentCount(value uint) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSegmentCount:"), value)
}








