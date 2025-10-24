// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ContentFilter] class.
var (
	ContentFilterClass     _ContentFilterClass
	ContentFilterClassOnce sync.Once
)

func getContentFilterClass() _ContentFilterClass {
	ContentFilterClassOnce.Do(func() {
		ContentFilterClass = _ContentFilterClass{objc.GetClass("SCContentFilter")}
	})
	return ContentFilterClass
}

type _ContentFilterClass struct {
	class objc.Class
}

// An interface definition for the [ContentFilter] class.
type IContentFilter interface {
	objectivec.IObject
	// properties:
	IncludeMenuBar() bool
	SetIncludeMenuBar(value bool)
	StreamType() StreamType
	ContentRect() objc.IObject /* cross-framework: Rect */
	SetContentRect(value objc.IObject /* cross-framework: Rect */)
	IncludedApplications() ISCRunningApplication
	SetIncludedApplications(value ISCRunningApplication)
	IncludedDisplays() ISCDisplay
	SetIncludedDisplays(value ISCDisplay)
	IncludedWindows() ISCWindow
	SetIncludedWindows(value ISCWindow)
	PointPixelScale() float32
	SetPointPixelScale(value float32)
	Style() ShareableContentStyle
	SetStyle(value ShareableContentStyle)
	// methods:
}

// An instance that filters the content a stream captures.
//
// Use a content filter to limit an object’s output to only that matching your filter criteria. Retrieve the displays, apps, and windows that your app can capture from an instance of .


// An instance that filters the content a stream captures.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentFilter
type ContentFilter struct {
	objectivec.Object
}

// ContentFilterFrom constructs a [ContentFilter] from an unsafe.Pointer.
//
// An instance that filters the content a stream captures.
func ContentFilterFrom(ptr unsafe.Pointer) ContentFilter {
	return ContentFilter{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _ContentFilterClass) Alloc() ContentFilter {
	rv := objc.Send[ContentFilter](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _ContentFilterClass) New() ContentFilter {
	rv := objc.Send[ContentFilter](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ContentFilter) Init() ContentFilter {
	rv := objc.Send[ContentFilter](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ContentFilter) Autorelease() ContentFilter {
	rv := objc.Send[ContentFilter](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewContentFilter creates a new ContentFilter instance.
func NewContentFilter() ContentFilter {
	return getContentFilterClass().New()
}



// Creates a filter that captures a display, excluding windows of the specified apps.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentFilter/init(display:excludingApplications:exceptingWindows:)
func NewContentFilterWithDisplayExcludingApplicationsExceptingWindows(display ISCDisplay, applications []IRunningApplication, exceptingWindows []IWindow) ContentFilter {
	instance := getContentFilterClass().Alloc()
	rv := objc.Send[ContentFilter](instance.ID, objc.Sel("initWithDisplay:excludingApplications:exceptingWindows:"), display, applications, exceptingWindows)
	rv.Autorelease()
	return rv
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentFilter/includeMenuBar
func (c_ ContentFilter) IncludeMenuBar() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("includeMenuBar"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentFilter/includeMenuBar
func (c_ ContentFilter) SetIncludeMenuBar(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIncludeMenuBar:"), value)
}


// The type of the streaming content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentFilter/streamType
func (c_ ContentFilter) StreamType() StreamType {
	rv := objc.Send[StreamType](c_.ID, objc.Sel("streamType"))
	return rv
}


// The size and location of the content to filter, in screen points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/sccontentfilter/contentrect
func (c_ ContentFilter) ContentRect() objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[corefoundation.Rect](c_.ID, objc.Sel("contentRect"))
	return rv
}


// The size and location of the content to filter, in screen points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/sccontentfilter/contentrect
func (c_ ContentFilter) SetContentRect(value objc.IObject /* cross-framework: Rect */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContentRect:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/sccontentfilter/includedapplications
func (c_ ContentFilter) IncludedApplications() ISCRunningApplication {
	rv := objc.Send[RunningApplication](c_.ID, objc.Sel("includedApplications"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/sccontentfilter/includedapplications
func (c_ ContentFilter) SetIncludedApplications(value ISCRunningApplication) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIncludedApplications:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/sccontentfilter/includeddisplays
func (c_ ContentFilter) IncludedDisplays() ISCDisplay {
	rv := objc.Send[Display](c_.ID, objc.Sel("includedDisplays"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/sccontentfilter/includeddisplays
func (c_ ContentFilter) SetIncludedDisplays(value ISCDisplay) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIncludedDisplays:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/sccontentfilter/includedwindows
func (c_ ContentFilter) IncludedWindows() ISCWindow {
	rv := objc.Send[Window](c_.ID, objc.Sel("includedWindows"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/sccontentfilter/includedwindows
func (c_ ContentFilter) SetIncludedWindows(value ISCWindow) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIncludedWindows:"), value)
}


// The scaling factor used to translate screen points into pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/sccontentfilter/pointpixelscale
func (c_ ContentFilter) PointPixelScale() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("pointPixelScale"))
	return rv
}


// The scaling factor used to translate screen points into pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/sccontentfilter/pointpixelscale
func (c_ ContentFilter) SetPointPixelScale(value float32) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPointPixelScale:"), value)
}


// The display style of the sharable content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/sccontentfilter/style
func (c_ ContentFilter) Style() ShareableContentStyle {
	rv := objc.Send[ShareableContentStyle](c_.ID, objc.Sel("style"))
	return rv
}


// The display style of the sharable content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/sccontentfilter/style
func (c_ ContentFilter) SetStyle(value ShareableContentStyle) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStyle:"), value)
}


