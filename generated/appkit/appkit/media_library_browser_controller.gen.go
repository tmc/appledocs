// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [MediaLibraryBrowserController] class.
var MediaLibraryBrowserControllerClass objc.Class

func init() {
	MediaLibraryBrowserControllerClass = objc.GetClass("NSMediaLibraryBrowserController")
}

type MediaLibraryBrowserController struct {
	objc.ID
}

func MediaLibraryBrowserControllerFrom(ptr unsafe.Pointer) MediaLibraryBrowserController {
	return MediaLibraryBrowserController{
		ID: objc.ID(ptr),
	}
}




