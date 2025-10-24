// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MediaLibraryBrowserController] class.
var (
	MediaLibraryBrowserControllerClass     _MediaLibraryBrowserControllerClass
	MediaLibraryBrowserControllerClassOnce sync.Once
)

func getMediaLibraryBrowserControllerClass() _MediaLibraryBrowserControllerClass {
	MediaLibraryBrowserControllerClassOnce.Do(func() {
		MediaLibraryBrowserControllerClass = _MediaLibraryBrowserControllerClass{objc.GetClass("NSMediaLibraryBrowserController")}
	})
	return MediaLibraryBrowserControllerClass
}

type _MediaLibraryBrowserControllerClass struct {
	class objc.Class
}

// An interface definition for the [MediaLibraryBrowserController] class.
type IMediaLibraryBrowserController interface {
	objectivec.IObject
	// properties:
	Frame() objc.IObject /* cross-framework: Rect */
	SetFrame(value objc.IObject /* cross-framework: Rect */)
	Visible() bool
	SetVisible(value bool)
	MediaLibraries() MediaLibrary
	SetMediaLibraries(value MediaLibrary)
	IsVisible() bool
	SetIsVisible(value bool)
	// methods:
	TogglePanel(sender objc.IObject)
}

// An object that configures and displays a Media Library Browser panel.
//
// From this panel a user can drag media into views in their app. The class provides a standard interface to the MediaLibrary framework content. For more information see , , , and in .


// An object that configures and displays a Media Library Browser panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMediaLibraryBrowserController
type MediaLibraryBrowserController struct {
	objectivec.Object
}

// MediaLibraryBrowserControllerFrom constructs a [MediaLibraryBrowserController] from an unsafe.Pointer.
//
// An object that configures and displays a Media Library Browser panel.
func MediaLibraryBrowserControllerFrom(ptr unsafe.Pointer) MediaLibraryBrowserController {
	return MediaLibraryBrowserController{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MediaLibraryBrowserControllerClass) Alloc() MediaLibraryBrowserController {
	rv := objc.Send[MediaLibraryBrowserController](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MediaLibraryBrowserControllerClass) New() MediaLibraryBrowserController {
	rv := objc.Send[MediaLibraryBrowserController](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MediaLibraryBrowserController) Init() MediaLibraryBrowserController {
	rv := objc.Send[MediaLibraryBrowserController](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MediaLibraryBrowserController) Autorelease() MediaLibraryBrowserController {
	rv := objc.Send[MediaLibraryBrowserController](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMediaLibraryBrowserController creates a new MediaLibraryBrowserController instance.
func NewMediaLibraryBrowserController() MediaLibraryBrowserController {
	return getMediaLibraryBrowserControllerClass().New()
}



// Returns the shared Media Library Browser instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMediaLibraryBrowserController/shared
func (mc _MediaLibraryBrowserControllerClass) SharedMediaLibraryBrowserController() MediaLibraryBrowserController {
	rv := objc.Send[MediaLibraryBrowserController](objc.ID(mc.class), objc.Sel("sharedMediaLibraryBrowserController"))
	return rv
}

// Toggles the visibility of the Media Library Browser.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMediaLibraryBrowserController/togglePanel(_:)
func (m_ MediaLibraryBrowserController) TogglePanel(sender objc.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("togglePanel:"), sender)
}


// The frame, in global coordinates, used to display the Media Library Browser panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMediaLibraryBrowserController/frame
func (m_ MediaLibraryBrowserController) Frame() objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[corefoundation.Rect](m_.ID, objc.Sel("frame"))
	return rv
}


// The frame, in global coordinates, used to display the Media Library Browser panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMediaLibraryBrowserController/frame
func (m_ MediaLibraryBrowserController) SetFrame(value objc.IObject /* cross-framework: Rect */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFrame:"), value)
}


// A Boolean value that determines whether the Media Library Browser panel is visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMediaLibraryBrowserController/isVisible
func (m_ MediaLibraryBrowserController) Visible() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("visible"))
	return rv
}


// A Boolean value that determines whether the Media Library Browser panel is visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMediaLibraryBrowserController/isVisible
func (m_ MediaLibraryBrowserController) SetVisible(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVisible:"), value)
}


// The media library that is in use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMediaLibraryBrowserController/mediaLibraries
func (m_ MediaLibraryBrowserController) MediaLibraries() MediaLibrary {
	rv := objc.Send[MediaLibrary](m_.ID, objc.Sel("mediaLibraries"))
	return rv
}


// The media library that is in use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMediaLibraryBrowserController/mediaLibraries
func (m_ MediaLibraryBrowserController) SetMediaLibraries(value MediaLibrary) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMediaLibraries:"), value)
}


// Returns the shared Media Library Browser instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMediaLibraryBrowserController/shared
func (m_ MediaLibraryBrowserController) SharedMediaLibraryBrowserController() IMediaLibraryBrowserController {
	rv := objc.Send[MediaLibraryBrowserController](m_.ID, objc.Sel("sharedMediaLibraryBrowserController"))
	return rv
}


// A Boolean value that determines whether the Media Library Browser panel is visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmedialibrarybrowsercontroller/isvisible
func (m_ MediaLibraryBrowserController) IsVisible() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isVisible"))
	return rv
}


// A Boolean value that determines whether the Media Library Browser panel is visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmedialibrarybrowsercontroller/isvisible
func (m_ MediaLibraryBrowserController) SetIsVisible(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsVisible:"), value)
}



