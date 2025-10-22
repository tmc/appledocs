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
	BestRepresentationForDevice(deviceDescription objectivec.IObject) ImageRep
	CompositeToPointFromRectOperation(point coregraphics.CGPoint, rect coregraphics.CGRect, operation ICompositingOperation)
	CompositeToPointFromRectOperationFraction(point coregraphics.CGPoint, rect coregraphics.CGRect, operation ICompositingOperation, fraction float64)
	CompositeToPointOperation(point coregraphics.CGPoint, operation ICompositingOperation)
	CompositeToPointOperationFraction(point coregraphics.CGPoint, operation ICompositingOperation, fraction float64)
	DrawAtPointFromRectOperationFraction(point coregraphics.CGPoint, fromRect coregraphics.CGRect, op ICompositingOperation, delta float64)
	DrawInRectFromRectOperationFraction(rect coregraphics.CGRect, fromRect coregraphics.CGRect, op ICompositingOperation, delta float64)
	LayerContentsForContentsScale(layerContentsScale float64) objc.ID
	ImageWithSymbolConfiguration(configuration IImageSymbolConfiguration) Image
	AccessibilityDescription() string
	SetAccessibilityDescription(value string)
	AlignmentRect() coregraphics.CGRect
	SetAlignmentRect(value coregraphics.CGRect)
	CapInsets() unsafe.Pointer
	SetCapInsets(value unsafe.Pointer)
	BackgroundColor() NSColor
	SetBackgroundColor(value IColor)
	CacheMode() unsafe.Pointer
	SetCacheMode(value unsafe.Pointer)
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	IsTemplate() bool
	SetIsTemplate(value bool)
	IsValid() bool
	SetIsValid(value bool)
	Locale() foundation.Locale
	SetLocale(value foundation.ILocale)
	MatchesOnMultipleResolution() bool
	SetMatchesOnMultipleResolution(value bool)
	MatchesOnlyOnBestFittingAxis() bool
	SetMatchesOnlyOnBestFittingAxis(value bool)
	PrefersColorMatch() bool
	SetPrefersColorMatch(value bool)
	Representations() NSImageRep
	SetRepresentations(value IImageRep)
	ResizingMode() unsafe.Pointer
	SetResizingMode(value unsafe.Pointer)
	Size() coregraphics.CGSize
	SetSize(value coregraphics.CGSize)
	SymbolConfiguration() ImageSymbolConfiguration
	SetSymbolConfiguration(value IImageSymbolConfiguration)
	TiffRepresentation() foundation.Data
	SetTiffRepresentation(value foundation.IData)
	UsesEPSOnResolutionMismatch() bool
	SetUsesEPSOnResolutionMismatch(value bool)
	Contents() unsafe.Pointer
	SetContents(value unsafe.Pointer)
	ContentsGravity() unsafe.Pointer
	SetContentsGravity(value unsafe.Pointer)
}

// A high-level interface for manipulating image data.
//
// You use instances of to load existing images, create new images, and draw the resulting image data into your views. Although you use this class predominantly for image-related operations, the class itself knows little about the underlying image data. Instead, it works in conjunction with one or more image representation objects (subclasses of ) to manage and render the image data. For the most part, these interactions are transparent. The class serves many purposes, providing support for the following tasks: Loading images stored on disk or at a specified URL. Drawing images into a view or graphics context. Providing the contents of a object. Creating new images based on a series of captured drawing commands. Producing versions of the image in a different format. The class itself is capable of managing image data in a variety of formats. The specific list of formats is dependent on the version of the operating system but includes many standard formats such as TIFF, JPEG, GIF, PNG, and PDF among others. AppKit manages each format using a specific type of image representation object, whose job is to manage the actual image data. You can get a list of supported formats using the methods described in Determining Supported Types of Images. For more information about how to use image objects in your app, see .
//
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




// Returns the image object associated with the specified name.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/init(named:)
func NewImageNamed(name unsafe.Pointer) Image {
	rv := objc.Send[Image](objc.ID(getImageClass().class), objc.Sel("imageNamed:"), name)
	return rv
}



// Creates a new image using the contents of the provided image.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/init(cgImage:size:)
func NewImageWithCGImageSize(cgImage coregraphics.CGImageRef, size coregraphics.CGSize) Image {
	instance := getImageClass().Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithCGImage:size:"), cgImage, size)
	rv.Autorelease()
	return rv
}



// Initializes and returns an image object using the provided image data.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/init(data:)
func NewImageWithData(data foundation.IData) Image {
	instance := getImageClass().Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithData:"), data)
	rv.Autorelease()
	return rv
}



