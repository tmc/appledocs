// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [ImageCopyToMatrix] class.
var (
	ImageCopyToMatrixClass     _ImageCopyToMatrixClass
	ImageCopyToMatrixClassOnce sync.Once
)

func getImageCopyToMatrixClass() _ImageCopyToMatrixClass {
	ImageCopyToMatrixClassOnce.Do(func() {
		ImageCopyToMatrixClass = _ImageCopyToMatrixClass{objc.GetClass("MPSImageCopyToMatrix")}
	})
	return ImageCopyToMatrixClass
}

type _ImageCopyToMatrixClass struct {
	class objc.Class
}





// An interface definition for the [ImageCopyToMatrix] class.
type IImageCopyToMatrix interface {
	IKernel
	

	// properties:
	DestinationMatrixBatchIndex() objectivec.IObject
	SetDestinationMatrixBatchIndex(value objectivec.IObject)
	DestinationMatrixOrigin() Origin get set /* not a class type */
	SetDestinationMatrixOrigin(value Origin get set /* not a class type */)
	DataLayout() DataLayout get /* not a class type */
	SetDataLayout(value DataLayout get /* not a class type */)
	FeatureChannels() int
	SetFeatureChannels(value int)


	

	// methods:
	Encode()
	EncodeToCommandBufferSourceImageDestinationMatrix(commandBuffer unsafe.Pointer, sourceImage IImage, destinationMatrix IMatrix)
	EncodeBatch()
	EncodeBatchToCommandBufferSourceImagesDestinationMatrix(commandBuffer unsafe.Pointer, sourceImages ImageBatch /* not a class type */, destinationMatrix IMatrix)


}





// Alloc allocates a new instance without initialization.
func (ic _ImageCopyToMatrixClass) Alloc() ImageCopyToMatrix {
	rv := objc.Send[ImageCopyToMatrix](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageCopyToMatrixClass) New() ImageCopyToMatrix {
	rv := objc.Send[ImageCopyToMatrix](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageCopyToMatrix) Init() ImageCopyToMatrix {
	rv := objc.Send[ImageCopyToMatrix](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageCopyToMatrix) Autorelease() ImageCopyToMatrix {
	rv := objc.Send[ImageCopyToMatrix](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageCopyToMatrix creates a new ImageCopyToMatrix instance.
func NewImageCopyToMatrix() ImageCopyToMatrix {
	return getImageCopyToMatrixClass().New()
}





// A class that copies image data to a matrix.
//
// This kernel copies image data to a object. The image data is stored in a row of a matrix. The specifies the order in which the feature channels in the image get stored in the matrix. If the stores a batch of images, the images are copied into multiple rows, one row per image. The number of elements in a row in the matrix must be greater than the image width multiplied its height multiplied by the number of in the image.


// A class that copies image data to a matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageCopyToMatrix
type ImageCopyToMatrix struct {
	Kernel
}

// ImageCopyToMatrixFrom constructs a [ImageCopyToMatrix] from an unsafe.Pointer.
//
// A class that copies image data to a matrix.
func ImageCopyToMatrixFrom(ptr unsafe.Pointer) ImageCopyToMatrix {
	return ImageCopyToMatrix{
		Kernel: KernelFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecopytomatrix/2873213-initwithcoder
func NewImageCopyToMatrixWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) ImageCopyToMatrix {
	instance := getImageCopyToMatrixClass().Alloc()
	rv := objc.Send[ImageCopyToMatrix](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecopytomatrix/2873210-initwithdevice
func NewImageCopyToMatrixWithDeviceDataLayout(device unsafe.Pointer, dataLayout DataLayout) ImageCopyToMatrix {
	instance := getImageCopyToMatrixClass().Alloc()
	rv := objc.Send[ImageCopyToMatrix](instance.ID, objc.Sel("initWithDevice:dataLayout:"), device, dataLayout)
	rv.Autorelease()
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecopytomatrix/2873212-encode
func (i_ ImageCopyToMatrix) Encode() {
	objc.Send[objc.ID](i_.ID, objc.Sel("encode"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecopytomatrix/2873212-encodetocommandbuffer
func (i_ ImageCopyToMatrix) EncodeToCommandBufferSourceImageDestinationMatrix(commandBuffer unsafe.Pointer, sourceImage IImage, destinationMatrix IMatrix) {
	objc.Send[objc.ID](i_.ID, objc.Sel("encodeToCommandBuffer:sourceImage:destinationMatrix:"), commandBuffer, sourceImage, destinationMatrix)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecopytomatrix/3013769-encodebatch
func (i_ ImageCopyToMatrix) EncodeBatch() {
	objc.Send[objc.ID](i_.ID, objc.Sel("encodeBatch"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecopytomatrix/3013769-encodebatchtocommandbuffer
func (i_ ImageCopyToMatrix) EncodeBatchToCommandBufferSourceImagesDestinationMatrix(commandBuffer unsafe.Pointer, sourceImages ImageBatch /* not a class type */, destinationMatrix IMatrix) {
	objc.Send[objc.ID](i_.ID, objc.Sel("encodeBatchToCommandBuffer:sourceImages:destinationMatrix:"), commandBuffer, sourceImages, destinationMatrix)
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecopytomatrix/2873211-destinationmatrixbatchindex
func (i_ ImageCopyToMatrix) DestinationMatrixBatchIndex() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("destinationMatrixBatchIndex"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecopytomatrix/2873211-destinationmatrixbatchindex
func (i_ ImageCopyToMatrix) SetDestinationMatrixBatchIndex(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDestinationMatrixBatchIndex:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecopytomatrix/2873215-destinationmatrixorigin
func (i_ ImageCopyToMatrix) DestinationMatrixOrigin() Origin get set /* not a class type */ {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("destinationMatrixOrigin"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecopytomatrix/2873215-destinationmatrixorigin
func (i_ ImageCopyToMatrix) SetDestinationMatrixOrigin(value Origin get set /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDestinationMatrixOrigin:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecopytomatrix/2873216-datalayout
func (i_ ImageCopyToMatrix) DataLayout() DataLayout get /* not a class type */ {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("dataLayout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecopytomatrix/2873216-datalayout
func (i_ ImageCopyToMatrix) SetDataLayout(value DataLayout get /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDataLayout:"), value)
}


// The number of feature channels per pixel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/featurechannels
func (i_ ImageCopyToMatrix) FeatureChannels() int {
	rv := objc.Send[int](i_.ID, objc.Sel("featureChannels"))
	return rv
}


// The number of feature channels per pixel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/featurechannels
func (i_ ImageCopyToMatrix) SetFeatureChannels(value int) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setFeatureChannels:"), value)
}







