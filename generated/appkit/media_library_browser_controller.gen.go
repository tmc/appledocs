// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MediaLibraryBrowserController] class.
var (
	mediaLibraryBrowserControllerClass     _MediaLibraryBrowserControllerClass
	mediaLibraryBrowserControllerClassOnce sync.Once
)

func getMediaLibraryBrowserControllerClass() _MediaLibraryBrowserControllerClass {
	mediaLibraryBrowserControllerClassOnce.Do(func() {
		mediaLibraryBrowserControllerClass = _MediaLibraryBrowserControllerClass{objc.GetClass("NSMediaLibraryBrowserController")}
	})
	return mediaLibraryBrowserControllerClass
}

type _MediaLibraryBrowserControllerClass struct {
	class objc.Class
}

// An interface definition for the [MediaLibraryBrowserController] class.
type IMediaLibraryBrowserController interface {
	objectivec.IObject
}

// An object that configures and displays a Media Library Browser panel. [Full Topic]
//
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




