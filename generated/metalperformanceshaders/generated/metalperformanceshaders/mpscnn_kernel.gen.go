// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNKernel */


/* debug [class_header]: Header for MPSCNNKernel */
// The class instance for the [CNNKernel] class.
var (
	CNNKernelClass     _CNNKernelClass
	CNNKernelClassOnce sync.Once
)

func getCNNKernelClass() _CNNKernelClass {
	CNNKernelClassOnce.Do(func() {
		CNNKernelClass = _CNNKernelClass{objc.GetClass("MPSCNNKernel")}
	})
	return CNNKernelClass
}

type _CNNKernelClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNKernel */
// An interface definition for the [CNNKernel] class.
type ICNNKernel interface {
	IKernel
	
/* debug [class_interface_properties]: Properties for CNNKernel */
	// properties:
	EdgeMode() ImageEdgeMode get set /* not a class type */
	SetEdgeMode(value ImageEdgeMode get set /* not a class type */)
	Offset() Offset get set /* not a class type */
	SetOffset(value Offset get set /* not a class type */)
	ClipRect() Region get set /* not a class type */
	SetClipRect(value Region get set /* not a class type */)
	DestinationFeatureChannelOffset() objectivec.IObject
	SetDestinationFeatureChannelOffset(value objectivec.IObject)
	IsBackwards() objectivec.IObject
	SetIsBackwards(value objectivec.IObject)
	KernelWidth() objectivec.IObject
	SetKernelWidth(value objectivec.IObject)
	StrideInPixelsY() objectivec.IObject
	SetStrideInPixelsY(value objectivec.IObject)
	KernelHeight() objectivec.IObject
	SetKernelHeight(value objectivec.IObject)
	DestinationImageAllocator() ImageAllocator get set /* not a class type */
	SetDestinationImageAllocator(value ImageAllocator get set /* not a class type */)
	StrideInPixelsX() objectivec.IObject
	SetStrideInPixelsX(value objectivec.IObject)
	Padding() Padding get set /* not a class type */
	SetPadding(value Padding get set /* not a class type */)
	DilationRateX() objectivec.IObject
	SetDilationRateX(value objectivec.IObject)
	IsStateModified() objectivec.IObject
	SetIsStateModified(value objectivec.IObject)
	DilationRateY() objectivec.IObject
	SetDilationRateY(value objectivec.IObject)
	SourceFeatureChannelOffset() objectivec.IObject
	SetSourceFeatureChannelOffset(value objectivec.IObject)
	SourceFeatureChannelMaxCount() objectivec.IObject
	SetSourceFeatureChannelMaxCount(value objectivec.IObject)
	Origin() MTLOrigin /* not a class type */
	SetOrigin(value MTLOrigin /* not a class type */)
	Size() objc.IObject /* cross-framework: MPSSize */
	SetSize(value objc.IObject /* cross-framework: MPSSize */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNKernel */
	// methods:
	Encode()
	EncodeToCommandBufferSourceImageDestinationImage(commandBuffer unsafe.Pointer, sourceImage IImage, destinationImage IImage)
	EncodeToCommandBufferSourceImage(commandBuffer unsafe.Pointer, sourceImage IImage) IImage
	EncodeBatch()
	EncodeBatchToCommandBufferSourceImagesDestinationStatesDestinationStateIsTemporary(commandBuffer unsafe.Pointer, sourceImages ImageBatch /* not a class type */, outStates StateBatch /* not a class type */, isTemporary bool) ImageBatch /* not a class type */
	EncodeBatchToCommandBufferSourceImagesDestinationStatesDestinationImages(commandBuffer unsafe.Pointer, sourceImages ImageBatch /* not a class type */, destinationStates StateBatch /* not a class type */, destinationImages ImageBatch /* not a class type */)
	EncodeBatchToCommandBufferSourceImages(commandBuffer unsafe.Pointer, sourceImages ImageBatch /* not a class type */) ImageBatch /* not a class type */
	DestinationImageDescriptor()
	DestinationImageDescriptorForSourceImagesSourceStates(sourceImages unsafe.Pointer, sourceStates unsafe.Pointer) IImageDescriptor
	IsResultStateReusedAcrossBatch()
	AppendBatchBarrier()
	EncodeToCommandBufferSourceImageDestinationStateDestinationImage(commandBuffer unsafe.Pointer, sourceImage IImage, destinationState IState, destinationImage IImage)
	EncodeToCommandBufferSourceImageDestinationStateDestinationStateIsTemporary(commandBuffer unsafe.Pointer, sourceImage IImage, outState objectivec.IObject, isTemporary bool) IImage
	EncodeBatchToCommandBufferSourceImagesDestinationImages(commandBuffer unsafe.Pointer, sourceImages ImageBatch /* not a class type */, destinationImages ImageBatch /* not a class type */)
	ResultStateBatch()
	ResultStateBatchForSourceImageSourceStatesDestinationImage(sourceImage ImageBatch /* not a class type */, sourceStates StateBatch /* not a class type */, destinationImage ImageBatch /* not a class type */) StateBatch /* not a class type */
	ResultState()
	ResultStateForSourceImageSourceStatesDestinationImage(sourceImage IImage, sourceStates unsafe.Pointer, destinationImage IImage) IState
	TemporaryResultState()
	TemporaryResultStateForCommandBufferSourceImageSourceStatesDestinationImage(commandBuffer unsafe.Pointer, sourceImage IImage, sourceStates unsafe.Pointer, destinationImage IImage) IState
	TemporaryResultStateBatch()
	TemporaryResultStateBatchForCommandBufferSourceImageSourceStatesDestinationImage(commandBuffer unsafe.Pointer, sourceImage ImageBatch /* not a class type */, sourceStates StateBatch /* not a class type */, destinationImage ImageBatch /* not a class type */) StateBatch /* not a class type */
	BatchEncodingStorageSize()
	BatchEncodingStorageSizeForSourceImageSourceStatesDestinationImage(sourceImage ImageBatch /* not a class type */, sourceStates StateBatch /* not a class type */, destinationImage ImageBatch /* not a class type */) uint
	EncodingStorageSize()
	EncodingStorageSizeForSourceImageSourceStatesDestinationImage(sourceImage IImage, sourceStates unsafe.Pointer, destinationImage IImage) uint
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNKernel */
// Alloc allocates a new instance without initialization.
func (cc _CNNKernelClass) Alloc() CNNKernel {
	rv := objc.Send[CNNKernel](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNKernelClass) New() CNNKernel {
	rv := objc.Send[CNNKernel](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNKernel) Init() CNNKernel {
	rv := objc.Send[CNNKernel](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNKernel) Autorelease() CNNKernel {
	rv := objc.Send[CNNKernel](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNKernel creates a new CNNKernel instance.
func NewCNNKernel() CNNKernel {
	return getCNNKernelClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNKernel */
// Base class for neural network layers.
//
// An object consumes one object and produces one object. The region overwritten in the destination image is described by the property. The top left corner of the region consumed (ignoring adjustments for filter size—for example, convolution filter size) is given by the property. The size of the region consumed is a function of the size of the property and any subsampling caused by pixel strides at work (for example, / in the class). Wherever the and properties would cause an pixel address not in the image to be read, the property is used to determine what value to read there. The or component of the , and properties indexes which images to use. If the object contains only a single image, then these values should be , , and . If the object contains multiple images, then the value of determines the number of images to process. Both the source and destination objects must have at least this many images. The value of refers to the starting image index of the source. Thus, the value of must be . Similarly, the value of determines the starting image index of the destination. Thus, the value of must be . The property can be used to control where the kernel will start writing in terms of feature channel dimension. For example, if the destination has 64 channels and th e kernel outputs 32 channels, channels 0-31 of the destination will be populated by the kernel (by default). But if you want the kernel to populate channels 32-63 of the destination, you can set the value of to 32. Suppose you have a source of dimensions , where is the number of channels, which goes through a convolution filter which produces the output and 1 which produces the output followed by concatenation which produces . You can achieve this by creating an object with dimensions and using this as the destination of both convolutions as follows: , this will output channels starting at channel of destination thus populating channels. , this will output channels starting at channel of destination thus populating channels.


// Base class for neural network layers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel
type CNNKernel struct {
	Kernel
}

// CNNKernelFrom constructs a [CNNKernel] from an unsafe.Pointer.
//
// Base class for neural network layers.
func CNNKernelFrom(ptr unsafe.Pointer) CNNKernel {
	return CNNKernel{
		Kernel: KernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNKernel */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2865655-initwithcoder
func NewCNNKernelWithCoderDevice(aDecoder Coder /* not a class type */, device unsafe.Pointer) CNNKernel {
	instance := getCNNKernelClass().Alloc()
	rv := objc.Send[CNNKernel](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNKernelWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2865653-initwithdevice
func NewCNNKernelWithDevice(device unsafe.Pointer) CNNKernel {
	instance := getCNNKernelClass().Alloc()
	rv := objc.Send[CNNKernel](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNKernelWithDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNKernel */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNKernel */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNKernel */

// Encodes a kernel into a command buffer. The ensuing operation proceeds out-of-place.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/1648919-encode
func (c_ CNNKernel) Encode() {
	objc.Send[objc.ID](c_.ID, objc.Sel("encode"))
}/* debug [instance_methods/method]: Encode */


// Encodes a kernel into a command buffer. The ensuing operation proceeds out-of-place.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/1648919-encodetocommandbuffer
func (c_ CNNKernel) EncodeToCommandBufferSourceImageDestinationImage(commandBuffer unsafe.Pointer, sourceImage IImage, destinationImage IImage) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeToCommandBuffer:sourceImage:destinationImage:"), commandBuffer, sourceImage, destinationImage)
}/* debug [instance_methods/method]: EncodeToCommandBufferSourceImageDestinationImage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2865642-encodetocommandbuffer
func (c_ CNNKernel) EncodeToCommandBufferSourceImage(commandBuffer unsafe.Pointer, sourceImage IImage) IImage {
	rv := objc.Send[Image](c_.ID, objc.Sel("encodeToCommandBuffer:sourceImage:"), commandBuffer, sourceImage)
	return rv
}/* debug [instance_methods/method]: EncodeToCommandBufferSourceImage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2942646-encodebatch
func (c_ CNNKernel) EncodeBatch() {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeBatch"))
}/* debug [instance_methods/method]: EncodeBatch */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2942646-encodebatchtocommandbuffer
func (c_ CNNKernel) EncodeBatchToCommandBufferSourceImagesDestinationStatesDestinationStateIsTemporary(commandBuffer unsafe.Pointer, sourceImages ImageBatch /* not a class type */, outStates StateBatch /* not a class type */, isTemporary bool) ImageBatch /* not a class type */ {
	rv := objc.Send[ImageBatch](c_.ID, objc.Sel("encodeBatchToCommandBuffer:sourceImages:destinationStates:destinationStateIsTemporary:"), commandBuffer, sourceImages, outStates, isTemporary)
	return rv
}/* debug [instance_methods/method]: EncodeBatchToCommandBufferSourceImagesDestinationStatesDestinationStateIsTemporary */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2942649-encodebatchtocommandbuffer
func (c_ CNNKernel) EncodeBatchToCommandBufferSourceImagesDestinationStatesDestinationImages(commandBuffer unsafe.Pointer, sourceImages ImageBatch /* not a class type */, destinationStates StateBatch /* not a class type */, destinationImages ImageBatch /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeBatchToCommandBuffer:sourceImages:destinationStates:destinationImages:"), commandBuffer, sourceImages, destinationStates, destinationImages)
}/* debug [instance_methods/method]: EncodeBatchToCommandBufferSourceImagesDestinationStatesDestinationImages */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2942651-encodebatchtocommandbuffer
func (c_ CNNKernel) EncodeBatchToCommandBufferSourceImages(commandBuffer unsafe.Pointer, sourceImages ImageBatch /* not a class type */) ImageBatch /* not a class type */ {
	rv := objc.Send[ImageBatch](c_.ID, objc.Sel("encodeBatchToCommandBuffer:sourceImages:"), commandBuffer, sourceImages)
	return rv
}/* debug [instance_methods/method]: EncodeBatchToCommandBufferSourceImages */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2942661-destinationimagedescriptor
func (c_ CNNKernel) DestinationImageDescriptor() {
	objc.Send[objc.ID](c_.ID, objc.Sel("destinationImageDescriptor"))
}/* debug [instance_methods/method]: DestinationImageDescriptor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2942661-destinationimagedescriptorforsou
func (c_ CNNKernel) DestinationImageDescriptorForSourceImagesSourceStates(sourceImages unsafe.Pointer, sourceStates unsafe.Pointer) IImageDescriptor {
	rv := objc.Send[ImageDescriptor](c_.ID, objc.Sel("destinationImageDescriptorForSourceImages:sourceStates:"), sourceImages, sourceStates)
	return rv
}/* debug [instance_methods/method]: DestinationImageDescriptorForSourceImagesSourceStates */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2942665-isresultstatereusedacrossbatch
func (c_ CNNKernel) IsResultStateReusedAcrossBatch() {
	objc.Send[objc.ID](c_.ID, objc.Sel("isResultStateReusedAcrossBatch"))
}/* debug [instance_methods/method]: IsResultStateReusedAcrossBatch */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2942671-appendbatchbarrier
func (c_ CNNKernel) AppendBatchBarrier() {
	objc.Send[objc.ID](c_.ID, objc.Sel("appendBatchBarrier"))
}/* debug [instance_methods/method]: AppendBatchBarrier */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2942672-encodetocommandbuffer
func (c_ CNNKernel) EncodeToCommandBufferSourceImageDestinationStateDestinationImage(commandBuffer unsafe.Pointer, sourceImage IImage, destinationState IState, destinationImage IImage) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeToCommandBuffer:sourceImage:destinationState:destinationImage:"), commandBuffer, sourceImage, destinationState, destinationImage)
}/* debug [instance_methods/method]: EncodeToCommandBufferSourceImageDestinationStateDestinationImage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2942680-encodetocommandbuffer
func (c_ CNNKernel) EncodeToCommandBufferSourceImageDestinationStateDestinationStateIsTemporary(commandBuffer unsafe.Pointer, sourceImage IImage, outState objectivec.IObject, isTemporary bool) IImage {
	rv := objc.Send[Image](c_.ID, objc.Sel("encodeToCommandBuffer:sourceImage:destinationState:destinationStateIsTemporary:"), commandBuffer, sourceImage, outState, isTemporary)
	return rv
}/* debug [instance_methods/method]: EncodeToCommandBufferSourceImageDestinationStateDestinationStateIsTemporary */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2942681-encodebatchtocommandbuffer
func (c_ CNNKernel) EncodeBatchToCommandBufferSourceImagesDestinationImages(commandBuffer unsafe.Pointer, sourceImages ImageBatch /* not a class type */, destinationImages ImageBatch /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeBatchToCommandBuffer:sourceImages:destinationImages:"), commandBuffer, sourceImages, destinationImages)
}/* debug [instance_methods/method]: EncodeBatchToCommandBufferSourceImagesDestinationImages */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2947931-resultstatebatch
func (c_ CNNKernel) ResultStateBatch() {
	objc.Send[objc.ID](c_.ID, objc.Sel("resultStateBatch"))
}/* debug [instance_methods/method]: ResultStateBatch */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2947931-resultstatebatchforsourceimage
func (c_ CNNKernel) ResultStateBatchForSourceImageSourceStatesDestinationImage(sourceImage ImageBatch /* not a class type */, sourceStates StateBatch /* not a class type */, destinationImage ImageBatch /* not a class type */) StateBatch /* not a class type */ {
	rv := objc.Send[StateBatch](c_.ID, objc.Sel("resultStateBatchForSourceImage:sourceStates:destinationImage:"), sourceImage, sourceStates, destinationImage)
	return rv
}/* debug [instance_methods/method]: ResultStateBatchForSourceImageSourceStatesDestinationImage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2947932-resultstate
func (c_ CNNKernel) ResultState() {
	objc.Send[objc.ID](c_.ID, objc.Sel("resultState"))
}/* debug [instance_methods/method]: ResultState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2947932-resultstateforsourceimage
func (c_ CNNKernel) ResultStateForSourceImageSourceStatesDestinationImage(sourceImage IImage, sourceStates unsafe.Pointer, destinationImage IImage) IState {
	rv := objc.Send[State](c_.ID, objc.Sel("resultStateForSourceImage:sourceStates:destinationImage:"), sourceImage, sourceStates, destinationImage)
	return rv
}/* debug [instance_methods/method]: ResultStateForSourceImageSourceStatesDestinationImage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2947937-temporaryresultstate
func (c_ CNNKernel) TemporaryResultState() {
	objc.Send[objc.ID](c_.ID, objc.Sel("temporaryResultState"))
}/* debug [instance_methods/method]: TemporaryResultState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2947937-temporaryresultstateforcommandbu
func (c_ CNNKernel) TemporaryResultStateForCommandBufferSourceImageSourceStatesDestinationImage(commandBuffer unsafe.Pointer, sourceImage IImage, sourceStates unsafe.Pointer, destinationImage IImage) IState {
	rv := objc.Send[State](c_.ID, objc.Sel("temporaryResultStateForCommandBuffer:sourceImage:sourceStates:destinationImage:"), commandBuffer, sourceImage, sourceStates, destinationImage)
	return rv
}/* debug [instance_methods/method]: TemporaryResultStateForCommandBufferSourceImageSourceStatesDestinationImage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2947939-temporaryresultstatebatch
func (c_ CNNKernel) TemporaryResultStateBatch() {
	objc.Send[objc.ID](c_.ID, objc.Sel("temporaryResultStateBatch"))
}/* debug [instance_methods/method]: TemporaryResultStateBatch */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2947939-temporaryresultstatebatchforcomm
func (c_ CNNKernel) TemporaryResultStateBatchForCommandBufferSourceImageSourceStatesDestinationImage(commandBuffer unsafe.Pointer, sourceImage ImageBatch /* not a class type */, sourceStates StateBatch /* not a class type */, destinationImage ImageBatch /* not a class type */) StateBatch /* not a class type */ {
	rv := objc.Send[StateBatch](c_.ID, objc.Sel("temporaryResultStateBatchForCommandBuffer:sourceImage:sourceStates:destinationImage:"), commandBuffer, sourceImage, sourceStates, destinationImage)
	return rv
}/* debug [instance_methods/method]: TemporaryResultStateBatchForCommandBufferSourceImageSourceStatesDestinationImage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/3237263-batchencodingstoragesize
func (c_ CNNKernel) BatchEncodingStorageSize() {
	objc.Send[objc.ID](c_.ID, objc.Sel("batchEncodingStorageSize"))
}/* debug [instance_methods/method]: BatchEncodingStorageSize */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/3237263-batchencodingstoragesizeforsourc
func (c_ CNNKernel) BatchEncodingStorageSizeForSourceImageSourceStatesDestinationImage(sourceImage ImageBatch /* not a class type */, sourceStates StateBatch /* not a class type */, destinationImage ImageBatch /* not a class type */) uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("batchEncodingStorageSizeForSourceImage:sourceStates:destinationImage:"), sourceImage, sourceStates, destinationImage)
	return rv
}/* debug [instance_methods/method]: BatchEncodingStorageSizeForSourceImageSourceStatesDestinationImage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/3237264-encodingstoragesize
func (c_ CNNKernel) EncodingStorageSize() {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodingStorageSize"))
}/* debug [instance_methods/method]: EncodingStorageSize */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/3237264-encodingstoragesizeforsourceimag
func (c_ CNNKernel) EncodingStorageSizeForSourceImageSourceStatesDestinationImage(sourceImage IImage, sourceStates unsafe.Pointer, destinationImage IImage) uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("encodingStorageSizeForSourceImage:sourceStates:destinationImage:"), sourceImage, sourceStates, destinationImage)
	return rv
}/* debug [instance_methods/method]: EncodingStorageSizeForSourceImageSourceStatesDestinationImage */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNKernel */

// The edge mode to use when texture reads stray off the edge of an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/1648826-edgemode
func (c_ CNNKernel) EdgeMode() ImageEdgeMode get set /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("edgeMode"))
	return rv
}/* debug [instance_properties/getter]: edgeMode */


// The edge mode to use when texture reads stray off the edge of an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/1648826-edgemode
func (c_ CNNKernel) SetEdgeMode(value ImageEdgeMode get set /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEdgeMode:"), value)
}/* debug [instance_properties/setter]: edgeMode */


// The position of the destination image's clip rectangle origin, relative to the source image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/1648835-offset
func (c_ CNNKernel) Offset() Offset get set /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("offset"))
	return rv
}/* debug [instance_properties/getter]: offset */


// The position of the destination image's clip rectangle origin, relative to the source image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/1648835-offset
func (c_ CNNKernel) SetOffset(value Offset get set /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOffset:"), value)
}/* debug [instance_properties/setter]: offset */


// An optional clip rectangle to use when writing data. Only the pixels in the clip rectangle will be overwritten.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/1648911-cliprect
func (c_ CNNKernel) ClipRect() Region get set /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("clipRect"))
	return rv
}/* debug [instance_properties/getter]: clipRect */


// An optional clip rectangle to use when writing data. Only the pixels in the clip rectangle will be overwritten.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/1648911-cliprect
func (c_ CNNKernel) SetClipRect(value Region get set /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setClipRect:"), value)
}/* debug [instance_properties/setter]: clipRect */


// The number of channels in the destination image to skip before writing output data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2097550-destinationfeaturechanneloffset
func (c_ CNNKernel) DestinationFeatureChannelOffset() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("destinationFeatureChannelOffset"))
	return rv
}/* debug [instance_properties/getter]: destinationFeatureChannelOffset */


