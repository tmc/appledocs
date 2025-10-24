// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CNNMultiaryKernel] class.
var (
	CNNMultiaryKernelClass     _CNNMultiaryKernelClass
	CNNMultiaryKernelClassOnce sync.Once
)

func getCNNMultiaryKernelClass() _CNNMultiaryKernelClass {
	CNNMultiaryKernelClassOnce.Do(func() {
		CNNMultiaryKernelClass = _CNNMultiaryKernelClass{objc.GetClass("MPSCNNMultiaryKernel")}
	})
	return CNNMultiaryKernelClass
}

type _CNNMultiaryKernelClass struct {
	class objc.Class
}





// An interface definition for the [CNNMultiaryKernel] class.
type ICNNMultiaryKernel interface {
	IKernel
	

	// properties:
	ClipRect() Region get set /* not a class type */
	SetClipRect(value Region get set /* not a class type */)
	DestinationFeatureChannelOffset() objectivec.IObject
	SetDestinationFeatureChannelOffset(value objectivec.IObject)
	DestinationImageAllocator() ImageAllocator get set /* not a class type */
	SetDestinationImageAllocator(value ImageAllocator get set /* not a class type */)
	IsBackwards() objectivec.IObject
	SetIsBackwards(value objectivec.IObject)
	IsStateModified() objectivec.IObject
	SetIsStateModified(value objectivec.IObject)
	Padding() Padding get set /* not a class type */
	SetPadding(value Padding get set /* not a class type */)
	SourceCount() objectivec.IObject
	SetSourceCount(value objectivec.IObject)


	

	// methods:
	AppendBatchBarrier()
	DestinationImageDescriptor()
	DestinationImageDescriptorForSourceImagesSourceStates(sourceImages unsafe.Pointer, sourceStates unsafe.Pointer) IImageDescriptor
	DilationRateXatIndex()
	DilationRateYatIndex()
	EdgeMode()
	EdgeModeAtIndex(index uint) ImageEdgeMode
	EncodeBatch()
	EncodeBatchToCommandBufferSourceImages(commandBuffer unsafe.Pointer, sourceImageBatches ImageBatch /* not a class type */) ImageBatch /* not a class type */
	EncodeBatchToCommandBufferSourceImagesDestinationImages(commandBuffer unsafe.Pointer, sourceImages ImageBatch /* not a class type */, destinationImages ImageBatch /* not a class type */)
	EncodeBatchToCommandBufferSourceImagesDestinationStatesDestinationStateIsTemporary(commandBuffer unsafe.Pointer, sourceImageBatches ImageBatch /* not a class type */, outState StateBatch /* not a class type */, isTemporary bool) ImageBatch /* not a class type */
	Encode()
	EncodeToCommandBufferSourceImages(commandBuffer unsafe.Pointer, sourceImages unsafe.Pointer) IImage
	EncodeToCommandBufferSourceImagesDestinationImage(commandBuffer unsafe.Pointer, sourceImages unsafe.Pointer, destinationImage IImage)
	EncodeToCommandBufferSourceImagesDestinationStateDestinationStateIsTemporary(commandBuffer unsafe.Pointer, sourceImages unsafe.Pointer, outState objectivec.IObject, isTemporary bool) IImage
	IsResultStateReusedAcrossBatch()
	KernelHeight()
	KernelHeightAtIndex(index uint) uint
	KernelWidth()
	KernelWidthAtIndex(index uint) uint
	Offset()
	OffsetAtIndex(index uint) objc.IObject /* cross-framework: MPSOffset */
	ResultStateBatch()
	ResultStateBatchForSourceImagesSourceStatesDestinationImage(sourceImages ImageBatch /* not a class type */, sourceStates StateBatch /* not a class type */, destinationImage ImageBatch /* not a class type */) StateBatch /* not a class type */
	ResultState()
	ResultStateForSourceImagesSourceStatesDestinationImage(sourceImages unsafe.Pointer, sourceStates unsafe.Pointer, destinationImage IImage) IState
	SetDilationRateX()
	SetDilationRateY()
	SetEdgeMode()
	SetKernelHeight()
	SetKernelWidth()
	SetOffset()
	SetSourceFeatureChannelMaxCount()
	SetSourceFeatureChannelOffset()
	SetStrideInPixelsX()
	SetStrideInPixelsY()
	SourceFeatureChannelMaxCount()
	SourceFeatureChannelMaxCountAtIndex(index uint) uint
	SourceFeatureChannelOffset()
	SourceFeatureChannelOffsetAtIndex(index uint) uint
	Stride()
	StrideInPixelsXatIndex(index uint) uint
	StrideInPixelsYatIndex(index uint) uint
	TemporaryResultStateBatch()
	TemporaryResultStateBatchForCommandBufferSourceImagesSourceStatesDestinationImage(commandBuffer unsafe.Pointer, sourceImage ImageBatch /* not a class type */, sourceStates StateBatch /* not a class type */, destinationImage ImageBatch /* not a class type */) StateBatch /* not a class type */
	TemporaryResultState()
	TemporaryResultStateForCommandBufferSourceImagesSourceStatesDestinationImage(commandBuffer unsafe.Pointer, sourceImage unsafe.Pointer, sourceStates unsafe.Pointer, destinationImage IImage) IState
	DilationRateXatIndexWithIndex(index uint) uint
	DilationRateYatIndexWithIndex(index uint) uint
	SetDilationRateXAtIndex(dilationRate uint, index uint)
	SetDilationRateYAtIndex(dilationRate uint, index uint)
	SetEdgeModeAtIndex(edgeMode ImageEdgeMode, index uint)
	SetKernelHeightAtIndex(height uint, index uint)
	SetKernelWidthAtIndex(width uint, index uint)
	SetOffsetAtIndex(offset objc.IObject /* cross-framework: MPSOffset */, index uint)
	SetSourceFeatureChannelMaxCountAtIndex(count uint, index uint)
	SetSourceFeatureChannelOffsetAtIndex(offset uint, index uint)
	SetStrideInPixelsXAtIndex(stride uint, index uint)
	SetStrideInPixelsYAtIndex(stride uint, index uint)


}





