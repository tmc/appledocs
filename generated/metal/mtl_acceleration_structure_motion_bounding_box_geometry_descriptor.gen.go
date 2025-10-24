// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLAccelerationStructureMotionBoundingBoxGeometryDescriptor */


/* debug [class_header]: Header for MTLAccelerationStructureMotionBoundingBoxGeometryDescriptor */
// The class instance for the [AccelerationStructureMotionBoundingBoxGeometryDescriptor] class.
var (
	AccelerationStructureMotionBoundingBoxGeometryDescriptorClass     _AccelerationStructureMotionBoundingBoxGeometryDescriptorClass
	AccelerationStructureMotionBoundingBoxGeometryDescriptorClassOnce sync.Once
)

func getAccelerationStructureMotionBoundingBoxGeometryDescriptorClass() _AccelerationStructureMotionBoundingBoxGeometryDescriptorClass {
	AccelerationStructureMotionBoundingBoxGeometryDescriptorClassOnce.Do(func() {
		AccelerationStructureMotionBoundingBoxGeometryDescriptorClass = _AccelerationStructureMotionBoundingBoxGeometryDescriptorClass{objc.GetClass("MTLAccelerationStructureMotionBoundingBoxGeometryDescriptor")}
	})
	return AccelerationStructureMotionBoundingBoxGeometryDescriptorClass
}