// The number of channels in the destination image to skip before writing output data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2097550-destinationfeaturechanneloffset
func (c_ CNNKernel) SetDestinationFeatureChannelOffset(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDestinationFeatureChannelOffset:"), value)
}/* debug [instance_properties/setter]: destinationFeatureChannelOffset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2865634-isbackwards
func (c_ CNNKernel) IsBackwards() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("isBackwards"))
	return rv
}/* debug [instance_properties/getter]: isBackwards */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2865634-isbackwards
func (c_ CNNKernel) SetIsBackwards(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsBackwards:"), value)
}/* debug [instance_properties/setter]: isBackwards */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2865637-kernelwidth
func (c_ CNNKernel) KernelWidth() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("kernelWidth"))
	return rv
}/* debug [instance_properties/getter]: kernelWidth */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2865637-kernelwidth
func (c_ CNNKernel) SetKernelWidth(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setKernelWidth:"), value)
}/* debug [instance_properties/setter]: kernelWidth */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2865644-strideinpixelsy
func (c_ CNNKernel) StrideInPixelsY() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("strideInPixelsY"))
	return rv
}/* debug [instance_properties/getter]: strideInPixelsY */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2865644-strideinpixelsy
func (c_ CNNKernel) SetStrideInPixelsY(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStrideInPixelsY:"), value)
}/* debug [instance_properties/setter]: strideInPixelsY */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2865648-kernelheight
func (c_ CNNKernel) KernelHeight() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("kernelHeight"))
	return rv
}/* debug [instance_properties/getter]: kernelHeight */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2865648-kernelheight
func (c_ CNNKernel) SetKernelHeight(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setKernelHeight:"), value)
}/* debug [instance_properties/setter]: kernelHeight */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2865650-destinationimageallocator
func (c_ CNNKernel) DestinationImageAllocator() ImageAllocator get set /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("destinationImageAllocator"))
	return rv
}/* debug [instance_properties/getter]: destinationImageAllocator */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2865650-destinationimageallocator
func (c_ CNNKernel) SetDestinationImageAllocator(value ImageAllocator get set /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDestinationImageAllocator:"), value)
}/* debug [instance_properties/setter]: destinationImageAllocator */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2865654-strideinpixelsx
func (c_ CNNKernel) StrideInPixelsX() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("strideInPixelsX"))
	return rv
}/* debug [instance_properties/getter]: strideInPixelsX */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2865654-strideinpixelsx
func (c_ CNNKernel) SetStrideInPixelsX(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStrideInPixelsX:"), value)
}/* debug [instance_properties/setter]: strideInPixelsX */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2865657-padding
func (c_ CNNKernel) Padding() Padding get set /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("padding"))
	return rv
}/* debug [instance_properties/getter]: padding */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2865657-padding
func (c_ CNNKernel) SetPadding(value Padding get set /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPadding:"), value)
}/* debug [instance_properties/setter]: padding */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2942669-dilationratex
func (c_ CNNKernel) DilationRateX() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("dilationRateX"))
	return rv
}/* debug [instance_properties/getter]: dilationRateX */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2942669-dilationratex
func (c_ CNNKernel) SetDilationRateX(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDilationRateX:"), value)
}/* debug [instance_properties/setter]: dilationRateX */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2942673-isstatemodified
func (c_ CNNKernel) IsStateModified() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("isStateModified"))
	return rv
}/* debug [instance_properties/getter]: isStateModified */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2942673-isstatemodified
func (c_ CNNKernel) SetIsStateModified(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsStateModified:"), value)
}/* debug [instance_properties/setter]: isStateModified */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2942679-dilationratey
func (c_ CNNKernel) DilationRateY() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("dilationRateY"))
	return rv
}/* debug [instance_properties/getter]: dilationRateY */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2942679-dilationratey
func (c_ CNNKernel) SetDilationRateY(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDilationRateY:"), value)
}/* debug [instance_properties/setter]: dilationRateY */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2942682-sourcefeaturechanneloffset
func (c_ CNNKernel) SourceFeatureChannelOffset() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("sourceFeatureChannelOffset"))
	return rv
}/* debug [instance_properties/getter]: sourceFeatureChannelOffset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2942682-sourcefeaturechanneloffset
func (c_ CNNKernel) SetSourceFeatureChannelOffset(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSourceFeatureChannelOffset:"), value)
}/* debug [instance_properties/setter]: sourceFeatureChannelOffset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2951917-sourcefeaturechannelmaxcount
func (c_ CNNKernel) SourceFeatureChannelMaxCount() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("sourceFeatureChannelMaxCount"))
	return rv
}/* debug [instance_properties/getter]: sourceFeatureChannelMaxCount */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2951917-sourcefeaturechannelmaxcount
func (c_ CNNKernel) SetSourceFeatureChannelMaxCount(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSourceFeatureChannelMaxCount:"), value)
}/* debug [instance_properties/setter]: sourceFeatureChannelMaxCount */


// The coordinates of the front upper-left corner of the region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRegion/origin
func (c_ CNNKernel) Origin() MTLOrigin /* not a class type */ {
	rv := objc.Send[Origin](c_.ID, objc.Sel("origin"))
	return rv
}/* debug [instance_properties/getter]: origin */


// The coordinates of the front upper-left corner of the region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRegion/origin
func (c_ CNNKernel) SetOrigin(value MTLOrigin /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOrigin:"), value)
}/* debug [instance_properties/setter]: origin */


// The size of the region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsregion/size
func (c_ CNNKernel) Size() objc.IObject /* cross-framework: MPSSize */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("size"))
	return rv
}/* debug [instance_properties/getter]: size */


// The size of the region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsregion/size
func (c_ CNNKernel) SetSize(value objc.IObject /* cross-framework: MPSSize */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSize:"), value)
}/* debug [instance_properties/setter]: size */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNKernel */


