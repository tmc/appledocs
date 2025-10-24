// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [MatrixCopyToImage] class.
var (
	MatrixCopyToImageClass     _MatrixCopyToImageClass
	MatrixCopyToImageClassOnce sync.Once
)

func getMatrixCopyToImageClass() _MatrixCopyToImageClass {
	MatrixCopyToImageClassOnce.Do(func() {
		MatrixCopyToImageClass = _MatrixCopyToImageClass{objc.GetClass("MPSMatrixCopyToImage")}
	})
	return MatrixCopyToImageClass
}

type _MatrixCopyToImageClass struct {
	class objc.Class
}





// An interface definition for the [MatrixCopyToImage] class.
type IMatrixCopyToImage interface {
	IKernel
	

	// properties:
	DataLayout() DataLayout get /* not a class type */
	SetDataLayout(value DataLayout get /* not a class type */)
	SourceMatrixBatchIndex() objectivec.IObject
	SetSourceMatrixBatchIndex(value objectivec.IObject)
	SourceMatrixOrigin() Origin get set /* not a class type */
	SetSourceMatrixOrigin(value Origin get set /* not a class type */)


	

	// methods:
	Encode()
	EncodeToCommandBufferSourceMatrixDestinationImage(commandBuffer unsafe.Pointer, sourceMatrix IMatrix, destinationImage IImage)
	EncodeBatch()
	EncodeBatchToCommandBufferSourceMatrixDestinationImages(commandBuffer unsafe.Pointer, sourceMatrix IMatrix, destinationImages ImageBatch /* not a class type */)


}





// Alloc allocates a new instance without initialization.
func (mc _MatrixCopyToImageClass) Alloc() MatrixCopyToImage {
	rv := objc.Send[MatrixCopyToImage](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MatrixCopyToImageClass) New() MatrixCopyToImage {
	rv := objc.Send[MatrixCopyToImage](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MatrixCopyToImage) Init() MatrixCopyToImage {
	rv := objc.Send[MatrixCopyToImage](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MatrixCopyToImage) Autorelease() MatrixCopyToImage {
	rv := objc.Send[MatrixCopyToImage](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMatrixCopyToImage creates a new MatrixCopyToImage instance.
func NewMatrixCopyToImage() MatrixCopyToImage {
	return getMatrixCopyToImageClass().New()
}





// A kernel that copies matrix data to a Metal Performance Shaders image.


// A kernel that copies matrix data to a Metal Performance Shaders image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixCopyToImage
type MatrixCopyToImage struct {
	Kernel
}

// MatrixCopyToImageFrom constructs a [MatrixCopyToImage] from an unsafe.Pointer.
//
// A kernel that copies matrix data to a Metal Performance Shaders image.
func MatrixCopyToImageFrom(ptr unsafe.Pointer) MatrixCopyToImage {
	return MatrixCopyToImage{
		Kernel: KernelFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixcopytoimage/2976459-initwithcoder
func NewMatrixCopyToImageWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) MatrixCopyToImage {
	instance := getMatrixCopyToImageClass().Alloc()
	rv := objc.Send[MatrixCopyToImage](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixcopytoimage/2976460-initwithdevice
func NewMatrixCopyToImageWithDeviceDataLayout(device unsafe.Pointer, dataLayout DataLayout) MatrixCopyToImage {
	instance := getMatrixCopyToImageClass().Alloc()
	rv := objc.Send[MatrixCopyToImage](instance.ID, objc.Sel("initWithDevice:dataLayout:"), device, dataLayout)
	rv.Autorelease()
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixcopytoimage/2976458-encode
func (m_ MatrixCopyToImage) Encode() {
	objc.Send[objc.ID](m_.ID, objc.Sel("encode"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixcopytoimage/2976458-encodetocommandbuffer
func (m_ MatrixCopyToImage) EncodeToCommandBufferSourceMatrixDestinationImage(commandBuffer unsafe.Pointer, sourceMatrix IMatrix, destinationImage IImage) {
	objc.Send[objc.ID](m_.ID, objc.Sel("encodeToCommandBuffer:sourceMatrix:destinationImage:"), commandBuffer, sourceMatrix, destinationImage)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixcopytoimage/3013770-encodebatch
func (m_ MatrixCopyToImage) EncodeBatch() {
	objc.Send[objc.ID](m_.ID, objc.Sel("encodeBatch"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixcopytoimage/3013770-encodebatchtocommandbuffer
func (m_ MatrixCopyToImage) EncodeBatchToCommandBufferSourceMatrixDestinationImages(commandBuffer unsafe.Pointer, sourceMatrix IMatrix, destinationImages ImageBatch /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("encodeBatchToCommandBuffer:sourceMatrix:destinationImages:"), commandBuffer, sourceMatrix, destinationImages)
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixcopytoimage/2976457-datalayout
func (m_ MatrixCopyToImage) DataLayout() DataLayout get /* not a class type */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("dataLayout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixcopytoimage/2976457-datalayout
func (m_ MatrixCopyToImage) SetDataLayout(value DataLayout get /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDataLayout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixcopytoimage/2976461-sourcematrixbatchindex
func (m_ MatrixCopyToImage) SourceMatrixBatchIndex() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("sourceMatrixBatchIndex"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixcopytoimage/2976461-sourcematrixbatchindex
func (m_ MatrixCopyToImage) SetSourceMatrixBatchIndex(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSourceMatrixBatchIndex:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixcopytoimage/2976462-sourcematrixorigin
func (m_ MatrixCopyToImage) SourceMatrixOrigin() Origin get set /* not a class type */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("sourceMatrixOrigin"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixcopytoimage/2976462-sourcematrixorigin
func (m_ MatrixCopyToImage) SetSourceMatrixOrigin(value Origin get set /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSourceMatrixOrigin:"), value)
}







