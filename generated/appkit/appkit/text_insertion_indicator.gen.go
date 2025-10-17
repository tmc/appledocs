// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TextInsertionIndicator] class.
var TextInsertionIndicatorClass objc.Class

func init() {
	TextInsertionIndicatorClass = objc.GetClass("NSTextInsertionIndicator")
}

type TextInsertionIndicator struct {
	objc.ID
}

func TextInsertionIndicatorFrom(ptr unsafe.Pointer) TextInsertionIndicator {
	return TextInsertionIndicator{
		ID: objc.ID(ptr),
	}
}



