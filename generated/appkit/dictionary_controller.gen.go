// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [DictionaryController] class.
var dictionaryControllerClass = _DictionaryControllerClass{objc.GetClass("NSDictionaryController")}

type _DictionaryControllerClass struct {
	class objc.Class
}

// A bindings-compatible controller that manages the display and editing of a dictionary of key-value pairs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDictionaryController

type DictionaryController struct {
	ArrayController
}

// DictionaryControllerFrom constructs a [DictionaryController] from an unsafe.Pointer.
//
// A bindings-compatible controller that manages the display and editing of a dictionary of key-value pairs.
func DictionaryControllerFrom(ptr unsafe.Pointer) DictionaryController {
	return DictionaryController{
		ArrayController: ArrayControllerFrom(ptr),
	}
}



