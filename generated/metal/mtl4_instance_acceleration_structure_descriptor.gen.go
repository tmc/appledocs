// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTL4InstanceAccelerationStructureDescriptor] class.
var (
	MTL4InstanceAccelerationStructureDescriptorClass     _MTL4InstanceAccelerationStructureDescriptorClass
	MTL4InstanceAccelerationStructureDescriptorClassOnce sync.Once
)

func getMTL4InstanceAccelerationStructureDescriptorClass() _MTL4InstanceAccelerationStructureDescriptorClass {
	MTL4InstanceAccelerationStructureDescriptorClassOnce.Do(func() {
		MTL4InstanceAccelerationStructureDescriptorClass = _MTL4InstanceAccelerationStructureDescriptorClass{objc.GetClass("MTL4InstanceAccelerationStructureDescriptor")}
	})
	return MTL4InstanceAccelerationStructureDescriptorClass
}

type _MTL4InstanceAccelerationStructureDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [MTL4InstanceAccelerationStructureDescriptor] class.
type IMTL4InstanceAccelerationStructureDescriptor interface {
	IMTL4AccelerationStructureDescriptor
}

// Descriptor for an instance acceleration structure.
//
// An instance acceleration structure references other acceleration structures, and provides the ability to “instantiate” them multiple times, each one with potentially a different transformation matrix. You specify the properties of the instances in the acceleration structure this descriptor builds by providing a buffer of via its property. Use a to mark residency of all buffers and acceleration structures this descriptor references when you build this acceleration structure.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4InstanceAccelerationStructureDescriptor
type MTL4InstanceAccelerationStructureDescriptor struct {
	MTL4AccelerationStructureDescriptor
}

// MTL4InstanceAccelerationStructureDescriptorFrom constructs a [MTL4InstanceAccelerationStructureDescriptor] from an unsafe.Pointer.
//
// Descriptor for an instance acceleration structure.
func MTL4InstanceAccelerationStructureDescriptorFrom(ptr unsafe.Pointer) MTL4InstanceAccelerationStructureDescriptor {
	return MTL4InstanceAccelerationStructureDescriptor{
		MTL4AccelerationStructureDescriptor: MTL4AccelerationStructureDescriptorFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTL4InstanceAccelerationStructureDescriptorClass) Alloc() MTL4InstanceAccelerationStructureDescriptor {
	rv := objc.Send[MTL4InstanceAccelerationStructureDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTL4InstanceAccelerationStructureDescriptorClass) New() MTL4InstanceAccelerationStructureDescriptor {
	rv := objc.Send[MTL4InstanceAccelerationStructureDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTL4InstanceAccelerationStructureDescriptor) Init() MTL4InstanceAccelerationStructureDescriptor {
	rv := objc.Send[MTL4InstanceAccelerationStructureDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTL4InstanceAccelerationStructureDescriptor) Autorelease() MTL4InstanceAccelerationStructureDescriptor {
	rv := objc.Send[MTL4InstanceAccelerationStructureDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4InstanceAccelerationStructureDescriptor creates a new MTL4InstanceAccelerationStructureDescriptor instance.
func NewMTL4InstanceAccelerationStructureDescriptor() MTL4InstanceAccelerationStructureDescriptor {
	return getMTL4InstanceAccelerationStructureDescriptorClass().New()
}




