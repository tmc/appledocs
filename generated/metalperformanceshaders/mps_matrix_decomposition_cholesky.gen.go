// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSMatrixDecompositionCholesky */


/* debug [class_header]: Header for MPSMatrixDecompositionCholesky */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MatrixDecompositionCholesky */
// An interface definition for the [MatrixDecompositionCholesky] class.
type IMatrixDecompositionCholesky interface {
	IMatrixUnaryKernel
	
/* debug [class_interface_properties]: Properties for MatrixDecompositionCholesky */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MatrixDecompositionCholesky */
	// methods:
	Encode()
	EncodeToCommandBufferSourceMatrixResultMatrixStatus(commandBuffer unsafe.Pointer, sourceMatrix IMatrix, resultMatrix IMatrix, status unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MatrixDecompositionCholesky */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MatrixDecompositionCholesky */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MatrixDecompositionCholesky */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixdecompositioncholesky/2867119-initwithdevice
func NewMatrixDecompositionCholeskyWithDeviceLowerOrder(device unsafe.Pointer, lower bool, order uint) MatrixDecompositionCholesky {
	instance := getMatrixDecompositionCholeskyClass().Alloc()
	rv := objc.Send[MatrixDecompositionCholesky](instance.ID, objc.Sel("initWithDevice:lower:order:"), device, lower, order)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMatrixDecompositionCholeskyWithDeviceLowerOrder */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MatrixDecompositionCholesky */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MatrixDecompositionCholesky */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MatrixDecompositionCholesky */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixdecompositioncholesky/2867004-encode
func (m_ MatrixDecompositionCholesky) Encode() {
	objc.Send[objc.ID](m_.ID, objc.Sel("encode"))
}/* debug [instance_methods/method]: Encode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixdecompositioncholesky/2867004-encodetocommandbuffer
func (m_ MatrixDecompositionCholesky) EncodeToCommandBufferSourceMatrixResultMatrixStatus(commandBuffer unsafe.Pointer, sourceMatrix IMatrix, resultMatrix IMatrix, status unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("encodeToCommandBuffer:sourceMatrix:resultMatrix:status:"), commandBuffer, sourceMatrix, resultMatrix, status)
}/* debug [instance_methods/method]: EncodeToCommandBufferSourceMatrixResultMatrixStatus */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MatrixDecompositionCholesky */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSMatrixDecompositionCholesky */


