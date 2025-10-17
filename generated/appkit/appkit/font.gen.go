// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Font] class.
var FontClass objc.Class

func init() {
	FontClass = objc.GetClass("NSFont")
}

type Font struct {
	objc.ID
}

func FontFrom(ptr unsafe.Pointer) Font {
	return Font{
		ID: objc.ID(ptr),
	}
}



