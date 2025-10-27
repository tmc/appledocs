// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [TensorReferenceType] class.
var (
	TensorReferenceTypeClass     _TensorReferenceTypeClass
	TensorReferenceTypeClassOnce sync.Once
)

func getTensorReferenceTypeClass() _TensorReferenceTypeClass {
	TensorReferenceTypeClassOnce.Do(func() {
		TensorReferenceTypeClass = _TensorReferenceTypeClass{objc.GetClass("MTLTensorReferenceType")}
	})
	return TensorReferenceTypeClass
}

type _TensorReferenceTypeClass struct {
	class objc.Class
}





// An interface definition for the [TensorReferenceType] class.
type ITensorReferenceType interface {
	IType
	

	// properties:
	Access() BindingAccess
	Dimensions() IMTLTensorExtents
	IndexType() DataType
	TensorDataType() TensorDataType
	MTLTensorDomain() foundation.foundation.INSString
	MTL_TENSOR_MAX_RANK() objectivec.IObject
	SetMTL_TENSOR_MAX_RANK(value objectivec.IObject)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (tc _TensorReferenceTypeClass) Alloc() TensorReferenceType {
	rv := objc.Send[TensorReferenceType](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TensorReferenceTypeClass) New() TensorReferenceType {
	rv := objc.Send[TensorReferenceType](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TensorReferenceType) Init() TensorReferenceType {
	rv := objc.Send[TensorReferenceType](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TensorReferenceType) Autorelease() TensorReferenceType {
	rv := objc.Send[TensorReferenceType](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTensorReferenceType creates a new TensorReferenceType instance.
func NewTensorReferenceType() TensorReferenceType {
	return getTensorReferenceTypeClass().New()
}





// An object that represents a tensor in the shading language in a struct or array.


// An object that represents a tensor in the shading language in a struct or array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTensorReferenceType
type TensorReferenceType struct {
	Type
}

// TensorReferenceTypeFrom constructs a [TensorReferenceType] from an unsafe.Pointer.
//
// An object that represents a tensor in the shading language in a struct or array.
func TensorReferenceTypeFrom(ptr unsafe.Pointer) TensorReferenceType {
	return TensorReferenceType{
		Type: TypeFrom(ptr),
	}
}

























// A value that represents the read/write permissions of the tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTensorReferenceType/access
func (t_ TensorReferenceType) Access() BindingAccess {
	rv := objc.Send[BindingAccess](t_.ID, objc.Sel("access"))
	return rv
}


// The array of sizes, in elements, one for each dimension of this tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTensorReferenceType/dimensions
func (t_ TensorReferenceType) Dimensions() IMTLTensorExtents {
	rv := objc.Send[TensorExtents](t_.ID, objc.Sel("dimensions"))
	return rv
}


// The data format you use for indexing into the tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTensorReferenceType/indexType
func (t_ TensorReferenceType) IndexType() DataType {
	rv := objc.Send[DataType](t_.ID, objc.Sel("indexType"))
	return rv
}


// The underlying data format of the tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTensorReferenceType/tensorDataType
func (t_ TensorReferenceType) TensorDataType() TensorDataType {
	rv := objc.Send[TensorDataType](t_.ID, objc.Sel("tensorDataType"))
	return rv
}


// An error domain for errors that pertain to creating a tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltensordomain
func (t_ TensorReferenceType) MTLTensorDomain() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("MTLTensorDomain"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl_tensor_max_rank
func (t_ TensorReferenceType) MTL_TENSOR_MAX_RANK() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](t_.ID, objc.Sel("MTL_TENSOR_MAX_RANK"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl_tensor_max_rank
func (t_ TensorReferenceType) SetMTL_TENSOR_MAX_RANK(value objectivec.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setMTL_TENSOR_MAX_RANK:"), value)
}








