// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTL4AccelerationStructureDescriptor] class.
var (
	MTL4AccelerationStructureDescriptorClass     _MTL4AccelerationStructureDescriptorClass
	MTL4AccelerationStructureDescriptorClassOnce sync.Once
)

func getMTL4AccelerationStructureDescriptorClass() _MTL4AccelerationStructureDescriptorClass {
	MTL4AccelerationStructureDescriptorClassOnce.Do(func() {
		MTL4AccelerationStructureDescriptorClass = _MTL4AccelerationStructureDescriptorClass{objc.GetClass("MTL4AccelerationStructureDescriptor")}
	})
	return MTL4AccelerationStructureDescriptorClass
}

type _MTL4AccelerationStructureDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [MTL4AccelerationStructureDescriptor] class.
type IMTL4AccelerationStructureDescriptor interface {
	IAccelerationStructureDescriptor
}

// Base class for Metal 4 acceleration structure descriptors.
//
// Don’t use this class directly. Use one of its subclasses instead.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureDescriptor
type MTL4AccelerationStructureDescriptor struct {
	AccelerationStructureDescriptor
}

// MTL4AccelerationStructureDescriptorFrom constructs a [MTL4AccelerationStructureDescriptor] from an unsafe.Pointer.
//
// Base class for Metal 4 acceleration structure descriptors.
func MTL4AccelerationStructureDescriptorFrom(ptr unsafe.Pointer) MTL4AccelerationStructureDescriptor {
	return MTL4AccelerationStructureDescriptor{
		AccelerationStructureDescriptor: AccelerationStructureDescriptorFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTL4AccelerationStructureDescriptorClass) Alloc() MTL4AccelerationStructureDescriptor {
	rv := objc.Send[MTL4AccelerationStructureDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTL4AccelerationStructureDescriptorClass) New() MTL4AccelerationStructureDescriptor {
	rv := objc.Send[MTL4AccelerationStructureDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTL4AccelerationStructureDescriptor) Init() MTL4AccelerationStructureDescriptor {
	rv := objc.Send[MTL4AccelerationStructureDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTL4AccelerationStructureDescriptor) Autorelease() MTL4AccelerationStructureDescriptor {
	rv := objc.Send[MTL4AccelerationStructureDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4AccelerationStructureDescriptor creates a new MTL4AccelerationStructureDescriptor instance.
func NewMTL4AccelerationStructureDescriptor() MTL4AccelerationStructureDescriptor {
	return getMTL4AccelerationStructureDescriptorClass().New()
}




