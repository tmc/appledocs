// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [CNNBinaryKernel] class.
type ICNNBinaryKernel interface {
	IKernel
	

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


}





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






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865640-initwithcoder
func NewCNNBinaryKernelWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) CNNBinaryKernel {
	instance := getCNNBinaryKernelClass().Alloc()
	rv := objc.Send[CNNBinaryKernel](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865629-initwithdevice
func NewCNNBinaryKernelWithDevice(device unsafe.Pointer) CNNBinaryKernel {
	instance := getCNNBinaryKernelClass().Alloc()
	rv := objc.Send[CNNBinaryKernel](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865632-encode
func (c_ CNNBinaryKernel) Encode() {
	objc.Send[objc.ID](c_.ID, objc.Sel("encode"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865632-encodetocommandbuffer
func (c_ CNNBinaryKernel) EncodeToCommandBufferPrimaryImageSecondaryImage(commandBuffer unsafe.Pointer, primaryImage IImage, secondaryImage IImage) IImage {
	rv := objc.Send[Image](c_.ID, objc.Sel("encodeToCommandBuffer:primaryImage:secondaryImage:"), commandBuffer, primaryImage, secondaryImage)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865636-encodetocommandbuffer
func (c_ CNNBinaryKernel) EncodeToCommandBufferPrimaryImageSecondaryImageDestinationImage(commandBuffer unsafe.Pointer, primaryImage IImage, secondaryImage IImage, destinationImage IImage) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeToCommandBuffer:primaryImage:secondaryImage:destinationImage:"), commandBuffer, primaryImage, secondaryImage, destinationImage)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2942642-appendbatchbarrier
func (c_ CNNBinaryKernel) AppendBatchBarrier() {
	objc.Send[objc.ID](c_.ID, objc.Sel("appendBatchBarrier"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2942650-encodebatch
func (c_ CNNBinaryKernel) EncodeBatch() {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeBatch"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2942650-encodebatchtocommandbuffer
func (c_ CNNBinaryKernel) EncodeBatchToCommandBufferPrimaryImagesSecondaryImages(commandBuffer unsafe.Pointer, primaryImage ImageBatch /* not a class type */, secondaryImage ImageBatch /* not a class type */) ImageBatch /* not a class type */ {
	rv := objc.Send[ImageBatch](c_.ID, objc.Sel("encodeBatchToCommandBuffer:primaryImages:secondaryImages:"), commandBuffer, primaryImage, secondaryImage)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2942659-isresultstatereusedacrossbatch
func (c_ CNNBinaryKernel) IsResultStateReusedAcrossBatch() {
	objc.Send[objc.ID](c_.ID, objc.Sel("isResultStateReusedAcrossBatch"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2942670-encodebatchtocommandbuffer
func (c_ CNNBinaryKernel) EncodeBatchToCommandBufferPrimaryImagesSecondaryImagesDestinationImages(commandBuffer unsafe.Pointer, primaryImages ImageBatch /* not a class type */, secondaryImages ImageBatch /* not a class type */, destinationImages ImageBatch /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeBatchToCommandBuffer:primaryImages:secondaryImages:destinationImages:"), commandBuffer, primaryImages, secondaryImages, destinationImages)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2942686-destinationimagedescriptor
func (c_ CNNBinaryKernel) DestinationImageDescriptor() {
	objc.Send[objc.ID](c_.ID, objc.Sel("destinationImageDescriptor"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2942686-destinationimagedescriptorforsou
func (c_ CNNBinaryKernel) DestinationImageDescriptorForSourceImagesSourceStates(sourceImages unsafe.Pointer, sourceStates unsafe.Pointer) IImageDescriptor {
	rv := objc.Send[ImageDescriptor](c_.ID, objc.Sel("destinationImageDescriptorForSourceImages:sourceStates:"), sourceImages, sourceStates)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2947930-resultstatebatch
func (c_ CNNBinaryKernel) ResultStateBatch() {
	objc.Send[objc.ID](c_.ID, objc.Sel("resultStateBatch"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2947930-resultstatebatchforprimaryimage
func (c_ CNNBinaryKernel) ResultStateBatchForPrimaryImageSecondaryImageSourceStatesDestinationImage(primaryImage ImageBatch /* not a class type */, secondaryImage ImageBatch /* not a class type */, sourceStates StateBatch /* not a class type */, destinationImage ImageBatch /* not a class type */) StateBatch /* not a class type */ {
	rv := objc.Send[StateBatch](c_.ID, objc.Sel("resultStateBatchForPrimaryImage:secondaryImage:sourceStates:destinationImage:"), primaryImage, secondaryImage, sourceStates, destinationImage)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2947933-temporaryresultstatebatch
func (c_ CNNBinaryKernel) TemporaryResultStateBatch() {
	objc.Send[objc.ID](c_.ID, objc.Sel("temporaryResultStateBatch"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2947933-temporaryresultstatebatchforcomm
func (c_ CNNBinaryKernel) TemporaryResultStateBatchForCommandBufferPrimaryImageSecondaryImageSourceStatesDestinationImage(commandBuffer unsafe.Pointer, primaryImage ImageBatch /* not a class type */, secondaryImage ImageBatch /* not a class type */, sourceStates StateBatch /* not a class type */, destinationImage ImageBatch /* not a class type */) StateBatch /* not a class type */ {
	rv := objc.Send[StateBatch](c_.ID, objc.Sel("temporaryResultStateBatchForCommandBuffer:primaryImage:secondaryImage:sourceStates:destinationImage:"), commandBuffer, primaryImage, secondaryImage, sourceStates, destinationImage)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2947934-encodebatchtocommandbuffer
func (c_ CNNBinaryKernel) EncodeBatchToCommandBufferPrimaryImagesSecondaryImagesDestinationStatesDestinationStateIsTemporary(commandBuffer unsafe.Pointer, primaryImages ImageBatch /* not a class type */, secondaryImages ImageBatch /* not a class type */, outState StateBatch /* not a class type */, isTemporary bool) ImageBatch /* not a class type */ {
	rv := objc.Send[ImageBatch](c_.ID, objc.Sel("encodeBatchToCommandBuffer:primaryImages:secondaryImages:destinationStates:destinationStateIsTemporary:"), commandBuffer, primaryImages, secondaryImages, outState, isTemporary)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2947935-resultstate
func (c_ CNNBinaryKernel) ResultState() {
	objc.Send[objc.ID](c_.ID, objc.Sel("resultState"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2947935-resultstateforprimaryimage
func (c_ CNNBinaryKernel) ResultStateForPrimaryImageSecondaryImageSourceStatesDestinationImage(primaryImage IImage, secondaryImage IImage, sourceStates unsafe.Pointer, destinationImage IImage) IState {
	rv := objc.Send[State](c_.ID, objc.Sel("resultStateForPrimaryImage:secondaryImage:sourceStates:destinationImage:"), primaryImage, secondaryImage, sourceStates, destinationImage)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2947936-encodetocommandbuffer
func (c_ CNNBinaryKernel) EncodeToCommandBufferPrimaryImageSecondaryImageDestinationStateDestinationStateIsTemporary(commandBuffer unsafe.Pointer, primaryImage IImage, secondaryImage IImage, outState objectivec.IObject, isTemporary bool) IImage {
	rv := objc.Send[Image](c_.ID, objc.Sel("encodeToCommandBuffer:primaryImage:secondaryImage:destinationState:destinationStateIsTemporary:"), commandBuffer, primaryImage, secondaryImage, outState, isTemporary)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2947938-temporaryresultstate
func (c_ CNNBinaryKernel) TemporaryResultState() {
	objc.Send[objc.ID](c_.ID, objc.Sel("temporaryResultState"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2947938-temporaryresultstateforcommandbu
func (c_ CNNBinaryKernel) TemporaryResultStateForCommandBufferPrimaryImageSecondaryImageSourceStatesDestinationImage(commandBuffer unsafe.Pointer, primaryImage IImage, secondaryImage IImage, sourceStates unsafe.Pointer, destinationImage IImage) IState {
	rv := objc.Send[State](c_.ID, objc.Sel("temporaryResultStateForCommandBuffer:primaryImage:secondaryImage:sourceStates:destinationImage:"), commandBuffer, primaryImage, secondaryImage, sourceStates, destinationImage)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/3237261-batchencodingstoragesize
func (c_ CNNBinaryKernel) BatchEncodingStorageSize() {
	objc.Send[objc.ID](c_.ID, objc.Sel("batchEncodingStorageSize"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/3237261-batchencodingstoragesizeforprima
func (c_ CNNBinaryKernel) BatchEncodingStorageSizeForPrimaryImageSecondaryImageSourceStatesDestinationImage(primaryImage ImageBatch /* not a class type */, secondaryImage ImageBatch /* not a class type */, sourceStates StateBatch /* not a class type */, destinationImage ImageBatch /* not a class type */) uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("batchEncodingStorageSizeForPrimaryImage:secondaryImage:sourceStates:destinationImage:"), primaryImage, secondaryImage, sourceStates, destinationImage)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/3237262-encodingstoragesize
func (c_ CNNBinaryKernel) EncodingStorageSize() {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodingStorageSize"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/3237262-encodingstoragesizeforprimaryima
func (c_ CNNBinaryKernel) EncodingStorageSizeForPrimaryImageSecondaryImageSourceStatesDestinationImage(primaryImage IImage, secondaryImage IImage, sourceStates unsafe.Pointer, destinationImage IImage) uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("encodingStorageSizeForPrimaryImage:secondaryImage:sourceStates:destinationImage:"), primaryImage, secondaryImage, sourceStates, destinationImage)
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865628-secondaryoffset
func (c_ CNNBinaryKernel) SecondaryOffset() Offset get set /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("secondaryOffset"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865628-secondaryoffset
func (c_ CNNBinaryKernel) SetSecondaryOffset(value Offset get set /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSecondaryOffset:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865630-padding
func (c_ CNNBinaryKernel) Padding() Padding get set /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("padding"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865630-padding
func (c_ CNNBinaryKernel) SetPadding(value Padding get set /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPadding:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865631-secondaryedgemode
func (c_ CNNBinaryKernel) SecondaryEdgeMode() ImageEdgeMode get set /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("secondaryEdgeMode"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865631-secondaryedgemode
func (c_ CNNBinaryKernel) SetSecondaryEdgeMode(value ImageEdgeMode get set /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSecondaryEdgeMode:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865639-secondarystrideinpixelsy
func (c_ CNNBinaryKernel) SecondaryStrideInPixelsY() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("secondaryStrideInPixelsY"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865639-secondarystrideinpixelsy
func (c_ CNNBinaryKernel) SetSecondaryStrideInPixelsY(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSecondaryStrideInPixelsY:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865641-cliprect
func (c_ CNNBinaryKernel) ClipRect() Region get set /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("clipRect"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865641-cliprect
func (c_ CNNBinaryKernel) SetClipRect(value Region get set /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setClipRect:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865643-destinationfeaturechanneloffset
func (c_ CNNBinaryKernel) DestinationFeatureChannelOffset() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("destinationFeatureChannelOffset"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865643-destinationfeaturechanneloffset
func (c_ CNNBinaryKernel) SetDestinationFeatureChannelOffset(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDestinationFeatureChannelOffset:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865645-primaryoffset
func (c_ CNNBinaryKernel) PrimaryOffset() Offset get set /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("primaryOffset"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865645-primaryoffset
func (c_ CNNBinaryKernel) SetPrimaryOffset(value Offset get set /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPrimaryOffset:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865646-primaryedgemode
func (c_ CNNBinaryKernel) PrimaryEdgeMode() ImageEdgeMode get set /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("primaryEdgeMode"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865646-primaryedgemode
func (c_ CNNBinaryKernel) SetPrimaryEdgeMode(value ImageEdgeMode get set /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPrimaryEdgeMode:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865649-secondarystrideinpixelsx
func (c_ CNNBinaryKernel) SecondaryStrideInPixelsX() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("secondaryStrideInPixelsX"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865649-secondarystrideinpixelsx
func (c_ CNNBinaryKernel) SetSecondaryStrideInPixelsX(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSecondaryStrideInPixelsX:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865651-destinationimageallocator
func (c_ CNNBinaryKernel) DestinationImageAllocator() ImageAllocator get set /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("destinationImageAllocator"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865651-destinationimageallocator
func (c_ CNNBinaryKernel) SetDestinationImageAllocator(value ImageAllocator get set /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDestinationImageAllocator:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865652-isbackwards
func (c_ CNNBinaryKernel) IsBackwards() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("isBackwards"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865652-isbackwards
func (c_ CNNBinaryKernel) SetIsBackwards(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsBackwards:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865656-primarystrideinpixelsy
func (c_ CNNBinaryKernel) PrimaryStrideInPixelsY() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("primaryStrideInPixelsY"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865656-primarystrideinpixelsy
func (c_ CNNBinaryKernel) SetPrimaryStrideInPixelsY(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPrimaryStrideInPixelsY:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865658-primarystrideinpixelsx
func (c_ CNNBinaryKernel) PrimaryStrideInPixelsX() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("primaryStrideInPixelsX"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865658-primarystrideinpixelsx
func (c_ CNNBinaryKernel) SetPrimaryStrideInPixelsX(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPrimaryStrideInPixelsX:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2942645-secondarydilationratex
func (c_ CNNBinaryKernel) SecondaryDilationRateX() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("secondaryDilationRateX"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2942645-secondarydilationratex
func (c_ CNNBinaryKernel) SetSecondaryDilationRateX(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSecondaryDilationRateX:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2942648-primarykernelheight
func (c_ CNNBinaryKernel) PrimaryKernelHeight() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("primaryKernelHeight"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2942648-primarykernelheight
func (c_ CNNBinaryKernel) SetPrimaryKernelHeight(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPrimaryKernelHeight:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2942654-secondarysourcefeaturechanneloff
func (c_ CNNBinaryKernel) SecondarySourceFeatureChannelOffset() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("secondarySourceFeatureChannelOffset"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2942654-secondarysourcefeaturechanneloff
func (c_ CNNBinaryKernel) SetSecondarySourceFeatureChannelOffset(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSecondarySourceFeatureChannelOffset:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2942656-primarysourcefeaturechanneloffse
func (c_ CNNBinaryKernel) PrimarySourceFeatureChannelOffset() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("primarySourceFeatureChannelOffset"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2942656-primarysourcefeaturechanneloffse
func (c_ CNNBinaryKernel) SetPrimarySourceFeatureChannelOffset(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPrimarySourceFeatureChannelOffset:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2942658-secondarykernelwidth
func (c_ CNNBinaryKernel) SecondaryKernelWidth() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("secondaryKernelWidth"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2942658-secondarykernelwidth
func (c_ CNNBinaryKernel) SetSecondaryKernelWidth(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSecondaryKernelWidth:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2942660-isstatemodified
func (c_ CNNBinaryKernel) IsStateModified() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("isStateModified"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2942660-isstatemodified
func (c_ CNNBinaryKernel) SetIsStateModified(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsStateModified:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2942662-primarydilationratey
func (c_ CNNBinaryKernel) PrimaryDilationRateY() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("primaryDilationRateY"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2942662-primarydilationratey
func (c_ CNNBinaryKernel) SetPrimaryDilationRateY(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPrimaryDilationRateY:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2942664-secondarykernelheight
func (c_ CNNBinaryKernel) SecondaryKernelHeight() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("secondaryKernelHeight"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2942664-secondarykernelheight
func (c_ CNNBinaryKernel) SetSecondaryKernelHeight(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSecondaryKernelHeight:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2942666-primarykernelwidth
func (c_ CNNBinaryKernel) PrimaryKernelWidth() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("primaryKernelWidth"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2942666-primarykernelwidth
func (c_ CNNBinaryKernel) SetPrimaryKernelWidth(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPrimaryKernelWidth:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2942667-secondarydilationratey
func (c_ CNNBinaryKernel) SecondaryDilationRateY() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("secondaryDilationRateY"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2942667-secondarydilationratey
func (c_ CNNBinaryKernel) SetSecondaryDilationRateY(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSecondaryDilationRateY:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2942677-primarydilationratex
func (c_ CNNBinaryKernel) PrimaryDilationRateX() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("primaryDilationRateX"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2942677-primarydilationratex
func (c_ CNNBinaryKernel) SetPrimaryDilationRateX(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPrimaryDilationRateX:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2951918-secondarysourcefeaturechannelmax
func (c_ CNNBinaryKernel) SecondarySourceFeatureChannelMaxCount() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("secondarySourceFeatureChannelMaxCount"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2951918-secondarysourcefeaturechannelmax
func (c_ CNNBinaryKernel) SetSecondarySourceFeatureChannelMaxCount(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSecondarySourceFeatureChannelMaxCount:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2951919-primarysourcefeaturechannelmaxco
func (c_ CNNBinaryKernel) PrimarySourceFeatureChannelMaxCount() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("primarySourceFeatureChannelMaxCount"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2951919-primarysourcefeaturechannelmaxco
func (c_ CNNBinaryKernel) SetPrimarySourceFeatureChannelMaxCount(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPrimarySourceFeatureChannelMaxCount:"), value)
}







