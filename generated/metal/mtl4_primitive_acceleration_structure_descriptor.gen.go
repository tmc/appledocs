// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MTL4PrimitiveAccelerationStructureDescriptor */


/* debug [class_header]: Header for MTL4PrimitiveAccelerationStructureDescriptor */
// The class instance for the [MTL4PrimitiveAccelerationStructureDescriptor] class.
var (
	MTL4PrimitiveAccelerationStructureDescriptorClass     _MTL4PrimitiveAccelerationStructureDescriptorClass
	MTL4PrimitiveAccelerationStructureDescriptorClassOnce sync.Once
)

func getMTL4PrimitiveAccelerationStructureDescriptorClass() _MTL4PrimitiveAccelerationStructureDescriptorClass {
	MTL4PrimitiveAccelerationStructureDescriptorClassOnce.Do(func() {
		MTL4PrimitiveAccelerationStructureDescriptorClass = _MTL4PrimitiveAccelerationStructureDescriptorClass{objc.GetClass("MTL4PrimitiveAccelerationStructureDescriptor")}
	})
	return MTL4PrimitiveAccelerationStructureDescriptorClass
}

type _MTL4PrimitiveAccelerationStructureDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTL4PrimitiveAccelerationStructureDescriptor */
// An interface definition for the [MTL4PrimitiveAccelerationStructureDescriptor] class.
type IMTL4PrimitiveAccelerationStructureDescriptor interface {
	IMTL4AccelerationStructureDescriptor
	
/* debug [class_interface_properties]: Properties for MTL4PrimitiveAccelerationStructureDescriptor */
	// properties:
	GeometryDescriptors() []MTL4AccelerationStructureGeometryDescriptor
	SetGeometryDescriptors(value []MTL4AccelerationStructureGeometryDescriptor)
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
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTL4PrimitiveAccelerationStructureDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTL4PrimitiveAccelerationStructureDescriptor */
// Alloc allocates a new instance without initialization.
func (mc _MTL4PrimitiveAccelerationStructureDescriptorClass) Alloc() MTL4PrimitiveAccelerationStructureDescriptor {
	rv := objc.Send[MTL4PrimitiveAccelerationStructureDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTL4PrimitiveAccelerationStructureDescriptorClass) New() MTL4PrimitiveAccelerationStructureDescriptor {
	rv := objc.Send[MTL4PrimitiveAccelerationStructureDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTL4PrimitiveAccelerationStructureDescriptor) Init() MTL4PrimitiveAccelerationStructureDescriptor {
	rv := objc.Send[MTL4PrimitiveAccelerationStructureDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTL4PrimitiveAccelerationStructureDescriptor) Autorelease() MTL4PrimitiveAccelerationStructureDescriptor {
	rv := objc.Send[MTL4PrimitiveAccelerationStructureDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4PrimitiveAccelerationStructureDescriptor creates a new MTL4PrimitiveAccelerationStructureDescriptor instance.
func NewMTL4PrimitiveAccelerationStructureDescriptor() MTL4PrimitiveAccelerationStructureDescriptor {
	return getMTL4PrimitiveAccelerationStructureDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTL4PrimitiveAccelerationStructureDescriptor */
// Descriptor for a primitive acceleration structure that directly references geometric shapes, such as triangles and bounding boxes.


// Descriptor for a primitive acceleration structure that directly references geometric shapes, such as triangles and bounding boxes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4PrimitiveAccelerationStructureDescriptor
type MTL4PrimitiveAccelerationStructureDescriptor struct {
	MTL4AccelerationStructureDescriptor
}

// MTL4PrimitiveAccelerationStructureDescriptorFrom constructs a [MTL4PrimitiveAccelerationStructureDescriptor] from an unsafe.Pointer.
//
// Descriptor for a primitive acceleration structure that directly references geometric shapes, such as triangles and bounding boxes.
func MTL4PrimitiveAccelerationStructureDescriptorFrom(ptr unsafe.Pointer) MTL4PrimitiveAccelerationStructureDescriptor {
	return MTL4PrimitiveAccelerationStructureDescriptor{
		MTL4AccelerationStructureDescriptor: MTL4AccelerationStructureDescriptorFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTL4PrimitiveAccelerationStructureDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTL4PrimitiveAccelerationStructureDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTL4PrimitiveAccelerationStructureDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTL4PrimitiveAccelerationStructureDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTL4PrimitiveAccelerationStructureDescriptor */

// Associates the array of geometry descriptors that comprise this primitive acceleration structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4PrimitiveAccelerationStructureDescriptor/geometryDescriptors
func (m_ MTL4PrimitiveAccelerationStructureDescriptor) GeometryDescriptors() []MTL4AccelerationStructureGeometryDescriptor {
	rv := objc.Send[[]MTL4AccelerationStructureGeometryDescriptor](m_.ID, objc.Sel("geometryDescriptors"))
	return rv
}/* debug [instance_properties/getter]: geometryDescriptors */


// Associates the array of geometry descriptors that comprise this primitive acceleration structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4PrimitiveAccelerationStructureDescriptor/geometryDescriptors
func (m_ MTL4PrimitiveAccelerationStructureDescriptor) SetGeometryDescriptors(value []MTL4AccelerationStructureGeometryDescriptor) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](m_.ID, objc.Sel("setGeometryDescriptors:"), nsArray)
}/* debug [instance_properties/setter]: geometryDescriptors */


// Configures the motion border mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4PrimitiveAccelerationStructureDescriptor/motionEndBorderMode
func (m_ MTL4PrimitiveAccelerationStructureDescriptor) MotionEndBorderMode() MotionBorderMode {
	rv := objc.Send[MotionBorderMode](m_.ID, objc.Sel("motionEndBorderMode"))
	return rv
}/* debug [instance_properties/getter]: motionEndBorderMode */


// Configures the motion border mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4PrimitiveAccelerationStructureDescriptor/motionEndBorderMode
func (m_ MTL4PrimitiveAccelerationStructureDescriptor) SetMotionEndBorderMode(value MotionBorderMode) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMotionEndBorderMode:"), value)
}/* debug [instance_properties/setter]: motionEndBorderMode */


// Configures the motion end time for this geometry.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4PrimitiveAccelerationStructureDescriptor/motionEndTime
func (m_ MTL4PrimitiveAccelerationStructureDescriptor) MotionEndTime() float32 {
	rv := objc.Send[float32](m_.ID, objc.Sel("motionEndTime"))
	return rv
}/* debug [instance_properties/getter]: motionEndTime */


// Configures the motion end time for this geometry.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4PrimitiveAccelerationStructureDescriptor/motionEndTime
func (m_ MTL4PrimitiveAccelerationStructureDescriptor) SetMotionEndTime(value float32) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMotionEndTime:"), value)
}/* debug [instance_properties/setter]: motionEndTime */


// Sets the motion keyframe count.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4PrimitiveAccelerationStructureDescriptor/motionKeyframeCount
func (m_ MTL4PrimitiveAccelerationStructureDescriptor) MotionKeyframeCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("motionKeyframeCount"))
	return rv
}/* debug [instance_properties/getter]: motionKeyframeCount */


// Sets the motion keyframe count.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4PrimitiveAccelerationStructureDescriptor/motionKeyframeCount
func (m_ MTL4PrimitiveAccelerationStructureDescriptor) SetMotionKeyframeCount(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMotionKeyframeCount:"), value)
}/* debug [instance_properties/setter]: motionKeyframeCount */


// Configures the behavior when the ray-tracing system samples the acceleration structure before the motion start time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4PrimitiveAccelerationStructureDescriptor/motionStartBorderMode
func (m_ MTL4PrimitiveAccelerationStructureDescriptor) MotionStartBorderMode() MotionBorderMode {
	rv := objc.Send[MotionBorderMode](m_.ID, objc.Sel("motionStartBorderMode"))
	return rv
}/* debug [instance_properties/getter]: motionStartBorderMode */


// Configures the behavior when the ray-tracing system samples the acceleration structure before the motion start time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4PrimitiveAccelerationStructureDescriptor/motionStartBorderMode
func (m_ MTL4PrimitiveAccelerationStructureDescriptor) SetMotionStartBorderMode(value MotionBorderMode) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMotionStartBorderMode:"), value)
}/* debug [instance_properties/setter]: motionStartBorderMode */


// Configures the motion start time for this geometry.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4PrimitiveAccelerationStructureDescriptor/motionStartTime
func (m_ MTL4PrimitiveAccelerationStructureDescriptor) MotionStartTime() float32 {
	rv := objc.Send[float32](m_.ID, objc.Sel("motionStartTime"))
	return rv
}/* debug [instance_properties/getter]: motionStartTime */


// Configures the motion start time for this geometry.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4PrimitiveAccelerationStructureDescriptor/motionStartTime
func (m_ MTL4PrimitiveAccelerationStructureDescriptor) SetMotionStartTime(value float32) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMotionStartTime:"), value)
}/* debug [instance_properties/setter]: motionStartTime */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTL4PrimitiveAccelerationStructureDescriptor */



