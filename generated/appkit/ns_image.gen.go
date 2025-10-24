// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
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
	BackgroundColor() objc.IObject /* cross-framework: Color */
	SetBackgroundColor(value objc.IObject /* cross-framework: Color */)
	CacheMode() unsafe.Pointer
	SetCacheMode(value unsafe.Pointer)
	CapInsets() objc.IObject /* cross-framework: EdgeInsets */
	SetCapInsets(value objc.IObject /* cross-framework: EdgeInsets */)
	Delegate() ImageDelegate /* not a class type */
	SetDelegate(value ImageDelegate /* not a class type */)
	IsTemplate() bool
	SetIsTemplate(value bool)
	IsValid() bool
	SetIsValid(value bool)
	Locale() objc.IObject /* cross-framework: Locale */
	SetLocale(value objc.IObject /* cross-framework: Locale */)
	MatchesOnMultipleResolution() bool
	SetMatchesOnMultipleResolution(value bool)
	MatchesOnlyOnBestFittingAxis() bool
	SetMatchesOnlyOnBestFittingAxis(value bool)
	PrefersColorMatch() bool
	SetPrefersColorMatch(value bool)
	Representations() IImageRep
	SetRepresentations(value IImageRep)
	ResizingMode() unsafe.Pointer
	SetResizingMode(value unsafe.Pointer)
	Size() objc.IObject /* cross-framework: Size */
	SetSize(value objc.IObject /* cross-framework: Size */)
	SymbolConfiguration() objc.IObject /* cross-framework: ImageSymbolConfiguration */
	SetSymbolConfiguration(value objc.IObject /* cross-framework: ImageSymbolConfiguration */)
	TiffRepresentation() objc.IObject /* cross-framework: Data */
	SetTiffRepresentation(value objc.IObject /* cross-framework: Data */)
	UsesEPSOnResolutionMismatch() bool
	SetUsesEPSOnResolutionMismatch(value bool)
	Contents() unsafe.Pointer
	SetContents(value unsafe.Pointer)
	ContentsGravity() LayerContentsGravity /* not a class type */
	SetContentsGravity(value LayerContentsGravity /* not a class type */)
	// methods:
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



// The image’s accessibility description.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/accessibilitydescription
func (i_ Image) AccessibilityDescription() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](i_.ID, objc.Sel("accessibilityDescription"))
	return rv
}


// The image’s accessibility description.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/accessibilitydescription
func (i_ Image) SetAccessibilityDescription(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setAccessibilityDescription:"), value)
}


// A rectangle that you can use to position the image during layout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/alignmentrect
func (i_ Image) AlignmentRect() objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[corefoundation.Rect](i_.ID, objc.Sel("alignmentRect"))
	return rv
}


// A rectangle that you can use to position the image during layout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/alignmentrect
func (i_ Image) SetAlignmentRect(value objc.IObject /* cross-framework: Rect */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setAlignmentRect:"), value)
}


// The background color for the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/backgroundcolor
func (i_ Image) BackgroundColor() objc.IObject /* cross-framework: Color */ {
	rv := objc.Send[Color](i_.ID, objc.Sel("backgroundColor"))
	return rv
}


// The background color for the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/backgroundcolor
func (i_ Image) SetBackgroundColor(value objc.IObject /* cross-framework: Color */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setBackgroundColor:"), value)
}


// The image’s caching mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/cachemode-swift.property
func (i_ Image) CacheMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("cacheMode"))
	return rv
}


// The image’s caching mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/cachemode-swift.property
func (i_ Image) SetCacheMode(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setCacheMode:"), value)
}


// The cap insets for the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/capinsets
func (i_ Image) CapInsets() objc.IObject /* cross-framework: EdgeInsets */ {
	rv := objc.Send[foundation.EdgeInsets](i_.ID, objc.Sel("capInsets"))
	return rv
}


// The cap insets for the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/capinsets
func (i_ Image) SetCapInsets(value objc.IObject /* cross-framework: EdgeInsets */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setCapInsets:"), value)
}


// The image’s delegate object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/delegate
func (i_ Image) Delegate() ImageDelegate /* not a class type */ {
	rv := objc.Send[ImageDelegate](i_.ID, objc.Sel("delegate"))
	return rv
}


// The image’s delegate object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/delegate
func (i_ Image) SetDelegate(value ImageDelegate /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDelegate:"), value)
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


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/locale
func (i_ Image) Locale() objc.IObject /* cross-framework: Locale */ {
	rv := objc.Send[foundation.Locale](i_.ID, objc.Sel("locale"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/locale
func (i_ Image) SetLocale(value objc.IObject /* cross-framework: Locale */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setLocale:"), value)
}


// A Boolean value that indicates whether image representations whose resolution is an integral multiple of the device resolution are a match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/matchesonmultipleresolution
func (i_ Image) MatchesOnMultipleResolution() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("matchesOnMultipleResolution"))
	return rv
}


// A Boolean value that indicates whether image representations whose resolution is an integral multiple of the device resolution are a match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/matchesonmultipleresolution
func (i_ Image) SetMatchesOnMultipleResolution(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMatchesOnMultipleResolution:"), value)
}


// A Boolean value that indicates whether the image matches only on the best fitting axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/matchesonlyonbestfittingaxis
func (i_ Image) MatchesOnlyOnBestFittingAxis() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("matchesOnlyOnBestFittingAxis"))
	return rv
}


// A Boolean value that indicates whether the image matches only on the best fitting axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/matchesonlyonbestfittingaxis
func (i_ Image) SetMatchesOnlyOnBestFittingAxis(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMatchesOnlyOnBestFittingAxis:"), value)
}


// A Boolean value that indicates whether the image prefers to choose image representations using color-matching or resolution-matching.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/preferscolormatch
func (i_ Image) PrefersColorMatch() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("prefersColorMatch"))
	return rv
}


// A Boolean value that indicates whether the image prefers to choose image representations using color-matching or resolution-matching.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/preferscolormatch
func (i_ Image) SetPrefersColorMatch(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPrefersColorMatch:"), value)
}


// An array containing all of the image object’s image representations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/representations
func (i_ Image) Representations() IImageRep {
	rv := objc.Send[ImageRep](i_.ID, objc.Sel("representations"))
	return rv
}


// An array containing all of the image object’s image representations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/representations
func (i_ Image) SetRepresentations(value IImageRep) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setRepresentations:"), value)
}


// The resizing mode for the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/resizingmode-swift.property
func (i_ Image) ResizingMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("resizingMode"))
	return rv
}


// The resizing mode for the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/resizingmode-swift.property
func (i_ Image) SetResizingMode(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setResizingMode:"), value)
}


// The size of the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/size
func (i_ Image) Size() objc.IObject /* cross-framework: Size */ {
	rv := objc.Send[corefoundation.Size](i_.ID, objc.Sel("size"))
	return rv
}


// The size of the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/size
func (i_ Image) SetSize(value objc.IObject /* cross-framework: Size */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSize:"), value)
}


// The configuration details for a symbol image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/symbolconfiguration-swift.property
func (i_ Image) SymbolConfiguration() objc.IObject /* cross-framework: ImageSymbolConfiguration */ {
	rv := objc.Send[ImageSymbolConfiguration](i_.ID, objc.Sel("symbolConfiguration"))
	return rv
}


// The configuration details for a symbol image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/symbolconfiguration-swift.property
func (i_ Image) SetSymbolConfiguration(value objc.IObject /* cross-framework: ImageSymbolConfiguration */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSymbolConfiguration:"), value)
}


// A data object containing TIFF data for all of the image representations in the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/tiffrepresentation
func (i_ Image) TiffRepresentation() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](i_.ID, objc.Sel("tiffRepresentation"))
	return rv
}


// A data object containing TIFF data for all of the image representations in the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/tiffrepresentation
func (i_ Image) SetTiffRepresentation(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setTiffRepresentation:"), value)
}


// A Boolean value that indicates whether EPS representations are preferred when no other representations match the resolution of the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/usesepsonresolutionmismatch
func (i_ Image) UsesEPSOnResolutionMismatch() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("usesEPSOnResolutionMismatch"))
	return rv
}


// A Boolean value that indicates whether EPS representations are preferred when no other representations match the resolution of the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimage/usesepsonresolutionmismatch
func (i_ Image) SetUsesEPSOnResolutionMismatch(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setUsesEPSOnResolutionMismatch:"), value)
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



