// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [Panel] class.
var panelClass = _PanelClass{objc.GetClass("NSPanel")}

type _PanelClass struct {
	class objc.Class
}

// An interface definition for the [Panel] class.
type IPanel interface {
	IWindow
}

// A special kind of window that typically performs a function that is auxiliary to the main window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPanel

type Panel struct {
	Window
}

// PanelFrom constructs a [Panel] from an unsafe.Pointer.
//
// A special kind of window that typically performs a function that is auxiliary to the main window.
func PanelFrom(ptr unsafe.Pointer) Panel {
	return Panel{
		Window: WindowFrom(ptr),
	}
}



