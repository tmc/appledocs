// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
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
	AccessibilityDescription() foundation.foundation.INSString
	SetAccessibilityDescription(value foundation.foundation.INSString)
	AlignmentRect() corefoundation.CGRect
	SetAlignmentRect(value corefoundation.CGRect)
	BackgroundColor() IColor
	SetBackgroundColor(value IColor)
	CacheMode() ImageCacheMode
	SetCacheMode(value ImageCacheMode)
	CapInsets() foundation.EdgeInsets
	SetCapInsets(value foundation.EdgeInsets)
	Template() bool
	SetTemplate(value bool)
	Valid() bool
	Locale() foundation.Locale
	MatchesOnMultipleResolution() bool
	SetMatchesOnMultipleResolution(value bool)
	MatchesOnlyOnBestFittingAxis() bool
	SetMatchesOnlyOnBestFittingAxis(value bool)
	PrefersColorMatch() bool
	SetPrefersColorMatch(value bool)
	Representations() []ImageRep
	ResizingMode() ImageResizingMode
	SetResizingMode(value ImageResizingMode)
	Size() corefoundation.CGSize
	SetSize(value corefoundation.CGSize)
	SymbolConfiguration() IImageSymbolConfiguration
	TIFFRepresentation() foundation.foundation.INSData
	UsesEPSOnResolutionMismatch() bool
	SetUsesEPSOnResolutionMismatch(value bool)
	IsTemplate() bool
	SetIsTemplate(value bool)
	IsValid() bool
	SetIsValid(value bool)
	ContentsGravity() LayerContentsGravity /* not a class type */
	SetContentsGravity(value LayerContentsGravity /* not a class type */)


	

	// methods:
	AddRepresentation(imageRep IImageRep)
	AddRepresentations(imageReps []ImageRep)
	BestRepresentationForRectContextHints(rect corefoundation.CGRect, referenceContext IGraphicsContext, hints foundation.IDictionary) IImageRep
	CGImageForProposedRectContextHints(proposedDestRect corefoundation.CGRect, referenceContext IGraphicsContext, hints foundation.IDictionary) ImageRef /* not a class type */
	DrawAtPointFromRectOperationFraction(point corefoundation.CGPoint, fromRect corefoundation.CGRect, op CompositingOperation, delta float64)
	DrawInRect(rect corefoundation.CGRect)
	DrawInRectFromRectOperationFraction(rect corefoundation.CGRect, fromRect corefoundation.CGRect, op CompositingOperation, delta float64)
	DrawInRectFromRectOperationFractionRespectFlippedHints(dstSpacePortionRect corefoundation.CGRect, srcSpacePortionRect corefoundation.CGRect, op CompositingOperation, requestedAlpha float64, respectContextIsFlipped bool, hints foundation.IDictionary)
	DrawRepresentationInRect(imageRep IImageRep, rect corefoundation.CGRect) bool
	HitTestRectWithImageDestinationRectContextHintsFlipped(testRectDestSpace corefoundation.CGRect, imageRectDestSpace corefoundation.CGRect, context IGraphicsContext, hints foundation.IDictionary, flipped bool) bool
	LayerContentsForContentsScale(layerContentsScale float64) objc.ID
	Name() ImageName
	Recache()
	RecommendedLayerContentsScale(preferredContentsScale float64) float64
	RemoveRepresentation(imageRep IImageRep)
	SetName(string_ ImageName) bool
	TIFFRepresentationUsingCompressionFactor(comp TIFFCompression, factor float32) foundation.Data
	ImageWithLocale(locale foundation.Locale) IImage
	ImageWithSymbolConfiguration(configuration IImageSymbolConfiguration) IImage


}





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






// Initializes and returns an image object using the specified file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/init(byReferencingFile:)
func NewImageByReferencingFile(fileName foundation.foundation.INSString) Image {
	instance := getImageClass().Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initByReferencingFile:"), fileName)
	rv.Autorelease()
	return rv
}


// Initializes and returns an image object using the specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/init(byReferencing:)
func NewImageByReferencingURL(url foundation.foundation.INSURL) Image {
	instance := getImageClass().Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initByReferencingURL:"), url)
	rv.Autorelease()
	return rv
}


// Returns the image object associated with the specified name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/init(named:)
func NewImageNamed(name ImageName) Image {
	rv := objc.Send[Image](objc.ID(getImageClass().class), objc.Sel("imageNamed:"), name)
	return rv
}


