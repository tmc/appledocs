
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [MediaLibraryBrowserController] class.
var MediaLibraryBrowserControllerClass _MediaLibraryBrowserControllerClass

func init() {
	MediaLibraryBrowserControllerClass = _MediaLibraryBrowserControllerClass{objc.GetClass("NSMediaLibraryBrowserController")}
}

type _MediaLibraryBrowserControllerClass struct {
	objc.Class
}

// An interface definition for the [MediaLibraryBrowserController] class.
type IMediaLibraryBrowserController interface {
	ID() objc.ID
}

type MediaLibraryBrowserController struct {
	id objc.ID
}

func MediaLibraryBrowserControllerFrom(ptr unsafe.Pointer) MediaLibraryBrowserController {
	return MediaLibraryBrowserController{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (m_ MediaLibraryBrowserController) ID() objc.ID {
	return m_.id
}

// Alloc allocates a new instance without initialization.
func (mc _MediaLibraryBrowserControllerClass) Alloc() MediaLibraryBrowserController {
	rv := objc.Send[MediaLibraryBrowserController](objc.ID(mc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (mc _MediaLibraryBrowserControllerClass) New() MediaLibraryBrowserController {
	rv := objc.Send[MediaLibraryBrowserController](objc.ID(mc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewMediaLibraryBrowserController creates and returns a new initialized instance.
func NewMediaLibraryBrowserController() MediaLibraryBrowserController {
	return MediaLibraryBrowserControllerClass.New()
}

// Init initializes the instance.
func (m_ MediaLibraryBrowserController) Init() MediaLibraryBrowserController {
	rv := objc.Send[MediaLibraryBrowserController](m_.ID(), selInit)
	return rv
}
