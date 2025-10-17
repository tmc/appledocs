// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TextContentManager] class.
var TextContentManagerClass objc.Class

func init() {
	TextContentManagerClass = objc.GetClass("NSTextContentManager")
}

type TextContentManager struct {
	objc.ID
}

func TextContentManagerFrom(ptr unsafe.Pointer) TextContentManager {
	return TextContentManager{
		ID: objc.ID(ptr),
	}
}



