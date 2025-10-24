// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLInstanceAccelerationStructureDescriptor */


/* debug [class_header]: Header for MTLInstanceAccelerationStructureDescriptor */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for InstanceAccelerationStructureDescriptor */
// An interface definition for the [InstanceAccelerationStructureDescriptor] class.
type IInstanceAccelerationStructureDescriptor interface {
	IAccelerationStructureDescriptor
	
/* debug [class_interface_properties]: Properties for InstanceAccelerationStructureDescriptor */
	// properties:
	InstanceCount() uint
	SetInstanceCount(value uint)
	InstancedAccelerationStructures() []objc.ID
	SetInstancedAccelerationStructures(value []objc.ID)
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
	MotionTransformBuffer() unsafe.Pointer
	SetMotionTransformBuffer(value unsafe.Pointer)
	MotionTransformBufferOffset() uint
	SetMotionTransformBufferOffset(value uint)
	MotionTransformCount() uint
	SetMotionTransformCount(value uint)
	MotionTransformStride() uint
	SetMotionTransformStride(value uint)
	MotionTransformType() TransformType
	SetMotionTransformType(value TransformType)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for InstanceAccelerationStructureDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for InstanceAccelerationStructureDescriptor */
// Alloc allocates a new instance without initialization.
func (ic _InstanceAccelerationStructureDescriptorClass) Alloc() InstanceAccelerationStructureDescriptor {
	rv := objc.Send[InstanceAccelerationStructureDescriptor](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for InstanceAccelerationStructureDescriptor */
// A description of an acceleration structure that derives from instances of primitive acceleration structures.
//
// Metal provides acceleration structures with a two-level hierarchy. The bottom layer consists of primitive acceleration structures, which instance acceleration structures in the top level reference.


// A description of an acceleration structure that derives from instances of primitive acceleration structures.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for InstanceAccelerationStructureDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for InstanceAccelerationStructureDescriptor */

// Creates an instance descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLInstanceAccelerationStructureDescriptor/descriptor
func (ic _InstanceAccelerationStructureDescriptorClass) Descriptor() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ic.class), objc.Sel("descriptor"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=Descriptor) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for InstanceAccelerationStructureDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for InstanceAccelerationStructureDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for InstanceAccelerationStructureDescriptor */

// The number of instances in the instance descriptor buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLInstanceAccelerationStructureDescriptor/instanceCount
func (i_ InstanceAccelerationStructureDescriptor) InstanceCount() uint {
	rv := objc.Send[uint](i_.ID, objc.Sel("instanceCount"))
	return rv
}/* debug [instance_properties/getter]: instanceCount */


// The number of instances in the instance descriptor buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLInstanceAccelerationStructureDescriptor/instanceCount
func (i_ InstanceAccelerationStructureDescriptor) SetInstanceCount(value uint) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setInstanceCount:"), value)
}/* debug [instance_properties/setter]: instanceCount */


// The bottom-level acceleration structures that instances use in the instance acceleration structure .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLInstanceAccelerationStructureDescriptor/instancedAccelerationStructures
func (i_ InstanceAccelerationStructureDescriptor) InstancedAccelerationStructures() []objc.ID {
	rv := objc.Send[[]objc.ID](i_.ID, objc.Sel("instancedAccelerationStructures"))
	return rv
}/* debug [instance_properties/getter]: instancedAccelerationStructures */


// The bottom-level acceleration structures that instances use in the instance acceleration structure .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLInstanceAccelerationStructureDescriptor/instancedAccelerationStructures
func (i_ InstanceAccelerationStructureDescriptor) SetInstancedAccelerationStructures(value []objc.ID) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](i_.ID, objc.Sel("setInstancedAccelerationStructures:"), nsArray)
}/* debug [instance_properties/setter]: instancedAccelerationStructures */


// A buffer that contains descriptions of each instance in the acceleration structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLInstanceAccelerationStructureDescriptor/instanceDescriptorBuffer
func (i_ InstanceAccelerationStructureDescriptor) InstanceDescriptorBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("instanceDescriptorBuffer"))
	return rv
}/* debug [instance_properties/getter]: instanceDescriptorBuffer */


// A buffer that contains descriptions of each instance in the acceleration structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLInstanceAccelerationStructureDescriptor/instanceDescriptorBuffer
func (i_ InstanceAccelerationStructureDescriptor) SetInstanceDescriptorBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setInstanceDescriptorBuffer:"), value)
}/* debug [instance_properties/setter]: instanceDescriptorBuffer */


// The offset, in bytes, to the descripton of the first instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLInstanceAccelerationStructureDescriptor/instanceDescriptorBufferOffset
func (i_ InstanceAccelerationStructureDescriptor) InstanceDescriptorBufferOffset() uint {
	rv := objc.Send[uint](i_.ID, objc.Sel("instanceDescriptorBufferOffset"))
	return rv
}/* debug [instance_properties/getter]: instanceDescriptorBufferOffset */


// The offset, in bytes, to the descripton of the first instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLInstanceAccelerationStructureDescriptor/instanceDescriptorBufferOffset
func (i_ InstanceAccelerationStructureDescriptor) SetInstanceDescriptorBufferOffset(value uint) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setInstanceDescriptorBufferOffset:"), value)
}/* debug [instance_properties/setter]: instanceDescriptorBufferOffset */


