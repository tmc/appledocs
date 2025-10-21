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
}

// Base class for all Metal 4 acceleration structure geometry descriptors.
//
// Don’t use this class directly. Use one of the derived classes instead.
//
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

// Alloc allocates a new instance without initialization.
func (mc _MTL4AccelerationStructureGeometryDescriptorClass) Alloc() MTL4AccelerationStructureGeometryDescriptor {
	rv := objc.Send[MTL4AccelerationStructureGeometryDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Assigns optional buffer containing data to associate with each primitive in this geometry.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuregeometrydescriptor/primitivedatabuffer
func (m_ MTL4AccelerationStructureGeometryDescriptor) PrimitiveDataBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("primitiveDataBuffer"))
	return rv
}


// SetPrimitiveDataBuffer sets the value of the primitiveDataBuffer property.
// Assigns optional buffer containing data to associate with each primitive in this geometry.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuregeometrydescriptor/primitivedatabuffer
func (m_ MTL4AccelerationStructureGeometryDescriptor) SetPrimitiveDataBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPrimitiveDataBuffer:"), value)
}

// Sets the size, in bytes, of the data for each primitive in the primitive data buffer
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuregeometrydescriptor/primitivedataelementsize
func (m_ MTL4AccelerationStructureGeometryDescriptor) PrimitiveDataElementSize() int {
	rv := objc.Send[int](m_.ID, objc.Sel("primitiveDataElementSize"))
	return rv
}


// SetPrimitiveDataElementSize sets the value of the primitiveDataElementSize property.
// Sets the size, in bytes, of the data for each primitive in the primitive data buffer

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuregeometrydescriptor/primitivedataelementsize
func (m_ MTL4AccelerationStructureGeometryDescriptor) SetPrimitiveDataElementSize(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPrimitiveDataElementSize:"), value)
}

// Assigns an optional label you can assign to this geometry for debugging purposes.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuregeometrydescriptor/label
func (m_ MTL4AccelerationStructureGeometryDescriptor) Label() string {
	rv := objc.Send[string](m_.ID, objc.Sel("label"))
	return rv
}


// SetLabel sets the value of the label property.
// Assigns an optional label you can assign to this geometry for debugging purposes.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuregeometrydescriptor/label
func (m_ MTL4AccelerationStructureGeometryDescriptor) SetLabel(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLabel:"), objc.String(value))
}

// A boolean value that indicates whether the ray-tracing system in Metal allows the invocation of intersection functions
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuregeometrydescriptor/allowduplicateintersectionfunctioninvocation
func (m_ MTL4AccelerationStructureGeometryDescriptor) AllowDuplicateIntersectionFunctionInvocation() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("allowDuplicateIntersectionFunctionInvocation"))
	return rv
}


// SetAllowDuplicateIntersectionFunctionInvocation sets the value of the allowDuplicateIntersectionFunctionInvocation property.
// A boolean value that indicates whether the ray-tracing system in Metal allows the invocation of intersection functions

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuregeometrydescriptor/allowduplicateintersectionfunctioninvocation
func (m_ MTL4AccelerationStructureGeometryDescriptor) SetAllowDuplicateIntersectionFunctionInvocation(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAllowDuplicateIntersectionFunctionInvocation:"), value)
}

// Sets the offset that this geometry contributes to determining the intersection function to invoke when a ray intersects it.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuregeometrydescriptor/intersectionfunctiontableoffset
func (m_ MTL4AccelerationStructureGeometryDescriptor) IntersectionFunctionTableOffset() int {
	rv := objc.Send[int](m_.ID, objc.Sel("intersectionFunctionTableOffset"))
	return rv
}


// SetIntersectionFunctionTableOffset sets the value of the intersectionFunctionTableOffset property.
// Sets the offset that this geometry contributes to determining the intersection function to invoke when a ray intersects it.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuregeometrydescriptor/intersectionfunctiontableoffset
func (m_ MTL4AccelerationStructureGeometryDescriptor) SetIntersectionFunctionTableOffset(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIntersectionFunctionTableOffset:"), value)
}

// Provides a hint to Metal that this geometry is opaque, potentially accelerating the ray/primitive intersection process.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuregeometrydescriptor/opaque
func (m_ MTL4AccelerationStructureGeometryDescriptor) Opaque() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("opaque"))
	return rv
}


// SetOpaque sets the value of the opaque property.
// Provides a hint to Metal that this geometry is opaque, potentially accelerating the ray/primitive intersection process.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuregeometrydescriptor/opaque
func (m_ MTL4AccelerationStructureGeometryDescriptor) SetOpaque(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOpaque:"), value)
}

// Defines the stride, in bytes, between each primitive’s data in the primitive data buffer
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuregeometrydescriptor/primitivedatastride
func (m_ MTL4AccelerationStructureGeometryDescriptor) PrimitiveDataStride() int {
	rv := objc.Send[int](m_.ID, objc.Sel("primitiveDataStride"))
	return rv
}


// SetPrimitiveDataStride sets the value of the primitiveDataStride property.
// Defines the stride, in bytes, between each primitive’s data in the primitive data buffer

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4accelerationstructuregeometrydescriptor/primitivedatastride
func (m_ MTL4AccelerationStructureGeometryDescriptor) SetPrimitiveDataStride(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPrimitiveDataStride:"), value)
}



