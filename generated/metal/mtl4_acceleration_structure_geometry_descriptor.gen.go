// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [MTL4AccelerationStructureGeometryDescriptor] class.
var (
	MTL4AccelerationStructureGeometryDescriptorClass     _MTL4AccelerationStructureGeometryDescriptorClass
	MTL4AccelerationStructureGeometryDescriptorClassOnce sync.Once
)

func getMTL4AccelerationStructureGeometryDescriptorClass() _MTL4AccelerationStructureGeometryDescriptorClass {
	MTL4AccelerationStructureGeometryDescriptorClassOnce.Do(func() {
		MTL4AccelerationStructureGeometryDescriptorClass = _MTL4AccelerationStructureGeometryDescriptorClass{objc.GetClass("MTL4AccelerationStructureGeometryDescriptor")}
	})
	return MTL4AccelerationStructureGeometryDescriptorClass
}

type _MTL4AccelerationStructureGeometryDescriptorClass struct {
	class objc.Class
}





// An interface definition for the [MTL4AccelerationStructureGeometryDescriptor] class.
type IMTL4AccelerationStructureGeometryDescriptor interface {
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
	PrimitiveDataBuffer() MTL4BufferRange
	SetPrimitiveDataBuffer(value MTL4BufferRange)
	PrimitiveDataElementSize() uint
	SetPrimitiveDataElementSize(value uint)
	PrimitiveDataStride() uint
	SetPrimitiveDataStride(value uint)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (mc _MTL4AccelerationStructureGeometryDescriptorClass) Alloc() MTL4AccelerationStructureGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureGeometryDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTL4AccelerationStructureGeometryDescriptorClass) New() MTL4AccelerationStructureGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureGeometryDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTL4AccelerationStructureGeometryDescriptor) Init() MTL4AccelerationStructureGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureGeometryDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTL4AccelerationStructureGeometryDescriptor) Autorelease() MTL4AccelerationStructureGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureGeometryDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4AccelerationStructureGeometryDescriptor creates a new MTL4AccelerationStructureGeometryDescriptor instance.
func NewMTL4AccelerationStructureGeometryDescriptor() MTL4AccelerationStructureGeometryDescriptor {
	return getMTL4AccelerationStructureGeometryDescriptorClass().New()
}





// Base class for all Metal 4 acceleration structure geometry descriptors.
//
// Don’t use this class directly. Use one of the derived classes instead.


// Base class for all Metal 4 acceleration structure geometry descriptors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureGeometryDescriptor
type MTL4AccelerationStructureGeometryDescriptor struct {
	objectivec.Object
}

// MTL4AccelerationStructureGeometryDescriptorFrom constructs a [MTL4AccelerationStructureGeometryDescriptor] from an unsafe.Pointer.
//
// Base class for all Metal 4 acceleration structure geometry descriptors.
func MTL4AccelerationStructureGeometryDescriptorFrom(ptr unsafe.Pointer) MTL4AccelerationStructureGeometryDescriptor {
	return MTL4AccelerationStructureGeometryDescriptor{objectivec.Object{objc.ID(ptr)}}
}

























// A boolean value that indicates whether the ray-tracing system in Metal allows the invocation of intersection functions more than once per ray-primitive intersection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureGeometryDescriptor/allowDuplicateIntersectionFunctionInvocation
func (m_ MTL4AccelerationStructureGeometryDescriptor) AllowDuplicateIntersectionFunctionInvocation() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("allowDuplicateIntersectionFunctionInvocation"))
	return rv
}


// A boolean value that indicates whether the ray-tracing system in Metal allows the invocation of intersection functions more than once per ray-primitive intersection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureGeometryDescriptor/allowDuplicateIntersectionFunctionInvocation
func (m_ MTL4AccelerationStructureGeometryDescriptor) SetAllowDuplicateIntersectionFunctionInvocation(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAllowDuplicateIntersectionFunctionInvocation:"), value)
}


// Sets the offset that this geometry contributes to determining the intersection function to invoke when a ray intersects it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureGeometryDescriptor/intersectionFunctionTableOffset
func (m_ MTL4AccelerationStructureGeometryDescriptor) IntersectionFunctionTableOffset() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("intersectionFunctionTableOffset"))
	return rv
}


// Sets the offset that this geometry contributes to determining the intersection function to invoke when a ray intersects it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureGeometryDescriptor/intersectionFunctionTableOffset
func (m_ MTL4AccelerationStructureGeometryDescriptor) SetIntersectionFunctionTableOffset(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIntersectionFunctionTableOffset:"), value)
}


// Assigns an optional label you can assign to this geometry for debugging purposes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureGeometryDescriptor/label
func (m_ MTL4AccelerationStructureGeometryDescriptor) Label() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("label"))
	return rv
}


// Assigns an optional label you can assign to this geometry for debugging purposes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureGeometryDescriptor/label
func (m_ MTL4AccelerationStructureGeometryDescriptor) SetLabel(value foundation.foundation.INSString) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLabel:"), value)
}


// Provides a hint to Metal that this geometry is opaque, potentially accelerating the ray/primitive intersection process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureGeometryDescriptor/opaque
func (m_ MTL4AccelerationStructureGeometryDescriptor) Opaque() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("opaque"))
	return rv
}


// Provides a hint to Metal that this geometry is opaque, potentially accelerating the ray/primitive intersection process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureGeometryDescriptor/opaque
func (m_ MTL4AccelerationStructureGeometryDescriptor) SetOpaque(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOpaque:"), value)
}


// Assigns optional buffer containing data to associate with each primitive in this geometry.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureGeometryDescriptor/primitiveDataBuffer
func (m_ MTL4AccelerationStructureGeometryDescriptor) PrimitiveDataBuffer() MTL4BufferRange {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("primitiveDataBuffer"))
	return rv
}


// Assigns optional buffer containing data to associate with each primitive in this geometry.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureGeometryDescriptor/primitiveDataBuffer
func (m_ MTL4AccelerationStructureGeometryDescriptor) SetPrimitiveDataBuffer(value MTL4BufferRange) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPrimitiveDataBuffer:"), value)
}


// Sets the size, in bytes, of the data for each primitive in the primitive data buffer references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureGeometryDescriptor/primitiveDataElementSize
func (m_ MTL4AccelerationStructureGeometryDescriptor) PrimitiveDataElementSize() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("primitiveDataElementSize"))
	return rv
}


// Sets the size, in bytes, of the data for each primitive in the primitive data buffer references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureGeometryDescriptor/primitiveDataElementSize
func (m_ MTL4AccelerationStructureGeometryDescriptor) SetPrimitiveDataElementSize(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPrimitiveDataElementSize:"), value)
}


// Defines the stride, in bytes, between each primitive’s data in the primitive data buffer references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureGeometryDescriptor/primitiveDataStride
func (m_ MTL4AccelerationStructureGeometryDescriptor) PrimitiveDataStride() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("primitiveDataStride"))
	return rv
}


// Defines the stride, in bytes, between each primitive’s data in the primitive data buffer references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureGeometryDescriptor/primitiveDataStride
func (m_ MTL4AccelerationStructureGeometryDescriptor) SetPrimitiveDataStride(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPrimitiveDataStride:"), value)
}








