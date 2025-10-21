// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ShareableContent] class.
var (
	ShareableContentClass     _ShareableContentClass
	ShareableContentClassOnce sync.Once
)

func getShareableContentClass() _ShareableContentClass {
	ShareableContentClassOnce.Do(func() {
		ShareableContentClass = _ShareableContentClass{objc.GetClass("SCShareableContent")}
	})
	return ShareableContentClass
}

type _ShareableContentClass struct {
	class objc.Class
}

// An interface definition for the [ShareableContent] class.
type IShareableContent interface {
	objectivec.IObject
}

// An instance that represents a set of displays, apps, and windows that your app can capture.
//
// Use the , , and properties to create a object that specifies what display content to capture. You apply the filter to an instance of to limit its output to only the content matching your filter.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCShareableContent
type ShareableContent struct {
	objectivec.Object
}

// ShareableContentFrom constructs a [ShareableContent] from an unsafe.Pointer.
//
// An instance that represents a set of displays, apps, and windows that your app can capture.
func ShareableContentFrom(ptr unsafe.Pointer) ShareableContent {
	return ShareableContent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _ShareableContentClass) Alloc() ShareableContent {
	rv := objc.Send[ShareableContent](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _ShareableContentClass) New() ShareableContent {
	rv := objc.Send[ShareableContent](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ ShareableContent) Init() ShareableContent {
	rv := objc.Send[ShareableContent](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ ShareableContent) Autorelease() ShareableContent {
	rv := objc.Send[ShareableContent](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewShareableContent creates a new ShareableContent instance.
func NewShareableContent() ShareableContent {
	return getShareableContentClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCShareableContent/getCurrentProcessShareableContent(completionHandler:)
func (sc _ShareableContentClass) GetCurrentProcessShareableContentWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("getCurrentProcessShareableContentWithCompletionHandler:"), completionHandler)
}

// Retrieves the displays, apps, and windows that match your criteria.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCShareableContent/getExcludingDesktopWindows(_:onScreenWindowsOnly:completionHandler:)
func (sc _ShareableContentClass) GetShareableContentExcludingDesktopWindowsOnScreenWindowsOnlyCompletionHandler(excludeDesktopWindows bool, onScreenWindowsOnly bool, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("getShareableContentExcludingDesktopWindows:onScreenWindowsOnly:completionHandler:"), excludeDesktopWindows, onScreenWindowsOnly, completionHandler)
}

// The apps available for capture.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCShareableContent/applications
func (s_ ShareableContent) Applications() []RunningApplication {
	rv := objc.Send[[]RunningApplication](s_.ID, objc.Sel("applications"))
	return rv
}

// The displays available for capture.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCShareableContent/displays
func (s_ ShareableContent) Displays() []Display {
	rv := objc.Send[[]Display](s_.ID, objc.Sel("displays"))
	return rv
}

// The windows available for capture.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCShareableContent/windows
func (s_ ShareableContent) Windows() []Window {
	rv := objc.Send[[]Window](s_.ID, objc.Sel("windows"))
	return rv
}



