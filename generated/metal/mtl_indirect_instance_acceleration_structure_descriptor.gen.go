// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
}

// A description of an acceleration structure that Metal derives from instances of primitive acceleration structures that the GPU can populate.
//
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

// Alloc allocates a new instance without initialization.
func (ic _IndirectInstanceAccelerationStructureDescriptorClass) Alloc() IndirectInstanceAccelerationStructureDescriptor {
	rv := objc.Send[IndirectInstanceAccelerationStructureDescriptor](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlindirectinstanceaccelerationstructuredescriptor/instancecountbuffer
func (i_ IndirectInstanceAccelerationStructureDescriptor) InstanceCountBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("instanceCountBuffer"))
	return rv
}


// SetInstanceCountBuffer sets the value of the instanceCountBuffer property.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlindirectinstanceaccelerationstructuredescriptor/instancecountbuffer
func (i_ IndirectInstanceAccelerationStructureDescriptor) SetInstanceCountBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setInstanceCountBuffer:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlindirectinstanceaccelerationstructuredescriptor/instancecountbufferoffset
func (i_ IndirectInstanceAccelerationStructureDescriptor) InstanceCountBufferOffset() int {
	rv := objc.Send[int](i_.ID, objc.Sel("instanceCountBufferOffset"))
	return rv
}


// SetInstanceCountBufferOffset sets the value of the instanceCountBufferOffset property.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlindirectinstanceaccelerationstructuredescriptor/instancecountbufferoffset
func (i_ IndirectInstanceAccelerationStructureDescriptor) SetInstanceCountBufferOffset(value int) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setInstanceCountBufferOffset:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlindirectinstanceaccelerationstructuredescriptor/instancedescriptorbuffer
func (i_ IndirectInstanceAccelerationStructureDescriptor) InstanceDescriptorBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("instanceDescriptorBuffer"))
	return rv
}


// SetInstanceDescriptorBuffer sets the value of the instanceDescriptorBuffer property.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlindirectinstanceaccelerationstructuredescriptor/instancedescriptorbuffer
func (i_ IndirectInstanceAccelerationStructureDescriptor) SetInstanceDescriptorBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setInstanceDescriptorBuffer:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlindirectinstanceaccelerationstructuredescriptor/instancedescriptorbufferoffset
func (i_ IndirectInstanceAccelerationStructureDescriptor) InstanceDescriptorBufferOffset() int {
	rv := objc.Send[int](i_.ID, objc.Sel("instanceDescriptorBufferOffset"))
	return rv
}


// SetInstanceDescriptorBufferOffset sets the value of the instanceDescriptorBufferOffset property.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlindirectinstanceaccelerationstructuredescriptor/instancedescriptorbufferoffset
func (i_ IndirectInstanceAccelerationStructureDescriptor) SetInstanceDescriptorBufferOffset(value int) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setInstanceDescriptorBufferOffset:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlindirectinstanceaccelerationstructuredescriptor/instancedescriptorstride
func (i_ IndirectInstanceAccelerationStructureDescriptor) InstanceDescriptorStride() int {
	rv := objc.Send[int](i_.ID, objc.Sel("instanceDescriptorStride"))
	return rv
}


// SetInstanceDescriptorStride sets the value of the instanceDescriptorStride property.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlindirectinstanceaccelerationstructuredescriptor/instancedescriptorstride
func (i_ IndirectInstanceAccelerationStructureDescriptor) SetInstanceDescriptorStride(value int) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setInstanceDescriptorStride:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlindirectinstanceaccelerationstructuredescriptor/instancedescriptortype
func (i_ IndirectInstanceAccelerationStructureDescriptor) InstanceDescriptorType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("instanceDescriptorType"))
	return rv
}


// SetInstanceDescriptorType sets the value of the instanceDescriptorType property.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlindirectinstanceaccelerationstructuredescriptor/instancedescriptortype
func (i_ IndirectInstanceAccelerationStructureDescriptor) SetInstanceDescriptorType(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setInstanceDescriptorType:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlindirectinstanceaccelerationstructuredescriptor/instancetransformationmatrixlayout
func (i_ IndirectInstanceAccelerationStructureDescriptor) InstanceTransformationMatrixLayout() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("instanceTransformationMatrixLayout"))
	return rv
}


