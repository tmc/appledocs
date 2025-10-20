// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AccelerationStructureBoundingBoxGeometryDescriptor] class.
var (
	AccelerationStructureBoundingBoxGeometryDescriptorClass     _AccelerationStructureBoundingBoxGeometryDescriptorClass
	AccelerationStructureBoundingBoxGeometryDescriptorClassOnce sync.Once
)

func getAccelerationStructureBoundingBoxGeometryDescriptorClass() _AccelerationStructureBoundingBoxGeometryDescriptorClass {
	AccelerationStructureBoundingBoxGeometryDescriptorClassOnce.Do(func() {
		AccelerationStructureBoundingBoxGeometryDescriptorClass = _AccelerationStructureBoundingBoxGeometryDescriptorClass{objc.GetClass("MTLAccelerationStructureBoundingBoxGeometryDescriptor")}
	})
	return AccelerationStructureBoundingBoxGeometryDescriptorClass
}

type _AccelerationStructureBoundingBoxGeometryDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [AccelerationStructureBoundingBoxGeometryDescriptor] class.
type IAccelerationStructureBoundingBoxGeometryDescriptor interface {
	IAccelerationStructureGeometryDescriptor
}

// A description of a list of bounding boxes to turn into an acceleration structure.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureBoundingBoxGeometryDescriptor
type AccelerationStructureBoundingBoxGeometryDescriptor struct {
	AccelerationStructureGeometryDescriptor
}

// AccelerationStructureBoundingBoxGeometryDescriptorFrom constructs a [AccelerationStructureBoundingBoxGeometryDescriptor] from an unsafe.Pointer.
//
// A description of a list of bounding boxes to turn into an acceleration structure.
func AccelerationStructureBoundingBoxGeometryDescriptorFrom(ptr unsafe.Pointer) AccelerationStructureBoundingBoxGeometryDescriptor {
	return AccelerationStructureBoundingBoxGeometryDescriptor{
		AccelerationStructureGeometryDescriptor: AccelerationStructureGeometryDescriptorFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AccelerationStructureBoundingBoxGeometryDescriptorClass) Alloc() AccelerationStructureBoundingBoxGeometryDescriptor {
	rv := objc.Send[AccelerationStructureBoundingBoxGeometryDescriptor](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AccelerationStructureBoundingBoxGeometryDescriptorClass) New() AccelerationStructureBoundingBoxGeometryDescriptor {
	rv := objc.Send[AccelerationStructureBoundingBoxGeometryDescriptor](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AccelerationStructureBoundingBoxGeometryDescriptor) Init() AccelerationStructureBoundingBoxGeometryDescriptor {
	rv := objc.Send[AccelerationStructureBoundingBoxGeometryDescriptor](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AccelerationStructureBoundingBoxGeometryDescriptor) Autorelease() AccelerationStructureBoundingBoxGeometryDescriptor {
	rv := objc.Send[AccelerationStructureBoundingBoxGeometryDescriptor](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAccelerationStructureBoundingBoxGeometryDescriptor creates a new AccelerationStructureBoundingBoxGeometryDescriptor instance.
func NewAccelerationStructureBoundingBoxGeometryDescriptor() AccelerationStructureBoundingBoxGeometryDescriptor {
	return getAccelerationStructureBoundingBoxGeometryDescriptorClass().New()
}




