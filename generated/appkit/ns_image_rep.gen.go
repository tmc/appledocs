// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	// properties:
	ColorSpaceName() objc.IObject /* cross-framework: ColorSpaceName */
	SetColorSpaceName(value objc.IObject /* cross-framework: ColorSpaceName */)
	Alpha() bool /* primitive/slice/pointer. */
	SetAlpha(value bool /* primitive/slice/pointer. */)
	Opaque() bool /* primitive/slice/pointer. */
	SetOpaque(value bool /* primitive/slice/pointer. */)
	LayoutDirection() ImageLayoutDirection
	SetLayoutDirection(value ImageLayoutDirection)
	PixelsHigh() int /* primitive/slice/pointer. */
	SetPixelsHigh(value int /* primitive/slice/pointer. */)
	BitsPerSample() int /* primitive/slice/pointer. */
	SetBitsPerSample(value int /* primitive/slice/pointer. */)
	HasAlpha() bool /* primitive/slice/pointer. */
	SetHasAlpha(value bool /* primitive/slice/pointer. */)
	IsOpaque() bool /* primitive/slice/pointer. */
	SetIsOpaque(value bool /* primitive/slice/pointer. */)
	PixelsWide() int /* primitive/slice/pointer. */
	SetPixelsWide(value int /* primitive/slice/pointer. */)
	Size() objc.IObject /* cross-framework: Size */
	SetSize(value objc.IObject /* cross-framework: Size */)
	// methods:
	Draw() bool /* primitive/slice/pointer. */
	DrawAtPoint(point objc.IObject /* cross-framework Point */) bool /* primitive/slice/pointer. */
	DrawInRect(rect objc.IObject /* cross-framework Rect */) bool /* primitive/slice/pointer. */
	DrawInRectFromRectOperationFractionRespectFlippedHints(dstSpacePortionRect objc.IObject /* cross-framework Rect */, srcSpacePortionRect objc.IObject /* cross-framework Rect */, op CompositingOperation, requestedAlpha float64 /* primitive/slice/pointer. */, respectContextIsFlipped bool /* primitive/slice/pointer. */, hints foundation.IDictionary /* already interface */) bool /* primitive/slice/pointer. */
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
func NewImageRepWithCoder(coder objc.IObject /* cross-framework Coder */) ImageRep {
	instance := getImageRepClass().Alloc()
	rv := objc.Send[ImageRep](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}


// Creates and returns an image representation object using the data at the specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageRep/init(contentsOf:)
func NewImageRepWithContentsOfURL(url objc.IObject /* cross-framework NSURL */) ImageRep {
	rv := objc.Send[ImageRep](objc.ID(getImageRepClass().class), objc.Sel("imageRepWithContentsOfURL:"), url)
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



// Returns a Boolean value that indicates whether the receiver can initialize itself from the data on the specified pasteboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageRep/canInit(with:)-56pum
func (ic _ImageRepClass) CanInitWithPasteboard(pasteboard IPasteboard) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](objc.ID(ic.class), objc.Sel("canInitWithPasteboard:"), pasteboard)
	return rv
}


// Returns a Boolean value that indicates whether the image representation can initialize itself from the specified data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageRep/canInit(with:)-6zv56
func (ic _ImageRepClass) CanInitWithData(data objc.IObject /* cross-framework NSData */) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](objc.ID(ic.class), objc.Sel("canInitWithData:"), data)
	return rv
}


// Returns the image representation subclass that handles the specified type of data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageRep/class(for:)
func (ic _ImageRepClass) ImageRepClassForData(data objc.IObject /* cross-framework NSData */) objc.Class {
	rv := objc.Send[objc.Class](objc.ID(ic.class), objc.Sel("imageRepClassForData:"), data)
	return rv
}


// Returns the image representation subclass that handles data with the specified type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageRep/class(forFileType:)
func (ic _ImageRepClass) ImageRepClassForFileType(type_ objc.IObject /* cross-framework NSString */) objc.Class {
	rv := objc.Send[objc.Class](objc.ID(ic.class), objc.Sel("imageRepClassForFileType:"), type_)
	return rv
}


// Returns the image representation subclass that handles data with the specified pasteboard type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageRep/class(forPasteboardType:)
func (ic _ImageRepClass) ImageRepClassForPasteboardType(type_ objc.IObject /* cross-framework PasteboardType */) objc.Class {
	rv := objc.Send[objc.Class](objc.ID(ic.class), objc.Sel("imageRepClassForPasteboardType:"), type_)
	return rv
}


