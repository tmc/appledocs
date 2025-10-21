// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
}

// An instance that filters the content a stream captures.
//
// Use a content filter to limit an object’s output to only that matching your filter criteria. Retrieve the displays, apps, and windows that your app can capture from an instance of .
//
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
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentFilter/init(display:excludingApplications:exceptingWindows:)
func NewContentFilterWithDisplayExcludingApplicationsExceptingWindows(display unsafe.Pointer, applications unsafe.Pointer, exceptingWindows unsafe.Pointer) ContentFilter {
	instance := getContentFilterClass().Alloc()
	rv := objc.Send[ContentFilter](instance.ID, objc.Sel("initWithDisplay:excludingApplications:exceptingWindows:"), display, applications, exceptingWindows)
	rv.Autorelease()
	return rv
}



// Creates a filter that captures the contents of a display, excluding the specified windows.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentFilter/init(display:excludingWindows:)
func NewContentFilterWithDisplayExcludingWindows(display unsafe.Pointer, excluded unsafe.Pointer) ContentFilter {
	instance := getContentFilterClass().Alloc()
	rv := objc.Send[ContentFilter](instance.ID, objc.Sel("initWithDisplay:excludingWindows:"), display, excluded)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentFilter/includeMenuBar
func (c_ ContentFilter) IncludeMenuBar() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("includeMenuBar"))
	return rv
}


// SetIncludeMenuBar sets the value of the includeMenuBar property.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentFilter/includeMenuBar
func (c_ ContentFilter) SetIncludeMenuBar(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIncludeMenuBar:"), value)
}


