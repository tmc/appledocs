// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [DocumentController] class.
var DocumentControllerClass objc.Class

func init() {
	DocumentControllerClass = objc.GetClass("NSDocumentController")
}

type DocumentController struct {
	objc.ID
}

func DocumentControllerFrom(ptr unsafe.Pointer) DocumentController {
	return DocumentController{
		ID: objc.ID(ptr),
	}
}