// Alloc allocates a new instance without initialization.
func (cc _CNNMultiaryKernelClass) Alloc() CNNMultiaryKernel {
	rv := objc.Send[CNNMultiaryKernel](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNMultiaryKernelClass) New() CNNMultiaryKernel {
	rv := objc.Send[CNNMultiaryKernel](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNMultiaryKernel) Init() CNNMultiaryKernel {
	rv := objc.Send[CNNMultiaryKernel](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNMultiaryKernel) Autorelease() CNNMultiaryKernel {
	rv := objc.Send[CNNMultiaryKernel](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNMultiaryKernel creates a new CNNMultiaryKernel instance.
func NewCNNMultiaryKernel() CNNMultiaryKernel {
	return getCNNMultiaryKernelClass().New()
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNMultiaryKernel
type CNNMultiaryKernel struct {
	Kernel
}

// CNNMultiaryKernelFrom constructs a [CNNMultiaryKernel] from an unsafe.Pointer.
func CNNMultiaryKernelFrom(ptr unsafe.Pointer) CNNMultiaryKernel {
	return CNNMultiaryKernel{
		Kernel: KernelFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043425-initwithcoder
func NewCNNMultiaryKernelWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) CNNMultiaryKernel {
	instance := getCNNMultiaryKernelClass().Alloc()
	rv := objc.Send[CNNMultiaryKernel](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043426-initwithdevice
func NewCNNMultiaryKernelWithDeviceSourceCount(device unsafe.Pointer, sourceCount uint) CNNMultiaryKernel {
	instance := getCNNMultiaryKernelClass().Alloc()
	rv := objc.Send[CNNMultiaryKernel](instance.ID, objc.Sel("initWithDevice:sourceCount:"), device, sourceCount)
	rv.Autorelease()
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043411-appendbatchbarrier
func (c_ CNNMultiaryKernel) AppendBatchBarrier() {
	objc.Send[objc.ID](c_.ID, objc.Sel("appendBatchBarrier"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043415-destinationimagedescriptor
func (c_ CNNMultiaryKernel) DestinationImageDescriptor() {
	objc.Send[objc.ID](c_.ID, objc.Sel("destinationImageDescriptor"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043415-destinationimagedescriptorforsou
func (c_ CNNMultiaryKernel) DestinationImageDescriptorForSourceImagesSourceStates(sourceImages unsafe.Pointer, sourceStates unsafe.Pointer) IImageDescriptor {
	rv := objc.Send[ImageDescriptor](c_.ID, objc.Sel("destinationImageDescriptorForSourceImages:sourceStates:"), sourceImages, sourceStates)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043416-dilationratexatindex
func (c_ CNNMultiaryKernel) DilationRateXatIndex() {
	objc.Send[objc.ID](c_.ID, objc.Sel("dilationRateXatIndex"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043417-dilationrateyatindex
func (c_ CNNMultiaryKernel) DilationRateYatIndex() {
	objc.Send[objc.ID](c_.ID, objc.Sel("dilationRateYatIndex"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043418-edgemode
func (c_ CNNMultiaryKernel) EdgeMode() {
	objc.Send[objc.ID](c_.ID, objc.Sel("edgeMode"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043418-edgemodeatindex
func (c_ CNNMultiaryKernel) EdgeModeAtIndex(index uint) ImageEdgeMode {
	rv := objc.Send[ImageEdgeMode](c_.ID, objc.Sel("edgeModeAtIndex:"), index)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043419-encodebatch
func (c_ CNNMultiaryKernel) EncodeBatch() {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeBatch"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043419-encodebatchtocommandbuffer
func (c_ CNNMultiaryKernel) EncodeBatchToCommandBufferSourceImages(commandBuffer unsafe.Pointer, sourceImageBatches ImageBatch /* not a class type */) ImageBatch /* not a class type */ {
	rv := objc.Send[ImageBatch](c_.ID, objc.Sel("encodeBatchToCommandBuffer:sourceImages:"), commandBuffer, sourceImageBatches)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043420-encodebatchtocommandbuffer
func (c_ CNNMultiaryKernel) EncodeBatchToCommandBufferSourceImagesDestinationImages(commandBuffer unsafe.Pointer, sourceImages ImageBatch /* not a class type */, destinationImages ImageBatch /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeBatchToCommandBuffer:sourceImages:destinationImages:"), commandBuffer, sourceImages, destinationImages)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043421-encodebatchtocommandbuffer
func (c_ CNNMultiaryKernel) EncodeBatchToCommandBufferSourceImagesDestinationStatesDestinationStateIsTemporary(commandBuffer unsafe.Pointer, sourceImageBatches ImageBatch /* not a class type */, outState StateBatch /* not a class type */, isTemporary bool) ImageBatch /* not a class type */ {
	rv := objc.Send[ImageBatch](c_.ID, objc.Sel("encodeBatchToCommandBuffer:sourceImages:destinationStates:destinationStateIsTemporary:"), commandBuffer, sourceImageBatches, outState, isTemporary)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043422-encode
func (c_ CNNMultiaryKernel) Encode() {
	objc.Send[objc.ID](c_.ID, objc.Sel("encode"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043422-encodetocommandbuffer
func (c_ CNNMultiaryKernel) EncodeToCommandBufferSourceImages(commandBuffer unsafe.Pointer, sourceImages unsafe.Pointer) IImage {
	rv := objc.Send[Image](c_.ID, objc.Sel("encodeToCommandBuffer:sourceImages:"), commandBuffer, sourceImages)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043423-encodetocommandbuffer
func (c_ CNNMultiaryKernel) EncodeToCommandBufferSourceImagesDestinationImage(commandBuffer unsafe.Pointer, sourceImages unsafe.Pointer, destinationImage IImage) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeToCommandBuffer:sourceImages:destinationImage:"), commandBuffer, sourceImages, destinationImage)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043424-encodetocommandbuffer
func (c_ CNNMultiaryKernel) EncodeToCommandBufferSourceImagesDestinationStateDestinationStateIsTemporary(commandBuffer unsafe.Pointer, sourceImages unsafe.Pointer, outState objectivec.IObject, isTemporary bool) IImage {
	rv := objc.Send[Image](c_.ID, objc.Sel("encodeToCommandBuffer:sourceImages:destinationState:destinationStateIsTemporary:"), commandBuffer, sourceImages, outState, isTemporary)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043428-isresultstatereusedacrossbatch
func (c_ CNNMultiaryKernel) IsResultStateReusedAcrossBatch() {
	objc.Send[objc.ID](c_.ID, objc.Sel("isResultStateReusedAcrossBatch"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043430-kernelheight
func (c_ CNNMultiaryKernel) KernelHeight() {
	objc.Send[objc.ID](c_.ID, objc.Sel("kernelHeight"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043430-kernelheightatindex
func (c_ CNNMultiaryKernel) KernelHeightAtIndex(index uint) uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("kernelHeightAtIndex:"), index)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043431-kernelwidth
func (c_ CNNMultiaryKernel) KernelWidth() {
	objc.Send[objc.ID](c_.ID, objc.Sel("kernelWidth"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043431-kernelwidthatindex
func (c_ CNNMultiaryKernel) KernelWidthAtIndex(index uint) uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("kernelWidthAtIndex:"), index)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043432-offset
func (c_ CNNMultiaryKernel) Offset() {
	objc.Send[objc.ID](c_.ID, objc.Sel("offset"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043432-offsetatindex
func (c_ CNNMultiaryKernel) OffsetAtIndex(index uint) objc.IObject /* cross-framework: MPSOffset */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("offsetAtIndex:"), index)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043434-resultstatebatch
func (c_ CNNMultiaryKernel) ResultStateBatch() {
	objc.Send[objc.ID](c_.ID, objc.Sel("resultStateBatch"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043434-resultstatebatchforsourceimages
func (c_ CNNMultiaryKernel) ResultStateBatchForSourceImagesSourceStatesDestinationImage(sourceImages ImageBatch /* not a class type */, sourceStates StateBatch /* not a class type */, destinationImage ImageBatch /* not a class type */) StateBatch /* not a class type */ {
	rv := objc.Send[StateBatch](c_.ID, objc.Sel("resultStateBatchForSourceImages:sourceStates:destinationImage:"), sourceImages, sourceStates, destinationImage)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043435-resultstate
func (c_ CNNMultiaryKernel) ResultState() {
	objc.Send[objc.ID](c_.ID, objc.Sel("resultState"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043435-resultstateforsourceimages
func (c_ CNNMultiaryKernel) ResultStateForSourceImagesSourceStatesDestinationImage(sourceImages unsafe.Pointer, sourceStates unsafe.Pointer, destinationImage IImage) IState {
	rv := objc.Send[State](c_.ID, objc.Sel("resultStateForSourceImages:sourceStates:destinationImage:"), sourceImages, sourceStates, destinationImage)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043436-setdilationratex
func (c_ CNNMultiaryKernel) SetDilationRateX() {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDilationRateX"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043437-setdilationratey
func (c_ CNNMultiaryKernel) SetDilationRateY() {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDilationRateY"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043438-setedgemode
func (c_ CNNMultiaryKernel) SetEdgeMode() {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEdgeMode"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043439-setkernelheight
func (c_ CNNMultiaryKernel) SetKernelHeight() {
	objc.Send[objc.ID](c_.ID, objc.Sel("setKernelHeight"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043440-setkernelwidth
func (c_ CNNMultiaryKernel) SetKernelWidth() {
	objc.Send[objc.ID](c_.ID, objc.Sel("setKernelWidth"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043441-setoffset
func (c_ CNNMultiaryKernel) SetOffset() {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOffset"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043442-setsourcefeaturechannelmaxcount
func (c_ CNNMultiaryKernel) SetSourceFeatureChannelMaxCount() {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSourceFeatureChannelMaxCount"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043443-setsourcefeaturechanneloffset
func (c_ CNNMultiaryKernel) SetSourceFeatureChannelOffset() {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSourceFeatureChannelOffset"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043444-setstrideinpixelsx
func (c_ CNNMultiaryKernel) SetStrideInPixelsX() {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStrideInPixelsX"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043445-setstrideinpixelsy
func (c_ CNNMultiaryKernel) SetStrideInPixelsY() {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStrideInPixelsY"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043447-sourcefeaturechannelmaxcount
func (c_ CNNMultiaryKernel) SourceFeatureChannelMaxCount() {
	objc.Send[objc.ID](c_.ID, objc.Sel("sourceFeatureChannelMaxCount"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043447-sourcefeaturechannelmaxcountatin
func (c_ CNNMultiaryKernel) SourceFeatureChannelMaxCountAtIndex(index uint) uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("sourceFeatureChannelMaxCountAtIndex:"), index)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043448-sourcefeaturechanneloffset
func (c_ CNNMultiaryKernel) SourceFeatureChannelOffset() {
	objc.Send[objc.ID](c_.ID, objc.Sel("sourceFeatureChannelOffset"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043448-sourcefeaturechanneloffsetatinde
func (c_ CNNMultiaryKernel) SourceFeatureChannelOffsetAtIndex(index uint) uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("sourceFeatureChannelOffsetAtIndex:"), index)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043449-stride
func (c_ CNNMultiaryKernel) Stride() {
	objc.Send[objc.ID](c_.ID, objc.Sel("stride"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043449-strideinpixelsxatindex
func (c_ CNNMultiaryKernel) StrideInPixelsXatIndex(index uint) uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("strideInPixelsXatIndex:"), index)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043450-strideinpixelsyatindex
func (c_ CNNMultiaryKernel) StrideInPixelsYatIndex(index uint) uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("strideInPixelsYatIndex:"), index)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043451-temporaryresultstatebatch
func (c_ CNNMultiaryKernel) TemporaryResultStateBatch() {
	objc.Send[objc.ID](c_.ID, objc.Sel("temporaryResultStateBatch"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043451-temporaryresultstatebatchforcomm
func (c_ CNNMultiaryKernel) TemporaryResultStateBatchForCommandBufferSourceImagesSourceStatesDestinationImage(commandBuffer unsafe.Pointer, sourceImage ImageBatch /* not a class type */, sourceStates StateBatch /* not a class type */, destinationImage ImageBatch /* not a class type */) StateBatch /* not a class type */ {
	rv := objc.Send[StateBatch](c_.ID, objc.Sel("temporaryResultStateBatchForCommandBuffer:sourceImages:sourceStates:destinationImage:"), commandBuffer, sourceImage, sourceStates, destinationImage)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043452-temporaryresultstate
func (c_ CNNMultiaryKernel) TemporaryResultState() {
	objc.Send[objc.ID](c_.ID, objc.Sel("temporaryResultState"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043452-temporaryresultstateforcommandbu
func (c_ CNNMultiaryKernel) TemporaryResultStateForCommandBufferSourceImagesSourceStatesDestinationImage(commandBuffer unsafe.Pointer, sourceImage unsafe.Pointer, sourceStates unsafe.Pointer, destinationImage IImage) IState {
	rv := objc.Send[State](c_.ID, objc.Sel("temporaryResultStateForCommandBuffer:sourceImages:sourceStates:destinationImage:"), commandBuffer, sourceImage, sourceStates, destinationImage)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNMultiaryKernel/dilationRateXatIndex(_:)
func (c_ CNNMultiaryKernel) DilationRateXatIndexWithIndex(index uint) uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("dilationRateXatIndex:"), index)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNMultiaryKernel/dilationRateYatIndex(_:)
func (c_ CNNMultiaryKernel) DilationRateYatIndexWithIndex(index uint) uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("dilationRateYatIndex:"), index)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNMultiaryKernel/setDilationRateX(_:at:)
func (c_ CNNMultiaryKernel) SetDilationRateXAtIndex(dilationRate uint, index uint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDilationRateX:atIndex:"), dilationRate, index)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNMultiaryKernel/setDilationRateY(_:at:)
func (c_ CNNMultiaryKernel) SetDilationRateYAtIndex(dilationRate uint, index uint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDilationRateY:atIndex:"), dilationRate, index)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNMultiaryKernel/setEdgeMode(_:at:)
func (c_ CNNMultiaryKernel) SetEdgeModeAtIndex(edgeMode ImageEdgeMode, index uint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEdgeMode:atIndex:"), edgeMode, index)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNMultiaryKernel/setKernelHeight(_:at:)
func (c_ CNNMultiaryKernel) SetKernelHeightAtIndex(height uint, index uint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setKernelHeight:atIndex:"), height, index)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNMultiaryKernel/setKernelWidth(_:at:)
func (c_ CNNMultiaryKernel) SetKernelWidthAtIndex(width uint, index uint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setKernelWidth:atIndex:"), width, index)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNMultiaryKernel/setOffset(_:at:)
func (c_ CNNMultiaryKernel) SetOffsetAtIndex(offset objc.IObject /* cross-framework: MPSOffset */, index uint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOffset:atIndex:"), offset, index)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNMultiaryKernel/setSourceFeatureChannelMaxCount(_:at:)
func (c_ CNNMultiaryKernel) SetSourceFeatureChannelMaxCountAtIndex(count uint, index uint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSourceFeatureChannelMaxCount:atIndex:"), count, index)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNMultiaryKernel/setSourceFeatureChannelOffset(_:at:)
func (c_ CNNMultiaryKernel) SetSourceFeatureChannelOffsetAtIndex(offset uint, index uint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSourceFeatureChannelOffset:atIndex:"), offset, index)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNMultiaryKernel/setStrideInPixelsX(_:at:)
func (c_ CNNMultiaryKernel) SetStrideInPixelsXAtIndex(stride uint, index uint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStrideInPixelsX:atIndex:"), stride, index)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNMultiaryKernel/setStrideInPixelsY(_:at:)
func (c_ CNNMultiaryKernel) SetStrideInPixelsYAtIndex(stride uint, index uint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStrideInPixelsY:atIndex:"), stride, index)
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043412-cliprect
func (c_ CNNMultiaryKernel) ClipRect() Region get set /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("clipRect"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043412-cliprect
func (c_ CNNMultiaryKernel) SetClipRect(value Region get set /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setClipRect:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043413-destinationfeaturechanneloffset
func (c_ CNNMultiaryKernel) DestinationFeatureChannelOffset() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("destinationFeatureChannelOffset"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043413-destinationfeaturechanneloffset
func (c_ CNNMultiaryKernel) SetDestinationFeatureChannelOffset(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDestinationFeatureChannelOffset:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043414-destinationimageallocator
func (c_ CNNMultiaryKernel) DestinationImageAllocator() ImageAllocator get set /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("destinationImageAllocator"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043414-destinationimageallocator
func (c_ CNNMultiaryKernel) SetDestinationImageAllocator(value ImageAllocator get set /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDestinationImageAllocator:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043427-isbackwards
func (c_ CNNMultiaryKernel) IsBackwards() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("isBackwards"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043427-isbackwards
func (c_ CNNMultiaryKernel) SetIsBackwards(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsBackwards:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043429-isstatemodified
func (c_ CNNMultiaryKernel) IsStateModified() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("isStateModified"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043429-isstatemodified
func (c_ CNNMultiaryKernel) SetIsStateModified(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsStateModified:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043433-padding
func (c_ CNNMultiaryKernel) Padding() Padding get set /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("padding"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043433-padding
func (c_ CNNMultiaryKernel) SetPadding(value Padding get set /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPadding:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043446-sourcecount
func (c_ CNNMultiaryKernel) SourceCount() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("sourceCount"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043446-sourcecount
func (c_ CNNMultiaryKernel) SetSourceCount(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSourceCount:"), value)
}







