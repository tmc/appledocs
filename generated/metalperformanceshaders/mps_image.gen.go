// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSImage */


/* debug [class_header]: Header for MPSImage */
// The class instance for the [Image] class.
var (
	ImageClass     _ImageClass
	ImageClassOnce sync.Once
)

func getImageClass() _ImageClass {
	ImageClassOnce.Do(func() {
		ImageClass = _ImageClass{objc.GetClass("MPSImage")}
	})
	return ImageClass
}

type _ImageClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Image */
// An interface definition for the [Image] class.
type IImage interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Image */
	// properties:
	Usage() TextureUsage get /* not a class type */
	SetUsage(value TextureUsage get /* not a class type */)
	PixelFormat() PixelFormat get /* not a class type */
	SetPixelFormat(value PixelFormat get /* not a class type */)
	PixelSize() objectivec.IObject
	SetPixelSize(value objectivec.IObject)
	Device() Device get /* not a class type */
	SetDevice(value Device get /* not a class type */)
	Precision() objectivec.IObject
	SetPrecision(value objectivec.IObject)
	Width() objectivec.IObject
	SetWidth(value objectivec.IObject)
	Label() objectivec.IObject
	SetLabel(value objectivec.IObject)
	NumberOfImages() objectivec.IObject
	SetNumberOfImages(value objectivec.IObject)
	FeatureChannels() objectivec.IObject
	SetFeatureChannels(value objectivec.IObject)
	Texture() Texture get /* not a class type */
	SetTexture(value Texture get /* not a class type */)
	TextureType() TextureType get /* not a class type */
	SetTextureType(value TextureType get /* not a class type */)
	Height() objectivec.IObject
	SetHeight(value objectivec.IObject)
	Parent() IMPSImage
	SetParent(value IMPSImage)
	FeatureChannelFormat() ImageFeatureChannelFormat get /* not a class type */
	SetFeatureChannelFormat(value ImageFeatureChannelFormat get /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Image */
	// methods:
	WriteBytes()
	ReadBytes()
	SubImage()
	SubImageWithFeatureChannelRange(range_ corefoundation.Range) IImage
	Synchronize()
	SynchronizeOnCommandBuffer(commandBuffer unsafe.Pointer)
	BatchRepresentation()
	BatchRepresentationWithSubRange(subRange corefoundation.Range) ImageBatch /* not a class type */
	ResourceSize()
	ReadBytesDataLayoutBytesPerRowBytesPerImageRegionFeatureChannelInfoImageIndex(dataBytes objectivec.IObject, dataLayout DataLayout, bytesPerRow uint, bytesPerImage uint, region objc.IObject /* cross-framework: MTLRegion */, featureChannelInfo objc.IObject /* cross-framework: MPSImageReadWriteParams */, imageIndex uint)
	ReadBytesDataLayoutBytesPerRowRegionFeatureChannelInfoImageIndex(dataBytes objectivec.IObject, dataLayout DataLayout, bytesPerRow uint, region objc.IObject /* cross-framework: MTLRegion */, featureChannelInfo objc.IObject /* cross-framework: MPSImageReadWriteParams */, imageIndex uint)
	ReadBytesDataLayoutImageIndex(dataBytes objectivec.IObject, dataLayout DataLayout, imageIndex uint)
	SetPurgeableState(state PurgeableState) PurgeableState
	WriteBytesDataLayoutBytesPerColumnBytesPerRowBytesPerImageRegionFeatureChannelInfoImageIndex(dataBytes objectivec.IObject, dataLayout DataLayout, bytesPerColumn uint, bytesPerRow uint, bytesPerImage uint, region objc.IObject /* cross-framework: MTLRegion */, featureChannelInfo objc.IObject /* cross-framework: MPSImageReadWriteParams */, imageIndex uint)
	WriteBytesDataLayoutBytesPerRowBytesPerImageRegionFeatureChannelInfoImageIndex(dataBytes objectivec.IObject, dataLayout DataLayout, bytesPerRow uint, bytesPerImage uint, region objc.IObject /* cross-framework: MTLRegion */, featureChannelInfo objc.IObject /* cross-framework: MPSImageReadWriteParams */, imageIndex uint)
	WriteBytesDataLayoutBytesPerRowRegionFeatureChannelInfoImageIndex(dataBytes objectivec.IObject, dataLayout DataLayout, bytesPerRow uint, region objc.IObject /* cross-framework: MTLRegion */, featureChannelInfo objc.IObject /* cross-framework: MPSImageReadWriteParams */, imageIndex uint)
	WriteBytesDataLayoutImageIndex(dataBytes objectivec.IObject, dataLayout DataLayout, imageIndex uint)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Image */
// Alloc allocates a new instance without initialization.
func (ic _ImageClass) Alloc() Image {
	rv := objc.Send[Image](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageClass) New() Image {
	rv := objc.Send[Image](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ Image) Init() Image {
	rv := objc.Send[Image](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ Image) Autorelease() Image {
	rv := objc.Send[Image](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImage creates a new Image instance.
func NewImage() Image {
	return getImageClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Image */
// A texture that may have more than four channels for use in convolutional neural networks.
//
// Some image types, such as those found in convolutional neural networks (CNN), differ from a standard texture in that they may have more than 4 channels per pixel. While the channels could hold RGBA data, they will more commonly hold a number of structural permutations upon an RGBA image as the neural network progresses. It is not uncommon for each pixel to have 32 or 64 channels in it. Since a standard object cannot have more than 4 channels, the additional channels are stored in slices of a 2D texture array (i.e. a texture of type ) such that 4 consecutive channels are stored in each slice of this array. If the number of feature channels is , the number of array slices needed is . For example, a 9-channel CNN image with a width of 3 and a height of 2 will be stored as follows: Thus, the width and height of the underlying 2D texture array is the same as the width and height of the object and the array length is equal to . (Channels marked with a are just for padding and should not contain or values.) An object can contain multiple CNN images for batch processing. In order to create an object that contains images, create an object with the property set to . The length of the 2D texture array (i.e. the number of slices) will be equal to , where consecutive slices of this array represent one image. Although an object can contain more than one image, the actual number of images among these processed by an object is controlled by the dimension of the property. (A kernel processes images from this collection.) The starting index of the image to process from the source object is given by . The starting index of the image in the destination object where this processed image is written to is given by . Thus, an object takes the image from the source at indices , processes each independently, and stores the result in the destination at indices respectively. Thus, should be , should be , and must be . For example, suppose an object takes an input image with 16 channels and outputs an image with 32 channels. The number of slices needed in the source 2D texture array is 4 and the number of slices needed in the destination 2D texture array is 8. Suppose the source batch size is 5 and the destination batch size is 4. Thus, the number of source slices will be and the number of destination slices will be . If you want to process image 2 and 3 of the source and store the result at index 1 and 2 in the destination, you can achieve this by setting , , and . The object will take, in this case, slices 4 and 5 of the source and produce slices 4 to 7 of the destination. Similarly, slices 6 and 7 will be used to produce slices 8 to 11 of the destination. All objects process images in the batch independently. That is, calling a object on a batch is formally the same as calling it on each image in the batch sequentially. Computational and GPU work submission overhead will be amortized over more work if batch processing is used. This is especially important for better performance on small images. If and (i.e. only one slice is needed to represent the image), the underlying metal texture type is chosen to be rather than as explained above. The framework also provides objects, intended for very short-lived image data that is produced and consumed immediately in the same object. They are a useful way to minimize CPU-side texture allocation costs and greatly reduce the amount of memory used by your image pipeline. Creation of the underlying texture may occur lazily in some cases. In general, you should avoid calling the property to avoid materializing memory for longer than necessary. When possible, use the other properties to get information about the object instead.


// A texture that may have more than four channels for use in convolutional neural networks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImage
type Image struct {
	objectivec.Object
}

// ImageFrom constructs a [Image] from an unsafe.Pointer.
//
// A texture that may have more than four channels for use in convolutional neural networks.
func ImageFrom(ptr unsafe.Pointer) Image {
	return Image{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Image */

// Initializes an empty image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/1648920-initwithdevice
func NewImageWithDeviceImageDescriptor(device unsafe.Pointer, imageDescriptor IImageDescriptor) Image {
	instance := getImageClass().Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithDevice:imageDescriptor:"), device, imageDescriptor)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageWithDeviceImageDescriptor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/2942493-initwithparentimage
func NewImageWithParentImageSliceRangeFeatureChannels(parent IImage, sliceRange corefoundation.Range, featureChannels uint) Image {
	instance := getImageClass().Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithParentImage:sliceRange:featureChannels:"), parent, sliceRange, featureChannels)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageWithParentImageSliceRangeFeatureChannels */


// Initializes an image from a texture. The user-allocated texture has been created for a specific number of feature channels and number of images.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/2097547-initwithtexture
func NewImageWithTextureFeatureChannels(texture unsafe.Pointer, featureChannels uint) Image {
	instance := getImageClass().Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithTexture:featureChannels:"), texture, featureChannels)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageWithTextureFeatureChannels */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Image */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/2867148-defaultallocator
func (ic _ImageClass) DefaultAllocator() {
	objc.Send[objc.ID](objc.ID(ic.class), objc.Sel("defaultAllocator"))
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DefaultAllocator) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Image */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Image */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/2867055-writebytes
func (i_ Image) WriteBytes() {
	objc.Send[objc.ID](i_.ID, objc.Sel("writeBytes"))
}/* debug [instance_methods/method]: WriteBytes */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/2867105-readbytes
func (i_ Image) ReadBytes() {
	objc.Send[objc.ID](i_.ID, objc.Sel("readBytes"))
}/* debug [instance_methods/method]: ReadBytes */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/2942488-subimage
func (i_ Image) SubImage() {
	objc.Send[objc.ID](i_.ID, objc.Sel("subImage"))
}/* debug [instance_methods/method]: SubImage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/2942488-subimagewithfeaturechannelrange
func (i_ Image) SubImageWithFeatureChannelRange(range_ corefoundation.Range) IImage {
	rv := objc.Send[Image](i_.ID, objc.Sel("subImageWithFeatureChannelRange:"), range_)
	return rv
}/* debug [instance_methods/method]: SubImageWithFeatureChannelRange */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/2942491-synchronize
func (i_ Image) Synchronize() {
	objc.Send[objc.ID](i_.ID, objc.Sel("synchronize"))
}/* debug [instance_methods/method]: Synchronize */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/2942491-synchronizeoncommandbuffer
func (i_ Image) SynchronizeOnCommandBuffer(commandBuffer unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("synchronizeOnCommandBuffer:"), commandBuffer)
}/* debug [instance_methods/method]: SynchronizeOnCommandBuffer */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/2942492-batchrepresentation
func (i_ Image) BatchRepresentation() {
	objc.Send[objc.ID](i_.ID, objc.Sel("batchRepresentation"))
}/* debug [instance_methods/method]: BatchRepresentation */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/2942492-batchrepresentationwithsubrange
func (i_ Image) BatchRepresentationWithSubRange(subRange corefoundation.Range) ImageBatch /* not a class type */ {
	rv := objc.Send[ImageBatch](i_.ID, objc.Sel("batchRepresentationWithSubRange:"), subRange)
	return rv
}/* debug [instance_methods/method]: BatchRepresentationWithSubRange */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/2942494-resourcesize
func (i_ Image) ResourceSize() {
	objc.Send[objc.ID](i_.ID, objc.Sel("resourceSize"))
}/* debug [instance_methods/method]: ResourceSize */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImage/readBytes(_:dataLayout:bytesPerRow:bytesPerImage:region:featureChannelInfo:imageIndex:)
func (i_ Image) ReadBytesDataLayoutBytesPerRowBytesPerImageRegionFeatureChannelInfoImageIndex(dataBytes objectivec.IObject, dataLayout DataLayout, bytesPerRow uint, bytesPerImage uint, region objc.IObject /* cross-framework: MTLRegion */, featureChannelInfo objc.IObject /* cross-framework: MPSImageReadWriteParams */, imageIndex uint) {
	objc.Send[objc.ID](i_.ID, objc.Sel("readBytes:dataLayout:bytesPerRow:bytesPerImage:region:featureChannelInfo:imageIndex:"), dataBytes, dataLayout, bytesPerRow, bytesPerImage, region, featureChannelInfo, imageIndex)
}/* debug [instance_methods/method]: ReadBytesDataLayoutBytesPerRowBytesPerImageRegionFeatureChannelInfoImageIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImage/readBytes(_:dataLayout:bytesPerRow:region:featureChannelInfo:imageIndex:)
func (i_ Image) ReadBytesDataLayoutBytesPerRowRegionFeatureChannelInfoImageIndex(dataBytes objectivec.IObject, dataLayout DataLayout, bytesPerRow uint, region objc.IObject /* cross-framework: MTLRegion */, featureChannelInfo objc.IObject /* cross-framework: MPSImageReadWriteParams */, imageIndex uint) {
	objc.Send[objc.ID](i_.ID, objc.Sel("readBytes:dataLayout:bytesPerRow:region:featureChannelInfo:imageIndex:"), dataBytes, dataLayout, bytesPerRow, region, featureChannelInfo, imageIndex)
}/* debug [instance_methods/method]: ReadBytesDataLayoutBytesPerRowRegionFeatureChannelInfoImageIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImage/readBytes(_:dataLayout:imageIndex:)
func (i_ Image) ReadBytesDataLayoutImageIndex(dataBytes objectivec.IObject, dataLayout DataLayout, imageIndex uint) {
	objc.Send[objc.ID](i_.ID, objc.Sel("readBytes:dataLayout:imageIndex:"), dataBytes, dataLayout, imageIndex)
}/* debug [instance_methods/method]: ReadBytesDataLayoutImageIndex */


// Set (or query) the purgeable state of the image’s underlying texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImage/setPurgeableState(_:)
func (i_ Image) SetPurgeableState(state PurgeableState) PurgeableState {
	rv := objc.Send[PurgeableState](i_.ID, objc.Sel("setPurgeableState:"), state)
	return rv
}/* debug [instance_methods/method]: SetPurgeableState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImage/writeBytes(_:dataLayout:bytesPerColumn:bytesPerRow:bytesPerImage:region:featureChannelInfo:imageIndex:)
func (i_ Image) WriteBytesDataLayoutBytesPerColumnBytesPerRowBytesPerImageRegionFeatureChannelInfoImageIndex(dataBytes objectivec.IObject, dataLayout DataLayout, bytesPerColumn uint, bytesPerRow uint, bytesPerImage uint, region objc.IObject /* cross-framework: MTLRegion */, featureChannelInfo objc.IObject /* cross-framework: MPSImageReadWriteParams */, imageIndex uint) {
	objc.Send[objc.ID](i_.ID, objc.Sel("writeBytes:dataLayout:bytesPerColumn:bytesPerRow:bytesPerImage:region:featureChannelInfo:imageIndex:"), dataBytes, dataLayout, bytesPerColumn, bytesPerRow, bytesPerImage, region, featureChannelInfo, imageIndex)
}/* debug [instance_methods/method]: WriteBytesDataLayoutBytesPerColumnBytesPerRowBytesPerImageRegionFeatureChannelInfoImageIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImage/writeBytes(_:dataLayout:bytesPerRow:bytesPerImage:region:featureChannelInfo:imageIndex:)
func (i_ Image) WriteBytesDataLayoutBytesPerRowBytesPerImageRegionFeatureChannelInfoImageIndex(dataBytes objectivec.IObject, dataLayout DataLayout, bytesPerRow uint, bytesPerImage uint, region objc.IObject /* cross-framework: MTLRegion */, featureChannelInfo objc.IObject /* cross-framework: MPSImageReadWriteParams */, imageIndex uint) {
	objc.Send[objc.ID](i_.ID, objc.Sel("writeBytes:dataLayout:bytesPerRow:bytesPerImage:region:featureChannelInfo:imageIndex:"), dataBytes, dataLayout, bytesPerRow, bytesPerImage, region, featureChannelInfo, imageIndex)
}/* debug [instance_methods/method]: WriteBytesDataLayoutBytesPerRowBytesPerImageRegionFeatureChannelInfoImageIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImage/writeBytes(_:dataLayout:bytesPerRow:region:featureChannelInfo:imageIndex:)
func (i_ Image) WriteBytesDataLayoutBytesPerRowRegionFeatureChannelInfoImageIndex(dataBytes objectivec.IObject, dataLayout DataLayout, bytesPerRow uint, region objc.IObject /* cross-framework: MTLRegion */, featureChannelInfo objc.IObject /* cross-framework: MPSImageReadWriteParams */, imageIndex uint) {
	objc.Send[objc.ID](i_.ID, objc.Sel("writeBytes:dataLayout:bytesPerRow:region:featureChannelInfo:imageIndex:"), dataBytes, dataLayout, bytesPerRow, region, featureChannelInfo, imageIndex)
}/* debug [instance_methods/method]: WriteBytesDataLayoutBytesPerRowRegionFeatureChannelInfoImageIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImage/writeBytes(_:dataLayout:imageIndex:)
func (i_ Image) WriteBytesDataLayoutImageIndex(dataBytes objectivec.IObject, dataLayout DataLayout, imageIndex uint) {
	objc.Send[objc.ID](i_.ID, objc.Sel("writeBytes:dataLayout:imageIndex:"), dataBytes, dataLayout, imageIndex)
}/* debug [instance_methods/method]: WriteBytesDataLayoutImageIndex */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Image */

// The intended usage of the underlying texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/1648828-usage
func (i_ Image) Usage() TextureUsage get /* not a class type */ {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("usage"))
	return rv
}/* debug [instance_properties/getter]: usage */


// The intended usage of the underlying texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/1648828-usage
func (i_ Image) SetUsage(value TextureUsage get /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setUsage:"), value)
}/* debug [instance_properties/setter]: usage */


// The pixel format of the underlying texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/1648844-pixelformat
func (i_ Image) PixelFormat() PixelFormat get /* not a class type */ {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("pixelFormat"))
	return rv
}/* debug [instance_properties/getter]: pixelFormat */


// The pixel format of the underlying texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/1648844-pixelformat
func (i_ Image) SetPixelFormat(value PixelFormat get /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPixelFormat:"), value)
}/* debug [instance_properties/setter]: pixelFormat */


// The number of bytes from the first byte of one pixel to the first byte of the next pixel, in storage order. (Includes padding.)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/1648854-pixelsize
func (i_ Image) PixelSize() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("pixelSize"))
	return rv
}/* debug [instance_properties/getter]: pixelSize */


// The number of bytes from the first byte of one pixel to the first byte of the next pixel, in storage order. (Includes padding.)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/1648854-pixelsize
func (i_ Image) SetPixelSize(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPixelSize:"), value)
}/* debug [instance_properties/setter]: pixelSize */


// The device on which the image will be used.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/1648857-device
func (i_ Image) Device() Device get /* not a class type */ {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("device"))
	return rv
}/* debug [instance_properties/getter]: device */


// The device on which the image will be used.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/1648857-device
func (i_ Image) SetDevice(value Device get /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDevice:"), value)
}/* debug [instance_properties/setter]: device */


// The number of bits of numeric precision available for each feature channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/1648880-precision
func (i_ Image) Precision() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("precision"))
	return rv
}/* debug [instance_properties/getter]: precision */


