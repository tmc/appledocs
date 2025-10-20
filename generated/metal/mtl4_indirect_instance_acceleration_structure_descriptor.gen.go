// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTL4IndirectInstanceAccelerationStructureDescriptor] class.
var (
	MTL4IndirectInstanceAccelerationStructureDescriptorClass     _MTL4IndirectInstanceAccelerationStructureDescriptorClass
	MTL4IndirectInstanceAccelerationStructureDescriptorClassOnce sync.Once
)

func getMTL4IndirectInstanceAccelerationStructureDescriptorClass() _MTL4IndirectInstanceAccelerationStructureDescriptorClass {
	MTL4IndirectInstanceAccelerationStructureDescriptorClassOnce.Do(func() {
		MTL4IndirectInstanceAccelerationStructureDescriptorClass = _MTL4IndirectInstanceAccelerationStructureDescriptorClass{objc.GetClass("MTL4IndirectInstanceAccelerationStructureDescriptor")}
	})
	return MTL4IndirectInstanceAccelerationStructureDescriptorClass
}

type _MTL4IndirectInstanceAccelerationStructureDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [MTL4IndirectInstanceAccelerationStructureDescriptor] class.
type IMTL4IndirectInstanceAccelerationStructureDescriptor interface {
	IMTL4AccelerationStructureDescriptor
}

// Descriptor for an “indirect” instance acceleration structure that allows providing the instance count and motion transform count indirectly, through buffer references.
//
// An instance acceleration structure references other acceleration structures, and provides the ability to “instantiate” them multiple times, each one with potentially a different transformation matrix. You specify the properties of the instances in the acceleration structure this descriptor builds by providing a buffer of via its property. Compared to , this descriptor allows you to provide the number of instances it references indirectly through a buffer reference, as well as the number of motion transforms. This enables you to determine these counts indirectly in the GPU timeline via a compute pipeline. Metal needs only to know the maximum possible number of instances and motion transforms to support, which you specify via the and properties. Use a to mark residency of all buffers and acceleration structures this descriptor references when you build this acceleration structure.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4IndirectInstanceAccelerationStructureDescriptor
type MTL4IndirectInstanceAccelerationStructureDescriptor struct {
	MTL4AccelerationStructureDescriptor
}

// MTL4IndirectInstanceAccelerationStructureDescriptorFrom constructs a [MTL4IndirectInstanceAccelerationStructureDescriptor] from an unsafe.Pointer.
//
// Descriptor for an “indirect” instance acceleration structure that allows providing the instance count and motion transform count indirectly, through buffer references.
func MTL4IndirectInstanceAccelerationStructureDescriptorFrom(ptr unsafe.Pointer) MTL4IndirectInstanceAccelerationStructureDescriptor {
	return MTL4IndirectInstanceAccelerationStructureDescriptor{
		MTL4AccelerationStructureDescriptor: MTL4AccelerationStructureDescriptorFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTL4IndirectInstanceAccelerationStructureDescriptorClass) Alloc() MTL4IndirectInstanceAccelerationStructureDescriptor {
	rv := objc.Send[MTL4IndirectInstanceAccelerationStructureDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTL4IndirectInstanceAccelerationStructureDescriptorClass) New() MTL4IndirectInstanceAccelerationStructureDescriptor {
	rv := objc.Send[MTL4IndirectInstanceAccelerationStructureDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTL4IndirectInstanceAccelerationStructureDescriptor) Init() MTL4IndirectInstanceAccelerationStructureDescriptor {
	rv := objc.Send[MTL4IndirectInstanceAccelerationStructureDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTL4IndirectInstanceAccelerationStructureDescriptor) Autorelease() MTL4IndirectInstanceAccelerationStructureDescriptor {
	rv := objc.Send[MTL4IndirectInstanceAccelerationStructureDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4IndirectInstanceAccelerationStructureDescriptor creates a new MTL4IndirectInstanceAccelerationStructureDescriptor instance.
func NewMTL4IndirectInstanceAccelerationStructureDescriptor() MTL4IndirectInstanceAccelerationStructureDescriptor {
	return getMTL4IndirectInstanceAccelerationStructureDescriptorClass().New()
}