type _AccelerationStructureMotionBoundingBoxGeometryDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AccelerationStructureMotionBoundingBoxGeometryDescriptor */
// An interface definition for the [AccelerationStructureMotionBoundingBoxGeometryDescriptor] class.
type IAccelerationStructureMotionBoundingBoxGeometryDescriptor interface {
	IAccelerationStructureGeometryDescriptor
	
/* debug [class_interface_properties]: Properties for AccelerationStructureMotionBoundingBoxGeometryDescriptor */
	// properties:
	BoundingBoxBuffers() []MotionKeyframeData
	SetBoundingBoxBuffers(value []MotionKeyframeData)
	BoundingBoxCount() uint
	SetBoundingBoxCount(value uint)
	BoundingBoxStride() uint
	SetBoundingBoxStride(value uint)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AccelerationStructureMotionBoundingBoxGeometryDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AccelerationStructureMotionBoundingBoxGeometryDescriptor */
// Alloc allocates a new instance without initialization.
func (ac _AccelerationStructureMotionBoundingBoxGeometryDescriptorClass) Alloc() AccelerationStructureMotionBoundingBoxGeometryDescriptor {
	rv := objc.Send[AccelerationStructureMotionBoundingBoxGeometryDescriptor](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AccelerationStructureMotionBoundingBoxGeometryDescriptorClass) New() AccelerationStructureMotionBoundingBoxGeometryDescriptor {
	rv := objc.Send[AccelerationStructureMotionBoundingBoxGeometryDescriptor](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AccelerationStructureMotionBoundingBoxGeometryDescriptor) Init() AccelerationStructureMotionBoundingBoxGeometryDescriptor {
	rv := objc.Send[AccelerationStructureMotionBoundingBoxGeometryDescriptor](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AccelerationStructureMotionBoundingBoxGeometryDescriptor) Autorelease() AccelerationStructureMotionBoundingBoxGeometryDescriptor {
	rv := objc.Send[AccelerationStructureMotionBoundingBoxGeometryDescriptor](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAccelerationStructureMotionBoundingBoxGeometryDescriptor creates a new AccelerationStructureMotionBoundingBoxGeometryDescriptor instance.
func NewAccelerationStructureMotionBoundingBoxGeometryDescriptor() AccelerationStructureMotionBoundingBoxGeometryDescriptor {
	return getAccelerationStructureMotionBoundingBoxGeometryDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AccelerationStructureMotionBoundingBoxGeometryDescriptor */
// A description of a list of bounding boxes, as motion keyframe data, to turn into an acceleration structure.


// A description of a list of bounding boxes, as motion keyframe data, to turn into an acceleration structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionBoundingBoxGeometryDescriptor
type AccelerationStructureMotionBoundingBoxGeometryDescriptor struct {
	AccelerationStructureGeometryDescriptor
}

// AccelerationStructureMotionBoundingBoxGeometryDescriptorFrom constructs a [AccelerationStructureMotionBoundingBoxGeometryDescriptor] from an unsafe.Pointer.
//
// A description of a list of bounding boxes, as motion keyframe data, to turn into an acceleration structure.
func AccelerationStructureMotionBoundingBoxGeometryDescriptorFrom(ptr unsafe.Pointer) AccelerationStructureMotionBoundingBoxGeometryDescriptor {
	return AccelerationStructureMotionBoundingBoxGeometryDescriptor{
		AccelerationStructureGeometryDescriptor: AccelerationStructureGeometryDescriptorFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AccelerationStructureMotionBoundingBoxGeometryDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AccelerationStructureMotionBoundingBoxGeometryDescriptor */

// Creates a new bounding box descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionBoundingBoxGeometryDescriptor/descriptor
func (ac _AccelerationStructureMotionBoundingBoxGeometryDescriptorClass) Descriptor() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ac.class), objc.Sel("descriptor"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=Descriptor) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AccelerationStructureMotionBoundingBoxGeometryDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AccelerationStructureMotionBoundingBoxGeometryDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AccelerationStructureMotionBoundingBoxGeometryDescriptor */

// A array of motion keyframes, each containing bounding box data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionBoundingBoxGeometryDescriptor/boundingBoxBuffers
func (a_ AccelerationStructureMotionBoundingBoxGeometryDescriptor) BoundingBoxBuffers() []MotionKeyframeData {
	rv := objc.Send[[]MotionKeyframeData](a_.ID, objc.Sel("boundingBoxBuffers"))
	return rv
}/* debug [instance_properties/getter]: boundingBoxBuffers */


// A array of motion keyframes, each containing bounding box data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionBoundingBoxGeometryDescriptor/boundingBoxBuffers
func (a_ AccelerationStructureMotionBoundingBoxGeometryDescriptor) SetBoundingBoxBuffers(value []MotionKeyframeData) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](a_.ID, objc.Sel("setBoundingBoxBuffers:"), nsArray)
}/* debug [instance_properties/setter]: boundingBoxBuffers */


// The number of bounding boxes in each bounding box buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionBoundingBoxGeometryDescriptor/boundingBoxCount
func (a_ AccelerationStructureMotionBoundingBoxGeometryDescriptor) BoundingBoxCount() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("boundingBoxCount"))
	return rv
}/* debug [instance_properties/getter]: boundingBoxCount */


// The number of bounding boxes in each bounding box buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionBoundingBoxGeometryDescriptor/boundingBoxCount
func (a_ AccelerationStructureMotionBoundingBoxGeometryDescriptor) SetBoundingBoxCount(value uint) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setBoundingBoxCount:"), value)
}/* debug [instance_properties/setter]: boundingBoxCount */


// The stride, in bytes, between bounding boxes in each buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionBoundingBoxGeometryDescriptor/boundingBoxStride
func (a_ AccelerationStructureMotionBoundingBoxGeometryDescriptor) BoundingBoxStride() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("boundingBoxStride"))
	return rv
}/* debug [instance_properties/getter]: boundingBoxStride */


// The stride, in bytes, between bounding boxes in each buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionBoundingBoxGeometryDescriptor/boundingBoxStride
func (a_ AccelerationStructureMotionBoundingBoxGeometryDescriptor) SetBoundingBoxStride(value uint) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setBoundingBoxStride:"), value)
}/* debug [instance_properties/setter]: boundingBoxStride */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLAccelerationStructureMotionBoundingBoxGeometryDescriptor */