// Creates and returns an image object whose contents are drawn using the specified block.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/init(size:flipped:drawingHandler:)
func NewImageWithSizeFlippedDrawingHandler(size coregraphics.CGSize, drawingHandlerShouldBeCalledWithFlippedContext bool, drawingHandler unsafe.Pointer) Image {
	rv := objc.Send[Image](objc.ID(getImageClass().class), objc.Sel("imageWithSize:flipped:drawingHandler:"), size, drawingHandlerShouldBeCalledWithFlippedContext, drawingHandler)
	return rv
}



// Creates a symbol image with the symbol name and variable value you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/init(symbolName:variableValue:)
func NewImageWithSymbolNameVariableValue(name string, value float64) Image {
	rv := objc.Send[Image](objc.ID(getImageClass().class), objc.Sel("imageWithSymbolName:variableValue:"), objc.String(name), value)
	return rv
}



// Creates a symbol image with the system symbol name and accessibility description you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/init(systemSymbolName:accessibilityDescription:)
func NewImageWithSystemSymbolNameAccessibilityDescription(name string, description string) Image {
	rv := objc.Send[Image](objc.ID(getImageClass().class), objc.Sel("imageWithSystemSymbolName:accessibilityDescription:"), objc.String(name), objc.String(description))
	return rv
}


// Returns the image object associated with the specified name.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/init(named:)
func (ic _ImageClass) ImageNamed(name unsafe.Pointer) Image {
	rv := objc.Send[Image](objc.ID(ic.class), objc.Sel("imageNamed:"), name)
	return rv
}

// Creates and returns an image object whose contents are drawn using the specified block.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/init(size:flipped:drawingHandler:)
func (ic _ImageClass) ImageWithSizeFlippedDrawingHandler(size coregraphics.CGSize, drawingHandlerShouldBeCalledWithFlippedContext bool, drawingHandler unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ic.class), objc.Sel("imageWithSize:flipped:drawingHandler:"), size, drawingHandlerShouldBeCalledWithFlippedContext, drawingHandler)
	return rv
}

// Creates a symbol image with the symbol name and variable value you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/init(symbolName:variableValue:)
func (ic _ImageClass) ImageWithSymbolNameVariableValue(name string, value float64) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ic.class), objc.Sel("imageWithSymbolName:variableValue:"), objc.String(name), value)
	return rv
}

// Creates a symbol image with the system symbol name and accessibility description you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/init(systemSymbolName:accessibilityDescription:)
func (ic _ImageClass) ImageWithSystemSymbolNameAccessibilityDescription(name string, description string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ic.class), objc.Sel("imageWithSystemSymbolName:accessibilityDescription:"), objc.String(name), objc.String(description))
	return rv
}

// Returns the best representation for the device with the specified characteristics.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/bestRepresentationForDevice:
func (i_ Image) BestRepresentationForDevice(deviceDescription objectivec.IObject) ImageRep {
	rv := objc.Send[ImageRep](i_.ID, objc.Sel("bestRepresentationForDevice:"), deviceDescription)
	return rv
}

// Composites a portion of the image to the specified point in the current coordinate system.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/compositeToPoint:fromRect:operation:
func (i_ Image) CompositeToPointFromRectOperation(point coregraphics.CGPoint, rect coregraphics.CGRect, operation ICompositingOperation) {
	objc.Send[objc.ID](i_.ID, objc.Sel("compositeToPoint:fromRect:operation:"), point, rect, operation)
}

// Composites a portion of the image at the specified opacity to the current coordinate system.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/compositeToPoint:fromRect:operation:fraction:
func (i_ Image) CompositeToPointFromRectOperationFraction(point coregraphics.CGPoint, rect coregraphics.CGRect, operation ICompositingOperation, fraction float64) {
	objc.Send[objc.ID](i_.ID, objc.Sel("compositeToPoint:fromRect:operation:fraction:"), point, rect, operation, fraction)
}

// Composites the entire image to the specified point in the current coordinate system.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/compositeToPoint:operation:
func (i_ Image) CompositeToPointOperation(point coregraphics.CGPoint, operation ICompositingOperation) {
	objc.Send[objc.ID](i_.ID, objc.Sel("compositeToPoint:operation:"), point, operation)
}

// Composites the entire image at the specified opacity in the current coordinate system.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/compositeToPoint:operation:fraction:
func (i_ Image) CompositeToPointOperationFraction(point coregraphics.CGPoint, operation ICompositingOperation, fraction float64) {
	objc.Send[objc.ID](i_.ID, objc.Sel("compositeToPoint:operation:fraction:"), point, operation, fraction)
}

// Draws all or part of the image at the specified point in the current coordinate system.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/draw(at:from:operation:fraction:)
func (i_ Image) DrawAtPointFromRectOperationFraction(point coregraphics.CGPoint, fromRect coregraphics.CGRect, op ICompositingOperation, delta float64) {
	objc.Send[objc.ID](i_.ID, objc.Sel("drawAtPoint:fromRect:operation:fraction:"), point, fromRect, op, delta)
}

