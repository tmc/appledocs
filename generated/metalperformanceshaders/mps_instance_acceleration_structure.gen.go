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