// Creates a new image using the contents of the provided image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/init(cgImage:size:)
func NewImageWithCGImageSize(cgImage ImageRef /* not a class type */, size corefoundation.CGSize) Image {
	instance := getImageClass().Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithCGImage:size:"), cgImage, size)
	rv.Autorelease()
	return rv
}


// Initializes and returns an image object from data in an unarchiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/init(coder:)
func NewImageWithCoder(coder foundation.foundation.INSCoder) Image {
	instance := getImageClass().Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}


// Initializes and returns an image object with the contents of the specified file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/init(contentsOfFile:)
func NewImageWithContentsOfFile(fileName foundation.foundation.INSString) Image {
	instance := getImageClass().Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithContentsOfFile:"), fileName)
	rv.Autorelease()
	return rv
}


// Initializes and returns an image object with the contents of the specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/init(contentsOf:)
func NewImageWithContentsOfURL(url foundation.foundation.INSURL) Image {
	instance := getImageClass().Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithContentsOfURL:"), url)
	rv.Autorelease()
	return rv
}


// Initializes and returns an image object using the provided image data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/init(data:)
func NewImageWithData(data foundation.foundation.INSData) Image {
	instance := getImageClass().Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithData:"), data)
	rv.Autorelease()
	return rv
}


// Initializes and returns an image object using the provided image data and ignoring the EXIF orientation tags.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/init(dataIgnoringOrientation:)
func NewImageWithDataIgnoringOrientation(data foundation.foundation.INSData) Image {
	instance := getImageClass().Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithDataIgnoringOrientation:"), data)
	rv.Autorelease()
	return rv
}


// Initializes the image object with a Carbon-style icon resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/init(iconRef:)
func NewImageWithIconRef(iconRef objectivec.IObject) Image {
	instance := getImageClass().Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithIconRef:"), iconRef)
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
func NewImageWithSize(size corefoundation.CGSize) Image {
	instance := getImageClass().Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithSize:"), size)
	rv.Autorelease()
	return rv
}


// Creates and returns an image object whose contents are drawn using the specified block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/init(size:flipped:drawingHandler:)
func NewImageWithSizeFlippedDrawingHandler(size corefoundation.CGSize, drawingHandlerShouldBeCalledWithFlippedContext bool, drawingHandler unsafe.Pointer) Image {
	rv := objc.Send[Image](objc.ID(getImageClass().class), objc.Sel("imageWithSize:flipped:drawingHandler:"), size, drawingHandlerShouldBeCalledWithFlippedContext, drawingHandler)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/init(symbolName:bundle:variableValue:)
func NewImageWithSymbolNameBundleVariableValue(name foundation.foundation.INSString, bundle foundation.Bundle, value float64) Image {
	rv := objc.Send[Image](objc.ID(getImageClass().class), objc.Sel("imageWithSymbolName:bundle:variableValue:"), name, bundle, value)
	return rv
}


// Creates a symbol image with the symbol name and variable value you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/init(symbolName:variableValue:)
func NewImageWithSymbolNameVariableValue(name foundation.foundation.INSString, value float64) Image {
	rv := objc.Send[Image](objc.ID(getImageClass().class), objc.Sel("imageWithSymbolName:variableValue:"), name, value)
	return rv
}


// Creates a symbol image with the system symbol name and accessibility description you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/init(systemSymbolName:accessibilityDescription:)
func NewImageWithSystemSymbolNameAccessibilityDescription(name foundation.foundation.INSString, description foundation.foundation.INSString) Image {
	rv := objc.Send[Image](objc.ID(getImageClass().class), objc.Sel("imageWithSystemSymbolName:accessibilityDescription:"), name, description)
	return rv
}


// Creates a symbol image with the system symbol name and variable value you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/init(systemSymbolName:variableValue:accessibilityDescription:)
func NewImageWithSystemSymbolNameVariableValueAccessibilityDescription(name foundation.foundation.INSString, value float64, description foundation.foundation.INSString) Image {
	rv := objc.Send[Image](objc.ID(getImageClass().class), objc.Sel("imageWithSystemSymbolName:variableValue:accessibilityDescription:"), name, value, description)
	return rv
}







// Tests whether the image can create an instance of itself using pasteboard data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/canInit(with:)
func (ic _ImageClass) CanInitWithPasteboard(pasteboard IPasteboard) bool {
	rv := objc.Send[bool](objc.ID(ic.class), objc.Sel("canInitWithPasteboard:"), pasteboard)
	return rv
}


// Returns an array of strings identifying the image types supported by the registered image representation objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/imageFileTypes()
func (ic _ImageClass) ImageFileTypes() []string {
	rv := objc.Send[[]string](objc.ID(ic.class), objc.Sel("imageFileTypes"))
	return rv
}


// Returns an array of strings identifying the pasteboard types supported directly by the registered image representation objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/imagePasteboardTypes()
func (ic _ImageClass) ImagePasteboardTypes() []string {
	rv := objc.Send[[]string](objc.ID(ic.class), objc.Sel("imagePasteboardTypes"))
	return rv
}


// Returns an array of strings identifying the file types supported directly by the registered image representation objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/imageUnfilteredFileTypes()
func (ic _ImageClass) ImageUnfilteredFileTypes() []string {
	rv := objc.Send[[]string](objc.ID(ic.class), objc.Sel("imageUnfilteredFileTypes"))
	return rv
}


// Returns an array of strings identifying the pasteboard types supported directly by the registered image representation objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/imageUnfilteredPasteboardTypes()
func (ic _ImageClass) ImageUnfilteredPasteboardTypes() []string {
	rv := objc.Send[[]string](objc.ID(ic.class), objc.Sel("imageUnfilteredPasteboardTypes"))
	return rv
}


// Returns the image object associated with the specified name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/init(named:)
func (ic _ImageClass) ImageNamed(name ImageName) IImage {
	rv := objc.Send[Image](objc.ID(ic.class), objc.Sel("imageNamed:"), name)
	return rv
}


// Creates and returns an image object whose contents are drawn using the specified block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/init(size:flipped:drawingHandler:)
func (ic _ImageClass) ImageWithSizeFlippedDrawingHandler(size corefoundation.CGSize, drawingHandlerShouldBeCalledWithFlippedContext bool, drawingHandler unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ic.class), objc.Sel("imageWithSize:flipped:drawingHandler:"), size, drawingHandlerShouldBeCalledWithFlippedContext, drawingHandler)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/init(symbolName:bundle:variableValue:)
func (ic _ImageClass) ImageWithSymbolNameBundleVariableValue(name foundation.foundation.INSString, bundle foundation.Bundle, value float64) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ic.class), objc.Sel("imageWithSymbolName:bundle:variableValue:"), name, bundle, value)
	return rv
}


