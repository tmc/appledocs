// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Image] class.
var (
	ImageClass     _ImageClass
	ImageClassOnce sync.Once
)

func getImageClass() _ImageClass {
	ImageClassOnce.Do(func() {
		ImageClass = _ImageClass{objc.GetClass("NSImage")}
	})
	return ImageClass
}

type _ImageClass struct {
	class objc.Class
}

// An interface definition for the [Image] class.
type IImage interface {
	objectivec.IObject
	// properties:
	AccessibilityDescription() objc.IObject /* cross-framework: NSString */
	SetAccessibilityDescription(value objc.IObject /* cross-framework: NSString */)
	AlignmentRect() objc.IObject /* cross-framework: Rect */
	SetAlignmentRect(value objc.IObject /* cross-framework: Rect */)
	BackgroundColor() IColor
	SetBackgroundColor(value IColor)
	CacheMode() ImageCacheMode
	SetCacheMode(value ImageCacheMode)
	CapInsets() objc.IObject /* cross-framework: EdgeInsets */
	SetCapInsets(value objc.IObject /* cross-framework: EdgeInsets */)
	Delegate() objc.ID
	SetDelegate(value objc.ID)
	Template() bool /* primitive/slice/pointer. */
	SetTemplate(value bool /* primitive/slice/pointer. */)
	Valid() bool /* primitive/slice/pointer. */
	Locale() objc.IObject /* cross-framework: Locale */
	MatchesOnMultipleResolution() bool /* primitive/slice/pointer. */
	SetMatchesOnMultipleResolution(value bool /* primitive/slice/pointer. */)
	MatchesOnlyOnBestFittingAxis() bool /* primitive/slice/pointer. */
	SetMatchesOnlyOnBestFittingAxis(value bool /* primitive/slice/pointer. */)
	PrefersColorMatch() bool /* primitive/slice/pointer. */
	SetPrefersColorMatch(value bool /* primitive/slice/pointer. */)
	Representations() []ImageRep /* primitive/slice/pointer. */
	ResizingMode() ImageResizingMode
	SetResizingMode(value ImageResizingMode)
	Size() objc.IObject /* cross-framework: Size */
	SetSize(value objc.IObject /* cross-framework: Size */)
	SymbolConfiguration() IImageSymbolConfiguration
	TIFFRepresentation() objc.IObject /* cross-framework: NSData */
	UsesEPSOnResolutionMismatch() bool /* primitive/slice/pointer. */
	SetUsesEPSOnResolutionMismatch(value bool /* primitive/slice/pointer. */)
	IsTemplate() bool /* primitive/slice/pointer. */
	SetIsTemplate(value bool /* primitive/slice/pointer. */)
	IsValid() bool /* primitive/slice/pointer. */
	SetIsValid(value bool /* primitive/slice/pointer. */)
	Contents() unsafe.Pointer
	SetContents(value unsafe.Pointer)
	ContentsGravity() LayerContentsGravity /* not a class type */
	SetContentsGravity(value LayerContentsGravity /* not a class type */)
	// methods:
	AddRepresentation(imageRep IImageRep)
	AddRepresentations(imageReps []ImageRep /* primitive/slice/pointer. */)
	BestRepresentationForRectContextHints(rect objc.IObject /* cross-framework Rect */, referenceContext IGraphicsContext, hints foundation.IDictionary /* already interface */) IImageRep
	CGImageForProposedRectContextHints(proposedDestRect objc.IObject /* cross-framework Rect */, referenceContext IGraphicsContext, hints foundation.IDictionary /* already interface */) ImageRef /* not a class type */
	DrawAtPointFromRectOperationFraction(point objc.IObject /* cross-framework Point */, fromRect objc.IObject /* cross-framework Rect */, op CompositingOperation, delta float64 /* primitive/slice/pointer. */)
	DrawInRect(rect objc.IObject /* cross-framework Rect */)
	DrawInRectFromRectOperationFraction(rect objc.IObject /* cross-framework Rect */, fromRect objc.IObject /* cross-framework Rect */, op CompositingOperation, delta float64 /* primitive/slice/pointer. */)
	DrawInRectFromRectOperationFractionRespectFlippedHints(dstSpacePortionRect objc.IObject /* cross-framework Rect */, srcSpacePortionRect objc.IObject /* cross-framework Rect */, op CompositingOperation, requestedAlpha float64 /* primitive/slice/pointer. */, respectContextIsFlipped bool /* primitive/slice/pointer. */, hints foundation.IDictionary /* already interface */)
	DrawRepresentationInRect(imageRep IImageRep, rect objc.IObject /* cross-framework Rect */) bool /* primitive/slice/pointer. */
	HitTestRectWithImageDestinationRectContextHintsFlipped(testRectDestSpace objc.IObject /* cross-framework Rect */, imageRectDestSpace objc.IObject /* cross-framework Rect */, context IGraphicsContext, hints foundation.IDictionary /* already interface */, flipped bool /* primitive/slice/pointer. */) bool /* primitive/slice/pointer. */
	LayerContentsForContentsScale(layerContentsScale float64 /* primitive/slice/pointer. */) objc.ID
	Name() objc.IObject /* cross-framework: ImageName */
	Recache()
	RecommendedLayerContentsScale(preferredContentsScale float64 /* primitive/slice/pointer. */) float64 /* primitive/slice/pointer. */
	RemoveRepresentation(imageRep IImageRep)
	SetName(string_ objc.IObject /* cross-framework ImageName */) bool /* primitive/slice/pointer. */
	TIFFRepresentationUsingCompressionFactor(comp TIFFCompression, factor float32 /* primitive/slice/pointer. */) objc.IObject /* cross-framework: Data */
	ImageWithLocale(locale objc.IObject /* cross-framework Locale */) IImage
	ImageWithSymbolConfiguration(configuration IImageSymbolConfiguration) IImage
}

// A high-level interface for manipulating image data.
//
// You use instances of to load existing images, create new images, and draw the resulting image data into your views. Although you use this class predominantly for image-related operations, the class itself knows little about the underlying image data. Instead, it works in conjunction with one or more image representation objects (subclasses of ) to manage and render the image data. For the most part, these interactions are transparent. The class serves many purposes, providing support for the following tasks: Loading images stored on disk or at a specified URL. Drawing images into a view or graphics context. Providing the contents of a object. Creating new images based on a series of captured drawing commands. Producing versions of the image in a different format. The class itself is capable of managing image data in a variety of formats. The specific list of formats is dependent on the version of the operating system but includes many standard formats such as TIFF, JPEG, GIF, PNG, and PDF among others. AppKit manages each format using a specific type of image representation object, whose job is to manage the actual image data. You can get a list of supported formats using the methods described in Determining Supported Types of Images. For more information about how to use image objects in your app, see .


// A high-level interface for manipulating image data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage
type Image struct {
	objectivec.Object
}

// ImageFrom constructs a [Image] from an unsafe.Pointer.
//
// A high-level interface for manipulating image data.
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



// Initializes and returns an image object using the specified file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/init(byReferencingFile:)
func NewImageByReferencingFile(fileName objc.IObject /* cross-framework NSString */) Image {
	instance := getImageClass().Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initByReferencingFile:"), fileName)
	rv.Autorelease()
	return rv
}


// Initializes and returns an image object using the specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/init(byReferencing:)
func NewImageByReferencingURL(url objc.IObject /* cross-framework NSURL */) Image {
	instance := getImageClass().Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initByReferencingURL:"), url)
	rv.Autorelease()
	return rv
}


// Returns the image object associated with the specified name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/init(named:)
func NewImageNamed(name objc.IObject /* cross-framework ImageName */) Image {
	rv := objc.Send[Image](objc.ID(getImageClass().class), objc.Sel("imageNamed:"), name)
	return rv
}


// Creates a new image using the contents of the provided image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/init(cgImage:size:)
func NewImageWithCGImageSize(cgImage ImageRef /* not a class type */, size objc.IObject /* cross-framework Size */) Image {
	instance := getImageClass().Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithCGImage:size:"), cgImage, size)
	rv.Autorelease()
	return rv
}


// Initializes and returns an image object from data in an unarchiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/init(coder:)
func NewImageWithCoder(coder objc.IObject /* cross-framework Coder */) Image {
	instance := getImageClass().Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}


// Initializes and returns an image object with the contents of the specified file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/init(contentsOfFile:)
func NewImageWithContentsOfFile(fileName objc.IObject /* cross-framework NSString */) Image {
	instance := getImageClass().Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithContentsOfFile:"), fileName)
	rv.Autorelease()
	return rv
}


// Initializes and returns an image object with the contents of the specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/init(contentsOf:)
func NewImageWithContentsOfURL(url objc.IObject /* cross-framework NSURL */) Image {
	instance := getImageClass().Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithContentsOfURL:"), url)
	rv.Autorelease()
	return rv
}


// Initializes and returns an image object using the provided image data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/init(data:)
func NewImageWithData(data objc.IObject /* cross-framework NSData */) Image {
	instance := getImageClass().Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithData:"), data)
	rv.Autorelease()
	return rv
}


// Initializes and returns an image object using the provided image data and ignoring the EXIF orientation tags.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/init(dataIgnoringOrientation:)
func NewImageWithDataIgnoringOrientation(data objc.IObject /* cross-framework NSData */) Image {
	instance := getImageClass().Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithDataIgnoringOrientation:"), data)
	rv.Autorelease()
	return rv
}


// Initializes and returns an image object with data from the specified pasteboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/init(pasteboard:)
func NewImageWithPasteboard(pasteboard IPasteboard) Image {
	instance := getImageClass().Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithPasteboard:"), pasteboard)
	rv.Autorelease()
	return rv
}


// Initializes and returns an image object with the specified dimensions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/init(size:)
func NewImageWithSize(size objc.IObject /* cross-framework Size */) Image {
	instance := getImageClass().Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithSize:"), size)
	rv.Autorelease()
	return rv
}


// Creates and returns an image object whose contents are drawn using the specified block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/init(size:flipped:drawingHandler:)
func NewImageWithSizeFlippedDrawingHandler(size objc.IObject /* cross-framework Size */, drawingHandlerShouldBeCalledWithFlippedContext bool /* primitive/slice/pointer. */, drawingHandler unsafe.Pointer) Image {
	rv := objc.Send[Image](objc.ID(getImageClass().class), objc.Sel("imageWithSize:flipped:drawingHandler:"), size, drawingHandlerShouldBeCalledWithFlippedContext, drawingHandler)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/init(symbolName:bundle:variableValue:)
func NewImageWithSymbolNameBundleVariableValue(name objc.IObject /* cross-framework NSString */, bundle objc.IObject /* cross-framework Bundle */, value float64 /* primitive/slice/pointer. */) Image {
	rv := objc.Send[Image](objc.ID(getImageClass().class), objc.Sel("imageWithSymbolName:bundle:variableValue:"), name, bundle, value)
	return rv
}


// Creates a symbol image with the symbol name and variable value you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/init(symbolName:variableValue:)
func NewImageWithSymbolNameVariableValue(name objc.IObject /* cross-framework NSString */, value float64 /* primitive/slice/pointer. */) Image {
	rv := objc.Send[Image](objc.ID(getImageClass().class), objc.Sel("imageWithSymbolName:variableValue:"), name, value)
	return rv
}


// Creates a symbol image with the system symbol name and accessibility description you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/init(systemSymbolName:accessibilityDescription:)
func NewImageWithSystemSymbolNameAccessibilityDescription(name objc.IObject /* cross-framework NSString */, description objc.IObject /* cross-framework NSString */) Image {
	rv := objc.Send[Image](objc.ID(getImageClass().class), objc.Sel("imageWithSystemSymbolName:accessibilityDescription:"), name, description)
	return rv
}


// Creates a symbol image with the system symbol name and variable value you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/init(systemSymbolName:variableValue:accessibilityDescription:)
func NewImageWithSystemSymbolNameVariableValueAccessibilityDescription(name objc.IObject /* cross-framework NSString */, value float64 /* primitive/slice/pointer. */, description objc.IObject /* cross-framework NSString */) Image {
	rv := objc.Send[Image](objc.ID(getImageClass().class), objc.Sel("imageWithSystemSymbolName:variableValue:accessibilityDescription:"), name, value, description)
	return rv
}



// Tests whether the image can create an instance of itself using pasteboard data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/canInit(with:)
func (ic _ImageClass) CanInitWithPasteboard(pasteboard IPasteboard) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](objc.ID(ic.class), objc.Sel("canInitWithPasteboard:"), pasteboard)
	return rv
}


