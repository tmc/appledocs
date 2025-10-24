// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [NDArrayMatrixMultiplication] class.
var (
	NDArrayMatrixMultiplicationClass     _NDArrayMatrixMultiplicationClass
	NDArrayMatrixMultiplicationClassOnce sync.Once
)

func getNDArrayMatrixMultiplicationClass() _NDArrayMatrixMultiplicationClass {
	NDArrayMatrixMultiplicationClassOnce.Do(func() {
		NDArrayMatrixMultiplicationClass = _NDArrayMatrixMultiplicationClass{objc.GetClass("MPSNDArrayMatrixMultiplication")}
	})
	return NDArrayMatrixMultiplicationClass
}

type _NDArrayMatrixMultiplicationClass struct {
	class objc.Class
}

// An interface definition for the [NDArrayMatrixMultiplication] class.
type INDArrayMatrixMultiplication interface {
	INDArrayMultiaryKernel
	// properties:
	Alpha() float64
	SetAlpha(value float64)
	Beta() float64
	SetBeta(value float64)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayMatrixMultiplication
type NDArrayMatrixMultiplication struct {
	NDArrayMultiaryKernel
}

// NDArrayMatrixMultiplicationFrom constructs a [NDArrayMatrixMultiplication] from an unsafe.Pointer.
func NDArrayMatrixMultiplicationFrom(ptr unsafe.Pointer) NDArrayMatrixMultiplication {
	return NDArrayMatrixMultiplication{
		NDArrayMultiaryKernel: NDArrayMultiaryKernelFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (nc _NDArrayMatrixMultiplicationClass) Alloc() NDArrayMatrixMultiplication {
	rv := objc.Send[NDArrayMatrixMultiplication](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NDArrayMatrixMultiplicationClass) New() NDArrayMatrixMultiplication {
	rv := objc.Send[NDArrayMatrixMultiplication](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NDArrayMatrixMultiplication) Init() NDArrayMatrixMultiplication {
	rv := objc.Send[NDArrayMatrixMultiplication](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NDArrayMatrixMultiplication) Autorelease() NDArrayMatrixMultiplication {
	rv := objc.Send[NDArrayMatrixMultiplication](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNDArrayMatrixMultiplication creates a new NDArrayMatrixMultiplication instance.
func NewNDArrayMatrixMultiplication() NDArrayMatrixMultiplication {
	return getNDArrayMatrixMultiplicationClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayMatrixMultiplication/alpha
func (n_ NDArrayMatrixMultiplication) Alpha() float64 {
	rv := objc.Send[float64](n_.ID, objc.Sel("alpha"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayMatrixMultiplication/alpha
func (n_ NDArrayMatrixMultiplication) SetAlpha(value float64) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setAlpha:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayMatrixMultiplication/beta
func (n_ NDArrayMatrixMultiplication) Beta() float64 {
	rv := objc.Send[float64](n_.ID, objc.Sel("beta"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayMatrixMultiplication/beta
func (n_ NDArrayMatrixMultiplication) SetBeta(value float64) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setBeta:"), value)
}



