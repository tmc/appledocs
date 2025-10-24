// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [MatrixSum] class.
var (
	MatrixSumClass     _MatrixSumClass
	MatrixSumClassOnce sync.Once
)

func getMatrixSumClass() _MatrixSumClass {
	MatrixSumClassOnce.Do(func() {
		MatrixSumClass = _MatrixSumClass{objc.GetClass("MPSMatrixSum")}
	})
	return MatrixSumClass
}

type _MatrixSumClass struct {
	class objc.Class
}





// An interface definition for the [MatrixSum] class.
type IMatrixSum interface {
	IKernel
	

	// properties:
	Columns() objectivec.IObject
	SetColumns(value objectivec.IObject)
	NeuronParameterB() objectivec.IObject
	SetNeuronParameterB(value objectivec.IObject)
	NeuronParameterC() objectivec.IObject
	SetNeuronParameterC(value objectivec.IObject)
	Count() objectivec.IObject
	SetCount(value objectivec.IObject)
	Transpose() objectivec.IObject
	SetTranspose(value objectivec.IObject)
	Rows() objectivec.IObject
	SetRows(value objectivec.IObject)
	NeuronParameterA() objectivec.IObject
	SetNeuronParameterA(value objectivec.IObject)
	ResultMatrixOrigin() Origin get set /* not a class type */
	SetResultMatrixOrigin(value Origin get set /* not a class type */)


	

	// methods:
	Encode()
	EncodeToCommandBufferSourceMatricesResultMatrixScaleVectorOffsetVectorBiasVectorStartIndex(buffer unsafe.Pointer, sourceMatrices unsafe.Pointer, resultMatrix IMatrix, scaleVector IVector, offsetVector IVector, biasVector IVector, startIndex uint)
	SetNeuronType()
	NeuronType()
	SetNeuronTypeParameterAParameterBParameterC(neuronType CNNNeuronType, parameterA float32, parameterB float32, parameterC float32)


}





// Alloc allocates a new instance without initialization.
func (mc _MatrixSumClass) Alloc() MatrixSum {
	rv := objc.Send[MatrixSum](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MatrixSumClass) New() MatrixSum {
	rv := objc.Send[MatrixSum](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MatrixSum) Init() MatrixSum {
	rv := objc.Send[MatrixSum](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MatrixSum) Autorelease() MatrixSum {
	rv := objc.Send[MatrixSum](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMatrixSum creates a new MatrixSum instance.
func NewMatrixSum() MatrixSum {
	return getMatrixSumClass().New()
}





// A kernel for performing a pointwise summation of a matrix.


// A kernel for performing a pointwise summation of a matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixSum
type MatrixSum struct {
	Kernel
}

// MatrixSumFrom constructs a [MatrixSum] from an unsafe.Pointer.
//
// A kernel for performing a pointwise summation of a matrix.
func MatrixSumFrom(ptr unsafe.Pointer) MatrixSum {
	return MatrixSum{
		Kernel: KernelFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsum/2935614-initwithcoder
func NewMatrixSumWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) MatrixSum {
	instance := getMatrixSumClass().Alloc()
	rv := objc.Send[MatrixSum](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsum/2935623-initwithdevice
func NewMatrixSumWithDeviceCountRowsColumnsTranspose(device unsafe.Pointer, count uint, rows uint, columns uint, transpose bool) MatrixSum {
	instance := getMatrixSumClass().Alloc()
	rv := objc.Send[MatrixSum](instance.ID, objc.Sel("initWithDevice:count:rows:columns:transpose:"), device, count, rows, columns, transpose)
	rv.Autorelease()
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsum/2935613-encode
func (m_ MatrixSum) Encode() {
	objc.Send[objc.ID](m_.ID, objc.Sel("encode"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsum/2935613-encodetocommandbuffer
func (m_ MatrixSum) EncodeToCommandBufferSourceMatricesResultMatrixScaleVectorOffsetVectorBiasVectorStartIndex(buffer unsafe.Pointer, sourceMatrices unsafe.Pointer, resultMatrix IMatrix, scaleVector IVector, offsetVector IVector, biasVector IVector, startIndex uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("encodeToCommandBuffer:sourceMatrices:resultMatrix:scaleVector:offsetVector:biasVector:startIndex:"), buffer, sourceMatrices, resultMatrix, scaleVector, offsetVector, biasVector, startIndex)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsum/2935617-setneurontype
func (m_ MatrixSum) SetNeuronType() {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNeuronType"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsum/2935625-neurontype
func (m_ MatrixSum) NeuronType() {
	objc.Send[objc.ID](m_.ID, objc.Sel("neuronType"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixSum/setNeuronType(_:parameterA:parameterB:parameterC:)
func (m_ MatrixSum) SetNeuronTypeParameterAParameterBParameterC(neuronType CNNNeuronType, parameterA float32, parameterB float32, parameterC float32) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNeuronType:parameterA:parameterB:parameterC:"), neuronType, parameterA, parameterB, parameterC)
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsum/2935615-columns
func (m_ MatrixSum) Columns() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("columns"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsum/2935615-columns
func (m_ MatrixSum) SetColumns(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setColumns:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsum/2935616-neuronparameterb
func (m_ MatrixSum) NeuronParameterB() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("neuronParameterB"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsum/2935616-neuronparameterb
func (m_ MatrixSum) SetNeuronParameterB(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNeuronParameterB:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsum/2935618-neuronparameterc
func (m_ MatrixSum) NeuronParameterC() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("neuronParameterC"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsum/2935618-neuronparameterc
func (m_ MatrixSum) SetNeuronParameterC(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNeuronParameterC:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsum/2935620-count
func (m_ MatrixSum) Count() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("count"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsum/2935620-count
func (m_ MatrixSum) SetCount(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCount:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsum/2935621-transpose
func (m_ MatrixSum) Transpose() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("transpose"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsum/2935621-transpose
func (m_ MatrixSum) SetTranspose(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTranspose:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsum/2935622-rows
func (m_ MatrixSum) Rows() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("rows"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsum/2935622-rows
func (m_ MatrixSum) SetRows(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRows:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsum/2935624-neuronparametera
func (m_ MatrixSum) NeuronParameterA() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("neuronParameterA"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsum/2935624-neuronparametera
func (m_ MatrixSum) SetNeuronParameterA(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNeuronParameterA:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsum/3152564-resultmatrixorigin
func (m_ MatrixSum) ResultMatrixOrigin() Origin get set /* not a class type */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("resultMatrixOrigin"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsum/3152564-resultmatrixorigin
func (m_ MatrixSum) SetResultMatrixOrigin(value Origin get set /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setResultMatrixOrigin:"), value)
}







