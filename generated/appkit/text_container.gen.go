// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TextContainer] class.
var TextContainerClass objc.Class

func init() {
	TextContainerClass = objc.GetClass("NSTextContainer")
}

type TextContainer struct {
	objc.ID
}

func TextContainerFrom(ptr unsafe.Pointer) TextContainer {
	return TextContainer{
		ID: objc.ID(ptr),
	}
}



