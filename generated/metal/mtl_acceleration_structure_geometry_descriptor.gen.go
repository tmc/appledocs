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
}

// A base class for descriptors that contain geometry data to convert into a ray-tracing acceleration structure.
//
// Don’t use this base class directly. Use one of the derived classes instead, as describes.
//
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

// Alloc allocates a new instance without initialization.
func (ac _AccelerationStructureGeometryDescriptorClass) Alloc() AccelerationStructureGeometryDescriptor {
	rv := objc.Send[AccelerationStructureGeometryDescriptor](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// A Boolean value that indicates whether Metal calls the ray-intersection test more than once per primitive on the structure.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuregeometrydescriptor/allowduplicateintersectionfunctioninvocation
func (a_ AccelerationStructureGeometryDescriptor) AllowDuplicateIntersectionFunctionInvocation() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("allowDuplicateIntersectionFunctionInvocation"))
	return rv
}


// SetAllowDuplicateIntersectionFunctionInvocation sets the value of the allowDuplicateIntersectionFunctionInvocation property.
// A Boolean value that indicates whether Metal calls the ray-intersection test more than once per primitive on the structure.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuregeometrydescriptor/allowduplicateintersectionfunctioninvocation
func (a_ AccelerationStructureGeometryDescriptor) SetAllowDuplicateIntersectionFunctionInvocation(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAllowDuplicateIntersectionFunctionInvocation:"), value)
}

// An index into the intersection table for determining which intersection function Metal calls when it intersects a ray with the acceleration structure.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuregeometrydescriptor/intersectionfunctiontableoffset
func (a_ AccelerationStructureGeometryDescriptor) IntersectionFunctionTableOffset() int {
	rv := objc.Send[int](a_.ID, objc.Sel("intersectionFunctionTableOffset"))
	return rv
}


// SetIntersectionFunctionTableOffset sets the value of the intersectionFunctionTableOffset property.
// An index into the intersection table for determining which intersection function Metal calls when it intersects a ray with the acceleration structure.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuregeometrydescriptor/intersectionfunctiontableoffset
func (a_ AccelerationStructureGeometryDescriptor) SetIntersectionFunctionTableOffset(value int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIntersectionFunctionTableOffset:"), value)
}

// A label for the geometry structure, suitable for debugging.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuregeometrydescriptor/label
func (a_ AccelerationStructureGeometryDescriptor) Label() string {
	rv := objc.Send[string](a_.ID, objc.Sel("label"))
	return rv
}


// SetLabel sets the value of the label property.
// A label for the geometry structure, suitable for debugging.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuregeometrydescriptor/label
func (a_ AccelerationStructureGeometryDescriptor) SetLabel(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setLabel:"), objc.String(value))
}

// A Boolean value that determines whether the geometry data in the acceleration structure needs to skip triangle-intersection tests.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuregeometrydescriptor/opaque
func (a_ AccelerationStructureGeometryDescriptor) Opaque() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("opaque"))
	return rv
}


// SetOpaque sets the value of the opaque property.
// A Boolean value that determines whether the geometry data in the acceleration structure needs to skip triangle-intersection tests.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuregeometrydescriptor/opaque
func (a_ AccelerationStructureGeometryDescriptor) SetOpaque(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setOpaque:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuregeometrydescriptor/primitivedatabuffer
func (a_ AccelerationStructureGeometryDescriptor) PrimitiveDataBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("primitiveDataBuffer"))
	return rv
}


// SetPrimitiveDataBuffer sets the value of the primitiveDataBuffer property.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuregeometrydescriptor/primitivedatabuffer
func (a_ AccelerationStructureGeometryDescriptor) SetPrimitiveDataBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPrimitiveDataBuffer:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuregeometrydescriptor/primitivedatabufferoffset
func (a_ AccelerationStructureGeometryDescriptor) PrimitiveDataBufferOffset() int {
	rv := objc.Send[int](a_.ID, objc.Sel("primitiveDataBufferOffset"))
	return rv
}


// SetPrimitiveDataBufferOffset sets the value of the primitiveDataBufferOffset property.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuregeometrydescriptor/primitivedatabufferoffset
func (a_ AccelerationStructureGeometryDescriptor) SetPrimitiveDataBufferOffset(value int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPrimitiveDataBufferOffset:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuregeometrydescriptor/primitivedataelementsize
func (a_ AccelerationStructureGeometryDescriptor) PrimitiveDataElementSize() int {
	rv := objc.Send[int](a_.ID, objc.Sel("primitiveDataElementSize"))
	return rv
}


// SetPrimitiveDataElementSize sets the value of the primitiveDataElementSize property.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuregeometrydescriptor/primitivedataelementsize
func (a_ AccelerationStructureGeometryDescriptor) SetPrimitiveDataElementSize(value int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPrimitiveDataElementSize:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuregeometrydescriptor/primitivedatastride
func (a_ AccelerationStructureGeometryDescriptor) PrimitiveDataStride() int {
	rv := objc.Send[int](a_.ID, objc.Sel("primitiveDataStride"))
	return rv
}


// SetPrimitiveDataStride sets the value of the primitiveDataStride property.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuregeometrydescriptor/primitivedatastride
func (a_ AccelerationStructureGeometryDescriptor) SetPrimitiveDataStride(value int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPrimitiveDataStride:"), value)
}