// The number of bits of numeric precision available for each feature channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/1648880-precision
func (i_ Image) SetPrecision(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPrecision:"), value)
}/* debug [instance_properties/setter]: precision */


// The formal width of the image, in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/1648884-width
func (i_ Image) Width() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("width"))
	return rv
}/* debug [instance_properties/getter]: width */


// The formal width of the image, in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/1648884-width
func (i_ Image) SetWidth(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setWidth:"), value)
}/* debug [instance_properties/setter]: width */


// A string to help identify this object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/1648899-label
func (i_ Image) Label() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("label"))
	return rv
}/* debug [instance_properties/getter]: label */


// A string to help identify this object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/1648899-label
func (i_ Image) SetLabel(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setLabel:"), value)
}/* debug [instance_properties/setter]: label */


// The number of images for batch processing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/1648900-numberofimages
func (i_ Image) NumberOfImages() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("numberOfImages"))
	return rv
}/* debug [instance_properties/getter]: numberOfImages */


// The number of images for batch processing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/1648900-numberofimages
func (i_ Image) SetNumberOfImages(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setNumberOfImages:"), value)
}/* debug [instance_properties/setter]: numberOfImages */


// The number of feature channels per pixel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/1648901-featurechannels
func (i_ Image) FeatureChannels() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("featureChannels"))
	return rv
}/* debug [instance_properties/getter]: featureChannels */


