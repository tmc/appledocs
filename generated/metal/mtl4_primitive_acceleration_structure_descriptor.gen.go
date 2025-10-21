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


// Configures the motion start time for this geometry.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4primitiveaccelerationstructuredescriptor/motionstarttime
func (m_ MTL4PrimitiveAccelerationStructureDescriptor) MotionStartTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("motionStartTime"))
	return rv
}


// SetMotionStartTime sets the value of the motionStartTime property.
// Configures the motion start time for this geometry.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4primitiveaccelerationstructuredescriptor/motionstarttime
func (m_ MTL4PrimitiveAccelerationStructureDescriptor) SetMotionStartTime(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMotionStartTime:"), value)
}

// Configures the motion end time for this geometry.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4primitiveaccelerationstructuredescriptor/motionendtime
func (m_ MTL4PrimitiveAccelerationStructureDescriptor) MotionEndTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("motionEndTime"))
	return rv
}


// SetMotionEndTime sets the value of the motionEndTime property.
// Configures the motion end time for this geometry.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4primitiveaccelerationstructuredescriptor/motionendtime
func (m_ MTL4PrimitiveAccelerationStructureDescriptor) SetMotionEndTime(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMotionEndTime:"), value)
}

// Associates the array of geometry descriptors that comprise this primitive acceleration structure.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4primitiveaccelerationstructuredescriptor/geometrydescriptors
func (m_ MTL4PrimitiveAccelerationStructureDescriptor) GeometryDescriptors() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("geometryDescriptors"))
	return rv
}


// SetGeometryDescriptors sets the value of the geometryDescriptors property.
// Associates the array of geometry descriptors that comprise this primitive acceleration structure.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4primitiveaccelerationstructuredescriptor/geometrydescriptors
func (m_ MTL4PrimitiveAccelerationStructureDescriptor) SetGeometryDescriptors(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGeometryDescriptors:"), value)
}

// Configures the motion border mode.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4primitiveaccelerationstructuredescriptor/motionendbordermode
func (m_ MTL4PrimitiveAccelerationStructureDescriptor) MotionEndBorderMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("motionEndBorderMode"))
	return rv
}


// SetMotionEndBorderMode sets the value of the motionEndBorderMode property.
// Configures the motion border mode.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4primitiveaccelerationstructuredescriptor/motionendbordermode
func (m_ MTL4PrimitiveAccelerationStructureDescriptor) SetMotionEndBorderMode(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMotionEndBorderMode:"), value)
}

// Configures the behavior when the ray-tracing system samples the acceleration structure before the motion start time.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4primitiveaccelerationstructuredescriptor/motionstartbordermode
func (m_ MTL4PrimitiveAccelerationStructureDescriptor) MotionStartBorderMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("motionStartBorderMode"))
	return rv
}


// SetMotionStartBorderMode sets the value of the motionStartBorderMode property.
// Configures the behavior when the ray-tracing system samples the acceleration structure before the motion start time.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4primitiveaccelerationstructuredescriptor/motionstartbordermode
func (m_ MTL4PrimitiveAccelerationStructureDescriptor) SetMotionStartBorderMode(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMotionStartBorderMode:"), value)
}

// Sets the motion keyframe count.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4primitiveaccelerationstructuredescriptor/motionkeyframecount
func (m_ MTL4PrimitiveAccelerationStructureDescriptor) MotionKeyframeCount() int {
	rv := objc.Send[int](m_.ID, objc.Sel("motionKeyframeCount"))
	return rv
}


// SetMotionKeyframeCount sets the value of the motionKeyframeCount property.
// Sets the motion keyframe count.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4primitiveaccelerationstructuredescriptor/motionkeyframecount
func (m_ MTL4PrimitiveAccelerationStructureDescriptor) SetMotionKeyframeCount(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMotionKeyframeCount:"), value)
}



