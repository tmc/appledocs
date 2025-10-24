// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MTL4AccelerationStructureBoundingBoxGeometryDescriptor */


/* debug [class_header]: Header for MTL4AccelerationStructureBoundingBoxGeometryDescriptor */
// The class instance for the [MTL4AccelerationStructureBoundingBoxGeometryDescriptor] class.
var (
	MTL4AccelerationStructureBoundingBoxGeometryDescriptorClass     _MTL4AccelerationStructureBoundingBoxGeometryDescriptorClass
	MTL4AccelerationStructureBoundingBoxGeometryDescriptorClassOnce sync.Once
)

func getMTL4AccelerationStructureBoundingBoxGeometryDescriptorClass() _MTL4AccelerationStructureBoundingBoxGeometryDescriptorClass {
	MTL4AccelerationStructureBoundingBoxGeometryDescriptorClassOnce.Do(func() {
		MTL4AccelerationStructureBoundingBoxGeometryDescriptorClass = _MTL4AccelerationStructureBoundingBoxGeometryDescriptorClass{objc.GetClass("MTL4AccelerationStructureBoundingBoxGeometryDescriptor")}
	})
	return MTL4AccelerationStructureBoundingBoxGeometryDescriptorClass
}

type _MTL4AccelerationStructureBoundingBoxGeometryDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTL4AccelerationStructureBoundingBoxGeometryDescriptor */
// An interface definition for the [MTL4AccelerationStructureBoundingBoxGeometryDescriptor] class.
type IMTL4AccelerationStructureBoundingBoxGeometryDescriptor interface {
	IMTL4AccelerationStructureGeometryDescriptor
	
/* debug [class_interface_properties]: Properties for MTL4AccelerationStructureBoundingBoxGeometryDescriptor */
	// properties:
	BoundingBoxBuffer() objc.IObject /* cross-framework: MTL4BufferRange */
	SetBoundingBoxBuffer(value objc.IObject /* cross-framework: MTL4BufferRange */)
	BoundingBoxCount() uint
	SetBoundingBoxCount(value uint)
	BoundingBoxStride() uint
	SetBoundingBoxStride(value uint)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTL4AccelerationStructureBoundingBoxGeometryDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTL4AccelerationStructureBoundingBoxGeometryDescriptor */
// Alloc allocates a new instance without initialization.
func (mc _MTL4AccelerationStructureBoundingBoxGeometryDescriptorClass) Alloc() MTL4AccelerationStructureBoundingBoxGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureBoundingBoxGeometryDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTL4AccelerationStructureBoundingBoxGeometryDescriptorClass) New() MTL4AccelerationStructureBoundingBoxGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureBoundingBoxGeometryDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTL4AccelerationStructureBoundingBoxGeometryDescriptor) Init() MTL4AccelerationStructureBoundingBoxGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureBoundingBoxGeometryDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTL4AccelerationStructureBoundingBoxGeometryDescriptor) Autorelease() MTL4AccelerationStructureBoundingBoxGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureBoundingBoxGeometryDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4AccelerationStructureBoundingBoxGeometryDescriptor creates a new MTL4AccelerationStructureBoundingBoxGeometryDescriptor instance.
func NewMTL4AccelerationStructureBoundingBoxGeometryDescriptor() MTL4AccelerationStructureBoundingBoxGeometryDescriptor {
	return getMTL4AccelerationStructureBoundingBoxGeometryDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTL4AccelerationStructureBoundingBoxGeometryDescriptor */
// Describes bounding-box geometry suitable for ray tracing.
//
// You use bounding boxes to implement procedural geometry for ray tracing, such as spheres or any other shape you define by using intersection functions. Use a to mark residency of all buffers this descriptor references when you build this acceleration structure.


// Describes bounding-box geometry suitable for ray tracing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureBoundingBoxGeometryDescriptor
type MTL4AccelerationStructureBoundingBoxGeometryDescriptor struct {
	MTL4AccelerationStructureGeometryDescriptor
}

// MTL4AccelerationStructureBoundingBoxGeometryDescriptorFrom constructs a [MTL4AccelerationStructureBoundingBoxGeometryDescriptor] from an unsafe.Pointer.
//
// Describes bounding-box geometry suitable for ray tracing.
func MTL4AccelerationStructureBoundingBoxGeometryDescriptorFrom(ptr unsafe.Pointer) MTL4AccelerationStructureBoundingBoxGeometryDescriptor {
	return MTL4AccelerationStructureBoundingBoxGeometryDescriptor{
		MTL4AccelerationStructureGeometryDescriptor: MTL4AccelerationStructureGeometryDescriptorFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTL4AccelerationStructureBoundingBoxGeometryDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTL4AccelerationStructureBoundingBoxGeometryDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTL4AccelerationStructureBoundingBoxGeometryDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTL4AccelerationStructureBoundingBoxGeometryDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTL4AccelerationStructureBoundingBoxGeometryDescriptor */

// References a buffer containing bounding box data in format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureBoundingBoxGeometryDescriptor/boundingBoxBuffer
func (m_ MTL4AccelerationStructureBoundingBoxGeometryDescriptor) BoundingBoxBuffer() objc.IObject /* cross-framework: MTL4BufferRange */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("boundingBoxBuffer"))
	return rv
}/* debug [instance_properties/getter]: boundingBoxBuffer */


// References a buffer containing bounding box data in format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureBoundingBoxGeometryDescriptor/boundingBoxBuffer
func (m_ MTL4AccelerationStructureBoundingBoxGeometryDescriptor) SetBoundingBoxBuffer(value objc.IObject /* cross-framework: MTL4BufferRange */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBoundingBoxBuffer:"), value)
}/* debug [instance_properties/setter]: boundingBoxBuffer */


// Describes the number of bounding boxes the contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureBoundingBoxGeometryDescriptor/boundingBoxCount
func (m_ MTL4AccelerationStructureBoundingBoxGeometryDescriptor) BoundingBoxCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("boundingBoxCount"))
	return rv
}/* debug [instance_properties/getter]: boundingBoxCount */


// Describes the number of bounding boxes the contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureBoundingBoxGeometryDescriptor/boundingBoxCount
func (m_ MTL4AccelerationStructureBoundingBoxGeometryDescriptor) SetBoundingBoxCount(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBoundingBoxCount:"), value)
}/* debug [instance_properties/setter]: boundingBoxCount */


// Assigns the stride, in bytes, between bounding boxes in the bounding box buffer references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureBoundingBoxGeometryDescriptor/boundingBoxStride
func (m_ MTL4AccelerationStructureBoundingBoxGeometryDescriptor) BoundingBoxStride() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("boundingBoxStride"))
	return rv
}/* debug [instance_properties/getter]: boundingBoxStride */


// Assigns the stride, in bytes, between bounding boxes in the bounding box buffer references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureBoundingBoxGeometryDescriptor/boundingBoxStride
func (m_ MTL4AccelerationStructureBoundingBoxGeometryDescriptor) SetBoundingBoxStride(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBoundingBoxStride:"), value)
}/* debug [instance_properties/setter]: boundingBoxStride */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTL4AccelerationStructureBoundingBoxGeometryDescriptor */



