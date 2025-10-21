// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AccelerationStructureMotionBoundingBoxGeometryDescriptor] class.
var (
	AccelerationStructureMotionBoundingBoxGeometryDescriptorClass     _AccelerationStructureMotionBoundingBoxGeometryDescriptorClass
	AccelerationStructureMotionBoundingBoxGeometryDescriptorClassOnce sync.Once
)

func getAccelerationStructureMotionBoundingBoxGeometryDescriptorClass() _AccelerationStructureMotionBoundingBoxGeometryDescriptorClass {
	AccelerationStructureMotionBoundingBoxGeometryDescriptorClassOnce.Do(func() {
		AccelerationStructureMotionBoundingBoxGeometryDescriptorClass = _AccelerationStructureMotionBoundingBoxGeometryDescriptorClass{objc.GetClass("MTLAccelerationStructureMotionBoundingBoxGeometryDescriptor")}
	})
	return AccelerationStructureMotionBoundingBoxGeometryDescriptorClass
}

type _AccelerationStructureMotionBoundingBoxGeometryDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [AccelerationStructureMotionBoundingBoxGeometryDescriptor] class.
type IAccelerationStructureMotionBoundingBoxGeometryDescriptor interface {
	IAccelerationStructureGeometryDescriptor
}

// A description of a list of bounding boxes, as motion keyframe data, to turn into an acceleration structure.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionBoundingBoxGeometryDescriptor
type AccelerationStructureMotionBoundingBoxGeometryDescriptor struct {
	AccelerationStructureGeometryDescriptor
}

// AccelerationStructureMotionBoundingBoxGeometryDescriptorFrom constructs a [AccelerationStructureMotionBoundingBoxGeometryDescriptor] from an unsafe.Pointer.
//
// A description of a list of bounding boxes, as motion keyframe data, to turn into an acceleration structure.
func AccelerationStructureMotionBoundingBoxGeometryDescriptorFrom(ptr unsafe.Pointer) AccelerationStructureMotionBoundingBoxGeometryDescriptor {
	return AccelerationStructureMotionBoundingBoxGeometryDescriptor{
		AccelerationStructureGeometryDescriptor: AccelerationStructureGeometryDescriptorFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AccelerationStructureMotionBoundingBoxGeometryDescriptorClass) Alloc() AccelerationStructureMotionBoundingBoxGeometryDescriptor {
	rv := objc.Send[AccelerationStructureMotionBoundingBoxGeometryDescriptor](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AccelerationStructureMotionBoundingBoxGeometryDescriptorClass) New() AccelerationStructureMotionBoundingBoxGeometryDescriptor {
	rv := objc.Send[AccelerationStructureMotionBoundingBoxGeometryDescriptor](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AccelerationStructureMotionBoundingBoxGeometryDescriptor) Init() AccelerationStructureMotionBoundingBoxGeometryDescriptor {
	rv := objc.Send[AccelerationStructureMotionBoundingBoxGeometryDescriptor](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AccelerationStructureMotionBoundingBoxGeometryDescriptor) Autorelease() AccelerationStructureMotionBoundingBoxGeometryDescriptor {
	rv := objc.Send[AccelerationStructureMotionBoundingBoxGeometryDescriptor](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAccelerationStructureMotionBoundingBoxGeometryDescriptor creates a new AccelerationStructureMotionBoundingBoxGeometryDescriptor instance.
func NewAccelerationStructureMotionBoundingBoxGeometryDescriptor() AccelerationStructureMotionBoundingBoxGeometryDescriptor {
	return getAccelerationStructureMotionBoundingBoxGeometryDescriptorClass().New()
}


// A array of motion keyframes, each containing bounding box data.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotionboundingboxgeometrydescriptor/boundingboxbuffers
func (a_ AccelerationStructureMotionBoundingBoxGeometryDescriptor) BoundingBoxBuffers() MTLMotionKeyframeData {
	rv := objc.Send[MTLMotionKeyframeData](a_.ID, objc.Sel("boundingBoxBuffers"))
	return rv
}


// SetBoundingBoxBuffers sets the value of the boundingBoxBuffers property.
// A array of motion keyframes, each containing bounding box data.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotionboundingboxgeometrydescriptor/boundingboxbuffers
func (a_ AccelerationStructureMotionBoundingBoxGeometryDescriptor) SetBoundingBoxBuffers(value IMTLMotionKeyframeData) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setBoundingBoxBuffers:"), value)
}

// The number of bounding boxes in each bounding box buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotionboundingboxgeometrydescriptor/boundingboxcount
func (a_ AccelerationStructureMotionBoundingBoxGeometryDescriptor) BoundingBoxCount() int {
	rv := objc.Send[int](a_.ID, objc.Sel("boundingBoxCount"))
	return rv
}


// SetBoundingBoxCount sets the value of the boundingBoxCount property.
// The number of bounding boxes in each bounding box buffer.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotionboundingboxgeometrydescriptor/boundingboxcount
func (a_ AccelerationStructureMotionBoundingBoxGeometryDescriptor) SetBoundingBoxCount(value int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setBoundingBoxCount:"), value)
}

// The stride, in bytes, between bounding boxes in each buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotionboundingboxgeometrydescriptor/boundingboxstride
func (a_ AccelerationStructureMotionBoundingBoxGeometryDescriptor) BoundingBoxStride() int {
	rv := objc.Send[int](a_.ID, objc.Sel("boundingBoxStride"))
	return rv
}


// SetBoundingBoxStride sets the value of the boundingBoxStride property.
// The stride, in bytes, between bounding boxes in each buffer.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotionboundingboxgeometrydescriptor/boundingboxstride
func (a_ AccelerationStructureMotionBoundingBoxGeometryDescriptor) SetBoundingBoxStride(value int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setBoundingBoxStride:"), value)
}



