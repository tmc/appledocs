// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
}

//
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

// Alloc allocates a new instance without initialization.
func (ac _AccelerationStructureCurveGeometryDescriptorClass) Alloc() AccelerationStructureCurveGeometryDescriptor {
	rv := objc.Send[AccelerationStructureCurveGeometryDescriptor](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




