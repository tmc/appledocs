// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MTL4IndirectInstanceAccelerationStructureDescriptor */


/* debug [class_header]: Header for MTL4IndirectInstanceAccelerationStructureDescriptor */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTL4IndirectInstanceAccelerationStructureDescriptor */
// An interface definition for the [MTL4IndirectInstanceAccelerationStructureDescriptor] class.
type IMTL4IndirectInstanceAccelerationStructureDescriptor interface {
	IMTL4AccelerationStructureDescriptor
	
/* debug [class_interface_properties]: Properties for MTL4IndirectInstanceAccelerationStructureDescriptor */
	// properties:
	InstanceCountBuffer() objc.IObject /* cross-framework: MTL4BufferRange */
	SetInstanceCountBuffer(value objc.IObject /* cross-framework: MTL4BufferRange */)
	InstanceDescriptorBuffer() objc.IObject /* cross-framework: MTL4BufferRange */
	SetInstanceDescriptorBuffer(value objc.IObject /* cross-framework: MTL4BufferRange */)
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
	MotionTransformBuffer() objc.IObject /* cross-framework: MTL4BufferRange */
	SetMotionTransformBuffer(value objc.IObject /* cross-framework: MTL4BufferRange */)
	MotionTransformCountBuffer() objc.IObject /* cross-framework: MTL4BufferRange */
	SetMotionTransformCountBuffer(value objc.IObject /* cross-framework: MTL4BufferRange */)
	MotionTransformStride() uint
	SetMotionTransformStride(value uint)
	MotionTransformType() TransformType
	SetMotionTransformType(value TransformType)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTL4IndirectInstanceAccelerationStructureDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTL4IndirectInstanceAccelerationStructureDescriptor */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTL4IndirectInstanceAccelerationStructureDescriptor */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTL4IndirectInstanceAccelerationStructureDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTL4IndirectInstanceAccelerationStructureDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTL4IndirectInstanceAccelerationStructureDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTL4IndirectInstanceAccelerationStructureDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTL4IndirectInstanceAccelerationStructureDescriptor */

// Provides a reference to a buffer containing the number of instances in the instance descriptor buffer, formatted as a 32-bit unsigned integer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4IndirectInstanceAccelerationStructureDescriptor/instanceCountBuffer
func (m_ MTL4IndirectInstanceAccelerationStructureDescriptor) InstanceCountBuffer() objc.IObject /* cross-framework: MTL4BufferRange */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("instanceCountBuffer"))
	return rv
}/* debug [instance_properties/getter]: instanceCountBuffer */


// Provides a reference to a buffer containing the number of instances in the instance descriptor buffer, formatted as a 32-bit unsigned integer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4IndirectInstanceAccelerationStructureDescriptor/instanceCountBuffer
func (m_ MTL4IndirectInstanceAccelerationStructureDescriptor) SetInstanceCountBuffer(value objc.IObject /* cross-framework: MTL4BufferRange */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInstanceCountBuffer:"), value)
}/* debug [instance_properties/setter]: instanceCountBuffer */


// Assigns a reference to a buffer containing instance descriptors for acceleration structures to reference.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4IndirectInstanceAccelerationStructureDescriptor/instanceDescriptorBuffer
func (m_ MTL4IndirectInstanceAccelerationStructureDescriptor) InstanceDescriptorBuffer() objc.IObject /* cross-framework: MTL4BufferRange */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("instanceDescriptorBuffer"))
	return rv
}/* debug [instance_properties/getter]: instanceDescriptorBuffer */


// Assigns a reference to a buffer containing instance descriptors for acceleration structures to reference.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4IndirectInstanceAccelerationStructureDescriptor/instanceDescriptorBuffer
func (m_ MTL4IndirectInstanceAccelerationStructureDescriptor) SetInstanceDescriptorBuffer(value objc.IObject /* cross-framework: MTL4BufferRange */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInstanceDescriptorBuffer:"), value)
}/* debug [instance_properties/setter]: instanceDescriptorBuffer */


// Sets the stride, in bytes, between instance descriptors in the instance descriptor buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4IndirectInstanceAccelerationStructureDescriptor/instanceDescriptorStride
func (m_ MTL4IndirectInstanceAccelerationStructureDescriptor) InstanceDescriptorStride() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("instanceDescriptorStride"))
	return rv
}/* debug [instance_properties/getter]: instanceDescriptorStride */


// Sets the stride, in bytes, between instance descriptors in the instance descriptor buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4IndirectInstanceAccelerationStructureDescriptor/instanceDescriptorStride
func (m_ MTL4IndirectInstanceAccelerationStructureDescriptor) SetInstanceDescriptorStride(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInstanceDescriptorStride:"), value)
}/* debug [instance_properties/setter]: instanceDescriptorStride */


// Controls the type of instance descriptor that the instance descriptor buffer references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4IndirectInstanceAccelerationStructureDescriptor/instanceDescriptorType
func (m_ MTL4IndirectInstanceAccelerationStructureDescriptor) InstanceDescriptorType() AccelerationStructureInstanceDescriptorType {
	rv := objc.Send[AccelerationStructureInstanceDescriptorType](m_.ID, objc.Sel("instanceDescriptorType"))
	return rv
}/* debug [instance_properties/getter]: instanceDescriptorType */


// Controls the type of instance descriptor that the instance descriptor buffer references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4IndirectInstanceAccelerationStructureDescriptor/instanceDescriptorType
func (m_ MTL4IndirectInstanceAccelerationStructureDescriptor) SetInstanceDescriptorType(value AccelerationStructureInstanceDescriptorType) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInstanceDescriptorType:"), value)
}/* debug [instance_properties/setter]: instanceDescriptorType */


