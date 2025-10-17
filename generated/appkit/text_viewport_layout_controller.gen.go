// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TextViewportLayoutController] class.
var textViewportLayoutControllerClass = _TextViewportLayoutControllerClass{objc.GetClass("NSTextViewportLayoutController")}

type _TextViewportLayoutControllerClass struct {
	class objc.Class
}

// Manages the layout process inside the viewport interacting with its delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextViewportLayoutController

type TextViewportLayoutController struct {
	objectivec.Object
}

// TextViewportLayoutControllerFrom constructs a [TextViewportLayoutController] from an unsafe.Pointer.
//
// Manages the layout process inside the viewport interacting with its delegate.
func TextViewportLayoutControllerFrom(ptr unsafe.Pointer) TextViewportLayoutController {
	return TextViewportLayoutController{objectivec.Object{objc.ID(ptr)}}
}



