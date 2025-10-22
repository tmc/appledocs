// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTL4AccelerationStructureBoundingBoxGeometryDescriptor] class.
var (
	MTL4AccelerationStructureBoundingBoxGeometryDescriptorClass     _MTL4AccelerationStructureBoundingBoxGeometryDescriptorClass
	MTL4AccelerationStructureBoundingBoxGeometryDescriptorClassOnce sync.Once
)

func getMTL4AccelerationStructureBoundingBoxGeometryDescriptorClass() _MTL4AccelerationStructureBoundingBoxGeometryDescriptorClass {
	MTL4AccelerationStructureBoundingBoxGeometryDescriptorClassOnce.Do(func() {
		MTL4AccelerationStructureBoundingBoxGeometryDescriptorClass = _MTL4AccelerationStructureBoundingBoxGeometryDescriptorClass{objc.GetClass("MTL4AccelerationStructureBoundingBoxGeometryDescriptor")}
	})
	return MTL4AccelerationStructureBoundingBoxGeometryDescriptorClass
}

type _MTL4AccelerationStructureBoundingBoxGeometryDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [MTL4AccelerationStructureBoundingBoxGeometryDescriptor] class.
type IMTL4AccelerationStructureBoundingBoxGeometryDescriptor interface {
	IMTL4AccelerationStructureGeometryDescriptor
	BoundingBoxBuffer() unsafe.Pointer
	SetBoundingBoxBuffer(value unsafe.Pointer)
	BoundingBoxCount() int
	SetBoundingBoxCount(value int)
	BoundingBoxStride() int
	SetBoundingBoxStride(value int)
}

// Describes bounding-box geometry suitable for ray tracing.
//
// You use bounding boxes to implement procedural geometry for ray tracing, such as spheres or any other shape you define by using intersection functions. Use a to mark residency of all buffers this descriptor references when you build this acceleration structure.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureBoundingBoxGeometryDescriptor
type MTL4AccelerationStructureBoundingBoxGeometryDescriptor struct {
	MTL4AccelerationStructureGeometryDescriptor
}

// MTL4AccelerationStructureBoundingBoxGeometryDescriptorFrom constructs a [MTL4AccelerationStructureBoundingBoxGeometryDescriptor] from an unsafe.Pointer.
//
// Describes bounding-box geometry suitable for ray tracing.
func MTL4AccelerationStructureBoundingBoxGeometryDescriptorFrom(ptr unsafe.Pointer) MTL4AccelerationStructureBoundingBoxGeometryDescriptor {
	return MTL4AccelerationStructureBoundingBoxGeometryDescriptor{
		MTL4AccelerationStructureGeometryDescriptor: MTL4AccelerationStructureGeometryDescriptorFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTL4AccelerationStructureBoundingBoxGeometryDescriptorClass) Alloc() MTL4AccelerationStructureBoundingBoxGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureBoundingBoxGeometryDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTL4AccelerationStructureBoundingBoxGeometryDescriptorClass) New() MTL4AccelerationStructureBoundingBoxGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureBoundingBoxGeometryDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTL4AccelerationStructureBoundingBoxGeometryDescriptor) Init() MTL4AccelerationStructureBoundingBoxGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureBoundingBoxGeometryDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTL4AccelerationStructureBoundingBoxGeometryDescriptor) Autorelease() MTL4AccelerationStructureBoundingBoxGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureBoundingBoxGeometryDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4AccelerationStructureBoundingBoxGeometryDescriptor creates a new MTL4AccelerationStructureBoundingBoxGeometryDescriptor instance.
func NewMTL4AccelerationStructureBoundingBoxGeometryDescriptor() MTL4AccelerationStructureBoundingBoxGeometryDescriptor {
	return getMTL4AccelerationStructureBoundingBoxGeometryDescriptorClass().New()
}


// References a buffer containing bounding box data in
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructureboundingboxgeometrydescriptor/boundingboxbuffer
func (m_ MTL4AccelerationStructureBoundingBoxGeometryDescriptor) BoundingBoxBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("boundingBoxBuffer"))
	return rv
}


// SetBoundingBoxBuffer sets the value of the boundingBoxBuffer property.
// References a buffer containing bounding box data in

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructureboundingboxgeometrydescriptor/boundingboxbuffer
func (m_ MTL4AccelerationStructureBoundingBoxGeometryDescriptor) SetBoundingBoxBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBoundingBoxBuffer:"), value)
}

// Describes the number of bounding boxes the
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructureboundingboxgeometrydescriptor/boundingboxcount
func (m_ MTL4AccelerationStructureBoundingBoxGeometryDescriptor) BoundingBoxCount() int {
	rv := objc.Send[int](m_.ID, objc.Sel("boundingBoxCount"))
	return rv
}


// SetBoundingBoxCount sets the value of the boundingBoxCount property.
// Describes the number of bounding boxes the

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructureboundingboxgeometrydescriptor/boundingboxcount
func (m_ MTL4AccelerationStructureBoundingBoxGeometryDescriptor) SetBoundingBoxCount(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBoundingBoxCount:"), value)
}

// Assigns the stride, in bytes, between bounding boxes in the bounding box buffer
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructureboundingboxgeometrydescriptor/boundingboxstride
func (m_ MTL4AccelerationStructureBoundingBoxGeometryDescriptor) BoundingBoxStride() int {
	rv := objc.Send[int](m_.ID, objc.Sel("boundingBoxStride"))
	return rv
}


// SetBoundingBoxStride sets the value of the boundingBoxStride property.
// Assigns the stride, in bytes, between bounding boxes in the bounding box buffer

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructureboundingboxgeometrydescriptor/boundingboxstride
func (m_ MTL4AccelerationStructureBoundingBoxGeometryDescriptor) SetBoundingBoxStride(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBoundingBoxStride:"), value)
}



