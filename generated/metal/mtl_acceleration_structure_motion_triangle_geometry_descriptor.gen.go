// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AccelerationStructureMotionTriangleGeometryDescriptor] class.
var (
	AccelerationStructureMotionTriangleGeometryDescriptorClass     _AccelerationStructureMotionTriangleGeometryDescriptorClass
	AccelerationStructureMotionTriangleGeometryDescriptorClassOnce sync.Once
)

func getAccelerationStructureMotionTriangleGeometryDescriptorClass() _AccelerationStructureMotionTriangleGeometryDescriptorClass {
	AccelerationStructureMotionTriangleGeometryDescriptorClassOnce.Do(func() {
		AccelerationStructureMotionTriangleGeometryDescriptorClass = _AccelerationStructureMotionTriangleGeometryDescriptorClass{objc.GetClass("MTLAccelerationStructureMotionTriangleGeometryDescriptor")}
	})
	return AccelerationStructureMotionTriangleGeometryDescriptorClass
}

type _AccelerationStructureMotionTriangleGeometryDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [AccelerationStructureMotionTriangleGeometryDescriptor] class.
type IAccelerationStructureMotionTriangleGeometryDescriptor interface {
	IAccelerationStructureGeometryDescriptor
}

// A description of a list of triangle primitives, as motion keyframe data, to turn into an acceleration structure.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionTriangleGeometryDescriptor
type AccelerationStructureMotionTriangleGeometryDescriptor struct {
	AccelerationStructureGeometryDescriptor
}

// AccelerationStructureMotionTriangleGeometryDescriptorFrom constructs a [AccelerationStructureMotionTriangleGeometryDescriptor] from an unsafe.Pointer.
//
// A description of a list of triangle primitives, as motion keyframe data, to turn into an acceleration structure.
func AccelerationStructureMotionTriangleGeometryDescriptorFrom(ptr unsafe.Pointer) AccelerationStructureMotionTriangleGeometryDescriptor {
	return AccelerationStructureMotionTriangleGeometryDescriptor{
		AccelerationStructureGeometryDescriptor: AccelerationStructureGeometryDescriptorFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AccelerationStructureMotionTriangleGeometryDescriptorClass) Alloc() AccelerationStructureMotionTriangleGeometryDescriptor {
	rv := objc.Send[AccelerationStructureMotionTriangleGeometryDescriptor](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AccelerationStructureMotionTriangleGeometryDescriptorClass) New() AccelerationStructureMotionTriangleGeometryDescriptor {
	rv := objc.Send[AccelerationStructureMotionTriangleGeometryDescriptor](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AccelerationStructureMotionTriangleGeometryDescriptor) Init() AccelerationStructureMotionTriangleGeometryDescriptor {
	rv := objc.Send[AccelerationStructureMotionTriangleGeometryDescriptor](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AccelerationStructureMotionTriangleGeometryDescriptor) Autorelease() AccelerationStructureMotionTriangleGeometryDescriptor {
	rv := objc.Send[AccelerationStructureMotionTriangleGeometryDescriptor](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAccelerationStructureMotionTriangleGeometryDescriptor creates a new AccelerationStructureMotionTriangleGeometryDescriptor instance.
func NewAccelerationStructureMotionTriangleGeometryDescriptor() AccelerationStructureMotionTriangleGeometryDescriptor {
	return getAccelerationStructureMotionTriangleGeometryDescriptorClass().New()
}