// Returns the image representation subclass that handles image data for the specified UTI.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageRep/class(forType:)
func (ic _ImageRepClass) ImageRepClassForType(type_ objc.IObject /* cross-framework NSString */) objc.Class {
	rv := objc.Send[objc.Class](objc.ID(ic.class), objc.Sel("imageRepClassForType:"), type_)
	return rv
}


// Returns the pasteboard types supported by the image representation class or one of its subclasses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageRep/imagePasteboardTypes()
func (ic _ImageRepClass) ImagePasteboardTypes() []string /* primitive/slice/pointer. */ {
	rv := objc.Send[[]string](objc.ID(ic.class), objc.Sel("imagePasteboardTypes"))
	return rv
}


// Creates and returns an array of image representation objects initialized using the contents of the pasteboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageRep/imageReps(with:)
func (ic _ImageRepClass) ImageRepsWithPasteboard(pasteboard IPasteboard) []ImageRep /* primitive/slice/pointer. */ {
	rv := objc.Send[[]ImageRep](objc.ID(ic.class), objc.Sel("imageRepsWithPasteboard:"), pasteboard)
	return rv
}


// Creates and returns an array of image representation objects initialized using the contents of the specified file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageRep/imageReps(withContentsOfFile:)
func (ic _ImageRepClass) ImageRepsWithContentsOfFile(filename objc.IObject /* cross-framework NSString */) []ImageRep /* primitive/slice/pointer. */ {
	rv := objc.Send[[]ImageRep](objc.ID(ic.class), objc.Sel("imageRepsWithContentsOfFile:"), filename)
	return rv
}


// Returns the list of file types supported directly by the image representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageRep/imageUnfilteredFileTypes()
func (ic _ImageRepClass) ImageUnfilteredFileTypes() []string /* primitive/slice/pointer. */ {
	rv := objc.Send[[]string](objc.ID(ic.class), objc.Sel("imageUnfilteredFileTypes"))
	return rv
}


// Returns the list of pasteboard types supported directly by the image representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageRep/imageUnfilteredPasteboardTypes()
func (ic _ImageRepClass) ImageUnfilteredPasteboardTypes() []string /* primitive/slice/pointer. */ {
	rv := objc.Send[[]string](objc.ID(ic.class), objc.Sel("imageUnfilteredPasteboardTypes"))
	return rv
}


// Creates and returns an image representation object using the data at the specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageRep/init(contentsOf:)
func (ic _ImageRepClass) ImageRepWithContentsOfURL(url objc.IObject /* cross-framework NSURL */) IImageRep {
	rv := objc.Send[ImageRep](objc.ID(ic.class), objc.Sel("imageRepWithContentsOfURL:"), url)
	return rv
}


// Creates and returns an image representation object using the contents of the specified pasteboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageRep/init(pasteboard:)
func (ic _ImageRepClass) ImageRepWithPasteboard(pasteboard IPasteboard) IImageRep {
	rv := objc.Send[ImageRep](objc.ID(ic.class), objc.Sel("imageRepWithPasteboard:"), pasteboard)
	return rv
}


// Returns an array containing the registered image representation classes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageRep/registeredClasses
func (ic _ImageRepClass) RegisteredImageRepClasses() []objc.Class /* not a class type */ {
	rv := objc.Send[[]objc.Class](objc.ID(ic.class), objc.Sel("registeredImageRepClasses"))
	return rv
}

// Implemented by subclasses to draw the image in the current coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageRep/draw()
func (i_ ImageRep) Draw() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](i_.ID, objc.Sel("draw"))
	return rv
}


// Draws the image representation’s image data at the specified point in the current coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageRep/draw(at:)
func (i_ ImageRep) DrawAtPoint(point objc.IObject /* cross-framework Point */) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](i_.ID, objc.Sel("drawAtPoint:"), point)
	return rv
}


// Draws the image, scaling it (as needed) to fit the specified rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageRep/draw(in:)
func (i_ ImageRep) DrawInRect(rect objc.IObject /* cross-framework Rect */) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](i_.ID, objc.Sel("drawInRect:"), rect)
	return rv
}


// Draws all or part of the image in the specified rectangle in the current coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageRep/draw(in:from:operation:fraction:respectFlipped:hints:)
func (i_ ImageRep) DrawInRectFromRectOperationFractionRespectFlippedHints(dstSpacePortionRect objc.IObject /* cross-framework Rect */, srcSpacePortionRect objc.IObject /* cross-framework Rect */, op CompositingOperation, requestedAlpha float64 /* primitive/slice/pointer. */, respectContextIsFlipped bool /* primitive/slice/pointer. */, hints foundation.IDictionary /* already interface */) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](i_.ID, objc.Sel("drawInRect:fromRect:operation:fraction:respectFlipped:hints:"), dstSpacePortionRect, srcSpacePortionRect, op, requestedAlpha, respectContextIsFlipped, hints)
	return rv
}


// The name of the color space used by the image data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageRep/colorSpaceName
func (i_ ImageRep) ColorSpaceName() objc.IObject /* cross-framework: ColorSpaceName */ {
	rv := objc.Send[ColorSpaceName](i_.ID, objc.Sel("colorSpaceName"))
	return rv
}


// The name of the color space used by the image data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageRep/colorSpaceName
func (i_ ImageRep) SetColorSpaceName(value objc.IObject /* cross-framework: ColorSpaceName */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setColorSpaceName:"), value)
}


// A Boolean value that indicates whether the image data has an alpha channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageRep/hasAlpha
func (i_ ImageRep) Alpha() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](i_.ID, objc.Sel("alpha"))
	return rv
}


// A Boolean value that indicates whether the image data has an alpha channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageRep/hasAlpha
func (i_ ImageRep) SetAlpha(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setAlpha:"), value)
}


// A Boolean value that indicates whether the image is opaque.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageRep/isOpaque
func (i_ ImageRep) Opaque() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](i_.ID, objc.Sel("opaque"))
	return rv
}


// A Boolean value that indicates whether the image is opaque.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageRep/isOpaque
func (i_ ImageRep) SetOpaque(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setOpaque:"), value)
}


// The layout direction for the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageRep/layoutDirection
func (i_ ImageRep) LayoutDirection() ImageLayoutDirection {
	rv := objc.Send[ImageLayoutDirection](i_.ID, objc.Sel("layoutDirection"))
	return rv
}


// The layout direction for the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageRep/layoutDirection
func (i_ ImageRep) SetLayoutDirection(value ImageLayoutDirection) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setLayoutDirection:"), value)
}


// The height of the image, measured in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageRep/pixelsHigh
func (i_ ImageRep) PixelsHigh() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](i_.ID, objc.Sel("pixelsHigh"))
	return rv
}


// The height of the image, measured in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageRep/pixelsHigh
func (i_ ImageRep) SetPixelsHigh(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPixelsHigh:"), value)
}


// Returns an array containing the registered image representation classes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageRep/registeredClasses
func (i_ ImageRep) RegisteredImageRepClasses() []objc.Class /* not a class type */ {
	rv := objc.Send[[]objc.Class](i_.ID, objc.Sel("registeredImageRepClasses"))
	return rv
}


// The number of bits per sample in the object (if the object is a planar image, this property contains the number of bits per sample per plane).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/bitspersample
func (i_ ImageRep) BitsPerSample() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](i_.ID, objc.Sel("bitsPerSample"))
	return rv
}


// The number of bits per sample in the object (if the object is a planar image, this property contains the number of bits per sample per plane).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/bitspersample
func (i_ ImageRep) SetBitsPerSample(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setBitsPerSample:"), value)
}


// A Boolean value that indicates whether the image data has an alpha channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/hasalpha
func (i_ ImageRep) HasAlpha() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](i_.ID, objc.Sel("hasAlpha"))
	return rv
}


// A Boolean value that indicates whether the image data has an alpha channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/hasalpha
func (i_ ImageRep) SetHasAlpha(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setHasAlpha:"), value)
}


// A Boolean value that indicates whether the image is opaque.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/isopaque
func (i_ ImageRep) IsOpaque() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](i_.ID, objc.Sel("isOpaque"))
	return rv
}


// A Boolean value that indicates whether the image is opaque.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/isopaque
func (i_ ImageRep) SetIsOpaque(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIsOpaque:"), value)
}


// The width of the image, measured in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/pixelswide
func (i_ ImageRep) PixelsWide() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](i_.ID, objc.Sel("pixelsWide"))
	return rv
}


// The width of the image, measured in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/pixelswide
func (i_ ImageRep) SetPixelsWide(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPixelsWide:"), value)
}


// The size of the image representation, measured in points in the user coordinate space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/size
func (i_ ImageRep) Size() objc.IObject /* cross-framework: Size */ {
	rv := objc.Send[Size](i_.ID, objc.Sel("size"))
	return rv
}


// The size of the image representation, measured in points in the user coordinate space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/size
func (i_ ImageRep) SetSize(value objc.IObject /* cross-framework: Size */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSize:"), value)
}


