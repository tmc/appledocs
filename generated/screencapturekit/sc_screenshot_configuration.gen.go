// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/uniformtypeidentifiers"
)

/* debug [class.gen.go]: Generating class SCScreenshotConfiguration */


/* debug [class_header]: Header for SCScreenshotConfiguration */
// The class instance for the [ScreenshotConfiguration] class.
var (
	ScreenshotConfigurationClass     _ScreenshotConfigurationClass
	ScreenshotConfigurationClassOnce sync.Once
)

func getScreenshotConfigurationClass() _ScreenshotConfigurationClass {
	ScreenshotConfigurationClassOnce.Do(func() {
		ScreenshotConfigurationClass = _ScreenshotConfigurationClass{objc.GetClass("SCScreenshotConfiguration")}
	})
	return ScreenshotConfigurationClass
}

type _ScreenshotConfigurationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ScreenshotConfiguration */
// An interface definition for the [ScreenshotConfiguration] class.
type IScreenshotConfiguration interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ScreenshotConfiguration */
	// properties:
	ContentType() uniformtypeidentifiers.UTType
	SetContentType(value uniformtypeidentifiers.UTType)
	DestinationRect() corefoundation.CGRect
	SetDestinationRect(value corefoundation.CGRect)
	DisplayIntent() ScreenshotDisplayIntent
	SetDisplayIntent(value ScreenshotDisplayIntent)
	DynamicRange() ScreenshotDynamicRange
	SetDynamicRange(value ScreenshotDynamicRange)
	FileURL() objc.IObject /* cross-framework: NSURL */
	SetFileURL(value objc.IObject /* cross-framework: NSURL */)
	Height() int
	SetHeight(value int)
	IgnoreClipping() bool
	SetIgnoreClipping(value bool)
	IgnoreShadows() bool
	SetIgnoreShadows(value bool)
	IncludeChildWindows() bool
	SetIncludeChildWindows(value bool)
	ShowsCursor() bool
	SetShowsCursor(value bool)
	SourceRect() corefoundation.CGRect
	SetSourceRect(value corefoundation.CGRect)
	Width() int
	SetWidth(value int)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ScreenshotConfiguration */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ScreenshotConfiguration */
// Alloc allocates a new instance without initialization.
func (sc _ScreenshotConfigurationClass) Alloc() ScreenshotConfiguration {
	rv := objc.Send[ScreenshotConfiguration](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _ScreenshotConfigurationClass) New() ScreenshotConfiguration {
	rv := objc.Send[ScreenshotConfiguration](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ ScreenshotConfiguration) Init() ScreenshotConfiguration {
	rv := objc.Send[ScreenshotConfiguration](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ ScreenshotConfiguration) Autorelease() ScreenshotConfiguration {
	rv := objc.Send[ScreenshotConfiguration](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewScreenshotConfiguration creates a new ScreenshotConfiguration instance.
func NewScreenshotConfiguration() ScreenshotConfiguration {
	return getScreenshotConfigurationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ScreenshotConfiguration */
// An object that contains screenshot properties such as output width, height, and image quality specifications.
//
// provides a default image capture configuration for . Only configure its properties if you need to customize the output. Additional options for customization include dynamic range settings, image reproduction optimizations, and ignoring user interface elements.


// An object that contains screenshot properties such as output width, height, and image quality specifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration
type ScreenshotConfiguration struct {
	objectivec.Object
}

// ScreenshotConfigurationFrom constructs a [ScreenshotConfiguration] from an unsafe.Pointer.
//
// An object that contains screenshot properties such as output width, height, and image quality specifications.
func ScreenshotConfigurationFrom(ptr unsafe.Pointer) ScreenshotConfiguration {
	return ScreenshotConfiguration{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ScreenshotConfiguration *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ScreenshotConfiguration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ScreenshotConfiguration */

// An array of uniform type identifiers that correspond to file formats the output image supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration/supportedContentTypes
func (sc _ScreenshotConfigurationClass) SupportedContentTypes() []uniformtypeidentifiers.UTType {
	rv := objc.Send[[]uniformtypeidentifiers.UTType](objc.ID(sc.class), objc.Sel("supportedContentTypes"))
	return rv
}/* debug [class_properties_class/property]: supportedContentTypes */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ScreenshotConfiguration */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ScreenshotConfiguration */

// A uniform type identifier that specifies the screenshot’s file format; HEIC, JPEG, or PNG.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration/contentType
func (s_ ScreenshotConfiguration) ContentType() uniformtypeidentifiers.UTType {
	rv := objc.Send[uniformtypeidentifiers.UTType](s_.ID, objc.Sel("contentType"))
	return rv
}/* debug [instance_properties/getter]: contentType */


// A uniform type identifier that specifies the screenshot’s file format; HEIC, JPEG, or PNG.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration/contentType
func (s_ ScreenshotConfiguration) SetContentType(value uniformtypeidentifiers.UTType) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setContentType:"), value)
}/* debug [instance_properties/setter]: contentType */


// A rectangle that specifies whether to output screenshots in a subset of the output image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration/destinationRect
func (s_ ScreenshotConfiguration) DestinationRect() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](s_.ID, objc.Sel("destinationRect"))
	return rv
}/* debug [instance_properties/getter]: destinationRect */


// A rectangle that specifies whether to output screenshots in a subset of the output image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration/destinationRect
func (s_ ScreenshotConfiguration) SetDestinationRect(value corefoundation.CGRect) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDestinationRect:"), value)
}/* debug [instance_properties/setter]: destinationRect */


// Specifies whether the screen capture uses attributes of the local or canonical display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration/displayIntent-swift.property
func (s_ ScreenshotConfiguration) DisplayIntent() ScreenshotDisplayIntent {
	rv := objc.Send[ScreenshotDisplayIntent](s_.ID, objc.Sel("displayIntent"))
	return rv
}/* debug [instance_properties/getter]: displayIntent */


// Specifies whether the screen capture uses attributes of the local or canonical display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration/displayIntent-swift.property
func (s_ ScreenshotConfiguration) SetDisplayIntent(value ScreenshotDisplayIntent) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDisplayIntent:"), value)
}/* debug [instance_properties/setter]: displayIntent */


// Specifies the type of image returned to the client; standard dynamic range, high dynamic range, or both.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration/dynamicRange-swift.property
func (s_ ScreenshotConfiguration) DynamicRange() ScreenshotDynamicRange {
	rv := objc.Send[ScreenshotDynamicRange](s_.ID, objc.Sel("dynamicRange"))
	return rv
}/* debug [instance_properties/getter]: dynamicRange */


// Specifies the type of image returned to the client; standard dynamic range, high dynamic range, or both.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration/dynamicRange-swift.property
func (s_ ScreenshotConfiguration) SetDynamicRange(value ScreenshotDynamicRange) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDynamicRange:"), value)
}/* debug [instance_properties/setter]: dynamicRange */


// Specifies the URL where the screenshot process saves the output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration/fileURL
func (s_ ScreenshotConfiguration) FileURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](s_.ID, objc.Sel("fileURL"))
	return rv
}/* debug [instance_properties/getter]: fileURL */


// Specifies the URL where the screenshot process saves the output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration/fileURL
func (s_ ScreenshotConfiguration) SetFileURL(value objc.IObject /* cross-framework: NSURL */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setFileURL:"), value)
}/* debug [instance_properties/setter]: fileURL */


// An integer value that specifies the output height, measured in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration/height
func (s_ ScreenshotConfiguration) Height() int {
	rv := objc.Send[int](s_.ID, objc.Sel("height"))
	return rv
}/* debug [instance_properties/getter]: height */


// An integer value that specifies the output height, measured in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration/height
func (s_ ScreenshotConfiguration) SetHeight(value int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setHeight:"), value)
}/* debug [instance_properties/setter]: height */


// A Boolean value that specifies whether to ignore framing on windows when using content filters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration/ignoreClipping
func (s_ ScreenshotConfiguration) IgnoreClipping() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("ignoreClipping"))
	return rv
}/* debug [instance_properties/getter]: ignoreClipping */


// A Boolean value that specifies whether to ignore framing on windows when using content filters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration/ignoreClipping
func (s_ ScreenshotConfiguration) SetIgnoreClipping(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIgnoreClipping:"), value)
}/* debug [instance_properties/setter]: ignoreClipping */


// A Boolean value that specifies whether to ignore framing on windows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration/ignoreShadows
func (s_ ScreenshotConfiguration) IgnoreShadows() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("ignoreShadows"))
	return rv
}/* debug [instance_properties/getter]: ignoreShadows */


// A Boolean value that specifies whether to ignore framing on windows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration/ignoreShadows
func (s_ ScreenshotConfiguration) SetIgnoreShadows(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIgnoreShadows:"), value)
}/* debug [instance_properties/setter]: ignoreShadows */


// A Boolean that specifies whether the screenshot captures subwindows of the included apps and windows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration/includeChildWindows
func (s_ ScreenshotConfiguration) IncludeChildWindows() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("includeChildWindows"))
	return rv
}/* debug [instance_properties/getter]: includeChildWindows */


// A Boolean that specifies whether the screenshot captures subwindows of the included apps and windows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration/includeChildWindows
func (s_ ScreenshotConfiguration) SetIncludeChildWindows(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIncludeChildWindows:"), value)
}/* debug [instance_properties/setter]: includeChildWindows */


// A Boolean value that specifies whether the pointer appears in the screenshot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration/showsCursor
func (s_ ScreenshotConfiguration) ShowsCursor() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("showsCursor"))
	return rv
}/* debug [instance_properties/getter]: showsCursor */


// A Boolean value that specifies whether the pointer appears in the screenshot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration/showsCursor
func (s_ ScreenshotConfiguration) SetShowsCursor(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setShowsCursor:"), value)
}/* debug [instance_properties/setter]: showsCursor */


// A rectangle that specifies that the screenshot only samples a subset of the frame input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration/sourceRect
func (s_ ScreenshotConfiguration) SourceRect() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](s_.ID, objc.Sel("sourceRect"))
	return rv
}/* debug [instance_properties/getter]: sourceRect */


// A rectangle that specifies that the screenshot only samples a subset of the frame input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration/sourceRect
func (s_ ScreenshotConfiguration) SetSourceRect(value corefoundation.CGRect) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSourceRect:"), value)
}/* debug [instance_properties/setter]: sourceRect */


// An array of uniform type identifiers that correspond to file formats the output image supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration/supportedContentTypes
func (s_ ScreenshotConfiguration) SupportedContentTypes() []uniformtypeidentifiers.UTType {
	rv := objc.Send[[]uniformtypeidentifiers.UTType](s_.ID, objc.Sel("supportedContentTypes"))
	return rv
}/* debug [instance_properties/getter]: supportedContentTypes */


// An integer value that specifies the output width in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration/width
func (s_ ScreenshotConfiguration) Width() int {
	rv := objc.Send[int](s_.ID, objc.Sel("width"))
	return rv
}/* debug [instance_properties/getter]: width */


// An integer value that specifies the output width in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration/width
func (s_ ScreenshotConfiguration) SetWidth(value int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setWidth:"), value)
}/* debug [instance_properties/setter]: width */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class SCScreenshotConfiguration */



