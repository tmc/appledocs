// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSMatrixMultiplication */


/* debug [class_header]: Header for MPSMatrixMultiplication */
// The class instance for the [MatrixMultiplication] class.
var (
	MatrixMultiplicationClass     _MatrixMultiplicationClass
	MatrixMultiplicationClassOnce sync.Once
)

func getMatrixMultiplicationClass() _MatrixMultiplicationClass {
	MatrixMultiplicationClassOnce.Do(func() {
		MatrixMultiplicationClass = _MatrixMultiplicationClass{objc.GetClass("MPSMatrixMultiplication")}
	})
	return MatrixMultiplicationClass
}

type _MatrixMultiplicationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MatrixMultiplication */
// An interface definition for the [MatrixMultiplication] class.
type IMatrixMultiplication interface {
	IKernel
	
/* debug [class_interface_properties]: Properties for MatrixMultiplication */
	// properties:
	LeftMatrixOrigin() Origin get set /* not a class type */
	SetLeftMatrixOrigin(value Origin get set /* not a class type */)
	ResultMatrixOrigin() Origin get set /* not a class type */
	SetResultMatrixOrigin(value Origin get set /* not a class type */)
	RightMatrixOrigin() Origin get set /* not a class type */
	SetRightMatrixOrigin(value Origin get set /* not a class type */)
	BatchStart() objectivec.IObject
	SetBatchStart(value objectivec.IObject)
	BatchSize() objectivec.IObject
	SetBatchSize(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MatrixMultiplication */
	// methods:
	Encode()
	EncodeToCommandBufferLeftMatrixRightMatrixResultMatrix(commandBuffer unsafe.Pointer, leftMatrix IMatrix, rightMatrix IMatrix, resultMatrix IMatrix)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MatrixMultiplication */
// Alloc allocates a new instance without initialization.
func (mc _MatrixMultiplicationClass) Alloc() MatrixMultiplication {
	rv := objc.Send[MatrixMultiplication](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MatrixMultiplicationClass) New() MatrixMultiplication {
	rv := objc.Send[MatrixMultiplication](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MatrixMultiplication) Init() MatrixMultiplication {
	rv := objc.Send[MatrixMultiplication](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MatrixMultiplication) Autorelease() MatrixMultiplication {
	rv := objc.Send[MatrixMultiplication](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMatrixMultiplication creates a new MatrixMultiplication instance.
func NewMatrixMultiplication() MatrixMultiplication {
	return getMatrixMultiplicationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MatrixMultiplication */
// A matrix multiplication kernel.
//
// An object computes the following operation: Where , _,_ and are matrices represented by objects, and and are scalar values of the same data type as the values of . and may each have an optional transposition operation applied. Matrices , , and are also referred to as the left input matrix, the right input matrix, and the result matrix respectively.


// A matrix multiplication kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixMultiplication
type MatrixMultiplication struct {
	Kernel
}

// MatrixMultiplicationFrom constructs a [MatrixMultiplication] from an unsafe.Pointer.
//
// A matrix multiplication kernel.
func MatrixMultiplicationFrom(ptr unsafe.Pointer) MatrixMultiplication {
	return MatrixMultiplication{
		Kernel: KernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MatrixMultiplication */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixmultiplication/2909034-initwithdevice
func NewMatrixMultiplicationWithDeviceResultRowsResultColumnsInteriorColumns(device unsafe.Pointer, resultRows uint, resultColumns uint, interiorColumns uint) MatrixMultiplication {
	instance := getMatrixMultiplicationClass().Alloc()
	rv := objc.Send[MatrixMultiplication](instance.ID, objc.Sel("initWithDevice:resultRows:resultColumns:interiorColumns:"), device, resultRows, resultColumns, interiorColumns)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMatrixMultiplicationWithDeviceResultRowsResultColumnsInteriorColumns */


// Initializes a matrix multiplication kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixmultiplication/2147845-initwithdevice
func NewMatrixMultiplicationWithDeviceTransposeLeftTransposeRightResultRowsResultColumnsInteriorColumnsAlphaBeta(device unsafe.Pointer, transposeLeft bool, transposeRight bool, resultRows uint, resultColumns uint, interiorColumns uint, alpha float64, beta float64) MatrixMultiplication {
	instance := getMatrixMultiplicationClass().Alloc()
	rv := objc.Send[MatrixMultiplication](instance.ID, objc.Sel("initWithDevice:transposeLeft:transposeRight:resultRows:resultColumns:interiorColumns:alpha:beta:"), device, transposeLeft, transposeRight, resultRows, resultColumns, interiorColumns, alpha, beta)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMatrixMultiplicationWithDeviceTransposeLeftTransposeRightResultRowsResultColumnsInteriorColumnsAlphaBeta */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MatrixMultiplication */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MatrixMultiplication */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MatrixMultiplication */

// Encodes a matrix multiplication kernel to a command buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixmultiplication/2147848-encode
func (m_ MatrixMultiplication) Encode() {
	objc.Send[objc.ID](m_.ID, objc.Sel("encode"))
}/* debug [instance_methods/method]: Encode */


// Encodes a matrix multiplication kernel to a command buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixmultiplication/2147848-encodetocommandbuffer
func (m_ MatrixMultiplication) EncodeToCommandBufferLeftMatrixRightMatrixResultMatrix(commandBuffer unsafe.Pointer, leftMatrix IMatrix, rightMatrix IMatrix, resultMatrix IMatrix) {
	objc.Send[objc.ID](m_.ID, objc.Sel("encodeToCommandBuffer:leftMatrix:rightMatrix:resultMatrix:"), commandBuffer, leftMatrix, rightMatrix, resultMatrix)
}/* debug [instance_methods/method]: EncodeToCommandBufferLeftMatrixRightMatrixResultMatrix */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MatrixMultiplication */

// The origin of the left input matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixmultiplication/2147846-leftmatrixorigin
func (m_ MatrixMultiplication) LeftMatrixOrigin() Origin get set /* not a class type */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("leftMatrixOrigin"))
	return rv
}/* debug [instance_properties/getter]: leftMatrixOrigin */


// The origin of the left input matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixmultiplication/2147846-leftmatrixorigin
func (m_ MatrixMultiplication) SetLeftMatrixOrigin(value Origin get set /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLeftMatrixOrigin:"), value)
}/* debug [instance_properties/setter]: leftMatrixOrigin */


// The origin of the result matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixmultiplication/2147847-resultmatrixorigin
func (m_ MatrixMultiplication) ResultMatrixOrigin() Origin get set /* not a class type */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("resultMatrixOrigin"))
	return rv
}/* debug [instance_properties/getter]: resultMatrixOrigin */


// The origin of the result matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixmultiplication/2147847-resultmatrixorigin
func (m_ MatrixMultiplication) SetResultMatrixOrigin(value Origin get set /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setResultMatrixOrigin:"), value)
}/* debug [instance_properties/setter]: resultMatrixOrigin */


// The origin of the right input matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixmultiplication/2147851-rightmatrixorigin
func (m_ MatrixMultiplication) RightMatrixOrigin() Origin get set /* not a class type */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("rightMatrixOrigin"))
	return rv
}/* debug [instance_properties/getter]: rightMatrixOrigin */


// The origin of the right input matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixmultiplication/2147851-rightmatrixorigin
func (m_ MatrixMultiplication) SetRightMatrixOrigin(value Origin get set /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRightMatrixOrigin:"), value)
}/* debug [instance_properties/setter]: rightMatrixOrigin */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixmultiplication/2873081-batchstart
func (m_ MatrixMultiplication) BatchStart() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("batchStart"))
	return rv
}/* debug [instance_properties/getter]: batchStart */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixmultiplication/2873081-batchstart
func (m_ MatrixMultiplication) SetBatchStart(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBatchStart:"), value)
}/* debug [instance_properties/setter]: batchStart */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixmultiplication/2873082-batchsize
func (m_ MatrixMultiplication) BatchSize() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("batchSize"))
	return rv
}/* debug [instance_properties/getter]: batchSize */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixmultiplication/2873082-batchsize
func (m_ MatrixMultiplication) SetBatchSize(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBatchSize:"), value)
}/* debug [instance_properties/setter]: batchSize */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSMatrixMultiplication */


