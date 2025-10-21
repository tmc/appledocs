// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
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
}

//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration
type ScreenshotConfiguration struct {
	objectivec.Object
}

// ScreenshotConfigurationFrom constructs a [ScreenshotConfiguration] from an unsafe.Pointer.
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


//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration/supportedContentTypes
func (sc _ScreenshotConfigurationClass) SupportedContentTypes() []UTType {
	rv := objc.Send[[]UTType](objc.ID(sc.class), objc.Sel("supportedContentTypes"))
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration/destinationRect
func (s_ ScreenshotConfiguration) DestinationRect() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](s_.ID, objc.Sel("destinationRect"))
	return rv
}


// SetDestinationRect sets the value of the destinationRect property.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration/destinationRect
func (s_ ScreenshotConfiguration) SetDestinationRect(value coregraphics.CGRect) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDestinationRect:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration/sourceRect
func (s_ ScreenshotConfiguration) SourceRect() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](s_.ID, objc.Sel("sourceRect"))
	return rv
}


// SetSourceRect sets the value of the sourceRect property.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration/sourceRect
func (s_ ScreenshotConfiguration) SetSourceRect(value coregraphics.CGRect) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSourceRect:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration/supportedContentTypes
func (s_ ScreenshotConfiguration) SupportedContentTypes() []UTType {
	rv := objc.Send[[]UTType](s_.ID, objc.Sel("supportedContentTypes"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scscreenshotconfiguration/contenttype
func (s_ ScreenshotConfiguration) ContentType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("contentType"))
	return rv
}


// SetContentType sets the value of the contentType property.
//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scscreenshotconfiguration/contenttype
func (s_ ScreenshotConfiguration) SetContentType(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setContentType:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scscreenshotconfiguration/displayintent-swift.property
func (s_ ScreenshotConfiguration) DisplayIntent() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("displayIntent"))
	return rv
}


// SetDisplayIntent sets the value of the displayIntent property.
//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scscreenshotconfiguration/displayintent-swift.property
func (s_ ScreenshotConfiguration) SetDisplayIntent(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDisplayIntent:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scscreenshotconfiguration/dynamicrange-swift.property
func (s_ ScreenshotConfiguration) DynamicRange() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("dynamicRange"))
	return rv
}


// SetDynamicRange sets the value of the dynamicRange property.
//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scscreenshotconfiguration/dynamicrange-swift.property
func (s_ ScreenshotConfiguration) SetDynamicRange(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDynamicRange:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scscreenshotconfiguration/fileurl
func (s_ ScreenshotConfiguration) FileURL() foundation.URL {
	rv := objc.Send[foundation.URL](s_.ID, objc.Sel("fileURL"))
	return rv
}


// SetFileURL sets the value of the fileURL property.
//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scscreenshotconfiguration/fileurl
func (s_ ScreenshotConfiguration) SetFileURL(value foundation.URL) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setFileURL:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scscreenshotconfiguration/height
func (s_ ScreenshotConfiguration) Height() int {
	rv := objc.Send[int](s_.ID, objc.Sel("height"))
	return rv
}


// SetHeight sets the value of the height property.
//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scscreenshotconfiguration/height
func (s_ ScreenshotConfiguration) SetHeight(value int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setHeight:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scscreenshotconfiguration/ignoreclipping
func (s_ ScreenshotConfiguration) IgnoreClipping() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("ignoreClipping"))
	return rv
}


// SetIgnoreClipping sets the value of the ignoreClipping property.
//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scscreenshotconfiguration/ignoreclipping
func (s_ ScreenshotConfiguration) SetIgnoreClipping(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIgnoreClipping:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scscreenshotconfiguration/ignoreshadows
func (s_ ScreenshotConfiguration) IgnoreShadows() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("ignoreShadows"))
	return rv
}


// SetIgnoreShadows sets the value of the ignoreShadows property.
//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scscreenshotconfiguration/ignoreshadows
func (s_ ScreenshotConfiguration) SetIgnoreShadows(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIgnoreShadows:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scscreenshotconfiguration/includechildwindows
func (s_ ScreenshotConfiguration) IncludeChildWindows() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("includeChildWindows"))
	return rv
}


// SetIncludeChildWindows sets the value of the includeChildWindows property.
//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scscreenshotconfiguration/includechildwindows
func (s_ ScreenshotConfiguration) SetIncludeChildWindows(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIncludeChildWindows:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scscreenshotconfiguration/showscursor
func (s_ ScreenshotConfiguration) ShowsCursor() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("showsCursor"))
	return rv
}


// SetShowsCursor sets the value of the showsCursor property.
//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scscreenshotconfiguration/showscursor
func (s_ ScreenshotConfiguration) SetShowsCursor(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setShowsCursor:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scscreenshotconfiguration/width
func (s_ ScreenshotConfiguration) Width() int {
	rv := objc.Send[int](s_.ID, objc.Sel("width"))
	return rv
}


// SetWidth sets the value of the width property.
//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scscreenshotconfiguration/width
func (s_ ScreenshotConfiguration) SetWidth(value int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setWidth:"), value)
}