// SetInstanceTransformationMatrixLayout sets the value of the instanceTransformationMatrixLayout property.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlindirectinstanceaccelerationstructuredescriptor/instancetransformationmatrixlayout
func (i_ IndirectInstanceAccelerationStructureDescriptor) SetInstanceTransformationMatrixLayout(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setInstanceTransformationMatrixLayout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlindirectinstanceaccelerationstructuredescriptor/maxinstancecount
func (i_ IndirectInstanceAccelerationStructureDescriptor) MaxInstanceCount() int {
	rv := objc.Send[int](i_.ID, objc.Sel("maxInstanceCount"))
	return rv
}


// SetMaxInstanceCount sets the value of the maxInstanceCount property.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlindirectinstanceaccelerationstructuredescriptor/maxinstancecount
func (i_ IndirectInstanceAccelerationStructureDescriptor) SetMaxInstanceCount(value int) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMaxInstanceCount:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlindirectinstanceaccelerationstructuredescriptor/maxmotiontransformcount
func (i_ IndirectInstanceAccelerationStructureDescriptor) MaxMotionTransformCount() int {
	rv := objc.Send[int](i_.ID, objc.Sel("maxMotionTransformCount"))
	return rv
}


// SetMaxMotionTransformCount sets the value of the maxMotionTransformCount property.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlindirectinstanceaccelerationstructuredescriptor/maxmotiontransformcount
func (i_ IndirectInstanceAccelerationStructureDescriptor) SetMaxMotionTransformCount(value int) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMaxMotionTransformCount:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlindirectinstanceaccelerationstructuredescriptor/motiontransformbuffer
func (i_ IndirectInstanceAccelerationStructureDescriptor) MotionTransformBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("motionTransformBuffer"))
	return rv
}


// SetMotionTransformBuffer sets the value of the motionTransformBuffer property.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlindirectinstanceaccelerationstructuredescriptor/motiontransformbuffer
func (i_ IndirectInstanceAccelerationStructureDescriptor) SetMotionTransformBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMotionTransformBuffer:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlindirectinstanceaccelerationstructuredescriptor/motiontransformbufferoffset
func (i_ IndirectInstanceAccelerationStructureDescriptor) MotionTransformBufferOffset() int {
	rv := objc.Send[int](i_.ID, objc.Sel("motionTransformBufferOffset"))
	return rv
}


// SetMotionTransformBufferOffset sets the value of the motionTransformBufferOffset property.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlindirectinstanceaccelerationstructuredescriptor/motiontransformbufferoffset
func (i_ IndirectInstanceAccelerationStructureDescriptor) SetMotionTransformBufferOffset(value int) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMotionTransformBufferOffset:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlindirectinstanceaccelerationstructuredescriptor/motiontransformcountbuffer
func (i_ IndirectInstanceAccelerationStructureDescriptor) MotionTransformCountBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("motionTransformCountBuffer"))
	return rv
}


// SetMotionTransformCountBuffer sets the value of the motionTransformCountBuffer property.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlindirectinstanceaccelerationstructuredescriptor/motiontransformcountbuffer
func (i_ IndirectInstanceAccelerationStructureDescriptor) SetMotionTransformCountBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMotionTransformCountBuffer:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlindirectinstanceaccelerationstructuredescriptor/motiontransformcountbufferoffset
func (i_ IndirectInstanceAccelerationStructureDescriptor) MotionTransformCountBufferOffset() int {
	rv := objc.Send[int](i_.ID, objc.Sel("motionTransformCountBufferOffset"))
	return rv
}


// SetMotionTransformCountBufferOffset sets the value of the motionTransformCountBufferOffset property.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlindirectinstanceaccelerationstructuredescriptor/motiontransformcountbufferoffset
func (i_ IndirectInstanceAccelerationStructureDescriptor) SetMotionTransformCountBufferOffset(value int) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMotionTransformCountBufferOffset:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlindirectinstanceaccelerationstructuredescriptor/motiontransformstride
func (i_ IndirectInstanceAccelerationStructureDescriptor) MotionTransformStride() int {
	rv := objc.Send[int](i_.ID, objc.Sel("motionTransformStride"))
	return rv
}


// SetMotionTransformStride sets the value of the motionTransformStride property.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlindirectinstanceaccelerationstructuredescriptor/motiontransformstride
func (i_ IndirectInstanceAccelerationStructureDescriptor) SetMotionTransformStride(value int) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMotionTransformStride:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlindirectinstanceaccelerationstructuredescriptor/motiontransformtype
func (i_ IndirectInstanceAccelerationStructureDescriptor) MotionTransformType() TransformType {
	rv := objc.Send[TransformType](i_.ID, objc.Sel("motionTransformType"))
	return rv
}


// SetMotionTransformType sets the value of the motionTransformType property.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlindirectinstanceaccelerationstructuredescriptor/motiontransformtype
func (i_ IndirectInstanceAccelerationStructureDescriptor) SetMotionTransformType(value TransformType) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMotionTransformType:"), value)
}



