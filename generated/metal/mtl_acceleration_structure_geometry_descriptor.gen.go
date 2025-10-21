// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [AccelerationStructureGeometryDescriptor] class.
var (
	AccelerationStructureGeometryDescriptorClass     _AccelerationStructureGeometryDescriptorClass
	AccelerationStructureGeometryDescriptorClassOnce sync.Once
)

func getAccelerationStructureGeometryDescriptorClass() _AccelerationStructureGeometryDescriptorClass {
	AccelerationStructureGeometryDescriptorClassOnce.Do(func() {
		AccelerationStructureGeometryDescriptorClass = _AccelerationStructureGeometryDescriptorClass{objc.GetClass("MTLAccelerationStructureGeometryDescriptor")}
	})
	return AccelerationStructureGeometryDescriptorClass
}

type _AccelerationStructureGeometryDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [AccelerationStructureGeometryDescriptor] class.
type IAccelerationStructureGeometryDescriptor interface {
	objectivec.IObject
}

// A base class for descriptors that contain geometry data to convert into a ray-tracing acceleration structure.
//
// Don’t use this base class directly. Use one of the derived classes instead, as describes.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureGeometryDescriptor
type AccelerationStructureGeometryDescriptor struct {
	objectivec.Object
}

// AccelerationStructureGeometryDescriptorFrom constructs a [AccelerationStructureGeometryDescriptor] from an unsafe.Pointer.
//
// A base class for descriptors that contain geometry data to convert into a ray-tracing acceleration structure.
func AccelerationStructureGeometryDescriptorFrom(ptr unsafe.Pointer) AccelerationStructureGeometryDescriptor {
	return AccelerationStructureGeometryDescriptor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AccelerationStructureGeometryDescriptorClass) Alloc() AccelerationStructureGeometryDescriptor {
	rv := objc.Send[AccelerationStructureGeometryDescriptor](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AccelerationStructureGeometryDescriptorClass) New() AccelerationStructureGeometryDescriptor {
	rv := objc.Send[AccelerationStructureGeometryDescriptor](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AccelerationStructureGeometryDescriptor) Init() AccelerationStructureGeometryDescriptor {
	rv := objc.Send[AccelerationStructureGeometryDescriptor](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AccelerationStructureGeometryDescriptor) Autorelease() AccelerationStructureGeometryDescriptor {
	rv := objc.Send[AccelerationStructureGeometryDescriptor](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAccelerationStructureGeometryDescriptor creates a new AccelerationStructureGeometryDescriptor instance.
func NewAccelerationStructureGeometryDescriptor() AccelerationStructureGeometryDescriptor {
	return getAccelerationStructureGeometryDescriptorClass().New()
}




