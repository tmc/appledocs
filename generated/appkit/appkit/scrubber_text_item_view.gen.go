// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ScrubberTextItemView] class.
var ScrubberTextItemViewClass objc.Class

func init() {
	ScrubberTextItemViewClass = objc.GetClass("NSScrubberTextItemView")
}

type ScrubberTextItemView struct {
	objc.ID
}

func ScrubberTextItemViewFrom(ptr unsafe.Pointer) ScrubberTextItemView {
	return ScrubberTextItemView{
		ID: objc.ID(ptr),
	}
}



