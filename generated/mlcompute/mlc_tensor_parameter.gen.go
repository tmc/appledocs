// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CTensorParameter] class.
var (
	CTensorParameterClass     _CTensorParameterClass
	CTensorParameterClassOnce sync.Once
)

func getCTensorParameterClass() _CTensorParameterClass {
	CTensorParameterClassOnce.Do(func() {
		CTensorParameterClass = _CTensorParameterClass{objc.GetClass("MLCTensorParameter")}
	})
	return CTensorParameterClass
}

type _CTensorParameterClass struct {
	class objc.Class
}

// An interface definition for the [CTensorParameter] class.
type ICTensorParameter interface {
	objectivec.IObject
}

// A tensor parameter object.
//
// Use a tensor parameter to describe input tensors that the optimizer updates during training.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensorParameter
type CTensorParameter struct {
	objectivec.Object
}

// CTensorParameterFrom constructs a [CTensorParameter] from an unsafe.Pointer.
//
// A tensor parameter object.
func CTensorParameterFrom(ptr unsafe.Pointer) CTensorParameter {
	return CTensorParameter{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CTensorParameterClass) Alloc() CTensorParameter {
	rv := objc.Send[CTensorParameter](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CTensorParameterClass) New() CTensorParameter {
	rv := objc.Send[CTensorParameter](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CTensorParameter) Init() CTensorParameter {
	rv := objc.Send[CTensorParameter](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CTensorParameter) Autorelease() CTensorParameter {
	rv := objc.Send[CTensorParameter](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCTensorParameter creates a new CTensorParameter instance.
func NewCTensorParameter() CTensorParameter {
	return getCTensorParameterClass().New()
}


// A Boolean that indicates whether this tensor parameter is updatable.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlctensorparameter/isupdatable
func (c_ CTensorParameter) IsUpdatable() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isUpdatable"))
	return rv
}


// SetIsUpdatable sets the value of the isUpdatable property.
// A Boolean that indicates whether this tensor parameter is updatable.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlctensorparameter/isupdatable
func (c_ CTensorParameter) SetIsUpdatable(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsUpdatable:"), value)
}

// The underlying tensor.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlctensorparameter/tensor
func (c_ CTensorParameter) Tensor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("tensor"))
	return rv
}


// SetTensor sets the value of the tensor property.
// The underlying tensor.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlctensorparameter/tensor
func (c_ CTensorParameter) SetTensor(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTensor:"), value)
}



