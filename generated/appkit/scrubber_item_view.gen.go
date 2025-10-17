// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ScrubberItemView] class.
var ScrubberItemViewClass objc.Class

func init() {
	ScrubberItemViewClass = objc.GetClass("NSScrubberItemView")
}

type ScrubberItemView struct {
	objc.ID
}

func ScrubberItemViewFrom(ptr unsafe.Pointer) ScrubberItemView {
	return ScrubberItemView{
		ID: objc.ID(ptr),
	}
}



