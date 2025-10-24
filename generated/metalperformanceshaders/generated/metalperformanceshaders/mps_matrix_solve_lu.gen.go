// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSMatrixSolveLU */


/* debug [class_header]: Header for MPSMatrixSolveLU */
// The class instance for the [MatrixSolveLU] class.
var (
	MatrixSolveLUClass     _MatrixSolveLUClass
	MatrixSolveLUClassOnce sync.Once
)

func getMatrixSolveLUClass() _MatrixSolveLUClass {
	MatrixSolveLUClassOnce.Do(func() {
		MatrixSolveLUClass = _MatrixSolveLUClass{objc.GetClass("MPSMatrixSolveLU")}
	})
	return MatrixSolveLUClass
}

type _MatrixSolveLUClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MatrixSolveLU */
// An interface definition for the [MatrixSolveLU] class.
type IMatrixSolveLU interface {
	IMatrixBinaryKernel
	
/* debug [class_interface_properties]: Properties for MatrixSolveLU */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MatrixSolveLU */
	// methods:
	Encode()
	EncodeToCommandBufferSourceMatrixRightHandSideMatrixPivotIndicesSolutionMatrix(commandBuffer unsafe.Pointer, sourceMatrix IMatrix, rightHandSideMatrix IMatrix, pivotIndices IMatrix, solutionMatrix IMatrix)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MatrixSolveLU */
// Alloc allocates a new instance without initialization.
func (mc _MatrixSolveLUClass) Alloc() MatrixSolveLU {
	rv := objc.Send[MatrixSolveLU](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MatrixSolveLUClass) New() MatrixSolveLU {
	rv := objc.Send[MatrixSolveLU](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MatrixSolveLU) Init() MatrixSolveLU {
	rv := objc.Send[MatrixSolveLU](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MatrixSolveLU) Autorelease() MatrixSolveLU {
	rv := objc.Send[MatrixSolveLU](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMatrixSolveLU creates a new MatrixSolveLU instance.
func NewMatrixSolveLU() MatrixSolveLU {
	return getMatrixSolveLUClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MatrixSolveLU */
// A kernel for computing the solution of a linear system of equations using an LU factorization.
//
// This kernel finds the solution matrix to the system , where: is or is the resulting matrix of solutions is the array of right hand sides for which the equations are to be solved


// A kernel for computing the solution of a linear system of equations using an LU factorization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixSolveLU
type MatrixSolveLU struct {
	MatrixBinaryKernel
}

// MatrixSolveLUFrom constructs a [MatrixSolveLU] from an unsafe.Pointer.
//
// A kernel for computing the solution of a linear system of equations using an LU factorization.
func MatrixSolveLUFrom(ptr unsafe.Pointer) MatrixSolveLU {
	return MatrixSolveLU{
		MatrixBinaryKernel: MatrixBinaryKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MatrixSolveLU */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsolvelu/2873005-initwithdevice
func NewMatrixSolveLUWithDeviceTransposeOrderNumberOfRightHandSides(device unsafe.Pointer, transpose bool, order uint, numberOfRightHandSides uint) MatrixSolveLU {
	instance := getMatrixSolveLUClass().Alloc()
	rv := objc.Send[MatrixSolveLU](instance.ID, objc.Sel("initWithDevice:transpose:order:numberOfRightHandSides:"), device, transpose, order, numberOfRightHandSides)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMatrixSolveLUWithDeviceTransposeOrderNumberOfRightHandSides */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MatrixSolveLU */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MatrixSolveLU */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MatrixSolveLU */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsolvelu/2867074-encode
func (m_ MatrixSolveLU) Encode() {
	objc.Send[objc.ID](m_.ID, objc.Sel("encode"))
}/* debug [instance_methods/method]: Encode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsolvelu/2867074-encodetocommandbuffer
func (m_ MatrixSolveLU) EncodeToCommandBufferSourceMatrixRightHandSideMatrixPivotIndicesSolutionMatrix(commandBuffer unsafe.Pointer, sourceMatrix IMatrix, rightHandSideMatrix IMatrix, pivotIndices IMatrix, solutionMatrix IMatrix) {
	objc.Send[objc.ID](m_.ID, objc.Sel("encodeToCommandBuffer:sourceMatrix:rightHandSideMatrix:pivotIndices:solutionMatrix:"), commandBuffer, sourceMatrix, rightHandSideMatrix, pivotIndices, solutionMatrix)
}/* debug [instance_methods/method]: EncodeToCommandBufferSourceMatrixRightHandSideMatrixPivotIndicesSolutionMatrix */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MatrixSolveLU */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSMatrixSolveLU */


