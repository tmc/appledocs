// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ImageRep] class.
var (
	ImageRepClass     _ImageRepClass
	ImageRepClassOnce sync.Once
)

func getImageRepClass() _ImageRepClass {
	ImageRepClassOnce.Do(func() {
		ImageRepClass = _ImageRepClass{objc.GetClass("NSImageRep")}
	})
	return ImageRepClass
}

type _ImageRepClass struct {
	class objc.Class
}

// An interface definition for the [ImageRep] class.
type IImageRep interface {
	objectivec.IObject
	// properties:
	BitsPerSample() int
	SetBitsPerSample(value int)
	ColorSpaceName() ColorSpaceName /* not a class type */
	SetColorSpaceName(value ColorSpaceName /* not a class type */)
	HasAlpha() bool
	SetHasAlpha(value bool)
	IsOpaque() bool
	SetIsOpaque(value bool)
	LayoutDirection() unsafe.Pointer
	SetLayoutDirection(value unsafe.Pointer)
	PixelsHigh() int
	SetPixelsHigh(value int)
	PixelsWide() int
	SetPixelsWide(value int)
	Size() objc.IObject /* cross-framework: Size */
	SetSize(value objc.IObject /* cross-framework: Size */)
	// methods:
}

// A semiabstract superclass that provides subclasses that you use to draw an image from a particular type of source data.
//
// The class is called “semiabstract” because it has some instance variables and implementation of its own, in addition to defining subclasses. Although an subclass can be used directly, it is typically accessed through an object, which manages a group of image representations, choosing the best one for the current output device.


// A semiabstract superclass that provides subclasses that you use to draw an image from a particular type of source data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageRep
type ImageRep struct {
	objectivec.Object
}

// ImageRepFrom constructs a [ImageRep] from an unsafe.Pointer.
//
// A semiabstract superclass that provides subclasses that you use to draw an image from a particular type of source data.
func ImageRepFrom(ptr unsafe.Pointer) ImageRep {
	return ImageRep{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _ImageRepClass) Alloc() ImageRep {
	rv := objc.Send[ImageRep](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _ImageRepClass) New() ImageRep {
	rv := objc.Send[ImageRep](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageRep) Init() ImageRep {
	rv := objc.Send[ImageRep](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageRep) Autorelease() ImageRep {
	rv := objc.Send[ImageRep](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageRep creates a new ImageRep instance.
func NewImageRep() ImageRep {
	return getImageRepClass().New()
}



// The number of bits per sample in the object (if the object is a planar image, this property contains the number of bits per sample per plane).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/bitspersample
func (i_ ImageRep) BitsPerSample() int {
	rv := objc.Send[int](i_.ID, objc.Sel("bitsPerSample"))
	return rv
}


// The number of bits per sample in the object (if the object is a planar image, this property contains the number of bits per sample per plane).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/bitspersample
func (i_ ImageRep) SetBitsPerSample(value int) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setBitsPerSample:"), value)
}


// The name of the color space used by the image data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/colorspacename
func (i_ ImageRep) ColorSpaceName() ColorSpaceName /* not a class type */ {
	rv := objc.Send[ColorSpaceName](i_.ID, objc.Sel("colorSpaceName"))
	return rv
}


// The name of the color space used by the image data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/colorspacename
func (i_ ImageRep) SetColorSpaceName(value ColorSpaceName /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setColorSpaceName:"), value)
}


// A Boolean value that indicates whether the image data has an alpha channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/hasalpha
func (i_ ImageRep) HasAlpha() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("hasAlpha"))
	return rv
}


// A Boolean value that indicates whether the image data has an alpha channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/hasalpha
func (i_ ImageRep) SetHasAlpha(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setHasAlpha:"), value)
}


// A Boolean value that indicates whether the image is opaque.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/isopaque
func (i_ ImageRep) IsOpaque() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("isOpaque"))
	return rv
}


// A Boolean value that indicates whether the image is opaque.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/isopaque
func (i_ ImageRep) SetIsOpaque(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIsOpaque:"), value)
}


// The layout direction for the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/layoutdirection
func (i_ ImageRep) LayoutDirection() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("layoutDirection"))
	return rv
}


// The layout direction for the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/layoutdirection
func (i_ ImageRep) SetLayoutDirection(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setLayoutDirection:"), value)
}


// The height of the image, measured in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/pixelshigh
func (i_ ImageRep) PixelsHigh() int {
	rv := objc.Send[int](i_.ID, objc.Sel("pixelsHigh"))
	return rv
}


// The height of the image, measured in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/pixelshigh
func (i_ ImageRep) SetPixelsHigh(value int) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPixelsHigh:"), value)
}


// The width of the image, measured in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/pixelswide
func (i_ ImageRep) PixelsWide() int {
	rv := objc.Send[int](i_.ID, objc.Sel("pixelsWide"))
	return rv
}


// The width of the image, measured in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/pixelswide
func (i_ ImageRep) SetPixelsWide(value int) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPixelsWide:"), value)
}


// The size of the image representation, measured in points in the user coordinate space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/size
func (i_ ImageRep) Size() objc.IObject /* cross-framework: Size */ {
	rv := objc.Send[corefoundation.Size](i_.ID, objc.Sel("size"))
	return rv
}


// The size of the image representation, measured in points in the user coordinate space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/size
func (i_ ImageRep) SetSize(value objc.IObject /* cross-framework: Size */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSize:"), value)
}