// The stride, in bytes, between instance descriptions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLInstanceAccelerationStructureDescriptor/instanceDescriptorStride
func (i_ InstanceAccelerationStructureDescriptor) InstanceDescriptorStride() uint {
	rv := objc.Send[uint](i_.ID, objc.Sel("instanceDescriptorStride"))
	return rv
}/* debug [instance_properties/getter]: instanceDescriptorStride */


// The stride, in bytes, between instance descriptions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLInstanceAccelerationStructureDescriptor/instanceDescriptorStride
func (i_ InstanceAccelerationStructureDescriptor) SetInstanceDescriptorStride(value uint) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setInstanceDescriptorStride:"), value)
}/* debug [instance_properties/setter]: instanceDescriptorStride */


// The format of the instance data in the descriptor buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLInstanceAccelerationStructureDescriptor/instanceDescriptorType
func (i_ InstanceAccelerationStructureDescriptor) InstanceDescriptorType() AccelerationStructureInstanceDescriptorType {
	rv := objc.Send[AccelerationStructureInstanceDescriptorType](i_.ID, objc.Sel("instanceDescriptorType"))
	return rv
}/* debug [instance_properties/getter]: instanceDescriptorType */


// The format of the instance data in the descriptor buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLInstanceAccelerationStructureDescriptor/instanceDescriptorType
func (i_ InstanceAccelerationStructureDescriptor) SetInstanceDescriptorType(value AccelerationStructureInstanceDescriptorType) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setInstanceDescriptorType:"), value)
}/* debug [instance_properties/setter]: instanceDescriptorType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLInstanceAccelerationStructureDescriptor/instanceTransformationMatrixLayout
func (i_ InstanceAccelerationStructureDescriptor) InstanceTransformationMatrixLayout() MatrixLayout {
	rv := objc.Send[MatrixLayout](i_.ID, objc.Sel("instanceTransformationMatrixLayout"))
	return rv
}/* debug [instance_properties/getter]: instanceTransformationMatrixLayout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLInstanceAccelerationStructureDescriptor/instanceTransformationMatrixLayout
func (i_ InstanceAccelerationStructureDescriptor) SetInstanceTransformationMatrixLayout(value MatrixLayout) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setInstanceTransformationMatrixLayout:"), value)
}/* debug [instance_properties/setter]: instanceTransformationMatrixLayout */


// A buffer that contains descriptions of each motion transform in the acceleration structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLInstanceAccelerationStructureDescriptor/motionTransformBuffer
func (i_ InstanceAccelerationStructureDescriptor) MotionTransformBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("motionTransformBuffer"))
	return rv
}/* debug [instance_properties/getter]: motionTransformBuffer */


// A buffer that contains descriptions of each motion transform in the acceleration structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLInstanceAccelerationStructureDescriptor/motionTransformBuffer
func (i_ InstanceAccelerationStructureDescriptor) SetMotionTransformBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMotionTransformBuffer:"), value)
}/* debug [instance_properties/setter]: motionTransformBuffer */


// The offset, in bytes, to the descripton of the first motion transform.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLInstanceAccelerationStructureDescriptor/motionTransformBufferOffset
func (i_ InstanceAccelerationStructureDescriptor) MotionTransformBufferOffset() uint {
	rv := objc.Send[uint](i_.ID, objc.Sel("motionTransformBufferOffset"))
	return rv
}/* debug [instance_properties/getter]: motionTransformBufferOffset */


// The offset, in bytes, to the descripton of the first motion transform.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLInstanceAccelerationStructureDescriptor/motionTransformBufferOffset
func (i_ InstanceAccelerationStructureDescriptor) SetMotionTransformBufferOffset(value uint) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMotionTransformBufferOffset:"), value)
}/* debug [instance_properties/setter]: motionTransformBufferOffset */


// The number of motion transforms in the motion transform buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLInstanceAccelerationStructureDescriptor/motionTransformCount
func (i_ InstanceAccelerationStructureDescriptor) MotionTransformCount() uint {
	rv := objc.Send[uint](i_.ID, objc.Sel("motionTransformCount"))
	return rv
}/* debug [instance_properties/getter]: motionTransformCount */


// The number of motion transforms in the motion transform buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLInstanceAccelerationStructureDescriptor/motionTransformCount
func (i_ InstanceAccelerationStructureDescriptor) SetMotionTransformCount(value uint) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMotionTransformCount:"), value)
}/* debug [instance_properties/setter]: motionTransformCount */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLInstanceAccelerationStructureDescriptor/motionTransformStride
func (i_ InstanceAccelerationStructureDescriptor) MotionTransformStride() uint {
	rv := objc.Send[uint](i_.ID, objc.Sel("motionTransformStride"))
	return rv
}/* debug [instance_properties/getter]: motionTransformStride */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLInstanceAccelerationStructureDescriptor/motionTransformStride
func (i_ InstanceAccelerationStructureDescriptor) SetMotionTransformStride(value uint) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMotionTransformStride:"), value)
}/* debug [instance_properties/setter]: motionTransformStride */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLInstanceAccelerationStructureDescriptor/motionTransformType
func (i_ InstanceAccelerationStructureDescriptor) MotionTransformType() TransformType {
	rv := objc.Send[TransformType](i_.ID, objc.Sel("motionTransformType"))
	return rv
}/* debug [instance_properties/getter]: motionTransformType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLInstanceAccelerationStructureDescriptor/motionTransformType
func (i_ InstanceAccelerationStructureDescriptor) SetMotionTransformType(value TransformType) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMotionTransformType:"), value)
}/* debug [instance_properties/setter]: motionTransformType */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLInstanceAccelerationStructureDescriptor */



