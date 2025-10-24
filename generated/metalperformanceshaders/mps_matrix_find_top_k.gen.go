// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [MatrixFindTopK] class.
var (
	MatrixFindTopKClass     _MatrixFindTopKClass
	MatrixFindTopKClassOnce sync.Once
)

func getMatrixFindTopKClass() _MatrixFindTopKClass {
	MatrixFindTopKClassOnce.Do(func() {
		MatrixFindTopKClass = _MatrixFindTopKClass{objc.GetClass("MPSMatrixFindTopK")}
	})
	return MatrixFindTopKClass
}

type _MatrixFindTopKClass struct {
	class objc.Class
}





// An interface definition for the [MatrixFindTopK] class.
type IMatrixFindTopK interface {
	IMatrixUnaryKernel
	

	// properties:
	SourceColumns() objectivec.IObject
	SetSourceColumns(value objectivec.IObject)
	IndexOffset() objectivec.IObject
	SetIndexOffset(value objectivec.IObject)
	NumberOfTopKValues() objectivec.IObject
	SetNumberOfTopKValues(value objectivec.IObject)
	SourceRows() objectivec.IObject
	SetSourceRows(value objectivec.IObject)


	

	// methods:
	Encode()
	EncodeToCommandBufferInputMatrixResultIndexMatrixResultValueMatrix(commandBuffer unsafe.Pointer, inputMatrix IMatrix, resultIndexMatrix IMatrix, resultValueMatrix IMatrix)
	CopyWithZoneDevice(zone Zone /* not a class type */, device unsafe.Pointer) objectivec.IObject


}





// Alloc allocates a new instance without initialization.
func (mc _MatrixFindTopKClass) Alloc() MatrixFindTopK {
	rv := objc.Send[MatrixFindTopK](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MatrixFindTopKClass) New() MatrixFindTopK {
	rv := objc.Send[MatrixFindTopK](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MatrixFindTopK) Init() MatrixFindTopK {
	rv := objc.Send[MatrixFindTopK](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MatrixFindTopK) Autorelease() MatrixFindTopK {
	rv := objc.Send[MatrixFindTopK](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMatrixFindTopK creates a new MatrixFindTopK instance.
func NewMatrixFindTopK() MatrixFindTopK {
	return getMatrixFindTopKClass().New()
}





// A kernel for computing the top-K values and their corresponding indices in a matrix.


// A kernel for computing the top-K values and their corresponding indices in a matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixFindTopK
type MatrixFindTopK struct {
	MatrixUnaryKernel
}

// MatrixFindTopKFrom constructs a [MatrixFindTopK] from an unsafe.Pointer.
//
// A kernel for computing the top-K values and their corresponding indices in a matrix.
func MatrixFindTopKFrom(ptr unsafe.Pointer) MatrixFindTopK {
	return MatrixFindTopK{
		MatrixUnaryKernel: MatrixUnaryKernelFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfindtopk/2935582-initwithcoder
func NewMatrixFindTopKWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) MatrixFindTopK {
	instance := getMatrixFindTopKClass().Alloc()
	rv := objc.Send[MatrixFindTopK](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfindtopk/2935575-initwithdevice
func NewMatrixFindTopKWithDeviceNumberOfTopKValues(device unsafe.Pointer, numberOfTopKValues uint) MatrixFindTopK {
	instance := getMatrixFindTopKClass().Alloc()
	rv := objc.Send[MatrixFindTopK](instance.ID, objc.Sel("initWithDevice:numberOfTopKValues:"), device, numberOfTopKValues)
	rv.Autorelease()
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfindtopk/2935579-encode
func (m_ MatrixFindTopK) Encode() {
	objc.Send[objc.ID](m_.ID, objc.Sel("encode"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfindtopk/2935579-encodetocommandbuffer
func (m_ MatrixFindTopK) EncodeToCommandBufferInputMatrixResultIndexMatrixResultValueMatrix(commandBuffer unsafe.Pointer, inputMatrix IMatrix, resultIndexMatrix IMatrix, resultValueMatrix IMatrix) {
	objc.Send[objc.ID](m_.ID, objc.Sel("encodeToCommandBuffer:inputMatrix:resultIndexMatrix:resultValueMatrix:"), commandBuffer, inputMatrix, resultIndexMatrix, resultValueMatrix)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfindtopk/2935581-copywithzone
func (m_ MatrixFindTopK) CopyWithZoneDevice(zone Zone /* not a class type */, device unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("copyWithZone:device:"), zone, device)
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfindtopk/2935573-sourcecolumns
func (m_ MatrixFindTopK) SourceColumns() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("sourceColumns"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfindtopk/2935573-sourcecolumns
func (m_ MatrixFindTopK) SetSourceColumns(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSourceColumns:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfindtopk/2935574-indexoffset
func (m_ MatrixFindTopK) IndexOffset() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("indexOffset"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfindtopk/2935574-indexoffset
func (m_ MatrixFindTopK) SetIndexOffset(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIndexOffset:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfindtopk/2935577-numberoftopkvalues
func (m_ MatrixFindTopK) NumberOfTopKValues() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("numberOfTopKValues"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfindtopk/2935577-numberoftopkvalues
func (m_ MatrixFindTopK) SetNumberOfTopKValues(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNumberOfTopKValues:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfindtopk/2935580-sourcerows
func (m_ MatrixFindTopK) SourceRows() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("sourceRows"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfindtopk/2935580-sourcerows
func (m_ MatrixFindTopK) SetSourceRows(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSourceRows:"), value)
}







