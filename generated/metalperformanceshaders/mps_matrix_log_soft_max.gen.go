// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [MatrixLogSoftMax] class.
var (
	MatrixLogSoftMaxClass     _MatrixLogSoftMaxClass
	MatrixLogSoftMaxClassOnce sync.Once
)

func getMatrixLogSoftMaxClass() _MatrixLogSoftMaxClass {
	MatrixLogSoftMaxClassOnce.Do(func() {
		MatrixLogSoftMaxClass = _MatrixLogSoftMaxClass{objc.GetClass("MPSMatrixLogSoftMax")}
	})
	return MatrixLogSoftMaxClass
}

type _MatrixLogSoftMaxClass struct {
	class objc.Class
}





// An interface definition for the [MatrixLogSoftMax] class.
type IMatrixLogSoftMax interface {
	IMatrixSoftMax
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (mc _MatrixLogSoftMaxClass) Alloc() MatrixLogSoftMax {
	rv := objc.Send[MatrixLogSoftMax](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MatrixLogSoftMaxClass) New() MatrixLogSoftMax {
	rv := objc.Send[MatrixLogSoftMax](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MatrixLogSoftMax) Init() MatrixLogSoftMax {
	rv := objc.Send[MatrixLogSoftMax](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MatrixLogSoftMax) Autorelease() MatrixLogSoftMax {
	rv := objc.Send[MatrixLogSoftMax](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMatrixLogSoftMax creates a new MatrixLogSoftMax instance.
func NewMatrixLogSoftMax() MatrixLogSoftMax {
	return getMatrixLogSoftMaxClass().New()
}





// A logarithmic softmax kernel that operates on matrices.


// A logarithmic softmax kernel that operates on matrices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixLogSoftMax
type MatrixLogSoftMax struct {
	MatrixSoftMax
}

// MatrixLogSoftMaxFrom constructs a [MatrixLogSoftMax] from an unsafe.Pointer.
//
// A logarithmic softmax kernel that operates on matrices.
func MatrixLogSoftMaxFrom(ptr unsafe.Pointer) MatrixLogSoftMax {
	return MatrixLogSoftMax{
		MatrixSoftMax: MatrixSoftMaxFrom(ptr),
	}
}