// Returns an array of strings identifying the image types supported by the registered image representation objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/imageFileTypes()
func (ic _ImageClass) ImageFileTypes() []string /* primitive/slice/pointer. */ {
	rv := objc.Send[[]string](objc.ID(ic.class), objc.Sel("imageFileTypes"))
	return rv
}


// Returns an array of strings identifying the pasteboard types supported directly by the registered image representation objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/imagePasteboardTypes()
func (ic _ImageClass) ImagePasteboardTypes() []string /* primitive/slice/pointer. */ {
	rv := objc.Send[[]string](objc.ID(ic.class), objc.Sel("imagePasteboardTypes"))
	return rv
}


// Returns an array of strings identifying the pasteboard types supported directly by the registered image representation objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/imageUnfilteredPasteboardTypes()
func (ic _ImageClass) ImageUnfilteredPasteboardTypes() []string /* primitive/slice/pointer. */ {
	rv := objc.Send[[]string](objc.ID(ic.class), objc.Sel("imageUnfilteredPasteboardTypes"))
	return rv
}


// Returns the image object associated with the specified name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/init(named:)
func (ic _ImageClass) ImageNamed(name objc.IObject /* cross-framework ImageName */) IImage {
	rv := objc.Send[Image](objc.ID(ic.class), objc.Sel("imageNamed:"), name)
	return rv
}


// Creates and returns an image object whose contents are drawn using the specified block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/init(size:flipped:drawingHandler:)
func (ic _ImageClass) ImageWithSizeFlippedDrawingHandler(size objc.IObject /* cross-framework Size */, drawingHandlerShouldBeCalledWithFlippedContext bool /* primitive/slice/pointer. */, drawingHandler unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ic.class), objc.Sel("imageWithSize:flipped:drawingHandler:"), size, drawingHandlerShouldBeCalledWithFlippedContext, drawingHandler)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/init(symbolName:bundle:variableValue:)
func (ic _ImageClass) ImageWithSymbolNameBundleVariableValue(name objc.IObject /* cross-framework NSString */, bundle objc.IObject /* cross-framework Bundle */, value float64 /* primitive/slice/pointer. */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ic.class), objc.Sel("imageWithSymbolName:bundle:variableValue:"), name, bundle, value)
	return rv
}


// Creates a symbol image with the symbol name and variable value you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/init(symbolName:variableValue:)
func (ic _ImageClass) ImageWithSymbolNameVariableValue(name objc.IObject /* cross-framework NSString */, value float64 /* primitive/slice/pointer. */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ic.class), objc.Sel("imageWithSymbolName:variableValue:"), name, value)
	return rv
}


// Creates a symbol image with the system symbol name and accessibility description you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/init(systemSymbolName:accessibilityDescription:)
func (ic _ImageClass) ImageWithSystemSymbolNameAccessibilityDescription(name objc.IObject /* cross-framework NSString */, description objc.IObject /* cross-framework NSString */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ic.class), objc.Sel("imageWithSystemSymbolName:accessibilityDescription:"), name, description)
	return rv
}


// Creates a symbol image with the system symbol name and variable value you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/init(systemSymbolName:variableValue:accessibilityDescription:)
func (ic _ImageClass) ImageWithSystemSymbolNameVariableValueAccessibilityDescription(name objc.IObject /* cross-framework NSString */, value float64 /* primitive/slice/pointer. */, description objc.IObject /* cross-framework NSString */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ic.class), objc.Sel("imageWithSystemSymbolName:variableValue:accessibilityDescription:"), name, value, description)
	return rv
}


// Returns an array of UTI strings identifying the image types supported by the registered image representation objects, either directly or through a user-installed filter service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/imageTypes
func (ic _ImageClass) ImageTypes() []string /* primitive/slice/pointer. */ {
	rv := objc.Send[[]string](objc.ID(ic.class), objc.Sel("imageTypes"))
	return rv
}

// Returns an array of UTI strings identifying the image types supported directly by the registered image representation objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/imageUnfilteredTypes
func (ic _ImageClass) ImageUnfilteredTypes() []string /* primitive/slice/pointer. */ {
	rv := objc.Send[[]string](objc.ID(ic.class), objc.Sel("imageUnfilteredTypes"))
	return rv
}

// Adds the specified image representation object to the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/addRepresentation(_:)
func (i_ Image) AddRepresentation(imageRep IImageRep) {
	objc.Send[objc.ID](i_.ID, objc.Sel("addRepresentation:"), imageRep)
}


// Adds an array of image representation objects to the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/addRepresentations(_:)
func (i_ Image) AddRepresentations(imageReps []ImageRep /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("addRepresentations:"), imageReps)
}


// Returns the best representation of the image for the specified rectangle using the provided hints.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/bestRepresentation(for:context:hints:)
func (i_ Image) BestRepresentationForRectContextHints(rect objc.IObject /* cross-framework Rect */, referenceContext IGraphicsContext, hints foundation.IDictionary /* already interface */) IImageRep {
	rv := objc.Send[ImageRep](i_.ID, objc.Sel("bestRepresentationForRect:context:hints:"), rect, referenceContext, hints)
	return rv
}


// Returns a Core Graphics image based on the contents of the current image object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/cgImage(forProposedRect:context:hints:)
func (i_ Image) CGImageForProposedRectContextHints(proposedDestRect objc.IObject /* cross-framework Rect */, referenceContext IGraphicsContext, hints foundation.IDictionary /* already interface */) ImageRef /* not a class type */ {
	rv := objc.Send[ImageRef](i_.ID, objc.Sel("CGImageForProposedRect:context:hints:"), proposedDestRect, referenceContext, hints)
	return rv
}


// Draws all or part of the image at the specified point in the current coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/draw(at:from:operation:fraction:)
func (i_ Image) DrawAtPointFromRectOperationFraction(point objc.IObject /* cross-framework Point */, fromRect objc.IObject /* cross-framework Rect */, op CompositingOperation, delta float64 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("drawAtPoint:fromRect:operation:fraction:"), point, fromRect, op, delta)
}


// Draws the image in the specified rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/draw(in:)
func (i_ Image) DrawInRect(rect objc.IObject /* cross-framework Rect */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("drawInRect:"), rect)
}


// Draws all or part of the image in the specified rectangle in the current coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/draw(in:from:operation:fraction:)
func (i_ Image) DrawInRectFromRectOperationFraction(rect objc.IObject /* cross-framework Rect */, fromRect objc.IObject /* cross-framework Rect */, op CompositingOperation, delta float64 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("drawInRect:fromRect:operation:fraction:"), rect, fromRect, op, delta)
}


// Draws all or part of the image in the specified rectangle respecting the hints and the orientation of the current coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/draw(in:from:operation:fraction:respectFlipped:hints:)
func (i_ Image) DrawInRectFromRectOperationFractionRespectFlippedHints(dstSpacePortionRect objc.IObject /* cross-framework Rect */, srcSpacePortionRect objc.IObject /* cross-framework Rect */, op CompositingOperation, requestedAlpha float64 /* primitive/slice/pointer. */, respectContextIsFlipped bool /* primitive/slice/pointer. */, hints foundation.IDictionary /* already interface */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("drawInRect:fromRect:operation:fraction:respectFlipped:hints:"), dstSpacePortionRect, srcSpacePortionRect, op, requestedAlpha, respectContextIsFlipped, hints)
}


// Draws the image using the specified image representation object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/drawRepresentation(_:in:)
func (i_ Image) DrawRepresentationInRect(imageRep IImageRep, rect objc.IObject /* cross-framework Rect */) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](i_.ID, objc.Sel("drawRepresentation:inRect:"), imageRep, rect)
	return rv
}


// Returns whether the destination rectangle would intersect a non-transparent portion of the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/hitTest(_:withDestinationRect:context:hints:flipped:)
func (i_ Image) HitTestRectWithImageDestinationRectContextHintsFlipped(testRectDestSpace objc.IObject /* cross-framework Rect */, imageRectDestSpace objc.IObject /* cross-framework Rect */, context IGraphicsContext, hints foundation.IDictionary /* already interface */, flipped bool /* primitive/slice/pointer. */) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](i_.ID, objc.Sel("hitTestRect:withImageDestinationRect:context:hints:flipped:"), testRectDestSpace, imageRectDestSpace, context, hints, flipped)
	return rv
}


// Returns an object that may be used as the contents of a layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/layerContents(forContentsScale:)
func (i_ Image) LayerContentsForContentsScale(layerContentsScale float64 /* primitive/slice/pointer. */) objc.ID {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("layerContentsForContentsScale:"), layerContentsScale)
	return rv
}


// Returns the name associated with the image, if any.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/name()
func (i_ Image) Name() objc.IObject /* cross-framework: ImageName */ {
	rv := objc.Send[ImageName](i_.ID, objc.Sel("name"))
	return rv
}


// Invalidates and frees offscreen caches of all image representations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/recache()
func (i_ Image) Recache() {
	objc.Send[objc.ID](i_.ID, objc.Sel("recache"))
}


// Returns the recommended layer contents scale for this image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/recommendedLayerContentsScale(_:)
func (i_ Image) RecommendedLayerContentsScale(preferredContentsScale float64 /* primitive/slice/pointer. */) float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](i_.ID, objc.Sel("recommendedLayerContentsScale:"), preferredContentsScale)
	return rv
}


// Removes and releases the specified image representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/removeRepresentation(_:)
func (i_ Image) RemoveRepresentation(imageRep IImageRep) {
	objc.Send[objc.ID](i_.ID, objc.Sel("removeRepresentation:"), imageRep)
}


// Registers the image object under the specified name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/setName(_:)
func (i_ Image) SetName(string_ objc.IObject /* cross-framework ImageName */) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](i_.ID, objc.Sel("setName:"), string_)
	return rv
}


// Returns a data object that contains TIFF data with the specified compression settings for all of the image representations in the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/tiffRepresentation(using:factor:)
func (i_ Image) TIFFRepresentationUsingCompressionFactor(comp TIFFCompression, factor float32 /* primitive/slice/pointer. */) objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[Data](i_.ID, objc.Sel("TIFFRepresentationUsingCompression:factor:"), comp, factor)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/withLocale(_:)
func (i_ Image) ImageWithLocale(locale objc.IObject /* cross-framework Locale */) IImage {
	rv := objc.Send[Image](i_.ID, objc.Sel("imageWithLocale:"), locale)
	return rv
}


// Creates a new symbol image with the specified configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/withSymbolConfiguration(_:)
func (i_ Image) ImageWithSymbolConfiguration(configuration IImageSymbolConfiguration) IImage {
	rv := objc.Send[Image](i_.ID, objc.Sel("imageWithSymbolConfiguration:"), configuration)
	return rv
}


// The image’s accessibility description.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/accessibilityDescription
func (i_ Image) AccessibilityDescription() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](i_.ID, objc.Sel("accessibilityDescription"))
	return rv
}


// The image’s accessibility description.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/accessibilityDescription
func (i_ Image) SetAccessibilityDescription(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setAccessibilityDescription:"), value)
}


// A rectangle that you can use to position the image during layout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/alignmentRect
func (i_ Image) AlignmentRect() objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[Rect](i_.ID, objc.Sel("alignmentRect"))
	return rv
}


// A rectangle that you can use to position the image during layout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/alignmentRect
func (i_ Image) SetAlignmentRect(value objc.IObject /* cross-framework: Rect */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setAlignmentRect:"), value)
}


// The background color for the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/backgroundColor
func (i_ Image) BackgroundColor() IColor {
	rv := objc.Send[Color](i_.ID, objc.Sel("backgroundColor"))
	return rv
}


// The background color for the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/backgroundColor
func (i_ Image) SetBackgroundColor(value IColor) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setBackgroundColor:"), value)
}


// The image’s caching mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/cacheMode-swift.property
func (i_ Image) CacheMode() ImageCacheMode {
	rv := objc.Send[ImageCacheMode](i_.ID, objc.Sel("cacheMode"))
	return rv
}


// The image’s caching mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/cacheMode-swift.property
func (i_ Image) SetCacheMode(value ImageCacheMode) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setCacheMode:"), value)
}


// The cap insets for the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/capInsets
func (i_ Image) CapInsets() objc.IObject /* cross-framework: EdgeInsets */ {
	rv := objc.Send[EdgeInsets](i_.ID, objc.Sel("capInsets"))
	return rv
}


// The cap insets for the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/capInsets
func (i_ Image) SetCapInsets(value objc.IObject /* cross-framework: EdgeInsets */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setCapInsets:"), value)
}


// The image’s delegate object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/delegate
func (i_ Image) Delegate() objc.ID {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("delegate"))
	return rv
}


// The image’s delegate object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/delegate
func (i_ Image) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDelegate:"), value)
}


// Returns an array of UTI strings identifying the image types supported by the registered image representation objects, either directly or through a user-installed filter service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/imageTypes
func (i_ Image) ImageTypes() []string /* primitive/slice/pointer. */ {
	rv := objc.Send[[]string](i_.ID, objc.Sel("imageTypes"))
	return rv
}


// Returns an array of UTI strings identifying the image types supported directly by the registered image representation objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/imageUnfilteredTypes
func (i_ Image) ImageUnfilteredTypes() []string /* primitive/slice/pointer. */ {
	rv := objc.Send[[]string](i_.ID, objc.Sel("imageUnfilteredTypes"))
	return rv
}


// A Boolean value that determines whether the image represents a template image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/isTemplate
func (i_ Image) Template() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](i_.ID, objc.Sel("template"))
	return rv
}


// A Boolean value that determines whether the image represents a template image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/isTemplate
func (i_ Image) SetTemplate(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setTemplate:"), value)
}


// A Boolean value that indicates whether it is possible to draw an image representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/isValid
func (i_ Image) Valid() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](i_.ID, objc.Sel("valid"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/locale
func (i_ Image) Locale() objc.IObject /* cross-framework: Locale */ {
	rv := objc.Send[Locale](i_.ID, objc.Sel("locale"))
	return rv
}


// A Boolean value that indicates whether image representations whose resolution is an integral multiple of the device resolution are a match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/matchesOnMultipleResolution
func (i_ Image) MatchesOnMultipleResolution() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](i_.ID, objc.Sel("matchesOnMultipleResolution"))
	return rv
}


// A Boolean value that indicates whether image representations whose resolution is an integral multiple of the device resolution are a match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/matchesOnMultipleResolution
func (i_ Image) SetMatchesOnMultipleResolution(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMatchesOnMultipleResolution:"), value)
}


// A Boolean value that indicates whether the image matches only on the best fitting axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/matchesOnlyOnBestFittingAxis
func (i_ Image) MatchesOnlyOnBestFittingAxis() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](i_.ID, objc.Sel("matchesOnlyOnBestFittingAxis"))
	return rv
}


// A Boolean value that indicates whether the image matches only on the best fitting axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/matchesOnlyOnBestFittingAxis
func (i_ Image) SetMatchesOnlyOnBestFittingAxis(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMatchesOnlyOnBestFittingAxis:"), value)
}


// A Boolean value that indicates whether the image prefers to choose image representations using color-matching or resolution-matching.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/prefersColorMatch
func (i_ Image) PrefersColorMatch() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](i_.ID, objc.Sel("prefersColorMatch"))
	return rv
}


// A Boolean value that indicates whether the image prefers to choose image representations using color-matching or resolution-matching.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/prefersColorMatch
func (i_ Image) SetPrefersColorMatch(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPrefersColorMatch:"), value)
}


// An array containing all of the image object’s image representations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/representations
func (i_ Image) Representations() []ImageRep /* primitive/slice/pointer. */ {
	rv := objc.Send[[]ImageRep](i_.ID, objc.Sel("representations"))
	return rv
}


// The resizing mode for the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/resizingMode-swift.property
func (i_ Image) ResizingMode() ImageResizingMode {
	rv := objc.Send[ImageResizingMode](i_.ID, objc.Sel("resizingMode"))
	return rv
}


// The resizing mode for the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/resizingMode-swift.property
func (i_ Image) SetResizingMode(value ImageResizingMode) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setResizingMode:"), value)
}


// The size of the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/size
func (i_ Image) Size() objc.IObject /* cross-framework: Size */ {
	rv := objc.Send[Size](i_.ID, objc.Sel("size"))
	return rv
}


// The size of the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/size
func (i_ Image) SetSize(value objc.IObject /* cross-framework: Size */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSize:"), value)
}


// The configuration details for a symbol image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/symbolConfiguration-swift.property
func (i_ Image) SymbolConfiguration() IImageSymbolConfiguration {
	rv := objc.Send[ImageSymbolConfiguration](i_.ID, objc.Sel("symbolConfiguration"))
	return rv
}


// A data object containing TIFF data for all of the image representations in the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/tiffRepresentation
func (i_ Image) TIFFRepresentation() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](i_.ID, objc.Sel("TIFFRepresentation"))
	return rv
}


// A Boolean value that indicates whether EPS representations are preferred when no other representations match the resolution of the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/usesEPSOnResolutionMismatch
func (i_ Image) UsesEPSOnResolutionMismatch() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](i_.ID, objc.Sel("usesEPSOnResolutionMismatch"))
	return rv
}


// A Boolean value that indicates whether EPS representations are preferred when no other representations match the resolution of the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/usesEPSOnResolutionMismatch
func (i_ Image) SetUsesEPSOnResolutionMismatch(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setUsesEPSOnResolutionMismatch:"), value)
}


// A Boolean value that determines whether the image represents a template image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/istemplate
func (i_ Image) IsTemplate() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](i_.ID, objc.Sel("isTemplate"))
	return rv
}


// A Boolean value that determines whether the image represents a template image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/istemplate
func (i_ Image) SetIsTemplate(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIsTemplate:"), value)
}


// A Boolean value that indicates whether it is possible to draw an image representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/isvalid
func (i_ Image) IsValid() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](i_.ID, objc.Sel("isValid"))
	return rv
}


// A Boolean value that indicates whether it is possible to draw an image representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/isvalid
func (i_ Image) SetIsValid(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIsValid:"), value)
}


// An object that provides the contents of the layer. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/contents
func (i_ Image) Contents() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("contents"))
	return rv
}


// An object that provides the contents of the layer. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/contents
func (i_ Image) SetContents(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setContents:"), value)
}


// A constant that specifies how the layer’s contents are positioned or scaled within its bounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/contentsGravity
func (i_ Image) ContentsGravity() LayerContentsGravity /* not a class type */ {
	rv := objc.Send[LayerContentsGravity](i_.ID, objc.Sel("contentsGravity"))
	return rv
}


// A constant that specifies how the layer’s contents are positioned or scaled within its bounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/contentsGravity
func (i_ Image) SetContentsGravity(value LayerContentsGravity /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setContentsGravity:"), value)
}


