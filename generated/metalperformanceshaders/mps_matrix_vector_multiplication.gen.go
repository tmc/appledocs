// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [MatrixVectorMultiplication] class.
var (
	MatrixVectorMultiplicationClass     _MatrixVectorMultiplicationClass
	MatrixVectorMultiplicationClassOnce sync.Once
)

func getMatrixVectorMultiplicationClass() _MatrixVectorMultiplicationClass {
	MatrixVectorMultiplicationClassOnce.Do(func() {
		MatrixVectorMultiplicationClass = _MatrixVectorMultiplicationClass{objc.GetClass("MPSMatrixVectorMultiplication")}
	})
	return MatrixVectorMultiplicationClass
}

type _MatrixVectorMultiplicationClass struct {
	class objc.Class
}





// An interface definition for the [MatrixVectorMultiplication] class.
type IMatrixVectorMultiplication interface {
	IMatrixBinaryKernel
	

	// properties:


	

	// methods:
	Encode()
	EncodeToCommandBufferInputMatrixInputVectorResultVector(commandBuffer unsafe.Pointer, inputMatrix IMatrix, inputVector IVector, resultVector IVector)


}





// Alloc allocates a new instance without initialization.
func (mc _MatrixVectorMultiplicationClass) Alloc() MatrixVectorMultiplication {
	rv := objc.Send[MatrixVectorMultiplication](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MatrixVectorMultiplicationClass) New() MatrixVectorMultiplication {
	rv := objc.Send[MatrixVectorMultiplication](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MatrixVectorMultiplication) Init() MatrixVectorMultiplication {
	rv := objc.Send[MatrixVectorMultiplication](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MatrixVectorMultiplication) Autorelease() MatrixVectorMultiplication {
	rv := objc.Send[MatrixVectorMultiplication](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMatrixVectorMultiplication creates a new MatrixVectorMultiplication instance.
func NewMatrixVectorMultiplication() MatrixVectorMultiplication {
	return getMatrixVectorMultiplicationClass().New()
}





// A matrix-vector multiplication kernel


// A matrix-vector multiplication kernel
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixVectorMultiplication
type MatrixVectorMultiplication struct {
	MatrixBinaryKernel
}

// MatrixVectorMultiplicationFrom constructs a [MatrixVectorMultiplication] from an unsafe.Pointer.
//
// A matrix-vector multiplication kernel
func MatrixVectorMultiplicationFrom(ptr unsafe.Pointer) MatrixVectorMultiplication {
	return MatrixVectorMultiplication{
		MatrixBinaryKernel: MatrixBinaryKernelFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixvectormultiplication/2909035-initwithdevice
func NewMatrixVectorMultiplicationWithDeviceRowsColumns(device unsafe.Pointer, rows uint, columns uint) MatrixVectorMultiplication {
	instance := getMatrixVectorMultiplicationClass().Alloc()
	rv := objc.Send[MatrixVectorMultiplication](instance.ID, objc.Sel("initWithDevice:rows:columns:"), device, rows, columns)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixvectormultiplication/2873083-initwithdevice
func NewMatrixVectorMultiplicationWithDeviceTransposeRowsColumnsAlphaBeta(device unsafe.Pointer, transpose bool, rows uint, columns uint, alpha float64, beta float64) MatrixVectorMultiplication {
	instance := getMatrixVectorMultiplicationClass().Alloc()
	rv := objc.Send[MatrixVectorMultiplication](instance.ID, objc.Sel("initWithDevice:transpose:rows:columns:alpha:beta:"), device, transpose, rows, columns, alpha, beta)
	rv.Autorelease()
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixvectormultiplication/2873084-encode
func (m_ MatrixVectorMultiplication) Encode() {
	objc.Send[objc.ID](m_.ID, objc.Sel("encode"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixvectormultiplication/2873084-encodetocommandbuffer
func (m_ MatrixVectorMultiplication) EncodeToCommandBufferInputMatrixInputVectorResultVector(commandBuffer unsafe.Pointer, inputMatrix IMatrix, inputVector IVector, resultVector IVector) {
	objc.Send[objc.ID](m_.ID, objc.Sel("encodeToCommandBuffer:inputMatrix:inputVector:resultVector:"), commandBuffer, inputMatrix, inputVector, resultVector)
}












