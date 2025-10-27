// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [MTL4IndirectInstanceAccelerationStructureDescriptor] class.
var (
	MTL4IndirectInstanceAccelerationStructureDescriptorClass     _MTL4IndirectInstanceAccelerationStructureDescriptorClass
	MTL4IndirectInstanceAccelerationStructureDescriptorClassOnce sync.Once
)

func getMTL4IndirectInstanceAccelerationStructureDescriptorClass() _MTL4IndirectInstanceAccelerationStructureDescriptorClass {
	MTL4IndirectInstanceAccelerationStructureDescriptorClassOnce.Do(func() {
		MTL4IndirectInstanceAccelerationStructureDescriptorClass = _MTL4IndirectInstanceAccelerationStructureDescriptorClass{objc.GetClass("MTL4IndirectInstanceAccelerationStructureDescriptor")}
	})
	return MTL4IndirectInstanceAccelerationStructureDescriptorClass
}

type _MTL4IndirectInstanceAccelerationStructureDescriptorClass struct {
	class objc.Class
}





// An interface definition for the [MTL4IndirectInstanceAccelerationStructureDescriptor] class.
type IMTL4IndirectInstanceAccelerationStructureDescriptor interface {
	IMTL4AccelerationStructureDescriptor
	

	// properties:
	InstanceCountBuffer() MTL4BufferRange
	SetInstanceCountBuffer(value MTL4BufferRange)
	InstanceDescriptorBuffer() MTL4BufferRange
	SetInstanceDescriptorBuffer(value MTL4BufferRange)
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
	MotionTransformBuffer() MTL4BufferRange
	SetMotionTransformBuffer(value MTL4BufferRange)
	MotionTransformCountBuffer() MTL4BufferRange
	SetMotionTransformCountBuffer(value MTL4BufferRange)
	MotionTransformStride() uint
	SetMotionTransformStride(value uint)
	MotionTransformType() TransformType
	SetMotionTransformType(value TransformType)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (mc _MTL4IndirectInstanceAccelerationStructureDescriptorClass) Alloc() MTL4IndirectInstanceAccelerationStructureDescriptor {
	rv := objc.Send[MTL4IndirectInstanceAccelerationStructureDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTL4IndirectInstanceAccelerationStructureDescriptorClass) New() MTL4IndirectInstanceAccelerationStructureDescriptor {
	rv := objc.Send[MTL4IndirectInstanceAccelerationStructureDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTL4IndirectInstanceAccelerationStructureDescriptor) Init() MTL4IndirectInstanceAccelerationStructureDescriptor {
	rv := objc.Send[MTL4IndirectInstanceAccelerationStructureDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTL4IndirectInstanceAccelerationStructureDescriptor) Autorelease() MTL4IndirectInstanceAccelerationStructureDescriptor {
	rv := objc.Send[MTL4IndirectInstanceAccelerationStructureDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4IndirectInstanceAccelerationStructureDescriptor creates a new MTL4IndirectInstanceAccelerationStructureDescriptor instance.
func NewMTL4IndirectInstanceAccelerationStructureDescriptor() MTL4IndirectInstanceAccelerationStructureDescriptor {
	return getMTL4IndirectInstanceAccelerationStructureDescriptorClass().New()
}





// Descriptor for an “indirect” instance acceleration structure that allows providing the instance count and motion transform count indirectly, through buffer references.
//
// An instance acceleration structure references other acceleration structures, and provides the ability to “instantiate” them multiple times, each one with potentially a different transformation matrix. You specify the properties of the instances in the acceleration structure this descriptor builds by providing a buffer of via its property. Compared to , this descriptor allows you to provide the number of instances it references indirectly through a buffer reference, as well as the number of motion transforms. This enables you to determine these counts indirectly in the GPU timeline via a compute pipeline. Metal needs only to know the maximum possible number of instances and motion transforms to support, which you specify via the and properties. Use a to mark residency of all buffers and acceleration structures this descriptor references when you build this acceleration structure.


// Descriptor for an “indirect” instance acceleration structure that allows providing the instance count and motion transform count indirectly, through buffer references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4IndirectInstanceAccelerationStructureDescriptor
type MTL4IndirectInstanceAccelerationStructureDescriptor struct {
	MTL4AccelerationStructureDescriptor
}

// MTL4IndirectInstanceAccelerationStructureDescriptorFrom constructs a [MTL4IndirectInstanceAccelerationStructureDescriptor] from an unsafe.Pointer.
//
// Descriptor for an “indirect” instance acceleration structure that allows providing the instance count and motion transform count indirectly, through buffer references.
func MTL4IndirectInstanceAccelerationStructureDescriptorFrom(ptr unsafe.Pointer) MTL4IndirectInstanceAccelerationStructureDescriptor {
	return MTL4IndirectInstanceAccelerationStructureDescriptor{
		MTL4AccelerationStructureDescriptor: MTL4AccelerationStructureDescriptorFrom(ptr),
	}
}

























// Provides a reference to a buffer containing the number of instances in the instance descriptor buffer, formatted as a 32-bit unsigned integer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4IndirectInstanceAccelerationStructureDescriptor/instanceCountBuffer
func (m_ MTL4IndirectInstanceAccelerationStructureDescriptor) InstanceCountBuffer() MTL4BufferRange {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("instanceCountBuffer"))
	return rv
}


// Provides a reference to a buffer containing the number of instances in the instance descriptor buffer, formatted as a 32-bit unsigned integer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4IndirectInstanceAccelerationStructureDescriptor/instanceCountBuffer
func (m_ MTL4IndirectInstanceAccelerationStructureDescriptor) SetInstanceCountBuffer(value MTL4BufferRange) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInstanceCountBuffer:"), value)
}


// Assigns a reference to a buffer containing instance descriptors for acceleration structures to reference.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4IndirectInstanceAccelerationStructureDescriptor/instanceDescriptorBuffer
func (m_ MTL4IndirectInstanceAccelerationStructureDescriptor) InstanceDescriptorBuffer() MTL4BufferRange {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("instanceDescriptorBuffer"))
	return rv
}


// Assigns a reference to a buffer containing instance descriptors for acceleration structures to reference.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4IndirectInstanceAccelerationStructureDescriptor/instanceDescriptorBuffer
func (m_ MTL4IndirectInstanceAccelerationStructureDescriptor) SetInstanceDescriptorBuffer(value MTL4BufferRange) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInstanceDescriptorBuffer:"), value)
}


// Sets the stride, in bytes, between instance descriptors in the instance descriptor buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4IndirectInstanceAccelerationStructureDescriptor/instanceDescriptorStride
func (m_ MTL4IndirectInstanceAccelerationStructureDescriptor) InstanceDescriptorStride() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("instanceDescriptorStride"))
	return rv
}