// Creates a symbol image with the symbol name and variable value you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/init(symbolName:variableValue:)
func (ic _ImageClass) ImageWithSymbolNameVariableValue(name foundation.foundation.INSString, value float64) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ic.class), objc.Sel("imageWithSymbolName:variableValue:"), name, value)
	return rv
}


// Creates a symbol image with the system symbol name and accessibility description you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/init(systemSymbolName:accessibilityDescription:)
func (ic _ImageClass) ImageWithSystemSymbolNameAccessibilityDescription(name foundation.foundation.INSString, description foundation.foundation.INSString) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ic.class), objc.Sel("imageWithSystemSymbolName:accessibilityDescription:"), name, description)
	return rv
}


// Creates a symbol image with the system symbol name and variable value you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/init(systemSymbolName:variableValue:accessibilityDescription:)
func (ic _ImageClass) ImageWithSystemSymbolNameVariableValueAccessibilityDescription(name foundation.foundation.INSString, value float64, description foundation.foundation.INSString) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ic.class), objc.Sel("imageWithSystemSymbolName:variableValue:accessibilityDescription:"), name, value, description)
	return rv
}







// Returns an array of UTI strings identifying the image types supported by the registered image representation objects, either directly or through a user-installed filter service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/imageTypes
func (ic _ImageClass) ImageTypes() []string {
	rv := objc.Send[[]string](objc.ID(ic.class), objc.Sel("imageTypes"))
	return rv
}

// Returns an array of UTI strings identifying the image types supported directly by the registered image representation objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/imageUnfilteredTypes
func (ic _ImageClass) ImageUnfilteredTypes() []string {
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
func (i_ Image) AddRepresentations(imageReps []ImageRep) {
	objc.Send[objc.ID](i_.ID, objc.Sel("addRepresentations:"), imageReps)
}


// Returns the best representation of the image for the specified rectangle using the provided hints.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/bestRepresentation(for:context:hints:)
func (i_ Image) BestRepresentationForRectContextHints(rect corefoundation.CGRect, referenceContext IGraphicsContext, hints foundation.IDictionary) IImageRep {
	rv := objc.Send[ImageRep](i_.ID, objc.Sel("bestRepresentationForRect:context:hints:"), rect, referenceContext, hints)
	return rv
}


// Returns a Core Graphics image based on the contents of the current image object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/cgImage(forProposedRect:context:hints:)
func (i_ Image) CGImageForProposedRectContextHints(proposedDestRect corefoundation.CGRect, referenceContext IGraphicsContext, hints foundation.IDictionary) ImageRef /* not a class type */ {
	rv := objc.Send[ImageRef](i_.ID, objc.Sel("CGImageForProposedRect:context:hints:"), proposedDestRect, referenceContext, hints)
	return rv
}


// Draws all or part of the image at the specified point in the current coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/draw(at:from:operation:fraction:)
func (i_ Image) DrawAtPointFromRectOperationFraction(point corefoundation.CGPoint, fromRect corefoundation.CGRect, op CompositingOperation, delta float64) {
	objc.Send[objc.ID](i_.ID, objc.Sel("drawAtPoint:fromRect:operation:fraction:"), point, fromRect, op, delta)
}


// Draws the image in the specified rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/draw(in:)
func (i_ Image) DrawInRect(rect corefoundation.CGRect) {
	objc.Send[objc.ID](i_.ID, objc.Sel("drawInRect:"), rect)
}


// Draws all or part of the image in the specified rectangle in the current coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/draw(in:from:operation:fraction:)
func (i_ Image) DrawInRectFromRectOperationFraction(rect corefoundation.CGRect, fromRect corefoundation.CGRect, op CompositingOperation, delta float64) {
	objc.Send[objc.ID](i_.ID, objc.Sel("drawInRect:fromRect:operation:fraction:"), rect, fromRect, op, delta)
}


// Draws all or part of the image in the specified rectangle respecting the hints and the orientation of the current coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/draw(in:from:operation:fraction:respectFlipped:hints:)
func (i_ Image) DrawInRectFromRectOperationFractionRespectFlippedHints(dstSpacePortionRect corefoundation.CGRect, srcSpacePortionRect corefoundation.CGRect, op CompositingOperation, requestedAlpha float64, respectContextIsFlipped bool, hints foundation.IDictionary) {
	objc.Send[objc.ID](i_.ID, objc.Sel("drawInRect:fromRect:operation:fraction:respectFlipped:hints:"), dstSpacePortionRect, srcSpacePortionRect, op, requestedAlpha, respectContextIsFlipped, hints)
}


// Draws the image using the specified image representation object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/drawRepresentation(_:in:)
func (i_ Image) DrawRepresentationInRect(imageRep IImageRep, rect corefoundation.CGRect) bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("drawRepresentation:inRect:"), imageRep, rect)
	return rv
}


// Returns whether the destination rectangle would intersect a non-transparent portion of the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/hitTest(_:withDestinationRect:context:hints:flipped:)
func (i_ Image) HitTestRectWithImageDestinationRectContextHintsFlipped(testRectDestSpace corefoundation.CGRect, imageRectDestSpace corefoundation.CGRect, context IGraphicsContext, hints foundation.IDictionary, flipped bool) bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("hitTestRect:withImageDestinationRect:context:hints:flipped:"), testRectDestSpace, imageRectDestSpace, context, hints, flipped)
	return rv
}


// Returns an object that may be used as the contents of a layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/layerContents(forContentsScale:)
func (i_ Image) LayerContentsForContentsScale(layerContentsScale float64) objc.ID {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("layerContentsForContentsScale:"), layerContentsScale)
	return rv
}


// Returns the name associated with the image, if any.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/name()
func (i_ Image) Name() ImageName {
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
func (i_ Image) RecommendedLayerContentsScale(preferredContentsScale float64) float64 {
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
func (i_ Image) SetName(string_ ImageName) bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("setName:"), string_)
	return rv
}


// Returns a data object that contains TIFF data with the specified compression settings for all of the image representations in the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/tiffRepresentation(using:factor:)
func (i_ Image) TIFFRepresentationUsingCompressionFactor(comp TIFFCompression, factor float32) foundation.Data {
	rv := objc.Send[foundation.Data](i_.ID, objc.Sel("TIFFRepresentationUsingCompression:factor:"), comp, factor)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/withLocale(_:)
func (i_ Image) ImageWithLocale(locale foundation.Locale) IImage {
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
func (i_ Image) AccessibilityDescription() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](i_.ID, objc.Sel("accessibilityDescription"))
	return rv
}


// The image’s accessibility description.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/accessibilityDescription
func (i_ Image) SetAccessibilityDescription(value foundation.foundation.INSString) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setAccessibilityDescription:"), value)
}


// A rectangle that you can use to position the image during layout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/alignmentRect
func (i_ Image) AlignmentRect() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](i_.ID, objc.Sel("alignmentRect"))
	return rv
}


// A rectangle that you can use to position the image during layout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/alignmentRect
func (i_ Image) SetAlignmentRect(value corefoundation.CGRect) {
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
func (i_ Image) CapInsets() foundation.EdgeInsets {
	rv := objc.Send[foundation.EdgeInsets](i_.ID, objc.Sel("capInsets"))
	return rv
}


// The cap insets for the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/capInsets
func (i_ Image) SetCapInsets(value foundation.EdgeInsets) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setCapInsets:"), value)
}


// Returns an array of UTI strings identifying the image types supported by the registered image representation objects, either directly or through a user-installed filter service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/imageTypes
func (i_ Image) ImageTypes() []string {
	rv := objc.Send[[]string](i_.ID, objc.Sel("imageTypes"))
	return rv
}


// Returns an array of UTI strings identifying the image types supported directly by the registered image representation objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/imageUnfilteredTypes
func (i_ Image) ImageUnfilteredTypes() []string {
	rv := objc.Send[[]string](i_.ID, objc.Sel("imageUnfilteredTypes"))
	return rv
}


// A Boolean value that determines whether the image represents a template image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/isTemplate
func (i_ Image) Template() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("template"))
	return rv
}


// A Boolean value that determines whether the image represents a template image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/isTemplate
func (i_ Image) SetTemplate(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setTemplate:"), value)
}


// A Boolean value that indicates whether it is possible to draw an image representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/isValid
func (i_ Image) Valid() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("valid"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/locale
func (i_ Image) Locale() foundation.Locale {
	rv := objc.Send[foundation.Locale](i_.ID, objc.Sel("locale"))
	return rv
}


// A Boolean value that indicates whether image representations whose resolution is an integral multiple of the device resolution are a match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/matchesOnMultipleResolution
func (i_ Image) MatchesOnMultipleResolution() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("matchesOnMultipleResolution"))
	return rv
}


// A Boolean value that indicates whether image representations whose resolution is an integral multiple of the device resolution are a match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/matchesOnMultipleResolution
func (i_ Image) SetMatchesOnMultipleResolution(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMatchesOnMultipleResolution:"), value)
}


// A Boolean value that indicates whether the image matches only on the best fitting axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/matchesOnlyOnBestFittingAxis
func (i_ Image) MatchesOnlyOnBestFittingAxis() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("matchesOnlyOnBestFittingAxis"))
	return rv
}


// A Boolean value that indicates whether the image matches only on the best fitting axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/matchesOnlyOnBestFittingAxis
func (i_ Image) SetMatchesOnlyOnBestFittingAxis(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMatchesOnlyOnBestFittingAxis:"), value)
}


// A Boolean value that indicates whether the image prefers to choose image representations using color-matching or resolution-matching.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/prefersColorMatch
func (i_ Image) PrefersColorMatch() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("prefersColorMatch"))
	return rv
}


// A Boolean value that indicates whether the image prefers to choose image representations using color-matching or resolution-matching.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/prefersColorMatch
func (i_ Image) SetPrefersColorMatch(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPrefersColorMatch:"), value)
}


// An array containing all of the image object’s image representations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/representations
func (i_ Image) Representations() []ImageRep {
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
func (i_ Image) Size() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](i_.ID, objc.Sel("size"))
	return rv
}


// The size of the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/size
func (i_ Image) SetSize(value corefoundation.CGSize) {
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
func (i_ Image) TIFFRepresentation() foundation.foundation.INSData {
	rv := objc.Send[foundation.NSData](i_.ID, objc.Sel("TIFFRepresentation"))
	return rv
}


// A Boolean value that indicates whether EPS representations are preferred when no other representations match the resolution of the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/usesEPSOnResolutionMismatch
func (i_ Image) UsesEPSOnResolutionMismatch() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("usesEPSOnResolutionMismatch"))
	return rv
}


// A Boolean value that indicates whether EPS representations are preferred when no other representations match the resolution of the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/usesEPSOnResolutionMismatch
func (i_ Image) SetUsesEPSOnResolutionMismatch(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setUsesEPSOnResolutionMismatch:"), value)
}


// A Boolean value that determines whether the image represents a template image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/istemplate
func (i_ Image) IsTemplate() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("isTemplate"))
	return rv
}


// A Boolean value that determines whether the image represents a template image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/istemplate
func (i_ Image) SetIsTemplate(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIsTemplate:"), value)
}


// A Boolean value that indicates whether it is possible to draw an image representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/isvalid
func (i_ Image) IsValid() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("isValid"))
	return rv
}


// A Boolean value that indicates whether it is possible to draw an image representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/isvalid
func (i_ Image) SetIsValid(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIsValid:"), value)
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







