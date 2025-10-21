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
	objectivec.IObject
}

// An object that represents a tensor in the shading language in a struct or array.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTensorReferenceType
type TensorReferenceType struct {
	objectivec.Object
}

// TensorReferenceTypeFrom constructs a [TensorReferenceType] from an unsafe.Pointer.
//
// An object that represents a tensor in the shading language in a struct or array.
func TensorReferenceTypeFrom(ptr unsafe.Pointer) TensorReferenceType {
	return TensorReferenceType{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TensorReferenceTypeClass) Alloc() TensorReferenceType {
	rv := objc.Send[TensorReferenceType](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The underlying data format of the tensor.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTensorReferenceType/tensorDataType
func (t_ TensorReferenceType) TensorDataType() TensorDataType {
	rv := objc.Send[TensorDataType](t_.ID, objc.Sel("tensorDataType"))
	return rv
}

// An error domain for errors that pertain to creating a tensor.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltensordomain
func (t_ TensorReferenceType) MTLTensorDomain() appkit.string {
	rv := objc.Send[appkit.string](t_.ID, objc.Sel("MTLTensorDomain"))
	return rv
}

// A value that represents the read/write permissions of the tensor.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltensorreferencetype/access
func (t_ TensorReferenceType) Access() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("access"))
	return rv
}


// SetAccess sets the value of the access property.
// A value that represents the read/write permissions of the tensor.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltensorreferencetype/access
func (t_ TensorReferenceType) SetAccess(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAccess:"), value)
}

// The array of sizes, in elements, one for each dimension of this tensor.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltensorreferencetype/dimensions
func (t_ TensorReferenceType) Dimensions() MTLTensorExtents {
	rv := objc.Send[MTLTensorExtents](t_.ID, objc.Sel("dimensions"))
	return rv
}


// SetDimensions sets the value of the dimensions property.
// The array of sizes, in elements, one for each dimension of this tensor.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltensorreferencetype/dimensions
func (t_ TensorReferenceType) SetDimensions(value IMTLTensorExtents) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDimensions:"), value)
}

// The data format you use for indexing into the tensor.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltensorreferencetype/indextype
func (t_ TensorReferenceType) IndexType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("indexType"))
	return rv
}


// SetIndexType sets the value of the indexType property.
// The data format you use for indexing into the tensor.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltensorreferencetype/indextype
func (t_ TensorReferenceType) SetIndexType(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIndexType:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl_tensor_max_rank
func (t_ TensorReferenceType) MTL_TENSOR_MAX_RANK() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("MTL_TENSOR_MAX_RANK"))
	return rv
}


// SetMTL_TENSOR_MAX_RANK sets the value of the MTL_TENSOR_MAX_RANK property.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl_tensor_max_rank
func (t_ TensorReferenceType) SetMTL_TENSOR_MAX_RANK(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setMTL_TENSOR_MAX_RANK:"), value)
}



