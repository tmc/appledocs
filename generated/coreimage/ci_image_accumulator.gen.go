// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/coregraphics"
)

// The class instance for the [ImageAccumulator] class.
var (
	ImageAccumulatorClass     _ImageAccumulatorClass
	ImageAccumulatorClassOnce sync.Once
)

func getImageAccumulatorClass() _ImageAccumulatorClass {
	ImageAccumulatorClassOnce.Do(func() {
		ImageAccumulatorClass = _ImageAccumulatorClass{objc.GetClass("CIImageAccumulator")}
	})
	return ImageAccumulatorClass
}

type _ImageAccumulatorClass struct {
	class objc.Class
}

// An interface definition for the [ImageAccumulator] class.
type IImageAccumulator interface {
	objectivec.IObject
	Clear()
	Image() unsafe.Pointer
	SetImage(image unsafe.Pointer)
	SetImageDirtyRect(image unsafe.Pointer, dirtyRect coregraphics.CGRect)
}

// An object that manages feedback-based image processing for tasks such as painting or fluid simulation.
//
// The class enables feedback-based image processing for such things as iterative painting operations or fluid dynamics simulations. You use objects in conjunction with other Core Image classes, such as , , , and , to take advantage of the built-in Core Image filters when processing images.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImageAccumulator
type ImageAccumulator struct {
	objectivec.Object
}

// ImageAccumulatorFrom constructs a [ImageAccumulator] from an unsafe.Pointer.
//
// An object that manages feedback-based image processing for tasks such as painting or fluid simulation.
func ImageAccumulatorFrom(ptr unsafe.Pointer) ImageAccumulator {
	return ImageAccumulator{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _ImageAccumulatorClass) Alloc() ImageAccumulator {
	rv := objc.Send[ImageAccumulator](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _ImageAccumulatorClass) New() ImageAccumulator {
	rv := objc.Send[ImageAccumulator](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageAccumulator) Init() ImageAccumulator {
	rv := objc.Send[ImageAccumulator](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageAccumulator) Autorelease() ImageAccumulator {
	rv := objc.Send[ImageAccumulator](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageAccumulator creates a new ImageAccumulator instance.
func NewImageAccumulator() ImageAccumulator {
	return getImageAccumulatorClass().New()
}


// Initializes an image accumulator with the specified extent, pixel format, and color space.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImageAccumulator/init(extent:format:colorSpace:)
func NewImageAccumulatorWithExtentFormatColorSpace(extent coregraphics.CGRect, format unsafe.Pointer, colorSpace coregraphics.CGColorSpaceRef) ImageAccumulator {
	instance := getImageAccumulatorClass().Alloc()
	rv := objc.Send[ImageAccumulator](instance.ID, objc.Sel("initWithExtent:format:colorSpace:"), extent, format, colorSpace)
	rv.Autorelease()
	return rv
}

// Initializes an image accumulator with the specified extent and pixel format.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImageAccumulator/init(extent:format:)
func NewImageAccumulatorWithExtentFormat(extent coregraphics.CGRect, format unsafe.Pointer) ImageAccumulator {
	instance := getImageAccumulatorClass().Alloc()
	rv := objc.Send[ImageAccumulator](instance.ID, objc.Sel("initWithExtent:format:"), extent, format)
	rv.Autorelease()
	return rv
}


// Creates an image accumulator with the specified extent and pixel format.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImageAccumulator/imageAccumulatorWithExtent:format:
func (ic _ImageAccumulatorClass) ImageAccumulatorWithExtentFormat(extent coregraphics.CGRect, format unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ic.class), objc.Sel("imageAccumulatorWithExtent:format:"), extent, format)
	return rv
}

// Creates an image accumulator with the specified extent, pixel format, and color space.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImageAccumulator/imageAccumulatorWithExtent:format:colorSpace:
func (ic _ImageAccumulatorClass) ImageAccumulatorWithExtentFormatColorSpace(extent coregraphics.CGRect, format unsafe.Pointer, colorSpace coregraphics.CGColorSpaceRef) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ic.class), objc.Sel("imageAccumulatorWithExtent:format:colorSpace:"), extent, format, colorSpace)
	return rv
}

// Resets the accumulator, discarding any pending updates and the current content.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImageAccumulator/clear()
func (i_ ImageAccumulator) Clear() {
	objc.Send[objc.ID](i_.ID, objc.Sel("clear"))
}

// Returns the current contents of the image accumulator.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImageAccumulator/image()
func (i_ ImageAccumulator) Image() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("image"))
	return rv
}

// Sets the contents of the image accumulator to the contents of the specified image object.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImageAccumulator/setImage(_:)
func (i_ ImageAccumulator) SetImage(image unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setImage:"), image)
}

// Updates an image accumulator with a subregion of an image object.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImageAccumulator/setImage(_:dirtyRect:)
func (i_ ImageAccumulator) SetImageDirtyRect(image unsafe.Pointer, dirtyRect coregraphics.CGRect) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setImage:dirtyRect:"), image, dirtyRect)
}

// The extent of the image associated with the image accumulator.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImageAccumulator/extent
func (i_ ImageAccumulator) Extent() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](i_.ID, objc.Sel("extent"))
	return rv
}

// The pixel format of the image accumulator.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImageAccumulator/format
func (i_ ImageAccumulator) Format() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("format"))
	return rv
}


