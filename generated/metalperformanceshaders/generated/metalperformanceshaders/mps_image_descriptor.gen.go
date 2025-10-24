// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSImageDescriptor */


/* debug [class_header]: Header for MPSImageDescriptor */
// The class instance for the [ImageDescriptor] class.
var (
	ImageDescriptorClass     _ImageDescriptorClass
	ImageDescriptorClassOnce sync.Once
)

func getImageDescriptorClass() _ImageDescriptorClass {
	ImageDescriptorClassOnce.Do(func() {
		ImageDescriptorClass = _ImageDescriptorClass{objc.GetClass("MPSImageDescriptor")}
	})
	return ImageDescriptorClass
}

type _ImageDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ImageDescriptor */
// An interface definition for the [ImageDescriptor] class.
type IImageDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ImageDescriptor */
	// properties:
	ChannelFormat() ImageFeatureChannelFormat get set /* not a class type */
	SetChannelFormat(value ImageFeatureChannelFormat get set /* not a class type */)
	Width() objectivec.IObject
	SetWidth(value objectivec.IObject)
	NumberOfImages() objectivec.IObject
	SetNumberOfImages(value objectivec.IObject)
	PixelFormat() PixelFormat get /* not a class type */
	SetPixelFormat(value PixelFormat get /* not a class type */)
	FeatureChannels() objectivec.IObject
	SetFeatureChannels(value objectivec.IObject)
	CpuCacheMode() CPUCacheMode get set /* not a class type */
	SetCpuCacheMode(value CPUCacheMode get set /* not a class type */)
	Usage() TextureUsage get set /* not a class type */
	SetUsage(value TextureUsage get set /* not a class type */)
	Height() objectivec.IObject
	SetHeight(value objectivec.IObject)
	StorageMode() StorageMode get set /* not a class type */
	SetStorageMode(value StorageMode get set /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ImageDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ImageDescriptor */
// Alloc allocates a new instance without initialization.
func (ic _ImageDescriptorClass) Alloc() ImageDescriptor {
	rv := objc.Send[ImageDescriptor](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageDescriptorClass) New() ImageDescriptor {
	rv := objc.Send[ImageDescriptor](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageDescriptor) Init() ImageDescriptor {
	rv := objc.Send[ImageDescriptor](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageDescriptor) Autorelease() ImageDescriptor {
	rv := objc.Send[ImageDescriptor](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageDescriptor creates a new ImageDescriptor instance.
func NewImageDescriptor() ImageDescriptor {
	return getImageDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ImageDescriptor */
// A description of the attributes used to create an .
//
// You use an to describe and create the properties of an such as its size, pixel format and CPU cache mode.


// A description of the attributes used to create an .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageDescriptor
type ImageDescriptor struct {
	objectivec.Object
}

// ImageDescriptorFrom constructs a [ImageDescriptor] from an unsafe.Pointer.
//
// A description of the attributes used to create an .
func ImageDescriptorFrom(ptr unsafe.Pointer) ImageDescriptor {
	return ImageDescriptor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ImageDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ImageDescriptor */

// Creates an image descriptor for a single image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagedescriptor/1648819-imagedescriptorwithchannelformat
func (ic _ImageDescriptorClass) ImageDescriptorWithChannelFormatWidthHeightFeatureChannels(channelFormat ImageFeatureChannelFormat, width uint, height uint, featureChannels uint) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ic.class), objc.Sel("imageDescriptorWithChannelFormat:width:height:featureChannels:"), channelFormat, width, height, featureChannels)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ImageDescriptorWithChannelFormatWidthHeightFeatureChannels) */


// Creates an image descriptor for an image container with options to set texture usage and batch size (number of images).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagedescriptor/1648893-imagedescriptorwithchannelformat
func (ic _ImageDescriptorClass) ImageDescriptorWithChannelFormatWidthHeightFeatureChannelsNumberOfImagesUsage(channelFormat ImageFeatureChannelFormat, width uint, height uint, featureChannels uint, numberOfImages uint, usage TextureUsage /* not a class type */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ic.class), objc.Sel("imageDescriptorWithChannelFormat:width:height:featureChannels:numberOfImages:usage:"), channelFormat, width, height, featureChannels, numberOfImages, usage)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ImageDescriptorWithChannelFormatWidthHeightFeatureChannelsNumberOfImagesUsage) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ImageDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ImageDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ImageDescriptor */

// The storage format to use for each channel in the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagedescriptor/1648818-channelformat
func (i_ ImageDescriptor) ChannelFormat() ImageFeatureChannelFormat get set /* not a class type */ {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("channelFormat"))
	return rv
}/* debug [instance_properties/getter]: channelFormat */


// The storage format to use for each channel in the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagedescriptor/1648818-channelformat
func (i_ ImageDescriptor) SetChannelFormat(value ImageFeatureChannelFormat get set /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setChannelFormat:"), value)
}/* debug [instance_properties/setter]: channelFormat */


// The width of the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagedescriptor/1648830-width
func (i_ ImageDescriptor) Width() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("width"))
	return rv
}/* debug [instance_properties/getter]: width */


