// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTL4AccelerationStructureMotionCurveGeometryDescriptor] class.
var (
	MTL4AccelerationStructureMotionCurveGeometryDescriptorClass     _MTL4AccelerationStructureMotionCurveGeometryDescriptorClass
	MTL4AccelerationStructureMotionCurveGeometryDescriptorClassOnce sync.Once
)

func getMTL4AccelerationStructureMotionCurveGeometryDescriptorClass() _MTL4AccelerationStructureMotionCurveGeometryDescriptorClass {
	MTL4AccelerationStructureMotionCurveGeometryDescriptorClassOnce.Do(func() {
		MTL4AccelerationStructureMotionCurveGeometryDescriptorClass = _MTL4AccelerationStructureMotionCurveGeometryDescriptorClass{objc.GetClass("MTL4AccelerationStructureMotionCurveGeometryDescriptor")}
	})
	return MTL4AccelerationStructureMotionCurveGeometryDescriptorClass
}

type _MTL4AccelerationStructureMotionCurveGeometryDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [MTL4AccelerationStructureMotionCurveGeometryDescriptor] class.
type IMTL4AccelerationStructureMotionCurveGeometryDescriptor interface {
	IMTL4AccelerationStructureGeometryDescriptor
}

// Describes motion curve geometry, suitable for motion ray tracing.
//
// Use a to mark residency of all buffers this descriptor references when you build this acceleration structure.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionCurveGeometryDescriptor
type MTL4AccelerationStructureMotionCurveGeometryDescriptor struct {
	MTL4AccelerationStructureGeometryDescriptor
}

// MTL4AccelerationStructureMotionCurveGeometryDescriptorFrom constructs a [MTL4AccelerationStructureMotionCurveGeometryDescriptor] from an unsafe.Pointer.
//
// Describes motion curve geometry, suitable for motion ray tracing.
func MTL4AccelerationStructureMotionCurveGeometryDescriptorFrom(ptr unsafe.Pointer) MTL4AccelerationStructureMotionCurveGeometryDescriptor {
	return MTL4AccelerationStructureMotionCurveGeometryDescriptor{
		MTL4AccelerationStructureGeometryDescriptor: MTL4AccelerationStructureGeometryDescriptorFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTL4AccelerationStructureMotionCurveGeometryDescriptorClass) Alloc() MTL4AccelerationStructureMotionCurveGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureMotionCurveGeometryDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTL4AccelerationStructureMotionCurveGeometryDescriptorClass) New() MTL4AccelerationStructureMotionCurveGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureMotionCurveGeometryDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTL4AccelerationStructureMotionCurveGeometryDescriptor) Init() MTL4AccelerationStructureMotionCurveGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureMotionCurveGeometryDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTL4AccelerationStructureMotionCurveGeometryDescriptor) Autorelease() MTL4AccelerationStructureMotionCurveGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureMotionCurveGeometryDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4AccelerationStructureMotionCurveGeometryDescriptor creates a new MTL4AccelerationStructureMotionCurveGeometryDescriptor instance.
func NewMTL4AccelerationStructureMotionCurveGeometryDescriptor() MTL4AccelerationStructureMotionCurveGeometryDescriptor {
	return getMTL4AccelerationStructureMotionCurveGeometryDescriptorClass().New()
}




