// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ScrubberArrangedView] class.
var scrubberArrangedViewClass = _ScrubberArrangedViewClass{objc.GetClass("NSScrubberArrangedView")}

type _ScrubberArrangedViewClass struct {
	class objc.Class
}

// An abstract base class for the views whose layout is managed by a scrubber. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberArrangedView

type ScrubberArrangedView struct {
	View
}

// ScrubberArrangedViewFrom constructs a [ScrubberArrangedView] from an unsafe.Pointer.
//
// An abstract base class for the views whose layout is managed by a scrubber.
func ScrubberArrangedViewFrom(ptr unsafe.Pointer) ScrubberArrangedView {
	return ScrubberArrangedView{
		View: ViewFrom(ptr),
	}
}

// Updates the layout of the arranged view to respect the provided layout attributes. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberArrangedView/apply(_:)
func (s_ ScrubberArrangedView) ApplyLayoutAttributes(layoutAttributes unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("applyLayoutAttributes:"), layoutAttributes)
}


