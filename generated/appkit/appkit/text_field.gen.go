// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TextField] class.
var TextFieldClass objc.Class

func init() {
	TextFieldClass = objc.GetClass("NSTextField")
}

type TextField struct {
	objc.ID
}

func TextFieldFrom(ptr unsafe.Pointer) TextField {
	return TextField{
		ID: objc.ID(ptr),
	}
}




