// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TextSelectionNavigation] class.
var TextSelectionNavigationClass objc.Class

func init() {
	TextSelectionNavigationClass = objc.GetClass("NSTextSelectionNavigation")
}

type TextSelectionNavigation struct {
	objc.ID
}

func TextSelectionNavigationFrom(ptr unsafe.Pointer) TextSelectionNavigation {
	return TextSelectionNavigation{
		ID: objc.ID(ptr),
	}
}



