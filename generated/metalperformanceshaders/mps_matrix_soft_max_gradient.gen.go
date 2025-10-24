// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [MatrixSoftMaxGradient] class.
var (
	MatrixSoftMaxGradientClass     _MatrixSoftMaxGradientClass
	MatrixSoftMaxGradientClassOnce sync.Once
)

func getMatrixSoftMaxGradientClass() _MatrixSoftMaxGradientClass {
	MatrixSoftMaxGradientClassOnce.Do(func() {
		MatrixSoftMaxGradientClass = _MatrixSoftMaxGradientClass{objc.GetClass("MPSMatrixSoftMaxGradient")}
	})
	return MatrixSoftMaxGradientClass
}

type _MatrixSoftMaxGradientClass struct {
	class objc.Class
}





// An interface definition for the [MatrixSoftMaxGradient] class.
type IMatrixSoftMaxGradient interface {
	IMatrixBinaryKernel
	

	// properties:
	SourceColumns() objectivec.IObject
	SetSourceColumns(value objectivec.IObject)
	SourceRows() objectivec.IObject
	SetSourceRows(value objectivec.IObject)


	

	// methods:
	CopyWithZoneDevice(zone Zone /* not a class type */, device unsafe.Pointer) objectivec.IObject
	Encode()
	EncodeToCommandBufferGradientMatrixForwardOutputMatrixResultMatrix(commandBuffer unsafe.Pointer, gradientMatrix IMatrix, forwardOutputMatrix IMatrix, resultMatrix IMatrix)


}





// Alloc allocates a new instance without initialization.
func (mc _MatrixSoftMaxGradientClass) Alloc() MatrixSoftMaxGradient {
	rv := objc.Send[MatrixSoftMaxGradient](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MatrixSoftMaxGradientClass) New() MatrixSoftMaxGradient {
	rv := objc.Send[MatrixSoftMaxGradient](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MatrixSoftMaxGradient) Init() MatrixSoftMaxGradient {
	rv := objc.Send[MatrixSoftMaxGradient](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MatrixSoftMaxGradient) Autorelease() MatrixSoftMaxGradient {
	rv := objc.Send[MatrixSoftMaxGradient](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMatrixSoftMaxGradient creates a new MatrixSoftMaxGradient instance.
func NewMatrixSoftMaxGradient() MatrixSoftMaxGradient {
	return getMatrixSoftMaxGradientClass().New()
}





// A gradient softmax kernel that operates on matrices.


// A gradient softmax kernel that operates on matrices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixSoftMaxGradient
type MatrixSoftMaxGradient struct {
	MatrixBinaryKernel
}

// MatrixSoftMaxGradientFrom constructs a [MatrixSoftMaxGradient] from an unsafe.Pointer.
//
// A gradient softmax kernel that operates on matrices.
func MatrixSoftMaxGradientFrom(ptr unsafe.Pointer) MatrixSoftMaxGradient {
	return MatrixSoftMaxGradient{
		MatrixBinaryKernel: MatrixBinaryKernelFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsoftmaxgradient/2966653-initwithcoder
func NewMatrixSoftMaxGradientWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) MatrixSoftMaxGradient {
	instance := getMatrixSoftMaxGradientClass().Alloc()
	rv := objc.Send[MatrixSoftMaxGradient](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsoftmaxgradient/2966654-initwithdevice
func NewMatrixSoftMaxGradientWithDevice(device unsafe.Pointer) MatrixSoftMaxGradient {
	instance := getMatrixSoftMaxGradientClass().Alloc()
	rv := objc.Send[MatrixSoftMaxGradient](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsoftmaxgradient/2966651-copywithzone
func (m_ MatrixSoftMaxGradient) CopyWithZoneDevice(zone Zone /* not a class type */, device unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("copyWithZone:device:"), zone, device)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsoftmaxgradient/2966652-encode
func (m_ MatrixSoftMaxGradient) Encode() {
	objc.Send[objc.ID](m_.ID, objc.Sel("encode"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsoftmaxgradient/2966652-encodetocommandbuffer
func (m_ MatrixSoftMaxGradient) EncodeToCommandBufferGradientMatrixForwardOutputMatrixResultMatrix(commandBuffer unsafe.Pointer, gradientMatrix IMatrix, forwardOutputMatrix IMatrix, resultMatrix IMatrix) {
	objc.Send[objc.ID](m_.ID, objc.Sel("encodeToCommandBuffer:gradientMatrix:forwardOutputMatrix:resultMatrix:"), commandBuffer, gradientMatrix, forwardOutputMatrix, resultMatrix)
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsoftmaxgradient/2966655-sourcecolumns
func (m_ MatrixSoftMaxGradient) SourceColumns() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("sourceColumns"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsoftmaxgradient/2966655-sourcecolumns
func (m_ MatrixSoftMaxGradient) SetSourceColumns(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSourceColumns:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsoftmaxgradient/2966656-sourcerows
func (m_ MatrixSoftMaxGradient) SourceRows() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("sourceRows"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsoftmaxgradient/2966656-sourcerows
func (m_ MatrixSoftMaxGradient) SetSourceRows(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSourceRows:"), value)
}







