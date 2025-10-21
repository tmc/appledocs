// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [NDArray] class.
var (
	NDArrayClass     _NDArrayClass
	NDArrayClassOnce sync.Once
)

func getNDArrayClass() _NDArrayClass {
	NDArrayClassOnce.Do(func() {
		NDArrayClass = _NDArrayClass{objc.GetClass("MPSNDArray")}
	})
	return NDArrayClass
}

type _NDArrayClass struct {
	class objc.Class
}

// An interface definition for the [NDArray] class.
type INDArray interface {
	objectivec.IObject
	ArrayViewWithShapeStrides(shape unsafe.Pointer, strides unsafe.Pointer) unsafe.Pointer
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArray
type NDArray struct {
	objectivec.Object
}

// NDArrayFrom constructs a [NDArray] from an unsafe.Pointer.
func NDArrayFrom(ptr unsafe.Pointer) NDArray {
	return NDArray{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (nc _NDArrayClass) Alloc() NDArray {
	rv := objc.Send[NDArray](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NDArrayClass) New() NDArray {
	rv := objc.Send[NDArray](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NDArray) Init() NDArray {
	rv := objc.Send[NDArray](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NDArray) Autorelease() NDArray {
	rv := objc.Send[NDArray](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNDArray creates a new NDArray instance.
func NewNDArray() NDArray {
	return getNDArrayClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArray/arrayView(withShape:strides:)
func (n_ NDArray) ArrayViewWithShapeStrides(shape unsafe.Pointer, strides unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("arrayViewWithShape:strides:"), shape, strides)
	return rv
}



