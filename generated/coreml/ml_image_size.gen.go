// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ImageSize] class.
var (
	ImageSizeClass     _ImageSizeClass
	ImageSizeClassOnce sync.Once
)

func getImageSizeClass() _ImageSizeClass {
	ImageSizeClassOnce.Do(func() {
		ImageSizeClass = _ImageSizeClass{objc.GetClass("MLImageSize")}
	})
	return ImageSizeClass
}

type _ImageSizeClass struct {
	class objc.Class
}

// An interface definition for the [ImageSize] class.
type IImageSize interface {
	objectivec.IObject
}

// The width and height of an image feature size.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLImageSize
type ImageSize struct {
	objectivec.Object
}

// ImageSizeFrom constructs a [ImageSize] from an unsafe.Pointer.
//
// The width and height of an image feature size.
func ImageSizeFrom(ptr unsafe.Pointer) ImageSize {
	return ImageSize{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _ImageSizeClass) Alloc() ImageSize {
	rv := objc.Send[ImageSize](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _ImageSizeClass) New() ImageSize {
	rv := objc.Send[ImageSize](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageSize) Init() ImageSize {
	rv := objc.Send[ImageSize](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageSize) Autorelease() ImageSize {
	rv := objc.Send[ImageSize](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageSize creates a new ImageSize instance.
func NewImageSize() ImageSize {
	return getImageSizeClass().New()
}


// The height of an image feature in pixels.
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlimagesize/pixelshigh
func (i_ ImageSize) PixelsHigh() int {
	rv := objc.Send[int](i_.ID, objc.Sel("pixelsHigh"))
	return rv
}


// SetPixelsHigh sets the value of the pixelsHigh property.
// The height of an image feature in pixels.

//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlimagesize/pixelshigh
func (i_ ImageSize) SetPixelsHigh(value int) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPixelsHigh:"), value)
}

// The width of an image feature in pixels.
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlimagesize/pixelswide
func (i_ ImageSize) PixelsWide() int {
	rv := objc.Send[int](i_.ID, objc.Sel("pixelsWide"))
	return rv
}


// SetPixelsWide sets the value of the pixelsWide property.
// The width of an image feature in pixels.

//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlimagesize/pixelswide
func (i_ ImageSize) SetPixelsWide(value int) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPixelsWide:"), value)
}

// An array of image sizes a model’s image feature accepts as input or produces as output.
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlimagesizeconstraint/enumeratedimagesizes
func (i_ ImageSize) EnumeratedImageSizes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("enumeratedImageSizes"))
	return rv
}


// SetEnumeratedImageSizes sets the value of the enumeratedImageSizes property.
// An array of image sizes a model’s image feature accepts as input or produces as output.

//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlimagesizeconstraint/enumeratedimagesizes
func (i_ ImageSize) SetEnumeratedImageSizes(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setEnumeratedImageSizes:"), value)
}



