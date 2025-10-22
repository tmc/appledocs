// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTL4InstanceAccelerationStructureDescriptor] class.
var (
	MTL4InstanceAccelerationStructureDescriptorClass     _MTL4InstanceAccelerationStructureDescriptorClass
	MTL4InstanceAccelerationStructureDescriptorClassOnce sync.Once
)

func getMTL4InstanceAccelerationStructureDescriptorClass() _MTL4InstanceAccelerationStructureDescriptorClass {
	MTL4InstanceAccelerationStructureDescriptorClassOnce.Do(func() {
		MTL4InstanceAccelerationStructureDescriptorClass = _MTL4InstanceAccelerationStructureDescriptorClass{objc.GetClass("MTL4InstanceAccelerationStructureDescriptor")}
	})
	return MTL4InstanceAccelerationStructureDescriptorClass
}

type _MTL4InstanceAccelerationStructureDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [MTL4InstanceAccelerationStructureDescriptor] class.
type IMTL4InstanceAccelerationStructureDescriptor interface {
	IMTL4AccelerationStructureDescriptor
	InstanceCount() int
	SetInstanceCount(value int)
	InstanceDescriptorBuffer() unsafe.Pointer
	SetInstanceDescriptorBuffer(value unsafe.Pointer)
	InstanceDescriptorStride() int
	SetInstanceDescriptorStride(value int)
	InstanceDescriptorType() unsafe.Pointer
	SetInstanceDescriptorType(value unsafe.Pointer)
	InstanceTransformationMatrixLayout() unsafe.Pointer
	SetInstanceTransformationMatrixLayout(value unsafe.Pointer)
	MotionTransformBuffer() unsafe.Pointer
	SetMotionTransformBuffer(value unsafe.Pointer)
	MotionTransformCount() int
	SetMotionTransformCount(value int)
	MotionTransformStride() int
	SetMotionTransformStride(value int)
	MotionTransformType() TransformType
	SetMotionTransformType(value TransformType)
}

// Descriptor for an instance acceleration structure.
//
// An instance acceleration structure references other acceleration structures, and provides the ability to “instantiate” them multiple times, each one with potentially a different transformation matrix. You specify the properties of the instances in the acceleration structure this descriptor builds by providing a buffer of via its property. Use a to mark residency of all buffers and acceleration structures this descriptor references when you build this acceleration structure.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4InstanceAccelerationStructureDescriptor
type MTL4InstanceAccelerationStructureDescriptor struct {
	MTL4AccelerationStructureDescriptor
}

// MTL4InstanceAccelerationStructureDescriptorFrom constructs a [MTL4InstanceAccelerationStructureDescriptor] from an unsafe.Pointer.
//
// Descriptor for an instance acceleration structure.
func MTL4InstanceAccelerationStructureDescriptorFrom(ptr unsafe.Pointer) MTL4InstanceAccelerationStructureDescriptor {
	return MTL4InstanceAccelerationStructureDescriptor{
		MTL4AccelerationStructureDescriptor: MTL4AccelerationStructureDescriptorFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTL4InstanceAccelerationStructureDescriptorClass) Alloc() MTL4InstanceAccelerationStructureDescriptor {
	rv := objc.Send[MTL4InstanceAccelerationStructureDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTL4InstanceAccelerationStructureDescriptorClass) New() MTL4InstanceAccelerationStructureDescriptor {
	rv := objc.Send[MTL4InstanceAccelerationStructureDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTL4InstanceAccelerationStructureDescriptor) Init() MTL4InstanceAccelerationStructureDescriptor {
	rv := objc.Send[MTL4InstanceAccelerationStructureDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTL4InstanceAccelerationStructureDescriptor) Autorelease() MTL4InstanceAccelerationStructureDescriptor {
	rv := objc.Send[MTL4InstanceAccelerationStructureDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4InstanceAccelerationStructureDescriptor creates a new MTL4InstanceAccelerationStructureDescriptor instance.
func NewMTL4InstanceAccelerationStructureDescriptor() MTL4InstanceAccelerationStructureDescriptor {
	return getMTL4InstanceAccelerationStructureDescriptorClass().New()
}


// Controls the number of instance descriptors in the instance descriptor buffer references.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4instanceaccelerationstructuredescriptor/instancecount
func (m_ MTL4InstanceAccelerationStructureDescriptor) InstanceCount() int {
	rv := objc.Send[int](m_.ID, objc.Sel("instanceCount"))
	return rv
}


// SetInstanceCount sets the value of the instanceCount property.
// Controls the number of instance descriptors in the instance descriptor buffer references.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4instanceaccelerationstructuredescriptor/instancecount
func (m_ MTL4InstanceAccelerationStructureDescriptor) SetInstanceCount(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInstanceCount:"), value)
}

// Assigns a reference to a buffer containing instance descriptors for acceleration structures to reference.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4instanceaccelerationstructuredescriptor/instancedescriptorbuffer
func (m_ MTL4InstanceAccelerationStructureDescriptor) InstanceDescriptorBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("instanceDescriptorBuffer"))
	return rv
}


// SetInstanceDescriptorBuffer sets the value of the instanceDescriptorBuffer property.
// Assigns a reference to a buffer containing instance descriptors for acceleration structures to reference.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4instanceaccelerationstructuredescriptor/instancedescriptorbuffer
func (m_ MTL4InstanceAccelerationStructureDescriptor) SetInstanceDescriptorBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInstanceDescriptorBuffer:"), value)
}

// Sets the stride, in bytes, between instance descriptors the instance descriptor buffer references.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4instanceaccelerationstructuredescriptor/instancedescriptorstride
func (m_ MTL4InstanceAccelerationStructureDescriptor) InstanceDescriptorStride() int {
	rv := objc.Send[int](m_.ID, objc.Sel("instanceDescriptorStride"))
	return rv
}


// SetInstanceDescriptorStride sets the value of the instanceDescriptorStride property.
// Sets the stride, in bytes, between instance descriptors the instance descriptor buffer references.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4instanceaccelerationstructuredescriptor/instancedescriptorstride
func (m_ MTL4InstanceAccelerationStructureDescriptor) SetInstanceDescriptorStride(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInstanceDescriptorStride:"), value)
}

// Sets the type of instance descriptor that the instance descriptor buffer references.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4instanceaccelerationstructuredescriptor/instancedescriptortype
func (m_ MTL4InstanceAccelerationStructureDescriptor) InstanceDescriptorType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("instanceDescriptorType"))
	return rv
}


// SetInstanceDescriptorType sets the value of the instanceDescriptorType property.
// Sets the type of instance descriptor that the instance descriptor buffer references.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4instanceaccelerationstructuredescriptor/instancedescriptortype
func (m_ MTL4InstanceAccelerationStructureDescriptor) SetInstanceDescriptorType(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInstanceDescriptorType:"), value)
}

// Specifies the layout for the transformation matrices in the instance descriptor buffer and the motion transformation matrix buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4instanceaccelerationstructuredescriptor/instancetransformationmatrixlayout
func (m_ MTL4InstanceAccelerationStructureDescriptor) InstanceTransformationMatrixLayout() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("instanceTransformationMatrixLayout"))
	return rv
}