// The width of the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagedescriptor/1648830-width
func (i_ ImageDescriptor) SetWidth(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setWidth:"), value)
}/* debug [instance_properties/setter]: width */


// The number of images for batch processing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagedescriptor/1648846-numberofimages
func (i_ ImageDescriptor) NumberOfImages() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("numberOfImages"))
	return rv
}/* debug [instance_properties/getter]: numberOfImages */


// The number of images for batch processing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagedescriptor/1648846-numberofimages
func (i_ ImageDescriptor) SetNumberOfImages(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setNumberOfImages:"), value)
}/* debug [instance_properties/setter]: numberOfImages */


// The pixel format for the underlying texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagedescriptor/1648913-pixelformat
func (i_ ImageDescriptor) PixelFormat() PixelFormat get /* not a class type */ {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("pixelFormat"))
	return rv
}/* debug [instance_properties/getter]: pixelFormat */


// The pixel format for the underlying texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagedescriptor/1648913-pixelformat
func (i_ ImageDescriptor) SetPixelFormat(value PixelFormat get /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPixelFormat:"), value)
}/* debug [instance_properties/setter]: pixelFormat */


// The number of feature channels per pixel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagedescriptor/1648918-featurechannels
func (i_ ImageDescriptor) FeatureChannels() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("featureChannels"))
	return rv
}/* debug [instance_properties/getter]: featureChannels */


// The number of feature channels per pixel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagedescriptor/1648918-featurechannels
func (i_ ImageDescriptor) SetFeatureChannels(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setFeatureChannels:"), value)
}/* debug [instance_properties/setter]: featureChannels */


// The CPU cache mode of the underlying texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagedescriptor/1648930-cpucachemode
func (i_ ImageDescriptor) CpuCacheMode() CPUCacheMode get set /* not a class type */ {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("cpuCacheMode"))
	return rv
}/* debug [instance_properties/getter]: cpuCacheMode */


// The CPU cache mode of the underlying texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagedescriptor/1648930-cpucachemode
func (i_ ImageDescriptor) SetCpuCacheMode(value CPUCacheMode get set /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setCpuCacheMode:"), value)
}/* debug [instance_properties/setter]: cpuCacheMode */


// Options to specify the intended usage of the underlying texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagedescriptor/1648937-usage
func (i_ ImageDescriptor) Usage() TextureUsage get set /* not a class type */ {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("usage"))
	return rv
}/* debug [instance_properties/getter]: usage */


// Options to specify the intended usage of the underlying texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagedescriptor/1648937-usage
func (i_ ImageDescriptor) SetUsage(value TextureUsage get set /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setUsage:"), value)
}/* debug [instance_properties/setter]: usage */


// The height of the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagedescriptor/1648947-height
func (i_ ImageDescriptor) Height() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("height"))
	return rv
}/* debug [instance_properties/getter]: height */


// The height of the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagedescriptor/1648947-height
func (i_ ImageDescriptor) SetHeight(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setHeight:"), value)
}/* debug [instance_properties/setter]: height */


// The storage mode of underlying texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagedescriptor/1648955-storagemode
func (i_ ImageDescriptor) StorageMode() StorageMode get set /* not a class type */ {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("storageMode"))
	return rv
}/* debug [instance_properties/getter]: storageMode */


// The storage mode of underlying texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagedescriptor/1648955-storagemode
func (i_ ImageDescriptor) SetStorageMode(value StorageMode get set /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setStorageMode:"), value)
}/* debug [instance_properties/setter]: storageMode */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSImageDescriptor */



