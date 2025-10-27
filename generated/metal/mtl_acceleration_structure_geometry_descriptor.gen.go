// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [AccelerationStructureGeometryDescriptor] class.
var (
	AccelerationStructureGeometryDescriptorClass     _AccelerationStructureGeometryDescriptorClass
	AccelerationStructureGeometryDescriptorClassOnce sync.Once
)

func getAccelerationStructureGeometryDescriptorClass() _AccelerationStructureGeometryDescriptorClass {
	AccelerationStructureGeometryDescriptorClassOnce.Do(func() {
		AccelerationStructureGeometryDescriptorClass = _AccelerationStructureGeometryDescriptorClass{objc.GetClass("MTLAccelerationStructureGeometryDescriptor")}
	})
	return AccelerationStructureGeometryDescriptorClass
}

type _AccelerationStructureGeometryDescriptorClass struct {
	class objc.Class
}





// An interface definition for the [AccelerationStructureGeometryDescriptor] class.
type IAccelerationStructureGeometryDescriptor interface {
	objectivec.IObject
	

	// properties:
	AllowDuplicateIntersectionFunctionInvocation() bool
	SetAllowDuplicateIntersectionFunctionInvocation(value bool)
	IntersectionFunctionTableOffset() uint
	SetIntersectionFunctionTableOffset(value uint)
	Label() foundation.foundation.INSString
	SetLabel(value foundation.foundation.INSString)
	Opaque() bool
	SetOpaque(value bool)
	PrimitiveDataBuffer() unsafe.Pointer
	SetPrimitiveDataBuffer(value unsafe.Pointer)
	PrimitiveDataBufferOffset() uint
	SetPrimitiveDataBufferOffset(value uint)
	PrimitiveDataElementSize() uint
	SetPrimitiveDataElementSize(value uint)
	PrimitiveDataStride() uint
	SetPrimitiveDataStride(value uint)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (ac _AccelerationStructureGeometryDescriptorClass) Alloc() AccelerationStructureGeometryDescriptor {
	rv := objc.Send[AccelerationStructureGeometryDescriptor](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AccelerationStructureGeometryDescriptorClass) New() AccelerationStructureGeometryDescriptor {
	rv := objc.Send[AccelerationStructureGeometryDescriptor](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AccelerationStructureGeometryDescriptor) Init() AccelerationStructureGeometryDescriptor {
	rv := objc.Send[AccelerationStructureGeometryDescriptor](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AccelerationStructureGeometryDescriptor) Autorelease() AccelerationStructureGeometryDescriptor {
	rv := objc.Send[AccelerationStructureGeometryDescriptor](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAccelerationStructureGeometryDescriptor creates a new AccelerationStructureGeometryDescriptor instance.
func NewAccelerationStructureGeometryDescriptor() AccelerationStructureGeometryDescriptor {
	return getAccelerationStructureGeometryDescriptorClass().New()
}





// A base class for descriptors that contain geometry data to convert into a ray-tracing acceleration structure.
//
// Don’t use this base class directly. Use one of the derived classes instead, as describes.


// A base class for descriptors that contain geometry data to convert into a ray-tracing acceleration structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureGeometryDescriptor
type AccelerationStructureGeometryDescriptor struct {
	objectivec.Object
}

// AccelerationStructureGeometryDescriptorFrom constructs a [AccelerationStructureGeometryDescriptor] from an unsafe.Pointer.
//
// A base class for descriptors that contain geometry data to convert into a ray-tracing acceleration structure.
func AccelerationStructureGeometryDescriptorFrom(ptr unsafe.Pointer) AccelerationStructureGeometryDescriptor {
	return AccelerationStructureGeometryDescriptor{objectivec.Object{objc.ID(ptr)}}
}

























// A Boolean value that indicates whether Metal calls the ray-intersection test more than once per primitive on the structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureGeometryDescriptor/allowDuplicateIntersectionFunctionInvocation
func (a_ AccelerationStructureGeometryDescriptor) AllowDuplicateIntersectionFunctionInvocation() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("allowDuplicateIntersectionFunctionInvocation"))
	return rv
}


// A Boolean value that indicates whether Metal calls the ray-intersection test more than once per primitive on the structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureGeometryDescriptor/allowDuplicateIntersectionFunctionInvocation
func (a_ AccelerationStructureGeometryDescriptor) SetAllowDuplicateIntersectionFunctionInvocation(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAllowDuplicateIntersectionFunctionInvocation:"), value)
}


// An index into the intersection table for determining which intersection function Metal calls when it intersects a ray with the acceleration structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureGeometryDescriptor/intersectionFunctionTableOffset
func (a_ AccelerationStructureGeometryDescriptor) IntersectionFunctionTableOffset() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("intersectionFunctionTableOffset"))
	return rv
}


// An index into the intersection table for determining which intersection function Metal calls when it intersects a ray with the acceleration structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureGeometryDescriptor/intersectionFunctionTableOffset
func (a_ AccelerationStructureGeometryDescriptor) SetIntersectionFunctionTableOffset(value uint) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIntersectionFunctionTableOffset:"), value)
}


// A label for the geometry structure, suitable for debugging.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureGeometryDescriptor/label
func (a_ AccelerationStructureGeometryDescriptor) Label() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("label"))
	return rv
}


// A label for the geometry structure, suitable for debugging.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureGeometryDescriptor/label
func (a_ AccelerationStructureGeometryDescriptor) SetLabel(value foundation.foundation.INSString) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setLabel:"), value)
}


// A Boolean value that determines whether the geometry data in the acceleration structure needs to skip triangle-intersection tests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureGeometryDescriptor/opaque
func (a_ AccelerationStructureGeometryDescriptor) Opaque() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("opaque"))
	return rv
}


// A Boolean value that determines whether the geometry data in the acceleration structure needs to skip triangle-intersection tests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureGeometryDescriptor/opaque
func (a_ AccelerationStructureGeometryDescriptor) SetOpaque(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setOpaque:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureGeometryDescriptor/primitiveDataBuffer
func (a_ AccelerationStructureGeometryDescriptor) PrimitiveDataBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("primitiveDataBuffer"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureGeometryDescriptor/primitiveDataBuffer
func (a_ AccelerationStructureGeometryDescriptor) SetPrimitiveDataBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPrimitiveDataBuffer:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureGeometryDescriptor/primitiveDataBufferOffset
func (a_ AccelerationStructureGeometryDescriptor) PrimitiveDataBufferOffset() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("primitiveDataBufferOffset"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureGeometryDescriptor/primitiveDataBufferOffset
func (a_ AccelerationStructureGeometryDescriptor) SetPrimitiveDataBufferOffset(value uint) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPrimitiveDataBufferOffset:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureGeometryDescriptor/primitiveDataElementSize
func (a_ AccelerationStructureGeometryDescriptor) PrimitiveDataElementSize() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("primitiveDataElementSize"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureGeometryDescriptor/primitiveDataElementSize
func (a_ AccelerationStructureGeometryDescriptor) SetPrimitiveDataElementSize(value uint) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPrimitiveDataElementSize:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureGeometryDescriptor/primitiveDataStride
func (a_ AccelerationStructureGeometryDescriptor) PrimitiveDataStride() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("primitiveDataStride"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureGeometryDescriptor/primitiveDataStride
func (a_ AccelerationStructureGeometryDescriptor) SetPrimitiveDataStride(value uint) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPrimitiveDataStride:"), value)
}








