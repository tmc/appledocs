// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTL4AccelerationStructureTriangleGeometryDescriptor] class.
var (
	MTL4AccelerationStructureTriangleGeometryDescriptorClass     _MTL4AccelerationStructureTriangleGeometryDescriptorClass
	MTL4AccelerationStructureTriangleGeometryDescriptorClassOnce sync.Once
)

func getMTL4AccelerationStructureTriangleGeometryDescriptorClass() _MTL4AccelerationStructureTriangleGeometryDescriptorClass {
	MTL4AccelerationStructureTriangleGeometryDescriptorClassOnce.Do(func() {
		MTL4AccelerationStructureTriangleGeometryDescriptorClass = _MTL4AccelerationStructureTriangleGeometryDescriptorClass{objc.GetClass("MTL4AccelerationStructureTriangleGeometryDescriptor")}
	})
	return MTL4AccelerationStructureTriangleGeometryDescriptorClass
}

type _MTL4AccelerationStructureTriangleGeometryDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [MTL4AccelerationStructureTriangleGeometryDescriptor] class.
type IMTL4AccelerationStructureTriangleGeometryDescriptor interface {
	IMTL4AccelerationStructureGeometryDescriptor
}

// Describes triangle geometry suitable for ray tracing.
//
// Use a to mark residency of all buffers this descriptor references when you build this acceleration structure.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureTriangleGeometryDescriptor
type MTL4AccelerationStructureTriangleGeometryDescriptor struct {
	MTL4AccelerationStructureGeometryDescriptor
}

// MTL4AccelerationStructureTriangleGeometryDescriptorFrom constructs a [MTL4AccelerationStructureTriangleGeometryDescriptor] from an unsafe.Pointer.
//
// Describes triangle geometry suitable for ray tracing.
func MTL4AccelerationStructureTriangleGeometryDescriptorFrom(ptr unsafe.Pointer) MTL4AccelerationStructureTriangleGeometryDescriptor {
	return MTL4AccelerationStructureTriangleGeometryDescriptor{
		MTL4AccelerationStructureGeometryDescriptor: MTL4AccelerationStructureGeometryDescriptorFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTL4AccelerationStructureTriangleGeometryDescriptorClass) Alloc() MTL4AccelerationStructureTriangleGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureTriangleGeometryDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTL4AccelerationStructureTriangleGeometryDescriptorClass) New() MTL4AccelerationStructureTriangleGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureTriangleGeometryDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTL4AccelerationStructureTriangleGeometryDescriptor) Init() MTL4AccelerationStructureTriangleGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureTriangleGeometryDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTL4AccelerationStructureTriangleGeometryDescriptor) Autorelease() MTL4AccelerationStructureTriangleGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureTriangleGeometryDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4AccelerationStructureTriangleGeometryDescriptor creates a new MTL4AccelerationStructureTriangleGeometryDescriptor instance.
func NewMTL4AccelerationStructureTriangleGeometryDescriptor() MTL4AccelerationStructureTriangleGeometryDescriptor {
	return getMTL4AccelerationStructureTriangleGeometryDescriptorClass().New()
}




