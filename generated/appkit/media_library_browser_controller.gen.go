// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MediaLibraryBrowserController] class.
var mediaLibraryBrowserControllerClass = _MediaLibraryBrowserControllerClass{objc.GetClass("NSMediaLibraryBrowserController")}

type _MediaLibraryBrowserControllerClass struct {
	class objc.Class
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



