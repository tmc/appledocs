// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TensorExtents] class.
var (
	TensorExtentsClass     _TensorExtentsClass
	TensorExtentsClassOnce sync.Once
)

func getTensorExtentsClass() _TensorExtentsClass {
	TensorExtentsClassOnce.Do(func() {
		TensorExtentsClass = _TensorExtentsClass{objc.GetClass("MTLTensorExtents")}
	})
	return TensorExtentsClass
}

type _TensorExtentsClass struct {
	class objc.Class
}

// An interface definition for the [TensorExtents] class.
type ITensorExtents interface {
	objectivec.IObject
}

// An array of length matching the rank, holding the dimensions of a tensor.
//
// Supports rank up to .
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTensorExtents
type TensorExtents struct {
	objectivec.Object
}

// TensorExtentsFrom constructs a [TensorExtents] from an unsafe.Pointer.
//
// An array of length matching the rank, holding the dimensions of a tensor.
func TensorExtentsFrom(ptr unsafe.Pointer) TensorExtents {
	return TensorExtents{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TensorExtentsClass) Alloc() TensorExtents {
	rv := objc.Send[TensorExtents](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TensorExtentsClass) New() TensorExtents {
	rv := objc.Send[TensorExtents](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TensorExtents) Init() TensorExtents {
	rv := objc.Send[TensorExtents](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TensorExtents) Autorelease() TensorExtents {
	rv := objc.Send[TensorExtents](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTensorExtents creates a new TensorExtents instance.
func NewTensorExtents() TensorExtents {
	return getTensorExtentsClass().New()
}


// Retrieves the extents for this object.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltensorextents/extents
func (t_ TensorExtents) Extents() int {
	rv := objc.Send[int](t_.ID, objc.Sel("extents"))
	return rv
}


// SetExtents sets the value of the extents property.
// Retrieves the extents for this object.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltensorextents/extents
func (t_ TensorExtents) SetExtents(value int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setExtents:"), value)
}

// An error domain for errors that pertain to creating a tensor.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltensordomain
func (t_ TensorExtents) MTLTensorDomain() string {
	rv := objc.Send[string](t_.ID, objc.Sel("MTLTensorDomain"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl_tensor_max_rank
func (t_ TensorExtents) MTL_TENSOR_MAX_RANK() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("MTL_TENSOR_MAX_RANK"))
	return rv
}


// SetMTL_TENSOR_MAX_RANK sets the value of the MTL_TENSOR_MAX_RANK property.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl_tensor_max_rank
func (t_ TensorExtents) SetMTL_TENSOR_MAX_RANK(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setMTL_TENSOR_MAX_RANK:"), value)
}

// Obtains the rank of the tensor.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltensorextents/rank
func (t_ TensorExtents) Rank() int {
	rv := objc.Send[int](t_.ID, objc.Sel("rank"))
	return rv
}


// SetRank sets the value of the rank property.
// Obtains the rank of the tensor.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltensorextents/rank
func (t_ TensorExtents) SetRank(value int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setRank:"), value)
}



