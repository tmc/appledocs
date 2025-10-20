// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTL4AccelerationStructureMotionTriangleGeometryDescriptor] class.
var (
	MTL4AccelerationStructureMotionTriangleGeometryDescriptorClass     _MTL4AccelerationStructureMotionTriangleGeometryDescriptorClass
	MTL4AccelerationStructureMotionTriangleGeometryDescriptorClassOnce sync.Once
)

func getMTL4AccelerationStructureMotionTriangleGeometryDescriptorClass() _MTL4AccelerationStructureMotionTriangleGeometryDescriptorClass {
	MTL4AccelerationStructureMotionTriangleGeometryDescriptorClassOnce.Do(func() {
		MTL4AccelerationStructureMotionTriangleGeometryDescriptorClass = _MTL4AccelerationStructureMotionTriangleGeometryDescriptorClass{objc.GetClass("MTL4AccelerationStructureMotionTriangleGeometryDescriptor")}
	})
	return MTL4AccelerationStructureMotionTriangleGeometryDescriptorClass
}

type _MTL4AccelerationStructureMotionTriangleGeometryDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [MTL4AccelerationStructureMotionTriangleGeometryDescriptor] class.
type IMTL4AccelerationStructureMotionTriangleGeometryDescriptor interface {
	IMTL4AccelerationStructureGeometryDescriptor
}

// Describes motion triangle geometry, suitable for motion ray tracing.
//
// Use a to mark residency of all buffers this descriptor references when you build this acceleration structure.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionTriangleGeometryDescriptor
type MTL4AccelerationStructureMotionTriangleGeometryDescriptor struct {
	MTL4AccelerationStructureGeometryDescriptor
}

// MTL4AccelerationStructureMotionTriangleGeometryDescriptorFrom constructs a [MTL4AccelerationStructureMotionTriangleGeometryDescriptor] from an unsafe.Pointer.
//
// Describes motion triangle geometry, suitable for motion ray tracing.
func MTL4AccelerationStructureMotionTriangleGeometryDescriptorFrom(ptr unsafe.Pointer) MTL4AccelerationStructureMotionTriangleGeometryDescriptor {
	return MTL4AccelerationStructureMotionTriangleGeometryDescriptor{
		MTL4AccelerationStructureGeometryDescriptor: MTL4AccelerationStructureGeometryDescriptorFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTL4AccelerationStructureMotionTriangleGeometryDescriptorClass) Alloc() MTL4AccelerationStructureMotionTriangleGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureMotionTriangleGeometryDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTL4AccelerationStructureMotionTriangleGeometryDescriptorClass) New() MTL4AccelerationStructureMotionTriangleGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureMotionTriangleGeometryDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTL4AccelerationStructureMotionTriangleGeometryDescriptor) Init() MTL4AccelerationStructureMotionTriangleGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureMotionTriangleGeometryDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTL4AccelerationStructureMotionTriangleGeometryDescriptor) Autorelease() MTL4AccelerationStructureMotionTriangleGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureMotionTriangleGeometryDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4AccelerationStructureMotionTriangleGeometryDescriptor creates a new MTL4AccelerationStructureMotionTriangleGeometryDescriptor instance.
func NewMTL4AccelerationStructureMotionTriangleGeometryDescriptor() MTL4AccelerationStructureMotionTriangleGeometryDescriptor {
	return getMTL4AccelerationStructureMotionTriangleGeometryDescriptorClass().New()
}