// Specifies the layout for the transformation matrices in the instance descriptor buffer and the motion transformation matrix buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4IndirectInstanceAccelerationStructureDescriptor/instanceTransformationMatrixLayout
func (m_ MTL4IndirectInstanceAccelerationStructureDescriptor) InstanceTransformationMatrixLayout() MatrixLayout {
	rv := objc.Send[MatrixLayout](m_.ID, objc.Sel("instanceTransformationMatrixLayout"))
	return rv
}/* debug [instance_properties/getter]: instanceTransformationMatrixLayout */


// Specifies the layout for the transformation matrices in the instance descriptor buffer and the motion transformation matrix buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4IndirectInstanceAccelerationStructureDescriptor/instanceTransformationMatrixLayout
func (m_ MTL4IndirectInstanceAccelerationStructureDescriptor) SetInstanceTransformationMatrixLayout(value MatrixLayout) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInstanceTransformationMatrixLayout:"), value)
}/* debug [instance_properties/setter]: instanceTransformationMatrixLayout */


// Controls the maximum number of instance descriptors the instance descriptor buffer can reference.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4IndirectInstanceAccelerationStructureDescriptor/maxInstanceCount
func (m_ MTL4IndirectInstanceAccelerationStructureDescriptor) MaxInstanceCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("maxInstanceCount"))
	return rv
}/* debug [instance_properties/getter]: maxInstanceCount */


// Controls the maximum number of instance descriptors the instance descriptor buffer can reference.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4IndirectInstanceAccelerationStructureDescriptor/maxInstanceCount
func (m_ MTL4IndirectInstanceAccelerationStructureDescriptor) SetMaxInstanceCount(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxInstanceCount:"), value)
}/* debug [instance_properties/setter]: maxInstanceCount */


// Controls the maximum number of motion transforms in the motion transform buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4IndirectInstanceAccelerationStructureDescriptor/maxMotionTransformCount
func (m_ MTL4IndirectInstanceAccelerationStructureDescriptor) MaxMotionTransformCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("maxMotionTransformCount"))
	return rv
}/* debug [instance_properties/getter]: maxMotionTransformCount */


// Controls the maximum number of motion transforms in the motion transform buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4IndirectInstanceAccelerationStructureDescriptor/maxMotionTransformCount
func (m_ MTL4IndirectInstanceAccelerationStructureDescriptor) SetMaxMotionTransformCount(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxMotionTransformCount:"), value)
}/* debug [instance_properties/setter]: maxMotionTransformCount */


// A buffer containing transformation information for instance motion keyframes, formatted according to the motion transform type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4IndirectInstanceAccelerationStructureDescriptor/motionTransformBuffer
func (m_ MTL4IndirectInstanceAccelerationStructureDescriptor) MotionTransformBuffer() objc.IObject /* cross-framework: MTL4BufferRange */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("motionTransformBuffer"))
	return rv
}/* debug [instance_properties/getter]: motionTransformBuffer */


// A buffer containing transformation information for instance motion keyframes, formatted according to the motion transform type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4IndirectInstanceAccelerationStructureDescriptor/motionTransformBuffer
func (m_ MTL4IndirectInstanceAccelerationStructureDescriptor) SetMotionTransformBuffer(value objc.IObject /* cross-framework: MTL4BufferRange */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMotionTransformBuffer:"), value)
}/* debug [instance_properties/setter]: motionTransformBuffer */


// Associates a buffer reference containing the number of motion transforms in the motion transform buffer, formatted as a 32-bit unsigned integer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4IndirectInstanceAccelerationStructureDescriptor/motionTransformCountBuffer
func (m_ MTL4IndirectInstanceAccelerationStructureDescriptor) MotionTransformCountBuffer() objc.IObject /* cross-framework: MTL4BufferRange */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("motionTransformCountBuffer"))
	return rv
}/* debug [instance_properties/getter]: motionTransformCountBuffer */


// Associates a buffer reference containing the number of motion transforms in the motion transform buffer, formatted as a 32-bit unsigned integer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4IndirectInstanceAccelerationStructureDescriptor/motionTransformCountBuffer
func (m_ MTL4IndirectInstanceAccelerationStructureDescriptor) SetMotionTransformCountBuffer(value objc.IObject /* cross-framework: MTL4BufferRange */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMotionTransformCountBuffer:"), value)
}/* debug [instance_properties/setter]: motionTransformCountBuffer */


// Sets the stride for motion transform.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4IndirectInstanceAccelerationStructureDescriptor/motionTransformStride
func (m_ MTL4IndirectInstanceAccelerationStructureDescriptor) MotionTransformStride() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("motionTransformStride"))
	return rv
}/* debug [instance_properties/getter]: motionTransformStride */


// Sets the stride for motion transform.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4IndirectInstanceAccelerationStructureDescriptor/motionTransformStride
func (m_ MTL4IndirectInstanceAccelerationStructureDescriptor) SetMotionTransformStride(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMotionTransformStride:"), value)
}/* debug [instance_properties/setter]: motionTransformStride */


// Sets the type of motion transforms, either as a matrix or individual components.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4IndirectInstanceAccelerationStructureDescriptor/motionTransformType
func (m_ MTL4IndirectInstanceAccelerationStructureDescriptor) MotionTransformType() TransformType {
	rv := objc.Send[TransformType](m_.ID, objc.Sel("motionTransformType"))
	return rv
}/* debug [instance_properties/getter]: motionTransformType */


// Sets the type of motion transforms, either as a matrix or individual components.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4IndirectInstanceAccelerationStructureDescriptor/motionTransformType
func (m_ MTL4IndirectInstanceAccelerationStructureDescriptor) SetMotionTransformType(value TransformType) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMotionTransformType:"), value)
}/* debug [instance_properties/setter]: motionTransformType */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTL4IndirectInstanceAccelerationStructureDescriptor */



