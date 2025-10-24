// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSMatrixSolveTriangular */


/* debug [class_header]: Header for MPSMatrixSolveTriangular */
// The class instance for the [MatrixSolveTriangular] class.
var (
	MatrixSolveTriangularClass     _MatrixSolveTriangularClass
	MatrixSolveTriangularClassOnce sync.Once
)

func getMatrixSolveTriangularClass() _MatrixSolveTriangularClass {
	MatrixSolveTriangularClassOnce.Do(func() {
		MatrixSolveTriangularClass = _MatrixSolveTriangularClass{objc.GetClass("MPSMatrixSolveTriangular")}
	})
	return MatrixSolveTriangularClass
}

type _MatrixSolveTriangularClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MatrixSolveTriangular */
// An interface definition for the [MatrixSolveTriangular] class.
type IMatrixSolveTriangular interface {
	IMatrixBinaryKernel
	
/* debug [class_interface_properties]: Properties for MatrixSolveTriangular */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MatrixSolveTriangular */
	// methods:
	Encode()
	EncodeToCommandBufferSourceMatrixRightHandSideMatrixSolutionMatrix(commandBuffer unsafe.Pointer, sourceMatrix IMatrix, rightHandSideMatrix IMatrix, solutionMatrix IMatrix)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MatrixSolveTriangular */
// Alloc allocates a new instance without initialization.
func (mc _MatrixSolveTriangularClass) Alloc() MatrixSolveTriangular {
	rv := objc.Send[MatrixSolveTriangular](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MatrixSolveTriangularClass) New() MatrixSolveTriangular {
	rv := objc.Send[MatrixSolveTriangular](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MatrixSolveTriangular) Init() MatrixSolveTriangular {
	rv := objc.Send[MatrixSolveTriangular](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MatrixSolveTriangular) Autorelease() MatrixSolveTriangular {
	rv := objc.Send[MatrixSolveTriangular](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMatrixSolveTriangular creates a new MatrixSolveTriangular instance.
func NewMatrixSolveTriangular() MatrixSolveTriangular {
	return getMatrixSolveTriangularClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MatrixSolveTriangular */
// A kernel for computing the solution of a linear system of equations using a triangular coefficient matrix.
//
// This kernel finds the solution matrix to the system or , where: A is either an upper or lower triangular matrix is either or is the resulting matrix of solutions is the array of right hand sides for which the equations are to be solved


// A kernel for computing the solution of a linear system of equations using a triangular coefficient matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixSolveTriangular
type MatrixSolveTriangular struct {
	MatrixBinaryKernel
}

// MatrixSolveTriangularFrom constructs a [MatrixSolveTriangular] from an unsafe.Pointer.
//
// A kernel for computing the solution of a linear system of equations using a triangular coefficient matrix.
func MatrixSolveTriangularFrom(ptr unsafe.Pointer) MatrixSolveTriangular {
	return MatrixSolveTriangular{
		MatrixBinaryKernel: MatrixBinaryKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MatrixSolveTriangular */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsolvetriangular/2873007-initwithdevice
func NewMatrixSolveTriangularWithDeviceRightUpperTransposeUnitOrderNumberOfRightHandSidesAlpha(device unsafe.Pointer, right bool, upper bool, transpose bool, unit bool, order uint, numberOfRightHandSides uint, alpha float64) MatrixSolveTriangular {
	instance := getMatrixSolveTriangularClass().Alloc()
	rv := objc.Send[MatrixSolveTriangular](instance.ID, objc.Sel("initWithDevice:right:upper:transpose:unit:order:numberOfRightHandSides:alpha:"), device, right, upper, transpose, unit, order, numberOfRightHandSides, alpha)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMatrixSolveTriangularWithDeviceRightUpperTransposeUnitOrderNumberOfRightHandSidesAlpha */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MatrixSolveTriangular */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MatrixSolveTriangular */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MatrixSolveTriangular */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsolvetriangular/2867027-encode
func (m_ MatrixSolveTriangular) Encode() {
	objc.Send[objc.ID](m_.ID, objc.Sel("encode"))
}/* debug [instance_methods/method]: Encode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsolvetriangular/2867027-encodetocommandbuffer
func (m_ MatrixSolveTriangular) EncodeToCommandBufferSourceMatrixRightHandSideMatrixSolutionMatrix(commandBuffer unsafe.Pointer, sourceMatrix IMatrix, rightHandSideMatrix IMatrix, solutionMatrix IMatrix) {
	objc.Send[objc.ID](m_.ID, objc.Sel("encodeToCommandBuffer:sourceMatrix:rightHandSideMatrix:solutionMatrix:"), commandBuffer, sourceMatrix, rightHandSideMatrix, solutionMatrix)
}/* debug [instance_methods/method]: EncodeToCommandBufferSourceMatrixRightHandSideMatrixSolutionMatrix */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MatrixSolveTriangular */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSMatrixSolveTriangular */


