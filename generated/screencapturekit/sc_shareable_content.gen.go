// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class SCShareableContent */


/* debug [class_header]: Header for SCShareableContent */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ShareableContent */
// An interface definition for the [ShareableContent] class.
type IShareableContent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ShareableContent */
	// properties:
	Applications() []RunningApplication
	Displays() []Display
	Windows() []Window
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ShareableContent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ShareableContent */
// Alloc allocates a new instance without initialization.
func (sc _ShareableContentClass) Alloc() ShareableContent {
	rv := objc.Send[ShareableContent](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ShareableContent */
// An instance that represents a set of displays, apps, and windows that your app can capture.
//
// Use the , , and properties to create a object that specifies what display content to capture. You apply the filter to an instance of to limit its output to only the content matching your filter.


// An instance that represents a set of displays, apps, and windows that your app can capture.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ShareableContent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ShareableContent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCShareableContent/getCurrentProcessShareableContent(completionHandler:)
func (sc _ShareableContentClass) GetCurrentProcessShareableContentWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("getCurrentProcessShareableContentWithCompletionHandler:"), completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=GetCurrentProcessShareableContentWithCompletionHandler) */


// Retrieves the displays, apps, and windows that match your criteria.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCShareableContent/getExcludingDesktopWindows(_:onScreenWindowsOnly:completionHandler:)
func (sc _ShareableContentClass) GetShareableContentExcludingDesktopWindowsOnScreenWindowsOnlyCompletionHandler(excludeDesktopWindows bool, onScreenWindowsOnly bool, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("getShareableContentExcludingDesktopWindows:onScreenWindowsOnly:completionHandler:"), excludeDesktopWindows, onScreenWindowsOnly, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=GetShareableContentExcludingDesktopWindowsOnScreenWindowsOnlyCompletionHandler) */


// Retrieves the displays, apps, and windows that are in front of the specified window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCShareableContent/getExcludingDesktopWindows(_:onScreenWindowsOnlyAbove:completionHandler:)
func (sc _ShareableContentClass) GetShareableContentExcludingDesktopWindowsOnScreenWindowsOnlyAboveWindowCompletionHandler(excludeDesktopWindows bool, window ISCWindow, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("getShareableContentExcludingDesktopWindows:onScreenWindowsOnlyAboveWindow:completionHandler:"), excludeDesktopWindows, window, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=GetShareableContentExcludingDesktopWindowsOnScreenWindowsOnlyAboveWindowCompletionHandler) */


// Retrieves the displays, apps, and windows that are behind the specified window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCShareableContent/getExcludingDesktopWindows(_:onScreenWindowsOnlyBelow:completionHandler:)
func (sc _ShareableContentClass) GetShareableContentExcludingDesktopWindowsOnScreenWindowsOnlyBelowWindowCompletionHandler(excludeDesktopWindows bool, window ISCWindow, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("getShareableContentExcludingDesktopWindows:onScreenWindowsOnlyBelowWindow:completionHandler:"), excludeDesktopWindows, window, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=GetShareableContentExcludingDesktopWindowsOnScreenWindowsOnlyBelowWindowCompletionHandler) */


// Retrieves the displays, apps, and windows that your app can capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCShareableContent/getWithCompletionHandler(_:)
func (sc _ShareableContentClass) GetShareableContentWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("getShareableContentWithCompletionHandler:"), completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=GetShareableContentWithCompletionHandler) */


// Retrieves any available sharable content information that matches the provided filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCShareableContent/info(for:)
func (sc _ShareableContentClass) InfoForFilter(filter ISCContentFilter) IShareableContentInfo {
	rv := objc.Send[ShareableContentInfo](objc.ID(sc.class), objc.Sel("infoForFilter:"), filter)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=InfoForFilter) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ShareableContent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ShareableContent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ShareableContent */

// The apps available for capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCShareableContent/applications
func (s_ ShareableContent) Applications() []RunningApplication {
	rv := objc.Send[[]RunningApplication](s_.ID, objc.Sel("applications"))
	return rv
}/* debug [instance_properties/getter]: applications */


// The displays available for capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCShareableContent/displays
func (s_ ShareableContent) Displays() []Display {
	rv := objc.Send[[]Display](s_.ID, objc.Sel("displays"))
	return rv
}/* debug [instance_properties/getter]: displays */


// The windows available for capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCShareableContent/windows
func (s_ ShareableContent) Windows() []Window {
	rv := objc.Send[[]Window](s_.ID, objc.Sel("windows"))
	return rv
}/* debug [instance_properties/getter]: windows */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class SCShareableContent */



