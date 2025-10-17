// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [DictionaryController] class.
var DictionaryControllerClass objc.Class

func init() {
	DictionaryControllerClass = objc.GetClass("NSDictionaryController")
}

type DictionaryController struct {
	objc.ID
}

func DictionaryControllerFrom(ptr unsafe.Pointer) DictionaryController {
	return DictionaryController{
		ID: objc.ID(ptr),
	}
}



