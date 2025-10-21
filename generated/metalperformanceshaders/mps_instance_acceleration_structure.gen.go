// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [InstanceAccelerationStructure] class.
var (
	InstanceAccelerationStructureClass     _InstanceAccelerationStructureClass
	InstanceAccelerationStructureClassOnce sync.Once
)

func getInstanceAccelerationStructureClass() _InstanceAccelerationStructureClass {
	InstanceAccelerationStructureClassOnce.Do(func() {
		InstanceAccelerationStructureClass = _InstanceAccelerationStructureClass{objc.GetClass("MPSInstanceAccelerationStructure")}
	})
	return InstanceAccelerationStructureClass
}

type _InstanceAccelerationStructureClass struct {
	class objc.Class
}

// An interface definition for the [InstanceAccelerationStructure] class.
type IInstanceAccelerationStructure interface {
	IAccelerationStructure
}

// An acceleration structure built over instances of other acceleration structures.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSInstanceAccelerationStructure
type InstanceAccelerationStructure struct {
	AccelerationStructure
}

// InstanceAccelerationStructureFrom constructs a [InstanceAccelerationStructure] from an unsafe.Pointer.
//
// An acceleration structure built over instances of other acceleration structures.
func InstanceAccelerationStructureFrom(ptr unsafe.Pointer) InstanceAccelerationStructure {
	return InstanceAccelerationStructure{
		AccelerationStructure: AccelerationStructureFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _InstanceAccelerationStructureClass) Alloc() InstanceAccelerationStructure {
	rv := objc.Send[InstanceAccelerationStructure](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _InstanceAccelerationStructureClass) New() InstanceAccelerationStructure {
	rv := objc.Send[InstanceAccelerationStructure](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ InstanceAccelerationStructure) Init() InstanceAccelerationStructure {
	rv := objc.Send[InstanceAccelerationStructure](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ InstanceAccelerationStructure) Autorelease() InstanceAccelerationStructure {
	rv := objc.Send[InstanceAccelerationStructure](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewInstanceAccelerationStructure creates a new InstanceAccelerationStructure instance.
func NewInstanceAccelerationStructure() InstanceAccelerationStructure {
	return getInstanceAccelerationStructureClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSInstanceAccelerationStructure/transformBuffer
func (i_ InstanceAccelerationStructure) TransformBuffer() objc.ID {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("transformBuffer"))
	return rv
}


// SetTransformBuffer sets the value of the transformBuffer property.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSInstanceAccelerationStructure/transformBuffer
func (i_ InstanceAccelerationStructure) SetTransformBuffer(value objc.ID) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setTransformBuffer:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsinstanceaccelerationstructure/accelerationstructures
func (i_ InstanceAccelerationStructure) AccelerationStructures() MPSPolygonAccelerationStructure {
	rv := objc.Send[MPSPolygonAccelerationStructure](i_.ID, objc.Sel("accelerationStructures"))
	return rv
}


// SetAccelerationStructures sets the value of the accelerationStructures property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsinstanceaccelerationstructure/accelerationstructures
func (i_ InstanceAccelerationStructure) SetAccelerationStructures(value IMPSPolygonAccelerationStructure) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setAccelerationStructures:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsinstanceaccelerationstructure/instancebuffer
func (i_ InstanceAccelerationStructure) InstanceBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("instanceBuffer"))
	return rv
}


// SetInstanceBuffer sets the value of the instanceBuffer property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsinstanceaccelerationstructure/instancebuffer
func (i_ InstanceAccelerationStructure) SetInstanceBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setInstanceBuffer:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsinstanceaccelerationstructure/instancebufferoffset
func (i_ InstanceAccelerationStructure) InstanceBufferOffset() int {
	rv := objc.Send[int](i_.ID, objc.Sel("instanceBufferOffset"))
	return rv
}


// SetInstanceBufferOffset sets the value of the instanceBufferOffset property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsinstanceaccelerationstructure/instancebufferoffset
func (i_ InstanceAccelerationStructure) SetInstanceBufferOffset(value int) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setInstanceBufferOffset:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsinstanceaccelerationstructure/instancecount
func (i_ InstanceAccelerationStructure) InstanceCount() int {
	rv := objc.Send[int](i_.ID, objc.Sel("instanceCount"))
	return rv
}


// SetInstanceCount sets the value of the instanceCount property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsinstanceaccelerationstructure/instancecount
func (i_ InstanceAccelerationStructure) SetInstanceCount(value int) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setInstanceCount:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsinstanceaccelerationstructure/maskbuffer
func (i_ InstanceAccelerationStructure) MaskBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("maskBuffer"))
	return rv
}


// SetMaskBuffer sets the value of the maskBuffer property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsinstanceaccelerationstructure/maskbuffer
func (i_ InstanceAccelerationStructure) SetMaskBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMaskBuffer:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsinstanceaccelerationstructure/maskbufferoffset
func (i_ InstanceAccelerationStructure) MaskBufferOffset() int {
	rv := objc.Send[int](i_.ID, objc.Sel("maskBufferOffset"))
	return rv
}


// SetMaskBufferOffset sets the value of the maskBufferOffset property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsinstanceaccelerationstructure/maskbufferoffset
func (i_ InstanceAccelerationStructure) SetMaskBufferOffset(value int) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMaskBufferOffset:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsinstanceaccelerationstructure/transformbufferoffset
func (i_ InstanceAccelerationStructure) TransformBufferOffset() int {
	rv := objc.Send[int](i_.ID, objc.Sel("transformBufferOffset"))
	return rv
}


// SetTransformBufferOffset sets the value of the transformBufferOffset property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsinstanceaccelerationstructure/transformbufferoffset
func (i_ InstanceAccelerationStructure) SetTransformBufferOffset(value int) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setTransformBufferOffset:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsinstanceaccelerationstructure/transformtype
func (i_ InstanceAccelerationStructure) TransformType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("transformType"))
	return rv
}


// SetTransformType sets the value of the transformType property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsinstanceaccelerationstructure/transformtype
func (i_ InstanceAccelerationStructure) SetTransformType(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setTransformType:"), value)
}



