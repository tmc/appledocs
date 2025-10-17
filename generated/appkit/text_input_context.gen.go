// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TextInputContext] class.
var TextInputContextClass objc.Class

func init() {
	TextInputContextClass = objc.GetClass("NSTextInputContext")
}

type TextInputContext struct {
	objc.ID
}

func TextInputContextFrom(ptr unsafe.Pointer) TextInputContext {
	return TextInputContext{
		ID: objc.ID(ptr),
	}
}



