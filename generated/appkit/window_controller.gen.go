// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [WindowController] class.
var windowControllerClass = _WindowControllerClass{objc.GetClass("NSWindowController")}

type _WindowControllerClass struct {
	class objc.Class
}

// A controller that manages a window, usually a window stored in a nib file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowController

type WindowController struct {
	Responder
}

// WindowControllerFrom constructs a [WindowController] from an unsafe.Pointer.
//
// A controller that manages a window, usually a window stored in a nib file.
func WindowControllerFrom(ptr unsafe.Pointer) WindowController {
	return WindowController{
		Responder: ResponderFrom(ptr),
	}
}



