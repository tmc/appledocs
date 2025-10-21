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
}

// Descriptor for an “indirect” instance acceleration structure that allows providing the instance count and motion transform count indirectly, through buffer references.
//
// An instance acceleration structure references other acceleration structures, and provides the ability to “instantiate” them multiple times, each one with potentially a different transformation matrix. You specify the properties of the instances in the acceleration structure this descriptor builds by providing a buffer of via its property. Compared to , this descriptor allows you to provide the number of instances it references indirectly through a buffer reference, as well as the number of motion transforms. This enables you to determine these counts indirectly in the GPU timeline via a compute pipeline. Metal needs only to know the maximum possible number of instances and motion transforms to support, which you specify via the and properties. Use a to mark residency of all buffers and acceleration structures this descriptor references when you build this acceleration structure.
//
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

// Alloc allocates a new instance without initialization.
func (mc _MTL4IndirectInstanceAccelerationStructureDescriptorClass) Alloc() MTL4IndirectInstanceAccelerationStructureDescriptor {
	rv := objc.Send[MTL4IndirectInstanceAccelerationStructureDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Associates a buffer reference containing the number of motion transforms in the motion transform buffer, formatted as a
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4indirectinstanceaccelerationstructuredescriptor/motiontransformcountbuffer
func (m_ MTL4IndirectInstanceAccelerationStructureDescriptor) MotionTransformCountBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("motionTransformCountBuffer"))
	return rv
}


// SetMotionTransformCountBuffer sets the value of the motionTransformCountBuffer property.
// Associates a buffer reference containing the number of motion transforms in the motion transform buffer, formatted as a

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4indirectinstanceaccelerationstructuredescriptor/motiontransformcountbuffer
func (m_ MTL4IndirectInstanceAccelerationStructureDescriptor) SetMotionTransformCountBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMotionTransformCountBuffer:"), value)
}

// A buffer containing transformation information for instance motion keyframes, formatted according
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4indirectinstanceaccelerationstructuredescriptor/motiontransformbuffer
func (m_ MTL4IndirectInstanceAccelerationStructureDescriptor) MotionTransformBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("motionTransformBuffer"))
	return rv
}


// SetMotionTransformBuffer sets the value of the motionTransformBuffer property.
// A buffer containing transformation information for instance motion keyframes, formatted according

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4indirectinstanceaccelerationstructuredescriptor/motiontransformbuffer
func (m_ MTL4IndirectInstanceAccelerationStructureDescriptor) SetMotionTransformBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMotionTransformBuffer:"), value)
}

// Controls the maximum number of instance descriptors the instance descriptor buffer can reference.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4indirectinstanceaccelerationstructuredescriptor/maxinstancecount
func (m_ MTL4IndirectInstanceAccelerationStructureDescriptor) MaxInstanceCount() int {
	rv := objc.Send[int](m_.ID, objc.Sel("maxInstanceCount"))
	return rv
}


// SetMaxInstanceCount sets the value of the maxInstanceCount property.
// Controls the maximum number of instance descriptors the instance descriptor buffer can reference.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4indirectinstanceaccelerationstructuredescriptor/maxinstancecount
func (m_ MTL4IndirectInstanceAccelerationStructureDescriptor) SetMaxInstanceCount(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxInstanceCount:"), value)
}

// Controls the maximum number of motion transforms in the motion transform buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4indirectinstanceaccelerationstructuredescriptor/maxmotiontransformcount
func (m_ MTL4IndirectInstanceAccelerationStructureDescriptor) MaxMotionTransformCount() int {
	rv := objc.Send[int](m_.ID, objc.Sel("maxMotionTransformCount"))
	return rv
}


// SetMaxMotionTransformCount sets the value of the maxMotionTransformCount property.
// Controls the maximum number of motion transforms in the motion transform buffer.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4indirectinstanceaccelerationstructuredescriptor/maxmotiontransformcount
func (m_ MTL4IndirectInstanceAccelerationStructureDescriptor) SetMaxMotionTransformCount(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxMotionTransformCount:"), value)
}

// Sets the type of motion transforms, either as a matrix or individual components.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4indirectinstanceaccelerationstructuredescriptor/motiontransformtype
func (m_ MTL4IndirectInstanceAccelerationStructureDescriptor) MotionTransformType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("motionTransformType"))
	return rv
}


// SetMotionTransformType sets the value of the motionTransformType property.
// Sets the type of motion transforms, either as a matrix or individual components.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4indirectinstanceaccelerationstructuredescriptor/motiontransformtype
func (m_ MTL4IndirectInstanceAccelerationStructureDescriptor) SetMotionTransformType(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMotionTransformType:"), value)
}

