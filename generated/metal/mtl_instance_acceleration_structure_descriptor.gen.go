// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [InstanceAccelerationStructureDescriptor] class.
var (
	InstanceAccelerationStructureDescriptorClass     _InstanceAccelerationStructureDescriptorClass
	InstanceAccelerationStructureDescriptorClassOnce sync.Once
)

func getInstanceAccelerationStructureDescriptorClass() _InstanceAccelerationStructureDescriptorClass {
	InstanceAccelerationStructureDescriptorClassOnce.Do(func() {
		InstanceAccelerationStructureDescriptorClass = _InstanceAccelerationStructureDescriptorClass{objc.GetClass("MTLInstanceAccelerationStructureDescriptor")}
	})
	return InstanceAccelerationStructureDescriptorClass
}

type _InstanceAccelerationStructureDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [InstanceAccelerationStructureDescriptor] class.
type IInstanceAccelerationStructureDescriptor interface {
	IAccelerationStructureDescriptor
}

// A description of an acceleration structure that derives from instances of primitive acceleration structures.
//
// Metal provides acceleration structures with a two-level hierarchy. The bottom layer consists of primitive acceleration structures, which instance acceleration structures in the top level reference.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLInstanceAccelerationStructureDescriptor
type InstanceAccelerationStructureDescriptor struct {
	AccelerationStructureDescriptor
}

// InstanceAccelerationStructureDescriptorFrom constructs a [InstanceAccelerationStructureDescriptor] from an unsafe.Pointer.
//
// A description of an acceleration structure that derives from instances of primitive acceleration structures.
func InstanceAccelerationStructureDescriptorFrom(ptr unsafe.Pointer) InstanceAccelerationStructureDescriptor {
	return InstanceAccelerationStructureDescriptor{
		AccelerationStructureDescriptor: AccelerationStructureDescriptorFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _InstanceAccelerationStructureDescriptorClass) Alloc() InstanceAccelerationStructureDescriptor {
	rv := objc.Send[InstanceAccelerationStructureDescriptor](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _InstanceAccelerationStructureDescriptorClass) New() InstanceAccelerationStructureDescriptor {
	rv := objc.Send[InstanceAccelerationStructureDescriptor](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ InstanceAccelerationStructureDescriptor) Init() InstanceAccelerationStructureDescriptor {
	rv := objc.Send[InstanceAccelerationStructureDescriptor](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ InstanceAccelerationStructureDescriptor) Autorelease() InstanceAccelerationStructureDescriptor {
	rv := objc.Send[InstanceAccelerationStructureDescriptor](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewInstanceAccelerationStructureDescriptor creates a new InstanceAccelerationStructureDescriptor instance.
func NewInstanceAccelerationStructureDescriptor() InstanceAccelerationStructureDescriptor {
	return getInstanceAccelerationStructureDescriptorClass().New()
}




