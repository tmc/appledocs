// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSMatrixDecompositionLU */


/* debug [class_header]: Header for MPSMatrixDecompositionLU */
// The class instance for the [MatrixDecompositionLU] class.
var (
	MatrixDecompositionLUClass     _MatrixDecompositionLUClass
	MatrixDecompositionLUClassOnce sync.Once
)

func getMatrixDecompositionLUClass() _MatrixDecompositionLUClass {
	MatrixDecompositionLUClassOnce.Do(func() {
		MatrixDecompositionLUClass = _MatrixDecompositionLUClass{objc.GetClass("MPSMatrixDecompositionLU")}
	})
	return MatrixDecompositionLUClass
}

type _MatrixDecompositionLUClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MatrixDecompositionLU */
// An interface definition for the [MatrixDecompositionLU] class.
type IMatrixDecompositionLU interface {
	IMatrixUnaryKernel
	
/* debug [class_interface_properties]: Properties for MatrixDecompositionLU */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MatrixDecompositionLU */
	// methods:
	Encode()
	EncodeToCommandBufferSourceMatrixResultMatrixPivotIndicesStatus(commandBuffer unsafe.Pointer, sourceMatrix IMatrix, resultMatrix IMatrix, pivotIndices IMatrix, status unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MatrixDecompositionLU */
// Alloc allocates a new instance without initialization.
func (mc _MatrixDecompositionLUClass) Alloc() MatrixDecompositionLU {
	rv := objc.Send[MatrixDecompositionLU](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MatrixDecompositionLUClass) New() MatrixDecompositionLU {
	rv := objc.Send[MatrixDecompositionLU](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MatrixDecompositionLU) Init() MatrixDecompositionLU {
	rv := objc.Send[MatrixDecompositionLU](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MatrixDecompositionLU) Autorelease() MatrixDecompositionLU {
	rv := objc.Send[MatrixDecompositionLU](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMatrixDecompositionLU creates a new MatrixDecompositionLU instance.
func NewMatrixDecompositionLU() MatrixDecompositionLU {
	return getMatrixDecompositionLUClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MatrixDecompositionLU */
// A kernel for computing the LU factorization of a matrix using partial pivoting with row interchanges.
//
// This kernel object computes an LU factorization, , where: is a matrix for which the LU factorization is to be computed is a unit lower triangular matrix is an upper triangular matrix is a permutation matrix


// A kernel for computing the LU factorization of a matrix using partial pivoting with row interchanges.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixDecompositionLU
type MatrixDecompositionLU struct {
	MatrixUnaryKernel
}

// MatrixDecompositionLUFrom constructs a [MatrixDecompositionLU] from an unsafe.Pointer.
//
// A kernel for computing the LU factorization of a matrix using partial pivoting with row interchanges.
func MatrixDecompositionLUFrom(ptr unsafe.Pointer) MatrixDecompositionLU {
	return MatrixDecompositionLU{
		MatrixUnaryKernel: MatrixUnaryKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MatrixDecompositionLU */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixdecompositionlu/2866960-initwithdevice
func NewMatrixDecompositionLUWithDeviceRowsColumns(device unsafe.Pointer, rows uint, columns uint) MatrixDecompositionLU {
	instance := getMatrixDecompositionLUClass().Alloc()
	rv := objc.Send[MatrixDecompositionLU](instance.ID, objc.Sel("initWithDevice:rows:columns:"), device, rows, columns)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMatrixDecompositionLUWithDeviceRowsColumns */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MatrixDecompositionLU */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MatrixDecompositionLU */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MatrixDecompositionLU */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixdecompositionlu/2867184-encode
func (m_ MatrixDecompositionLU) Encode() {
	objc.Send[objc.ID](m_.ID, objc.Sel("encode"))
}/* debug [instance_methods/method]: Encode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixdecompositionlu/2867184-encodetocommandbuffer
func (m_ MatrixDecompositionLU) EncodeToCommandBufferSourceMatrixResultMatrixPivotIndicesStatus(commandBuffer unsafe.Pointer, sourceMatrix IMatrix, resultMatrix IMatrix, pivotIndices IMatrix, status unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("encodeToCommandBuffer:sourceMatrix:resultMatrix:pivotIndices:status:"), commandBuffer, sourceMatrix, resultMatrix, pivotIndices, status)
}/* debug [instance_methods/method]: EncodeToCommandBufferSourceMatrixResultMatrixPivotIndicesStatus */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MatrixDecompositionLU */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSMatrixDecompositionLU */


