// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [MatrixSolveCholesky] class.
var (
	MatrixSolveCholeskyClass     _MatrixSolveCholeskyClass
	MatrixSolveCholeskyClassOnce sync.Once
)

func getMatrixSolveCholeskyClass() _MatrixSolveCholeskyClass {
	MatrixSolveCholeskyClassOnce.Do(func() {
		MatrixSolveCholeskyClass = _MatrixSolveCholeskyClass{objc.GetClass("MPSMatrixSolveCholesky")}
	})
	return MatrixSolveCholeskyClass
}

type _MatrixSolveCholeskyClass struct {
	class objc.Class
}





// An interface definition for the [MatrixSolveCholesky] class.
type IMatrixSolveCholesky interface {
	IMatrixBinaryKernel
	

	// properties:


	

	// methods:
	Encode()
	EncodeToCommandBufferSourceMatrixRightHandSideMatrixSolutionMatrix(commandBuffer unsafe.Pointer, sourceMatrix IMatrix, rightHandSideMatrix IMatrix, solutionMatrix IMatrix)


}





// Alloc allocates a new instance without initialization.
func (mc _MatrixSolveCholeskyClass) Alloc() MatrixSolveCholesky {
	rv := objc.Send[MatrixSolveCholesky](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MatrixSolveCholeskyClass) New() MatrixSolveCholesky {
	rv := objc.Send[MatrixSolveCholesky](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MatrixSolveCholesky) Init() MatrixSolveCholesky {
	rv := objc.Send[MatrixSolveCholesky](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MatrixSolveCholesky) Autorelease() MatrixSolveCholesky {
	rv := objc.Send[MatrixSolveCholesky](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMatrixSolveCholesky creates a new MatrixSolveCholesky instance.
func NewMatrixSolveCholesky() MatrixSolveCholesky {
	return getMatrixSolveCholeskyClass().New()
}





// A kernel for computing the solution of a linear system of equations using a Cholesky factorization.
//
// This kernel finds the solution matrix to the system , where: is a symmetric positive-definite matrix is the resulting matrix of solutions is the array of right-hand-sides for which the equations are to be solved


// A kernel for computing the solution of a linear system of equations using a Cholesky factorization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixSolveCholesky
type MatrixSolveCholesky struct {
	MatrixBinaryKernel
}

// MatrixSolveCholeskyFrom constructs a [MatrixSolveCholesky] from an unsafe.Pointer.
//
// A kernel for computing the solution of a linear system of equations using a Cholesky factorization.
func MatrixSolveCholeskyFrom(ptr unsafe.Pointer) MatrixSolveCholesky {
	return MatrixSolveCholesky{
		MatrixBinaryKernel: MatrixBinaryKernelFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsolvecholesky/2873006-initwithdevice
func NewMatrixSolveCholeskyWithDeviceUpperOrderNumberOfRightHandSides(device unsafe.Pointer, upper bool, order uint, numberOfRightHandSides uint) MatrixSolveCholesky {
	instance := getMatrixSolveCholeskyClass().Alloc()
	rv := objc.Send[MatrixSolveCholesky](instance.ID, objc.Sel("initWithDevice:upper:order:numberOfRightHandSides:"), device, upper, order, numberOfRightHandSides)
	rv.Autorelease()
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsolvecholesky/2866957-encode
func (m_ MatrixSolveCholesky) Encode() {
	objc.Send[objc.ID](m_.ID, objc.Sel("encode"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsolvecholesky/2866957-encodetocommandbuffer
func (m_ MatrixSolveCholesky) EncodeToCommandBufferSourceMatrixRightHandSideMatrixSolutionMatrix(commandBuffer unsafe.Pointer, sourceMatrix IMatrix, rightHandSideMatrix IMatrix, solutionMatrix IMatrix) {
	objc.Send[objc.ID](m_.ID, objc.Sel("encodeToCommandBuffer:sourceMatrix:rightHandSideMatrix:solutionMatrix:"), commandBuffer, sourceMatrix, rightHandSideMatrix, solutionMatrix)
}












