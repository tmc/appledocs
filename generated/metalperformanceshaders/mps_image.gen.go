// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [Image] class.
type IImage interface {
	objectivec.IObject
}

// A texture that may have more than four channels for use in convolutional neural networks.
//
// Some image types, such as those found in convolutional neural networks (CNN), differ from a standard texture in that they may have more than 4 channels per pixel. While the channels could hold RGBA data, they will more commonly hold a number of structural permutations upon an RGBA image as the neural network progresses. It is not uncommon for each pixel to have 32 or 64 channels in it. Since a standard object cannot have more than 4 channels, the additional channels are stored in slices of a 2D texture array (i.e. a texture of type ) such that 4 consecutive channels are stored in each slice of this array. If the number of feature channels is , the number of array slices needed is . For example, a 9-channel CNN image with a width of 3 and a height of 2 will be stored as follows: Thus, the width and height of the underlying 2D texture array is the same as the width and height of the object and the array length is equal to . (Channels marked with a are just for padding and should not contain or values.) An object can contain multiple CNN images for batch processing. In order to create an object that contains images, create an object with the property set to . The length of the 2D texture array (i.e. the number of slices) will be equal to , where consecutive slices of this array represent one image. Although an object can contain more than one image, the actual number of images among these processed by an object is controlled by the dimension of the property. (A kernel processes images from this collection.) The starting index of the image to process from the source object is given by . The starting index of the image in the destination object where this processed image is written to is given by . Thus, an object takes the image from the source at indices , processes each independently, and stores the result in the destination at indices respectively. Thus, should be , should be , and must be . For example, suppose an object takes an input image with 16 channels and outputs an image with 32 channels. The number of slices needed in the source 2D texture array is 4 and the number of slices needed in the destination 2D texture array is 8. Suppose the source batch size is 5 and the destination batch size is 4. Thus, the number of source slices will be and the number of destination slices will be . If you want to process image 2 and 3 of the source and store the result at index 1 and 2 in the destination, you can achieve this by setting , , and . The object will take, in this case, slices 4 and 5 of the source and produce slices 4 to 7 of the destination. Similarly, slices 6 and 7 will be used to produce slices 8 to 11 of the destination. All objects process images in the batch independently. That is, calling a object on a batch is formally the same as calling it on each image in the batch sequentially. Computational and GPU work submission overhead will be amortized over more work if batch processing is used. This is especially important for better performance on small images. If and (i.e. only one slice is needed to represent the image), the underlying metal texture type is chosen to be rather than as explained above. The framework also provides objects, intended for very short-lived image data that is produced and consumed immediately in the same object. They are a useful way to minimize CPU-side texture allocation costs and greatly reduce the amount of memory used by your image pipeline. Creation of the underlying texture may occur lazily in some cases. In general, you should avoid calling the property to avoid materializing memory for longer than necessary. When possible, use the other properties to get information about the object instead.
//
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

// Alloc allocates a new instance without initialization.
func (ic _ImageClass) Alloc() Image {
	rv := objc.Send[Image](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The underlying texture.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImage/texture
func (i_ Image) Texture() objc.ID {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("texture"))
	return rv
}