// SetInstanceTransformationMatrixLayout sets the value of the instanceTransformationMatrixLayout property.
// Specifies the layout for the transformation matrices in the instance descriptor buffer and the motion transformation matrix buffer.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4instanceaccelerationstructuredescriptor/instancetransformationmatrixlayout
func (m_ MTL4InstanceAccelerationStructureDescriptor) SetInstanceTransformationMatrixLayout(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInstanceTransformationMatrixLayout:"), value)
}

// A buffer containing transformation information for instance motion keyframes, formatted according
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4instanceaccelerationstructuredescriptor/motiontransformbuffer
func (m_ MTL4InstanceAccelerationStructureDescriptor) MotionTransformBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("motionTransformBuffer"))
	return rv
}


// SetMotionTransformBuffer sets the value of the motionTransformBuffer property.
// A buffer containing transformation information for instance motion keyframes, formatted according

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4instanceaccelerationstructuredescriptor/motiontransformbuffer
func (m_ MTL4InstanceAccelerationStructureDescriptor) SetMotionTransformBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMotionTransformBuffer:"), value)
}

// Controls the total number of motion transforms in the motion transform buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4instanceaccelerationstructuredescriptor/motiontransformcount
func (m_ MTL4InstanceAccelerationStructureDescriptor) MotionTransformCount() int {
	rv := objc.Send[int](m_.ID, objc.Sel("motionTransformCount"))
	return rv
}


// SetMotionTransformCount sets the value of the motionTransformCount property.
// Controls the total number of motion transforms in the motion transform buffer.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4instanceaccelerationstructuredescriptor/motiontransformcount
func (m_ MTL4InstanceAccelerationStructureDescriptor) SetMotionTransformCount(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMotionTransformCount:"), value)
}

// Specify the stride for motion transform.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4instanceaccelerationstructuredescriptor/motiontransformstride
func (m_ MTL4InstanceAccelerationStructureDescriptor) MotionTransformStride() int {
	rv := objc.Send[int](m_.ID, objc.Sel("motionTransformStride"))
	return rv
}


// SetMotionTransformStride sets the value of the motionTransformStride property.
// Specify the stride for motion transform.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4instanceaccelerationstructuredescriptor/motiontransformstride
func (m_ MTL4InstanceAccelerationStructureDescriptor) SetMotionTransformStride(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMotionTransformStride:"), value)
}

// Controls the type of motion transforms, either as a matrix or individual components.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4instanceaccelerationstructuredescriptor/motiontransformtype
func (m_ MTL4InstanceAccelerationStructureDescriptor) MotionTransformType() TransformType {
	rv := objc.Send[TransformType](m_.ID, objc.Sel("motionTransformType"))
	return rv
}


// SetMotionTransformType sets the value of the motionTransformType property.
// Controls the type of motion transforms, either as a matrix or individual components.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4instanceaccelerationstructuredescriptor/motiontransformtype
func (m_ MTL4InstanceAccelerationStructureDescriptor) SetMotionTransformType(value TransformType) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMotionTransformType:"), value)
}