// Sets the stride, in bytes, between instance descriptors in the instance descriptor buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4IndirectInstanceAccelerationStructureDescriptor/instanceDescriptorStride
func (m_ MTL4IndirectInstanceAccelerationStructureDescriptor) SetInstanceDescriptorStride(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInstanceDescriptorStride:"), value)
}


// Controls the type of instance descriptor that the instance descriptor buffer references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4IndirectInstanceAccelerationStructureDescriptor/instanceDescriptorType
func (m_ MTL4IndirectInstanceAccelerationStructureDescriptor) InstanceDescriptorType() AccelerationStructureInstanceDescriptorType {
	rv := objc.Send[AccelerationStructureInstanceDescriptorType](m_.ID, objc.Sel("instanceDescriptorType"))
	return rv
}


// Controls the type of instance descriptor that the instance descriptor buffer references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4IndirectInstanceAccelerationStructureDescriptor/instanceDescriptorType
func (m_ MTL4IndirectInstanceAccelerationStructureDescriptor) SetInstanceDescriptorType(value AccelerationStructureInstanceDescriptorType) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInstanceDescriptorType:"), value)
}


// Specifies the layout for the transformation matrices in the instance descriptor buffer and the motion transformation matrix buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4IndirectInstanceAccelerationStructureDescriptor/instanceTransformationMatrixLayout
func (m_ MTL4IndirectInstanceAccelerationStructureDescriptor) InstanceTransformationMatrixLayout() MatrixLayout {
	rv := objc.Send[MatrixLayout](m_.ID, objc.Sel("instanceTransformationMatrixLayout"))
	return rv
}


