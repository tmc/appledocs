// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTL4PrimitiveAccelerationStructureDescriptor] class.
var (
	MTL4PrimitiveAccelerationStructureDescriptorClass     _MTL4PrimitiveAccelerationStructureDescriptorClass
	MTL4PrimitiveAccelerationStructureDescriptorClassOnce sync.Once
)

func getMTL4PrimitiveAccelerationStructureDescriptorClass() _MTL4PrimitiveAccelerationStructureDescriptorClass {
	MTL4PrimitiveAccelerationStructureDescriptorClassOnce.Do(func() {
		MTL4PrimitiveAccelerationStructureDescriptorClass = _MTL4PrimitiveAccelerationStructureDescriptorClass{objc.GetClass("MTL4PrimitiveAccelerationStructureDescriptor")}
	})
	return MTL4PrimitiveAccelerationStructureDescriptorClass
}

type _MTL4PrimitiveAccelerationStructureDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [MTL4PrimitiveAccelerationStructureDescriptor] class.
type IMTL4PrimitiveAccelerationStructureDescriptor interface {
	IMTL4AccelerationStructureDescriptor
}

// Descriptor for a primitive acceleration structure that directly references geometric shapes, such as triangles and bounding boxes.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4PrimitiveAccelerationStructureDescriptor
type MTL4PrimitiveAccelerationStructureDescriptor struct {
	MTL4AccelerationStructureDescriptor
}

// MTL4PrimitiveAccelerationStructureDescriptorFrom constructs a [MTL4PrimitiveAccelerationStructureDescriptor] from an unsafe.Pointer.
//
// Descriptor for a primitive acceleration structure that directly references geometric shapes, such as triangles and bounding boxes.
func MTL4PrimitiveAccelerationStructureDescriptorFrom(ptr unsafe.Pointer) MTL4PrimitiveAccelerationStructureDescriptor {
	return MTL4PrimitiveAccelerationStructureDescriptor{
		MTL4AccelerationStructureDescriptor: MTL4AccelerationStructureDescriptorFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTL4PrimitiveAccelerationStructureDescriptorClass) Alloc() MTL4PrimitiveAccelerationStructureDescriptor {
	rv := objc.Send[MTL4PrimitiveAccelerationStructureDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTL4PrimitiveAccelerationStructureDescriptorClass) New() MTL4PrimitiveAccelerationStructureDescriptor {
	rv := objc.Send[MTL4PrimitiveAccelerationStructureDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTL4PrimitiveAccelerationStructureDescriptor) Init() MTL4PrimitiveAccelerationStructureDescriptor {
	rv := objc.Send[MTL4PrimitiveAccelerationStructureDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTL4PrimitiveAccelerationStructureDescriptor) Autorelease() MTL4PrimitiveAccelerationStructureDescriptor {
	rv := objc.Send[MTL4PrimitiveAccelerationStructureDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4PrimitiveAccelerationStructureDescriptor creates a new MTL4PrimitiveAccelerationStructureDescriptor instance.
func NewMTL4PrimitiveAccelerationStructureDescriptor() MTL4PrimitiveAccelerationStructureDescriptor {
	return getMTL4PrimitiveAccelerationStructureDescriptorClass().New()
}