// Draws all or part of the image in the specified rectangle in the current coordinate system.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/draw(in:from:operation:fraction:)
func (i_ Image) DrawInRectFromRectOperationFraction(rect coregraphics.CGRect, fromRect coregraphics.CGRect, op ICompositingOperation, delta float64) {
	objc.Send[objc.ID](i_.ID, objc.Sel("drawInRect:fromRect:operation:fraction:"), rect, fromRect, op, delta)
}

// Returns an object that may be used as the contents of a layer.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/layerContents(forContentsScale:)
func (i_ Image) LayerContentsForContentsScale(layerContentsScale float64) objc.ID {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("layerContentsForContentsScale:"), layerContentsScale)
	return rv
}

// Creates a new symbol image with the specified configuration.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/withSymbolConfiguration(_:)
func (i_ Image) ImageWithSymbolConfiguration(configuration IImageSymbolConfiguration) Image {
	rv := objc.Send[Image](i_.ID, objc.Sel("imageWithSymbolConfiguration:"), configuration)
	return rv
}

// The image’s accessibility description.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/accessibilityDescription
func (i_ Image) AccessibilityDescription() string {
	rv := objc.Send[string](i_.ID, objc.Sel("accessibilityDescription"))
	return rv
}


// SetAccessibilityDescription sets the value of the accessibilityDescription property.
// The image’s accessibility description.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/accessibilityDescription
func (i_ Image) SetAccessibilityDescription(value string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setAccessibilityDescription:"), objc.String(value))
}

// A rectangle that you can use to position the image during layout.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/alignmentRect
func (i_ Image) AlignmentRect() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](i_.ID, objc.Sel("alignmentRect"))
	return rv
}


// SetAlignmentRect sets the value of the alignmentRect property.
// A rectangle that you can use to position the image during layout.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/alignmentRect
func (i_ Image) SetAlignmentRect(value coregraphics.CGRect) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setAlignmentRect:"), value)
}

// The cap insets for the image.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/capInsets
func (i_ Image) CapInsets() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("capInsets"))
	return rv
}


// SetCapInsets sets the value of the capInsets property.
// The cap insets for the image.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/capInsets
func (i_ Image) SetCapInsets(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setCapInsets:"), value)
}

// The background color for the image.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/backgroundcolor
func (i_ Image) BackgroundColor() NSColor {
	rv := objc.Send[NSColor](i_.ID, objc.Sel("backgroundColor"))
	return rv
}


// SetBackgroundColor sets the value of the backgroundColor property.
// The background color for the image.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/backgroundcolor
func (i_ Image) SetBackgroundColor(value IColor) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setBackgroundColor:"), value)
}

// The image’s caching mode.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/cachemode-swift.property
func (i_ Image) CacheMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("cacheMode"))
	return rv
}


// SetCacheMode sets the value of the cacheMode property.
// The image’s caching mode.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/cachemode-swift.property
func (i_ Image) SetCacheMode(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setCacheMode:"), value)
}

// The image’s delegate object.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/delegate
func (i_ Image) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The image’s delegate object.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/delegate
func (i_ Image) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDelegate:"), value)
}

// A Boolean value that determines whether the image represents a template image.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/istemplate
func (i_ Image) IsTemplate() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("isTemplate"))
	return rv
}


// SetIsTemplate sets the value of the isTemplate property.
// A Boolean value that determines whether the image represents a template image.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/istemplate
func (i_ Image) SetIsTemplate(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIsTemplate:"), value)
}

// A Boolean value that indicates whether it is possible to draw an image representation.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/isvalid
func (i_ Image) IsValid() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("isValid"))
	return rv
}


// SetIsValid sets the value of the isValid property.
// A Boolean value that indicates whether it is possible to draw an image representation.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/isvalid
func (i_ Image) SetIsValid(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIsValid:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/locale
func (i_ Image) Locale() foundation.Locale {
	rv := objc.Send[foundation.Locale](i_.ID, objc.Sel("locale"))
	return rv
}


// SetLocale sets the value of the locale property.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/locale
func (i_ Image) SetLocale(value foundation.ILocale) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setLocale:"), value)
}

// A Boolean value that indicates whether image representations whose resolution is an integral multiple of the device resolution are a match.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/matchesonmultipleresolution
func (i_ Image) MatchesOnMultipleResolution() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("matchesOnMultipleResolution"))
	return rv
}


// SetMatchesOnMultipleResolution sets the value of the matchesOnMultipleResolution property.
// A Boolean value that indicates whether image representations whose resolution is an integral multiple of the device resolution are a match.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/matchesonmultipleresolution
func (i_ Image) SetMatchesOnMultipleResolution(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMatchesOnMultipleResolution:"), value)
}

// A Boolean value that indicates whether the image matches only on the best fitting axis.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/matchesonlyonbestfittingaxis
func (i_ Image) MatchesOnlyOnBestFittingAxis() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("matchesOnlyOnBestFittingAxis"))
	return rv
}


// SetMatchesOnlyOnBestFittingAxis sets the value of the matchesOnlyOnBestFittingAxis property.
// A Boolean value that indicates whether the image matches only on the best fitting axis.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/matchesonlyonbestfittingaxis
func (i_ Image) SetMatchesOnlyOnBestFittingAxis(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMatchesOnlyOnBestFittingAxis:"), value)
}

// A Boolean value that indicates whether the image prefers to choose image representations using color-matching or resolution-matching.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/preferscolormatch
func (i_ Image) PrefersColorMatch() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("prefersColorMatch"))
	return rv
}


// SetPrefersColorMatch sets the value of the prefersColorMatch property.
// A Boolean value that indicates whether the image prefers to choose image representations using color-matching or resolution-matching.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/preferscolormatch
func (i_ Image) SetPrefersColorMatch(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPrefersColorMatch:"), value)
}

// An array containing all of the image object’s image representations.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/representations
func (i_ Image) Representations() NSImageRep {
	rv := objc.Send[NSImageRep](i_.ID, objc.Sel("representations"))
	return rv
}


// SetRepresentations sets the value of the representations property.
// An array containing all of the image object’s image representations.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/representations
func (i_ Image) SetRepresentations(value IImageRep) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setRepresentations:"), value)
}

// The resizing mode for the image.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/resizingmode-swift.property
func (i_ Image) ResizingMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("resizingMode"))
	return rv
}


// SetResizingMode sets the value of the resizingMode property.
// The resizing mode for the image.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/resizingmode-swift.property
func (i_ Image) SetResizingMode(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setResizingMode:"), value)
}

// The size of the image.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/size
func (i_ Image) Size() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](i_.ID, objc.Sel("size"))
	return rv
}


// SetSize sets the value of the size property.
// The size of the image.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/size
func (i_ Image) SetSize(value coregraphics.CGSize) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSize:"), value)
}

// The configuration details for a symbol image.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/symbolconfiguration-swift.property
func (i_ Image) SymbolConfiguration() ImageSymbolConfiguration {
	rv := objc.Send[ImageSymbolConfiguration](i_.ID, objc.Sel("symbolConfiguration"))
	return rv
}


// SetSymbolConfiguration sets the value of the symbolConfiguration property.
// The configuration details for a symbol image.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/symbolconfiguration-swift.property
func (i_ Image) SetSymbolConfiguration(value IImageSymbolConfiguration) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSymbolConfiguration:"), value)
}

// A data object containing TIFF data for all of the image representations in the image.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/tiffrepresentation
func (i_ Image) TiffRepresentation() foundation.Data {
	rv := objc.Send[foundation.Data](i_.ID, objc.Sel("tiffRepresentation"))
	return rv
}


// SetTiffRepresentation sets the value of the tiffRepresentation property.
// A data object containing TIFF data for all of the image representations in the image.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/tiffrepresentation
func (i_ Image) SetTiffRepresentation(value foundation.IData) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setTiffRepresentation:"), value)
}

// A Boolean value that indicates whether EPS representations are preferred when no other representations match the resolution of the device.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/usesepsonresolutionmismatch
func (i_ Image) UsesEPSOnResolutionMismatch() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("usesEPSOnResolutionMismatch"))
	return rv
}


// SetUsesEPSOnResolutionMismatch sets the value of the usesEPSOnResolutionMismatch property.
// A Boolean value that indicates whether EPS representations are preferred when no other representations match the resolution of the device.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/usesepsonresolutionmismatch
func (i_ Image) SetUsesEPSOnResolutionMismatch(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setUsesEPSOnResolutionMismatch:"), value)
}

// An object that provides the contents of the layer. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/contents
func (i_ Image) Contents() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("contents"))
	return rv
}


// SetContents sets the value of the contents property.
// An object that provides the contents of the layer. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/contents
func (i_ Image) SetContents(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setContents:"), value)
}

// A constant that specifies how the layer’s contents are positioned or scaled within its bounds.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/contentsGravity
func (i_ Image) ContentsGravity() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("contentsGravity"))
	return rv
}


// SetContentsGravity sets the value of the contentsGravity property.
// A constant that specifies how the layer’s contents are positioned or scaled within its bounds.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/contentsGravity
func (i_ Image) SetContentsGravity(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setContentsGravity:"), value)
}


