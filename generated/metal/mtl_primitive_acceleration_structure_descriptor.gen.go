// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
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
	

	// properties:
	GeometryDescriptors() []AccelerationStructureGeometryDescriptor
	SetGeometryDescriptors(value []AccelerationStructureGeometryDescriptor)
	MotionEndBorderMode() MotionBorderMode
	SetMotionEndBorderMode(value MotionBorderMode)
	MotionEndTime() float32
	SetMotionEndTime(value float32)
	MotionKeyframeCount() uint
	SetMotionKeyframeCount(value uint)
	MotionStartBorderMode() MotionBorderMode
	SetMotionStartBorderMode(value MotionBorderMode)
	MotionStartTime() float32
	SetMotionStartTime(value float32)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (pc _PrimitiveAccelerationStructureDescriptorClass) Alloc() PrimitiveAccelerationStructureDescriptor {
	rv := objc.Send[PrimitiveAccelerationStructureDescriptor](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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





// A description of an acceleration structure that contains geometry primitives.
//
// Metal provides acceleration structures with a two-level hierarchy. The bottom layer consists of primitive acceleration structures, which instance acceleration structures in the top level reference.


// A description of an acceleration structure that contains geometry primitives.
//
// [Full Topic]
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










// Creates a new primitive descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPrimitiveAccelerationStructureDescriptor/descriptor
func (pc _PrimitiveAccelerationStructureDescriptorClass) Descriptor() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(pc.class), objc.Sel("descriptor"))
	return rv
}

















// An array that contains the individual pieces of geometry that compose the acceleration structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPrimitiveAccelerationStructureDescriptor/geometryDescriptors
func (p_ PrimitiveAccelerationStructureDescriptor) GeometryDescriptors() []AccelerationStructureGeometryDescriptor {
	rv := objc.Send[[]AccelerationStructureGeometryDescriptor](p_.ID, objc.Sel("geometryDescriptors"))
	return rv
}


// An array that contains the individual pieces of geometry that compose the acceleration structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPrimitiveAccelerationStructureDescriptor/geometryDescriptors
func (p_ PrimitiveAccelerationStructureDescriptor) SetGeometryDescriptors(value []AccelerationStructureGeometryDescriptor) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](p_.ID, objc.Sel("setGeometryDescriptors:"), nsArray)
}


// The mode to use when handling timestamps after the end time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPrimitiveAccelerationStructureDescriptor/motionEndBorderMode
func (p_ PrimitiveAccelerationStructureDescriptor) MotionEndBorderMode() MotionBorderMode {
	rv := objc.Send[MotionBorderMode](p_.ID, objc.Sel("motionEndBorderMode"))
	return rv
}


// The mode to use when handling timestamps after the end time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPrimitiveAccelerationStructureDescriptor/motionEndBorderMode
func (p_ PrimitiveAccelerationStructureDescriptor) SetMotionEndBorderMode(value MotionBorderMode) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMotionEndBorderMode:"), value)
}


// The end time for the range of motion that the keyframe data describes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPrimitiveAccelerationStructureDescriptor/motionEndTime
func (p_ PrimitiveAccelerationStructureDescriptor) MotionEndTime() float32 {
	rv := objc.Send[float32](p_.ID, objc.Sel("motionEndTime"))
	return rv
}


// The end time for the range of motion that the keyframe data describes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPrimitiveAccelerationStructureDescriptor/motionEndTime
func (p_ PrimitiveAccelerationStructureDescriptor) SetMotionEndTime(value float32) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMotionEndTime:"), value)
}


// The number of keyframes in the geometry data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPrimitiveAccelerationStructureDescriptor/motionKeyframeCount
func (p_ PrimitiveAccelerationStructureDescriptor) MotionKeyframeCount() uint {
	rv := objc.Send[uint](p_.ID, objc.Sel("motionKeyframeCount"))
	return rv
}


// The number of keyframes in the geometry data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPrimitiveAccelerationStructureDescriptor/motionKeyframeCount
func (p_ PrimitiveAccelerationStructureDescriptor) SetMotionKeyframeCount(value uint) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMotionKeyframeCount:"), value)
}


// The mode to use when handling timestamps before the start time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPrimitiveAccelerationStructureDescriptor/motionStartBorderMode
func (p_ PrimitiveAccelerationStructureDescriptor) MotionStartBorderMode() MotionBorderMode {
	rv := objc.Send[MotionBorderMode](p_.ID, objc.Sel("motionStartBorderMode"))
	return rv
}


// The mode to use when handling timestamps before the start time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPrimitiveAccelerationStructureDescriptor/motionStartBorderMode
func (p_ PrimitiveAccelerationStructureDescriptor) SetMotionStartBorderMode(value MotionBorderMode) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMotionStartBorderMode:"), value)
}


// The start time for the range of motion that the keyframe data describes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPrimitiveAccelerationStructureDescriptor/motionStartTime
func (p_ PrimitiveAccelerationStructureDescriptor) MotionStartTime() float32 {
	rv := objc.Send[float32](p_.ID, objc.Sel("motionStartTime"))
	return rv
}


// The start time for the range of motion that the keyframe data describes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPrimitiveAccelerationStructureDescriptor/motionStartTime
func (p_ PrimitiveAccelerationStructureDescriptor) SetMotionStartTime(value float32) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMotionStartTime:"), value)
}