// Specifies the layout for the transformation matrices in the instance descriptor buffer and the motion transformation matrix buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4indirectinstanceaccelerationstructuredescriptor/instancetransformationmatrixlayout
func (m_ MTL4IndirectInstanceAccelerationStructureDescriptor) InstanceTransformationMatrixLayout() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("instanceTransformationMatrixLayout"))
	return rv
}


// SetInstanceTransformationMatrixLayout sets the value of the instanceTransformationMatrixLayout property.
// Specifies the layout for the transformation matrices in the instance descriptor buffer and the motion transformation matrix buffer.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4indirectinstanceaccelerationstructuredescriptor/instancetransformationmatrixlayout
func (m_ MTL4IndirectInstanceAccelerationStructureDescriptor) SetInstanceTransformationMatrixLayout(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInstanceTransformationMatrixLayout:"), value)
}

// Sets the stride, in bytes, between instance descriptors in the instance descriptor buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4indirectinstanceaccelerationstructuredescriptor/instancedescriptorstride
func (m_ MTL4IndirectInstanceAccelerationStructureDescriptor) InstanceDescriptorStride() int {
	rv := objc.Send[int](m_.ID, objc.Sel("instanceDescriptorStride"))
	return rv
}


// SetInstanceDescriptorStride sets the value of the instanceDescriptorStride property.
// Sets the stride, in bytes, between instance descriptors in the instance descriptor buffer.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4indirectinstanceaccelerationstructuredescriptor/instancedescriptorstride
func (m_ MTL4IndirectInstanceAccelerationStructureDescriptor) SetInstanceDescriptorStride(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInstanceDescriptorStride:"), value)
}

// Provides a reference to a buffer containing the number of instances in the instance descriptor buffer, formatted as a
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4indirectinstanceaccelerationstructuredescriptor/instancecountbuffer
func (m_ MTL4IndirectInstanceAccelerationStructureDescriptor) InstanceCountBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("instanceCountBuffer"))
	return rv
}


// SetInstanceCountBuffer sets the value of the instanceCountBuffer property.
// Provides a reference to a buffer containing the number of instances in the instance descriptor buffer, formatted as a

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4indirectinstanceaccelerationstructuredescriptor/instancecountbuffer
func (m_ MTL4IndirectInstanceAccelerationStructureDescriptor) SetInstanceCountBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInstanceCountBuffer:"), value)
}

// Controls the type of instance descriptor that the instance descriptor buffer references.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4indirectinstanceaccelerationstructuredescriptor/instancedescriptortype
func (m_ MTL4IndirectInstanceAccelerationStructureDescriptor) InstanceDescriptorType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("instanceDescriptorType"))
	return rv
}


// SetInstanceDescriptorType sets the value of the instanceDescriptorType property.
// Controls the type of instance descriptor that the instance descriptor buffer references.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4indirectinstanceaccelerationstructuredescriptor/instancedescriptortype
func (m_ MTL4IndirectInstanceAccelerationStructureDescriptor) SetInstanceDescriptorType(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInstanceDescriptorType:"), value)
}

// Sets the stride for motion transform.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4indirectinstanceaccelerationstructuredescriptor/motiontransformstride
func (m_ MTL4IndirectInstanceAccelerationStructureDescriptor) MotionTransformStride() int {
	rv := objc.Send[int](m_.ID, objc.Sel("motionTransformStride"))
	return rv
}


// SetMotionTransformStride sets the value of the motionTransformStride property.
// Sets the stride for motion transform.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4indirectinstanceaccelerationstructuredescriptor/motiontransformstride
func (m_ MTL4IndirectInstanceAccelerationStructureDescriptor) SetMotionTransformStride(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMotionTransformStride:"), value)
}

// Assigns a reference to a buffer containing instance descriptors for acceleration structures to reference.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4indirectinstanceaccelerationstructuredescriptor/instancedescriptorbuffer
func (m_ MTL4IndirectInstanceAccelerationStructureDescriptor) InstanceDescriptorBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("instanceDescriptorBuffer"))
	return rv
}


// SetInstanceDescriptorBuffer sets the value of the instanceDescriptorBuffer property.
// Assigns a reference to a buffer containing instance descriptors for acceleration structures to reference.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4indirectinstanceaccelerationstructuredescriptor/instancedescriptorbuffer
func (m_ MTL4IndirectInstanceAccelerationStructureDescriptor) SetInstanceDescriptorBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInstanceDescriptorBuffer:"), value)
}



