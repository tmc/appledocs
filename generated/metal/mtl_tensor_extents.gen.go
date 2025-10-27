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
	

	// properties:
	Rank() uint
	MTLTensorDomain() foundation.foundation.INSString
	Extents() int
	SetExtents(value int)
	MTL_TENSOR_MAX_RANK() objectivec.IObject
	SetMTL_TENSOR_MAX_RANK(value objectivec.IObject)


	

	// methods:
	ExtentAtDimensionIndex(dimensionIndex uint) int


}





// Alloc allocates a new instance without initialization.
func (tc _TensorExtentsClass) Alloc() TensorExtents {
	rv := objc.Send[TensorExtents](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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





// An array of length matching the rank, holding the dimensions of a tensor.
//
// Supports rank up to .


// An array of length matching the rank, holding the dimensions of a tensor.
//
// [Full Topic]
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






// Creates a new tensor extents with the rank and extent values you provide.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTensorExtents/initWithRank:values:
func NewTensorExtentsWithRankValues(rank uint, values int) TensorExtents {
	instance := getTensorExtentsClass().Alloc()
	rv := objc.Send[TensorExtents](instance.ID, objc.Sel("initWithRank:values:"), rank, values)
	rv.Autorelease()
	return rv
}

















// Returns the extent at an index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTensorExtents/extentAtDimensionIndex:
func (t_ TensorExtents) ExtentAtDimensionIndex(dimensionIndex uint) int {
	rv := objc.Send[int](t_.ID, objc.Sel("extentAtDimensionIndex:"), dimensionIndex)
	return rv
}







// Obtains the rank of the tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTensorExtents/rank
func (t_ TensorExtents) Rank() uint {
	rv := objc.Send[uint](t_.ID, objc.Sel("rank"))
	return rv
}


// An error domain for errors that pertain to creating a tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltensordomain
func (t_ TensorExtents) MTLTensorDomain() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("MTLTensorDomain"))
	return rv
}


// Retrieves the extents for this object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltensorextents/extents
func (t_ TensorExtents) Extents() int {
	rv := objc.Send[int](t_.ID, objc.Sel("extents"))
	return rv
}


// Retrieves the extents for this object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltensorextents/extents
func (t_ TensorExtents) SetExtents(value int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setExtents:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl_tensor_max_rank
func (t_ TensorExtents) MTL_TENSOR_MAX_RANK() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](t_.ID, objc.Sel("MTL_TENSOR_MAX_RANK"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl_tensor_max_rank
func (t_ TensorExtents) SetMTL_TENSOR_MAX_RANK(value objectivec.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setMTL_TENSOR_MAX_RANK:"), value)
}







