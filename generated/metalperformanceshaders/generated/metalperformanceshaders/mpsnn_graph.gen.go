// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSNNGraph */


/* debug [class_header]: Header for MPSNNGraph */
// The class instance for the [Graph] class.
var (
	GraphClass     _GraphClass
	GraphClassOnce sync.Once
)

func getGraphClass() _GraphClass {
	GraphClassOnce.Do(func() {
		GraphClass = _GraphClass{objc.GetClass("MPSNNGraph")}
	})
	return GraphClass
}

type _GraphClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Graph */
// An interface definition for the [Graph] class.
type IGraph interface {
	IKernel
	
/* debug [class_interface_properties]: Properties for Graph */
	// properties:
	DestinationImageAllocator() ImageAllocator get set /* not a class type */
	SetDestinationImageAllocator(value ImageAllocator get set /* not a class type */)
	IntermediateImageHandles() Handle get /* not a class type */
	SetIntermediateImageHandles(value Handle get /* not a class type */)
	SourceImageHandles() Handle get /* not a class type */
	SetSourceImageHandles(value Handle get /* not a class type */)
	SourceStateHandles() Handle get /* not a class type */
	SetSourceStateHandles(value Handle get /* not a class type */)
	OutputStateIsTemporary() objectivec.IObject
	SetOutputStateIsTemporary(value objectivec.IObject)
	ResultHandle() Handle get /* not a class type */
	SetResultHandle(value Handle get /* not a class type */)
	ResultStateHandles() Handle get /* not a class type */
	SetResultStateHandles(value Handle get /* not a class type */)
	Format() ImageFeatureChannelFormat get set /* not a class type */
	SetFormat(value ImageFeatureChannelFormat get set /* not a class type */)
	ResultImageIsNeeded() objectivec.IObject
	SetResultImageIsNeeded(value objectivec.IObject)
	DestinationFeatureChannelOffset() int
	SetDestinationFeatureChannelOffset(value int)
	Offset() objc.IObject /* cross-framework: MPSOffset */
	SetOffset(value objc.IObject /* cross-framework: MPSOffset */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Graph */
	// methods:
	Encode()
	EncodeToCommandBufferSourceImagesSourceStatesIntermediateImagesDestinationStates(commandBuffer unsafe.Pointer, sourceImages unsafe.Pointer, sourceStates unsafe.Pointer, intermediateImages unsafe.Pointer, destinationStates unsafe.Pointer) IImage
	EncodeToCommandBufferSourceImages(commandBuffer unsafe.Pointer, sourceImages unsafe.Pointer) IImage
	ExecuteAsync()
	ExecuteAsyncWithSourceImagesCompletionHandler(sourceImages unsafe.Pointer, handler GraphCompletionHandler /* not a class type */) IImage
	EncodeBatch()
	EncodeBatchToCommandBufferSourceImagesSourceStatesIntermediateImagesDestinationStates(commandBuffer unsafe.Pointer, sourceImages ImageBatch /* not a class type */, sourceStates StateBatch /* not a class type */, intermediateImages ImageBatch /* not a class type */, destinationStates StateBatch /* not a class type */) ImageBatch /* not a class type */
	EncodeBatchToCommandBufferSourceImagesSourceStates(commandBuffer unsafe.Pointer, sourceImages ImageBatch /* not a class type */, sourceStates StateBatch /* not a class type */) ImageBatch /* not a class type */
	ReloadFromDataSources()
	ReadCountForSourceImage()
	ReadCountForSourceImageAtIndex(index uint) uint
	ReadCountForSourceState()
	ReadCountForSourceStateAtIndex(index uint) uint
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Graph */
// Alloc allocates a new instance without initialization.
func (gc _GraphClass) Alloc() Graph {
	rv := objc.Send[Graph](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GraphClass) New() Graph {
	rv := objc.Send[Graph](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ Graph) Init() Graph {
	rv := objc.Send[Graph](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ Graph) Autorelease() Graph {
	rv := objc.Send[Graph](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGraph creates a new Graph instance.
func NewGraph() Graph {
	return getGraphClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Graph */
// An optimized representation of a graph of neural network image and filter nodes.
//
// Once you have prepared a graph of , , and, if needed, objects, you may initialize a using the image node that you wish to appear as the result. The graph object will introspect the graph representation and determine which nodes are needed for inputs, and which nodes are produced as output state (if any). Nodes which are not needed to calculate the result image node are ignored. Some nodes may be internally concatenated with other nodes for better performance. During construction, the graph attached to the result node will be parsed and reduced to an optimized representation. This representation may be saved using the protocol for later recall. When decoding a using a , it will be created against the system default . If you would like to set the device, your should conform to the protocol.


// An optimized representation of a graph of neural network image and filter nodes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGraph
type Graph struct {
	Kernel
}

// GraphFrom constructs a [Graph] from an unsafe.Pointer.
//
// An optimized representation of a graph of neural network image and filter nodes.
func GraphFrom(ptr unsafe.Pointer) Graph {
	return Graph{
		Kernel: KernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Graph */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/2867043-initwithcoder
func NewGraphWithCoderDevice(aDecoder Coder /* not a class type */, device unsafe.Pointer) Graph {
	instance := getGraphClass().Alloc()
	rv := objc.Send[Graph](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewGraphWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/2867077-initwithdevice
func NewGraphWithDeviceResultImage(device unsafe.Pointer, resultImage IImageNode) Graph {
	instance := getGraphClass().Alloc()
	rv := objc.Send[Graph](instance.ID, objc.Sel("initWithDevice:resultImage:"), device, resultImage)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewGraphWithDeviceResultImage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/2953955-initwithdevice
func NewGraphWithDeviceResultImageResultImageIsNeeded(device unsafe.Pointer, resultImage IImageNode, resultIsNeeded bool) Graph {
	instance := getGraphClass().Alloc()
	rv := objc.Send[Graph](instance.ID, objc.Sel("initWithDevice:resultImage:resultImageIsNeeded:"), device, resultImage, resultIsNeeded)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewGraphWithDeviceResultImageResultImageIsNeeded */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/3037385-initwithdevice
func NewGraphWithDeviceResultImagesResultsAreNeeded(device unsafe.Pointer, resultImages unsafe.Pointer, areResultsNeeded objectivec.IObject) Graph {
	instance := getGraphClass().Alloc()
	rv := objc.Send[Graph](instance.ID, objc.Sel("initWithDevice:resultImages:resultsAreNeeded:"), device, resultImages, areResultsNeeded)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewGraphWithDeviceResultImagesResultsAreNeeded */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Graph */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/2953953-graphwithdevice
func (gc _GraphClass) GraphWithDeviceResultImage(device unsafe.Pointer, resultImage IImageNode) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(gc.class), objc.Sel("graphWithDevice:resultImage:"), device, resultImage)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=GraphWithDeviceResultImage) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/2953956-graphwithdevice
func (gc _GraphClass) GraphWithDeviceResultImageResultImageIsNeeded(device unsafe.Pointer, resultImage IImageNode, resultIsNeeded bool) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(gc.class), objc.Sel("graphWithDevice:resultImage:resultImageIsNeeded:"), device, resultImage, resultIsNeeded)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=GraphWithDeviceResultImageResultImageIsNeeded) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/3037384-graphwithdevice
func (gc _GraphClass) GraphWithDeviceResultImagesResultsAreNeeded(device unsafe.Pointer, resultImages unsafe.Pointer, areResultsNeeded objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(gc.class), objc.Sel("graphWithDevice:resultImages:resultsAreNeeded:"), device, resultImages, areResultsNeeded)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=GraphWithDeviceResultImagesResultsAreNeeded) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Graph */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Graph */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/2867011-encode
func (g_ Graph) Encode() {
	objc.Send[objc.ID](g_.ID, objc.Sel("encode"))
}/* debug [instance_methods/method]: Encode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/2867011-encodetocommandbuffer
func (g_ Graph) EncodeToCommandBufferSourceImagesSourceStatesIntermediateImagesDestinationStates(commandBuffer unsafe.Pointer, sourceImages unsafe.Pointer, sourceStates unsafe.Pointer, intermediateImages unsafe.Pointer, destinationStates unsafe.Pointer) IImage {
	rv := objc.Send[Image](g_.ID, objc.Sel("encodeToCommandBuffer:sourceImages:sourceStates:intermediateImages:destinationStates:"), commandBuffer, sourceImages, sourceStates, intermediateImages, destinationStates)
	return rv
}/* debug [instance_methods/method]: EncodeToCommandBufferSourceImagesSourceStatesIntermediateImagesDestinationStates */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/2867036-encodetocommandbuffer
func (g_ Graph) EncodeToCommandBufferSourceImages(commandBuffer unsafe.Pointer, sourceImages unsafe.Pointer) IImage {
	rv := objc.Send[Image](g_.ID, objc.Sel("encodeToCommandBuffer:sourceImages:"), commandBuffer, sourceImages)
	return rv
}/* debug [instance_methods/method]: EncodeToCommandBufferSourceImages */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/2890826-executeasync
func (g_ Graph) ExecuteAsync() {
	objc.Send[objc.ID](g_.ID, objc.Sel("executeAsync"))
}/* debug [instance_methods/method]: ExecuteAsync */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/2890826-executeasyncwithsourceimages
func (g_ Graph) ExecuteAsyncWithSourceImagesCompletionHandler(sourceImages unsafe.Pointer, handler GraphCompletionHandler /* not a class type */) IImage {
	rv := objc.Send[Image](g_.ID, objc.Sel("executeAsyncWithSourceImages:completionHandler:"), sourceImages, handler)
	return rv
}/* debug [instance_methods/method]: ExecuteAsyncWithSourceImagesCompletionHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/2942459-encodebatch
func (g_ Graph) EncodeBatch() {
	objc.Send[objc.ID](g_.ID, objc.Sel("encodeBatch"))
}/* debug [instance_methods/method]: EncodeBatch */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/2942459-encodebatchtocommandbuffer
func (g_ Graph) EncodeBatchToCommandBufferSourceImagesSourceStatesIntermediateImagesDestinationStates(commandBuffer unsafe.Pointer, sourceImages ImageBatch /* not a class type */, sourceStates StateBatch /* not a class type */, intermediateImages ImageBatch /* not a class type */, destinationStates StateBatch /* not a class type */) ImageBatch /* not a class type */ {
	rv := objc.Send[ImageBatch](g_.ID, objc.Sel("encodeBatchToCommandBuffer:sourceImages:sourceStates:intermediateImages:destinationStates:"), commandBuffer, sourceImages, sourceStates, intermediateImages, destinationStates)
	return rv
}/* debug [instance_methods/method]: EncodeBatchToCommandBufferSourceImagesSourceStatesIntermediateImagesDestinationStates */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/2953952-encodebatchtocommandbuffer
func (g_ Graph) EncodeBatchToCommandBufferSourceImagesSourceStates(commandBuffer unsafe.Pointer, sourceImages ImageBatch /* not a class type */, sourceStates StateBatch /* not a class type */) ImageBatch /* not a class type */ {
	rv := objc.Send[ImageBatch](g_.ID, objc.Sel("encodeBatchToCommandBuffer:sourceImages:sourceStates:"), commandBuffer, sourceImages, sourceStates)
	return rv
}/* debug [instance_methods/method]: EncodeBatchToCommandBufferSourceImagesSourceStates */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/2976512-reloadfromdatasources
func (g_ Graph) ReloadFromDataSources() {
	objc.Send[objc.ID](g_.ID, objc.Sel("reloadFromDataSources"))
}/* debug [instance_methods/method]: ReloadFromDataSources */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/3037386-readcountforsourceimage
func (g_ Graph) ReadCountForSourceImage() {
	objc.Send[objc.ID](g_.ID, objc.Sel("readCountForSourceImage"))
}/* debug [instance_methods/method]: ReadCountForSourceImage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/3037386-readcountforsourceimageatindex
func (g_ Graph) ReadCountForSourceImageAtIndex(index uint) uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("readCountForSourceImageAtIndex:"), index)
	return rv
}/* debug [instance_methods/method]: ReadCountForSourceImageAtIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/3037387-readcountforsourcestate
func (g_ Graph) ReadCountForSourceState() {
	objc.Send[objc.ID](g_.ID, objc.Sel("readCountForSourceState"))
}/* debug [instance_methods/method]: ReadCountForSourceState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/3037387-readcountforsourcestateatindex
func (g_ Graph) ReadCountForSourceStateAtIndex(index uint) uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("readCountForSourceStateAtIndex:"), index)
	return rv
}/* debug [instance_methods/method]: ReadCountForSourceStateAtIndex */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Graph */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/2866998-destinationimageallocator
func (g_ Graph) DestinationImageAllocator() ImageAllocator get set /* not a class type */ {
	rv := objc.Send[objc.ID](g_.ID, objc.Sel("destinationImageAllocator"))
	return rv
}/* debug [instance_properties/getter]: destinationImageAllocator */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/2866998-destinationimageallocator
func (g_ Graph) SetDestinationImageAllocator(value ImageAllocator get set /* not a class type */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDestinationImageAllocator:"), value)
}/* debug [instance_properties/setter]: destinationImageAllocator */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/2867000-intermediateimagehandles
func (g_ Graph) IntermediateImageHandles() Handle get /* not a class type */ {
	rv := objc.Send[objc.ID](g_.ID, objc.Sel("intermediateImageHandles"))
	return rv
}/* debug [instance_properties/getter]: intermediateImageHandles */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/2867000-intermediateimagehandles
func (g_ Graph) SetIntermediateImageHandles(value Handle get /* not a class type */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setIntermediateImageHandles:"), value)
}/* debug [instance_properties/setter]: intermediateImageHandles */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/2867012-sourceimagehandles
func (g_ Graph) SourceImageHandles() Handle get /* not a class type */ {
	rv := objc.Send[objc.ID](g_.ID, objc.Sel("sourceImageHandles"))
	return rv
}/* debug [instance_properties/getter]: sourceImageHandles */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/2867012-sourceimagehandles
func (g_ Graph) SetSourceImageHandles(value Handle get /* not a class type */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setSourceImageHandles:"), value)
}/* debug [instance_properties/setter]: sourceImageHandles */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/2867056-sourcestatehandles
func (g_ Graph) SourceStateHandles() Handle get /* not a class type */ {
	rv := objc.Send[objc.ID](g_.ID, objc.Sel("sourceStateHandles"))
	return rv
}/* debug [instance_properties/getter]: sourceStateHandles */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/2867056-sourcestatehandles
func (g_ Graph) SetSourceStateHandles(value Handle get /* not a class type */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setSourceStateHandles:"), value)
}/* debug [instance_properties/setter]: sourceStateHandles */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/2867094-outputstateistemporary
func (g_ Graph) OutputStateIsTemporary() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](g_.ID, objc.Sel("outputStateIsTemporary"))
	return rv
}/* debug [instance_properties/getter]: outputStateIsTemporary */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/2867094-outputstateistemporary
func (g_ Graph) SetOutputStateIsTemporary(value objectivec.IObject) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setOutputStateIsTemporary:"), value)
}/* debug [instance_properties/setter]: outputStateIsTemporary */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/2867123-resulthandle
func (g_ Graph) ResultHandle() Handle get /* not a class type */ {
	rv := objc.Send[objc.ID](g_.ID, objc.Sel("resultHandle"))
	return rv
}/* debug [instance_properties/getter]: resultHandle */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/2867123-resulthandle
func (g_ Graph) SetResultHandle(value Handle get /* not a class type */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setResultHandle:"), value)
}/* debug [instance_properties/setter]: resultHandle */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/2867149-resultstatehandles
func (g_ Graph) ResultStateHandles() Handle get /* not a class type */ {
	rv := objc.Send[objc.ID](g_.ID, objc.Sel("resultStateHandles"))
	return rv
}/* debug [instance_properties/getter]: resultStateHandles */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/2867149-resultstatehandles
func (g_ Graph) SetResultStateHandles(value Handle get /* not a class type */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setResultStateHandles:"), value)
}/* debug [instance_properties/setter]: resultStateHandles */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/2953133-format
func (g_ Graph) Format() ImageFeatureChannelFormat get set /* not a class type */ {
	rv := objc.Send[objc.ID](g_.ID, objc.Sel("format"))
	return rv
}/* debug [instance_properties/getter]: format */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/2953133-format
func (g_ Graph) SetFormat(value ImageFeatureChannelFormat get set /* not a class type */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setFormat:"), value)
}/* debug [instance_properties/setter]: format */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/2953954-resultimageisneeded
func (g_ Graph) ResultImageIsNeeded() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](g_.ID, objc.Sel("resultImageIsNeeded"))
	return rv
}/* debug [instance_properties/getter]: resultImageIsNeeded */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/2953954-resultimageisneeded
func (g_ Graph) SetResultImageIsNeeded(value objectivec.IObject) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setResultImageIsNeeded:"), value)
}/* debug [instance_properties/setter]: resultImageIsNeeded */


// The number of channels in the destination image to skip before writing output data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/destinationfeaturechanneloffset
func (g_ Graph) DestinationFeatureChannelOffset() int {
	rv := objc.Send[int](g_.ID, objc.Sel("destinationFeatureChannelOffset"))
	return rv
}/* debug [instance_properties/getter]: destinationFeatureChannelOffset */


// The number of channels in the destination image to skip before writing output data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/destinationfeaturechanneloffset
func (g_ Graph) SetDestinationFeatureChannelOffset(value int) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDestinationFeatureChannelOffset:"), value)
}/* debug [instance_properties/setter]: destinationFeatureChannelOffset */


// The position of the destination image’s clip rectangle origin, relative to the source image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/offset
func (g_ Graph) Offset() objc.IObject /* cross-framework: MPSOffset */ {
	rv := objc.Send[objc.ID](g_.ID, objc.Sel("offset"))
	return rv
}/* debug [instance_properties/getter]: offset */


// The position of the destination image’s clip rectangle origin, relative to the source image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/offset
func (g_ Graph) SetOffset(value objc.IObject /* cross-framework: MPSOffset */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setOffset:"), value)
}/* debug [instance_properties/setter]: offset */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNGraph */


