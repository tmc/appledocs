// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MTL4InstanceAccelerationStructureDescriptor */


/* debug [class_header]: Header for MTL4InstanceAccelerationStructureDescriptor */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTL4InstanceAccelerationStructureDescriptor */
// An interface definition for the [MTL4InstanceAccelerationStructureDescriptor] class.
type IMTL4InstanceAccelerationStructureDescriptor interface {
	IMTL4AccelerationStructureDescriptor
	
/* debug [class_interface_properties]: Properties for MTL4InstanceAccelerationStructureDescriptor */
	// properties:
	InstanceCount() uint
	SetInstanceCount(value uint)
	InstanceDescriptorBuffer() objc.IObject /* cross-framework: MTL4BufferRange */
	SetInstanceDescriptorBuffer(value objc.IObject /* cross-framework: MTL4BufferRange */)
	InstanceDescriptorStride() uint
	SetInstanceDescriptorStride(value uint)
	InstanceDescriptorType() AccelerationStructureInstanceDescriptorType
	SetInstanceDescriptorType(value AccelerationStructureInstanceDescriptorType)
	InstanceTransformationMatrixLayout() MatrixLayout
	SetInstanceTransformationMatrixLayout(value MatrixLayout)
	MotionTransformBuffer() objc.IObject /* cross-framework: MTL4BufferRange */
	SetMotionTransformBuffer(value objc.IObject /* cross-framework: MTL4BufferRange */)
	MotionTransformCount() uint
	SetMotionTransformCount(value uint)
	MotionTransformStride() uint
	SetMotionTransformStride(value uint)
	MotionTransformType() TransformType
	SetMotionTransformType(value TransformType)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTL4InstanceAccelerationStructureDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTL4InstanceAccelerationStructureDescriptor */
// Alloc allocates a new instance without initialization.
func (mc _MTL4InstanceAccelerationStructureDescriptorClass) Alloc() MTL4InstanceAccelerationStructureDescriptor {
	rv := objc.Send[MTL4InstanceAccelerationStructureDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTL4InstanceAccelerationStructureDescriptor */
// Descriptor for an instance acceleration structure.
//
// An instance acceleration structure references other acceleration structures, and provides the ability to “instantiate” them multiple times, each one with potentially a different transformation matrix. You specify the properties of the instances in the acceleration structure this descriptor builds by providing a buffer of via its property. Use a to mark residency of all buffers and acceleration structures this descriptor references when you build this acceleration structure.


// Descriptor for an instance acceleration structure.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTL4InstanceAccelerationStructureDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTL4InstanceAccelerationStructureDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTL4InstanceAccelerationStructureDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTL4InstanceAccelerationStructureDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTL4InstanceAccelerationStructureDescriptor */

// Controls the number of instance descriptors in the instance descriptor buffer references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4InstanceAccelerationStructureDescriptor/instanceCount
func (m_ MTL4InstanceAccelerationStructureDescriptor) InstanceCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("instanceCount"))
	return rv
}/* debug [instance_properties/getter]: instanceCount */


// Controls the number of instance descriptors in the instance descriptor buffer references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4InstanceAccelerationStructureDescriptor/instanceCount
func (m_ MTL4InstanceAccelerationStructureDescriptor) SetInstanceCount(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInstanceCount:"), value)
}/* debug [instance_properties/setter]: instanceCount */


// Assigns a reference to a buffer containing instance descriptors for acceleration structures to reference.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4InstanceAccelerationStructureDescriptor/instanceDescriptorBuffer
func (m_ MTL4InstanceAccelerationStructureDescriptor) InstanceDescriptorBuffer() objc.IObject /* cross-framework: MTL4BufferRange */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("instanceDescriptorBuffer"))
	return rv
}/* debug [instance_properties/getter]: instanceDescriptorBuffer */


// Assigns a reference to a buffer containing instance descriptors for acceleration structures to reference.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4InstanceAccelerationStructureDescriptor/instanceDescriptorBuffer
func (m_ MTL4InstanceAccelerationStructureDescriptor) SetInstanceDescriptorBuffer(value objc.IObject /* cross-framework: MTL4BufferRange */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInstanceDescriptorBuffer:"), value)
}/* debug [instance_properties/setter]: instanceDescriptorBuffer */


// Sets the stride, in bytes, between instance descriptors the instance descriptor buffer references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4InstanceAccelerationStructureDescriptor/instanceDescriptorStride
func (m_ MTL4InstanceAccelerationStructureDescriptor) InstanceDescriptorStride() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("instanceDescriptorStride"))
	return rv
}/* debug [instance_properties/getter]: instanceDescriptorStride */


// Sets the stride, in bytes, between instance descriptors the instance descriptor buffer references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4InstanceAccelerationStructureDescriptor/instanceDescriptorStride
func (m_ MTL4InstanceAccelerationStructureDescriptor) SetInstanceDescriptorStride(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInstanceDescriptorStride:"), value)
}/* debug [instance_properties/setter]: instanceDescriptorStride */


// Sets the type of instance descriptor that the instance descriptor buffer references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4InstanceAccelerationStructureDescriptor/instanceDescriptorType
func (m_ MTL4InstanceAccelerationStructureDescriptor) InstanceDescriptorType() AccelerationStructureInstanceDescriptorType {
	rv := objc.Send[AccelerationStructureInstanceDescriptorType](m_.ID, objc.Sel("instanceDescriptorType"))
	return rv
}/* debug [instance_properties/getter]: instanceDescriptorType */


// Sets the type of instance descriptor that the instance descriptor buffer references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4InstanceAccelerationStructureDescriptor/instanceDescriptorType
func (m_ MTL4InstanceAccelerationStructureDescriptor) SetInstanceDescriptorType(value AccelerationStructureInstanceDescriptorType) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInstanceDescriptorType:"), value)
}/* debug [instance_properties/setter]: instanceDescriptorType */


// Specifies the layout for the transformation matrices in the instance descriptor buffer and the motion transformation matrix buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4InstanceAccelerationStructureDescriptor/instanceTransformationMatrixLayout
func (m_ MTL4InstanceAccelerationStructureDescriptor) InstanceTransformationMatrixLayout() MatrixLayout {
	rv := objc.Send[MatrixLayout](m_.ID, objc.Sel("instanceTransformationMatrixLayout"))
	return rv
}/* debug [instance_properties/getter]: instanceTransformationMatrixLayout */


// Specifies the layout for the transformation matrices in the instance descriptor buffer and the motion transformation matrix buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4InstanceAccelerationStructureDescriptor/instanceTransformationMatrixLayout
func (m_ MTL4InstanceAccelerationStructureDescriptor) SetInstanceTransformationMatrixLayout(value MatrixLayout) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInstanceTransformationMatrixLayout:"), value)
}/* debug [instance_properties/setter]: instanceTransformationMatrixLayout */


// A buffer containing transformation information for instance motion keyframes, formatted according to the motion transform type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4InstanceAccelerationStructureDescriptor/motionTransformBuffer
func (m_ MTL4InstanceAccelerationStructureDescriptor) MotionTransformBuffer() objc.IObject /* cross-framework: MTL4BufferRange */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("motionTransformBuffer"))
	return rv
}/* debug [instance_properties/getter]: motionTransformBuffer */


// A buffer containing transformation information for instance motion keyframes, formatted according to the motion transform type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4InstanceAccelerationStructureDescriptor/motionTransformBuffer
func (m_ MTL4InstanceAccelerationStructureDescriptor) SetMotionTransformBuffer(value objc.IObject /* cross-framework: MTL4BufferRange */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMotionTransformBuffer:"), value)
}/* debug [instance_properties/setter]: motionTransformBuffer */


// Controls the total number of motion transforms in the motion transform buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4InstanceAccelerationStructureDescriptor/motionTransformCount
func (m_ MTL4InstanceAccelerationStructureDescriptor) MotionTransformCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("motionTransformCount"))
	return rv
}/* debug [instance_properties/getter]: motionTransformCount */


// Controls the total number of motion transforms in the motion transform buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4InstanceAccelerationStructureDescriptor/motionTransformCount
func (m_ MTL4InstanceAccelerationStructureDescriptor) SetMotionTransformCount(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMotionTransformCount:"), value)
}/* debug [instance_properties/setter]: motionTransformCount */


// Specify the stride for motion transform.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4InstanceAccelerationStructureDescriptor/motionTransformStride
func (m_ MTL4InstanceAccelerationStructureDescriptor) MotionTransformStride() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("motionTransformStride"))
	return rv
}/* debug [instance_properties/getter]: motionTransformStride */


// Specify the stride for motion transform.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4InstanceAccelerationStructureDescriptor/motionTransformStride
func (m_ MTL4InstanceAccelerationStructureDescriptor) SetMotionTransformStride(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMotionTransformStride:"), value)
}/* debug [instance_properties/setter]: motionTransformStride */


// Controls the type of motion transforms, either as a matrix or individual components.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4InstanceAccelerationStructureDescriptor/motionTransformType
func (m_ MTL4InstanceAccelerationStructureDescriptor) MotionTransformType() TransformType {
	rv := objc.Send[TransformType](m_.ID, objc.Sel("motionTransformType"))
	return rv
}/* debug [instance_properties/getter]: motionTransformType */


// Controls the type of motion transforms, either as a matrix or individual components.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4InstanceAccelerationStructureDescriptor/motionTransformType
func (m_ MTL4InstanceAccelerationStructureDescriptor) SetMotionTransformType(value TransformType) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMotionTransformType:"), value)
}/* debug [instance_properties/setter]: motionTransformType */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTL4InstanceAccelerationStructureDescriptor */



