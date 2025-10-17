// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [BackgroundExtensionView] class.
var BackgroundExtensionViewClass objc.Class

func init() {
	BackgroundExtensionViewClass = objc.GetClass("NSBackgroundExtensionView")
}

type BackgroundExtensionView struct {
	objc.ID
}

func BackgroundExtensionViewFrom(ptr unsafe.Pointer) BackgroundExtensionView {
	return BackgroundExtensionView{
		ID: objc.ID(ptr),
	}
}




