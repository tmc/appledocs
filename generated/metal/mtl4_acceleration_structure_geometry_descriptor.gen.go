// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTL4AccelerationStructureGeometryDescriptor] class.
var (
	MTL4AccelerationStructureGeometryDescriptorClass     _MTL4AccelerationStructureGeometryDescriptorClass
	MTL4AccelerationStructureGeometryDescriptorClassOnce sync.Once
)

func getMTL4AccelerationStructureGeometryDescriptorClass() _MTL4AccelerationStructureGeometryDescriptorClass {
	MTL4AccelerationStructureGeometryDescriptorClassOnce.Do(func() {
		MTL4AccelerationStructureGeometryDescriptorClass = _MTL4AccelerationStructureGeometryDescriptorClass{objc.GetClass("MTL4AccelerationStructureGeometryDescriptor")}
	})
	return MTL4AccelerationStructureGeometryDescriptorClass
}

type _MTL4AccelerationStructureGeometryDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [MTL4AccelerationStructureGeometryDescriptor] class.
type IMTL4AccelerationStructureGeometryDescriptor interface {
	objectivec.IObject
}

// Base class for all Metal 4 acceleration structure geometry descriptors.
//
// Don’t use this class directly. Use one of the derived classes instead.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureGeometryDescriptor
type MTL4AccelerationStructureGeometryDescriptor struct {
	objectivec.Object
}

// MTL4AccelerationStructureGeometryDescriptorFrom constructs a [MTL4AccelerationStructureGeometryDescriptor] from an unsafe.Pointer.
//
// Base class for all Metal 4 acceleration structure geometry descriptors.
func MTL4AccelerationStructureGeometryDescriptorFrom(ptr unsafe.Pointer) MTL4AccelerationStructureGeometryDescriptor {
	return MTL4AccelerationStructureGeometryDescriptor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTL4AccelerationStructureGeometryDescriptorClass) Alloc() MTL4AccelerationStructureGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureGeometryDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTL4AccelerationStructureGeometryDescriptorClass) New() MTL4AccelerationStructureGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureGeometryDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTL4AccelerationStructureGeometryDescriptor) Init() MTL4AccelerationStructureGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureGeometryDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTL4AccelerationStructureGeometryDescriptor) Autorelease() MTL4AccelerationStructureGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureGeometryDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4AccelerationStructureGeometryDescriptor creates a new MTL4AccelerationStructureGeometryDescriptor instance.
func NewMTL4AccelerationStructureGeometryDescriptor() MTL4AccelerationStructureGeometryDescriptor {
	return getMTL4AccelerationStructureGeometryDescriptorClass().New()
}




