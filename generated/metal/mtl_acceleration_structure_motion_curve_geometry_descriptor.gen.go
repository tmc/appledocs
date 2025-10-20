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




