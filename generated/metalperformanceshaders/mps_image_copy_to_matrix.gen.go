// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSImageCopyToMatrix */


/* debug [class_header]: Header for MPSImageCopyToMatrix */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ImageCopyToMatrix */
// An interface definition for the [ImageCopyToMatrix] class.
type IImageCopyToMatrix interface {
	IKernel
	
/* debug [class_interface_properties]: Properties for ImageCopyToMatrix */
	// properties:
	DestinationMatrixBatchIndex() objectivec.IObject
	SetDestinationMatrixBatchIndex(value objectivec.IObject)
	DestinationMatrixOrigin() Origin get set /* not a class type */
	SetDestinationMatrixOrigin(value Origin get set /* not a class type */)
	DataLayout() DataLayout get /* not a class type */
	SetDataLayout(value DataLayout get /* not a class type */)
	FeatureChannels() int
	SetFeatureChannels(value int)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ImageCopyToMatrix */
	// methods:
	Encode()
	EncodeToCommandBufferSourceImageDestinationMatrix(commandBuffer unsafe.Pointer, sourceImage IImage, destinationMatrix IMatrix)
	EncodeBatch()
	EncodeBatchToCommandBufferSourceImagesDestinationMatrix(commandBuffer unsafe.Pointer, sourceImages ImageBatch /* not a class type */, destinationMatrix IMatrix)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ImageCopyToMatrix */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ImageCopyToMatrix */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ImageCopyToMatrix */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecopytomatrix/2873213-initwithcoder
func NewImageCopyToMatrixWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) ImageCopyToMatrix {
	instance := getImageCopyToMatrixClass().Alloc()
	rv := objc.Send[ImageCopyToMatrix](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageCopyToMatrixWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecopytomatrix/2873210-initwithdevice
func NewImageCopyToMatrixWithDeviceDataLayout(device unsafe.Pointer, dataLayout DataLayout) ImageCopyToMatrix {
	instance := getImageCopyToMatrixClass().Alloc()
	rv := objc.Send[ImageCopyToMatrix](instance.ID, objc.Sel("initWithDevice:dataLayout:"), device, dataLayout)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageCopyToMatrixWithDeviceDataLayout */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ImageCopyToMatrix */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ImageCopyToMatrix */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ImageCopyToMatrix */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecopytomatrix/2873212-encode
func (i_ ImageCopyToMatrix) Encode() {
	objc.Send[objc.ID](i_.ID, objc.Sel("encode"))
}/* debug [instance_methods/method]: Encode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecopytomatrix/2873212-encodetocommandbuffer
func (i_ ImageCopyToMatrix) EncodeToCommandBufferSourceImageDestinationMatrix(commandBuffer unsafe.Pointer, sourceImage IImage, destinationMatrix IMatrix) {
	objc.Send[objc.ID](i_.ID, objc.Sel("encodeToCommandBuffer:sourceImage:destinationMatrix:"), commandBuffer, sourceImage, destinationMatrix)
}/* debug [instance_methods/method]: EncodeToCommandBufferSourceImageDestinationMatrix */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecopytomatrix/3013769-encodebatch
func (i_ ImageCopyToMatrix) EncodeBatch() {
	objc.Send[objc.ID](i_.ID, objc.Sel("encodeBatch"))
}/* debug [instance_methods/method]: EncodeBatch */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecopytomatrix/3013769-encodebatchtocommandbuffer
func (i_ ImageCopyToMatrix) EncodeBatchToCommandBufferSourceImagesDestinationMatrix(commandBuffer unsafe.Pointer, sourceImages ImageBatch /* not a class type */, destinationMatrix IMatrix) {
	objc.Send[objc.ID](i_.ID, objc.Sel("encodeBatchToCommandBuffer:sourceImages:destinationMatrix:"), commandBuffer, sourceImages, destinationMatrix)
}/* debug [instance_methods/method]: EncodeBatchToCommandBufferSourceImagesDestinationMatrix */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ImageCopyToMatrix */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecopytomatrix/2873211-destinationmatrixbatchindex
func (i_ ImageCopyToMatrix) DestinationMatrixBatchIndex() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("destinationMatrixBatchIndex"))
	return rv
}/* debug [instance_properties/getter]: destinationMatrixBatchIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecopytomatrix/2873211-destinationmatrixbatchindex
func (i_ ImageCopyToMatrix) SetDestinationMatrixBatchIndex(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDestinationMatrixBatchIndex:"), value)
}/* debug [instance_properties/setter]: destinationMatrixBatchIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecopytomatrix/2873215-destinationmatrixorigin
func (i_ ImageCopyToMatrix) DestinationMatrixOrigin() Origin get set /* not a class type */ {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("destinationMatrixOrigin"))
	return rv
}/* debug [instance_properties/getter]: destinationMatrixOrigin */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecopytomatrix/2873215-destinationmatrixorigin
func (i_ ImageCopyToMatrix) SetDestinationMatrixOrigin(value Origin get set /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDestinationMatrixOrigin:"), value)
}/* debug [instance_properties/setter]: destinationMatrixOrigin */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecopytomatrix/2873216-datalayout
func (i_ ImageCopyToMatrix) DataLayout() DataLayout get /* not a class type */ {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("dataLayout"))
	return rv
}/* debug [instance_properties/getter]: dataLayout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecopytomatrix/2873216-datalayout
func (i_ ImageCopyToMatrix) SetDataLayout(value DataLayout get /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDataLayout:"), value)
}/* debug [instance_properties/setter]: dataLayout */


// The number of feature channels per pixel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/featurechannels
func (i_ ImageCopyToMatrix) FeatureChannels() int {
	rv := objc.Send[int](i_.ID, objc.Sel("featureChannels"))
	return rv
}/* debug [instance_properties/getter]: featureChannels */


// The number of feature channels per pixel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/featurechannels
func (i_ ImageCopyToMatrix) SetFeatureChannels(value int) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setFeatureChannels:"), value)
}/* debug [instance_properties/setter]: featureChannels */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSImageCopyToMatrix */


