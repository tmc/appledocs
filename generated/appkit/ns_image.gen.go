// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/vision"
)

/* debug [class.gen.go]: Generating class NSImage */


/* debug [class_header]: Header for NSImage */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Image */
// An interface definition for the [Image] class.
type IImage interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Image */
	// properties:
	AccessibilityDescription() objc.IObject /* cross-framework: NSString */
	SetAccessibilityDescription(value objc.IObject /* cross-framework: NSString */)
	AlignmentRect() Rect /* not a class type */
	SetAlignmentRect(value Rect /* not a class type */)
	BackgroundColor() IColor
	SetBackgroundColor(value IColor)
	CacheMode() ImageCacheMode
	SetCacheMode(value ImageCacheMode)
	CapInsets() foundation.EdgeInsets
	SetCapInsets(value foundation.EdgeInsets)
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
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
	Size() Size /* not a class type */
	SetSize(value Size /* not a class type */)
	SymbolConfiguration() IImageSymbolConfiguration
	TIFFRepresentation() objc.IObject /* cross-framework: NSData */
	UsesEPSOnResolutionMismatch() bool
	SetUsesEPSOnResolutionMismatch(value bool)
	IsTemplate() bool
	SetIsTemplate(value bool)
	IsValid() bool
	SetIsValid(value bool)
	ContentsGravity() LayerContentsGravity /* not a class type */
	SetContentsGravity(value LayerContentsGravity /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Image */
	// methods:
	AddRepresentation(imageRep IImageRep)
	AddRepresentations(imageReps []ImageRep)
	BestRepresentationForRectContextHints(rect Rect /* not a class type */, referenceContext IGraphicsContext, hints foundation.IDictionary) IImageRep
	CGImageForProposedRectContextHints(proposedDestRect Rect /* not a class type */, referenceContext IGraphicsContext, hints foundation.IDictionary) ImageRef /* not a class type */
	DrawAtPointFromRectOperationFraction(point vision.Point, fromRect Rect /* not a class type */, op CompositingOperation, delta float64)
	DrawInRect(rect Rect /* not a class type */)
	DrawInRectFromRectOperationFraction(rect Rect /* not a class type */, fromRect Rect /* not a class type */, op CompositingOperation, delta float64)
	DrawInRectFromRectOperationFractionRespectFlippedHints(dstSpacePortionRect Rect /* not a class type */, srcSpacePortionRect Rect /* not a class type */, op CompositingOperation, requestedAlpha float64, respectContextIsFlipped bool, hints foundation.IDictionary)
	DrawRepresentationInRect(imageRep IImageRep, rect Rect /* not a class type */) bool
	HitTestRectWithImageDestinationRectContextHintsFlipped(testRectDestSpace Rect /* not a class type */, imageRectDestSpace Rect /* not a class type */, context IGraphicsContext, hints foundation.IDictionary, flipped bool) bool
	LayerContentsForContentsScale(layerContentsScale float64) objc.ID
	Name() ImageName /* typedef */
	Recache()
	RecommendedLayerContentsScale(preferredContentsScale float64) float64
	RemoveRepresentation(imageRep IImageRep)
	SetName(string_ ImageName /* typedef */) bool
	TIFFRepresentationUsingCompressionFactor(comp TIFFCompression, factor float32) foundation.Data
	ImageWithLocale(locale foundation.Locale) IImage
	ImageWithSymbolConfiguration(configuration IImageSymbolConfiguration) IImage
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Image */

// Initializes and returns an image object using the specified file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/init(byReferencingFile:)
func NewImageByReferencingFile(fileName objc.IObject /* cross-framework: NSString */) Image {
	instance := getImageClass().Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initByReferencingFile:"), fileName)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageByReferencingFile */


// Initializes and returns an image object using the specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/init(byReferencing:)
func NewImageByReferencingURL(url objc.IObject /* cross-framework: NSURL */) Image {
	instance := getImageClass().Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initByReferencingURL:"), url)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageByReferencingURL */


// Returns the image object associated with the specified name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/init(named:)
func NewImageNamed(name ImageName /* typedef */) Image {
	rv := objc.Send[Image](objc.ID(getImageClass().class), objc.Sel("imageNamed:"), name)
	return rv
}/* debug [class_init_methods/constructor]: NewImageNamed */


// Creates a new image using the contents of the provided image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/init(cgImage:size:)
func NewImageWithCGImageSize(cgImage ImageRef /* not a class type */, size Size /* not a class type */) Image {
	instance := getImageClass().Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithCGImage:size:"), cgImage, size)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageWithCGImageSize */


// Initializes and returns an image object from data in an unarchiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/init(coder:)
func NewImageWithCoder(coder foundation.Coder) Image {
	instance := getImageClass().Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageWithCoder */


// Initializes and returns an image object with the contents of the specified file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/init(contentsOfFile:)
func NewImageWithContentsOfFile(fileName objc.IObject /* cross-framework: NSString */) Image {
	instance := getImageClass().Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithContentsOfFile:"), fileName)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageWithContentsOfFile */


// Initializes and returns an image object with the contents of the specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/init(contentsOf:)
func NewImageWithContentsOfURL(url objc.IObject /* cross-framework: NSURL */) Image {
	instance := getImageClass().Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithContentsOfURL:"), url)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageWithContentsOfURL */


// Initializes and returns an image object using the provided image data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/init(data:)
func NewImageWithData(data objc.IObject /* cross-framework: NSData */) Image {
	instance := getImageClass().Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithData:"), data)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageWithData */


// Initializes and returns an image object using the provided image data and ignoring the EXIF orientation tags.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/init(dataIgnoringOrientation:)
func NewImageWithDataIgnoringOrientation(data objc.IObject /* cross-framework: NSData */) Image {
	instance := getImageClass().Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithDataIgnoringOrientation:"), data)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageWithDataIgnoringOrientation */


// Initializes the image object with a Carbon-style icon resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/init(iconRef:)
func NewImageWithIconRef(iconRef objectivec.IObject) Image {
	instance := getImageClass().Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithIconRef:"), iconRef)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageWithIconRef */


// Initializes and returns an image object with data from the specified pasteboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/init(pasteboard:)
func NewImageWithPasteboard(pasteboard IPasteboard) Image {
	instance := getImageClass().Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithPasteboard:"), pasteboard)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageWithPasteboard */


// Initializes and returns an image object with the specified dimensions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/init(size:)
func NewImageWithSize(size Size /* not a class type */) Image {
	instance := getImageClass().Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithSize:"), size)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageWithSize */


// Creates and returns an image object whose contents are drawn using the specified block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/init(size:flipped:drawingHandler:)
func NewImageWithSizeFlippedDrawingHandler(size Size /* not a class type */, drawingHandlerShouldBeCalledWithFlippedContext bool, drawingHandler unsafe.Pointer) Image {
	rv := objc.Send[Image](objc.ID(getImageClass().class), objc.Sel("imageWithSize:flipped:drawingHandler:"), size, drawingHandlerShouldBeCalledWithFlippedContext, drawingHandler)
	return rv
}/* debug [class_init_methods/constructor]: NewImageWithSizeFlippedDrawingHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/init(symbolName:bundle:variableValue:)
func NewImageWithSymbolNameBundleVariableValue(name objc.IObject /* cross-framework: NSString */, bundle foundation.Bundle, value float64) Image {
	rv := objc.Send[Image](objc.ID(getImageClass().class), objc.Sel("imageWithSymbolName:bundle:variableValue:"), name, bundle, value)
	return rv
}/* debug [class_init_methods/constructor]: NewImageWithSymbolNameBundleVariableValue */


// Creates a symbol image with the symbol name and variable value you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/init(symbolName:variableValue:)
func NewImageWithSymbolNameVariableValue(name objc.IObject /* cross-framework: NSString */, value float64) Image {
	rv := objc.Send[Image](objc.ID(getImageClass().class), objc.Sel("imageWithSymbolName:variableValue:"), name, value)
	return rv
}/* debug [class_init_methods/constructor]: NewImageWithSymbolNameVariableValue */


// Creates a symbol image with the system symbol name and accessibility description you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/init(systemSymbolName:accessibilityDescription:)
func NewImageWithSystemSymbolNameAccessibilityDescription(name objc.IObject /* cross-framework: NSString */, description objc.IObject /* cross-framework: NSString */) Image {
	rv := objc.Send[Image](objc.ID(getImageClass().class), objc.Sel("imageWithSystemSymbolName:accessibilityDescription:"), name, description)
	return rv
}/* debug [class_init_methods/constructor]: NewImageWithSystemSymbolNameAccessibilityDescription */


// Creates a symbol image with the system symbol name and variable value you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/init(systemSymbolName:variableValue:accessibilityDescription:)
func NewImageWithSystemSymbolNameVariableValueAccessibilityDescription(name objc.IObject /* cross-framework: NSString */, value float64, description objc.IObject /* cross-framework: NSString */) Image {
	rv := objc.Send[Image](objc.ID(getImageClass().class), objc.Sel("imageWithSystemSymbolName:variableValue:accessibilityDescription:"), name, value, description)
	return rv
}/* debug [class_init_methods/constructor]: NewImageWithSystemSymbolNameVariableValueAccessibilityDescription */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Image */

// Tests whether the image can create an instance of itself using pasteboard data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/canInit(with:)
func (ic _ImageClass) CanInitWithPasteboard(pasteboard IPasteboard) bool {
	rv := objc.Send[bool](objc.ID(ic.class), objc.Sel("canInitWithPasteboard:"), pasteboard)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CanInitWithPasteboard) */


// Returns an array of strings identifying the image types supported by the registered image representation objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/imageFileTypes()
func (ic _ImageClass) ImageFileTypes() []string {
	rv := objc.Send[[]string](objc.ID(ic.class), objc.Sel("imageFileTypes"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ImageFileTypes) */


// Returns an array of strings identifying the pasteboard types supported directly by the registered image representation objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/imagePasteboardTypes()
func (ic _ImageClass) ImagePasteboardTypes() []string {
	rv := objc.Send[[]string](objc.ID(ic.class), objc.Sel("imagePasteboardTypes"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ImagePasteboardTypes) */


// Returns an array of strings identifying the file types supported directly by the registered image representation objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/imageUnfilteredFileTypes()
func (ic _ImageClass) ImageUnfilteredFileTypes() []string {
	rv := objc.Send[[]string](objc.ID(ic.class), objc.Sel("imageUnfilteredFileTypes"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ImageUnfilteredFileTypes) */


// Returns an array of strings identifying the pasteboard types supported directly by the registered image representation objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/imageUnfilteredPasteboardTypes()
func (ic _ImageClass) ImageUnfilteredPasteboardTypes() []string {
	rv := objc.Send[[]string](objc.ID(ic.class), objc.Sel("imageUnfilteredPasteboardTypes"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ImageUnfilteredPasteboardTypes) */


// Returns the image object associated with the specified name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/init(named:)
func (ic _ImageClass) ImageNamed(name ImageName /* typedef */) IImage {
	rv := objc.Send[Image](objc.ID(ic.class), objc.Sel("imageNamed:"), name)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ImageNamed) */


// Creates and returns an image object whose contents are drawn using the specified block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/init(size:flipped:drawingHandler:)
func (ic _ImageClass) ImageWithSizeFlippedDrawingHandler(size Size /* not a class type */, drawingHandlerShouldBeCalledWithFlippedContext bool, drawingHandler unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ic.class), objc.Sel("imageWithSize:flipped:drawingHandler:"), size, drawingHandlerShouldBeCalledWithFlippedContext, drawingHandler)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ImageWithSizeFlippedDrawingHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/init(symbolName:bundle:variableValue:)
func (ic _ImageClass) ImageWithSymbolNameBundleVariableValue(name objc.IObject /* cross-framework: NSString */, bundle foundation.Bundle, value float64) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ic.class), objc.Sel("imageWithSymbolName:bundle:variableValue:"), name, bundle, value)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ImageWithSymbolNameBundleVariableValue) */


// Creates a symbol image with the symbol name and variable value you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/init(symbolName:variableValue:)
func (ic _ImageClass) ImageWithSymbolNameVariableValue(name objc.IObject /* cross-framework: NSString */, value float64) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ic.class), objc.Sel("imageWithSymbolName:variableValue:"), name, value)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ImageWithSymbolNameVariableValue) */


// Creates a symbol image with the system symbol name and accessibility description you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/init(systemSymbolName:accessibilityDescription:)
func (ic _ImageClass) ImageWithSystemSymbolNameAccessibilityDescription(name objc.IObject /* cross-framework: NSString */, description objc.IObject /* cross-framework: NSString */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ic.class), objc.Sel("imageWithSystemSymbolName:accessibilityDescription:"), name, description)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ImageWithSystemSymbolNameAccessibilityDescription) */


// Creates a symbol image with the system symbol name and variable value you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/init(systemSymbolName:variableValue:accessibilityDescription:)
func (ic _ImageClass) ImageWithSystemSymbolNameVariableValueAccessibilityDescription(name objc.IObject /* cross-framework: NSString */, value float64, description objc.IObject /* cross-framework: NSString */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ic.class), objc.Sel("imageWithSystemSymbolName:variableValue:accessibilityDescription:"), name, value, description)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ImageWithSystemSymbolNameVariableValueAccessibilityDescription) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Image */

// Returns an array of UTI strings identifying the image types supported by the registered image representation objects, either directly or through a user-installed filter service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/imageTypes
func (ic _ImageClass) ImageTypes() []string {
	rv := objc.Send[[]string](objc.ID(ic.class), objc.Sel("imageTypes"))
	return rv
}/* debug [class_properties_class/property]: imageTypes */

// Returns an array of UTI strings identifying the image types supported directly by the registered image representation objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/imageUnfilteredTypes
func (ic _ImageClass) ImageUnfilteredTypes() []string {
	rv := objc.Send[[]string](objc.ID(ic.class), objc.Sel("imageUnfilteredTypes"))
	return rv
}/* debug [class_properties_class/property]: imageUnfilteredTypes */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Image */

// Adds the specified image representation object to the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/addRepresentation(_:)
func (i_ Image) AddRepresentation(imageRep IImageRep) {
	objc.Send[objc.ID](i_.ID, objc.Sel("addRepresentation:"), imageRep)
}/* debug [instance_methods/method]: AddRepresentation */


// Adds an array of image representation objects to the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/addRepresentations(_:)
func (i_ Image) AddRepresentations(imageReps []ImageRep) {
	objc.Send[objc.ID](i_.ID, objc.Sel("addRepresentations:"), imageReps)
}/* debug [instance_methods/method]: AddRepresentations */


// Returns the best representation of the image for the specified rectangle using the provided hints.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/bestRepresentation(for:context:hints:)
func (i_ Image) BestRepresentationForRectContextHints(rect Rect /* not a class type */, referenceContext IGraphicsContext, hints foundation.IDictionary) IImageRep {
	rv := objc.Send[ImageRep](i_.ID, objc.Sel("bestRepresentationForRect:context:hints:"), rect, referenceContext, hints)
	return rv
}/* debug [instance_methods/method]: BestRepresentationForRectContextHints */


// Returns a Core Graphics image based on the contents of the current image object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/cgImage(forProposedRect:context:hints:)
func (i_ Image) CGImageForProposedRectContextHints(proposedDestRect Rect /* not a class type */, referenceContext IGraphicsContext, hints foundation.IDictionary) ImageRef /* not a class type */ {
	rv := objc.Send[ImageRef](i_.ID, objc.Sel("CGImageForProposedRect:context:hints:"), proposedDestRect, referenceContext, hints)
	return rv
}/* debug [instance_methods/method]: CGImageForProposedRectContextHints */


// Draws all or part of the image at the specified point in the current coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/draw(at:from:operation:fraction:)
func (i_ Image) DrawAtPointFromRectOperationFraction(point vision.Point, fromRect Rect /* not a class type */, op CompositingOperation, delta float64) {
	objc.Send[objc.ID](i_.ID, objc.Sel("drawAtPoint:fromRect:operation:fraction:"), point, fromRect, op, delta)
}/* debug [instance_methods/method]: DrawAtPointFromRectOperationFraction */


// Draws the image in the specified rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/draw(in:)
func (i_ Image) DrawInRect(rect Rect /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("drawInRect:"), rect)
}/* debug [instance_methods/method]: DrawInRect */


// Draws all or part of the image in the specified rectangle in the current coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/draw(in:from:operation:fraction:)
func (i_ Image) DrawInRectFromRectOperationFraction(rect Rect /* not a class type */, fromRect Rect /* not a class type */, op CompositingOperation, delta float64) {
	objc.Send[objc.ID](i_.ID, objc.Sel("drawInRect:fromRect:operation:fraction:"), rect, fromRect, op, delta)
}/* debug [instance_methods/method]: DrawInRectFromRectOperationFraction */


// Draws all or part of the image in the specified rectangle respecting the hints and the orientation of the current coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/draw(in:from:operation:fraction:respectFlipped:hints:)
func (i_ Image) DrawInRectFromRectOperationFractionRespectFlippedHints(dstSpacePortionRect Rect /* not a class type */, srcSpacePortionRect Rect /* not a class type */, op CompositingOperation, requestedAlpha float64, respectContextIsFlipped bool, hints foundation.IDictionary) {
	objc.Send[objc.ID](i_.ID, objc.Sel("drawInRect:fromRect:operation:fraction:respectFlipped:hints:"), dstSpacePortionRect, srcSpacePortionRect, op, requestedAlpha, respectContextIsFlipped, hints)
}/* debug [instance_methods/method]: DrawInRectFromRectOperationFractionRespectFlippedHints */


// Draws the image using the specified image representation object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/drawRepresentation(_:in:)
func (i_ Image) DrawRepresentationInRect(imageRep IImageRep, rect Rect /* not a class type */) bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("drawRepresentation:inRect:"), imageRep, rect)
	return rv
}/* debug [instance_methods/method]: DrawRepresentationInRect */


// Returns whether the destination rectangle would intersect a non-transparent portion of the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/hitTest(_:withDestinationRect:context:hints:flipped:)
func (i_ Image) HitTestRectWithImageDestinationRectContextHintsFlipped(testRectDestSpace Rect /* not a class type */, imageRectDestSpace Rect /* not a class type */, context IGraphicsContext, hints foundation.IDictionary, flipped bool) bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("hitTestRect:withImageDestinationRect:context:hints:flipped:"), testRectDestSpace, imageRectDestSpace, context, hints, flipped)
	return rv
}/* debug [instance_methods/method]: HitTestRectWithImageDestinationRectContextHintsFlipped */


// Returns an object that may be used as the contents of a layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/layerContents(forContentsScale:)
func (i_ Image) LayerContentsForContentsScale(layerContentsScale float64) objc.ID {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("layerContentsForContentsScale:"), layerContentsScale)
	return rv
}/* debug [instance_methods/method]: LayerContentsForContentsScale */


// Returns the name associated with the image, if any.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/name()
func (i_ Image) Name() ImageName /* typedef */ {
	rv := objc.Send[foundation.NSString](i_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_methods/method]: Name */


// Invalidates and frees offscreen caches of all image representations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/recache()
func (i_ Image) Recache() {
	objc.Send[objc.ID](i_.ID, objc.Sel("recache"))
}/* debug [instance_methods/method]: Recache */


// Returns the recommended layer contents scale for this image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/recommendedLayerContentsScale(_:)
func (i_ Image) RecommendedLayerContentsScale(preferredContentsScale float64) float64 {
	rv := objc.Send[float64](i_.ID, objc.Sel("recommendedLayerContentsScale:"), preferredContentsScale)
	return rv
}/* debug [instance_methods/method]: RecommendedLayerContentsScale */


// Removes and releases the specified image representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/removeRepresentation(_:)
func (i_ Image) RemoveRepresentation(imageRep IImageRep) {
	objc.Send[objc.ID](i_.ID, objc.Sel("removeRepresentation:"), imageRep)
}/* debug [instance_methods/method]: RemoveRepresentation */


// Registers the image object under the specified name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/setName(_:)
func (i_ Image) SetName(string_ ImageName /* typedef */) bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("setName:"), string_)
	return rv
}/* debug [instance_methods/method]: SetName */


// Returns a data object that contains TIFF data with the specified compression settings for all of the image representations in the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/tiffRepresentation(using:factor:)
func (i_ Image) TIFFRepresentationUsingCompressionFactor(comp TIFFCompression, factor float32) foundation.Data {
	rv := objc.Send[foundation.Data](i_.ID, objc.Sel("TIFFRepresentationUsingCompression:factor:"), comp, factor)
	return rv
}/* debug [instance_methods/method]: TIFFRepresentationUsingCompressionFactor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/withLocale(_:)
func (i_ Image) ImageWithLocale(locale foundation.Locale) IImage {
	rv := objc.Send[Image](i_.ID, objc.Sel("imageWithLocale:"), locale)
	return rv
}/* debug [instance_methods/method]: ImageWithLocale */


// Creates a new symbol image with the specified configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/withSymbolConfiguration(_:)
func (i_ Image) ImageWithSymbolConfiguration(configuration IImageSymbolConfiguration) IImage {
	rv := objc.Send[Image](i_.ID, objc.Sel("imageWithSymbolConfiguration:"), configuration)
	return rv
}/* debug [instance_methods/method]: ImageWithSymbolConfiguration */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Image */

// The image’s accessibility description.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/accessibilityDescription
func (i_ Image) AccessibilityDescription() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](i_.ID, objc.Sel("accessibilityDescription"))
	return rv
}/* debug [instance_properties/getter]: accessibilityDescription */


// The image’s accessibility description.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/accessibilityDescription
func (i_ Image) SetAccessibilityDescription(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setAccessibilityDescription:"), value)
}/* debug [instance_properties/setter]: accessibilityDescription */


// A rectangle that you can use to position the image during layout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/alignmentRect
func (i_ Image) AlignmentRect() Rect /* not a class type */ {
	rv := objc.Send[Rect](i_.ID, objc.Sel("alignmentRect"))
	return rv
}/* debug [instance_properties/getter]: alignmentRect */


// A rectangle that you can use to position the image during layout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/alignmentRect
func (i_ Image) SetAlignmentRect(value Rect /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setAlignmentRect:"), value)
}/* debug [instance_properties/setter]: alignmentRect */


// The background color for the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/backgroundColor
func (i_ Image) BackgroundColor() IColor {
	rv := objc.Send[Color](i_.ID, objc.Sel("backgroundColor"))
	return rv
}/* debug [instance_properties/getter]: backgroundColor */


// The background color for the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/backgroundColor
func (i_ Image) SetBackgroundColor(value IColor) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setBackgroundColor:"), value)
}/* debug [instance_properties/setter]: backgroundColor */


// The image’s caching mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/cacheMode-swift.property
func (i_ Image) CacheMode() ImageCacheMode {
	rv := objc.Send[ImageCacheMode](i_.ID, objc.Sel("cacheMode"))
	return rv
}/* debug [instance_properties/getter]: cacheMode */


// The image’s caching mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/cacheMode-swift.property
func (i_ Image) SetCacheMode(value ImageCacheMode) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setCacheMode:"), value)
}/* debug [instance_properties/setter]: cacheMode */


// The cap insets for the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/capInsets
func (i_ Image) CapInsets() foundation.EdgeInsets {
	rv := objc.Send[foundation.EdgeInsets](i_.ID, objc.Sel("capInsets"))
	return rv
}/* debug [instance_properties/getter]: capInsets */


// The cap insets for the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/capInsets
func (i_ Image) SetCapInsets(value foundation.EdgeInsets) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setCapInsets:"), value)
}/* debug [instance_properties/setter]: capInsets */


// The image’s delegate object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/delegate
func (i_ Image) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The image’s delegate object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/delegate
func (i_ Image) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// Returns an array of UTI strings identifying the image types supported by the registered image representation objects, either directly or through a user-installed filter service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/imageTypes
func (i_ Image) ImageTypes() []string {
	rv := objc.Send[[]string](i_.ID, objc.Sel("imageTypes"))
	return rv
}/* debug [instance_properties/getter]: imageTypes */


// Returns an array of UTI strings identifying the image types supported directly by the registered image representation objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/imageUnfilteredTypes
func (i_ Image) ImageUnfilteredTypes() []string {
	rv := objc.Send[[]string](i_.ID, objc.Sel("imageUnfilteredTypes"))
	return rv
}/* debug [instance_properties/getter]: imageUnfilteredTypes */


// A Boolean value that determines whether the image represents a template image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/isTemplate
func (i_ Image) Template() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("template"))
	return rv
}/* debug [instance_properties/getter]: template */


// A Boolean value that determines whether the image represents a template image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/isTemplate
func (i_ Image) SetTemplate(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setTemplate:"), value)
}/* debug [instance_properties/setter]: template */


// A Boolean value that indicates whether it is possible to draw an image representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/isValid
func (i_ Image) Valid() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("valid"))
	return rv
}/* debug [instance_properties/getter]: valid */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/locale
func (i_ Image) Locale() foundation.Locale {
	rv := objc.Send[foundation.Locale](i_.ID, objc.Sel("locale"))
	return rv
}/* debug [instance_properties/getter]: locale */


// A Boolean value that indicates whether image representations whose resolution is an integral multiple of the device resolution are a match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/matchesOnMultipleResolution
func (i_ Image) MatchesOnMultipleResolution() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("matchesOnMultipleResolution"))
	return rv
}/* debug [instance_properties/getter]: matchesOnMultipleResolution */


// A Boolean value that indicates whether image representations whose resolution is an integral multiple of the device resolution are a match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/matchesOnMultipleResolution
func (i_ Image) SetMatchesOnMultipleResolution(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMatchesOnMultipleResolution:"), value)
}/* debug [instance_properties/setter]: matchesOnMultipleResolution */


// A Boolean value that indicates whether the image matches only on the best fitting axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/matchesOnlyOnBestFittingAxis
func (i_ Image) MatchesOnlyOnBestFittingAxis() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("matchesOnlyOnBestFittingAxis"))
	return rv
}/* debug [instance_properties/getter]: matchesOnlyOnBestFittingAxis */


// A Boolean value that indicates whether the image matches only on the best fitting axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/matchesOnlyOnBestFittingAxis
func (i_ Image) SetMatchesOnlyOnBestFittingAxis(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMatchesOnlyOnBestFittingAxis:"), value)
}/* debug [instance_properties/setter]: matchesOnlyOnBestFittingAxis */


// A Boolean value that indicates whether the image prefers to choose image representations using color-matching or resolution-matching.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/prefersColorMatch
func (i_ Image) PrefersColorMatch() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("prefersColorMatch"))
	return rv
}/* debug [instance_properties/getter]: prefersColorMatch */


// A Boolean value that indicates whether the image prefers to choose image representations using color-matching or resolution-matching.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/prefersColorMatch
func (i_ Image) SetPrefersColorMatch(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPrefersColorMatch:"), value)
}/* debug [instance_properties/setter]: prefersColorMatch */


// An array containing all of the image object’s image representations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/representations
func (i_ Image) Representations() []ImageRep {
	rv := objc.Send[[]ImageRep](i_.ID, objc.Sel("representations"))
	return rv
}/* debug [instance_properties/getter]: representations */


// The resizing mode for the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/resizingMode-swift.property
func (i_ Image) ResizingMode() ImageResizingMode {
	rv := objc.Send[ImageResizingMode](i_.ID, objc.Sel("resizingMode"))
	return rv
}/* debug [instance_properties/getter]: resizingMode */


// The resizing mode for the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/resizingMode-swift.property
func (i_ Image) SetResizingMode(value ImageResizingMode) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setResizingMode:"), value)
}/* debug [instance_properties/setter]: resizingMode */


// The size of the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/size
func (i_ Image) Size() Size /* not a class type */ {
	rv := objc.Send[Size](i_.ID, objc.Sel("size"))
	return rv
}/* debug [instance_properties/getter]: size */


// The size of the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/size
func (i_ Image) SetSize(value Size /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSize:"), value)
}/* debug [instance_properties/setter]: size */


// The configuration details for a symbol image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/symbolConfiguration-swift.property
func (i_ Image) SymbolConfiguration() IImageSymbolConfiguration {
	rv := objc.Send[ImageSymbolConfiguration](i_.ID, objc.Sel("symbolConfiguration"))
	return rv
}/* debug [instance_properties/getter]: symbolConfiguration */


// A data object containing TIFF data for all of the image representations in the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/tiffRepresentation
func (i_ Image) TIFFRepresentation() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](i_.ID, objc.Sel("TIFFRepresentation"))
	return rv
}/* debug [instance_properties/getter]: TIFFRepresentation */


// A Boolean value that indicates whether EPS representations are preferred when no other representations match the resolution of the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/usesEPSOnResolutionMismatch
func (i_ Image) UsesEPSOnResolutionMismatch() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("usesEPSOnResolutionMismatch"))
	return rv
}/* debug [instance_properties/getter]: usesEPSOnResolutionMismatch */


// A Boolean value that indicates whether EPS representations are preferred when no other representations match the resolution of the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/usesEPSOnResolutionMismatch
func (i_ Image) SetUsesEPSOnResolutionMismatch(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setUsesEPSOnResolutionMismatch:"), value)
}/* debug [instance_properties/setter]: usesEPSOnResolutionMismatch */


// A Boolean value that determines whether the image represents a template image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/istemplate
func (i_ Image) IsTemplate() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("isTemplate"))
	return rv
}/* debug [instance_properties/getter]: isTemplate */


// A Boolean value that determines whether the image represents a template image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/istemplate
func (i_ Image) SetIsTemplate(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIsTemplate:"), value)
}/* debug [instance_properties/setter]: isTemplate */


// A Boolean value that indicates whether it is possible to draw an image representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/isvalid
func (i_ Image) IsValid() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("isValid"))
	return rv
}/* debug [instance_properties/getter]: isValid */


// A Boolean value that indicates whether it is possible to draw an image representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/isvalid
func (i_ Image) SetIsValid(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIsValid:"), value)
}/* debug [instance_properties/setter]: isValid */


// A constant that specifies how the layer’s contents are positioned or scaled within its bounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/contentsGravity
func (i_ Image) ContentsGravity() LayerContentsGravity /* not a class type */ {
	rv := objc.Send[LayerContentsGravity](i_.ID, objc.Sel("contentsGravity"))
	return rv
}/* debug [instance_properties/getter]: contentsGravity */


// A constant that specifies how the layer’s contents are positioned or scaled within its bounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/contentsGravity
func (i_ Image) SetContentsGravity(value LayerContentsGravity /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setContentsGravity:"), value)
}/* debug [instance_properties/setter]: contentsGravity */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSImage */


