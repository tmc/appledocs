// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PrimitiveAccelerationStructureDescriptor] class.
var (
	PrimitiveAccelerationStructureDescriptorClass     _PrimitiveAccelerationStructureDescriptorClass
	PrimitiveAccelerationStructureDescriptorClassOnce sync.Once
)

func getPrimitiveAccelerationStructureDescriptorClass() _PrimitiveAccelerationStructureDescriptorClass {
	PrimitiveAccelerationStructureDescriptorClassOnce.Do(func() {
		PrimitiveAccelerationStructureDescriptorClass = _PrimitiveAccelerationStructureDescriptorClass{objc.GetClass("MTLPrimitiveAccelerationStructureDescriptor")}
	})
	return PrimitiveAccelerationStructureDescriptorClass
}

type _PrimitiveAccelerationStructureDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [PrimitiveAccelerationStructureDescriptor] class.
type IPrimitiveAccelerationStructureDescriptor interface {
	IAccelerationStructureDescriptor
}

// A description of an acceleration structure that contains geometry primitives.
//
// Metal provides acceleration structures with a two-level hierarchy. The bottom layer consists of primitive acceleration structures, which instance acceleration structures in the top level reference.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPrimitiveAccelerationStructureDescriptor
type PrimitiveAccelerationStructureDescriptor struct {
	AccelerationStructureDescriptor
}

// PrimitiveAccelerationStructureDescriptorFrom constructs a [PrimitiveAccelerationStructureDescriptor] from an unsafe.Pointer.
//
// A description of an acceleration structure that contains geometry primitives.
func PrimitiveAccelerationStructureDescriptorFrom(ptr unsafe.Pointer) PrimitiveAccelerationStructureDescriptor {
	return PrimitiveAccelerationStructureDescriptor{
		AccelerationStructureDescriptor: AccelerationStructureDescriptorFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PrimitiveAccelerationStructureDescriptorClass) Alloc() PrimitiveAccelerationStructureDescriptor {
	rv := objc.Send[PrimitiveAccelerationStructureDescriptor](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PrimitiveAccelerationStructureDescriptorClass) New() PrimitiveAccelerationStructureDescriptor {
	rv := objc.Send[PrimitiveAccelerationStructureDescriptor](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PrimitiveAccelerationStructureDescriptor) Init() PrimitiveAccelerationStructureDescriptor {
	rv := objc.Send[PrimitiveAccelerationStructureDescriptor](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PrimitiveAccelerationStructureDescriptor) Autorelease() PrimitiveAccelerationStructureDescriptor {
	rv := objc.Send[PrimitiveAccelerationStructureDescriptor](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPrimitiveAccelerationStructureDescriptor creates a new PrimitiveAccelerationStructureDescriptor instance.
func NewPrimitiveAccelerationStructureDescriptor() PrimitiveAccelerationStructureDescriptor {
	return getPrimitiveAccelerationStructureDescriptorClass().New()
}




