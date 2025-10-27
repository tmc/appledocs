// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [IndirectInstanceAccelerationStructureDescriptor] class.
var (
	IndirectInstanceAccelerationStructureDescriptorClass     _IndirectInstanceAccelerationStructureDescriptorClass
	IndirectInstanceAccelerationStructureDescriptorClassOnce sync.Once
)

func getIndirectInstanceAccelerationStructureDescriptorClass() _IndirectInstanceAccelerationStructureDescriptorClass {
	IndirectInstanceAccelerationStructureDescriptorClassOnce.Do(func() {
		IndirectInstanceAccelerationStructureDescriptorClass = _IndirectInstanceAccelerationStructureDescriptorClass{objc.GetClass("MTLIndirectInstanceAccelerationStructureDescriptor")}
	})
	return IndirectInstanceAccelerationStructureDescriptorClass
}

type _IndirectInstanceAccelerationStructureDescriptorClass struct {
	class objc.Class
}





// An interface definition for the [IndirectInstanceAccelerationStructureDescriptor] class.
type IIndirectInstanceAccelerationStructureDescriptor interface {
	IAccelerationStructureDescriptor
	

	// properties:
	InstanceCountBuffer() unsafe.Pointer
	SetInstanceCountBuffer(value unsafe.Pointer)
	InstanceCountBufferOffset() uint
	SetInstanceCountBufferOffset(value uint)
	InstanceDescriptorBuffer() unsafe.Pointer
	SetInstanceDescriptorBuffer(value unsafe.Pointer)
	InstanceDescriptorBufferOffset() uint
	SetInstanceDescriptorBufferOffset(value uint)
	InstanceDescriptorStride() uint
	SetInstanceDescriptorStride(value uint)
	InstanceDescriptorType() AccelerationStructureInstanceDescriptorType
	SetInstanceDescriptorType(value AccelerationStructureInstanceDescriptorType)
	InstanceTransformationMatrixLayout() MatrixLayout
	SetInstanceTransformationMatrixLayout(value MatrixLayout)
	MaxInstanceCount() uint
	SetMaxInstanceCount(value uint)
	MaxMotionTransformCount() uint
	SetMaxMotionTransformCount(value uint)
	MotionTransformBuffer() unsafe.Pointer
	SetMotionTransformBuffer(value unsafe.Pointer)
	MotionTransformBufferOffset() uint
	SetMotionTransformBufferOffset(value uint)
	MotionTransformCountBuffer() unsafe.Pointer
	SetMotionTransformCountBuffer(value unsafe.Pointer)
	MotionTransformCountBufferOffset() uint
	SetMotionTransformCountBufferOffset(value uint)
	MotionTransformStride() uint
	SetMotionTransformStride(value uint)
	MotionTransformType() TransformType
	SetMotionTransformType(value TransformType)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (ic _IndirectInstanceAccelerationStructureDescriptorClass) Alloc() IndirectInstanceAccelerationStructureDescriptor {
	rv := objc.Send[IndirectInstanceAccelerationStructureDescriptor](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _IndirectInstanceAccelerationStructureDescriptorClass) New() IndirectInstanceAccelerationStructureDescriptor {
	rv := objc.Send[IndirectInstanceAccelerationStructureDescriptor](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ IndirectInstanceAccelerationStructureDescriptor) Init() IndirectInstanceAccelerationStructureDescriptor {
	rv := objc.Send[IndirectInstanceAccelerationStructureDescriptor](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ IndirectInstanceAccelerationStructureDescriptor) Autorelease() IndirectInstanceAccelerationStructureDescriptor {
	rv := objc.Send[IndirectInstanceAccelerationStructureDescriptor](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewIndirectInstanceAccelerationStructureDescriptor creates a new IndirectInstanceAccelerationStructureDescriptor instance.
func NewIndirectInstanceAccelerationStructureDescriptor() IndirectInstanceAccelerationStructureDescriptor {
	return getIndirectInstanceAccelerationStructureDescriptorClass().New()
}





// A description of an acceleration structure that Metal derives from instances of primitive acceleration structures that the GPU can populate.


// A description of an acceleration structure that Metal derives from instances of primitive acceleration structures that the GPU can populate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectInstanceAccelerationStructureDescriptor
type IndirectInstanceAccelerationStructureDescriptor struct {
	AccelerationStructureDescriptor
}

// IndirectInstanceAccelerationStructureDescriptorFrom constructs a [IndirectInstanceAccelerationStructureDescriptor] from an unsafe.Pointer.
//
// A description of an acceleration structure that Metal derives from instances of primitive acceleration structures that the GPU can populate.
func IndirectInstanceAccelerationStructureDescriptorFrom(ptr unsafe.Pointer) IndirectInstanceAccelerationStructureDescriptor {
	return IndirectInstanceAccelerationStructureDescriptor{
		AccelerationStructureDescriptor: AccelerationStructureDescriptorFrom(ptr),
	}
}










// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectInstanceAccelerationStructureDescriptor/descriptor
func (ic _IndirectInstanceAccelerationStructureDescriptorClass) Descriptor() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ic.class), objc.Sel("descriptor"))
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectInstanceAccelerationStructureDescriptor/instanceCountBuffer
func (i_ IndirectInstanceAccelerationStructureDescriptor) InstanceCountBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("instanceCountBuffer"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectInstanceAccelerationStructureDescriptor/instanceCountBuffer
func (i_ IndirectInstanceAccelerationStructureDescriptor) SetInstanceCountBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setInstanceCountBuffer:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectInstanceAccelerationStructureDescriptor/instanceCountBufferOffset
func (i_ IndirectInstanceAccelerationStructureDescriptor) InstanceCountBufferOffset() uint {
	rv := objc.Send[uint](i_.ID, objc.Sel("instanceCountBufferOffset"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectInstanceAccelerationStructureDescriptor/instanceCountBufferOffset
func (i_ IndirectInstanceAccelerationStructureDescriptor) SetInstanceCountBufferOffset(value uint) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setInstanceCountBufferOffset:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectInstanceAccelerationStructureDescriptor/instanceDescriptorBuffer
func (i_ IndirectInstanceAccelerationStructureDescriptor) InstanceDescriptorBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("instanceDescriptorBuffer"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectInstanceAccelerationStructureDescriptor/instanceDescriptorBuffer
func (i_ IndirectInstanceAccelerationStructureDescriptor) SetInstanceDescriptorBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setInstanceDescriptorBuffer:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectInstanceAccelerationStructureDescriptor/instanceDescriptorBufferOffset
func (i_ IndirectInstanceAccelerationStructureDescriptor) InstanceDescriptorBufferOffset() uint {
	rv := objc.Send[uint](i_.ID, objc.Sel("instanceDescriptorBufferOffset"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectInstanceAccelerationStructureDescriptor/instanceDescriptorBufferOffset
func (i_ IndirectInstanceAccelerationStructureDescriptor) SetInstanceDescriptorBufferOffset(value uint) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setInstanceDescriptorBufferOffset:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectInstanceAccelerationStructureDescriptor/instanceDescriptorStride
func (i_ IndirectInstanceAccelerationStructureDescriptor) InstanceDescriptorStride() uint {
	rv := objc.Send[uint](i_.ID, objc.Sel("instanceDescriptorStride"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectInstanceAccelerationStructureDescriptor/instanceDescriptorStride
func (i_ IndirectInstanceAccelerationStructureDescriptor) SetInstanceDescriptorStride(value uint) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setInstanceDescriptorStride:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectInstanceAccelerationStructureDescriptor/instanceDescriptorType
func (i_ IndirectInstanceAccelerationStructureDescriptor) InstanceDescriptorType() AccelerationStructureInstanceDescriptorType {
	rv := objc.Send[AccelerationStructureInstanceDescriptorType](i_.ID, objc.Sel("instanceDescriptorType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectInstanceAccelerationStructureDescriptor/instanceDescriptorType
func (i_ IndirectInstanceAccelerationStructureDescriptor) SetInstanceDescriptorType(value AccelerationStructureInstanceDescriptorType) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setInstanceDescriptorType:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectInstanceAccelerationStructureDescriptor/instanceTransformationMatrixLayout
func (i_ IndirectInstanceAccelerationStructureDescriptor) InstanceTransformationMatrixLayout() MatrixLayout {
	rv := objc.Send[MatrixLayout](i_.ID, objc.Sel("instanceTransformationMatrixLayout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectInstanceAccelerationStructureDescriptor/instanceTransformationMatrixLayout
func (i_ IndirectInstanceAccelerationStructureDescriptor) SetInstanceTransformationMatrixLayout(value MatrixLayout) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setInstanceTransformationMatrixLayout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectInstanceAccelerationStructureDescriptor/maxInstanceCount
func (i_ IndirectInstanceAccelerationStructureDescriptor) MaxInstanceCount() uint {
	rv := objc.Send[uint](i_.ID, objc.Sel("maxInstanceCount"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectInstanceAccelerationStructureDescriptor/maxInstanceCount
func (i_ IndirectInstanceAccelerationStructureDescriptor) SetMaxInstanceCount(value uint) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMaxInstanceCount:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectInstanceAccelerationStructureDescriptor/maxMotionTransformCount
func (i_ IndirectInstanceAccelerationStructureDescriptor) MaxMotionTransformCount() uint {
	rv := objc.Send[uint](i_.ID, objc.Sel("maxMotionTransformCount"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectInstanceAccelerationStructureDescriptor/maxMotionTransformCount
func (i_ IndirectInstanceAccelerationStructureDescriptor) SetMaxMotionTransformCount(value uint) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMaxMotionTransformCount:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectInstanceAccelerationStructureDescriptor/motionTransformBuffer
func (i_ IndirectInstanceAccelerationStructureDescriptor) MotionTransformBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("motionTransformBuffer"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectInstanceAccelerationStructureDescriptor/motionTransformBuffer
func (i_ IndirectInstanceAccelerationStructureDescriptor) SetMotionTransformBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMotionTransformBuffer:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectInstanceAccelerationStructureDescriptor/motionTransformBufferOffset
func (i_ IndirectInstanceAccelerationStructureDescriptor) MotionTransformBufferOffset() uint {
	rv := objc.Send[uint](i_.ID, objc.Sel("motionTransformBufferOffset"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectInstanceAccelerationStructureDescriptor/motionTransformBufferOffset
func (i_ IndirectInstanceAccelerationStructureDescriptor) SetMotionTransformBufferOffset(value uint) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMotionTransformBufferOffset:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectInstanceAccelerationStructureDescriptor/motionTransformCountBuffer
func (i_ IndirectInstanceAccelerationStructureDescriptor) MotionTransformCountBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("motionTransformCountBuffer"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectInstanceAccelerationStructureDescriptor/motionTransformCountBuffer
func (i_ IndirectInstanceAccelerationStructureDescriptor) SetMotionTransformCountBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMotionTransformCountBuffer:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectInstanceAccelerationStructureDescriptor/motionTransformCountBufferOffset
func (i_ IndirectInstanceAccelerationStructureDescriptor) MotionTransformCountBufferOffset() uint {
	rv := objc.Send[uint](i_.ID, objc.Sel("motionTransformCountBufferOffset"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectInstanceAccelerationStructureDescriptor/motionTransformCountBufferOffset
func (i_ IndirectInstanceAccelerationStructureDescriptor) SetMotionTransformCountBufferOffset(value uint) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMotionTransformCountBufferOffset:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectInstanceAccelerationStructureDescriptor/motionTransformStride
func (i_ IndirectInstanceAccelerationStructureDescriptor) MotionTransformStride() uint {
	rv := objc.Send[uint](i_.ID, objc.Sel("motionTransformStride"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectInstanceAccelerationStructureDescriptor/motionTransformStride
func (i_ IndirectInstanceAccelerationStructureDescriptor) SetMotionTransformStride(value uint) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMotionTransformStride:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectInstanceAccelerationStructureDescriptor/motionTransformType
func (i_ IndirectInstanceAccelerationStructureDescriptor) MotionTransformType() TransformType {
	rv := objc.Send[TransformType](i_.ID, objc.Sel("motionTransformType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectInstanceAccelerationStructureDescriptor/motionTransformType
func (i_ IndirectInstanceAccelerationStructureDescriptor) SetMotionTransformType(value TransformType) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMotionTransformType:"), value)
}








