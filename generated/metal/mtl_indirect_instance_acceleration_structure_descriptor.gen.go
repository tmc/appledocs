// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [IndirectInstanceAccelerationStructureDescriptor] class.
var (
	IndirectInstanceAccelerationStructureDescriptorClass     _IndirectInstanceAccelerationStructureDescriptorClass
	IndirectInstanceAccelerationStructureDescriptorClassOnce sync.Once
)

func getIndirectInstanceAccelerationStructureDescriptorClass() _IndirectInstanceAccelerationStructureDescriptorClass {
	IndirectInstanceAccelerationStructureDescriptorClassOnce.Do(func() {
		IndirectInstanceAccelerationStructureDescriptorClass = _IndirectInstanceAccelerationStructureDescriptorClass{objc.GetClass("MTLIndirectInstanceAccelerationStructureDescriptor")}
	})
	return IndirectInstanceAccelerationStructureDescriptorClass
}

type _IndirectInstanceAccelerationStructureDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [IndirectInstanceAccelerationStructureDescriptor] class.
type IIndirectInstanceAccelerationStructureDescriptor interface {
	IAccelerationStructureDescriptor
}

// A description of an acceleration structure that Metal derives from instances of primitive acceleration structures that the GPU can populate.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectInstanceAccelerationStructureDescriptor
type IndirectInstanceAccelerationStructureDescriptor struct {
	AccelerationStructureDescriptor
}

// IndirectInstanceAccelerationStructureDescriptorFrom constructs a [IndirectInstanceAccelerationStructureDescriptor] from an unsafe.Pointer.
//
// A description of an acceleration structure that Metal derives from instances of primitive acceleration structures that the GPU can populate.
func IndirectInstanceAccelerationStructureDescriptorFrom(ptr unsafe.Pointer) IndirectInstanceAccelerationStructureDescriptor {
	return IndirectInstanceAccelerationStructureDescriptor{
		AccelerationStructureDescriptor: AccelerationStructureDescriptorFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _IndirectInstanceAccelerationStructureDescriptorClass) Alloc() IndirectInstanceAccelerationStructureDescriptor {
	rv := objc.Send[IndirectInstanceAccelerationStructureDescriptor](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _IndirectInstanceAccelerationStructureDescriptorClass) New() IndirectInstanceAccelerationStructureDescriptor {
	rv := objc.Send[IndirectInstanceAccelerationStructureDescriptor](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ IndirectInstanceAccelerationStructureDescriptor) Init() IndirectInstanceAccelerationStructureDescriptor {
	rv := objc.Send[IndirectInstanceAccelerationStructureDescriptor](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ IndirectInstanceAccelerationStructureDescriptor) Autorelease() IndirectInstanceAccelerationStructureDescriptor {
	rv := objc.Send[IndirectInstanceAccelerationStructureDescriptor](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewIndirectInstanceAccelerationStructureDescriptor creates a new IndirectInstanceAccelerationStructureDescriptor instance.
func NewIndirectInstanceAccelerationStructureDescriptor() IndirectInstanceAccelerationStructureDescriptor {
	return getIndirectInstanceAccelerationStructureDescriptorClass().New()
}




