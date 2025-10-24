// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class SCContentFilter */


/* debug [class_header]: Header for SCContentFilter */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ContentFilter */
// An interface definition for the [ContentFilter] class.
type IContentFilter interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ContentFilter */
	// properties:
	ContentRect() corefoundation.CGRect
	IncludedApplications() []RunningApplication
	IncludedDisplays() []Display
	IncludedWindows() []Window
	IncludeMenuBar() bool
	SetIncludeMenuBar(value bool)
	PointPixelScale() float32
	StreamType() StreamType
	Style() ShareableContentStyle
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ContentFilter */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ContentFilter */
// Alloc allocates a new instance without initialization.
func (cc _ContentFilterClass) Alloc() ContentFilter {
	rv := objc.Send[ContentFilter](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ContentFilter */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ContentFilter */

// Creates a filter that captures only the specified window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentFilter/init(desktopIndependentWindow:)
func NewContentFilterWithDesktopIndependentWindow(window ISCWindow) ContentFilter {
	instance := getContentFilterClass().Alloc()
	rv := objc.Send[ContentFilter](instance.ID, objc.Sel("initWithDesktopIndependentWindow:"), window)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewContentFilterWithDesktopIndependentWindow */


// Creates a filter that captures a display, excluding windows of the specified apps.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentFilter/init(display:excludingApplications:exceptingWindows:)
func NewContentFilterWithDisplayExcludingApplicationsExceptingWindows(display ISCDisplay, applications []RunningApplication, exceptingWindows []Window) ContentFilter {
	instance := getContentFilterClass().Alloc()
	rv := objc.Send[ContentFilter](instance.ID, objc.Sel("initWithDisplay:excludingApplications:exceptingWindows:"), display, applications, exceptingWindows)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewContentFilterWithDisplayExcludingApplicationsExceptingWindows */


// Creates a filter that captures the contents of a display, excluding the specified windows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentFilter/init(display:excludingWindows:)
func NewContentFilterWithDisplayExcludingWindows(display ISCDisplay, excluded []Window) ContentFilter {
	instance := getContentFilterClass().Alloc()
	rv := objc.Send[ContentFilter](instance.ID, objc.Sel("initWithDisplay:excludingWindows:"), display, excluded)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewContentFilterWithDisplayExcludingWindows */


// Creates a filter that captures a display, including only windows of the specified apps.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentFilter/init(display:including:exceptingWindows:)
func NewContentFilterWithDisplayIncludingApplicationsExceptingWindows(display ISCDisplay, applications []RunningApplication, exceptingWindows []Window) ContentFilter {
	instance := getContentFilterClass().Alloc()
	rv := objc.Send[ContentFilter](instance.ID, objc.Sel("initWithDisplay:includingApplications:exceptingWindows:"), display, applications, exceptingWindows)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewContentFilterWithDisplayIncludingApplicationsExceptingWindows */


// Creates a filter that captures only specific windows from a display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentFilter/init(display:including:)
func NewContentFilterWithDisplayIncludingWindows(display ISCDisplay, includedWindows []Window) ContentFilter {
	instance := getContentFilterClass().Alloc()
	rv := objc.Send[ContentFilter](instance.ID, objc.Sel("initWithDisplay:includingWindows:"), display, includedWindows)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewContentFilterWithDisplayIncludingWindows */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ContentFilter */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ContentFilter */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ContentFilter */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ContentFilter */

// The size and location of the content to filter, in screen points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentFilter/contentRect
func (c_ ContentFilter) ContentRect() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](c_.ID, objc.Sel("contentRect"))
	return rv
}/* debug [instance_properties/getter]: contentRect */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentFilter/includedApplications
func (c_ ContentFilter) IncludedApplications() []RunningApplication {
	rv := objc.Send[[]RunningApplication](c_.ID, objc.Sel("includedApplications"))
	return rv
}/* debug [instance_properties/getter]: includedApplications */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentFilter/includedDisplays
func (c_ ContentFilter) IncludedDisplays() []Display {
	rv := objc.Send[[]Display](c_.ID, objc.Sel("includedDisplays"))
	return rv
}/* debug [instance_properties/getter]: includedDisplays */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentFilter/includedWindows
func (c_ ContentFilter) IncludedWindows() []Window {
	rv := objc.Send[[]Window](c_.ID, objc.Sel("includedWindows"))
	return rv
}/* debug [instance_properties/getter]: includedWindows */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentFilter/includeMenuBar
func (c_ ContentFilter) IncludeMenuBar() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("includeMenuBar"))
	return rv
}/* debug [instance_properties/getter]: includeMenuBar */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentFilter/includeMenuBar
func (c_ ContentFilter) SetIncludeMenuBar(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIncludeMenuBar:"), value)
}/* debug [instance_properties/setter]: includeMenuBar */


// The scaling factor used to translate screen points into pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentFilter/pointPixelScale
func (c_ ContentFilter) PointPixelScale() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("pointPixelScale"))
	return rv
}/* debug [instance_properties/getter]: pointPixelScale */


// The type of the streaming content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentFilter/streamType
func (c_ ContentFilter) StreamType() StreamType {
	rv := objc.Send[StreamType](c_.ID, objc.Sel("streamType"))
	return rv
}/* debug [instance_properties/getter]: streamType */


// The display style of the sharable content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentFilter/style
func (c_ ContentFilter) Style() ShareableContentStyle {
	rv := objc.Send[ShareableContentStyle](c_.ID, objc.Sel("style"))
	return rv
}/* debug [instance_properties/getter]: style */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class SCContentFilter */


