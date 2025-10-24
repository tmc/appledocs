// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [MatrixDecompositionCholesky] class.
var (
	MatrixDecompositionCholeskyClass     _MatrixDecompositionCholeskyClass
	MatrixDecompositionCholeskyClassOnce sync.Once
)

func getMatrixDecompositionCholeskyClass() _MatrixDecompositionCholeskyClass {
	MatrixDecompositionCholeskyClassOnce.Do(func() {
		MatrixDecompositionCholeskyClass = _MatrixDecompositionCholeskyClass{objc.GetClass("MPSMatrixDecompositionCholesky")}
	})
	return MatrixDecompositionCholeskyClass
}

type _MatrixDecompositionCholeskyClass struct {
	class objc.Class
}





// An interface definition for the [MatrixDecompositionCholesky] class.
type IMatrixDecompositionCholesky interface {
	IMatrixUnaryKernel
	

	// properties:


	

	// methods:
	Encode()
	EncodeToCommandBufferSourceMatrixResultMatrixStatus(commandBuffer unsafe.Pointer, sourceMatrix IMatrix, resultMatrix IMatrix, status unsafe.Pointer)


}





// Alloc allocates a new instance without initialization.
func (mc _MatrixDecompositionCholeskyClass) Alloc() MatrixDecompositionCholesky {
	rv := objc.Send[MatrixDecompositionCholesky](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MatrixDecompositionCholeskyClass) New() MatrixDecompositionCholesky {
	rv := objc.Send[MatrixDecompositionCholesky](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MatrixDecompositionCholesky) Init() MatrixDecompositionCholesky {
	rv := objc.Send[MatrixDecompositionCholesky](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MatrixDecompositionCholesky) Autorelease() MatrixDecompositionCholesky {
	rv := objc.Send[MatrixDecompositionCholesky](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMatrixDecompositionCholesky creates a new MatrixDecompositionCholesky instance.
func NewMatrixDecompositionCholesky() MatrixDecompositionCholesky {
	return getMatrixDecompositionCholeskyClass().New()
}





// A kernel for computing the Cholesky factorization of a matrix.
//
// This kernel computes one of the following factorizations of a matrix : where: is a symmetric positive-definite matrix for which the factorization is to be computed is the lower triangular matrix is the upper triangular matrix


// A kernel for computing the Cholesky factorization of a matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixDecompositionCholesky
type MatrixDecompositionCholesky struct {
	MatrixUnaryKernel
}

// MatrixDecompositionCholeskyFrom constructs a [MatrixDecompositionCholesky] from an unsafe.Pointer.
//
// A kernel for computing the Cholesky factorization of a matrix.
func MatrixDecompositionCholeskyFrom(ptr unsafe.Pointer) MatrixDecompositionCholesky {
	return MatrixDecompositionCholesky{
		MatrixUnaryKernel: MatrixUnaryKernelFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixdecompositioncholesky/2867119-initwithdevice
func NewMatrixDecompositionCholeskyWithDeviceLowerOrder(device unsafe.Pointer, lower bool, order uint) MatrixDecompositionCholesky {
	instance := getMatrixDecompositionCholeskyClass().Alloc()
	rv := objc.Send[MatrixDecompositionCholesky](instance.ID, objc.Sel("initWithDevice:lower:order:"), device, lower, order)
	rv.Autorelease()
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixdecompositioncholesky/2867004-encode
func (m_ MatrixDecompositionCholesky) Encode() {
	objc.Send[objc.ID](m_.ID, objc.Sel("encode"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixdecompositioncholesky/2867004-encodetocommandbuffer
func (m_ MatrixDecompositionCholesky) EncodeToCommandBufferSourceMatrixResultMatrixStatus(commandBuffer unsafe.Pointer, sourceMatrix IMatrix, resultMatrix IMatrix, status unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("encodeToCommandBuffer:sourceMatrix:resultMatrix:status:"), commandBuffer, sourceMatrix, resultMatrix, status)
}












