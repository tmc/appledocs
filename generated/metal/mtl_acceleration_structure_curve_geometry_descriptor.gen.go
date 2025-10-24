// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLAccelerationStructureCurveGeometryDescriptor */


/* debug [class_header]: Header for MTLAccelerationStructureCurveGeometryDescriptor */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AccelerationStructureCurveGeometryDescriptor */
// An interface definition for the [AccelerationStructureCurveGeometryDescriptor] class.
type IAccelerationStructureCurveGeometryDescriptor interface {
	IAccelerationStructureGeometryDescriptor
	
/* debug [class_interface_properties]: Properties for AccelerationStructureCurveGeometryDescriptor */
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
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AccelerationStructureCurveGeometryDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AccelerationStructureCurveGeometryDescriptor */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AccelerationStructureCurveGeometryDescriptor */


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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AccelerationStructureCurveGeometryDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AccelerationStructureCurveGeometryDescriptor */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/descriptor
func (ac _AccelerationStructureCurveGeometryDescriptorClass) Descriptor() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ac.class), objc.Sel("descriptor"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=Descriptor) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AccelerationStructureCurveGeometryDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AccelerationStructureCurveGeometryDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AccelerationStructureCurveGeometryDescriptor */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/controlPointBuffer
func (a_ AccelerationStructureCurveGeometryDescriptor) ControlPointBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("controlPointBuffer"))
	return rv
}/* debug [instance_properties/getter]: controlPointBuffer */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/controlPointBuffer
func (a_ AccelerationStructureCurveGeometryDescriptor) SetControlPointBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setControlPointBuffer:"), value)
}/* debug [instance_properties/setter]: controlPointBuffer */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/controlPointBufferOffset
func (a_ AccelerationStructureCurveGeometryDescriptor) ControlPointBufferOffset() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("controlPointBufferOffset"))
	return rv
}/* debug [instance_properties/getter]: controlPointBufferOffset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/controlPointBufferOffset
func (a_ AccelerationStructureCurveGeometryDescriptor) SetControlPointBufferOffset(value uint) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setControlPointBufferOffset:"), value)
}/* debug [instance_properties/setter]: controlPointBufferOffset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/controlPointCount
func (a_ AccelerationStructureCurveGeometryDescriptor) ControlPointCount() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("controlPointCount"))
	return rv
}/* debug [instance_properties/getter]: controlPointCount */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/controlPointCount
func (a_ AccelerationStructureCurveGeometryDescriptor) SetControlPointCount(value uint) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setControlPointCount:"), value)
}/* debug [instance_properties/setter]: controlPointCount */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/controlPointFormat
func (a_ AccelerationStructureCurveGeometryDescriptor) ControlPointFormat() AttributeFormat {
	rv := objc.Send[AttributeFormat](a_.ID, objc.Sel("controlPointFormat"))
	return rv
}/* debug [instance_properties/getter]: controlPointFormat */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/controlPointFormat
func (a_ AccelerationStructureCurveGeometryDescriptor) SetControlPointFormat(value AttributeFormat) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setControlPointFormat:"), value)
}/* debug [instance_properties/setter]: controlPointFormat */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/controlPointStride
func (a_ AccelerationStructureCurveGeometryDescriptor) ControlPointStride() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("controlPointStride"))
	return rv
}/* debug [instance_properties/getter]: controlPointStride */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/controlPointStride
func (a_ AccelerationStructureCurveGeometryDescriptor) SetControlPointStride(value uint) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setControlPointStride:"), value)
}/* debug [instance_properties/setter]: controlPointStride */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/curveBasis
func (a_ AccelerationStructureCurveGeometryDescriptor) CurveBasis() CurveBasis {
	rv := objc.Send[CurveBasis](a_.ID, objc.Sel("curveBasis"))
	return rv
}/* debug [instance_properties/getter]: curveBasis */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/curveBasis
func (a_ AccelerationStructureCurveGeometryDescriptor) SetCurveBasis(value CurveBasis) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCurveBasis:"), value)
}/* debug [instance_properties/setter]: curveBasis */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/curveEndCaps
func (a_ AccelerationStructureCurveGeometryDescriptor) CurveEndCaps() CurveEndCaps {
	rv := objc.Send[CurveEndCaps](a_.ID, objc.Sel("curveEndCaps"))
	return rv
}/* debug [instance_properties/getter]: curveEndCaps */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/curveEndCaps
func (a_ AccelerationStructureCurveGeometryDescriptor) SetCurveEndCaps(value CurveEndCaps) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCurveEndCaps:"), value)
}/* debug [instance_properties/setter]: curveEndCaps */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/curveType
func (a_ AccelerationStructureCurveGeometryDescriptor) CurveType() CurveType {
	rv := objc.Send[CurveType](a_.ID, objc.Sel("curveType"))
	return rv
}/* debug [instance_properties/getter]: curveType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/curveType
func (a_ AccelerationStructureCurveGeometryDescriptor) SetCurveType(value CurveType) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCurveType:"), value)
}/* debug [instance_properties/setter]: curveType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/indexBuffer
func (a_ AccelerationStructureCurveGeometryDescriptor) IndexBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("indexBuffer"))
	return rv
}/* debug [instance_properties/getter]: indexBuffer */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/indexBuffer
func (a_ AccelerationStructureCurveGeometryDescriptor) SetIndexBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIndexBuffer:"), value)
}/* debug [instance_properties/setter]: indexBuffer */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/indexBufferOffset
func (a_ AccelerationStructureCurveGeometryDescriptor) IndexBufferOffset() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("indexBufferOffset"))
	return rv
}/* debug [instance_properties/getter]: indexBufferOffset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/indexBufferOffset
func (a_ AccelerationStructureCurveGeometryDescriptor) SetIndexBufferOffset(value uint) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIndexBufferOffset:"), value)
}/* debug [instance_properties/setter]: indexBufferOffset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/indexType
func (a_ AccelerationStructureCurveGeometryDescriptor) IndexType() IndexType {
	rv := objc.Send[IndexType](a_.ID, objc.Sel("indexType"))
	return rv
}/* debug [instance_properties/getter]: indexType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/indexType
func (a_ AccelerationStructureCurveGeometryDescriptor) SetIndexType(value IndexType) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIndexType:"), value)
}/* debug [instance_properties/setter]: indexType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/radiusBuffer
func (a_ AccelerationStructureCurveGeometryDescriptor) RadiusBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("radiusBuffer"))
	return rv
}/* debug [instance_properties/getter]: radiusBuffer */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/radiusBuffer
func (a_ AccelerationStructureCurveGeometryDescriptor) SetRadiusBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRadiusBuffer:"), value)
}/* debug [instance_properties/setter]: radiusBuffer */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/radiusBufferOffset
func (a_ AccelerationStructureCurveGeometryDescriptor) RadiusBufferOffset() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("radiusBufferOffset"))
	return rv
}/* debug [instance_properties/getter]: radiusBufferOffset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/radiusBufferOffset
func (a_ AccelerationStructureCurveGeometryDescriptor) SetRadiusBufferOffset(value uint) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRadiusBufferOffset:"), value)
}/* debug [instance_properties/setter]: radiusBufferOffset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/radiusFormat
func (a_ AccelerationStructureCurveGeometryDescriptor) RadiusFormat() AttributeFormat {
	rv := objc.Send[AttributeFormat](a_.ID, objc.Sel("radiusFormat"))
	return rv
}/* debug [instance_properties/getter]: radiusFormat */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/radiusFormat
func (a_ AccelerationStructureCurveGeometryDescriptor) SetRadiusFormat(value AttributeFormat) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRadiusFormat:"), value)
}/* debug [instance_properties/setter]: radiusFormat */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/radiusStride
func (a_ AccelerationStructureCurveGeometryDescriptor) RadiusStride() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("radiusStride"))
	return rv
}/* debug [instance_properties/getter]: radiusStride */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/radiusStride
func (a_ AccelerationStructureCurveGeometryDescriptor) SetRadiusStride(value uint) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRadiusStride:"), value)
}/* debug [instance_properties/setter]: radiusStride */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/segmentControlPointCount
func (a_ AccelerationStructureCurveGeometryDescriptor) SegmentControlPointCount() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("segmentControlPointCount"))
	return rv
}/* debug [instance_properties/getter]: segmentControlPointCount */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/segmentControlPointCount
func (a_ AccelerationStructureCurveGeometryDescriptor) SetSegmentControlPointCount(value uint) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSegmentControlPointCount:"), value)
}/* debug [instance_properties/setter]: segmentControlPointCount */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/segmentCount
func (a_ AccelerationStructureCurveGeometryDescriptor) SegmentCount() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("segmentCount"))
	return rv
}/* debug [instance_properties/getter]: segmentCount */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureCurveGeometryDescriptor/segmentCount
func (a_ AccelerationStructureCurveGeometryDescriptor) SetSegmentCount(value uint) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSegmentCount:"), value)
}/* debug [instance_properties/setter]: segmentCount */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLAccelerationStructureCurveGeometryDescriptor */



