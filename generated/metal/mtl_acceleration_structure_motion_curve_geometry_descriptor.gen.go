// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLAccelerationStructureMotionCurveGeometryDescriptor */


/* debug [class_header]: Header for MTLAccelerationStructureMotionCurveGeometryDescriptor */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AccelerationStructureMotionCurveGeometryDescriptor */
// An interface definition for the [AccelerationStructureMotionCurveGeometryDescriptor] class.
type IAccelerationStructureMotionCurveGeometryDescriptor interface {
	IAccelerationStructureGeometryDescriptor
	
/* debug [class_interface_properties]: Properties for AccelerationStructureMotionCurveGeometryDescriptor */
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
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AccelerationStructureMotionCurveGeometryDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AccelerationStructureMotionCurveGeometryDescriptor */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AccelerationStructureMotionCurveGeometryDescriptor */


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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AccelerationStructureMotionCurveGeometryDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AccelerationStructureMotionCurveGeometryDescriptor */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor/descriptor
func (ac _AccelerationStructureMotionCurveGeometryDescriptorClass) Descriptor() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ac.class), objc.Sel("descriptor"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=Descriptor) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AccelerationStructureMotionCurveGeometryDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AccelerationStructureMotionCurveGeometryDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AccelerationStructureMotionCurveGeometryDescriptor */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor/controlPointBuffers
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) ControlPointBuffers() []MotionKeyframeData {
	rv := objc.Send[[]MotionKeyframeData](a_.ID, objc.Sel("controlPointBuffers"))
	return rv
}/* debug [instance_properties/getter]: controlPointBuffers */


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
}/* debug [instance_properties/setter]: controlPointBuffers */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor/controlPointCount
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) ControlPointCount() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("controlPointCount"))
	return rv
}/* debug [instance_properties/getter]: controlPointCount */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor/controlPointCount
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) SetControlPointCount(value uint) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setControlPointCount:"), value)
}/* debug [instance_properties/setter]: controlPointCount */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor/controlPointFormat
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) ControlPointFormat() AttributeFormat {
	rv := objc.Send[AttributeFormat](a_.ID, objc.Sel("controlPointFormat"))
	return rv
}/* debug [instance_properties/getter]: controlPointFormat */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor/controlPointFormat
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) SetControlPointFormat(value AttributeFormat) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setControlPointFormat:"), value)
}/* debug [instance_properties/setter]: controlPointFormat */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor/controlPointStride
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) ControlPointStride() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("controlPointStride"))
	return rv
}/* debug [instance_properties/getter]: controlPointStride */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor/controlPointStride
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) SetControlPointStride(value uint) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setControlPointStride:"), value)
}/* debug [instance_properties/setter]: controlPointStride */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor/curveBasis
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) CurveBasis() CurveBasis {
	rv := objc.Send[CurveBasis](a_.ID, objc.Sel("curveBasis"))
	return rv
}/* debug [instance_properties/getter]: curveBasis */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor/curveBasis
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) SetCurveBasis(value CurveBasis) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCurveBasis:"), value)
}/* debug [instance_properties/setter]: curveBasis */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor/curveEndCaps
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) CurveEndCaps() CurveEndCaps {
	rv := objc.Send[CurveEndCaps](a_.ID, objc.Sel("curveEndCaps"))
	return rv
}/* debug [instance_properties/getter]: curveEndCaps */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor/curveEndCaps
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) SetCurveEndCaps(value CurveEndCaps) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCurveEndCaps:"), value)
}/* debug [instance_properties/setter]: curveEndCaps */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor/curveType
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) CurveType() CurveType {
	rv := objc.Send[CurveType](a_.ID, objc.Sel("curveType"))
	return rv
}/* debug [instance_properties/getter]: curveType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor/curveType
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) SetCurveType(value CurveType) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCurveType:"), value)
}/* debug [instance_properties/setter]: curveType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor/indexBuffer
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) IndexBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("indexBuffer"))
	return rv
}/* debug [instance_properties/getter]: indexBuffer */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor/indexBuffer
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) SetIndexBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIndexBuffer:"), value)
}/* debug [instance_properties/setter]: indexBuffer */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor/indexBufferOffset
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) IndexBufferOffset() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("indexBufferOffset"))
	return rv
}/* debug [instance_properties/getter]: indexBufferOffset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor/indexBufferOffset
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) SetIndexBufferOffset(value uint) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIndexBufferOffset:"), value)
}/* debug [instance_properties/setter]: indexBufferOffset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor/indexType
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) IndexType() IndexType {
	rv := objc.Send[IndexType](a_.ID, objc.Sel("indexType"))
	return rv
}/* debug [instance_properties/getter]: indexType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor/indexType
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) SetIndexType(value IndexType) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIndexType:"), value)
}/* debug [instance_properties/setter]: indexType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor/radiusBuffers
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) RadiusBuffers() []MotionKeyframeData {
	rv := objc.Send[[]MotionKeyframeData](a_.ID, objc.Sel("radiusBuffers"))
	return rv
}/* debug [instance_properties/getter]: radiusBuffers */


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
}/* debug [instance_properties/setter]: radiusBuffers */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor/radiusFormat
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) RadiusFormat() AttributeFormat {
	rv := objc.Send[AttributeFormat](a_.ID, objc.Sel("radiusFormat"))
	return rv
}/* debug [instance_properties/getter]: radiusFormat */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor/radiusFormat
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) SetRadiusFormat(value AttributeFormat) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRadiusFormat:"), value)
}/* debug [instance_properties/setter]: radiusFormat */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor/radiusStride
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) RadiusStride() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("radiusStride"))
	return rv
}/* debug [instance_properties/getter]: radiusStride */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor/radiusStride
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) SetRadiusStride(value uint) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRadiusStride:"), value)
}/* debug [instance_properties/setter]: radiusStride */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor/segmentControlPointCount
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) SegmentControlPointCount() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("segmentControlPointCount"))
	return rv
}/* debug [instance_properties/getter]: segmentControlPointCount */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor/segmentControlPointCount
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) SetSegmentControlPointCount(value uint) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSegmentControlPointCount:"), value)
}/* debug [instance_properties/setter]: segmentControlPointCount */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor/segmentCount
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) SegmentCount() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("segmentCount"))
	return rv
}/* debug [instance_properties/getter]: segmentCount */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionCurveGeometryDescriptor/segmentCount
func (a_ AccelerationStructureMotionCurveGeometryDescriptor) SetSegmentCount(value uint) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSegmentCount:"), value)
}/* debug [instance_properties/setter]: segmentCount */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLAccelerationStructureMotionCurveGeometryDescriptor */



