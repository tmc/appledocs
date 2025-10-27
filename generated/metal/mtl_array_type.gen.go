// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [ArrayType] class.
var (
	ArrayTypeClass     _ArrayTypeClass
	ArrayTypeClassOnce sync.Once
)

func getArrayTypeClass() _ArrayTypeClass {
	ArrayTypeClassOnce.Do(func() {
		ArrayTypeClass = _ArrayTypeClass{objc.GetClass("MTLArrayType")}
	})
	return ArrayTypeClass
}

type _ArrayTypeClass struct {
	class objc.Class
}





// An interface definition for the [ArrayType] class.
type IArrayType interface {
	IType
	

	// properties:
	ArgumentIndexStride() uint
	ArrayLength() uint
	ElementType() DataType
	Stride() uint


	

	// methods:
	ElementArrayType() IArrayType
	ElementPointerType() IPointerType
	ElementStructType() IStructType
	ElementTensorReferenceType() ITensorReferenceType
	ElementTextureReferenceType() ITextureReferenceType


}





// Alloc allocates a new instance without initialization.
func (ac _ArrayTypeClass) Alloc() ArrayType {
	rv := objc.Send[ArrayType](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _ArrayTypeClass) New() ArrayType {
	rv := objc.Send[ArrayType](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ ArrayType) Init() ArrayType {
	rv := objc.Send[ArrayType](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ ArrayType) Autorelease() ArrayType {
	rv := objc.Send[ArrayType](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewArrayType creates a new ArrayType instance.
func NewArrayType() ArrayType {
	return getArrayTypeClass().New()
}





// A description of an array.
//
// An instance provides details about an array parameter. Don’t create instances directly; other reflection instances contain properties to determine if a parameter is an array and to obtain the instance that describes the array.


// A description of an array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArrayType
type ArrayType struct {
	Type
}

// ArrayTypeFrom constructs a [ArrayType] from an unsafe.Pointer.
//
// A description of an array.
func ArrayTypeFrom(ptr unsafe.Pointer) ArrayType {
	return ArrayType{
		Type: TypeFrom(ptr),
	}
}




















// Provides a description of the underlying type when an array holds other arrays as its elements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArrayType/element()
func (a_ ArrayType) ElementArrayType() IArrayType {
	rv := objc.Send[ArrayType](a_.ID, objc.Sel("elementArrayType"))
	return rv
}


// Provides a description of the underlying pointer type when an array holds pointers as its elements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArrayType/elementPointerType()
func (a_ ArrayType) ElementPointerType() IPointerType {
	rv := objc.Send[PointerType](a_.ID, objc.Sel("elementPointerType"))
	return rv
}


// Provides a description of the underlying struct type when an array holds structs as its elements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArrayType/elementStructType()
func (a_ ArrayType) ElementStructType() IStructType {
	rv := objc.Send[StructType](a_.ID, objc.Sel("elementStructType"))
	return rv
}


// Provides a description of the underlying tensor type when this array holds tensors as its elements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArrayType/elementTensorReferenceType()
func (a_ ArrayType) ElementTensorReferenceType() ITensorReferenceType {
	rv := objc.Send[TensorReferenceType](a_.ID, objc.Sel("elementTensorReferenceType"))
	return rv
}


// Provides a description of the underlying texture type when an array holds textures as its elements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArrayType/elementTextureReferenceType()
func (a_ ArrayType) ElementTextureReferenceType() ITextureReferenceType {
	rv := objc.Send[TextureReferenceType](a_.ID, objc.Sel("elementTextureReferenceType"))
	return rv
}







// The stride, in bytes, between argument indices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArrayType/argumentIndexStride
func (a_ ArrayType) ArgumentIndexStride() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("argumentIndexStride"))
	return rv
}


// The number of elements in the array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArrayType/arrayLength
func (a_ ArrayType) ArrayLength() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("arrayLength"))
	return rv
}


// The data type of the array’s elements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArrayType/elementType
func (a_ ArrayType) ElementType() DataType {
	rv := objc.Send[DataType](a_.ID, objc.Sel("elementType"))
	return rv
}


// The stride between array elements, in bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArrayType/stride
func (a_ ArrayType) Stride() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("stride"))
	return rv
}








