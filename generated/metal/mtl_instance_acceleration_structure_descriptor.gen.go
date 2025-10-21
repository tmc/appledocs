// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [InstanceAccelerationStructureDescriptor] class.
var (
	InstanceAccelerationStructureDescriptorClass     _InstanceAccelerationStructureDescriptorClass
	InstanceAccelerationStructureDescriptorClassOnce sync.Once
)

func getInstanceAccelerationStructureDescriptorClass() _InstanceAccelerationStructureDescriptorClass {
	InstanceAccelerationStructureDescriptorClassOnce.Do(func() {
		InstanceAccelerationStructureDescriptorClass = _InstanceAccelerationStructureDescriptorClass{objc.GetClass("MTLInstanceAccelerationStructureDescriptor")}
	})
	return InstanceAccelerationStructureDescriptorClass
}

type _InstanceAccelerationStructureDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [InstanceAccelerationStructureDescriptor] class.
type IInstanceAccelerationStructureDescriptor interface {
	IAccelerationStructureDescriptor
}

// A description of an acceleration structure that derives from instances of primitive acceleration structures.
//
// Metal provides acceleration structures with a two-level hierarchy. The bottom layer consists of primitive acceleration structures, which instance acceleration structures in the top level reference.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLInstanceAccelerationStructureDescriptor
type InstanceAccelerationStructureDescriptor struct {
	AccelerationStructureDescriptor
}

// InstanceAccelerationStructureDescriptorFrom constructs a [InstanceAccelerationStructureDescriptor] from an unsafe.Pointer.
//
// A description of an acceleration structure that derives from instances of primitive acceleration structures.
func InstanceAccelerationStructureDescriptorFrom(ptr unsafe.Pointer) InstanceAccelerationStructureDescriptor {
	return InstanceAccelerationStructureDescriptor{
		AccelerationStructureDescriptor: AccelerationStructureDescriptorFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _InstanceAccelerationStructureDescriptorClass) Alloc() InstanceAccelerationStructureDescriptor {
	rv := objc.Send[InstanceAccelerationStructureDescriptor](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _InstanceAccelerationStructureDescriptorClass) New() InstanceAccelerationStructureDescriptor {
	rv := objc.Send[InstanceAccelerationStructureDescriptor](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ InstanceAccelerationStructureDescriptor) Init() InstanceAccelerationStructureDescriptor {
	rv := objc.Send[InstanceAccelerationStructureDescriptor](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ InstanceAccelerationStructureDescriptor) Autorelease() InstanceAccelerationStructureDescriptor {
	rv := objc.Send[InstanceAccelerationStructureDescriptor](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewInstanceAccelerationStructureDescriptor creates a new InstanceAccelerationStructureDescriptor instance.
func NewInstanceAccelerationStructureDescriptor() InstanceAccelerationStructureDescriptor {
	return getInstanceAccelerationStructureDescriptorClass().New()
}


// The number of instances in the instance descriptor buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlinstanceaccelerationstructuredescriptor/instancecount
func (i_ InstanceAccelerationStructureDescriptor) InstanceCount() int {
	rv := objc.Send[int](i_.ID, objc.Sel("instanceCount"))
	return rv
}


// SetInstanceCount sets the value of the instanceCount property.
// The number of instances in the instance descriptor buffer.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlinstanceaccelerationstructuredescriptor/instancecount
func (i_ InstanceAccelerationStructureDescriptor) SetInstanceCount(value int) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setInstanceCount:"), value)
}

// A buffer that contains descriptions of each instance in the acceleration structure.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlinstanceaccelerationstructuredescriptor/instancedescriptorbuffer
func (i_ InstanceAccelerationStructureDescriptor) InstanceDescriptorBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("instanceDescriptorBuffer"))
	return rv
}


// SetInstanceDescriptorBuffer sets the value of the instanceDescriptorBuffer property.
// A buffer that contains descriptions of each instance in the acceleration structure.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlinstanceaccelerationstructuredescriptor/instancedescriptorbuffer
func (i_ InstanceAccelerationStructureDescriptor) SetInstanceDescriptorBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setInstanceDescriptorBuffer:"), value)
}

// The offset, in bytes, to the descripton of the first instance.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlinstanceaccelerationstructuredescriptor/instancedescriptorbufferoffset
func (i_ InstanceAccelerationStructureDescriptor) InstanceDescriptorBufferOffset() int {
	rv := objc.Send[int](i_.ID, objc.Sel("instanceDescriptorBufferOffset"))
	return rv
}


// SetInstanceDescriptorBufferOffset sets the value of the instanceDescriptorBufferOffset property.
// The offset, in bytes, to the descripton of the first instance.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlinstanceaccelerationstructuredescriptor/instancedescriptorbufferoffset
func (i_ InstanceAccelerationStructureDescriptor) SetInstanceDescriptorBufferOffset(value int) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setInstanceDescriptorBufferOffset:"), value)
}

// The stride, in bytes, between instance descriptions.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlinstanceaccelerationstructuredescriptor/instancedescriptorstride
func (i_ InstanceAccelerationStructureDescriptor) InstanceDescriptorStride() int {
	rv := objc.Send[int](i_.ID, objc.Sel("instanceDescriptorStride"))
	return rv
}


// SetInstanceDescriptorStride sets the value of the instanceDescriptorStride property.
// The stride, in bytes, between instance descriptions.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlinstanceaccelerationstructuredescriptor/instancedescriptorstride
func (i_ InstanceAccelerationStructureDescriptor) SetInstanceDescriptorStride(value int) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setInstanceDescriptorStride:"), value)
}

