// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [AccelerationStructureBoundingBoxGeometryDescriptor] class.
var (
	AccelerationStructureBoundingBoxGeometryDescriptorClass     _AccelerationStructureBoundingBoxGeometryDescriptorClass
	AccelerationStructureBoundingBoxGeometryDescriptorClassOnce sync.Once
)

func getAccelerationStructureBoundingBoxGeometryDescriptorClass() _AccelerationStructureBoundingBoxGeometryDescriptorClass {
	AccelerationStructureBoundingBoxGeometryDescriptorClassOnce.Do(func() {
		AccelerationStructureBoundingBoxGeometryDescriptorClass = _AccelerationStructureBoundingBoxGeometryDescriptorClass{objc.GetClass("MTLAccelerationStructureBoundingBoxGeometryDescriptor")}
	})
	return AccelerationStructureBoundingBoxGeometryDescriptorClass
}

type _AccelerationStructureBoundingBoxGeometryDescriptorClass struct {
	class objc.Class
}





// An interface definition for the [AccelerationStructureBoundingBoxGeometryDescriptor] class.
type IAccelerationStructureBoundingBoxGeometryDescriptor interface {
	IAccelerationStructureGeometryDescriptor
	

	// properties:
	BoundingBoxBuffer() unsafe.Pointer
	SetBoundingBoxBuffer(value unsafe.Pointer)
	BoundingBoxBufferOffset() uint
	SetBoundingBoxBufferOffset(value uint)
	BoundingBoxCount() uint
	SetBoundingBoxCount(value uint)
	BoundingBoxStride() uint
	SetBoundingBoxStride(value uint)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (ac _AccelerationStructureBoundingBoxGeometryDescriptorClass) Alloc() AccelerationStructureBoundingBoxGeometryDescriptor {
	rv := objc.Send[AccelerationStructureBoundingBoxGeometryDescriptor](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AccelerationStructureBoundingBoxGeometryDescriptorClass) New() AccelerationStructureBoundingBoxGeometryDescriptor {
	rv := objc.Send[AccelerationStructureBoundingBoxGeometryDescriptor](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AccelerationStructureBoundingBoxGeometryDescriptor) Init() AccelerationStructureBoundingBoxGeometryDescriptor {
	rv := objc.Send[AccelerationStructureBoundingBoxGeometryDescriptor](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AccelerationStructureBoundingBoxGeometryDescriptor) Autorelease() AccelerationStructureBoundingBoxGeometryDescriptor {
	rv := objc.Send[AccelerationStructureBoundingBoxGeometryDescriptor](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAccelerationStructureBoundingBoxGeometryDescriptor creates a new AccelerationStructureBoundingBoxGeometryDescriptor instance.
func NewAccelerationStructureBoundingBoxGeometryDescriptor() AccelerationStructureBoundingBoxGeometryDescriptor {
	return getAccelerationStructureBoundingBoxGeometryDescriptorClass().New()
}





// A description of a list of bounding boxes to turn into an acceleration structure.


// A description of a list of bounding boxes to turn into an acceleration structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureBoundingBoxGeometryDescriptor
type AccelerationStructureBoundingBoxGeometryDescriptor struct {
	AccelerationStructureGeometryDescriptor
}

// AccelerationStructureBoundingBoxGeometryDescriptorFrom constructs a [AccelerationStructureBoundingBoxGeometryDescriptor] from an unsafe.Pointer.
//
// A description of a list of bounding boxes to turn into an acceleration structure.
func AccelerationStructureBoundingBoxGeometryDescriptorFrom(ptr unsafe.Pointer) AccelerationStructureBoundingBoxGeometryDescriptor {
	return AccelerationStructureBoundingBoxGeometryDescriptor{
		AccelerationStructureGeometryDescriptor: AccelerationStructureGeometryDescriptorFrom(ptr),
	}
}










// Creates a new bounding box descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureBoundingBoxGeometryDescriptor/descriptor
func (ac _AccelerationStructureBoundingBoxGeometryDescriptorClass) Descriptor() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ac.class), objc.Sel("descriptor"))
	return rv
}

















// A buffer that contains bounding box data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureBoundingBoxGeometryDescriptor/boundingBoxBuffer
func (a_ AccelerationStructureBoundingBoxGeometryDescriptor) BoundingBoxBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("boundingBoxBuffer"))
	return rv
}


// A buffer that contains bounding box data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureBoundingBoxGeometryDescriptor/boundingBoxBuffer
func (a_ AccelerationStructureBoundingBoxGeometryDescriptor) SetBoundingBoxBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setBoundingBoxBuffer:"), value)
}


// The offset, in bytes, to the first bounding box in the buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureBoundingBoxGeometryDescriptor/boundingBoxBufferOffset
func (a_ AccelerationStructureBoundingBoxGeometryDescriptor) BoundingBoxBufferOffset() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("boundingBoxBufferOffset"))
	return rv
}


// The offset, in bytes, to the first bounding box in the buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureBoundingBoxGeometryDescriptor/boundingBoxBufferOffset
func (a_ AccelerationStructureBoundingBoxGeometryDescriptor) SetBoundingBoxBufferOffset(value uint) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setBoundingBoxBufferOffset:"), value)
}


// The number of bounding boxes in the bounding box buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureBoundingBoxGeometryDescriptor/boundingBoxCount
func (a_ AccelerationStructureBoundingBoxGeometryDescriptor) BoundingBoxCount() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("boundingBoxCount"))
	return rv
}


// The number of bounding boxes in the bounding box buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureBoundingBoxGeometryDescriptor/boundingBoxCount
func (a_ AccelerationStructureBoundingBoxGeometryDescriptor) SetBoundingBoxCount(value uint) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setBoundingBoxCount:"), value)
}


// The stride, in bytes, between bounding boxes in the buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureBoundingBoxGeometryDescriptor/boundingBoxStride
func (a_ AccelerationStructureBoundingBoxGeometryDescriptor) BoundingBoxStride() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("boundingBoxStride"))
	return rv
}


// The stride, in bytes, between bounding boxes in the buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureBoundingBoxGeometryDescriptor/boundingBoxStride
func (a_ AccelerationStructureBoundingBoxGeometryDescriptor) SetBoundingBoxStride(value uint) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setBoundingBoxStride:"), value)
}








