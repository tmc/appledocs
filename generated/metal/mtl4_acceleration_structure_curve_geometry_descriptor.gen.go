// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTL4AccelerationStructureCurveGeometryDescriptor] class.
var (
	MTL4AccelerationStructureCurveGeometryDescriptorClass     _MTL4AccelerationStructureCurveGeometryDescriptorClass
	MTL4AccelerationStructureCurveGeometryDescriptorClassOnce sync.Once
)

func getMTL4AccelerationStructureCurveGeometryDescriptorClass() _MTL4AccelerationStructureCurveGeometryDescriptorClass {
	MTL4AccelerationStructureCurveGeometryDescriptorClassOnce.Do(func() {
		MTL4AccelerationStructureCurveGeometryDescriptorClass = _MTL4AccelerationStructureCurveGeometryDescriptorClass{objc.GetClass("MTL4AccelerationStructureCurveGeometryDescriptor")}
	})
	return MTL4AccelerationStructureCurveGeometryDescriptorClass
}

type _MTL4AccelerationStructureCurveGeometryDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [MTL4AccelerationStructureCurveGeometryDescriptor] class.
type IMTL4AccelerationStructureCurveGeometryDescriptor interface {
	IMTL4AccelerationStructureGeometryDescriptor
}

// Describes curve geometry suitable for ray tracing.
//
// Use a to mark residency of all buffers this descriptor references when you build this acceleration structure.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureCurveGeometryDescriptor
type MTL4AccelerationStructureCurveGeometryDescriptor struct {
	MTL4AccelerationStructureGeometryDescriptor
}

// MTL4AccelerationStructureCurveGeometryDescriptorFrom constructs a [MTL4AccelerationStructureCurveGeometryDescriptor] from an unsafe.Pointer.
//
// Describes curve geometry suitable for ray tracing.
func MTL4AccelerationStructureCurveGeometryDescriptorFrom(ptr unsafe.Pointer) MTL4AccelerationStructureCurveGeometryDescriptor {
	return MTL4AccelerationStructureCurveGeometryDescriptor{
		MTL4AccelerationStructureGeometryDescriptor: MTL4AccelerationStructureGeometryDescriptorFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTL4AccelerationStructureCurveGeometryDescriptorClass) Alloc() MTL4AccelerationStructureCurveGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureCurveGeometryDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTL4AccelerationStructureCurveGeometryDescriptorClass) New() MTL4AccelerationStructureCurveGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureCurveGeometryDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) Init() MTL4AccelerationStructureCurveGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureCurveGeometryDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTL4AccelerationStructureCurveGeometryDescriptor) Autorelease() MTL4AccelerationStructureCurveGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureCurveGeometryDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4AccelerationStructureCurveGeometryDescriptor creates a new MTL4AccelerationStructureCurveGeometryDescriptor instance.
func NewMTL4AccelerationStructureCurveGeometryDescriptor() MTL4AccelerationStructureCurveGeometryDescriptor {
	return getMTL4AccelerationStructureCurveGeometryDescriptorClass().New()
}




