// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PrimitiveAccelerationStructureDescriptor] class.
var (
	PrimitiveAccelerationStructureDescriptorClass     _PrimitiveAccelerationStructureDescriptorClass
	PrimitiveAccelerationStructureDescriptorClassOnce sync.Once
)

func getPrimitiveAccelerationStructureDescriptorClass() _PrimitiveAccelerationStructureDescriptorClass {
	PrimitiveAccelerationStructureDescriptorClassOnce.Do(func() {
		PrimitiveAccelerationStructureDescriptorClass = _PrimitiveAccelerationStructureDescriptorClass{objc.GetClass("MTLPrimitiveAccelerationStructureDescriptor")}
	})
	return PrimitiveAccelerationStructureDescriptorClass
}

type _PrimitiveAccelerationStructureDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [PrimitiveAccelerationStructureDescriptor] class.
type IPrimitiveAccelerationStructureDescriptor interface {
	IAccelerationStructureDescriptor
}

// A description of an acceleration structure that contains geometry primitives.
//
// Metal provides acceleration structures with a two-level hierarchy. The bottom layer consists of primitive acceleration structures, which instance acceleration structures in the top level reference.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPrimitiveAccelerationStructureDescriptor
type PrimitiveAccelerationStructureDescriptor struct {
	AccelerationStructureDescriptor
}

// PrimitiveAccelerationStructureDescriptorFrom constructs a [PrimitiveAccelerationStructureDescriptor] from an unsafe.Pointer.
//
// A description of an acceleration structure that contains geometry primitives.
func PrimitiveAccelerationStructureDescriptorFrom(ptr unsafe.Pointer) PrimitiveAccelerationStructureDescriptor {
	return PrimitiveAccelerationStructureDescriptor{
		AccelerationStructureDescriptor: AccelerationStructureDescriptorFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PrimitiveAccelerationStructureDescriptorClass) Alloc() PrimitiveAccelerationStructureDescriptor {
	rv := objc.Send[PrimitiveAccelerationStructureDescriptor](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PrimitiveAccelerationStructureDescriptorClass) New() PrimitiveAccelerationStructureDescriptor {
	rv := objc.Send[PrimitiveAccelerationStructureDescriptor](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PrimitiveAccelerationStructureDescriptor) Init() PrimitiveAccelerationStructureDescriptor {
	rv := objc.Send[PrimitiveAccelerationStructureDescriptor](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PrimitiveAccelerationStructureDescriptor) Autorelease() PrimitiveAccelerationStructureDescriptor {
	rv := objc.Send[PrimitiveAccelerationStructureDescriptor](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPrimitiveAccelerationStructureDescriptor creates a new PrimitiveAccelerationStructureDescriptor instance.
func NewPrimitiveAccelerationStructureDescriptor() PrimitiveAccelerationStructureDescriptor {
	return getPrimitiveAccelerationStructureDescriptorClass().New()
}


// An array that contains the individual pieces of geometry that compose the acceleration structure.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlprimitiveaccelerationstructuredescriptor/geometrydescriptors
func (p_ PrimitiveAccelerationStructureDescriptor) GeometryDescriptors() MTLAccelerationStructureGeometryDescriptor {
	rv := objc.Send[MTLAccelerationStructureGeometryDescriptor](p_.ID, objc.Sel("geometryDescriptors"))
	return rv
}


// SetGeometryDescriptors sets the value of the geometryDescriptors property.
// An array that contains the individual pieces of geometry that compose the acceleration structure.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlprimitiveaccelerationstructuredescriptor/geometrydescriptors
func (p_ PrimitiveAccelerationStructureDescriptor) SetGeometryDescriptors(value IMTLAccelerationStructureGeometryDescriptor) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setGeometryDescriptors:"), value)
}

// The mode to use when handling timestamps after the end time.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlprimitiveaccelerationstructuredescriptor/motionendbordermode
func (p_ PrimitiveAccelerationStructureDescriptor) MotionEndBorderMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("motionEndBorderMode"))
	return rv
}


// SetMotionEndBorderMode sets the value of the motionEndBorderMode property.
// The mode to use when handling timestamps after the end time.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlprimitiveaccelerationstructuredescriptor/motionendbordermode
func (p_ PrimitiveAccelerationStructureDescriptor) SetMotionEndBorderMode(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMotionEndBorderMode:"), value)
}

// The end time for the range of motion that the keyframe data describes.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlprimitiveaccelerationstructuredescriptor/motionendtime
func (p_ PrimitiveAccelerationStructureDescriptor) MotionEndTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("motionEndTime"))
	return rv
}


// SetMotionEndTime sets the value of the motionEndTime property.
// The end time for the range of motion that the keyframe data describes.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlprimitiveaccelerationstructuredescriptor/motionendtime
func (p_ PrimitiveAccelerationStructureDescriptor) SetMotionEndTime(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMotionEndTime:"), value)
}

// The number of keyframes in the geometry data.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlprimitiveaccelerationstructuredescriptor/motionkeyframecount
func (p_ PrimitiveAccelerationStructureDescriptor) MotionKeyframeCount() int {
	rv := objc.Send[int](p_.ID, objc.Sel("motionKeyframeCount"))
	return rv
}


// SetMotionKeyframeCount sets the value of the motionKeyframeCount property.
// The number of keyframes in the geometry data.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlprimitiveaccelerationstructuredescriptor/motionkeyframecount
func (p_ PrimitiveAccelerationStructureDescriptor) SetMotionKeyframeCount(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMotionKeyframeCount:"), value)
}

// The mode to use when handling timestamps before the start time.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlprimitiveaccelerationstructuredescriptor/motionstartbordermode
func (p_ PrimitiveAccelerationStructureDescriptor) MotionStartBorderMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("motionStartBorderMode"))
	return rv
}


// SetMotionStartBorderMode sets the value of the motionStartBorderMode property.
// The mode to use when handling timestamps before the start time.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlprimitiveaccelerationstructuredescriptor/motionstartbordermode
func (p_ PrimitiveAccelerationStructureDescriptor) SetMotionStartBorderMode(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMotionStartBorderMode:"), value)
}

// The start time for the range of motion that the keyframe data describes.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlprimitiveaccelerationstructuredescriptor/motionstarttime
func (p_ PrimitiveAccelerationStructureDescriptor) MotionStartTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("motionStartTime"))
	return rv
}


// SetMotionStartTime sets the value of the motionStartTime property.
// The start time for the range of motion that the keyframe data describes.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlprimitiveaccelerationstructuredescriptor/motionstarttime
func (p_ PrimitiveAccelerationStructureDescriptor) SetMotionStartTime(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMotionStartTime:"), value)
}



