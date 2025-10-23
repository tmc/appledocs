// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/foundation"
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
	Draw() bool
	DrawAtPoint(point coregraphics.CGPoint) bool
	DrawInRect(rect coregraphics.CGRect) bool
	DrawInRectFromRectOperationFractionRespectFlippedHints(dstSpacePortionRect coregraphics.CGRect, srcSpacePortionRect coregraphics.CGRect, op ICompositingOperation, requestedAlpha float64, respectContextIsFlipped bool, hints unsafe.Pointer) bool
	BitsPerSample() int
	SetBitsPerSample(value int)
	ColorSpaceName() ColorSpaceName
	SetColorSpaceName(value IColorSpaceName)
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
	Size() coregraphics.CGSize
	SetSize(value coregraphics.CGSize)
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



// Creates and returns an image representation object from data in an unarchiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageRep/init(coder:)
func NewImageRepWithCoder(coder foundation.ICoder) ImageRep {
	instance := getImageRepClass().Alloc()
	rv := objc.Send[ImageRep](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}


// Creates and returns an image representation object using the contents of the specified pasteboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageRep/init(pasteboard:)
func NewImageRepWithPasteboard(pasteboard IPasteboard) ImageRep {
	rv := objc.Send[ImageRep](objc.ID(getImageRepClass().class), objc.Sel("imageRepWithPasteboard:"), pasteboard)
	return rv
}



// Returns the image representation subclass that handles the specified type of data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageRep/class(for:)
func (ic _ImageRepClass) ImageRepClassForData(data foundation.IData) objc.Class {
	rv := objc.Send[objc.Class](objc.ID(ic.class), objc.Sel("imageRepClassForData:"), data)
	return rv
}


// Returns the list of pasteboard types supported directly by the image representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageRep/imageUnfilteredPasteboardTypes()
func (ic _ImageRepClass) ImageUnfilteredPasteboardTypes() []string {
	rv := objc.Send[[]string](objc.ID(ic.class), objc.Sel("imageUnfilteredPasteboardTypes"))
	return rv
}


// Creates and returns an image representation object using the contents of the specified pasteboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageRep/init(pasteboard:)
func (ic _ImageRepClass) ImageRepWithPasteboard(pasteboard IPasteboard) ImageRep {
	rv := objc.Send[ImageRep](objc.ID(ic.class), objc.Sel("imageRepWithPasteboard:"), pasteboard)
	return rv
}


// Implemented by subclasses to draw the image in the current coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageRep/draw()
func (i_ ImageRep) Draw() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("draw"))
	return rv
}


// Draws the image representation’s image data at the specified point in the current coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageRep/draw(at:)
func (i_ ImageRep) DrawAtPoint(point coregraphics.CGPoint) bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("drawAtPoint:"), point)
	return rv
}


// Draws the image, scaling it (as needed) to fit the specified rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageRep/draw(in:)
func (i_ ImageRep) DrawInRect(rect coregraphics.CGRect) bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("drawInRect:"), rect)
	return rv
}


// Draws all or part of the image in the specified rectangle in the current coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageRep/draw(in:from:operation:fraction:respectFlipped:hints:)
func (i_ ImageRep) DrawInRectFromRectOperationFractionRespectFlippedHints(dstSpacePortionRect coregraphics.CGRect, srcSpacePortionRect coregraphics.CGRect, op ICompositingOperation, requestedAlpha float64, respectContextIsFlipped bool, hints unsafe.Pointer) bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("drawInRect:fromRect:operation:fraction:respectFlipped:hints:"), dstSpacePortionRect, srcSpacePortionRect, op, requestedAlpha, respectContextIsFlipped, hints)
	return rv
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
func (i_ ImageRep) ColorSpaceName() ColorSpaceName {
	rv := objc.Send[ColorSpaceName](i_.ID, objc.Sel("colorSpaceName"))
	return rv
}


// The name of the color space used by the image data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/colorspacename
func (i_ ImageRep) SetColorSpaceName(value IColorSpaceName) {
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
func (i_ ImageRep) Size() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](i_.ID, objc.Sel("size"))
	return rv
}


// The size of the image representation, measured in points in the user coordinate space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/size
func (i_ ImageRep) SetSize(value coregraphics.CGSize) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSize:"), value)
}


