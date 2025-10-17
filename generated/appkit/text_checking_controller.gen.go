// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TextCheckingController] class.
var textCheckingControllerClass = _TextCheckingControllerClass{objc.GetClass("NSTextCheckingController")}

type _TextCheckingControllerClass struct {
	class objc.Class
}

// An interface definition for the [TextCheckingController] class.
type ITextCheckingController interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextCheckingController

type TextCheckingController struct {
	objectivec.Object
}

// TextCheckingControllerFrom constructs a [TextCheckingController] from an unsafe.Pointer.
func TextCheckingControllerFrom(ptr unsafe.Pointer) TextCheckingController {
	return TextCheckingController{objectivec.Object{objc.ID(ptr)}}
}