// The number of feature channels per pixel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/1648901-featurechannels
func (i_ Image) SetFeatureChannels(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setFeatureChannels:"), value)
}/* debug [instance_properties/setter]: featureChannels */


// The underlying texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/1648903-texture
func (i_ Image) Texture() Texture get /* not a class type */ {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("texture"))
	return rv
}/* debug [instance_properties/getter]: texture */


// The underlying texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/1648903-texture
func (i_ Image) SetTexture(value Texture get /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setTexture:"), value)
}/* debug [instance_properties/setter]: texture */


// The type of the underlying texture, typically or .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/1648948-texturetype
func (i_ Image) TextureType() TextureType get /* not a class type */ {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("textureType"))
	return rv
}/* debug [instance_properties/getter]: textureType */


// The type of the underlying texture, typically or .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/1648948-texturetype
func (i_ Image) SetTextureType(value TextureType get /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setTextureType:"), value)
}/* debug [instance_properties/setter]: textureType */


// The formal height of the image, in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/1648952-height
func (i_ Image) Height() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("height"))
	return rv
}/* debug [instance_properties/getter]: height */


// The formal height of the image, in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/1648952-height
func (i_ Image) SetHeight(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setHeight:"), value)
}/* debug [instance_properties/setter]: height */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/2942490-parent
func (i_ Image) Parent() IMPSImage {
	rv := objc.Send[Image](i_.ID, objc.Sel("parent"))
	return rv
}/* debug [instance_properties/getter]: parent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/2942490-parent
func (i_ Image) SetParent(value IMPSImage) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setParent:"), value)
}/* debug [instance_properties/setter]: parent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/3131715-featurechannelformat
func (i_ Image) FeatureChannelFormat() ImageFeatureChannelFormat get /* not a class type */ {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("featureChannelFormat"))
	return rv
}/* debug [instance_properties/getter]: featureChannelFormat */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/3131715-featurechannelformat
func (i_ Image) SetFeatureChannelFormat(value ImageFeatureChannelFormat get /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setFeatureChannelFormat:"), value)
}/* debug [instance_properties/setter]: featureChannelFormat */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSImage */


