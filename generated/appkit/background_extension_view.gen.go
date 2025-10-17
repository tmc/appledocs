// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [BackgroundExtensionView] class.
var backgroundExtensionViewClass = _BackgroundExtensionViewClass{objc.GetClass("NSBackgroundExtensionView")}

type _BackgroundExtensionViewClass struct {
	class objc.Class
}

// An interface definition for the [BackgroundExtensionView] class.
type IBackgroundExtensionView interface {
	IView
}

// A view that extends content to fill its own bounds. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBackgroundExtensionView

type BackgroundExtensionView struct {
	View
}

// BackgroundExtensionViewFrom constructs a [BackgroundExtensionView] from an unsafe.Pointer.
//
// A view that extends content to fill its own bounds.
func BackgroundExtensionViewFrom(ptr unsafe.Pointer) BackgroundExtensionView {
	return BackgroundExtensionView{
		View: ViewFrom(ptr),
	}
}



