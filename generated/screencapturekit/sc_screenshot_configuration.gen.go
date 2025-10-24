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

// An interface definition for the [ScreenshotConfiguration] class.
type IScreenshotConfiguration interface {
	objectivec.IObject
	// properties:
	DisplayIntent() ScreenshotDisplayIntent
	SetDisplayIntent(value ScreenshotDisplayIntent)
	ContentType() objc.IObject /* cross-framework: UTType */
	SetContentType(value objc.IObject /* cross-framework: UTType */)
	DestinationRect() objc.IObject /* cross-framework: Rect */
	SetDestinationRect(value objc.IObject /* cross-framework: Rect */)
	DynamicRange() unsafe.Pointer
	SetDynamicRange(value unsafe.Pointer)
	FileURL() objc.IObject /* cross-framework: URL */
	SetFileURL(value objc.IObject /* cross-framework: URL */)
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
	SourceRect() objc.IObject /* cross-framework: Rect */
	SetSourceRect(value objc.IObject /* cross-framework: Rect */)
	Width() int
	SetWidth(value int)
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (sc _ScreenshotConfigurationClass) Alloc() ScreenshotConfiguration {
	rv := objc.Send[ScreenshotConfiguration](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration/supportedContentTypes
func (sc _ScreenshotConfigurationClass) SupportedContentTypes() []objc.IObject /* cross-framework: UTType */ {
	rv := objc.Send[[]uniformtypeidentifiers.UTType](objc.ID(sc.class), objc.Sel("supportedContentTypes"))
	return rv
}

// Specifies whether the screen capture uses attributes of the local or canonical display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration/displayIntent-swift.property
func (s_ ScreenshotConfiguration) DisplayIntent() ScreenshotDisplayIntent {
	rv := objc.Send[ScreenshotDisplayIntent](s_.ID, objc.Sel("displayIntent"))
	return rv
}


// Specifies whether the screen capture uses attributes of the local or canonical display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration/displayIntent-swift.property
func (s_ ScreenshotConfiguration) SetDisplayIntent(value ScreenshotDisplayIntent) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDisplayIntent:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration/supportedContentTypes
func (s_ ScreenshotConfiguration) SupportedContentTypes() []objc.IObject /* cross-framework: UTType */ {
	rv := objc.Send[[]uniformtypeidentifiers.UTType](s_.ID, objc.Sel("supportedContentTypes"))
	return rv
}


// A uniform type identifier that specifies the screenshot’s file format; HEIC, JPEG, or PNG.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scscreenshotconfiguration/contenttype
func (s_ ScreenshotConfiguration) ContentType() objc.IObject /* cross-framework: UTType */ {
	rv := objc.Send[uniformtypeidentifiers.UTType](s_.ID, objc.Sel("contentType"))
	return rv
}


// A uniform type identifier that specifies the screenshot’s file format; HEIC, JPEG, or PNG.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scscreenshotconfiguration/contenttype
func (s_ ScreenshotConfiguration) SetContentType(value objc.IObject /* cross-framework: UTType */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setContentType:"), value)
}


// A rectangle that specifies whether to output screenshots in a subset of the output image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scscreenshotconfiguration/destinationrect
func (s_ ScreenshotConfiguration) DestinationRect() objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[corefoundation.Rect](s_.ID, objc.Sel("destinationRect"))
	return rv
}


// A rectangle that specifies whether to output screenshots in a subset of the output image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scscreenshotconfiguration/destinationrect
func (s_ ScreenshotConfiguration) SetDestinationRect(value objc.IObject /* cross-framework: Rect */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDestinationRect:"), value)
}


// Specifies the type of image returned to the client; standard dynamic range, high dynamic range, or both.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scscreenshotconfiguration/dynamicrange-swift.property
func (s_ ScreenshotConfiguration) DynamicRange() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("dynamicRange"))
	return rv
}


// Specifies the type of image returned to the client; standard dynamic range, high dynamic range, or both.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scscreenshotconfiguration/dynamicrange-swift.property
func (s_ ScreenshotConfiguration) SetDynamicRange(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDynamicRange:"), value)
}


// Specifies the URL where the screenshot process saves the output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scscreenshotconfiguration/fileurl
func (s_ ScreenshotConfiguration) FileURL() objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[foundation.URL](s_.ID, objc.Sel("fileURL"))
	return rv
}


// Specifies the URL where the screenshot process saves the output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scscreenshotconfiguration/fileurl
func (s_ ScreenshotConfiguration) SetFileURL(value objc.IObject /* cross-framework: URL */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setFileURL:"), value)
}


// An integer value that specifies the output height, measured in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scscreenshotconfiguration/height
func (s_ ScreenshotConfiguration) Height() int {
	rv := objc.Send[int](s_.ID, objc.Sel("height"))
	return rv
}


// An integer value that specifies the output height, measured in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scscreenshotconfiguration/height
func (s_ ScreenshotConfiguration) SetHeight(value int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setHeight:"), value)
}


// A Boolean value that specifies whether to ignore framing on windows when using content filters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scscreenshotconfiguration/ignoreclipping
func (s_ ScreenshotConfiguration) IgnoreClipping() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("ignoreClipping"))
	return rv
}


// A Boolean value that specifies whether to ignore framing on windows when using content filters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scscreenshotconfiguration/ignoreclipping
func (s_ ScreenshotConfiguration) SetIgnoreClipping(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIgnoreClipping:"), value)
}


// A Boolean value that specifies whether to ignore framing on windows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scscreenshotconfiguration/ignoreshadows
func (s_ ScreenshotConfiguration) IgnoreShadows() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("ignoreShadows"))
	return rv
}


// A Boolean value that specifies whether to ignore framing on windows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scscreenshotconfiguration/ignoreshadows
func (s_ ScreenshotConfiguration) SetIgnoreShadows(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIgnoreShadows:"), value)
}


// A Boolean that specifies whether the screenshot captures subwindows of the included apps and windows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scscreenshotconfiguration/includechildwindows
func (s_ ScreenshotConfiguration) IncludeChildWindows() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("includeChildWindows"))
	return rv
}


// A Boolean that specifies whether the screenshot captures subwindows of the included apps and windows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scscreenshotconfiguration/includechildwindows
func (s_ ScreenshotConfiguration) SetIncludeChildWindows(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIncludeChildWindows:"), value)
}


// A Boolean value that specifies whether the pointer appears in the screenshot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scscreenshotconfiguration/showscursor
func (s_ ScreenshotConfiguration) ShowsCursor() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("showsCursor"))
	return rv
}


// A Boolean value that specifies whether the pointer appears in the screenshot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scscreenshotconfiguration/showscursor
func (s_ ScreenshotConfiguration) SetShowsCursor(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setShowsCursor:"), value)
}


// A rectangle that specifies that the screenshot only samples a subset of the frame input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scscreenshotconfiguration/sourcerect
func (s_ ScreenshotConfiguration) SourceRect() objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[corefoundation.Rect](s_.ID, objc.Sel("sourceRect"))
	return rv
}


// A rectangle that specifies that the screenshot only samples a subset of the frame input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scscreenshotconfiguration/sourcerect
func (s_ ScreenshotConfiguration) SetSourceRect(value objc.IObject /* cross-framework: Rect */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSourceRect:"), value)
}


// An integer value that specifies the output width in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scscreenshotconfiguration/width
func (s_ ScreenshotConfiguration) Width() int {
	rv := objc.Send[int](s_.ID, objc.Sel("width"))
	return rv
}


// An integer value that specifies the output width in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scscreenshotconfiguration/width
func (s_ ScreenshotConfiguration) SetWidth(value int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setWidth:"), value)
}



