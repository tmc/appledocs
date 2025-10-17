// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ScrubberImageItemView] class.
var ScrubberImageItemViewClass objc.Class

func init() {
	ScrubberImageItemViewClass = objc.GetClass("NSScrubberImageItemView")
}

type ScrubberImageItemView struct {
	objc.ID
}

func ScrubberImageItemViewFrom(ptr unsafe.Pointer) ScrubberImageItemView {
	return ScrubberImageItemView{
		ID: objc.ID(ptr),
	}
}




