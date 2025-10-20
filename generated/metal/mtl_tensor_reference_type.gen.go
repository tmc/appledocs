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
func (t_ TensorReferenceType) TensorDataType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("tensorDataType"))
	return rv
}