// The format of the instance data in the descriptor buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlinstanceaccelerationstructuredescriptor/instancedescriptortype
func (i_ InstanceAccelerationStructureDescriptor) InstanceDescriptorType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("instanceDescriptorType"))
	return rv
}


// SetInstanceDescriptorType sets the value of the instanceDescriptorType property.
// The format of the instance data in the descriptor buffer.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlinstanceaccelerationstructuredescriptor/instancedescriptortype
func (i_ InstanceAccelerationStructureDescriptor) SetInstanceDescriptorType(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setInstanceDescriptorType:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlinstanceaccelerationstructuredescriptor/instancetransformationmatrixlayout
func (i_ InstanceAccelerationStructureDescriptor) InstanceTransformationMatrixLayout() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("instanceTransformationMatrixLayout"))
	return rv
}


// SetInstanceTransformationMatrixLayout sets the value of the instanceTransformationMatrixLayout property.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlinstanceaccelerationstructuredescriptor/instancetransformationmatrixlayout
func (i_ InstanceAccelerationStructureDescriptor) SetInstanceTransformationMatrixLayout(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setInstanceTransformationMatrixLayout:"), value)
}

// The bottom-level acceleration structures that instances use in the instance acceleration structure .
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlinstanceaccelerationstructuredescriptor/instancedaccelerationstructures
func (i_ InstanceAccelerationStructureDescriptor) InstancedAccelerationStructures() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("instancedAccelerationStructures"))
	return rv
}


// SetInstancedAccelerationStructures sets the value of the instancedAccelerationStructures property.
// The bottom-level acceleration structures that instances use in the instance acceleration structure .

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlinstanceaccelerationstructuredescriptor/instancedaccelerationstructures
func (i_ InstanceAccelerationStructureDescriptor) SetInstancedAccelerationStructures(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setInstancedAccelerationStructures:"), value)
}

// A buffer that contains descriptions of each motion transform in the acceleration structure.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlinstanceaccelerationstructuredescriptor/motiontransformbuffer
func (i_ InstanceAccelerationStructureDescriptor) MotionTransformBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("motionTransformBuffer"))
	return rv
}


// SetMotionTransformBuffer sets the value of the motionTransformBuffer property.
// A buffer that contains descriptions of each motion transform in the acceleration structure.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlinstanceaccelerationstructuredescriptor/motiontransformbuffer
func (i_ InstanceAccelerationStructureDescriptor) SetMotionTransformBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMotionTransformBuffer:"), value)
}

// The offset, in bytes, to the descripton of the first motion transform.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlinstanceaccelerationstructuredescriptor/motiontransformbufferoffset
func (i_ InstanceAccelerationStructureDescriptor) MotionTransformBufferOffset() int {
	rv := objc.Send[int](i_.ID, objc.Sel("motionTransformBufferOffset"))
	return rv
}


// SetMotionTransformBufferOffset sets the value of the motionTransformBufferOffset property.
// The offset, in bytes, to the descripton of the first motion transform.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlinstanceaccelerationstructuredescriptor/motiontransformbufferoffset
func (i_ InstanceAccelerationStructureDescriptor) SetMotionTransformBufferOffset(value int) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMotionTransformBufferOffset:"), value)
}

// The number of motion transforms in the motion transform buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlinstanceaccelerationstructuredescriptor/motiontransformcount
func (i_ InstanceAccelerationStructureDescriptor) MotionTransformCount() int {
	rv := objc.Send[int](i_.ID, objc.Sel("motionTransformCount"))
	return rv
}


// SetMotionTransformCount sets the value of the motionTransformCount property.
// The number of motion transforms in the motion transform buffer.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlinstanceaccelerationstructuredescriptor/motiontransformcount
func (i_ InstanceAccelerationStructureDescriptor) SetMotionTransformCount(value int) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMotionTransformCount:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlinstanceaccelerationstructuredescriptor/motiontransformstride
func (i_ InstanceAccelerationStructureDescriptor) MotionTransformStride() int {
	rv := objc.Send[int](i_.ID, objc.Sel("motionTransformStride"))
	return rv
}


// SetMotionTransformStride sets the value of the motionTransformStride property.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlinstanceaccelerationstructuredescriptor/motiontransformstride
func (i_ InstanceAccelerationStructureDescriptor) SetMotionTransformStride(value int) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMotionTransformStride:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlinstanceaccelerationstructuredescriptor/motiontransformtype
func (i_ InstanceAccelerationStructureDescriptor) MotionTransformType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("motionTransformType"))
	return rv
}


// SetMotionTransformType sets the value of the motionTransformType property.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlinstanceaccelerationstructuredescriptor/motiontransformtype
func (i_ InstanceAccelerationStructureDescriptor) SetMotionTransformType(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMotionTransformType:"), value)
}