// Specifies the layout for the transformation matrices in the instance descriptor buffer and the motion transformation matrix buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4IndirectInstanceAccelerationStructureDescriptor/instanceTransformationMatrixLayout
func (m_ MTL4IndirectInstanceAccelerationStructureDescriptor) SetInstanceTransformationMatrixLayout(value MatrixLayout) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInstanceTransformationMatrixLayout:"), value)
}


// Controls the maximum number of instance descriptors the instance descriptor buffer can reference.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4IndirectInstanceAccelerationStructureDescriptor/maxInstanceCount
func (m_ MTL4IndirectInstanceAccelerationStructureDescriptor) MaxInstanceCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("maxInstanceCount"))
	return rv
}


// Controls the maximum number of instance descriptors the instance descriptor buffer can reference.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4IndirectInstanceAccelerationStructureDescriptor/maxInstanceCount
func (m_ MTL4IndirectInstanceAccelerationStructureDescriptor) SetMaxInstanceCount(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxInstanceCount:"), value)
}


// Controls the maximum number of motion transforms in the motion transform buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4IndirectInstanceAccelerationStructureDescriptor/maxMotionTransformCount
func (m_ MTL4IndirectInstanceAccelerationStructureDescriptor) MaxMotionTransformCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("maxMotionTransformCount"))
	return rv
}


// Controls the maximum number of motion transforms in the motion transform buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4IndirectInstanceAccelerationStructureDescriptor/maxMotionTransformCount
func (m_ MTL4IndirectInstanceAccelerationStructureDescriptor) SetMaxMotionTransformCount(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxMotionTransformCount:"), value)
}


// A buffer containing transformation information for instance motion keyframes, formatted according to the motion transform type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4IndirectInstanceAccelerationStructureDescriptor/motionTransformBuffer
func (m_ MTL4IndirectInstanceAccelerationStructureDescriptor) MotionTransformBuffer() MTL4BufferRange {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("motionTransformBuffer"))
	return rv
}


// A buffer containing transformation information for instance motion keyframes, formatted according to the motion transform type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4IndirectInstanceAccelerationStructureDescriptor/motionTransformBuffer
func (m_ MTL4IndirectInstanceAccelerationStructureDescriptor) SetMotionTransformBuffer(value MTL4BufferRange) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMotionTransformBuffer:"), value)
}


// Associates a buffer reference containing the number of motion transforms in the motion transform buffer, formatted as a 32-bit unsigned integer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4IndirectInstanceAccelerationStructureDescriptor/motionTransformCountBuffer
func (m_ MTL4IndirectInstanceAccelerationStructureDescriptor) MotionTransformCountBuffer() MTL4BufferRange {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("motionTransformCountBuffer"))
	return rv
}


// Associates a buffer reference containing the number of motion transforms in the motion transform buffer, formatted as a 32-bit unsigned integer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4IndirectInstanceAccelerationStructureDescriptor/motionTransformCountBuffer
func (m_ MTL4IndirectInstanceAccelerationStructureDescriptor) SetMotionTransformCountBuffer(value MTL4BufferRange) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMotionTransformCountBuffer:"), value)
}


// Sets the stride for motion transform.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4IndirectInstanceAccelerationStructureDescriptor/motionTransformStride
func (m_ MTL4IndirectInstanceAccelerationStructureDescriptor) MotionTransformStride() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("motionTransformStride"))
	return rv
}


// Sets the stride for motion transform.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4IndirectInstanceAccelerationStructureDescriptor/motionTransformStride
func (m_ MTL4IndirectInstanceAccelerationStructureDescriptor) SetMotionTransformStride(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMotionTransformStride:"), value)
}


// Sets the type of motion transforms, either as a matrix or individual components.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4IndirectInstanceAccelerationStructureDescriptor/motionTransformType
func (m_ MTL4IndirectInstanceAccelerationStructureDescriptor) MotionTransformType() TransformType {
	rv := objc.Send[TransformType](m_.ID, objc.Sel("motionTransformType"))
	return rv
}


// Sets the type of motion transforms, either as a matrix or individual components.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4IndirectInstanceAccelerationStructureDescriptor/motionTransformType
func (m_ MTL4IndirectInstanceAccelerationStructureDescriptor) SetMotionTransformType(value TransformType) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMotionTransformType:"), value)
}








