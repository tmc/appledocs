// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptor] class.
var (
	MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptorClass     _MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptorClass
	MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptorClassOnce sync.Once
)

func getMTL4AccelerationStructureMotionBoundingBoxGeometryDescriptorClass() _MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptorClass {
	MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptorClassOnce.Do(func() {
		MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptorClass = _MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptorClass{objc.GetClass("MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptor")}
	})
	return MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptorClass
}

type _MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptor] class.
type IMTL4AccelerationStructureMotionBoundingBoxGeometryDescriptor interface {
	IMTL4AccelerationStructureGeometryDescriptor
}

// Describes motion bounding box geometry, suitable for motion ray tracing.
//
// You use bounding boxes to implement procedural geometry for ray tracing, such as spheres or any other shape you define by using intersection functions. Use a to mark residency of all buffers this descriptor references when you build this acceleration structure.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptor
type MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptor struct {
	MTL4AccelerationStructureGeometryDescriptor
}

// MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptorFrom constructs a [MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptor] from an unsafe.Pointer.
//
// Describes motion bounding box geometry, suitable for motion ray tracing.
func MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptorFrom(ptr unsafe.Pointer) MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptor {
	return MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptor{
		MTL4AccelerationStructureGeometryDescriptor: MTL4AccelerationStructureGeometryDescriptorFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptorClass) Alloc() MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptorClass) New() MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptor) Init() MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptor) Autorelease() MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4AccelerationStructureMotionBoundingBoxGeometryDescriptor creates a new MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptor instance.
func NewMTL4AccelerationStructureMotionBoundingBoxGeometryDescriptor() MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptor {
	return getMTL4AccelerationStructureMotionBoundingBoxGeometryDescriptorClass().New()
}


// Configures a reference to a buffer where each entry contains a reference to a buffer of bounding boxes.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuremotionboundingboxgeometrydescriptor/boundingboxbuffers
func (m_ MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptor) BoundingBoxBuffers() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("boundingBoxBuffers"))
	return rv
}


// SetBoundingBoxBuffers sets the value of the boundingBoxBuffers property.
// Configures a reference to a buffer where each entry contains a reference to a buffer of bounding boxes.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuremotionboundingboxgeometrydescriptor/boundingboxbuffers
func (m_ MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptor) SetBoundingBoxBuffers(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBoundingBoxBuffers:"), value)
}

// Declares the number of bounding boxes in each buffer that
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuremotionboundingboxgeometrydescriptor/boundingboxcount
func (m_ MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptor) BoundingBoxCount() int {
	rv := objc.Send[int](m_.ID, objc.Sel("boundingBoxCount"))
	return rv
}


// SetBoundingBoxCount sets the value of the boundingBoxCount property.
// Declares the number of bounding boxes in each buffer that

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuremotionboundingboxgeometrydescriptor/boundingboxcount
func (m_ MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptor) SetBoundingBoxCount(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBoundingBoxCount:"), value)
}

// Declares the stride, in bytes, between bounding boxes in the bounding box buffers each entry in
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuremotionboundingboxgeometrydescriptor/boundingboxstride
func (m_ MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptor) BoundingBoxStride() int {
	rv := objc.Send[int](m_.ID, objc.Sel("boundingBoxStride"))
	return rv
}


// SetBoundingBoxStride sets the value of the boundingBoxStride property.
// Declares the stride, in bytes, between bounding boxes in the bounding box buffers each entry in

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuremotionboundingboxgeometrydescriptor/boundingboxstride
func (m_ MTL4AccelerationStructureMotionBoundingBoxGeometryDescriptor) SetBoundingBoxStride(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBoundingBoxStride:"), value)
}



