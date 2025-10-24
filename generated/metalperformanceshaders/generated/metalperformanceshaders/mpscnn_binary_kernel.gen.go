// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNBinaryKernel */


/* debug [class_header]: Header for MPSCNNBinaryKernel */
// The class instance for the [CNNBinaryKernel] class.
var (
	CNNBinaryKernelClass     _CNNBinaryKernelClass
	CNNBinaryKernelClassOnce sync.Once
)

func getCNNBinaryKernelClass() _CNNBinaryKernelClass {
	CNNBinaryKernelClassOnce.Do(func() {
		CNNBinaryKernelClass = _CNNBinaryKernelClass{objc.GetClass("MPSCNNBinaryKernel")}
	})
	return CNNBinaryKernelClass
}

type _CNNBinaryKernelClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNBinaryKernel */
// An interface definition for the [CNNBinaryKernel] class.
type ICNNBinaryKernel interface {
	IKernel
	
/* debug [class_interface_properties]: Properties for CNNBinaryKernel */
	// properties:
	SecondaryOffset() Offset get set /* not a class type */
	SetSecondaryOffset(value Offset get set /* not a class type */)
	Padding() Padding get set /* not a class type */
	SetPadding(value Padding get set /* not a class type */)
	SecondaryEdgeMode() ImageEdgeMode get set /* not a class type */
	SetSecondaryEdgeMode(value ImageEdgeMode get set /* not a class type */)
	SecondaryStrideInPixelsY() objectivec.IObject
	SetSecondaryStrideInPixelsY(value objectivec.IObject)
	ClipRect() Region get set /* not a class type */
	SetClipRect(value Region get set /* not a class type */)
	DestinationFeatureChannelOffset() objectivec.IObject
	SetDestinationFeatureChannelOffset(value objectivec.IObject)
	PrimaryOffset() Offset get set /* not a class type */
	SetPrimaryOffset(value Offset get set /* not a class type */)
	PrimaryEdgeMode() ImageEdgeMode get set /* not a class type */
	SetPrimaryEdgeMode(value ImageEdgeMode get set /* not a class type */)
	SecondaryStrideInPixelsX() objectivec.IObject
	SetSecondaryStrideInPixelsX(value objectivec.IObject)
	DestinationImageAllocator() ImageAllocator get set /* not a class type */
	SetDestinationImageAllocator(value ImageAllocator get set /* not a class type */)
	IsBackwards() objectivec.IObject
	SetIsBackwards(value objectivec.IObject)
	PrimaryStrideInPixelsY() objectivec.IObject
	SetPrimaryStrideInPixelsY(value objectivec.IObject)
	PrimaryStrideInPixelsX() objectivec.IObject
	SetPrimaryStrideInPixelsX(value objectivec.IObject)
	SecondaryDilationRateX() objectivec.IObject
	SetSecondaryDilationRateX(value objectivec.IObject)
	PrimaryKernelHeight() objectivec.IObject
	SetPrimaryKernelHeight(value objectivec.IObject)
	SecondarySourceFeatureChannelOffset() objectivec.IObject
	SetSecondarySourceFeatureChannelOffset(value objectivec.IObject)
	PrimarySourceFeatureChannelOffset() objectivec.IObject
	SetPrimarySourceFeatureChannelOffset(value objectivec.IObject)
	SecondaryKernelWidth() objectivec.IObject
	SetSecondaryKernelWidth(value objectivec.IObject)
	IsStateModified() objectivec.IObject
	SetIsStateModified(value objectivec.IObject)
	PrimaryDilationRateY() objectivec.IObject
	SetPrimaryDilationRateY(value objectivec.IObject)
	SecondaryKernelHeight() objectivec.IObject
	SetSecondaryKernelHeight(value objectivec.IObject)
	PrimaryKernelWidth() objectivec.IObject
	SetPrimaryKernelWidth(value objectivec.IObject)
	SecondaryDilationRateY() objectivec.IObject
	SetSecondaryDilationRateY(value objectivec.IObject)
	PrimaryDilationRateX() objectivec.IObject
	SetPrimaryDilationRateX(value objectivec.IObject)
	SecondarySourceFeatureChannelMaxCount() objectivec.IObject
	SetSecondarySourceFeatureChannelMaxCount(value objectivec.IObject)
	PrimarySourceFeatureChannelMaxCount() objectivec.IObject
	SetPrimarySourceFeatureChannelMaxCount(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNBinaryKernel */
	// methods:
	Encode()
	EncodeToCommandBufferPrimaryImageSecondaryImage(commandBuffer unsafe.Pointer, primaryImage IImage, secondaryImage IImage) IImage
	EncodeToCommandBufferPrimaryImageSecondaryImageDestinationImage(commandBuffer unsafe.Pointer, primaryImage IImage, secondaryImage IImage, destinationImage IImage)
	AppendBatchBarrier()
	EncodeBatch()
	EncodeBatchToCommandBufferPrimaryImagesSecondaryImages(commandBuffer unsafe.Pointer, primaryImage ImageBatch /* not a class type */, secondaryImage ImageBatch /* not a class type */) ImageBatch /* not a class type */
	IsResultStateReusedAcrossBatch()
	EncodeBatchToCommandBufferPrimaryImagesSecondaryImagesDestinationImages(commandBuffer unsafe.Pointer, primaryImages ImageBatch /* not a class type */, secondaryImages ImageBatch /* not a class type */, destinationImages ImageBatch /* not a class type */)
	DestinationImageDescriptor()
	DestinationImageDescriptorForSourceImagesSourceStates(sourceImages unsafe.Pointer, sourceStates unsafe.Pointer) IImageDescriptor
	ResultStateBatch()
	ResultStateBatchForPrimaryImageSecondaryImageSourceStatesDestinationImage(primaryImage ImageBatch /* not a class type */, secondaryImage ImageBatch /* not a class type */, sourceStates StateBatch /* not a class type */, destinationImage ImageBatch /* not a class type */) StateBatch /* not a class type */
	TemporaryResultStateBatch()
	TemporaryResultStateBatchForCommandBufferPrimaryImageSecondaryImageSourceStatesDestinationImage(commandBuffer unsafe.Pointer, primaryImage ImageBatch /* not a class type */, secondaryImage ImageBatch /* not a class type */, sourceStates StateBatch /* not a class type */, destinationImage ImageBatch /* not a class type */) StateBatch /* not a class type */
	EncodeBatchToCommandBufferPrimaryImagesSecondaryImagesDestinationStatesDestinationStateIsTemporary(commandBuffer unsafe.Pointer, primaryImages ImageBatch /* not a class type */, secondaryImages ImageBatch /* not a class type */, outState StateBatch /* not a class type */, isTemporary bool) ImageBatch /* not a class type */
	ResultState()
	ResultStateForPrimaryImageSecondaryImageSourceStatesDestinationImage(primaryImage IImage, secondaryImage IImage, sourceStates unsafe.Pointer, destinationImage IImage) IState
	EncodeToCommandBufferPrimaryImageSecondaryImageDestinationStateDestinationStateIsTemporary(commandBuffer unsafe.Pointer, primaryImage IImage, secondaryImage IImage, outState objectivec.IObject, isTemporary bool) IImage
	TemporaryResultState()
	TemporaryResultStateForCommandBufferPrimaryImageSecondaryImageSourceStatesDestinationImage(commandBuffer unsafe.Pointer, primaryImage IImage, secondaryImage IImage, sourceStates unsafe.Pointer, destinationImage IImage) IState
	BatchEncodingStorageSize()
	BatchEncodingStorageSizeForPrimaryImageSecondaryImageSourceStatesDestinationImage(primaryImage ImageBatch /* not a class type */, secondaryImage ImageBatch /* not a class type */, sourceStates StateBatch /* not a class type */, destinationImage ImageBatch /* not a class type */) uint
	EncodingStorageSize()
	EncodingStorageSizeForPrimaryImageSecondaryImageSourceStatesDestinationImage(primaryImage IImage, secondaryImage IImage, sourceStates unsafe.Pointer, destinationImage IImage) uint
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNBinaryKernel */
// Alloc allocates a new instance without initialization.
func (cc _CNNBinaryKernelClass) Alloc() CNNBinaryKernel {
	rv := objc.Send[CNNBinaryKernel](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNBinaryKernelClass) New() CNNBinaryKernel {
	rv := objc.Send[CNNBinaryKernel](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNBinaryKernel) Init() CNNBinaryKernel {
	rv := objc.Send[CNNBinaryKernel](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNBinaryKernel) Autorelease() CNNBinaryKernel {
	rv := objc.Send[CNNBinaryKernel](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNBinaryKernel creates a new CNNBinaryKernel instance.
func NewCNNBinaryKernel() CNNBinaryKernel {
	return getCNNBinaryKernelClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNBinaryKernel */
// A convolution neural network kernel.


// A convolution neural network kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryKernel
type CNNBinaryKernel struct {
	Kernel
}

// CNNBinaryKernelFrom constructs a [CNNBinaryKernel] from an unsafe.Pointer.
//
// A convolution neural network kernel.
func CNNBinaryKernelFrom(ptr unsafe.Pointer) CNNBinaryKernel {
	return CNNBinaryKernel{
		Kernel: KernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNBinaryKernel */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865640-initwithcoder
func NewCNNBinaryKernelWithCoderDevice(aDecoder Coder /* not a class type */, device unsafe.Pointer) CNNBinaryKernel {
	instance := getCNNBinaryKernelClass().Alloc()
	rv := objc.Send[CNNBinaryKernel](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNBinaryKernelWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865629-initwithdevice
func NewCNNBinaryKernelWithDevice(device unsafe.Pointer) CNNBinaryKernel {
	instance := getCNNBinaryKernelClass().Alloc()
	rv := objc.Send[CNNBinaryKernel](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNBinaryKernelWithDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNBinaryKernel */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNBinaryKernel */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNBinaryKernel */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865632-encode
func (c_ CNNBinaryKernel) Encode() {
	objc.Send[objc.ID](c_.ID, objc.Sel("encode"))
}/* debug [instance_methods/method]: Encode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865632-encodetocommandbuffer
func (c_ CNNBinaryKernel) EncodeToCommandBufferPrimaryImageSecondaryImage(commandBuffer unsafe.Pointer, primaryImage IImage, secondaryImage IImage) IImage {
	rv := objc.Send[Image](c_.ID, objc.Sel("encodeToCommandBuffer:primaryImage:secondaryImage:"), commandBuffer, primaryImage, secondaryImage)
	return rv
}/* debug [instance_methods/method]: EncodeToCommandBufferPrimaryImageSecondaryImage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865636-encodetocommandbuffer
func (c_ CNNBinaryKernel) EncodeToCommandBufferPrimaryImageSecondaryImageDestinationImage(commandBuffer unsafe.Pointer, primaryImage IImage, secondaryImage IImage, destinationImage IImage) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeToCommandBuffer:primaryImage:secondaryImage:destinationImage:"), commandBuffer, primaryImage, secondaryImage, destinationImage)
}/* debug [instance_methods/method]: EncodeToCommandBufferPrimaryImageSecondaryImageDestinationImage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2942642-appendbatchbarrier
func (c_ CNNBinaryKernel) AppendBatchBarrier() {
	objc.Send[objc.ID](c_.ID, objc.Sel("appendBatchBarrier"))
}/* debug [instance_methods/method]: AppendBatchBarrier */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2942650-encodebatch
func (c_ CNNBinaryKernel) EncodeBatch() {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeBatch"))
}/* debug [instance_methods/method]: EncodeBatch */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2942650-encodebatchtocommandbuffer
func (c_ CNNBinaryKernel) EncodeBatchToCommandBufferPrimaryImagesSecondaryImages(commandBuffer unsafe.Pointer, primaryImage ImageBatch /* not a class type */, secondaryImage ImageBatch /* not a class type */) ImageBatch /* not a class type */ {
	rv := objc.Send[ImageBatch](c_.ID, objc.Sel("encodeBatchToCommandBuffer:primaryImages:secondaryImages:"), commandBuffer, primaryImage, secondaryImage)
	return rv
}/* debug [instance_methods/method]: EncodeBatchToCommandBufferPrimaryImagesSecondaryImages */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2942659-isresultstatereusedacrossbatch
func (c_ CNNBinaryKernel) IsResultStateReusedAcrossBatch() {
	objc.Send[objc.ID](c_.ID, objc.Sel("isResultStateReusedAcrossBatch"))
}/* debug [instance_methods/method]: IsResultStateReusedAcrossBatch */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2942670-encodebatchtocommandbuffer
func (c_ CNNBinaryKernel) EncodeBatchToCommandBufferPrimaryImagesSecondaryImagesDestinationImages(commandBuffer unsafe.Pointer, primaryImages ImageBatch /* not a class type */, secondaryImages ImageBatch /* not a class type */, destinationImages ImageBatch /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeBatchToCommandBuffer:primaryImages:secondaryImages:destinationImages:"), commandBuffer, primaryImages, secondaryImages, destinationImages)
}/* debug [instance_methods/method]: EncodeBatchToCommandBufferPrimaryImagesSecondaryImagesDestinationImages */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2942686-destinationimagedescriptor
func (c_ CNNBinaryKernel) DestinationImageDescriptor() {
	objc.Send[objc.ID](c_.ID, objc.Sel("destinationImageDescriptor"))
}/* debug [instance_methods/method]: DestinationImageDescriptor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2942686-destinationimagedescriptorforsou
func (c_ CNNBinaryKernel) DestinationImageDescriptorForSourceImagesSourceStates(sourceImages unsafe.Pointer, sourceStates unsafe.Pointer) IImageDescriptor {
	rv := objc.Send[ImageDescriptor](c_.ID, objc.Sel("destinationImageDescriptorForSourceImages:sourceStates:"), sourceImages, sourceStates)
	return rv
}/* debug [instance_methods/method]: DestinationImageDescriptorForSourceImagesSourceStates */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2947930-resultstatebatch
func (c_ CNNBinaryKernel) ResultStateBatch() {
	objc.Send[objc.ID](c_.ID, objc.Sel("resultStateBatch"))
}/* debug [instance_methods/method]: ResultStateBatch */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2947930-resultstatebatchforprimaryimage
func (c_ CNNBinaryKernel) ResultStateBatchForPrimaryImageSecondaryImageSourceStatesDestinationImage(primaryImage ImageBatch /* not a class type */, secondaryImage ImageBatch /* not a class type */, sourceStates StateBatch /* not a class type */, destinationImage ImageBatch /* not a class type */) StateBatch /* not a class type */ {
	rv := objc.Send[StateBatch](c_.ID, objc.Sel("resultStateBatchForPrimaryImage:secondaryImage:sourceStates:destinationImage:"), primaryImage, secondaryImage, sourceStates, destinationImage)
	return rv
}/* debug [instance_methods/method]: ResultStateBatchForPrimaryImageSecondaryImageSourceStatesDestinationImage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2947933-temporaryresultstatebatch
func (c_ CNNBinaryKernel) TemporaryResultStateBatch() {
	objc.Send[objc.ID](c_.ID, objc.Sel("temporaryResultStateBatch"))
}/* debug [instance_methods/method]: TemporaryResultStateBatch */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2947933-temporaryresultstatebatchforcomm
func (c_ CNNBinaryKernel) TemporaryResultStateBatchForCommandBufferPrimaryImageSecondaryImageSourceStatesDestinationImage(commandBuffer unsafe.Pointer, primaryImage ImageBatch /* not a class type */, secondaryImage ImageBatch /* not a class type */, sourceStates StateBatch /* not a class type */, destinationImage ImageBatch /* not a class type */) StateBatch /* not a class type */ {
	rv := objc.Send[StateBatch](c_.ID, objc.Sel("temporaryResultStateBatchForCommandBuffer:primaryImage:secondaryImage:sourceStates:destinationImage:"), commandBuffer, primaryImage, secondaryImage, sourceStates, destinationImage)
	return rv
}/* debug [instance_methods/method]: TemporaryResultStateBatchForCommandBufferPrimaryImageSecondaryImageSourceStatesDestinationImage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2947934-encodebatchtocommandbuffer
func (c_ CNNBinaryKernel) EncodeBatchToCommandBufferPrimaryImagesSecondaryImagesDestinationStatesDestinationStateIsTemporary(commandBuffer unsafe.Pointer, primaryImages ImageBatch /* not a class type */, secondaryImages ImageBatch /* not a class type */, outState StateBatch /* not a class type */, isTemporary bool) ImageBatch /* not a class type */ {
	rv := objc.Send[ImageBatch](c_.ID, objc.Sel("encodeBatchToCommandBuffer:primaryImages:secondaryImages:destinationStates:destinationStateIsTemporary:"), commandBuffer, primaryImages, secondaryImages, outState, isTemporary)
	return rv
}/* debug [instance_methods/method]: EncodeBatchToCommandBufferPrimaryImagesSecondaryImagesDestinationStatesDestinationStateIsTemporary */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2947935-resultstate
func (c_ CNNBinaryKernel) ResultState() {
	objc.Send[objc.ID](c_.ID, objc.Sel("resultState"))
}/* debug [instance_methods/method]: ResultState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2947935-resultstateforprimaryimage
func (c_ CNNBinaryKernel) ResultStateForPrimaryImageSecondaryImageSourceStatesDestinationImage(primaryImage IImage, secondaryImage IImage, sourceStates unsafe.Pointer, destinationImage IImage) IState {
	rv := objc.Send[State](c_.ID, objc.Sel("resultStateForPrimaryImage:secondaryImage:sourceStates:destinationImage:"), primaryImage, secondaryImage, sourceStates, destinationImage)
	return rv
}/* debug [instance_methods/method]: ResultStateForPrimaryImageSecondaryImageSourceStatesDestinationImage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2947936-encodetocommandbuffer
func (c_ CNNBinaryKernel) EncodeToCommandBufferPrimaryImageSecondaryImageDestinationStateDestinationStateIsTemporary(commandBuffer unsafe.Pointer, primaryImage IImage, secondaryImage IImage, outState objectivec.IObject, isTemporary bool) IImage {
	rv := objc.Send[Image](c_.ID, objc.Sel("encodeToCommandBuffer:primaryImage:secondaryImage:destinationState:destinationStateIsTemporary:"), commandBuffer, primaryImage, secondaryImage, outState, isTemporary)
	return rv
}/* debug [instance_methods/method]: EncodeToCommandBufferPrimaryImageSecondaryImageDestinationStateDestinationStateIsTemporary */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2947938-temporaryresultstate
func (c_ CNNBinaryKernel) TemporaryResultState() {
	objc.Send[objc.ID](c_.ID, objc.Sel("temporaryResultState"))
}/* debug [instance_methods/method]: TemporaryResultState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2947938-temporaryresultstateforcommandbu
func (c_ CNNBinaryKernel) TemporaryResultStateForCommandBufferPrimaryImageSecondaryImageSourceStatesDestinationImage(commandBuffer unsafe.Pointer, primaryImage IImage, secondaryImage IImage, sourceStates unsafe.Pointer, destinationImage IImage) IState {
	rv := objc.Send[State](c_.ID, objc.Sel("temporaryResultStateForCommandBuffer:primaryImage:secondaryImage:sourceStates:destinationImage:"), commandBuffer, primaryImage, secondaryImage, sourceStates, destinationImage)
	return rv
}/* debug [instance_methods/method]: TemporaryResultStateForCommandBufferPrimaryImageSecondaryImageSourceStatesDestinationImage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/3237261-batchencodingstoragesize
func (c_ CNNBinaryKernel) BatchEncodingStorageSize() {
	objc.Send[objc.ID](c_.ID, objc.Sel("batchEncodingStorageSize"))
}/* debug [instance_methods/method]: BatchEncodingStorageSize */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/3237261-batchencodingstoragesizeforprima
func (c_ CNNBinaryKernel) BatchEncodingStorageSizeForPrimaryImageSecondaryImageSourceStatesDestinationImage(primaryImage ImageBatch /* not a class type */, secondaryImage ImageBatch /* not a class type */, sourceStates StateBatch /* not a class type */, destinationImage ImageBatch /* not a class type */) uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("batchEncodingStorageSizeForPrimaryImage:secondaryImage:sourceStates:destinationImage:"), primaryImage, secondaryImage, sourceStates, destinationImage)
	return rv
}/* debug [instance_methods/method]: BatchEncodingStorageSizeForPrimaryImageSecondaryImageSourceStatesDestinationImage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/3237262-encodingstoragesize
func (c_ CNNBinaryKernel) EncodingStorageSize() {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodingStorageSize"))
}/* debug [instance_methods/method]: EncodingStorageSize */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/3237262-encodingstoragesizeforprimaryima
func (c_ CNNBinaryKernel) EncodingStorageSizeForPrimaryImageSecondaryImageSourceStatesDestinationImage(primaryImage IImage, secondaryImage IImage, sourceStates unsafe.Pointer, destinationImage IImage) uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("encodingStorageSizeForPrimaryImage:secondaryImage:sourceStates:destinationImage:"), primaryImage, secondaryImage, sourceStates, destinationImage)
	return rv
}/* debug [instance_methods/method]: EncodingStorageSizeForPrimaryImageSecondaryImageSourceStatesDestinationImage */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNBinaryKernel */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865628-secondaryoffset
func (c_ CNNBinaryKernel) SecondaryOffset() Offset get set /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("secondaryOffset"))
	return rv
}/* debug [instance_properties/getter]: secondaryOffset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865628-secondaryoffset
func (c_ CNNBinaryKernel) SetSecondaryOffset(value Offset get set /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSecondaryOffset:"), value)
}/* debug [instance_properties/setter]: secondaryOffset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865630-padding
func (c_ CNNBinaryKernel) Padding() Padding get set /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("padding"))
	return rv
}/* debug [instance_properties/getter]: padding */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865630-padding
func (c_ CNNBinaryKernel) SetPadding(value Padding get set /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPadding:"), value)
}/* debug [instance_properties/setter]: padding */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865631-secondaryedgemode
func (c_ CNNBinaryKernel) SecondaryEdgeMode() ImageEdgeMode get set /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("secondaryEdgeMode"))
	return rv
}/* debug [instance_properties/getter]: secondaryEdgeMode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865631-secondaryedgemode
func (c_ CNNBinaryKernel) SetSecondaryEdgeMode(value ImageEdgeMode get set /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSecondaryEdgeMode:"), value)
}/* debug [instance_properties/setter]: secondaryEdgeMode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865639-secondarystrideinpixelsy
func (c_ CNNBinaryKernel) SecondaryStrideInPixelsY() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("secondaryStrideInPixelsY"))
	return rv
}/* debug [instance_properties/getter]: secondaryStrideInPixelsY */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865639-secondarystrideinpixelsy
func (c_ CNNBinaryKernel) SetSecondaryStrideInPixelsY(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSecondaryStrideInPixelsY:"), value)
}/* debug [instance_properties/setter]: secondaryStrideInPixelsY */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865641-cliprect
func (c_ CNNBinaryKernel) ClipRect() Region get set /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("clipRect"))
	return rv
}/* debug [instance_properties/getter]: clipRect */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865641-cliprect
func (c_ CNNBinaryKernel) SetClipRect(value Region get set /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setClipRect:"), value)
}/* debug [instance_properties/setter]: clipRect */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865643-destinationfeaturechanneloffset
func (c_ CNNBinaryKernel) DestinationFeatureChannelOffset() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("destinationFeatureChannelOffset"))
	return rv
}/* debug [instance_properties/getter]: destinationFeatureChannelOffset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865643-destinationfeaturechanneloffset
func (c_ CNNBinaryKernel) SetDestinationFeatureChannelOffset(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDestinationFeatureChannelOffset:"), value)
}/* debug [instance_properties/setter]: destinationFeatureChannelOffset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865645-primaryoffset
func (c_ CNNBinaryKernel) PrimaryOffset() Offset get set /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("primaryOffset"))
	return rv
}/* debug [instance_properties/getter]: primaryOffset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865645-primaryoffset
func (c_ CNNBinaryKernel) SetPrimaryOffset(value Offset get set /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPrimaryOffset:"), value)
}/* debug [instance_properties/setter]: primaryOffset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865646-primaryedgemode
func (c_ CNNBinaryKernel) PrimaryEdgeMode() ImageEdgeMode get set /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("primaryEdgeMode"))
	return rv
}/* debug [instance_properties/getter]: primaryEdgeMode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865646-primaryedgemode
func (c_ CNNBinaryKernel) SetPrimaryEdgeMode(value ImageEdgeMode get set /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPrimaryEdgeMode:"), value)
}/* debug [instance_properties/setter]: primaryEdgeMode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865649-secondarystrideinpixelsx
func (c_ CNNBinaryKernel) SecondaryStrideInPixelsX() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("secondaryStrideInPixelsX"))
	return rv
}/* debug [instance_properties/getter]: secondaryStrideInPixelsX */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865649-secondarystrideinpixelsx
func (c_ CNNBinaryKernel) SetSecondaryStrideInPixelsX(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSecondaryStrideInPixelsX:"), value)
}/* debug [instance_properties/setter]: secondaryStrideInPixelsX */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865651-destinationimageallocator
func (c_ CNNBinaryKernel) DestinationImageAllocator() ImageAllocator get set /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("destinationImageAllocator"))
	return rv
}/* debug [instance_properties/getter]: destinationImageAllocator */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865651-destinationimageallocator
func (c_ CNNBinaryKernel) SetDestinationImageAllocator(value ImageAllocator get set /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDestinationImageAllocator:"), value)
}/* debug [instance_properties/setter]: destinationImageAllocator */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865652-isbackwards
func (c_ CNNBinaryKernel) IsBackwards() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("isBackwards"))
	return rv
}/* debug [instance_properties/getter]: isBackwards */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865652-isbackwards
func (c_ CNNBinaryKernel) SetIsBackwards(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsBackwards:"), value)
}/* debug [instance_properties/setter]: isBackwards */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865656-primarystrideinpixelsy
func (c_ CNNBinaryKernel) PrimaryStrideInPixelsY() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("primaryStrideInPixelsY"))
	return rv
}/* debug [instance_properties/getter]: primaryStrideInPixelsY */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865656-primarystrideinpixelsy
func (c_ CNNBinaryKernel) SetPrimaryStrideInPixelsY(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPrimaryStrideInPixelsY:"), value)
}/* debug [instance_properties/setter]: primaryStrideInPixelsY */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865658-primarystrideinpixelsx
func (c_ CNNBinaryKernel) PrimaryStrideInPixelsX() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("primaryStrideInPixelsX"))
	return rv
}/* debug [instance_properties/getter]: primaryStrideInPixelsX */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865658-primarystrideinpixelsx
func (c_ CNNBinaryKernel) SetPrimaryStrideInPixelsX(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPrimaryStrideInPixelsX:"), value)
}/* debug [instance_properties/setter]: primaryStrideInPixelsX */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2942645-secondarydilationratex
func (c_ CNNBinaryKernel) SecondaryDilationRateX() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("secondaryDilationRateX"))
	return rv
}/* debug [instance_properties/getter]: secondaryDilationRateX */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2942645-secondarydilationratex
func (c_ CNNBinaryKernel) SetSecondaryDilationRateX(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSecondaryDilationRateX:"), value)
}/* debug [instance_properties/setter]: secondaryDilationRateX */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2942648-primarykernelheight
func (c_ CNNBinaryKernel) PrimaryKernelHeight() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("primaryKernelHeight"))
	return rv
}/* debug [instance_properties/getter]: primaryKernelHeight */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2942648-primarykernelheight
func (c_ CNNBinaryKernel) SetPrimaryKernelHeight(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPrimaryKernelHeight:"), value)
}/* debug [instance_properties/setter]: primaryKernelHeight */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2942654-secondarysourcefeaturechanneloff
func (c_ CNNBinaryKernel) SecondarySourceFeatureChannelOffset() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("secondarySourceFeatureChannelOffset"))
	return rv
}/* debug [instance_properties/getter]: secondarySourceFeatureChannelOffset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2942654-secondarysourcefeaturechanneloff
func (c_ CNNBinaryKernel) SetSecondarySourceFeatureChannelOffset(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSecondarySourceFeatureChannelOffset:"), value)
}/* debug [instance_properties/setter]: secondarySourceFeatureChannelOffset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2942656-primarysourcefeaturechanneloffse
func (c_ CNNBinaryKernel) PrimarySourceFeatureChannelOffset() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("primarySourceFeatureChannelOffset"))
	return rv
}/* debug [instance_properties/getter]: primarySourceFeatureChannelOffset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2942656-primarysourcefeaturechanneloffse
func (c_ CNNBinaryKernel) SetPrimarySourceFeatureChannelOffset(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPrimarySourceFeatureChannelOffset:"), value)
}/* debug [instance_properties/setter]: primarySourceFeatureChannelOffset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2942658-secondarykernelwidth
func (c_ CNNBinaryKernel) SecondaryKernelWidth() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("secondaryKernelWidth"))
	return rv
}/* debug [instance_properties/getter]: secondaryKernelWidth */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2942658-secondarykernelwidth
func (c_ CNNBinaryKernel) SetSecondaryKernelWidth(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSecondaryKernelWidth:"), value)
}/* debug [instance_properties/setter]: secondaryKernelWidth */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2942660-isstatemodified
func (c_ CNNBinaryKernel) IsStateModified() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("isStateModified"))
	return rv
}/* debug [instance_properties/getter]: isStateModified */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2942660-isstatemodified
func (c_ CNNBinaryKernel) SetIsStateModified(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsStateModified:"), value)
}/* debug [instance_properties/setter]: isStateModified */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2942662-primarydilationratey
func (c_ CNNBinaryKernel) PrimaryDilationRateY() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("primaryDilationRateY"))
	return rv
}/* debug [instance_properties/getter]: primaryDilationRateY */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2942662-primarydilationratey
func (c_ CNNBinaryKernel) SetPrimaryDilationRateY(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPrimaryDilationRateY:"), value)
}/* debug [instance_properties/setter]: primaryDilationRateY */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2942664-secondarykernelheight
func (c_ CNNBinaryKernel) SecondaryKernelHeight() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("secondaryKernelHeight"))
	return rv
}/* debug [instance_properties/getter]: secondaryKernelHeight */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2942664-secondarykernelheight
func (c_ CNNBinaryKernel) SetSecondaryKernelHeight(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSecondaryKernelHeight:"), value)
}/* debug [instance_properties/setter]: secondaryKernelHeight */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2942666-primarykernelwidth
func (c_ CNNBinaryKernel) PrimaryKernelWidth() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("primaryKernelWidth"))
	return rv
}/* debug [instance_properties/getter]: primaryKernelWidth */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2942666-primarykernelwidth
func (c_ CNNBinaryKernel) SetPrimaryKernelWidth(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPrimaryKernelWidth:"), value)
}/* debug [instance_properties/setter]: primaryKernelWidth */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2942667-secondarydilationratey
func (c_ CNNBinaryKernel) SecondaryDilationRateY() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("secondaryDilationRateY"))
	return rv
}/* debug [instance_properties/getter]: secondaryDilationRateY */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2942667-secondarydilationratey
func (c_ CNNBinaryKernel) SetSecondaryDilationRateY(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSecondaryDilationRateY:"), value)
}/* debug [instance_properties/setter]: secondaryDilationRateY */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2942677-primarydilationratex
func (c_ CNNBinaryKernel) PrimaryDilationRateX() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("primaryDilationRateX"))
	return rv
}/* debug [instance_properties/getter]: primaryDilationRateX */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2942677-primarydilationratex
func (c_ CNNBinaryKernel) SetPrimaryDilationRateX(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPrimaryDilationRateX:"), value)
}/* debug [instance_properties/setter]: primaryDilationRateX */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2951918-secondarysourcefeaturechannelmax
func (c_ CNNBinaryKernel) SecondarySourceFeatureChannelMaxCount() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("secondarySourceFeatureChannelMaxCount"))
	return rv
}/* debug [instance_properties/getter]: secondarySourceFeatureChannelMaxCount */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2951918-secondarysourcefeaturechannelmax
func (c_ CNNBinaryKernel) SetSecondarySourceFeatureChannelMaxCount(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSecondarySourceFeatureChannelMaxCount:"), value)
}/* debug [instance_properties/setter]: secondarySourceFeatureChannelMaxCount */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2951919-primarysourcefeaturechannelmaxco
func (c_ CNNBinaryKernel) PrimarySourceFeatureChannelMaxCount() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("primarySourceFeatureChannelMaxCount"))
	return rv
}/* debug [instance_properties/getter]: primarySourceFeatureChannelMaxCount */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2951919-primarysourcefeaturechannelmaxco
func (c_ CNNBinaryKernel) SetPrimarySourceFeatureChannelMaxCount(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPrimarySourceFeatureChannelMaxCount:"), value)
}/* debug [instance_properties/setter]: primarySourceFeatureChannelMaxCount */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNBinaryKernel */


