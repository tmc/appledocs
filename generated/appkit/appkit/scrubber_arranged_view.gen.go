// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ScrubberArrangedView] class.
var ScrubberArrangedViewClass objc.Class

func init() {
	ScrubberArrangedViewClass = objc.GetClass("NSScrubberArrangedView")
}

type ScrubberArrangedView struct {
	objc.ID
}

func ScrubberArrangedViewFrom(ptr unsafe.Pointer) ScrubberArrangedView {
	return ScrubberArrangedView{
		ID: objc.ID(ptr),
	}
}


// Updates the layout of the arranged view to respect the provided layout attributes. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubberArrangedView/apply(_:)
func (s_ ScrubberArrangedView) ApplyLayoutAttributes(layoutAttributes unsafe.Pointer) {
	sel := objc.RegisterName("applyLayoutAttributes:")
	s_.ID.Send(sel, layoutAttributes)
}

