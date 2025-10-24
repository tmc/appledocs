// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSMatrixVectorMultiplication */


/* debug [class_header]: Header for MPSMatrixVectorMultiplication */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MatrixVectorMultiplication */
// An interface definition for the [MatrixVectorMultiplication] class.
type IMatrixVectorMultiplication interface {
	IMatrixBinaryKernel
	
/* debug [class_interface_properties]: Properties for MatrixVectorMultiplication */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MatrixVectorMultiplication */
	// methods:
	Encode()
	EncodeToCommandBufferInputMatrixInputVectorResultVector(commandBuffer unsafe.Pointer, inputMatrix IMatrix, inputVector IVector, resultVector IVector)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MatrixVectorMultiplication */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MatrixVectorMultiplication */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MatrixVectorMultiplication */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixvectormultiplication/2909035-initwithdevice
func NewMatrixVectorMultiplicationWithDeviceRowsColumns(device unsafe.Pointer, rows uint, columns uint) MatrixVectorMultiplication {
	instance := getMatrixVectorMultiplicationClass().Alloc()
	rv := objc.Send[MatrixVectorMultiplication](instance.ID, objc.Sel("initWithDevice:rows:columns:"), device, rows, columns)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMatrixVectorMultiplicationWithDeviceRowsColumns */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixvectormultiplication/2873083-initwithdevice
func NewMatrixVectorMultiplicationWithDeviceTransposeRowsColumnsAlphaBeta(device unsafe.Pointer, transpose bool, rows uint, columns uint, alpha float64, beta float64) MatrixVectorMultiplication {
	instance := getMatrixVectorMultiplicationClass().Alloc()
	rv := objc.Send[MatrixVectorMultiplication](instance.ID, objc.Sel("initWithDevice:transpose:rows:columns:alpha:beta:"), device, transpose, rows, columns, alpha, beta)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMatrixVectorMultiplicationWithDeviceTransposeRowsColumnsAlphaBeta */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MatrixVectorMultiplication */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MatrixVectorMultiplication */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MatrixVectorMultiplication */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixvectormultiplication/2873084-encode
func (m_ MatrixVectorMultiplication) Encode() {
	objc.Send[objc.ID](m_.ID, objc.Sel("encode"))
}/* debug [instance_methods/method]: Encode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixvectormultiplication/2873084-encodetocommandbuffer
func (m_ MatrixVectorMultiplication) EncodeToCommandBufferInputMatrixInputVectorResultVector(commandBuffer unsafe.Pointer, inputMatrix IMatrix, inputVector IVector, resultVector IVector) {
	objc.Send[objc.ID](m_.ID, objc.Sel("encodeToCommandBuffer:inputMatrix:inputVector:resultVector:"), commandBuffer, inputMatrix, inputVector, resultVector)
}/* debug [instance_methods/method]: EncodeToCommandBufferInputMatrixInputVectorResultVector */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MatrixVectorMultiplication */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSMatrixVectorMultiplication */


