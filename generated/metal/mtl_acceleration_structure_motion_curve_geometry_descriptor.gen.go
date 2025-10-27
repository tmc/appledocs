// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
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
	

	// properties:
	ControlPointBuffers() []MotionKeyframeData
	SetControlPointBuffers(value []MotionKeyframeData)
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
	RadiusBuffers() []MotionKeyframeData
	SetRadiusBuffers(value []MotionKeyframeData)
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
func (ac _AccelerationStructureMotionCurveGeometryDescriptorClass) Alloc() AccelerationStructureMotionCurveGeometryDescriptor {
	rv := objc.Send[AccelerationStructureMotionCurveGeometryDescriptor](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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







// [Full Topic]
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










// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor/descriptor
func (ac _AccelerationStructureMotionCurveGeometryDescriptorClass) Descriptor() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ac.class), objc.Sel("descriptor"))
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor/controlPointBuffers
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) ControlPointBuffers() []MotionKeyframeData {
	rv := objc.Send[[]MotionKeyframeData](a_.ID, objc.Sel("controlPointBuffers"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor/controlPointBuffers
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) SetControlPointBuffers(value []MotionKeyframeData) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](a_.ID, objc.Sel("setControlPointBuffers:"), nsArray)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor/controlPointCount
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) ControlPointCount() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("controlPointCount"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor/controlPointCount
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) SetControlPointCount(value uint) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setControlPointCount:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor/controlPointFormat
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) ControlPointFormat() AttributeFormat {
	rv := objc.Send[AttributeFormat](a_.ID, objc.Sel("controlPointFormat"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor/controlPointFormat
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) SetControlPointFormat(value AttributeFormat) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setControlPointFormat:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor/controlPointStride
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) ControlPointStride() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("controlPointStride"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor/controlPointStride
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) SetControlPointStride(value uint) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setControlPointStride:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor/curveBasis
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) CurveBasis() CurveBasis {
	rv := objc.Send[CurveBasis](a_.ID, objc.Sel("curveBasis"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor/curveBasis
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) SetCurveBasis(value CurveBasis) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCurveBasis:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor/curveEndCaps
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) CurveEndCaps() CurveEndCaps {
	rv := objc.Send[CurveEndCaps](a_.ID, objc.Sel("curveEndCaps"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor/curveEndCaps
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) SetCurveEndCaps(value CurveEndCaps) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCurveEndCaps:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor/curveType
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) CurveType() CurveType {
	rv := objc.Send[CurveType](a_.ID, objc.Sel("curveType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor/curveType
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) SetCurveType(value CurveType) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCurveType:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor/indexBuffer
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) IndexBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("indexBuffer"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor/indexBuffer
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) SetIndexBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIndexBuffer:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor/indexBufferOffset
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) IndexBufferOffset() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("indexBufferOffset"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor/indexBufferOffset
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) SetIndexBufferOffset(value uint) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIndexBufferOffset:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor/indexType
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) IndexType() IndexType {
	rv := objc.Send[IndexType](a_.ID, objc.Sel("indexType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor/indexType
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) SetIndexType(value IndexType) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIndexType:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor/radiusBuffers
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) RadiusBuffers() []MotionKeyframeData {
	rv := objc.Send[[]MotionKeyframeData](a_.ID, objc.Sel("radiusBuffers"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor/radiusBuffers
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) SetRadiusBuffers(value []MotionKeyframeData) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](a_.ID, objc.Sel("setRadiusBuffers:"), nsArray)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor/radiusFormat
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) RadiusFormat() AttributeFormat {
	rv := objc.Send[AttributeFormat](a_.ID, objc.Sel("radiusFormat"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor/radiusFormat
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) SetRadiusFormat(value AttributeFormat) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRadiusFormat:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor/radiusStride
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) RadiusStride() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("radiusStride"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor/radiusStride
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) SetRadiusStride(value uint) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRadiusStride:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor/segmentControlPointCount
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) SegmentControlPointCount() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("segmentControlPointCount"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor/segmentControlPointCount
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) SetSegmentControlPointCount(value uint) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSegmentControlPointCount:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor/segmentCount
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) SegmentCount() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("segmentCount"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor/segmentCount
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) SetSegmentCount(value uint) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSegmentCount:"), value)
}








