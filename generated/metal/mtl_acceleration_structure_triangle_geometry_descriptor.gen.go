// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AccelerationStructureTriangleGeometryDescriptor] class.
var (
	AccelerationStructureTriangleGeometryDescriptorClass     _AccelerationStructureTriangleGeometryDescriptorClass
	AccelerationStructureTriangleGeometryDescriptorClassOnce sync.Once
)

func getAccelerationStructureTriangleGeometryDescriptorClass() _AccelerationStructureTriangleGeometryDescriptorClass {
	AccelerationStructureTriangleGeometryDescriptorClassOnce.Do(func() {
		AccelerationStructureTriangleGeometryDescriptorClass = _AccelerationStructureTriangleGeometryDescriptorClass{objc.GetClass("MTLAccelerationStructureTriangleGeometryDescriptor")}
	})
	return AccelerationStructureTriangleGeometryDescriptorClass
}

type _AccelerationStructureTriangleGeometryDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [AccelerationStructureTriangleGeometryDescriptor] class.
type IAccelerationStructureTriangleGeometryDescriptor interface {
	IAccelerationStructureGeometryDescriptor
}

// A description of a list of triangle primitives to turn into an acceleration structure.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureTriangleGeometryDescriptor
type AccelerationStructureTriangleGeometryDescriptor struct {
	AccelerationStructureGeometryDescriptor
}

// AccelerationStructureTriangleGeometryDescriptorFrom constructs a [AccelerationStructureTriangleGeometryDescriptor] from an unsafe.Pointer.
//
// A description of a list of triangle primitives to turn into an acceleration structure.
func AccelerationStructureTriangleGeometryDescriptorFrom(ptr unsafe.Pointer) AccelerationStructureTriangleGeometryDescriptor {
	return AccelerationStructureTriangleGeometryDescriptor{
		AccelerationStructureGeometryDescriptor: AccelerationStructureGeometryDescriptorFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AccelerationStructureTriangleGeometryDescriptorClass) Alloc() AccelerationStructureTriangleGeometryDescriptor {
	rv := objc.Send[AccelerationStructureTriangleGeometryDescriptor](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AccelerationStructureTriangleGeometryDescriptorClass) New() AccelerationStructureTriangleGeometryDescriptor {
	rv := objc.Send[AccelerationStructureTriangleGeometryDescriptor](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AccelerationStructureTriangleGeometryDescriptor) Init() AccelerationStructureTriangleGeometryDescriptor {
	rv := objc.Send[AccelerationStructureTriangleGeometryDescriptor](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AccelerationStructureTriangleGeometryDescriptor) Autorelease() AccelerationStructureTriangleGeometryDescriptor {
	rv := objc.Send[AccelerationStructureTriangleGeometryDescriptor](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAccelerationStructureTriangleGeometryDescriptor creates a new AccelerationStructureTriangleGeometryDescriptor instance.
func NewAccelerationStructureTriangleGeometryDescriptor() AccelerationStructureTriangleGeometryDescriptor {
	return getAccelerationStructureTriangleGeometryDescriptorClass().New()
}




